# Override the project's PCR policy on one node.
#
# The block is managed only once it is set: a policy set with
# `zcli edge-node pcr-policy set` is kept while the block is absent, and
# removing the block after it was set resets the node to the project (or EVE)
# defaults, like `zcli edge-node pcr-policy reset`.
#
# For a node in a deployment project, a deployment PCR device policy that
# matches the node takes precedence, and the project policy is not used.
resource "zedcloud_edgenode" "node17" {
  name       = "node-17"
  title      = "node-17"
  model_id   = zedcloud_model.m.id
  project_id = zedcloud_project.fleet.id
  # ...

  pcr_policy {
    pcr_indices = [0, 2, 3, 4, 6, 7, 8, 9, 13, 14]
    # policy_version is managed by the provider: it continues from the node's
    # current version (including one set with zcli) and increases only when
    # pcr_indices change.
  }
}
