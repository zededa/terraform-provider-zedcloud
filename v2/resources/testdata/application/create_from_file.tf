resource "zedcloud_application" "test_tf_provider" {
    name = "test_tf_provider-ubuntu-all-ip"
    title = "test_tf_provider-ubuntu-all-ip"
    description = "test_tf_provider-ubuntu-all-ip"
    manifest_file = "./testdata/application/create_manifest.json"
    user_defined_version = "1.1"
    origin_type = "ORIGIN_LOCAL"
}
