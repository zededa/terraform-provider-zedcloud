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
  endpoint = "zedcontrol.alpha.zededa.net"
  token = "zgf9l4jf:zZuHGJNR_HAR5vllQYdD62aaU7giDXGRFalhxMyYOwCCZ0WyJE0yzZhEg0mryaRrc45-rA-jsxDyJQ6odoMO5PJ0hSMEihwdsuwYBCMJOELY7ilNPV2TnjBVMrkpEqhE-1LoO6OwNci_LL1baGuVgbn7OHYrlLB3GeRuwsh4MJI="
}

action "zedcloud_reboot_node" "reboot_node_1" {
  config {
    node_id = "4f1994cc-86b2-4eef-84fa-dcce2191a837"
  }
}
