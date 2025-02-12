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
  tag_level_settings {
    flow_log_transmission =  "NETWORK_INSTANCE_FLOW_LOG_TRANSMISSION_DISABLED"
    interface_ordering = "INTERFACE_ORDERING_DISABLED"
  }
  # attestation_policy {
  #   # required
  #   title = "tf-attestation-policy"
  #   name  = "tf-attestation-policy"
  #   type  = "POLICY_TYPE_ATTESTATION"

  #   attestation_policy {
  #     type = "ATTEST_POLICY_TYPE_ACCEPT"
  #   }
  # }
  # edgeview_policy {
  #   title = "tf-edgeview-policy"
  #   name  = "tf-edgeview-policy"
  #   edgeview_policy {
  #     # required
  #     edgeview_allow = false
  #     # optional
  #     access_allow_change = true
  #     edgeviewcfg {
  #       dev_policy {
  #         allow_dev = true
  #       }
  #       ext_policy {
  #         allow_ext = false
  #       }
  #       generation_id = 0
  #       jwt_info {
  #         allow_sec = 18000
  #         disp_url  = "zedcloud.local.zededa.net/api/v1/edge-view"
  #         encrypt   = false
  #         num_inst  = 1
  #       }
  #     }
  #     max_expire_sec = 2592000
  #     max_inst       = 3
  #   }
  #   type = "POLICY_TYPE_EDGEVIEW"
  # }
}

data "zedcloud_project" "test_tf_provider" {
  name  = "test_tf_provider"
  type  = "TAG_TYPE_PROJECT"
  title = "test_tf_provider"
  depends_on = [
    zedcloud_project.test_tf_provider
  ]
  tag_level_settings {
    flow_log_transmission =  "NETWORK_INSTANCE_FLOW_LOG_TRANSMISSION_DISABLED"
    interface_ordering = "INTERFACE_ORDERING_DISABLED"
  }
}

resource "zedcloud_brand" "test_tf_provider" {
  name        = "qemu"
  title       = "QEMU"
  description = "qemu"
  origin_type = "ORIGIN_LOCAL"
}

resource "zedcloud_model" "test_tf_provider" {
  brand_id    = zedcloud_brand.test_tf_provider.id
  name        = "test_tf_provider-create_edgenode"
  title       = "test_tf_provider-create_edgenode"
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

resource "zedcloud_datastore" "test_datastore" {
  depends_on = [
    zedcloud_project.test_tf_provider
  ]
  # required
  ds_fqdn             = "docker://docker.io"
  ds_path             = ""
  ds_type             = "DATASTORE_TYPE_CONTAINERREGISTRY"
  name                = "dockerhub"
  title               = "dockerhub"
  description         = "dockerhub"
  region              = "eu"
  project_access_list = [zedcloud_project.test_tf_provider.id]
}

resource "zedcloud_datastore" "alpha_ds" {
  depends_on = [
    zedcloud_project.test_tf_provider
  ]
  # required
  ds_fqdn             = "http://147.75.33.217"
  ds_path             = "images"
  ds_type             = "DATASTORE_TYPE_HTTP"
  name                = "alpha_ds"
  title               = "alpha_ds"
  description         = "alpha_ds"
  region              = "eu"
  project_access_list = [zedcloud_project.test_tf_provider.id]
}

resource "zedcloud_image" "alpine_image" {
  depends_on = [
    zedcloud_project.test_tf_provider,
    zedcloud_datastore.test_datastore
  ]
  name                = "alpine"
  datastore_id        = zedcloud_datastore.test_datastore.id
  image_arch          = "ARM64"
  image_format        = "CONTAINER"
  image_rel_url       = "alpine:latest"
  image_size_bytes    = 0
  image_type          = "IMAGE_TYPE_APPLICATION"
  title               = "alpine"
  project_access_list = [zedcloud_project.test_tf_provider.id]
}

resource "zedcloud_image" "silverpeak_cloudinit_azure_ds" {
  datastore_id = zedcloud_datastore.alpha_ds.id
  image_type = "IMAGE_TYPE_APPLICATION"
  image_arch = "AMD64"
  image_format = "ISO"
  image_sha256 = "290b2df3264618f8760fe115fc36b6e480ee9125699d4781867090e25eb65865"
  image_size_bytes =  1319
  name = "norfolk-cloudinit.iso"
  title = "norfolk-cloudinit.iso"
  project_access_list = [zedcloud_project.test_tf_provider.id]
  image_rel_url = "norfolk-cloudinit.iso"
}

# data "zedcloud_image" "alpine_image" {
#   name         = "alpine_image"
#   title        = "alpine"
#   datastore_id = zedcloud_datastore.test_datastore.id
#   image_arch   = "ARM64"
#   image_format = "CONTAINER"
#   depends_on = [
#     zedcloud_image.alpine_image
#   ]
# }

resource "zedcloud_edgenode" "test_tf_provider" {
  onboarding_key = "" # placeholder
  serialno       = "2293dbe8-29ce-420c-8264-962857efc46b"
  # required
  name       = "test_tf_provider"
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
  name     = "test_tf_provider"
  title    = "test_tf_provider"
  model_id = zedcloud_model.test_tf_provider.id
  depends_on = [
    zedcloud_edgenode.test_tf_provider
  ]
}

resource "zedcloud_application" "alpine_vm_app" {
  name                 = "alpine_vm_app"
  title                = "alpine_vm_app"
  user_defined_version = "24.0.4"
  depends_on = [ zedcloud_image.alpine_image ]
  project_access_list = [zedcloud_project.test_tf_provider.id]
  manifest {
    ac_kind             = "VMManifest"
    ac_version          = "1.2.0"
    name                = "alpine_24"
    vmmode              = "HV_HVM"
    enablevnc           = true
    app_type            = "APP_TYPE_VM"
    deployment_type     = "DEPLOYMENT_TYPE_STAND_ALONE"
    cpu_pinning_enabled = false
    resources {
      name  = "cpus"
      value = 2
    }
    resources {
      name  = "memory"
      value = 8000000
    }
    resources {
      name  = "storage"
      value = 102400000
    }
    images {
      imagename   = zedcloud_image.alpine_image.name
      imageid     = zedcloud_image.alpine_image.id
      imageformat = zedcloud_image.alpine_image.image_format
      cleartext   = false
      drvtype     = "HDD"
      ignorepurge = true
      maxsize     = 102400000
      target      = "Disk"
    }
    configuration {
      custom_config {
        add      = true
        name     = "cloud-config"
        override = true
        template = ""
      }
    }
  }
}

data "zedcloud_application" "alpine_vm_app" {
  name  = "alpine_vm_app"
  title = "alpine_vm_app"
  depends_on = [
    zedcloud_application.alpine_vm_app
  ]
}

# resource "zedcloud_project_deployment" "tf_deployment_project" {
#   # required
#   name  = "tf_deployment_project"
#   title = "tf_deployment_project"

#   # optional
#   type = "TAG_TYPE_DEPLOYMENT"

#   depends_on = [ 
#     zedcloud_application.test_tf_provider,
#     zedcloud_brand.test_tf_provider
#   ]

#   deployment {
#     name = "depl_project"
#     title = "deployment_project"
#     deployment_tag = "depl:1234"

#     device_policies {
#         meta_data {
#             name = "tf-edge-node"
#             title = "tf-edge-node"
#             policy_target_condition = {
#               "depl": "123"
#             }
#         }
#         policy_sub_type = "DEVICE_POLICY_TYPE_ATTESTATION"
#         attestation_policy {
#           type = "DEVICE_ATTEST_POLICY_TYPE_ACCEPT"
#         }
#     }

#     network_inst_policies {
#       meta_data {
#         name = "tf-network-instance"
#         title = "tf-network-instanceß"
#         policy_target_condition = {
#           "depl": "123"
#         }
#       }
#       net_inst_config {
#         port = "uplink"
#         kind = "NETWORK_INSTANCE_KIND_LOCAL"
#         type = "NETWORK_INSTANCE_DHCP_TYPE_V4"
#       }
#     }

#     app_inst_policies {
#       meta_data {
#         name = "tf-app-instance"
#         title = "tf-app-instance"
#         policy_target_condition = {
#           "depl": "123"
#         }
#       }
#       app_inst_config {
#         naming_scheme = "APP_NAMING_SCHEMEV2_PROJECT_APP_DEVICE"
#         name_project_part = "depl-project"
#         start_delay_in_seconds = 0
#         name_app_part = "test_tf_provider"
#         bundle_id = zedcloud_application.test_tf_provider.id
#         origin_type = "ORIGIN_LOCAL"
#         interfaces {
#           intforder = 1
#           intfname = "indirect"
#           directattach = false
#           netinstname = ""
#           privateip = false
#         }
#         tags = {
#           "depl": "123"
#         }
#         manifest_json {
#           name = "tf-app-instance"
#           ac_kind = "VMManifest"
#           ac_version = "1.2.0"
#           deployment_type = "DEPLOYMENT_TYPE_STAND_ALONE"
#           app_type = "APP_TYPE_VM"
#           owner {
#             user = "testuser"
#             group = "testgroup"
#             company = "Zededa Inc."
#             website = "https://www.zededa.com"
#             email = "test@zededa.com"
#           }
#           description = "test app instance"
#           cpu_pinning_enabled = false
#           images {
#             imagename = zedcloud_image.test_image.name
#             imageid = zedcloud_image.test_image.id
#             imageformat = "CONTAINER"
#             preserve = false
#             readonly = false
#             ignorepurge = false
#             cleartext = false
#           }
#           interfaces {}
#           desc {
#             category = "APP_CATEGORY_CLOUD_APPLICATION"
#             app_category = "APP_CATEGORY_UNSPECIFIED"
#             support = "support"
#           }
#           vmmode = "HV_HVM"
#           enablevnc = false
#           resources {
#             name = "resourceType"
#             value = "Custom"
#           }
#           resources {
#             name = "cpus"
#             value = "2"
#           }
#           resources {
#             name = "memory"
#             value = "1024000.00"
#           }
#           configuration {
#             custom_config {
#               name = "custom_config_name"
#               add = false
#               override = false
#               allow_storage_resize = false
#               field_delimiter = ""
#               template = ""
#             }
#           }
#         }
#       }
#     }

#     volume_inst_policies {
#       meta_data {
#         name = "tf-volume-instance"
#         title = "tf-volume-instance"
#         policy_target_condition = {
#           "depl": "123"
#         }
#       }
#       vol_inst_config {
#         type = "VOLUME_INSTANCE_TYPE_BLOCKSTORAGE"
#         accessmode = "VOLUME_INSTANCE_ACCESS_MODE_READWRITE"
#         size_bytes = 1048576
#         cleartext= true
#       }
#     }
#   }

#   edgeview_policy {
#     title = "tf-deployment-project"
#     name = "tf-edgeview-policy"
#     type = "POLICY_TYPE_EDGEVIEW"

#     edgeview_policy {
#       edgeviewcfg {
#         jwt_info {
#           disp_url = ""
#           allow_sec = 18000
#           num_inst = 1
#           encrypt = false
#         }
#         dev_policy {
#           allow_dev = true
#         }
#         app_policy {
#           allow_app = true
#         }
#       }
#       access_allow_change = true
#       max_expire_sec = 2592000
#       max_inst = 3
#       edgeview_allow = true
#     }
#   }
# }