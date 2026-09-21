# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# Consumed by the acceptance suite, which needs the real cluster's identity
# injected via ZEDCLOUD_TEST_* environment variables.

output "cluster_id" { value = zedcloud_edgenode_cluster.CLUSTER.id }
output "cluster_name" { value = zedcloud_edgenode_cluster.CLUSTER.name }

output "device_ids" { value = [for n in zedcloud_edgenode.EN : n.id] }
output "device_names" { value = [for n in zedcloud_edgenode.EN : n.name] }
output "serialnos" { value = [for n in zedcloud_edgenode.EN : n.serialno] }

# Comma-joined, for the __NODE_IDS__ fixture token.
output "device_ids_csv" { value = join(",", [for n in zedcloud_edgenode.EN : n.id]) }

output "project_id" { value = zedcloud_project.PROJECT.id }
output "model_id" { value = zedcloud_model.QEMU_VM.id }
output "mgmt_network_name" { value = zedcloud_network.mgmt_dhcp.name }
output "cluster_network_name" {
  value = local.second_nic ? zedcloud_network.cluster_dhcp[0].name : null
}
output "cluster_interface" { value = local.cluster_interface }
output "node_count" { value = var.node_count }

# ---------------------------------------------------------------------------
# What the controller assigned. Reading these back is the cheapest check that
# the cluster really formed rather than merely being accepted -- seed_node_ip
# lands inside cluster_prefix, and is_master is true on the seed.
#
# NOTE: the provider's SetClusterResourceData DISCARDS seed_node_id,
# seed_node_ip and admin_state from the EdgeNodeClusterConfigSummary response
# (they are not in the resource schema at all), so they cannot be read off the
# cluster resource. They ARE present on each device's computed
# edge_node_cluster block, which is where these read from.
# ---------------------------------------------------------------------------

output "cluster_assignment" {
  description = "Per-node cluster state as the controller populated it (CI-834 surface)"
  value = [
    for n in zedcloud_edgenode.EN : {
      name              = n.name
      cluster_interface = n.cluster_interface
      cluster = length(n.edge_node_cluster) > 0 ? {
        id             = n.edge_node_cluster[0].id
        cluster_prefix = n.edge_node_cluster[0].cluster_prefix
        is_master      = n.edge_node_cluster[0].is_master
        seed_node_id   = n.edge_node_cluster[0].seed_node_id
        seed_node_ip   = n.edge_node_cluster[0].seed_node_ip
      } : null
    }
  ]
  # token is in this block too and is a shared secret -- do not output it.
  sensitive = true
}

# ---------------------------------------------------------------------------
# Post-mortem handles.
#
# The console logs are the ONLY place EVE's silent failures (TLS trust, console
# misconfiguration, a kvm-flavour image that never reports kubevirt) show up.
# ALWAYS collect these, for every node, on failure.
# ---------------------------------------------------------------------------

output "install_console_logs" {
  value = [for i in zedamigo_installed_edge_node.INSTALL : i.serial_console_log]
}

output "run_console_logs" {
  value = [for v in zedamigo_edge_node.VM : v.serial_console_log]
}

output "install_success" {
  value = [for i in zedamigo_installed_edge_node.INSTALL : i.success]
}

# ssh -p <port> root@localhost ON THE TARGET HOST (SLIRP forwards bind there).
#
# This is the main debugging handle for a kubevirt node, and it is what the
# kube-readiness barrier uses. Useful once in:
#   eve exec kube ls -l /var/lib/all_components_initialized
#   eve exec kube kubectl get nodes
#   eve exec kube kubectl -n kubevirt get pods
output "eve_ssh_ports" {
  value = [for v in zedamigo_edge_node.VM : v.ssh_port]
}

output "kube_ready_waits" {
  value = [
    for w in zedamigo_wait_until.kube_ready : {
      attempts = w.attempts
      elapsed  = w.elapsed
      last     = trimspace(w.stdout)
    }
  ]
}

output "cluster_bridges" {
  value = local.second_nic ? {
    a = zedamigo_bridge.A[0].name
    b = zedamigo_bridge.B[0].name
    c = zedamigo_bridge.C[0].name
  } : null
}

output "cluster_taps" {
  value = {
    a = [for t in zedamigo_tap.A : t.name]
    b = [for t in zedamigo_tap.B : t.name]
    c = [for t in zedamigo_tap.C : t.name]
  }
}

# NOTE: `cpus_reserved` is NOT a number -- it is the reserved CPU *set* (e.g.
# a pinning list), which is why an earlier `sum()` over it failed with
# "argument must be list, set, or tuple of number values". `cpus_reserved_count`
# is the numeric one. Both are emitted raw rather than summed, so a type
# surprise in either cannot break `tofu output` again.
output "reserved_cpus" {
  value = [for s in zedamigo_host_reservation.SLOT : s.cpus_reserved]
}

output "reserved_cpu_counts" {
  value = [for s in zedamigo_host_reservation.SLOT : s.cpus_reserved_count]
}

output "onboard_waits" {
  value = [
    for w in zedamigo_wait_until.onboarded : {
      attempts = w.attempts
      elapsed  = w.elapsed
      last     = trimspace(w.stdout)
    }
  ]
}

output "cluster_ready_wait" {
  value = {
    attempts = zedamigo_wait_until.cluster_ready.attempts
    elapsed  = zedamigo_wait_until.cluster_ready.elapsed
    last     = trimspace(zedamigo_wait_until.cluster_ready.stdout)
  }
}
