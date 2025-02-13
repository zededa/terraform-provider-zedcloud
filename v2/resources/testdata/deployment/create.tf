resource "zedcloud_project" "test_tf_project" {
  # required
  name  = "test_tf_project-deployment"
  title = "test_tf_project-deployment"
  type = "TAG_TYPE_DEPLOYMENT"
  tag_level_settings {
    flow_log_transmission = "NETWORK_INSTANCE_FLOW_LOG_TRANSMISSION_DISABLED"
    interface_ordering    = "INTERFACE_ORDERING_DISABLED"
  }
}

resource "zedcloud_brand" "test_tf_provider" {
  name        = "test_tf_provider-qemu"
  title       = "test_tf_provider-QEMU"
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
    zedcloud_project.test_tf_project
  ]
  # required
  ds_fqdn             = "docker://docker.io"
  ds_path             = ""
  ds_type             = "DATASTORE_TYPE_CONTAINERREGISTRY"
  name                = "test_tf_provider-dockerhub"
  title               = "test_tf_provider-dockerhub"
  description         = "test_tf_provider-dockerhub"
  region              = "eu"
  project_access_list = [zedcloud_project.test_tf_project.id]
}

resource "zedcloud_image" "test_image" {
  depends_on = [
    zedcloud_project.test_tf_project,
    zedcloud_datastore.test_datastore
  ]
  name                = "test_tf_provider-alpine"
  datastore_id        = zedcloud_datastore.test_datastore.id
  image_arch          = "ARM64"
  image_format        = "CONTAINER"
  image_rel_url       = "alpine:latest"
  image_size_bytes    = 1024
  image_type          = "IMAGE_TYPE_APPLICATION"
  title               = "alpine"
  project_access_list = [zedcloud_project.test_tf_project.id]
}


resource "zedcloud_application" "alpine_vm_app" {
  name                 = "test_tf_provider-alpine_vm_app"
  title                = "test_tf_provider-alpine_vm_app"
  user_defined_version = "24.0.4"
  depends_on           = [zedcloud_image.test_image]
  project_access_list  = [zedcloud_project.test_tf_project.id]
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
  }
}

resource "zedcloud_deployment" "tf_deployment" {
  depends_on = [
    zedcloud_application.alpine_vm_app,
    zedcloud_brand.test_tf_provider,
    zedcloud_project.test_tf_project
  ]

  name           = "test_tf_provider-deployment"
  title          = "test_tf_provider-deployment"
  deployment_tag = "depl:1234"

  project_id = zedcloud_project.test_tf_project.id

  device_policies {
    meta_data {
      name  = "test_tf_provider-edge-node-policy"
      title = "test_tf_provider-edge-node-policy"
      policy_target_condition = {
        "depl" : "1234"
      }
      tags = {
        "depl" : "1234"
      }
    }
    policy_sub_type = "DEVICE_POLICY_TYPE_ATTESTATION"
    attestation_policy {
      type = "DEVICE_ATTEST_POLICY_TYPE_ACCEPT"
    }
  }

  network_inst_policies {
    meta_data {
      name  = "test_tf_provider-network-instance-policy"
      title = "test_tf_provider-network-instance-policy"
      policy_target_condition = {
        "depl" : "1234"
      }
    }
    net_inst_config {
      port = "uplink"
      kind = "NETWORK_INSTANCE_KIND_LOCAL"
      type = "NETWORK_INSTANCE_DHCP_TYPE_V4"
    }
  }

  app_inst_policies {
    meta_data {
      name  = "test_tf_provider-app-instance-policy"
      title = "test_tf_provider-app-instance-policy"
      policy_target_condition = {
        "depl" : "1234"
      }
      tags = {
        "depl" : "1234"
      }
    }
    app_inst_config {
      naming_scheme          = "APP_NAMING_SCHEMEV2_PROJECT_APP_DEVICE"
      name_project_part      = "depl-project"
      start_delay_in_seconds = 0
      name_app_part          = "test_tf_provider"
      bundle_id              = zedcloud_application.alpine_vm_app.id
      origin_type            = "ORIGIN_LOCAL"
      interfaces {
        intforder    = 1
        intfname     = "indirect"
        directattach = false
        netinstname  = ""
        privateip    = false
      }
      tags = {
        "depl" : "1234"
      }
      manifest_json {
        name            = "tf-app-instance"
        ac_kind         = "VMManifest"
        ac_version      = "1.2.0"
        deployment_type = "DEPLOYMENT_TYPE_STAND_ALONE"
        app_type        = "APP_TYPE_VM"
        owner {
          user    = "testuser"
          group   = "testgroup"
          company = "Zededa Inc."
          website = "https://www.zededa.com"
          email   = "test@zededa.com"
        }
        description         = "test app instance"
        cpu_pinning_enabled = false
        images {
          imagename   = zedcloud_image.test_image.name
          imageid     = zedcloud_image.test_image.id
          maxsize = "0"
          imageformat = "CONTAINER"
          preserve    = false
          readonly    = false
          ignorepurge = false
          cleartext   = false
        }
        desc {
          category     = "APP_CATEGORY_CLOUD_APPLICATION"
          app_category = "APP_CATEGORY_UNSPECIFIED"
          support      = "support"
        }
        vmmode    = "HV_HVM"
        enablevnc = false
        resources {
          name  = "resourceType"
          value = "Custom"
        }
        resources {
          name  = "cpus"
          value = "2"
        }
        resources {
          name  = "memory"
          value = "1024000.00"
        }
        configuration {
          custom_config {
            name                 = "custom_config_name"
            add                  = false
            override             = false
            allow_storage_resize = false
            field_delimiter      = ""
            template             = ""
          }
        }
      }
    }
  }

  volume_inst_policies {
    meta_data {
      name  = "test_tf_provider-volume-instance-policy"
      title = "test_tf_provider-volume-instance-policy"
      policy_target_condition = {
        "depl" : "1234"
      }
      tags = {
        "depl" : "1234"
      }
    }
    vol_inst_config {
      type       = "VOLUME_INSTANCE_TYPE_BLOCKSTORAGE"
      accessmode = "VOLUME_INSTANCE_ACCESS_MODE_READWRITE"
      size_bytes = "1048576"
      cleartext  = true
    }
  }

  edgeview_policy {
    meta_data {
      name  = "test_tf_provider-edgeview-policy"
      title = "test_tf_provider-edgeview-policy"
      policy_target_condition = {
        "depl" : "1234"
      }
    }
    edgeviewcfg {
      jwt_info {
        disp_url  = "zedcloud.local.zededa.net/api/v1/edge-view"
        allow_sec = 18000
        num_inst  = 1
        encrypt   = false
      }
      dev_policy {
        allow_dev = true
      }
      app_policy {
        allow_app = true
      }

      ext_policy {
        allow_ext = false
      }
    }
    access_allow_change = true
    max_expire_sec      = 2592000
    max_inst            = 3
    edgeview_allow      = true
  }
}

