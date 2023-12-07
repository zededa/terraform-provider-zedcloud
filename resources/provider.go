package resources

import (
	"context"
	"errors"
	"strings"

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
			"zedcloud_edgenode":             NodeDataSource(),
			"zedcloud_network":              NetworkDataSource(),
			"zedcloud_network_instance":     NetworkInstanceDataSource(),
			"zedcloud_application":          ApplicationDataSource(),
			"zedcloud_application_instance": ApplicationInstanceDataSource(),
			"zedcloud_image":                ImageDataSource(),
			"zedcloud_datastore":            DatastoreDataSource(),
			"zedcloud_volume_instance":      VolumeInstanceDataSource(),
			"zedcloud_project":              ProjectDataSource(),
			"zedcloud_patch_envelope":       PatchEnvelopeDataSource(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"zedcloud_edgenode":             NodeResource(),
			"zedcloud_network":              NetworkResource(),
			"zedcloud_network_instance":     NetworkInstanceResource(),
			"zedcloud_application":          ApplicationResource(),
			"zedcloud_application_instance": ApplicationInstanceResource(),
			"zedcloud_image":                ImageResource(),
			"zedcloud_datastore":            DatastoreResource(),
			"zedcloud_volume_instance":      VolumeInstanceResource(),
			"zedcloud_project":              ProjectResource(),
			"zedcloud_patch_envelope":       PatchEnvelopeResource(),
		},
	}
}

func ProviderConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	zedCloudURL := d.Get("zedcloud_url").(string)

	token, tokenIsSet := d.Get("zedcloud_token").(string)
	if !tokenIsSet || token == "" {
		return nil, diag.FromErr(errors.New("zedcloud API key must be set"))
	}

	if strings.HasPrefix(zedCloudURL, "https://") {
		zedCloudURL = strings.TrimPrefix(zedCloudURL, "https://")
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
