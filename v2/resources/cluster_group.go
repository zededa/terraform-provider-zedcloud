package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	cluster_groups "github.com/zededa/terraform-provider-zedcloud/v2/client/cluster_group"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

/*
ClusterGroups cluster groups API
*/

func ClusterGroupResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateClusterGroup,
		UpdateContext: schema.NoopContext,
		ReadContext:   schema.NoopContext,
		DeleteContext: DeleteClusterGroup,
		Schema:        zschema.ClusterGroupSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func ClusterGroupDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: schema.NoopContext,
		Schema:      zschema.ClusterGroupSchema(),
	}
}

func CreateClusterGroup(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.ClusterGroupModel(d)
	params := cluster_groups.NewCreateClusterGroupParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, _, err := client.ClusterGroup.CreateClusterGroup(params, nil)
	if err != nil {
		log.Printf("[TRACE] CreateClusterGroup error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("CreateClusterGroup error: %s", err)...)
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

	return diags
}

func DeleteClusterGroup(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := cluster_groups.NewDeleteClusterGroupParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	clusterGroupIdVal, clusterGroupIdIsSet := d.GetOk("id")
	if clusterGroupIdIsSet {
		params.ClusterGroupID = clusterGroupIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	_, err := client.ClusterGroup.DeleteClusterGroup(params, nil)
	if err != nil {
		log.Printf("[TRACE] DeleteClusterGroup error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("DeleteClusterGroup error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
