
resource "zedcloud_project" "test_tf_provider" {
  # required
  name  = "test_tf_provider_newproject"
  title = "title"

  # optional
  type = "TAG_TYPE_PROJECT"
  tag_level_settings {
    flow_log_transmission = "NETWORK_INSTANCE_FLOW_LOG_TRANSMISSION_DISABLED"
    interface_ordering    = "INTERFACE_ORDERING_DISABLED"
  }
}

data "zedcloud_project" "test_tf_provider" {
  name  = "test_tf_provider_newproject"
  title = "title"
  type = "TAG_TYPE_PROJECT"
  depends_on = [
    zedcloud_project.test_tf_provider
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


resource "zedcloud_edgenode" "test_tf_provider" {
  onboarding_key = "" # placeholder
  serialno       = "2293dbe8-29ce-420c-8264-962858efc46b"
  # required
  name       = "test_tf_provider_newedgenode"
  model_id   = zedcloud_model.test_tf_provider.id
  project_id = zedcloud_project.test_tf_provider.id
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
    "tag-key-1" = "tag-value-1"
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

data "zedcloud_edgenode" "test_tf_provider" {
  name     = "test_tf_provider_newedgenode"
  title    = "test_tf_provider-create_edgenode-title"
  model_id = zedcloud_model.test_tf_provider.id
  depends_on = [
    zedcloud_edgenode.test_tf_provider
  ]
  interfaces {}
}

resource "zedcloud_asset_group" "test_tf_provider" {
  name        = "test_tf_provider_assetgroup"
  description = "This is an example asset group"
  project_id  = data.zedcloud_project.test_tf_provider.id
  asset_ids {
    ids = [ data.zedcloud_edgenode.test_tf_provider.id ]
  }
}

resource "zedcloud_datastore" "open_ds_provider" {
  # required
  ds_fqdn             = "http://147.75.33.217"
  ds_path             = "images"
  ds_type             = "DATASTORE_TYPE_HTTP"
  name                = "open_ds_provider"
  title               = "open_ds_provider"
  description         = "open_ds_provider"
  region              = "eu"
  # project_access_list = [zedcloud_project.test_tf_provider.id]
}

resource "zedcloud_image" "open_alpine_image" {
  depends_on = [
    zedcloud_datastore.open_ds_provider
  ]
  name                = "openalpine"
  datastore_id        = zedcloud_datastore.open_ds_provider.id
  image_arch          = "ARM64"
  image_format        = "CONTAINER"
  image_rel_url       = "alpine:latest"
  image_size_bytes    = 0
  image_type          = "IMAGE_TYPE_APPLICATION"
  title               = "openalpine"
}


resource "zedcloud_app_profile" "test_tf_provider_app_profile" {
  depends_on = [ zedcloud_image.open_alpine_image ]
  name = "test_tf_provider_app_profile"
  title = "test_tf_provider_app_profile"

  app_policies {
    meta_data {
      name = "test_tf_provider_default-app-policy"
      title = "test_tf_provider_default-app-policy"
    }

    app_config {
      bundle_version = 0
      cpus = 2

      manifest_json {
        ac_kind             = "VMManifest"
        ac_version          = "1.2.0"
        name                = "test_tf_provider_default-app-policy"
        vmmode              = "HV_HVM"
        enablevnc           = false
        app_type            = "APP_TYPE_VM"
        deployment_type     = "DEPLOYMENT_TYPE_STAND_ALONE"
        cpu_pinning_enabled = false

        images {
          imagename   = zedcloud_image.open_alpine_image.name
          imageid     = zedcloud_image.open_alpine_image.id
          imageformat = zedcloud_image.open_alpine_image.image_format
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
      name = "test_tf_provider_default-network-policy"
      title = "test_tf_provider_default-network-policy"
    }

    network_config {
      port = "1234"
      kind = "NETWORK_INSTANCE_KIND_LOCAL"
      type = "NETWORK_INSTANCE_DHCP_TYPE_V4"
    }
  }
}


data "zedcloud_asset_group" "test_tf_provider" {
  name     = "test_tf_provider_assetgroup"
  title    = "test_tf_provider_assetgroup"
  depends_on = [
    zedcloud_asset_group.test_tf_provider
  ]
}

data "zedcloud_app_profile" "test_tf_provider_app_profile" {
  name     = "test_tf_provider_app_profile"
  title    = "test_tf_provider_app_profile"
  depends_on = [
    zedcloud_app_profile.test_tf_provider_app_profile
  ]
}

resource "zedcloud_profile_deployment" "test_tf_provider" {
  name = "test_tf_provider"
  title = "test_tf_provider"
  app_profile_info {
    app_profile_id = zedcloud_app_profile.test_tf_provider_app_profile.id
    version = 1
  }
  target_asset_group {
    group_id  = zedcloud_asset_group.test_tf_provider.id
  }
  project_id = zedcloud_project.test_tf_provider.id
  status = ""
  depends_on = [
    zedcloud_app_profile.test_tf_provider_app_profile,
    zedcloud_asset_group.test_tf_provider
  ]
}
