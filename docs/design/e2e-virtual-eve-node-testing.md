# E2E acceptance testing against virtual EVE nodes in CI

Status: draft plan — **iteration 1 implemented and validated 2026-08-19** (see §0)
Ticket: TBD
Repos: `zededa/terraform-provider-zedcloud`, `andrei-zededa/terraform-provider-zedamigo`, `andrei-zededa/zededa-berlin-lab`

---

## 0. Validation log — what iteration 1 proved, and what it corrected

A virtual EVE node was created on `h-m-dl20` and onboarded to alpha end-to-end. The
config lives in [`test/e2e/`](../../test/e2e/); see its README for the exact procedure.
Result: `ADMIN_STATE_REGISTERED` / `RUN_STATE_ONLINE`, `install_success = true`.

Six assumptions in this document turned out to be wrong. They are corrected in place
below; collected here so reviewers of the earlier draft can see the delta.

| § | Assumption in the draft | Reality |
|---|---|---|
| 3.3 | `zedcontrold.alpha.zededa.net` might be the device endpoint | **It does not resolve.** The pair is `zedcontrol.alpha` (control) and `zedcloud.alpha` (device), both Cloudflare-fronted with a Google Trust Services cert — so **no `tls_ca` and no `additional_hosts` needed** |
| 5.4 | zedamigo has no wait-for-online primitive; build a Go `waitonline` helper | **`zedamigo_wait_until` exists** (v0.13.0+). Runs the probe on the target, so it works unchanged for local and remote targets. The custom helper is unnecessary |
| 7.2 | Serialization needs a bespoke `lab-lock.sh` ticket lock | **`zedamigo_host_reservation` exists** (v0.13.0+) and the capacity tree is already deployed on `h-m-dl20` (16 CPU slots, 58 GB slots, LVM devices), flock-atomic and cross-config. It is admission control, not a queue — so it complements rather than replaces the FIFO design, but the custom lock should be built *on top of* it, not instead of it |
| 4.2 | Nested virtualization needs enabling | **Already enabled** (`kvm_intel nested=Y`). The host has drifted ahead of the `zededa-berlin-lab` checkout (host rev `21a8f219`, clone `48adfea`) |
| 4.5 | Host capacity unknown | **16 cores, 62 GB RAM, 399 GB free on `/`.** `/dev/kvm` is `0666`, so no group membership needed |
| 10.3 | 30–60 min per run, ~1 PR/hour | **~3m10s** total (installer 12s cached, install VM 1m34s, boot→ONLINE 1m21s). Roughly an order of magnitude cheaper than assumed. Per-PR is comfortably viable and the warm-node pool (§3.4B) is probably unnecessary |

Also settled: **the onboarding certificate is a non-issue on alpha.** The 403
`register` → "device not found" blocker documented against the local cluster
(§10.2) did not occur; the default key `5d0767ee-…` is authorized. This was the
highest-risk item in the plan (task 0.2) and it is now closed.

Two new findings:

- **The reservation lock is single-user.** `/var/lib/zedamigo/reservations/.lock` is
  mode `0644` owned by `andrei` inside a group-writable directory, and the provider
  opens it for writing — so cooperative reservation works for exactly one user and
  everyone else gets `Permission denied`. Needs `0664`/`g+w`, declared in the lab's
  `configuration.nix` so it survives a `nixos-rebuild`. Until then `test/e2e/` uses a
  per-user tree with no cross-user coordination.
- **Two GitHub Actions service accounts already exist** on the host —
  `gh-run-tf-zedamigo` and `gh-run-tf-zedcloud`, both `nologin` with `/var/empty`
  homes, in the `zedamigo_reservations` group. No runner service is registered yet.
  Whatever CI identity this design lands on should use `gh-run-tf-zedcloud` rather
  than inventing another.

What iteration 1 did **not** cover: the SSH remote-target topology (OpenTofu ran on
the lab host, `target = localhost`) and the network path from a GitHub-hosted runner
(§3.2 is still entirely open).

### Iteration 3 — runner decision reversed: self-hosted on the lab host

**This supersedes the runner choice in §3.1 and most of §3.2 and §7.2.** The original
decision was a GitHub-hosted runner reaching the lab over SSH, which made the network
path the central problem. Three transports were evaluated for it — a self-hosted
Headscale on EC2 (built, validated, ~$22/mo), Tailscale SaaS (~$12/mo, needs a
tailnet), and Cloudflare Tunnel ($0, already used internally) — before the simpler
question surfaced: put the runner *on* the lab host and there is no network path to
solve at all.

Decisive evidence that this was always the intent: `hosts/h-m-dl20/configuration.nix`
already contains `gh-run-tf-zedcloud` and `gh-run-tf-zedamigo` service accounts in a
`gh-runners` group, with restricted passwordless sudo for exactly `ip`/`kill`/`taskset`
(zedamigo's set), membership in `zedamigo_reservations`, and committed
age-encrypted tokens (`d39af16 feat: Add secrets for Github Runners`). Only the
services were missing. Confirmed no runner was registered (`runners=0`).

What this removes:

| Was needed | Now |
|---|---|
| A network path from GitHub's egress (§3.2) | Gone — the runner is already inside |
| `hosts/hs-aws-euc1/` (EC2, EIP, EBS, Route53, nginx, ACME) | Unnecessary |
| The FIFO ticket lock, `lab-lock.sh` (§7.2) | Unnecessary — GitHub dispatches jobs to a runner label serially, in arrival order |
| SSH, SFTP, `known_hosts`, `--build-host` | Gone — `target = "localhost"`, local executor |
| ~$22/month | $0 |

Implemented: `hosts/h-m-dl20/github-runners.nix` in `zededa-berlin-lab` (validated by
a full `nix eval` of the system closure, not just a parse) and
`.github/workflows/e2e.yml` here.

**The one new risk, and it is serious.** This repository is **public**
(`gh repo view` → `PUBLIC`). A self-hosted runner on a public repo lets any fork PR
execute code on the lab host. That is *strictly worse* than the hosted-runner design,
for a reason worth stating plainly: GitHub withholds repository secrets from fork
`pull_request` workflows, so a hosted runner would hand a fork no credentials and it
could reach nothing — whereas here the attacker's code is already on the target and
needs no secret at all.

The mitigation is in the workflow and is not optional:

```yaml
if: github.event.pull_request.head.repo.full_name == github.repository
```

plus repo-scoped (not org-scoped) registration, `ephemeral = true`, and GitHub's
"require approval for all outside collaborators". With those, the trust boundary is
"who has push access" — roughly the same set who could read secrets anyway.

Also fixed here: the reservation `.lock` permission bug from §0 is now declarative
(`systemd.tmpfiles` declares it `0664 root:zedamigo_reservations`), so `test/e2e/`
uses the shared capacity tree again rather than a per-user workaround.

Still open: §6's `__SUFFIX__` rollout across the 45 existing fixtures, and whether
`h-m-dl20` should host the runner for `terraform-provider-zedamigo` too (the account
and token exist; the service is declared but disabled).

### Iteration 2 — first real-node acceptance test (2026-08-19)

`TestApplicationInstance_RealNode` deploys an nginx container onto the live node and
waits until the device reports it running. **Passes in 58s**, including teardown:

```
real node "tf_e2e_en_it1" resolved to 723397fb-... (RUN_STATE_ONLINE)
attempt 1: swState=SW_STATE_UNSPECIFIED runState=RUN_STATE_UNKNOWN
attempt 2: swState=SW_STATE_UNSPECIFIED runState=RUN_STATE_UNKNOWN
attempt 3: swState=SW_STATE_RUNNING     runState=RUN_STATE_ONLINE
--- PASS: TestApplicationInstance_RealNode (58.47s)
```

New code: `v2/testing/realnode.go` (env contract, name→UUID resolution, placeholder
expansion, `WaitForAppInstanceSwState`), the fixture
`v2/resources/testdata/application_instance/create_real_node.tf`, and the test
`v2/resources/application_instance_real_node_test.go`. Teardown verified clean —
zero `test_tf_real*` leftovers across instances, apps, projects, images, datastores
and network instances.

This validates the §6.3 capability-gating design: `ZEDCLOUD_ACC_REAL_NODE` plus
`ZEDCLOUD_TEST_NODE_NAME`/`ZEDCLOUD_TEST_SUFFIX`, skipping when unset so the default
suite is untouched. It also means §6.5's 90-minute timeout is generous: a real-node
test costs ~1 minute, not tens.

Two further corrections:

| § | Assumption | Reality |
|---|---|---|
| 6.1 | Inject the node via `__NODE_ID__` **or** the `zedcloud_edgenode` data source | The **data source cannot be used at all** — see below. Node identity is injected as `__NODE_ID__`, resolved from a *name* by the helper so callers still pass a name |
| 5.4 | The wait helper should live in `v2/testing/cmd/waitonline` for the workflow | Unnecessary — `zedamigo_wait_until` covers the workflow side and `WaitForAppInstanceSwState` covers the test side. No standalone command needed |

**New defect found: `zedcloud_edgenode` data source is unusable.** `NodeDataSource()`
(`v2/resources/node.go:33`) returns `Schema: zschema.Node()` — the *resource* schema —
so the data source inherits every write-side `Required` field. A name-only lookup
fails with `The argument "model_id" is required`, plus `project_id`, `title` and
"At least 1 interfaces blocks are required". Data sources need a read-only schema
variant. Worth a small standalone PR; it is not blocking, since the helper resolves
names against the API.

Five fixture-level API constraints, all discovered empirically and now recorded in
comments in the fixture (none are documented anywhere):

1. Container image `name` must be a plain identifier — `/` and `:` are rejected with
   "Name field contains invalid characters". The pull reference goes in
   **`image_rel_url`** (`"library/nginx:stable-alpine"`), and downstream references
   (`images.imagename`, `drives.imagename`) use the image *name*.
2. `vminfo.memory` is Computed-only; setting it fails with "Value for unconfigurable
   attribute". Memory comes from the manifest's `resources` block.
3. App-instance `interfaces` requires `privateip` alongside `intfname`/`netinstname`.
4. An application's `project_access_list` must match its images' scope, else
   "app project list cannot be set to all projects if all of it's images do not set
   their project list to all projects".
5. `datastore_id_list` on an application is compose-only — "datastore refs are only
   supported for docker compose type app".

Remaining from §6: the `__SUFFIX__` rollout across the existing 45 fixtures, and
un-skipping `TestApplicationInstance_CreateCompose`. The scaffolding they need now
exists.

---

## 1. Goal

Replace the current PR test job — which runs acceptance tests against a controller with
**no real devices** — with one that:

1. Creates one or more **virtual EVE-OS edge nodes** as QEMU/KVM VMs on the Berlin lab host
   `h-m-dl20`.
2. Onboards them into the **alpha** Zedcloud cluster.
3. Runs the **full** `v2/resources` acceptance suite against those real, online devices.
4. Tears everything down, on every pull request.

Decisions already taken (see §11 for the ones still open):

| Decision | Choice |
|---|---|
| Runner | **GitHub-hosted** (`ubuntu-latest`); the zedamigo provider drives `h-m-dl20` over SSH |
| Trigger | **Every pull request** |
| Scope | **Full real-node acceptance suite**, not just a smoke test |
| Capacity | **One run at a time** on the lab host; queued runs served **first-come-first-served** (§7.2) |
| Lifecycle | Virtual EVE nodes are **destroyed after every run**, including on cancellation (§8) |

---

## 2. Summary of findings

### 2.1 What already exists — more than expected

**The zedamigo provider already supports exactly the topology we chose.** Commits
`810ff20` ("Support running against a remote host over SSH") and `5372a17` ("Add SSH jump host
(ProxyJump) support") added an `exec.Executor` abstraction with an `SSHExecutor` implementation.
The provider plugin process runs on the runner; every command, file operation, process kill and
socket dial happens on the target over SSH + SFTP. All paths in the config are *target* paths.
Configuration is `provider "zedamigo" { target = "..." ssh { ... } }`, and every SSH
attribute has an environment fallback — 11 of the 12 as `ZEDAMIGO_SSH_*`, plus
`ZEDAMIGO_REMOTE_BINARY_PATH` for `remote_binary_path` — purpose-built for CI secrets
(`internal/provider/ssh_config.go`).

**The lab host is already provisioned for a CI identity.** `zededa-berlin-lab`
`hosts/h-m-dl20/configuration.nix` defines a `github-runner` user that is deliberately *not* in
`wheel`, is in the `kvm` group, has `autoSubUidGidRange` for rootless podman, and has exactly
three passwordless sudo commands:

```nix
{ command = "/run/current-system/sw/bin/ip";      options = [ "NOPASSWD" ]; }
{ command = "/run/current-system/sw/bin/kill";    options = [ "NOPASSWD" ]; }
{ command = "/run/current-system/sw/bin/taskset"; options = [ "NOPASSWD" ]; }
```

That list is not arbitrary — it covers most of what zedamigo invokes under sudo: `ip`
(`netns_resource.go`, `tap_resource.go`, and bridge/vlan/lag via
`netns_helpers.go:buildIPCommand`), `kill` (`internal/exec/ssh.go:648-653`) and `taskset` (the
`cpu_pins` path, which uses plain `sudo`, not `sudo -n` —
`internal/hypervisor/qemu_cpupin_linux.go:110-112`). Someone clearly worked out a
least-privilege set for this use case.

It is **not complete**, though: zedamigo also self-invokes its own binary under `sudo -n` for
the embedded daemons — `tap_resource.go:578-580` (tap-mover),
`dhcp_server_resource.go:520-523`, `dhcp6_server_resource.go:459-461`,
`radv_resource.go:503-505`. Those are only reached if we use the host-networking resources
(§6.4), so phase 1 is unaffected, but phase 2 needs a fourth sudoers entry for the
bootstrapped binary path. The `github-runner` SSH key is still the placeholder
`ssh-ed25519 AAAA...REPLACE_ME`.

**There is a proven, debugged reference config.** `andrei-zededa/local-en/` is a complete
two-provider (zedamigo + zedcloud) stack that builds an EVE installer, installs it, boots the
VM and onboards it — with `CLAUDE.md` and `RUNBOOK.md` documenting every failure mode hit along
the way. It targets macOS/vfkit against a *local* cluster, so it is not directly reusable, but
its findings are (private TLS CA, `additional_hosts`, the `depends_on` ordering, the onboarding
cert blocker, teardown recipe). This is the single most valuable input to this work.

**Alpha access already exists.** Branch `add-tf-agent-skill` in this repo carries
`.agents/skills/terraform/scripts/set_api_token.sh`, which logs into
`zedcontrol.alpha.zededa.net` as `ivan-icd-test@zededa.com` (enterprise
`AAFlABBtJ0mP_lhJjIyVzjzgMXiR`) and mints a 90-day API token, and
`clean_tf_objects_on_alpha.sh`, which sweeps `test_tf%`-named leftovers straight out of the
alpha Postgres. Both are useful; both are also a problem — see §10.1.

**The Berlin lab is not on the public internet.** `h-m-dl20` is `10.208.13.94`. The documented
access path is the Berlin office L2TP/IPsec VPN, wrapped in
`zededa-berlin-lab/berlin-vpn-in-a-container/` (strongswan + xl2tpd + pppd + squid + an sshd on
port 11022 used as a jump host). Every documented remote command in the lab README goes through
`ProxyJump=root@localhost:11022` (the one exception is the "rebuild on the host itself"
alternative, which assumes you are already inside the lab network). **This is the load-bearing
unknown for a GitHub-hosted runner** — see §3.2.

### 2.2 Not related, despite the name

The `tf-actions` and `tf-actions-hackaton` branches in this repo are about **Terraform provider
Actions** (`provider.ProviderWithActions`, `v2/actions/node/RebootAction.go`), i.e. the
plugin-framework feature, not GitHub Actions. They also migrate `v2/main.go` from SDKv2 to
plugin-framework. Unrelated to this work, but the name collision will confuse people in
review — worth saying so out loud in the PR description.

### 2.3 The four hard constraints

Everything below follows from these.

**(a) The acceptance suite has no concept of a real device.** There are 41 acceptance test
functions in `v2/resources/*_test.go`; 19 of the 45 fixtures create their own
`zedcloud_edgenode` against a synthetic
`zedcloud_brand` + `zedcloud_model` (`origin_type = "ORIGIN_LOCAL"`), with invented serials
(`"def-netinst"`, `"2293dbe8-29ce-420c-8264-962857efc46b"`) and `onboarding_key = ""`. There
is **zero** run-state awareness: grepping `v2/resources/` for `RunState`, `WaitForState`,
`StateChangeConf` or `PollInterval` returns nothing. Nothing waits for a device to check in;
nothing asserts an app instance actually runs.

Consequence: pointing this suite at a real node changes *nothing* unless the fixtures are
rewritten to reference that node. "Full real-node suite" is mostly a **test-suite refactor**,
not a CI plumbing job. Budget accordingly.

**(b) Fixtures are static HCL with hardcoded, colliding names.** `v2/resources/testdata/*.tf`
are read verbatim by `testhelper.MustGetTestInput` and handed to `TestStep.Config` — no
templating, no `variable` blocks, no `fmt.Sprintf`, no env reads. `name =
"test_tf_provider-create_edgenode"` appears **34 times**; `serialno =
"2293dbe8-29ce-420c-8264-962857efc46b"` is shared by `application_instance/create.tf` and
`patch_reference_update/create.tf`. There is no `t.Parallel()` anywhere, so the suite is saved
today only by Go's sequential default.

Consequence: two concurrent runs against the same enterprise **will** collide on unique-name
constraints. On a per-PR trigger this is not hypothetical.

**(c) Six tests are gated on `CI`, one of them for exactly our reason.**
`application_instance_test.go:65-71`:

```go
// TestApplicationInstance_CreateCompose is a test case for creating an ApplicationInstance using a Docker Compose file.
// It's skipped until we decide how to provide a real device for testing.

func TestApplicationInstance_CreateCompose(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping docker compose test for CI environment")
	}
```

Also skipped on `CI`: `TestProfileDeployment_Create` (an acknowledged flake —
*"failing in CI but passing locally"*) and four `TestEnterprise_*` tests
(`TestEnterprise_Create`, `_WhiteLabeling`, `_WhiteLabelingControllerHostURL`,
`_WhiteLabelingLegacyAttributes`) whose reason is **undocumented** — the skip message says only
*"Skipping enterprise test for CI environment"*; a permissions gap on child-enterprise creation
is the likely cause but needs confirming. The current workflow sets `CI: true`, so all six are
dead today. `CI` is the wrong axis — it conflates "runs in automation" with "has a real device"
and with "has elevated permissions".

**(d) Timings do not fit a 20-minute budget.** `make test` uses `-timeout 20m` for the whole
package. Add EVE install (a synchronous, *unbounded* QEMU run inside
`zedamigo_installed_edge_node`), boot, onboarding and a wait-for-online poll, and the realistic
end-to-end is 30-60 minutes. `local-en/RUNBOOK.md` records ~5-10 min just for
installer-build → install VM → runtime VM, before onboarding.

---

## 3. Proposed architecture

### 3.1 Topology

```
GitHub-hosted runner (ubuntu-latest, ephemeral)
├── go build ./v2                      -> terraform-provider-zedcloud (under test)
├── terraform / tofu                   -> runs the e2e config
│    ├── provider "zedcloud"  ──────── HTTPS ───────▶ zedcontrol.alpha.zededa.net  (control API)
│    └── provider "zedamigo"  ── SSH (see §3.2) ───▶ github-runner@10.208.13.94
│                                                     └── QEMU/KVM: EVE-OS VM
│                                                          │ SLIRP NAT, host uplink
│                                                          └── HTTPS ▶ <device API>.alpha.zededa.net
└── go test ./v2/resources/...         -> acceptance suite, real node injected via env
```

Key property: **the provider binary under test runs on the runner**, so the tests exercise the
build from the PR. The zedamigo provider also runs on the runner and only reaches into the lab
over SSH. Nothing needs to be installed on `h-m-dl20` per run except the zedamigo helper binary
(auto-bootstrapped, §4.3).

### 3.2 The network path — decide this first

A GitHub-hosted runner cannot reach `10.208.13.94`. Three options, in order of preference:

**Option 1 (recommended): Tailscale subnet router on `h-m-dl20`.**
Add `services.tailscale` to `configuration.nix` and advertise `10.208.13.0/24` (or just
`10.208.13.94/32`). In the workflow, `tailscale/github-action@v3` with an ephemeral, tagged
auth key or OIDC. Userspace WireGuard, no kernel modules, no privileged container, no VPN
credentials in secrets, works on every hosted runner, and gives a stable name
(`h-m-dl20.<tailnet>.ts.net`) so `known_hosts` pinning is trivial. Cost: a new dependency and
an ACL/tag policy to agree with IT.

**Option 2: run `berlin-vpn-in-a-container` on the runner.** It already exists and is
documented. But it needs `--privileged`, `--cap-add SYS_MODULE` and a `/lib/modules` bind mount
to bring up IPsec/L2TP — on Azure-backed hosted runners that is unreliable at best, and the
credentials (`ZEDEDA_VPN_PSK`/`USER`/`PASS`) become repo secrets used on every PR. Treat as a
fallback and **spike it before committing** (a 15-line throwaway workflow answers the question).

**Option 3: publish a narrow SSH ingress.** A Cloudflare Tunnel / `cloudflared` or an
autossh reverse tunnel from `h-m-dl20` to a small public jump host, restricted to the
`github-runner` key. More moving parts than Option 1 with fewer benefits.

If none of the three clears review, the self-hosted-runner topology (register the runner *as*
`github-runner` on `h-m-dl20`) sidesteps the problem entirely and should be reconsidered — it
is also what the existing `github-runner` user and its sudo rules look designed for.

Whichever is chosen, the zedamigo provider config stays the same; only how the runner resolves
and routes to the host changes:

```hcl
provider "zedamigo" {
  target   = var.lab_host              # ZEDAMIGO_TARGET
  use_sudo = true                      # needed once we add eth1 (§6.4); harmless before that
  lib_path = "/home/github-runner/.local/state/zedamigo-gh-${var.run_id}"

  ssh {
    user             = "github-runner"
    private_key      = var.lab_ssh_key   # ZEDAMIGO_SSH_PRIVATE_KEY, from secrets
    known_hosts_file = "./known_hosts"   # written in a workflow step; do NOT use insecure mode
    # proxy_jump = var.lab_jump_host     # only if Option 2/3
  }
}
```

Note: a hosted runner has no `~/.ssh/known_hosts`, and `ssh_config.go:300` **fails closed** if
neither `known_hosts_file`, `host_key` nor `insecure_ignore_host_key` is set. Write the host key
into the workspace from a repo variable (not a secret — host public keys are not secret) and
point `known_hosts_file` at it.

### 3.3 Endpoint naming — resolve before writing HCL

Zedcloud exposes **two different hostnames** and mixing them up is the classic failure:

| Purpose | Consumer | Value on alpha |
|---|---|---|
| Control / REST API | `provider "zedcloud" { zedcloud_url }`, `TF_VAR_zedcloud_url` | `zedcontrol.alpha.zededa.net` (confirmed by `set_api_token.sh`) |
| Device API (`/config/server` inside EVE) | `zedamigo_eve_installer.cluster` | **TBD — see below** |

The request names `zedcontrold.alpha.zededa.net`. That is neither of the two names seen in the
codebase: the control API is `zedcontrol.alpha...` and the device-facing convention in
`local-en` is `zedcloud.<env>.zededa.net`. `zedcontrold` (with a `d`, plausibly "device") may
well be the correct alpha device endpoint, but it is unverified. **Confirm with the cluster
owners before writing the config, and keep the two as separate Terraform variables** — the
top-level zedamigo `README.md:85-88` and `examples/other/1_node_3_vlans/README.md` contradict
each other on this exact point, and `local-en/CLAUDE.md:46` resolves it in favour of two
distinct values.

### 3.4 Warm pool vs. create-per-PR

"Every PR" plus "one shared lab host" plus "30-60 min per run" is a queueing problem. Two
shapes:

**(A) Create and destroy per PR.** Simplest, fully hermetic, no state to persist beyond the
job. But every PR pays the full install+onboard cost, and concurrent PRs serialize behind
`concurrency:`. On a busy day PR feedback is hours.

**(B) Warm node + lease (recommended for per-PR).** A scheduled workflow keeps a small number of
onboarded nodes alive on `h-m-dl20`, re-created nightly. A PR job *leases* one, runs the suite
against it inside a per-run project, and releases it. Node lifecycle cost disappears from the PR
critical path; PR runtime drops to roughly the suite runtime.

Because the lab is a single-slot resource (§7.2), the pool is **N=1** — and that is the point:
one node alive continuously costs the same RAM as one active run, but every PR after the first
skips the installer build, install VM and onboarding wait entirely. The saving is most of the
critical path, not a marginal one. The trade-off is that state accumulated by a previous PR's
suite can leak into the next one, so per-run isolation (§5.3) has to be genuinely airtight
before this is safe — which is another reason to build (A) first.

The test-suite work in §6 is identical either way — the node identity arrives via environment
variables regardless of who created it. **Build (A) first** because it is the only way to prove
the node lifecycle works; add (B) once green. Do not design (A) in a way that blocks (B): keep
"create the node" and "run the suite" in separate jobs with the node's identity passed as job
outputs.

---

## 4. Host preparation — changes to `zededa-berlin-lab`

All in `hosts/h-m-dl20/configuration.nix` unless noted. These are the concrete gaps between
what the host provides today and what zedamigo requires.

### 4.1 Replace the placeholder SSH key

```nix
github-runner = [ "ssh-ed25519 AAAA<real CI key> github-runner" ];
```

Generate a dedicated ed25519 keypair for this purpose only. Private key → repo secret
`LAB_SSH_PRIVATE_KEY`; host's public key → repo variable `LAB_SSH_HOST_KEY`.

### 4.2 Enable nested virtualization

`hardware-configuration.nix` loads `kvm-intel`, but nesting is not enabled. Without it, EVE can
boot and onboard, but **cannot run app instances** — which kills
`TestApplicationInstance_*`, the most valuable real-node tests.

```nix
boot.extraModprobeConfig = "options kvm_intel nested=1";
```

Verify after rebuild: `cat /sys/module/kvm_intel/parameters/nested` → `Y`.

### 4.3 Tool inventory

zedamigo's `Configure()` hard-fails if any of these is missing from `PATH`, and warns for the
rest:

| Tool | Required? | Present on `h-m-dl20`? |
|---|---|---|
| `bash` | fatal (`Configure`, `provider.go:280`) | yes |
| `qemu-system-x86_64`, `qemu-img` | fatal | yes (`qemu`) |
| `ip` | fatal | yes (`iproute2`) |
| `docker` | fatal | yes — via `virtualisation.podman.dockerCompat`, which symlinks `docker` → `podman` |
| `sudo` | fatal when `use_sudo = true` | yes |
| `curl`, `grep`, `sed`, `mktemp`, `unzip`, `tar` | needed by `install.sh:22` (not checked by `Configure`) | yes |
| `taskset` | warn | yes (via the NixOS base system, not an explicit `systemPackages` entry) |
| `swtpm` | warn | **no** — fine, we do not use `zedamigo_swtpm` (it is WIP anyway) |
| `genisoimage` | warn | **no** — fine, we do not use `zedamigo_cloud_init_iso` |
| OVMF firmware | not needed | embedded in the provider binary, extracted over SFTP to `<lib_path>/embedded_ovmf/` on the target — note this lives inside the per-run `lib_path` that §5.3/§8 propose wiping, so it is re-extracted each run |

Rootless podman is fine for what we need: `zedamigo_eve_installer` runs
`docker run --network none --rm -v <dir>/config:/in -v <dir>/out:/out docker.io/lfedge/eve:<tag> ...`
(fully qualified — `eve_installer_resource.go:294` — which matters on podman, since it does not
default to Docker Hub),
which works rootless. No `docker` group is needed and no podman socket has to be enabled —
zedamigo shells out to the CLI, not the API socket.

### 4.4 Pre-pull the EVE image

The `lfedge/eve:<tag>` pull is multi-GB and dominates `zedamigo_eve_installer` on a cold cache.
Add a small systemd timer (or a step in the nightly pool workflow) that keeps the pinned tag
warm for the `github-runner` user, and **pin the tag in a variable** so it is not silently
re-pulled per run.

### 4.5 Disk and capacity

qcow2 images live under `~/.local/state/zedamigo/...` on `/` (the NVMe root). The LVM
`reserve*` LVs on `vg_sdb`/`vg_sdc`/`vg_sdd` are unformatted and only usable as raw block
devices via `disk { type = "device" }` — not what we want. Confirm NVMe free space: each node
needs a 20 GB sparse overlay plus the installer artefacts; a warm pool of 3 plus a PR run is
comfortable but not free.

The install VM is hard-coded to `-m 4096 -smp 4,cores=2` and the runtime VM defaults to
`4G`/4 vCPU, so budget ~4-8 GB and 4 vCPU per node. **The host's actual core count and RAM are
not recorded in the repo** — measure it (§11), but note that the single-slot constraint (§7.2)
means only ever one run's worth of VMs is live, so the sizing question is "can it host one node
comfortably", not "how many in parallel". Since the install and runtime VMs briefly overlap
during `apply`, size for two.

### 4.6 Monitoring is already there

`monitoring.nix` runs VictoriaMetrics + VictoriaLogs + node_exporter + a `process_exporter`
whose match rules group by the QEMU `-name` argument. That gives per-VM CPU/RSS from
`process_exporter`, and per-VM disk IO from node_exporter's `diskstats` on each VM's
device-mapper node — for free. Worth wiring a dashboard link into the workflow summary; CI
flakes on this host will mostly be resource flakes, and this is how they get diagnosed. Note the
stack is loopback-only, so reaching it needs the same SSH tunnel as everything else.

---

## 5. The e2e Terraform config

New directory in **this** repo: `test/e2e/` (a sibling of the existing `test/retry/`), so the
config is versioned with the provider it tests.

```
test/e2e/
├── terraform.tf        # required_providers + both provider blocks
├── vars.tf             # lab_host, ssh key, cluster URLs, eve_tag, run_id, node_count
├── cluster_objects.tf  # zedcloud_brand / model / project / network / edgenode
├── node.tf             # zedamigo disk_image / eve_installer / installed_edge_node / edge_node
├── outputs.tf          # serials, device UUIDs, project id, model id, ssh_port
└── README.md
```

Base it on `terraform-provider-zedamigo/examples/other/1_node_1_container` (which has **no**
`host_networking.tf` — SLIRP only, no root needed) and layer on the `local-en` corrections.

### 5.1 Serial number flow

On Linux/QEMU, choose the serial yourself and let it flow *down*: the cloud object is the source
of truth, QEMU stamps it via `-smbios type=1,serial=<serial_no>`, EVE reads it as the hardware
serial.

```hcl
resource "zedcloud_edgenode" "EN" {
  name           = "tf_e2e_en_${var.run_id}"
  title          = "tf_e2e_en_${var.run_id}"   # Required, alongside name/model_id/project_id
  serialno       = "TF-E2E-${var.run_id}"
  onboarding_key = var.onboarding_key      # 5d0767ee-0547-4569-b530-387e526f8cb9
  model_id       = zedcloud_model.QEMU_VM.id
  project_id     = zedcloud_project.PROJECT.id
  admin_state    = "ADMIN_STATE_ACTIVE"

  interfaces {
    intfname   = "eth0"
    intf_usage = "ADAPTER_USAGE_MANAGEMENT"
    netname    = zedcloud_network.mgmt_dhcp.name
    ztype      = "IO_TYPE_ETH"
  }
}

resource "zedamigo_installed_edge_node" "INSTALL" {
  serial_no       = zedcloud_edgenode.EN.serialno          # cloud object -> VM
  installer_iso   = zedamigo_eve_installer.eve.filename
  disk_image_base = zedamigo_disk_image.empty.filename
}

resource "zedamigo_edge_node" "VM" {
  serial_no       = zedamigo_installed_edge_node.INSTALL.serial_no
  disk_image_base = zedamigo_installed_edge_node.INSTALL.disk_image
  ovmf_vars_src   = zedamigo_installed_edge_node.INSTALL.ovmf_vars
  depends_on      = [zedcloud_edgenode.EN]   # cloud record must exist before EVE starts registering
}
```

Do **not** use the macOS "soft-serial flip" (`serialno = ...INSTALL.soft_serial`). `local-en`
documents it as a macOS/vfkit-specific workaround; the apparent reason is that
Virtualization.framework offers no SMBIOS serial to stamp, but that rationale is our inference,
not something the source states.

### 5.2 Things to carry over from `local-en`

- `grub_cfg` **must match** `serial_type`: `console=hvc0` for the default `virtio`,
  `console=ttyS0` for `serial`. A mismatch produces an empty console log, so
  `installed_edge_node.success` never flips true and `soft_serial` is never extracted — and the
  failure is silent.
- `tls_ca` — if alpha does not serve a publicly-trusted certificate, EVE's pre-onboarding
  connectivity check fails **silently**. Check first:
  `openssl s_client -connect <device-endpoint>:443 -servername <device-endpoint>`. If it is the
  private Zededa chain, commit the intermediate CA under `test/e2e/cluster_certs/` and wire it
  as `tls_ca`. Alpha is a shared long-lived cluster so this may well be a real cert — verify,
  don't assume.
- `additional_hosts` — not needed if alpha resolves publicly. Omit unless proven necessary.
- `depends_on` on the VM, as above.
- Avoid `cpu_pins` unless there is a reason to want it: there was a VM-start race between the
  socket-tailer and `taskset` (fixed in `f7342b7`), it needs the `taskset` sudoers entry, and it
  pins us to specific host cores, which conflicts with running a pool.

### 5.3 Per-run isolation

Thread `var.run_id = "gh${{ github.run_id }}"` through **every** name: cloud objects, VM names,
`lib_path`. Two payoffs: concurrent runs cannot collide, and orphan cleanup becomes
`rm -rf ~/.local/state/zedamigo-gh-*` filtered by age (§8).

If `zedamigo_local_datastore` is ever used, override `listen` per run — the default `:8080` is a
guaranteed conflict.

### 5.4 Wait for ONLINE

`installed_edge_node.success` only means the *installer* finished; `zedamigo_edge_node` returns
as soon as QEMU is detached, and all of EVE's boot → `register` → `uuid` → `certs` → `config`
handshake happens after `apply` returns. So a barrier is required, or the apply "succeeds"
against a node that never onboarded.

**Use `zedamigo_wait_until`** (v0.13.0+). It runs an idempotent single-shot probe on the
provider `target` on an `interval` until it exits 0 or `timeout` expires, and fails the apply
otherwise. Crucially the probe runs *on the target*, so the same config works whether the target
is localhost or remote — which a `local-exec` provisioner could not do. Per-attempt logs land in
`<lib_path>/wait_until/<id>/attempts/NNNN/` on the target and are deliberately retained on
failure.

The working probe (validated: 5 attempts, 1m21s to ONLINE) polls
`GET /api/v1/devices/id/{id}/status` for `runState == RUN_STATE_ONLINE`; the progression is
`RUN_STATE_UNPROVISIONED` → `RUN_STATE_PROVISIONED` → `RUN_STATE_ONLINE`. See
`test/e2e/node.tf`. Do not put a retry loop inside the script — `interval`/`timeout` own the
retrying.

Regardless, always publish `zedamigo_edge_node.serial_console_log` and the install console log as
artifacts: EVE fails silently on TLS trust and console misconfiguration, and those logs are the
only place it shows up.

One caveat for CI: the probe needs the API token, and the script is written to
`<lib_path>/wait_until/<id>/wait_until.sh` on the target. Either accept that (the lab host is
already trusted with the token) or have the probe read the token from a file placed out of band.

---

## 6. Test-suite changes — `v2/resources`

This is the bulk of the work. Design goal: **the same suite runs in both modes** —
API-only (today's behaviour, no device) and real-node — selected by environment, with the
smallest possible diff to the 18 directories of fixtures.

### 6.1 Placeholder expansion in `MustGetTestInput`

`v2/testing/testhelper.go` currently returns fixture bytes verbatim. Add a token-expansion pass:

```go
// MustGetTestInput reads ./testdata/<path> and expands __TOKEN__ placeholders
// from the test environment. Unknown tokens are a hard failure so a fixture can
// never silently apply with a literal placeholder.
func MustGetTestInput(t *testing.T, path string) string { ... }
```

Tokens (deliberately **not** `${...}`, which would break `terraform fmt`/`validate` on the
fixtures):

| Token | API-only mode | Real-node mode |
|---|---|---|
| `__SUFFIX__` | `""` | `_gh<run_id>` |
| `__NODE_SERIALNO__` | today's literal | the real node's serial |
| `__NODE_NAME__` | `test_tf_provider` | the real node's name |
| `__MODEL_ID__` / `__PROJECT_ID__` | created by the fixture | the e2e config's IDs |

`__SUFFIX__` alone fixes constraint (b): appending it to all 34 duplicated names makes every run
namespaced, which is a prerequisite for a per-PR trigger regardless of real nodes.

### 6.2 Fixture rewrite

Mechanical, but touches 45 `.tf` files (11 of which carry the duplicated `create_edgenode`
name, and 19 of which declare a `zedcloud_edgenode` at all). Sequence that keeps it reviewable:

1. **Commit 1 — `__SUFFIX__` only.** Add the expansion helper; `sed` every `name =`/`title =`
   in `testdata/**/*.tf` to append `__SUFFIX__`. With `__SUFFIX__=""` the suite behaves exactly
   as today, so this is provably a no-op refactor. Golden `.yaml` comparisons ignore names? They
   do not — check `testXAttributes` ignore lists and extend where a name is compared.
2. **Commit 2 — node identity tokens.** Replace hardcoded `serialno` with `__NODE_SERIALNO__`
   and, in the fixtures that create a node purely as a dependency, switch to referencing the
   pre-existing node via `__NODE_NAME__`/ID injection.
3. **Commit 3 — per-resource real-node variants** where the API-only and real-node shapes
   genuinely differ (e.g. `network_instance` `port = "eth1"`, §6.4).

### 6.3 Replace `CI` gating with capability gating

Split the one overloaded `CI` axis into three explicit ones:

| Env var | Gates |
|---|---|
| `ZEDCLOUD_ACC_REAL_NODE=1` | tests that need an onboarded device — un-skips `TestApplicationInstance_CreateCompose` |
| `ZEDCLOUD_ACC_ELEVATED=1` | `TestEnterprise_*` (need operator/distributor permissions) |
| *(unchanged skip)* | `TestProfileDeployment_Create` — this is a **real flake**, not a capability gap. Keep it skipped, but file a ticket; hiding it behind a new flag would launder a known bug. |

Add `testhelper.RealNode(t) RealNodeInfo` returning serial / device UUID / project ID / model ID
from the environment, `t.Skip`-ing when unset. Add `testhelper.WaitForRunState(t, deviceID,
"ONLINE", timeout)` — the same poll helper §5.4 needs, so the workflow and the tests share one
implementation.

### 6.4 The `eth1` problem

Every `network_instance` fixture sets `port = "eth1"`, and the test models declare only `eth0`
in `io_member_list`. The controller accepts this today because nothing checks realizability. A
real single-NIC SLIRP node **cannot** realize it.

Two paths:

- **Phase 1: leave it API-only.** Keep the network-instance tests on a synthetic node. Honest,
  and unblocks everything else.
- **Phase 2: give the VM a second NIC.** `zedamigo_bridge` + `zedamigo_tap` +
  `zedamigo_dhcp_server`, attached via `extra_qemu_args`, exactly as
  `examples/other/1_node_3_vlans` and `2_edge_nodes_bridged` do. This needs `use_sudo = true`
  and it is the point at which the `github-runner` sudoers rule becomes insufficient:
  `zedamigo_dhcp_server` and the `zedamigo_tap` mover daemon self-invoke the provider binary
  under `sudo -n` (§2.1), so a fourth allow-list entry for the bootstrapped binary path is
  needed — or use the netns variants, which go through `sudo ip netns exec …`. Names must carry
  `__SUFFIX__` because interface names are kernel-global.

Do not conflate the two. Phase 1 delivers real onboarding + real app instances; phase 2
delivers real network instances.

### 6.5 Timeout and ordering

- Raise `-timeout` from `20m` to `90m` for the real-node target; keep `20m` for the API-only
  target.
- Add a `make test-e2e` target rather than overloading `make test`, so the existing PR job keeps
  working unchanged while this is built.
- Keep the suite sequential. Do **not** add `t.Parallel()` — with one shared device and a
  ~120-entry ignore-list of API-assigned defaults (45 `cmpopts.IgnoreFields` blocks),
  parallelism will produce noise, not speed.

---

## 7. The workflow

New file `.github/workflows/e2e.yml`. Leave `test.yml` and `go.yml` alone.

```yaml
name: E2E (virtual EVE node)
on: [pull_request]

# PER-BRANCH coalescing only. Do NOT put the lab in a concurrency group (see §7.2).
concurrency:
  group: e2e-${{ github.ref }}
  cancel-in-progress: true       # a new push to the same PR supersedes its own older run

jobs:
  # Single-slot FIFO lock on the lab host. Nothing else runs until this returns.
  lock:
    runs-on: ubuntu-latest
    timeout-minutes: 180
    steps:
      - uses: actions/checkout@v4
      - uses: tailscale/github-action@v3
        with: { oauth-client-id: "${{ secrets.TS_OAUTH_CLIENT_ID }}", oauth-secret: "${{ secrets.TS_OAUTH_SECRET }}", tags: "tag:ci" }
      - name: Take a ticket and wait for our turn
        run: scripts/lab-lock.sh acquire "gh${{ github.run_id }}" --timeout 150m

  node:
    needs: lock
    runs-on: ubuntu-latest
    timeout-minutes: 45
    outputs:
      serialno:  ${{ steps.tf.outputs.serialno }}
      device_id: ${{ steps.tf.outputs.device_id }}
      project_id: ${{ steps.tf.outputs.project_id }}
      model_id:  ${{ steps.tf.outputs.model_id }}
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5            # 1.24+, not the 1.21 the old workflows pin
      - uses: hashicorp/setup-terraform@v3   # pin it; do not rely on the SDK auto-download
        with: { terraform_wrapper: false }

      # --- network path (§3.2, Option 1) ---
      - uses: tailscale/github-action@v3
        with:
          oauth-client-id: ${{ secrets.TS_OAUTH_CLIENT_ID }}
          oauth-secret:    ${{ secrets.TS_OAUTH_SECRET }}
          tags: tag:ci

      - name: Pin lab host key
        run: printf '%s\n' "${{ vars.LAB_SSH_HOST_KEY }}" > test/e2e/known_hosts

      - name: Build provider under test
        run: make build                      # -> v2/terraform-provider-zedcloud_<ver>

      - name: Create node
        id: tf
        working-directory: test/e2e
        env:
          TF_VAR_run_id:      gh${{ github.run_id }}
          TF_VAR_lab_host:    ${{ vars.LAB_HOST }}
          TF_VAR_lab_ssh_key: ${{ secrets.LAB_SSH_PRIVATE_KEY }}
          TF_VAR_zedcloud_url:   ${{ vars.TF_VAR_ZEDCLOUD_URL }}     # zedcontrol.alpha...
          TF_VAR_eve_cluster_host: ${{ vars.EVE_CLUSTER_HOST }}      # device API, §3.3
          TF_VAR_zedcloud_token: ${{ secrets.TF_VAR_ZEDCLOUD_TOKEN }}
        run: |
          terraform init -input=false
          terraform apply -auto-approve
          terraform output -json | tee ../../outputs.json
      - name: Wait for ONLINE
        # NB: v2/ is a separate Go module, so this must be `go run -C v2`
        run: go run -C v2 ./testing/cmd/waitonline -device "${{ steps.tf.outputs.device_id }}" -timeout 15m

  suite:
    needs: node
    runs-on: ubuntu-latest
    timeout-minutes: 120
    env:
      TF_ACC: 1
      CI: true
      ZEDCLOUD_ACC_REAL_NODE: 1
      ZEDCLOUD_TEST_SUFFIX: _gh${{ github.run_id }}
      ZEDCLOUD_TEST_NODE_SERIALNO: ${{ needs.node.outputs.serialno }}
      ZEDCLOUD_TEST_NODE_ID:       ${{ needs.node.outputs.device_id }}
      TF_VAR_zedcloud_url:   ${{ vars.TF_VAR_ZEDCLOUD_URL }}
      TF_VAR_zedcloud_token: ${{ secrets.TF_VAR_ZEDCLOUD_TOKEN }}
    steps: [ checkout, setup-go, setup-terraform, "make test-e2e" ]

  teardown:
    needs: [lock, node, suite]
    if: always()                             # non-negotiable
    runs-on: ubuntu-latest
    steps:
      - tailscale, checkout, restore state artifact
      - run: terraform destroy -auto-approve
      - run: ssh github-runner@$LAB_HOST 'bash -s' < scripts/sweep_zedamigo.sh gh${{ github.run_id }}
      - if: always()
        run: scripts/lab-lock.sh release "gh${{ github.run_id }}"   # last thing, always
      - uses: actions/upload-artifact@v4     # serial console logs + terraform.log, always
```

Two details that matter more than they look:

- **State must survive between the `node` and `teardown` jobs.** Options: keep everything in one
  job (simplest, but no parallel `suite`), upload `terraform.tfstate` as an artifact between
  jobs, or use a real backend. `terraform-provider-zedamigo/misc/tf_state_backend_http` exists
  but its own README says no locking, no auth, no TLS — **do not use it**. Prefer a proper
  backend (S3/GCS) or the artifact hand-off plus the belt-and-braces sweep. A lost state file
  means permanently orphaned VMs on a long-lived host.
- **The destroy path must not depend on the runner surviving.** `if: always()` does run on
  cancellation, but only within GitHub's cancellation grace period — a force-killed runner
  skips teardown entirely. The host-side reaper in §7.2 is what actually guarantees the nodes
  die.

### 7.2 Serialization: the lab is a single-slot resource

**Requirement:** the Berlin lab can host one run at a time; with several PRs open, they run
one after another in the order they arrived (first-come-first-served by run).

**GitHub's `concurrency:` cannot express this.** A concurrency group holds at most **one**
pending run. When a third run arrives while one is running and one is pending, the *pending*
one is **cancelled** and replaced — even with `cancel-in-progress: false`, which only protects
the *running* job. So a global `group: lab-h-m-dl20` would give newest-wins, silently starving
the PR that has been waiting longest. That is the opposite of the requirement, and it is why
§7's `concurrency` block is scoped to `github.ref` and used only for per-branch coalescing
(where newest-wins *is* what you want: a new push to the same PR should supersede its own
in-flight run).

Cross-PR serialization therefore needs an explicit queue.

**Start from `zedamigo_host_reservation`, not from scratch.** zedamigo v0.13.0+ ships
cooperative capacity reservation, and the tree is already deployed on `h-m-dl20`
(`/var/lib/zedamigo/reservations`: 16 CPU slots, 58 GB slots, the LVM volumes as devices).
Claims are atomic via `flock`, durable across reboots until the resource is destroyed, visible
to *any* config on the host including a human's, and each slot records who holds it. That is
most of what the bespoke lock below was invented to do, already built and already tested
(validated in iteration 1: `reserved_cpus = [0,1,2,3]`).

Its one gap: it is **admission control, not a queue** — if capacity is short the apply *fails*
rather than waiting, which would turn a busy lab into spurious red PRs. So the remaining work is
a *waiting* layer on top, not a whole lock. The cheapest version is a `zedamigo_wait_until`
barrier that polls for free capacity before the `host_reservation` resource claims it; that
reuses two shipped resources and needs no custom scripts. A FIFO ticket queue is only necessary
if fairness across PRs turns out to matter in practice — measure first (§9 phase 4).

Note also that the lock file's permissions currently make reservation single-user (§0), so this
depends on that being fixed in the lab's `configuration.nix`.

If a bespoke queue does prove necessary, the shape below is the proposal: a **ticket lock on the
lab host**, driven over SSH by `scripts/lab-lock.sh` (`acquire` / `release` / `status` / `reap`),
with state under `/home/gh-run-tf-zedcloud/.local/state/e2e-lab-lock/`:

```
queue        # append-only ticket list, one "<run_id> <epoch>" per line, mutated under flock
lease        # "<run_id> <epoch_acquired>" — who currently owns the lab
hb/<run_id>  # heartbeat file, mtime touched every 30s by waiters AND the holder
```

- **`acquire <run_id>`** appends a ticket under `flock`, then polls every 15s until (a) no live
  lease exists and (b) our run is the oldest live ticket. Then it writes `lease` under `flock`
  and returns. Ordering is by ticket position, so it is FIFO by arrival.
- **Heartbeats are the liveness signal.** Both waiters and the holder touch `hb/<run_id>` every
  30s. `reap` (called at the top of every poll iteration, and by the janitor) drops any ticket
  or lease whose heartbeat is older than 2 minutes. This is what makes the design survive the
  cases that matter:
  - a **waiting** run is cancelled → its stale ticket is dropped, the queue advances, no
    human involvement;
  - the **holding** run is force-killed before teardown → the lease is broken after 2 minutes
    *and* the reaper immediately runs `sweep_zedamigo.sh <run_id>`, so the orphaned QEMU
    processes and `lib_path` go with it.
- No GitHub credentials are needed on the host — liveness is inferred from heartbeats, not
  from the Actions API.
- **Humans can take the lock too.** `lab-lock.sh acquire manual` before an interactive debugging
  session stops CI from stomping on the host mid-session; `release manual` hands it back. This
  is a genuine advantage of putting the lock on the host rather than in GitHub: it covers *all*
  users of the lab, not just Actions.

Alternatives considered:

- **`softprops/turnstyle`** (block until earlier runs of the same workflow finish) gets FIFO
  by run number with no host-side state, but it only knows about this one workflow, cannot see
  a human using the lab, and has no way to sweep after a killed holder. Reasonable fallback if
  the host script proves fiddly; not the primary choice.
- **A single self-hosted runner** would give all of this for free — GitHub queues jobs for a
  runner label and dispatches them FIFO, one at a time, with no lock, no polling and no wasted
  minutes. This was considered and set aside in favour of GitHub-hosted runners (§3.1). Worth
  revisiting if the lock proves to be a maintenance burden, since the new single-slot constraint
  is *exactly* what one runner models naturally.

**Cost to be aware of:** a queued run occupies a GitHub-hosted runner while it polls. The `lock`
job is deliberately a separate, near-empty job so a waiter burns one minimal runner rather than
holding a fully provisioned one, but the minutes are still billed on a private repo. If queue
depth becomes a problem, the knobs are: skip draft PRs, run only on `ready_for_review` +
subsequent pushes, or fall back to a nightly-plus-label trigger. Cap the wait at 150 minutes and
fail with an explicit "lab busy, re-run when free" message rather than hanging until the job
timeout.

### 7.3 Provider resolution

Both providers must resolve to local builds, and the repo's existing `dev.tfrc`
(`dev_overrides` + `direct {}`) actively **breaks zedamigo resolution** — this is recorded in
`local-en/RUNBOOK.md`. Use a dedicated `test/e2e/e2e.tfrc` with a `filesystem_mirror` pattern
copied from `local-en/zedamigo.tfrc`, set `TF_CLI_CONFIG_FILE` to it for the e2e config only,
and leave `dev.tfrc` alone for the in-process acceptance tests. Those tests do not need it for
provider resolution — `Providers: testAccProviders` makes the SDK reattach the in-process
provider — but `testhelper.CheckEnv` (`v2/testing/testhelper.go:67-74`) **hard-fails** unless
`TF_CLI_CONFIG_FILE` is set or a `dev.tfrc` exists in the CWD, so it must stay set to
*something*.

Also: zedamigo's remote-binary bootstrap **refuses versions `""`, `dev` and `test`**
(`ssh_config.go:366-368`). So either pin a released zedamigo version and let `install.sh
--binary-only` bootstrap it (requires outbound HTTPS from `h-m-dl20` to
`raw.githubusercontent.com` and `github.com`), or pre-place a binary and set
`ssh.remote_binary_path`. The bootstrap path is preferable — one less thing to keep in sync.

---

## 8. Cleanup and orphan control

Four layers, because any one of them will fail eventually.

1. **`terraform destroy`** in an `if: always()` job. Handles the happy and most unhappy paths.
2. **A host sweep script** (`scripts/sweep_zedamigo.sh`, run over SSH), keyed on the per-run
   `lib_path`:
   ```bash
   pkill -f "zedamigo-gh-${RUN_ID}/edge_nodes/" || true
   pkill -f "terraform-provider-zedamigo -socket-tailer" || true
   rm -rf "$HOME/.local/state/zedamigo-gh-${RUN_ID}"
   ```
   Note that `zedamigo_edge_node.Delete` treats a failed QMP `quit` as a *warning* and removes
   the state directory anyway — so an orphaned QEMU process with its files deleted is a normal,
   expected outcome, not an exotic one. The `pkill` is load-bearing.
3. **The lock reaper** (§7.2). Because it is driven by heartbeats on the host rather than by the
   workflow, it is the only layer that still works when the runner is force-killed — the case
   `if: always()` does *not* cover. Breaking a stale lease also triggers layer 2 for that run's
   id, so a cancelled run's VMs die within ~2 minutes and the lab slot is freed for the next
   ticket in the queue. This is the layer that makes "one run at a time" safe in practice.
4. **A scheduled janitor** (nightly): `lab-lock.sh reap`, sweep any `zedamigo-gh-*` directory
   older than 4 hours and any matching QEMU process, and delete cloud objects whose names match
   `tf_e2e_%`/`test_tf%` older than 24h. Belt and braces for whatever the first three missed.

For the cloud side, prefer an **API-based** sweeper over
`add-tf-agent-skill`'s `clean_tf_objects_on_alpha.sh`, which issues `DELETE` straight against
the alpha Postgres — and not only for `test_tf%`: it also sweeps `openalpine%`, `open_ds%`,
`qemu100_`, `default-app%` and `alphagroup13`. Deleting controller rows behind the API's back
will desynchronise caches and mask exactly the kind of bug this suite exists to catch. The
provider already tags test traffic with `User-Agent: zededa-terraform-provider/testbuild`
(`v2/resources/provider.go:140,148-150`) — a usable server-side signal if a controller-side
sweep endpoint is ever wanted.

---

## 9. Work breakdown

Phases are ordered so each one ends somewhere shippable.

### Phase 0 — Unblock (do before writing code)

| # | Task | Owner |
|---|---|---|
| 0.1 | Confirm the device-facing endpoint name for alpha (`zedcontrold.alpha` vs `zedcloud.alpha`) and whether it serves a publicly-trusted cert | cluster owners |
| 0.2 | **Manually onboard one virtual EVE node to alpha**, from a laptop, using the `local-en` recipe adapted to Linux/QEMU. Confirms the onboarding key `5d0767ee-…` is authorized in enterprise `AAFlABBtJ0mP_lhJjIyVzjzgMXiR`. This was the hard blocker on the local cluster (`register` → 403 "device not found") and it is **not fixable from Terraform** | us |
| 0.3 | Agree the network path (§3.2). Spike Tailscale on a throwaway branch | us + IT |
| 0.4 | Measure `h-m-dl20`: core count, RAM, NVMe free space; decide pool size | us |
| 0.5 | **Rotate the credentials committed to `add-tf-agent-skill`** (§10.1) | us |

Phase 0 is not optional overhead. 0.2 in particular is the difference between "three weeks of
work" and "three weeks of work that cannot possibly succeed".

### Phase 1 — Host + network (`zededa-berlin-lab`)

1.1 Real `github-runner` SSH key. 1.2 `nested=1`. 1.3 Tailscale subnet router (or chosen
alternative). 1.4 EVE image pre-pull. 1.5 `nixos-rebuild switch`, then verify from a hosted
runner that a trivial workflow can `ssh github-runner@<host> 'qemu-img --version && docker
version && ls /dev/kvm'`.

**Exit criterion:** a hosted runner can run commands on `h-m-dl20`.

### Phase 2 — Node lifecycle (`test/e2e/`)

2.1 The Terraform config (§5). 2.2 `e2e.tfrc`. 2.3 The wait-for-online helper + `waitonline`
command. 2.4 `sweep_zedamigo.sh`. 2.5 `lab-lock.sh` (`acquire`/`release`/`status`/`reap`) plus the
heartbeat/reaper logic (§7.2) — test it standalone by racing three fake tickets before it goes
anywhere near a workflow. 2.6 A minimal `e2e.yml` that takes the lock, creates a node, waits for
ONLINE, destroys it and releases — no test suite yet.

**Exit criterion:** green workflow that onboards and destroys a node, with the serial console log
uploaded as an artifact on failure — and two concurrently-triggered runs demonstrably executing
one after the other, in arrival order, with neither cancelled.

### Phase 3 — Suite refactor (`v2/resources`, `v2/testing`)

3.1 `__SUFFIX__` expansion + fixture `sed` (provable no-op at `__SUFFIX__=""`). 3.2
`testhelper.RealNode` + `WaitForRunState`. 3.3 Replace `CI` gating with
`ZEDCLOUD_ACC_REAL_NODE` / `ZEDCLOUD_ACC_ELEVATED`. 3.4 Node-identity tokens in fixtures. 3.5
`make test-e2e` with a 90m timeout. 3.6 Un-skip `TestApplicationInstance_CreateCompose` and make
it assert the app actually reaches running state.

**Exit criterion:** the suite passes in both modes; `CreateCompose` passes against a real node.

### Phase 4 — Per-PR viability

4.1 Wire `e2e.yml` to `pull_request` with per-branch `concurrency` + the host lock (§7.2).
4.2 The nightly janitor (including `lab-lock.sh reap`). 4.3 Measure wall-clock and queue depth
over ~20 real PRs; publish `lab-lock.sh status` somewhere visible. 4.4 If PR latency is
unacceptable — with one slot it probably will be — implement the warm node (§3.4 option B).

### Phase 5 — Second NIC (optional, §6.4)

5.1 `zedamigo_bridge`/`tap`/`dhcp_server` with `__SUFFIX__`-scoped names. 5.2 Model
`io_member_list` gains `eth1`. 5.3 Real network-instance tests, asserting realization.

Rough shape: Phase 0 is days of other people's calendars; Phases 1-2 about a week; Phase 3 is
the big one (two to three weeks, mostly fixture archaeology); Phases 4-5 iterative.

---

## 10. Risks

### 10.1 Committed credentials — fix this week, independent of everything else

Branch `add-tf-agent-skill` contains, in plaintext:

- `set_api_token.sh` — alpha UI password for `ivan-icd-test@zededa.com`, and the local-cluster
  password for `watchdog-c20-local@zededa.com`
- `clean_tf_objects_on_alpha.sh` — Postgres credentials for both the `alpha` and `local`
  databases

These are committed on the branch `add-tf-agent-skill`, which in the local clone is **not**
present on `origin` (unlike `tf-actions`, which is). Confirm it was never pushed — then rotate
all four regardless, move them to GitHub secrets, and rewrite the scripts to read from the
environment. Do this independently of whether this design ships. `set_api_token.sh` also mints
90-day tokens (`expires: 7776000`); shorten that for anything used by CI.

### 10.2 The onboarding certificate

The highest-probability project-killer. On the local cluster this manifested as
`POST /api/v2/edgedevice/register` → 403 "device not found", because no onboarding cert for key
`5d0767ee-0547-4569-b530-387e526f8cb9` was registered in the enterprise. Confirmed dead ends
(all documented in `local-en/RUNBOOK.md`): the zedcloud provider exposes only `onboarding_key`,
never the cert; REST `PUT` silently drops `onboarding.pemCert`; `zcli edge-node create
--onboarding-certificate` works but only at create time and is mutually exclusive with
`--onboarding-key`; zedamigo cannot inject `onboard.cert.pem`; and the `lf-edge/eve` repo's
bundled cert **expired in 2020**. Alpha is long-lived and shared so it is *probably* seeded —
but "probably" is why task 0.2 exists.

### 10.3 One host, one slot, per-PR trigger

Single point of failure and a hard serialization bottleneck: with a 30-60 minute run and one
slot, throughput is at best ~1 PR/hour, and a queue of five PRs means the last one waits half a
day. Mitigations: the FIFO ticket lock (§7.2), the four-layer cleanup (§8), and the warm node
(§3.4B) to cut per-PR cost. Beyond that the only real lever is queue depth — skip drafts, or
demote to nightly-plus-label if the queue is consistently deep.

Two operational consequences to accept up front: a lab outage blocks the e2e check on **every**
PR, and a single stuck run blocks everyone behind it. So (a) keep this check **non-required** in
branch protection until it has been stable for several weeks, and (b) make `lab-lock.sh status`
trivially runnable by a human, because "who is holding the lab and since when" will be the most
frequently asked question about this system.

### 10.4 Silent failure modes

EVE fails *quietly* on TLS trust and on console misconfiguration. Both surface only in the
serial console log. Always upload `zedamigo_edge_node.serial_console_log` and
`installed_edge_node.serial_console_log` as artifacts — on success as well as failure, at least
initially.

### 10.5 Fixture refactor blast radius

Touching 45 fixture files across 18 resource directories risks breaking the tests everyone
relies on. This
is why Phase 3.1 is structured as a provable no-op (`__SUFFIX__=""` → byte-identical config) and
lands before any semantic change.

### 10.6 Flakes we already know about

`TestProfileDeployment_Create` (skipped on `CI`, passes locally). Nine steps across seven files
set `ExpectNonEmptyPlan: true`, masking non-converging diffs. The provider retries GET 404s
using `retryablehttp`'s defaults (4 attempts, exponential backoff — `RetryMax` is never set
explicitly, `provider.go:188`), which inflates wall clock on failure paths. And the
recent commit history — CI-723, CI-790, CI-865, CI-875 — is a steady stream of exactly the
drift/diff bugs a real-node suite is supposed to catch earlier. That is the argument *for* this
work; it is also a warning that the first few real-node runs will surface a queue of genuine
bugs, and the suite will look flaky while they are fixed. Plan for that politically, not just
technically.

---

## 11. Open questions

1. **Device API hostname on alpha.** Is `zedcontrold.alpha.zededa.net` the device-facing
   endpoint (`/config/server`), with `zedcontrol.alpha.zededa.net` remaining the control API?
   Does it present a publicly-trusted certificate?
2. **Is the default EVE onboarding key authorized in enterprise `AAFlABBtJ0mP_lhJjIyVzjzgMXiR`?**
   (Task 0.2.)
3. **Which EVE tag do we pin?** `local-en` used `14.5.3-lts-kvm-arm64`; the zedamigo examples
   range across `14.5.3`, `16.0.0` and `16.0.1` LTS tags plus an `rc1` and a PR build — the
   template we would copy (`1_node_1_container`) uses `16.0.0-lts-kvm-<arch>`. Should CI track
   the LTS the product supports, or the newest?
   Testing a provider against a stale EVE has limited value; testing against a moving tag makes
   CI non-reproducible. Suggest pinning, with a scheduled bump PR.
4. **`h-m-dl20` capacity** — cores, RAM, NVMe free. Determines pool size and whether concurrent
   runs are ever viable.
5. **Is Tailscale acceptable** in the lab network, and who owns the tailnet ACLs?
6. **Terraform state backend** — is there an existing org S3/GCS bucket we should use, or is the
   artifact hand-off good enough?
7. **Does alpha tolerate per-PR churn?** We would be creating and deleting projects, models,
   brands, networks, images and devices on a shared cluster on every PR. Confirm with its owners,
   and confirm nobody else's tests rely on `test_tf%` objects.
8. **Should the e2e config live here or in zedamigo?** Proposal is `test/e2e/` in this repo
   (versioned with the provider under test). The alternative — a third repo — decouples but
   adds a cross-repo pinning problem.
9. **`terraform` or `tofu`?** The host has `opentofu` installed; the workflows use
   `hashicorp/setup-terraform`. Pick one for the e2e config.

---

## 12. Reference index

**`zededa/terraform-provider-zedcloud`**
- `.github/workflows/test.yml`, `go.yml` — current PR jobs; `go.yml`'s test step is commented out
- `Makefile` — `test`, `test-run` targets; `TF_ACC=1`, `TF_CLI_CONFIG_FILE`, `-timeout 20m`
- `v2/testing/testhelper.go` — `MustGetTestInput`, `MustGetExpectedOutput`, `CheckEnv`
- `v2/resources/provider.go:23,34,40,140,148-150,188,197-222` — `defaultURL`, env defaults,
  `testbuild` User-Agent, the retryable client, the 404-retry policy
- `v2/resources/provider_test.go` — `testAccProviders` wiring
- `v2/resources/node_test.go:199,262` — golden-diff ignore list, 404-shape handling
- `v2/resources/application_instance_test.go:65-71` — *"skipped until we decide how to provide a
  real device"*
- `v2/resources/testdata/**` — 18 dirs, 45 `.tf` fixtures + 37 `.yaml` goldens, all static
- `dev.tfrc` — `dev_overrides` to `./v2`; breaks zedamigo resolution
- `test/retry/` — manual mitmproxy 404 fault injection (not CI)
- `docs/design/NFR-238-…md` — design-doc convention followed here
- branches: `tf-actions` (provider Actions, unrelated — the only one of these on `origin`),
  `add-tf-agent-skill` (alpha scripts, §10.1 — local-only in this clone), `add-claude`
  (local-only), `tf-actions-hackaton` (local-only, empty vs `main`)

**`andrei-zededa/terraform-provider-zedamigo`**
- `internal/provider/ssh_config.go` — `ssh{}` schema, `ZEDAMIGO_SSH_*` env fallbacks,
  fail-closed host-key check (`:300`), `bootstrapRemoteBinary` (`:365-390`), version refusal
  (`:366-368`)
- `internal/exec/{exec,ssh,local}.go` — the Executor abstraction, `sudo -n kill`
- `internal/hypervisor/qemu.go:177-283` — install/runtime QEMU args, `-smbios … serial=`
- `internal/hypervisor/forwards.go:32-38,66-73` — SLIRP nic0, port forwards on remote targets
- `internal/provider/eve_installer_resource.go:291-302` — the `docker run … lfedge/eve` call
- `internal/provider/installed_edge_node_resource.go:356-373` — `success` / `soft_serial`
  scraping from the console log
- `internal/provider/edge_node_resource.go:357,611-635` — random `ssh_port`, delete-path QMP
  warning
- `examples/other/1_node_1_container/` — SLIRP-only template (no `host_networking.tf`)
- `examples/other/{1_node_3_vlans,2_edge_nodes_bridged}/` — bridge/tap/VLAN patterns,
  `config_suffix` convention
- `scripts/e2e_lag.sh` — rootless e2e harness; `dev_overrides` + `PATH` stubs worth stealing
- `misc/tf_state_backend_http/` — **do not use** (no locking/auth/TLS, per its own README)
- `CLAUDE.md`, `README.md:85-88` (endpoint naming; contradicted by the examples)

**`andrei-zededa/zededa-berlin-lab`**
- `hosts/h-m-dl20/configuration.nix` — `github-runner` user, sudo allow-list, podman
  `dockerCompat`, `kvm` group, LVM udev rule
- `hosts/h-m-dl20/hardware-configuration.nix` — `kvm-intel`, no `nested=1`
- `hosts/h-m-dl20/monitoring.nix` — VictoriaMetrics/Logs, per-VM process_exporter grouping
- `hosts/h-m-dl20/README.md` — `nixos-rebuild` deploy procedure, `ProxyJump=root@localhost:11022`
- `berlin-vpn-in-a-container/` — the VPN jump host (§3.2 option 2)

**`andrei-zededa/local-en`** (not a repo we own, but the primary reference)
- `CLAUDE.md` — endpoint split (`:46`), TLS CA extraction (§4), orphan recovery (`:102-108`),
  teardown (final section)
- `RUNBOOK.md` — the onboarding-cert blocker and every confirmed dead end
- `edge_node.tf`, `cluster_objects.tf` — the working two-provider wiring
- `zedamigo.tfrc` — `filesystem_mirror` pattern that avoids the `dev.tfrc` conflict
