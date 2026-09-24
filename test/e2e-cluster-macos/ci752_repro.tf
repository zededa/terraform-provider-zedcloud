# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# ===========================================================================
# CI-752 — opt in with `-var repro_ci752=true`
#
# https://zededa.atlassian.net/browse/CI-752
#   "Terraform - Errors after creating persistent storage to cluster"
#
# THE REPORT. A `zedcloud_volume_instance` is created against an edge-node
# cluster and applies cleanly. Every subsequent plan then shows, forever:
#
#     ~ resource "zedcloud_volume_instance" "ignition_vol1_persist" {
#         - device_id = "19a0f34f-..." -> null
#
#     ~ resource "zedcloud_network_instance" "demo_lan_1" {
#         - device_id = "931308d5-..." -> null
#
# The two ids are DIFFERENT nodes of the same three-node cluster, which is the
# detail the reporter noticed ("we appear to be creating an initial ID for
# cluster, and the individual node elements have their own ID") and the tell
# that the values are server-assigned rather than anything Terraform wrote.
#
# ROOT CAUSE — the same defect as CI-709, on the volume-instance path.
#
# SERVER. Create a volume instance with only `edgeNodeCluster` and no
# `deviceId` and the controller picks a node for it and persists the choice:
# srvs/seine/volinstproc.go:383-397 (`bindVolInstData`) calls
# `clusterproc.GetEdgeNodeForClusterObject`, which selects at RANDOM among the
# eligible, non-tie-breaker nodes (clusterproc.go:2843,
# `rand.Intn(len(eligibleNodes))`), and volinstproc.go:435 writes
# `volinst.DeviceId = dev.Id`. So every GET returns a populated `deviceId`, and
# on a multi-node cluster different objects legitimately land on different
# nodes — exactly the asymmetry in the ticket.
#
# PROVIDER, before the fix. `device_id` was `Optional` and NOT `Computed`.
# Refresh wrote the server's value into state, the configuration had none,
# Terraform planned `device_id -> null`, and Update PUT the whole model back.
# The diff never converges: `volinstUpdate` (volinstproc.go:1726-1734) only
# copies title and description, so the empty `deviceId` in the PUT is silently
# ignored and the server's value is still there on the next GET.
#
# THE FIX: PR #234 (44f74c4a), `Optional + Computed` on
# volume_instance.device_id, network_instance.{device_id,cluster_id} and
# application_instance.{device_id,cluster_id,app_type}. Merged to main, NOT in
# any release tag yet — v2.8.0 predates it.
#
# ---------------------------------------------------------------------------
# WHAT THIS FILE ADDS OVER test/e2e-cluster/ci709_repro.tf
#
# The Linux harness already covers this shape on a real 3-node cluster. This
# is the macOS/vfkit equivalent so the volume-instance path can be exercised
# without the Berlin lab, and it is deliberately NARROWER:
#
#   - a volume instance and a network instance, no application instance. The
#     app-instance leg needs a datastore, image and app definition, and the
#     CI-752 report is about volumes.
#   - it depends on `zedcloud_edgenode_cluster.CLUSTER` ONLY, never on
#     `zedamigo_wait_until.cluster_ready`. On macOS the cluster config never
#     reaches EVE (README "Two blockers found", B), so `cluster_ready` cannot
#     pass here. It does not have to: `deviceId` is assigned by the CONTROLLER
#     at create time, on the config plane, with no device round trip. The
#     cluster OBJECT existing is the whole prerequisite.
#
# ---------------------------------------------------------------------------
# EXPECTED RESULT, SINCE #234: both resources apply and then plan clean.
#
#   ZC_RUN_ID=<fresh> ./run.sh apply -var repro_ci752=true \
#     -target=zedcloud_volume_instance.CI752 \
#     -target=zedcloud_network_instance.CI752
#
#   ZC_RUN_ID=<same> ./run.sh plan -var repro_ci752=true -json \
#     | jq -r 'select(.type=="planned_change") | .change.resource.addr'
#
# A clean plan is only meaningful if the controller actually assigned a node —
# otherwise the test is vacuous. Check the raw API first:
#
#   VOL=$(./run.sh output -raw ci752_volume_instance_id)
#   curl -s --cacert ~/zedcloud-local/certs/zededa-root-ca.pem \
#     -H "Authorization: Bearer $TF_VAR_zedcloud_token" \
#     "https://zedcontrol.local.zededa.net/api/v1/volumes/instances/id/$VOL" \
#     | jq -c '{deviceId, edgeNodeCluster}'
#
# `deviceId` must be non-empty. If it is empty, nothing here is being tested.
# ===========================================================================

variable "repro_ci752" {
  type        = bool
  description = <<-EOT
    Create a cluster-scoped volume instance and network instance to exercise
    CI-752.

    Off by default: these need a formed cluster object, and on macOS the
    cluster leg is reached only after the long kube_ready barrier.
  EOT
  default     = false
}

locals {
  ci752 = var.repro_ci752 ? 1 : 0

  # eth0 is the only NIC on this topology (one gvproxy interface; no bridges
  # or taps exist on macOS), so it is both the management uplink and the only
  # port a network instance can bind to.
  ci752_port = "eth0"
}

# ---------------------------------------------------------------------------
# The volume instance — the resource named in CI-752.
# ---------------------------------------------------------------------------

resource "zedcloud_volume_instance" "CI752" {
  count = local.ci752

  name        = "tf_encm_ci752_vol_${var.run_id}"
  title       = "tf_encm_ci752_vol_${var.run_id}"
  description = "CI-752: cluster-scoped volume instance, device_id left to the controller"

  # EMPTYDIR rather than the ticket's BLOCKSTORAGE: BLOCKSTORAGE without a
  # content tree goes down volinstproc.go's size-validation path and, with a
  # content tree, down `alignDesignatedNodeWithContentTree`, which copies the
  # device from the content-tree volume instead of resolving one. EMPTYDIR
  # takes the plain `GetEdgeNodeForClusterObject` branch that the bug is about.
  type       = "VOLUME_INSTANCE_TYPE_EMPTYDIR"
  size_bytes = "1073741824"
  accessmode = "VOLUME_INSTANCE_ACCESS_MODE_READWRITE"
  cleartext  = true
  implicit   = false

  # *** THE POINT OF THE TEST: device_id is deliberately NOT set. ***
  #
  # The create guard in v2/resources/volume_instance.go:44-49 requires either
  # device_id or edge_node_cluster, so supplying only the cluster is the
  # supported cluster-scoped form, not a contrivance.
  edge_node_cluster {
    id = zedcloud_edgenode_cluster.CLUSTER.id
  }

  # Deliberately NOT depending on zedamigo_wait_until.cluster_ready — see the
  # header. The cluster object is the only prerequisite.
  depends_on = [zedcloud_edgenode_cluster.CLUSTER]
}

# ---------------------------------------------------------------------------
# The network instance, the ticket's second perpetual diff. Cheap to include
# and it exercises the second of the six fields #234 changed (`cluster_id`,
# which only this resource and application_instance carry).
# ---------------------------------------------------------------------------

resource "zedcloud_network_instance" "CI752" {
  count = local.ci752

  name  = "tf_encm_ci752_ni_${var.run_id}"
  title = "tf_encm_ci752_ni_${var.run_id}"

  kind = "NETWORK_INSTANCE_KIND_LOCAL"
  type = "NETWORK_INSTANCE_DHCP_TYPE_V4"
  port = local.ci752_port

  device_default = false
  dhcp           = false

  # Again: no device_id, and no cluster_id.
  edge_node_cluster {
    id = zedcloud_edgenode_cluster.CLUSTER.id
  }

  depends_on = [zedcloud_edgenode_cluster.CLUSTER]
}

# ---------------------------------------------------------------------------
# Outputs — the handles the verification steps in the header need.
# ---------------------------------------------------------------------------

#
# Guarded on `length(...)`, not on `local.ci752`. A staged build uses
# `-target=zedcloud_edgenode_cluster.CLUSTER` to reach the cluster without the
# cluster_ready barrier, and with -target the resources are excluded while the
# variable is still true -- so `local.ci752 == 1 ? ...[0]` fails the apply with
# "Invalid index ... empty tuple" after all the real work has succeeded.
output "ci752_volume_instance_id" {
  value = one(zedcloud_volume_instance.CI752[*].id)
}

output "ci752_network_instance_id" {
  value = one(zedcloud_network_instance.CI752[*].id)
}

# What the CONTROLLER chose. On a one-node cluster both must be the single
# node; on a multi-node cluster they may differ, and in the reporter's case
# they did.
output "ci752_assigned_devices" {
  description = "Controller-assigned device_id per cluster-scoped instance"
  value = length(zedcloud_volume_instance.CI752) == 0 ? null : {
    volume_instance  = zedcloud_volume_instance.CI752[0].device_id
    network_instance = one(zedcloud_network_instance.CI752[*].device_id)
    cluster_id       = one(zedcloud_network_instance.CI752[*].cluster_id)
    the_only_node    = zedcloud_edgenode.EN.id
  }
}
