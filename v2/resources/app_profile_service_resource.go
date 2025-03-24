package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	app_profile_service "github.com/zededa/terraform-provider-zedcloud/v2/client/app_profile_service"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

/*
AppProfileService app profile service API
*/

func AppProfileResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateAppProfile,
		DeleteContext: DeleteAppProfile,
		ReadContext:   ReadAppProfile,
		UpdateContext: UpdateAppProfile,
		Schema:        zschema.AppProfileReadSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func AppProfileDataResource() *schema.Resource {
	return &schema.Resource{
		Schema:      zschema.AppProfileSchema(),
		ReadContext: ReadAppProfile,
	}
}

func CreateAppProfile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.AppProfileModel(d)
	params := app_profile_service.NewAppProfileServiceCreateAppProfileParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.AppProfile.AppProfileServiceCreateAppProfile(params, nil)
	if err != nil {
		log.Printf("[TRACE] AppProfileService.AppProfileServiceCreateAppProfile error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("AppProfileService.AppProfileServiceCreateAppProfile error: %s", err)...)
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

	// note, we need to set the ID in any case, even GetByName endpoint seems to require the ID
	// but doesn't return any error if it's not set.
	d.SetId(responseData.ObjectID)

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if diags := ReadAppProfile(ctx, d, m); diags.HasError() {
		return diags
	}

	return diags
}

func ReadAppProfile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if _, isSet := d.GetOk("name"); isSet {
		return GetAppProfileByName(ctx, d, m)
	}
	return GetAppProfileByID(ctx, d, m)
}

func GetAppProfileByID(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := app_profile_service.NewAppProfileServiceGetAppProfileParams()

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

	resp, err := client.AppProfile.AppProfileServiceGetAppProfile(params, nil)
	if err != nil {
		log.Printf("[TRACE] AppProfileService.AppProfileServiceGetAppProfile error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("AppProfileService.AppProfileServiceGetAppProfile error: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	zschema.SetAppProfileReadResourceData(d, respModel)
	d.SetId(respModel.ID)

	return diags
}

func GetAppProfileByName(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := app_profile_service.NewAppProfileServiceGetAppProfileByNameParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	nameVal, nameIsSet := d.GetOk("name")
	if nameIsSet {
		params.Name = nameVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: name")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.AppProfile.AppProfileServiceGetAppProfileByName(params, nil)
	if err != nil {
		log.Printf("[TRACE] AppProfileService.AppProfileServiceGetAppProfileByName error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("AppProfileService.AppProfileServiceGetAppProfileByName error: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	zschema.SetAppProfileReadResourceData(d, respModel)
	d.SetId(respModel.ID)

	return diags
}

func UpdateAppProfile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	params := app_profile_service.NewAppProfileServiceUpdateAppProfileParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.AppProfileModel(d))
	// AppProfileServiceUpdateAppProfileBody

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.AppProfile.AppProfileServiceUpdateAppProfile(params, nil)
	if err != nil {
		log.Printf("[TRACE] AppProfileService.AppProfileServiceUpdateAppProfile error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("AppProfileService.AppProfileServiceUpdateAppProfile error: %s", err)...)
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

	// note, we need to set the ID in any case, even GetByName endpoint seems to require items
	// but doesn't return any error if it's not set.
	d.SetId(responseData.ObjectID)

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := ReadAppProfile(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	return diags
}

func DeleteAppProfile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := app_profile_service.NewAppProfileServiceDeleteAppProfileParams()

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

	_, err := client.AppProfile.AppProfileServiceDeleteAppProfile(params, nil)
	if err != nil {
		log.Printf("[TRACE] AppProfileService.AppProfileServiceDeleteAppProfile error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("AppProfileService.AppProfileServiceDeleteAppProfile error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
