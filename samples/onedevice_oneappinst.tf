// Set env variable TF_VAR_username to pass username to terraform plan/apply
variable "username" {
    description = "ZEDCloud username"
    sensitive = true
    type        = string
}

// Set env variable TF_VAR_password to pass password to terraform plan/apply
variable "password" {
    description = "ZEDCloud password"
    sensitive = true
    type        = string
}

// Set env variable TF_VAR_zedcloud_token to pass ZEDCloud API Token to
// terraform plan/apply
variable "zedcloud_token" {
    description = "ZEDCloud API Token"
    sensitive = true
    type        = string
}

provider "zedcloud" {
    zedcloud_url = ""
    token = var.zedcloud_token
    username = var.username
    password = var.password
}

data "zedcloud_edgenode" "Sample-Device-Data" {
    name = "node1"
}

data "zedcloud_image" "app-image-1" {
    name = "app-image-1"
}

resource "zedcloud_image" "TFR-app-image-1" {
    name = "TFR-app-image-1"
    datastore_id = data.zedcloud_image.app-image-1.datastore_id
    description = data.zedcloud_image.app-image-1.description
    image_arch = data.zedcloud_image.app-image-1.image_arch
    image_format = data.zedcloud_image.app-image-1.image_format
    image_local = data.zedcloud_image.app-image-1.image_local
    image_rel_url = data.zedcloud_image.app-image-1.image_rel_url
    image_sha_256 = data.zedcloud_image.app-image-1.image_sha_256
    image_size_bytes = data.zedcloud_image.app-image-1.image_size_bytes
    image_type = data.zedcloud_image.app-image-1.image_type
    title = data.zedcloud_image.app-image-1.title
}

resource "zedcloud_edgenode" "Sample-Device" {
    name = "TF-Sample-Device"
    asset_id = "sample asset id"
    client_ip = ""
    config_items = {
        "timer.config.interval" = "100"
        "debug.disable.dhcp.all-ones.netmask" = "true"
        // Disable the network-fallback option
        "network.fallback.any.eth" = "disabled"
        // Set the reboot-on-no-network timer to 1 year
        "timer.reboot.no.network" = "31536000"
        // Set the config checkpoint timer to 1 year
        "timer.use.config.checkpoint" = "31536000"
    }
    description = "Sample device Description"
    dev_location {
        city = "San Jose"
        country = "USA"
        free_loc = ""
        hostname = "sample-device"
        latlong = ""
        loc = ""
        org = "Zededa Engineering"
        postal = "95116"
        region = "Alameda County"
        underlay_ip = "10.1.1.1"
    }
    // The specified EVE image needs to be present in the Zedcloud ( Image Object)
    eve_image_version = "6.8.2-kvm-amd64"

    interface {
      intf_name = "eth0"
      intf_usage = "ADAPTER_USAGE_MANAGEMENT"
    }
    interface {
      intf_name = "eth1"
      intf_usage = "ADAPTER_USAGE_UNSPECIFIED"
    }
    tags = {
        "tag1" = "tag-value1"
        "tag2" = "tag-value2"
    }
    title = "Sample Device Title Test"

    lifecycle {
        prevent_destroy = true
    }
}

data "zedcloud_network" "sample-network-data" {
    name = "defaultIPv4-net"
}

data "zedcloud_network" "wifi_network" {
    name = "wifi_network"
}

data "zedcloud_network" "test-lte-network" {
    name = "test-lte-network"
}

data "zedcloud_network" "test-network-static" {
    name = "test-network-static"
}

resource "zedcloud_network" "sample-network" {
    name = "TF-Test-Network"
    enterprise_default = false
    description = "Test Network for TF testing"
    kind = "NETWORK_KIND_V4"
    title = "TF-Test-Network"
    project_id   = "99999888-bbbb-aaaa-1111-bbbbbbbbbbbb"
    dns_list {
        addrs = [ "10.10.10.1", "10.2.2.2"]
        hostname = "www.example.com"
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
            crypto {
                identity = "test identity"
                password = "test password"
            }
            crypto_key = ""
            encrypted_secret {
                "secret1" = "value1"
                "secret2" = "value2"
            }
            identity = ""
            key_scheme = ""
            priority = 5
            secret {}
            wifi_ssid = "TF-Test-Wifi-SSID"
        }
    }
}

data "zedcloud_edgeapp" "ubuntu-all-ip-data" {
    name = "ubuntu-all-ip"
}

resource "zedcloud_edgeapp" "ubuntu-all-ip" {
    name = "ubuntu-all-ip"
    title = "ubuntu-all-ip"
    description = "ubuntu-all-ip"
    manifest_file = "./TFTest-ubuntu-xenial-16.04.json"
    user_defined_version = "1.1"
}

resource "zedcloud_edgeapp" "tftest-ubuntu-xenial" {
    name = "TFTest-ubuntu-xenial-16.04"
    description = "TFTest-ubuntu-xenial-16.04"
    manifest_file = "./TFTest-ubuntu-xenial-16.04.json"
    title = "TFTest-ubuntu-xenial-16.04"
    user_defined_version = "3"
}

data "zedcloud_network_instance" "default-nwinst-data" {
    name = "Sample-datasource-edgenode"
}

resource "zedcloud_network_instance" "default-nwinst" {
    name = "edge-node-1-ni-default-nwinst"
    title = "edge-node-1-ni-default-nwinst"
    description = "edge-node-1-ni-default-nwinst"
    device_default = true
    port = "eth1"
    device_id = zedcloud_edgenode.Sample-Device.id
}

resource "zedcloud_network_instance" "TF-Sample-test-full" {
    name = "TF-Sample-test-full"
    title = "edge-node-1-ni-TF-Sample-test-full"
    description = "edge-node-1-ni-TF-Sample-test-full"
    device_default = true
    device_id = zedcloud_edgenode.Sample-Device.id
    dns_list {
        addrs = [
            "10.1.1.1",
            "10.1.1.2"
        ]
        hostname = "wwww.ns1.example.com"
    }
    dns_list {
        addrs = [
            "10.1.2.1",
            "10.1.2.2"
        ]
        hostname = "wwww.ns2.example.com"
    }
    ip {
        dhcp_range {
            end = "10.10.1.0"
            start = "10.10.1.255"
        }
        dns = [
            "www.ns1.example.com",
            "www.ns2.example.com"
        ]
        domain = "example.com"
        gateway = "10.1.0.1"
        mask = "255.255.255.0"
        ntp = "10.1.0.2"
        subnet = "10.1.0.0"
    }
    opaque {
        oconfig = "Test OConfig"
        type = "OPAQUE_CONFIG_TYPE_UNSPECIFIED"
    }
    kind = "NETWORK_INSTANCE_KIND_LOCAL"
    network_policy_id = "tft-test-network-policy-id"
    oconfig = "tft-test-oconfig"
    port = "eth1"
    port_tags = {
        "port-tag1" = "port-tag-value-1"
        "port-tag2" = "port-tag-value-2"
    }
    tags = {
        "ni-tag1" = "ni-tag-value-1"
        "ni-tag2" = "ni-tag-value-2"
    }
    type = "NETWORK_INSTANCE_DHCP_TYPE_V4"
}

data "zedcloud_edgeapp_instance" "Sample-EdgeAppInstance-data" {
    name = "TF-Sample-EdgeAppInstance1"
}

resource "zedcloud_edgeapp_instance" "Sample-EdgeAppInstance" {
    name = "TF-Sample-Resource-EdgeAppInstance"
    activate = true
    app_type = "APP_TYPE_VM"
    app_id = zedcloud_edgeapp.ubuntu-all-ip.id
    //app_user_defined_version = "5"
    description = "TerraForm Sample App Instance - TESTING Update"
    device_id = zedcloud_edgenode.Sample-Device.id
    interfaces {
        intfname = "indirect"
        netinstname = zedcloud_network_instance.default-nwinst.name
        acl {
            matches {
                type = "ip"
                value = "0.0.0.0/0"
            }
            id = "1"
            name = ""
        }
        acl {
            matches {
                type = "protocol"
                value = "tcp"
            }
            matches {
                type = "lport"
                value = "8022"
            }
            matches {
                type = "ip"
                value = "0.0.0.0/0"
            }
            actions {
                drop = false
                limit = false
                limit_rate = 0
                limit_unit = ""
                limit_burst = 0
                portmap = true
                map_params {
                    port = 22
                }
            }
            id = 2
            name = ""
        }
    }
    title = "TerraForm Sample App Instance"
    remote_console = true
    vminfo {
        cpus = 1
        memory = 512
        mode = "HV_HVM"
        vnc = true
    }
    depends_on = [
        zedcloud_edgeapp.ubuntu-all-ip,
        zedcloud_edgenode.Sample-Device,
        zedcloud_network_instance.default-nwinst
    ]
}

resource "zedcloud_volume_instance" "TF-sample-vol-inst-RO-content-tree" {
    accessmode  = "VOLUME_INSTANCE_ACCESS_MODE_READONLY"
    description = "Immutable Volume"
    device_id   = "99999999-bbbb-aaaa-1111-bbbbbbbbbbbb"
    id          = "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
    image       = "alpine-base"
    implicit    = true
    multiattach = false
    name        = "sample-volinst-1"
    project_id  = "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
    size_bytes  = "140443648"
    title       = "Sample Title"
    type        = "VOLUME_INSTANCE_TYPE_CONTENT_TREE"
}
