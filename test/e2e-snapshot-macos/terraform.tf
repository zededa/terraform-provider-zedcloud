# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# ---------------------------------------------------------------------------
# PLE-576 end-to-end: immediate app instance snapshots against a LOCAL
# zedcloud, on one virtual EVE node on this Mac (vfkit + zedamigo).
#
# Derived from test/e2e-cluster-macos (read its README for the macOS traps:
# leaked TF_VAR_*, the trailing newline in additional_hosts, no eve_nuke_disks,
# the TPM-less EdgeNodeCert flap). Differences:
#
#   - EVE is the KVM flavour (`-kvm-`), not `-k-`: the controller refuses
#     snapshots on KubeVirt devices.
#   - no kube readiness and no cluster: snapshots are rejected for app
#     instances on an edge node cluster (PLE-576 D3).
#   - app.tf deploys a small VM app and drives the snapshot resources in
#     stages (var.stage), see README.md.
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
