// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	zedcloudapi "github.com/zededa/zedcloud-api"
)

var DefaultZedcloudUrl = "https://zedcontrol.zededa.net"

var ProviderSchema = &schema.Provider{
	ConfigureContextFunc: providerConfigure,
	Schema:               zschemas.ProviderSchema,

	DataSourcesMap: map[string]*schema.Resource{
		"zedcloud_datastore":        getDataStoreDataSourceSchema(),
		"zedcloud_edgeapp":          getEdgeAppDataSourceSchema(),
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

// Client holds metadata / config for use by Terraform resources
type Client struct {
	// ZedcloudUrl Ex: "https://zedcontrol.zededa.net"
	ZedcloudUrl string
	// Token = API Token to be used
	//  Either Token OR Username / Password must be specified.
	Token    string
	Username string
	Password string
	Client   *zedcloudapi.Client
}

func Provider() *schema.Provider {
	return ProviderSchema
}

// providerConfigure parses the config into the Terraform provider meta object
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	zedcloudUrl := rdEntryStr(d, "zedcloud_url")
	token := rdEntryStr(d, "token")
	username := rdEntryStr(d, "username")
	password := rdEntryStr(d, "password")

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if zedcloudUrl == "" {
		zedcloudUrl = DefaultZedcloudUrl
	}

	if token != "" {
		username = ""
		password = ""
		log.Println("Token Specified. Using Token for Authentication")
	} else if username == "" && password == "" {
		err := fmt.Errorf("Neither token(%s) nor username(%s)/password(%s) specified",
			token, username, password)
		log.Printf("%s", err)
		return nil, diag.Errorf(err.Error())
	}
	log.Printf("providerConfigure - Configuring Provider. Creating new client")

	providerClient := Client{
		ZedcloudUrl: zedcloudUrl,
		Token:       token,
		Username:    username,
		Password:    password,
	}
	var err error
	providerClient.Client, err = zedcloudapi.NewClient(zedcloudUrl, token, username, password)
	if err != nil {
		log.Printf("Failed to create new client. err: %+v", err)
		return nil, diag.Errorf(err.Error())
	}
	return providerClient, diags
}
