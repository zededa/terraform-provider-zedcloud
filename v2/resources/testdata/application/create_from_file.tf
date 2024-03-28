resource "zedcloud_application" "test_tf_provider" {
    name = "ubuntu-all-ip"
    title = "ubuntu-all-ip"
    description = "ubuntu-all-ip"
    manifest_file = "./testdata/application/create_manifest.json"
    user_defined_version = "1.1"
    origin_type = "ORIGIN_LOCAL"
}
