package resources

import (
	"context"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	kubernetes_snapshots "github.com/zededa/terraform-provider-zedcloud/v2/client/kubernetes_snapshots"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

/*
KubernetesSnapshots kubernetes snapshots API
*/

func KubernetesSnapshots() *schema.Resource {
	return &schema.Resource{
		CreateContext: KubernetesSnapshots_CreateSnapshot,
		ReadContext:   KubernetesSnapshots_ListSnapshots,
		DeleteContext: KubernetesSnapshots_DeleteSnapshot,
		UpdateContext: KubernetesSnapshots_RestoreSnapshot,
		Schema:        zschema.KubernetesSnapshotsSchema(),
	}
}

func DataResourceKubernetesSnapshots() *schema.Resource {
	return &schema.Resource{
		ReadContext: KubernetesSnapshots_ListSnapshots,
		Schema:      zschema.KubernetesSnapshotsSchema(),
	}
}

func KubernetesSnapshots_ListSnapshots(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := kubernetes_snapshots.NewKubernetesSnapshotsListSnapshotsParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	clusterIdVal, clusterIdIsSet := d.GetOk("cluster_id")
	if clusterIdIsSet {
		params.ClusterID = clusterIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: clusterId")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.KubernetesSnapshots.KubernetesSnapshotsListSnapshots(params, nil)
	if err != nil {
		log.Printf("[TRACE] KubernetesSnapshots.KubernetesSnapshotsListSnapshots error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("KubernetesSnapshots.KubernetesSnapshotsListSnapshots error: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	if len(respModel.Snapshots) == 0 {
		return append(diags, diag.Errorf("no snapshots found")...)
	}
	// limit output to a single result
	result := respModel.Snapshots[0]

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	d.SetId(result.Name)
	zschema.SetSnapshotGetResponseResourceData(d, result)
	return diags
}

func KubernetesSnapshots_CreateSnapshot(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := kubernetes_snapshots.NewKubernetesSnapshotsCreateSnapshotParams()

	clusterIdVal, clusterIdIsSet := d.GetOk("cluster_id")
	if clusterIdIsSet {
		params.ClusterID = clusterIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: clusterId")...)
		return diags
	}

	body := kubernetes_snapshots.KubernetesSnapshotsCreateSnapshotBody{}
	if v, ok := d.GetOk("snapshot_timeout"); ok {
		body.SnapshotTimeout = int32(v.(int))
	}
	params.SetBody(body)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.KubernetesSnapshots.KubernetesSnapshotsCreateSnapshot(params, nil)
	if err != nil {
		log.Printf("[TRACE] KubernetesSnapshots.KubernetesSnapshotsCreateSnapshot error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("KubernetesSnapshots.KubernetesSnapshotsCreateSnapshot error: %s", err)...)
		return diags
	}

	responseData := resp.GetPayload()
	if responseData != nil && responseData.OperationID != "" {
		d.SetId(responseData.OperationID)
	}

	return diags
}

func KubernetesSnapshots_RestoreSnapshot(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := kubernetes_snapshots.NewKubernetesSnapshotsRestoreSnapshotParams()

	clusterIdVal, clusterIdIsSet := d.GetOk("cluster_id")
	if clusterIdIsSet {
		params.ClusterID = clusterIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: clusterId")...)
		return diags
	}

	nameVal, nameIsSet := d.GetOk("name")
	if nameIsSet {
		params.Name = nameVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: name")...)
		return diags
	}

	body := kubernetes_snapshots.KubernetesSnapshotsRestoreSnapshotBody{}
	if v, ok := d.GetOk("snapshot_timeout"); ok {
		body.SnapshotTimeout = int32(v.(int))
	}
	params.SetBody(body)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.KubernetesSnapshots.KubernetesSnapshotsRestoreSnapshot(params, nil)
	if err != nil {
		log.Printf("[TRACE] KubernetesSnapshots.KubernetesSnapshotsRestoreSnapshot error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("KubernetesSnapshots.KubernetesSnapshotsRestoreSnapshot error: %s", err)...)
		return diags
	}

	responseData := resp.GetPayload()
	if responseData != nil && responseData.OperationID != "" {
		d.SetId(responseData.OperationID)
	}

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := KubernetesSnapshots_ListSnapshots(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	return diags
}

func KubernetesSnapshots_DeleteSnapshot(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := kubernetes_snapshots.NewKubernetesSnapshotsDeleteSnapshotParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	clusterIdVal, clusterIdIsSet := d.GetOk("cluster_id")
	if clusterIdIsSet {
		params.ClusterID = clusterIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: clusterId")...)
		return diags
	}

	nameVal, nameIsSet := d.GetOk("name")
	if nameIsSet {
		params.Name = nameVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: name")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	_, err := client.KubernetesSnapshots.KubernetesSnapshotsDeleteSnapshot(params, nil)
	if err != nil {
		log.Printf("[TRACE] KubernetesSnapshots.KubernetesSnapshotsDeleteSnapshot error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("KubernetesSnapshots.KubernetesSnapshotsDeleteSnapshot error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
