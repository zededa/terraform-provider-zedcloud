package provider

import (
	"context"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/client"
	"github.com/zededa/terraform-provider/resources"
)

const zedCloudProductionURL = "https://zedcontrol.zededa.net"

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureContextFunc: providerConfigure,
		Schema: map[string]*schema.Schema{
			"zedcloud_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ZEDCloud url. Ex: https://zedcontrol.zededa.net",
			},
			"zedcloud_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ZEDCloud API token",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ZEDCloud user",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ZEDCloud user password ",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"zedcloud_edgenode": resources.DataResourceEdgeNodeConfiguration(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"zedcloud_edgenode": resources.EdgeNodeConfiguration(),
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	zedCloudURL, zedCloudIsSet := d.Get("zedcloud_url").(string)
	if !zedCloudIsSet {
		zedCloudURL = zedCloudProductionURL
	}

	// token, tokenIsSet := d.Get("zedcloud_token").(string)
	// username, usernameIsSet := d.Get("username").(string)
	// password, passwordIsSet := d.Get("password").(string)

	// credentialsSet := usernameIsSet && passwordIsSet
	// if !tokenIsSet && !credentialsSet {
	// 	return nil, diag.FromErr(errors.New("bearer token or credentials must be set"))
	// }

	transport := httptransport.New(zedCloudURL, "/api", []string{"https"})
	transport.SetDebug(true)

	token := d.Get("zedcloud_token").(string)

	// var auth runtime.ClientAuthInfoWriter
	// auth = httptransport.BearerToken(token)
	// transport.DefaultAuthentication = httptransport.Compose(auth, httptransport.BasicAuth(username, password))
	transport.DefaultAuthentication = BearerToken(token)

	return client.New(transport, strfmt.Default), nil
}

// BearerToken provides a header based oauth2 bearer access token auth info writer
func BearerToken(token string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		r.SetHeaderParam("Authorization", "Bearer "+token)
		return nil
	})
}
