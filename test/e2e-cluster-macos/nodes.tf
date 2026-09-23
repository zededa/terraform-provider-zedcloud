# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# ---------------------------------------------------------------------------
# Phase 1: an empty qcow2 for the node.
#
# No zedamigo_host_reservation phase: that resource is Linux-only (it needs
# util-linux flock and a shared reservation tree). Nothing else on this Mac
# competes for the CPUs, so admission control has no job here.
# ---------------------------------------------------------------------------

resource "zedamigo_disk_image" "empty" {
  name    = "tf_encm_disk_${var.run_id}"
  size_mb = var.node_disk_mb
}

# ---------------------------------------------------------------------------
# Phase 2: the EVE-OS installer, customised for the LOCAL controller.
#
# Four attributes here exist only because the controller is local and private:
#
#   cluster            the device API host EVE writes to /config/server
#   additional_hosts   maps that host to an address the GUEST can reach --
#                      the Mac's /etc/hosts entry (127.0.0.1) is useless
#                      inside the VM
#   tls_ca             the private CA behind the ingress certificate;
#                      without it EVE's TLS check fails before onboarding and
#                      says almost nothing about why
#   object_signing_ca  the CA behind the controller's config-signing cert
#
# format = "raw" and installer_raw are the macOS path. grub_cfg MUST use
# console=hvc0: vfkit only offers virtio-serial, and a mismatch yields an
# EMPTY console log, which fails silently.
#
# Requires Docker (Desktop is running, 29.8.0) -- the installer is built by
# running the lfedge/eve container image.
# ---------------------------------------------------------------------------

resource "zedamigo_eve_installer" "eve" {
  name    = "EVE-K_${var.run_id}_${lower(var.eve_arch)}"
  format  = "raw"
  tag     = var.eve_tag
  cluster = var.eve_cluster_host

  authorized_keys = local.ssh_pub_key

  # *** THE TRAILING NEWLINE IS LOAD-BEARING. ***
  #
  # zedamigo writes this string to /config/hosts verbatim, and EVE APPENDS its
  # own entry to that file once it has a device UUID. Without a final newline
  # the two run together:
  #
  #   192.168.127.254 zedcloud.local.zededa.net127.0.0.1 9aa846b3-...
  #
  # so the name becomes "zedcloud.local.zededa.net127.0.0.1" and resolution
  # breaks. Measured 2026-09-21, and the failure mode is nasty: the node
  # REGISTERS fine (that happens before EVE appends its line) and then goes
  # silent -- adminState ADMIN_STATE_REGISTERED, runState stuck at
  # RUN_STATE_PROVISIONED forever, and curl from inside pillar returning 000.
  #
  # A heredoc rather than "...\n" so the newline cannot be lost in editing.
  additional_hosts = <<-EOH
    ${var.controller_host_ip} ${var.eve_cluster_host}
  EOH

  tls_ca            = local.tls_ca
  object_signing_ca = local.object_signing_ca

  # *** DO NOT ADD eve_nuke_disks HERE. ***
  #
  # The Linux harness passes `eve_nuke_disks=vda,sda,vdb,sdb`, which is safe
  # there because the installer is an ISO on a CD-ROM device. On macOS the
  # installer is a RAW DISK attached as the second virtio-blk device -- vfkit
  # gets `--device virtio-blk,path=<target>` then
  # `--device virtio-blk,path=<installer>` -- so vdb IS THE INSTALLER.
  #
  # Nuking it destroys the installer's own CONFIG partition mid-install. What
  # that looks like, measured 2026-09-21:
  #
  #   - the install still reports "EVE-OS installation completed" and
  #     zedamigo's `success` is true
  #   - the installer .raw provably contains the custom config
  #     (grep -a 'zedcloud.local.zededa.net' finds it)
  #   - the INSTALLED disk contains none of it (grep finds 0)
  #   - the node boots with EVE's STOCK /config: server = zedcloud.zededa.net,
  #     no authorized_keys, no hosts, no v2tlsbaseroot-certificates.pem
  #   - so it dials the public controller, never onboards, and sshd is
  #     unreachable -- with nothing in any log naming the cause
  #   - the install console log's one hint is "EVE install server : " with an
  #     empty value
  #
  # Upstream's macOS example carries no eve_nuke_disks, which is the tell.
  # It is unnecessary here anyway: zedamigo_disk_image.empty is a fresh image
  # on every run, so there are no stale partitions to clear.
  grub_cfg = <<-EOF
   set_getty
   set_global dom0_extra_args "$dom0_extra_args console=hvc0 hv_console=hvc0 dom0_console=hvc0"
   EOF

  # GUARD AGAINST A LEAKED TF_VAR_eve_cluster_host.
  #
  # This fired for real on 2026-09-21: a ~/.zshrc for the Berlin harness
  # exported zedcloud.alpha.zededa.net, the installer baked it, and the node
  # dialled the local controller's IP while presenting alpha's hostname. The
  # local ingress had no route or certificate for that name, so the node sat
  # in RUN_STATE_PROVISIONED with nothing in any log explaining it -- an hour
  # to diagnose, and a full rebuild to fix, because /config/server is baked.
  #
  # Compares everything after the first label: "local.zededa.net" from
  # zedcontrol.local.zededa.net must equal "local.zededa.net" from
  # zedcloud.local.zededa.net.
  lifecycle {
    precondition {
      condition = (
        join(".", slice(split(".", var.eve_cluster_host), 1, length(split(".", var.eve_cluster_host))))
        ==
        join(".", slice(split(".", var.zedcloud_url), 1, length(split(".", var.zedcloud_url))))
      )
      error_message = <<-EOM
        eve_cluster_host and zedcloud_url point at DIFFERENT deployments:
          zedcloud_url     = ${var.zedcloud_url}      (control API)
          eve_cluster_host = ${var.eve_cluster_host}  (device API, baked into /config/server)

        Almost always a leaked TF_VAR_eve_cluster_host from a shell set up for
        another environment. Use ./run.sh, which scrubs TF_VAR_*.
      EOM
    }
  }
}

# ---------------------------------------------------------------------------
# Phase 3: run the installer.
#
# Blocks until the install VM powers off, with NO timeout in the provider --
# if an install hangs, the apply hangs. `success` is derived by grepping the
# install console log for "EVE-OS installation completed".
#
# serial_no is a placeholder: the guest ignores it on macOS. The value that
# matters comes back out as `soft_serial` and feeds zedcloud_edgenode.EN.
# ---------------------------------------------------------------------------

resource "zedamigo_installed_edge_node" "INSTALL" {
  name            = "tf_encm_install_${var.run_id}"
  serial_no       = var.install_serial_placeholder
  installer_raw   = zedamigo_eve_installer.eve.filename
  disk_image_base = zedamigo_disk_image.empty.filename
}

# ---------------------------------------------------------------------------
# Phase 4: boot the installed disk.
#
# use_gvproxy is the default on macOS but is set explicitly because the guest's
# view of the host (var.controller_host_ip) depends on it.
#
# No TPM. Virtualization.framework provides no virtual TPM device, so
# swtpm_socket is ignored on macOS and EVE runs with no TPM-backed device
# certificate, no sealed vault and empty PCRs. Onboarding uses the soft device
# certificate instead. If something later insists on a TPM, that is a hard
# reason to move to a Linux host.
#
# depends_on zedcloud_edgenode.EN is load-bearing: the cloud record must exist
# before EVE boots and starts POSTing to /api/v2/edgedevice/register.
# ---------------------------------------------------------------------------

resource "zedamigo_edge_node" "VM" {
  name      = "tf_encm_vm_${var.run_id}"
  serial_no = zedamigo_installed_edge_node.INSTALL.serial_no

  cpus = var.node_cpus
  mem  = "${var.node_mem_gb}G"

  disk_image_base = zedamigo_installed_edge_node.INSTALL.disk_image
  ovmf_vars_src   = zedamigo_installed_edge_node.INSTALL.ovmf_vars

  use_gvproxy = true

  # Ignored on macOS (vfkit has no socket-based serial device); the console log
  # is written to a file either way. Left set so the attribute matches the
  # Linux harness.
  serial_port_server = true

  depends_on = [zedcloud_edgenode.EN]
}

# ---------------------------------------------------------------------------
# Phase 5: wait for onboarding.
#
# RUN_STATE_SUSPECT IS NOT AN ERROR. the controller's device cache sets it when
# EVE has sent an info message but no metrics message has been processed yet:
# the state becomes online only once a metrics message arrives. On a node busy pulling gigabytes of k3s/kubevirt images that can last
# well past the dormant window. An earlier version of this probe demanded
# ONLINE and burned 35 minutes on a perfectly healthy node.
#
# The cluster API gates on AdminState == DEVICE_REGISTERED and never looks at
# runState, so that is the condition that matters.
# ---------------------------------------------------------------------------

resource "zedamigo_wait_until" "onboarded" {
  triggers = {
    node_id   = zedamigo_edge_node.VM.id
    device_id = zedcloud_edgenode.EN.id
  }

  timeout         = var.onboard_timeout
  interval        = "20s"
  attempt_timeout = "30s"

  script = <<-EOT
    set -u

    STATUS=$(curl -s ${local.curl_ca} --max-time 15 \
      -H "Authorization: Bearer ${var.zedcloud_token}" \
      "https://${var.zedcloud_url}/api/v1/devices/id/${zedcloud_edgenode.EN.id}/status")

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
# Phase 6: KUBE READINESS.
#
# Being online does NOT mean Kubernetes is up -- k3s, kubevirt and longhorn
# install and start afterwards over tens of minutes, and creating the cluster
# before that finishes "fails or stalls" (upstream's words).
#
# Ask the NODE, not the controller. Polling the controller for
# OptionalCapabilities.hvTypeKubevirt looks right (it is exactly what
# the controller's live node-capability check gates on) but zedcloud only populates that field
# `if dinfo.GetOptionalCapabilities() != nil`, so "not reported yet" and "not
# capable" are indistinguishable from outside -- both are simply absent.
#
# The probe runs on THIS Mac and SSHes to the gvproxy-forwarded localhost port,
# using var.ssh_key_file, whose public half is baked into the installer.
#
# DO NOT swallow stderr. An earlier version ended the ssh call with
# `>/dev/null 2>&1`, so an auth failure and "still starting" were reported
# identically -- and it burned 25 minutes on `Permission denied (publickey)`
# while looking like a slow boot.
# ===========================================================================

resource "zedamigo_wait_until" "kube_ready" {
  triggers = {
    onboarded = zedamigo_wait_until.onboarded.id
    node_id   = zedamigo_edge_node.VM.id
  }

  timeout         = var.kube_ready_timeout
  interval        = "10s"
  attempt_timeout = "60s"

  script = <<-EOT
    set -u

    PORT=${zedamigo_edge_node.VM.ssh_port}

    SSH="ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null \
             -o ConnectTimeout=5 -o BatchMode=yes -o LogLevel=ERROR \
             -i ${local.ssh_key_path} -o IdentitiesOnly=yes \
             -p $PORT root@localhost"

    # MARKER FILES ARE VERSION-DEPENDENT. Upstream polls
    # /var/lib/all_components_initialized, which DOES NOT EXIST on EVE 17.0.0 --
    # that release uses per-component markers instead. Accept either shape.
    OUT=$($SSH 'eve exec kube sh -c "ls /var/lib/all_components_initialized || (ls /var/lib/kubevirt_initialized && ls /var/lib/multus_initialized)"' 2>&1)
    RC=$?

    if [ $RC -eq 0 ]; then
      # Markers alone are not enough: they appear while the stack is still
      # unhealthy. DiskPressure in particular leaves kubevirt cycling forever
      # and EVE then never advertises hvTypeKubevirt.
      COND=$($SSH 'eve exec kube kubectl get node -o jsonpath={.items[0].status.conditions[*].type}={.items[0].status.conditions[*].status}' 2>&1 | grep -v DEPRECATION)
      PRESSURE=$($SSH 'eve exec kube kubectl get node --no-headers' 2>&1 | grep -ci "NotReady\|SchedulingDisabled")

      if printf '%s' "$COND" | grep -qi "DiskPressure=True"; then
        echo "port $PORT: markers present but node has DISK PRESSURE."
        echo "  Longhorn needs 25% of /persist free. Raise var.node_disk_mb."
        exit 1
      fi

      if [ "$PRESSURE" != "0" ]; then
        echo "port $PORT: markers present but node not Ready yet"
        exit 1
      fi

      echo "port $PORT: kube components initialized and node Ready"
      exit 0
    fi

    case "$OUT" in
      *"Permission denied"*|*"Host key verification"*|*"Too many authentication"*)
        echo "port $PORT: SSH AUTH FAILED -- this is NOT a kube problem."
        echo "  $OUT"
        echo "  The installer bakes in ${local.ssh_key_path}.pub; the probe"
        echo "  presents ${local.ssh_key_path}. If they disagree, rebuild."
        ;;
      *"Connection refused"*|*"onnection closed"*|*"No route"*|*"timed out"*)
        echo "port $PORT: sshd not up yet (rc=$RC): $OUT"
        ;;
      *)
        echo "port $PORT: kube not ready yet (rc=$RC, eve_tag=${var.eve_tag}): $OUT"
        ;;
    esac
    exit 1
  EOT
}
