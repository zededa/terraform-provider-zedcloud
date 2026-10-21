# Take an immediate snapshot of an edge app instance, for example before a
# risky change. Change a trigger value to take a new one.
#
# The controller keeps one snapshot per app instance: a new snapshot (from any
# resource, the UI, zcli, or an update with snapshot_on_update) replaces the
# previous one, and Terraform then drops the replaced resource from state.
# Use one zedcloud_app_instance_snapshot per app instance, and don't combine
# it with snapshot_on_update on the same zedcloud_application_instance: each
# would delete the other's snapshot.
#
# Before taking the snapshot, the provider checks that the edge node's API
# capability level is at least API_CAPABILITY_VOLUME_SNAPSHOTS_IMMEDIATE. The
# edge node restarts the app to take the snapshot. Not supported for app
# instances deployed to an edge node cluster.
resource "zedcloud_app_instance_snapshot" "before_upgrade" {
  app_instance_id = zedcloud_application_instance.app.id

  triggers = {
    app_version = "1.4.0"
  }

  timeouts {
    create = "20m"
  }
}
