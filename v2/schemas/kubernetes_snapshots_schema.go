package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func KubernetesSnapshotsModel(d *schema.ResourceData) *models.SnapshotGetResponse {
	return SnapshotGetResponseModel(d)
}

func KubernetesSnapshotsSchema() map[string]*schema.Schema {
	s := SnapshotGetResponseSchema()
	s["snapshot_timeout"] = &schema.Schema{
		Description: `Timeout in seconds for the snapshot operation (create/restore)`,
		Type:        schema.TypeInt,
		Optional:    true,
	}
	s["x_request_id"] = &schema.Schema{
		Description: `User-Agent specified id to track a request`,
		Type:        schema.TypeString,
		Optional:    true,
	}
	return s
}
