package node

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/zededa/terraform-provider-zedcloud/v2/client/node"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

//var _ action.Action = (*RebootAction)(nil)

type RebootAction struct {
	client node.ClientService
}

func NewRebootAction(client node.ClientService) *RebootAction {
	return &RebootAction{client: client}
}

type rebootInput struct {
	ID types.String `tfsdk:"node_id"`
}

func (m *RebootAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_reboot_node"
}

func (m *RebootAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"node_id": schema.StringAttribute{
				Required:    true,
				Description: "The unique identifier of the edge node to reboot.",
			},
		},
	}
}

func (m *RebootAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var input rebootInput
	diags := req.Config.Get(ctx, &input)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := node.NewEdgeNodeConfigurationRebootParams()
	params.ID = input.ID.ValueString()

	//params := node.EdgeNodeConfigurationRebootParams{
	//	ID: input.ID.String(),
	//}

	response, err := m.client.EdgeNodeConfigurationReboot(params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error rebooting edge node", err.Error())
		return
	}
	if response == nil {
		resp.Diagnostics.AddError("Error rebooting edge node", "API response is nil")
		return
	}

	responseData := response.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.ErrorCode != nil && *err.ErrorCode == models.ErrorCodeSuccess {
				continue
			}
			diags.AddError("Error rebooting edge node", err.Details)
		}
		if diags.HasError() {
			return
		}
	}
}
