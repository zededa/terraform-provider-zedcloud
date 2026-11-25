# UE-203 / ENG-3000: interface ACLs on zedcloud_application_instance against
# a controller with zedcloud #8548. No EVE node: this checks what the
# controller stores and what the provider plans. Steps are driven by
# var.instances (see README.md).

terraform {
  required_providers {
    zedcloud = { source = "zededa/zedcloud" }
  }
}

provider "zedcloud" {
  zedcloud_url   = var.zedcloud_url
  zedcloud_token = var.zedcloud_token
  async_writes   = var.async_writes
}

variable "zedcloud_url" { type = string }
variable "zedcloud_token" {
  type      = string
  sensitive = true
}
variable "run_id" {
  type    = string
  default = "ue203"
}
variable "async_writes" {
  type    = bool
  default = true
}

# Instance name suffix -> project ("on" = interface ordering enabled, "off" =
# disabled), eth0 ACL rule keys (null = no acls block) and cpus.
variable "instances" {
  type = map(object({
    project = string
    acls    = optional(list(string))
    cpus    = optional(number, 1)
  }))
}

locals {
  projects = toset(["on", "off"])
  # Rule keys usable in var.instances.
  rules = {
    allow = { m = [["ip", "0.0.0.0/0"]], port = null }
    p6022 = { m = [["protocol", "tcp"], ["lport", "6022"], ["ip", "0.0.0.0/0"]], port = 22 }
    p7022 = { m = [["protocol", "tcp"], ["lport", "7022"], ["ip", "0.0.0.0/0"]], port = 22 }
    p8022 = { m = [["protocol", "tcp"], ["lport", "8022"], ["ip", "0.0.0.0/0"]], port = 22 }
    p9022 = { m = [["protocol", "tcp"], ["lport", "9022"], ["ip", "0.0.0.0/0"]], port = 22 }
    p9080 = { m = [["protocol", "tcp"], ["lport", "9080"], ["ip", "0.0.0.0/0"]], port = 80 }
  }
}

resource "zedcloud_project" "P" {
  for_each = local.projects
  name     = "${var.run_id}_project_${each.key}"
  title    = "${var.run_id}_project_${each.key}"
  type     = "TAG_TYPE_PROJECT"
  tag_level_settings {
    flow_log_transmission = "NETWORK_INSTANCE_FLOW_LOG_TRANSMISSION_UNSPECIFIED"
    interface_ordering    = each.key == "on" ? "INTERFACE_ORDERING_ENABLED" : "INTERFACE_ORDERING_DISABLED"
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
    "Cpus"    = "4"
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
  for_each   = local.projects
  name       = "${var.run_id}_net_${each.key}"
  title      = "${var.run_id}_net_${each.key}"
  kind       = "NETWORK_KIND_V4"
  project_id = zedcloud_project.P[each.key].id
  mtu        = 1500
  ip { dhcp = "NETWORK_DHCP_TYPE_CLIENT" }
}

resource "zedcloud_edgenode" "EN" {
  for_each       = var.instances
  name           = "${var.run_id}_en_${each.key}"
  title          = "${var.run_id}_en_${each.key}"
  serialno       = "${var.run_id}-serial-${each.key}"
  onboarding_key = "5d0767ee-0547-4569-b530-387e526f8cb9"
  model_id       = zedcloud_model.M.id
  project_id     = zedcloud_project.P[each.value.project].id
  admin_state    = "ADMIN_STATE_ACTIVE"
  interfaces {
    intfname   = "eth0"
    intf_usage = "ADAPTER_USAGE_MANAGEMENT"
    net_dhcp   = "NETWORK_DHCP_TYPE_CLIENT"
    cost       = 0
    netname    = zedcloud_network.N[each.value.project].name
    ztype      = "IO_TYPE_ETH"
    tags       = {}
  }
  tags = {}
}

resource "zedcloud_network_instance" "NI" {
  for_each       = var.instances
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
  project_access_list = sort([for p in zedcloud_project.P : p.id]) # the controller sorts it
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
  project_access_list = sort([for p in zedcloud_project.P : p.id]) # the controller sorts it
}

# Bundle with the ENG-3000 eth0 rules: allow 0.0.0.0/0, 9022->22, 9080->80.
resource "zedcloud_application" "APP" {
  name                 = "${var.run_id}_app"
  title                = "${var.run_id}_app"
  user_defined_version = "1.0"
  origin_type          = "ORIGIN_LOCAL"
  project_access_list  = sort([for p in zedcloud_project.P : p.id])

  manifest {
    ac_kind         = "VMManifest"
    ac_version      = "1.2.0"
    name            = "${var.run_id}_app"
    app_type        = "APP_TYPE_VM"
    deployment_type = "DEPLOYMENT_TYPE_STAND_ALONE"
    vmmode          = "HV_HVM"
    enablevnc       = false

    owner {
      user    = "UE-203 test"
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
      custom_config { add = false }
    }
  }
}

resource "zedcloud_application_instance" "AI" {
  for_each   = var.instances
  name       = "${var.run_id}_ai_${each.key}"
  title      = "${var.run_id}_ai_${each.key}"
  activate   = true
  app_id     = zedcloud_application.APP.id
  device_id  = zedcloud_edgenode.EN[each.key].id
  project_id = zedcloud_project.P[each.value.project].id
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
      for_each = each.value.acls == null ? [] : each.value.acls
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
    cpus = each.value.cpus
    vnc  = false
    mode = "HV_HVM"
  }

  remote_console    = false
  is_secret_updated = false
  logs { access = false }
  # manifest_info is left out: the controller fills it (and its details
  # change on app updates), which is unrelated to UE-203.
}

output "instance_ids" {
  value = { for k, v in zedcloud_application_instance.AI : k => v.id }
}
