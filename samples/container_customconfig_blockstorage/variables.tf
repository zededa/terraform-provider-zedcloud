variable "username" {
    description = "ZEDCloud username"
    sensitive = true
    type        = string
}

variable "password" {
    description = "ZEDCloud password"
    sensitive = true
    type        = string
}


variable "deviceaname" {
    description = "device to use"
    sensitive = false
    type        = string
}

variable "datastorename"{
    description = "Name of the datastore to use"
    sensitive = false
    type        = string
}

variable "imagename"{
    description = "Name of the image to use"
    sensitive = false
    type        = string
}

variable "zedcloud_url" {
    description = "zedcloud url"
    sensitive = false
    type        = string
}


variable "blockstoragename"{
    description = "Nameof the block storage"
    sensitive = false
    type      = string

}

variable "networkinstancename"{
    description = "Nameof the network instance"
    sensitive = false
    type      = string

}

variable "tag"{
    description = "Name of the tag"
    sensitive = false
    type      = string
}

variable "edgeappinstancename"{
    description = "Name of the tag"
    sensitive = false
    type      = string
}

variable "edgeappbundlename"{
    description = "Name of the tag"
    sensitive = false
    type      = string
}



variable "custom_config_name"{
    description = "Name of the custom config"
    sensitive = false
    type      = string
}


variable "custom_config_field_delimter"{
    description = "Name of the custom config"
    sensitive = false
    type      = string
}


variable "custom_config_template"{
    description = "Name of the custom config"
    sensitive = false
    type      = string
}


variable "variablegroup_name"{
    description = "Name of the custom config"
    sensitive = false
    type      = string
}


variable "variablegroup_label"{
    description = "Name of the custom config"
    sensitive = false
    type      = string
}


variable "ip"{
    description = "Ip Address to be used according to config"
    sensitive = false
    type      = string
}

variable "size_bytes"{
    description = "Size of the block storage"
    sensitive = false
    type      = string
}


