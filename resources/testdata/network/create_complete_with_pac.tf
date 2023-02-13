resource "zedcloud_network" "complete_with_pac" {
		# computed
    # id =

		# required
    name = "zedcloud_network.complete_with_pac.name"
		project_id = "4754cd0f-82d7-4e06-a68f-ff9e23e75ccf"
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
        # not supported by API
        # exceptions = "1.1.1.1"

        # related to the PAC file
        # network_proxy = true
        # network_proxy_url = "http://www.proxy.com"
        # network_proxy_certs = []
        pacfile = "testpacfile"
   }

    wireless {
        cellular {
            apn = "1234"
            # not supported by API
            # location_tracking = true
        }
        type = "NETWORK_WIRELESS_TYPE_WIFI"
        wifi {
            key_scheme = "NETWORK_WIFIKEY_SCHEME_WPAPSK"
            priority = 5
            wifi_ssid = "test-SSID"
        }
    }
}

resource "zedcloud_network" "TF-sample-network-full-cfg" {
    name = "TF-Test-Network"
    description = "Test Network for TF testing"
    dns_list {
        addrs = [ "10.10.10.1", "10.2.2.2"]
        hostname = "www.example.com"
    }
    enterprise_default = false
    ip {
        dhcp = "NETWORK_DHCP_TYPE_STATIC"

        dhcp_range {
            start = "172.25.24.1"
            end = "172.25.24.3"
        }
        // Configure Optional DNS server
        dns = [
            "172.25.24.254"
        ]
        domain = "example.com"
        gateway = "172.25.24.254"
        mask = "255.255.0.0"
        ntp = ""
        subnet = "172.25.24.0/22"
    }
    kind = "NETWORK_KIND_V4"
    project_id = "b8552e75-54ee-4607-a715-1469dd132004"
    proxy {
      // Exceptions to Proxy server
      exceptions = "1.1.1.1"
      // Related to the PAC file
      network_proxy = false
      network_proxy_url = "www.proxy.com"
      network_proxy_certs = []
      pacfile = "testpacfile"
      // Specify Proxy Host for each protocol (HTTP, HTTPS, FTP, SOCKS)
      proxy {
        port = 5555
        proto = "NETWORK_PROXY_PROTO_HTTPS"
        server = "example.com"
      }
      proxy {
        port = 5556
        proto = "NETWORK_PROXY_PROTO_HTTPS"
        server = "example.com"
      }
    }
    title = "TF-sample-network-full-cfg"
    wireless {
        cellular_cfg {
            apn = "1234"
        }
        type = ""
        wifi_cfg {
            key_scheme = ""
            priority = 5
            wifi_ssid = "TF-Test-Wifi-SSID"
        }
    }
}
