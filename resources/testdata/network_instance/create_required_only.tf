resource "zedcloud_edgenode" "required_only" {
		name = "required_only"
		model_id = "2f716b55-2639-486c-9a2f-55a2e94146a6"
		project_id = "4754cd0f-82d7-4e06-a68f-ff9e23e75ccf"
		title = "required_only-title"
}

data "zedcloud_edgenode" "required_only" {
		name = "required_only"
		model_id = "2f716b55-2639-486c-9a2f-55a2e94146a6"
		project_id = "4754cd0f-82d7-4e06-a68f-ff9e23e75ccf"
		title = "required_only-title"
}

resource "zedcloud_network_instance" "required_only" {
    name = "zedcloud_network_instance-required_only-name"
    title = "zedcloud_network_instance-required_only-title"
    kind = "NETWORK_INSTANCE_KIND_SWITCH"
    port = "eth1"
    device_id = data.zedcloud_edgenode.required_only.id
}
