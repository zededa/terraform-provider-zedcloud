package resources

import (
	"context"
	"errors"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	identity_access_management "github.com/zededa/terraform-provider-zedcloud/v2/client/identity_access_management"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
	models "github.com/zededa/terraform-provider-zedcloud/v2/models"
)

/*
IdentityAccessManagement identity access management API
*/

func LogoutResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: IdentityAccessManagement_Logout,
		ReadContext:   schema.NoopContext,
		UpdateContext: schema.NoopContext,
		DeleteContext: schema.NoopContext,
		Schema: zschema.AAAFrontendLogoutRequestSchema(),
	}
}

func LogoutDataSource() *schema.Resource {
	return &schema.Resource{
		Schema: zschema.AAAFrontendLogoutRequestSchema(),
	}
}

func IdentityAccessManagement_Logout(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.AAAFrontendLogoutRequestModel(d)
	params := identity_access_management.NewIdentityAccessManagementLogoutParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementLogout(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.ErrorCode != nil && *err.ErrorCode == models.ErrorCodeSuccess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if diags.HasError() {
			return diags
		}
	}

	return diags
}