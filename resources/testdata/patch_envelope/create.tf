data "zedcloud_project" "test_tf_provider" {
  name  = "test_tf_provider"
  type  = "TAG_TYPE_PROJECT"
  title = "test_tf_provider"
}

resource "zedcloud_image"  "test_tf_provider" {
  name = "test_tf_provider"
  title = "test_tf_provider"
  image_arch = "UNSPECIFIED"
  image_format = "RAW"
  datastore_id = "bfa12a79-7dc8-4639-8909-882e8c44eaf4"
  image_rel_url = "test_url"
  image_size_bytes = 5000
  image_type = "IMAGE_TYPE_ARTIFACT"
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
      image_id = resource.zedcloud_image.test_tf_provider.id
      image_name = resource.zedcloud_image.test_tf_provider.name
      file_name_to_use = "new_filename"
    }
  }
  project_name = data.zedcloud_project.test_tf_provider.name
  project_id = data.zedcloud_project.test_tf_provider.id
  depends_on = [
    zedcloud_image.test_tf_provider
    ]
}