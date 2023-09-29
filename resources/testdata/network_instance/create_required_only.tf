resource "zedcloud_edgenode" "test_tf_provider" {
		name = "test_tf_provider-create_edgenode"
		model_id = "2f716b55-2639-486c-9a2f-55a2e94146a6"
		project_id = "4754cd0f-82d7-4e06-a68f-ff9e23e75ccf"
		title = "title"
}

data "zedcloud_edgenode" "test_tf_provider" {
		name = "test_tf_provider-create_edgenode"
        model_id = zedcloud_edgenode.test_tf_provider.model_id
        title = zedcloud_edgenode.test_tf_provider.title
    depends_on = [
        zedcloud_edgenode.test_tf_provider
    ]
}

resource "zedcloud_network_instance" "required_only" {
    depends_on = [
        data.zedcloud_edgenode.test_tf_provider
    ]

    name = "name"
    title = "title"
    kind = "NETWORK_INSTANCE_KIND_SWITCH"
    port = "eth1"
    device_id = data.zedcloud_edgenode.test_tf_provider.id
}
