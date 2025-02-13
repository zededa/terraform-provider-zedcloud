package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AppInstEdgeNodeClusterModel(d *schema.ResourceData) *models.AppInstEdgeNodeCluster {
	designatedNodeID, _ := d.Get("designated_node_id").(string)
	id, _ := d.Get("id").(string)
	return &models.AppInstEdgeNodeCluster{
		DesignatedNodeID: designatedNodeID,
		ID:               id,
	}
}

func AppInstEdgeNodeClusterModelFromMap(m map[string]interface{}) *models.AppInstEdgeNodeCluster {
	designatedNodeID := m["designated_node_id"].(string)
	id := m["id"].(string)
	return &models.AppInstEdgeNodeCluster{
		DesignatedNodeID: designatedNodeID,
		ID:               id,
	}
}

func SetAppInstEdgeNodeClusterResourceData(d *schema.ResourceData, m *models.AppInstEdgeNodeCluster) {
	d.Set("designated_node_id", m.DesignatedNodeID)
	d.Set("id", m.ID)
}

func SetAppInstEdgeNodeClusterSubResourceData(m []*models.AppInstEdgeNodeCluster) (d []*map[string]interface{}) {
	for _, AppInstEdgeNodeClusterModel := range m {
		if AppInstEdgeNodeClusterModel != nil {
			properties := make(map[string]interface{})
			properties["designated_node_id"] = AppInstEdgeNodeClusterModel.DesignatedNodeID
			properties["id"] = AppInstEdgeNodeClusterModel.ID
			d = append(d, &properties)
		}
	}
	return
}

func AppInstEdgeNodeClusterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"designated_node_id": {
			Description: `Id of the designated node`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `Id of the edge node cluster`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAppInstEdgeNodeClusterPropertyFields() (t []string) {
	return []string{
		"designated_node_id",
		"id",
	}
}
