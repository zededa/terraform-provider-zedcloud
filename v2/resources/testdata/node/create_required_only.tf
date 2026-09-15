// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

resource "zedcloud_project" "test_tf_provider" {
	# required
	name = "test_tf_provider-create_node_1__SUFFIX__"
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


resource "zedcloud_datastore"  "test_tf_provider" {
	depends_on = [
		zedcloud_project.test_tf_provider
	]
	# required
	ds_fqdn = "https://my-datastore.my-company.com"
	ds_path = "download/AMD64"
	ds_type = "DATASTORE_TYPE_AZUREBLOB"
	name = "test_tf_provider-test_datastore__SUFFIX__"
	title = "test_tf_provider-test_datastore__SUFFIX__"
	description = "test_tf_provider-test_datastore"
	region = "eu"
	project_access_list = [zedcloud_project.test_tf_provider.id]
}

resource "zedcloud_image" "test_tf_provider" {
	depends_on = [
		zedcloud_datastore.test_tf_provider,
		zedcloud_project.test_tf_provider
	]
	name = "test_tf_provider-create_edgenode__SUFFIX__"
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
	name = "test_tf_provider-create_edgenode__SUFFIX__"
	title = "test_tf_provider-create_edgenode__SUFFIX__"
	description = "description"
	origin_type = "ORIGIN_LOCAL"
}

resource "zedcloud_model" "test_tf_provider" {
	brand_id = zedcloud_brand.test_tf_provider.id
	name = "test_tf_provider-create_edgenode__SUFFIX__"
	title = "test_tf_provider-create_edgenode__SUFFIX__"
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

resource "zedcloud_network" "test_tf_provider" {
	depends_on = [
		zedcloud_project.test_tf_provider
	]
	name = "test_tf_provider-required_only-net__SUFFIX__"
	project_id = zedcloud_project.test_tf_provider.id
	title = "test_tf_provider-required_only-net__SUFFIX__"
	kind = "NETWORK_KIND_V4"
	ip {
		dhcp = "NETWORK_DHCP_TYPE_STATIC"
		dhcp_range {
			start = "172.25.24.1"
			end = "172.25.24.3"
		}
		dns = ["172.25.24.254"]
		domain = "example.com"
		gateway = "172.25.24.254"
		mask = "255.255.0.0"
		ntp = ""
		subnet = "172.25.24.0/22"
	}
}

resource "zedcloud_edgenode" "required_only" {
	depends_on = [
		zedcloud_project.test_tf_provider,
		zedcloud_model.test_tf_provider,
		zedcloud_network.test_tf_provider
	]
	name = "test_tf_provider-required_only__SUFFIX__"
	model_id = zedcloud_model.test_tf_provider.id
	project_id = zedcloud_project.test_tf_provider.id
	title = "required_only-title"
	interfaces {
		cost = 0
		intf_usage = "ADAPTER_USAGE_MANAGEMENT"
		intfname = "ethernet0"
		netname = zedcloud_network.test_tf_provider.name
		netid = zedcloud_network.test_tf_provider.id
		net_dhcp = zedcloud_network.test_tf_provider.ip[0].dhcp
		ipaddr = "172.25.24.2"
		tags = {}
	}
}
