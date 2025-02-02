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

func HardwareModelResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateHardwareModel,
		DeleteContext: DeleteHardwareModel,
		ReadContext: GetHardwareModel,
		UpdateContext: UpdateHardwareModel,
		Schema: zschema.SysModelSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func HardwareModelDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: GetHardwareModel,
		Schema: zschema.SysModelSchema(),
	}
}

func GetHardwareModel(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if _, isSet := d.GetOk("name"); isSet {
		return GetHardwareModelByName(ctx, d, m)
	}
	return GetHardwareModelById(ctx, d, m)
}

func GetHardwareModelById(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := hardware_model.NewHardwareModelGetHardwareModelParams()

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

	resp, err := client.HardwareModel.HardwareModelGetHardwareModel(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	respModel := resp.GetPayload()
	zschema.SetSysModelResourceData(d, respModel)
	d.SetId(respModel.ID)

	return diags
}

func GetHardwareModelByName(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := hardware_model.NewHardwareModelGetHardwareModelByNameParams()

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

	resp, err := client.HardwareModel.HardwareModelGetHardwareModelByName(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	respModel := resp.GetPayload()
	zschema.SetSysModelResourceData(d, respModel)
	d.SetId(respModel.ID)

	return diags
}

func CreateHardwareModel(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.SysModelModel(d)
	params := hardware_model.NewHardwareModelCreateHardwareModelParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.HardwareModel.HardwareModelCreateHardwareModel(params, nil)
	if err != nil {
		log.Printf("[TRACE] model create error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("model create error: %s", err)...)
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
	if errs := GetHardwareModelByName(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func UpdateHardwareModel(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	params := hardware_model.NewHardwareModelUpdateHardwareModelParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.SysModelModel(d))
	// HardwareModelUpdateHardwareModelBody

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
	resp, err := client.HardwareModel.HardwareModelUpdateHardwareModel(params, nil)
	if err != nil {
		log.Printf("[TRACE] model update error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("model update error: %s", err)...)
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
	return GetHardwareModelByName(ctx, d, m)
}

func DeleteHardwareModel(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := hardware_model.NewHardwareModelDeleteHardwareModelParams()

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

	_, err := client.HardwareModel.HardwareModelDeleteHardwareModel(params, nil)
	if err != nil {
		log.Printf("[TRACE] model delete error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("model delete error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
