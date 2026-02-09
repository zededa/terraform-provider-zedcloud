resource "zedcloud_project" "test_tf_provider" {
    # required
    name = "test_tf_provider-netinst_fwd_lldp_3"
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
    name = "test_tf_provider-ds_fwd_lldp_3"
    title = "test_tf_provider-ds_fwd_lldp_3"
    description = "test_tf_provider-ds_fwd_lldp_3"
    region = "eu"
    project_access_list = [zedcloud_project.test_tf_provider.id]
}

resource "zedcloud_image" "test_tf_provider" {
    depends_on = [
        zedcloud_datastore.test_tf_provider,
        zedcloud_project.test_tf_provider
    ]
    name = "test_tf_provider-img_fwd_lldp_3"
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
    name = "test_tf_provider-brand_fwd_lldp_3"
    title = "test_tf_provider-brand_fwd_lldp_3"
    description = "description"
    origin_type = "ORIGIN_LOCAL"
}

resource "zedcloud_model" "test_tf_provider" {
    brand_id = zedcloud_brand.test_tf_provider.id
    name = "test_tf_provider-model_fwd_lldp_3"
    title = "test_tf_provider-model_fwd_lldp_3"
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
    name = "test_tf_provider-node_fwd_lldp_3"
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

# Test network instance with forward_lldp enabled
resource "zedcloud_network_instance" "test_tf_netinst_forward_lldp" {
    depends_on = [
        zedcloud_edgenode.test_tf_provider
    ]

    # required
    device_id = zedcloud_edgenode.test_tf_provider.id
    name = "test_tf_provider-netinst_fwd_lldp_3"
    title = "test_tf_provider-netinst_fwd_lldp_3"
    kind = "NETWORK_INSTANCE_KIND_SWITCH"
    port = "eth1"

    # optional
    description = "Test network instance with LLDP forwarding enabled"
    device_default = false
    
    # Test forward_lldp flag set to true
    forward_lldp = true

    tags = {
        "test-type" = "forward-lldp-test"
    }
}
