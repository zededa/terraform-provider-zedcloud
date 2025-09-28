resource "zedcloud_project" "test_tf_provider" {
    # required
    name = "test_tf_provider-create_net_1"
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

resource "zedcloud_network" "required_only" {
    depends_on = [
        zedcloud_project.test_tf_provider
    ]
    name = "zedcloud_network.required_only.name"
    title = "zedcloud_network.required_only.title"
    project_id = zedcloud_project.test_tf_provider.id
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
}


resource "zedcloud_network" "adapter_spec_network" {
    depends_on = [
        zedcloud_project.test_tf_provider
    ]
    name = "adapter specific network"
    title = "adapter specific network"
    project_id = zedcloud_project.test_tf_provider.id
    ip {
        dhcp = "NETWORK_DHCP_TYPE_STATIC_ADAPTER_SPECIFIC"
    }
}