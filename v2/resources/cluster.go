package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	cluster "github.com/zededa/terraform-provider-zedcloud/v2/client/cluster"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

/*
 edge node cluster configuration API
*/

func ClusterResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateCluster,
		UpdateContext: UpdateCluster,
		ReadContext: GetCluster,
		DeleteContext: DeleteCluster,
		Schema: zschema.ClusterSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func ClusterDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: GetCluster,
		Schema: zschema.ClusterSchema(),
	}
}

func GetCluster(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if _, nameIsSet := d.GetOk("name"); nameIsSet {
		return GetClusterByName(ctx, d, m)
	}
	
	return GetClusterByID(ctx, d, m)
}

func GetClusterByID(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := cluster.NewGetParams()

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

	resp, err := client.Cluster.GetCluster(params, nil)
	if err != nil {
		log.Printf("[TRACE] edge node cluster read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("edge node cluster read error: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	zschema.SetClusterResourceData(d, respModel)

	return diags
}

func GetClusterByName(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := cluster.NewGetByNameParams()

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

	resp, err := client.Cluster.GetClusterByName(params, nil)
	if err != nil {
		log.Printf("[TRACE] edge node cluster read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("edge node cluster read error: %s", err)...)
	}

	respModel := resp.GetPayload()
	zschema.SetClusterResourceData(d, respModel)

	return diags
}

func CreateCluster(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.ClusterModel(d)
	params := cluster.NewCreateParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.Cluster.CreateCluster(params, nil)
	if err != nil {
		log.Printf("[TRACE] edge node cluster create error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("edge node cluster create error: %s", err)...)
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
	if errs := GetCluster(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func UpdateCluster(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	params := cluster.NewUpdateParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.ClusterModel(d))
	// UpdateClusterBody

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
	resp, err := client.Cluster.UpdateCluster(params, nil)
	if err != nil {
		log.Printf("[TRACE] edge node cluster update error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("edge node cluster update error: %s", err)...)
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
	if errs := GetCluster(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	return diags
}

func DeleteCluster(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := cluster.NewDeleteParams()

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

	_, err := client.Cluster.DeleteCluster(params, nil)
	if err != nil {
		log.Printf("[TRACE] edge node cluster delete error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("edge node cluster error: %s", err)...)
		return diags
	}

	d.SetId("")
	
	return diags
}
