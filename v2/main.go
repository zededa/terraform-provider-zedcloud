// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package main

// Run the docs generation tool, check its repository for more information on how it works and how docs can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/zededa/terraform-provider-zedcloud/v2/resources"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version   string = "dev"
	debugMode bool
)

func main() {
	err := providerserver.Serve(context.Background(), resources.NewActionsProvider, providerserver.ServeOpts{
		// This address must match the source attribute in your Terraform configuration
		Address: "hashicorp.com/edu/example",
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
