package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/volume_instance"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

func VolumeInstanceResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateVolumeInstance,
		ReadContext:   ReadVolumeInstance,
		UpdateContext: UpdateVolumeInstance,
		DeleteContext: DeleteVolumeInstance,
		Schema:        zschema.VolumeInstance(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func VolumeInstanceDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: ReadVolumeInstance,
		Schema:      zschema.VolumeInstance(),
	}
}

func CreateVolumeInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.VolumeInstanceModel(d)
	params := config.CreateParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.VolumeInstance.Create(params, nil)
	if err != nil {
		log.Printf("[TRACE] volume instance create error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("volume instance create error: %s", err)...)
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
	if errs := readVolumeInstanceByName(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func ReadVolumeInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if _, isSet := d.GetOk("name"); isSet {
		return readVolumeInstanceByName(ctx, d, m)
	}
	return readVolumeInstanceByID(ctx, d, m)
}

func readVolumeInstanceByID(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.GetByIDParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	id, isSet := d.GetOk("id")
	if isSet {
		params.ID = id.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.VolumeInstance.GetByID(params, nil)
	if err != nil {
		log.Printf("[TRACE] volume instance error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("volume instance read error: %s", err)...)
		return diags
	}

	volume := resp.GetPayload()
	zschema.SetVolumeInstanceResourceData(d, volume)
	d.SetId(volume.ID)

	return diags
}

func readVolumeInstanceByName(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.GetByNameParams()

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

	resp, err := client.VolumeInstance.GetByName(params, nil)
	if err != nil {
		log.Printf("[TRACE] volume instance error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("volume instance read error: %s", err)...)
		return diags
	}

	volume := resp.GetPayload()
	zschema.SetVolumeInstanceResourceData(d, volume)
	d.SetId(volume.ID)

	return diags
}

func UpdateVolumeInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.UpdateParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.VolumeInstanceModel(d))

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.VolumeInstance.Update(params, nil)
	if err != nil {
		log.Printf("[TRACE] volume instance update error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("volume instance update error: %s", err)...)
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
	if errs := readVolumeInstanceByName(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func DeleteVolumeInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.DeleteParams()

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

	_, err := client.VolumeInstance.Delete(params, nil)
	if err != nil {
		log.Printf("[TRACE] volume instance delete error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("volume instance delete error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}

func GetVolumeInstanceByID(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.GetByIDParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		params.ID = idVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.VolumeInstance.GetByID(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	volume := resp.GetPayload()
	zschema.SetVolumeInstanceResourceData(d, volume)
	d.SetId(volume.ID)

	return diags
}
