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
  zedcloud_url = "zedcloud.local.zededa.net"
  zedcloud_token = "c8xBwliOAXyrnTRh46dWDREM0lXZHqsrWWTyDMAM4UQ4gnoLBhXXUS6Pq4NaOuPCJsDBOLJAhhdooeLOipufmRMWeYc1c5oC-iU2B8cRnJRlBYHvxP5-Rf7XjwzdhsczJ6CXvzQU4lVwRvgEEWo_W06kVM5zmLHAExPWXUJ3ncE="
}


resource "zedcloud_zasset_groups" "example" {
  name        = "example-asset-group"
  description = "My example asset group"
  # Add other required fields as per your schema, e.g.:
  asset_ids {
    ids = [ "asset-id-1", "asset-id-2" ]
  }
  # tags        = ["tag1", "tag2"]
}

resource "zedcloud_zasset_groups" "example2" {
  name        = "example-asset-group2"
  description = "My example asset group"
  # Add other required fields as per your schema, e.g.:
  asset_ids {
    ids = [ "asset-id-3", "asset-id-4" ]
  }
  # tags        = ["tag1", "tag2"]
}

data "zedcloud_zasset_groups" "example2" {
  name = zedcloud_zasset_groups.example2.name
  depends_on = [
    zedcloud_zasset_groups.example2
  ]
}