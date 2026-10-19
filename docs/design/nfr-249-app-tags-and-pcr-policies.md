# Design: NFR-249 — Edge app tags and PCR policies via Terraform

**Ticket:** [NFR-249](https://zededa.atlassian.net/browse/NFR-249) — *Terraform Provider support for marketplace app tags and PCR policies*
**Customer:** customer request (via ZD-3895). Follow-up to the two controller features built for that customer:
[NFR-96](https://zededa.atlassian.net/browse/NFR-96) (tags on marketplace Edge apps, API epic
[ZEDCLOUD-3310](https://zededa.atlassian.net/browse/ZEDCLOUD-3310)) and
[NFR-189](https://zededa.atlassian.net/browse/NFR-189) (flexible PCR policy, epic
[ZEDCLOUD-2157](https://zededa.atlassian.net/browse/ZEDCLOUD-2157)).
**Repos:** `zededa/terraform-provider-zedcloud` @ `e5dddd7e` (origin/main), `zededa/zedcloud` @ `6197300bc2` (main), `zededa/zcli` @ `79b1feb`
**Date:** 2026-10-08 (revised 2026-10-13 with the zcli analysis, §1.3; 2026-10-16 with the tags minimum controller
version §1.1, EVE `PolicyID` semantics §8 Q1, the "follow the current API" ground rule §8, and in-place deployment update §4.5)
**Author:** Ivan Curachin
**Status:** Implemented on `nfr-249-app-tags-and-pcr-policies`, 2026-10-16, and tested against the local controller. See §9 for where the
implementation differs from this design. Controller gaps are tracked as non-blocking follow-ups (§6.2).

**Headline: no zedcloud work is needed. The provider follows the current API (§8 ground rule).** The work is:

- new attributes on four existing resources: `zedcloud_application.tags`, `zedcloud_project.pcr_policy`,
  `zedcloud_edgenode.pcr_policy` (the customer's must-have), and the `zedcloud_deployment` PCR device policy;
- **in-place `zedcloud_deployment` update** (today a stub that does nothing), using the API's "new version" PUT (§4.5);
- one small new client (job status, used to wait for deployment updates). No new resources.

Minimum controller: **zedcloud 20.0.7 / 21.0.0+** for tags; 19.2.24+ for PCR policies. Estimate **9–11 days**.

Both features are already live in the controller API. The provider's swagger and models are partly stale, and its
schemas don't expose either feature. Two findings matter more than the feature gap itself:

1. **Terraform already deletes this data today.** `zedcloud_application` and `zedcloud_edgenode`
   updates send the whole object. The controller treats a missing `tags` / `pcrPolicy` as "clear
   it". So any `terraform apply` that touches an app wipes tags set in the UI, and any apply that
   touches an edge node wipes a device PCR policy set with zcli (§1.4).
2. **On a device GET, `pcrPolicy` can come from a deployment, not the device**, and the response doesn't say which. The provider has to handle this or it will copy deployment policy into device config (§4.3).

---

## 1. Problem Statement

The customer manages tenants with Terraform and wants to:

- set **tags on Edge applications** (marketplace apps), so their portal can filter apps by tag;
- set **PCR policies** (which TPM PCRs are used to seal the volume key) on **projects** and on
  **individual edge nodes**.

Both are possible through the UI/API/zcli, but not through `terraform-provider-zedcloud`.

### 1.1 What was done in zedcloud: Edge app tags (NFR-96)

Commit `b64a96fdaa` ([zedcloud#8255](https://github.com/zededa/zedcloud/pull/8255), 2026-07-06) on main, cherry-picked to
`20.0-point-release` as `b96e857c9b`. **Minimum controller version: zedcloud 20.0.7** (tagged 2026-08-05). It's also
in every later 20.0.x (through 20.0.20 at the time of writing) and in **21.0.0+**. Not in 19.x or 20.0.0–20.0.6.
Verified per tag: at 20.0.6 the migration, the `tags` field in `app.proto`, gilas `matchAnyTag`, the seine `app.Tags` binding and the
swagger `tags` property are all missing. At 20.0.7 and 21.0.0 all are present. The cherry-pick touches the same 16 files with
the same line counts as the original. (The "20.0.3" in ZEDCLOUD-3310 is the **old UI** version, not the controller's;
Cortex UI is 1.1.4.)
Spec: `zedcloud/docs/design/app-filtering-by-tags/NFR-96-tags-spec.md`.

| Thing | Actually | Where |
|---|---|---|
| Field | `map<string,string> tags` on `App` (19), `App_HttpRequest` (19), `AppSummary` (16) | `libs/zmsg/app/app.proto` |
| Storage | `user_config.app_data.tags jsonb` + GIN index | `srvs/seine/migrations/20260703120000_addtags.tx.up.sql` |
| Create/update | `bindAppData` sets `app.Tags = uInput.GetTags()` unconditionally; `appdb.Update` writes the whole row (`bun:"tags,nullzero"`), so **omitting `tags` on PUT clears them** | `srvs/seine/appproc.go:206`, `srvs/seine/appdata.go:267` |
| Validation | `zutils.ValidateTag`: ≤32 tags; key 3–512 chars, no `< > % & ? / \ ' " ` `` ` ``; value 3–256 chars, no `' " ` `` ` `` `\` | `libs/zutils/validate.go:330` |
| Filtering | `GET /v1/apps?tags={"k":"v"}&matchAnyTag=true` (JSON object in a query param; AND by default, OR with `matchAnyTag`) | `srvs/gilas/url.go:123`, `srvs/gilas/app_handlers.go`, `srvs/seine/appproc.go:1522` |
| Scope | Edge app (bundle) only. Not app instances, not app profiles | — |

### 1.2 What was done in zedcloud: flexible PCR policy (NFR-189)

Commit `49f0a66098` ([zedcloud#7763](https://github.com/zededa/zedcloud/pull/7763), 2026-03-12). NFR-189 says
it was released with **ZedCloud 19.2.24 / EVE 16.0.1-LTS**. Feature/systest doc:
`zedcloud/docs/feature-docs/flexible-pcr-policy-systest.md`. zcli has `edge-node pcr-policy show|set|reset`
and `project create|update --policy-config=pcr:<file>`.

There are **three** places the API accepts a PCR policy:

| Level | API field | Type | Stored as |
|---|---|---|---|
| Project | `Tag.pcrPolicy` (19) | `PolicyConfig{type: POLICY_TYPE_PCR, pcrPolicy: PCRPolicy{pcrIndices, policyVersion}}` | A row in the policy table, attached to the project (same path as attestation/PNAC policies) |
| Edge node | `DeviceConfig.pcrPolicy` (82) | `PolicyConfig` (same as above) | `device_data.pcrpolicyconfig BYTEA`, the proto-marshalled `PolicyConfig` |
| Deployment (device policy) | `DevicePolicy.pcrPolicy` (4) with `policySubType = DEVICE_POLICY_TYPE_PCR` | `DevicePCRPolicy{pcrIndices, policyVersion}` | Deployment policy row; versioned by the server (`-pcr-vN`) |

Resolution, as `thames` computes it (`srvs/thames/pcrpolicy.go`): **device → project → none (EVE defaults)**.
The resolved `policyVersion` is sent to EVE as `PolicyID`.

**For nodes in a deployment project the order is different.** thames gets the device through seine's device read
(`fetchAttestPolicy` → `fetchDeviceInfo`, `srvs/thames/devutils.go:199-227`). When the device has a
`DeploymentTag`, it returns **no project info** (`pInfo = nil`). In that read, a matching deployment
`DEVICE_POLICY_TYPE_PCR` policy has already replaced the device-level one. So for these nodes:
**deployment PCR device policy → node-level policy → EVE defaults**. Project-level PCR is ignored, and node-level
only takes effect if the deployment has no PCR device policy matching the node.

Validation is the same in all three places (`validatePCRIndices` in seine, the gilas deployment pre-check, and thames):

| Rule | |
|---|---|
| Non-empty | `pcrIndices` must not be empty |
| Range | 0–15 |
| PCR 0 | required |
| PCR 5 | rejected (volatile) |
| Duplicates | rejected |

Behaviour that matters for Terraform (verified in code):

| Behaviour | Where | Consequence |
|---|---|---|
| Device PUT with no `pcrPolicy` sets the column to NULL (**reset**) | `srvs/seine/devproc.go:1235-1249` (`bindDeviceData`) | The provider has to send the policy on every update or it gets cleared |
| Device GET: if the device's project is a deployment with a matching `DEVICE_POLICY_TYPE_PCR` device policy, **`pcrPolicy` is overwritten** with the deployment's policy (`type` + `pcrPolicy` only, no id or name) | `srvs/seine/devproc.go:~3776-3790` | Read can't tell device-level and deployment-derived apart (§4.3) |
| Project update with no `pcrPolicy` does nothing. With one, it calls `createOrUpdatePolicy` | `srvs/seine/tagproc.go:1454` | **A project PCR policy can't be removed** through the API once set |
| Project create/attach: the server overwrites the policy `name` with `<project>_<random15>` | `srvs/seine/tagproc.go:666`, `:1912` | `name` must be Computed |
| `policyVersion` for **deployment** policies is server-managed (1, then +1 on each update) | `srvs/seine/deploymentpolicyproc.go:1013-1054` | Computed in Terraform |
| `policyVersion` for **project/device** policies: **client-managed**. seine stores whatever it receives and thames passes it to EVE as `PolicyID` | seine (no increment); zcli (§1.3) | The systest doc's "server ignores user-supplied policyVersion" is wrong. The provider must set it (§4.2, §4.3) |
| Index order: the server doesn't sort (`validatePCRIndices`). **zcli sorts before sending** (§1.3), so that's where the systest doc's "sorted" claim comes from | zcli `_validate_pcr_indices` | Treat the indices as a **set** in Terraform and send them sorted |

### 1.3 How zcli manages PCR policies (reference client)

Repo: `zededa/zcli` @ `79b1feb`. PRs [#1146](https://github.com/zededa/zcli/pull/1146) and [#1148](https://github.com/zededa/zcli/pull/1148)
(`d57feea`, `d35fcbd`). Code: `zcli/modules/edge-node.py:269-415, 1353-1462`; `zcli/modules/project.py:222, 238-269, 450, 456-505`;
template `zcli/config-templates/pcr-policy-config.json`.

**Edge node (`edge-node pcr-policy show|set|reset`)**: read-modify-write of the whole device:

| Step | What zcli does | Where |
|---|---|---|
| Validate | Range 0–15, not 5, no duplicates, must contain 0. Returns the indices **sorted and de-duplicated** | `_validate_pcr_indices` (`edge-node.py:320`) |
| Read | `GET devices/name/{name}` (or by id) and takes `pcrPolicy` from the response | `edge-node.py:1401-1408` |
| Version | **`new_version = existing.pcrPolicy.policyVersion + 1`** (0 → 1 when none), computed by the client on every `set`, even if the indices didn't change (with `--force`) | `edge-node.py:1409` |
| Build | `{type: POLICY_TYPE_PCR, pcrPolicy: {pcrIndices, policyVersion}}`, **keeping the existing `id` and `name`** if the GET had them. Never sends `title` | `_build_pcr_policy` (`edge-node.py:405`) |
| Write | `PUT devices/id/{id}` with the whole device object | `edge-node.py:1423-1425` |
| Reset | `device.pop('pcrPolicy')` then `PUT` the whole device. **Leaving the field out is the reset**, which confirms §1.2 | `edge-node.py:1457-1459` |
| Safety | Shows the diff and warns about deviations from the default `[0,1,2,3,4,6,7,8,9,13,14]`. Prompts unless `--force --confirm-risks` | `edge-node.py:340-403` |

**Project (`project create|update --policy-config=pcr:<file>`)**: a generic policy path with nothing PCR-specific:

| Step | What zcli does | Where |
|---|---|---|
| Create | Sends `{pcrPolicy: {type, pcrPolicy: <file.pcrPolicy>}}` exactly as in the file. No index validation on the client | `fetch_policy_config` (`project.py:238`) |
| Update | Requires the project to **already have** a `pcrPolicy` (`"No such pcrPolicy is attached to this project"`), then sends `{id: <existing id>, type, pcrPolicy: <file>}` | `process_policy_update` (`project.py:456`) |
| Version | **Taken from the file as-is.** The template ships `"policyVersion": 1`, and zcli never bumps it. Repeated updates send version 1 with different indices unless the user edits the number | template, `project.py:495-499` |
| Removal | No command. Same API gap as §1.2 | — |

**Deployment device policies (`--dev-policy-config`)**: raw JSON passthrough, nothing PCR-specific. The server manages versions there.

What this tells us:

- **Policy version for edge nodes is owned by the client, and the convention is "previous + 1".** The provider should do the same (§4.3).
- **Policy version for projects has no convention.** zcli leaves it to the user, so in practice it stays at 1. The provider can do better by bumping it on change (§4.2). EVE confirmed it applies a policy when either the indices or `PolicyID` change, so zcli's practice is functionally fine. The version is bookkeeping (§8 Q1).
- **zcli has the same deployment-vs-device problem.** `pcr-policy show` prints a deployment-derived policy as "device-level", and `set` bumps from the deployment's version. The provider isn't worse than the reference client here, and the controller fix in §6.2 item 2 helps both.
- **Only zcli sends id/name on the device path, and it's harmless:** seine stores the device `PolicyConfig` as an opaque blob. The provider can leave them out.
- zcli `project update` can't **attach** a PCR policy to a project that doesn't have one, even though the server supports it (`createOrUpdatePolicy`). That's a zcli gap, not an API one. The provider sends the block either way.

### 1.4 Provider ground truth

| Thing | State | Where |
|---|---|---|
| `zedcloud_application` `tags` | **Missing.** Not in schema or model. `swagger/zedge_app_service.swagger.json` predates NFR-96 (0 matches for `tags`/`matchAnyTag` on App) | `v2/schemas/application.go:258`, `v2/models/app.go` |
| App update | Sends `ApplicationModel(d)` (full object, no tags) → **clears UI-set tags** on every app update | `v2/resources/application.go:210, 237` |
| `zedcloud_project` PCR policy | **Missing.** `models.Tag` and `models.Policy` have no `PcrPolicy`; `PolicyType` enum lacks `POLICY_TYPE_PCR` | `v2/models/tag.go`, `v2/models/policy_config.go`, `v2/models/policy_type.go` |
| `zedcloud_edgenode` PCR policy | **Half there.** `models.Node.PcrPolicy *Policy` exists (from a partial regen in `b21396f8`), but `schemas/node.go` doesn't expose it | `v2/models/device_config.go:148` |
| Edge node update | Sends `NodeModel(d)` (full object, `PcrPolicy` nil) → **clears zcli-set device PCR policy** on every node update | `v2/resources/node.go:297-304` |
| `zedcloud_deployment` PCR device policy | **Missing.** `models.DevicePolicy` has no `PcrPolicy`, and `DevicePolicyType` lacks `DEVICE_POLICY_TYPE_PCR` | `v2/models/device_policy.go`, `v2/models/device_policy_type.go`, `v2/schemas/device_policy.go` |
| Node swagger | **Already current** (`PCRPolicy`, `DevicePCRPolicy`, `POLICY_TYPE_PCR`, `DEVICE_POLICY_TYPE_PCR` all present) | `swagger/zedge_node_service.swagger.json:13659, 17178, 17687` |
| Closest precedent | PNAC ([#204](https://github.com/zededa/terraform-provider-zedcloud/pull/204), `a64dd172`) added a policy at **all three levels**: project policy, node, and deployment device policy. Mirror its file set and tests | `v2/schemas/{tag,policy_config,device_policy}.go`, `v2/resources/testdata/{project,node}/*pnac*` |

---

## 2. Proposed Solution

No new resources. Add attributes to four existing resources (application, project, edge node,
deployment), and fix the update paths so they stop deleting data Terraform doesn't manage.

```hcl
resource "zedcloud_application" "app" {
  name  = "edge-router"
  title = "Edge router"
  # ...
  tags = {
    customer = "acme"
    tier     = "gold"
  }
}

resource "zedcloud_project" "p" {
  name  = "fleet-a"
  title = "Fleet A"
  type  = "TAG_TYPE_PROJECT"

  pcr_policy {
    title = "fleet-a-pcr"
    type  = "POLICY_TYPE_PCR"
    pcr_policy {
      pcr_indices = [0, 1, 2, 3, 4, 6, 7, 8, 9, 13, 14]
    }
  }
}

resource "zedcloud_edgenode" "n" {
  name       = "node-17"
  project_id = zedcloud_project.p.id
  # ...
  pcr_policy {
    pcr_indices = [0, 2, 3, 4, 6, 7, 8, 9, 13, 14]   # e.g. drop PCR 1 on this node only
  }
}
```

---

## 3. Alternatives Considered

| Option | Why not / why |
|---|---|
| **A. Separate resources** (`zedcloud_project_pcr_policy`, `zedcloud_edgenode_pcr_policy`, `zedcloud_application_tags`) | Lets Terraform manage the field without owning the whole parent, the same way `aws_ec2_tag` does. But the controller has no PATCH: each would need a read-modify-write of the whole parent and would fight the parent resource. Too much for three fields, and nothing else in this provider works that way. Rejected for now. |
| **B. Attributes on the parent resource (chosen)** | Matches how the API models it and follows the PNAC precedent. Combined with "preserve when not configured" (§4.3) it doesn't delete data managed outside Terraform. |
| **C. Edge node `pcr_policy` reuses the full `Policy()` block** (like project policies) | Consistent with projects, but it drags in ~15 unrelated nested blocks (`app_policy`, `network_policy`, ...) plus `name`/`title`/`type`, and the device API ignores all of them. **Rejected:** use a small dedicated block, and have the provider wrap it in `PolicyConfig{type: POLICY_TYPE_PCR}`. |
| **D. Project `pcr_policy` as a small block too** | Nicer HCL, but it differs from `attestation_policy` and `port_based_network_access_control_policy` on the same resource, and the project policy path needs the policy `id`/`name` round-tripped for updates. **Rejected:** keep the `Policy()` wrapper, like PNAC. |
| **E. `zedcloud_applications` data source filtering by tags** (`tags` + `match_any_tag`) | It would need a list data source the provider doesn't have yet. Customer (2026-10-14): "could be useful, but we do not have a use case at the moment". **Deferred** (§8 Q4). |

---

## 4. Interface / Schema Design

### 4.1 `zedcloud_application.tags`

| Attribute | Type | Mode | Notes |
|---|---|---|---|
| `tags` | map(string) | Optional, Computed | Same description as `zedcloud_project.tags`. Plan-time `ValidateDiagFunc` mirrors `zutils.ValidateTag`: ≤32 entries; key 3–512 and value 3–256 chars; key must not contain `< > % & ? / \ ' " ` `` ` ``; value must not contain `' " ` `` ` `` `\`. |

- **Model:** add `Tags map[string]string \`json:"tags,omitempty"\`` to `models.Application` and
  `models.AppSummary` (by hand, or regenerate from the controller's current `zedge_app_service.swagger.json`; see §5).
- **Expand/flatten:** `ApplicationModel`, `EdgeApplicationModelFromMap`, `SetApplicationResourceData`,
  `SetApplicationSubResourceData`, `GetApplicationPropertyFields`.
- **Update semantics (fixes the wipe):**
  - `tags` set in config: send it (authoritative; an empty map `{}` clears).
  - `tags` **not in config** (`GetRawConfig().GetAttr("tags").IsNull()`): send the server's
    current tags. In practice, carry the value from state, because `Computed` keeps the last read.
    UI-managed tags survive `terraform apply`.
  - Trade-off: with Optional+Computed, removing the `tags` argument entirely leaves the tags
    in place (documented; to clear them, set `tags = {}`). This is the usual SDKv2 behaviour for
    attributes that are partly managed outside Terraform.
- **Data source** `zedcloud_application` gets `tags` for free (same schema).
- **Older controllers (< 20.0.7 on the 20.0 line; all 19.x):** `omitempty` means nothing is sent unless the user sets
  `tags`. If they set it on an old controller, gilas may reject the unknown field or silently drop it. Document the
  minimum controller version (§8 Q3).

### 4.2 `zedcloud_project.pcr_policy`

The same pattern as `attestation_policy` / `port_based_network_access_control_policy`:

```text
pcr_policy            TypeList(MaxItems 1) of Policy()      Optional, Computed
  └─ pcr_policy       TypeList(MaxItems 1) of PCRPolicy()   (new nested block inside Policy())
       ├─ pcr_indices     TypeSet(int)  Required   elem: 0–15, != 5
       ├─ policy_version  int           Computed   (provider-managed, see below)
       └─ id              string        Computed
```

- **`policy_version`:** the controller doesn't manage it for projects (§1.2), and zcli just sends whatever's in the file (§1.3).
  The provider owns it: on create it sends `1`. On update it sends `<version from state> + 1` **only when
  `pcr_indices` changed**, and otherwise sends the stored value. A `CustomizeDiff` marks
  `policy_version` as `SetNewComputed` when the indices change, so the plan shows "(known after apply)".
  This is stricter than zcli and matches the edge-node convention.

- **Models:** `models.PCRPolicy` (new: `ID`, `PcrIndices []int64`, `PolicyVersion int64`),
  `models.Policy.PcrPolicy *PCRPolicy`, `models.Tag.PcrPolicy *Policy`, and `POLICY_TYPE_PCR` in
  `PolicyType` (also add the missing `POLICY_TYPE_LOC` and `POLICY_TYPE_CONFIGURATION_LOCK`
  while there, so the enum matches the proto).
- **Schemas:** `v2/schemas/pcr_policy.go` (new; `PCRPolicyModel`, `PCRPolicyModelFromMap`,
  `SetPCRPolicySubResourceData`, `PCRPolicySchema`), plus wiring in `policy_config.go` and `tag.go`
  (model, from-map, set, sub-resource set, schema, property fields). Copy the PNAC diff in `a64dd172`.
- **Validation:** element `ValidateFunc` for range and `!= 5`. A `CustomizeDiff` on the resource
  checks that PCR 0 is in the set (SDKv2 can't validate a whole `TypeSet` with `ValidateFunc`). The server checks everything again.
- **`type`:** default to / require `POLICY_TYPE_PCR`. Error at plan time on any other value.
- **`name`:** Computed (the server sets it to `<project>_<random>`).
- **Removal:** the API can't remove a project PCR policy (§1.2). With `Optional+Computed`,
  removing the block is a no-op and causes no diff. Document it. Removing it needs a controller fix (§6.2).

### 4.3 `zedcloud_edgenode.pcr_policy`

A small dedicated block. The provider wraps it in `PolicyConfig` on the wire:

```text
pcr_policy            TypeList(MaxItems 1)   Optional
  ├─ pcr_indices      TypeSet(int)           Required   elem: 0–15, != 5; CustomizeDiff: must contain 0
  └─ policy_version   int                    Computed   (provider-managed, same as zcli)
```

Wire mapping: `Node.PcrPolicy = &Policy{Type: POLICY_TYPE_PCR, PcrPolicy: &PCRPolicy{PcrIndices (sorted), PolicyVersion}}`.
This is the same payload zcli's `_build_pcr_policy` produces (§1.3). Leave `name`/`title` unset, as zcli does
(seine stores the blob as-is). **Check:** the `Policy` model marks `name`/`title`/`type` as `Required`,
so it serialises `"name": null` / `"title": null`. Confirm gilas accepts this, or add `omitempty`
handling. zcli and systest TC-04 both leave the fields out entirely, so leaving them out is known to work.

**`policy_version`:** follow zcli's convention. Set it to `<version in the GET done at node.go:280> + 1`
when `pcr_indices` changed (or the policy is new: `0 + 1`), and otherwise send the fetched value unchanged.
Unlike zcli, don't bump when nothing changed. EVE treats a `PolicyID` change alone as a new policy (§8 Q1), so a
bump on an apply with no changes would make it re-apply the same PCR set for no reason.
Basing it on the live GET instead of state also covers changes zcli made in between, so the version keeps increasing
across both clients.

**The hard part is ownership.** A device GET returns a `pcrPolicy` that is one of:

| Source | How it looks on GET |
|---|---|
| Device-level (set by TF/zcli/API) | `PolicyConfig` exactly as stored (whatever the writer sent) |
| Deployment device policy (overrides the device-level one in the response) | `{type: POLICY_TYPE_PCR, pcrPolicy: {pcrIndices, policyVersion}}`, no `id`/`name` |
| None | absent |

And a device PUT without `pcrPolicy` **clears** the device-level policy. Proposed behaviour:

| Config | Update sends | Read sets state |
|---|---|---|
| `pcr_policy` in config | The configured policy (authoritative) | The value from GET (drift detection) |
| `pcr_policy` removed from config (it was in prior state from config) | `nil` → server reset (same as `zcli edge-node pcr-policy reset`) | Absent |
| `pcr_policy` never in config | **`existingNode.PcrPolicy` as fetched** (already done at `node.go:280`), so a zcli-set policy survives | Not stored, so no diff |

How this works without `Computed`: **Read writes `pcr_policy` to state only if prior state already has it**,
which means Terraform manages it (it got there from config on create or update). An unmanaged
policy never enters state, so the plan shows no diff. A managed policy removed from config shows
a diff (state has it, config doesn't), and the update sends `nil`. Cost: `terraform import`
doesn't adopt an existing policy. Add the block to config and the first apply adopts it. Document this.

Known limitation (document it, §8 Q2): for a node in a **deployment** project with a PCR device
policy, the GET shows the deployment's policy. If the user also sets `pcr_policy` on the node, the
plan shows a perpetual diff (state shows the deployment's indices). And if `pcr_policy` was never in config,
the preserve path writes the deployment's policy back as the node's device-level policy. That's
harmless while the deployment policy exists, because the read shows the deployment's policy and so
does thames. It becomes visible if the deployment policy is later removed. A clean fix needs the controller (§6.2 item 2).

~~Heuristic: skip the preserve path when the fetched policy has no `id`/`name`.~~ **Dropped after reading zcli:**
a device-level policy set by zcli also has no `id`/`name` (`_build_pcr_policy` copies them only if they
already exist, and they never do on the device path). So the response looks the same either way, and the heuristic would drop real
zcli policies. Preserving the policy as fetched is the safer failure mode: it never deletes a policy, and the worst
case is a stale device-level copy of a deployment policy. zcli `pcr-policy set` already does this,
because it also writes back what it read. If this matters, the reliable provider-side check is:
the node's project is `TAG_TYPE_DEPLOYMENT` and its deployment has a `DEVICE_POLICY_TYPE_PCR` device
policy whose target matches the node's tags. That costs two extra GETs per update. Not proposed for v1.

Also apply the preserve logic in `publishBaseOS` / `applyBaseOS`, which send `newNode` too (`node.go:522-564`).

### 4.4 `zedcloud_deployment` device policy (in scope)

Not in the NFR text ("edge nodes and projects"), but it's the third API surface, and the
PNAC precedent already added `port_based_network_access_control_policy` here. Small:

- `models.DevicePCRPolicy` (new), `models.DevicePolicy.PcrPolicy`, `DEVICE_POLICY_TYPE_PCR` in `DevicePolicyType`.
- `v2/schemas/device_policy.go`: a `pcr_policy` block (`pcr_indices` set, `policy_version` Computed:
  server-managed here).
- Validation as above. gilas also rejects bad indices synchronously with a 400 (`deploymenthandlers.go`).

**Decision (NFR-249 comment, 2026-10-13):** in scope. Sergey: there's no reason to skip it given the small effort.

**Customer's answer (2026-10-14):** "We need at least to configure it at the **edge node level**. We use also project
deployments but we cannot update existing deployments since this is not supported from Terraform."

What that means:

- **Edge-node PCR (§4.3) is the must-have.** For their deployment-project nodes it takes effect as long as their
  deployments carry no PCR device policy (§1.2). That's true today, since Terraform can't add one. Document the order.
- **Their "can't update deployments" is a separate, existing provider gap, not an API one.** `UpdateDeployment`
  is a stub that only returns a warning (`v2/resources/deployment.go:343-351`), and no attribute is `ForceNew`. So
  editing a `zedcloud_deployment` shows a plan, "applies" with a warning, and changes nothing on the controller.
  The controller does support it: `PUT /v2/projects/id/{projectId}/deployments/id/{id}` (gilas `deploymentUpdate` →
  seine `deploymentUpdate`) creates **version N+1** with a new id. It only allows this from the latest version. The provider already
  calls this endpoint (`client.Deployment.CreateNewVersion`), but only from `CreateDeployment`, when a second
  `zedcloud_deployment` resource is declared for a project that already has one (`deployment.go:203-221`).
- Adding a `pcr_policy` block to `zedcloud_deployment` without fixing update gives the customer very little: they
  could set it only by declaring a new deployment resource. **Decision (2026-10-16): the API supports updates, so
  NFR-249 implements in-place deployment update** (§4.5). That makes the deployment PCR policy usable.

### 4.5 `zedcloud_deployment` in-place update (in scope)

**What the API does on update** (verified in code):

| Step | Behaviour | Where |
|---|---|---|
| Route | `PUT /v2/projects/id/{projectId}/deployments/id/{id}`, with the body's `id` matching the URL | gilas `project_handlers.go:72`, `deploymenthandlers.go:606-623` |
| Version check | Rejected (400) unless `{id}` is the **latest** version in the project (`DepVersion == max(depversion)`) | seine `deploymentproc.go:627-680` |
| Result | A **new deployment row**: new UUID, `DepVersion + 1`, name `<name>-v<N+1>`, a job created. Returns `202` with `objectId` = **new id** and `jobId`. The old version row stays as history | seine `deploymentproc.go:681-760` |
| Policies | Applied **asynchronously** (`go updateDeploymentPolicies(...)`) after the 202. It's a full replace: policy types missing from the request are removed. Each changed policy gets a new version (`bumpDeploymentPolicyVersion`) | gilas `deploymenthandlers.go:646-700`, seine `deploymentpolicyproc.go:1008-1054` |
| `deploymentTag` | Required on update (an empty value is rejected) | seine `deploymentproc.go:683-692` |

The provider already calls this endpoint (`client.Deployment.CreateNewVersion`, `v2/client/deployment/client.go:262`),
just from the wrong place. `UpdateDeployment` is a stub that only returns a warning (`v2/resources/deployment.go:343-351`).

**Provider design:**

1. `UpdateDeployment`:
   - `model := DeploymentModel(d)` (full body, so the controller's full-replace semantics match the config), with `model.ID = d.Id()`.
   - `client.Deployment.CreateNewVersion(projectId, d.Id(), model)`.
   - Map a 400 version mismatch to a clear error: *"deployment `<id>` is not the latest version in project `<p>`
     (someone created a newer version outside Terraform). Run `terraform apply -refresh-only` or re-import the latest version."*
   - `d.SetId(resp.ObjectID)`: the resource moves to the new version. SDKv2 allows changing the id in Update.
   - **Wait for the async policy job.** Poll `GET /v1/jobs/id/{jobId}` (in the existing job swagger; add a small hand-written
     client like `v2/client/artifact`) until it succeeds or fails, bounded by `timeouts { update }` (default 5 min). A failed job is an error,
     with the job's per-object status in the message. Without the wait, the read below races the goroutine and
     shows stale policies, which gives a flaky plan.
   - Then `GetDeploymentByID` to refresh state.
2. **Schema:**
   - `project_id` becomes `ForceNew`; a deployment can't move between projects.
   - `deployment_tag`: the API requires it on update, so make it Required. Check the existing fixtures: if create also
     sends it, there's no behaviour change.
   - Remove the "does not support updates" text from the resource `Description`.
   - Add `Timeouts{Update}`.
3. **Name round-trip:** the controller stores `<name>-vN`. The existing `TestDeployment_Create` (`testdata/deployment/create.yaml`) expects the configured
   name back after create, so the read path already returns it without the suffix. **Verify this for v2+** in the update
   acceptance test. If the suffix shows up, add a `DiffSuppressFunc` that strips `-v\d+$`.
4. **Destroy:** deletes the **current** version only (unchanged). Older versions stay as controller history, the same as
   updates made in the UI or zcli.
5. **Existing "second resource = new version" path in `CreateDeployment`** (`deployment.go:173-221`): keep it for
   compatibility, but document in-place update as the recommended way.

Estimate: **2–3 d** (update path, job client and wait, error mapping, acceptance test).

---

## 5. Affected Components

### 5.1 `terraform-provider-zedcloud` (required)

| Change | Files |
|---|---|
| Swagger refresh (app) | `swagger/zedge_app_service.swagger.json` (`make download-specs` against a 21.x controller). Node swagger is already current |
| Models | `v2/models/app.go`, `app_summary.go`, `policy_config.go`, `pcr_policy.go` (new), `tag.go`, `policy_type.go`; `device_policy.go`, `device_p_c_r_policy.go` (new), `device_policy_type.go`. Hand-edit, as with PNAC, rather than regenerating all of v2 |
| Schemas | `v2/schemas/application.go`, `pcr_policy.go` (new), `policy_config.go`, `tag.go`, `node.go` (+ small `node_pcr_policy.go`), `device_policy.go` |
| Resources | `v2/resources/application.go` (tags preserve), `node.go` (pcr preserve/reset in `UpdateNode`, `publishBaseOS`, `applyBaseOS`), `project.go` (`CustomizeDiff`); fix `project.go:224` (`err != nil` should be `errs != nil`) while there |
| Deployment update (§4.5) | `v2/resources/deployment.go` (`UpdateDeployment`, timeouts, description), `v2/schemas/deployment.go` (`project_id` ForceNew, `deployment_tag`), `v2/client/job/` (new, small: get job by id), `v2/client/zedcloud_api.go` |
| Tests | see §7 |
| Docs | `v2/examples/resources/{zedcloud_application,zedcloud_project,zedcloud_edgenode}/resource.tf`, then `make docs` |

Rough size: **tags 1–1.5 d, project PCR 1.5 d, edge node PCR 2 d (ownership logic), deployment PCR 1 d,
deployment in-place update 2–3 d, tests and docs 2–2.5 d → 9–11 days**, including acceptance runs.

### 5.2 `zedcloud` (recommended, not blocking)

See §6.2.

---

## 6. Security and correctness considerations

### 6.1 Provider

- No new credentials or endpoints. App tags need the Edge-app permission, and PCR policies need project/device permission, which are the same permissions the resources already need.
- Tag values are validated client-side with the controller's injection-safe rules (§4.1), so bad input fails at plan time rather than as a 400.
- PCR policies affect whether an encrypted device can unseal its vault. A wrong set (for example, adding a PCR the
  BIOS changes on update) can leave a node unable to unseal after a firmware update. The docs should give the
  recommended default set `[0,1,2,3,4,6,7,8,9,13,14]` and the PCR table, and say that the
  change takes effect on the next seal. Unlike zcli, the provider won't prompt, so put a warning in the resource docs.

### 6.2 Controller follow-ups found during analysis

1. **No way to remove a project PCR policy** (`tagUpdate` only handles a non-nil `pcrPolicy`). Suggest: an empty `pcrPolicy` (or `pcrIndices: []`) on update detaches and deletes it.
2. **Device GET mixes device-level and deployment-derived PCR policy** into one field. Suggest: return the deployment-derived one in a separate read-only field (for example `effectivePcrPolicy`), or don't overwrite `pcrPolicy`.
3. **Device and app PUT treat a missing field as "clear".** That's consistent with the rest of the API but makes every client do read-modify-write. Not worth changing the API for; the provider handles it (§4.1, §4.3).
4. **`policyVersion` for project/device policies is client-managed** (seine stores it as-is; zcli bumps it for devices, never for projects). The systest doc says the server assigns it. Fix the doc, or move the bump into seine so all clients behave the same. It's sent to EVE as `PolicyID`.
5. **Systest doc says indices are sorted.** Only zcli sorts; `validatePCRIndices` doesn't. Fix the doc, or sort in seine.
6. **zcli** (not the controller): `project update --policy-config=pcr:` can't attach a PCR policy to a project that has none, and it never bumps `policyVersion` (§1.3). `pcr-policy show` labels a deployment-derived policy as "device-level".

---

## 7. Testing Strategy

- **Unit (no network):**
  - `tags` validator: limits, forbidden characters (quotes, backslash, `<>%&?/`), 33 entries rejected.
  - PCR element validator and CustomizeDiff: `[1,2]` (no 0), `[0,5]`, `[0,16]`, `[]` all rejected; `[0]` accepted.
  - Expand/flatten round-trip for `Policy.PcrPolicy`, `Tag.PcrPolicy`, `Node.PcrPolicy`, `Application.Tags`.
  - Node update body builder: configured, removed, never configured, and never configured + deployment-shaped existing policy.
  - `policy_version`: new policy → 1; indices changed → fetched + 1 (including after a zcli bump in between); indices unchanged in a different order → no bump; payload indices sorted (matches zcli).
- **Acceptance (`make test-run case=...`, `__SUFFIX__` fixtures):**
  - `TestApplication_Tags`: create with tags, update a value, set `{}` and verify they're cleared. Then set tags out-of-band via API, apply an unrelated change, and **check the tags survive**.
  - `TestProject_PCRPolicy`: create with `pcr_policy`, update the indices, plan shows no diff when the indices are reordered.
  - `TestNode_PCRPolicy`: set, update, remove (GET has no `pcrPolicy`). Set out-of-band, apply an unrelated node change, and check the policy survives. Copy `testdata/node/*_with_pnac.*`.
  - `TestDeployment_PCRDevicePolicy`, with `policy_version` going 1 → 2 on update.
  - `TestDeployment_Update`: create, then change `title`, an app-instance policy and the PCR device policy in place. Expect a new `id`, an
    incremented `revision`, the configured `name` unchanged in state, no diff on the next plan, and the old version id still readable. Then remove a
    policy type from config and check the controller drops it (full replace).
  - `TestDeployment_UpdateNotLatest`: create a newer version outside Terraform (API PUT), then apply a change and expect the
    "not the latest version" error.
  - Unit: job-wait helper (success, failure with per-object status, timeout) against `httptest.Server`.
- **Manual:** zcli `edge-node pcr-policy show` after a Terraform apply. The customer's portal filters apps by the Terraform-set tags.

---

## 8. Open Questions

**Ground rule (2026-10-16):** the provider follows the **current** zedcloud API. If the API doesn't support
something, Terraform doesn't implement it or work around it. Controller gaps stay in §6.2 as follow-ups and
don't block NFR-249.

1. ~~**Who owns `policyVersion` for project and device policies?**~~ **Mostly answered by zcli (§1.3):**
   the client owns it. For edge nodes the convention is `previous + 1` from the live GET. For projects there is no convention (it stays at 1).
   Decision: the provider manages it as Computed and bumps it only when the indices change (§4.2, §4.3).
   **EVE side answered (NFR-249, 2026-10-16):** `PolicyID` is mostly bookkeeping. EVE treats a change to
   **either** the PCR indices **or** `PolicyID` as a new policy and applies it. So projects that zcli updated
   without a bump *did* get their new indices applied. `PolicyID` only needs to increase when the policy changes.
   This confirms the provider rule: bump only when `pcr_indices` change, and never on an apply with no changes. A
   bump alone would make EVE re-apply the same PCR set. **Closed.**
2. ~~**Deployment vs device PCR on a GET**~~ **Closed by the ground rule.** The API returns one `pcrPolicy` with
   no source marker, so the provider takes it as-is. Preserve as fetched when `pcr_policy` isn't in config (§4.3),
   and document the deployment-project behaviour as a known limitation. No extra GETs and no heuristics. A
   separate effective-policy field stays a controller follow-up (§6.2 item 2).
3. ~~**Minimum controller version** for tags~~ **Answered (zedcloud tag analysis, 2026-10-16):** zedcloud **20.0.7**
   (cherry-pick `b96e857c9b`) on the 20.0 line, and 21.0.0+. Goes in the `zedcloud_application` docs and the provider changelog (§1.1).
4. ~~**Tag-filtered list data source**~~ **Answered (customer, 2026-10-14):** only *setting* tags is needed. Filtering "could be useful"
   but has no use case yet. **Deferred**; not part of NFR-249.
5. **Deployments:** The customer needs edge-node level (must). They use deployment projects but can't update them from Terraform (§4.4).
   **Closed (2026-10-16):** the API supports updates (`PUT /v2/projects/id/{projectId}/deployments/id/{id}` creates version N+1), so
   per the ground rule **NFR-249 implements in-place `zedcloud_deployment` update** (§4.5).
6. ~~**Project PCR removal**~~ **Closed by the ground rule.** The API can't remove a project PCR policy, so
   Terraform doesn't offer it. `zedcloud_project.pcr_policy` is Optional+Computed, so removing the block is a no-op
   and shows no diff (§4.2). The resource docs say so explicitly. (Sergey's `zcli edge-node pcr-policy reset` is the
   **edge-node** reset, which the API does support and the design implements, §4.3.) Project removal stays a
   controller follow-up (§6.2 item 1).
7. ~~**Systest doc vs code** (`policyVersion` assignment, index sorting)~~ **Closed by the ground rule.** The provider
   follows the code: the client sets `policyVersion` (§4.2, §4.3), and the provider sorts the indices before sending. Fixing
   the doc stays a controller follow-up (§6.2 items 4–5).

---

## 9. Implementation notes (2026-10-16)

This section lists what was built and where it differs from §2–§7.

**Files**

| Area | Files |
|---|---|
| Models | `v2/models/app.go`, `app_summary.go` (`Tags`); `p_c_r_policy.go`, `device_p_c_r_policy.go` (new); `policy_config.go`, `tag.go`, `device_policy.go` (`PcrPolicy`); `policy_type.go` (+`POLICY_TYPE_PCR`, `_LOC`, `_CONFIGURATION_LOCK`), `device_policy_type.go` (+`DEVICE_POLICY_TYPE_PCR`) |
| Schemas | `v2/schemas/app_tags.go` (validator), `pcr_policy.go` (PCR schemas, validation, version rule, node wrapping), `deployment_diff.go` (diff suppression), plus wiring in `application.go`, `policy_config.go`, `tag.go`, `node.go`, `device_policy.go`, `deployment.go` |
| Resources | `v2/resources/managed.go` (managed-only helpers), `application.go`, `project.go`, `node.go`, `deployment.go` |
| Client | `v2/client/job/` (new, hand-written: `Get`, `Wait`), wired into `v2/client/zedcloud_api.go` |
| Tests | Unit: `schemas/app_tags_test.go`, `schemas/pcr_policy_test.go`, `resources/pcr_policy_test.go`, `client/job/client_test.go`. Acceptance: `resources/nfr249_test.go` with fixtures in `resources/testdata/nfr249/` |
| Docs | `v2/examples/resources/{zedcloud_application,zedcloud_project,zedcloud_edgenode,zedcloud_deployment}/resource.tf`, regenerated `docs/resources/` and `docs/data-sources/` for those four (`make docs`) |

**Differences from the design**

- **`zedcloud_application.tags` uses "managed once configured" semantics, not Optional+Computed (§4.1).** The attribute is
  Optional. Read writes it to state only when it is already in state (it came from config). Update sends the app's current
  tags when the attribute is in neither the old nor the new state. Effect: UI-set tags cause no diff and survive applies;
  removing the argument after it was set clears the tags (the design had "set `tags = {}`", which Optional+Computed in
  SDKv2 cannot express). The data source still returns all tags. Same mechanism as the edge node PCR policy.
- **SDKv2 detail:** clearing an attribute in Read must use an empty value (`map[string]interface{}{}`, `[]interface{}{}`),
  not `nil`: the SDK's map writer keeps the previously written keys for `nil`. Found by the acceptance test.
- **`policy_version` (project, node) is Computed only.** It is set at apply time. The plan does not show "(known after
  apply)": `SetNewComputed` doesn't work on nested keys, and the legacy SDK tolerates the value changing during apply.
- **Deployment PCR heuristic:** not implemented, as decided in §8 Q2. The node keeps whatever the GET returns.
- **Deployments need a new `deployment_tag` for every version.** Not in the design. `user_config.deployment_data` has
  `unique (enterpriseid, deploymenttag, project_id)`, so a new version with an unchanged tag fails with a bare `409`.
  A `CustomizeDiff` fails the plan when an existing deployment changes without a new `deployment_tag`, and explains that
  nodes move to the new version by changing their tag (which is the controller's rollout model). `deployment_tag` stays
  Optional instead of becoming Required (§4.5): a missing tag only matters for updates, and the plan check covers that.
- **The "not the latest version" check is done by the provider before the PUT.** The controller's 400 for it has
  `"details": null`, so its message cannot be matched. The provider lists the project's versions first.
- **Destroy deletes the older versions too.** Not in the design. While any version exists, the project cannot be
  deleted (`409`, "One or more Deployments are still attached to Project"), so destroy removes the current version and
  every older one. Newer versions (created outside Terraform) are left alone.
- **Perpetual diffs on `zedcloud_deployment` are fixed.** They were harmless while update was a no-op, but would now
  create a new version on every apply. Two causes, both suppressed in `schemas/deployment_diff.go`:
  empty maps in nested blocks (`+ tags = {}`, `+ port_tags = {}`, `+ license_list = {}`, ...) and device policy names the
  controller decorates (`<name>-attestation`, `<name>-pcr`, sometimes `<name>-pcr-pcr`). The provider also strips the
  decoration before sending names on update, so they don't grow. `TestDeployment_Create` and `TestProject_PNAC_V2_Create`
  used to expect a non-empty plan after apply; they now require an empty one.
- **Waiting for the update job** is skipped when the deployment has no policies: gilas then sends no policy message and
  the job would never complete.
- **"Full replace" of policy types (§4.5) is not what the controller does.** gilas only sends policy types present in the
  request; the provider always sends the whole config, so the behaviour is the same for Terraform users.
- **`application_test.go`:** the golden comparison now uses `cmpopts.EquateEmpty()`, because controllers with tags
  return `tags: {}`.

**Verified on the local controller (2026-10-16)**

| Test | Result |
|---|---|
| `TestApplication_Tags`: create with tags, change, remove (cleared), UI-set tags survive an unrelated apply with an empty plan, invalid key rejected at plan | PASS |
| `TestProject_PCRPolicy`: create (version 1), reorder (no diff), change (version 2), PCR 0 missing / PCR 5 rejected at plan | PASS |
| `TestNode_PCRPolicy`: zcli-style policy set out of band survives an unrelated apply; Terraform takes over at zcli's version + 1; reorder keeps the version; removing the block resets the node | PASS |
| `TestDeployment_UpdateInPlace`: create; change without a new tag rejected at plan; two in-place updates (new id each, clean plan); out-of-band newer version gives the "not the latest version" error; destroy removes all versions and the project | PASS |
| Regression: `TestApplication_*`, `TestArtifact_*`, `TestProject_*`, `TestNode_*` (incl. `TestNode_PNAC_Onboard`), `TestDeployment_Create` | PASS |
| Unit: validators, version rule, node update decision, job client (complete, fail with decoded per-object status, timeout), schema round-trips, `TestProvider_InternalValidate`, fixture token checks | PASS |

**Verified end to end on a virtual EVE node with a vTPM (Berlin lab + alpha, 2026-10-19)**

Node `tf_e2e_en_nfr249`: QEMU/KVM on `h-m-dl20`, EVE `16.0.1-lts-kvm-amd64`, a software TPM (`zedamigo_swtpm`, with
`swtpm` taken from `nix shell nixpkgs#swtpm`, since it is not installed on the host), onboarded to alpha. The config is
`test/e2e` plus a `zedamigo_swtpm` resource (wired into `swtpm_socket` of the install and run VMs), and two variables
that set `pcr_policy` on the project and on the node. Terraform ran on the lab host with this branch's provider
(cross-compiled, `dev_overrides`). After each change the node was rebooted through the API so it attests again. On the
device, the policy EVE received is in `/run/zedagent/EncryptedVaultKeyFromController/*.json` (`PolicyPcr`), published
by zedagent after a successful attestation (`attesttask.go`).

| Step (Terraform change) | Controller | EVE after reboot (`PolicyPcr`) | Node |
|---|---|---|---|
| Baseline, no policies | — | `PolicyPresent: false` (first attestation; no escrowed key yet) | ONLINE, vault unlocked |
| Project `pcr_policy = [0,1,2,3,4,6,7,8,9,13,14]` | project policy v1 | `[0,1,2,3,4,6,7,8,9,13,14]`, `ID: 1` | ONLINE, vault unlocked |
| Node `pcr_policy = [0,1,2,3,4,6,7]` | device policy v1 | `[0,1,2,3,4,6,7]`, `ID: 1` (node overrides project) | ONLINE, vault unlocked |
| Node `pcr_policy = [0,7]` | device policy v2 | `[0,7]`, `ID: 2` | ONLINE, vault unlocked |
| Node block removed | device policy cleared | `[0,1,2,3,4,6,7,8,9,13,14]`, `ID: 1` (back to the project policy) | ONLINE, vault unlocked |
| zcli-style out-of-band set (`[0,4,7]`, v3), then an apply that updates the node without a `pcr_policy` block | device policy still `[0,4,7]` v3 | — | — |

Against alpha, `TestApplication_Tags`, `TestProject_PCRPolicy`, `TestNode_PCRPolicy`, `TestDeployment_UpdateInPlace`
and `TestDeployment_Create` also pass.

`TestApplicationInstance_RealNode` (an nginx app instance on the node) could not be judged: on alpha, every
`GET /v1/apps/instances/id/{id}/status` returned `500` during the run, for all instances in the enterprise, not only the
test's, so the wait for `SW_STATE_RUNNING` timed out. This is an alpha problem, unrelated to this change (the app
instance resource is untouched).

Found on the way, not caused by this change: `test/e2e` has a perpetual diff on `zedcloud_edgenode.EN`
(`interfaces.net_dhcp`, `NETWORK_DHCP_TYPE_CLIENT` → `NETWORK_DHCP_TYPE_UNSPECIFIED`), so every apply updates the node.
That happened to be useful here (it exercised the "keep the zcli policy" path on each apply). **Fixed** in the same
branch: `test/e2e` now sets `net_dhcp` from the network, and gained a `vtpm` option (a `zedamigo_swtpm` wired into both
VMs). A fresh `test/e2e` run with `vtpm = true` (run_id `pr1`, 2026-10-19) onboarded in 1m21s with a TPM on EVE, and the
post-apply plan was empty. And the lab host's zedamigo mirror gained `LICENSE`/`README.md` on 2026-09-17, which breaks the
recorded lock file hash for existing `test/e2e` working directories (`tofu init` fails until the zedamigo entry is
removed from `.terraform.lock.hcl`).
- **Controller findings to add to §6.2:** the `deploymenttag` uniqueness (and the bare 409 for it); the bare 400 for "not
  the latest version"; device policy name decoration growing on update (`-pcr-pcr`); deployment policy names are unique
  per enterprise, so a policy left behind by a failed run makes later deployments with the same policy name fail, and the
  controller then deletes the new deployment (the create returns 202 and the read 404s).

