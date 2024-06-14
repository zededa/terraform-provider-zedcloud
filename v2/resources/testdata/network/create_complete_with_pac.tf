resource "zedcloud_project" "test_tf_provider" {
    # required
    name = "test_tf_provider-create_net_3"
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
