# UE-156 — 3-node edge-node cluster on virtual EVE nodes

Status: implementation plan
Ticket: [UE-156](https://zededa.atlassian.net/browse/UE-156) (parent [UE-131](https://zededa.atlassian.net/browse/UE-131))
Blocks: [CI-834](https://zededa.atlassian.net/browse/CI-834), [CI-709](https://zededa.atlassian.net/browse/CI-709)
Related: [UE-157](https://zededa.atlassian.net/browse/UE-157) (spike: arm64 kubevirt EVE, for a laptop-local loop — does not gate this)
Ref: [`e2e-virtual-eve-node-testing.md`](./e2e-virtual-eve-node-testing.md) §6.4, [`ue-139-fixture-suffix-tokens.md`](./ue-139-fixture-suffix-tokens.md), [`ue-140-compose-real-node.md`](./ue-140-compose-real-node.md) §3.2
Scope: **harness + reproduction only.** Provider fixes stay in their own tickets.
Artifacts: [`test/e2e-cluster/`](../../test/e2e-cluster/) — a complete, unvalidated first cut of the configuration described here.

---

## 0. Headline: the Terraform is the easy part

`zedcloud_edgenode_cluster` has existed since `6f88cb33` (zedcloud-956, Feb 2025). The
resource works, the API is complete, and the HCL for a 3-node cluster is about twenty lines.
Andrei has had it working since February — the state dump on CI-834 is proof.

What does not exist is a harness that can *stand up three virtual EVE nodes that are eligible
to be clustered*. Eligibility is the whole problem, and it decomposes into four gates, of
which **three are host/image concerns rather than Terraform concerns**:

| gate | enforced by | current harness |
|---|---|---|
| all nodes report `hvTypeKubevirt` | the controller's live node-capability check | ✗ pinned to a `kvm` image |
| cluster interface on a shared L2 segment | the controller's cluster-interface resolution + physics | ✗ one SLIRP NIC, no guest-to-guest path |
| `use_sudo` for tap/dhcp resources | `zedamigo` self-invoking under `sudo -n` | ✗ 3 NOPASSWD entries, needs 4 |
| 3 × node capacity | `zedamigo_host_reservation` (fail, not queue) | ✗ sized for 1 node |

So this ticket is 10% HCL and 90% making the lab able to host a cluster. The ordering below
front-loads the two gates that can invalidate everything downstream.

**Recommendation: resolve §1 (the EVE image) before writing or reviewing any more code.** If
Zededa has no kubevirt EVE build the harness can pin, the rest of this plan does not execute
and the ticket becomes "obtain a kubevirt EVE image".

---

## 1. Gate 1 — the EVE image must report KubeVirt support (**BLOCKING**)

### 1.1 The gate

the controller's live node-capability check issues a **live** `DeviceStatusReq`
over Kafka/gRPC to every candidate node at cluster-create time, bounded by
`clusterNodesInfoQueryTimeout`:

```go
if len(deviceInfoList) != len(nodeIDs) {
    return fmt.Errorf("failed to validate nodes, found nodes: %d, expected: %d", ...)
}
for _, devInfo := range deviceInfoList {
    if devInfo.GetOptionalCapabilities() == nil ||
       !devInfo.GetOptionalCapabilities().GetHvTypeKubevirt() {
        return fmt.Errorf("node %s does not support kubevirt", devInfo.GetId())
    }
}
```

`libs/zmsg/device/devstatus.proto:522`:

```proto
message OptionalCapabilities {
    bool hvTypeKubevirt = 1;   // "not supported by all eve flavors"
    bool etcdSnapshot   = 3;   // eve-k
}
```

Two consequences that are easy to miss:

- This is a **live query, not a cached capability**. All three nodes must be online
  *simultaneously* at create time, or the create fails with
  `timed out after %s querying cluster node status`. That is why
  `zedcloud_edgenode_cluster` must depend on all three onboarding barriers, not just on the
  node records.
- It is checked **only at cluster create**. A node can onboard perfectly, pass every other
  check, and then fail the cluster create with a flat 400 that reads like a configuration
  error. But polling this field is NOT a usable readiness barrier — see §1.2.2. The barrier
  that works is `zedamigo_wait_until.kube_ready`, which asks the node over SSH.

### 1.2 The problem

`test/e2e/vars.tf` pins `eve_tag = "16.0.1-lts-kvm-amd64"`. The `kvm` and `xen` flavours do
not report `hvTypeKubevirt`. **The existing harness image cannot form a cluster.**

**The fix is `17.0.0-lts-k-amd64`** — see §1.2.1 for why the tag is spelled `-k-` and why
an earlier revision of this document got the available images badly wrong.

### 1.2.1 Architecture: arm64 exists — the flavour is spelled `-k-`

**CORRECTION (2026-09-17).** An earlier revision of this section claimed there was no arm64
KubeVirt EVE image and that clustering was therefore x86_64-only. That was wrong, and the
error was a naming assumption: **the kubevirt flavour was renamed from `-kubevirt-` to
`-k-`.** Searching Docker Hub for `kubevirt` finds only the old, abandoned naming (newest
`15.10.0-kubevirt-amd64`).

What actually exists:

```
16.0.0-lts-k-amd64   16.0.1-lts-k-amd64   16.0.2-lts-k-amd64
17.0.0-lts-k-amd64   17.0.0-lts-k-arm64
```

EVE master's Makefile agrees — the kube packages are gated on `ifeq ($(HV),k)`, not
`HV=kubevirt`. The upstream zedamigo reference config
([`examples/other/edge_node_cluster`](https://github.com/andrei-zededa/terraform-provider-zedamigo/tree/main/examples/other/edge_node_cluster))
pins `17.0.0-lts-k-<arch>`.

Consequences:

- `test/e2e-cluster/vars.tf` pins `17.0.0-lts-k-amd64`.
- **An Apple Silicon loop is not ruled out.** `17.0.0-lts-k-arm64` is published, so no build
  work is needed to try one. What remains unexplored is the vfkit path: no Linux bridges or
  taps on macOS, so the multi-node shared-L2 design here does not transfer, and a single-node
  cluster is the realistic target there — which is sufficient for CI-834 and CI-709 (§9.1).
- [UE-157](https://zededa.atlassian.net/browse/UE-157) was opened to spike *building* arm64
  kubevirt EVE. **That premise is void** — the image is published. Repurpose or close it.

Two clarifications worth keeping, because it is easy to reach for the wrong explanation:

- **Hardware nested virtualization is not the blocker.** Nested virt (M3 and later, exposed
  through Virtualization.framework) lets a guest run a hypervisor, which a virtual EVE node
  genuinely needs in order to run VM-type app instances — a real capability gain for this
  harness. It is independent of architecture, and it was never what blocked clustering.
- **The gate is the build flavour, not actual nested VM execution.**
  the controller's live node-capability check only checks that EVE *reports* `hvTypeKubevirt`.

### 1.2.2 The capability is not a usable readiness signal

Also learned the hard way. A `15.10.0-kubevirt-amd64` node onboarded cleanly, reached
`RUN_STATE_ONLINE` / `ADMIN_STATE_REGISTERED`, and then sat for 20+ minutes with
`optionalCapabilities: null` and **zero** k3s / kubelet / longhorn / etcd in its console
log — the Kubernetes stack never started at all.

Beyond the wrong image, polling the controller for the capability is the wrong test in
principle: the controller's device-info handler only populates the field
`if dinfo.GetOptionalCapabilities() != nil`, so "not reported yet" and "not capable" are
indistinguishable from outside — both are simply absent.

Upstream instead asks the node, over SSH:

```
eve exec kube ls -l /var/lib/all_components_initialized
```

at 10s intervals with a 20m budget, noting that without that barrier cluster formation
"fails or stalls". `test/e2e-cluster` now does the same, which makes
`var.edge_node_ssh_pub_key` **mandatory** rather than a convenience.

One hypothesis that proved wrong along the way: the missing `swtpm` on h-m-dl20. The
upstream config uses no TPM at all — no `zedamigo_swtpm`, no `swtpm_socket`.


### 1.3 Action — RESOLVED

**Use `17.0.0-lts-k-amd64`** (or `-arm64`), matching the upstream zedamigo reference config.
`test/e2e-cluster/vars.tf` pins it. This is no longer a blocker, and no EVE-team question is
needed to proceed.

Resource shape is also settled by the reference config rather than guessed: **6 vCPU / 16 GB /
~30 GB per node**, with CPU pinning. See §4 for the capacity consequence on h-m-dl20.

### 1.4 How `test_clusters.py` sidesteps this, and why we cannot

`zedcloud/tests/apiTest/tests/test_clusters.py` is the only existing end-to-end cluster test.
It uses `zobjects/zdevice.py:DeviceSimulate` — a pure-protobuf fake device, no VM — and
satisfies the gate by simply asserting the capability:

```python
dev_sim_obj.send_info(kube_virt_supported=True)     # -> hv_type_kubevirt = True
```

That is a fine unit-level test of the *controller*, and a useless test of the *provider*: it
never exercises a real config push, a real Terraform refresh, or a real `plan`. It is precisely
why CI-834 and CI-709 escaped. We need real EVE.

Its sequencing is still worth copying verbatim:

1. brand → model → project → networks → devices with two interfaces each
2. encryption certs, then `send_info(kube_virt_supported=True)`
3. negative cases (2 nodes → 400, active app instance → 400, bad cluster interface → 400)
4. deactivate app instances, then create the 3-node cluster → 200
5. `wait_for_cluster_config`: poll until `config['cluster']['clusterId']` appears on all three

Step 3 is a cheap and valuable addition to the acceptance test — see §7.3.

---

## 2. Gate 2 — a shared L2 segment for the cluster interface

### 2.1 Why SLIRP cannot carry it

`test/e2e/node.tf` gives each VM exactly one NIC, on QEMU SLIRP (user-mode NAT). SLIRP hands
every guest the same isolated `10.0.2.15/24` behind its own NAT with **no guest-to-guest
path**. Three SLIRP-only nodes cannot see each other at all. No Zedcloud configuration changes
that.

`test/e2e/cluster_objects.tf` is explicit about the consequence:

> One management NIC. The VM is created with a single QEMU SLIRP interface, so declaring more
> here would describe hardware that does not exist.

And `ue-140-compose-real-node.md` §3.2 reached the same conclusion from a different direction,
with corroboration from a working customer config:

> The customer's model declares both `eth0` and `eth1`, the node assigns `eth0` to management
> and `eth1` to `ADAPTER_USAGE_APP_SHARED`… **Treat the second NIC as mandatory**, not as a
> thing to try shortcuts around. Binding a switch NI to the sole management port is how you
> lose the node.

Zedcloud's own virtual device model, `spec/template_l1_ZedVirtual-4G.json`, declares `eth0`
(MANAGEMENT) **and** `eth1` — for the same reason.

### 2.2 The topology

Keep eth0 on SLIRP so the outbound path to the device API is unchanged from `test/e2e` (a
verified-working path — `/api/v2/edgedevice/ping` → 200 from the lab host). Add eth1 as a tap
on a shared host bridge:

```
+-------------- host bridge cbr-<run_id> (10.77.77.0/24) --------------+
|            |                      |                      |          |
|     ctap-<run_id>-1        ctap-<run_id>-2        ctap-<run_id>-3    |
|            |                      |                      |          |
|         VM 1 eth1              VM 2 eth1              VM 3 eth1     |
|         VM 1 eth0              VM 2 eth0              VM 3 eth0     |
|              (SLIRP -> host NAT -> device API)                      |
+---------------------------------------------------------------------+
```

This is `e2e-virtual-eve-node-testing.md` §6.4 **Phase 2** arriving, as predicted there:

> `zedamigo_bridge` + `zedamigo_tap` + `zedamigo_dhcp_server`, attached via `extra_qemu_args`,
> exactly as `examples/other/1_node_3_vlans` and `2_edge_nodes_bridged` do. This needs
> `use_sudo = true`… Names must carry `__SUFFIX__` because interface names are kernel-global.

Three matching changes on the Zedcloud side, all in `cloud_objects.tf`:

1. a second `io_member_list` block on the model for `eth1`, `ADAPTER_USAGE_APP_SHARED` — without
   it the node never offers the adapter and the controller's cluster-interface resolution returns
   `node %s does not have interface eth1`
2. a second `interfaces` block on each `zedcloud_edgenode`, `intf_usage =
   "ADAPTER_USAGE_APP_SHARED"` — `ADAPTER_USAGE_UNSPECIFIED` is rejected with
   `node %s: interface %s has no usage configured`
3. `shared_labels` left **unset on all three nodes** —
   the controller's shared-label check requires the set to match across every node, and "unset
   everywhere" is the only value that trivially satisfies it

### 2.3 Naming: kernel-global, not per-state

Bridge and tap names live in the host's network namespace, not in Terraform state, and are
capped at `IFNAMSIZ` = 16 bytes. Two concurrent runs with the same `run_id` collide on the
*host*, not merely in Zedcloud — and the failure mode is a half-configured bridge, not a clean
error.

`test/e2e-cluster/vars.tf` therefore validates `length(var.run_id) <= 9`, which keeps
`ctap-<run_id>-N` at ≤ 16, and uses short `cbr-` / `ctap-` / `cdhcp-` prefixes rather than the
`tf_enc_` prefix used for cloud objects.

### 2.4 Open question: one bridge or three?

Andrei's working config (CI-834 state dump) has **three** bridges and three taps per node:
`zedamigo_bridge.BRIDGE_{201,301,401}` + `zedamigo_tap.TAP_{A,B,C}_{201,301,401}` +
`zedamigo_dhcp_server.DHCP_{201,301,401}`, matching the four-interface node in the ticket's
HCL (eth0 management + eth1/eth2/eth3 `APP_SHARED`), with `cluster_interface = "eth1"`.

Nothing in the validation code requires more than the one interface named in
`cluster_interface`, so `test/e2e-cluster/` uses a single bridge on eth1. Whether the extra two
segments are load-bearing for k3s/kubevirt in practice, or just an artefact of Andrei's wider
test config, is unresolved. **Ask him** — this is the cheapest of the four open questions to
close and the answer changes `net_cluster.tf` materially.

### 2.5 Unverified schemas

`terraform-provider-zedamigo` is external (`github.com/andrei-zededa/terraform-provider-zedamigo`,
not on any public registry, resolved via a filesystem mirror). `test/e2e` uses only
`host_reservation`, `disk_image`, `eve_installer`, `installed_edge_node`, `edge_node` and
`wait_until` — so the argument names for `zedamigo_bridge`, `zedamigo_tap` and
`zedamigo_dhcp_server` in `net_cluster.tf` are **inferred from the provider's naming
conventions**, not confirmed against the binary.

Confirm before the first apply:

```sh
tofu providers schema -json \
  | jq '.provider_schemas["localhost/andrei-zededa/zedamigo"].resource_schemas
        | with_entries(select(.key | test("bridge|tap|dhcp")))
        | map_values(.block.attributes | keys)'
```

or, faster, get Andrei's config.

---

## 3. Gate 3 — sudo

`zedamigo_tap` (`tap_resource.go:578-580`), `zedamigo_dhcp_server` (`dhcp_server_resource.go:520-523`),
`zedamigo_dhcp6_server` and `zedamigo_radv` self-invoke the provider binary under `sudo -n`.
`test/e2e` sets `use_sudo = false` and explains why it can: it uses no host networking
resources. `test/e2e-cluster` cannot.

`zededa-berlin-lab/hosts/h-m-dl20/configuration.nix` gives the `github-runner` user exactly
three NOPASSWD commands:

```
/run/current-system/sw/bin/ip
/run/current-system/sw/bin/kill
/run/current-system/sw/bin/taskset
```

A fourth entry for the bootstrapped provider binary path is required, in the lab repo, before
this applies in CI. The alternative, noted in §6.4 of the parent design doc, is to use the
`netns` variants, which route through `sudo ip netns exec` and therefore fit the existing `ip`
allowance — worth evaluating if amending sudoers is contentious, though it changes the topology
from "shared bridge" to "shared bridge inside a netns" and adds a layer to debug.

Locally this is a non-issue with passwordless sudo.

**This is a cross-repo change and should be raised as soon as §1 is settled**, because it is
the item most likely to sit in someone else's queue.

---

## 4. Gate 4 — host capacity

`test/e2e` reserves 4 vCPU / 8 GB / 20 GB for one `kvm` node and measures ~3m10s end to end on
h-m-dl20 (16 cores / 62 GB / 399 GB free).

Three kubevirt nodes are a different animal. A kubevirt EVE node runs k3s, kubevirt and
longhorn on top of EVE. `test/e2e-cluster/vars.tf` provisionally sets 4 vCPU / 16 GB / 60 GB,
i.e. **12 vCPU / 48 GB / 180 GB** for the set — most of the host. Those numbers are a guess
pending §1.3 question 3.

Three things to keep in view:

- `zedamigo_host_reservation` is **admission control, not a queue**. Over-subscription *fails
  the apply*; it does not wait. So a concurrent single-node `test/e2e` run will not fit
  alongside this, and the CI concurrency group must cover both directories, not just this one.
- The three `zedamigo_installed_edge_node` install VMs run in parallel up to Terraform's
  default `-parallelism=10`, **on top of** the reservations for the real VMs. True peak load is
  higher than the reservation implies. `-parallelism=1` is the escape hatch.
- Total wall time is dominated by cluster formation, not by the VMs:

  | step | estimate |
  |---|---|
  | `eve_installer` (built once, shared across nodes) | 15–60s |
  | 3 × install (parallel) | 3–6m |
  | 3 × onboard to `RUN_STATE_ONLINE` | 2–5m |
  | cluster create → all nodes `READY` | **10–25m** |
  | **total** | **~20–40m** |

  Hence `onboard_timeout = "35m"`, `cluster_ready_timeout = "30m"`, and a CI `timeout-minutes`
  well above the 60 used for the single-node workflow.

---

## 5. The Terraform itself

Everything above is why the HCL looks the way it does. The server-side checklist it satisfies,
all 400s:

the controller's cluster logic:

```go
const (
    defaultClusterPrefix          = "10.244.244.2/28"
    the required node count         = 3
    requiredMinimumOfZKSNodes     = 1
    preferableAmountOfZKSNodes    = 3
    notAvailableAmountOfZKSNodes  = 2
    maximumAmountOfZKSMasterNodes = 3
)
```

- the controller's cluster-field validation: 1 or 3+ nodes; **2 is never allowed** (no quorum, split-brain).
  Mirrored client-side as a validation on `var.node_count` so the failure is a plan error
  rather than an API round-trip.
- the controller's master-node check: 3+ nodes → at least 3 and at most 3 of type
  `EDGE_NODE_CLUSTER_NODE_TYPE_SERVER`. All three are SERVER, stated explicitly in
  `cluster.tf` rather than left to the schema default, because it is a hard requirement rather
  than a preference.
- the controller's prefix-overlap check: `cluster_prefix` must not overlap any network instance subnet on the
  member nodes. Hence `cluster_prefix` (`10.244.244.2/28`) and `cluster_bridge_subnet`
  (`10.77.77.0/24`) are disjoint, and both are variables so a collision with something else in
  the lab is a one-line fix.
- the controller's prefix assignment gives each node a distinct address out of the prefix (a `/28` yields
  ~13 usable node prefixes). Server-generated: the per-node `cluster_prefix` in a `nodes` block
  is `Computed` — do not set it.
- Seed node election and the 32-byte cluster token are server-side.

the controller's cluster-membership checks:

| check | error |
|---|---|
| all node IDs resolve | `not enough nodes for cluster formation, found nodes: %d` |
| not already clustered | `node %s is already part of a cluster` |
| `AdminState == DEVICE_REGISTERED` | `node %s is not registered` |
| `ClusterID == ""` | `node %s is a part of a cluster instance` |
| ≤ 1 `tie-breaker` node | `only one tie-breaker node is allowed in the cluster` |
| the controller's cluster-interface resolution resolves | `node %s does not have interface %s` |
| interface usage set | `node %s: interface %s has no usage configured` |
| no auto-deploy policy | `app instance %s is configured for auto-deployment` |
| no active app instance | `app instance %s is active on the node` |

`DEVICE_REGISTERED` is the subtle one. `ADMIN_STATE_ACTIVE` on the Terraform resource is a
*precondition* for onboarding, not evidence of it — the controller's registration handler flips
`AdminState` to `DEVICE_REGISTERED` only once EVE has actually completed the register
handshake. So `zedcloud_edgenode_cluster` must `depends_on` the
`zedamigo_wait_until.onboarded` barriers. Depending on `zedcloud_edgenode.EN` alone would
create the cluster against three records whose VMs are still booting, and fail.

### 5.1 Readiness

Create returns 200 once the config is accepted and pushed; the k3s/etcd cluster forms
afterwards, on the nodes, over minutes.

```
GET /api/v1/cluster/id/{id}/status -> EdgeNodeClusterStatus
  { id, runState, nodes: [ { id, name, readyState } ] }
readyState in EDGE_NODE_CLUSTER_NODE_READY_STATE_{UNSPECIFIED,READY,NOT_READY,UNKNOWN}
```

`UNKNOWN` means "in the cluster config but has not reported a Ready condition yet" — the normal
state for the first several minutes, not an error.

The provider generates **no client method** for this endpoint (only
Create/Read/Update/Delete/GetByName exist), so readiness is a `zedamigo_wait_until` curl probe
in `cluster.tf`, and the Go-side equivalent is a new `WaitForClusterNodesReady` in
`v2/testing/realnode.go` using the existing `apiGet` — same shape as
`WaitForAppInstanceSwState`.

### 5.2 Destroy ordering

`canClusterBeDeleted` blocks cluster deletion while an app instance is active on any member
node. Terraform's dependency graph handles this for resources inside the config; anything
created out of band wedges the destroy. The sweep step must also clear bridges and taps, since
`zedamigo_edge_node.Delete` treats a failed QMP quit as a warning and removes the state
directory anyway — orphan QEMU processes are an expected outcome, not an anomaly.

---

## 6. File-by-file plan

`test/e2e-cluster/` is a **sibling** of `test/e2e`, not a modification of it. Rationale: the
two differ in `use_sudo`, sudoers requirements, EVE flavour and capacity envelope, so
parameterising one directory over `node_count` would make the single-node path — which is
green today and gating PRs — carry the cluster path's risks. Cost is some duplication in
`terraform.tf` / `vars.tf` / the brand-model-project preamble; that is the cheaper trade.

Delivered (first cut, unvalidated):

| file | contents |
|---|---|
| `terraform.tf` | provider requirements; `use_sudo = true` and why |
| `vars.tf` | endpoints, `run_id` + IFNAMSIZ validation, `node_count` + quorum validation, `eve_tag` kubevirt warning, capacity, timeouts, subnets |
| `cloud_objects.tf` | brand, model (**two** NICs), project, two networks, `count`-ed edge-node records with the commented-out CI-834 workaround |
| `net_cluster.tf` | bridge + taps + DHCP; schemas flagged unverified |
| `nodes.tf` | the 5 zedamigo phases × `node_count`, shared installer, plus the kubevirt capability probe |
| `cluster.tf` | `zedcloud_edgenode_cluster` + readiness barrier, with the full server-side checklist inline |
| `outputs.tf` | ids for the acceptance suite, controller-assigned cluster state, per-node console logs |
| `e2e-cluster.tfrc` | CLI config resolving zedamigo from the filesystem mirror |
| `README.md` | prerequisites, local run, blockers, expected timings |

Still to write:

| file | change |
|---|---|
| `v2/testing/realnode.go` | `RealCluster(t) RealClusterInfo{ID, Name, NodeIDs, Suffix}` gated on `ZEDCLOUD_ACC_REAL_CLUSTER`; `WaitForClusterNodesReady(t, id, timeout)` polling `cluster/id/{id}/status` via `apiGet` |
| `v2/testing/fixtures_test.go` | **add tokens to `knownTokens`** — currently `{SUFFIX, NODE_ID}` only, and `TestFixtureTokensAreKnown` fails the build on anything else. Either `NODE_ID_1..3` or a single comma-joined `NODE_IDS` (see the `device_ids_csv` output) |
| `v2/resources/testdata/edgenode_cluster/create_real_nodes.tf` | new fixture |
| `v2/resources/edgenode_cluster_real_nodes_test.go` | `TestEdgeNodeCluster_RealNodes` |
| `.github/workflows/e2e-cluster.yml` | new job on `[self-hosted, berlin-lab]`; `workflow_dispatch` only at first (see §8) |
| `test/e2e/README.md` | cross-reference |
| `zededa-berlin-lab/hosts/h-m-dl20/github-runners.nix` | fourth sudoers entry (**other repo**) |

Note there is currently **no `make test-e2e` target** — §6.5 of the parent design doc proposed
one and it was never added; CI invokes `go test` directly. Adding `make test-e2e-cluster` here
is optional and probably better done once, for both directories, separately.

---

## 7. The acceptance test

`TestEdgeNodeCluster_RealNodes`, gated on `ZEDCLOUD_ACC_REAL_CLUSTER` and skipped otherwise,
consuming `cluster_id` / `device_ids_csv` from the `tofu` outputs via `ZEDCLOUD_TEST_*` env
vars — the same shape as `TestApplicationInstance_RealNode`.

### 7.1 CI-834: no spurious diff on cluster membership

The bug: when a node joins a cluster the controller populates `cluster_interface` and the
`edge_node_cluster` block on the device object. The provider declared both `Optional`, saw them
as a diff, and tried to null them out — which breaks the cluster.

Fixed in `c932b75d` (Feb 2026) by changing both to `Computed` in `v2/schemas/node.go:748,835`.
The fix is a two-line diff and is in the current HEAD lineage, so **this test is a regression
guard, not a bug reproduction**. It is still the highest-value assertion in the suite, because
the fix has never been exercised against a real cluster.

Assertion: after the cluster is formed, a refresh + plan over the three `zedcloud_edgenode`
resources is empty.

`cloud_objects.tf` deliberately ships the `lifecycle { ignore_changes = [cluster_interface,
edge_node_cluster] }` workaround **commented out**. Needing it *is* the failure.

One field was left behind by that fix and is worth asserting separately: `cluster_id` on the
node is still `Optional`, not `Computed` (`v2/schemas/node.go:739-743`). It is the obvious
candidate for residual churn in the same family.

### 7.2 CI-709: cluster-scoped instances lose `device_id`

Root cause, both halves:

*Server side* — the controller's network-instance create path. Create a network
instance with only `edge_node_cluster` and no `device_id`, and the controller:

1. sees an empty `device_id` alongside a non-empty cluster id,
2. resolves a designated node for that cluster,
3. assigns it as the instance's device, and
4. persists both the device id and the designated-node id.

So `device_id` comes back populated on every subsequent GET. The app-instance
create path does the same for app instances.

*Provider side* — `device_id` is `Optional` and **not `Computed`** on all three of
`network_instance`, `volume_instance` and `application_instance`
(`v2/schemas/network_instance.go:412`, `volume_instance.go:~245`, `application_instance.go:586`),
and `UpdateNetworkInstance` PUTs `zschema.NetworkInstanceModel(d)` wholesale with no per-field
change detection. Refresh writes the server's value into state, config has none, Terraform
plans to remove it, and the update sends it back empty. `app_type` is the same shape and worse:
it carries `Default: "APP_TYPE_UNSPECIFIED"`, so config always asserts a value and overwrites
whatever the cluster flow set — exactly the `APP_TYPE_VM -> APP_TYPE_UNSPECIFIED` line in the
ticket. The existing `DiffSuppressFunc` on `edge_node_cluster`
(`v2/schemas/schema_helpers.go:391-423`) only covers that block's `id`; it does nothing for
`device_id` or `app_type`.

The fix is `Optional + Computed` on both (and arguably `ConflictsWith: edge_node_cluster` on
`device_id`) — **out of scope here**. In scope: a test that **fails**, pinning the behaviour so
whoever fixes CI-709 has a green/red signal.

Assertion: create a cluster-scoped `zedcloud_network_instance` and
`zedcloud_application_instance` with `edge_node_cluster { id = <cluster id> }` and no
`device_id`; a second plan is empty and a second apply does not 409.

Note the create-time XOR guard (`v2/resources/network_instance.go:44-49`, identical in
volume/app): `either device_id or edge_node_cluster has to be specified`. Supplying only
`edge_node_cluster` is the supported cluster-scoped form, so the fixture is legitimate, not a
contrivance.

### 7.3 Cheap negative cases, adapted from `test_clusters.py`

Each is a plan/apply expected to fail with a specific 400, and each costs no additional VM
time once the three nodes are up:

- 2 nodes → the controller's cluster-field validation rejects (also caught client-side by the `node_count`
  validation, so this one needs a raw-API or fixture-level test)
- `cluster_interface = "eth0"` where eth0 is MANAGEMENT → should surface a usage/interface error
- `cluster_interface = "eth9"` → `node %s does not have interface eth9`
- an active app instance on a member node → `app instance %s is active on the node`
- `cluster_prefix` overlapping a network instance subnet → the controller's prefix-overlap check

### 7.4 Fixture tokens

`v2/testing/fixtures_test.go:22` declares `knownTokens = {"SUFFIX", "NODE_ID"}` and
`TestFixtureTokensAreKnown` walks `../resources/testdata` failing the build on any other
`__...__` placeholder. A three-node fixture needs either `__NODE_ID_1__`…`__NODE_ID_3__` or a
single `__NODE_IDS__`.

Prefer **`__NODE_IDS__`**, comma-joined, split in the fixture with `split(",", ...)`: one token
instead of three, one env var instead of three, and it does not hard-code the node count into
the token set. `outputs.tf` already exposes `device_ids_csv` for this.

Also relevant from UE-139: `Suffix()` returns `ZEDCLOUD_TEST_SUFFIX` **verbatim** with no
generated fallback, and `deployment_tag` / `tags` map values were deliberately left
un-tokenised. If this suite ever runs concurrently with the single-node one, that deferral
needs revisiting — cluster objects carry tags too.

---

## 8. CI integration

A separate workflow, `e2e-cluster.yml`, not a job added to `e2e.yml`. Reasons:

- It needs the fourth sudoers entry; `e2e.yml` must keep working without it.
- 20–40m vs ~10m, so it cannot share `timeout-minutes` sensibly.
- It consumes most of the host, so it must not run concurrently with `e2e.yml`.

Shape, following `e2e.yml`:

```yaml
name: E2E (3-node cluster)
on:
  workflow_dispatch:            # dispatch-only at first
jobs:
  e2e-cluster:
    runs-on: [self-hosted, berlin-lab]
    timeout-minutes: 120
    continue-on-error: true     # informational until it is reliably green
    concurrency:
      group: e2e-berlin-lab     # SHARED with e2e.yml -- capacity, not correctness
      cancel-in-progress: false
```

Two deltas from `e2e.yml` that matter:

- The `concurrency` group must be **shared with `e2e.yml`** and must not be per-`ref`.
  `e2e.yml` uses `e2e-${{ github.ref }}` with `cancel-in-progress: true`, which serialises a
  single branch and relies on there being one runner. That is adequate for one 4-vCPU node; it
  is not adequate once a job wants 12 vCPU and 48 GB. And `cancel-in-progress: false` here,
  because cancelling mid-apply leaves orphan VMs, bridges and taps.
- The teardown and sweep steps must cover three VMs and the host networking:

  ```sh
  tofu destroy -auto-approve
  pkill -f "zedamigo/edge_nodes/" || true
  ip link show | grep -E "cbr-|ctap-"    # must be empty
  ```

  UE-142 (nightly janitor) should learn about bridges and taps too.

Diagnostics on failure: install and run console logs for **all three** nodes, plus the final
`cluster/id/{id}/status` payload. The console logs are the only place EVE's silent failures
show up — and with a kubevirt image there is a new one to expect: a node that boots and
onboards fine but never reports `hvTypeKubevirt`.

Move to `pull_request` only after several consecutive green dispatch runs, and only alongside
UE-141 — a 40-minute job that consumes the whole lab host is a different proposition from a
10-minute one, and probably wants `pull_request` + a path filter or a label gate rather than
running on every PR.

---

## 9. Sequencing

§1.3 is resolved (`17.0.0-lts-k-amd64`), so the critical path is now much shorter than the
original framing suggested.

### 9.1 Lead with a 1-node cluster

**A single-node cluster reproduces both target bugs, and it needs none of §2, §3 or §4.**

the controller's cluster-field validation allows 1 node or 3+ (only 2 is rejected — no quorum), and
the controller explicitly permits it: a one-node cluster has one master. Critically,
**there is no node-count guard anywhere on the device mutation path** — `devproc.go:1567,1601`
set `dinfo.ClusterInterface` and the full `dinfo.EdgeNodeCluster` block, and
the controller's prefix assignment runs, identically for a 1-node cluster. So a single node gives:

- **CI-834's exact surface** — the controller populates `cluster_interface` and
  `edge_node_cluster` on the device object, which is the entire bug.
- **CI-709's exact surface** — a cluster-scoped network instance / app instance with
  `edge_node_cluster { id = ... }` and no `device_id`, where the controller assigns one and the
  provider tries to null it.

With one node there are no peers, so the shared L2 segment — and with it the bridge/tap
resources, `use_sudo`, and the sudoers change — drops out entirely. One third of the capacity,
too.

Worth trying first, before any of that: the controller's cluster-membership checks only requires the cluster
interface to have `intf_usage != ADAPTER_USAGE_UNSPECIFIED`, and `ADAPTER_USAGE_MANAGEMENT`
satisfies that. With no peer traffic, **`cluster_interface = "eth0"` on a single node may
simply work** — reducing the whole thing to the existing `test/e2e` harness plus a kubevirt tag
plus ~20 lines of HCL. Untested, and it is what §2.1 warns against for real deployments, but it
is a ten-minute experiment with a large payoff.

### 9.2 Order of work

| # | step | blocked by | can start now? |
|---|---|---|---|
| 1 | `test/e2e-cluster` with `node_count = 1`, `eve_tag = "17.0.0-lts-k-amd64"`, an SSH key set; confirm it onboards **and `kube_ready` passes** | — | **yes, do this first** |
| 2 | 1-node `zedcloud_edgenode_cluster`, `cluster_interface = "eth0"` (§9.1) | 1 | no — but ~10 min once 1 lands |
| 3 | If 2 is rejected on the interface: second NIC for one node — one `zedamigo_bridge` + one `zedamigo_tap`, no DHCP needed | 2 | no |
| 4 | **CI-834 regression assertion** on the 1-node cluster (§7.1) | 2 or 3 | no — *this is the unblock milestone* |
| 5 | **CI-709 reproduction** on the 1-node cluster (§7.2) | 4 | no |
| 6 | Get Andrei's working config; confirm zedamigo schemas + one-vs-three bridges (§2.4, §2.5) | — | **yes, in parallel from day one** |
| 7 | Confirm whether sudoers actually blocks you (§9.3) | — | **yes, in parallel** |
| 8 | Fourth sudoers entry on h-m-dl20 (§3) | 7 | only if 7 says it is needed (other repo, other queue) |
| 9 | Size the nodes for kubevirt; verify 3 fit (§4) | 1 | no |
| 10 | Finalise `net_cluster.tf` against the real schemas | 6 | no |
| 11 | Scale to 3 nodes → all onboarded → cluster `READY` | 8, 9, 10 | no |
| 12 | `WaitForClusterNodesReady` + `RealCluster` helpers | 2 | partly (write against the API, test later) |
| 13 | `__NODE_IDS__` token + fixture + `TestEdgeNodeCluster_RealNodes` (§7.1) | 12 | partly |
| 14 | Negative cases (§7.3) | 13 | no |
| 15 | `e2e-cluster.yml`, dispatch-only | 8, 11 | no |
| 16 | Promote to `pull_request` | 15 + several green runs | no |

Steps 1–5 are the unblock path and have **no external dependencies**. Steps 6–8 are the
external ones and should be raised on day one so they resolve in the background while the
1-node path proceeds.

### 9.3 "Local" means the lab, driven by hand — and sudoers may not apply

The iterative loop runs on the Berlin lab host
`h-m-dl20` over SSH with your own `tofu` state in your own directory — a hand-driven dev loop,
just not on a laptop. It is already x86_64 Linux with KVM, 16 cores / 62 GB, zedamigo 0.13.1
and the reservation tree in place, and it onboards a virtual EVE node in ~3 minutes today.

One thing to check early, because it can remove the biggest external dependency from the
critical path: **the three-entry NOPASSWD restriction in §3 applies to the `github-runner`
user.** If a personal account on that host has ordinary sudo, `use_sudo = true` works for
interactive runs and the lab-repo change is only needed when this goes into CI. Five minutes
to confirm, and it would reorder the whole plan.

---

## 10. Known provider defects (documented, not fixed here)

Found while reading the cluster code. Listed so they are not rediscovered mid-debug; each
wants its own ticket.

**Actively broken:**

1. `v2/resources/cluster.go:sanitizeClusterNodes` (~:283) type-asserts a `*schema.Set` to
   `[]interface{}`. `nodes` **is** a `TypeSet`, so `oldList` is always nil, `oldIDs` is always
   empty, and **every** node's `ClusterPrefix` is cleared on any `nodes` change. Written
   against an earlier list-typed schema; the eng-2255 node-replacement fix it implements is
   therefore inoperative. *Practical consequence: do not mutate `nodes` in place — recreate.*
2. `GetClusterByName` (~:108-120) appends diags on error but **does not return**, then
   dereferences `resp.GetPayload()` — a nil-pointer panic path.
3. `CreateCluster` line ~165: `if errs := GetCluster(ctx, d, m); err != nil` — tests the stale
   `err`, not `errs`. Same copy-paste at `network_instance.go:86` and `:231`.
4. `nodes` is a `TypeSet` whose element carries a `Computed` `cluster_prefix`. Set membership
   hashes over all element attributes, so a server-assigned prefix changes the hash and causes
   plan churn on `nodes` itself. Wants `TypeList`, or a `Set:` function hashing `id` only.
5. `v2/schemas/cluster_node.go:19` — `ClusterNodeModel` reads `d.GetOk("resourceLabels")`,
   camelCase, never matches. Harmless only because the function is unused.

**Missing / wrong:**

6. `name` has no `ForceNew` despite the API documenting it immutable.
7. Read prefers GET-by-name over GET-by-id for managed resources
   (`cluster.go:41-47`) — a rename out of band, or a name collision, silently binds to a
   different object.
8. Schema omits `cluster_type` (`EDGE_NODE_CLUSTER_TYPE_STANDART` — sic — vs `_ZKS`, and the
   swagger says STANDART is *"a cluster that has exactly 3 nodes"*), `manifest` / `manifest_url`,
   and the read-only `seed_node_id` / `seed_node_ip` / `admin_state` that
   `EdgeNodeClusterConfigSummary` already returns and `SetClusterResourceData` discards.
9. No client methods generated for `GET /v1/cluster` (query), `/id/{id}/status`,
   `/id/{id}/ports`, `/id/{id}/raw/status`, `PUT /id/{id}/nodes`, `/id/{id}/upgrade`,
   `/id/{id}/upgrade/status`, `GET /node/id/{id}`, `/available/projects`. The status one is the
   painful omission — there is no first-class way to know a cluster is ready.
10. `node_type`'s description claims SERVER is the only supported value;
    `EDGE_NODE_CLUSTER_NODE_TYPE_AGENT` is also valid per the swagger enum.
11. CI-709: `device_id` and `app_type` want `Optional + Computed` on net/vol/app instances
    (§7.2).
12. `cluster_id` on `zedcloud_edgenode` is still `Optional` where the CI-834 fix made its two
    siblings `Computed` (§7.1).
13. Zero test coverage: no cluster acceptance test, no `v2/resources/testdata/cluster*`, and
    the only cluster-adjacent unit test is
    `v2/schemas/edge_node_cluster_config_test.go` (from the CI-875 manifest-panic fix).
    `main.tf.example:543-568` has a commented-out 3-node example — the only HCL in the repo
    that has ever described a cluster.

Note also that `test/e2e/cluster_objects.tf` is **not** an edge-node-cluster fixture despite
the name — "cluster" there means the Zedcloud control plane (the alpha controller). Renaming it
`cloud_objects.tf`, as `test/e2e-cluster/` does, would save the next reader the confusion.
