# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

output "cluster_id" { value = zedcloud_edgenode_cluster.CLUSTER.id }
output "cluster_name" { value = zedcloud_edgenode_cluster.CLUSTER.name }

output "device_id" { value = zedcloud_edgenode.EN.id }
output "device_name" { value = zedcloud_edgenode.EN.name }

# The soft serial EVE generated at install time. On macOS this is the source of
# truth for the device's identity -- the cloud record follows it, not the
# reverse.
output "soft_serial" { value = zedamigo_installed_edge_node.INSTALL.soft_serial }

output "project_id" { value = zedcloud_project.PROJECT.id }
output "model_id" { value = zedcloud_model.VFKIT_VM.id }
output "cluster_interface" { value = local.cluster_interface }

# CI-834 surface: what the controller wrote back onto the device object.
output "cluster_assignment" {
  description = "Per-node cluster state as the controller populated it"
  value = {
    cluster_interface = zedcloud_edgenode.EN.cluster_interface
    cluster = length(zedcloud_edgenode.EN.edge_node_cluster) > 0 ? {
      id             = zedcloud_edgenode.EN.edge_node_cluster[0].id
      cluster_prefix = zedcloud_edgenode.EN.edge_node_cluster[0].cluster_prefix
      is_master      = zedcloud_edgenode.EN.edge_node_cluster[0].is_master
      seed_node_id   = zedcloud_edgenode.EN.edge_node_cluster[0].seed_node_id
      seed_node_ip   = zedcloud_edgenode.EN.edge_node_cluster[0].seed_node_ip
    } : null
  }
  # The block also carries the cluster token, a shared secret.
  sensitive = true
}

# ---------------------------------------------------------------------------
# Post-mortem handles. The console logs are the ONLY place EVE's silent
# failures (TLS trust against the local CA, console misconfiguration, a
# non-kubevirt image) show up. Collect these on every failure.
# ---------------------------------------------------------------------------

output "install_console_log" { value = zedamigo_installed_edge_node.INSTALL.serial_console_log }
output "run_console_log" { value = zedamigo_edge_node.VM.serial_console_log }
output "install_success" { value = zedamigo_installed_edge_node.INSTALL.success }

# ssh -i ~/.ssh/eve_probe_ed25519 -p <port> root@localhost
#   eve exec kube kubectl get nodes
#   eve exec kube kubectl -n kubevirt get pods
#   curl -sk https://zedcloud.local.zededa.net/api/v2/edgedevice/ping
output "eve_ssh_port" { value = zedamigo_edge_node.VM.ssh_port }
output "eve_port_forwards" { value = zedamigo_edge_node.VM.nic0_port_forwards }

output "onboard_wait" {
  value = {
    attempts = zedamigo_wait_until.onboarded.attempts
    elapsed  = zedamigo_wait_until.onboarded.elapsed
    last     = trimspace(zedamigo_wait_until.onboarded.stdout)
  }
}

output "kube_ready_wait" {
  value = {
    attempts = zedamigo_wait_until.kube_ready.attempts
    elapsed  = zedamigo_wait_until.kube_ready.elapsed
    last     = trimspace(zedamigo_wait_until.kube_ready.stdout)
  }
}

output "cluster_ready_wait" {
  value = {
    attempts = zedamigo_wait_until.cluster_ready.attempts
    elapsed  = zedamigo_wait_until.cluster_ready.elapsed
    last     = trimspace(zedamigo_wait_until.cluster_ready.stdout)
  }
}
