# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# ---------------------------------------------------------------------------
# Cluster endpoints.
#
# These are TWO DIFFERENT hostnames and mixing them up is the classic failure:
#   - zedcloud_url      -> the CONTROL/REST API the zedcloud provider calls
#   - eve_cluster_host  -> the DEVICE API, written into EVE's /config/server
#
# Verified on alpha (2026-08-19):
#   zedcontrol.alpha.zededa.net   resolves, control API, token works
#   zedcloud.alpha.zededa.net     resolves, device API, /api/v2/edgedevice/ping -> 200
#   zedcontrold.alpha.zededa.net  DOES NOT RESOLVE -- do not use
#
# Both are behind Cloudflare with a Google Trust Services cert for
# *.alpha.zededa.net, i.e. publicly trusted. That is why this config needs no
# `tls_ca` and no `additional_hosts` (unlike the local-cluster setup, where
# EVE's pre-onboarding TLS check failed silently against a private Zededa CA).
# ---------------------------------------------------------------------------

variable "zedcloud_url" {
  type        = string
  description = "Control/REST API host, e.g. zedcontrol.alpha.zededa.net"
  default     = "zedcontrol.alpha.zededa.net"
}

variable "zedcloud_token" {
  type        = string
  description = "Zedcloud API token (TF_VAR_zedcloud_token)"
  sensitive   = true
}

variable "eve_cluster_host" {
  type        = string
  description = "Device API host baked into EVE /config/server"
  default     = "zedcloud.alpha.zededa.net"
}

# ---------------------------------------------------------------------------
# Identity / isolation
# ---------------------------------------------------------------------------

# Suffixed onto every object name so concurrent runs and leftovers cannot
# collide. In CI this becomes "gh<run_id>".
variable "run_id" {
  type        = string
  description = "Unique per-run suffix for every created object"
}

# CN of EVE's default onboarding certificate. Alpha has 15 devices in
# RUN_STATE_ONLINE, so this key is expected to be authorized here.
variable "onboarding_key" {
  type        = string
  description = "Onboarding key (CN of EVE's default onboard cert)"
  default     = "5d0767ee-0547-4569-b530-387e526f8cb9"
}

variable "edge_node_ssh_pub_key" {
  type        = string
  description = "SSH public key baked into the EVE image and set as debug.enable.ssh"
  default     = ""
}

# ---------------------------------------------------------------------------
# Node shape
# ---------------------------------------------------------------------------

variable "eve_tag" {
  type        = string
  description = "lfedge/eve Docker tag. Pinned deliberately: a moving tag makes CI non-reproducible."
  default     = "16.0.1-lts-kvm-amd64"
}

variable "eve_arch" {
  type    = string
  default = "amd64"
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
  type    = number
  default = 20000
}

# Overall budget for EVE to boot, register and report in.
variable "onboard_timeout" {
  type    = string
  default = "25m"
}

# Root of the zedamigo capacity-reservation tree. The host-wide default,
# /var/lib/zedamigo/reservations, is what gives cross-user/cross-config
# coordination and is what CI should ultimately use -- but see the KNOWN ISSUE
# note on zedamigo_host_reservation in node.tf.
variable "reservations_path" {
  type        = string
  description = "Reservation tree root on the target"
  default     = "/var/lib/zedamigo/reservations"
}
