package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func NetInstEdgeNodeClusterModel(d *schema.ResourceData) *models.NetInstEdgeNodeCluster {
	designatedNodeID, _ := d.Get("designated_node_id").(string)
	id, _ := d.Get("id").(string)
	return &models.NetInstEdgeNodeCluster{
		DesignatedNodeID: designatedNodeID,
		ID:               id,
	}
}

func NetInstEdgeNodeClusterModelFromMap(m map[string]interface{}) *models.NetInstEdgeNodeCluster {
	designatedNodeID := m["designated_node_id"].(string)
	id := m["id"].(string)
	return &models.NetInstEdgeNodeCluster{
		DesignatedNodeID: designatedNodeID,
		ID:               id,
	}
}

func SetNetInstEdgeNodeClusterResourceData(d *schema.ResourceData, m *models.NetInstEdgeNodeCluster) {
	d.Set("designated_node_id", m.DesignatedNodeID)
	d.Set("id", m.ID)
}

func SetNetInstEdgeNodeClusterSubResourceData(m []*models.NetInstEdgeNodeCluster) (d []*map[string]interface{}) {
	for _, NetInstEdgeNodeClusterModel := range m {
		if NetInstEdgeNodeClusterModel != nil {
			properties := make(map[string]interface{})
			properties["designated_node_id"] = NetInstEdgeNodeClusterModel.DesignatedNodeID
			properties["id"] = NetInstEdgeNodeClusterModel.ID
			d = append(d, &properties)
		}
	}
	return
}

func NetInstEdgeNodeClusterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"designated_node_id": {
			Description: `Id of the designated Edge Node`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `Id of the Edge Node Cluster`,
			Type:        schema.TypeString,
			Required:     true,
		},
	}
}

func GetNetInstEdgeNodeClusterPropertyFields() (t []string) {
	return []string{
		"designated_node_id",
		"id",
	}
}
