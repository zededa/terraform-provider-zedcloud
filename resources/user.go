package resources

import (
	"context"
	"errors"
	models "github.com/zededa/terraform-provider-zedcloud/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/schemas"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/client"
	identity_access_management "github.com/zededa/terraform-provider-zedcloud/client/identity_access_management"
)

/*
User management API
*/

func UserResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: IdentityAccessManagement_CreateUser,
		DeleteContext: IdentityAccessManagement_DeleteUser,
		ReadContext:   IdentityAccessManagement_GetUser,
		UpdateContext: IdentityAccessManagement_UpdateUser2,
		Schema:        zschema.DetailedUserSchema(),
	}
}

func UserDataSource() *schema.Resource {
	return &schema.Resource{
		Schema: zschema.DetailedUserSchema(),
	}
}

func IdentityAccessManagement_GetUser(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var user *models.DetailedUser

	if _, isSet := d.GetOk("username"); isSet {
		user, diags = getUserByName(ctx, d, m)
	} else if _, isSet := d.GetOk("id"); isSet {
		user, diags = getUserById(ctx, d, m)
	}

	if diags.HasError() {
		return diags
	}

	d.SetId(user.ID)

	return diags
}

func getUserById(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.DetailedUser, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := identity_access_management.NewIdentityAccessManagementGetUserParams()

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

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementGetUser(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return nil, append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	respModel := resp.GetPayload()
	zschema.SetDetailedUserResourceData(d, respModel)

	return respModel, diags
}

func getUserByName(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.DetailedUser, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := identity_access_management.NewIdentityAccessManagementGetUserByNameParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	nameVal, nameIsSet := d.GetOk("username")
	if nameIsSet {
		params.Name = nameVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: username")...)
		return nil, diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementGetUserByName(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return nil, append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	respModel := resp.GetPayload()
	zschema.SetDetailedUserResourceData(d, respModel)

	return respModel, diags
}

func IdentityAccessManagement_CreateUser(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.DetailedUserModel(d)
	params := identity_access_management.NewIdentityAccessManagementCreateUserParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementCreateUser(params, nil)
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

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if diags := IdentityAccessManagement_GetUser(ctx, d, m); diags.HasError() {
		return diags
	}

	return diags
}

func IdentityAccessManagement_UpdateUser2(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	params := identity_access_management.NewIdentityAccessManagementUpdateUser2Params()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.DetailedUserModel(d))
	// IdentityAccessManagementUpdateUser2Body

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
	resp, err := client.IdentityAccessManagement.IdentityAccessManagementUpdateUser2(params, nil)
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

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := IdentityAccessManagement_GetUser(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	d.Partial(false)

	return diags
}

func IdentityAccessManagement_DeleteUser(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := identity_access_management.NewIdentityAccessManagementDeleteUserParams()

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

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementDeleteUser(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
