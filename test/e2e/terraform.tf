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

    # Resolved from a dev_overrides entry pointing at the locally built
    # provider -- see the generated e2e.tfrc in .github/workflows/e2e.yml,
    # and test/e2e/README.md for the manual equivalent.
    #
    # NO VERSION CONSTRAINT, deliberately. dev_overrides ignores version
    # constraints (with a warning), so a pin here would be decorative and
    # would imply this comes from the registry. It used to: iteration 1
    # pinned ">= 2.7.0" from the public registry on the grounds that we were
    # proving out the node lifecycle rather than the provider. That stopped
    # being true once the acceptance suite ran -- every object this config
    # creates (project, brand, model, network instance, image, edgenode) is
    # part of what a pull request can break, so it has to be the PR's binary
    # creating them.
    zedcloud = {
      source = "zededa/zedcloud"
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
