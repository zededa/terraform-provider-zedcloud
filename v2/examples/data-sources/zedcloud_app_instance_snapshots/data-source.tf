# List the snapshots of an edge app instance, immediate and app-update ones.
data "zedcloud_app_instance_snapshots" "app" {
  app_instance_id = zedcloud_application_instance.app.id
}

output "snapshot_names" {
  value = data.zedcloud_app_instance_snapshots.app.snapshots[*].name
}
