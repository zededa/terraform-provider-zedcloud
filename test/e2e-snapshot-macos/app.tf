# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# ---------------------------------------------------------------------------
# PLE-576: a small VM app on the node, and the snapshot resources, driven in
# stages by var.stage (see vars.tf and README.md).
#
# A VM app rather than a container: EVE snapshots the app's volumes, and a
# qcow2 VM disk is the case the feature is built for.
# ---------------------------------------------------------------------------

locals {
  app_on  = var.stage >= 1
  snap_on = var.stage >= 2 && var.keep_snapshot_resource
}

resource "zedcloud_datastore" "IMG" {
  count = local.app_on ? 1 : 0

  name    = "tf_snap_ds_${var.run_id}"
  title   = "tf_snap_ds_${var.run_id}"
  ds_type = "DATASTORE_TYPE_HTTPS"
  ds_fqdn = var.vm_image.fqdn
  ds_path = var.vm_image.path

  project_access_list = [zedcloud_project.PROJECT.id]
}

resource "zedcloud_image" "VM" {
  count = local.app_on ? 1 : 0

  name             = "tf_snap_img_${var.run_id}"
  title            = "tf_snap_img_${var.run_id}"
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
  name           = "tf_snap_ni_${var.run_id}"
  title          = "tf_snap_ni_${var.run_id}"
  kind           = "NETWORK_INSTANCE_KIND_LOCAL"
  port           = "eth0"
  type           = "NETWORK_INSTANCE_DHCP_TYPE_V4"
  device_default = false
  dhcp           = false

  depends_on = [zedamigo_wait_until.onboarded]
}

resource "zedcloud_application" "VM" {
  count = local.app_on ? 1 : 0

  name                 = "tf_snap_app_${var.run_id}"
  title                = "tf_snap_app_${var.run_id}"
  user_defined_version = "1.0"
  origin_type          = "ORIGIN_LOCAL"
  project_access_list  = [zedcloud_project.PROJECT.id]

  manifest {
    ac_kind         = "VMManifest"
    ac_version      = "1.2.0"
    name            = "tf_snap_app_${var.run_id}"
    app_type        = "APP_TYPE_VM"
    deployment_type = "DEPLOYMENT_TYPE_STAND_ALONE"
    vmmode          = "HV_HVM"
    enablevnc       = false

    owner {
      user    = "PLE-576 e2e"
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
      acls {
        matches {
          type  = "ip"
          value = "0.0.0.0/0"
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
  count = local.app_on ? 1 : 0

  name        = "tf_snap_appinst_${var.run_id}"
  title       = "tf_snap_appinst_${var.run_id}"
  description = var.app_description

  activate   = true
  app_id     = zedcloud_application.VM[0].id
  device_id  = zedcloud_edgenode.EN.id
  project_id = zedcloud_project.PROJECT.id
  app_type   = "APP_TYPE_VM"

  # PLE-576
  allow_snapshots    = true
  max_snapshots      = var.max_snapshots
  snapshot_on_update = var.snapshot_on_update

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
    acls {
      matches {
        type  = "ip"
        value = "0.0.0.0/0"
      }
    }
  }

  vminfo {
    cpus = var.app_cpus
    vnc  = false
    mode = "HV_HVM"
  }

  remote_console    = false
  is_secret_updated = false

  # Bump to purge: EVE replaces the app's volumes, and that is when it takes
  # an app-update snapshot requested with snapshot_on_update.
  dynamic "purge" {
    for_each = var.purge_counter > 0 ? [1] : []
    content {
      counter = var.purge_counter
    }
  }

  logs {
    access = false
  }

  manifest_info {
    transition_action = "INSTANCE_TA_NONE"
  }
}

# Stage 2: an immediate snapshot of the running app instance.
resource "zedcloud_app_instance_snapshot" "S" {
  count = local.snap_on ? 1 : 0

  app_instance_id = zedcloud_application_instance.VM[0].id
  triggers        = { take = var.snapshot_take }
}

# Changing cpus restarts the app, which is what makes EVE take an app-update
# snapshot requested by snapshot_on_update.
variable "purge_counter" {
  type    = number
  default = 0
}

variable "max_snapshots" {
  type    = number
  default = null
}

variable "app_cpus" {
  type    = number
  default = 1
}

variable "snapshot_take" {
  type    = string
  default = "1"
}

# Stage 4: roll back to the snapshot.
variable "rollback_snapshot_name" {
  type        = string
  description = "Snapshot to roll back to. Empty = the stage-2 zedcloud_app_instance_snapshot."
  default     = ""
}

variable "rollback_trigger" {
  type    = string
  default = "1"
}

resource "zedcloud_app_instance_snapshot_activation" "R" {
  count = var.stage >= 4 ? 1 : 0

  app_instance_id = zedcloud_application_instance.VM[0].id
  snapshot_name   = var.rollback_snapshot_name != "" ? var.rollback_snapshot_name : zedcloud_app_instance_snapshot.S[0].name
  triggers        = { run = var.rollback_trigger }
}

data "zedcloud_app_instance_snapshots" "ALL" {
  count           = local.app_on ? 1 : 0
  app_instance_id = zedcloud_application_instance.VM[0].id
}
