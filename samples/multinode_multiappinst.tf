provider "zedcloud" {
	hostname = ""
    token = ""
    username = ""
    password = ""
}

locals {
    num_nodes = 2
    node_list = toset(["edge-node-1" , "edge-node-2"])
}

resource "zedcloud_edgenode" "sample-edge-node" {
    for_each = local.node_list
    name = each.value
    title = format("Edge Node %s", each.value)
    description = format("Edge Node Description %s", each.value)
    eveimage_version = "abc"
}

resource "zedcloud_network_instance" "default-nwinst" {
    for_each = local.node_list
    name = format("TF-default-%s", each.value)
    title = format("Network Instance default-%s", each.value)
    description = format("Network Instance Description default-%s", each.value)
    device_default = true
    port = "eth0"
    edgenode_id = zedcloud_edgenode.sample-edge-node[each.key].id
}

resource "zedcloud_edgeapp" "TF-Test-ubuntu-all-ip" {
    name = "TF-Test-ubuntu-all-ip"
    title = "TF-Test-ubuntu-all-ip"
    description = "TF-Test-ubuntu-all-ip"
    manifest_file = "./TFTest-ubuntu-xenial-16.04.json"
}

resource "zedcloud_edgeapp" "tftest-ubuntu-xenial" {
    name = "TFTest-ubuntu-xenial-16.04"
    manifest_file = "./TFTest-ubuntu-xenial-16.04.json"
    title = "TFTest-ubuntu-xenial-16.04"
    description = "TFTest-ubuntu-xenial-16.04"
}

resource "zedcloud_edgeappinstance" "Sample-EdgeAppInstance" {
    for_each = local.node_list
    name = format("TF-Sample-EdgeAppInstance-%s", each.value)
    activate = false
    edgeapp_type="VM"
    edgeapp_id=zedcloud_edgeapp.TF-Test-ubuntu-all-ip.id
    edgenode_id=zedcloud_edgenode.sample-edge-node[each.key].id
    title="TerraForm Sample App Instance"
    description="TerraForm Sample App Instance"
    network_instance=zedcloud_network_instance.default-nwinst[each.key].name
    depends_on = [
        zedcloud_edgeapp.TF-Test-ubuntu-all-ip,
        zedcloud_edgenode.sample-edge-node,
        zedcloud_network_instance.default-nwinst,
    ]
}
