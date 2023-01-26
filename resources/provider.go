package resources

import (
	"context"
	"errors"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureContextFunc: ProviderConfigure,
		Schema: map[string]*schema.Schema{
			"zedcloud_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ZEDCloud url. Ex: https://zedcontrol.zededa.net",
				Default:     "zedcontrol.local.zededa.net",
				DefaultFunc: schema.EnvDefaultFunc("TF_VAR_zedcloud_url", nil),
			},
			"zedcloud_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ZEDCloud API token",
				DefaultFunc: schema.EnvDefaultFunc("TF_VAR_zedcloud_token", nil),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"zedcloud_edgenode": EdgeNodeDataSource(),
			"zedcloud_network":  NetworkDataSource(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"zedcloud_edgenode": EdgeNodeResource(),
			"zedcloud_network":  NetworkResource(),
		},
	}
}

func ProviderConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	zedCloudURL := d.Get("zedcloud_url").(string)

	token, tokenIsSet := d.Get("zedcloud_token").(string)
	if !tokenIsSet || token == "" {
		return nil, diag.FromErr(errors.New("zedcloud API key must be set"))
	}

	transport := httptransport.New(zedCloudURL, "/api", []string{"https"})
	transport.SetDebug(false)
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
