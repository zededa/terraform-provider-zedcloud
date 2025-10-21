package resources

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	identity_access_management "github.com/zededa/terraform-provider-zedcloud/v2/client/identity_access_management"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

/*
IdentityAccessManagement identity access management API
*/

func LoginExternalOAuth2CallbackResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: IdentityAccessManagement_LoginExternalOAuth2Callback,
		ReadContext: schema.NoopContext,
		UpdateContext: schema.NoopContext,
		DeleteContext: schema.NoopContext,
		Schema: zschema.AAAFrontendOAUTHCallbackRequestSchema(),
	}
}

func LoginExternalOAuth2CallbackDataSource() *schema.Resource {
	return &schema.Resource{
		Schema: zschema.AAAFrontendOAUTHCallbackRequestSchema(),
	}
}

func IdentityAccessManagement_LoginExternalOAuth2Callback(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.AAAFrontendOAUTHCallbackRequestModel(d)
	params := identity_access_management.NewIdentityAccessManagementLoginExternalOAuth2CallbackParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementLoginExternalOAuth2Callback(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	return diags
}