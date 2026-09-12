package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SnapshotGetResponseModel(d *schema.ResourceData) *models.SnapshotGetResponse {
	clusterID, _ := d.Get("cluster_id").(string)
	createdAt, _ := d.Get("created_at").(string)
	k3sVersion, _ := d.Get("k3s_version").(string)
	location, _ := d.Get("location").(string)
	name, _ := d.Get("name").(string)
	nodeName, _ := d.Get("node_name").(string)
	size, _ := d.Get("size").(string)
	var status *models.SnapshotStatusInfo // SnapshotStatusInfo
	statusInterface, statusIsSet := d.GetOk("status")
	if statusIsSet && statusInterface != nil {
		statusMap := statusInterface.([]interface{})
		if len(statusMap) > 0 {
			status = SnapshotStatusInfoModelFromMap(statusMap[0].(map[string]interface{}))
		}
	}
	return &models.SnapshotGetResponse{
		ClusterID:  clusterID,
		CreatedAt:  createdAt,
		K3sVersion: k3sVersion,
		Location:   location,
		Name:       name,
		NodeName:   nodeName,
		Size:       size,
		Status:     status,
	}
}

func SnapshotGetResponseModelFromMap(m map[string]interface{}) *models.SnapshotGetResponse {
	clusterID := m["cluster_id"].(string)
	createdAt := m["created_at"].(string)
	k3sVersion := m["k3s_version"].(string)
	location := m["location"].(string)
	name := m["name"].(string)
	nodeName := m["node_name"].(string)
	size := m["size"].(string)
	var status *models.SnapshotStatusInfo // SnapshotStatusInfo
	statusInterface, statusIsSet := m["status"]
	if statusIsSet && statusInterface != nil {
		statusMap := statusInterface.([]interface{})
		if len(statusMap) > 0 {
			status = SnapshotStatusInfoModelFromMap(statusMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.SnapshotGetResponse{
		ClusterID:  clusterID,
		CreatedAt:  createdAt,
		K3sVersion: k3sVersion,
		Location:   location,
		Name:       name,
		NodeName:   nodeName,
		Size:       size,
		Status:     status,
	}
}

func SetSnapshotGetResponseResourceData(d *schema.ResourceData, m *models.SnapshotGetResponse) {
	d.Set("cluster_id", m.ClusterID)
	d.Set("created_at", m.CreatedAt)
	d.Set("k3s_version", m.K3sVersion)
	d.Set("location", m.Location)
	d.Set("name", m.Name)
	d.Set("node_name", m.NodeName)
	d.Set("size", m.Size)
	d.Set("status", SetSnapshotStatusInfoSubResourceData([]*models.SnapshotStatusInfo{m.Status}))
}

func SetSnapshotGetResponseSubResourceData(m []*models.SnapshotGetResponse) (d []*map[string]interface{}) {
	for _, SnapshotGetResponseModel := range m {
		if SnapshotGetResponseModel != nil {
			properties := make(map[string]interface{})
			properties["cluster_id"] = SnapshotGetResponseModel.ClusterID
			properties["created_at"] = SnapshotGetResponseModel.CreatedAt
			properties["k3s_version"] = SnapshotGetResponseModel.K3sVersion
			properties["location"] = SnapshotGetResponseModel.Location
			properties["name"] = SnapshotGetResponseModel.Name
			properties["node_name"] = SnapshotGetResponseModel.NodeName
			properties["size"] = SnapshotGetResponseModel.Size
			properties["status"] = SetSnapshotStatusInfoSubResourceData([]*models.SnapshotStatusInfo{SnapshotGetResponseModel.Status})
			d = append(d, &properties)
		}
	}
	return
}

func SnapshotGetResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_id": {
			Description: `ID of the cluster`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"created_at": {
			Description: `When the snapshot record was created (Unix timestamp)`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"k3s_version": {
			Description: `k3s Version`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"location": {
			Description: `Snapshot Location in ZKS server`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `Name of the snapshot`,
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
		},

		"node_name": {
			Description: `Name of the node where the snapshot was taken`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"size": {
			Description: `Size of the snapshot in bytes`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"status": {
			Description: `Status information for the snapshot`,
			Type:        schema.TypeList,
			Elem: &schema.Resource{
				Schema: SnapshotStatusInfoSchema(),
			},
			Computed: true,
		},
	}
}

func GetSnapshotGetResponsePropertyFields() (t []string) {
	return []string{
		"cluster_id",
		"created_at",
		"k3s_version",
		"location",
		"name",
		"node_name",
		"size",
		"status",
	}
}
