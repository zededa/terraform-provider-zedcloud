
resource "zedcloud_project" "test_tf_provider_for_tags" {
  # required
  name  = "test_tf_provider_newproject_for_tags"
  title = "title"

  # optional
  type = "TAG_TYPE_PROJECT"
  tag_level_settings {
    flow_log_transmission = "NETWORK_INSTANCE_FLOW_LOG_TRANSMISSION_DISABLED"
    interface_ordering    = "INTERFACE_ORDERING_DISABLED"
  }
}

data "zedcloud_project" "test_tf_provider_for_tags" {
  name  = "test_tf_provider_newproject_for_tags"
  title = "title"
  type = "TAG_TYPE_PROJECT"
  depends_on = [
    zedcloud_project.test_tf_provider_for_tags
  ]
}

resource "zedcloud_brand" "test_tf_provider" {
  name        = "qemu100"
  title       = "QEMU100"
  description = "qemu100"
  origin_type = "ORIGIN_LOCAL"
}

resource "zedcloud_model" "test_tf_provider" {
  brand_id    = zedcloud_brand.test_tf_provider.id
  name        = "test_tf_provider-create_edgenode100"
  title       = "test_tf_provider-create_edgenode100"
  type        = "AMD64"
  origin_type = "ORIGIN_LOCAL"
  state       = "SYS_MODEL_STATE_ACTIVE"
  attr = {
    memory  = "8G"
    storage = "100G"
    Cpus    = "4"
  }
  io_member_list {
    ztype     = "IO_TYPE_ETH"
    phylabel  = "firstEth"
    usage     = "ADAPTER_USAGE_MANAGEMENT"
    assigngrp = "eth0"
    phyaddrs = {
      Ifname  = "eth0"
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


resource "zedcloud_edgenode" "test_tf_provider_for_tags" {
  onboarding_key = "" # placeholder
  serialno       = "2293dbe8-29ce-420c-8364-962858efc46b"
  # required
  name       = "test_tf_provider_newedgenode_for_tags"
  model_id   = zedcloud_model.test_tf_provider.id
  project_id = data.zedcloud_project.test_tf_provider_for_tags.id
  title      = "test_tf_provider-create_edgenode-title"

  admin_state = "ADMIN_STATE_ACTIVE"
  asset_id    = "asset_id"
  config_item {
    bool_value   = true
    float_value  = 1.0
    key          = "key"
    string_value = "string"
    uint32_value = 32
    uint64_value = 64
    value_type   = "value type"
  }
  deployment_tag = "depl_tag"
  description    = "description"
  dev_location {
    city        = "berlin"
    country     = "germany"
    freeloc     = "freeloc"
    hostname    = "hostname"
    loc         = "52.520008, 13.404954"
    org         = "zededa"
    postal      = "10115"
    region      = "europe/west"
    underlay_ip = ""
  }
  generate_soft_serial = false
  tags = {
    "tag-for-tags" = "value-for-tags"
  }
  token = "token"
  interfaces {
    cost       = 255
    intf_usage = "ADAPTER_USAGE_MANAGEMENT"
    intfname   = "defaultIPv4"
    ipaddr     = "127.0.0.1"
    # macaddr = "00:00:00:00:00:00"
    tags = {
      "system_interface_1_key" = "system_interface_1_value"
      "system_interface_2_key" = "system_interface_2_value"
    }
  }
}

data "zedcloud_edgenode" "test_tf_provider_for_tags" {
  name     = "test_tf_provider_newedgenode_for_tags"
  title    = "test_tf_provider-create_edgenode-title"
  model_id = zedcloud_model.test_tf_provider.id
  tags = {
    "tag-key-1" = "tag-value-1"
  }
  depends_on = [
    zedcloud_edgenode.test_tf_provider_for_tags
  ]
}

resource "zedcloud_asset_group" "test_tf_provider_for_tags" {
  name        = "alphagroup13_for_tags"
  description = "This is an example asset group"
  project_id  = data.zedcloud_project.test_tf_provider_for_tags.id
  
  asset_tags {
    asset_tag {
      tag = {
        key = "tag-for-tags"
        value = "value-for-tags"
      }
    }
  }
}
