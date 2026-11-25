# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# ---------------------------------------------------------------------------
# UE-203 / ENG-3000: VM app instances whose port maps differ from the
# bundle's, on one node. See vars.tf (stage) and README.md.
#
# The bundle's eth0 rules: allow 0.0.0.0/0, 9022->22, 9080->80.
#   X (stage >= 1): allow, var.x_ssh_port->22, 9080->80   (edited default)
#   Y (stage >= 2): allow, 7022->22, 7080->80             (both moved)
# Before #8548 the controller added 9022->22 back to X, and Y got 9022 and
# 9080 too, so EVE failed Y with "overlapping portmaps".
# ---------------------------------------------------------------------------

locals {
  app_on = var.stage >= 1

  rules = {
    allow = { m = [["ip", "0.0.0.0/0"]], port = null }
    xssh  = { m = [["protocol", "tcp"], ["lport", var.x_ssh_port], ["ip", "0.0.0.0/0"]], port = 22 }
    p9022 = { m = [["protocol", "tcp"], ["lport", "9022"], ["ip", "0.0.0.0/0"]], port = 22 }
    p9080 = { m = [["protocol", "tcp"], ["lport", "9080"], ["ip", "0.0.0.0/0"]], port = 80 }
    p7022 = { m = [["protocol", "tcp"], ["lport", "7022"], ["ip", "0.0.0.0/0"]], port = 22 }
    p7080 = { m = [["protocol", "tcp"], ["lport", "7080"], ["ip", "0.0.0.0/0"]], port = 80 }
  }

  instances = merge(
    var.stage >= 1 ? { x = ["allow", "xssh", "p9080"] } : {},
    var.stage >= 2 ? { y = ["allow", "p7022", "p7080"] } : {},
  )
}

resource "zedcloud_datastore" "IMG" {
  count = local.app_on ? 1 : 0

  name    = "tf_acl_ds_${var.run_id}"
  title   = "tf_acl_ds_${var.run_id}"
  ds_type = "DATASTORE_TYPE_HTTPS"
  ds_fqdn = var.vm_image.fqdn
  ds_path = var.vm_image.path

  project_access_list = [zedcloud_project.PROJECT.id]
}

resource "zedcloud_image" "VM" {
  count = local.app_on ? 1 : 0

  name             = "tf_acl_img_${var.run_id}"
  title            = "tf_acl_img_${var.run_id}"
  datastore_id     = zedcloud_datastore.IMG[0].id
  image_rel_url    = var.vm_image.file
  image_sha256     = var.vm_image.sha256
  image_size_bytes = var.vm_image.size
  image_arch       = "ARM64"
  image_format     = "QCOW2"
  image_type       = "IMAGE_TYPE_APPLICATION"

  project_access_list = [zedcloud_project.PROJECT.id]
}

resource "zedcloud_network_instance" "LOCAL" {
  count = local.app_on ? 1 : 0

  device_id      = zedcloud_edgenode.EN.id
  name           = "tf_acl_ni_${var.run_id}"
  title          = "tf_acl_ni_${var.run_id}"
  kind           = "NETWORK_INSTANCE_KIND_LOCAL"
  port           = "eth0"
  type           = "NETWORK_INSTANCE_DHCP_TYPE_V4"
  device_default = false
  dhcp           = false

  depends_on = [zedamigo_wait_until.onboarded]
}

resource "zedcloud_application" "VM" {
  count = local.app_on ? 1 : 0

  name                 = "tf_acl_app_${var.run_id}"
  title                = "tf_acl_app_${var.run_id}"
  user_defined_version = "1.0"
  origin_type          = "ORIGIN_LOCAL"
  project_access_list  = [zedcloud_project.PROJECT.id]

  manifest {
    ac_kind         = "VMManifest"
    ac_version      = "1.2.0"
    name            = "tf_acl_app_${var.run_id}"
    app_type        = "APP_TYPE_VM"
    deployment_type = "DEPLOYMENT_TYPE_STAND_ALONE"
    vmmode          = "HV_HVM"
    enablevnc       = false

    owner {
      user    = "UE-203 e2e"
      email   = "noreply@zededa.com"
      website = "www.zededa.com"
    }

    images {
      imagename   = zedcloud_image.VM[0].name
      imageid     = zedcloud_image.VM[0].id
      imageformat = "QCOW2"
      maxsize     = "0"
      drvtype     = "HDD"
      target      = "Disk"
      cleartext   = false
      ignorepurge = false
      preserve    = false
      readonly    = false
    }

    interfaces {
      name         = "eth0"
      directattach = false
      dynamic "acls" {
        for_each = ["allow", "p9022", "p9080"]
        content {
          dynamic "matches" {
            for_each = local.rules[acls.value].m
            content {
              type  = matches.value[0]
              value = matches.value[1]
            }
          }
          dynamic "actions" {
            for_each = local.rules[acls.value].port == null ? [] : [local.rules[acls.value].port]
            content {
              portmap = true
              portmapto { app_port = actions.value }
            }
          }
        }
      }
    }

    desc {
      category     = "APP_CATEGORY_OTHERS"
      app_category = "APP_CATEGORY_OTHERS"
    }

    resources {
      name  = "resourceType"
      value = "Tiny"
    }
    resources {
      name  = "cpus"
      value = "1"
    }
    resources {
      name  = "memory"
      value = "262144.00"
    }

    configuration {
      custom_config {
        add = false
      }
    }
  }
}

resource "zedcloud_application_instance" "VM" {
  for_each = local.instances

  name       = "tf_acl_appinst_${each.key}_${var.run_id}"
  title      = "tf_acl_appinst_${each.key}_${var.run_id}"
  activate   = true
  app_id     = zedcloud_application.VM[0].id
  device_id  = zedcloud_edgenode.EN.id
  project_id = zedcloud_project.PROJECT.id
  app_type   = "APP_TYPE_VM"

  drives {
    imagename = zedcloud_image.VM[0].name
    maxsize   = "0"
    preserve  = false
    readonly  = false
    drvtype   = "HDD"
    target    = "Disk"
  }

  interfaces {
    intfname    = "eth0"
    intforder   = 1
    netinstname = zedcloud_network_instance.LOCAL[0].name
    privateip   = false

    dynamic "acls" {
      for_each = each.value
      content {
        dynamic "matches" {
          for_each = local.rules[acls.value].m
          content {
            type  = matches.value[0]
            value = matches.value[1]
          }
        }
        dynamic "actions" {
          for_each = local.rules[acls.value].port == null ? [] : [local.rules[acls.value].port]
          content {
            portmap = true
            mapparams { port = actions.value }
          }
        }
      }
    }
  }

  vminfo {
    cpus = 1
    vnc  = false
    mode = "HV_HVM"
  }

  remote_console    = false
  is_secret_updated = false

  logs {
    access = false
  }
}
