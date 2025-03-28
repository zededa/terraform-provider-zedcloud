package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	profile_deployment_service "github.com/zededa/terraform-provider-zedcloud/v2/client/profile_deployment_service"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

/*
ProfileDeploymentService profile deployment service API
*/

func ProfileDeploymentResource() *schema.Resource {
	return &schema.Resource{

		CreateContext: CreateProfileDeployment,
		DeleteContext: DeleteProfileDeployment,
		ReadContext:   ReadProfileDeployment,
		UpdateContext: UpdateProfileDeployment,
		Schema:        zschema.ProfileDeploymentSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func ProfileDeploymentDataResource() *schema.Resource {
	return &schema.Resource{
		Schema:      zschema.ProfileDeploymentSchema(),
		ReadContext: ReadProfileDeployment,
	}
}

func ReadProfileDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if _, isSet := d.GetOk("name"); isSet {
		return GetProfileDeploymentByName(ctx, d, m)
	}
	return GetProfileDeploymentByID(ctx, d, m)
}

func GetProfileDeploymentByID(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := profile_deployment_service.NewProfileDeploymentServiceGetProfileDeploymentParams()

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

	resp, err := client.ProfileDeployment.ProfileDeploymentServiceGetProfileDeployment(params, nil)
	if err != nil {
		log.Printf("[TRACE] ProfileDeploymentService.ProfileDeploymentServiceGetProfileDeployment error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("ProfileDeploymentService.ProfileDeploymentServiceGetProfileDeployment error: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	zschema.SetProfileDeploymentResourceData(d, respModel)
	d.SetId(respModel.ID)
	return diags
}

func GetProfileDeploymentByName(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := profile_deployment_service.NewProfileDeploymentServiceGetProfileDeploymentByNameParams()

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

	resp, err := client.ProfileDeployment.ProfileDeploymentServiceGetProfileDeploymentByName(params, nil)
	if err != nil {
		log.Printf("[TRACE] ProfileDeploymentService.ProfileDeploymentServiceGetProfileDeploymentByName error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("ProfileDeploymentService.ProfileDeploymentServiceGetProfileDeploymentByName error: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	zschema.SetProfileDeploymentResourceData(d, respModel)
	d.SetId(respModel.ID)

	return diags
}

func CreateProfileDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.ProfileDeploymentModel(d)
	params := profile_deployment_service.NewProfileDeploymentServiceCreateProfileDeploymentParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.ProfileDeployment.ProfileDeploymentServiceCreateProfileDeployment(params, nil)
	if err != nil {
		log.Printf("[TRACE] ProfileDeploymentService.ProfileDeploymentServiceCreateProfileDeployment error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("ProfileDeploymentService.ProfileDeploymentServiceCreateProfileDeployment error: %s", err)...)
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
	if errs := ReadProfileDeployment(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	return diags
}

func UpdateProfileDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	params := profile_deployment_service.NewProfileDeploymentServiceUpdateProfileDeploymentParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.ProfileDeploymentModel(d))
	// ProfileDeploymentServiceUpdateProfileDeploymentBody

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
	resp, err := client.ProfileDeployment.ProfileDeploymentServiceUpdateProfileDeployment(params, nil)
	if err != nil {
		log.Printf("[TRACE] ProfileDeploymentService.ProfileDeploymentServiceUpdateProfileDeployment error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("ProfileDeploymentService.ProfileDeploymentServiceUpdateProfileDeployment error: %s", err)...)
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
	if errs := ReadProfileDeployment(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	return diags
}

func DeleteProfileDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := profile_deployment_service.NewProfileDeploymentServiceDeleteProfileDeploymentParams()

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

	_, err := client.ProfileDeployment.ProfileDeploymentServiceDeleteProfileDeployment(params, nil)
	if err != nil {
		log.Printf("[TRACE] ProfileDeploymentService.ProfileDeploymentServiceDeleteProfileDeployment error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("ProfileDeploymentService.ProfileDeploymentServiceDeleteProfileDeployment error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
