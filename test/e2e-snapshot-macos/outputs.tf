# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

output "device_id" { value = zedcloud_edgenode.EN.id }
output "device_name" { value = zedcloud_edgenode.EN.name }
output "soft_serial" { value = zedamigo_installed_edge_node.INSTALL.soft_serial }
output "project_id" { value = zedcloud_project.PROJECT.id }

output "install_console_log" { value = zedamigo_installed_edge_node.INSTALL.serial_console_log }
output "run_console_log" { value = zedamigo_edge_node.VM.serial_console_log }
output "install_success" { value = zedamigo_installed_edge_node.INSTALL.success }
output "eve_ssh_port" { value = zedamigo_edge_node.VM.ssh_port }

output "onboard_wait" {
  value = {
    attempts = zedamigo_wait_until.onboarded.attempts
    elapsed  = zedamigo_wait_until.onboarded.elapsed
    last     = trimspace(zedamigo_wait_until.onboarded.stdout)
  }
}

output "app_instance_id" {
  value = try(zedcloud_application_instance.VM[0].id, null)
}

output "app_instance_snapshot_settings" {
  value = try({
    allow_snapshots    = zedcloud_application_instance.VM[0].allow_snapshots
    max_snapshots      = zedcloud_application_instance.VM[0].max_snapshots
    snapshot_on_update = zedcloud_application_instance.VM[0].snapshot_on_update
    description        = zedcloud_application_instance.VM[0].description
  }, null)
}

output "snapshot" {
  value = try({
    id    = zedcloud_app_instance_snapshot.S[0].id
    name  = zedcloud_app_instance_snapshot.S[0].name
    state = zedcloud_app_instance_snapshot.S[0].state
  }, null)
}

output "activation" {
  value = try({
    id           = zedcloud_app_instance_snapshot_activation.R[0].id
    state        = zedcloud_app_instance_snapshot_activation.R[0].state
    activated_at = zedcloud_app_instance_snapshot_activation.R[0].activated_at
  }, null)
}

output "snapshots" {
  value = try(data.zedcloud_app_instance_snapshots.ALL[0].snapshots, null)
}
