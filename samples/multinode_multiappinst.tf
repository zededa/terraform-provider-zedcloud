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

locals {
    num_nodes = 2
    node_list = toset(["edge-node-1" , "edge-node-2"])
}

resource "zedcloud_edgenode" "sample-edge-node" {
    for_each = local.node_list
    name = each.value
    title = format("Edge Node %s", each.value)
    description = format("Edge Node Description %s", each.value)
    eve_image_version = "abc"
}

resource "zedcloud_network_instance" "default-nwinst" {
    for_each = local.node_list
    name = format("TF-default-%s", each.value)
    title = format("Network Instance default-%s", each.value)
    description = format("Network Instance Description default-%s", each.value)
    device_default = true
    port = "eth0"
    device_id = zedcloud_edgenode.sample-edge-node[each.key].id
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

resource "zedcloud_edgeapp_instance" "Sample-EdgeAppInstance" {
    for_each = local.node_list
    name = format("TF-Sample-EdgeAppInstance-%s", each.value)
    activate = false
    app_type = "APP_TYPE_VM"
    app_id = zedcloud_edgeapp.TF-Test-ubuntu-all-ip.id
    device_id=zedcloud_edgenode.sample-edge-node[each.key].id
    title="TerraForm Sample App Instance"
    description="TerraForm Sample App Instance"
    depends_on = [
        zedcloud_edgeapp.TF-Test-ubuntu-all-ip,
        zedcloud_edgenode.sample-edge-node,
        zedcloud_network_instance.default-nwinst,
    ]
}
