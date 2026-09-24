package resources

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	cluster "github.com/zededa/terraform-provider-zedcloud/v2/client/cluster"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

/*
 edge node cluster configuration API
*/

func ClusterResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateCluster,
		UpdateContext: UpdateCluster,
		ReadContext:   GetCluster,
		DeleteContext: DeleteCluster,
		Schema:        zschema.ClusterSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func ClusterDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: GetCluster,
		Schema:      zschema.ClusterSchema(),
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

	if errs := readClusterUpgradeStatus(ctx, d, m); errs.HasError() {
		return append(diags, errs...)
	}

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

	if respModel != nil && respModel.ID != "" {
		d.SetId(respModel.ID)
	}

	if errs := readClusterUpgradeStatus(ctx, d, m); errs.HasError() {
		return append(diags, errs...)
	}

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

	// Nodes are only cluster members once the cluster object exists, so the
	// EVE-OS rollout has to come after the create, not as part of its body.
	if errs := setClusterBaseImage(ctx, d, m); errs.HasError() {
		return append(diags, errs...)
	}

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

	clusterModel := sanitizeClusterNodes(d)
	params.SetBody(clusterModel)

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

	if errs := setClusterBaseImage(ctx, d, m); errs.HasError() {
		return append(diags, errs...)
	}

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := GetCluster(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	return diags
}

// setClusterBaseImage triggers a cluster-scoped EVE-OS upgrade when the
// configured base_image differs from what the controller is already rolling
// out.
//
// This mirrors setBaseImage() on the node resource: compare first, call second.
// The comparison matters more here than it does for a node, because
// PUT /v1/cluster/id/{id}/upgrade is not idempotent in the way Terraform wants
// -- re-requesting an image while nodes are still moving is answered with 409.
func setClusterBaseImage(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	desired := zschema.ClusterBaseImageModel(d)
	if desired == nil || desired.ImageName == nil || *desired.ImageName == "" {
		// base_image absent from the config. Terraform does not own the
		// cluster's EVE-OS version, so leave whatever is there alone.
		return diags
	}

	id, idIsSet := d.GetOk("id")
	if !idIsSet || id.(string) == "" {
		return append(diags, diag.Errorf("missing client parameter: id")...)
	}
	clusterID := id.(string)

	client := m.(*api_client.ZedcloudAPI)

	current, statusDiags := getClusterUpgradeStatus(d, client, clusterID)
	if statusDiags.HasError() {
		return append(diags, statusDiags...)
	}
	if current != nil {
		inProgress := false
		alreadyRequested := false
		for _, node := range current.Nodes {
			if node == nil {
				continue
			}
			if node.UpgradeableEveOs == *desired.ImageName {
				alreadyRequested = true
			}
			if node.Status != nil && *node.Status == models.EdgeNodeClusterUpgradeStatusSTATUSINPROGRESS {
				inProgress = true
			}
		}

		if alreadyRequested && !inProgress {
			// The controller already has this image; nothing to ask for.
			return diags
		}
		if inProgress {
			if alreadyRequested {
				// Our image, still rolling. Re-requesting it would only earn a
				// 409, and the rollout advances on its own as each node reports
				// in, so report progress and let it run.
				return append(diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  "edge node cluster EVE-OS upgrade still in progress",
					Detail: fmt.Sprintf(
						"cluster %s is still rolling out %q. The controller upgrades member nodes one at a time; "+
							"re-run a plan later, or watch the upgrade_status attribute, to see it finish.",
						clusterID, *desired.ImageName),
				})
			}
			return append(diags, diag.Errorf(
				"cluster %s has an EVE-OS upgrade in progress, so it cannot be asked for %q yet. "+
					"Wait for the running rollout to reach STATUS_COMPLETED or STATUS_FAILED on every node "+
					"(see the upgrade_status attribute) and apply again.",
				clusterID, *desired.ImageName)...)
		}
	}

	params := cluster.NewUpgradeParams()
	params.ID = clusterID
	params.SetBody(desired)
	if xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id"); xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	resp, err := client.Cluster.UpgradeCluster(params, nil)
	if err != nil {
		log.Printf("[TRACE] edge node cluster upgrade error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			return append(diags, ds...)
		}
		return append(diags, diag.Errorf("edge node cluster upgrade error: %s", err)...)
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, respErr := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if respErr.ErrorCode != nil && *respErr.ErrorCode == models.ErrorCodeSuccess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(respErr.Details))...)
		}
	}

	return diags
}

// readClusterUpgradeStatus refreshes upgrade_status and base_image in state.
//
// A cluster that has never been upgraded has no status rows at all, and older
// controllers do not serve the endpoint. Neither is a reason to fail a read, so
// both degrade to leaving the attributes untouched.
func readClusterUpgradeStatus(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	id, idIsSet := d.GetOk("id")
	if !idIsSet || id.(string) == "" {
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)
	status, statusDiags := getClusterUpgradeStatus(d, client, id.(string))
	if statusDiags.HasError() {
		return statusDiags
	}

	zschema.SetClusterUpgradeResourceData(d, status)

	return diags
}

func getClusterUpgradeStatus(
	d *schema.ResourceData,
	client *api_client.ZedcloudAPI,
	clusterID string,
) (*models.EdgeNodeClusterUpgradeStatusResp, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := cluster.NewUpgradeStatusParams()
	params.ID = clusterID
	if xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id"); xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	resp, err := client.Cluster.GetClusterUpgradeStatus(params, nil)
	if err != nil {
		// A cluster with no rollout history, or a controller that predates the
		// endpoint, answers 404. That is not an error for a read.
		if isStatusNotFound(err) {
			return nil, diags
		}
		log.Printf("[TRACE] edge node cluster upgrade status read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			return nil, append(diags, ds...)
		}
		return nil, append(diags, diag.Errorf("edge node cluster upgrade status read error: %s", err)...)
	}

	return resp.GetPayload(), diags
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
		if isStatusNotFound(err) {
			d.SetId("")
			return diags
		}
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

func sanitizeClusterNodes(d *schema.ResourceData) *models.Cluster {
	clusterModel := zschema.ClusterModel(d)

	if d.HasChange("nodes") {
		oldRaw, _ := d.GetChange("nodes")

		oldList, _ := oldRaw.([]interface{})

		oldIDs := make(map[string]struct{})
		for _, o := range oldList {
			m, _ := o.(map[string]interface{})
			if id, ok := m["id"].(string); ok && id != "" {
				oldIDs[id] = struct{}{}
			}
		}

		for _, n := range clusterModel.Nodes {
			// If a node has no old ID (new replacement), clear its prefix
			if _, existed := oldIDs[n.ID]; !existed {
				n.ClusterPrefix = ""
			}
		}
	}

	return clusterModel
}
