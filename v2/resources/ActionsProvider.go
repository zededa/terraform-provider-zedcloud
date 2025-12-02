package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &actionsProvider{}

func NewActionsProvider() provider.Provider {
	return &actionsProvider{}
}

type actionsProvider struct {
}

type actionsProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

func (p *actionsProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "example" // This prefix will be used for resources: `example_item`
	resp.Version = "1.0.0"
}

func (p *actionsProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "API endpoint for the Example Service.",
				Optional:            true,
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
	clientConfig := config.Endpoint.ValueString()
	if clientConfig == "" {
		clientConfig = "https://api.example.com"
	}

	resp.ResourceData = clientConfig // Passes this string to the Resource's client field
	resp.DataSourceData = clientConfig
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
