// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var ProviderSchema = map[string]*schema.Schema{
	"zedcloud_url": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "ZEDCloud url. Ex: https://zedcontrol.zededa.net",
	},
	"token": {
		Type:        schema.TypeString,
		Optional:    true,
		Sensitive:   true,
		Description: "API token to be used to login to ZEDCloud. One of token OR username/password must be specified.",
	},
	"username": {
		Type:        schema.TypeString,
		Optional:    true,
		Sensitive:   true,
		Description: "Username to be used to login to ZEDCloud. One of token OR username/password must be specified.",
	},
	"password": {
		Type:        schema.TypeString,
		Optional:    true,
		Sensitive:   true,
		Description: "Password to be used to login to ZEDCloud. One of token OR username/password must be specified.",
	},
}
