# A deployment with a PCR device policy.
#
# The controller versions deployments: changing this resource creates a new
# version with a new id, and every version needs its own deployment_tag (it
# must be unique within the project). Nodes run the version whose tag they
# carry, so moving nodes to the new tag rolls the change out. Older versions
# stay on the controller until the resource is destroyed.
resource "zedcloud_deployment" "fleet" {
  name           = "fleet-deployment"
  title          = "Fleet deployment"
  project_id     = zedcloud_project.fleet_deployment.id
  deployment_tag = "depl:v2" # was "depl:v1"; a change needs a new tag

  device_policies {
    meta_data {
      name  = "fleet-pcr"
      title = "Fleet PCR policy"
      policy_target_condition = {
        "depl" = "v2"
      }
    }
    policy_sub_type = "DEVICE_POLICY_TYPE_PCR"
    pcr_policy {
      pcr_indices = [0, 1, 2, 3, 4, 6, 7, 8, 9, 13, 14]
      # policy_version is managed by the controller.
    }
  }
}
