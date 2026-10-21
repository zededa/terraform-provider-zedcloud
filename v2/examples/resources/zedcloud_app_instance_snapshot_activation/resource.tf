# Roll an edge app instance back to a snapshot. Apply the rollback in its own
# `terraform apply`, with no pending changes to the app instance.
#
# The resource waits until the edge node reports the snapshot as ACTIVE.
# Destroying it does not undo the rollback.
#
# The snapshot must be CREATED, and a snapshot can be rolled back to only once:
# the controller cannot send the edge node a second rollback to the same
# snapshot. To roll back again later, take a new snapshot and point
# snapshot_name at it. The provider refuses both cases before anything is sent.
#
# After the rollback, zedcloud_application_instance shows the rolled-back
# configuration as drift on the next plan, like any other change made outside
# Terraform. To keep the rollback, change the configuration to match it
# (`terraform apply -refresh-only` records it in state, but the next plan
# still reverts it while the configuration differs). Running `terraform apply`
# with the old configuration pushes it again and undoes the rollback.
resource "zedcloud_app_instance_snapshot_activation" "rollback" {
  app_instance_id = zedcloud_application_instance.app.id
  snapshot_name   = zedcloud_app_instance_snapshot.before_upgrade.name

  # Change to roll back again.
  triggers = {
    reason = "INC-1234"
  }
}

# Roll back to the snapshot taken by an update with snapshot_on_update. The edge
# node takes that snapshot only on an update that replaces the app's volumes
# (for example a purge); check it is CREATED first.
data "zedcloud_app_instance_snapshots" "app" {
  app_instance_id = zedcloud_application_instance.app.id
}

resource "zedcloud_app_instance_snapshot_activation" "undo_update" {
  app_instance_id = zedcloud_application_instance.app.id
  snapshot_name   = data.zedcloud_app_instance_snapshots.app.snapshots[0].name
}
