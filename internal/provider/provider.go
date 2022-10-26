// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/internal/datasources"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	zedcloudapi "github.com/zededa/zedcloud-api"
)

const defaultZedcloudUrl = "https://zedcontrol.zededa.net"

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureContextFunc: configure,
		Schema:               zschemas.ProviderSchema,
		DataSourcesMap: map[string]*schema.Resource{
			"zedcloud_datastore":        datasources.NewDataStore(),
			"zedcloud_edgeapp":          datasources.NewEdgeApp(),
			"zedcloud_edgeapp_instance": getAppInstDataSourceSchema(),
			"zedcloud_edgenode":         getEdgeNodeDataSourceSchema(),
			"zedcloud_image":            getImageDataSourceSchema(),
			"zedcloud_network":          getNetworkDataSourceSchema(),
			"zedcloud_network_instance": getNetInstDataSourceSchema(),
			"zedcloud_volume_instance":  getVolumeInstanceDataSourceSchema(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"zedcloud_edgeapp":          getEdgeAppResourceSchema(),
			"zedcloud_edgeapp_instance": getAppInstResourceSchema(),
			"zedcloud_edgenode":         getEdgeNodeResourceSchema(),
			"zedcloud_image":            getImageResourceSchema(),
			"zedcloud_network":          getNetworkResourceSchema(),
			"zedcloud_network_instance": getNetInstResourceSchema(),
			"zedcloud_volume_instance":  getVolumeInstanceResourceSchema(),
		},
	}
}

// Token or Username and Password must be specified.
type Client struct {
	*zedcloudapi.Client
	zedcloudURL string // https://zedcontrol.zededa.net
	token       string // API Token
	username    string
	password    string
}

// configure parses the config into the Terraform provider meta object
func configure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := getStr(d, "token")
	username := getStr(d, "username")
	password := getStr(d, "password")
	zedcloudUrl := getStr(d, "zedcloud_url")
	if zedcloudUrl == "" {
		zedcloudUrl = defaultZedcloudUrl
	}

	var diags diag.Diagnostics
	if token != "" {
		username = ""
		password = ""
		log.Println("enable authentication by token")
	} else if username == "" && password == "" {
		return nil, diag.Errorf("require either token or username and password")
	}

	apiClient, err := zedcloudapi.NewClient(zedcloudUrl, token, username, password)
	if err != nil {
		return nil, diag.FromErr(fmt.Errorf("could not initialize api client: %w", err))
	}
	client := &Client{
		Client:      apiClient,
		zedcloudURL: zedcloudUrl,
		token:       token,
		username:    username,
		password:    password,
	}

	log.Println("successfully configured provider")
	return client, diags
}
