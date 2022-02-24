terraform {
  required_providers {
    zedcloud = {
      source = "zededa/zedcloud"
      version = "1.0.1"
    }
  }
}

provider "zedcloud" {
    zedcloud_url =  var.zedcloud_url
    username = var.username
    password = var.password
}

data "zedcloud_datastore" "image_store" {
   name = var.datastorename
}

data "zedcloud_image" "vm_image"{
    name = var.imagename
}

data "zedcloud_edgenode" "sample_testedgedevice" {
   name = var.deviceaname
}


resource "zedcloud_network_instance" "sample_localnetworkinstance" {
    title = "Local network instance"
    name = "sample_localnetworkinstance"
    description = "sample local network instance"
    device_default = false
    port = "uplink"
    device_id = data.zedcloud_edgenode.sample_testedgedevice.id
    kind = "NETWORK_INSTANCE_KIND_LOCAL"
    type = "NETWORK_INSTANCE_DHCP_TYPE_V4"
}


resource "zedcloud_volume_instance" "sample_contenttree" {
    accessmode  = "VOLUME_INSTANCE_ACCESS_MODE_READONLY"
    description = "Immutable Volume"
    device_id   = data.zedcloud_edgenode.sample_testedgedevice.id
    image       = data.zedcloud_image.vm_image.name
    multiattach = false
    name        =  var.contenttreename
    title       = "This is the volume"
    type        = "VOLUME_INSTANCE_TYPE_CONTENT_TREE"

    depends_on = [
       data.zedcloud_image.vm_image
    ]
}


resource "zedcloud_volume_instance" "sample_blockstorage" {
    accessmode  = "VOLUME_INSTANCE_ACCESS_MODE_READONLY"
    description = "Mutable Volume"
    device_id   = data.zedcloud_edgenode.sample_testedgedevice.id
    multiattach = false
    name        = var.blockstoragename
    title       = "This is the volume"
    type        = "VOLUME_INSTANCE_TYPE_BLOCKSTORAGE"
    content_tree_id = zedcloud_volume_instance.sample_contenttree.id
    label        = var.tag
    
    depends_on = [
       data.zedcloud_image.vm_image,
       zedcloud_volume_instance.sample_contenttree
    ]
}

resource "zedcloud_edgeapp" "centos-test-app" {
    name = var.edgebundlename
    manifest_file = "./app1-manifest.json"
    title = "CentOs-image"
    description = "Centos image"
    depends_on = [
       data.zedcloud_image.vm_image
    ]
}


resource "zedcloud_edgeapp_instance" "sample-EdgeAppInstance" {
    name =  var.edgeappinstancename
    activate = true
    app_id = zedcloud_edgeapp.centos-test-app.id
    app_type = "APP_TYPE_VM"
    description = "sampletemplatetest"
    device_id = data.zedcloud_edgenode.sample_testedgedevice.id
    deployment_type = "DEPLOYMENT_TYPE_STAND_ALONE"
    interface { 
        intfname = "eth0"
        netinstname = zedcloud_network_instance.sample_localnetworkinstance.name
    }
    title = "samples"

    
    depends_on = [
        zedcloud_edgeapp.centos-test-app,
        zedcloud_network_instance.sample_localnetworkinstance,
        zedcloud_volume_instance.sample_contenttree,
        zedcloud_volume_instance.sample_blockstorage   
    ]
}












