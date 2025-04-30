resource "zedcloud_project" "test_tf_runtime_project" {
  # required
  name = "test_tf_runtime_project"
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

resource "zedcloud_brand" "test_tf_runtime_brand" {
  name = "test_tf_runtime_brand"
  title = "test_tf_runtime_brand"
  description = "description"
  origin_type = "ORIGIN_LOCAL"
}

resource "zedcloud_model" "test_tf_runtime_model" {
  brand_id = zedcloud_brand.test_tf_runtime_brand.id
  name = "test_tf_runtime_model"
  title = "test_tf_runtime_model"
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
}

resource "zedcloud_edgenode" "test_tf_runtime_node" {
  onboarding_key = "" # placeholder
  serialno = "2293dbe8-29ce-420c-8264-962857efc46c"
  # required
  name = "test_tf_runtime_node"
  model_id = zedcloud_model.test_tf_runtime_model.id
  project_id = zedcloud_project.test_tf_runtime_project.id
  title = "test_tf_provider-create_edgenode-title"
  interfaces {
    cost = 255
    intf_usage = "ADAPTER_USAGE_MANAGEMENT"
    intfname = "defaultIPv4"
    ipaddr = "127.0.0.1"
    macaddr = "00:00:00:00:00:00"
    tags = {
      "system_interface_1_key" = "system_interface_1_value"
      "system_interface_2_key" = "system_interface_2_value"
    }
  }
  interfaces {
    intfname = "eth0"
    intf_usage = "ADAPTER_USAGE_UNSPECIFIED"
  }
  admin_state = "ADMIN_STATE_ACTIVE"
  asset_id = "asset_id"
  config_item {
    bool_value = true
    float_value = 1.0
    key = "key"
    string_value = "string"
    uint32_value = 32
    uint64_value = 64
    value_type = "value type"
  }
  # cpu = 2
  deployment_tag = "depl_tag"
  # deprecated = false
  description = "description"
  dev_location {
    city = "berlin"
    country = "germany"
    freeloc = "freeloc"
    hostname = "hostname"
    loc = "52.520008, 13.404954"
    org = "zededa"
    postal = "10115"
    region = "europe/west"
    underlay_ip = ""
  }
  generate_soft_serial = false
  tags = {
    "tag-key-1" = "tag-value-1"
  }
  token = "token_runtime"
}

resource "zedcloud_datastore" "test_datastore_runtime" {
  ds_fqdn             = "ftp:://fake-datastore.com"
  ds_path             = ""
  ds_type             = "DATASTORE_TYPE_SFTP"
  name                = "test_tf_provider-ftp-runtime"
  title               = "test_tf_provider-ftp-runtime"
  description         = "test_tf_provider-ftp-runtime"
  region              = "eu"
}

resource "zedcloud_datastore" "test_datastore_compose" {
  ds_fqdn             = "docker://gcr.io"
  ds_path             = ""
  ds_type             = "DATASTORE_TYPE_CONTAINERREGISTRY"
  name                = "test_tf_provider-docker-compose"
  title               = "test_tf_provider-docker-compose"
  description         = "test_tf_provider-docker-compose"
  secret {
    api_key = "test_api_key"
    api_passwd = "test_api_passwd"
  }
  lifecycle {
    ignore_changes = [
      secret,
    ]
  }
}

resource "zedcloud_image" "test_runtime_image" {
  name                = "test_tf_provider-runtime-image"
  datastore_id        = zedcloud_datastore.test_datastore_runtime.id
  image_arch          = "AMD64"
  image_format        = "QCOW2"
  image_type          = "IMAGE_TYPE_DOCKER_RUNTIME"
  image_size_bytes    = "0"
  image_rel_url       = "test_tf_provider-runtime-image"
  title               = "tes_tf_runtime"
}

resource "zedcloud_image" "test_compose_image" {
  name                = "test_tf_provider-compose-image"
  datastore_id        = zedcloud_datastore.test_datastore_runtime.id
  image_format        = "RAW"
  image_type          = "IMAGE_TYPE_DOCKER_COMPOSE_TAR"
  image_size_bytes    = "0"
  image_rel_url       = "test_tf_provider-compose-image"
  title               = "tes_tf_compose"
}

resource "zedcloud_application" "test_tf_app_runtime" {
  name                 = "test_tf_app_runtime"
  title                = "runtime_app_bundle"
  description          = "runtime app bundle"
  user_defined_version = "1.1"
  origin_type          = "ORIGIN_LOCAL"
  manifest {
    ac_kind = "VMManifest"
    ac_version = "1.2.0"
    app_type = "APP_TYPE_VM"
    deployment_type = "DEPLOYMENT_TYPE_DOCKER_RUNTIME"
    name = "docker_runtime_app_bundle"
    owner {
      user = "Test User"
      email = "test@zededa.com"
      website = "www.zededa.com"
    }
    images {
      imagename = zedcloud_image.test_runtime_image.name
      imageid = zedcloud_image.test_runtime_image.id
      cleartext = false
      ignorepurge = false
      imageformat = "QCOW2"
    }
    interfaces {
      name = "eth0"
      directattach = false
      acls  {
        matches {
          type = "ip"
          value = "0.0.0.0/0"
        }
      }
    }
    interfaces {
      name = "eth1"
      directattach = false
      acls {
        matches {
          type = "ip"
          value = "0.0.0.0/0"
        }
      }
    }
    desc {
      category = "APP_CATEGORY_OTHERS"
      app_category = "APP_CATEGORY_OTHERS"
      agreement_list = {
        "AGREEMENT_LIST_KEY" = "AGREEMENT_LIST_VALUE"
      }
      license_list = {
        "APACHE_LICENSE_V2" = "https://choosealicense.com/licenses/apache-2.0/"
      }
      screenshot_list = {
        "screenshot_list_key" = "screenshot_list_value"
      }
      logo = {
        "logo" = "301dd8fd-1228-11eb-89dd-028e313eda78_logo"
      }
      # logo": null,
    }
    vmmode = "HV_PV"
    resources {
      name = "resourceType"
      value = "Tiny"
    }
    resources {
      name = "cpus"
      value = "1"
    }
    resources {
      name = "memory"
      value = "524288.00"
    }
    configuration {
      custom_config {
        add = false
      }
    }
    runtime_version = "RuntimeVersion_V1"
    runtime_protocol_version = "RuntimeProtocolVersion_V1"
    persistent_runtime_size_bytes = 5368709120
  }
}

resource "zedcloud_application" "test_tf_app_compose" {
  name                 = "test_tf_app_compose"
  title                = "compose_app_bundle"
  description          = "compose app bundle"
  user_defined_version = "1.1"
  origin_type          = "ORIGIN_LOCAL"
  datastore_id_list = [
    zedcloud_datastore.test_datastore_compose.id
  ]
  manifest {
    ac_kind         = "ComposeManifest"
    ac_version      = "1.2.0"
    app_type        = "APP_TYPE_DOCKER_COMPOSE"
    deployment_type = "DEPLOYMENT_TYPE_STAND_ALONE"
    name            = "docker_compose_app_bundle"
    owner {
      user    = "Test User"
      email   = "test@zededa.com"
      website = "www.zededa.com"
    }
    desc {
      app_category = "APP_CATEGORY_INDUSTRIAL"
      category = "APP_CATEGORY_INDUSTRIAL"
      logo = {
        "logo": "301dd8fd-1228-11eb-89dd-028e313eda78_logo"
      }
      license_list = {
        "APACHE_LICENSE_V2": "https://choosealicense.com/licenses/apache-2.0/"
      }
      agreement_list = {
        "AGREEMENT_LIST_KEY": "AGREEMENT_LIST_VALUE"
      }
      screenshot_list = {
        "screenshot_list_key": "screenshot_list_value"
      }
    }
    docker_compose_yaml_text = "IyBCYXNlIGltYWdlCkZST00gYWxwaW5lOmxhdGVzdAoKIyBEZWZpbmUgYnVpbGQgYXJndW1lbnRzIGZvciBhdXRoZW50aWNhdGlvbgpBUkcgUkVHSVNUUllfVVNFUgpBUkcgUkVHSVNUUllfUEFTUwoKIyBMb2cgaW4gdG8gdGhlIHByaXZhdGUgcmVnaXN0cnkKUlVOIGVjaG8gIiRSRUdJU1RSWV9QQVNTIiB8IGRvY2tlciBsb2dpbiBteS1wcml2YXRlLXJlZ2lzdHJ5LmNvbSAtdSAiJFJFR0lTVFJZX1VTRVIiIC0tcGFzc3dvcmQtc3RkaW4KCiMgUHVsbCB0aGUgcHJpdmF0ZSBpbWFnZQpSVU4gZG9ja2VyIHB1bGwgbXktcHJpdmF0ZS1yZWdpc3RyeS5jb20vbXktcHJpdmF0ZS1pbWFnZTpsYXRlc3QKCiMgU2V0IHdvcmtpbmcgZGlyZWN0b3J5CldPUktESVIgL2FwcAoKIyBDb3B5IG5lY2Vzc2FyeSBmaWxlcwpDT1BZIC4gLgoKIyBSdW4gdGhlIGFwcGxpY2F0aW9uCkNNRCBbIi4vbWFpbiJd"
  }
}

resource "zedcloud_network_instance" "test_tf_netinst_runtime_1" {
  device_id = zedcloud_edgenode.test_tf_runtime_node.id
  name = "test_tf_netinst_runtime_1"
  title = "title"
  kind = "NETWORK_INSTANCE_KIND_SWITCH"
  port = "eth0"

  # optional
  description = "switch_network_inst_eth0"
  type = "NETWORK_INSTANCE_DHCP_TYPE_V4"
  device_default = false
  dhcp = false
  dns_list {
    addrs = [
      "10.1.2.2",
      "10.1.2.1"
    ]
    hostname = "wwww.ns2.example.com"
  }
  dns_list {
    addrs = [
      "10.1.1.1",
      "10.1.1.2"
    ]
    hostname = "wwww.ns1.example.com"
  }
  opaque {
    oconfig = "Test OConfig"
    type = "OPAQUE_CONFIG_TYPE_UNSPECIFIED"
  }
  mtu = 1500
}

resource "zedcloud_network_instance" "test_tf_netinst_runtime_2" {
  device_id = zedcloud_edgenode.test_tf_runtime_node.id
  name = "test_tf_netinst_runtime_2"
  title = "title"
  kind = "NETWORK_INSTANCE_KIND_LOCAL"

  # optional
  description = "local_airgapped_network_inst"
  type = "NETWORK_INSTANCE_DHCP_TYPE_V4"
  device_default = false
  dhcp = false
  dns_list {
    addrs = [
      "10.1.2.2",
      "10.1.2.1"
    ]
    hostname = "wwww.ns2.example.com"
  }
  dns_list {
    addrs = [
      "10.1.1.1",
      "10.1.1.2"
    ]
    hostname = "wwww.ns1.example.com"
  }
  opaque {
    oconfig = "Test OConfig"
    type = "OPAQUE_CONFIG_TYPE_UNSPECIFIED"
  }
  mtu = 1500
}

resource "zedcloud_application_instance"  "test_tf_appinst_runtime" {
  name        = "test_tf_appinst_runtime"
  title       = "tf test"
  description = "tf test"
  app_id      = zedcloud_application.test_tf_app_runtime.id
  app_type    = "APP_TYPE_VM"
  device_id   = zedcloud_edgenode.test_tf_runtime_node.id
  collect_stats_ip_addr = "10.1.255.254"
  interfaces {
    netinstname  = zedcloud_network_instance.test_tf_netinst_runtime_1.name
    intfname     = "eth0"
    intforder = 1
    privateip = false
    netinsttag = {
      "fake_tag" = "fake_value"
    }
  }
  interfaces {
    netinstname  = zedcloud_network_instance.test_tf_netinst_runtime_2.name
    intfname     = "eth1"
    intforder = 2
    privateip = false
    netinsttag = {
      "fake_tag" = "fake_value"
    }
  }
  tags = {
    "tag-key-1" = "tag-value-1"
  }
  logs {
    access = false
  }
  manifest_info {
    transition_action = "INSTANCE_TA_NONE"
  }
  custom_config {
    add = false
    allow_storage_resize = false
    override = false
  }
  vminfo {
    cpus = 1
    enable_oem_win_license_key = false
    mode = "HV_PV"
    vnc  = false
  }
}

resource "zedcloud_application_instance"  "test_tf_appinst_compose" {
  depends_on = [
    zedcloud_application_instance.test_tf_appinst_runtime,
  ]
  name        = "test_tf_appinst_compose"
  title       = "tf test"
  description = "tf test"
  app_id      = zedcloud_application.test_tf_app_compose.id
  app_type    = zedcloud_application.test_tf_app_compose.manifest[0].app_type
  device_id   = zedcloud_edgenode.test_tf_runtime_node.id
  tags = {
    "tag-key-1" = "tag-value-1"
  }
  logs {
    access = false
  }
  manifest_info {
    transition_action = "INSTANCE_TA_NONE"
  }
  custom_config {
    add = false
    allow_storage_resize = false
    override = false
  }
}