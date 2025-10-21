package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/application_instance"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

func ApplicationInstanceResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateApplicationInstance,
		ReadContext:   ReadApplicationInstance,
		UpdateContext: UpdateApplicationInstance,
		DeleteContext: DeleteApplicationInstance,
		Schema:        zschema.ApplicationInstance(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func ApplicationInstanceDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: ReadApplicationInstance,
		Schema:      zschema.ApplicationInstance(),
	}
}

func CreateApplicationInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.ApplicationInstanceModel(d)
	params := config.CreateParams()
	params.SetBody(model)

	_, isDevIDSet := d.GetOk("device_id")
	_, isEdgeNodeClusterSet := d.GetOk("edge_node_cluster")
	if !isDevIDSet && !isEdgeNodeClusterSet {
		diags = append(diags, diag.Errorf("either device_id or edge_node_cluster has to be specified")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.ApplicationInstance.Create(params, nil)
	if err != nil {
		log.Printf("[TRACE] edge application instance create error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("edge application instance create error: %s", err)...)
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
	if errs := readApplicationInstanceByName(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func ReadApplicationInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if _, isSet := d.GetOk("name"); isSet {
		return readApplicationInstanceByName(ctx, d, m)
	}
	return readApplicationInstanceByID(ctx, d, m)
}

func readApplicationInstanceByID(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	resp, err := client.ApplicationInstance.GetByID(params, nil)
	if err != nil {
		log.Printf("[TRACE] edge application instance read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("edge application instance read error: %s", err)...)
		return diags
	}

	appInstance := resp.GetPayload()
	zschema.SetApplicationInstanceResourceData(d, appInstance)
	d.SetId(appInstance.ID)

	return diags

}

func readApplicationInstanceByName(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	resp, err := client.ApplicationInstance.GetByName(params, nil)
	if err != nil {
		log.Printf("[TRACE] edge application instance read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("edge application instance read error: %s", err)...)
		return diags
	}

	appInstance := resp.GetPayload()
	zschema.SetApplicationInstanceResourceData(d, appInstance)
	d.SetId(appInstance.ID)

	return diags

}

func UpdateApplicationInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.UpdateParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.ApplicationInstanceModel(d))

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.ApplicationInstance.Update(params, nil)
	if err != nil {
		log.Printf("[TRACE] edge application instance update error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("edge application instance update error: %s", err)...)
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
	if errs := readApplicationInstanceByName(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func DeleteApplicationInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	_, err := client.ApplicationInstance.Delete(params, nil)
	if err != nil {
		log.Printf("[TRACE] edge application instance delete error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("edge application instance delete error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
