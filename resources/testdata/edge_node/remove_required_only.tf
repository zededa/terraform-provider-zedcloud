// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0
//

terraform {
  required_providers {
    zedcloud = {
      source = "zededa/zedcloud"
      //version = "0.0.7-alpha"
    }
  }
}

provider "zedcloud" {
	  zedcloud_url = "zedcontrol.local.zededa.net"
    username = var.username
    password = var.password
    zedcloud_token = var.zedcloud_token
}
