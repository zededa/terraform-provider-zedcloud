# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# ---------------------------------------------------------------------------
# Per-node VM lifecycle. Same phases as test/e2e/node.tf, multiplied by
# var.node_count, plus the extra NICs from net_cluster.tf and a kube-readiness
# barrier that test/e2e does not need.
#
# Shaped after the upstream reference config
# (andrei-zededa/terraform-provider-zedamigo, examples/other/edge_node_cluster,
# edge_nodes.tf) -- the only known-working 3-node topology.
# ---------------------------------------------------------------------------

locals {
  # Block-device mode when node_disk_devs is supplied; qcow2 files otherwise.
  use_disk_devs = length(var.node_disk_devs) > 0

  # Per-node vCPU counts. var.node_cpus may be shorter than node_count, in which
  # case the last entry repeats. Default [6, 6, 4] = 16, which is exactly what
  # h-m-dl20 has.
  node_cpus = [
    for i in range(var.node_count) :
    i < length(var.node_cpus) ? var.node_cpus[i] : var.node_cpus[length(var.node_cpus) - 1]
  ]
}

# ---------------------------------------------------------------------------
# Phase 0: cooperative capacity reservation.
#
# Claims are atomic via flock against the shared tree and visible to any other
# config on the host. ADMISSION CONTROL, NOT A QUEUE: if capacity is
# unavailable the apply FAILS rather than waiting.
#
# CAPACITY: with var.cpu_pinning (the default, matching upstream), the sum of
# local.node_cpus is a hard constraint. h-m-dl20 has 16 CPUs, hence the default
# [6, 6, 4] rather than upstream's 6/6/6 -- see var.node_cpus.
#
# DISKS: one reserve per node, spread across the three physical drives
# (vg_sdb / vg_sdc / vg_sdd) -- see var.node_disk_devs.
# ---------------------------------------------------------------------------

resource "zedamigo_host_reservation" "SLOT" {
  count = var.node_count

  path = var.reservations_path
  cpus = local.node_cpus[count.index]
  mem  = var.node_mem_gb

  # Reserve a host block device per node when one was supplied.
  devs = local.use_disk_devs ? [var.node_disk_devs[count.index]] : []
}

# ---------------------------------------------------------------------------
# Phase 1: an empty qcow2 per node -- only in file mode.
#
# Upstream uses raw host block devices instead, because Longhorn and etcd on a
# qcow2 FILE are slow enough to matter for a kubevirt node. h-m-dl20 has no LVM
# volume groups today, so file mode is the default here; supply
# var.node_disk_devs once storage exists.
# ---------------------------------------------------------------------------

resource "zedamigo_disk_image" "empty" {
  count = local.use_disk_devs ? 0 : var.node_count

  name    = "tf_enc_disk_${var.run_id}_${count.index + 1}"
  size_mb = var.node_disk_mb
}

# ---------------------------------------------------------------------------
# Phase 2: build ONE EVE-OS installer, shared by all nodes.
#
# The installer only carries /config (server endpoint + authorized_keys), which
# is identical across nodes -- node identity comes from the SMBIOS serial at
# boot, not from the image. So one installer, built once, installed N times.
#
# grub_cfg MUST match serial_type. The default serial_type is `virtio`, whose
# console is hvc0. Getting this wrong yields an EMPTY console log, so
# installed_edge_node.success never flips true -- and it fails SILENTLY.
#
# `eve_nuke_disks` comes from upstream and matters when disks are REUSED (block
# devices, or a re-run over an existing image): without it EVE can find stale
# partitions and refuse to install cleanly.
# ---------------------------------------------------------------------------

resource "zedamigo_eve_installer" "eve" {
  name            = "EVE-K_${var.run_id}_${lower(var.eve_arch)}"
  tag             = var.eve_tag
  cluster         = var.eve_cluster_host
  authorized_keys = var.edge_node_ssh_pub_key

  grub_cfg = <<-EOF
   set_getty
   set_global dom0_extra_args "$dom0_extra_args console=hvc0 hv_console=hvc0 dom0_console=hvc0 eve_nuke_disks=vda,sda,vdb,sdb"
   EOF
}

# ---------------------------------------------------------------------------
# Phase 3: run the installer, once per node.
#
# Boots a throwaway VM and BLOCKS until it powers off, with NO timeout in the
# provider -- if an install hangs, the apply hangs. `success` is derived by
# grepping the install console log for "EVE-OS installation completed".
#
# These run in parallel up to Terraform's -parallelism (default 10), ON TOP OF
# the reservations for the real VMs, so true peak load exceeds what the
# reservation implies. Pass -parallelism=1 if the host thrashes.
# ---------------------------------------------------------------------------

resource "zedamigo_installed_edge_node" "INSTALL" {
  count = var.node_count

  name          = "tf_enc_install_${var.run_id}_${count.index + 1}"
  serial_no     = zedcloud_edgenode.EN[count.index].serialno
  installer_iso = zedamigo_eve_installer.eve.filename

  # File mode: install onto the qcow2 image.
  disk_image_base = local.use_disk_devs ? null : zedamigo_disk_image.empty[count.index].filename

  # Block-device mode: install straight onto the reserved device.
  dynamic "disk" {
    for_each = local.use_disk_devs ? [1] : []

    content {
      type     = "device"
      source   = zedamigo_host_reservation.SLOT[count.index].devs_reserved[0]
      format   = "raw"
      drive_if = "virtio"
      options  = ["cache=none", "aio=io_uring", "discard=unmap", "detect-zeroes=unmap"]
    }
  }
}

# ---------------------------------------------------------------------------
# Phase 4: boot the installed disks. EVE connects out and onboards.
#
# eth0 is QEMU SLIRP (the nic0 default): 10.0.2.15 behind NAT, inheriting the
# host's outbound connectivity, plus a host-forwarded SSH port. No root needed
# for that path.
#
# eth1/eth2/eth3 are taps on bridges A/B/C -- MULTI-NODE ONLY. Attached via
# extra_qemu_args using the single-argument `-nic tap,...` form upstream uses.
# MACs are derived from run_id and node index so they are stable across
# re-applies (EVE keys some state off them) and unique across concurrent runs.
#
# `depends_on` is load-bearing: the cloud record must exist before EVE boots and
# starts POSTing to /api/v2/edgedevice/register, otherwise the first attempts
# fail against a device the controller does not know yet.
# ---------------------------------------------------------------------------

resource "zedamigo_edge_node" "VM" {
  count = var.node_count

  name      = "tf_enc_vm_${var.run_id}_${count.index + 1}"
  serial_no = zedamigo_installed_edge_node.INSTALL[count.index].serial_no

  cpus     = zedamigo_host_reservation.SLOT[count.index].cpus_reserved_count
  cpu_pins = var.cpu_pinning ? zedamigo_host_reservation.SLOT[count.index].cpus_reserved : null
  mem      = "${zedamigo_host_reservation.SLOT[count.index].mem_reserved_total_gb}G"

  ovmf_vars_src      = zedamigo_installed_edge_node.INSTALL[count.index].ovmf_vars
  serial_port_server = true

  # File mode: boot the installed qcow2.
  disk_image_base = local.use_disk_devs ? null : zedamigo_installed_edge_node.INSTALL[count.index].disk_image

  # Block-device mode: boot the same device the installer wrote to.
  dynamic "disk" {
    for_each = local.use_disk_devs ? [1] : []

    content {
      type     = "device"
      source   = zedamigo_installed_edge_node.INSTALL[count.index].disk[0].source
      format   = "raw"
      drive_if = "virtio"
      options  = ["cache=none", "aio=io_uring", "discard=unmap", "detect-zeroes=unmap"]
    }
  }

  extra_qemu_args = local.second_nic ? [
    "-nic", "tap,id=vmnet1,ifname=${zedamigo_tap.A[count.index].name},script=no,downscript=no,model=virtio,mac=1E:94:C2:${substr(md5(var.run_id), 0, 2)}:A${count.index + 1}:01",
    "-nic", "tap,id=vmnet2,ifname=${zedamigo_tap.B[count.index].name},script=no,downscript=no,model=virtio,mac=1E:94:C2:${substr(md5(var.run_id), 0, 2)}:B${count.index + 1}:02",
    "-nic", "tap,id=vmnet3,ifname=${zedamigo_tap.C[count.index].name},script=no,downscript=no,model=virtio,mac=1E:94:C2:${substr(md5(var.run_id), 0, 2)}:C${count.index + 1}:03",
  ] : []

  depends_on = [
    zedcloud_edgenode.EN,
    zedamigo_host_reservation.SLOT,
    zedamigo_tap.A,
    zedamigo_tap.B,
    zedamigo_tap.C,
    zedamigo_dhcp_server.A,
    zedamigo_dhcp_server.B,
  ]
}

# ---------------------------------------------------------------------------
# Phase 5: wait for each node to onboard.
#
# zedamigo_edge_node returns as soon as QEMU is detached -- EVE's boot,
# register -> uuid -> certs -> config handshake all happen AFTER apply would
# otherwise finish. Without this barrier the apply "succeeds" against nodes that
# never onboarded, and the cluster create then fails with
# "node %s is not registered" (validateNodesForCluster requires
# AdminState == DEVICE_REGISTERED, which only onboarding sets).
# ---------------------------------------------------------------------------

resource "zedamigo_wait_until" "onboarded" {
  count = var.node_count

  triggers = {
    node_id   = zedamigo_edge_node.VM[count.index].id
    device_id = zedcloud_edgenode.EN[count.index].id
  }

  timeout         = var.onboard_timeout
  interval        = "20s"
  attempt_timeout = "30s"

  # WHAT THIS BARRIER SHOULD ACTUALLY REQUIRE.
  #
  # An earlier version demanded RUN_STATE_ONLINE and timed out after 35 minutes
  # on a node sitting in RUN_STATE_SUSPECT -- which is NOT an error state.
  # srvs/ganges/devcache.go:254 sets SUSPECT when EVE has sent an info message
  # but no metrics message has been processed yet: "keep opstate as suspect here
  # and make it online only upon arrival of metrics msg". On a node busy pulling
  # gigabytes of k3s/kubevirt images on first boot, metrics can lag well past
  # the dormant window (DefaultMetricsIntervalSec * 3).
  #
  # And the cluster API does not care: validateNodesForCluster gates on
  # AdminState == DEVICE_REGISTERED, never on runState. So the meaningful
  # condition is "registered and talking", with kube_ready doing the real
  # gating afterwards.
  #
  # ONLINE   -> reporting normally
  # SUSPECT  -> info received, metrics pending (accepted)
  # anything else with adminState REGISTERED -> keep waiting
  script = <<-EOT
    set -u
    PATH=/run/current-system/sw/bin:/usr/bin:/bin

    STATUS=$(curl -s --max-time 15 \
      -H "Authorization: Bearer ${var.zedcloud_token}" \
      "https://${var.zedcloud_url}/api/v1/devices/id/${zedcloud_edgenode.EN[count.index].id}/status")

    RUN=$(printf '%s' "$STATUS" | jq -r '.runState // "none"')
    ADM=$(printf '%s' "$STATUS" | jq -r '.adminState // "none"')
    echo "runState=$RUN adminState=$ADM"

    if [ "$ADM" = "ADMIN_STATE_REGISTERED" ]; then
      case "$RUN" in
        RUN_STATE_ONLINE|RUN_STATE_SUSPECT) exit 0 ;;
      esac
    fi
    exit 1
  EOT
}

# ===========================================================================
# Phase 6: KUBE READINESS — the barrier this config previously got wrong.
#
# Being RUN_STATE_ONLINE does NOT mean the Kubernetes stack is up. k3s,
# kubevirt and longhorn install and start afterwards, over many minutes, and
# creating the cluster before that finishes "fails or stalls" (upstream's
# words).
#
# WHAT NOT TO DO — and what this config did on its first attempt:
# poll the CONTROLLER for OptionalCapabilities.hvTypeKubevirt. That looks right
# (srvs/seine/devproc.go:validateClusterNodesInfo gates cluster creation on
# exactly that field) but it is a poor barrier:
#
#   - zedcloud only populates it `if dinfo.GetOptionalCapabilities() != nil`
#     (srvs/yamuna/devproc.go:deviceFillInfo), so "not reported yet" and "not
#     capable" are indistinguishable from outside -- both are simply absent.
#   - observed on a node running the WRONG (old-naming) image: online and
#     registered for 20+ minutes with optionalCapabilities: null forever, and
#     zero k3s/kubelet/longhorn/etcd in the console log.
#
# WHAT UPSTREAM DOES, and what this now does: ask the NODE directly, over SSH,
# whether EVE finished initialising its kube components.
#
# This needs working SSH INTO EACH NODE, FROM THE TARGET. Two variables govern
# that, and they are separate on purpose:
#
#   var.node_ssh_authorized_key   -> the `debug.enable.ssh` config item, which
#                                    EVE applies LIVE. Change it and a running
#                                    node accepts the new key without a rebuild.
#   var.node_ssh_private_key_file -> the key the probe presents, ON THE TARGET.
#
# Agent forwarding is NOT a workable substitute: with `forward_agent = true`
# set, a check on h-m-dl20 found SSH_AUTH_SOCK unset and no keys on the host, so
# the probe failed `Permission denied (publickey)` for its entire budget.
#
# The probe runs ON the zedamigo target, so root@localhost:<ssh_port> resolves
# to the QEMU host-forwarded port on the machine that actually runs the VMs --
# which is what keeps this working when `target` points at a remote host.
# ===========================================================================

resource "zedamigo_wait_until" "kube_ready" {
  count = var.node_count

  triggers = {
    onboarded = zedamigo_wait_until.onboarded[count.index].id
    node_id   = zedamigo_edge_node.VM[count.index].id
  }

  timeout  = var.kube_ready_timeout
  interval = "10s"
  # Backstop only: ssh carries its own ConnectTimeout, so one attempt is
  # bounded well below this.
  attempt_timeout = "60s"

  # DO NOT swallow stderr here. An earlier version ended the ssh invocation with
  # `>/dev/null 2>&1`, so an SSH AUTH failure and "kube is still starting" were
  # reported identically as "not ready yet" -- and it burned a full 25-minute
  # budget on `Permission denied (publickey)` while looking like a slow boot.
  # Distinguishing the two is the whole value of this probe.
  script = <<-EOT
    set -u
    PATH=/run/current-system/sw/bin:/usr/bin:/bin

    PORT=${zedamigo_edge_node.VM[count.index].ssh_port}
    N=${count.index + 1}
    KEY_OPT=""
    if [ -n "${var.node_ssh_private_key_file == null ? "" : var.node_ssh_private_key_file}" ]; then
      KEY_OPT="-i ${var.node_ssh_private_key_file == null ? "" : var.node_ssh_private_key_file} -o IdentitiesOnly=yes"
    fi

    SSH="ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null \
             -o ConnectTimeout=5 -o BatchMode=yes -o LogLevel=ERROR $KEY_OPT \
             -p $PORT root@localhost"

    # MARKER FILES ARE VERSION-DEPENDENT. Upstream polls
    # /var/lib/all_components_initialized, which DOES NOT EXIST on EVE 17.0.0 --
    # that release uses per-component markers instead:
    #   kubevirt_initialized  multus_initialized  debuguser-initialized
    # So accept either shape.
    OUT=$($SSH 'eve exec kube sh -c "ls /var/lib/all_components_initialized || (ls /var/lib/kubevirt_initialized && ls /var/lib/multus_initialized)"' 2>&1)
    RC=$?

    if [ $RC -eq 0 ]; then
      # Markers alone are not enough: they appear while the stack is still
      # unhealthy. DiskPressure in particular leaves kubevirt cycling forever
      # and EVE then never advertises hvTypeKubevirt -- observed with a 30 GB
      # qcow2 giving /persist 5.1 GB. Require the node to be actually healthy.
      COND=$($SSH 'eve exec kube kubectl get node -o jsonpath={.items[0].status.conditions[*].type}={.items[0].status.conditions[*].status}' 2>&1 | grep -v DEPRECATION)
      PRESSURE=$($SSH 'eve exec kube kubectl get node --no-headers' 2>&1 | grep -ci "NotReady\|SchedulingDisabled")

      if printf '%s' "$COND" | grep -qi "DiskPressure=True"; then
        echo "node $N (port $PORT): markers present but node has DISK PRESSURE."
        echo "  Longhorn needs 25% of /persist free. A 30 GB disk yields ~5 GB"
        echo "  /persist and fails. Use var.node_disk_devs (100 GB reserves) or"
        echo "  raise var.node_disk_mb."
        exit 1
      fi

      if [ "$PRESSURE" != "0" ]; then
        echo "node $N (port $PORT): markers present but node not Ready yet"
        exit 1
      fi

      echo "node $N (port $PORT): kube components initialized and node Ready"
      exit 0
    fi

    # Separate "cannot get in" from "in, but kube not ready".
    case "$OUT" in
      *"Permission denied"*|*"Host key verification"*|*"Too many authentication"*)
        echo "node $N (port $PORT): SSH AUTH FAILED -- this is NOT a kube problem."
        echo "  $OUT"
        echo "  Fix var.node_ssh_authorized_key / var.node_ssh_private_key_file;"
        echo "  the probe runs on the zedamigo target and needs a key THERE."
        ;;
      *"Connection refused"*|*"onnection closed"*|*"No route"*|*"timed out"*)
        echo "node $N (port $PORT): sshd not up yet (rc=$RC): $OUT"
        ;;
      *)
        echo "node $N (port $PORT): kube not ready yet (rc=$RC, eve_tag=${var.eve_tag}): $OUT"
        ;;
    esac
    exit 1
  EOT
}
