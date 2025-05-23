package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	identity_access_management "github.com/zededa/terraform-provider-zedcloud/v2/client/identity_access_management"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

/*
IdentityAccessManagement identity access management API
*/

func RoleResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: IdentityAccessManagement_CreateRole,
		DeleteContext: IdentityAccessManagement_DeleteRole,
		ReadContext:   IdentityAccessManagement_GetRole,
		UpdateContext: IdentityAccessManagement_UpdateRole,
		Schema:        zschema.RoleSchema(),
	}
}

func RoleDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext:   IdentityAccessManagement_GetRole,
		Schema: zschema.RoleSchema(),
	}
}

func IdentityAccessManagement_GetRole(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var role *models.Role

	if _, isSet := d.GetOk("name"); isSet {
		role, diags = getRoleByName(ctx, d, m)
	} else if _, isSet := d.GetOk("id"); isSet {
		role, diags = getRoleById(ctx, d, m)
	}

	if diags.HasError() {
		return diags
	}

	d.SetId(role.ID)

	return diags
}

func getRoleById(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.Role, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := identity_access_management.NewIdentityAccessManagementGetRoleParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return nil, diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementGetRole(params, nil)
	if err != nil {
		log.Printf("[TRACE] role read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return nil, diags
		}

		diags = append(diags, diag.Errorf("role read error: %s", err)...)
		return nil, diags
	}

	respModel := resp.GetPayload()
	zschema.SetRoleResourceData(d, respModel)
	d.SetId(respModel.ID)

	return respModel, diags
}

func getRoleByName(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.Role, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := identity_access_management.NewIdentityAccessManagementGetRoleByNameParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	nameVal, nameIsSet := d.GetOk("name")
	if nameIsSet {
		params.Name = nameVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: name")...)
		return nil, diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementGetRoleByName(params, nil)
	if err != nil {
		log.Printf("[TRACE] role read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return nil, diags
		}

		diags = append(diags, diag.Errorf("role read error: %s", err)...)
		return nil, diags
	}

	respModel := resp.GetPayload()
	zschema.SetRoleResourceData(d, respModel)
	d.SetId(respModel.ID)

	return respModel, diags
}

func IdentityAccessManagement_CreateRole(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.RoleModel(d)
	params := identity_access_management.NewIdentityAccessManagementCreateRoleParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementCreateRole(params, nil)
	if err != nil {
		log.Printf("[TRACE] role create error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("role create error: %s", err)...)
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

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if diags := IdentityAccessManagement_GetRole(ctx, d, m); diags.HasError() {
		return diags
	}

	return diags
}

func IdentityAccessManagement_UpdateRole(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	params := identity_access_management.NewIdentityAccessManagementUpdateRoleParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.RoleModel(d))
	// IdentityAccessManagementUpdateRoleBody

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	// makes a bulk update for all properties that were changed
	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.IdentityAccessManagement.IdentityAccessManagementUpdateRole(params, nil)
	if err != nil {
		log.Printf("[TRACE] role update error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("role update error: %s", err)...)
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

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if diags := IdentityAccessManagement_GetRole(ctx, d, m); diags.HasError() {
		return diags
	}

	d.Partial(false)

	return diags
}

func IdentityAccessManagement_DeleteRole(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := identity_access_management.NewIdentityAccessManagementDeleteRoleParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	_, err := client.IdentityAccessManagement.IdentityAccessManagementDeleteRole(params, nil)
	if err != nil {
		log.Printf("[TRACE] role delete error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("role delete error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
