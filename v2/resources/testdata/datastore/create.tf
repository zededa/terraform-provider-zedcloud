resource "zedcloud_project" "test_tf_provider" {
    # required
    name = "test_tf_provider-create_datastore"
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
    tag_level_settings {
        flow_log_transmission = "NETWORK_INSTANCE_FLOW_LOG_TRANSMISSION_UNSPECIFIED"
        interface_ordering = "INTERFACE_ORDERING_ENABLED"
    }
}

resource "zedcloud_datastore"  "test_datastore" {
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