# Design: ENG-3000 — Editable interface ACLs on `zedcloud_application_instance`

**Ticket:** [ENG-3000](https://zededa.atlassian.net/browse/ENG-3000) — *[NFR-116] Payload sends removed ACL rules even after modifying*
(Terraform question in [comment 106827](https://zededa.atlassian.net/browse/ENG-3000?focusedCommentId=106827))
**Controller fix:** zedcloud [#8548](https://github.com/zededa/zedcloud/pull/8548) — *fix(seine): treat bundle ACLs as defaults on ordering-enabled projects (ENG-3000)*, merged as `9974421e`
**Repos:** `zededa/terraform-provider-zedcloud` @ `4c68a184` (origin/main), `zededa/zedcloud` @ `pr-8548` (`b56b4510`)
**Repro:** `test/eng3000-repro/` on branch `eng-3000-tf-acl-repro`
**Date:** 2026-10-27
**Author:** Ivan Curachin
**Status:** Draft — decisions D1–D5 proposed, not yet agreed

**Headline:** after #8548 the controller stores an instance's interface ACL list exactly as
sent (empty list = bundle defaults) on projects with interface ordering enabled. The provider
can't use this, because `interfaces` has a blanket diff suppressor (CI-201) that hides every
change to an interface after create. The proposal: stop suppressing ACL changes, make
`interfaces[].acls` Optional+Computed so "omitted" means "bundle defaults", and clear ACL diffs
in `CustomizeDiff` on ordering-disabled projects, where the controller still ignores them.
Every other interface field stays suppressed (NFR-116 interface add/remove/rebind is a
separate ticket).

---

## 1. Problem Statement

### 1.1 What users hit

ENG-3000: a VM app bundle defines eth0 rules `allow 0.0.0.0/0`, `9022->22`, `9080->80`. The
user deploys an instance and changes `9022->22` to `8022->22`. Before #8548, the controller
appended the bundle rules behind the request, so the instance forwarded **both** 8022 and 9022
to port 22. A second instance with moved ports then failed on EVE with
`overlapping portmaps`.

Yuri asked whether Terraform behaves the same. It does (run on 2026-10-26 against a pre-fix
local cluster, `test/eng3000-repro/README.md`):

| Case | Instance ACLs in TF | Stored by controller (pre-#8548) | `tofu plan` after |
|---|---|---|---|
| a: edits a bundle rule | allow, 8022->22, 9080->80 | 8022->22, allow, **9022->22**, 9080->80 | No changes |
| b: `acls` omitted | none | allow, 9022->22, 9080->80 | No changes |
| c: echoes the bundle | allow, 9022->22, 9080->80 | allow, 9022->22, 9080->80 | No changes |

Changing case a's `8022` to `7022` in config afterwards also planned **No changes**.

### 1.2 What #8548 changed in the controller

Only on projects with `interface_ordering = INTERFACE_ORDERING_ENABLED`:

| Path | Non-empty ACL list | Empty (or omitted) list |
|---|---|---|
| Create, optional interface re-add | stored as sent | bundle ACLs |
| Update (interface matched by name) | replaces the stored list as sent | resets to bundle ACLs |
| Bundle refresh | untouched | untouched |

- A rule sent without an id that matches a stored rule keeps its id (`reuseStoredACLIds`), so
  resending an unchanged id-less list is a no-op. The provider never sends ids, so this is the
  path it takes.
- An empty list can't be stored on an interface whose bundle defines ACLs.
- Ordering-disabled projects are unchanged: ACL edits on update are ignored, and create still
  appends the bundle ACLs (folding echoed copies, ENG-3001).

### 1.3 What the provider does today

- `interfaces` (`v2/schemas/application_instance.go:667`) has
  `DiffSuppressFunc: suppressAppInterfacesChangeAfterCreation("interfaces")`
  (`v2/schemas/schema_helpers.go:270`, from CI-201, `70c2ab39`, 2023). It returns `true` for
  every key once the interface list in state is non-empty. In SDKv2 a top-level suppressor is
  called for every nested attribute key (`schemaMap.diff`, `schema.go:1140`), so
  `interfaces.0.acls.1.matches.1.value` is suppressed too. That's why case a shows no drift.
- `interfaces[].acls` is a plain Optional `TypeList` (`v2/schemas/app_interface.go:219`). The
  ACE `id` is already Computed.
- Create and Update send the config's `interfaces` in full (`ApplicationInstanceModel(d)`),
  on both the v2 async route and the v1 route. Read writes the controller's interfaces back
  into state.

Against a controller with #8548, this already has two side effects, both hidden by the
suppressor:

1. Any unrelated update (say `vminfo.cpus`) now replaces the stored ACLs with the config's.
   For case a that silently *fixes* the instance, but the plan never said ACLs would change.
2. If config omits `acls`, every update sends an empty list, which resets the interface to the
   bundle defaults and wipes any rules set in Cortex.

So with the fixed controller, the provider changes device firewall rules without showing it in
the plan. That's the main reason to do this now rather than wait for the NFR-116 interface work.

---

## 2. Proposed Solution

1. **Stop suppressing ACL keys.** `suppressAppInterfacesChangeAfterCreation` returns `false`
   for keys under `interfaces.<n>.acls`. Everything else under `interfaces` stays suppressed
   exactly as today (D1).
2. **`acls` becomes Optional+Computed** on `zedcloud_application_instance` only. Omitting
   `acls` means "whatever the controller stores", which is the bundle defaults after create.
   No diff, and Update sends the stored list back, which `reuseStoredACLIds` makes a no-op (D2).
3. **Ordering-disabled projects:** `CustomizeDiff` looks up the instance's project (cached per
   provider process, like `projectLookups`). If interface ordering is off, it calls
   `d.Clear("interfaces.<n>.acls")` for each interface, which keeps today's behaviour there (D4).
4. **Read-back check:** after Update, if the stored ACL list of an interface differs from what
   was sent on an ordering-enabled project, return a **warning** naming the interface. That
   points users at a controller without #8548 instead of leaving them with an unexplained
   permanent diff (D5).
5. **Docs and release notes:** document the defaults semantics and the first-plan-after-upgrade
   effect (§5.3).

The resulting user-facing behaviour, on ordering-enabled projects:

| Config | Create | Later plan/apply |
|---|---|---|
| `acls` omitted | bundle defaults stored | no diff |
| `acls` listed | stored as written | edits plan and apply as normal list changes |
| `acls` listed, then the block removed | — | no diff; stored rules kept (D3) |
| `acls` equal to the bundle rules | stored as written | no diff |

---

## 3. Alternatives Considered

| Alternative | Why not |
|---|---|
| **Delete the suppressor entirely** | Unsuppresses every interface field. CI-201 added it because the controller fills in interface fields (MAC, IP, order, bundle ACLs) the config doesn't carry, so the old perpetual diffs come back. Interface add/remove/rebind (NFR-116) also needs its own design: gating on device capability, restarts, order rules. |
| **Keep `acls` Optional-only, and treat empty config as "reset to defaults"** | State always holds the defaults, config holds nothing, so every plan shows a diff. Apply sends an empty list, the controller resets to the same defaults, and the diff repeats forever. Suppressing that case in the diff suppressor would also suppress a real reset. |
| **Compare against the bundle in `CustomizeDiff` so removing `acls` resets to defaults** | Needs an app GET per instance per plan, and a byte-for-byte copy of seine's `buildACEsFromAcls` + `aceContentEqual` to decide what "equal to the defaults" means. It also can't produce the diff: `SetNew`/`SetNewComputed` only work on top-level keys and `interfaces` isn't Computed. Possible as a follow-up (D3). |
| **Track "ACLs came from config" in a hidden Computed attribute** | Makes "remove the block → reset to defaults" work without an API call, but adds a state-only attribute with import edge cases (unknown after import) for a rare operation. Kept as the D3 fallback. |
| **Make `acls` Computed on the shared `AppInterface()` schema** | The same schema is used by `app_config`, `profile_app_config` and `app_inst_config` (deployments, app profiles). Their controller semantics weren't part of #8548. Only the app-instance copy changes. |
| **Fail the plan on ordering-disabled projects when ACLs differ** | Configs that work today would break: on those projects create still appends the bundle ACLs, so state differs from config whenever the bundle has rules the config doesn't. |

---

## 4. Interface / API Design

### 4.1 Schema

In `ApplicationInstanceResourceSchema()` (`v2/schemas/application_instance.go:867`), replace the
`interfaces` element with a copy of `AppInterface()` whose `acls` has `Computed: true`:

```go
s["interfaces"].Elem = &schema.Resource{Schema: appInstanceInterfaceSchema()}

func appInstanceInterfaceSchema() map[string]*schema.Schema {
	s := AppInterface()            // fresh map each call, so the shared schema is untouched
	s["acls"].Computed = true      // omitted = controller's list (bundle defaults on create)
	s["acls"].Description = "..."  // see §4.4
	return s
}
```

The data source (`ApplicationInstanceDataSource`) uses `ApplicationInstance()` and doesn't
change.

### 4.2 Diff suppressor

```go
var aclKey = regexp.MustCompile(`^interfaces\.\d+\.acls(\.|$)`)

func suppressAppInterfacesChangeAfterCreation(mapKey string) schema.SchemaDiffSuppressFunc {
	return func(key, oldValue, newValue string, d *schema.ResourceData) bool {
		if aclKey.MatchString(key) {
			return false // ENG-3000: ACL lists are editable; CustomizeDiff gates them
		}
		... unchanged ...
	}
}
```

The suppressor stays scoped to this resource (it's only used at `application_instance.go:675`).

### 4.3 `CustomizeDiff`

`applicationInstanceCustomizeDiff` (now only `checkAppInstanceSnapshotSettings`) gains:

```go
func gateInterfaceACLDiffs(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
	if d.Id() == "" {
		return nil // create: the controller handles both project kinds
	}
	changed := aclChangedInterfaces(d) // indexes n where HasChange("interfaces.n.acls")
	if len(changed) == 0 {
		return nil // no API call on the common path
	}
	enabled, err := projectInterfaceOrdering(ctx, m, d.Get("project_id").(string))
	if err != nil {
		return err // or: log and keep the diff (D4 open point)
	}
	if enabled {
		return nil
	}
	for _, n := range changed {
		if err := d.Clear(fmt.Sprintf("interfaces.%d.acls", n)); err != nil {
			return err
		}
	}
	return nil
}
```

`ResourceDiff.Clear` accepts nested keys, provided the leaf is Computed (`checkKey(key, "Clear",
true)`), which §4.1 makes `acls`. `projectInterfaceOrdering` GETs the project once per provider
process and reads `tag_level_settings.interface_ordering`, using the same singleflight cache
shape as `projectLookups` (`v2/resources/application_instance.go:209`).

### 4.4 Docs wording for `interfaces.acls`

> Interface ACLs. On projects with interface ordering enabled, the bundle's ACLs for this
> interface are defaults: omit `acls` to use them, or list the rules the interface should
> have. A listed set replaces the defaults completely, so repeat any bundle rule you want to
> keep. Removing the block later keeps the current rules; to go back to the defaults, list
> the bundle's rules. On projects with interface ordering disabled, ACL changes after create
> are ignored by the controller and are not shown in plans.

---

## 5. Affected Components

### 5.1 Files

| File | Change |
|---|---|
| `v2/schemas/schema_helpers.go` | ACL key exemption in `suppressAppInterfacesChangeAfterCreation` |
| `v2/schemas/application_instance.go` | `appInstanceInterfaceSchema()` with Computed `acls` |
| `v2/resources/application_instance_snapshots.go` (or a new `application_instance_acls.go`) | `gateInterfaceACLDiffs`, wired into `applicationInstanceCustomizeDiff` |
| `v2/resources/application_instance.go` | project-ordering lookup/cache; read-back warning after Update (D5) |
| `docs/resources/application_instance.md` | regenerated with `make docs` (tfplugindocs) after the §4.4 description change |
| GitHub release notes | §5.3 (the repo has no CHANGELOG file) |
| `test/eng3000-repro/` | turned into a repeatable check (§7) |

### 5.2 Not affected

- `zedcloud_application` (bundle ACLs): unchanged.
- Deployments and app profiles (`app_config`, `profile_app_config`): unchanged; whether they
  have the same append problem is an open question (§8).
- Every interface field other than `acls`: still suppressed after create.

### 5.3 Upgrade effect

- **Ordering-enabled instances created before #8548** that store a merged list (bundle rules
  behind the config's) will show an ACL diff on the first plan after upgrading. Applying it
  removes the bundle rules the user had edited away. That's the intended fix, but it changes
  device firewall rules, so the release notes must say so.
- **Instances whose ACLs were edited in Cortex** while the config omits `acls`: no diff, and
  the Cortex rules are kept (an improvement over today, where any update resets them, §1.3).
  If the config *lists* ACLs, Terraform now shows and reverts the Cortex edits, as it would
  for any other attribute.
- **Controllers without #8548:** ACL edits on an ordering-enabled project come back with the
  bundle rules appended, so the diff never settles. D5's warning makes this visible. Release
  notes should name the minimum controller release (§8).

---

## 6. Security Considerations

ACLs are the app's inbound firewall and port maps on the edge node, so this is
security-relevant:

- **Today** the provider can't remove an exposed port map that the bundle defines, and with a
  fixed controller it can change ACLs without showing them in the plan (§1.3). Both get better.
- **After the change**, a config that lists ACLs fully defines them: a bundle rule not repeated
  in config is **not** applied. That's what #8548 decided, and it's documented in §4.4. Users
  relying on "bundle rules are always added" lose that on ordering-enabled projects.
- A bundle ACL change no longer reaches existing instances on refresh (#8548). Tightening a
  bundle rule needs a config change or a reset on each instance.
- No new credentials, endpoints or data. The project lookup uses the existing API token.

---

## 7. Testing Strategy

**Unit (no controller):**
- Suppressor: ACL keys (`interfaces.0.acls.#`, `interfaces.1.acls.2.matches.0.value`) return
  `false`; `interfaces.0.macaddr`, `interfaces.#` keep today's result.
- `gateInterfaceACLDiffs` with a stubbed project lookup: ordering on → diff kept; ordering off
  → `interfaces.n.acls` cleared and other diffs kept; create → no lookup; no ACL change → no
  lookup.
- Schema: `AppInterface()["acls"].Computed` is still `false` (shared schema untouched); the
  app-instance copy is `true`.

**Local cluster with #8548** (extend `test/eng3000-repro/`, ordering-enabled project unless noted):

| # | Steps | Expected |
|---|---|---|
| 1 | create a (edited default) | stored as written, plan clean |
| 2 | a: change 8022 to 7022 | plan shows one ACL change; apply; stored 7022; plan clean |
| 3 | b: create with `acls` omitted | bundle defaults stored, plan clean |
| 4 | b: change `vminfo.cpus` | ACLs untouched (stored list sent back) |
| 5 | c: create echoing the bundle | stored once, plan clean |
| 6 | a: remove the `acls` block | plan clean, rules kept (D3) |
| 7 | ordering-disabled project, edit ACLs | plan clean (cleared), controller unchanged |
| 8 | ACLs edited in Cortex, config omits `acls` | plan clean; config lists `acls` → diff back to config |
| 9 | instance stored pre-fix with a merged list | first plan shows the removal; apply; clean |

Run each with `async_writes` on (v2) and off (v1). Repeat 1–2 against a controller without
#8548 to check the D5 warning.

**e2e:** add case 2 to the virtual-EVE e2e suite so the device actually reports the new port
map and `overlapping portmaps` no longer happens for a second instance on moved ports.

---

## 8. Decisions and Open Questions

### 8.1 Proposed

- **D1 Scope.** ACLs only. The rest of the interface (add/remove/rebind/order/MAC, NFR-116)
  gets its own provider ticket; #8548 lists it as a follow-up.
- **D2 Omitted = defaults/kept.** `acls` Optional+Computed on the app-instance schema only.
- **D3 No reset by removing the block.** To go back to the defaults, list the bundle rules (or
  recreate the instance). Revisit with the hidden-attribute or bundle-compare option (§3) if
  users ask.
- **D4 Ordering-disabled projects.** Clear ACL diffs in `CustomizeDiff`, keeping today's
  behaviour. Needs one cached project GET per project per run, only when an ACL diff exists.
- **D5 Read-back warning.** Warn after Update when the stored ACLs differ from what was sent on
  an ordering-enabled project.

### 8.2 Open

1. Which zedcloud release first ships #8548? Needed for the release notes and for D5's text.
2. D4 when the project lookup fails: fail the plan, or keep the diff and log? Leaning towards
   keeping the diff, since the server ignores the update anyway.
3. Ids in plans: when the list changes, the plan shows the old ids at each position until
   apply returns the real ones. SDKv2 tolerates this, but the plan output is misleading. Accept,
   or hide ids from the plan?
4. Do deployments and app profiles (`app_config`, `profile_app_config`) have the same bundle
   ACL append? If yes, separate tickets.
5. Should the provider read the project setting at all, or should the controller report
   whether ACL edits are honoured (for example a field on the instance)? A controller field
   would remove the extra GET and the project-toggle edge case.
