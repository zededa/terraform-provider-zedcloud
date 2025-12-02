package resources

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"strings"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/zededa/terraform-provider-zedcloud/v2/actions/node"
	node_client "github.com/zededa/terraform-provider-zedcloud/v2/client/node"
)

var _ provider.ProviderWithActions = &actionsProvider{}

func NewActionsProvider() provider.Provider {
	return &actionsProvider{}
}

type actionsProvider struct {
	model     actionsProviderModel
	transport runtime.ClientTransport
}

type actionsProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
	Token    types.String `tfsdk:"token"`
}

func (p *actionsProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "zedcloud" // This prefix will be used for resources: `example_item`
	resp.Version = "1.0.0"
}

func (p *actionsProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "API endpoint for the Actions Service.",
				Optional:            true,
			},
			"token": schema.StringAttribute{
				MarkdownDescription: "API token for the Actions Service.",
				Required:            true,
			},
		},
	}
}

func (p *actionsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config actionsProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 💡 Context value: In a real provider, you would instantiate an *APIClient* struct
	// here using the config values (config.Endpoint.ValueString()) and pass it.
	// We'll pass a simple endpoint string for this example.
	endpoint := config.Endpoint.ValueString()
	//if endpoint == "" {
	//	endpoint = "https://zedcontrol.alpha.zededa.net"
	//}
	token := config.Token.ValueString()
	if len(strings.TrimSpace(token)) == 0 {
		diags.AddError("Configuration Error", "token must be provided")
		return
	}
	//resp.Diagnostics.AddError(token, "Not Implemented")
	//return
	p.model = config

	dialer := &net.Dialer{
		Resolver: &net.Resolver{
			PreferGo: false,
		},
	}
	httpTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext:     dialer.DialContext,
	}
	transport := httptransport.New(endpoint, "/api", []string{"https"})
	transport.DefaultAuthentication = BearerToken(token)
	transport.Transport = httpTransport
	p.transport = transport

	//resp.ResourceData = endpoint // Passes this string to the Resource's client field
	//resp.DataSourceData = endpoint
}

// Resources returns the list of resources implemented by the provider.
func (p *actionsProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		//NewItemResource, // Function to instantiate the resource
	}
}

// DataSources returns the list of data sources implemented by the provider.
func (p *actionsProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil // No data sources for this simple example
}

func (p *actionsProvider) Actions(_ context.Context) []func() action.Action {
	return []func() action.Action{
		p.RebootAction,
	}
}

func (p *actionsProvider) RebootAction() action.Action {
	return node.NewRebootAction(node_client.New(p.transport, strfmt.Default))
}
