resource "zedcloud_project" "test_tf_provider" {
  # required
  name = "test_tf_provider-create_net_2"
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


resource "zedcloud_network" "complete_with_proxy" {
  depends_on = [
    zedcloud_project.test_tf_provider
  ]
  # required
  name = "zedcloud_network.complete_with_proxy.name"
  title = "zedcloud_network.complete_with_proxy.title"
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

  # optional
  description = "terraform test network"

  # not supported
  dns_list {
    addrs = [ "10.2.2.2", "10.10.10.1" ]
    hostname = "www.example.com"
  }

  enterprise_default = false

  proxy {
    proxies {
      port = 5557
      proto = "NETWORK_PROXY_PROTO_SOCKS"
      server = "example.com"
    }
    proxies {
      port = 5555
      proto = "NETWORK_PROXY_PROTO_HTTP"
      server = "example.com"
    }
    proxies {
      port = 5556
      proto = "NETWORK_PROXY_PROTO_HTTPS"
      server = "example.com"
    }
  }

  wireless {
    cellular {
      apn = "1234"
      # not supported by API
      location_tracking = true
    }
    type = "NETWORK_WIRELESS_TYPE_CELLULAR"
  }
}
