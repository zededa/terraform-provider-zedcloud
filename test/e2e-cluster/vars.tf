# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# ---------------------------------------------------------------------------
# Endpoints. These are TWO DIFFERENT hostnames -- mixing them up is the classic
# failure:
#   - zedcloud_url      -> the CONTROL/REST API the zedcloud provider calls
#   - eve_cluster_host  -> the DEVICE API, written into EVE's /config/server
#
# Verified on alpha:
#   zedcontrol.alpha.zededa.net   control API, token works
#   zedcloud.alpha.zededa.net     device API, /api/v2/edgedevice/ping -> 200
#   zedcontrold.alpha.zededa.net  DOES NOT RESOLVE -- do not use
# ---------------------------------------------------------------------------

variable "zedcloud_url" {
  type        = string
  description = "Control/REST API host, e.g. zedcontrol.alpha.zededa.net"
  default     = "zedcontrol.alpha.zededa.net"
}

variable "zedcloud_token" {
  type        = string
  description = "Zedcloud API token (TF_VAR_zedcloud_token)"
  sensitive   = true
}

variable "eve_cluster_host" {
  type        = string
  description = "Device API host baked into EVE /config/server"
  default     = "zedcloud.alpha.zededa.net"
}

# ---------------------------------------------------------------------------
# Identity / isolation
# ---------------------------------------------------------------------------

variable "run_id" {
  type        = string
  description = <<-EOT
    Unique per-run suffix for every created object. In CI this is "gh<run_id>".

    IMPORTANT for this config specifically: bridge and tap names are
    KERNEL-GLOBAL, not namespaced per Terraform state. Two concurrent runs with
    the same run_id will collide on the host's network interfaces, not just in
    Zedcloud. Interface names are also length-limited (IFNAMSIZ = 16), which is
    why the bridge/tap names use short prefixes and run_id is capped below.
  EOT

  validation {
    # Longest generated name is "ctapC-<run_id>-N" = 8 + len(run_id) <= 16.
    condition     = length(var.run_id) <= 8
    error_message = "run_id must be <= 8 chars: bridge/tap names are kernel-global and capped at IFNAMSIZ (16)."
  }
}

variable "onboarding_key" {
  type        = string
  description = "Onboarding key (CN of EVE's default onboard cert)"
  default     = "5d0767ee-0547-4569-b530-387e526f8cb9"
}

variable "node_ssh_authorized_key" {
  type        = string
  description = <<-EOT
    SSH public key for the `debug.enable.ssh` CONFIG ITEM only.

    Deliberately separate from var.edge_node_ssh_pub_key, which is baked into
    the installer image. Changing the installer key replaces the installer, the
    install and therefore the whole node. Changing this one updates only
    `zedcloud_edgenode`.

    *** IMPORTANT LIMITATION, measured 2026-09-17. ***

    This works to change the key on an ALREADY-RUNNING node: set it, apply, and
    EVE picks the new key up within a minute or two. It does NOT appear to take
    effect on a FRESH node, where EVE's sshd uses the `authorized_keys` baked
    into /config by the installer. On a newly built node the probe key was
    refused for 25 minutes while the installer key worked immediately:

      probe key (config item) -> Permission denied (publickey)
      installer key           -> INSTALLER_KEY_OK

    So for a clean build, make var.edge_node_ssh_pub_key the TARGET-SIDE probe
    key and leave this null. Use this only as the no-rebuild escape hatch for a
    node that is already up.

    THIS SHOULD BE A KEY WHOSE PRIVATE HALF LIVES ON THE ZEDAMIGO TARGET,
    because the kube-readiness probe runs there and SSHes into each node from
    there. Relying on a forwarded agent does not work in practice -- see
    var.node_ssh_private_key_file.

      ssh-keygen -t ed25519 -N "" -f ~/.ssh/eve_probe_ed25519   # on the target
      TF_VAR_node_ssh_authorized_key="$(ssh lab 'cat ~/.ssh/eve_probe_ed25519.pub')"

    Null falls back to var.edge_node_ssh_pub_key.
  EOT
  sensitive   = true
  default     = null
}

variable "node_ssh_private_key_file" {
  type        = string
  description = <<-EOT
    Private key file ON THE ZEDAMIGO TARGET that the kube-readiness probe uses
    to SSH into each EVE node. Null = rely on the target's default keys / agent.

    Why not agent forwarding: `forward_agent = true` is set on the provider, but
    a probe run on h-m-dl20 found `SSH_AUTH_SOCK` unset and the host has no keys
    of its own, so every attempt failed `Permission denied (publickey)` for the
    probe's full 25-minute budget. An explicit key on the target removes that
    whole dependency -- and it is also what CI will need, where there is no
    agent to forward.
  EOT
  default     = "~/.ssh/eve_probe_ed25519"
}

variable "edge_node_ssh_pub_key" {
  type        = string
  description = <<-EOT
    SSH public key baked into the EVE INSTALLER image.

    *** THIS SHOULD BE THE TARGET-SIDE PROBE KEY. ***

    EVE's sshd on a freshly installed node uses the authorized_keys baked in
    here, NOT the `debug.enable.ssh` config item (see
    var.node_ssh_authorized_key). Since the kube-readiness probe runs on the
    zedamigo target, the key baked in here must be one the target holds:

      ssh-keygen -t ed25519 -N "" -f ~/.ssh/eve_probe_ed25519   # on the target
      TF_VAR_edge_node_ssh_pub_key="$(ssh lab 'cat ~/.ssh/eve_probe_ed25519.pub')"

    Using a laptop-side key here is what caused the probe to fail authentication
    for its entire budget on an otherwise healthy node.

    Changing this replaces the installer and rebuilds every node. To change the
    key on an already-running node without a rebuild, use
    var.node_ssh_authorized_key instead.

    *** THIS IS NOW MANDATORY, not a convenience. ***

    The kube-readiness barrier polls each node OVER SSH for
    /var/lib/all_components_initialized -- there is no cloud-side signal that
    works (see nodes.tf). Without a key there is no way to know when EVE's
    Kubernetes stack is up, and creating the cluster before then "fails or
    stalls" (upstream's words).

    Generate one on the zedamigo target if needed:
      ssh-keygen -t ed25519 -N ""
  EOT
  sensitive   = true

  validation {
    condition     = length(trimspace(var.edge_node_ssh_pub_key)) > 0
    error_message = "edge_node_ssh_pub_key is required: the kube-readiness probe polls each node over SSH."
  }
}

# ---------------------------------------------------------------------------
# Node shape
# ---------------------------------------------------------------------------

variable "node_count" {
  type        = number
  description = <<-EOT
    Number of virtual EVE nodes to stand up and cluster.

    Zedcloud enforces this hard (srvs/seine/cluster/clusterproc.go):
      requiredAmountOfNodes = 3
      1 node  -> allowed (single-node cluster, 1 master)
      2 nodes -> rejected by validateClusterFields / validateMasterNodes
                 ("cluster must not have 2 number of nodes")
      3+      -> allowed; validateMasterNodes requires exactly 3 SERVER nodes

    NOTE: the upstream reference config carries a commented-out
    `tie-breaker = true` tag "for 2-node ENC configs", which suggests a
    tie-breaker path may exist. The validation code as read rejects 2
    unconditionally, so this config does too. Revisit only with evidence.
  EOT
  default     = 3

  validation {
    condition     = var.node_count == 1 || var.node_count >= 3
    error_message = "Zedcloud rejects 2-node clusters (no quorum). Use 1 or >= 3."
  }
}

variable "eve_tag" {
  type        = string
  description = <<-EOT
    lfedge/eve Docker tag. Pinned deliberately: a moving tag makes CI
    non-reproducible.

    *** MUST BE A KUBEVIRT FLAVOUR, AND THE FLAVOUR IS SPELLED `-k-`. ***

    The flavour was RENAMED from `-kubevirt-` to `-k-`. Searching Docker Hub for
    "kubevirt" finds only the old, abandoned naming (newest: 15.10.0) and leads
    to the wrong conclusion that no current build exists. It does:

      16.0.0-lts-k-amd64   16.0.1-lts-k-amd64   16.0.2-lts-k-amd64
      17.0.0-lts-k-amd64   17.0.0-lts-k-arm64

    EVE master's Makefile agrees -- the kube packages are gated on
    `ifeq ($(HV),k)`, not on HV=kubevirt.

    ARM64 EXISTS (17.0.0-lts-k-arm64), so clustering is not inherently
    x86_64-only and an Apple Silicon / vfkit loop is not ruled out.

    A `15.10.0-kubevirt-amd64` node was observed to onboard fine and then never
    start its Kubernetes stack (no k3s/kubelet/longhorn/etcd in the console log)
    and never report OptionalCapabilities. Do not use the old naming.

    Matches the upstream zedamigo reference config:
    examples/other/edge_node_cluster.
  EOT
  default     = "17.0.0-lts-k-amd64"
}

variable "eve_arch" {
  type    = string
  default = "amd64"
}

# Matches the upstream reference config (6 CPU / 16 GB / ~30 GB). A kubevirt
# node runs k3s + kubevirt + longhorn on top of EVE, so it is materially heavier
# than the kvm flavour test/e2e uses (4 / 8 / 20000).
#
# CAPACITY WARNING: node_count * node_cpus must fit the host.
# 3 x 6 = 18 vCPU does NOT fit h-m-dl20's 16 cores. Either drop to 5, accept
# oversubscription by clearing cpu_pinning, or use a bigger host.
# zedamigo_host_reservation is admission control, not a queue -- it FAILS the
# apply rather than waiting.
variable "node_cpus" {
  type        = list(number)
  description = <<-EOT
    vCPUs per node, indexed by node. Shorter than node_count is fine -- the last
    entry repeats for the remaining nodes.

    Default [6, 6, 4] is what fits h-m-dl20: it has 16 CPUs total, and 6+6+4=16.
    Upstream's own config uses 6 for all three on a larger machine.

    Per Andrei: "EVE-K should work with 4 as well, especially if you declare it
    as a tie-breaker node" -- see var.tie_breaker_node_index. So the 4-CPU node
    should be the tie-breaker if you use one.

    Note this leaves the host no dedicated CPUs. With cpu_pinning the VM threads
    are bound but the host can still schedule on them, so it works; it is simply
    tight.
  EOT
  default     = [6, 6, 4]

  validation {
    condition     = length(var.node_cpus) > 0
    error_message = "node_cpus must have at least one entry."
  }
}

variable "node_mem_gb" {
  type    = number
  default = 16
}

variable "tie_breaker_node_index" {
  type        = number
  description = <<-EOT
    0-based index of the node to tag `tie-breaker`, or null for none.

    validateNodesForCluster permits AT MOST ONE tie-breaker node
    ("only one tie-breaker node is allowed in the cluster"), and clusterproc
    records it as TieBreakerNodeId in the config pushed to EVE.

    Andrei suggests tagging the smaller (4-CPU) node, which with the default
    node_cpus is index 2. Left null by default because the exact effect on a
    3-node cluster is not something this config has verified -- turn it on
    deliberately, not by accident.
  EOT
  default     = null

  validation {
    condition = var.tie_breaker_node_index == null || (
      var.tie_breaker_node_index >= 0 && var.tie_breaker_node_index < var.node_count
    )
    error_message = "tie_breaker_node_index must be null or a valid 0-based node index."
  }
}

variable "node_disk_mb" {
  type        = number
  description = <<-EOT
    Size of the qcow2 node disk, used ONLY when var.node_disk_devs is empty.

    *** 30000 (upstream's number, for a 100 GB BLOCK DEVICE) IS A TRAP HERE. ***

    Measured on EVE 17.0.0 with a 30 GB qcow2: EVE's partition layout gave
    /persist just 5.1 GB, and Longhorn requires 25% of its storage free. With
    k3s + kubevirt + CDI + longhorn images pulled, 400 MB was left, so the node
    went into DiskPressure:

      Kubernetes node ... has pressure: KubeletHasDiskPressure
      Disk default-disk-... (/persist/vault/volumes) is not schedulable:
        CurrentAvailable = 419430400 <= MinimalAvailable = 1375671296

    That evicted multus, which broke every subsequent pod sandbox with
    "Multus: error getting pod: Unauthorized", which left virt-api and
    virt-controller cycling Error/ContainerCreating forever -- so EVE never
    advertised hvTypeKubevirt and the cluster could never be created.

    Prefer var.node_disk_devs (100 GB reserves). This default exists so the
    qcow2 fallback is not itself a trap.
  EOT
  default     = 102400
}

variable "cpu_pinning" {
  type        = bool
  description = <<-EOT
    Pin each VM to the CPUs its reservation claimed (upstream does this).
    Better and more predictable performance, but it makes the CPU count a hard
    constraint rather than a soft one -- with pinning on, 3 x 6 vCPU simply does
    not fit 16 cores. Turn off to oversubscribe a smaller host.
  EOT
  default     = true
}

variable "node_disk_devs" {
  type        = list(string)
  description = <<-EOT
    Host block devices to reserve and use as the node disks, one per node.
    Passed to zedamigo_host_reservation.devs and consumed via the edge node's
    `disk` block as type="device", format="raw".

    Upstream uses block devices rather than qcow2 files because Longhorn and
    etcd on a FILE are slow enough to matter for a kubevirt node.

h-m-dl20 has three volume groups on three separate physical drives, each
    with three reserves (`tree /var/lib/zedamigo/reservations/devs/`), sized:

      /dev/vg_sd{b,c,d}/reserve1   100 GB
      /dev/vg_sd{b,c,d}/reserve2   100 GB
      /dev/vg_sd{b,c,d}/reserve3   247 GB

    100 GB is comfortably above what the kube stack needs. Do NOT fall back to
    a small qcow2: a 30 GB image gave /persist only 5.1 GB and put the node into
    Longhorn DiskPressure, which silently prevents the kubevirt capability from
    ever being reported. See var.node_disk_mb.

    The default below takes reserve1 from each VG, i.e. ONE NODE PER PHYSICAL
    DRIVE -- per Andrei, "best to spread the nodes over different physical
    drives for best performance". Use reserve2/reserve3 for a concurrent run.

    (An earlier `sudo vgs` on the host returned nothing, which is what led this
    config to default to qcow2 files at first. The reservation tree is the
    interface to look at, not the LVM tooling.)

    Set to [] to fall back to plain qcow2 files (zedamigo_disk_image) -- fine
    for a quick single-node run, slow for a real cluster.

    Must be empty or exactly node_count entries.
  EOT
  default = [
    "/dev/vg_sdb/reserve1",
    "/dev/vg_sdc/reserve1",
    "/dev/vg_sdd/reserve1",
  ]

  validation {
    condition     = length(var.node_disk_devs) == 0 || length(var.node_disk_devs) >= var.node_count
    error_message = "node_disk_devs must be empty or contain at least node_count entries."
  }
}

# ---------------------------------------------------------------------------
# Timeouts
# ---------------------------------------------------------------------------

variable "onboard_timeout" {
  type        = string
  description = "Budget for EVE to boot, register and report RUN_STATE_ONLINE"
  default     = "35m"
}

variable "kube_ready_timeout" {
  type        = string
  description = <<-EOT
    Budget for every node to finish bringing up its Kubernetes stack, i.e. for
    /var/lib/all_components_initialized to exist inside the kube container.

    Upstream uses 20m at 10s intervals. This is a genuinely slow step -- k3s,
    kubevirt and longhorn all install and start -- and it is NOT the same thing
    as the node being online.

    MEASURED on h-m-dl20, EVE 17.0.0-lts-k, FIRST boot: ~33 minutes just to
    reach `kubectl get nodes` = Ready, with kubevirt/CDI/longhorn still
    converging after that. A 25m budget expired before kube had even begun
    pulling images (/persist 5.6 MB used, no kubectl in the container yet).
    60m is not generous here, it is realistic for a cold node.
  EOT
  default     = "60m"
}

variable "cluster_ready_timeout" {
  type        = string
  description = <<-EOT
    Budget for all nodes to reach EDGE_NODE_CLUSTER_NODE_READY_STATE_READY
    after the cluster object is created.
  EOT
  default     = "30m"
}

variable "reservations_path" {
  type        = string
  description = "zedamigo reservation tree root on the target"
  default     = "/var/lib/zedamigo/reservations"
}

# ---------------------------------------------------------------------------
# Remote zedamigo target
#
# zedamigo can drive a remote machine over SSH, which means tofu runs on your
# laptop and the VMs run on the lab host -- no scp of this directory, no state
# living on a shared box. Per Andrei this now works:
#
#   provider "zedamigo" {
#     use_sudo = true
#     target   = "172.16.2.40"
#     ssh { user = "andrei", use_agent = true, forward_agent = true }
#   }
#
# Leave zedamigo_target null to run locally (tofu executing ON the target),
# which is what the earlier runs did.
#
# BERLIN LAB REACHABILITY -- established 2026-09-17 by trying it.
#
# zedamigo dials the target directly:
#
#   Error: Can't resolve remote lib_path
#   can't resolve remote state dir: ssh dial 10.208.13.94:22:
#   dial tcp 10.208.13.94:22: i/o timeout
#
# It is a Go SSH client, so it does NOT read ~/.ssh/config and will never pick
# up a ProxyCommand. The VPN container's HTTP CONNECT proxy on :3128 is
# therefore useless to it -- remote mode needs a real TCP path to port 22.
#
# The provider's `ssh` block does offer two ways out (verified against the
# 0.13.1 schema):
#
#   proxy_jump -- an SSH jump host, i.e. `ssh -J`. The VPN container advertises
#     one on localhost:11022, which would make this the clean answer:
#         zedamigo_ssh_proxy_jump = "root@localhost:11022"
#     but that port was observed CLOSED even with the tunnel up, so the
#     container's sshd needs looking at first.
#
#   port -- combine with a local forward for a path that works today:
#         ssh -N -L 2222:127.0.0.1:22 lab      # via the CONNECT proxy
#         zedamigo_target   = "127.0.0.1"
#         zedamigo_ssh_port = 2222
#
# Local mode (tofu on the host) remains the proven fallback and survives a VPN
# drop if started under tmux. Every path in this config is already a path on the
# target, so switching modes changes nothing else.
# ---------------------------------------------------------------------------

variable "zedamigo_target" {
  type        = string
  description = "Host zedamigo should drive over SSH. Null = run locally on the target."
  default     = null
}

variable "zedamigo_ssh_user" {
  type        = string
  description = "SSH user for the remote zedamigo target."
  default     = null
}

variable "zedamigo_ssh_use_agent" {
  type        = bool
  description = "Use the local SSH agent to authenticate to the zedamigo target."
  default     = true
}

variable "zedamigo_ssh_forward_agent" {
  type        = bool
  description = <<-EOT
    Forward the SSH agent to the target. Needed because the kube-readiness probe
    runs ON the target and itself SSHes into each EVE node -- with agent
    forwarding it can use your key rather than needing one staged on the host.
  EOT
  default     = true
}

variable "zedamigo_ssh_port" {
  type        = number
  description = <<-EOT
    SSH port on the zedamigo target. Null = 22.

    Set this together with zedamigo_target = "127.0.0.1" to reach the lab
    through a local forward, which is the route that works today given the
    CONNECT-proxy-only path:

      ssh -N -L 2222:127.0.0.1:22 lab
  EOT
  default     = null
}

variable "zedamigo_ssh_insecure_ignore_host_key" {
  type        = bool
  description = <<-EOT
    Skip host-key verification for the zedamigo target. Escape hatch only.

    The failure this exists for looks alarming and usually is not: Go's
    `knownhosts` does not retry across host-key algorithms the way OpenSSH
    does, so if the server negotiates ssh-rsa while known_hosts only holds the
    ssh-ed25519 entry a previous `ssh` recorded, you get
    "knownhosts: key mismatch".

    Prefer recording every host key type instead:
      ssh-keygen -R '[127.0.0.1]:2222'
      ssh-keyscan -p 2222 127.0.0.1 >> ~/.ssh/known_hosts
  EOT
  default     = null
}

variable "zedamigo_ssh_proxy_jump" {
  type        = string
  description = <<-EOT
    SSH jump host for reaching the target, i.e. `ssh -J`. Null = direct.

    The Berlin VPN container advertises a jump host on localhost:11022, which
    would make this the clean answer:

      zedamigo_ssh_proxy_jump = "root@localhost:11022"

    That port was observed closed even with the tunnel up, so check the
    container's sshd before relying on it.
  EOT
  default     = null
}

# ---------------------------------------------------------------------------
# Cluster network
# ---------------------------------------------------------------------------

variable "cluster_interface" {
  type        = string
  description = <<-EOT
    Logical label of the interface that carries cluster traffic. Leave null to
    let the config pick: "eth0" for a single-node cluster, "eth1" for
    multi-node (which is what the upstream reference uses).

    Constraints (resolveClusterSysInterface / validateNodesForCluster, each a
    400): the label must resolve against the node's interfaces / bond_adapters /
    vlan_adapters, and that interface's intf_usage must not be
    ADAPTER_USAGE_UNSPECIFIED.

    MULTI-NODE cannot be eth0: eth0 is the QEMU SLIRP management uplink, which
    gives each guest an isolated 10.0.2.15/24 behind its own NAT with no
    guest-to-guest path.

    SINGLE-NODE can very likely be eth0 -- no peers, no traffic to carry, and
    ADAPTER_USAGE_MANAGEMENT satisfies the usage check. Still UNVERIFIED.
  EOT
  default     = null
}

# Three segments, mirroring the upstream reference config, which gives each node
# eth1/eth2/eth3 on bridges A/B/C. Only eth1 (bridge A) carries cluster traffic;
# B and C exist because that is the shape known to work, and app/network
# instances in later tests will want somewhere to live.
variable "cluster_bridge_a_subnet" {
  type        = string
  description = "Subnet for bridge A (eth1) -- the cluster interface segment. Must not overlap cluster_prefix."
  default     = "10.99.1.0/24"
}

variable "cluster_bridge_b_subnet" {
  type        = string
  description = "Subnet for bridge B (eth2), addressed + DHCP like A."
  default     = "10.99.2.0/24"
}

variable "cluster_prefix" {
  type        = string
  description = <<-EOT
    The edge-node-cluster overlay prefix. The provider schema defaults it to
    10.244.244.2/28 and assignClusterPrefixes hands each node a distinct address
    out of it (a /28 gives ~13 usable node prefixes).

    clusterPrefixOverlap rejects a prefix overlapping any network instance subnet
    on the member nodes -- so keep this away from the bridge subnets above.

    NOTE: the upstream reference config does NOT set this at all and relies on
    the default. Set to null here to do the same.
  EOT
  default     = null
}
