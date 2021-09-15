// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	zedcloudapi "github.com/zededa/zedcloud-api"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
)

var DefaultZedcloudUrl = "https://zedcontrol.zededa.net"

var ProviderSchema = &schema.Provider{
	ConfigureContextFunc: providerConfigure,
	Schema: zschemas.ProviderSchema,

	DataSourcesMap: map[string]*schema.Resource{
		"zedcloud_edgeapp":          getEdgeAppDataSourceSchema(),
		"zedcloud_edgeapp_instance":  getAppInstDataSourceSchema(),
		"zedcloud_edgenode":         getEdgeNodeDataSourceSchema(),
		"zedcloud_image":            getImageDataSourceSchema(),
		"zedcloud_network":          getNetworkDataSourceSchema(),
		"zedcloud_network_instance": getNetInstDataSourceSchema(),
		"zedcloud_volume_instance": getVolumeInstanceDataSourceSchema(),
	},

	ResourcesMap: map[string]*schema.Resource{
		"zedcloud_edgeapp":          getEdgeAppResourceSchema(),
		"zedcloud_edgeapp_instance":  getAppInstResourceSchema(),
		"zedcloud_edgenode":         getEdgeNodeResourceSchema(),
		"zedcloud_image":            getImageResourceSchema(),
		"zedcloud_network":          getNetworkResourceSchema(),
		"zedcloud_network_instance": getNetInstResourceSchema(),
		"zedcloud_volume_instance": getVolumeInstanceResourceSchema(),
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

func getErrMsgPrefix(name, id, objectType, action string) string {
	return fmt.Sprintf("[ERROR] %s %s ( id: %s) %s Failed.",
		objectType, name, id, action)
}

func Provider() *schema.Provider {
	return ProviderSchema
}

// marshalData is used to ensure the data is put into a format Terraform can output
func marshalData(d *schema.ResourceData, vals map[string]interface{}) {
	for k, v := range vals {
		if k == "id" {
			d.SetId(v.(string))
		} else {
			str, ok := v.(string)
			if ok {
				d.Set(k, str)
			} else {
				d.Set(k, v)
			}
		}
	}
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
		Token:    token,
		Username: username,
		Password: password,
	}
	var err error
	providerClient.Client, err = zedcloudapi.NewClient(zedcloudUrl, token, username, password)
	if err != nil {
		log.Printf("Failed to create new client. err: %+v", err)
		return nil, diag.Errorf(err.Error())
	}
	return providerClient, diags
}
