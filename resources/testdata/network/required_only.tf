resource "zedcloud_network" "required_only" {
		# computed
    # id =

		# required
    name = "zedcloud_network.required_only"
		title = "zedcloud_network.required_only.title"
		project_id = "4754cd0f-82d7-4e06-a68f-ff9e23e75ccf"
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
