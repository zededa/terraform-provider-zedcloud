resource "zedcloud_project" "test_tf_provider" {
    # required
    name = "test_tf_provider-create_netinst_2"
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
    name = "test_tf_provider-test_datastore"
    title = "test_tf_provider-test_datastore"
    description = "test_tf_provider-test_datastore"
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
    name = "test_tf_provider-create_edgenode"
    model_id = zedcloud_model.test_tf_provider.id
    project_id = zedcloud_project.test_tf_provider.id
    title = "title"
    interfaces {
        cost = 0
        intf_usage = "ADAPTER_USAGE_MANAGEMENT"
        intfname = "ethernet0"
        tags = {}
    }
}

resource "zedcloud_network_instance" "test_tf_netinst_complete" {
    depends_on = [
        zedcloud_edgenode.test_tf_provider
    ]

    # required
    device_id = zedcloud_edgenode.test_tf_provider.id
    name = "test_tf_provider-netinst_complete"
    title = "test_tf_provider-netinst_complete"
    kind = "NETWORK_INSTANCE_KIND_LOCAL"
    port = "eth1"

    # optional
    description = "test_tf_provider-netinst_complete-description"
    type = "NETWORK_INSTANCE_DHCP_TYPE_V4"
    device_default = false
    dhcp = false
    dns_list {
        addrs = [
            "10.1.2.2",
            "10.1.2.1"
        ]
        hostname = "wwww.ns2.example.com"
    }
    dns_list {
        addrs = [
            "10.1.1.1",
            "10.1.1.2"
        ]
        hostname = "wwww.ns1.example.com"
    }
    ip {
        domain = "htttp://example.com"
        gateway = "10.0.20.1"
        subnet = "10.0.20.0/24"
        dhcp_range {
            end = "10.0.20.100"
            start = "10.0.20.50"
        }
        mask = "255.255.255.0"
        ntp = "10.1.0.2"
    }
    opaque {
        oconfig = "Test OConfig"
        type = "OPAQUE_CONFIG_TYPE_UNSPECIFIED"
    }
    tags = {
        "ni-tag1" = "ni-tag-value-1"
        "ni-tag2" = "ni-tag-value-2"
    }
    static_routes {
      gateway = "10.0.20.1"
      probe_config {
          enable_gateway_ping = true
          ping_max_cost = 50
          prefer_lower_cost = true
          custom_probe_config {
            probe_method = "CONNECTIVITY_PROBE_METHOD_TCP"
            probe_endpoint {
                host = "example.com"
                port = 8080
            }
          }
      }
    }
}
