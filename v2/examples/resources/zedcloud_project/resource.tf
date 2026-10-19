# A project whose devices seal their TPM volume key to a fixed set of PCRs.
#
# PCR 0 is required and PCR 5 (volatile) is not allowed; indices range 0-15.
# The recommended set is [0, 1, 2, 3, 4, 6, 7, 8, 9, 13, 14]. A set that
# includes PCRs which change on firmware or bootloader updates can leave a node
# unable to unseal its vault after such an update.
#
# Node-level policies (zedcloud_edgenode.pcr_policy) override this one. The
# controller cannot remove a project PCR policy once it is set: deleting the
# block leaves the policy in place.
resource "zedcloud_project" "fleet" {
  name  = "fleet-a"
  title = "Fleet A"
  type  = "TAG_TYPE_PROJECT"

  pcr_policy {
    title = "fleet-a-pcr"
    type  = "POLICY_TYPE_PCR"

    pcr_policy {
      pcr_indices = [0, 1, 2, 3, 4, 6, 7, 8, 9, 13, 14]
      # policy_version is managed by the provider: 1 on create, +1 whenever
      # pcr_indices change.
    }
  }
}
