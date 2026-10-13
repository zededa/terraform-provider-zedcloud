package resources

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/go-cty/cty"
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
			"request_timeout": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Timeout for a single HTTP request to ZEDCloud, as a Go duration (for example `60s`). " +
					"Resources with a `timeouts` block are bounded by it as a whole; this bounds each attempt. " +
					"Defaults to `60s`, or the `ZEDCLOUD_REQUEST_TIMEOUT` environment variable.",
				DefaultFunc:      schema.EnvDefaultFunc("ZEDCLOUD_REQUEST_TIMEOUT", "60s"),
				ValidateDiagFunc: validatePositiveDuration,
			},
			"max_retries": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "How many times a failed read is retried. Writes are only retried when the request " +
					"provably did not reach ZEDCloud. Defaults to `4`, or the `ZEDCLOUD_MAX_RETRIES` environment variable.",
				DefaultFunc: schema.EnvDefaultFunc("ZEDCLOUD_MAX_RETRIES", 4),
			},
			"retry_wait_max": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Longest backoff between retries, as a Go duration. Defaults to `30s`, or the " +
					"`ZEDCLOUD_RETRY_WAIT_MAX` environment variable.",
				DefaultFunc:      schema.EnvDefaultFunc("ZEDCLOUD_RETRY_WAIT_MAX", "30s"),
				ValidateDiagFunc: validatePositiveDuration,
			},
			"async_writes": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: "Create, update and delete `zedcloud_application_instance` through the job-based " +
					"`/v2/apps/instances` API and wait for the job, instead of a synchronous call that can time out " +
					"under load. Defaults to `true`, or the `ZEDCLOUD_ASYNC_WRITES` environment variable.",
				DefaultFunc: schema.EnvDefaultFunc("ZEDCLOUD_ASYNC_WRITES", true),
			},
			"job_poll_interval_max": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Advanced. Longest interval between job status polls when `async_writes` is on, as a " +
					"Go duration. Polls start at 2s and back off to this value. Defaults to `10s`, or the " +
					"`ZEDCLOUD_JOB_POLL_INTERVAL_MAX` environment variable.",
				DefaultFunc:      schema.EnvDefaultFunc("ZEDCLOUD_JOB_POLL_INTERVAL_MAX", "10s"),
				ValidateDiagFunc: validatePositiveDuration,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"zedcloud_edgenode":                  NodeDataSource(),
			"zedcloud_network":                   NetworkDataSource(),
			"zedcloud_network_instance":          NetworkInstanceDataSource(),
			"zedcloud_application":               ApplicationDataSource(),
			"zedcloud_application_instance":      ApplicationInstanceDataSource(),
			"zedcloud_image":                     ImageDataSource(),
			"zedcloud_datastore":                 DatastoreDataSource(),
			"zedcloud_volume_instance":           VolumeInstanceDataSource(),
			"zedcloud_project":                   ProjectDataSource(),
			"zedcloud_deployment":                DeploymentDataSource(),
			"zedcloud_user":                      UserDataSource(),
			"zedcloud_role":                      RoleDataSource(),
			"zedcloud_credential":                CredentialDataSource(),
			"zedcloud_patch_envelope":            PatchEnvelopeDataSource(),
			"zedcloud_model":                     HardwareModelDataSource(),
			"zedcloud_brand":                     HardwareBrandDataSource(),
			"zedcloud_auth_profile":              AuthProfileDataSource(),
			"zedcloud_edgenode_cluster":          ClusterDataSource(),
			"zedcloud_enterprise":                EnterpriseDataSource(),
			"zedcloud_app_profile":               AppProfileDataResource(),
			"zedcloud_asset_group":               AssetGroupDataResource(),
			"zedcloud_profile_deployment":        ProfileDeploymentDataResource(),
			"zedcloud_kubernetes_cluster":        ZksInstanceDataSource(),
			"zedcloud_kubernetes_cluster_import": ZksInstanceDataSource(),
			"zedcloud_kubernetes_cluster_status": ZksInstanceStatusDataSource(),
			"zedcloud_kubernetes_cluster_group":  ClusterGroupDataSource(),
			"zedcloud_cluster_group_manifest":    ClusterGroupManifestDataSource(),
			"zedcloud_helm_chart_management":     HelmChartManagementDataSource(),
			"zedcloud_private_helm_repository":   PrivateHelmRepositoryDataSource(),
			"zedcloud_kubernetes_deployment":     KubernetesDeploymentDataSource(),
			"zedcloud_kubernetes_gitops":         KubernetesGitOpsDataSource(),
			"zedcloud_kubernetes_secret":         KubernetesSecretDataSource(),
			"zedcloud_cep_profile":               CEPProfileDataSource(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"zedcloud_edgenode":                  NodeResource(),
			"zedcloud_network":                   NetworkResource(),
			"zedcloud_network_instance":          NetworkInstanceResource(),
			"zedcloud_application":               ApplicationResource(),
			"zedcloud_artifact":                  ArtifactResource(),
			"zedcloud_application_instance":      ApplicationInstanceResource(),
			"zedcloud_image":                     ImageResource(),
			"zedcloud_datastore":                 DatastoreResource(),
			"zedcloud_volume_instance":           VolumeInstanceResource(),
			"zedcloud_project":                   ProjectResource(),
			"zedcloud_deployment":                DeploymentResource(),
			"zedcloud_user":                      UserResource(),
			"zedcloud_role":                      RoleResource(),
			"zedcloud_credential":                CredentialResource(),
			"zedcloud_patch_envelope":            PatchEnvelopeResource(),
			"zedcloud_patch_reference_update":    PatchEnvelopeReferenceUpdate(),
			"zedcloud_model":                     HardwareModelResource(),
			"zedcloud_brand":                     HardwareBrandResource(),
			"zedcloud_auth_profile":              AuthProfileResource(),
			"zedcloud_edgenode_cluster":          ClusterResource(),
			"zedcloud_enterprise":                EnterpriseResource(),
			"zedcloud_app_profile":               AppProfileResource(),
			"zedcloud_asset_group":               AssetGroupResource(),
			"zedcloud_profile_deployment":        ProfileDeploymentResource(),
			"zedcloud_kubernetes_cluster":        ZksInstanceResource(),
			"zedcloud_kubernetes_cluster_import": ZksInstanceImportResource(),
			"zedcloud_kubernetes_cluster_status": ZksInstanceStatusResource(),
			"zedcloud_kubernetes_cluster_group":  ClusterGroupResource(),
			"zedcloud_cluster_group_manifest":    ClusterGroupManifestResource(),
			"zedcloud_helm_chart_management":     HelmChartManagementResource(),
			"zedcloud_private_helm_repository":   PrivateHelmRepositoryResource(),
			"zedcloud_kubernetes_deployment":     KubernetesDeploymentResource(),
			"zedcloud_kubernetes_gitops":         KubernetesGitOpsResource(),
			"zedcloud_kubernetes_secret":         KubernetesSecretResource(),
			"zedcloud_cep_profile":               CEPProfileResource(),
		},
	}
}

type HttpTransportWrapper struct {
	http.RoundTripper
	providerVersion string
	userAgent       string
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
	wrapper.userAgent = fmt.Sprintf("zededa-terraform-provider/%s", wrapper.providerVersion)

	return wrapper
}

// RoundTrip adds the provider's User-Agent. It must not write shared state:
// it runs concurrently for every in-flight request (it used to overwrite the
// package-level version string on each call, a data race at high
// parallelism that also corrupted the version of later wrappers).
func (h *HttpTransportWrapper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", h.userAgent)
	req.Header.Add("X-Custom-User-Agent", h.userAgent)
	resp, err := h.RoundTripper.RoundTrip(req)
	if err == nil {
		rewriteNonJSONError(resp)
	}
	return resp, err
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

	opts, settings, diags := providerOptions(d)
	if diags.HasError() {
		return nil, diags
	}

	// Calls that do not carry a Terraform context get go-openapi's
	// DefaultTimeout as their whole budget, retries included. It used to be
	// the library's 30s; size it from the retry settings instead (UE-177 §4.1).
	httptransport.DefaultTimeout = opts.callBudget()

	zedCloudURL = strings.TrimPrefix(zedCloudURL, "https://")
	stdRetryClient := getRetryClient(opts).StandardClient()
	stdRetryClient.Transport = NewHttpTransportWrapper(stdRetryClient.Transport)
	transport := httptransport.NewWithClient(zedCloudURL, "/api", []string{"https"}, stdRetryClient)
	httpSessionDebugEnabled := envVarIsEnabled("TF_HTTP_SESSION_DEBUG")
	if httpSessionDebugEnabled {
		transport.SetDebug(true)
		transport.SetLogger(HTLogger{})
	}
	transport.DefaultAuthentication = BearerToken(token)

	c := client.New(transport, strfmt.Default)
	c.Settings = settings
	return c, nil
}

// providerOptions reads the HTTP and async-write settings from the provider
// block.
func providerOptions(d *schema.ResourceData) (httpOptions, *client.Settings, diag.Diagnostics) {
	opts := defaultHTTPOptions()
	settings := client.DefaultSettings()
	var diags diag.Diagnostics

	if v, ok := d.Get("request_timeout").(string); ok && v != "" {
		dur, err := time.ParseDuration(v)
		if err != nil {
			diags = append(diags, diag.Errorf("request_timeout: %s", err)...)
		} else {
			opts.requestTimeout = dur
		}
	}
	if v, ok := d.Get("retry_wait_max").(string); ok && v != "" {
		dur, err := time.ParseDuration(v)
		if err != nil {
			diags = append(diags, diag.Errorf("retry_wait_max: %s", err)...)
		} else {
			opts.retryWaitMax = dur
		}
	}
	if v, ok := d.Get("max_retries").(int); ok {
		if v < 0 {
			diags = append(diags, diag.Errorf("max_retries must not be negative")...)
		} else {
			opts.maxRetries = v
		}
	}
	if v, ok := d.Get("async_writes").(bool); ok {
		settings.AsyncWrites = v
	}
	if v, ok := d.Get("job_poll_interval_max").(string); ok && v != "" {
		dur, err := time.ParseDuration(v)
		if err != nil {
			diags = append(diags, diag.Errorf("job_poll_interval_max: %s", err)...)
		} else {
			settings.JobPollIntervalMax = dur
			if settings.JobPollIntervalMin > dur {
				settings.JobPollIntervalMin = dur
			}
		}
	}
	return opts, settings, diags
}

func validatePositiveDuration(v interface{}, path cty.Path) diag.Diagnostics {
	s, _ := v.(string)
	dur, err := time.ParseDuration(s)
	if err != nil {
		return diag.Diagnostics{{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("invalid duration %q", s),
			Detail:        "Use a Go duration such as 30s or 2m.",
			AttributePath: path,
		}}
	}
	if dur <= 0 {
		return diag.Diagnostics{{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("duration %q must be positive", s),
			AttributePath: path,
		}}
	}
	return nil
}

// BearerToken provides a header based oauth2 bearer access token auth info writer
func BearerToken(token string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		r.SetHeaderParam("Authorization", "Bearer "+token)
		return nil
	})
}
