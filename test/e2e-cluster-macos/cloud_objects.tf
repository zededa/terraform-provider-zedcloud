# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

locals {
  ssh_key_path = pathexpand(var.ssh_key_file)
  ssh_pub_key  = trimspace(file("${local.ssh_key_path}.pub"))

  tls_ca_path       = pathexpand(var.tls_ca_file)
  tls_ca            = file(local.tls_ca_path)
  object_signing_ca = file(pathexpand(var.object_signing_ca_file))

  # Every zedamigo_wait_until probe curls the control API, and that API is
  # served with a PRIVATE certificate. The Linux harness gets away with a bare
  # `curl -s` because alpha's certificate is publicly trusted; here curl
  # rejects it, prints nothing, and the probe fails forever with no hint --
  # measured as a 30-minute onboard timeout on a node that was ONLINE and
  # REGISTERED the whole time.
  #
  # --cacert rather than -k, so a genuinely wrong certificate still fails.
  curl_ca = "--cacert ${local.tls_ca_path}"

  # eth0 is both the management uplink and the cluster interface.
  #
  # Zedcloud allows a 1-node cluster (the controller: "1 node cluster: 1 master
  # is fine") and the controller's cluster-membership checks only requires the cluster interface's
  # intf_usage != ADAPTER_USAGE_UNSPECIFIED -- ADAPTER_USAGE_MANAGEMENT
  # satisfies that. With no peers there is no cluster traffic to carry, so the
  # single gvproxy NIC is enough. Proven on the Linux harness (node_count = 1).
  cluster_interface = "eth0"
}

resource "zedcloud_brand" "QEMU" {
  name        = "tf_encm_brand_${var.run_id}"
  title       = "tf_encm_brand_${var.run_id}"
  origin_type = "ORIGIN_LOCAL"
}

resource "zedcloud_model" "VFKIT_VM" {
  name           = "tf_encm_model_${var.run_id}"
  title          = "tf_encm_model_${var.run_id}"
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

  # One NIC only. No bridges or taps exist on macOS, so there is nothing for a
  # second adapter to attach to and declaring one would describe hardware the
  # VM does not have.
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
  name  = "tf_encm_project_${var.run_id}"
  title = "tf_encm_project_${var.run_id}"
  type  = "TAG_TYPE_PROJECT"

  tag_level_settings {
    flow_log_transmission = "NETWORK_INSTANCE_FLOW_LOG_TRANSMISSION_UNSPECIFIED"
    interface_ordering    = "INTERFACE_ORDERING_ENABLED"
  }
}

resource "zedcloud_network" "mgmt_dhcp" {
  name       = "tf_encm_mgmt_${var.run_id}"
  title      = "tf_encm_mgmt_${var.run_id}"
  kind       = "NETWORK_KIND_V4"
  project_id = zedcloud_project.PROJECT.id
  mtu        = 1500

  ip { dhcp = "NETWORK_DHCP_TYPE_CLIENT" }
}

# ---------------------------------------------------------------------------
# The edge-node record.
#
# *** THE SERIAL FLOWS UPWARDS HERE, NOT DOWNWARDS. ***
#
# On Linux/QEMU the serial is chosen in this resource and QEMU stamps it via
# `-smbios type=1,serial=...`, which EVE reads through dmidecode as the
# hardware serial. Apple's Virtualization.framework has no equivalent, so
# vfkit cannot set one; EVE instead derives a "soft serial" (a UUID from the
# device certificate) at install time. So the install must run FIRST and this
# record takes its serialno from the install's soft_serial output.
#
# That inverts the dependency: INSTALL -> EN -> VM, where on Linux it is
# EN -> INSTALL -> VM. No cycle, because zedamigo_edge_node.VM depends on both.
#
# admin_state = ADMIN_STATE_ACTIVE is MANDATORY: the controller's registration handler returns 404
# "Device is not activated" for a device still in DEVICE_CREATED.
# ---------------------------------------------------------------------------

resource "zedcloud_edgenode" "EN" {
  name  = "tf_encm_en_${var.run_id}"
  title = "tf_encm_en_${var.run_id}"

  serialno       = zedamigo_installed_edge_node.INSTALL.soft_serial
  onboarding_key = var.onboarding_key
  model_id       = zedcloud_model.VFKIT_VM.id
  project_id     = zedcloud_project.PROJECT.id
  admin_state    = "ADMIN_STATE_ACTIVE"

  interfaces {
    intfname   = "eth0"
    intf_usage = "ADAPTER_USAGE_MANAGEMENT"
    net_dhcp   = "NETWORK_DHCP_TYPE_CLIENT"
    cost       = 0
    netname    = zedcloud_network.mgmt_dhcp.name
    ztype      = "IO_TYPE_ETH"
    tags       = {}
  }

  # `uint64_value = "0"` is load-bearing: without it Zedcloud keeps reporting a
  # value the config does not set and every plan shows a diff.
  #
  # Same key as the installer's authorized_keys on purpose. The config item
  # does NOT take effect on a fresh node (sshd uses the baked-in file), but it
  # does on a running one, so setting both means the key keeps working after
  # EVE switches over.
  config_item {
    key          = "debug.enable.ssh"
    string_value = local.ssh_pub_key
    uint64_value = "0"
  }

  tags = {}

  # CI-834 GUARD -- deliberately commented out. With a provider carrying the
  # CI-834 fix, cluster_interface and edge_node_cluster are Computed and
  # produce no diff. Needing to uncomment this is CI-834 regressing.
  #
  # lifecycle {
  #   ignore_changes = [cluster_interface, edge_node_cluster]
  # }
}
