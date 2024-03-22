package resources

import (
	"context"
	"errors"
	"github.com/zededa/terraform-provider-zedcloud/models"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/client"
	"github.com/zededa/terraform-provider-zedcloud/client/identity_access_management"
	zschema "github.com/zededa/terraform-provider-zedcloud/schemas"
)

func CredentialResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: IdentityAccessManagement_CreateCredential,
		DeleteContext: IdentityAccessManagement_DeleteCredential,
		ReadContext:   schema.NoopContext,
		UpdateContext: IdentityAccessManagement_UpdateCredential,
		Schema:        zschema.CredentialSchema(),
	}
}

func CredentialDataSource() *schema.Resource {
	return &schema.Resource{
		Schema: zschema.CredentialsSchema(),
	}
}

func IdentityAccessManagement_CreateCredential(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.CredentialModel(d)
	params := identity_access_management.NewIdentityAccessManagementCreateCredentialParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementCreateCredential(params, nil)
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

	d.SetId(responseData.ObjectID)

	return diags
}

func IdentityAccessManagement_UpdateCredential(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	params := identity_access_management.NewIdentityAccessManagementUpdateCredentialParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.CredentialModel(d))
	// models.Credential

	// makes a bulk update for all properties that were changed
	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.IdentityAccessManagement.IdentityAccessManagementUpdateCredential(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
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

	d.Partial(false)

	return diags
}

func IdentityAccessManagement_DeleteCredential(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	// due to bug in the zedlcoud implementation of credential deletion, we just reset the credential's id

	d.SetId("")
	return diags
}
