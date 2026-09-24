# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# CI-836 — opt in with `-var repro_ci836=true`
#
# Two halves of one ticket.
#
# 1. THE FEATURE. Once nodes are in an edge-node cluster, EVE-OS has to be
#    upgraded through `PUT /v1/cluster/id/{id}/upgrade` rather than by setting
#    base_image per node, because the controller rolls the image out one node
#    at a time and migrates workloads as it goes. Terraform had no way to ask
#    for that at all. It does now: base_image on zedcloud_edgenode_cluster.
#
# 2. THE DIFF. That rollout writes baseImage onto every member node's device
#    config (srvs/seine/devproc.go clusterNodeUpgrade -> devBaseImagePublish +
#    devBaseImageApply), the provider refreshes it into state, and a config
#    that never mentions base_image was dragged into a permanent diff by it:
#
#      - base_image {
#          - activate   = true -> null
#          - image_name = "16.5.0-k-amd64" -> null
#        }
#
#    Permanent because PUT /v1/devices/id/{id} refuses to clear baseImage
#    (devproc.go: dev.BaseImage = exists.BaseImage), so the apply is a no-op
#    and the diff returns on the next plan. Fixed by the base_image
#    DiffSuppressFunc on the node schema; the unit regression test is
#    v2/resources/node_base_image_test.go, and this is the live counterpart.
#
# WHAT THIS NEEDS
#
# An IMAGE_TYPE_EVE image in IMAGE_STATUS_READY that the cluster can be pointed
# at. Pass its name in `ci836_eve_image`. It does NOT have to be downloadable:
# the bug is entirely on the config plane -- what the controller writes onto the
# device objects and what the provider then plans -- and asserting that needs
# no actual install or reboot. A placeholder image object is enough, and is a
# great deal quicker than a real rollout.
#
# HOW TO RUN
#
#   # 1. the cluster, as usual
#   ZC_RUN_ID=<fresh> ./run.sh apply -auto-approve \
#     -target=zedcloud_edgenode_cluster.CLUSTER
#
#   # 2. ask for the upgrade
#   ZC_RUN_ID=<same> ./run.sh apply -auto-approve \
#     -var repro_ci836=true -var ci836_eve_image=<image name> \
#     -target=zedcloud_edgenode_cluster.CLUSTER
#
#   # 3. the assertion: the node picked up base_image, and the plan is STILL
#   #    clean, because zedcloud_edgenode.EN never mentions base_image
#   ZC_RUN_ID=<same> ./run.sh plan \
#     -var repro_ci836=true -var ci836_eve_image=<image name>
#   # must say "No changes" for zedcloud_edgenode.EN
#
# Measured 2026-09-24 against the local controller, EVE 17.0.0-lts-k-arm64,
# with a placeholder image: the upgrade was accepted, the device picked up
# baseImage (revision 5 -> 6), and the node's plan stayed clean. Reverting the
# #201 DiffSuppressFunc and re-planning against the same objects reproduced the
# ticket's diff exactly.

variable "repro_ci836" {
  type        = bool
  description = <<-EOT
    Drive a cluster-scoped EVE-OS upgrade to exercise CI-836.

    Off by default: it needs a formed cluster object, and on macOS the cluster
    leg is reached only after the long kube_ready barrier.
  EOT
  default     = false
}

variable "ci836_eve_image" {
  type        = string
  description = <<-EOT
    Name of an IMAGE_TYPE_EVE / IMAGE_STATUS_READY image to roll out across the
    cluster. Required when repro_ci836 is true.

    A placeholder image object is sufficient -- see the header of this file.
  EOT
  default     = ""
}

check "ci836_image_name_supplied" {
  assert {
    condition     = !var.repro_ci836 || var.ci836_eve_image != ""
    error_message = "repro_ci836 = true needs -var ci836_eve_image=<name of a READY IMAGE_TYPE_EVE image>."
  }
}

# The upgrade rides on the EXISTING cluster, as a dynamic base_image block in
# cluster.tf gated on these variables. It cannot be a second
# zedcloud_edgenode_cluster resource: a node belongs to exactly one cluster, so
# a second one naming the same node would be rejected.
#
# With repro_ci836 = false the dynamic block produces nothing, so base_image is
# absent from the raw config -- which is the state the second half of the ticket
# is about, and what keeps the node's plan clean.

output "ci836_upgrade_status" {
  description = "Per-node progress of the cluster EVE-OS rollout."
  value       = var.repro_ci836 ? zedcloud_edgenode_cluster.CLUSTER.upgrade_status : null
}
