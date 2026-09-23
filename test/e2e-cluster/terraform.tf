# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

terraform {
  required_providers {
    # Resolved from the on-host filesystem mirror at
    # ~/.terraform.d/plugins/localhost/andrei-zededa/zedamigo/<ver>/<os>_<arch>/
    # (see e2e-cluster.tfrc). Not on a public registry.
    zedamigo = {
      source  = "localhost/andrei-zededa/zedamigo"
      version = "0.13.1"
    }

    # NO VERSION CONSTRAINT, deliberately -- resolved via dev_overrides against
    # the locally built provider. Same reasoning as test/e2e/terraform.tf: every
    # object this config creates is part of what a pull request can break, so it
    # has to be the PR's binary creating them.
    zedcloud = {
      source = "zededa/zedcloud"
    }
  }
}

provider "zedamigo" {
  # Only multi-node needs sudo: the bridge/tap/dhcp resources in
  # net_cluster.tf self-invoke the provider binary under `sudo -n`, and those
  # resources are gated on node_count > 1. A single-node run touches no host
  # networking and needs no privilege, exactly like test/e2e.
  #
  # CI NOTE: on h-m-dl20 the `github-runner` user has exactly three NOPASSWD
  # entries (ip, kill, taskset), which is NOT enough for the multi-node path --
  # a fourth entry for the bootstrapped provider binary path is required in
  # zededa-berlin-lab/hosts/h-m-dl20/github-runners.nix.
  #
  # INTERACTIVE RUNS ARE UNAFFECTED. Verified 2026-09-16: `ivan` on h-m-dl20
  # has `(ALL : ALL) SETENV: NOPASSWD: ALL`, so multi-node works by hand today.
  # The sudoers change is a CI-only prerequisite.
  use_sudo = var.node_count > 1

  # Remote mode: tofu runs on your laptop, the VMs run on the lab host. Set
  # var.zedamigo_target (+ var.zedamigo_ssh_user) to enable; leave null to run
  # tofu ON the target, which is what the earlier runs did.
  #
  # Every path in this config is already a path on the TARGET, so switching
  # modes changes nothing else.
  #
  # zedamigo dials the target DIRECTLY over TCP:22 with a Go SSH client, so it
  # never reads ~/.ssh/config and cannot use a ProxyCommand. The Berlin lab is
  # only reachable via the VPN container's HTTP CONNECT proxy, which is useless
  # to it. Use `zedamigo_ssh_proxy_jump`, or a local forward plus
  # `zedamigo_ssh_port` -- see the long note in vars.tf.
  target = var.zedamigo_target

  # `dynamic` inside a provider block is valid and confirmed working here
  # (OpenTofu 1.x): the provider configured far enough to attempt the SSH dial.
  dynamic "ssh" {
    for_each = var.zedamigo_target == null ? [] : [1]

    content {
      user          = var.zedamigo_ssh_user
      use_agent     = var.zedamigo_ssh_use_agent
      forward_agent = var.zedamigo_ssh_forward_agent
      port          = var.zedamigo_ssh_port
      proxy_jump    = var.zedamigo_ssh_proxy_jump

      # Escape hatch for the Go-knownhosts algorithm-mismatch trap; see vars.tf.
      insecure_ignore_host_key = var.zedamigo_ssh_insecure_ignore_host_key
    }
  }
}

provider "zedcloud" {
  zedcloud_url   = var.zedcloud_url
  zedcloud_token = var.zedcloud_token
}
