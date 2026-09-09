# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# ---------------------------------------------------------------------------
# Cooperative capacity reservation.
#
# The Berlin lab host is shared, so claim the CPUs and RAM this node will use
# before creating it. The reservation tree is pre-created by the operator at
# /var/lib/zedamigo/reservations (16 CPU slots, 58 GB slots on h-m-dl20) and is
# group-writable by `zedamigo_reservations`, which the login user belongs to --
# so no sudo is needed. Claims are atomic via flock and are visible to any other
# Terraform configuration on the host, including a human's.
#
# NOTE: this is admission control, not a queue. If capacity is unavailable the
# apply FAILS rather than waiting.
# ---------------------------------------------------------------------------

# Claims capacity from the shared tree at /var/lib/zedamigo/reservations, so
# CI, andrei's manual runs and anyone else's cannot oversubscribe the host.
#
# HISTORY: this used to need a per-user tree. The shared tree's `.lock` was
# created on first use with the creating process's umask -- 0644, owned by
# whoever ran first -- so every other member of zedamigo_reservations failed
# with "Permission denied ... is /var/lib/zedamigo/reservations writable?",
# and cooperative reservation worked for exactly one user. Now fixed
# declaratively in the lab repo (hosts/h-m-dl20/github-runners.nix declares the
# lock 0664 root:zedamigo_reservations via tmpfiles), so the default shared
# path is correct again.
#
# Override var.reservations_path only if testing against a host without that
# fix; a per-user tree keeps this resource exercised but gives no cross-user
# coordination.
resource "zedamigo_host_reservation" "SLOT" {
  path = var.reservations_path

  cpus = var.node_cpus
  mem  = var.node_mem_gb
}

# ---------------------------------------------------------------------------
# Phase 1: an empty qcow2 that EVE will be installed onto.
# ---------------------------------------------------------------------------

resource "zedamigo_disk_image" "empty" {
  name    = "tf_e2e_disk_${var.run_id}"
  size_mb = var.node_disk_mb
}

# ---------------------------------------------------------------------------
# Phase 2: build a custom EVE-OS installer.
#
# This runs `docker run --network none ... docker.io/lfedge/eve:<tag>` on the
# target, which writes the /config directory into the image -- `cluster` becomes
# /config/server, i.e. the DEVICE API endpoint.
#
# grub_cfg MUST match serial_type. The default serial_type is `virtio`, whose
# console is hvc0. Getting this wrong yields an EMPTY console log, so
# installed_edge_node.success never flips true -- and it fails silently.
# ---------------------------------------------------------------------------

resource "zedamigo_eve_installer" "eve" {
  name            = "tf_e2e_installer_${var.run_id}"
  tag             = var.eve_tag
  cluster         = var.eve_cluster_host
  authorized_keys = var.edge_node_ssh_pub_key

  grub_cfg = <<-EOF
   set_getty
   set_global dom0_extra_args "$dom0_extra_args console=hvc0 hv_console=hvc0 dom0_console=hvc0"
   EOF
}

# ---------------------------------------------------------------------------
# Phase 3: run the installer.
#
# This boots a throwaway VM and blocks until it powers off. There is NO timeout
# in the provider -- if the install hangs, the apply hangs. `success` is derived
# by grepping the install console log for "EVE-OS installation completed".
# ---------------------------------------------------------------------------

resource "zedamigo_installed_edge_node" "INSTALL" {
  name            = "tf_e2e_install_${var.run_id}"
  serial_no       = zedcloud_edgenode.EN.serialno
  installer_iso   = zedamigo_eve_installer.eve.filename
  disk_image_base = zedamigo_disk_image.empty.filename
}

# ---------------------------------------------------------------------------
# Phase 4: boot the installed disk. EVE connects out and onboards.
#
# Uplink is QEMU SLIRP (the nic0 default): the guest gets 10.0.2.15 behind NAT
# and inherits the host's outbound connectivity. No bridge, tap, netns or root
# required -- and the lab host reaches the device API (verified: /api/v2/
# edgedevice/ping -> 200).
#
# `depends_on` is load-bearing: the cloud record must exist before EVE boots and
# starts POSTing to /api/v2/edgedevice/register, otherwise the first attempts
# fail against a device the controller does not know yet.
# ---------------------------------------------------------------------------

resource "zedamigo_edge_node" "VM" {
  name               = "tf_e2e_vm_${var.run_id}"
  cpus               = var.node_cpus
  mem                = "${var.node_mem_gb}G"
  serial_no          = zedamigo_installed_edge_node.INSTALL.serial_no
  disk_image_base    = zedamigo_installed_edge_node.INSTALL.disk_image
  ovmf_vars_src      = zedamigo_installed_edge_node.INSTALL.ovmf_vars
  serial_port_server = true

  depends_on = [
    zedcloud_edgenode.EN,
    zedamigo_host_reservation.SLOT,
  ]
}

# ---------------------------------------------------------------------------
# Phase 5: wait for the node to actually come up.
#
# zedamigo_edge_node returns as soon as QEMU is detached -- EVE's boot,
# register -> uuid -> certs -> config handshake all happen AFTER apply would
# otherwise finish. Without this barrier the apply "succeeds" against a node
# that never onboarded.
#
# The probe runs ON the target (the lab host), which is what makes this work
# unchanged when the provider later points at a remote target.
#
# Deliberately a single-shot probe with no inner loop: `interval` and `timeout`
# own the retrying.
# ---------------------------------------------------------------------------

resource "zedamigo_wait_until" "onboarded" {
  triggers = {
    node_id   = zedamigo_edge_node.VM.id
    device_id = zedcloud_edgenode.EN.id
  }

  timeout         = var.onboard_timeout
  interval        = "20s"
  attempt_timeout = "30s"

  # Succeeds once the controller reports the device ONLINE. The progression is
  # UNPROVISIONED -> PROVISIONED (registered, not yet reporting) -> ONLINE.
  script = <<-EOT
    set -u
    PATH=/run/current-system/sw/bin:/usr/bin:/bin

    STATUS=$(curl -s --max-time 15 \
      -H "Authorization: Bearer ${var.zedcloud_token}" \
      "https://${var.zedcloud_url}/api/v1/devices/id/${zedcloud_edgenode.EN.id}/status")

    RUN=$(printf '%s' "$STATUS" | jq -r '.runState // "none"')
    echo "runState=$RUN"

    if [ "$RUN" = "RUN_STATE_ONLINE" ]; then
      exit 0
    fi

    exit 1
  EOT
}
