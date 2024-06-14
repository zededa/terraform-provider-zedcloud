// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0
//

terraform {
  required_providers {
    zedcloud = {
      source = "zededa/zedcloud"
      //version = "0.0.7-alpha"
    }
  }
}

resource "zedcloud_project" "test_tf_provider" {
  # required
  name  = "test_tf_provider"
  title = "title"

  # optional
  type = "TAG_TYPE_PROJECT"
  attestation_policy {
    # required
    title = "title"
    type  = "POLICY_TYPE_ATTESTATION"

    attestation_policy {
      type = "ATTEST_POLICY_TYPE_ACCEPT"
    }
  }
  edgeview_policy {
    edgeview_policy {
      # required
      edgeview_allow = false
      # optional
      access_allow_change = true
      edgeviewcfg {
        dev_policy {
          allow_dev = true
        }
        ext_policy {
          allow_ext = false
        }
        generation_id = 0
        jwt_info {
          allow_sec = 18000
          disp_url  = "zedcloud.local.zededa.net/api/v1/edge-view"
          encrypt   = false
          num_inst = 1
        }
      }
      max_expire_sec = 2592000
      max_inst       = 3
    }
    type = "POLICY_TYPE_EDGEVIEW"
  }
}

data "zedcloud_project" "test_tf_provider" {
  name  = "test_tf_provider"
  type  = "TAG_TYPE_PROJECT"
  title = "test_tf_provider"
  depends_on = [
    zedcloud_project.test_tf_provider
  ]
}

resource "zedcloud_edgenode" "test_tf_provider" {
  onboarding_key = "" # placeholder
  serialno       = "2293dbe8-29ce-420c-8264-962857efc46b"
  # required
  name       = "test_tf_provider"
  model_id   = "ab66ab26-cbc2-4d66-829e-7bedb59aabba"
  project_id = data.zedcloud_project.test_tf_provider.id
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
  description = "description"
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
}

data "zedcloud_edgenode" "test_tf_provider" {
  name     = "test_tf_provider"
  title    = "test_tf_provider"
  model_id = "ab66ab26-cbc2-4d66-829e-7bedb59aabba"
  depends_on = [
    zedcloud_edgenode.test_tf_provider
  ]
}

resource "zedcloud_application" "test_tf_provider" {
  name                 = "test_tf_provider"
  title                = "ubuntu-all-ip"
  description          = "ubuntu-all-ip"
  user_defined_version = "1.1"
  origin_type          = "ORIGIN_LOCAL"
  manifest {
    # computed
    ac_kind = "VMManifest"

    # optional
    name       = "xenial-amd64-docker-20180725"
    ac_version = "1.2.0"
    app_type   = "APP_TYPE_VM"
    configuration {
      custom_config {
        add                  = false
        allow_storage_resize = false
        field_delimiter      = ","
        name                 = "custom_config_name"
        override             = false
        template             = "dGVzdA=="
        variable_groups {
          condition {
            # optional
            name     = "condition"
            operator = "CONDITION_OPERATOR_EQUALTO"
            value    = "val"
          }
          name     = "var_group"
          required = false
        }
      }
    }
    cpu_pinning_enabled = false
    deployment_type     = "DEPLOYMENT_TYPE_K3S"
    desc {
      category     = "Infrastructure"
      os           = "Zenix"
      app_category = "APP_CATEGORY_OTHERS"
      support      = "support"
    }
    description  = "description"
    display_name = "display_name"
    enablevnc    = false
    interfaces {
      name         = "indirect"
      directattach = false
      privateip    = false
      acls {
        matches {
          type  = "protocol"
          value = "tcp"
        }
        matches {
          type  = "lport"
          value = "8022"
        }
        actions {
          portmap = true
          portmapto {
            app_port = 22
          }
        }
      }
      acls {
        matches {
          type  = "host"
          value = ""
        }
      }
    }
    owner {
      company = "Zededa Inc."
      email   = "test@zededa.com"
      group   = "testgroup"
      user    = "testuser"
      website = "http://www.zededa.com"
    }
    resources {
      name  = "cpus"
      value = "2"
    }
    resources {
      name  = "memory"
      value = "1024000"
    }
    vmmode = "HV_HVM"
  }
}

data "zedcloud_application" "test_tf_provider" {
  name  = "test_tf_provider"
  title = "test_tf_provider"
  depends_on = [
    zedcloud_application.test_tf_provider
  ]
}

resource "zedcloud_volume_instance" "test_tf_provider" {
  device_id   = data.zedcloud_edgenode.test_tf_provider.id
  name        = "test_tf_provider"
  title       = "test_title"
  description = "test_description"
  type        = "VOLUME_INSTANCE_TYPE_BLOCKSTORAGE"

  multiattach = true
  cleartext   = true
  accessmode  = "VOLUME_INSTANCE_ACCESS_MODE_READWRITE"
  size_bytes  = "1024"
  implicit    = false
  label       = "label"
  tags = {
    "tag-key-1" = "tag-value-1"
    "tag-key-2" = "tag-value-2"
  }
  depends_on = [
    data.zedcloud_edgenode.test_tf_provider
  ]
}

data "zedcloud_volume_instance" "test_tf_provider" {
  name      = "test_tf_provider"
  title     = "test_tf_provider"
  type      = "VOLUME_INSTANCE_TYPE_CONTENT_TREE"
  device_id = data.zedcloud_edgenode.test_tf_provider.id
  depends_on = [
    zedcloud_volume_instance.test_tf_provider
  ]
}

resource "zedcloud_network_instance" "test_tf_provider" {
  depends_on = [
    zedcloud_edgenode.test_tf_provider,
  ]
  # required
  device_id = data.zedcloud_edgenode.test_tf_provider.id
  name      = "test_tf_provider"
  title     = "title"
  kind      = "NETWORK_INSTANCE_KIND_LOCAL"
  port      = "eth1"

  # optional
  description    = "zedcloud_network_instance-complete-description"
  type           = "NETWORK_INSTANCE_DHCP_TYPE_V4"
  device_default = false
  dhcp           = false
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
    type    = "OPAQUE_CONFIG_TYPE_UNSPECIFIED"
  }
  tags = {
    "ni-tag1" = "ni-tag-value-1"
    "ni-tag2" = "ni-tag-value-2"
  }
}

data "zedcloud_network_instance" "test_tf_provider" {
  name      = "test_tf_provider"
  title     = "test_tf_provider"
  kind      = "NETWORK_INSTANCE_KIND_LOCAL"
  device_id = data.zedcloud_edgenode.test_tf_provider.id
  depends_on = [
    zedcloud_network_instance.test_tf_provider
  ]
}

resource "zedcloud_application_instance" "test_tf_provider" {
  depends_on = [
    zedcloud_application.test_tf_provider,
    zedcloud_edgenode.test_tf_provider,
    zedcloud_network_instance.test_tf_provider
  ]

  # required
  name        = "test_tf_provider"
  title       = "tf test"
  description = "tf test"
  activate    = "true"
  app_id      = zedcloud_application.test_tf_provider.id
  device_id   = zedcloud_edgenode.test_tf_provider.id
  drives {
    # required
    drvtype   = "UNSPECIFIED"
    imagename = "ubuntu-tiny"
    maxsize   = "0"
    preserve  = false
    readonly  = false
    target    = "Disk"

    # optional
    cleartext   = false
    ignorepurge = false
  }
  interfaces {
    directattach = false
    privateip    = false
    netinstname  = zedcloud_network_instance.test_tf_provider.name
    intfname     = "indirect"
    acls {
      actions {
        drop       = false
        limit      = false
        limitburst = 0
        limitrate  = 0
        portmap    = true

        mapparams {
          port = 22
        }
      }

      matches {
        type  = "protocol"
        value = "tcp"
      }
      matches {
        type  = "lport"
        value = "8022"
      }
    }
    acls {
      #   id = 2
      matches {
        type = "host"
      }
    }
  }
  # optional
  remote_console        = false
  is_secret_updated     = false
  collect_stats_ip_addr = "true"
  app_policy_id         = ""
  app_type              = "APP_TYPE_VM"
  cluster_id            = ""
  custom_config {
    add                  = false
    allow_storage_resize = false
    override             = false
  }
  logs {
    access = false
  }
  manifest_info {
    transition_action = "INSTANCE_TA_NONE"
  }
  start_delay_in_seconds = 120
  user_data              = ""
  vminfo {
    # required
    cpus = 2
    mode = "HV_HVM"
    vnc  = false
    # optional
  }
  tags = {
    "tag-key-1" = "tag-value-1"
  }
}

data "zedcloud_application_instance" "test_tf_provider" {
  depends_on = [
    zedcloud_application_instance.test_tf_provider
  ]
  name      = "test_tf_provider"
  title     = "test_tf_provider"
  device_id = zedcloud_edgenode.test_tf_provider.id
  app_id    = zedcloud_application.test_tf_provider.id
}


resource "zedcloud_image"  "test_tf_provider" {
  name = "test_tf_provider"
  title = "test_tf_provider"
  image_arch = "UNSPECIFIED"
  image_format = "RAW"
  datastore_id = "55c029bd-5d0a-46dd-b0ab-c6c9895b0d9a"
  image_rel_url = "test_url"
  image_size_bytes = 5000
  image_type = "IMAGE_TYPE_ARTIFACT"
  image_sha256 = "815fa14f544f2e955af18848667320fefd90d82e42497f429b53cc57cc76fb5c"
  project_access_list = []
}


resource "zedcloud_patch_envelope" "test_tf_provider" {
  name   = "test_tf_provider"
  action = "PATCH_ENVELOPE_ACTION_ACTIVATE"
  title  = "test_tf_provider"
  description = "test tf provider"
  user_defined_version = "1.0"

  artifacts {
    format = "OpaqueObjectCategoryInline"
    base64_artifact {
      base64_data      = "YXJ0aWZhY3QgZGF0YQ=="
      file_name_to_use = "testfile"
    }
  }
  artifacts {
    format = "OpaqueObjectCategoryExternalBinary"
    binary_artifact {
      image_id = zedcloud_image.test_tf_provider.id
      image_name = zedcloud_image.test_tf_provider.name
      file_name_to_use = "new_filename"
    }
  }
  project_name = zedcloud_project.test_tf_provider.name
  project_id = zedcloud_project.test_tf_provider.id
  depends_on = [
    zedcloud_image.test_tf_provider
  ]
}

resource "zedcloud_patch_reference_update" "test_tf_provider" {
  app_inst_id_list = [zedcloud_application_instance.test_tf_provider.id]
  patchenvelope_id = zedcloud_patch_envelope.test_tf_provider.id
  project_id = zedcloud_project.test_tf_provider.id
  depends_on = [
    zedcloud_application_instance.test_tf_provider,
    zedcloud_patch_envelope.test_tf_provider,
    zedcloud_project.test_tf_provider
  ]
}

