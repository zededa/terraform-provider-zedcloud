# Copyright (c) Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

# ===========================================================================
# CI-709 — opt in with `TF_VAR_repro_ci709=true`
#
# https://zededa.atlassian.net/browse/CI-709
#
# THE BUG, as reported: deploy a network instance (and app instance) onto a
# cluster, everything applies cleanly, and then every subsequent `plan` shows a
# large spurious diff --
#
#     ~ resource "zedcloud_network_instance" "demo_net_0" {
#         - device_id = "08d09c84-..." -> null
#
# -- and a follow-up `apply` fails with HTTP 409 Conflict.
#
# ROOT CAUSE, confirmed by reading both sides:
#
# SERVER: srvs/seine/netinstproc.go:100-153. Create an NI with only
# `edge_node_cluster` and no `device_id`, and the controller resolves a
# designated node and PERSISTS the choice:
#
#     if uInput.GetDeviceId() == "" && uInput.GetEdgeNodeCluster().GetId() != "" {
#         nodeID, err := clusterproc.GetEdgeNodeForClusterObject(...)
#         uInput.DeviceId = nodeID
#         if netinst.DesignatedNodeID == "" { netinst.DesignatedNodeID = nodeID }
#     }
#     ...
#     netinst.DeviceId = dev.Id
#
# So `device_id` comes back populated on every subsequent GET.
# (appinstproc.go:2739 does the same for app instances.)
#
# PROVIDER, before the fix: `device_id` was Optional and NOT Computed, and the
# Update path PUTs the whole model with no per-field change detection. Refresh
# wrote the server's value into state, the config had none, Terraform planned to
# remove it, and the update sent it back empty.
#
# `app_type` on application_instance was the same shape and worse: it carried
# `Default: "APP_TYPE_UNSPECIFIED"`, so the config always asserted a value and
# overwrote whatever the cluster flow set -- the `APP_TYPE_VM ->
# APP_TYPE_UNSPECIFIED` line in the ticket.
#
# THE FIX: PR #234 (commit 44f74c4a, merged), six fields to `Optional +
# Computed` across network_instance, volume_instance and application_instance.
# `app_type` additionally LOST its `Default`, because SDKv2 rejects Default and
# Computed together.
#
# ---------------------------------------------------------------------------
# EXPECTED RESULT, SINCE #234: EVERYTHING HERE APPLIES AND THEN PLANS CLEAN.
#
# (The header used to read "IT SHOULD FAIL". That was correct before #234 and
# is why this file is opt-in: a permanent diff here would break the CI-834
# clean-plan assertion for an unrelated reason.)
#
#   1. `tofu apply` with TF_VAR_repro_ci709=true   -> succeeds
#   2. `tofu plan`                                 -> 0 non-no-op changes
#
# Do NOT use `tofu plan -detailed-exitcode` as the assertion: it returns 2 for
# output-only drift as well, so it reports 2 on a clean plan of this config.
# Count non-`no-op` entries in `resource_changes` instead.
#
# A clean plan can also mean the objects never got a server-assigned device_id,
# which would make this test vacuous. Check the raw API before celebrating:
#
#   NI=$(tofu output -json ci709_network_instance | jq -r .id)
#   curl -s -H "Authorization: Bearer $TF_VAR_zedcloud_token" \
#     "https://$TF_VAR_zedcloud_url/api/v1/netinsts/id/$NI" \
#     | jq -c '{deviceId, clusterId, edgeNodeCluster}'
#
# The Go-level equivalents live in v2/resources/edgenode_cluster_test.go
# (`TestClusterScopedInstances_RealNode` and the schema/diff layers beside it),
# which assert the same thing against the live objects without a Terraform run.
#
# WHY ALL THREE RESOURCES ARE HERE
#
# #234 changed six fields across three resources, but CI covers none of this
# shape: no fixture anywhere creates a cluster-scoped instance
# (`grep -l edge_node_cluster v2/resources/testdata/*/*.tf`), `test/e2e` forms
# no cluster, and every app-instance fixture sets `app_type` explicitly -- so
# the removed Default is invisible to it too. Hence a network instance, a volume
# instance AND an application instance, each supplying only
# `edge_node_cluster`, and the app instance additionally omitting `app_type`.
# ===========================================================================

variable "repro_ci709" {
  type        = bool
  description = <<-EOT
    Create cluster-scoped network, volume and application instances to exercise
    CI-709.

    Off by default. Before PR #234 these produced a permanent diff, which would
    break the CI-834 clean-plan assertion for an unrelated reason. Since #234
    they should plan clean, but they still cost an image push and a few minutes,
    so they stay opt-in.

    Requires a formed cluster: implies node_count >= 1 and a completed
    cluster_ready.
  EOT
  default     = false
}

locals {
  ci709 = var.repro_ci709 ? 1 : 0

  # eth0 for the app-facing network instance, on BOTH topologies -- not
  # local.cluster_interface.
  #
  # Two reasons. eth0 is the QEMU SLIRP management NIC, so it is the only port
  # with outbound connectivity; bridges A/B/C are host-only, so a container on
  # eth1/2/3 could never pull an image. And putting an application network
  # instance on the CLUSTER interface invites clusterPrefixOverlap (the
  # controller rejects a cluster_prefix overlapping any NI subnet on a member
  # node), which is a different failure than the one under test.
  #
  # An NI on the management port is what test/e2e does, against a model that
  # likewise declares eth0 as ADAPTER_USAGE_MANAGEMENT.
  ci709_port = "eth0"
}

# ---------------------------------------------------------------------------
# The network instance. This is the resource named in the ticket.
# ---------------------------------------------------------------------------

resource "zedcloud_network_instance" "CI709" {
  count = local.ci709

  name  = "tf_enc_ci709_ni_${var.run_id}"
  title = "tf_enc_ci709_ni_${var.run_id}"

  kind = "NETWORK_INSTANCE_KIND_LOCAL"
  type = "NETWORK_INSTANCE_DHCP_TYPE_V4"
  port = local.ci709_port

  device_default = false
  dhcp           = false

  # *** THE POINT OF THE TEST: device_id is deliberately NOT set. ***
  #
  # The create-time guard in v2/resources/network_instance.go:44-49 requires
  # either device_id or edge_node_cluster, so supplying only the cluster is the
  # supported cluster-scoped form -- not a contrivance. The controller then
  # picks the designated node itself, and that is the value the provider used
  # to try to erase. All three resources below share the identical guard.
  #
  # designated_node_id is also left unset, so the controller chooses
  # (GetEdgeNodeForClusterObject). Setting it explicitly is the other half of
  # the API surface and worth a second test.
  edge_node_cluster {
    id = zedcloud_edgenode_cluster.CLUSTER.id
  }

  # The cluster must exist AND have converged: validateNodesForCluster rejects
  # objects on a cluster whose nodes are not ready, and an instance on a
  # half-formed cluster is a different failure mode than the one under test.
  depends_on = [
    zedcloud_edgenode_cluster.CLUSTER,
    zedamigo_wait_until.cluster_ready,
  ]
}

# ---------------------------------------------------------------------------
# The volume instance. `device_id` only -- volume_instance has no app_type.
# ---------------------------------------------------------------------------

resource "zedcloud_volume_instance" "CI709" {
  count = local.ci709

  name        = "tf_enc_ci709_vol_${var.run_id}"
  title       = "tf_enc_ci709_vol_${var.run_id}"
  description = "CI-709: cluster-scoped volume instance, device_id left to the controller"

  type       = "VOLUME_INSTANCE_TYPE_EMPTYDIR"
  size_bytes = "1073741824"
  accessmode = "VOLUME_INSTANCE_ACCESS_MODE_READWRITE"
  cleartext  = true
  implicit   = false

  # Again: no device_id.
  edge_node_cluster {
    id = zedcloud_edgenode_cluster.CLUSTER.id
  }

  depends_on = [
    zedcloud_edgenode_cluster.CLUSTER,
    zedamigo_wait_until.cluster_ready,
  ]
}

# ---------------------------------------------------------------------------
# An application to instantiate. Modelled on
# v2/resources/testdata/application_instance/create_real_node.tf, which is
# known-good against a real node, so the shapes here are copied rather than
# guessed.
# ---------------------------------------------------------------------------

resource "zedcloud_datastore" "CI709" {
  count = local.ci709

  name        = "tf_enc_ci709_ds_${var.run_id}"
  title       = "tf_enc_ci709_ds_${var.run_id}"
  description = "Docker Hub, for the CI-709 cluster-scoped app instance"

  ds_type = "DATASTORE_TYPE_CONTAINERREGISTRY"
  ds_fqdn = "docker://docker.io"
  ds_path = ""

  project_access_list = [zedcloud_project.PROJECT.id]
}

resource "zedcloud_image" "CI709" {
  count      = local.ci709
  depends_on = [zedcloud_datastore.CI709]

  # A plain identifier: the controller rejects "/" and ":" in `name` with
  # "Name field contains invalid characters". The pull reference goes in
  # image_rel_url, and downstream references use the NAME.
  name             = "tf_enc_ci709_img_${var.run_id}"
  title            = "tf_enc_ci709_img_${var.run_id}"
  datastore_id     = zedcloud_datastore.CI709[0].id
  image_rel_url    = "library/nginx:stable-alpine"
  image_arch       = upper(var.eve_arch)
  image_format     = "CONTAINER"
  image_type       = "IMAGE_TYPE_APPLICATION"
  image_size_bytes = 0

  project_access_list = [zedcloud_project.PROJECT.id]
}

resource "zedcloud_application" "CI709" {
  count      = local.ci709
  depends_on = [zedcloud_image.CI709]

  name                 = "tf_enc_ci709_app_${var.run_id}"
  title                = "tf_enc_ci709_app_${var.run_id}"
  description          = "nginx container for the CI-709 cluster-scoped app instance"
  user_defined_version = "1.0"
  origin_type          = "ORIGIN_LOCAL"

  # Must match the image's project scope. Unset means "all projects", which the
  # controller rejects while the image is scoped to one.
  project_access_list = [zedcloud_project.PROJECT.id]

  manifest {
    ac_kind         = "PodManifest"
    ac_version      = "1.2.0"
    name            = "tf_enc_ci709_app_${var.run_id}"
    app_type        = "APP_TYPE_CONTAINER"
    deployment_type = "DEPLOYMENT_TYPE_STAND_ALONE"
    vmmode          = "HV_NOHYPER"

    owner {
      user    = "Terraform e2e-cluster"
      email   = "noreply@zededa.com"
      website = "www.zededa.com"
    }

    images {
      imagename   = zedcloud_image.CI709[0].name
      imageid     = zedcloud_image.CI709[0].id
      imageformat = "CONTAINER"
      maxsize     = "0"
      mountpath   = "/"
      cleartext   = false
      ignorepurge = false
      preserve    = false
      readonly    = false
    }

    interfaces {
      name         = local.ci709_port
      directattach = false
      acls {
        matches {
          type  = "ip"
          value = "0.0.0.0/0"
        }
      }
    }

    desc {
      category     = "APP_CATEGORY_OTHERS"
      app_category = "APP_CATEGORY_OTHERS"
    }

    resources {
      name  = "resourceType"
      value = "Tiny"
    }
    resources {
      name  = "cpus"
      value = "1"
    }
    resources {
      name  = "memory"
      value = "524288.00"
    }

    configuration {
      custom_config {
        add = false
      }
    }
  }
}

# ---------------------------------------------------------------------------
# The application instance -- the resource that carries BOTH halves of #234.
# ---------------------------------------------------------------------------

resource "zedcloud_application_instance" "CI709" {
  count = local.ci709

  depends_on = [
    zedcloud_application.CI709,
    zedcloud_network_instance.CI709,
    zedcloud_image.CI709,
    zedamigo_wait_until.cluster_ready,
  ]

  name        = "tf_enc_ci709_appinst_${var.run_id}"
  title       = "tf_enc_ci709_appinst_${var.run_id}"
  description = "CI-709: cluster-scoped app instance, device_id and app_type left to the controller"

  app_id = zedcloud_application.CI709[0].id

  # activate = false ON PURPOSE.
  #
  # CI-709 is entirely config-plane: the controller assigns device_id at create
  # regardless (appinstproc.go:2739), so activating adds a multi-minute image
  # pull and a possible device-side error state without strengthening the
  # assertion. Proving a workload actually RUNS is
  # TestApplicationInstance_RealNode's job, on a standalone node.
  activate = "false"

  # *** BOTH POINTS OF THIS RESOURCE. ***
  #
  #   device_id  NOT set -- the controller resolves the designated node
  #   app_type   NOT set -- #234 removed its `Default: "APP_TYPE_UNSPECIFIED"`
  #                         and made it Computed. Every fixture in v2/ sets
  #                         app_type, so nothing else anywhere exercises a
  #                         config that omits it. If the controller rejects an
  #                         empty app_type on create, or defaults it to
  #                         something the provider then fights over, this is
  #                         where it shows up.
  edge_node_cluster {
    id = zedcloud_edgenode_cluster.CLUSTER.id
  }

  drives {
    imagename = zedcloud_image.CI709[0].name
    maxsize   = "0"
    mountpath = "/"
    preserve  = false
    readonly  = false
    drvtype   = "UNSPECIFIED"
    target    = "Disk"
  }

  # intfname / netinstname / privateip are all Required by the app-instance
  # interface schema.
  interfaces {
    intfname    = local.ci709_port
    intforder   = 1
    netinstname = zedcloud_network_instance.CI709[0].name
    privateip   = false
    acls {
      matches {
        type  = "ip"
        value = "0.0.0.0/0"
      }
    }
  }

  # `memory` and `vnc_display` are Computed -- memory comes from the
  # application manifest's `resources` block, so setting it here fails with
  # "Value for unconfigurable attribute".
  vminfo {
    cpus = 1
    vnc  = false
    mode = "HV_NOHYPER"
  }

  remote_console    = false
  is_secret_updated = false

  logs {
    access = false
  }

  manifest_info {
    transition_action = "INSTANCE_TA_NONE"
  }
}

# ===========================================================================
# MEASURED 2026-09-21 on the enc3a 3-node cluster. Two findings.
#
# 1. The controller spreads cluster-scoped objects across the cluster.
#    One apply, three different nodes:
#
#      network instance      -> 109a4ea2  (node 3)
#      volume instance       -> 78f386c4  (node 2)
#      application instance  -> bf3a4abb  (node 1)
#
#    So `device_id` is genuinely resolved per object, not a constant.
#
# 2. *** device_id IS NOT STABLE. THE CONTROLLER RE-HOMES OBJECTS. ***
#
#    Minutes after the apply, with nobody touching anything, the network
#    instance had MOVED:
#
#      state, written at apply : device_id = 109a4ea2   (node 3)
#      API, four minutes later : deviceId  = bf3a4abb   (node 1)
#
#    bf3a4abb is the application instance's node. Creating the app instance
#    made the controller move its network instance to co-locate them -- an app
#    instance and the NI its interface binds to must be on the same device.
#
#    This is the strongest argument for the #234 fix, and it is easy to miss.
#    `Computed` is not merely ergonomic here: without it, this re-assignment
#    produces a spurious diff for a user who has changed NOTHING, at a moment
#    the user cannot predict, and the follow-up apply then fights the
#    controller for ownership of a field the controller owns. With it, the
#    follow-up plan was clean -- 0 non-no-op resource changes.
#
# 3. `app_type` came back "APP_TYPE_CONTAINER" although the configuration
#    never set it: the controller derives it from the application manifest.
#    Pre-#234 the provider's `Default: "APP_TYPE_UNSPECIFIED"` would have
#    asserted UNSPECIFIED against that on every plan -- the ticket's
#    `APP_TYPE_VM -> APP_TYPE_UNSPECIFIED` line, with CONTAINER in place of VM.
#
# 4. `cluster_id` stayed empty on both the NI and the app instance; the cluster
#    is carried in `edgeNodeCluster.id` instead. So of the six fields #234
#    changed, four are exercised with real server-assigned values here and the
#    two `cluster_id`s are not populated by this flow.
# ===========================================================================

# ---------------------------------------------------------------------------
# Outputs: the server-assigned fields, which are the whole CI-709 surface.
#
# Read these rather than `tofu state show` -- the on-disk state for these
# resources is written at create time, and a later refresh only lives in the
# plan unless an apply persists it. Note finding 2 above: the output and the
# live API can legitimately disagree, and that is not drift in the Terraform
# sense.
# ---------------------------------------------------------------------------

output "ci709_network_instance" {
  description = "Server-assigned fields on the cluster-scoped NI"
  value = var.repro_ci709 ? {
    id         = zedcloud_network_instance.CI709[0].id
    name       = zedcloud_network_instance.CI709[0].name
    device_id  = zedcloud_network_instance.CI709[0].device_id
    cluster_id = zedcloud_network_instance.CI709[0].cluster_id
  } : null
}

output "ci709_volume_instance" {
  description = "Server-assigned fields on the cluster-scoped volume instance"
  value = var.repro_ci709 ? {
    id        = zedcloud_volume_instance.CI709[0].id
    name      = zedcloud_volume_instance.CI709[0].name
    device_id = zedcloud_volume_instance.CI709[0].device_id
  } : null
}

output "ci709_application_instance" {
  description = "Server-assigned fields on the cluster-scoped app instance -- app_type is the second half of #234"
  value = var.repro_ci709 ? {
    id         = zedcloud_application_instance.CI709[0].id
    name       = zedcloud_application_instance.CI709[0].name
    device_id  = zedcloud_application_instance.CI709[0].device_id
    cluster_id = zedcloud_application_instance.CI709[0].cluster_id
    app_type   = zedcloud_application_instance.CI709[0].app_type
  } : null
}
