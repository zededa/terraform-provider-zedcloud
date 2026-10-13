# Design: UE-177: Provider robustness at high parallelism (async v2 writes, timeouts, drift)

**Tickets:** [UE-177](https://zededa.atlassian.net/browse/UE-177) (provider), driven by
[ENG-3015](https://zededa.atlassian.net/browse/ENG-3015) (seine serialization, API side fixed in
[zedcloud#8516](https://github.com/zededa/zedcloud/pull/8516)). Related: [CI-1004](https://zededa.atlassian.net/browse/CI-1004)
(VW, `-parallelism=160`), [ENG-3016](https://zededa.atlassian.net/browse/ENG-3016) (orphaned instances),
[ENG-2987](https://zededa.atlassian.net/browse/ENG-2987) (VNC 409).
**Repos:** `terraform-provider-zedcloud` @ `4b1ae91e` (branch `nfr-19-edge-app-logo`, last tag `v2.8.0`),
`zedcloud` @ `cdedb70283` (main, = PR #8516 merge, 2026-10-12)
**Date:** 2026-10-12
**Author:** Ivan Curachin
**Status:** Implemented on branch `ue-177-high-parallelism` (pending review and load test)

**Summary:** Most of what UE-177 asks for is provider work that does not depend on the API. Only the
async ("v2") write path depends on ENG-3015. Yuri's comment on ENG-3015 says the provider must move to
the v2 job-based routes, because server-side throughput fixes don't help a client that still gives up
after 30 s. The plan has two releases:

- **2.8.1 (no API dependency):** honour Terraform `ctx` and add `timeouts {}`, stop retrying
  non-idempotent writes, remove resources on 404 during refresh, fix the read-after-create bug, fix the
  perpetual diff on `zedcloud_application_instance`, and report non-JSON errors correctly.
- **2.9.0 (API dependency):** app-instance create, update and delete through `/api/v2/apps/instances`
  plus job polling (`retry.StateChangeConf`). Edge-app delete stays on the sync v1 route. #8516 adds no
  async bundle route; it makes v1 bundle delete parallel server-side (§3.1).

---

## 1. Problem statement

At `-parallelism=160` (VW, CI-1004), Terraform 1.16.4 with provider 2.8.0:

| Symptom | Cause (server, ENG-3015) | Cause (provider, UE-177) |
|---|---|---|
| 159/200 app deletes and ~80% of instance creates/updates fail with `context deadline exceeded`, even though the server finishes the work | Bundle DELETE is serialized (~0.73 s each); instance create/update on one device is slow | Hard 30 s deadline per call with 1 attempt. No `timeouts {}`. Sync API. |
| Every later plan fails with `giving up after 5 attempt(s): 404` (159 resources) until `terraform state rm` | n/a | Read returns an error on 404 instead of removing the resource |
| Every instance plans `~ update` right after a clean apply, which re-PUTs everything and multiplies load | n/a | Server-filled nested fields are not `Computed` |
| 17× `Root object was present, but now absent` on create; instance is orphaned (ENG-3016) | Read-after-create is not yet visible | Read-after-create error dropped (`errs` vs `err`); no not-found retry |
| `invalid character '<' looking for beginning of value` | LB/Cloudflare 504 HTML page | go-swagger decodes every error body as `ZsrvResponse` JSON |

## 2. Current behaviour (code)

All paths are under `v2/`.

### 2.1 Timeout and retry
- The client is go-swagger `ZedcloudAPI` on `httptransport.NewWithClient(..., retryablehttp.StandardClient())`
  (`resources/provider.go:155-178`).
- No resource calls `params.WithContext(ctx)`. go-openapi then applies `DefaultTimeout = 30s`
  (`vendor/github.com/go-openapi/runtime/client/runtime.go:216`, `:447-463`) to the **whole** call,
  including every retryablehttp retry and backoff.
- If a context is set, the runtime uses `context.WithCancel` and applies **no** 30 s timeout. That is the
  hook we use.
- `CheckRetry` (`provider.go:197-224`):
  - When `resp == nil` (transport error or timeout) it calls `DefaultRetryPolicy`, so **POST/PUT/DELETE
    are retried on a client timeout**. With a per-attempt timeout that would duplicate creates. Today the
    shared 30 s ctx hides this.
  - It does not retry a non-GET response.
  - It retries a GET 404 (CI-723 read-after-write) 4 times, ~15 s.
- `retryablehttp` defaults: `RetryMax=4`, `RetryWaitMin=1s`, `RetryWaitMax=30s`.

### 2.2 `zedcloud_application` (`resources/application.go`)
- `POST /v1/apps`, `GET /v1/apps/name/{name}`, `PUT /v1/apps/id/{id}`, `DELETE /v1/apps/id/{id}` (404 → success).
- No `Timeouts`. Read on 404 → `diag.Errorf` (`:157`, `:196`).
- The data source **reuses `ReadApplication`**, so a resource-only 404 change must not leak into the data
  source.

### 2.3 `zedcloud_application_instance` (`resources/application_instance.go`)
- `POST /v1/apps/instances` → `SetId(objectId)` → read by name (`:86`).
- `PUT /v1/apps/instances/id/{id}` with revision lock (`currentApplicationInstanceRevision`, `:256`) → read by name (`:246`).
- `DELETE /v1/apps/instances/id/{id}` (404 → success).
- **Bug at `:86` and `:246`:** `if errs := readApplicationInstanceByName(...); err != nil` tests `err`, which
  is always nil. Read-after-write errors are silently dropped.
- Read by name does `d.SetId(appInstance.ID)`. If the server returns an empty payload, the ID becomes `""`.
  This is the likely source of "Root object was present, but now absent" (**to verify** with
  `TF_HTTP_SESSION_DEBUG`).

### 2.4 Schema fields behind the perpetual diff

| Planned `-> null` | Where | Today |
|---|---|---|
| `bundle_version`, `next_bundle_version` | `manifest_info` block (`schemas/application_instance.go:680`, `schemas/manifest_info.go:111,126`) | Leaf fields are Computed, but the **block** is Optional and not Computed |
| `cpus`, `memory`, `mode`, `disable_v_t_p_m`, `enable_oem_win_license_key` | `vminfo` block (`:782`, `schemas/vm.go:87-139`) | Block Optional; `cpus` Required, `memory` Computed only, booleans Optional only |
| `access` | `logs` block (`:671`, `schemas/app_instance_logs.go:37`) | Block Optional; `access` Required |
| `tags` (`"chm" = "node"`, inherited from device) | top level (`:755`) | Optional TypeMap, not Computed |

If the user omits a block, the server fills it from the bundle or device. On refresh it lands in state and
the plan tries to null it.

### 2.5 Error decoding
Generated readers `Consume(body, *ZsrvResponse)` for every status. If the body isn't JSON, the raw decode
error is returned, `ZsrvResponderToDiags` can't classify it, and the HTTP status is lost (`resources/util.go:71-101`).

### 2.6 Existing patterns to reuse
- `resources/artifact.go` (UE-176): `schema.ResourceTimeout` + context-aware **hand-written** client
  (`client/artifact/client.go`) + poll loop `waitForArtifact`. This is the model for the job client.
- `models.ZsrvResponse` already has `JobID`/`ObjectID`. No job client or model is generated.
  `swagger/zedge_job_service.swagger.json` exists but is not generated.

## 3. Server contract (zedcloud `main`; #8516 still to confirm)

| Op | Route | Sync response | Job |
|---|---|---|---|
| Create instance | `POST /api/v2/apps/instances` (body = v1 `AppInstance`) | **200** `ZsrvResponse`; **job id in `objectId`** (not `jobId`) | name `<inst>_create_<jobId>` |
| Update instance | `PUT /api/v2/apps/instances/id/{id}[?makeSnapshot=]` (body `id` must equal path id, else 400) | 200, job id in `objectId` | `app_instance_update_<id>_<5>` |
| Delete instance | `DELETE /api/v2/apps/instances/id/{id}` | 200, job id in `objectId`; **404 synchronously** if absent | `<id>_delete_<jobId>` |
| Poll | `GET /api/v1/jobs/id/{id}/objectType/OBJECT_TYPE_EDGE_APP_INSTANCE` | `JobConfig` | — |

Sources: `srvs/gilas/appinstance_handlers.go:86-133`, `srvs/seine/appinstproc.go:3191,3780,5146`,
`srvs/gilas/job_handlers.go:21-57`, `libs/zmsg/bulk/bulk.proto:59-100`, apitest `tests/apiTest/tests/test_app_instance.py:2858-2938`.

- `JobConfig.status`: `INITIALIZED`, `READY`, `INPROGRESS` → `COMPLETED` | `FAILED` | `ABORTED`.
- Error text is in `errorDetailOfJob` (plain text on read; also set on success). Per-object result is in
  `individualOpStatus` (name → `SUCCESS` or error).
- The job has **no instance id**. After create, read by name (unique per enterprise).
- **Use the `/objectType/OBJECT_TYPE_EDGE_APP_INSTANCE` form.** Without it, RBAC checks device access
  (`getJobAccessType`, `job_handlers.go:167-181`).
- **Duplicates surface late:** v2 create returns 200 first and fails the **job** on a name or id conflict.
  REST callers get no idempotency key.
- No job TTL was found, and jobs are never cleaned up. That isn't a provider concern, but worth raising.
- Edge apps (bundles): **no `/api/v2/apps` routes on main**, including after #8516.

### 3.1 What zedcloud#8516 changed (`cdedb70283`, merged 2026-10-12)

The change is server-side only. All 13 files are under `srvs/seine/`. No gilas route, proto, or swagger
changes, so **the API contract the provider sees is unchanged.** Four commits:

1. **VNC slot claimed atomically, outside the create tx** (`devproc.go`, `devdata.go`, `appinstproc.go`).
   - Before: each instance create rewrote the whole device row with a version check inside the create
     transaction. That held the device row lock until commit, so creates on one device ran one at a
     time, and the losers got 409 "Version mismatch" (ENG-2987).
   - Now: the slot is set or cleared with single-column SQL on `vncportusage` and pre-claimed before the
     tx. It is released if the tx fails, including when ctx expires.
   - `VNCPortUsage` is marked `update_invalid`, so full device updates can't write a stale mask back.
   - Cluster placement and policy-driven creates still claim inside their tx.
2. **Bundle delete: batched policy reads in `bundleAttachedToProject`** (`appproc.go`, `policydata.go`).
   The N+1 read of every project plus one policy per attached policy is replaced by one
   `PolicyDB.ReadByIDs` query. Before, this took ~0.73 s per delete.
3. **Kafka app-bundle writes on a bounded keyed pool** (`keyedpool.go`, `appproc.go`).
   - Before, seine ran app CREATE/UPDATE/DELETE inline on its single Kafka goroutine, so every other
     message, reads included, queued behind them.
   - Now there are 8 workers keyed by tenant + app id (or name). Writes to the same app stay in order.
   - Each worker has a 32-entry queue, and Kafka backpressure applies when it is full. The worker count
     stays below the 15-connection PG pool.
4. **Review fixes:** the VNC pre-claim device is now scope-checked through a project-scoped read, and the
   compensating release runs on `bgReconcileCtx` so it survives an expired deadline.

**What this means for the provider:**
- Bundle delete throughput goes up by roughly 8× from parallelism, plus a faster per-delete path. At 160
  concurrent deletes, the queue (8×32 = 256) absorbs the burst. Tail latency should drop well under the
  measured 117 s, but **it is not re-measured yet**.
- Same-device instance creates no longer serialize on the device row.
- Neither change touches the request/response shape. The provider issues in UE-177 (fixed 30 s ctx, 404
  on refresh, drift, read-after-create, error decoding) all remain.
- Yuri's v2 job-polling ask is still the way to make instance writes immune to tail latency.
  Bundle writes have no async route, so for `zedcloud_application` the provider relies on §4.1 timeouts.

## 4. Design

### 4.1 Timeouts and context (2.8.1, all resources via a helper; both app resources first)
1. Add `Timeouts: &schema.ResourceTimeout{Create: 20m, Update: 20m, Delete: 20m, Read: 5m, Default: 20m}`
   to `zedcloud_application` and `zedcloud_application_instance`. SDK v2 already puts the deadline from
   `d.Timeout(...)` on the `ctx` passed to `*Context` funcs.
2. Pass `ctx` on every call: `params.WithContext(ctx)`. This removes the hidden 30 s deadline, which now
   belongs to Terraform.
3. **Per-attempt timeout:** `retryClient.HTTPClient.Timeout = request_timeout`. Add a new provider attribute
   `request_timeout` (default `"60s"`, env `ZEDCLOUD_REQUEST_TIMEOUT`). This stops one hung request from
   eating the whole resource timeout.
4. **Retry policy for writes.** In `CheckRetry`, for `resp == nil` on a non-GET, retry only
   connection-establishment errors (dial / connection refused / TLS handshake). Never retry a
   timeout or reset after the request was sent. Without this, change 3 creates duplicate POSTs.
5. Add provider attributes `max_retries` (default 4) and `retry_wait_max` (default 30s). This is optional,
   but cheap, and lets VW tune behaviour without a new release.
6. Introduce helper `withCtx(ctx, params)` / a small wrapper, so resources migrate one at a time. Out of
   scope for this ticket: converting the ~40 other resources. Do that mechanically in a follow-up.

**As implemented (4.1):** the ~40 resources that don't pass `ctx` yet are no longer stuck with the hidden
30 s: `ProviderConfigure` sets go-openapi's `DefaultTimeout` to
`(max_retries+1) × (request_timeout + retry_wait_max)`, which is 450 s with the defaults.
`zedcloud_application` and `zedcloud_application_instance` pass `ctx` and get Terraform's timeouts.
`checkRetry` reads the method from the `*url.Error` op, because retryablehttp passes no request when
there is no response.

### 4.2 404 on refresh → remove from state (2.8.1)
- Split each resource's Read into:
  - `resourceXRead` (Terraform refresh): on `isStatusNotFound(err)`, log
    `[WARN] ... not found, removing from state`, `d.SetId("")`, return nil.
  - `readXAfterWrite`: used by Create and Update; never removes the resource (see 4.3).
- The data sources keep the current behaviour (error on 404), so `DataSource.ReadContext` must stop
  pointing at the resource Read.
- Keep the CI-723 GET-404 transport retry. It costs ~15 s per vanished resource on refresh, which is
  acceptable. Reduce it to 2 attempts **only** on refresh, via a ctx value, if this proves slow at 160.

### 4.3 Read-after-write (2.8.1)
- Fix the `errs` vs `err` bug (`application_instance.go:86,246`). Check `application.go` too.
- `readXAfterWrite` wraps the GET-by-name in `retry.RetryContext(ctx, 2*time.Minute, ...)`:
  - **Retryable:** not-found, or a 200 with an empty `ID`.
  - **Non-retryable:** everything else.
- Never `SetId("")` from this path. Create has already set the id from `objectId`, so on final failure
  the resource stays in state (tainted) instead of being orphaned. This fixes the provider half of ENG-3016.

### 4.4 Perpetual diff on `zedcloud_application_instance` (2.8.1)
- Make `manifest_info`, `vminfo` and `logs` **Optional + Computed** blocks.
- **As implemented:** only the blocks become O+C. The nested schemas (`VMSchema`, `AppInstanceLogsSchema`,
  `ManifestInfoSchema`) are shared with `zedcloud_app_profile` and the app-inst config schemas, so their
  fields are left unchanged. The reported diff is the whole block planned `-> null` because the config
  omits it, and block-level O+C fixes that. Revisit the inner fields only if a diff shows up with a block
  that *is* configured.
- `tags`: make it **Optional + Computed**. In seine, device tags are *copied once at create* and stored on
  the instance, and are never merged on read (Q3). Facts:
  - In `appInstCreateCommon` (`srvs/seine/appinstproc.go:2933-2936`), when the request has no tags,
    `appInst.Tags = dev.Tags` copies the whole device map. If the user sends any tags, no device tags
    are added.
  - The copy is persisted in `app_inst.tags` (`AppInstData.Tags`, `appinstdata.go:542`). Reads return
    the stored value (`appinstproc.go:253,397`).
  - Update (`appInstUpdate` → `bindAppInstData`, `:224`) takes tags from the request only, and the row is
    written in full (`appinstdata.go:419-427`; `tags` is not `update_invalid`). **A PUT without tags
    clears them.** Today's provider does exactly that on the "drift" apply, silently erasing the
    inherited tags.
  - Later device tag changes do not propagate to existing instances. Deployment-driven instances
    (`autoappdeploy.go:489`) copy them the same way.
  - Consequence: O+C state holds the stored value, and the next PUT sends it back unchanged. Trade-off:
    with SDK v2 O+C, removing `tags` from config (or `tags = {}`) can't clear them. Document this. If
    clearing is needed later, add an explicit flag.
- Trade-off: with O+C blocks, deleting a block from config no longer resets it server-side. Document this
  in `docs/resources/application_instance.md`.
- **Regression guard:** add an acceptance fixture with a *minimal* instance (no `vminfo`, `logs`,
  `manifest_info` or `tags`) on a tagged device. `resource.Test` fails on a non-empty post-apply plan.
- Also check `zedcloud_application` for the `ac_kind` `PodManifest` vs `VMManifest` diff. The fix is in the
  zededa-terraform skill example, which lives in another repo; tracked here only as a note.

### 4.5 Non-JSON error bodies (2.8.1)
In `HttpTransportWrapper.RoundTrip`, for `status >= 400` where the `Content-Type` isn't JSON, or the body
isn't valid JSON, rewrite the response before it reaches go-swagger:
- `Content-Type: application/json`
- body
  `{"httpStatusCode":<code>,"httpStatusMsg":"<status>","error":[{"details":"<first 512 bytes, stripped of HTML>"}]}`

Typed `ZsrvResponder` handling then works unchanged, and users see `HTTP 504 Gateway Time-out: …`.
retryablehttp still sees the original status code, because rewriting happens after `CheckRetry`. The
wrapper sits *outside* the retry client (verified).

**As implemented, also:** `retryClient.ErrorHandler = giveUpWithLastResponse`. When retries are exhausted
on an HTTP status (GET 5xx or 404), the last response is passed through instead of the bare
`giving up after N attempt(s)` error. Without this, a GET that keeps getting a 504 page never reaches
the rewrite. `isStatusNotFound` still works, because typed 404 responses implement `Code()`.

### 4.6 Async v2 writes for app instances (2.9.0)
New hand-written, context-aware client `v2/client/jobs` and v2 instance methods (follow `client/artifact`).
They reuse the same `retryablehttp` client and auth. Hand-written, not generated, because the response
puts the job id in `objectId` and the job service isn't in the generated client.

```
POST   /v2/apps/instances               -> ZsrvResponse{objectId=jobId}
PUT    /v2/apps/instances/id/{id}       -> ZsrvResponse{objectId=jobId}
DELETE /v2/apps/instances/id/{id}       -> ZsrvResponse{objectId=jobId} | 404
GET    /v1/jobs/id/{jobId}/objectType/OBJECT_TYPE_EDGE_APP_INSTANCE -> JobConfig
```

Shared waiter `waitForJob(ctx, c, jobID, timeout) (*JobConfig, error)` using `retry.StateChangeConf`:
- **States:**
  - `Pending: INITIALIZED, READY, INPROGRESS, ""`
  - `Target: COMPLETED`
  - `FAILED` / `ABORTED` → error carrying `errorDetailOfJob` plus non-`SUCCESS` entries of
    `individualOpStatus`
- **Polling:** `Delay 1s`, `MinTimeout 2s`, `PollInterval` growing to 10s with jitter. Worst case at 160
  parallel is ~16-80 GET/s, which is fine for the jobs read path. Confirm with ENG (**Q4**).
- **Job GET errors:**
  - transient (5xx / timeout): keep polling
  - 404 on the job: retry up to `NotFoundChecks: 5`
  - 403: fail fast with an RBAC hint

Resource flows:
- **Create** (the provider mints the id; see Q5):
  1. Generate a UUID and send it as `id` in the POST body. Seine keeps a caller-supplied id
     (`appInstCreateCommon`, `appinstproc.go:2743`, generates one only when it's empty) and validates it
     (`:6331`). app-profiles already does this for its v2 creates (`deployment.go`, `zutils.GenerateId()`).
  2. `d.SetId(uuid)` **before** waiting, so the instance is never orphaned. If the wait times out or
     the job fails, return the error. The SDK keeps the resource in state as tainted, and the next apply
     refreshes it by id and replaces it or adopts the result.
  3. `waitForJob(d.Timeout(Create))`, then `readAfterWrite` (by name, cross-checked against the minted id).
  - If the job fails because the name already exists and `GET by name` returns a **different** id, it's
    a real collision. Raise an error that asks for `terraform import`. Never adopt by name.
- **Update:**
  1. Revision lock as today.
  2. PUT v2.
  3. Wait (Update timeout).
  4. Read after write.
- **Delete:**
  1. DELETE v2. 404 → done.
  2. Wait (Delete timeout).
  3. Confirm with GET by id → 404 (short retry).
- **projectId (found on the local controller):** v2 create rejects a request without `projectId` up
  front with 400 "project id is empty". v1 derives the project from the device. When `project_id` isn't
  configured, the provider reads it from the edge node (or the edge node cluster) before the v2 POST.
- **Id format:** the minted id is an RFC 4122 v4 UUID, like the ones the controller generates.
- **Response parsing:** take the job id from `ZsrvResponse.objectId`. Also accept `jobID` when the body
  is an `AppInstanceServiceResp`. gilas returns that shape whenever seine fills `jobID`
  (`appinstance_handlers.go:345-347`). Today seine doesn't fill it on these paths, but it costs nothing to
  handle.
- **Mode switch:** provider attribute `async_writes` (bool, default `true` in 2.9.0). With `false`, the v1
  sync path is used (plus 4.1 timeouts), for controllers older than the v2 routes. If the v2 route itself
  is missing (404/405 with no ZsrvResponse body on POST), fall back to v1 once per provider instance and
  log a warning. Don't auto-detect on DELETE, because a 404 there is ambiguous.

### 4.7 Edge-app (`zedcloud_application`) delete
- #8516 adds no async route (§3.1), so this resource stays on sync v1 and 4.1 is the fix.
- Per-request timeout: the default `request_timeout` of 60 s should cover post-#8516 latency. Confirm it
  with the load test. Before #8516, p90 was 106 s at N=160.
- Because a timed-out DELETE is no longer retried (4.1.4), a delete that completed server-side after the
  client gave up is handled on the next run: Delete gets 404, which counts as success, and refresh drops
  the resource (4.2).
- Nothing on main plans `/api/v2/apps` async routes. The gRPC migration plan
  (`docs/design/kafka-sync-to-grpc/implementation-plan-phase-g7.md`) moves app-bundle writes to sync gRPC,
  not to jobs. If that changes, reuse `waitForJob` (4.6).

## 5. Work breakdown

| # | Item | Release | Size |
|---|---|---|---|
| 1 | ctx propagation, `Timeouts`, `request_timeout`/`max_retries`, safe `CheckRetry` (4.1) | 2.8.1 | M |
| 2 | 404 → remove on refresh; split resource/data source Read (4.2) | 2.8.1 | S |
| 3 | `errs` bug + read-after-write retry (4.3) | 2.8.1 | S |
| 4 | O+C schema for `manifest_info`/`vminfo`/`logs`/`tags` (4.4) | 2.8.1 | M |
| 5 | Non-JSON error rewrite (4.5) | 2.8.1 | S |
| 6 | `client/jobs` + v2 instance client + `waitForJob` (4.6) | 2.9.0 | M |
| 7 | Instance create, update and delete on v2 + `async_writes` + fallback (4.6) | 2.9.0 | M |
| 8 | Edge-app delete: verify v1 + timeouts at N=160 (4.7); no async route exists | 2.8.1 | XS |
| 9 | Docs (`timeouts`, new provider attrs, minimum cloud version), `make docs` | both | S |

## 6. Testing
- **Unit tests** (`httptest.Server`, no cloud):
  - `waitForJob`: completed, failed (error text surfaced), aborted, ctx deadline, 404 then found, 5xx
    then completed.
  - `CheckRetry`: no POST/PUT retry on timeout; retry on dial error; GET 404 still retried.
  - 504 HTML → `HTTP 504 Gateway Time-out` diag.
  - Read 404 → `Id()==""` for the resource, error for the data source.
- **Acceptance** (`TF_ACC`):
  - minimal-instance fixture, empty plan after apply (4.4)
  - `CheckDestroy` after a delete out of band (4.2)
  - existing `TestApplication*` / `TestApplicationInstance*` on both `async_writes=true` and `false`
- **Fault injection:** extend `test/retry/simulate_404.py` (mitmproxy) to delay responses past
  `request_timeout`, return Cloudflare 504 HTML, and 404 the first N GETs after create.
- **Load (sign-off):** re-run the CI-1004 config `ci1004-tf160/{container,vm}_main.tf` on hummingbird at
  `-parallelism=160`. Pass criteria:
  - 200 app create and destroy, and 120 instance create, update, recreate and destroy, with 0 failures
  - a second `plan` shows no changes
  - no orphans (`GET /v1/apps/instances` count matches state)

### 6.1 Local results (2026-10-12, minikube controller, `-parallelism=160`)

Harness: 2 registered-only edge nodes, N edge apps (VM), N instances (`activate=false`), with
mitmproxy in front of the controller for fault injection. Registry 2.8.0 vs this branch.

| Scenario | 2.8.0 | branch |
|---|---|---|
| 200 apps / 120 instances: create, re-apply, delete | ok | ok |
| `plan` right after apply | **120 to change** (perpetual diff) | no changes |
| 35 s response latency on writes (server completes) | 40/40 creates time out, 40 orphans, re-apply 409 ×40 | creates take ~38 s, no errors |
| Deletes under 35 s latency | 10 time out; every later plan, apply and destroy fails with 404 | ok |
| 30% Cloudflare-style HTML 504 on reads (48 injected) | – | clean plan (retried) |
| 10 instances deleted out of band | – | plan: 10 to add, converges |
| Job reads delayed past a 10 s `timeouts.create` | – | 20 job-timeout errors → tainted → next apply replaces 20, no orphans |

Found and fixed during these runs:
- The local `znginx` `istio-proxy` sidecar (128 MiB) is OOM-killed by bursts from both versions;
  raised to 512 MiB locally for the comparison. Async mode made the bursts worse, so the first job
  poll is now delayed by a jittered interval, and the device→project lookup for v2 creates is cached
  per provider run, with concurrent lookups for the same device sharing one read.
- `HttpTransportWrapper.RoundTrip` wrote the package-level `version` on every request: a data race
  under parallelism (`go test -race`). The User-Agent is now built once.

Not covered locally: real controller latency (the local controller answers in milliseconds), which
still needs hummingbird.

## 7. Open questions
1. ~~What exactly does #8516 add?~~ **Answered (§3.1):** seine-only throughput fixes (atomic VNC claim, batched
   policy reads, 8-worker keyed pool for bundle writes). No new routes, and the v2 instance routes are
   unchanged. Async bundle routes are not planned on main: the G7 gRPC plan moves bundle writes to
   **sync** gRPC (§4.7). Sync v1 plus timeouts is the answer for `zedcloud_application`.
2. ~~Minimum cloud version for the v2 routes~~ **Answered:** `/api/v2/apps/instances` was added in
   `ff938c4c60` (#4726, 2022-06) and first shipped in tag **8.5.0**, so every supported controller has it.
   `async_writes` can default to `true`, and the v1 fallback is only a safety valve. Separately, #8516 is
   not yet in any release tag (latest are 21.1.2 / 20.0.20). The VW load re-test (§6) needs the next
   release that includes it, or hummingbird running main.
3. ~~Is `AppInstance.tags` merged with device tags?~~ **Answered:** copied at create, only when the request
   has none (`appinstproc.go:2933`); stored; not merged on read; cleared by a PUT without tags. So the
   field is Optional + Computed (§4.4).
4. **Jobs read load and retention.** Mostly answered.
   - **Per-read cost:** one primary-key lookup on `job_data` by `(id, enterpriseid)`
     (`jobproc.go:414-425`). It's cheap.
   - **Routing depends on a cluster flag.** With `SEINE_GRPC_MIGRATION_ENABLED` set, gilas serves job
     reads over gRPC (`job_handlers.go:141-149`). Without it, they go over Kafka and are handled **inline
     on seine's single consumer goroutine** (`seine.go:648-650`). That's the same goroutine #8516 had to
     offload bundle writes from.
   - **Provider consequence:** poll conservatively. Start at 2 s, back off to 10 s, add jitter. That's
     ~16-30 req/s at parallelism 160 rather than 80.
   - **Retention:** no TTL, cleanup job or retention policy exists. The migration `20260508120000_addjobdataindexes`
     only adds indexes. Every Terraform create, update and delete leaves a `job_data` row forever. Not a
     provider blocker; raise it with ENG.
   - **About the flag:**
     - Introduced in `13b198fa11` (ZEDCLOUD-3301, 2026-07-31). Job reads over gRPC came the day before in
       `ce8fad6a21` (Phase 6e).
     - First shipped in **21.0.0**, so controllers before 21.0 always serve job reads over Kafka.
     - It is checked by presence (`os.Getenv(...) != ""`) and is **off by default**
       (`docs/design/kafka-sync-to-grpc/implementation-plan-phase6.md` §1.1).
   - **Decision (2026-10-12):** the flag is **not set** in hummingbird's Vault (checked by Ivan), so the
     design assumes it is off everywhere. Job polls run on seine's single Kafka goroutine.
     - Keep polling conservative: start at 2 s, back off to 10 s, add jitter.
     - Expose a hidden provider setting `job_poll_interval_max` (default 10 s) so the interval can be
       tuned without a release.
     - The hummingbird load test (§6) therefore covers the worst case. Watch seine Kafka consumer lag and
       the latency of unrelated reads during the test. If unrelated reads degrade, raise the max interval.
5. ~~Adoption after create timeout / client-supplied id~~ **Answered: no server change needed.**
   - v2 create keeps a caller-supplied `id` (`appinstproc.go:2743`) and validates it (`:6331`).
     app-profiles already relies on this.
   - The provider mints the UUID and sets it in state before waiting, so there are no orphans and no
     adoption by name (§4.6).
   - Caveat: true server-side dedupe of a *retried* POST (`idempotentWriteEnabled`, `:3201`) is gated on
     the gRPC `x-zedcloud-idempotency-key`, which REST can't send. A duplicate POST with the same id
     fails the job ("already exists"), and the provider then sees that the existing instance has its own
     id. So the provider must still not retry POSTs on timeout (§4.1.4).
6. ~~Move the job id to `jobId`?~~ **Answered: keep `objectId`.** The apitests (`test_app_instance.py:2858-2938`) and
   app-profiles (`zedcloud_steps.go`) already read the job id from `objectId`, so changing it would break
   them. The provider reads `objectId`, and falls back to `jobID` if gilas ever returns
   `AppInstanceServiceResp` (§4.6).

**No open questions remain.** `SEINE_GRPC_MIGRATION_ENABLED` is treated as off (Q4). Nice to have: ENG
ticket for `job_data` retention, since every Terraform write now leaves a job row.
