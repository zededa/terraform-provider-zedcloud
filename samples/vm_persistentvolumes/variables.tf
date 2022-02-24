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

variable "contenttreename"{
    description = "Name of the content tree"
    sensitive = false
    type      = string

}

variable "blockstoragename"{
    description = "Nameof the block storage"
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

variable "edgebundlename"{
    description = "Name of the tag"
    sensitive = false
    type      = string
}
