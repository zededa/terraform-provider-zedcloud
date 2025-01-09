resource "zedcloud_project" "test_tf_provider" {
  # required
  name = "test_tf_provider-create_edgenode"
  title = "title"

  # optional
  type = "TAG_TYPE_PROJECT"
  attestation_policy {
    # required
    title = "title"
    type = "POLICY_TYPE_ATTESTATION"

    attestation_policy {
      # required
      type = "ATTEST_POLICY_TYPE_ACCEPT"
    }
  }
  tag_level_settings {
    flow_log_transmission = "NETWORK_INSTANCE_FLOW_LOG_TRANSMISSION_UNSPECIFIED"
    interface_ordering = "INTERFACE_ORDERING_ENABLED"
  }
}

resource "zedcloud_brand" "test_tf_provider" {
  name = "test_tf_provider-create_edgenode"
  title = "test_tf_provider-create_edgenode"
  description = "description"
  origin_type = "ORIGIN_LOCAL"
}

resource "zedcloud_model" "test_tf_provider" {
  brand_id = zedcloud_brand.test_tf_provider.id
  name = "test_tf_provider-create_edgenode"
  title = "test_tf_provider-create_edgenode"
  type = "AMD64"
  origin_type = "ORIGIN_LOCAL"
  state = "SYS_MODEL_STATE_ACTIVE"
  attr = {
    memory = "8G"
    storage = "100G"
    Cpus = "4"
  }
  io_member_list {
    ztype = "IO_TYPE_ETH"
    phylabel =  "firstEth"
    usage = "ADAPTER_USAGE_MANAGEMENT"
    assigngrp = "eth0"
    phyaddrs = {
      Ifname = "eth0"
      PciLong = "0000:02:00.0"
    }
    logicallabel = "ethernet0"
    usage_policy = {
      FreeUplink = true
    }
    cost = 0
  }
  depends_on = [
    zedcloud_brand.test_tf_provider
  ]
}

resource "zedcloud_edgenode" "test_tf_provider" {
  name = "test_tf_provider-create_edgenode"
  model_id = zedcloud_model.test_tf_provider.id
  project_id = zedcloud_project.test_tf_provider.id
  title = "title"
  depends_on = [
    zedcloud_model.test_tf_provider,
    zedcloud_project.test_tf_provider
  ]
}

resource "zedcloud_volume_instance"  "test_volume_instance" {
  device_id = zedcloud_edgenode.test_tf_provider.id
  name = "test-name"
  title = "test-title"
  description = "test-description"
  type = "VOLUME_INSTANCE_TYPE_EMPTYDIR"
  multiattach = true
  cleartext = true
  accessmode = "VOLUME_INSTANCE_ACCESS_MODE_READWRITE"
  size_bytes = "1024"
  implicit = false
  label = "label"
  tags = {
    "tag-key-1" = "tag-value-1"
    "tag-key-2" = "tag-value-2"
  }
}
