
resource "zedcloud_project" "test_tf_provider_deploy_via_tags" {
  # required
  name  = "test_tf_provider_newproject_deploy_via_tags"
  title = "title"

  # optional
  type = "TAG_TYPE_PROJECT"
  tag_level_settings {
    flow_log_transmission = "NETWORK_INSTANCE_FLOW_LOG_TRANSMISSION_DISABLED"
    interface_ordering    = "INTERFACE_ORDERING_DISABLED"
  }
}

data "zedcloud_project" "test_tf_provider_deploy_via_tags" {
  name  = "test_tf_provider_newproject_deploy_via_tags"
  title = "title"
  type = "TAG_TYPE_PROJECT"
  depends_on = [
    zedcloud_project.test_tf_provider_deploy_via_tags
  ]
}

resource "zedcloud_brand" "test_tf_provider_deploy_via_tags" {
  name        = "test_tf_provider-qemu100_deploy_via_tags"
  title       = "QEMU100"
  description = "qemu100"
  origin_type = "ORIGIN_LOCAL"
}

resource "zedcloud_model" "test_tf_provider_deploy_via_tags" {
  brand_id    = zedcloud_brand.test_tf_provider_deploy_via_tags.id
  name        = "test_tf_provider-create_edgenode100_deploy_via_tags"
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
    zedcloud_brand.test_tf_provider_deploy_via_tags
  ]
}


resource "zedcloud_edgenode" "test_tf_provider_deploy_via_tags" {
  onboarding_key = "" # placeholder
  serialno       = "2293dbe8-29ce-420c-8264-962858edc46b"
  # required
  name       = "test_tf_provider_newedgenode_deploy_via_tags"
  model_id   = zedcloud_model.test_tf_provider_deploy_via_tags.id
  project_id = zedcloud_project.test_tf_provider_deploy_via_tags.id
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
    "tag-key-100" = "tag-value-100"
  }
  token = "token_profiledeploymenttags"
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

resource "zedcloud_asset_group" "test_tf_provider_deploy_via_tags" {
  name        = "test_tf_provider_assetgroup_deploy_via_tags"
  description = "This is an example asset group"
  project_id  = data.zedcloud_project.test_tf_provider_deploy_via_tags.id
  asset_tags {
    asset_tag {
      tag = {
        key = "tag-key-100"
        value = "tag-value-100"
      }
    }
  }
}

resource "zedcloud_datastore" "open_ds_provider_deploy_via_tags" {
  # required
  ds_fqdn             = "http://147.75.33.217"
  ds_path             = "images"
  ds_type             = "DATASTORE_TYPE_HTTP"
  name                = "test_tf_provider-open_ds_provider_deploy_via_tags"
  title               = "open_ds_provider"
  description         = "open_ds_provider"
  region              = "eu"
  # project_access_list = [zedcloud_project.test_tf_provider.id]
}

resource "zedcloud_image" "open_alpine_image_deploy_via_tags" {
  depends_on = [
    zedcloud_datastore.open_ds_provider_deploy_via_tags
  ]
  name                = "test_tf_provider-open_alpine_image_deploy_via_tags"
  datastore_id        = zedcloud_datastore.open_ds_provider_deploy_via_tags.id
  image_arch          = "ARM64"
  image_format        = "CONTAINER"
  image_rel_url       = "alpine:latest"
  image_size_bytes    = 0
  image_type          = "IMAGE_TYPE_APPLICATION"
  title               = "openalpine"
}


resource "zedcloud_app_profile" "test_tf_provider_app_profile_deploy_via_tags" {
  depends_on = [ zedcloud_image.open_alpine_image_deploy_via_tags ]
  name = "test_tf_provider_app_profile_deploy_via_tags"
  title = "test_tf_provider_app_profile"

  app_policies {
    meta_data {
      name = "test_tf_provider_default-app-policy_deploy_via_tags"
      title = "test_tf_provider_default-app-policy"
    }

    app_config {
      bundle_version = 0
      cpus = 2

      manifest_json {
        ac_kind             = "VMManifest"
        ac_version          = "1.2.0"
        name                = "test_tf_provider_default-app-policy_deploy_via_tags"
        vmmode              = "HV_HVM"
        enablevnc           = false
        app_type            = "APP_TYPE_VM"
        deployment_type     = "DEPLOYMENT_TYPE_STAND_ALONE"
        cpu_pinning_enabled = false

        images {
          imagename   = zedcloud_image.open_alpine_image_deploy_via_tags.name
          imageid     = zedcloud_image.open_alpine_image_deploy_via_tags.id
          imageformat = zedcloud_image.open_alpine_image_deploy_via_tags.image_format
          cleartext   = false
          drvtype     = "HDD"
          ignorepurge = true
          maxsize     = 102400000
          target      = "Disk"
        }
      }
    }
  }

  network_policies {
    meta_data {
      name = "test_tf_provider_default-network-policy_deploy_via_tags"
      title = "test_tf_provider_default-network-policy"
    }

    network_config {
      port = "1234"
      kind = "NETWORK_INSTANCE_KIND_LOCAL"
      type = "NETWORK_INSTANCE_DHCP_TYPE_V4"
    }
  }
}


data "zedcloud_asset_group" "test_tf_provider_deploy_via_tags" {
  name     = "test_tf_provider_assetgroup_deploy_via_tags"
  title    = "test_tf_provider_assetgroup"
  depends_on = [
    zedcloud_asset_group.test_tf_provider_deploy_via_tags
  ]
}

data "zedcloud_app_profile" "test_tf_provider_app_profile_deploy_via_tags" {
  name     = "test_tf_provider_app_profile_deploy_via_tags"
  title    = "test_tf_provider_app_profile"
  depends_on = [
    zedcloud_app_profile.test_tf_provider_app_profile_deploy_via_tags
  ]
}

resource "zedcloud_profile_deployment" "test_tf_provider_deploy_via_tags" {
  name = "test_tf_provider_deploy_via_tags"
  title = "test_tf_provider"
  app_profile_info {
    app_profile_id = zedcloud_app_profile.test_tf_provider_app_profile_deploy_via_tags.id
    version = 1
  }
  target_asset_group {
    group_id  = zedcloud_asset_group.test_tf_provider_deploy_via_tags.id
  }
  project_id = zedcloud_project.test_tf_provider_deploy_via_tags.id
  status = ""
  depends_on = [
    zedcloud_app_profile.test_tf_provider_app_profile_deploy_via_tags,
    zedcloud_asset_group.test_tf_provider_deploy_via_tags
  ]
}
