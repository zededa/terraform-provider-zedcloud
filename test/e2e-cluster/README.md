<!--
Copyright (c) Zededa, Inc.
SPDX-License-Identifier: Apache-2.0
-->

# e2e: 3-node edge-node cluster on virtual EVE nodes

Tracked by **[UE-156](https://zededa.atlassian.net/browse/UE-156)** under epic
[UE-131](https://zededa.atlassian.net/browse/UE-131). Exists to make
[CI-834](https://zededa.atlassian.net/browse/CI-834) and
[CI-709](https://zededa.atlassian.net/browse/CI-709) reproducible.

This is the sibling of [`../e2e`](../e2e), which stands up **one** virtual EVE
node. This directory stands up **three** and forms a `zedcloud_edgenode_cluster`
over them.

> **Status: WORKING for `node_count = 1`. Applied successfully 2026-09-17 on
> h-m-dl20; CI-834 verified against a real cluster.**
>
> ```
> cluster_id        = 6d249175-7f0e-4911-8f73-53b6de6c4998
> cluster_interface = "eth0"
> onboarded         = 1m22s   (5 attempts)
> kube_ready        = 52s     (6 attempts)
> cluster_ready     = 1m1s    ready=1/1, ..._READY_STATE_READY
> ```
>
> Two things this settles:
>
> - **The single-node `cluster_interface = "eth0"` shortcut works.** No second
>   NIC, no bridges, no taps, no `use_sudo`. That was the load-bearing
>   assumption behind starting at one node, and the API accepts it.
> - **CI-834's fix holds against a real cluster.** The controller populates
>   `clusterInterface: "eth0"` and the full `edgeNodeCluster` block
>   (`clusterPrefix 10.244.244.1/28`, `seedNodeIp 10.244.244.1`,
>   `isMaster: true`) on the device object; the provider refreshes them into
>   state and a subsequent plan reports **"No changes"**. Previously only ever
>   tested against simulated devices.
>
> Total time from `apply` to a formed cluster: **~4 minutes** with the 100 GB
> block device. The 33-minute k3s startup seen earlier was a symptom of the
> undersized disk, not normal behaviour.
>
> **CI-709 is also reproduced**, via `ci709_repro.tf`
> (`TF_VAR_repro_ci709=true`). A cluster-scoped network instance with no
> `device_id` produces a **permanent** diff, not a one-off:
>
> ```
> ~ zedcloud_network_instance.CI709[0] will be updated in-place
>     - device_id = "eb897188-..." -> null
> ```
>
> It survives being applied: the provider PUTs `device_id` empty, the controller
> re-resolves the designated node and re-persists it, and the next plan shows
> the same diff again. Indefinitely. That is the ticket's complaint exactly, and
> it explains the reported 409s.
>
> **Not yet done:** `node_count = 3` has never been applied, and the
> app-instance half of CI-709 (needs a datastore + image + app) is uncovered.
> See [What the first attempt got wrong](#what-the-first-attempt-got-wrong) for
> the seven traps that stood between the first attempt and this one.
>
> Also holding: `tofu fmt -check` clean, `tofu validate` passing for
> `node_count` 1 and 3, and the negative cases rejected as designed
> (`node_count = 2` by the quorum rule, an over-long `run_id` by the IFNAMSIZ
> rule).
>
> The shape of this config comes from
> [`examples/other/edge_node_cluster`](https://github.com/andrei-zededa/terraform-provider-zedamigo/tree/main/examples/other/edge_node_cluster)
> in the zedamigo repo — the only known-working multi-node topology. An earlier
> version of this directory diverged from it and did not work; see
> [What the first attempt got wrong](#what-the-first-attempt-got-wrong).

## What the first attempt got wrong

Recorded because both mistakes are easy to repeat and cost most of a day.

### 1. The kubevirt flavour is spelled `-k-`, not `-kubevirt-`

Searching Docker Hub for `kubevirt` finds only the **old, abandoned** naming,
newest `15.10.0-kubevirt-amd64`, and leads to the false conclusion that no
current kubevirt build exists — and that none exists for arm64 at all. What
actually exists:

```
16.0.0-lts-k-amd64   16.0.1-lts-k-amd64   16.0.2-lts-k-amd64
17.0.0-lts-k-amd64   17.0.0-lts-k-arm64
```

EVE master's Makefile agrees: the kube packages are gated on `ifeq ($(HV),k)`.

Consequences:

- `var.eve_tag` is now `17.0.0-lts-k-amd64`.
- A `15.10.0-kubevirt-amd64` node was observed to onboard, reach
  `RUN_STATE_ONLINE` / `ADMIN_STATE_REGISTERED`, and then sit for 20+ minutes
  with `optionalCapabilities: null` and **zero** k3s / kubelet / longhorn / etcd
  in its console log. The kube stack simply never started.
- **arm64 exists**, so clustering is not inherently x86_64-only and an Apple
  Silicon / vfkit loop is not ruled out. [UE-157](https://zededa.atlassian.net/browse/UE-157)
  was opened to spike *building* an arm64 kubevirt EVE; that is unnecessary.

### 2. Readiness must be asked of the node, not the controller

The first version polled the controller for
`OptionalCapabilities.hvTypeKubevirt`, since
`srvs/seine/devproc.go:validateClusterNodesInfo` gates cluster creation on
exactly that field. It is still a poor barrier:

- zedcloud only populates it `if dinfo.GetOptionalCapabilities() != nil`
  (`srvs/yamuna/devproc.go:deviceFillInfo`), so "not reported yet" and "not
  capable" are indistinguishable from outside — both are simply absent.
- it says nothing about *when* the stack is ready, and creating the cluster too
  early "fails or stalls" (upstream's words).

Upstream asks the node directly, over SSH:

```
eve exec kube ls -l /var/lib/all_components_initialized
```

at 10s intervals with a 20m budget. That is what `zedamigo_wait_until.kube_ready`
now does — **and it is why `var.edge_node_ssh_pub_key` is now mandatory rather
than a convenience.**

### 3. Smaller things, all from the reference config

- taps need `group = "kvm"` (so QEMU can open them), `state = "up"`, and `mtu`
  as a **string**; all three were missing
- `net_dhcp = "NETWORK_DHCP_TYPE_CLIENT"` set explicitly on every interface
- `-nic tap,...,model=virtio,mac=...` single-argument form, not `-netdev` +
  `-device`
- `eve_nuke_disks=vda,sda,vdb,sdb` in `grub_cfg`
- three bridges (A addressed + DHCP, B addressed + DHCP, C pure L2) feeding
  eth1/eth2/eth3 — not one bridge on eth1
- 6 vCPU / 16 GB / 30 GB per node with `cpu_pins`, not 4 / 16 / 60
- upstream uses **raw host block devices** rather than qcow2 files, because
  Longhorn and etcd on a file are slow enough to matter

One theory that proved **wrong**: the missing `swtpm` on h-m-dl20. The upstream
config uses no TPM at all — no `zedamigo_swtpm`, no `swtpm_socket`. The plan-time
`swtpm` warning really is irrelevant.

### 4. Disk size decides whether the kube stack ever works

The second apply used the right image, onboarded, ran k3s — and still never
reported `hvTypeKubevirt`. Cause, found by getting inside the node:

```
/dev/sda9   5.1G  2.0G  2.9G  41%  /hostfs/persist

Kubernetes node ... has pressure: KubeletHasDiskPressure
Disk default-disk-... (/persist/vault/volumes) is not schedulable:
  CurrentAvailable = 419430400 <= MinimalAvailable = 1375671296 (25% of Storage Max)
```

A **30 GB qcow2 gives `/persist` only 5.1 GB**, Longhorn requires 25% of its
storage free, and after pulling k3s + kubevirt + CDI + longhorn images there was
400 MB left. The node went into DiskPressure → multus was **Evicted** → every
subsequent pod sandbox failed with `Multus: error getting pod: Unauthorized` →
`virt-api` and `virt-controller` cycled `Error`/`ContainerCreating` indefinitely.
KubeVirt's CR read `Deployed` the whole time while nothing in it actually ran, so
EVE never advertised the capability.

**Use `var.node_disk_devs`** — the h-m-dl20 reserves are 100 GB
(`reserve1`/`reserve2`) and 247 GB (`reserve3`). The qcow2 fallback default is
now 100 GB rather than upstream's 30 GB, because upstream's 30 GB is paired with
a 100 GB *block device*, not with a 30 GB image.

### 5. Readiness markers are version-specific

EVE 17.0.0 has **no `/var/lib/all_components_initialized`** — the file upstream
polls. It uses per-component markers instead:

```
kubevirt_initialized   multus_initialized   debuguser-initialized
```

So the probe accepts either shape, and additionally requires the node not to be
in `DiskPressure` and not `NotReady` — because the markers appear while the
stack is still broken, which is exactly how the disk problem stayed invisible.

### 6. zedamigo does not propagate replacement — changing the installer needs `-replace` on the whole chain

`zedamigo_eve_installer` has no in-place update path, and the provider does not
mark its attributes `ForceNew`. So changing `authorized_keys` (or `tag`,
`cluster`, `grub_cfg`) makes Terraform plan an *update*, which the provider then
rejects outright:

```
Error: EVE-OS Installer Resource Update Error
Update is not supported.
```

Worse, `-replace` on the installer **alone** is a trap. Its `filename` is
deterministic, so nothing downstream sees a diff:

```
# zedamigo_eve_installer.eve         will be replaced, as requested
# zedamigo_installed_edge_node.INSTALL[0]  will be updated in-place   <-- WRONG
# zedamigo_edge_node.VM[0]                 will be updated in-place   <-- WRONG
```

The install never re-runs, so the node keeps the **old** baked SSH key and the
probe still cannot authenticate — an hour wasted for no change. Replace the
whole chain:

```sh
tofu apply \
  -replace='zedamigo_eve_installer.eve' \
  -replace='zedamigo_installed_edge_node.INSTALL[0]' \
  -replace='zedamigo_edge_node.VM[0]' \
  -auto-approve
```

A correct plan shows all three `will be replaced` plus
`zedamigo_wait_until.onboarded[0] must be replaced`, and **nothing** "updated
in-place". Add the extra indices for a multi-node run.

### 7. Never reinstall EVE against an already-registered device object

This one cost two runs and is the least obvious failure in the list.

Replacing the zedamigo chain (`-replace` on installer / installed_edge_node /
edge_node) gives the node a **new device certificate**, but leaves the
`zedcloud_edgenode` object untouched — and the controller already holds that
object in `DEVICE_REGISTERED` with the *old* cert hash bound to it.
`srvs/seine/seine.go:processRegRequest` answers re-registration of an already
registered device with `304 already registered`, so the reinstalled node never
completes onboarding and never receives its config.

The symptoms look like anything but a certificate problem:

```
runState = RUN_STATE_SUSPECT          # for hours, not minutes
/persist/vault                        # does not exist
kubectl                               # not present in the kube container
/var/log/k3s-install.log  (4 lines, frozen):
  2026-09-17 14:30:31 : Waiting for DeviceName from controller...
```

kube-init blocks on the device name, so the whole Kubernetes stack never starts
and no amount of waiting helps.

**Rule: a fresh VM needs a fresh device object.** Change `var.run_id` (which
renames every object and changes the serial) or destroy the device object first.
Do not `-replace` only the zedamigo side.

```sh
tofu destroy -auto-approve
pkill -f "$HOME/.local/state/zedamigo/edge_nodes/" || true   # see "Sweeping orphans" below
export TF_VAR_run_id="n$(date +%H%M%S)"
tofu apply -auto-approve
```

Check for orphans before starting: `pgrep -af qemu-system-x86_64`. A VM from an
earlier run stays alive across `destroy` and keeps holding its CPU/RAM
reservation, which on a 16-CPU host will make the next reservation fail.

### 8. What the console log is not

Two of the diagnoses above were initially wrong because they leaned on
"no k3s/kubevirt/longhorn strings in `serial_console_run.log`". **EVE does not
log kube activity to the serial console.** The log goes quiet shortly after the
base containerd shims start; k3s ran happily for 33 minutes leaving no trace in
it. Absence there is not evidence. Use SSH and `kubectl`.

## Two modes

One directory, two shapes, selected by `var.node_count`:

| | `node_count = 1` | `node_count = 3` |
|---|---|---|
| cluster interface | `eth0` (management NIC) | `eth1` (dedicated) |
| NICs per VM | 1 (QEMU SLIRP) | 4 (SLIRP + taps on bridges A/B/C) |
| bridges / taps / DHCP | **none** | 3 bridges, 9 taps, 2 DHCP |
| `zedamigo` `use_sudo` | `false` | `true` |
| host sudoers change | **not needed** | needed *for CI only* (see below) |
| capacity | 6 vCPU / 16 GB | 6+6+4 = **16 vCPU** / 48 GB |
| disks | qcow2 or one reserve | one reserve per physical drive |
| fits h-m-dl20 (16 cores) | yes | yes, exactly |
| reproduces CI-834 / CI-709 | **yes, both** | yes |

**Sized for h-m-dl20 on Andrei's advice.** Upstream runs 6/6/6 on a larger
machine; h-m-dl20 has 16 CPUs, so `var.node_cpus` defaults to **[6, 6, 4]**.
Per Andrei, "EVE-K should work with 4 as well, especially if you declare it as a
tie-breaker node" — hence `var.tie_breaker_node_index`, which is opt-in and
would be index 2 (the 4-CPU node). With `cpu_pinning = true` the sum is a hard
reservation constraint, so this uses the box exactly.

**Disks: h-m-dl20 does have block devices** — three volume groups on three
separate physical drives, three reserves each:

```
/dev/vg_sdb/reserve{1,2,3}
/dev/vg_sdc/reserve{1,2,3}
/dev/vg_sdd/reserve{1,2,3}
```

`var.node_disk_devs` defaults to `reserve1` from each VG, i.e. **one node per
physical drive**, which is what Andrei recommends for performance. Use
`reserve2` / `reserve3` for a concurrent run. Set the variable to `[]` to fall
back to qcow2 files — acceptable for a quick single-node run, slow for a real
cluster because of Longhorn and etcd.

(An earlier `sudo vgs` here returned nothing, which is why this config briefly
defaulted to qcow2. `tree /var/lib/zedamigo/reservations/devs/` is the interface
to look at, not the LVM tooling.)

**Start with `node_count = 1`.** It reproduces both target bugs and needs none
of the multi-node machinery. `validateClusterFields` allows 1 node or 3+ (only 2
is rejected — no quorum), `validateMasterNodes` says *"1 node cluster: 1 master
is fine"*, and there is **no node-count guard on the device mutation path** —
`devproc.go:1567,1601` populate `dinfo.ClusterInterface` and
`dinfo.EdgeNodeCluster` identically for one node as for three. So the controller
writes exactly the fields CI-834 is about.

The single-node `cluster_interface = "eth0"` choice rests on
`validateNodesForCluster` only requiring `intf_usage != ADAPTER_USAGE_UNSPECIFIED`,
which `ADAPTER_USAGE_MANAGEMENT` satisfies. With no peers there is no traffic to
carry. **This is the one untested assumption in the single-node path** — it is
the first thing an apply will confirm or refute.

## Why this is a separate directory and not `count = 3` on `../e2e`

Four things change, and three of them change the host's configuration rather
than just the Terraform:

| | `../e2e` | here (3 nodes) |
|---|---|---|
| NICs per VM | 1 (QEMU SLIRP) | 4 (SLIRP + taps on bridges A/B/C) |
| `zedamigo` `use_sudo` | `false` | **`true`** |
| host sudoers | 3 NOPASSWD entries suffice | needs a 4th *for CI only* |
| EVE flavour | `16.0.1-lts-kvm-amd64` | **must be a `-k-` build** |
| readiness signal | `RUN_STATE_ONLINE` | that **plus** kube components initialized, over SSH |
| host capacity | 4 vCPU / 8 GB | 16 vCPU / 48 GB + 3 block devices |

SLIRP is the crux. It gives every guest the same isolated `10.0.2.15/24` behind
its own NAT with **no guest-to-guest path**, so three SLIRP-only nodes cannot
see each other regardless of Zedcloud configuration. The cluster interface needs
a real shared L2 segment: host bridges, one tap per VM per bridge, DHCP on the
addressed ones.

## Lab environment — verified 2026-09-16

Measured on `h-m-dl20`, not assumed. Things that cost time if you rediscover them:

| fact | detail |
|---|---|
| **`ivan` has full sudo** | `(ALL : ALL) SETENV: NOPASSWD: ALL`. The three-entry restriction is on `github-runner`, so **multi-node works by hand today**; the lab-repo sudoers change is a CI-only prerequisite. |
| host | Linux 6.18.48 x86_64, 16 cores, 62 GB, 389 GB free, `/dev/kvm` writable, no VMs running |
| `zedamigo` | 0.13.1 `linux_amd64` installed at `~/.terraform.d/plugins/…` |
| reservation tree | present; `.lock` is `0664 root:zedamigo_reservations`, so the shared-tree permission bug is fixed |
| **`make` is MISSING** | so is `gcc`. Build the provider with `CGO_ENABLED=0 go build -o terraform-provider-zedcloud .` from `v2/` — this is why the CI workflow sets `CGO_ENABLED: "0"`. |
| present | `go`, `tofu`, `qemu-system-x86_64`, `qemu-img`, `docker`, `ip`, `flock`, `jq`, `curl`, `unzip` |

### Getting in

`nc -X connect` does **not** work — squid answers `HTTP/1.1 200 Connection
established` and macOS `nc` reports that as an error. That is the misparse the
lab README warns about. With `socat` absent, a stdlib Python `ProxyCommand`
works and installs nothing:

```
ssh -o ProxyCommand="/usr/bin/python3 /path/to/berlin-proxy.py %h %p" \
    -i ~/.ssh/ivan-berlin-lab ivan@10.208.13.94
```

where the script opens `CONNECT %h:%p` to `127.0.0.1:3128`, checks for ` 200`,
then shuttles stdin/stdout. `brew install socat` is the tidier long-term fix.

Also note `brew install tofu` installs **Tofu.app**, a macOS SSH/serial terminal.
OpenTofu is `brew install opentofu`. Local OpenTofu is optional anyway — the
lab host has it, and that is where this config runs.

### `dev_overrides` and `tofu init` do not mix

With `dev_overrides` set for `zededa/zedcloud`, `tofu init` still tries to
resolve that provider from the registry and fails with *"provider
registry.opentofu.org/zededa/zedcloud was not found in any of the search
locations"* — OpenTofu itself says *"Skip tofu init when using provider
development overrides."* But `init` is still needed to install `zedamigo` from
the filesystem mirror.

Two ways out. Either `init` first with a mirror-only tfrc and then
`plan`/`apply` with the `dev_overrides` one, or — simpler, and what was used to
validate this config — stage the locally built binary **into the mirror** and
skip `dev_overrides` entirely:

```sh
M=~/.terraform.d/plugins/registry.opentofu.org/zededa/zedcloud/99.0.0/linux_amd64
mkdir -p "$M"
cp ~/tfz/v2/terraform-provider-zedcloud "$M/terraform-provider-zedcloud_v99.0.0"
```

then include `registry.opentofu.org/zededa/zedcloud` in `filesystem_mirror` and
exclude it from `direct`. `terraform.tf` pins no version for `zedcloud`, so the
fake `99.0.0` resolves. Re-copy the binary after every rebuild.

### Two provider warnings you can ignore

`tofu plan` emits these at provider-configure time, regardless of which
resources the config declares — `../e2e` produces them too:

```
Warning: Can't find the `swtpm` executable.
Warning: Can't find the `genisoimage` executable ...
```

Neither binary is installed on `h-m-dl20`, and neither is needed: `swtpm` is
only used by `zedamigo_swtpm`, `genisoimage` only by
`zedamigo_cloud_init_iso`, and this config uses neither. `swtpm_socket` on
`zedamigo_edge_node` is likewise never set — EVE can use a vTPM for key storage
but onboards fine without one.

Ignore them and check the plan's resource shape instead — a single node should
produce no `zedamigo_bridge`, `zedamigo_tap` or `zedamigo_dhcp_server` at all.

### `tofu validate` does not check variable validations

`node_count = 2` passes `tofu validate` and is only rejected at `plan` time.
So a clean `validate` says nothing about the guard rails — use `plan` for those.

## Blockers

### 1. ~~The EVE image~~ — RESOLVED

Use `17.0.0-lts-k-amd64` (or `-arm64`). The `-k-` naming is the current
kubevirt flavour; see [What the first attempt got wrong](#1-the-kubevirt-flavour-is-spelled--k--not--kubevirt-).
The section below is kept because the *gate* it describes is still the thing
that has to be satisfied.

<details>
<summary>Why KubeVirt is required at all</summary>

### The KubeVirt gate

`srvs/seine/devproc.go:validateClusterNodesInfo` issues a **live**
`DeviceStatusReq` to every candidate node at cluster-create time and returns
`400 node %s does not support kubevirt` unless each one reports
`OptionalCapabilities.hvTypeKubevirt` (`device/devstatus.proto:522`). The `kvm`
and `xen` flavours do not report it.

`../e2e` pins `16.0.1-lts-kvm-amd64`, which **cannot** form a cluster — a `kvm`
flavour, not a `k` one.

Note the gate is checked **only at cluster create**, and it is a live query, so
every node must be up *simultaneously* or the create fails with
`timed out after %s querying cluster node status`. That is why
`zedcloud_edgenode_cluster` depends on all the per-node barriers rather than on
the node records.

The capability appears only once EVE has finished bringing up its Kubernetes
stack, which is why the barrier that matters is `kube_ready` (asked of the node
over SSH), not a poll of this field. See
[What the first attempt got wrong](#2-readiness-must-be-asked-of-the-node-not-the-controller).

</details>

### 2. ~~`zedamigo` schemas unverified~~ — RESOLVED

Verified against zedamigo 0.13.1 `linux_amd64` on `h-m-dl20`, 2026-09-16, via
`tofu providers schema -json`. `net_cluster.tf` now uses the real arguments.
Three earlier guesses were wrong, recorded here so the mistake is not repeated:

| resource | guessed | actual |
|---|---|---|
| `zedamigo_bridge` | `addr` | **`ipv4_address`** |
| `zedamigo_tap` | `bridge` | **`master`** |
| `zedamigo_dhcp_server` | `subnet` / `range_start` / `range_end` / `gateway` | **`netmask`** / **`pool { start, end }`** block / **`router`** |

Full verified surface:

```
zedamigo_bridge       name(req), ipv4_address, ipv6_address, mac_address, mtu,
                      netns, state, enslaved_interfaces(set)
zedamigo_tap          name(req), master, ipv4_address, mtu, netns, group,
                      owner, state, mover_status(computed)
zedamigo_dhcp_server  interface(req), netmask(req), router(req),
                      server_id(req), nameserver(req), lease_time, netns
                      + block pool { start, end }        (single)
                      + block static_route { to, via }   (list)
zedamigo_edge_node    serial_no(req), cpus, mem, disk_image_base, ovmf_vars_src,
                      extra_qemu_args(list), nic0, use_gvproxy, serial_type,
                      serial_port_server, cpu_pins, swtpm_socket, drive_if,
                      disk_size_mb, + block disk
```

Note `router` and `nameserver` are **required** with no way to omit them, which
is why `net_cluster.tf` points both at the bridge's own address — a well-formed
offer that leads nowhere, so the default route stays on eth0/SLIRP. If EVE
installs that router and it outranks SLIRP, onboarding will break in a way that
looks like a TLS fault; the fallback is a static/link-local eth1.

Full resource list, for reference: `bridge`, `cloud_init_iso`, `dhcp6_server`,
`dhcp_server`, `disk_image`, `edge_node`, `eve_installer`, `host_reservation`,
`installed_edge_node`, `internet_monitor`, `lag`, `local_datastore`,
`monitor_system_usage`, `netns`, `radv`, `swtpm`, `tap`, `virtual_machine`,
`vlan`, `vm`, `wait_until`.

Still worth doing: Andrei has a **working** 3-node config (linked from CI-834,
commit [`5d4fcc6e`](https://github.com/andrei-zededa/terraform-provider-zedcloud/commit/5d4fcc6e390d35de1ae7e5ae237528d2d2e83f06))
using **three** bridges — one per extra NIC, eth1/eth2/eth3 — where this config
uses a single bridge on eth1. Whether one shared segment suffices is UE-156 open
question #4, and asking him is cheaper than finding out by experiment.

### 3. ~~Host sudoers~~ — NOT a blocker for interactive work

`zedamigo_tap` and `zedamigo_dhcp_server` self-invoke the provider binary under
`sudo -n`, so the multi-node path needs sudo. **Verified 2026-09-16: `ivan` on
`h-m-dl20` has `(ALL : ALL) SETENV: NOPASSWD: ALL`** — the three-entry
restriction (`ip`, `kill`, `taskset`) applies to the `github-runner` user only.

So multi-node runs by hand today. The fourth NOPASSWD entry in
`zededa-berlin-lab/hosts/h-m-dl20/github-runners.nix` is needed **only** to run
this in CI, and is off the critical path. (Alternative for CI: the `netns`
variants, which route through the already-permitted `sudo ip`.)

Single-node needs no sudo at all — `use_sudo = var.node_count > 1`.

## Prerequisites

Host toolchain (same as `../e2e`, all hard-checked by the CI workflow):

```
go make tofu qemu-system-x86_64 qemu-img docker ip flock jq curl unzip
```

plus a writable `/dev/kvm`, nested virtualisation, and the `zedamigo` provider
installed into `~/.terraform.d/plugins/localhost/andrei-zededa/zedamigo/0.13.1/<os>_<arch>/`:

```sh
curl -fsSL https://raw.githubusercontent.com/andrei-zededa/terraform-provider-zedamigo/v0.13.1/install.sh \
  | bash -s -- --binary-only 0.13.1
```

Capacity: `node_count * node_cpus` vCPU and `node_count * node_mem_gb` GB.
Defaults are 12 vCPU / 48 GB, against h-m-dl20's 16 cores / 62 GB.
`zedamigo_host_reservation` is **admission control, not a queue** — if capacity
is unavailable the apply *fails* rather than waiting, and a concurrent `../e2e`
run will not fit alongside this one.

Note that the three install VMs (`zedamigo_installed_edge_node`) run in parallel
up to Terraform's default `-parallelism=10`, on top of the reservations for the
real VMs, so peak load is higher than the reservation implies. Use
`-parallelism=1` if the host thrashes.

## Two ways to run it

The VMs always live on `h-m-dl20` (that is where the reservation tree and the
block devices are). What changes is where `tofu` itself runs.

### Remote mode — tofu on your laptop

zedamigo drives the target over SSH, so nothing needs copying to the host and
the state stays with you. **Verified working 2026-09-17** from macOS against
h-m-dl20: full 14-resource plan, no errors.

Three non-obvious prerequisites, each of which produced a confusing failure
before being understood:

**1. A TCP path to port 22.** zedamigo uses a Go SSH client, so it does **not**
read `~/.ssh/config` and can never use a `ProxyCommand`. The Berlin lab is only
reachable via the VPN container's HTTP CONNECT proxy, which is invisible to it:

```
can't resolve remote state dir: ssh dial 10.208.13.94:22:
dial tcp 10.208.13.94:22: i/o timeout
```

Fix with a local forward (the container's own jump host on `:11022`, which
`ssh.proxy_jump` could use instead, was found closed):

```sh
ssh -N -L 2222:127.0.0.1:22 lab &     # 'lab' goes through the CONNECT proxy
export TF_VAR_zedamigo_target="127.0.0.1"
export TF_VAR_zedamigo_ssh_port=2222
```

**2. All host key types in `known_hosts`.** Go's `knownhosts` does not retry
across host-key algorithms the way OpenSSH does. If the server offers `ssh-rsa`
and `known_hosts` only holds the `ssh-ed25519` entry that a previous `ssh`
recorded, you get:

```
ssh: handshake failed: knownhosts: key mismatch
```

which reads like a MITM warning and is not one. Record every type:

```sh
ssh-keygen -R '[127.0.0.1]:2222'
ssh-keyscan -p 2222 127.0.0.1 >> ~/.ssh/known_hosts
```

(`var.zedamigo_ssh_insecure_ignore_host_key` and `known_hosts_file` exist in the
provider schema as escape hatches; recording the keys is the better answer.)

**3. TWO keys in the SSH agent**, not one:

| key | for |
|---|---|
| the lab key (`ivan-berlin-lab`) | zedamigo authenticating to h-m-dl20 |
| the key matching `var.edge_node_ssh_pub_key` | the `kube_ready` probe, which runs **on the target** and SSHes into each EVE node via the forwarded agent |

```sh
ssh-add --apple-use-keychain ~/.ssh/ivan-berlin-lab ~/.ssh/id_ed25519
ssh-add -l    # expect 2
```

Miss the second one and `forward_agent = true` forwards an *empty* agent —
`kube_ready` then fails for its full 25-minute budget with nothing obviously
wrong, which is the same shape of mystery as the original wrong-image failure.

**Still prefer local mode for the actual apply.** The forward rides a tunnel
that dropped three times in one afternoon; if it dies mid-apply zedamigo loses
its target. Local mode under `tmux` survives that.

### Local mode — tofu on the host

Leave `zedamigo_target` unset. This is what the earlier runs did, and it
survives a VPN drop if you start it inside `tmux`.

One-time setup, if `~/tfz` and the staged provider are not already there:

```sh
git clone https://github.com/zededa/terraform-provider-zedcloud.git ~/tfz
cd ~/tfz/v2 && CGO_ENABLED=0 go build -o terraform-provider-zedcloud .   # no make, no gcc

M=~/.terraform.d/plugins/registry.opentofu.org/zededa/zedcloud/99.0.0/linux_amd64
mkdir -p "$M" && cp ~/tfz/v2/terraform-provider-zedcloud "$M/terraform-provider-zedcloud_v99.0.0"
```

Then:

```sh
cd ~/tfz/test/e2e-cluster

# See "dev_overrides and tofu init do not mix" above -- this tfrc must use
# filesystem_mirror for BOTH providers, not dev_overrides.
# One-time: the tfrc is per-user (absolute mirror path) and gitignored.
cp -n e2e-cluster.tfrc.example e2e-cluster.tfrc   # then edit the mirror path
export TF_CLI_CONFIG_FILE="$PWD/e2e-cluster.tfrc"
export TF_VAR_zedcloud_url="zedcontrol.alpha.zededa.net"
export TF_VAR_zedcloud_token="<api token>"

# MANDATORY -- the kube-readiness probe polls the node over SSH.
# Generate one on the host first if needed: ssh-keygen -t ed25519 -N ""
export TF_VAR_edge_node_ssh_pub_key="$(cat ~/.ssh/id_ed25519.pub)"

export TF_VAR_node_count=1
export TF_VAR_run_id="l$(date +%H%M%S)"     # <= 8 chars, see below

tofu init
tofu plan          # single node: no bridges, taps or DHCP
tofu apply -auto-approve
```

Watch for these in order, because each is a separate failure mode:

1. `zedamigo_wait_until.onboarded[0]` → `runState=RUN_STATE_ONLINE`.
   If the install hangs instead, check
   `tofu output -raw install_console_logs` — an empty log means the
   `grub_cfg` / `serial_type` mismatch, which fails silently.
2. `zedamigo_wait_until.kube_ready[0]` → `kube components initialized`.
   **Slowest step, up to 25m** — k3s, kubevirt and longhorn installing. If it
   times out, get on the node and look rather than guessing:
   ```sh
   PORT=$(tofu output -json eve_ssh_ports | jq -r '.[0]')
   ssh -p "$PORT" root@localhost 'eve exec kube kubectl get nodes; eve exec kube kubectl -n kubevirt get pods'
   ```
   Zero k3s activity in `tofu output -raw run_console_logs` alongside a silent
   console is the signature of the wrong EVE image — check `var.eve_tag` really
   is a `-k-` build.
3. `zedcloud_edgenode_cluster.CLUSTER` → created. A 400 mentioning the
   interface means `cluster_interface = "eth0"` is not accepted after all; set
   `TF_VAR_cluster_interface=eth1`, which on a single node means supplying
   `TF_VAR_node_count=3`-style networking by hand — easier to just go to three
   nodes at that point.
4. `zedamigo_wait_until.cluster_ready` → `ready=1/1`.

Then the assertion this whole exercise exists for:

```sh
# CI-834 assertion.
#
# Run `tofu apply` FIRST, even when nothing changed. `plan` refreshes in memory
# but does not write state, so immediately after the initial apply the stored
# state still holds the pre-cluster values and the plan reports
# "Changes to Outputs: ~ cluster_assignment". That is the output catching up,
# not a resource diff -- but it makes -detailed-exitcode return 2 and looks
# like a failure.
tofu apply -auto-approve          # settles outputs, touches no infrastructure
tofu plan -detailed-exitcode      # NOW expect 0 / "No changes"

# For the acceptance test, assert on RESOURCE changes rather than the exit
# code, which conflates output-only drift with real drift:
tofu plan -out=/tmp/p.tfplan >/dev/null
tofu show -json /tmp/p.tfplan \
  | jq '[.resource_changes[]? | select(.change.actions != ["no-op"])] | length'
# expect 0

tofu output cluster_id
tofu output cluster_ready_wait
tofu output cluster_assignment     # sensitive: per-node controller-assigned state
```

Scaling up afterwards is `TF_VAR_node_count=3` plus a fresh `run_id`.

`run_id` is capped at 9 characters by a variable validation, because bridge and
tap names are **kernel-global** (not namespaced per Terraform state) and capped
at `IFNAMSIZ` = 16 bytes. Two concurrent runs with the same `run_id` collide on
the host's network interfaces, not merely in Zedcloud.

Teardown:

```sh
tofu destroy -auto-approve

# Sweep orphans. zedamigo_edge_node.Delete treats a failed QMP quit as a
# warning and removes the state dir anyway, so orphan QEMU processes are an
# expected outcome, not an anomaly.
# SCOPE THIS TO YOUR OWN HOME -- see the warning below.
pkill -f "$HOME/.local/state/zedamigo/edge_nodes/" || true
pgrep -u "$(whoami)" -f zedamigo/edge_nodes | wc -l    # expect 0
ip link show | grep -E "cbr-|ctap-"     # should be empty
```

**Destroy ordering matters.** `canClusterBeDeleted` blocks cluster deletion
while an app instance is active on any member node, so anything deployed onto
the cluster must be deactivated first. Terraform's dependency graph handles this
for resources in this config; anything created out of band will wedge the
destroy.

## Expected timings

Extrapolated from `../e2e`'s measured single-node numbers (12s installer /
1m34s install / 1m21s onboard = ~3m10s total, kvm flavour on h-m-dl20).
A kubevirt image is larger and boots slower, and cluster formation is k3s + etcd
converging across three nodes:

| step | estimate |
|---|---|
| `eve_installer` (once, shared) | 15–60s (longer on a cold image pull) |
| 3 × `installed_edge_node` (parallel) | 3–6m |
| 3 × onboard to `RUN_STATE_ONLINE` | 2–5m |
| cluster create → all nodes `READY` | **10–25m** |
| **total** | **~20–40m** |

Hence `var.onboard_timeout = "35m"` and
`var.cluster_ready_timeout = "30m"`, and why the CI job needs
`timeout-minutes` well above the 60 used for the single-node workflow.

## What the acceptance test asserts

### CI-834 — DONE, verified on a real 3-node cluster 2026-09-21

`v2/resources/edgenode_cluster_test.go`, three layers:

| test | needs the lab? | what it pins |
|---|---|---|
| `TestEdgeNodeSchema_ClusterFieldsAreComputed` | no | `cluster_interface` and `edge_node_cluster` are `Computed` in `v2/schemas/node.go` — the two lines `c932b75d` changed |
| `TestEdgeNode_ClusterFieldsProduceNoDiff` | no | the SDK's diff drops neither field when the config omits them, **plus a negative control** that reverts both to `Optional` and requires the removals to reappear |
| `TestEdgeNodeCluster_RealNode` | yes | the same, against the values the controller actually wrote on a live clustered node |

Run the last one against this harness's output:

```shell
ZEDCLOUD_ACC_REAL_NODE=1 \
ZEDCLOUD_TEST_NODE_NAME=$(tofu output -json device_names | jq -r '.[0]') \
  go test ./resources/ -run TestEdgeNodeCluster_RealNode -v
```

**`NewComputed` is not the bug, and conflating the two is the trap.** The
first version of the real-node test reported any non-identity diff entry and
duly "failed" on all three nodes of a healthy cluster with
`cluster_interface: "eth1" -> "" (removed=false computed=true)`. That is the
ordinary treatment of a `Computed` attribute — Terraform prints
"(known after apply)", Read returns the same value, the plan is empty.
CI-834 is a *drop*: `removed=true computed=false`. The test predicate
distinguishes them and the negative control keeps it honest.

Measured on `enc3a` (3 nodes, EVE `17.0.0-lts-k-amd64`, alpha):

```
non-no-op resource changes in the plan: 0

refreshed values, all three nodes:
  tf_enc_en_enc3a_1  iface=eth1  cluster=tf_enc_cluster_enc3a  prefix=10.244.244.3/28  master=true
  tf_enc_en_enc3a_2  iface=eth1  cluster=tf_enc_cluster_enc3a  prefix=10.244.244.1/28  master=true
  tf_enc_en_enc3a_3  iface=eth1  cluster=tf_enc_cluster_enc3a  prefix=10.244.244.2/28  master=true
```

Two things about reading that result:

- **`tofu state show` will show no cluster fields, and that is not a failure.**
  The on-disk state for `zedcloud_edgenode` was written when the node was
  created, *before* the cluster existed. `plan` refreshes in memory without
  persisting, so the values appear in the plan's `prior_state`, not in
  `terraform.tfstate`. Reading the stale file is what produced one false
  "not fixed" conclusion here. Use
  `tofu show -json <plan> | jq '.prior_state...'`.
- **Do not use `tofu plan -detailed-exitcode` as the assertion.** It returns 2
  for output-only drift as well as for real changes, so it reports 2 on a clean
  plan of this config. Count non-`no-op` entries in `resource_changes`.

`cloud_objects.tf` deliberately leaves the
`lifecycle { ignore_changes = [...] }` workaround commented out — needing it
*is* the failure.

### Still to write

1. **CI-709**: create a cluster-scoped `zedcloud_network_instance` and
   `zedcloud_application_instance` using
   `edge_node_cluster { id = <cluster id> }` with **no** `device_id`, then
   `plan` again. Today the controller assigns `device_id` server-side
   (`srvs/seine/netinstproc.go:153` persists `netinst.DeviceId = dev.Id` after
   resolving the designated node) while the provider declares `device_id` as
   `Optional` and PUTs the full model on update — so the controller's choice
   round-trips as empty, and the follow-up apply 409s. The network-instance
   half is fixed and merged (PR #234, `Optional + Computed` on `device_id`,
   `cluster_id` and `app_type`); the **application-instance half is still
   untested against a real cluster**.
2. Cluster `/status` reports all three nodes
   `EDGE_NODE_CLUSTER_NODE_READY_STATE_READY`. `zedamigo_wait_until.cluster_ready`
   already asserts this during apply; there is no Go-level equivalent because
   the provider has no client method for the cluster status endpoint.

Adding that fixture requires new tokens in `knownTokens` in
`v2/testing/fixtures_test.go` — it currently allows only `SUFFIX` and
`NODE_ID`, and `TestFixtureTokensAreKnown` fails the build on anything else.
Either `__NODE_ID_1__`…`__NODE_ID_3__` or a single comma-joined `__NODE_IDS__`
(see the `device_ids_csv` output).

## Server-side constraints encoded in this config

All from `zedcloud`; every one is a 400 if violated.

`srvs/seine/cluster/clusterproc.go`:

- `requiredAmountOfNodes = 3`. Node counts of 1 or 3+ are allowed; **2 is never
  allowed** (no quorum, split-brain). Enforced client-side by a validation on
  `var.node_count`.
- `validateMasterNodes`: for 3+ nodes, at least 3 and at most 3 of type
  `EDGE_NODE_CLUSTER_NODE_TYPE_SERVER`. All three nodes here are SERVER.
- `clusterPrefixOverlap`: `cluster_prefix` must not overlap any network instance
  subnet on the member nodes. Hence `cluster_prefix` (`10.244.244.2/28`) and
  the bridge subnets (`10.99.1.0/24`, `10.99.2.0/24`) are disjoint. Note
  `var.cluster_prefix` defaults to `null` so the provider's own default
  (`10.244.244.2/28`) applies, matching upstream.
- `assignClusterPrefixes` gives each node a distinct address out of
  `cluster_prefix`; a `/28` yields ~13 usable node prefixes. Server-generated —
  the per-node `cluster_prefix` in a `nodes` block is `Computed`, do not set it.
- The seed node and the 32-byte cluster token are chosen server-side.

`srvs/seine/devproc.go:validateNodesForCluster`:

- every node id resolves, is not already in a cluster, and has an empty
  `ClusterID`
- `node.AdminState == DEVICE_REGISTERED` — the node must have **onboarded**, not
  merely be `ADMIN_STATE_ACTIVE`. This is why `zedcloud_edgenode_cluster`
  depends on the `zedamigo_wait_until.onboarded` barriers and not just on the
  node records.
- at most one `tie-breaker`-tagged node
- `resolveClusterSysInterface` resolves `cluster_interface` against the node's
  `interfaces` / `bond_adapters` / `vlan_adapters`, and that interface's
  `intf_usage` is not `ADAPTER_USAGE_UNSPECIFIED`. Hence `eth1` is declared both
  in the model's `io_member_list` and as an `interfaces` block with
  `ADAPTER_USAGE_APP_SHARED`.
- no app instance with an auto-deployment policy, and **no active app instance**
  on any node

`validateClusterNodesSharedLabels`: shared labels on the cluster interface must
match across all nodes. `shared_labels` is left unset on every node, which is
the only value that trivially satisfies this.

## Provider gaps worth knowing while using this

Not fixed here (UE-156 is harness + reproduction only), but they shape the
config:

- **No client method for cluster status.** Only Create / Read / Update / Delete /
  GetByName are generated. `/v1/cluster/id/{id}/status`, `/ports`,
  `/raw/status`, `PUT .../nodes`, `/upgrade` and
  `/v1/cluster/node/id/{id}` have none — which is why readiness is a
  `zedamigo_wait_until` curl probe rather than a data source.
- **`seed_node_id` / `seed_node_ip` / `admin_state` are dropped.**
  `SetClusterResourceData` reads an `EdgeNodeClusterConfigSummary` but the
  schema has no fields for them. Read them off each device's computed
  `edge_node_cluster` block instead — that is what `outputs.tf` does.
- **`sanitizeClusterNodes` is broken.** It type-asserts a `*schema.Set` to
  `[]interface{}`, so the old-node set is always empty and **every** node's
  `ClusterPrefix` is cleared on any `nodes` change. Avoid mutating `nodes` in
  place; recreate the cluster instead.
- **`nodes` is a `TypeSet` containing a `Computed` `cluster_prefix`.** Set
  membership hashes over all element attributes, so server-assigned prefixes can
  cause plan churn on `nodes` itself. If you see that, it is this, not your
  config.
- **`name` has no `ForceNew`** despite being API-immutable, and Read prefers
  GET-by-name over GET-by-id for managed resources.
- `zedcloud_edgenode` **data source is unusable** for lookups (UE-143) — it
  reuses the resource schema, so a name-only lookup fails with
  `The argument "model_id" is required`. That is why the probes here go against
  the raw REST API.

### Sweeping orphans: scope the pattern to your own home

`zedamigo_edge_node.Delete` treats a failed QMP quit as a warning and removes
the state dir anyway, so leftover QEMU processes after a `destroy` are an
expected outcome rather than an anomaly. They must be swept, or they keep
holding their CPU/RAM reservation.

**But do not use the bare pattern:**

```sh
pkill -f "zedamigo/edge_nodes/"        # DON'T -- matches every user on the host
```

h-m-dl20 is shared, and *every* user's zedamigo state path contains
`zedamigo/edge_nodes/` -- `/home/andrei/.local/state/zedamigo/edge_nodes/...`
matches exactly as well as yours. Observed 2026-09-17: that pattern killed the
running VMs of another user's cluster, and only failed on two of them because
of Unix permissions:

```
pkill: killing pid 82214 failed: Operation not permitted
pkill: killing pid 82217 failed: Operation not permitted
```

It fails harmlessly for an unprivileged user, but **every interactive user on
this host has full NOPASSWD sudo**, and under `sudo` it would take out a
colleague's cluster with no warning. Use:

```sh
pkill -f "$HOME/.local/state/zedamigo/edge_nodes/" || true
pgrep -u "$(whoami)" -f zedamigo/edge_nodes | wc -l    # expect 0
```

The same bare pattern appears in `.github/workflows/e2e.yml:362` and
`test/e2e/README.md`. It is only safe there because the `github-runner` user
cannot signal other users' processes -- worth tightening anyway.

### Orphaned reservations block the next run, and the error blames the host

Sweeping the QEMU processes is not enough: the **reservation** survives them.
`zedamigo_host_reservation` releases its claim on destroy, so an apply that is
killed -- or a state file that is thrown away -- leaves the claim behind. The
tree is host-global rather than per-state, so the next run dies in admission
control:

```
Error: insufficient free CPUs under /var/lib/zedamigo/reservations/cpus/unit:
requested 6, only 2 free
```

with **no VM running and nothing in any state file**. Measured 2026-09-21: 14
of 16 CPUs and 48 GB of RAM were held by three dead reservation ids from
earlier runs of this very directory. The message names neither the holder nor
the claim path, so it reads as "the host is busy" when the host is idle.

The claims are plain text files, one per unit, and releasing one is truncating
it. Each records `<reservation-id> <user> <user> <host> <host> <directory>`:

```sh
for i in $(seq 0 15); do printf '%2s %s\n' "$i" "$(cat /var/lib/zedamigo/reservations/cpus/unit/$i)"; done
```

`./release-orphans.sh` does it safely -- it only truncates claims naming both
your user and your harness directory, skips everyone else's, and refuses to run
while any of your VM processes are alive. `DRY_RUN=1` to look first.

**Check this before blaming capacity.** `node_cpus` defaults to `[6, 6, 4]`,
exactly h-m-dl20's 16 cores, so a single stale claim is enough to fail the run.


## Files

| file | contents |
|---|---|
| `terraform.tf` | provider requirements; `use_sudo = var.node_count > 1` and why |
| `vars.tf` | endpoints, `run_id` + IFNAMSIZ validation, `node_count` + quorum validation, `eve_tag` (the `-k-` naming), **mandatory** SSH key, capacity, `cpu_pinning`, optional `node_disk_devs`, timeouts, subnets |
| `cloud_objects.tf` | brand, model (**four** NICs when multi-node), project, two networks, `count`-ed edge-node records with the commented-out CI-834 workaround |
| `net_cluster.tf` | bridges A/B/C + 9 taps + 2 DHCP servers; verified schemas; multi-node only |
| `nodes.tf` | the zedamigo phases × `node_count`, shared installer, file-or-block-device disks, and the SSH-based kube-readiness barrier |
| `cluster.tf` | `zedcloud_edgenode_cluster` and the readiness barrier, with the full server-side checklist |
| `outputs.tf` | ids for the acceptance suite, controller-assigned cluster state, per-node console logs |
| `e2e-cluster.tfrc.example` | Template CLI config: resolves zedamigo from the filesystem mirror, and documents why `dev_overrides` cannot be used here. Copy to `e2e-cluster.tfrc` (gitignored) and set the absolute mirror path. |
| `release-orphans.sh` | Releases reservations left behind by a killed apply, scoped to your own user and directory. Run it on the target when a run fails with "insufficient free CPUs". |
