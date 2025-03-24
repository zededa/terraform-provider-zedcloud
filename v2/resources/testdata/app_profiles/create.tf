resource "zedcloud_datastore" "open_ds" {
  # required
  ds_fqdn             = "http://147.75.33.217"
  ds_path             = "images"
  ds_type             = "DATASTORE_TYPE_HTTP"
  name                = "open_ds"
  title               = "open_ds"
  description         = "open_ds"
  region              = "eu"
  # project_access_list = [zedcloud_project.test_tf_provider.id]
}

resource "zedcloud_image" "open_alpine_image" {
  depends_on = [
    zedcloud_datastore.open_ds
  ]
  name                = "openalpine"
  datastore_id        = zedcloud_datastore.open_ds.id
  image_arch          = "ARM64"
  image_format        = "CONTAINER"
  image_rel_url       = "alpine:latest"
  image_size_bytes    = 0
  image_type          = "IMAGE_TYPE_APPLICATION"
  title               = "openalpine"
}


resource "zedcloud_app_profile" "test_tf_provider" {
  depends_on = [ zedcloud_image.open_alpine_image ]
  name = "test_tf_provider"
  title = "test_tf_provider"

  app_policies {
    meta_data {
      name = "default-app-policy"
      title = "default-app-policy"
    }

    app_config {
      bundle_version = 0
      cpus = 2

      manifest_json {
        ac_kind             = "VMManifest"
        ac_version          = "1.2.0"
        name                = "test_tf_provider"
        vmmode              = "HV_HVM"
        enablevnc           = false
        app_type            = "APP_TYPE_VM"
        deployment_type     = "DEPLOYMENT_TYPE_STAND_ALONE"
        cpu_pinning_enabled = false

        images {
          imagename   = zedcloud_image.open_alpine_image.name
          imageid     = zedcloud_image.open_alpine_image.id
          imageformat = zedcloud_image.open_alpine_image.image_format
          cleartext   = false
          drvtype     = "HDD"
          ignorepurge = true
          maxsize     = 102400000
          target      = "Disk"
        }
      }
    }
  }

  network_policies {
    meta_data {
      name = "default-network-policy"
      title = "default-network-policy"
    }

    network_config {
      port = "1234"
      kind = "NETWORK_INSTANCE_KIND_LOCAL"
      type = "NETWORK_INSTANCE_DHCP_TYPE_V4"
    }
  }
}
