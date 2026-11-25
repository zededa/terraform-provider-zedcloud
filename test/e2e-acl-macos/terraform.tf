# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# ---------------------------------------------------------------------------
# UE-203 / ENG-3000 end-to-end: app instance interface ACLs on a virtual EVE
# node (vfkit + zedamigo) onboarded to a LOCAL zedcloud with #8548.
#
# Node setup copied from test/e2e-snapshot-macos (KVM EVE, one node, no
# cluster); read its README and test/e2e-cluster-macos's for the macOS traps.
# app.tf deploys VM app instances whose port maps differ from the bundle's.
# ---------------------------------------------------------------------------

terraform {
  required_providers {
    zedamigo = {
      source  = "localhost/andrei-zededa/zedamigo"
      version = "0.13.1"
    }

    zedcloud = {
      source = "zededa/zedcloud"
    }
  }
}

# No target, no sudo: tofu, vfkit and the VM all run on this Mac, and none of
# the resources used here touch host networking.
provider "zedamigo" {}

provider "zedcloud" {
  zedcloud_url   = var.zedcloud_url
  zedcloud_token = var.zedcloud_token
}
