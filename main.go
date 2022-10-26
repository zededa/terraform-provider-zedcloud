// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/zededa/terraform-provider-zedcloud/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
