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

resource "zedcloud_datastore"  "test_tf_provider" {
  depends_on = [
    zedcloud_project.test_tf_provider
  ]
  # required
  ds_fqdn = "my-datastore.my-company.com"
  ds_path = "download/AMD64"
  ds_type = "DATASTORE_TYPE_AZUREBLOB"
  name = "test_tf_provider-test-datastore"
  title = "test_tf_provider-title"
  description = "description"
  region = "eu"
  project_access_list = [zedcloud_project.test_tf_provider.id]
}

resource "zedcloud_image" "test_tf_provider" {
  depends_on = [
    zedcloud_datastore.test_tf_provider,
    zedcloud_project.test_tf_provider
  ]
  name = "test_tf_provider-create_edgenode"
  datastore_id = zedcloud_datastore.test_tf_provider.id
  image_arch = "AMD64"
  image_format = "CONTAINER"
  image_rel_url = "test_url"
  image_size_bytes = 0
  image_type =  "IMAGE_TYPE_APPLICATION"
  title = "test"
  project_access_list = [zedcloud_project.test_tf_provider.id]
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
  onboarding_key = ""
  serialno = "2293dbe8-29ce-420c-8264-962857efc46b"

  # required
  name = "test_tf_provider"
  model_id = zedcloud_model.test_tf_provider.id
  project_id = zedcloud_project.test_tf_provider.id
  title = "test_tf_provider-create_edgenode-title"

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
  token = "token"
  depends_on = [
    zedcloud_model.test_tf_provider,
    zedcloud_project.test_tf_provider
  ]
}

data "zedcloud_edgenode" "test_tf_provider" {
  name = "test_tf_provider"
  model_id = zedcloud_edgenode.test_tf_provider.model_id
  title = zedcloud_edgenode.test_tf_provider.title
  depends_on = [
    zedcloud_edgenode.test_tf_provider
  ]
}

resource "zedcloud_application" "test_tf_provider" {
  name = "test_tf_provider"
  title = "test_tf_provider-ubuntu-all-ip"
  description = "test_tf_provider-ubuntu-all-ip"
  user_defined_version = "1.1"
  origin_type = "ORIGIN_LOCAL"
  manifest {
    name = "xenial-amd64-docker-20180725"
    ac_version = "1.2.0"
    ac_kind = "VMManifest"
    app_type = "APP_TYPE_VM"
    configuration {
      custom_config {
        add = false
        allow_storage_resize = false
        field_delimiter = ","
        name = "custom_config_name"
        override = false
        template = "dGVzdA=="
        variable_groups {
          condition {
            # optional
            name = "condition"
            operator = "CONDITION_OPERATOR_EQUALTO"
            value = "val"
          }
          name = "var_group"
          required = false
        }
      }
    }
    cpu_pinning_enabled = false
    deployment_type = "DEPLOYMENT_TYPE_K3S"
    desc {
      category = "APP_CATEGORY_OTHERS"
      os = "Zenix"
      app_category = "APP_CATEGORY_OTHERS"
      support = "support"
    }
    description = "description"
    display_name = "display_name"
    enablevnc = false
    owner {
      company = "Zededa Inc."
      email = "test@zededa.com"
      group = "testgroup"
      user = "testuser"
      website = "http://www.zededa.com"
    }
    resources {
      name = "cpus"
      value = "2"
    }
    resources {
      name = "memory"
      value = "1024000"
    }
    vmmode = "HV_HVM"
  }
}

data "zedcloud_application" "test_tf_provider" {
  name = "test_tf_provider"
  title = zedcloud_application.test_tf_provider.title
  depends_on = [
    zedcloud_application.test_tf_provider
  ]
}

resource "zedcloud_volume_instance"  "test_tf_provider" {
  device_id = data.zedcloud_edgenode.test_tf_provider.id
  name = "test_tf_provider"
  title = "test_title"
  description = "test_description"
  type = "VOLUME_INSTANCE_TYPE_BLOCKSTORAGE"

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
  depends_on = [
    data.zedcloud_edgenode.test_tf_provider
  ]
}

resource "zedcloud_network_instance" "test_tf_provider" {
  depends_on = [
    zedcloud_edgenode.test_tf_provider,
  ]
  # required
  device_id = data.zedcloud_edgenode.test_tf_provider.id
  name = "test_tf_provider"
  title = "title"
  kind = "NETWORK_INSTANCE_KIND_LOCAL"
  port = "eth1"

  # optional
  description = "zedcloud_network_instance-complete-description"
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
  tags = {
    "ni-tag1" = "ni-tag-value-1"
    "ni-tag2" = "ni-tag-value-2"
  }
}

resource "zedcloud_application_instance"  "test_tf_provider" {
  depends_on = [
    zedcloud_application.test_tf_provider,
    zedcloud_edgenode.test_tf_provider,
    zedcloud_network_instance.test_tf_provider,
    zedcloud_image.test_tf_provider
  ]

  # required
  name = "test_tf_provider"
  title = "tf test"
  description = "tf test"
  activate = "true"
  app_id = zedcloud_application.test_tf_provider.id
  device_id = zedcloud_edgenode.test_tf_provider.id

  drives {
    # required
    drvtype = "UNSPECIFIED"
    imagename = "test_tf_provider-create_edgenode"
    #             "ubuntu-tiny"
    maxsize = "0"
    preserve = false
    readonly = false
    target = "Disk"
  }
  # optional
  remote_console = true
  is_secret_updated = false
  collect_stats_ip_addr = "true"
  app_policy_id = ""
  app_type = "APP_TYPE_VM"
  cluster_id = ""
  custom_config {
    add = false
    allow_storage_resize = false
    override = false
  }
  logs {
    access = false
  }
  manifest_info {
    transition_action = "INSTANCE_TA_NONE"
  }
  start_delay_in_seconds = 120
  user_data = ""
  vminfo {
    # required
    cpus = 2
    mode = "HV_HVM"
    vnc = false
    enable_oem_win_license_key = false
  }
  tags = {
    "tag-key-1" = "tag-value-1"
  }
}

