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

func LoginExternalResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: IdentityAccessManagement_LoginExternal,
		ReadContext: schema.NoopContext,
		UpdateContext: schema.NoopContext,
		DeleteContext: schema.NoopContext,
		Schema: zschema.AAAFrontendLoginWithOauthRequestSchema(),
	}
}

func LoginExternalDataSource() *schema.Resource {
	return &schema.Resource{
		Schema: zschema.AAAFrontendLoginWithOauthRequestSchema(),
	}
}

func IdentityAccessManagement_LoginExternal(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.AAAFrontendLoginWithOauthRequestModel(d)
	params := identity_access_management.NewIdentityAccessManagementLoginExternalParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementLoginExternal(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	return diags
}
