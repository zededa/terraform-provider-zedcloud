# ENG-3000 repro from the Terraform side: bundle (manifest) ACLs vs
# instance ACLs. No EVE node needed -- this is controller-side storage only.
#
#   A: instance edits a bundle portmap (9022->22 becomes 8022->22)
#   B: instance omits acls entirely
#   C: instance echoes the bundle ACLs verbatim (control)

terraform {
  required_providers {
    zedcloud = { source = "zededa/zedcloud" }
  }
}

provider "zedcloud" {
  zedcloud_url   = var.zedcloud_url
  zedcloud_token = var.zedcloud_token
}

variable "zedcloud_url" { type = string }
variable "zedcloud_token" {
  type      = string
  sensitive = true
}
variable "run_id" {
  type    = string
  default = "eng3000"
}

locals {
  cases = toset(["a", "b", "c"])
}

resource "zedcloud_project" "P" {
  name  = "${var.run_id}_project"
  title = "${var.run_id}_project"
  type  = "TAG_TYPE_PROJECT"
  tag_level_settings {
    flow_log_transmission = "NETWORK_INSTANCE_FLOW_LOG_TRANSMISSION_UNSPECIFIED"
    interface_ordering    = "INTERFACE_ORDERING_ENABLED"
  }
}

resource "zedcloud_brand" "B" {
  name        = "${var.run_id}_brand"
  title       = "${var.run_id}_brand"
  origin_type = "ORIGIN_LOCAL"
}

resource "zedcloud_model" "M" {
  name           = "${var.run_id}_model"
  title          = "${var.run_id}_model"
  origin_type    = "ORIGIN_LOCAL"
  brand_id       = zedcloud_brand.B.id
  type           = "ARM64"
  product_status = "production"
  state          = "SYS_MODEL_STATE_ACTIVE"
  attr = {
    "Cpus"    = "2"
    "memory"  = "4096M"
    "storage" = "20000M"
  }
  io_member_list {
    assigngrp    = "eth0"
    cbattr       = {}
    cost         = 0
    logicallabel = "eth0"
    phyaddrs     = { Ifname = "eth0" }
    phylabel     = "eth0"
    usage        = "ADAPTER_USAGE_MANAGEMENT"
    usage_policy = {}
    ztype        = "IO_TYPE_ETH"
  }
}

resource "zedcloud_network" "N" {
  name       = "${var.run_id}_net"
  title      = "${var.run_id}_net"
  kind       = "NETWORK_KIND_V4"
  project_id = zedcloud_project.P.id
  mtu        = 1500
  ip { dhcp = "NETWORK_DHCP_TYPE_CLIENT" }
}

resource "zedcloud_edgenode" "EN" {
  for_each       = local.cases
  name           = "${var.run_id}_en_${each.key}"
  title          = "${var.run_id}_en_${each.key}"
  serialno       = "${var.run_id}-serial-${each.key}"
  onboarding_key = "5d0767ee-0547-4569-b530-387e526f8cb9"
  model_id       = zedcloud_model.M.id
  project_id     = zedcloud_project.P.id
  admin_state    = "ADMIN_STATE_ACTIVE"
  interfaces {
    intfname   = "eth0"
    intf_usage = "ADAPTER_USAGE_MANAGEMENT"
    net_dhcp   = "NETWORK_DHCP_TYPE_CLIENT"
    cost       = 0
    netname    = zedcloud_network.N.name
    ztype      = "IO_TYPE_ETH"
    tags       = {}
  }
  tags = {}
}

resource "zedcloud_network_instance" "NI" {
  for_each       = local.cases
  device_id      = zedcloud_edgenode.EN[each.key].id
  name           = "${var.run_id}_ni_${each.key}"
  title          = "${var.run_id}_ni_${each.key}"
  kind           = "NETWORK_INSTANCE_KIND_LOCAL"
  port           = "eth0"
  type           = "NETWORK_INSTANCE_DHCP_TYPE_V4"
  device_default = false
  dhcp           = false
}

resource "zedcloud_datastore" "DS" {
  name                = "${var.run_id}_ds"
  title               = "${var.run_id}_ds"
  ds_type             = "DATASTORE_TYPE_HTTPS"
  ds_fqdn             = "https://download.cirros-cloud.net"
  ds_path             = "0.6.2"
  project_access_list = [zedcloud_project.P.id]
}

resource "zedcloud_image" "IMG" {
  name                = "${var.run_id}_img"
  title               = "${var.run_id}_img"
  datastore_id        = zedcloud_datastore.DS.id
  image_rel_url       = "cirros-0.6.2-aarch64-disk.img"
  image_sha256        = "6a2222bd7ba89384d35378b5065bd5e838857f25db08c7abbff99b6d9bb48986"
  image_size_bytes    = "24969728"
  image_arch          = "ARM64"
  image_format        = "QCOW2"
  image_type          = "IMAGE_TYPE_APPLICATION"
  project_access_list = [zedcloud_project.P.id]
}

# Bundle with the ticket's eth0 rules: allow 0.0.0.0/0, 9022->22, 9080->80.
resource "zedcloud_application" "APP" {
  name                 = "${var.run_id}_app"
  title                = "${var.run_id}_app"
  user_defined_version = "1.0"
  origin_type          = "ORIGIN_LOCAL"
  project_access_list  = [zedcloud_project.P.id]

  manifest {
    ac_kind         = "VMManifest"
    ac_version      = "1.2.0"
    name            = "${var.run_id}_app"
    app_type        = "APP_TYPE_VM"
    deployment_type = "DEPLOYMENT_TYPE_STAND_ALONE"
    vmmode          = "HV_HVM"
    enablevnc       = false

    owner {
      user    = "ENG-3000 repro"
      email   = "noreply@zededa.com"
      website = "www.zededa.com"
    }

    images {
      imagename   = zedcloud_image.IMG.name
      imageid     = zedcloud_image.IMG.id
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
      acls {
        matches {
          type  = "protocol"
          value = "tcp"
        }
        matches {
          type  = "lport"
          value = "9022"
        }
        matches {
          type  = "ip"
          value = "0.0.0.0/0"
        }
        actions {
          portmap = true
          portmapto { app_port = 22 }
        }
      }
      acls {
        matches {
          type  = "protocol"
          value = "tcp"
        }
        matches {
          type  = "lport"
          value = "9080"
        }
        matches {
          type  = "ip"
          value = "0.0.0.0/0"
        }
        actions {
          portmap = true
          portmapto { app_port = 80 }
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
      custom_config { add = false }
    }
  }
}

locals {
  # Instance-side rules per case. b = none, c = exact bundle copy.
  allow = [{ m = [["ip", "0.0.0.0/0"]], port = null }]
  pm = {
    p8022 = { m = [["protocol", "tcp"], ["lport", "8022"], ["ip", "0.0.0.0/0"]], port = 22 }
    p9022 = { m = [["protocol", "tcp"], ["lport", "9022"], ["ip", "0.0.0.0/0"]], port = 22 }
    p9080 = { m = [["protocol", "tcp"], ["lport", "9080"], ["ip", "0.0.0.0/0"]], port = 80 }
  }
  inst_acls = {
    a = concat(local.allow, [local.pm.p8022, local.pm.p9080])
    b = []
    c = concat(local.allow, [local.pm.p9022, local.pm.p9080])
  }
}

resource "zedcloud_application_instance" "AI" {
  for_each   = local.cases
  name       = "${var.run_id}_ai_${each.key}"
  title      = "${var.run_id}_ai_${each.key}"
  activate   = true
  app_id     = zedcloud_application.APP.id
  device_id  = zedcloud_edgenode.EN[each.key].id
  project_id = zedcloud_project.P.id
  app_type   = "APP_TYPE_VM"

  drives {
    imagename = zedcloud_image.IMG.name
    maxsize   = "0"
    preserve  = false
    readonly  = false
    drvtype   = "HDD"
    target    = "Disk"
  }

  interfaces {
    intfname    = "eth0"
    intforder   = 1
    netinstname = zedcloud_network_instance.NI[each.key].name
    privateip   = false

    dynamic "acls" {
      for_each = local.inst_acls[each.key]
      content {
        dynamic "matches" {
          for_each = acls.value.m
          content {
            type  = matches.value[0]
            value = matches.value[1]
          }
        }
        dynamic "actions" {
          for_each = acls.value.port == null ? [] : [acls.value.port]
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
  logs { access = false }
  manifest_info { transition_action = "INSTANCE_TA_NONE" }
}

output "instance_ids" {
  value = { for k, v in zedcloud_application_instance.AI : k => v.id }
}
