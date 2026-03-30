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
		interface_ordering = "INTERFACE_ORDERING_ENABLED"
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
	name = "test_tf_provider-amd64"
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

resource "zedcloud_network" "complete_with_pac" {
    depends_on = [
        zedcloud_project.test_tf_provider
    ]
    # required
    name = "zedcloud_network.complete_with_pac.name"
    project_id = zedcloud_project.test_tf_provider.id
    title = "zedcloud_network.complete_with_pac.title"
    ip {
        dhcp = "NETWORK_DHCP_TYPE_STATIC"
        dhcp_range {
            start = "172.25.24.1"
            end = "172.25.24.3"
        }
        dns = [
            "172.25.24.254",
            "172.25.24.255"
        ]
        domain = "example.com"
        gateway = "172.25.24.254"
        mask = "255.255.0.0"
        ntp = ""
        subnet = "172.25.24.0/22"
    }
    kind = "NETWORK_KIND_V4"

    # optional
    description = "terraform test network"

    # not supported
    dns_list {
        addrs = [ "10.10.10.1", "10.2.2.2"]
        hostname = "www.example.com"
    }

    enterprise_default = false

    proxy {
        pacfile = "testpacfile"
    }

    wireless {
        cellular {
            apn = "1234"
        }
        type = "NETWORK_WIRELESS_TYPE_WIFI"
        wifi {
            key_scheme = "NETWORK_WIFIKEY_SCHEME_WPAPSK"
            priority = 5
            wifi_ssid = "test-SSID"
        }
    }
}

resource "zedcloud_edgenode" "test_tf_provider" {
	depends_on = [
		zedcloud_project.test_tf_provider,
		zedcloud_model.test_tf_provider,
		zedcloud_network.complete_with_pac
	]
	name        = "test_tf_provider"
	model_id    = zedcloud_model.test_tf_provider.id
	project_id  = zedcloud_project.test_tf_provider.id
	title       = "test_tf_provider-title"
	description = "device with default network instance - updated"
	serialno    = "def-netinst"
	admin_state = "ADMIN_STATE_CREATED"

	interfaces {
		intfname   = "Audio"
		intf_usage = "ADAPTER_USAGE_UNSPECIFIED"
		ztype      = "IO_TYPE_AUDIO"
	}
	interfaces {
		intfname   = "COM4"
		intf_usage = "ADAPTER_USAGE_UNSPECIFIED"
		ztype      = "IO_TYPE_COM"
	}
	interfaces {
		intfname   = "ethernet0"
		intf_usage = "ADAPTER_USAGE_MANAGEMENT"
		ztype      = "IO_TYPE_ETH"
	}
	interfaces {
		intfname   = "ethernet1"
		intf_usage = "ADAPTER_USAGE_UNSPECIFIED"
		ztype      = "IO_TYPE_ETH"
	}
	interfaces {
		intfname   = "ethernet2"
		intf_usage = "ADAPTER_USAGE_UNSPECIFIED"
		ztype      = "IO_TYPE_ETH"
	}
	interfaces {
		intfname   = "ethernet3"
		intf_usage = "ADAPTER_USAGE_UNSPECIFIED"
		ztype      = "IO_TYPE_ETH"
	}
	interfaces {
		intfname   = "ethernet4"
		intf_usage = "ADAPTER_USAGE_UNSPECIFIED"
		ztype      = "IO_TYPE_ETH"
	}

	vlan_adapters {
		logical_label    = "test123"
		lower_layer_name = "ethernet0"
		vlan_id          = 1
		interface {
			netname    = zedcloud_network.complete_with_pac.name
			intf_usage = "ADAPTER_USAGE_APP_SHARED"
			ipaddr     = "172.25.24.2"
			intfname   = "test123"
			netid      = zedcloud_network.complete_with_pac.id
		}
	}
	vlan_adapters {
		logical_label    = "test1234"
		lower_layer_name = "ethernet0"
		vlan_id          = 2
		interface {
			netname    = zedcloud_network.complete_with_pac.name
			intf_usage = "ADAPTER_USAGE_APP_SHARED"
			ipaddr     = "172.25.24.2"
			intfname   = ""
			netid      = zedcloud_network.complete_with_pac.id
		}
	}
	vlan_adapters {
		logical_label    = "test12345"
		lower_layer_name = "bondTest2"
		vlan_id          = 2
		interface {
			netname    = zedcloud_network.complete_with_pac.name
			intf_usage = "ADAPTER_USAGE_APP_SHARED"
			ipaddr     = "172.25.24.2"
			intfname   = ""
			netid      = zedcloud_network.complete_with_pac.id
		}
	}

	bond_adapter {
		logical_label     = "bondTest"
		lower_layer_names = ["ethernet2", "ethernet4"]
		bond_mode         = "BOND_MODE_802_3AD"
		lacp_rate         = "LACP_RATE_FAST"
		arp {
			interval = 2
		}
		interface {
			netname    = zedcloud_network.complete_with_pac.name
			intf_usage = "ADAPTER_USAGE_APP_SHARED"
			ipaddr     = "172.25.24.2"
			intfname   = ""
			netid      = zedcloud_network.complete_with_pac.id
		}
	}
	bond_adapter {
		logical_label     = "bondTest2"
		lower_layer_names = ["ethernet3"]
		bond_mode         = "BOND_MODE_ACTIVE_BACKUP"
	}

	tags = {
		"update-key" = "update-value"
	}
}
