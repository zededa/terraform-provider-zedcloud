terraform {
  required_providers {
    zedcloud = {
      source = "zededa/zedcloud"
      version = "1.0.1"
    }
  }
}

provider "zedcloud" {
    zedcloud_url = var.zedcloud_url
    username = var.username   // Read as enviornment variable or from file
    password = var.password    // Read as environment vairable or from file
}

data "zedcloud_datastore" "image_store" {
   name = var.datastorename  //Can be read as variable
}

data "zedcloud_image" "vm_image"{
    name = var.imagename // Can be read as variable
}

data "zedcloud_edgenode" "test_testedgedevice" {
   name =  var.deviceaname // Read as enviornment variable or from file
}

resource "zedcloud_network_instance" "test_localnetworkinstance_em" {
    title = "Local network instance"
    name = var.networkinstancename  //var.name can be read as variable
    description = "test local network instance"
    device_default = false
    port = "uplink"
    device_id = data.zedcloud_edgenode.test_testedgedevice.id
    kind = "NETWORK_INSTANCE_KIND_LOCAL"
    type = "NETWORK_INSTANCE_DHCP_TYPE_V4"
}



resource "zedcloud_volume_instance" "test_blockstorage_em" {
    accessmode  = "VOLUME_INSTANCE_ACCESS_MODE_READONLY"
    description = "mutable volume"
    device_id   = data.zedcloud_edgenode.test_testedgedevice.id
    multiattach = false
    name        =  var.blockstoragename
    title       = "This is the  mutable volume"
    type        = "VOLUME_INSTANCE_TYPE_BLOCKSTORAGE"
    size_bytes =  var.size_bytes
    label        = var.tag //Can be used as variable
    
    depends_on = [
       data.zedcloud_image.vm_image
    ]
}



resource "zedcloud_edgeapp" "container_test_em" {
    name = var.edgeappbundlename    # can be read as variable. var.name
    manifest_file = "./container-manifest.json"
    title = "container_manifest_test"
    description = "Container with custom config."
    depends_on = [
       data.zedcloud_image.vm_image
    ]
}


resource "zedcloud_edgeapp_instance" "ContainerInstance_em" {
    name = var.edgeappinstancename   # can be read as variable  var.name
    activate = true
    app_id = zedcloud_edgeapp.container_test_em.id
    app_type = "APP_TYPE_CONTAINER"
    description = "container custom config test"
    device_id = data.zedcloud_edgenode.test_testedgedevice.id
    deployment_type = "DEPLOYMENT_TYPE_STAND_ALONE"
    interface { 
        intfname = "eth0"
        netinstname = zedcloud_network_instance.test_localnetworkinstance_em.name
    }
    logs{
        access = true
    }
    title = "ContainerInstance"
    custom_config {
        name = var.custom_config_name
        add = "true"
        override  = "false"
        field_delimiter = var.custom_config_field_delimter
        template =  var.custom_config_template
        variable_group {
            name = "Default Group 1"
            variable {
                name = var.variablegroup_name
                label = var.variablegroup_label
                required = true
                value = var.ip   # ip of the device. can be read as variable //var.ip
                format = "VARIABLE_FORMAT_TEXT"
                encode = "FILE_ENCODING_UNSPECIFIED"

            }

        }

    }
    depends_on = [
        zedcloud_edgeapp.container_test_em,
        zedcloud_network_instance.test_localnetworkinstance_em,
    ]
}







