# E2E testing on virtual EVE nodes — final implementation plan

Epic: [UE-131](https://zededa.atlassian.net/browse/UE-131) — Terraform test pipeline on virtual EVE nodes
Status: **design settled, partially implemented.** This is the plan of record.

Supersedes the runner and transport decisions in
[`e2e-virtual-eve-node-testing.md`](./e2e-virtual-eve-node-testing.md); that document
remains the detailed design and validation log. The rejected network transports are
summarised in the decisions table in §2 below — the Headscale and Tailscale trial
runbooks were never committed.

---

## 1. What we are building

Give `terraform-provider-zedcloud` end-to-end test coverage against a **real, onboarded
EVE-OS edge node**, running as a QEMU VM on the Berlin lab host `h-m-dl20`, onboarded
into the alpha cluster, exercised on every pull request.

The gap today: all 41 acceptance tests fabricate a `zedcloud_edgenode` against a
synthetic brand/model with an invented serial. Nothing ever checks in, so an application
instance is "created" without ever running, and the drift/diff bugs that dominate recent
commit history (CI-723, CI-790, CI-865, CI-875) are exactly the class a real device would
catch earlier.

## 2. Architecture

```
  h-m-dl20 (Berlin lab, NixOS, 16 cores / 62 GB)
  ├── github-runner  (self-hosted, ephemeral, repo-scoped)
  │     └── job: tofu apply  ->  zedamigo (target = localhost)
  │                              ├── EVE installer  (podman, lfedge/eve)
  │                              ├── install VM     (QEMU, synchronous)
  │                              └── runtime VM     (QEMU, detached)
  │           go test          ->  acceptance suite, in-process provider
  └── /var/lib/zedamigo/reservations   (cooperative CPU/RAM claims)

  EVE VM  ──HTTPS──>  zedcloud.alpha.zededa.net      (device API, onboarding)
  runner  ──HTTPS──>  zedcontrol.alpha.zededa.net    (control API, tests)
```

Both cluster endpoints are public. **There is no private network hop**, which is what
makes this design simple: the runner is already inside the lab.

### Decisions and why

| Decision | Rationale |
|---|---|
| **Self-hosted runner on `h-m-dl20`** | Removes the network-path problem entirely. Also gives "one job at a time, in arrival order" for free — GitHub dispatches serially per runner label. The accounts, group, sudo rules and age-encrypted tokens already existed in the lab repo; only the services were missing |
| Rejected: hosted runner + tunnel/mesh | Evaluated Headscale on EC2 (built and validated, ~$22/mo + ops), Tailscale SaaS (~$12/mo, needs a tailnet), Cloudflare Tunnel ($0, already used internally). All solved a problem that disappears if the runner moves |
| `zedamigo` with `target = "localhost"` | Local executor — no SSH, no SFTP over the network, no `--build-host`. Faster and fewer moving parts |
| Cooperative capacity reservation | `zedamigo_host_reservation` against the shared tree, so CI cannot oversubscribe the host against a human's manual run |
| `zedamigo_wait_until` for readiness | Provider-native barrier; replaces the custom polling helper the earlier design proposed |
| Pinned EVE tag (`16.0.1-lts-kvm-amd64`) | A moving tag makes CI non-reproducible; a stale one tests nothing. Pin, with a scheduled bump |
| Per-run `run_id` suffix on every object | Alpha is shared (~233 devices, ~334 projects). Without this, concurrent or leftover runs collide on unique-name constraints |

## 3. Measured facts

Not estimates — observed during iterations 1 and 2.

| | |
|---|---|
| EVE installer (image cached) | 12s |
| Install VM (synchronous) | 1m34s |
| Boot → `RUN_STATE_ONLINE` | 1m21s (5 polls) |
| **Node lifecycle total** | **~3m10s** |
| Real-node app-instance test | 58s, including teardown |
| Host capacity | 16 cores, 62 GB RAM, 399 GB free, nested virt on |

The earlier 30–60 minute estimate was wrong by an order of magnitude — it came from a
macOS/arm64 emulated run. A per-PR trigger is comfortably viable and the "warm node pool"
idea is unnecessary.

## 4. The security constraint

`zededa/terraform-provider-zedcloud` is **PUBLIC**. A self-hosted runner on a public repo
lets any fork PR execute code on a host inside the lab network.

This is *strictly worse* than the hosted-runner design, for a reason worth stating
plainly: GitHub withholds repository secrets from fork `pull_request` workflows, so a
hosted runner hands a fork nothing useful. Here the attacker's code is already on the
target and needs no secret.

Mandatory, not optional — and **listed in order of what actually does the work**, which is
not the order you would guess:

1. **Approval for all external contributors.** GitHub → Settings → Actions → General →
   "Approval for running fork pull request workflows from contributors" → *Require
   approval for all external contributors*. This is the boundary: nothing from outside the
   `zededa` org runs until someone with write access clicks approve.
   Set on 2026-09-04; verify with
   `gh api repos/zededa/terraform-provider-zedcloud/actions/permissions/fork-pr-contributor-approval`,
   which must read `all_external_contributors`. The GitHub **default** is
   `first_time_contributors`, which exempts anyone who has ever had a commit merged — not
   sufficient here.
2. Runner registered **per repository**, never org-wide, so no other repo can target it.
3. `ephemeral = true` — fresh registration per job, so one job cannot poison the next.
4. `if: github.event.pull_request.head.repo.full_name == github.repository` in `e2e.yml`.

### Why the workflow `if:` gate is #4 and not #1

An earlier version of this section led with the `if:` gate and described the approval
requirement as an additional measure. That was backwards, and worth recording so the
mistake is not repeated.

For `pull_request` events GitHub runs the workflow **from the pull request's merge commit,
not from the base branch** — which is exactly why you can edit CI inside a PR and see the
change take effect in that same PR. A fork therefore ships its *own* copy of
`.github/workflows/e2e.yml` and can simply delete the `if:` line. It does not even need to
touch `e2e.yml`: adding `runs-on: [self-hosted, berlin-lab]` to `go.yml`, which has no
gate, reaches the same runner.

GitHub says this plainly: "forks of your public repository can potentially run dangerous
code on your self-hosted runner machine by creating a pull request that executes the code
in a workflow", and "potentially malicious user-controlled workflow code will execute
automatically if the user is allowed to bypass approval in the set approval policy."

So the gate is a guardrail, not a boundary. It keeps honest fork PRs from queueing against
a runner they cannot use, which is worth having for noise — but it stops nobody who does
not want to be stopped. Keep it; do not rely on it.

With control 1 in place the trust boundary is "who has push access" — roughly the same set
who could read repository secrets anyway. That is the argument for accepting the residual
risk. Without control 1, controls 2–4 do not establish any boundary at all.

Ref: [Secure use reference](https://docs.github.com/en/actions/reference/security/secure-use).

## 5. Work already done

| Item | Where |
|---|---|
| Node lifecycle Terraform config | `test/e2e/` (validated, node reached ONLINE) |
| Real-node test helpers | `v2/testing/realnode.go` |
| Real-node app-instance fixture + test | `v2/resources/testdata/application_instance/create_real_node.tf`, `application_instance_real_node_test.go` (passing) |
| Runner NixOS config | `zededa-berlin-lab/hosts/h-m-dl20/github-runners.nix` (deployed; runner online) |
| E2E workflow | `.github/workflows/e2e.yml` (dispatched on PR #228; fails at the toolchain guard — see T4) |
| Reservation `.lock` permission fix | in `github-runners.nix` via tmpfiles |
| Runner credential as a GitHub App | `githubApp` in `github-runners.nix` + `gh-run-tf-zedcloud-app-key.age` |
| Actions settings hardened | `all_external_contributors`; `GITHUB_TOKEN` default read-only |

Removed as superseded: `hosts/hs-aws-euc1/` (Headscale on EC2) and
`test/e2e/headscale-trial/`. Neither was ever pushed, so the removal is not visible in
either repo's history.

## 6. Work remaining — the ticket breakdown

Ordered by dependency. Each is independently reviewable and leaves the tree working.

| Ticket | Summary | Depends on | Status |
|---|---|---|---|
| [UE-135](https://zededa.atlassian.net/browse/UE-135) | T1 Runner credential (shipped as a GitHub App, **not** a PAT) | — | Done |
| [UE-136](https://zededa.atlassian.net/browse/UE-136) | T2 Deploy the runner to h-m-dl20 | UE-135 | Done |
| [UE-137](https://zededa.atlassian.net/browse/UE-137) | T3 Harden Actions settings (public repo) | with UE-136 | Controls set; fork test outstanding |
| [UE-138](https://zededa.atlassian.net/browse/UE-138) | T4 First green e2e run | UE-136, UE-137 | In progress — 3 tools missing on the runner |
| [UE-139](https://zededa.atlassian.net/browse/UE-139) | T5 `__SUFFIX__` across 45 fixtures | — (parallel) | |
| [UE-140](https://zededa.atlassian.net/browse/UE-140) | T6 Un-skip `CreateCompose` | UE-138 | |
| [UE-141](https://zededa.atlassian.net/browse/UE-141) | T7 Promote e2e to a required check | UE-138, UE-139 | |
| [UE-142](https://zededa.atlassian.net/browse/UE-142) | T8 Nightly janitor | UE-138 | |
| [UE-143](https://zededa.atlassian.net/browse/UE-143) | T9 Fix `zedcloud_edgenode` data source (bug) | — (independent) | |
| [UE-144](https://zededa.atlassian.net/browse/UE-144) | T10 Rotate committed credentials | — (do soon) | |

Critical path: UE-135 → UE-136/137 → UE-138 → UE-141.

### T1 — Provision the runner credential *(done — as a GitHub App, not a PAT)*
This ticket was planned as a fine-grained PAT and **shipped as a GitHub App instead**. The
reasoning is in `github-runners.nix` under "WHY NOT A PAT", and is worth repeating because
the PAT plan looks adequate until you count the years: with `ephemeral = true` the runner
re-registers after every job, so the credential has to still be valid months from now. A
registration token lasts one hour and a fine-grained PAT at most 366 days — both eventually
leave the service in a restart loop, failing in a way that reads like a runner bug rather
than an expired credential. A PAT merely moves that cliff a year out, onto a date nobody
will remember. An App private key does not expire, so the failure mode is designed out.
It is also not tied to a person: a PAT acts as its creator in the audit log and dies when
they are offboarded.

App id 4807245 (`zededa-berlin-lab-runner`), key in
`secrets/gh-run-tf-zedcloud-app-key.age`. `secrets/gh-run-tf-zedcloud.age` (the old PAT
path) is no longer read by the runner. Requires a nixpkgs newer than the 2026-08-10
revision the flake was pinned to — `services.github-runners.<name>.githubApp` exists in
neither 25.11 nor 26.05.

### T2 — Deploy the runner to `h-m-dl20` *(done)*
`nixos-rebuild switch` with `github-runners.nix`. Registered as `h-m-dl20-tf-zedcloud`
with labels `self-hosted, Linux, X64, berlin-lab, kvm, eve`; confirmed online and picking
up jobs.

### T3 — Harden the repo's Actions settings *(controls set; one test outstanding)*
Done: `all_external_contributors`, repo-scoped registration confirmed,
`default_workflow_permissions: read`, `can_approve_pull_request_reviews: false`.

Outstanding: the fork test, and it must be run as the *hostile* case to mean anything —
from a throwaway fork, delete the `if:` gate **and** point `go.yml` at
`runs-on: [self-hosted, berlin-lab]`, then confirm the run sits in "waiting for approval"
and nothing dispatches. An unmodified fork passes this test even with the approval policy
wide open, so testing that way proves nothing. See §4.

### T4 — First green e2e run
First dispatch (PR #228) reached the runner and stopped at `Verify host toolchain` with
`missing: docker`. Three tools are missing, all the same root cause — **systemd units do
not inherit `environment.systemPackages`, so anything the job shells out to must be in
`extraPackages`**:

| Tool | Why it is missing |
|---|---|
| `docker` | Comes from `virtualisation.podman.dockerCompat`, which writes the symlink into the *system* profile. Needs an explicit shim in `runnerTools`, e.g. `(pkgs.writeShellScriptBin "docker" ''exec ${pkgs.podman}/bin/podman "$@"'')` |
| `make` | Absent from `runnerTools` **and** `environment.systemPackages` — no fallback at all. `make build` fails with a bare `command not found` |
| a Terraform CLI | The plugin-SDK test driver checks `TF_ACC_TERRAFORM_PATH`, then searches for a binary named `terraform`, then **downloads an unpinned Terraform from `releases.hashicorp.com`**. OpenTofu is never discovered on its own, so the `tofu` entry in the toolchain guard did not cover the acceptance step at all |

The third one does **not** get fixed by installing `terraform`. nixpkgs' `terraform` is
BUSL-licensed and this flake does not allow unfree packages, so adding it would be a policy
change smuggled in as a CI fix. `e2e.yml` sets `TF_ACC_TERRAFORM_PATH` to the `tofu` on the
runner's PATH instead — the documented route, and `tofu version -json` still reports a
`terraform_version` key for exactly this compatibility. The toolchain guard deliberately
does not check for `terraform`, and a comment there says so, because the obvious "fix"
is the wrong one.

Rootless podman is still the *predicted* failure point beyond these — the EVE installer
runs `docker run` as a system user with no login session — but the guard fires before
reaching it, so that remains untested.

### T5 — `__SUFFIX__` rollout across existing fixtures
Add token expansion to the 45 `testdata/*.tf` files so every created object is per-run
unique. Provably a no-op at `__SUFFIX__=""`, so it lands safely before any semantic
change. Prerequisite for running the wider suite per-PR.

### T6 — Un-skip `TestApplicationInstance_CreateCompose`
Its skip comment says "until we decide how to provide a real device for testing" — that
decision is now made. Re-gate from `CI` to `ZEDCLOUD_ACC_REAL_NODE`.

### T7 — Promote e2e to a required check
The `pull_request` trigger this ticket originally described is **already live** — it landed
with the workflow itself, ahead of the first green run. What remains is the other half:
the job currently carries `continue-on-error: true`, so it reports success whatever
happens. Remove that once the job has been stable for several weeks, and only then add it
to branch protection as a required check.

Note that `main` presently requires **no** status checks at all — not `go.yml`, not
`test.yml`. Making e2e the first required check is a larger change in habit than it looks,
and probably wants the hosted checks required first.

### T8 — Nightly janitor
Sweep `lib_path` dirs and orphaned QEMU processes older than a few hours, and delete
`tf_e2e_%` / `test_tf%` cloud objects older than a day. `zedamigo_edge_node.Delete` treats
a failed QMP quit as a warning and removes state anyway, so orphans are an expected
outcome, not an exotic one.

This is also the backstop for a leak the workflow cannot close on its own. The acceptance
test's six alpha objects (project, datastore, image, network instance, application, app
instance) are destroyed by the SDK *in process*, at the end of `resource.Test`. A
`go test -timeout 40m` expiry, the job's own `timeout-minutes: 60`, or a cancellation all
kill that process before destroy runs, and the workflow's `Destroy` step only tears down
`test/e2e/`, not the test's objects. So `test_tf_real_*` leftovers are an expected outcome
on any abnormal exit, and the janitor must match that prefix as well as `tf_e2e_%`.

### T9 — Fix the `zedcloud_edgenode` data source (independent)
`NodeDataSource()` returns `Schema: zschema.Node()` — the *resource* schema — so it
inherits every write-side `Required` field and a name-only lookup fails with
`The argument "model_id" is required`. Data sources need a read-only schema variant. Found
while building T-earlier work; not blocking, but a real defect.

### T10 — Rotate committed credentials (independent, do soon)
Branch `add-tf-agent-skill` contains plaintext alpha and local-cluster passwords plus
Postgres credentials for both databases. Local-only in this clone; confirm it was never
pushed, then rotate regardless.

## 7. Open questions

1. Which EVE tag should CI track — the supported LTS, or newest? Pinned to `16.0.1` now.
2. Does alpha's owner accept per-PR churn (projects, models, brands, networks, images,
   devices created and deleted on every run)?
3. Should `h-m-dl20` also host the runner for `terraform-provider-zedamigo`? Account,
   token and a disabled service already exist.
4. The long-lived EVE node from iteration 1 is still running (2 QEMU processes, ~8 GB and
   a CPU reservation). Keep for manual iteration, or destroy now that CI creates its own?
