# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# ---------------------------------------------------------------------------
# Controller endpoints.
#
# The local deployment mirrors alpha's split: zedcontrol.* is the control REST
# API, zedcloud.* is the EVE device API. Verified 2026-09-21:
#
#   https://zedcontrol.local.zededa.net/api/v1/enterprises    -> 200 (with a
#                                                                local token)
#   https://zedcloud.local.zededa.net/api/v1/version          -> 200
#   https://zedcloud.local.zededa.net/api/v2/edgedevice/ping  -> 200
#   https://zedcloud.local.zededa.net/api/v1/edgedevice/ping  -> 404  (v2 only)
#
# *** THESE TWO MUST BE THE SAME DEPLOYMENT. ***
#
# A ~/.zshrc set up for the Berlin harness exports
# TF_VAR_eve_cluster_host=zedcloud.alpha.zededa.net, which silently overrode
# the default below on the first build here: the node got
# /config/server = zedcloud.alpha.zededa.net and an /etc/hosts entry pointing
# that name at the LOCAL controller, so it dialled the right address with the
# wrong SNI, the local ingress had no route or certificate for it, and the node
# sat in RUN_STATE_PROVISIONED saying nothing. Terraform emits no diagnostic
# for an env var that merely supplies a variable.
#
# Hence run.sh (which scrubs TF_VAR_*) and the precondition in nodes.tf.
# ---------------------------------------------------------------------------

variable "zedcloud_url" {
  type        = string
  description = "Control/REST API host used by the zedcloud provider"
  default     = "zedcontrol.local.zededa.net"
}

variable "zedcloud_token" {
  type        = string
  description = <<-EOT
    Zedcloud API token for the LOCAL controller (TF_VAR_zedcloud_token).

    A token minted against alpha will NOT work here, and the failure is
    specific enough to recognise:

      HTTP 401  error[0].details = "Session Cache miss"

    The local controller keeps sessions in memory, so redeploying it (or
    restarting the API pods, which ArgoCD does routinely) invalidates every
    previously issued token. Re-login to the local UI and re-export.
  EOT
  sensitive   = true
}

variable "eve_cluster_host" {
  type        = string
  description = "Device API host baked into EVE /config/server"
  default     = "zedcloud.local.zededa.net"
}

variable "controller_host_ip" {
  type        = string
  description = <<-EOT
    The address, AS SEEN FROM INSIDE THE EVE GUEST, at which this Mac's
    controller is reachable. Written into the node's /etc/hosts via the
    installer's additional_hosts.

    Why this is needed at all: zedcloud.local.zededa.net resolves to 127.0.0.1
    through /etc/hosts on the Mac, and `minikube tunnel` binds the istio
    ingress LoadBalancer on loopback only. A guest cannot reach the host's
    127.0.0.1, and it cannot use the host's /etc/hosts either.

    192.168.127.254 is gvisor-tap-vsock's (gvproxy's) host address on its
    default 192.168.127.0/24 network -- the gateway is .1 and the host end is
    .254. Traffic to it is dialled by the gvproxy process running on this Mac,
    so it lands on the Mac's own loopback and reaches the tunnel without
    rebinding anything.

    UNVERIFIED as of 2026-09-21. If onboarding never starts, SSH into the node
    (see README) and check:

      curl -sk https://zedcloud.local.zededa.net/api/v2/edgedevice/ping

    If that fails, try 192.168.127.1 and the Mac's LAN address in turn, set
    this accordingly, and REBUILD -- additional_hosts is baked at installer
    build time, so changing it replaces the installer and the node.
  EOT
  default     = "192.168.127.254"
}

# ---------------------------------------------------------------------------
# Trust. This is the thing that silently killed the earlier local-controller
# attempt recorded in docs/design/ue-156-3-node-cluster-e2e.md: EVE's
# pre-onboarding TLS check failed against the private Zededa CA and reported
# nothing useful.
#
# Verified 2026-09-21: the local ingress serves
#   leaf  CN=zedcloud.local.zededa.net
#   int   CN=Zededa Inc. Intermediat CA1
#   root  CN=Zededa Inc. Root CA         (not sent, held locally)
# and `openssl s_client -CAfile zededa-root-ca.pem` verifies OK.
# ---------------------------------------------------------------------------

variable "tls_ca_file" {
  type        = string
  description = <<-EOT
    PEM of the CA that signed the controller's TLS server certificate. Becomes
    /config/v2tlsbaseroot-certificates.pem on the node.
    See https://github.com/lf-edge/eve/blob/master/docs/REGISTRATION.md
  EOT
  default     = "~/zedcloud-local/certs/zededa-root-ca.pem"
}

variable "object_signing_ca_file" {
  type        = string
  description = <<-EOT
    PEM of the CA that signed the controller's OBJECT SIGNING certificate --
    a different certificate from the TLS one, used to verify the signed config
    EVE fetches. Becomes /config/root-certificate.pem on the node.

    Defaults to the same Zededa root as tls_ca_file, which is the expected
    arrangement for this deployment but has NOT been confirmed by parsing
    /api/v2/edgedevice/certs (the response is a protobuf blob and a naive DER
    scan found no certificates in it). If onboarding gets as far as fetching
    config and then fails on signature verification, this is the variable to
    look at first.
  EOT
  default     = "~/zedcloud-local/certs/zededa-root-ca.pem"
}

# ---------------------------------------------------------------------------
# Identity
# ---------------------------------------------------------------------------

variable "run_id" {
  type        = string
  description = <<-EOT
    Unique per-run suffix for every created object.

    No IFNAMSIZ constraint here (this config creates no host interfaces), but
    keep it short anyway so object names stay readable.

    A FRESH run_id IS REQUIRED FOR A FRESH BUILD. Reinstalling EVE against an
    already-registered device object leaves the node stuck forever on
    "Waiting for DeviceName from controller...".
  EOT
}

variable "onboarding_key" {
  type        = string
  description = <<-EOT
    CN of EVE's default onboarding certificate. The controller's join key is
    literally "<onboarding_key>:<serialno>" (the controller's join-key derivation).

    This is the stock lf-edge value. Whether the LOCAL controller has it
    registered for your enterprise is unverified -- if registration is refused,
    check the onboarding keys in the local UI.
  EOT
  default     = "5d0767ee-0547-4569-b530-387e526f8cb9"
}

# ---------------------------------------------------------------------------
# Node shape
# ---------------------------------------------------------------------------

variable "eve_tag" {
  type        = string
  description = <<-EOT
    lfedge/eve tag. MUST be the KVM flavour (`-kvm-`): the controller rejects
    snapshots on devices that report hvTypeKubevirt, which the `-k-` flavour
    does. The node must also report API_CAPABILITY_VOLUME_SNAPSHOTS_IMMEDIATE.
  EOT
  default     = "17.0.0-lts-kvm-arm64"
}

variable "eve_arch" {
  type    = string
  default = "arm64"
}

variable "node_cpus" {
  type    = number
  default = 4
}

variable "node_mem_gb" {
  type    = number
  default = 8
}

variable "node_disk_mb" {
  type        = number
  description = "Node disk size. Sparse, so the size costs nothing until used."
  default     = 40960
}

# ---------------------------------------------------------------------------
# SSH
# ---------------------------------------------------------------------------

variable "ssh_key_file" {
  type        = string
  description = <<-EOT
    Private key on THIS Mac used by the kube-readiness probe to SSH into the
    node. Its ".pub" sibling is baked into the installer as authorized_keys.

    It must be the INSTALLER key, not the `debug.enable.ssh` config item: on a
    freshly installed node EVE's sshd uses the baked-in authorized_keys and
    ignores the config item. Getting this wrong costs a full probe budget of
    `Permission denied (publickey)` that reads exactly like a slow boot.

      ssh-keygen -t ed25519 -N "" -f ~/.ssh/eve_probe_ed25519
  EOT
  default     = "~/.ssh/eve_probe_ed25519"
}

# ---------------------------------------------------------------------------
# Timeouts
# ---------------------------------------------------------------------------

variable "onboard_timeout" {
  type    = string
  default = "35m"
}

variable "install_serial_placeholder" {
  type        = string
  description = <<-EOT
    serial_no passed to zedamigo_installed_edge_node. Ignored by the guest on
    macOS -- Virtualization.framework cannot set an SMBIOS serial -- but the
    attribute is required. Upstream's macOS example passes "1234567890".
  EOT
  default     = "1234567890"
}

# ---------------------------------------------------------------------------
# UE-203 test stages, see README.md.
# ---------------------------------------------------------------------------

variable "stage" {
  type        = number
  description = <<-EOT
    How far app.tf goes:
      0  node only
      1  + instance X: edits the bundle's 9022->22 port map to var.x_ssh_port->22
      2  + instance Y on the same node with both port maps moved (7022, 7080)
  EOT
  default     = 0
}

variable "x_ssh_port" {
  type        = string
  description = "Node port of instance X's ssh port map. The bundle's is 9022; change it to test an ACL update."
  default     = "8022"
}

variable "vm_image" {
  type = object({
    fqdn   = string
    path   = string
    file   = string
    sha256 = string
    size   = string
  })
  description = "A small arm64 qcow2 VM image reachable from the node over the internet."
  default = {
    fqdn   = "https://download.cirros-cloud.net"
    path   = "0.6.2"
    file   = "cirros-0.6.2-aarch64-disk.img"
    sha256 = "6a2222bd7ba89384d35378b5065bd5e838857f25db08c7abbff99b6d9bb48986"
    size   = "24969728"
  }
}

