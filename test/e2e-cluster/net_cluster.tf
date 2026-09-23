# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# ===========================================================================
# HOST NETWORKING — multi-node only
#
# Every resource here is gated on `var.node_count > 1`. A single-node cluster
# has no peers to reach, so it needs none of this: no bridges, no taps, no
# DHCP, and therefore no `use_sudo`.
#
# WHY IT IS NEEDED (for 2+ nodes):
#
# test/e2e gives each VM one NIC on QEMU SLIRP (user-mode NAT). SLIRP hands
# every guest the same isolated 10.0.2.15/24 behind its own NAT, with NO
# guest-to-guest path. Nodes on SLIRP cannot see each other regardless of how
# Zedcloud is configured, so the cluster interface needs a real shared segment.
#
# SHAPE — mirrors the upstream reference config
# (andrei-zededa/terraform-provider-zedamigo, examples/other/edge_node_cluster,
# host_networking.tf), which is the only known-working 3-node topology:
#
#   bridge A  10.99.1.1/24 + DHCP  -> each node's eth1  <- CLUSTER INTERFACE
#   bridge B  10.99.2.1/24 + DHCP  -> each node's eth2
#   bridge C  no IP, no DHCP, pure L2 -> each node's eth3
#
#   eth0 on every node stays on QEMU SLIRP for management, which preserves the
#   verified-working outbound path to the device API.
#
# Only eth1 carries cluster traffic. B and C exist because that is the shape
# known to work and because app/network instances in later tests need somewhere
# to live. Whether a single bridge would suffice is untested -- upstream uses
# three, so this uses three.
#
# ---------------------------------------------------------------------------
# SCHEMAS VERIFIED against zedamigo 0.13.1 on h-m-dl20 (2026-09-16) via
# `tofu providers schema -json`, and cross-checked against the upstream
# reference config (2026-09-17). Three earlier guesses were wrong:
#
#   zedamigo_bridge       addr   -> ipv4_address
#   zedamigo_tap          bridge -> master
#   zedamigo_dhcp_server  subnet/range_start/range_end/gateway
#                                -> netmask / pool{start,end} / router
#
# Also required and previously MISSING here, per upstream:
#   - taps need `group = "kvm"` so QEMU (running in that group) can open them
#   - bridges and taps need `state = "up"` -- they are not brought up implicitly
#   - `mtu` is a STRING ("1500"), not a number
# ---------------------------------------------------------------------------
# NAME LENGTH: bridge/tap names are kernel-global and capped at IFNAMSIZ (16).
# Longest here is "ctapC-<run_id>-N" = 8 + len(run_id), hence run_id <= 8.
# ===========================================================================

locals {
  cluster_net_enabled = var.node_count > 1
  cluster_net_count   = local.cluster_net_enabled ? 1 : 0
  tap_count           = local.cluster_net_enabled ? var.node_count : 0

  br_a = "cbrA-${var.run_id}"
  br_b = "cbrB-${var.run_id}"
  br_c = "cbrC-${var.run_id}"

  # Host holds .1; guests get .50-.99.
  a_host    = cidrhost(var.cluster_bridge_a_subnet, 1)
  a_first   = cidrhost(var.cluster_bridge_a_subnet, 50)
  a_last    = cidrhost(var.cluster_bridge_a_subnet, 99)
  a_mask    = cidrnetmask(var.cluster_bridge_a_subnet)
  a_preflen = split("/", var.cluster_bridge_a_subnet)[1]

  b_host    = cidrhost(var.cluster_bridge_b_subnet, 1)
  b_first   = cidrhost(var.cluster_bridge_b_subnet, 50)
  b_last    = cidrhost(var.cluster_bridge_b_subnet, 99)
  b_mask    = cidrnetmask(var.cluster_bridge_b_subnet)
  b_preflen = split("/", var.cluster_bridge_b_subnet)[1]
}

# ---------------------------------------------------------------------------
# Bridge A — the cluster interface segment (eth1).
#
# Taps attach via their own `master`, so `enslaved_interfaces` is deliberately
# left unset; setting both would give two resources a claim on the same
# relationship and produce perpetual diffs.
# ---------------------------------------------------------------------------

resource "zedamigo_bridge" "A" {
  count = local.cluster_net_count

  name         = local.br_a
  mtu          = "1500"
  state        = "up"
  ipv4_address = "${local.a_host}/${local.a_preflen}"
}

resource "zedamigo_tap" "A" {
  count = local.tap_count

  name   = "ctapA-${var.run_id}-${count.index + 1}"
  mtu    = "1500"
  state  = "up"
  group  = "kvm"
  master = zedamigo_bridge.A[0].name
}

resource "zedamigo_dhcp_server" "A" {
  count = local.cluster_net_count

  interface  = zedamigo_bridge.A[0].name
  server_id  = local.a_host
  netmask    = local.a_mask
  router     = local.a_host
  nameserver = "9.9.9.9"
  lease_time = 86400

  pool {
    start = local.a_first
    end   = local.a_last
  }

  depends_on = [zedamigo_tap.A]
}

# ---------------------------------------------------------------------------
# Bridge B — second app-shared segment (eth2). Addressed + DHCP, like A.
# ---------------------------------------------------------------------------

resource "zedamigo_bridge" "B" {
  count = local.cluster_net_count

  name         = local.br_b
  mtu          = "1500"
  state        = "up"
  ipv4_address = "${local.b_host}/${local.b_preflen}"
}

resource "zedamigo_tap" "B" {
  count = local.tap_count

  name   = "ctapB-${var.run_id}-${count.index + 1}"
  mtu    = "1500"
  state  = "up"
  group  = "kvm"
  master = zedamigo_bridge.B[0].name
}

resource "zedamigo_dhcp_server" "B" {
  count = local.cluster_net_count

  interface  = zedamigo_bridge.B[0].name
  server_id  = local.b_host
  netmask    = local.b_mask
  router     = local.b_host
  nameserver = "9.9.9.9"
  lease_time = 86400

  pool {
    start = local.b_first
    end   = local.b_last
  }

  depends_on = [zedamigo_tap.B]
}

# ---------------------------------------------------------------------------
# Bridge C — pure L2 (eth3). No host address, no DHCP server, exactly as
# upstream. The nodes' eth3 is still declared as a DHCP client in Zedcloud and
# simply never gets a lease; that is what the reference config does.
# ---------------------------------------------------------------------------

resource "zedamigo_bridge" "C" {
  count = local.cluster_net_count

  name  = local.br_c
  mtu   = "1500"
  state = "up"
}

resource "zedamigo_tap" "C" {
  count = local.tap_count

  name   = "ctapC-${var.run_id}-${count.index + 1}"
  mtu    = "1500"
  state  = "up"
  group  = "kvm"
  master = zedamigo_bridge.C[0].name
}
