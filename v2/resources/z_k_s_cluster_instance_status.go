package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	"github.com/zededa/terraform-provider-zedcloud/v2/client/z_k_s_cluster_instance"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

func ZksInstanceStatusResource() *schema.Resource {
	statusSchema := zschema.ZksInstanceStatusSchema()

	statusSchema["x_request_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}

	return &schema.Resource{
		ReadContext:   GetZKSInstanceStatus,
		CreateContext: schema.NoopContext,
		UpdateContext: schema.NoopContext,
		DeleteContext: schema.NoopContext,
		Schema:        statusSchema,
	}
}

func ZksInstanceStatusDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: GetZKSInstanceStatus,
		Schema:      zschema.ZksInstanceStatusSchema(),
	}
}

func GetZKSInstanceStatus(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	zksIdVal, zksIdIsSet := d.GetOk("id")
	if !zksIdIsSet {
		return diag.Errorf("missing client parameter: id")
	}
	zksID := zksIdVal.(string)

	client := m.(*api_client.ZedcloudAPI)

	params := z_k_s_cluster_instance.NewGetZKSInstancesStatusParams()
	params.Context = ctx
	params.ZksID = &zksID

	if xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id"); xRequestIdIsSet {
		val := xRequestIdVal.(string)
		params.XRequestID = &val
	}

	// The client parameter struct does not seem to expose 'zksid' filtering properties
	// despite API documentation suggesting support. We fetch all and filter client-side.
	resp, err := client.ZksClusterInstance.GetZKSInstancesStatus(params, nil)
	if err != nil {
		return diag.Errorf("GetZKSInstancesStatus error: %s", err)
	}

	if resp.Payload == nil || resp.Payload.List == nil {
		return diag.Errorf("ZKS Instance status list is empty")
	}

	var foundStatus *models.ZksInstanceStatus
	for _, status := range resp.Payload.List {
		if status.ID == zksID {
			foundStatus = status
			break
		}
	}

	if foundStatus == nil {
		return diag.Errorf("ZKS Instance status not found for id: %s", zksID)
	}

	zschema.SetZksInstanceStatusResourceData(d, foundStatus)
	d.SetId(foundStatus.ID)

	return diags
}
