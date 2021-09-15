// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var ProviderSchema = map[string]*schema.Schema{
    "zedcloud_url": {
        Type:     schema.TypeString,
        Optional: true,
    },

    "token": {
        Type:     schema.TypeString,
        Optional: true,
    },
    "username": {
        Type:     schema.TypeString,
        Optional: true,
    },
    "password": {
        Type:     schema.TypeString,
        Optional: true,
    },
}
