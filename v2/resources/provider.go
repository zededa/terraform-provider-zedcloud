package resources

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/client"
)

var (
	version    string = "dev"
	defaultURL string = "zedcontrol.local.zededa.net"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureContextFunc: ProviderConfigure,
		Schema: map[string]*schema.Schema{
			"zedcloud_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ZEDCloud url. Ex: https://zedcontrol.zededa.net",
				DefaultFunc: schema.EnvDefaultFunc("TF_VAR_zedcloud_url", defaultURL),
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
			"zedcloud_user":                 UserDataSource(),
			"zedcloud_role":                 RoleDataSource(),
			"zedcloud_credential":           CredentialDataSource(),
			"zedcloud_patch_envelope":       PatchEnvelopeDataSource(),
			"zedcloud_model":                HardwareModelDataSource(),
			"zedcloud_brand":                HardwareBrandDataSource(),
			"zedcloud_auth_profile":         AuthProfileDataSource(),
			"zedcloud_edgenode_cluster":     ClusterDataSource(),
			"zedcloud_enterprise":           EnterpriseDataSource(),
			"zedcloud_app_profile":          AppProfileDataResource(),
			"zedcloud_asset_group":          AssetGroupDataResource(),
			"zedcloud_profile_deployment":   ProfileDeploymentDataResource(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"zedcloud_edgenode":               NodeResource(),
			"zedcloud_network":                NetworkResource(),
			"zedcloud_network_instance":       NetworkInstanceResource(),
			"zedcloud_application":            ApplicationResource(),
			"zedcloud_application_instance":   ApplicationInstanceResource(),
			"zedcloud_image":                  ImageResource(),
			"zedcloud_datastore":              DatastoreResource(),
			"zedcloud_volume_instance":        VolumeInstanceResource(),
			"zedcloud_project":                ProjectResource(),
			"zedcloud_deployment":             DeploymentResource(),
			"zedcloud_user":                   UserResource(),
			"zedcloud_role":                   RoleResource(),
			"zedcloud_credential":             CredentialResource(),
			"zedcloud_patch_envelope":         PatchEnvelopeResource(),
			"zedcloud_patch_reference_update": PatchEnvelopeReferenceUpdate(),
			"zedcloud_model":                  HardwareModelResource(),
			"zedcloud_brand":                  HardwareBrandResource(),
			"zedcloud_auth_profile":           AuthProfileResource(),
			"zedcloud_edgenode_cluster":       ClusterResource(),
			"zedcloud_enterprise":             EnterpriseResource(),
			"zedcloud_app_profile":            AppProfileResource(),
			"zedcloud_asset_group":            AssetGroupResource(),
			"zedcloud_profile_deployment":     ProfileDeploymentResource(),
		},
	}
}

type HttpTransportWrapper struct {
	http.RoundTripper
	providerVersion string
}

func NewHttpTransportWrapper(rt http.RoundTripper) *HttpTransportWrapper {
	wrapper := &HttpTransportWrapper{RoundTripper: rt}
	execName, err := os.Executable()
	if err != nil {
		wrapper.providerVersion = "v0.0.0"
	} else {
		if !strings.Contains(execName, "resources.test") {
			execNameParts := strings.Split(execName, "_")
			if len(execNameParts) != 2 {
				if version != "dev" {
					wrapper.providerVersion = fmt.Sprintf("v%s", version)
				} else {
					wrapper.providerVersion = "v0.0.0"
				}
			} else {
				wrapper.providerVersion = execNameParts[1]
			}
		} else {
			wrapper.providerVersion = "testbuild"
		}
	}

	return wrapper
}

func (h *HttpTransportWrapper) RoundTrip(req *http.Request) (*http.Response, error) {
	version = fmt.Sprintf("zededa-terraform-provider/%s", h.providerVersion)
	req.Header.Add("User-Agent", version)
	req.Header.Add("X-Custom-User-Agent", version)
	return h.RoundTripper.RoundTrip(req)
}

func ProviderConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	zedCloudURL, urlIsSet := d.Get("zedcloud_url").(string)
	if !urlIsSet || zedCloudURL == "" {
		return nil, diag.FromErr(errors.New("zedcloud URL must be set"))
	}

	token, tokenIsSet := d.Get("zedcloud_token").(string)
	if !tokenIsSet || token == "" {
		return nil, diag.FromErr(errors.New("zedcloud API key must be set"))
	}

	if strings.HasPrefix(zedCloudURL, "https://") {
		zedCloudURL = strings.TrimPrefix(zedCloudURL, "https://")
	}
	retryClient := getRetryClient()
	transport := httptransport.NewWithClient(zedCloudURL, "/api", []string{"https"}, retryClient.StandardClient())
	httpSessionDebugEnabled := envVarIsEnabled("TF_HTTP_SESSION_DEBUG")
	if httpSessionDebugEnabled {
		transport.SetDebug(true)
		transport.SetLogger(HTLogger{})
	}
	transport.DefaultAuthentication = BearerToken(token)
	transport.Transport = NewHttpTransportWrapper(transport.Transport)

	return client.New(transport, strfmt.Default), nil
}

// BearerToken provides a header based oauth2 bearer access token auth info writer
func BearerToken(token string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		r.SetHeaderParam("Authorization", "Bearer "+token)
		return nil
	})
}

func getRetryClient() *retryablehttp.Client {
	retryClient := retryablehttp.NewClient()
	retryClient.CheckRetry = func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		if resp.StatusCode == http.StatusNotFound {
			return true, fmt.Errorf("unexpected HTTP status %s", resp.Status)
		}
		return retryablehttp.DefaultRetryPolicy(ctx, resp, err)
	}
	return retryClient
}
