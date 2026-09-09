# Real-node application instance fixture. Driven by
# TestApplicationInstance_RealNode; see also .github/workflows/e2e.yml.
#
# Unlike application_instance/create.tf, this fixture does NOT fabricate an edge
# node. It targets an existing, onboarded node and deploys a real container onto
# it, so the test can assert the workload actually reaches SW_STATE_RUNNING on
# the device.
#
# The node is referenced by injected UUID rather than through the
# `zedcloud_edgenode` DATA SOURCE, because that data source is currently unusable
# for lookups: NodeDataSource() reuses zschema.Node() -- the resource schema --
# so it inherits every write-side Required field and rejects a name-only lookup
# with "The argument \"model_id\" is required" (likewise project_id, title, and
# at least one interfaces block). testhelper.RealNode resolves the name to a UUID
# against the API instead. Switch this back to the data source once its schema is
# split from the resource's.
#
# Placeholders (expanded by testhelper.MustGetTestInputWithVars):
#   __NODE_ID__    UUID of the live onboarded edge node
#   __SUFFIX__     per-run suffix; keeps names unique on a shared enterprise
#
# Everything here is modelled on a container app already running on alpha, so the
# shapes are known-good rather than guessed:
#   datastore  DATASTORE_TYPE_CONTAINERREGISTRY, ds_fqdn "docker://docker.io",
#              empty ds_path
#   image      plain `name`, pull reference in `image_rel_url`
#   manifest   ac_kind "PodManifest", app_type APP_TYPE_CONTAINER,
#              deployment_type DEPLOYMENT_TYPE_STAND_ALONE, vmmode HV_NOHYPER
#   instance   binds to the network instance by `netinstname`

resource "zedcloud_project" "real" {
  name  = "test_tf_real__SUFFIX__"
  title = "test_tf_real__SUFFIX__"
  type  = "TAG_TYPE_PROJECT"
}

# ---------------------------------------------------------------------------
# Docker Hub as a container registry. The edge node reaches it through QEMU
# SLIRP NAT, inheriting the lab host's outbound connectivity.
# ---------------------------------------------------------------------------
resource "zedcloud_datastore" "real" {
  name        = "test_tf_real_ds__SUFFIX__"
  title       = "test_tf_real_ds__SUFFIX__"
  description = "Docker Hub, for real-node acceptance testing"

  ds_type = "DATASTORE_TYPE_CONTAINERREGISTRY"
  ds_fqdn = "docker://docker.io"
  ds_path = ""

  project_access_list = [zedcloud_project.real.id]
}

# For a container-registry datastore:
#   name          a plain identifier -- the controller rejects "/" and ":" with
#                 "Name field contains invalid characters"
#   image_rel_url the actual pull reference, "<repo>/<name>:<tag>", resolved
#                 relative to the datastore's ds_fqdn
#   image_size_bytes stays 0; the device discovers the real size when it pulls
#
# Downstream references (manifest images.imagename, instance drives.imagename)
# use the image NAME, not the pull reference.
resource "zedcloud_image" "real" {
  depends_on = [zedcloud_datastore.real]

  name             = "test_tf_real_img__SUFFIX__"
  title            = "test_tf_real_img__SUFFIX__"
  datastore_id     = zedcloud_datastore.real.id
  image_rel_url    = "library/nginx:stable-alpine"
  image_arch       = "AMD64"
  image_format     = "CONTAINER"
  image_type       = "IMAGE_TYPE_APPLICATION"
  image_size_bytes = 0

  project_access_list = [zedcloud_project.real.id]
}

# ---------------------------------------------------------------------------
# A local network instance on eth0.
#
# NOTE eth0, not eth1: the VM is created with a single QEMU SLIRP NIC, so eth1
# does not exist. The rest of the suite uses eth1 against synthetic nodes, where
# nothing ever checks whether the port is real.
# ---------------------------------------------------------------------------
resource "zedcloud_network_instance" "real" {
  device_id = "__NODE_ID__"
  name      = "test_tf_real_ni__SUFFIX__"
  title     = "test_tf_real_ni__SUFFIX__"
  kind      = "NETWORK_INSTANCE_KIND_LOCAL"
  port      = "eth0"

  type           = "NETWORK_INSTANCE_DHCP_TYPE_V4"
  device_default = false
  dhcp           = false
}

resource "zedcloud_application" "real" {
  depends_on = [zedcloud_image.real]

  name                 = "test_tf_real_app__SUFFIX__"
  title                = "test_tf_real_app__SUFFIX__"
  description          = "nginx container for real-node acceptance testing"
  user_defined_version = "1.0"
  origin_type          = "ORIGIN_LOCAL"

  # No datastore_id_list: the controller rejects it with "datastore refs are only
  # supported for docker compose type app". The datastore reaches the device via
  # the image reference instead.

  # Must match the image's project scope. Leaving this unset means "all
  # projects", which the controller rejects with "app project list cannot be set
  # to all projects if all of it's images do not set their project list to all
  # projects" while the image is scoped to just this project.
  project_access_list = [zedcloud_project.real.id]

  manifest {
    ac_kind         = "PodManifest"
    ac_version      = "1.2.0"
    name            = "test_tf_real_app__SUFFIX__"
    app_type        = "APP_TYPE_CONTAINER"
    deployment_type = "DEPLOYMENT_TYPE_STAND_ALONE"

    # Containers run without a hypervisor.
    vmmode = "HV_NOHYPER"

    owner {
      user    = "Terraform acceptance tests"
      email   = "noreply@zededa.com"
      website = "www.zededa.com"
    }

    images {
      imagename   = zedcloud_image.real.name
      imageid     = zedcloud_image.real.id
      imageformat = "CONTAINER"
      maxsize     = "0"
      mountpath   = "/"
      cleartext   = false
      ignorepurge = false
      preserve    = false
      readonly    = false
    }

    # An ACL is required; a permissive one keeps the test about deployment
    # rather than about firewall semantics.
    interfaces {
      name         = "eth0"
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
# The app instance. activate = true is what makes the controller push the
# workload to the device, which is the whole point of this test.
# ---------------------------------------------------------------------------
resource "zedcloud_application_instance" "real" {
  depends_on = [
    zedcloud_application.real,
    zedcloud_network_instance.real,
    zedcloud_image.real,
  ]

  name        = "test_tf_real_appinst__SUFFIX__"
  title       = "test_tf_real_appinst__SUFFIX__"
  description = "nginx container instance for real-node acceptance testing"

  activate  = "true"
  app_id    = zedcloud_application.real.id
  device_id = "__NODE_ID__"
  app_type  = "APP_TYPE_CONTAINER"

  drives {
    imagename = zedcloud_image.real.name
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
    intfname    = "eth0"
    intforder   = 1
    netinstname = zedcloud_network_instance.real.name
    privateip   = false
    acls {
      matches {
        type  = "ip"
        value = "0.0.0.0/0"
      }
    }
  }

  # Only cpus/vnc (required), mode and the two feature flags are configurable
  # here. `memory` and `vnc_display` are Computed -- memory is derived from the
  # application manifest's `resources` block above, so setting it here fails with
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
