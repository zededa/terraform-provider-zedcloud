# Design: NFR-249 — Edge app tags and PCR policies via Terraform

**Ticket:** [NFR-249](https://zededa.atlassian.net/browse/NFR-249) — *Terraform Provider support for marketplace app tags and PCR policies*
**Customer:** Speedcast (via ZD-3895). Follow-up to the two controller features built for them:
[NFR-96](https://zededa.atlassian.net/browse/NFR-96) (tags on marketplace Edge apps, API epic
[ZEDCLOUD-3310](https://zededa.atlassian.net/browse/ZEDCLOUD-3310)) and
[NFR-189](https://zededa.atlassian.net/browse/NFR-189) (flexible PCR policy, epic
[ZEDCLOUD-2157](https://zededa.atlassian.net/browse/ZEDCLOUD-2157)).
**Repos:** `zededa/terraform-provider-zedcloud` @ `e5dddd7e` (origin/main), `zededa/zedcloud` @ `6197300bc2` (main)
**Date:** 2026-10-08
**Author:** Ivan Curachin
**Status:** Draft

**Headline: no zedcloud work is needed to unblock this. The provider work is three schema
additions on existing resources, with no new resources or clients.** Both features are already live in the
controller API. The provider's swagger and models are partly stale, and its schemas don't expose
either feature. Two findings matter more than the feature gap itself:

1. **Terraform already deletes this data today.** `zedcloud_application` and `zedcloud_edgenode`
   updates send the whole object. The controller treats a missing `tags` / `pcrPolicy` as "clear
   it". So any `terraform apply` that touches an app wipes tags set in the UI, and any apply that
   touches an edge node wipes a device PCR policy set with zcli (§1.3).
2. **On a device GET, `pcrPolicy` can come from a deployment, not the device**, and the response doesn't say which. The provider has to handle this or it will copy deployment policy into device config (§4.3).

---

## 1. Problem Statement

Speedcast manages tenants with Terraform and wants to:

- set **tags on Edge applications** (marketplace apps), so their portal can filter apps by tag;
- set **PCR policies** (which TPM PCRs are used to seal the volume key) on **projects** and on
  **individual edge nodes**.

Both are possible through the UI/API/zcli, but not through `terraform-provider-zedcloud`.

### 1.1 What was done in zedcloud: Edge app tags (NFR-96)

Commit `b64a96fdaa` ([zedcloud#8255](https://github.com/zededa/zedcloud/pull/8255), 2026-07-06), first in tag **21.0.0**.
Per ZEDCLOUD-3310 the feature is available in old UI 20.0.3 and Cortex UI 1.1.4.
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
| `policyVersion` for **project/device** policies: no server-side increment found in seine; whatever the client sends is stored and sent to EVE | — | The systest doc says "server ignores user-supplied policyVersion". The code doesn't show that. **Open question §8 Q1** |
| Index order: the systest doc says the server sorts. `validatePCRIndices` doesn't sort | — | Treat the indices as a **set** in Terraform |

### 1.3 Provider ground truth

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

No new resources. Add attributes to three existing resources (four with the optional deployment
part), and fix the update paths so they stop deleting data Terraform doesn't manage.

```hcl
resource "zedcloud_application" "app" {
  name  = "speedcast-router"
  title = "Speedcast router"
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
| **E. `zedcloud_applications` data source filtering by tags** (`tags` + `match_any_tag`) | Not asked for (the customer wants to *set* tags), and it would need a list data source the provider doesn't have yet. **Deferred** (§8 Q4). |

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
- **Older controllers (< 21.0 / 20.0.3):** `omitempty` means nothing is sent unless the user sets
  `tags`. If they set it on an old controller, gilas may reject the unknown field or silently drop it. Document the
  minimum controller version (§8 Q3).

### 4.2 `zedcloud_project.pcr_policy`

The same pattern as `attestation_policy` / `port_based_network_access_control_policy`:

```text
pcr_policy            TypeList(MaxItems 1) of Policy()      Optional, Computed
  └─ pcr_policy       TypeList(MaxItems 1) of PCRPolicy()   (new nested block inside Policy())
       ├─ pcr_indices     TypeSet(int)  Required   elem: 0–15, != 5
       ├─ policy_version  int           Optional, Computed   (see §8 Q1)
       └─ id              string        Computed
```

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
  └─ policy_version   int                    Optional, Computed   (see §8 Q1)
```

Wire mapping: `Node.PcrPolicy = &Policy{Type: POLICY_TYPE_PCR, PcrPolicy: &PCRPolicy{PcrIndices, PolicyVersion}}`.
Leave `name`/`title` unset (the device path doesn't use them). **Check:** the `Policy` model marks
`name`/`title`/`type` as `Required` and so serialises `"name": null`. Confirm gilas accepts this,
or add `omitempty` handling (TC-04 in the systest doc sends only `type` and `pcrPolicy`, so it should).

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
In the provider, skip the preserve path when the fetched policy has no `id` and no `name` (the deployment-derived shape).
This is a heuristic until the controller exposes the source.

Also apply the preserve logic in `publishBaseOS` / `applyBaseOS`, which send `newNode` too (`node.go:522-564`).

### 4.4 `zedcloud_deployment` device policy (optional phase)

Not in the NFR text ("edge nodes and projects"), but it's the third API surface, and the
PNAC precedent already added `port_based_network_access_control_policy` here. Small:

- `models.DevicePCRPolicy` (new), `models.DevicePolicy.PcrPolicy`, `DEVICE_POLICY_TYPE_PCR` in `DevicePolicyType`.
- `v2/schemas/device_policy.go`: a `pcr_policy` block (`pcr_indices` set, `policy_version` Computed:
  server-managed here).
- Validation as above. gilas also rejects bad indices synchronously with a 400 (`deploymenthandlers.go`).

Recommendation: include it if Speedcast uses deployments (`TAG_TYPE_DEPLOYMENT` projects). Ask them (§8 Q5).

---

## 5. Affected Components

### 5.1 `terraform-provider-zedcloud` (required)

| Change | Files |
|---|---|
| Swagger refresh (app) | `swagger/zedge_app_service.swagger.json` (`make download-specs` against a 21.x controller). Node swagger is already current |
| Models | `v2/models/app.go`, `app_summary.go`, `policy_config.go`, `pcr_policy.go` (new), `tag.go`, `policy_type.go`; optional `device_policy.go`, `device_p_c_r_policy.go` (new), `device_policy_type.go`. Hand-edit, as with PNAC, rather than regenerating all of v2 |
| Schemas | `v2/schemas/application.go`, `pcr_policy.go` (new), `policy_config.go`, `tag.go`, `node.go` (+ small `node_pcr_policy.go`); optional `device_policy.go` |
| Resources | `v2/resources/application.go` (tags preserve), `node.go` (pcr preserve/reset in `UpdateNode`, `publishBaseOS`, `applyBaseOS`), `project.go` (`CustomizeDiff`); fix `project.go:224` (`err != nil` should be `errs != nil`) while there |
| Tests | see §7 |
| Docs | `v2/examples/resources/{zedcloud_application,zedcloud_project,zedcloud_edgenode}/resource.tf`, then `make docs` |

Rough size: **tags 1–1.5 d, project PCR 1.5 d, edge node PCR 2 d (ownership logic), deployment PCR 1 d
(optional), tests and docs 1.5–2 d → 6–8 days**, including acceptance runs.

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
4. **`policyVersion` ownership for project/device policies is unclear.** The doc says server-assigned, the code says client-supplied (§8 Q1). It's sent to EVE as `PolicyID`, so it matters.
5. **Systest doc says indices are sorted.** They aren't in `validatePCRIndices`. Fix either the doc or the code.

---

## 7. Testing Strategy

- **Unit (no network):**
  - `tags` validator: limits, forbidden characters (quotes, backslash, `<>%&?/`), 33 entries rejected.
  - PCR element validator and CustomizeDiff: `[1,2]` (no 0), `[0,5]`, `[0,16]`, `[]` all rejected; `[0]` accepted.
  - Expand/flatten round-trip for `Policy.PcrPolicy`, `Tag.PcrPolicy`, `Node.PcrPolicy`, `Application.Tags`.
  - Node update body builder: configured, removed, never configured, and never configured + deployment-shaped existing policy.
- **Acceptance (`make test-run case=...`, `__SUFFIX__` fixtures):**
  - `TestApplication_Tags`: create with tags, update a value, set `{}` and verify they're cleared. Then set tags out-of-band via API, apply an unrelated change, and **check the tags survive**.
  - `TestProject_PCRPolicy`: create with `pcr_policy`, update the indices, plan shows no diff when the indices are reordered.
  - `TestNode_PCRPolicy`: set, update, remove (GET has no `pcrPolicy`). Set out-of-band, apply an unrelated node change, and check the policy survives. Copy `testdata/node/*_with_pnac.*`.
  - Optional `TestDeployment_PCRDevicePolicy`, with `policy_version` going 1 → 2 on update.
- **Manual:** zcli `edge-node pcr-policy show` after a Terraform apply. The Speedcast portal filters apps by the Terraform-set tags.

---

## 8. Open Questions

1. **Who owns `policyVersion` for project and device policies?** If the client owns it (as the code
   suggests), should the provider bump it when `pcr_indices` change (like zcli presumably does)
   and leave it Optional+Computed for overrides? Does EVE use `PolicyID` to detect a change, and does
   it need to be monotonic? *Owner: zedcloud (rishabh) / EVE security.* Proposed default: the provider sets
   `prior + 1` when the indices change and the user didn't set it.
2. **Deployment vs device PCR on a GET** (§4.3). Is the "no id/name" heuristic acceptable until
   §6.2 item 2 lands?
3. **Minimum controller version** for tags: 21.0.0, or the 20.0.3 cherry-pick? This goes in the docs and the
   provider changelog.
4. **Tag-filtered list data source** (`zedcloud_applications` with `tags`, `match_any_tag`): does Speedcast
   need it in Terraform, or only in their portal through the API? Deferred unless they ask.
5. **Deployments:** does Speedcast use deployment projects with PCR device policies? If so, include §4.4.
6. **Project PCR removal:** OK to ship with "can't remove" documented, or block on §6.2 item 1?
