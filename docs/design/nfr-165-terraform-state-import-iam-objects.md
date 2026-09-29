# NFR-165 — Terraform state import for IAM objects (users, custom roles, auth profiles)

**Ticket:** [UE-166](https://zededa.atlassian.net/browse/UE-166) — implementation
**Driven by:** [NFR-165](https://zededa.atlassian.net/browse/NFR-165) — *Support Terraform state import for user accounts and custom roles*
**Follow-ups:** [UE-167](https://zededa.atlassian.net/browse/UE-167) (regenerate stale models, §12.1/§12.2),
[UE-168](https://zededa.atlassian.net/browse/UE-168) (auth_profile import, §7/§12.6)
**Repo:** `zededa/terraform-provider-zedcloud`, checked against branch `CI-836-cluster-eve-os-upgrade` @ `8b633ead`
**Date:** 2026-09-28
**Author:** Ivan Curachin
**Status:** Implemented on `nfr-165-iam-state-import`, except §3.5 and §7 — see §12 for
what implementation changed about this design

**Goal.** Make `terraform import zedcloud_user.alice …` and `terraform import zedcloud_role.readonly …`
work, and work *cleanly* — an import followed by `terraform plan` must produce an empty diff.
Customer (Speedcast) currently has to delete and recreate accounts to bring them under
Terraform management.

**Headline: this is a provider-only change. No zedcloud API work is required.** Every read
endpoint the import path needs already exists in `zedge_user_service.proto` and is already
wired into the generated client. The reason import fails today is one missing struct field.

---

## 1. Verified ground truth

Measured against the tree, not taken from the ticket.

| Thing | Actually |
|---|---|
| Resources registered in `ResourcesMap` | **34** |
| Resources that set `Importer:` | **28** |
| Resources missing `Importer:` | **6** — `zedcloud_user`, `zedcloud_role`, `zedcloud_auth_profile`, `zedcloud_credential`, `zedcloud_cluster_group_manifest`, `zedcloud_kubernetes_cluster_status` |
| Of those 6, real importable objects | **3** — user, role, auth_profile. The other 3 are no-op/write-only pseudo-resources (§6) |
| Import pattern used by the other 28 | `Importer: &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext}` |
| Acceptance tests exercising import | **1** — `resources/cep_profile_test.go:61` |
| Generated resource docs under `v2/docs/resources/` | **0** (directory is empty) |

Reproduce:

```bash
cd v2
sed -n '/ResourcesMap/,/^\t\t},/p' resources/provider.go | grep -c '"zedcloud_'
grep -l 'Importer:' resources/*.go | grep -v _test | wc -l
grep -rn 'ImportState' resources/*_test.go
ls docs/resources/
```

### 1.1 The API endpoints already exist

| Need | HTTP | Generated client method | Status |
|---|---|---|---|
| Read user by ID | `GET /v1/users/id/{id}` | `IdentityAccessManagementGetUser` | ✅ |
| Read user by name | `GET /v1/users/name/{name}` | `IdentityAccessManagementGetUserByName` | ✅ |
| Read role by ID | `GET /v1/roles/id/{id}` | `IdentityAccessManagementGetRole` | ✅ |
| Read role by name | `GET /v1/roles/name/{name}` | `IdentityAccessManagementGetRoleByName` | ✅ |
| Read auth profile by ID | `GET /v1/authorization/profiles/id/{id}` | `IdentityAccessManagementGetAuthProfile` | ✅ |
| Read auth profile by name | `GET /v1/authorization/profiles/name/{name}` | `IdentityAccessManagementGetAuthProfileByName` | ✅ |
| List | `GET /v1/users`, `/v1/roles`, `/v1/authorization/profiles` | `Query*` | ✅ |

The proto declares a `404` response on `GetUser` / `GetRole`, so "object was deleted out of
band" is expressible server-side — the provider just does not act on it (§3.3).

Reproduce:

```bash
python3 - <<'EOF'
import json
d = json.load(open('swagger/zedge_user_service.swagger.json'))
for p, ops in d['paths'].items():
    for m, o in ops.items():
        if 'user' in p or 'role' in p:
            print(m.upper(), p, '|', o.get('operationId'))
EOF
ls v2/client/identity_access_management/ | grep -E 'get_(user|role)(_by_name)?_parameters'
```

### 1.2 Why import fails today — exactly

`UserResource()` (`resources/user.go:21`) and `RoleResource()` (`resources/role.go:21`) return a
`*schema.Resource` with `CreateContext` / `ReadContext` / `UpdateContext` / `DeleteContext` /
`Schema` set, and no `Importer`. The SDK treats a resource without an `Importer` as
non-importable and Terraform errors with *"resource zedcloud_user doesn't support import"*
before any provider code runs.

The passthrough importer is sufficient here, because the plumbing already lines up:

1. `ImportStatePassthroughContext` returns the `ResourceData` unchanged, with only the ID set.
2. `GRPCProviderServer.ImportResourceState` then does `is.Attributes["id"] = is.ID`
   (`vendor/github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema/grpc_provider.go:1096`),
   so the imported state carries the `id` **attribute**, not just the state ID.
3. Terraform core calls `ReadResource` against that state. `DetailedUserSchema()` and
   `RoleSchema()` both declare `"id"` as a `Computed` `TypeString`, so `d.GetOk("id")`
   resolves.
4. `IdentityAccessManagement_GetUser` / `_GetRole` already branch to the read-by-ID path when
   `username` / `name` is unset and `id` is set.

So the mechanism works. What does *not* work cleanly is the state that comes back — §3.

---

## 2. Proposed solution

Three layers, in dependency order:

1. **Make the resources importable** — add `Importer` to user, role and auth_profile. Use a
   custom `StateContext` rather than bare passthrough so the import ID may be either the
   28-character system ID or the human-readable name (§4.1).
2. **Harden the read path** — nil-guard, prefer `d.Id()` over the name lookup, and treat 404
   as "gone" (§3.1–3.3). Without these, import appears to work and then misbehaves on the
   next `plan`.
3. **Fix round-trip fidelity** — three schema/marshalling bugs that make an imported user
   produce a diff that never converges (§3.4–3.6).

Layer 1 alone satisfies the letter of the ticket. Layers 2 and 3 are what make it usable;
shipping 1 without 3 would hand Speedcast a permanent-diff bug instead of a missing feature.

---

## 3. Defects found, and what each one costs

### 3.1 Nil dereference when neither lookup key is set

`resources/user.go:38-55`, `resources/role.go:38-55`, `resources/auth_profile.go:37-53` all
follow:

```go
if _, isSet := d.GetOk("username"); isSet {
    user, diags = getUserByName(ctx, d, m)
} else if _, isSet := d.GetOk("id"); isSet {
    user, diags = getUserById(ctx, d, m)
}
if diags.HasError() { return diags }
d.SetId(user.ID)     // user is nil if neither branch ran
```

**Cost:** provider panic (crash, not a diagnostic) rather than an error message. Reachable
today via a data source with neither argument; reachable on import if the ID plumbing ever
regresses. Fix: `else { return diag.Errorf(...) }`.

### 3.2 Name lookup is preferred over ID

After a successful import, state holds *both* `id` and `username`. Every subsequent refresh
therefore takes the **by-name** branch. Rename the user in config and Read queries the *new*,
non-existent name — the provider reports a lookup error instead of detecting drift on the
existing object.

**Cost:** rename is impossible to plan; the practical workaround is `terraform state rm` +
re-import, which is exactly the delete-and-recreate cycle this ticket exists to remove.

Fix: prefer `d.Id()` whenever it is non-empty; fall back to name lookup only for the
data-source / first-read case.

### 3.3 No 404 → `d.SetId("")`

`ZsrvResponderToDiags` (`resources/util.go`) turns every non-2xx into an error diagnostic. A
user deleted in the UI makes `terraform plan` fail instead of planning a re-create.

**Cost:** breaks the normal "reconcile after out-of-band change" workflow that import users
specifically rely on. Note `40ec46a5` already did the equivalent for *delete* (CI-723); this
is the read-side counterpart.

Fix: detect the 404 responder type, `d.SetId("")`, return no error.

### 3.4 `custom_user_input` is never sent — wrong map key

`schemas/detailed_user.go:33`:

```go
customUserInputInterface, customUserInputIsSet := d.GetOk("customUserInput")
```

The schema key is `custom_user_input`. `d.GetOk("customUserInput")` always misses, so
`DetailedUser.CustomUserInput` is always sent empty on create and update.

**Cost:** *this is the one that bites hardest on import.* Read populates
`custom_user_input` from the API; the config declares it; Update silently drops it; the next
plan shows the same diff again. Permanent non-convergence.

(`DetailedUserModelFromMap` at line 116 uses the correct key — only the `ResourceData` path is
wrong.)

### 3.5 `SetDetailedUserResourceData` drops three fields — DEFERRED, see §12.1

`DetailedUser` in the swagger carries `allowedEntpProjectList`, `zksSyncStatus` and
`tokenZksSyncStatus`. None appear in `DetailedUserSchema()` and none are set by
`SetDetailedUserResourceData`.

**Cost:** `allowedEntpProjectList` is per-enterprise-project role assignment — real
configuration. Importing a user who has it produces state that silently omits it, and a
subsequent apply wipes it. The two `zksSyncStatus` fields are read-only status and should be
`Computed`-only, or left out deliberately with a comment.

### 3.6 `last_login_time` / `last_logout_time` are mistyped and mis-classified

Schema declares `TypeString` + `Optional` + `ValidateFunc: validation.IsRFC3339Time` +
`DiffSuppressFunc: supress()`. The model builder does:

```go
lastLoginTime, _ := d.Get("last_login_time").(strfmt.DateTime)   // always fails
```

`d.Get` on a `TypeString` returns `string`; the assertion to `strfmt.DateTime` never succeeds,
so the zero value is always sent. The setter writes `m.LastLoginTime.String()`, which for the
zero `strfmt.DateTime` yields `1970-01-01T00:00:00.000Z`.

**Cost:** low today only because `supress()` unconditionally returns `true`, hiding the diff.
These are server-observed timestamps: they should be `Computed`, dropped from the model
builder, and the blanket diff suppression removed.

### 3.7 No `ForceNew` on immutable identifiers

`zedcloud_role.name` is documented in the API as *"Name cannot be changed once created"*, and
`zedcloud_user.username` is effectively immutable. Neither schema sets `ForceNew`.

**Cost:** after import, any config/remote mismatch on those fields plans an in-place update
that the API rejects at apply time, with a server error rather than a plan-time replacement.

---

## 4. Interface / API design

No proto, REST or schema-breaking changes. Two user-visible surfaces:

### 4.1 Import ID accepts ID *or* name

Administrators know `alice@corp.com`, not `AAGFABAEqnH4je5PHZTXSmHOs-XC`. Both `GetUserByName`
and `GetRoleByName` already exist, and the system ID is regex-separable from a username:

- user/role/profile ID pattern: `^[0-9A-Za-z_=-]{28}$`
- username pattern: `[a-zA-Z0-9][a-zA-Z0-9_.-]+`, min length 3, and in practice an email
- role name pattern: `[a-zA-Z0-9][a-zA-Z0-9_.-]+`, min length 3

A 28-char token matching the ID charset is treated as an ID; anything else is looked up by
name and the resolved ID is written into state.

```go
var objectIDRe = regexp.MustCompile(`^[0-9A-Za-z_=-]{28}$`)

func importUserByIDOrName(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
    raw := d.Id()
    if objectIDRe.MatchString(raw) {
        d.SetId(raw)
        return []*schema.ResourceData{d}, nil
    }
    // resolve via GET /v1/users/name/{name}, then d.SetId(resolved.ID)
}
```

Ambiguity is theoretically possible (a 28-character username in the ID charset) but excluded
in practice: usernames are email addresses and contain `@`. The doc will state the ID-wins
rule explicitly.

```
terraform import zedcloud_user.alice  alice@corp.com
terraform import zedcloud_user.alice  AAGFABAEqnH4je5PHZTXSmHOs-XC
terraform import zedcloud_role.ro     readonly-operators
terraform import zedcloud_auth_profile.okta  corp-okta
```

### 4.2 Schema changes (additive / classification only)

| Resource | Attribute | From | To |
|---|---|---|---|
| `zedcloud_user` | `last_login_time` | `Optional` + `supress()` | `Computed` |
| `zedcloud_user` | `last_logout_time` | `Optional` + `supress()` | `Computed` |
| `zedcloud_user` | `allowed_entp_project_list` | *absent* | new, `Optional` |
| `zedcloud_user` | `zks_sync_status`, `token_zks_sync_status` | *absent* | new, `Computed` |
| `zedcloud_user` | `username` | `Required` | `Required` + `ForceNew` |
| `zedcloud_role` | `name` | `Required` | `Required` + `ForceNew` |

`Optional` → `Computed` on the two timestamps is technically a breaking change for anyone who
set them in HCL. Nobody sensibly does — they are server-observed and the current diff
suppression means any configured value is inert. Call it out in the release notes.

---

## 5. Affected components

Provider only. No `zededa/zedcloud` change.

| File | Change |
|---|---|
| `v2/resources/user.go` | add `Importer`; nil-guard; prefer `d.Id()`; 404 handling |
| `v2/resources/role.go` | same |
| `v2/resources/auth_profile.go` | same |
| `v2/schemas/detailed_user.go` | fix `customUserInput` key; add 3 missing fields to schema + setter; reclassify timestamps; drop the bogus `strfmt.DateTime` assertions; `ForceNew` on `username` |
| `v2/schemas/role.go` | `ForceNew` on `name` (schema + setter are otherwise complete and correct) |
| `v2/resources/util.go` | shared `isNotFound(err) bool` helper for the 404 path |
| `v2/resources/user_test.go`, `role_test.go` | add import steps |
| `v2/resources/testdata/iam/` | role fixture (currently role + user share `user.create.tf`) |
| `v2/docs/resources/` | generated docs incl. import syntax — see §8 |

`v2/schemas/role.go` needs no data-fidelity work: `RoleSchema()` covers every field of
`models.Role` including `scopes` and `project_tags`, and `SetRoleResourceData` sets all nine.
Role is the cheap half of this ticket.

---

## 6. Explicitly out of scope

The ticket's postscript asks for *"all objects to support state import."* After this work the
remaining three gaps are all deliberate:

| Resource | Why it cannot be imported |
|---|---|
| `zedcloud_credential` | `ReadContext: schema.NoopContext`. The API never returns credential material — by design. Nothing to import into state. |
| `zedcloud_cluster_group_manifest` | Read-only projection; `Create`/`Update`/`Delete` are all `schema.NoopContext`. |
| `zedcloud_kubernetes_cluster_status` | Same — a status view, not a managed object. |

Recommend replying on the ticket that `zedcloud_credential` is permanently non-importable for
security reasons, so "all objects" can be closed honestly rather than left open.

---

## 7. Security considerations

- **Auth profiles carry secrets.** `models.OAUTHProfile` has `ClientSecret`, `CryptoKey` and
  `EncryptedSecrets`. Import will write whatever the API returns into plaintext Terraform
  state. Before enabling import on `zedcloud_auth_profile`, confirm what
  `GET /v1/authorization/profiles/id/{id}` actually returns for those fields; if they come
  back populated, mark them `Sensitive: true` and add them to `ImportStateVerifyIgnore`. If
  the API redacts them (the proto summaries say *"without security details"* for the role/user
  equivalents), document that the operator must supply them in config after import.
- **Users and roles carry no secrets.** `DetailedUser` has no credential field —
  passwords live behind `zedcloud_credential`, which is not importable. Importing a user
  leaks nothing beyond what `GET /v1/users/id/{id}` already returns to the same caller.
- **No new authz surface.** Import reuses the same authenticated GETs the data sources
  already call; a caller who cannot read the object cannot import it.
- **Roles grant privilege.** Importing a role brings its `scopes` under Terraform management,
  so a later careless apply can widen or narrow permissions. This is inherent to managing RBAC
  as code; the `ForceNew` on `name` (§3.7) at least prevents silent identity swaps.

---

## 8. Testing strategy

Follow the one existing precedent, `resources/cep_profile_test.go:61`: append an import step
to the existing create/update acceptance test rather than writing a standalone one. That
reuses the object the earlier steps created and gives `ImportStateVerify` a known-good
baseline.

```go
// Import
{
    ResourceName:      "zedcloud_user.test_tf_provider",
    ImportState:       true,
    ImportStateVerify: true,
},
```

Coverage to add:

1. **`TestUser_Create`** — import by ID with `ImportStateVerify: true`, no ignore list. If an
   ignore list turns out to be needed, each entry is a §3 bug that has not been fixed.
2. **`TestRole_Create`** — same. Needs a dedicated role fixture; today the role lives inside
   `testdata/iam/user.create.tf` as a dependency of the user.
3. **Import by name** — second import step with
   `ImportStateIdFunc` returning the username / role name, verifying §4.1 resolution.
4. **Regression for §3.4** — a user fixture with `custom_user_input` set; assert it survives
   create → read → update → import. This fails on `main` today.
5. **Regression for §3.3** — delete the object out of band via the API client, then
   `RefreshState` and assert the plan proposes a re-create instead of erroring.
6. **`TestProvider_InternalValidate`** — `provider_test.go` has no `InternalValidate` test
   today (it only wires up `testProvider` for the acceptance suite). Add the standard
   one-liner; it runs without a tenant and catches a malformed `Importer` or schema at
   `go test` time rather than at `terraform apply` time.

Unit-level: the schema fixes in §3.4–3.6 are testable without a live tenant, in the style of
`3595da57` (`test(schemas): regression coverage for CI-834 and CI-709`) — build a
`schema.TestResourceDataRaw`, round-trip through `DetailedUserModel` /
`SetDetailedUserResourceData`, assert equality. Do this first; it is fast and it pins the
exact defects.

Docs: `v2/docs/resources/` is empty, so there is currently nowhere that documents import
syntax for *any* resource. Out of scope to fix wholesale, but the three resources touched here
should get hand-written `import` examples, and `make docs` / `tfplugindocs` being unwired
deserves its own ticket.

---

## 9. Alternatives considered

**Bare `ImportStatePassthroughContext`, matching the other 28 resources.** Cheapest, and
consistent. Rejected as the primary interface because it forces the operator to first discover
a 28-character opaque ID — via `GET /v1/users`, the UI, or a `zedcloud_user` data source —
before they can import. That is real friction for the exact workflow the customer asked for.
The custom `StateContext` falls back to passthrough behaviour for anything ID-shaped, so
nothing is lost.

**Ship layer 1 only (importer, no read-path or fidelity fixes).** ~half a day. Rejected:
`terraform import` would succeed and the very next `plan` would show a diff that never
converges (§3.4), which is a worse customer experience than the current honest error.

**Fix it as a data-source workflow instead** — document `data "zedcloud_user"` +
`moved`/`removed` blocks. Rejected: data sources cannot bring an object under management; the
customer would still be unable to `terraform apply` changes to an existing account.

**Do the same sweep for all 34 resources in one pass.** Rejected for this ticket — 28 already
have importers, and auditing each one's round-trip fidelity is a much larger piece of work.
The §3 bug classes (wrong map keys, missing setter fields, `Optional` where `Computed`
belongs) are almost certainly not unique to `detailed_user.go`, and a generator-wide audit is
worth its own ticket.

---

## 10. Sequencing

| # | Work | Est. | Depends on |
|---|---|---|---|
| 1 | Schema round-trip unit tests for user + role (red) | 0.5d | — |
| 2 | §3.4 `customUserInput` key, §3.5 missing fields, §3.6 timestamps | 0.5d | 1 |
| 3 | §3.1 nil-guard, §3.2 prefer `d.Id()`, §3.3 404 → gone | 0.5d | — |
| 4 | §3.7 `ForceNew` on `username` / `name` | 0.25d | 2 |
| 5 | `Importer` on user + role, ID-or-name `StateContext` | 0.5d | 3 |
| 6 | Acceptance import steps + role fixture | 1–1.5d | 5 |
| 7 | `zedcloud_auth_profile` — same treatment, after the §7 secret audit | 0.5d | 5 |
| 8 | Import docs for the three resources | 0.25d | 5 |

**Total 4–4.5 days.** Steps 1–6 close the ticket as written (users + custom roles); step 7 is
the natural extension and step 8 is what makes it discoverable.

Lead with steps 1–2. The `customUserInput` defect is independently shippable, is a bug on
`main` today regardless of import, and writing its test first establishes the harness the rest
of the work uses.

---

## 11. Open questions

1. **Does `GET /v1/authorization/profiles/id/{id}` return `clientSecret` / `cryptoKey` in
   clear?** Gates step 7 and determines whether those fields need `Sensitive: true`. Needs a
   live call against a tenant — cannot be answered from the tree.
2. **Does `GET /v1/users/id/{id}` populate `allowedEntpProjectList`?** The swagger declares it;
   whether the handler fills it is unverified. If it does not, §3.5 becomes a zedcloud-side
   ask after all — the only candidate in this whole design for an API change.
3. **Is a 28-character username realistically possible?** If enterprise SSO can produce
   non-email usernames, the §4.1 heuristic needs an explicit `id:` / `name:` prefix escape
   hatch instead. Confirm with IAM.
4. **Is `Optional` → `Computed` on `last_login_time` / `last_logout_time` acceptable as a
   minor-version change,** or does it need a major bump? No known consumer sets them.
5. **Should `zedcloud_credential` be formally documented as non-importable** in the provider
   docs, or just answered on the ticket? — *partly answered: `TestProviderResourcesAreImportable`
   now encodes it in the test suite with the reason attached.*

---

## 12. What implementation changed about this design

Recorded against the branch `nfr-165-iam-state-import`. Everything in §1–§11 above is as
written unless it appears here.

### 12.1 §3.5 is not implementable as designed — deferred

The three missing fields cannot be added to the schema, because **the generated model does
not have them**. `models.DetailedUser` has no `AllowedEntpProjectList`, `ZksSyncStatus` or
`TokenZksSyncStatus`, and the `ProjectList` and `ZksSyncStatus` model types do not exist in
the tree at all.

The cause is that the checked-in models are stale relative to the checked-in swagger:

```bash
git log -1 --format='%h %ad' --date=short -- v2/models/detailed_user.go   # fe2050af 2024-03-28
git log -1 --format='%h %ad' --date=short -- swagger/zedge_user_service.swagger.json  # 43f3590b 2025-11-14
grep -c 'AllowedEntpProjectList' v2/models/detailed_user.go               # 0
ls v2/models/ | grep -iE 'project_list|zks_sync'                          # no match
```

Closing this gap means regenerating models from the current swagger (`make swagger-download`
then the `swagger generate` targets), which rewrites a large share of the 556 files under
`v2/models/` and is far too wide a blast radius for this ticket. Hand-writing two new model
types into generated files is worse: the next `make` clobbers them.

Also note `allowedEntpProjectList` carries `"title": "internal"` in the swagger, which is
plausibly why the generator was pointed at a filtered spec in the first place — it may not be
customer-facing configuration at all.

**Action:** file a separate ticket to regenerate the models against the current swagger, and
re-evaluate §3.5 there. The rest of NFR-165 does not depend on it. Users who have
`allowedEntpProjectList` set will import without it, exactly as they do today — this is a
pre-existing gap the ticket does not widen.

### 12.2 New defect: every `map[string]string` attribute is dropped (was not in §3)

`§3.4` turned out to be one instance of a generator-wide bug class, not a one-off typo.
`d.Get` on a `TypeMap` returns `map[string]interface{}`, never `map[string]string`, so the
generated form

```go
projectTags, _ := d.Get("project_tags").(map[string]string)   // always nil
```

silently drops the attribute, and the `...ModelFromMap` variant

```go
projectTags := m["project_tags"].(map[string]string)          // panics
```

would panic outright if it were ever reached with a live map. This hits `zedcloud_role`'s
`project_tags` directly, which is in scope here, and six other schema files besides:

```bash
grep -rn '\.(map\[string\]string)' v2/schemas/
```

Fixed for `role.go` and `detailed_user.go` with a shared `stringMap` helper in
`schemas/schema_helpers.go` that accepts both shapes. The other call sites are all on
device status/summary structs and are left alone — same follow-up ticket as §12.1.

### 12.3 New defect: user state was never refreshed after an update

`IdentityAccessManagement_UpdateUser2` ended with

```go
if errs := IdentityAccessManagement_GetUser(ctx, d, m); err != nil {
```

testing `err` — the update error, necessarily `nil` at that point — instead of `errs`. The
post-update refresh result was discarded and state was never repopulated. Fixed to
`errs.HasError()`. Not import-specific, but it would have masked import round-trip problems
during testing.

### 12.4 §3.3 needed no new helper

`isStatusNotFound` already exists in `resources/util.go`, added by `40ec46a5` (CI-723) for the
delete paths. The read paths just needed to call it. No new helper was written.

### 12.5 §8 was wrong about the docs

Resource docs *are* generated, into the repo-root `docs/` directory — all 34 resources are
there. `v2/docs/resources/` is the empty staging directory that `tfplugindocs` writes through,
which is what the original audit saw. `make docs` works as-is.

Import documentation was therefore added the idiomatic way, as
`v2/examples/resources/zedcloud_{user,role}/import.sh`, which `tfplugindocs` renders into an
`## Import` section automatically. No custom templates needed. `make docs` touches only
`docs/resources/{user,role}.md` and `docs/data-sources/user.md`.

### 12.6 Added a guard that did not exist in the plan

`TestProviderResourcesAreImportable` walks `ResourcesMap` and fails on any resource with no
`Importer` that is not on an explicit exclusion list carrying a written reason. This turns
"all objects should support state import" from a ticket comment into something the suite
enforces, and it is what will remind whoever picks up §7 to delete the `zedcloud_auth_profile`
entry.

### 12.7 New defect, found only by running against a controller: unvalidated `custom_user_input` keys

Fixing §3.4 meant `custom_user_input` reached the API for the first time — and the API
rejected it:

```
Error: Zedcloud API call returned HTTP status code: 400
Error: Invalid characters in the key
```

The controller requires a custom-parameter key to be **exactly three `_`-separated
segments** from a restricted alphabet (`ValidateCustomParam` / `validCustomParam` in
zedcloud `libs/zutils/validate.go:792`):

```go
regexp.MustCompile("^[a-zA-Z0-9-.%@#:~!=]+_[a-zA-Z0-9-.%@#:~!=]+_[a-zA-Z0-9-.%@#:~!=]+$")
```

So `department` is invalid; `acme_ui_theme` is valid. Nothing in the swagger, the provider
schema or the docs said so, and the failure named neither the attribute nor the offending
key. This was undiscoverable before the §3.4 fix, because the map never left the provider.

Added `validateCustomUserInput` as a `ValidateDiagFunc` mirroring the controller's rules
(key shape, 1024-char key and value limits, non-empty), so the error arrives at plan time
against the specific key. The schema description now states the format, which `make docs`
propagates to `docs/resources/user.md`.

This is a good argument for the follow-up ticket in §12.1 to also carry "surface the
controller's validation rules in the provider schema" — this will not be the only one.

### 12.8 Verified against a live controller

Run against `zedcontrol.local.zededa.net` on 2026-09-28, provider rebuilt from this branch:

| Suite | Result |
|---|---|
| `go test ./...` (unit, no `TF_ACC`) | ok — resources, schemas, testing, pkg/compare |
| `TestRole_CreateAndImport` | **PASS** (24.7s) — create, import by ID, import by name, destroy |
| `TestUser_Create` | **PASS** (24.2s) — create, import by ID, import by username, destroy |
| `TestUser\|TestRole\|TestEnterprise` | **PASS** (142.4s) — no collateral damage to neighbouring IAM tests |

`ImportStateVerify` is on for all four import steps with **no `ImportStateVerifyIgnore`**,
which is the outcome §8 asked for: every attribute the Read path produces matches what the
create step put in state.

The request log confirms the §3.2 ordering fix in practice — the by-name GET appears only
for the create step and the import-by-username step, each immediately followed by a GET
by ID, while every refresh goes straight to `/v1/users/id/{id}`.

### 12.9 Effort

Steps 1–6 and 8 landed in well under the 4–4.5 day estimate, because §3.5 — the one item that
needed model regeneration — turned out to be the deferred one. Step 7 (`auth_profile`) remains
blocked on the §11.1 secret audit, which needs a live tenant.
