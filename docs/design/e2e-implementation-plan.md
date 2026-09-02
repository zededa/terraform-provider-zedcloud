# E2E testing on virtual EVE nodes — final implementation plan

Epic: [UE-131](https://zededa.atlassian.net/browse/UE-131) — Terraform test pipeline on virtual EVE nodes
Status: **design settled, partially implemented.** This is the plan of record.

Supersedes the runner and transport decisions in
[`e2e-virtual-eve-node-testing.md`](./e2e-virtual-eve-node-testing.md); that document
remains the detailed design and validation log. Rejected network transports are recorded
in [`tailscale-trial-runbook.md`](./tailscale-trial-runbook.md).

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

Mandatory, not optional:

1. `if: github.event.pull_request.head.repo.full_name == github.repository`
2. Runner registered **per repository**, never org-wide
3. `ephemeral = true` — fresh registration per job
4. GitHub → Settings → Actions → require approval for all outside collaborators

With these the trust boundary is "who has push access" — roughly the same set who could
read secrets anyway.

## 5. Work already done

| Item | Where |
|---|---|
| Node lifecycle Terraform config | `test/e2e/` (validated, node reached ONLINE) |
| Real-node test helpers | `v2/testing/realnode.go` |
| Real-node app-instance fixture + test | `v2/resources/testdata/application_instance/create_real_node.tf`, `application_instance_real_node_test.go` (passing) |
| Runner NixOS config | `zededa-berlin-lab/hosts/h-m-dl20/github-runners.nix` (validated by full `nix eval`) |
| E2E workflow | `.github/workflows/e2e.yml` (YAML validated; never executed) |
| Reservation `.lock` permission fix | in `github-runners.nix` via tmpfiles |

Removed as superseded: `hosts/hs-aws-euc1/` (Headscale on EC2) and
`test/e2e/headscale-trial/`.

## 6. Work remaining — the ticket breakdown

Ordered by dependency. Each is independently reviewable and leaves the tree working.

| Ticket | Summary | Depends on |
|---|---|---|
| [UE-135](https://zededa.atlassian.net/browse/UE-135) | T1 Runner token as a PAT | — |
| [UE-136](https://zededa.atlassian.net/browse/UE-136) | T2 Deploy the runner to h-m-dl20 | UE-135 |
| [UE-137](https://zededa.atlassian.net/browse/UE-137) | T3 Harden Actions settings (public repo) | with UE-136 |
| [UE-138](https://zededa.atlassian.net/browse/UE-138) | T4 First green e2e run | UE-136, UE-137 |
| [UE-139](https://zededa.atlassian.net/browse/UE-139) | T5 `__SUFFIX__` across 45 fixtures | — (parallel) |
| [UE-140](https://zededa.atlassian.net/browse/UE-140) | T6 Un-skip `CreateCompose` | UE-138 |
| [UE-141](https://zededa.atlassian.net/browse/UE-141) | T7 Enable `pull_request` trigger | UE-138, UE-139 |
| [UE-142](https://zededa.atlassian.net/browse/UE-142) | T8 Nightly janitor | UE-138 |
| [UE-143](https://zededa.atlassian.net/browse/UE-143) | T9 Fix `zedcloud_edgenode` data source (bug) | — (independent) |
| [UE-144](https://zededa.atlassian.net/browse/UE-144) | T10 Rotate committed credentials | — (do soon) |

Critical path: UE-135 → UE-136/137 → UE-138 → UE-141.

### T1 — Provision the runner token as a PAT
Confirm what `secrets/gh-run-tf-zedcloud.age` holds. `ephemeral = true` re-registers after
every job, and **registration tokens expire after one hour**, so a registration token
would break the second job of the day and leave the service restart-looping. Needs a
fine-grained PAT scoped to the repo with `Administration: read and write`.
*Blocks everything else.*

### T2 — Deploy the runner to `h-m-dl20`
`nixos-rebuild switch` with `github-runners.nix`. Confirm the runner appears in the repo's
runner list with labels `self-hosted, berlin-lab, kvm, eve`. Requires the current flake
from andrei — the local clone is behind, and rebuilding from a stale checkout would revert
his work.

### T3 — Harden the repo's Actions settings
Require approval for outside collaborators; confirm the runner is repo-scoped. Pairs with
the workflow's fork gate.

### T4 — First green e2e run
Trigger via `workflow_dispatch`, iterate on what the sandboxed unit actually permits.
Rootless podman is the most likely failure point: the EVE installer runs `docker run` as a
system user with no login session.

### T5 — `__SUFFIX__` rollout across existing fixtures
Add token expansion to the 45 `testdata/*.tf` files so every created object is per-run
unique. Provably a no-op at `__SUFFIX__=""`, so it lands safely before any semantic
change. Prerequisite for running the wider suite per-PR.

### T6 — Un-skip `TestApplicationInstance_CreateCompose`
Its skip comment says "until we decide how to provide a real device for testing" — that
decision is now made. Re-gate from `CI` to `ZEDCLOUD_ACC_REAL_NODE`.

### T7 — Enable the PR trigger
Flip `e2e.yml` from `workflow_dispatch`-only to `pull_request`. Keep the check
**non-required** in branch protection until it has been stable for several weeks.

### T8 — Nightly janitor
Sweep `lib_path` dirs and orphaned QEMU processes older than a few hours, and delete
`tf_e2e_%` / `test_tf%` cloud objects older than a day. `zedamigo_edge_node.Delete` treats
a failed QMP quit as a warning and removes state anyway, so orphans are an expected
outcome, not an exotic one.

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
