# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# ===========================================================================
# THE EDGE-NODE CLUSTER -- one node.
#
# Zedcloud allows this (the controller's cluster logic):
#   a required node count of three, but 1 is permitted and the controller's master-node check
#   explicitly permits a single-node cluster with one master. 2 is rejected -- no
#   quorum. There is no node-count guard on the device mutation path, so the
#   controller writes the same cluster fields for one node as for three, which
#   is what makes a single node a usable reproduction vehicle for CI-834 and
#   CI-709.
#
# Gates this still has to clear, all server-side:
#   the controller's cluster-membership checks  -- node AdminState == DEVICE_REGISTERED (i.e.
#     onboarded, not merely ADMIN_STATE_ACTIVE); cluster_interface resolves;
#     its intf_usage != ADAPTER_USAGE_UNSPECIFIED; no active app instance
#   the controller's live node-capability check -- a LIVE status query; the node must report
#     OptionalCapabilities.hvTypeKubevirt or it is "node %s does not support
#     kubevirt"
# ===========================================================================

resource "zedcloud_edgenode_cluster" "CLUSTER" {
  name        = "tf_encm_cluster_${var.run_id}"
  title       = "tf_encm_cluster_${var.run_id}"
  description = "UE-156 macOS/vfkit: 1-node edge-node cluster against the local controller"
  project_id  = zedcloud_project.PROJECT.id

  cluster_prefix = var.cluster_prefix

  nodes {
    id                = zedcloud_edgenode.EN.id
    cluster_interface = local.cluster_interface
    node_type         = "EDGE_NODE_CLUSTER_NODE_TYPE_SERVER"
  }

  # CI-836: the cluster-scoped EVE-OS upgrade. Opt in with
  # `-var repro_ci836=true -var ci836_eve_image=<name>`; see ci836_repro.tf.
  #
  # Off by default the block is absent from the config entirely, which is the
  # shape the diff half of CI-836 is about -- the controller may still put
  # base_image on the member nodes, and the node plans must stay clean.
  dynamic "base_image" {
    for_each = var.repro_ci836 && var.ci836_eve_image != "" ? [1] : []
    content {
      image_name = var.ci836_eve_image
      activate   = true
    }
  }

  # Depending on zedcloud_edgenode.EN alone is NOT enough -- that record exists
  # long before the VM finishes booting, and both validations above need a
  # live, kube-capable node.
  depends_on = [
    zedamigo_wait_until.onboarded,
    zedamigo_wait_until.kube_ready,
  ]
}

# ---------------------------------------------------------------------------
# Cluster readiness.
#
# The API returns 200 once the config is accepted and pushed; k3s/etcd form
# afterwards, on the node, over minutes.
#
#   GET /api/v1/cluster/id/{id}/status -> { id, runState, nodes: [...] }
#   readyState UNKNOWN means "in the config, has not reported Ready yet" --
#   the normal state for the first few minutes, not a failure.
#
# A zedamigo_wait_until rather than a data source because the provider has no
# client method for the cluster status endpoint.
# ---------------------------------------------------------------------------

resource "zedamigo_wait_until" "cluster_ready" {
  triggers = {
    cluster_id = zedcloud_edgenode_cluster.CLUSTER.id
  }

  timeout         = var.cluster_ready_timeout
  interval        = "30s"
  attempt_timeout = "30s"

  script = <<-EOT
    set -u

    STATUS=$(curl -s ${local.curl_ca} --max-time 25 \
      -H "Authorization: Bearer ${var.zedcloud_token}" \
      "https://${var.zedcloud_url}/api/v1/cluster/id/${zedcloud_edgenode_cluster.CLUSTER.id}/status")

    RUN=$(printf '%s' "$STATUS" | jq -r '.runState // "none"')
    TOTAL=$(printf '%s' "$STATUS" | jq -r '(.nodes // []) | length')
    READY=$(printf '%s' "$STATUS" | jq -r '
      [ (.nodes // [])[]
        | select(.readyState == "EDGE_NODE_CLUSTER_NODE_READY_STATE_READY") ] | length')

    echo "runState=$RUN ready=$READY/$TOTAL"
    printf '%s' "$STATUS" | jq -r '(.nodes // [])[] | "  \(.name) \(.readyState)"'

    [ "$READY" = "1" ] && exit 0
    exit 1
  EOT
}
