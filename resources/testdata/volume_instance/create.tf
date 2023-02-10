resource "zedcloud_edgenode" "test_tf_provider" {
		name = "test_tf_provider-create_edgenode"
		model_id = "2f716b55-2639-486c-9a2f-55a2e94146a6"
		project_id = "4754cd0f-82d7-4e06-a68f-ff9e23e75ccf"
		title = "title"
}

data "zedcloud_edgenode" "test_tf_provider" {
		name = "test_tf_provider-create_edgenode"
    depends_on = [
        zedcloud_edgenode.test_tf_provider
    ]
}

resource "zedcloud_volume_instance"  "test_volume_instance" {
    device_id = data.zedcloud_edgenode.test_tf_provider.id
    name = "test-name"
    title = "test-title"
    description = "test-description"
    type = "VOLUME_INSTANCE_TYPE_EMPTYDIR"
    multiattach = true
    cleartext = true
    accessmode = "VOLUME_INSTANCE_ACCESS_MODE_INVALID"
    size_bytes = "1024"
    implicit = false
    label = "label"
    tags = {
        "tag-key-1" = "tag-value-1"
        "tag-key-2" = "tag-value-2"
    }
}
