data "zedcloud_edgenode" "complete" {
		name = "test_tf_provider-create_edgenode"
}

resource "zedcloud_volume_instance"  "test_volume_instance" {
    device_id = data.zedcloud_edgenode.complete.id
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
