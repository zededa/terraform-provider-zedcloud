package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func VolInstEdgeNodeClusterModel(d *schema.ResourceData) *models.VolInstEdgeNodeCluster {
	designatedNodeID, _ := d.Get("designated_node_id").(string)
	id, _ := d.Get("id").(string)
	return &models.VolInstEdgeNodeCluster{
		DesignatedNodeID: designatedNodeID,
		ID:               id,
	}
}

func VolInstEdgeNodeClusterModelFromMap(m map[string]interface{}) *models.VolInstEdgeNodeCluster {
	designatedNodeID := m["designated_node_id"].(string)
	id := m["id"].(string)
	return &models.VolInstEdgeNodeCluster{
		DesignatedNodeID: designatedNodeID,
		ID:               id,
	}
}

func SetVolInstEdgeNodeClusterResourceData(d *schema.ResourceData, m *models.VolInstEdgeNodeCluster) {
	d.Set("designated_node_id", m.DesignatedNodeID)
	d.Set("id", m.ID)
}

func SetVolInstEdgeNodeClusterSubResourceData(m []*models.VolInstEdgeNodeCluster) (d []*map[string]interface{}) {
	for _, VolInstEdgeNodeClusterModel := range m {
		if VolInstEdgeNodeClusterModel != nil {
			properties := make(map[string]interface{})
			properties["designated_node_id"] = VolInstEdgeNodeClusterModel.DesignatedNodeID
			properties["id"] = VolInstEdgeNodeClusterModel.ID
			d = append(d, &properties)
		}
	}
	return
}

func VolInstEdgeNodeClusterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"designated_node_id": {
			Description: `designated node ID`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `edge node cluster ID`,
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}

func GetVolInstEdgeNodeClusterPropertyFields() (t []string) {
	return []string{
		"designated_node_id",
		"id",
	}
}
