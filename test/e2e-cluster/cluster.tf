# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# ===========================================================================
# THE EDGE-NODE CLUSTER
#
# Everything above this file exists to satisfy the gates this one resource
# imposes. Checklist, all enforced server-side:
#
#   the controller's cluster logic
#     a required node count of three; 1 or >=3 nodes, NEVER 2 (no quorum)
#     the controller's master-node check: >=3 nodes -> exactly 3 of type ..._SERVER
#     the controller's prefix-overlap check: cluster_prefix must not overlap any network
#       instance subnet on the member nodes
#     the controller's prefix assignment: each node gets a distinct address out of
#       cluster_prefix; a /28 gives ~13 usable node prefixes
#     a seed node is elected and seed_node_id / seed_node_ip recorded
#     cluster token is 32 random bytes, generated server-side
#
#   the controller's cluster-membership checks  (each is a 400)
#     every node id resolves
#     node not already in a cluster, node.ClusterID empty
#     node.AdminState == DEVICE_REGISTERED   <- must have ONBOARDED, not just
#                                               ADMIN_STATE_ACTIVE
#     at most one 'tie-breaker'-tagged node
#     the controller's cluster-interface resolution resolves cluster_interface against the node's
#       interfaces / bond_adapters / vlan_adapters
#     that interface's intf_usage != ADAPTER_USAGE_UNSPECIFIED
#     no app instance with an auto-deployment policy on any node
#     NO ACTIVE APP INSTANCE on any node
#
#   the controller's live node-capability check
#     LIVE DeviceStatusReq to all nodes; every one must report
#     OptionalCapabilities.hvTypeKubevirt, or:
#       "node %s does not support kubevirt"
#     and if a node is offline:
#       "timed out after %s querying cluster node status"
#
#   the controller's shared-label check
#     shared labels on the cluster interface must match across all nodes
# ===========================================================================

resource "zedcloud_edgenode_cluster" "CLUSTER" {
  name        = "tf_enc_cluster_${var.run_id}"
  title       = "tf_enc_cluster_${var.run_id}"
  description = "UE-156 e2e: ${var.node_count}-node edge-node cluster on virtual EVE nodes"
  project_id  = zedcloud_project.PROJECT.id

  # Overlay prefix for intra-cluster traffic. Must not overlap the bridge
  # subnets or any NI subnet on the member nodes.
  #
  # null by default, which lets the provider's own default (10.244.244.2/28)
  # apply -- the upstream reference config does not set this at all.
  cluster_prefix = var.cluster_prefix

  # One block per node.
  #
  # node_type is explicit even though the provider defaults it: for a 3-node
  # cluster the controller's master-node check requires EXACTLY three SERVER nodes, so making
  # it implicit here hides a hard requirement.
  #
  # resource_labels is left unset -- see the shared-labels note above.
  #
  # cluster_prefix inside a nodes block is Computed (the controller assigns a
  # per-node address); do not set it.
  dynamic "nodes" {
    for_each = zedcloud_edgenode.EN

    content {
      id                = nodes.value.id
      cluster_interface = local.cluster_interface
      node_type         = "EDGE_NODE_CLUSTER_NODE_TYPE_SERVER"
    }
  }

  # Load-bearing, three times over:
  #
  #  1. the controller's cluster-membership checks requires every node to be DEVICE_REGISTERED,
  #     which only happens once EVE has onboarded. Depending on
  #     zedcloud_edgenode.EN alone is NOT enough -- those records exist long
  #     before the VMs finish booting.
  #
  #  2. the controller's live node-capability check issues a LIVE status query to every node and
  #     fails the whole create if any one is unreachable, so all must be up
  #     simultaneously -- and it requires each to report
  #     OptionalCapabilities.hvTypeKubevirt.
  #
  #  3. That capability only appears once EVE has brought up its Kubernetes
  #     stack, which lags onboarding by many minutes. Creating the cluster
  #     before then "fails or stalls" (upstream). kube_ready is the barrier
  #     that actually establishes this -- see the long comment in nodes.tf for
  #     why polling the controller for the capability is the wrong test.
  depends_on = [
    zedamigo_wait_until.onboarded,
    zedamigo_wait_until.kube_ready,
  ]
}

# ---------------------------------------------------------------------------
# Cluster readiness barrier.
#
# The API returns 200 on create once the config is accepted and pushed to the
# nodes; the actual k3s/etcd cluster forms afterwards, on the nodes, over
# minutes. Anything that inspects the cluster (or deploys onto it) has to wait
# for this, not for the create.
#
#   GET /api/v1/cluster/id/{id}/status -> EdgeNodeClusterStatus
#     { id, runState, nodes: [ { id, name, readyState } ] }
#   readyState in EDGE_NODE_CLUSTER_NODE_READY_STATE_{
#     UNSPECIFIED, READY, NOT_READY, UNKNOWN }
#
# UNKNOWN means "in the cluster config but has not reported a Ready condition
# yet" -- i.e. the normal state for the first few minutes.
#
# This is a zedamigo_wait_until probe rather than a provider data source
# because the provider has NO client method for the cluster status endpoint
# (only Create/Read/Update/Delete/GetByName are generated). The Go-side
# equivalent for acceptance tests is a WaitForClusterNodesReady helper in
# v2/testing/realnode.go using the existing apiGet.
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
    PATH=/run/current-system/sw/bin:/usr/bin:/bin

    STATUS=$(curl -s --max-time 25 \
      -H "Authorization: Bearer ${var.zedcloud_token}" \
      "https://${var.zedcloud_url}/api/v1/cluster/id/${zedcloud_edgenode_cluster.CLUSTER.id}/status")

    RUN=$(printf '%s' "$STATUS" | jq -r '.runState // "none"')
    TOTAL=$(printf '%s' "$STATUS" | jq -r '(.nodes // []) | length')
    READY=$(printf '%s' "$STATUS" | jq -r '
      [ (.nodes // [])[]
        | select(.readyState == "EDGE_NODE_CLUSTER_NODE_READY_STATE_READY") ] | length')

    echo "runState=$RUN ready=$READY/$TOTAL"
    printf '%s' "$STATUS" | jq -r '(.nodes // [])[] | "  \(.name) \(.readyState)"'

    [ "$READY" = "${var.node_count}" ] && exit 0
    exit 1
  EOT
}
