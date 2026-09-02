# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# ---------------------------------------------------------------------------
# Control-plane objects. Every name carries var.run_id so runs cannot collide
# on the shared alpha enterprise (which already holds ~233 devices and ~334
# projects belonging to other people).
# ---------------------------------------------------------------------------

resource "zedcloud_brand" "QEMU" {
  name        = "tf_e2e_brand_${var.run_id}"
  title       = "tf_e2e_brand_${var.run_id}"
  origin_type = "ORIGIN_LOCAL"
}

resource "zedcloud_model" "QEMU_VM" {
  name           = "tf_e2e_model_${var.run_id}"
  title          = "tf_e2e_model_${var.run_id}"
  origin_type    = "ORIGIN_LOCAL"
  brand_id       = zedcloud_brand.QEMU.id
  type           = upper(var.eve_arch)
  product_status = "production"
  state          = "SYS_MODEL_STATE_ACTIVE"

  attr = {
    "Cpus"    = tostring(var.node_cpus)
    "memory"  = "${var.node_mem_gb * 1024}M"
    "storage" = "${var.node_disk_mb}M"
  }

  # One management NIC. The VM is created with a single QEMU SLIRP interface,
  # so declaring more here would describe hardware that does not exist.
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
}

resource "zedcloud_project" "PROJECT" {
  name  = "tf_e2e_project_${var.run_id}"
  title = "tf_e2e_project_${var.run_id}"
  type  = "TAG_TYPE_PROJECT"

  tag_level_settings {
    flow_log_transmission = "NETWORK_INSTANCE_FLOW_LOG_TRANSMISSION_UNSPECIFIED"
    interface_ordering    = "INTERFACE_ORDERING_ENABLED"
  }
}

resource "zedcloud_network" "mgmt_dhcp" {
  name       = "tf_e2e_net_${var.run_id}"
  title      = "tf_e2e_net_${var.run_id}"
  kind       = "NETWORK_KIND_V4"
  project_id = zedcloud_project.PROJECT.id
  mtu        = 1500

  ip { dhcp = "NETWORK_DHCP_TYPE_CLIENT" }
}

# ---------------------------------------------------------------------------
# The edge-node record.
#
# SERIAL FLOW (Linux/QEMU): we choose the serial here and it flows DOWNWARDS --
# QEMU stamps it via `-smbios type=1,serial=<serial_no>` and EVE reads it as the
# hardware serial. The cloud object stays the source of truth.
#
# Do NOT use the macOS "soft-serial flip" (serialno = ...INSTALL.soft_serial);
# that exists only because Apple's Virtualization.framework offers no SMBIOS
# serial to stamp.
# ---------------------------------------------------------------------------

resource "zedcloud_edgenode" "EN" {
  name  = "tf_e2e_en_${var.run_id}"
  title = "tf_e2e_en_${var.run_id}"

  serialno       = "TF-E2E-${upper(var.run_id)}"
  onboarding_key = var.onboarding_key
  model_id       = zedcloud_model.QEMU_VM.id
  project_id     = zedcloud_project.PROJECT.id
  admin_state    = "ADMIN_STATE_ACTIVE"

  interfaces {
    intfname   = "eth0"
    intf_usage = "ADAPTER_USAGE_MANAGEMENT"
    cost       = 0
    netname    = zedcloud_network.mgmt_dhcp.name
    ztype      = "IO_TYPE_ETH"
    tags       = {}
  }

  # Lets us shell into EVE for post-mortems once it is onboarded.
  config_item {
    key          = "debug.enable.ssh"
    string_value = var.edge_node_ssh_pub_key
    uint64_value = "0"
  }

  tags = {}
}
