package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
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

func AuthProfileResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: IdentityAccessManagement_CreateAuthProfile,
		DeleteContext: IdentityAccessManagement_DeleteAuthProfile,
		ReadContext: IdentityAccessManagement_GetAuthProfile,
		UpdateContext: IdentityAccessManagement_UpdateAuthProfile,
		Schema: zschema.AuthorizationProfileSchema(),
	}
}

func AuthProfileDataSource() *schema.Resource {
	return &schema.Resource{
		Schema: zschema.AuthorizationProfileSchema(),
	}
}

func IdentityAccessManagement_GetAuthProfile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var user *models.AuthorizationProfile

	if _, isSet := d.GetOk("name"); isSet {
		user, diags = GetAuthProfileByName(ctx, d, m)
	} else if _, isSet := d.GetOk("id"); isSet {
		user, diags = GetAuthProfileById(ctx, d, m)
	}

	if diags.HasError() {
		return diags
	}

	d.SetId(user.ID)

	return diags
}

func GetAuthProfileById(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.AuthorizationProfile, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := identity_access_management.NewIdentityAccessManagementGetAuthProfileParams()

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

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementGetAuthProfile(params, nil)
	if err != nil {
		log.Printf("[TRACE] auth profile read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return nil, diags
		}

		diags = append(diags, diag.Errorf("auth profile read error: %s", err)...)
		return nil, diags
	}

	respModel := resp.GetPayload()
	zschema.SetAuthorizationProfileResourceData(d, respModel)

	return respModel, diags
}

func GetAuthProfileByName(ctx context.Context, d *schema.ResourceData, m interface{})(*models.AuthorizationProfile, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := identity_access_management.NewIdentityAccessManagementGetAuthProfileByNameParams()

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

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementGetAuthProfileByName(params, nil)
	if err != nil {
		log.Printf("[TRACE] auth profile read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return nil, diags
		}

		diags = append(diags, diag.Errorf("auth profile read error: %s", err)...)
		return nil, diags
	}

	respModel := resp.GetPayload()
	zschema.SetAuthorizationProfileResourceData(d, respModel)

	return respModel, diags
}


func IdentityAccessManagement_CreateAuthProfile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.AuthorizationProfileModel(d)
	params := identity_access_management.NewIdentityAccessManagementCreateAuthProfileParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementCreateAuthProfile(params, nil)
	if err != nil {
		log.Printf("[TRACE] auth profile create error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("auth profile create error: %s", err)...)
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
	if diags := IdentityAccessManagement_GetAuthProfile(ctx, d, m); diags.HasError() {
		return diags
	}

	return diags
}


func IdentityAccessManagement_UpdateAuthProfile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	params := identity_access_management.NewIdentityAccessManagementUpdateAuthProfileParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.AuthorizationProfileModel(d))
	// IdentityAccessManagementUpdateAuthProfileBody

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
	resp, err := client.IdentityAccessManagement.IdentityAccessManagementUpdateAuthProfile(params, nil)
	if err != nil {
		log.Printf("[TRACE] auth profile update error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("auth profile update error: %s", err)...)
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
	if errs := IdentityAccessManagement_GetUser(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	d.Partial(false)

	return diags
}


func IdentityAccessManagement_DeleteAuthProfile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := identity_access_management.NewIdentityAccessManagementDeleteAuthProfileParams()

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

	_, err := client.IdentityAccessManagement.IdentityAccessManagementDeleteAuthProfile(params, nil)
	if err != nil {
		log.Printf("[TRACE] auth profile delete error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("auth profile delete error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}