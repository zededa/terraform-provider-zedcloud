# Design: PLE-576 — Immediate app instance snapshots in Terraform

**Ticket:** [PLE-576](https://zededa.atlassian.net/browse/PLE-576) — *terraform: implement snapshot operations*
(epic [PLE-518](https://zededa.atlassian.net/browse/PLE-518) — *Application snapshots v2.0*)
**Reference implementation:** zcli [#1105](https://github.com/zededa/zcli/pull/1105) — *feat: Implement immediate snapshots (ple-556)*, commit `765db61`
**Repos:** `zededa/terraform-provider-zedcloud` @ `c9a03747` (origin/main), `zededa/zedcloud` (main), `zededa/zcli` @ `79b1feb`
**Branch:** `ple-576-app-instance-snapshots`
**Date:** 2026-10-15
**Author:** Ivan Curachin
**Status:** Draft. Decided 2026-10-15: activation and the capability pre-check are in scope; edge-node-cluster app instances are not supported; a rollback shows up as ordinary drift; `allow_snapshots` / `max_snapshots` are in scope; new resources use the `zedcloud_app_instance_` prefix (§8.1). Remaining open questions in §8.2.

**Headline:** zcli gained four commands (`snapshot-create`, `snapshot-activate`, `snapshot-show`,
`snapshot-delete`) on `edge-app-instance`. The provider has none of this today. The proposal is
one new resource for the snapshot itself, one action-style resource for activation (rollback),
one data source for listing, and two new attributes on `zedcloud_application_instance`
(`allow_snapshots`, `max_snapshots`) for automatic app-update snapshots. Two of the five endpoints we need are not in the published
swagger, so the provider needs a small hand-written client, like the one for `zedcloud_artifact`.

---

## 1. Problem Statement

PLE-576 has no description. The ticket comments say it is about *"the support of the immediate
snapshots in the terraform provider"*. The scope is therefore: give Terraform users the same
operations zcli#1105 gave CLI users.

An *immediate* snapshot is a point-in-time copy of an app instance's config and volumes, taken
on demand. *Activating* it rolls the app instance back to that copy. This is different from the
older *app-update* snapshots, which EVE takes automatically when an app instance is updated
(controlled by `allowSnapshots` / `maxSnapshots` on the app instance). Those two settings aren't in zcli #1105
but are in scope for PLE-576 (D5, §4.7), so Terraform users can turn on both kinds of snapshot.

### 1.1 What zcli#1105 added

| zcli command | HTTP call | Notes |
|---|---|---|
| `snapshot-create --app-instance=<n> [-f]` | `POST /v1/apps/instances/id/{id}/snapshot` body `{}` | Without `-f`, zcli first reads `GET /v1/devices/id/{deviceId}/status` and aborts if `rawStatus` doesn't contain `API_CAPABILITY_VOLUME_SNAPSHOTS_IMMEDIATE` |
| `snapshot-activate <snap> --app-instance=<n>` | `PUT /v1/apps/instances/id/{id}/snapshot/name/{snap}/activate` | zcli never sends `?force=` |
| `snapshot-show [<snap>] --app-instance=<n>` | `GET .../id/{id}/snapshots` or `GET .../id/{id}/snapshot/name/{snap}` | zcli strips `config.encryptedSecrets` before printing |
| `snapshot-delete <snap> --app-instance=<n>` | `DELETE /v1/apps/instances/id/{id}/snapshot/name/{snap}` | |

Source: `zcli/modules/edge-app-instance.py` (`snapshot_create` … `snapshot_delete`), templates
`zcli/resources/edge-app-instance-snapshot*.mustache`.

### 1.2 Verified controller behaviour

| Thing | Actually | Where |
|---|---|---|
| Create route | gilas-only route, **not in swagger** | `zedcloud/srvs/gilas/appinstance_handlers.go:32` |
| Create always immediate | gilas hard-codes `Immediate: true` | `srvs/gilas/appinstance_snapshot_handlers.go:126-157` |
| Snapshot name | **Server-generated**: `snapshot-<YYYYMMDDhhmmss>-<8 hex>`. The caller can't choose it | same, `:144-145` |
| Create response | `ZsrvResponse` with `objectId` = snapshot UUID and `objectName` = snapshot name (ENG-1676) | `srvs/seine/aisnapproc.go:105-110, 296` |
| One snapshot per app instance | Create **deletes every existing snapshot** of the app instance first | `seine/aisnapproc.go:217-226` (`deleteSnapshots`, `:428`) |
| Create pre-checks (server) | App instance exists; device isn't KubeVirt; app instance isn't in `RUN_STATE_ERROR`. **No** EVE capability check | `seine/aisnapproc.go:139-215` |
| Activate route | gilas-only, **not in swagger**; optional `?force=true\|false` | `gilas/appinstance_handlers.go:71`, `appinstance_snapshot_handlers.go:184-215` |
| Activate semantics | Sets `activate=true`, applies the snapshot config to the app instance (`applySnapshot`). If already `activate` and no `force` → no-op success | `seine/aisnapproc.go:648-710` |
| Get / list / delete | In swagger (`EdgeApplicationInstanceStatus_*AppInstanceSnapshot*`) | `libs/zmsg/zapiservices/zedge_app_service.proto:1498-1540` |
| Get by snapshot UUID | `GET /v1/apps/instances/snapshot/id/{id}` (in swagger) | `gilas/appinstance_handlers.go:56` |
| Get / list payload | Only `id`, `name`, `config` are filled. **`state`, `activate`, `immediate` are always empty** even though the message has them | `seine/aisnapproc.go:328-426` |
| State | Separate endpoint, served by **ganges**: `GET /v1/apps/instances/snapshot/id/{id}/state` → `{state, timestamp}` | `ganges/api/service/useragent/service.go:52` |
| State machine | `CREATING → CREATED \| CREATE_FAILED`, `ROLLING_BACK → ACTIVE \| ROLLBACK_FAILED`, `DELETING → DELETED \| DELETE_FAILED`. Transitions are driven by **device reports** (`ZInfoSnapshot`), so they need an online device | `ganges/aisnapproc.go:88-130`, enum `libs/zmsg/app/appinst.proto:207-219` |
| Permission | `ACCESS_TYPE_APP_INSTANCE` for all operations (same as managing app instances) | `gilas/appinstance_snapshot_handlers.go` |

### 1.3 What the provider has today

- No snapshot support in `v2/` (models, schemas, resources, client).
- The generated `v2/client/application_instance` client is built from `swagger/zedge_app_service.swagger.json`, which
  contains only get/list/delete for snapshots.
- The swagger's `AppInstance` definition **has** `allowSnapshots`, `maxSnapshots` and `availableSnapshots`, but
  `v2/models/app_instance.go` (generated, then hand-maintained) doesn't, and neither does `v2/schemas/application_instance.go`.
  Side effect today: the provider's app instance update is a full `PUT` without these fields, and seine copies
  them straight from the request (`appinstproc.go:228-229`), so **any Terraform update resets `allowSnapshots` to false**
  if someone turned it on in the UI. Adding the fields (§4.7) fixes that.
- No device-status call in `v2/`: the swagger has `EdgeNodeStatus_GetEdgeNodeStatus` (`GET /v1/devices/id/{id}/status`), and `v2/models/device_status_msg.go` already has `RawStatus`, but `v2/client/node` doesn't expose the operation. Needed for the capability pre-check (§4.4).
- Precedent for a hand-written client: `v2/client/artifact/client.go` (NFR-19), wired into `ZedcloudAPI` in `v2/client/zedcloud_api.go`.
- Precedent for an action-style resource: `zedcloud_patch_reference_update` (`v2/resources/patch_reference_update.go`).

---

## 2. Proposed Solution

```hcl
# Turn on automatic app-update snapshots on the app instance (§4.7).
resource "zedcloud_application_instance" "app" {
  # ... existing attributes ...
  allow_snapshots = true
  max_snapshots   = 1
}

# Take an immediate snapshot. Re-taken whenever a trigger value changes.
resource "zedcloud_app_instance_snapshot" "before_upgrade" {
  app_instance_id = zedcloud_application_instance.app.id

  triggers = {
    app_version = zedcloud_application.app.manifest[0].app_version
  }

  timeouts { create = "20m" }
}

# Roll the app instance back to that snapshot. One-shot: runs on create
# (and again when a trigger changes). Destroy does not undo the rollback.
resource "zedcloud_app_instance_snapshot_activation" "rollback" {
  app_instance_id = zedcloud_application_instance.app.id
  snapshot_name   = zedcloud_app_instance_snapshot.before_upgrade.name
  force           = false

  triggers = { ticket = "INC-1234" }
}

# Read-only listing (covers snapshot-show without a name).
data "zedcloud_app_instance_snapshots" "app" {
  app_instance_id = zedcloud_application_instance.app.id
}
```

| zcli | Terraform |
|---|---|
| `snapshot-create` | `zedcloud_app_instance_snapshot` (create) |
| `snapshot-show <snap>` | `zedcloud_app_instance_snapshot` (read / import) |
| `snapshot-show` (list) | `data.zedcloud_app_instance_snapshots` |
| `snapshot-delete` | `zedcloud_app_instance_snapshot` (destroy) |
| `snapshot-activate` | `zedcloud_app_instance_snapshot_activation` (create) |
| `-f` on create | `skip_capability_check = true` |
| *(UI only today)* app instance snapshot settings | `zedcloud_application_instance.allow_snapshots` / `max_snapshots` |

---

## 3. Alternatives considered

| Option | Why not (or why) |
|---|---|
| **A. `activate = true` attribute on the snapshot resource** | Flipping a bool would silently roll back production. Can't activate snapshots Terraform didn't create (e.g. app-update snapshots). Un-setting it can't be undone. Rejected. |
| **B. Snapshot block / `take_snapshot` flag inside `zedcloud_application_instance`** | Mixes a one-shot operation into a long-lived resource; every app-instance read would also need snapshot reads; ties snapshot lifetime to the app instance. Rejected. |
| **C. Terraform "actions" (TF ≥ 1.14)** | The natural fit for activate, but actions exist only in `terraform-plugin-framework`. This provider is SDKv2 (`terraform-plugin-sdk/v2 v2.24.1`) with no mux. Revisit if/when we add a framework server. |
| **D. Generate client from patched swagger** | Create and activate are gilas-only routes and the state route lives in ganges; a hand-patched swagger is lost on the next `make swagger`. Better: hand-written client now + zedcloud ticket to publish the routes (§6). |
| **E. Separate snapshot resource + activation resource + list data source** | **Chosen.** Matches the API shape, keeps rollback explicit, follows existing precedents. |

---

## 4. Detailed Design

### 4.1 Client: `v2/client/app_instance_snapshot` (hand-written)

Same style as `v2/client/artifact`: uses the shared `runtime.ClientTransport` so auth, retries
and `X-Request-Id` behave like the generated clients. Package doc lists the controller quirks from §1.2.

```go
type Snapshot struct {
    ID        string          `json:"id"`
    Name      string          `json:"name"`
    State     string          `json:"state,omitempty"`     // empty from get/list today
    Activate  bool            `json:"activate,omitempty"`  // idem
    Immediate bool            `json:"immediate,omitempty"` // idem
    Config    json.RawMessage `json:"config,omitempty"`
}
type SnapshotState struct { State string `json:"state"`; Timestamp string `json:"timestamp"` }

type ClientService interface {
    Create(ctx, appInstID string) (*models.ZsrvResponse, error)                 // POST .../id/{id}/snapshot
    Get(ctx, snapID string) (*Snapshot, error)                                   // GET .../snapshot/id/{id}
    GetByName(ctx, appInstID, name string) (*Snapshot, error)                    // GET .../id/{id}/snapshot/name/{n}
    List(ctx, appInstID string) ([]Snapshot, error)                              // GET .../id/{id}/snapshots
    State(ctx, snapID string) (*SnapshotState, error)                            // GET .../snapshot/id/{id}/state
    Activate(ctx, appInstID, name string, force bool) (*models.ZsrvResponse, error) // PUT .../activate?force=
    Delete(ctx, appInstID, name string) (*models.ZsrvResponse, error)            // DELETE .../snapshot/name/{n}
}
var ErrNotFound = errors.New("snapshot not found")
```

- Wire as `ZedcloudAPI.AppInstanceSnapshot` next to `Artifact`.
- Map 404 to `ErrNotFound`; decode `ZsrvResponse` errors via existing `zsrvResponseDiags`.
- Drop `config.encryptedSecrets` before it reaches state (zcli does the same).

### 4.2 Resource `zedcloud_app_instance_snapshot`

**Schema**

| Attribute | Type | Mode | Notes |
|---|---|---|---|
| `app_instance_id` | string | Required, ForceNew | |
| `triggers` | map(string) | Optional, ForceNew | Arbitrary values; a change re-takes the snapshot |
| `wait_for_ready` | bool | Optional, default `true` | Poll until `CREATED` |
| `skip_capability_check` | bool | Optional, default `false` | Mirrors zcli `-f`; skips §4.4 |
| `id` | string | Computed | Snapshot UUID |
| `name` | string | Computed | Server-generated |
| `state` | string | Computed | From the ganges state endpoint |
| `config_json` | string | Computed, Sensitive | Snapshot's app-instance config as JSON, secrets stripped. JSON string rather than the full `AppInstance` nested schema (huge, and only informational) |

Timeouts: `create` 20m, `delete` 10m (state transitions depend on the device checking in).

**Create**
1. Cluster guard, §4.4a (always, not affected by `skip_capability_check`).
2. Capability pre-check, §4.4 (unless `skip_capability_check`). Fails before any API write.
3. `Create(app_instance_id)`. Take `id`/`name` from `objectId`/`objectName`.
   Fallback for controllers older than ENG-1676 (empty `objectId`): `List(app_instance_id)` — there is exactly one snapshot after create (§1.2).
4. `SetId(id)` immediately, so a timeout below leaves a tainted resource instead of an orphan.
5. If `wait_for_ready`: poll `State(id)` until `CREATED`; `CREATE_FAILED` → error (include snapshot error text if exposed); timeout → error.
6. Read.

**Read** — `Get(id)` + `State(id)`. `ErrNotFound` → `SetId("")` (normal case: a newer snapshot on the same app instance replaced it, see §5.1).

**Update** — none; everything is ForceNew except `wait_for_ready`/`skip_capability_check` (plain `Optional`, no API call).

**Delete** — `Delete(app_instance_id, name)`; 404 is success. Poll `Get` until `ErrNotFound` or `DELETE_FAILED` within the delete timeout.

**Import** — by snapshot UUID (`terraform import zedcloud_app_instance_snapshot.x <uuid>`); `app_instance_id` comes from `config.id`.

### 4.3 Resource `zedcloud_app_instance_snapshot_activation`

| Attribute | Type | Mode | Notes |
|---|---|---|---|
| `app_instance_id` | string | Required, ForceNew | |
| `snapshot_name` | string | Required, ForceNew | Any snapshot, Terraform-managed or not |
| `force` | bool | Optional, ForceNew, default `false` | Sent as `?force=true` — re-applies a snapshot already marked `activate` |
| `triggers` | map(string) | Optional, ForceNew | Re-run the rollback |
| `id` | string | Computed | `<app_instance_id>/<snapshot_name>` |
| `snapshot_id` | string | Computed | |
| `state` | string | Computed | Last state from the ganges state endpoint (`ACTIVE` after a successful apply) |
| `activated_at` | string | Computed | `timestamp` from the state endpoint when `ACTIVE` was seen |

Timeouts: `create` 20m (the rollback finishes only after the device reports back).

- **Create:**
  1. Cluster guard (§4.4a).
  2. `GetByName` → `snapshot_id`; fail early with a clear message if missing. `State(snapshot_id)` → `prev_state`.
  3. `Activate(app_instance_id, snapshot_name, force)`. If `prev_state` is `ROLLBACK_FAILED`, send `force=true`
     even when `force = false`: the controller left `activate=true` on the failed attempt, so without `force`
     the retry would be a silent no-op (§1.2 "Activate semantics").
  4. `SetId` (so a failure below leaves a tainted resource; the next apply retries the rollback).
  5. **Always wait** for the rollback result (D4, no opt-out): poll `State(snapshot_id)` until
     `ACTIVE` → success; `ROLLBACK_FAILED` → error; create timeout → error saying the device hasn't confirmed the
     rollback yet (device offline?) and that the controller config has already been changed.
  6. Store `state` and `activated_at`.
- **Read:** refresh `state`; if the snapshot is gone keep the resource (the rollback already happened) and only update `state = ""`. Not removing it avoids a surprise second rollback on the next apply.
- **Delete:** no API call; returns a warning diagnostic *"rollback cannot be undone; removing from state only"*.
- **Update:** none (all ForceNew).

### 4.4 Capability pre-check (in scope)

zcli refuses to create a snapshot unless the device's status contains
`API_CAPABILITY_VOLUME_SNAPSHOTS_IMMEDIATE`. The controller doesn't check this. Without the
check, a create on an old EVE would stay in `CREATING` until our timeout: a slow, unclear failure.
The provider does the same check, on by default.

**Flow (snapshot resource Create, before the POST):**
1. `GET /v1/apps/instances/id/{app_instance_id}` (existing generated client) → `deviceId`. The same read
   serves the cluster guard (§4.4a), so this costs no extra call.
2. Cluster app instances never get here (§4.4a), so `deviceId` is always the node to check.
   An empty `deviceId` on a non-cluster app instance → error (shouldn't happen; avoids sending a bad status request).
3. `GET /v1/devices/id/{deviceId}/status` → `rawStatus`.
4. `strings.Contains(rawStatus, "API_CAPABILITY_VOLUME_SNAPSHOTS_IMMEDIATE")`, the same substring test zcli uses.
   (`rawStatus` is EVE's JSON-encoded info message. Parsing it would be stricter, but its shape belongs to EVE and
   isn't in our models; a substring match stays in step with zcli.)
5. Missing → **error** diagnostic (zcli only prints and exits 0; Terraform must fail the apply):
   *"edge node <name/id> does not report API_CAPABILITY_VOLUME_SNAPSHOTS_IMMEDIATE; upgrade EVE-OS or set skip_capability_check = true"*.
   Empty `rawStatus` (node never reported, or offline since onboarding) → same error, worded as "capability unknown".

**Where the calls live:** add `DeviceRawStatus(ctx, deviceID)` to the
hand-written snapshot client and unmarshal into the existing `models.DeviceStatusMsg`. That's cheaper than regenerating
`v2/client/node` for one operation. If someone regenerates the node client later, switch to the generated method.

**Activation:** no capability check. A snapshot can only exist if the create succeeded on that node, and zcli doesn't
check on activate either. The capability result is not stored in state; it's checked again on every create.

### 4.4a Edge node cluster guard (in scope)

Immediate snapshots are **not supported for app instances deployed to an edge-node cluster** (decided
2026-10-15, §8.1 D3). The provider rejects them up front instead of relying on the controller.
(The controller does reject KubeVirt devices, §1.2, and cluster nodes usually run KubeVirt, but that is a side
effect, not a cluster rule, and the error it returns doesn't mention clusters.)

- **Test:** the app instance's `clusterID` is non-empty (`models.AppInstance.ClusterID`, already in `v2/models/app_instance.go`).
- **Where:**
  - `zedcloud_app_instance_snapshot` Create and `zedcloud_app_instance_snapshot_activation` Create.
  - Also in `CustomizeDiff` of both resources **when `app_instance_id` is known at plan time** (existing app instance),
    so `terraform plan` already fails. When the ID is unknown (app instance created in the same apply), the Create check catches it.
- **Not bypassable:** `skip_capability_check` does not skip it. It is a product rule, not a device-capability guess.
- **Error:** *"application instance <name> is deployed to edge node cluster <clusterID>; immediate snapshots are not supported for edge node clusters"*.
- **Read / Delete:** no guard, so a snapshot that somehow exists (e.g. created before the app instance was moved to a
  cluster) can still be refreshed and destroyed.
- **Data source:** no guard; listing is read-only and returns whatever the controller has (normally nothing).

### 4.5 Data source `zedcloud_app_instance_snapshots`

Input `app_instance_id`. Output `snapshots` list of `{id, name, state}`; `state` comes from one
`State()` call per snapshot (usually 1 item, since create keeps only one). Optional `name` filter
gives `snapshot-show <name>` parity.

### 4.7 `allow_snapshots` / `max_snapshots` on `zedcloud_application_instance` (in scope, D5)

**What they do on the controller**

| Field | Behaviour | Where |
|---|---|---|
| `allowSnapshots` (bool) | Stored as-is from every create/update request. Required for app-update snapshots: a non-immediate snapshot create is rejected when it's false and `maxSnapshots > 0` | `seine/appinstproc.go:228`, `aisnapproc.go:150-160` |
| `maxSnapshots` (uint32) | Stored as-is; sent to the device. **On read, 0 is returned as 1** for non-docker-compose apps (`defaultMaxSnapshots`) | `appinstproc.go:229, 106, 409-411`; `purus/appinst.go:388` |
| Docker-compose apps | `allowSnapshots = true` or `maxSnapshots != 0` → validation error | `appinstproc.go:6449-6462` |
| When an app-update snapshot is taken | When an app instance update (`PUT /v2/apps/instances/id/{id}`) or refresh is sent with `?makeSnapshot=true`. The provider's update uses the v1 route and never sends it | `gilas/appinstance_handlers.go:115, 495-500, 720-725`; `seine/appinstproc.go:5379-5393` |
| `availableSnapshots` | Read-only list; not exposed, `data.zedcloud_app_instance_snapshots` covers it | |

**Provider changes**

| Attribute | Type | Mode | Notes |
|---|---|---|---|
| `allow_snapshots` | bool | Optional, Computed | Computed so a value set in the UI is kept and sent back on update (fixes the reset bug in §1.3) |
| `max_snapshots` | int | Optional, Computed, `validation.IntAtLeast(0)` | Computed because the controller reads 0 back as 1 |

- `v2/models/app_instance.go`: add `AllowSnapshots bool` (json `allowSnapshots,omitempty`) and `MaxSnapshots int64`
  (json `maxSnapshots,omitempty`). Hand-edit, as was done for the docker-compose fields; matches the swagger definition.
- `v2/schemas/application_instance.go`: schema, `ApplicationInstanceModel`, `ApplicationInstanceModelFromMap`,
  `SetApplicationInstanceResourceData`, `SetApplicationInstanceSubResourceData`, `GetApplicationInstancePropertyFields`.
  Data source gets them as Computed automatically through the shared schema.
- **Plan-time validation** (`CustomizeDiff`), when `app_type` is known:
  - `APP_TYPE_DOCKER_COMPOSE` with `allow_snapshots = true` or `max_snapshots > 0` → error (the controller rejects it anyway; this fails at `plan`).
    For docker-compose apps both attributes also get a `DiffSuppressFunc` like the existing `diffSuppressIfFieldValueEqual("app_type", "APP_TYPE_DOCKER_COMPOSE")`, so the Computed values don't cause noise.
  - Edge-node-cluster app instances with `allow_snapshots = true` → error, by extension of D3 (§8.2 Q1 asks to confirm).
- **Optional+Computed caveat** (document it): removing `allow_snapshots` from the config does **not** turn it off; the old value
  stays. To turn it off, set `allow_snapshots = false` explicitly.
- **Out of scope here:** sending `?makeSnapshot=true` from Terraform. That needs the provider to switch its update to the v2
  route, which changes the async update path from UE-177. Listed as §8.2 Q2.

### 4.6 Files

| File | Change |
|---|---|
| `v2/client/app_instance_snapshot/client.go`, `client_test.go` | New hand-written client (snapshot routes + `DeviceRawStatus`) + httptest tests |
| `v2/resources/app_instance_snapshot_guards.go`, `_test.go` | §4.4 capability check and §4.4a cluster guard (shared by both resources) |
| `v2/client/zedcloud_api.go` | Wire `AppInstanceSnapshot` |
| `v2/resources/app_instance_snapshot.go`, `_test.go` | Snapshot resource |
| `v2/resources/app_instance_snapshot_activation.go`, `_test.go` | Activation resource |
| `v2/resources/app_instance_snapshots_data_source.go` | List data source |
| `v2/resources/provider.go` | Register 2 resources + 1 data source |
| `v2/examples/resources/zedcloud_app_instance_snapshot*/` | Examples |
| `docs/resources/app_instance_snapshot*.md`, `docs/data-sources/app_instance_snapshots.md` | Docs, including §5 caveats |
| `v2/models/app_instance.go` | `AllowSnapshots`, `MaxSnapshots` (§4.7) |
| `v2/schemas/application_instance.go`, `v2/resources/application_instance.go` | `allow_snapshots`, `max_snapshots` schema, mapping, `CustomizeDiff` (§4.7) |
| `v2/resources/application_instance_test.go`, `testdata/application_instance/*.tf` | Tests for the new attributes |
| `docs/resources/application_instance.md`, `docs/data-sources/application_instance.md` | Document the new attributes and the Optional+Computed caveat |

---

## 5. Behaviour and edge cases

### 5.1 Only one snapshot per app instance
Every create deletes the app instance's previous snapshots. Consequences:
- Two `zedcloud_app_instance_snapshot` resources on the same app instance fight: each apply that creates one makes the other vanish; its next refresh removes it from state and the following plan re-creates it. Document as *"one resource per app instance"*. A `CustomizeDiff` can't detect this across resources; we could add a warning at create time if `List` returns a snapshot whose id isn't ours.
- `create_before_destroy` is safe: the new create removes the old one server-side, the old delete then gets 404 = success.
- An activation resource pointing at a replaced snapshot fails at create time with "snapshot not found" (good), and keeps its state if already applied (§4.3).

### 5.2 Rollback is reflected as ordinary drift (decided, D4)

Rule: **if a rollback was requested and the controller rolled the app instance back, Terraform state must show the
rolled-back config, the same way it shows any other drift.** So the provider must (1) check the rollback status and
(2) only treat it as done when it succeeded.

How this works with SDKv2:
- The activation resource confirms the rollback (§4.3 step 5: waits for `ACTIVE`, fails on `ROLLBACK_FAILED`).
  A failed or unconfirmed rollback fails the apply, so nothing claims success.
- One resource can't write another resource's state in SDKv2, so `zedcloud_application_instance` picks up the
  rolled-back config on its **next refresh** (the next `plan`/`apply`, or `terraform apply -refresh-only`). Its Read
  already reads the full config from the controller, so no new code is needed there; the rollback appears like any
  out-of-band change.
- No special handling: no diff suppression, no auto-`ignore_changes`. As with other drift, the user picks:
  - keep the rollback → update the HCL to match (or `ignore_changes` for those fields), or
  - discard it → `terraform apply` pushes the HCL config again, which undoes the rollback.
  The resource docs and the example spell this out, including `terraform apply -refresh-only` to accept the rollback into state.
- Ordering: if the same apply also changes `zedcloud_application_instance`, the two writes race and the result is
  undefined. The docs recommend running a rollback in its own apply, with no pending app instance changes.

Verification: an acceptance test (§7) checks that after the activation apply, `terraform plan` on the app instance
shows the rolled-back values as drift, and that `apply -refresh-only` makes the plan empty.

### 5.3 App-update snapshots
Turned on with `allow_snapshots` / `max_snapshots` (§4.7). They live in the same table as immediate snapshots, so
**creating an immediate snapshot deletes the app-update snapshots too** (§5.1); the docs warn about this. The activation
resource and the data source work with both kinds, since they only need a snapshot name.

### 5.4 Offline devices
State transitions need device reports. With an offline device, waits time out. For the snapshot resource the error suggests `wait_for_ready = false`.
The activation resource always waits (D4): its error says the rollback is requested but unconfirmed, and the resource is tainted so the next apply checks again.

### 5.5 KubeVirt / error-state / cluster app instances
KubeVirt devices and error-state app instances are rejected by the controller (§1.2); we surface the `ZsrvResponse` error as-is.
Cluster app instances are rejected by the provider before any API write (§4.4a).

---

## 6. zedcloud follow-ups (non-blocking)

1. Add `google.api.http` annotations for create (`POST .../snapshot`) and activate (`PUT .../activate`) so they appear in swagger; then the provider can switch to a generated client.
2. Fill `state` / `activate` / `immediate` in get/list responses (seine leaves them empty; zcli's table prints blanks too).
3. Check `API_CAPABILITY_VOLUME_SNAPSHOTS_IMMEDIATE` server-side so every client (UI, zcli, Terraform, API users) gets the same guard. The provider keeps its own check until this ships.
4. Reject create/activate for cluster app instances server-side with an explicit error (today only the indirect KubeVirt check applies). zcli has no cluster guard either.

---

## 7. Testing

- **Client unit tests** (httptest): every route and verb, `force` query param, 404 → `ErrNotFound`, `ZsrvResponse` error decoding, `objectId` fallback to `List`, `encryptedSecrets` stripping, `DeviceRawStatus`.
- **Capability unit tests:** capability present → proceeds; absent → error and **no POST sent**; empty `rawStatus` → error; `skip_capability_check = true` → no status call.
- **Cluster guard unit tests:** cluster app instance → error on snapshot create and on activation create, **no POST/PUT sent**, even with `skip_capability_check = true`; `CustomizeDiff` fails the plan when `app_instance_id` is known and skips the check when it is unknown; Read/Delete of an existing snapshot still work.
- **Resource unit tests** (mock `ClientService`): create → poll `CREATING→CREATED`; `CREATE_FAILED`; timeout leaves ID set; read 404 removes from state; delete 404 ok; activation `ROLLBACK_FAILED` → error and resource tainted; retry after `ROLLBACK_FAILED` sends `force=true`; activation timeout message; activation delete is a no-op with a warning; import.
- **Acceptance** (real node, gated like `application_instance_real_node_test.go`): create → change app instance → activate (wait for `ACTIVE`) → `plan` shows the rolled-back values as drift on `zedcloud_application_instance` → `apply -refresh-only` → empty plan → destroy on an EVE with the immediate-snapshot capability; second create replaces the first; activation with `force = true` re-applies.
- **App instance attribute tests:** create/update/read round-trip of `allow_snapshots` / `max_snapshots`; `max_snapshots`
  unset → reads back 1 with no diff; update without the attributes in config keeps the controller value (regression test for the
  reset bug in §1.3); docker-compose + `allow_snapshots = true` fails at plan; no diff for docker-compose apps; import.
- `go vet`, `golangci-lint`, `make docs` (if generated).

---

## 8. Decisions and open questions

### 8.1 Decided (2026-10-15)

| # | Question | Decision |
|---|---|---|
| D1 | Should Terraform support activation (rollback)? | **Yes.** `zedcloud_app_instance_snapshot_activation` as in §4.3: a one-shot resource whose destroy only removes it from state. |
| D2 | Capability pre-check? | **Yes.** On by default for snapshot create, same test as zcli, `skip_capability_check` as the `-f` escape hatch (§4.4). Errors instead of zcli's print-and-exit. |
| D3 | Edge node clusters? | **Not supported.** Provider rejects snapshot create and activation for cluster app instances at plan time when possible, otherwise at apply, regardless of `skip_capability_check` (§4.4a). |
| D4 | Rollback vs. Terraform state? | **Ordinary drift.** After a rollback, check its status; if it succeeded, the rolled-back config must show in Terraform state like any other drift (§5.2). Activation always waits for `ACTIVE`. |
| D5 | `allow_snapshots` / `max_snapshots` in PLE-576? | **Yes.** Added to `zedcloud_application_instance` as Optional+Computed (§4.7). |
| D6 | Resource names? | **Shorter prefix:** `zedcloud_app_instance_snapshot`, `zedcloud_app_instance_snapshot_activation`, `data.zedcloud_app_instance_snapshots`. |

### 8.2 Still open

1. **`allow_snapshots` on edge-node-cluster app instances:** D3 rules out immediate snapshots for clusters. Should the provider
   also reject `allow_snapshots = true` for cluster app instances (current proposal, §4.7), or are app-update snapshots
   allowed there?
2. **`makeSnapshot` on update:** should Terraform be able to *take* an app-update snapshot when it updates an app instance
   (e.g. a `snapshot_on_update` attribute that sends `?makeSnapshot=true`)? Without it, `allow_snapshots` only matters for
   updates made from the UI or other clients. Adding it means moving the provider's update to the v2 route, so the
   proposal is a follow-up ticket.
