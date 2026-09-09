# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# Consumed by the acceptance suite in later iterations, which needs the real
# node's identity injected via ZEDCLOUD_TEST_* environment variables.
output "device_id" { value = zedcloud_edgenode.EN.id }
output "device_name" { value = zedcloud_edgenode.EN.name }
output "serialno" { value = zedcloud_edgenode.EN.serialno }
output "project_id" { value = zedcloud_project.PROJECT.id }
output "model_id" { value = zedcloud_model.QEMU_VM.id }
output "network_name" { value = zedcloud_network.mgmt_dhcp.name }

# Post-mortem handles. The console logs are the ONLY place EVE's silent failures
# (TLS trust, console misconfiguration) show up, so always collect them.
output "install_console_log" { value = zedamigo_installed_edge_node.INSTALL.serial_console_log }
output "run_console_log" { value = zedamigo_edge_node.VM.serial_console_log }
output "install_success" { value = zedamigo_installed_edge_node.INSTALL.success }
output "soft_serial" { value = zedamigo_installed_edge_node.INSTALL.soft_serial }

# ssh -p <port> root@localhost on the LAB HOST (SLIRP forwards bind there).
output "eve_ssh_port" { value = zedamigo_edge_node.VM.ssh_port }

output "reserved_cpus" { value = zedamigo_host_reservation.SLOT.cpus_reserved }

output "onboard_wait" {
  value = {
    attempts = zedamigo_wait_until.onboarded.attempts
    elapsed  = zedamigo_wait_until.onboarded.elapsed
    last     = trimspace(zedamigo_wait_until.onboarded.stdout)
  }
}
