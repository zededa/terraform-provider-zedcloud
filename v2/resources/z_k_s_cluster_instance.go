package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/client/z_k_s_cluster_instance"

	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

/*
ZksClusterInstances z k s cluster instances API
*/

func ZksInstanceResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateZKSInstance,
		ReadContext:   GetZKSInstance,
		UpdateContext: UpdateZKSInstance,
		DeleteContext: DeleteZKSInstance,
		Schema:        zschema.ZksInstanceSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func ZksInstanceDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: GetZKSInstance,
		Schema:      zschema.ZksInstanceSchema(),
	}
}

func CreateZKSInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if _, ok := d.GetOk("nodes"); !ok {
		return append(diags, diag.Errorf("nodes are required for ZKS instance creation")...)
	}

	model := zschema.ZksInstanceModel(d)
	params := z_k_s_cluster_instance.NewCreateZKSInstanceParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, _, err := client.ZksClusterInstance.CreateZKSInstance(params, nil)
	if err != nil {
		log.Printf("[TRACE] CreateZKSInstance error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("CreateZKSInstance error: %s", err)...)
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
	if errs := GetZKSInstance(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	return diags
}

func GetZKSInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var zksInstance *models.ZksInstance
	var diags diag.Diagnostics

	if _, isSet := d.GetOk("name"); isSet {
		zksInstance, diags = getZKSInstanceByName(ctx, d, m)
	} else if _, isSet := d.GetOk("id"); isSet {
		zksInstance, diags = getZKSInstanceById(ctx, d, m)
	}

	if diags.HasError() {
		return diags
	}

	zschema.SetZksInstanceResourceData(d, zksInstance)
	d.SetId(zksInstance.ID)

	return diags
}

func getZKSInstanceById(ctx context.Context, d *schema.ResourceData, m interface{},
) (*models.ZksInstance, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := z_k_s_cluster_instance.NewGetZKSInstanceParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	zksIdVal, zksIdIsSet := d.GetOk("id")
	if zksIdIsSet {
		params.ZksID = zksIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return nil, diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.ZksClusterInstance.GetZKSInstance(params, nil)
	if err != nil {
		log.Printf("[TRACE] GetZKSInstance error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return nil, diags
		}

		diags = append(diags, diag.Errorf("GetZKSInstance error: %s", err)...)
		return nil, diags
	}

	respModel := resp.GetPayload()
	zschema.SetZksInstanceResourceData(d, respModel)

	return respModel, diags
}

func getZKSInstanceByName(ctx context.Context, d *schema.ResourceData, m interface{},
) (*models.ZksInstance, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := z_k_s_cluster_instance.NewGetZKSInstanceByNameParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	zksNameVal, zksNameIsSet := d.GetOk("name")
	if zksNameIsSet {
		params.ZksName = zksNameVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: name")...)
		return nil, diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.ZksClusterInstance.GetZKSInstanceByName(params, nil)
	if err != nil {
		log.Printf("[TRACE] GetZKSInstanceByName error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return nil, diags
		}

		diags = append(diags, diag.Errorf("GetZKSInstanceByName error: %s", err)...)
		return nil, diags
	}

	respModel := resp.GetPayload()
	zschema.SetZksInstanceResourceData(d, respModel)

	return respModel, diags
}

func UpdateZKSInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	updateMetadata := d.HasChange("title") || d.HasChange("description") || d.HasChange("tags")

	client := m.(*api_client.ZedcloudAPI)

	zksIdVal, zksIdIsSet := d.GetOk("id")
	zksID := ""
	if zksIdIsSet {
		zksID = zksIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	if updateMetadata {
		model := zschema.ZksInstanceModel(d)

		params := z_k_s_cluster_instance.NewUpdateZKSInstanceParams()
		params.ZksID = zksID
		params.SetBody(model)

		resp, err := client.ZksClusterInstance.UpdateZKSInstance(params, nil)
		if err != nil {
			log.Printf("[TRACE] UpdateZKSInstance error: %s", spew.Sdump(err))
			if ds, ok := ZsrvResponderToDiags(err); ok {
				diags = append(diags, ds...)
				return diags
			}

			diags = append(diags, diag.Errorf("UpdateZKSInstance error: %s", err)...)
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
	}

	if d.HasChange("nodes") {
		body := z_k_s_cluster_instance.UpdateZKSInstanceNodesBody{}

		if diags = fillUpdateZksInstancesNodesBody(d, &body); diags != nil {
			return diags
		}
		if len(body.Nodes) == 0 {
			return diags
		}

		params := z_k_s_cluster_instance.NewUpdateZKSInstanceNodesParams()
		params.ZksID = zksID
		params.SetBody(body)

		// makes a bulk update for all properties that were changed
		resp, err := client.ZksClusterInstance.UpdateZKSInstanceNodes(params, nil)
		if err != nil {
			log.Printf("[TRACE] UpdateZKSInstance error: %s", spew.Sdump(err))
			if ds, ok := ZsrvResponderToDiags(err); ok {
				diags = append(diags, ds...)
				return diags
			}

			diags = append(diags, diag.Errorf("UpdateZKSInstance error: %s", err)...)
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
	}

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := GetZKSInstance(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	d.Partial(false)

	return diags
}

func fillUpdateZksInstancesNodesBody(d *schema.ResourceData,
	body *z_k_s_cluster_instance.UpdateZKSInstanceNodesBody) diag.Diagnostics {
	oldNodes, newNodes := d.GetChange("nodes")
	if oldNodes != nil && newNodes != nil {
		oldNodesRaw := oldNodes.(*schema.Set).List()
		newNodesRaw := newNodes.(*schema.Set).List()
		if len(oldNodesRaw) != len(newNodesRaw) {
			oldNodeMap := make(map[string]*models.ZksInstanceNode, len(oldNodesRaw))
			for _, nodeRaw := range oldNodesRaw {
				node := nodeRaw.(map[string]interface{})
				oldNodeMap[node["id"].(string)] = &models.ZksInstanceNode{
					ID:               node["id"].(string),
					ClusterInterface: node["cluster_interface"].(string),
				}
			}
			newNodeMap := make(map[string]*models.ZksInstanceNode, len(newNodesRaw))
			for _, nodeRaw := range newNodesRaw {
				node := nodeRaw.(map[string]interface{})
				newNodeMap[node["id"].(string)] = &models.ZksInstanceNode{
					ID:               node["id"].(string),
					ClusterInterface: node["cluster_interface"].(string),
				}
			}
			// some nodes are removed
			if len(oldNodeMap) > len(newNodeMap) {
				body.Action = models.NodeActionNODEACTIONREMOVENODES.Pointer()
				for id, node := range oldNodeMap {
					if _, ok := newNodeMap[id]; !ok {
						body.Nodes = append(body.Nodes, node)
					}
				}
				// some nodes are added
			} else if len(oldNodeMap) < len(newNodeMap) {
				body.Action = models.NodeActionNODEACTIONADDNODES.Pointer()
				for id, node := range newNodeMap {
					if _, ok := oldNodeMap[id]; !ok {
						body.Nodes = append(body.Nodes, node)
					}
				}
			}
		}
	} else {
		return diag.Errorf("nodes cannot be updated in place")
	}

	return nil
}

func DeleteZKSInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := z_k_s_cluster_instance.NewDeleteZKSInstanceParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	zksIdVal, zksIdIsSet := d.GetOk("id")
	if zksIdIsSet {
		params.ZksID = zksIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	_, err := client.ZksClusterInstance.DeleteZKSInstance(params, nil)
	if err != nil {
		log.Printf("[TRACE] DeleteZKSInstance error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("DeleteZKSInstance error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
