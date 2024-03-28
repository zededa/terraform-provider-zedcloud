// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package main

// Run the docs generation tool, check its repository for more information on how it works and how docs can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/zededa/terraform-provider-zedcloud/v2/resources"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version   string = "dev"
	debugMode bool
)

func main() {
	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	plugin.Serve(&plugin.ServeOpts{
		Debug:        debugMode,
		ProviderFunc: resources.Provider,
	})
}
