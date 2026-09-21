# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# ---------------------------------------------------------------------------
# Control-plane objects. Every name carries var.run_id so runs cannot collide
# on the shared alpha enterprise.
#
# DIFFERENCES FROM test/e2e/cluster_objects.tf:
#   - node records are created with count = var.node_count
#   - MULTI-NODE ONLY: the model declares a second NIC (eth1) and each node
#     gets a second `interfaces` block for it
#
# For node_count = 1 this file is intentionally near-identical to test/e2e:
# one SLIRP NIC, eth0 as both management and cluster interface.
# ---------------------------------------------------------------------------

locals {
  # eth0 for a single node (no peers, so the management NIC suffices);
  # eth1 for multi-node (SLIRP has no guest-to-guest path). See vars.tf.
  cluster_interface = coalesce(
    var.cluster_interface,
    var.node_count > 1 ? "eth1" : "eth0",
  )

  # Extra NICs only exist, and are only declared to Zedcloud, for multi-node.
  # Declaring them for a single node would describe hardware the VM lacks.
  second_nic = var.node_count > 1

  # eth1/eth2/eth3 on bridges A/B/C, mirroring the upstream reference config.
  # Only eth1 carries cluster traffic; the other two exist because that is the
  # shape known to work and later tests need somewhere to put instances.
  extra_nics = local.second_nic ? ["eth1", "eth2", "eth3"] : []
}

resource "zedcloud_brand" "QEMU" {
  name        = "tf_enc_brand_${var.run_id}"
  title       = "tf_enc_brand_${var.run_id}"
  origin_type = "ORIGIN_LOCAL"
}

resource "zedcloud_model" "QEMU_VM" {
  name           = "tf_enc_model_${var.run_id}"
  title          = "tf_enc_model_${var.run_id}"
  origin_type    = "ORIGIN_LOCAL"
  brand_id       = zedcloud_brand.QEMU.id
  type           = upper(var.eve_arch)
  product_status = "production"
  state          = "SYS_MODEL_STATE_ACTIVE"

  attr = {
    # One model object serves every node, but nodes may differ in vCPU count
    # (default [6, 6, 4]), so advertise the largest.
    "Cpus"    = tostring(max(var.node_cpus...))
    "memory"  = "${var.node_mem_gb * 1024}M"
    "storage" = "${var.node_disk_mb}M"
  }

  # eth0 -- management uplink, QEMU SLIRP (user-mode NAT).
  io_member_list {
    assigngrp    = "eth0"
    cbattr       = {}
    cost         = 0
    logicallabel = "eth0"
    phyaddrs     = { Ifname = "eth0" }
    phylabel     = "eth0"
    usage        = "ADAPTER_USAGE_MANAGEMENT"
    usage_policy = {}
    ztype        = "IO_TYPE_ETH"
  }

  # eth1/eth2/eth3 -- MULTI-NODE ONLY. Backed by taps on bridges A/B/C (see
  # net_cluster.tf); eth1 is the cluster interface.
  #
  # The model MUST declare these or the node will not offer the adapter and
  # resolveClusterSysInterface fails with
  #   "node %s does not have interface eth1".
  #
  # Zedcloud's own virtual model (spec/template_l1_ZedVirtual-4G.json) declares
  # eth0 + eth1 for the same reason; upstream's QEMU model declares four.
  dynamic "io_member_list" {
    for_each = local.extra_nics

    content {
      assigngrp    = io_member_list.value
      cbattr       = {}
      cost         = 0
      logicallabel = io_member_list.value
      phyaddrs     = { Ifname = io_member_list.value }
      phylabel     = io_member_list.value
      usage        = "ADAPTER_USAGE_APP_SHARED"
      usage_policy = {}
      ztype        = "IO_TYPE_ETH"
    }
  }
}

resource "zedcloud_project" "PROJECT" {
  name  = "tf_enc_project_${var.run_id}"
  title = "tf_enc_project_${var.run_id}"
  type  = "TAG_TYPE_PROJECT"

  tag_level_settings {
    flow_log_transmission = "NETWORK_INSTANCE_FLOW_LOG_TRANSMISSION_UNSPECIFIED"
    interface_ordering    = "INTERFACE_ORDERING_ENABLED"
  }
}

# Management network: DHCP client on eth0 (SLIRP hands out 10.0.2.15).
resource "zedcloud_network" "mgmt_dhcp" {
  name       = "tf_enc_mgmt_${var.run_id}"
  title      = "tf_enc_mgmt_${var.run_id}"
  kind       = "NETWORK_KIND_V4"
  project_id = zedcloud_project.PROJECT.id
  mtu        = 1500

  ip { dhcp = "NETWORK_DHCP_TYPE_CLIENT" }
}

# Cluster network: DHCP client on eth1, served by zedamigo_dhcp_server on the
# host bridge. MULTI-NODE ONLY -- a single node uses eth0/mgmt_dhcp.
#
# A separate zedcloud_network object rather than reusing mgmt_dhcp, so the two
# interfaces are independently debuggable and so the cluster segment can later
# be switched to static addressing without touching the management path.
resource "zedcloud_network" "cluster_dhcp" {
  count = local.second_nic ? 1 : 0

  name       = "tf_enc_clusternet_${var.run_id}"
  title      = "tf_enc_clusternet_${var.run_id}"
  kind       = "NETWORK_KIND_V4"
  project_id = zedcloud_project.PROJECT.id
  mtu        = 1500

  ip { dhcp = "NETWORK_DHCP_TYPE_CLIENT" }
}

# ---------------------------------------------------------------------------
# The edge-node records, one per virtual node.
#
# SERIAL FLOW (Linux/QEMU): the serial is chosen HERE and flows DOWNWARDS --
# QEMU stamps it via `-smbios type=1,serial=<serial_no>` and EVE reads it as
# the hardware serial. The cloud object stays the source of truth.
#
# The controller's join key is literally "<onboarding_key>:<serialno>"
# (srvs/seine/devproc.go:GetOnCertCN), so these must be unique per node.
#
# admin_state = ADMIN_STATE_ACTIVE is MANDATORY: processRegRequest returns 404
# "Device is not activated" for a device still in DEVICE_CREATED.
# ---------------------------------------------------------------------------

resource "zedcloud_edgenode" "EN" {
  count = var.node_count

  name  = "tf_enc_en_${var.run_id}_${count.index + 1}"
  title = "tf_enc_en_${var.run_id}_${count.index + 1}"

  serialno       = "TF-ENC-${upper(var.run_id)}-${count.index + 1}"
  onboarding_key = var.onboarding_key
  model_id       = zedcloud_model.QEMU_VM.id
  project_id     = zedcloud_project.PROJECT.id
  admin_state    = "ADMIN_STATE_ACTIVE"

  # `net_dhcp` is set explicitly here, as upstream does. Leaving it implicit
  # invites a diff against whatever Zedcloud fills in.
  interfaces {
    intfname   = "eth0"
    intf_usage = "ADAPTER_USAGE_MANAGEMENT"
    net_dhcp   = "NETWORK_DHCP_TYPE_CLIENT"
    cost       = 0
    netname    = zedcloud_network.mgmt_dhcp.name
    ztype      = "IO_TYPE_ETH"
    tags       = {}
  }

  # eth1/eth2/eth3 -- MULTI-NODE ONLY. eth1 is the cluster interface.
  #
  # For a single node local.cluster_interface is "eth0", already declared above
  # with ADAPTER_USAGE_MANAGEMENT -- which satisfies the only usage check
  # validateNodesForCluster makes. So no extra blocks are needed, and adding
  # them would describe NICs the VM does not have.
  #
  # intf_usage must NOT be ADAPTER_USAGE_UNSPECIFIED -- validateNodesForCluster
  # rejects that with "node %s: interface %s has no usage configured".
  #
  # All three point at the SAME zedcloud_network, exactly as upstream does: the
  # network object only says "be a DHCP client", and which host bridge each NIC
  # actually lands on is decided by the QEMU tap wiring, not by Zedcloud.
  #
  # shared_labels is deliberately left unset on ALL nodes:
  # validateClusterNodesSharedLabels requires the shared-label set on the
  # cluster interface to match across every node, and "unset everywhere" is the
  # only value that trivially satisfies it.
  dynamic "interfaces" {
    for_each = local.extra_nics

    content {
      intfname   = interfaces.value
      intf_usage = "ADAPTER_USAGE_APP_SHARED"
      net_dhcp   = "NETWORK_DHCP_TYPE_CLIENT"
      cost       = 0
      netname    = zedcloud_network.cluster_dhcp[0].name
      ztype      = "IO_TYPE_ETH"
      tags       = {}
    }
  }

  # `uint64_value = "0"` is load-bearing: without it Zedcloud keeps reporting a
  # value the config does not set and every plan shows a diff. Upstream has the
  # same note.
  # Uses node_ssh_authorized_key, NOT the installer key, so the SSH key on a
  # running node can be changed without rebuilding it. See vars.tf.
  config_item {
    key          = "debug.enable.ssh"
    string_value = coalesce(var.node_ssh_authorized_key, var.edge_node_ssh_pub_key)
    uint64_value = "0"
  }

  # `tie-breaker` on at most ONE node -- validateNodesForCluster rejects more
  # with "only one tie-breaker node is allowed in the cluster", and clusterproc
  # records it as TieBreakerNodeId in the config pushed to EVE.
  #
  # Opt-in via var.tie_breaker_node_index. Andrei's suggestion is to tag the
  # smaller (4-CPU) node, which with the default node_cpus is index 2.
  tags = (
    var.tie_breaker_node_index != null && var.tie_breaker_node_index == count.index
    ? { tie-breaker = "true" }
    : {}
  )

  # CI-834 GUARD.
  #
  # Once a node joins a cluster the controller populates cluster_interface and
  # the edge_node_cluster block on the device object. Provider >= the CI-834 fix
  # (commit c932b75d) marks both Computed, so they no longer produce a diff --
  # this lifecycle block is NOT needed with a fixed provider and is left
  # commented out on purpose.
  #
  # Uncomment it ONLY to work around an older provider build. If you have to
  # uncomment it to get a clean plan, that is CI-834 regressing and the test in
  # v2/resources/edgenode_cluster_real_nodes_test.go should be failing.
  #
  # lifecycle {
  #   ignore_changes = [cluster_interface, edge_node_cluster]
  # }
}
