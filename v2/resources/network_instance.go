package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/network_instance"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

func NetworkInstanceResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateNetworkInstance,
		ReadContext:   ReadNetworkInstance,
		UpdateContext: UpdateNetworkInstance,
		DeleteContext: DeleteNetworkInstance,
		Schema:        zschema.NetworkInstance(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func NetworkInstanceDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: ReadNetworkInstance,
		Schema:      zschema.NetworkInstance(),
	}
}

func CreateNetworkInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.NetworkInstanceModel(d)
	params := config.CreateNetworkInstanceParams()
	params.SetBody(model)

	_, isDevIDSet := d.GetOk("device_id")
	_, isEdgeNodeClusterSet := d.GetOk("edge_node_cluster")
	if !isDevIDSet && !isEdgeNodeClusterSet {
		diags = append(diags, diag.Errorf("either device_id or edge_node_cluster has to be specified")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.NetworkInstance.Create(params, nil)
	if err != nil {
		log.Printf("[TRACE] network instance create error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("network instance create error: %s", err)...)
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
	if errs := ReadNetworkInstance(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func ReadNetworkInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	_, isSet := d.GetOk("name")
	if isSet {
		return readNetworkInstanceByName(ctx, d, m)
	}
	return readNetworkInstanceByID(ctx, d, m)
}

func readNetworkInstanceByID(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	resp, err := client.NetworkInstance.GetByID(params, nil)
	if err != nil {
		log.Printf("[TRACE] network instance read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("network instance read error: %s", err)...)
		return diags
	}

	networkInstance := resp.GetPayload()
	zschema.SetNetworkInstanceResourceData(d, networkInstance)
	d.SetId(networkInstance.ID)

	return diags
}

func readNetworkInstanceByName(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	resp, err := client.NetworkInstance.GetByName(params, nil)
	if err != nil {
		log.Printf("[TRACE] network instance read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("network instance read error: %s", err)...)
		return diags
	}

	networkInstance := resp.GetPayload()
	zschema.SetNetworkInstanceResourceData(d, networkInstance)
	d.SetId(networkInstance.ID)

	return diags
}

func UpdateNetworkInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.NetworkInstanceModel(d)
	params := config.NewEdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams()
	params.SetBody(model)

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

	// makes a bulk update for all properties that were changed
	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.NetworkInstance.EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstance(params, nil)
	if err != nil {
		log.Printf("[TRACE] network instance update error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("network instance update error: %s", err)...)
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
	if errs := ReadNetworkInstance(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func DeleteNetworkInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.DeleteNetworkInstanceParams()

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

	_, err := client.NetworkInstance.Delete(params, nil)
	if err != nil {
		log.Printf("[TRACE] network instance delete error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("network instance delete error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
