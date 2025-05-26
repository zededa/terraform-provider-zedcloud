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

resource "zedcloud_project" "test_tf_provider" {
  # required
  name  = "test_tf_provider"
  title = "title"

  # optional
  type = "TAG_TYPE_PROJECT"
  tag_level_settings {
    flow_log_transmission = "NETWORK_INSTANCE_FLOW_LOG_TRANSMISSION_DISABLED"
    interface_ordering    = "INTERFACE_ORDERING_DISABLED"
  }
}

resource "zedcloud_role" "tf_jmp_role" {
  depends_on = [ zedcloud_project.test_tf_provider ]
  name        = "TF-JMP-Role"
  title       = "TF JMP Role"
  description = "Terraform-created role for JMP Project"
  type        = "USER_ROLE_USER_DEFINED"
  state       = "ROLE_STATE_ACTIVE"

  scopes {
    access_device = "PermissionAccessCreateReadUpdateDelete"
    project_filter = [
      zedcloud_project.test_tf_provider.id
    ]
  }
}
