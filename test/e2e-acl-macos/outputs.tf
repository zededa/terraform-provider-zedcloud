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

output "app_instance_ids" {
  value = { for k, v in zedcloud_application_instance.VM : k => v.id }
}
