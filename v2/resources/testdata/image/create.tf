resource "zedcloud_image" "test_image" {
    name = "test_tf_provider"
    datastore_id = "aa7eb10b-7c1e-4f07-87c3-b30331e5d451"
    image_arch = "AMD64"
    image_format = "CONTAINER"
    image_rel_url = "test_url"
    image_size_bytes = 0
    image_type =  "IMAGE_TYPE_APPLICATION"
    title = "test"
    project_access_list = ["4754cd0f-82d7-4e06-a68f-ff9e23e75ccf"]
}

