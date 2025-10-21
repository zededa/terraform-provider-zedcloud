package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	hardware_model "github.com/zededa/terraform-provider-zedcloud/v2/client/hardware_model"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

/*
HardwareModel hardware model API
*/

func HardwareBrandResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateHardwareBrand,
		DeleteContext: DeleteHardwareBrand,
		ReadContext: GetHardwareBrand,
		UpdateContext: UpdateHardwareBrand,
		Schema: zschema.SysBrandSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func HardwareBrandDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: GetHardwareBrand,
		Schema: zschema.SysBrandSchema(),
	}
}

func GetHardwareBrand(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
if _, isSet := d.GetOk("name"); isSet {
		return GetHardwareBrandByName(ctx, d, m)
	}
	return GetHardwareBrandById(ctx, d, m)
}

func GetHardwareBrandById(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := hardware_model.NewHardwareModelGetHardwareBrandParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	enterpriseIdVal, enterpriseIdIsSet := d.GetOk("enterprise_id")
	if enterpriseIdIsSet {
		params.EnterpriseID = enterpriseIdVal.(*string)
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

	resp, err := client.HardwareModel.HardwareModelGetHardwareBrand(params, nil)
	if err != nil {
		log.Printf("[TRACE] brand read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("brand read error: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	zschema.SetSysBrandResourceData(d, respModel)
	d.SetId(respModel.ID)

	return diags
}

func GetHardwareBrandByName(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := hardware_model.NewHardwareModelGetHardwareBrandByNameParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	enterpriseIdVal, enterpriseIdIsSet := d.GetOk("enterprise_id")
	if enterpriseIdIsSet {
		params.EnterpriseID = enterpriseIdVal.(*string)
	}

	nameVal, nameIsSet := d.GetOk("name")
	if nameIsSet {
		params.Name = nameVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: name")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.HardwareModel.HardwareModelGetHardwareBrandByName(params, nil)
	if err != nil {
		log.Printf("[TRACE] brand read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("brand read error: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	zschema.SetSysBrandResourceData(d, respModel)
	d.SetId(respModel.ID)

	return diags
}

func CreateHardwareBrand(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.SysBrandModel(d)
	params := hardware_model.NewHardwareModelCreateHardwareBrandParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.HardwareModel.HardwareModelCreateHardwareBrand(params, nil)
	if err != nil {
		log.Printf("[TRACE] brand create error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("brand create error: %s", err)...)
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

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := GetHardwareBrandByName(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func UpdateHardwareBrand(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	params := hardware_model.NewHardwareModelUpdateHardwareBrandParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.SysBrandModel(d))
	// HardwareModelUpdateHardwareBrandBody

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
	resp, err := client.HardwareModel.HardwareModelUpdateHardwareBrand(params, nil)
	if err != nil {
		log.Printf("[TRACE] brand update error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("brand update error: %s", err)...)
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
	if errs := GetHardwareBrandByName(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func DeleteHardwareBrand(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := hardware_model.NewHardwareModelDeleteHardwareBrandParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	enterpriseIdVal, enterpriseIdIsSet := d.GetOk("enterprise_id")
	if enterpriseIdIsSet {
		params.EnterpriseID = enterpriseIdVal.(*string)
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

	_, err := client.HardwareModel.HardwareModelDeleteHardwareBrand(params, nil)
	if err != nil {
		log.Printf("[TRACE] brand delete error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("brand delete error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}