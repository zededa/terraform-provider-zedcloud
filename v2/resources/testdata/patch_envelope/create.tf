resource "zedcloud_project" "test_tf_provider" {
  # required
  name = "test_tf_provider-create_patch_envelope"
  title = "title"

  # optional
  type = "TAG_TYPE_PROJECT"
  attestation_policy {
    # required
    title = "title"
    type = "POLICY_TYPE_ATTESTATION"

    attestation_policy {
      # required
      type = "ATTEST_POLICY_TYPE_ACCEPT"
    }
  }
}


resource "zedcloud_datastore"  "test_tf_provider" {
  depends_on = [
    zedcloud_project.test_tf_provider
  ]
  # required
  ds_fqdn = "my-datastore.my-company.com"
  ds_path = "download/AMD64"
  ds_type = "DATASTORE_TYPE_AZUREBLOB"
  name = "test"
  title = "title"
  description = "description"
  region = "eu"
  project_access_list = [zedcloud_project.test_tf_provider.id]
}

resource "zedcloud_image"  "test_tf_provider" {
  depends_on = [
    zedcloud_datastore.test_tf_provider
  ]
  name = "test_tf_provider"
  title = "test_tf_provider"
  image_arch = "UNSPECIFIED"
  image_format = "RAW"
  datastore_id = zedcloud_datastore.test_tf_provider.id
  image_rel_url = "test_url"
  image_size_bytes = 5000
  image_type = "IMAGE_TYPE_ARTIFACT"
  project_access_list = [zedcloud_project.test_tf_provider.id]
}

resource "zedcloud_patch_envelope" "test_tf_provider" {
  name   = "test_tf_provider"
  action = "PATCH_ENVELOPE_ACTION_STORE"
  title  = "test_tf_provider"
  description = "test tf provider"
  user_defined_version = "1.0"

  artifacts {
    format = "OpaqueObjectCategoryInline"
    base64_artifact {
      base64_data      = "YXJ0aWZhY3QgZGF0YQ=="
      file_name_to_use = "testfile"
    }
  }
  artifacts {
    format = "OpaqueObjectCategoryExternalBinary"
    binary_artifact {
      image_id = zedcloud_image.test_tf_provider.id
      image_name = zedcloud_image.test_tf_provider.name
      file_name_to_use = "new_filename"
    }
  }
  project_name = zedcloud_project.test_tf_provider.name
  project_id = zedcloud_project.test_tf_provider.id
  depends_on = [
    zedcloud_image.test_tf_provider,
    zedcloud_project.test_tf_provider
  ]
}