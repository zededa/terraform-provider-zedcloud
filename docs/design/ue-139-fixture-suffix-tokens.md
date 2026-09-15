# UE-139 — `__SUFFIX__` token expansion across the existing test fixtures

**Ticket:** [UE-139](https://zededa.atlassian.net/browse/UE-139) (T5, epic [UE-131](https://zededa.atlassian.net/browse/UE-131))
**Repo:** `zededa/terraform-provider-zedcloud`, checked against branch `ue-137-self-hosted-gh-runner` @ `9986cfe`
**Design refs:** `docs/design/e2e-virtual-eve-node-testing.md` §6.1, §6.2, §6.5

**Goal.** Make every object the acceptance suite creates carry a per-run suffix, so two
runs against the shared alpha enterprise cannot collide on unique-name constraints — the
precondition for the per-PR trigger in [UE-141](https://zededa.atlassian.net/browse/UE-141).
Land it as a change that is *provably* a no-op with the suffix empty.

---

## 1. Verified ground truth

Measured against the tree, not taken from the ticket. Numbers differ slightly because
UE-138 landed `create_real_node.tf` in the meantime.

| Thing | Ticket says | Actually |
|---|---|---|
| `.tf` fixtures under `v2/resources/testdata/**` | 45 | **46** (45 + `application_instance/create_real_node.tf`, already tokenised) |
| Fixtures needing rewrite | 45 | **44** |
| `name =` / `title =` lines in fixtures | — | **535** total; **345** are object names to suffix, **128** are structural/external (must not change), 2 are HCL references (`= zedcloud_application.x.title`), rest already tokenised |
| `name = "test_tf_provider-create_edgenode"` | 34× | **33×** as `name`, plus **20×** as `title` |
| Shared `serialno` | 2 fixtures | confirmed — `2293dbe8-…-962857efc46b` in `application_instance/create.tf:96` and `patch_reference_update/create.tf:98` |
| `MustGetTestInput` call sites | — | **47** across 25 test files |
| `cmpopts.IgnoreFields` blocks | ~45 / ~120 entries | **49** blocks across 18 test files |
| Golden `.yaml` files | — | **37**; **69 lines** need a token (68 exact + 1 derived) |
| Literal name/title assertions in Go | — | **33** inline `TestCheckResourceAttr(…, "name"/"title", "literal")` in 5 files, plus 1 multi-line in `deployment_data_source_test.go:30` |

Reproduce:

```bash
cd v2/resources
find testdata -name '*.tf' | wc -l
grep -rhoE '^\s*(name|title)\s*=\s*".*"' testdata --include='*.tf' | sort | uniq -c | sort -rn
grep -rn 'MustGetTestInput' --include='*.go' .. | grep -v '/testing/' | wc -l
grep -rn 'IgnoreFields' --include='*_test.go' . | wc -l
```

Machinery that already exists and this ticket only has to reach:

- `testhelper.MustGetTestInputWithVars` (`v2/testing/realnode.go`) — expands `__TOKEN__`,
  hard-fails on a leftover `__…__`. No fixture contains a stray `__` today, so the guard
  will not false-positive.
- `testhelper.EnvTestSuffix = "ZEDCLOUD_TEST_SUFFIX"`.
- `.github/workflows/e2e.yml:277` already exports `ZEDCLOUD_TEST_SUFFIX: _${{ env.RUN_ID }}`
  for the real-node step.
- `application_instance/create_real_node.tf` is the worked precedent for *which* attributes
  get a suffix and which do not.

---

## 2. Three corrections to the ticket's approach

### 2.1 `sed` every `name =` / `title =` is wrong — 128 of 535 lines must not change

The fixtures use `name` for four unrelated things. Suffixing the wrong ones does not
produce a collision, it produces a broken apply or a workload that never starts:

| Class | Count | Examples | Why it must not change |
|---|---|---|---|
| Interface names | 3 | `name = "eth0"`, `"eth1"` | kernel-global on the device; a suffixed port does not exist |
| Manifest resource keys | 23 | `cpus`, `memory`, `resourceType`, `bootparam` | key names the controller interprets, not labels |
| Custom-config keys | 12 | `custom_config_name`, `var_group`, `condition`, `indirect` | ditto |
| External artifacts | 9 | `xenial-amd64-docker-20180725`, `alpine_24`, `docker_compose_app_bundle` | must resolve in the datastore; suffix ⇒ image pull fails |
| Filler | 81 | `title = "title"` (66×), `"test"` (11×) | no uniqueness constraint; churn for nothing |

`create_real_node.tf` already draws this line: it suffixes resource-level `name`/`title`
and `manifest.name`, and leaves `resources { name = "cpus" }` and
`interfaces { name = "eth0" }` alone.

**So: a nesting-aware rewriter with an explicit allowlist, not `sed`.** §5 has it; dry-run
output is 345 lines across 44 files, and the diff is reviewable line by line.

### 2.2 Golden `.yaml` names cannot be blind-appended — one is controller-derived

The ticket's "check the `IgnoreFields` lists and extend where a name is asserted" is the
right instinct but the wrong remedy: adding `"Name"`/`"Title"` to 18 ignore lists deletes
the assertion that the provider sends the name it was given. Only `project_test.go` does
this today (3 places, for `models.Policy` where the controller owns the name).

Better: give the goldens the same token and expand them at load time (§3.3). 69 lines,
36 files. Of those, one is **derived** and needs the token mid-string:

```yaml
# testdata/project/create.yaml:31  — controller builds this from the project's name
- name: test_tf_provider.edgeviewPolicy
+ name: test_tf_provider__SUFFIX__.edgeviewPolicy
```

Verified as derived (the string appears in no `.tf` literal). The similar-looking
`deployment/create.yaml:7` `…-edge-node-policy-attestation` and
`project/create_with_pnac.yaml:6` `…_pnac_attest` are built from *nested* policy blocks,
which the allowlist does not suffix — they stay unchanged. That consistency is a reason to
keep nested names out of scope.

### 2.3 Names are not the only collision axis — three others are reachable per-PR

Acceptance criterion 3 ("two runs with different suffixes, back to back, no collisions")
is not met by names alone:

1. **`serialno`** — `application_instance/create.tf` and `patch_reference_update/create.tf`
   share one, and a device serial is unique per enterprise. Two runs ⇒ the second fails to
   create its node. 13 `serialno` literals total.
2. **`iam/user.create.tf`** — `email = "user1@example.com"`, `username = "user2@example.com"`.
   Appending at end-of-value yields an invalid address; needs
   `user1+__SUFFIX__@example.com`.
3. **Tag selectors** — `deployment_tag = "depl_tag"`, `token = "token_profiledeploymenttags"`,
   and the `tags = { depl = "1234" }` / `policyTargetCondition` maps in
   `deployment/create.tf`, `profile_deployments/*`, `asset_groups/*`. Two runs sharing
   `depl_tag` cross-match each other's devices. This one fails *quietly* — as a confusing
   golden diff, not a 409 — which makes it the most expensive to debug and the easiest to
   miss.

These sit in commit 5 (§4). Beyond the ticket's literal wording, inside its acceptance
criteria.

---

## 3. Design decisions

### 3.1 Suffix source: env verbatim, empty means empty

Add to `v2/testing`:

```go
// Suffix returns the per-run object-name suffix, or "" when unset.
//
// Deliberately NOT generated when absent, unlike RealNode's fallback: the
// no-op proof for the fixture rewrite is "unset => byte-identical config",
// which a generated value would make impossible. CI sets it (e2e.yml);
// `make test` leaves it empty and behaves exactly as before.
//
// Charset: object names are validated against [a-zA-Z0-9][a-zA-Z0-9_.-]+
// (68 fields in swagger/), maxLength 256. Longest existing fixture name is 55
// chars, so `_gh<run_id>` is safe. Underscore, dot and hyphen only.
func Suffix() string { return os.Getenv(EnvTestSuffix) }
```

`RealNode()` keeps its generated `_<unix>` fallback — real-node tests are opt-in and always
want isolation.

### 3.2 Placeholder form stays `__SUFFIX__`

`${SUFFIX}` is HCL interpolation and would break `terraform fmt` / `validate` on the
fixtures. Already reasoned in `realnode.go`; restating it here because it is the first thing
a reviewer will want to change.

### 3.3 Goldens expand through the same mechanism

Add `MustGetExpectedOutputWithVars(t, path, vars, i)` mirroring `MustGetTestInputWithVars`
(same `__TOKEN__` replacement, same leftover-token hard failure) applied to the YAML text
before `YAMLToJSON`. Keeps names asserted; no ignore-list weakening.

### 3.4 One PR, six commits

Reviewability matters more than commit count here: commit 2 is 345 mechanical lines and
nobody reads that diff carefully unless the surrounding commits are small.

---

## 4. What landed

Six planned commits, with two deliberate departures from the plan above, both noted
inline.

**Helpers** (`v2/testing/testhelper.go`, `realnode.go`). `Suffix()` reads
`ZEDCLOUD_TEST_SUFFIX` verbatim; `expandTokens`/`mustReadTestdata` are shared by all four
loaders, so the unexpanded-token guard exists once.

**Departure 1 — expansion moved into the existing loaders instead of the call sites.**
`MustGetTestInput` and `MustGetExpectedOutput` now expand `__SUFFIX__` themselves, which is
what design §6.1 originally proposed. This drops the planned 47-call-site commit entirely:
less churn, and a test added later cannot forget the suffix. `MustGetTestInputWithVars`
stays for the real-node fixture's `__NODE_ID__`, and takes its vars explicitly with no env
fallback, so `RealNode`'s generated suffix still flows.

**Fixtures.** `v2/testing/scripts/suffix_fixtures.py --apply`, kept in the tree as the
record of which `name`/`title` values are object names. 369 lines across 44 fixtures — 345
`name`/`title`, plus 24 identity values from the axes below.

**Goldens.** 91 lines across 36 files. The golden pass matches on the *value* rather than
the key, keyed off exactly what the fixture pass tokenised, and that caught ten lines a
`name:`/`title:`-only pass would have missed — `imagename`, `netname`, `imageName`,
`projectName`, `nameAppPart`: controller echoes of a suffixed object name. It also needs
one guard: several fixtures set `description` to the same string as the object's name, and
`description` is *not* tokenised in the `.tf`, so the golden copy must not be either
(`DENY_KEYS`). Without that the suffix would be asserted where the provider never sent it.

**Go assertions.** A `suffixed()` helper in `provider_test.go` (rather than 30×
`fmt.Sprintf` with a per-test local) wraps the 30 inline `TestCheckResourceAttr` name/title
literals plus `deployment_data_source_test.go:30`. Three assertions are deliberately left
literal: `"title"` ×2 and `"required_only-title"`, whose fixture values are untouched
filler.

**Non-name axes** — `serialno` (13), `token` (10) and the `iam/user.create.tf`
email/username, whose token goes in the local part (`user1__SUFFIX__@example.com`) so it
still collapses to a valid address when empty.

**Departure 2 — the tag-selector axis is deferred, not done.** `deployment_tag` and the
`tags` / `policy_target_condition` / `asset_tags` map values are left alone. They are a
*cross-matching* hazard (run A's deployment also targeting run B's device), not a
uniqueness failure, so they cannot break the back-to-back case in AC3, and the single
self-hosted runner serialises CI jobs. Doing it properly means keeping the string form
(`deployment_tag = "depl:1234"`) and the map form (`tags = { "depl" : "1234" }`) in lockstep
across 12 fixtures and their goldens — a change with its own review surface. Worth a
follow-up ticket under UE-131 before anything runs two suites at once.

**New test.** `v2/testing/fixtures_test.go` walks every fixture and golden and fails on a
placeholder no loader expands, so a typo'd `__SUFIX__` is caught by `go test ./...` with no
cluster rather than mid-apply against a shared controller. It also pins that `Suffix()` does
not generate a fallback.

**Then CI.** Nothing to change in `e2e.yml` for the real-node step. When
[UE-141](https://zededa.atlassian.net/browse/UE-141) widens the job to the full suite, set
`ZEDCLOUD_TEST_SUFFIX: _gh${{ github.run_id }}` on it. `make test` (`Makefile:49`) stays
untouched — empty suffix, today's behaviour.

---

## 5. The rewriter

`v2/testing/scripts/suffix_fixtures.py`, committed rather than run-and-discarded: it is the
record of which `name`/`title` values are object names, and the thing to rerun (and to read
`NEVER` / `ALLOWED_PATHS` in) after adding fixtures.

```bash
python3 v2/testing/scripts/suffix_fixtures.py v2/resources/testdata            # dry run
python3 v2/testing/scripts/suffix_fixtures.py --apply v2/resources/testdata
git diff --stat -- v2/resources/testdata
```

Two passes. The fixture pass is nesting-aware: `name`/`title` are suffixed as direct
resource attributes and inside `manifest`, never inside `resources` / `interfaces` /
`custom_config`; `serialno` and `token` only at resource level; `NEVER` catches values that
are structural keys or external artifacts; `MIDSTRING` handles values where the token cannot
go last (the user fixture's email addresses). The golden pass keys off exactly what the
fixture pass tokenised, matching on value rather than key, with `DENY_KEYS` for `description`
and `DERIVED` for the one controller-built name.

The suffix goes *inside* the quotes, so `name  = ` alignment is untouched and `tofu fmt`
sees no change — verified below.

---

## 6. Verification — what was actually run

Against the **local** cluster (`zedcontrol.local.zededa.net`), not alpha.

**Byte-identity, mechanically (the AC1 proof).** All 80 changed fixture and golden files,
with `__SUFFIX__` stripped, are byte-identical to their pre-change content:

```bash
for f in $(git diff --name-only <base> -- v2/resources/testdata); do
  diff -q <(sed 's/__SUFFIX__//g' "$f") <(git show "<base>:$f") >/dev/null || echo "DIFF $f"
done
# checked 80 changed files, 0 differ with SUFFIX=''
```

**Formatting.** 35 fixtures were already not `tofu fmt`-clean at the base commit, and the
same 35 are not clean now — zero newly unformatted (compared via a `git worktree` of the
base). Pre-existing, and untouched by tokenisation.

**Build.** `go build ./... && go vet ./...` clean inside `v2/`. (The repo-root
`go build ./...` fails on a pre-existing `vendor/modules.txt` inconsistency, unrelated.)
`gofmt -l` flags 13 files, of which the two this change touches — `project_test.go`,
`asset_group_service_resource_test.go` — were already unformatted at the base commit (import
ordering).

**Unit tests, no cluster needed.** `go test ./testing/ -run 'TestFixtureTokensAreKnown|TestSuffixIsEnvVerbatim'`
passes; it reports 474 `__SUFFIX__` and 3 `__NODE_ID__` placeholders and no unknown ones.

**Single test, both modes.** `TestDatastore_Create` passes with the suffix unset (objects
created as `test_tf_provider-test_datastore`, exactly as before) and with
`ZEDCLOUD_TEST_SUFFIX=_ue139a` (created as `test_tf_provider-test_datastore_ue139a`), the
golden comparison passing in both — which is the golden tokenisation working end to end.

**First attempt at the full suite: discarded.** It ran 65 passed / 4 failed, but the local
cluster degraded partway through (`dial tcp 10.101.228.195:5432: connection refused`, then a
node flap), so its tail measured the cluster rather than the change. §6.3 covers the resize
that fixed that; §6.2 has the runs that count. Worth keeping one detail from it: even across
those failures there were **zero** golden mismatches and **zero** name collisions.

### 6.1 What the suffixed runs found

A suffixed run finds what an unsuffixed one structurally cannot, and it found three defects
in this change plus one over-reach. All are fixed.

1. **Golden names written as list items were never tokenised.** The rewriter's YAML matcher
   wanted `key: value` at line start, but a golden's first key inside a list item reads
   `- imagename: x`. Two image references stayed unsuffixed
   (`app_profiles/create.yaml`, `deployment/create.yaml`).

2. **Two fixtures referenced their own image by a hardcoded name.**
   `application_instance/create.tf` and `patch_reference_update/create.tf` wire a drive with
   `imagename = "test_tf_provider..."` rather than `zedcloud_image.x.name`, so suffixing the
   image left the drive pointing at an image that does not exist.

3. **A nested manifest name was left behind.** `app_profiles/create.tf` repeats the app
   profile's own name inside `manifest_json`, which the controller stores and returns, so
   `VMManifest.Name` came back unsuffixed against a tokenised golden.

   The fix is deliberately narrow. Blanket-suffixing every `manifest_json.name` would also
   hit `deployment/create.tf`, whose manifest name (`tf-app-instance`) is a label of its own
   that appears nowhere else and comes back verbatim -- tokenising it would have broken a
   passing test. So a manifest's `name`/`title` is treated as a *reference*: it moves only
   when the same value is tokenised elsewhere in the same file. `TestAppProfile_Create` and
   `TestDeployment_Create` both pass under that rule.

4. **`nameAppPart` was tokenised but should not be.** `deployment/create.tf` sets
   `name_app_part` as a naming-scheme component, not as a reference to the app, so the
   controller echoes it unsuffixed -- the same trap as `description`. Both are in the golden
   pass's `DENY_KEYS`.

The rule these share: **a golden field may carry the token only when the fixture field it
reflects carries it too.** A sweep for every fixture key holding an untokenised copy of a
suffixed name found exactly four -- `description`, `name_app_part`, `image_rel_url` (a
datastore path, absent from the goldens) and `imagename` -- so the classification is believed
complete, established by reading the tree rather than by more 20-minute runs.

`TestGoldensTokeniseEveryFixtureName` closes the gap statically in the direction that can be
checked: for every name the fixtures tokenise, it fails on any golden asserting the same
value without the token, under any key, list item or not. It reproduces the misses when the
fix is reverted.

### 6.2 Acceptance criteria: met

Three full runs of `v2/resources` back to back on the resized cluster (§6.3), each with
**zero transport errors** and the node never leaving `Ready`:

| | AC1 (no suffix) | AC2 (`_ue139a`) | AC3 (`_ue139b`) |
|---|---|---|---|
| pass / fail / skip | 43 / 3 / 1 | 38 / 4 / 1 | 43 / 3 / 1 |
| elapsed | 18m49s | 47m34s | 17m16s |
| name collisions | 0 | 0 | 0 |

**AC1 — byte-identical with the suffix unset.** Proven mechanically first: all 80 changed
fixture and golden files, with `__SUFFIX__` stripped, are byte-identical to the base commit.
The suite then ran with the failure set below.

**AC2 — every created object carries the suffix.** Audited from the run's own API traffic:
**96 of 96** distinct object names carried `_ue139a`. AC3 reported 99 of 101, and both
exceptions are artefacts of the audit, not misses: `test-name` is a fixture literal
deliberately left alone, and `user2_ue139b@example.com` carries the token mid-string.

**AC3 — two suffixes back to back.** AC3 ran immediately after AC2 against the objects AC2
had just created and destroyed, with no collisions and the same pass set as AC1.

Every failure is accounted for, none in this change:

| Test | Verdict |
|---|---|
| `TestApplicationInstance_CreateCompose` | needs a real device -- the point of [UE-140](https://zededa.atlassian.net/browse/UE-140). Fails on the base commit too. |
| `TestProfileDeployment_Create` | the flake the design doc already calls out. Fails in all three runs and on base. |
| `TestPatchEnvelope_Create` | pre-existing flake: the base commit fails it **3 of 3** with the same `OpaqueObjectCategoryInline` vs `ExternalBinary` artifact diff -- the fixture declares one inline and one external artifact and the comparison pairs them by position. No name involved. Passed in AC3. |
| `TestAppProfile_Create` | was ours (finding 3). Passes with a suffix after the fix. |

### 6.3 The cluster had to be resized first

The first attempt at these runs measured the cluster rather than the change: during a
suffixed run the minikube node went `NodeNotReady` for ~23 minutes, taking every namespace
with it, and afterwards the API alternated between `503`, 10-second responses and timeouts.

Cause, from `docker stats`: the minikube container was at **22.39 GiB of a 23.44 GiB limit
(95.5%)** with CPU pinned at 808% of 8 cores, while the sum of pod memory *limits* was
22492Mi (~22.0 GiB) -- no headroom for what 70 acceptance tests add. Worse, the kubelet
reports node capacity as the **Docker VM's** 31.8 GiB rather than the container's cgroup
limit, so it never sees MemoryPressure and never evicts; the container simply walks into its
wall and the kernel kills pods inside it.

`tools/zedcloud-local/.env` in `sre-helm` now sets `MINIKUBE_RAM=28000` (from 24000) and
`MINIKUBE_CPUS=12` (from 8) -- the largest that still fits inside the existing 31.8 GiB
Docker VM with slack. After `minikube delete` + `make cloud-all`, the same suite runs at
**77-84% of 27.34 GiB** with CPU around 200% of 1200%, and all three runs above completed
with no node flap. The capacity-reporting gap is narrowed (4.5 GiB rather than 8.4) but not
closed: to close it, raise Docker Desktop's VM to ~40 GiB and `MINIKUBE_RAM` to 36000.

The suite stays sequential; no `t.Parallel()` (design §6.5).

---

## 7. Risks and open questions

- **Nested policy names stay unsuffixed** (`deployment/create.tf`, `project/create_with_pnac*.tf`).
  Assumed scoped to their parent object. If the controller enforces global policy-name
  uniqueness, AC3 fails on those two fixtures — the fix is one entry in `ALLOWED_PATHS`
  plus the corresponding golden lines. Worth confirming empirically during AC3 rather than
  reasoning about it. `project_test.go:367` already ignores `models.Policy.Name`, which
  hints the controller owns that name anyway.
- **Enterprise / white-labeling fixtures** (`iam/enterprise.*`) create cluster-scoped
  objects and are skipped in CI today. Suffix them for consistency, but they will not be
  exercised until [UE-140](https://zededa.atlassian.net/browse/UE-140) splits the `CI` gate
  into `ZEDCLOUD_ACC_ELEVATED`.
- **`title` suffixing is optional.** Titles carry no uniqueness constraint, and skipping
  them would halve the diff. Doing it anyway (design §6.1, and the `create_real_node.tf`
  precedent) makes every object attributable to a run in the UI, which is worth the churn
  the first time someone stares at 400 leftover objects on alpha.
- **`TestProfileDeployment_Create` is a known flake**, not a capability gap. If it fails
  during AC2/AC3, that is not this ticket. Do not "fix" it by widening a gate.
- **Out of scope:** `t.Parallel()`; `__NODE_SERIALNO__` / `__NODE_NAME__` / `__MODEL_ID__`
  (design §6.2 commit 2, belongs with UE-140); the `eth1` problem (design §6.4);
  the `zedcloud_edgenode` data source defect ([UE-143](https://zededa.atlassian.net/browse/UE-143)).

## 8. Follow-ups this left behind

1. **Tag-selector suffixing** (§4, departure 2). Needed before two suites can run
   concurrently; not needed for back-to-back, which AC3 demonstrates. Candidate ticket under
   UE-131.
2. **`ZEDCLOUD_TEST_SUFFIX` on the CI job** when [UE-141](https://zededa.atlassian.net/browse/UE-141)
   widens `e2e.yml` past the real-node test. The real-node step already sets it.
3. **`TestPatchEnvelope_Create` is a real flake** — the base commit fails it 3 of 3 on a
   healthy cluster, comparing two artifacts by position. Worth its own ticket; do not launder
   it behind a capability gate (the same argument the design doc makes for
   `TestProfileDeployment_Create`).
4. **Sweep leftovers** from the runs' failed destroys ([UE-142](https://zededa.atlassian.net/browse/UE-142)
   is the durable answer; until then, by suffix prefix).
5. **Close the kubelet capacity gap** if the cluster still wobbles: Docker Desktop VM to
   ~40 GiB and `MINIKUBE_RAM=36000` (§6.3).
