resource "zedcloud_project" "test_tf_provider" {
	# required
	name = "test_tf_provider-create_node_2"
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

resource "zedcloud_image" "test_tf_provider" {
	depends_on = [
		zedcloud_datastore.test_tf_provider,
		zedcloud_project.test_tf_provider
	]
	name = "test_tf_provider-create_edgenode"
	datastore_id = zedcloud_datastore.test_tf_provider.id
	image_arch = "AMD64"
	image_format = "CONTAINER"
	image_rel_url = "test_url"
	image_size_bytes = 0
	image_type =  "IMAGE_TYPE_APPLICATION"
	title = "test"
	project_access_list = [zedcloud_project.test_tf_provider.id]
}

resource "zedcloud_brand" "test_tf_provider" {
	name = "test_tf_provider-create_edgenode"
	title = "test_tf_provider-create_edgenode"
	description = "description"
	origin_type = "ORIGIN_LOCAL"
}

resource "zedcloud_model" "test_tf_provider" {
	brand_id = zedcloud_brand.test_tf_provider.id
	name = "test_tf_provider-create_edgenode"
	title = "test_tf_provider-create_edgenode"
	type = "AMD64"
	origin_type = "ORIGIN_LOCAL"
	state = "SYS_MODEL_STATE_ACTIVE"
	attr = {
		memory = "8G"
		storage = "100G"
		Cpus = "4"
	}
	io_member_list {
		ztype = "IO_TYPE_ETH"
		phylabel =  "firstEth"
		usage = "ADAPTER_USAGE_MANAGEMENT"
		assigngrp = "eth0"
		phyaddrs = {
			Ifname = "eth0"
			PciLong = "0000:02:00.0"
		}
		logicallabel = "ethernet0"
		usage_policy = {
			FreeUplink = true
		}
		cost = 0
	}
	depends_on = [
		zedcloud_brand.test_tf_provider
	]
}

resource "zedcloud_edgenode" "test_tf_provider" {
	depends_on = [
		zedcloud_project.test_tf_provider,
		zedcloud_model.test_tf_provider
	]
	name = "test_tf_provider"
	model_id = zedcloud_model.test_tf_provider.id
	project_id = zedcloud_project.test_tf_provider.id
	title = "test_tf_provider-title"

	# optional
	onboarding_key = ""
	serialno = "d6aebfa5-56b6-4b66-9d8e-6552b0e2b45b"
	admin_state = "ADMIN_STATE_INACTIVE"
	asset_id = "asset_id"
	deployment_tag = "depl_tag_update"
	description = "description_update"
	generate_soft_serial = false
	token = "token_update"
	site_pictures = []
	interfaces {
		cost = 255
		intf_usage = "ADAPTER_USAGE_MANAGEMENT"
		intfname = "defaultIPv4"
		ipaddr = "127.0.0.1"
		macaddr = "00:00:00:00:00:00"
		tags = {
			"system_interface_1_key" = "system_interface_1_value"
			"system_interface_2_key" = "system_interface_2_value"
		}
	}
	interfaces {
		intfname = "eth0"
		intf_usage = "ADAPTER_USAGE_UNSPECIFIED"
	}

	config_item {
		bool_value = false
		float_value = 1.0
		key = "key"
		string_value = "string_update"
		uint32_value = 32
		uint64_value = 64
		value_type = "value type"
	}

	dev_location {
		city = "berlin"
		country = "germany"
		freeloc = "freeloc"
		hostname = "hostname"
		loc = "52.520008, 13.404954"
		org = "zededa"
		postal = "10115"
		region = "europe/west"
		underlay_ip = ""
	}

	tags = {
		"tag-key-1" = "tag-value-1"
		"tag-key-2" = "tag-value-2"
		"tag-key-2_update" = "tag-value-2_update"
	}
}