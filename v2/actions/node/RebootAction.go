package node

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/client/node"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

type RebootAction struct {
	client *node.Client
}

func NewRebootAction(client *node.Client) *RebootAction {
	return &RebootAction{client: client}
}

type rebootInput struct {
	ID string `tfsdk:"id"`
}

func (m *RebootAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
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

	params := node.EdgeNodeConfigurationRebootParams{
		ID: input.ID,
	}

	response, err := m.client.EdgeNodeConfigurationReboot(&params, nil)
	if err != nil {
		diags.AddError("Error rebooting edge node", err.Error())
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
