# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# ---------------------------------------------------------------------------
# macOS / vfkit sibling of test/e2e-cluster.
#
# Deliberately a SEPARATE DIRECTORY rather than a flag on test/e2e-cluster.
# Six things differ structurally, not by a boolean:
#
#   1. zedamigo_host_reservation is Linux-only, so there is no reservation
#      phase at all and cpus/mem come straight from variables.
#   2. bridge / tap / dhcp_server are Linux-only, so multi-node is impossible
#      here -- a single-node cluster on eth0 is the ONLY shape.
#   3. The serial flow is REVERSED. Apple's Virtualization.framework cannot
#      stamp an SMBIOS serial, so EVE generates a soft serial at install time
#      and the cloud object has to take its serialno FROM the install, not the
#      other way round. That inverts a dependency edge.
#   4. The installer is built as `raw`, not `iso`, and consumed via
#      installer_raw.
#   5. Networking is gvproxy, not SLIRP.
#   6. It points at a LOCAL controller, which needs tls_ca / object_signing_ca
#      baked into the installer and an /etc/hosts entry for the host.
#
# Folding all six into test/e2e-cluster would make that config unreadable and
# would put an unverified macOS path into the PR that carries the verified
# Linux one.
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
