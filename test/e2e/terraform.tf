# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

terraform {
  required_providers {
    # Resolved from the on-host filesystem mirror at
    # ~/.terraform.d/plugins/localhost/andrei-zededa/zedamigo/<ver>/linux_amd64/
    # (see e2e.tfrc). Not on a public registry.
    zedamigo = {
      source  = "localhost/andrei-zededa/zedamigo"
      version = "0.13.1"
    }

    # Public registry. Iteration 1 uses the released provider because we are
    # proving out the node lifecycle, not the provider under test. Later
    # iterations run the acceptance suite, which loads the provider in-process
    # and therefore does not use this binary at all.
    zedcloud = {
      source  = "zededa/zedcloud"
      version = ">= 2.7.0"
    }
  }
}

provider "zedamigo" {
  # Iteration 1 runs OpenTofu on the lab host itself, so the target is the
  # local machine and no ssh{} block is needed. Switching to a GitHub-hosted
  # runner later means setting `target` + `ssh {}` here and changing nothing
  # else -- every path in this config is already a path on the target.
  #
  # target = var.lab_host
  # ssh { user = "gh-run-tf-zedcloud" ... }

  # No host networking resources are used (the node gets its uplink from QEMU
  # SLIRP), so the provider needs no sudo.
  use_sudo = false
}

provider "zedcloud" {
  zedcloud_url   = var.zedcloud_url
  zedcloud_token = var.zedcloud_token
}
