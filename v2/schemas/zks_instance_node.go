package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZksInstanceNodeModel(d *schema.ResourceData) *models.ZksInstanceNode {
	clusterInterface, _ := d.Get("cluster_interface").(string)
	id, _ := d.Get("id").(string)
	return &models.ZksInstanceNode{
		ClusterInterface: clusterInterface, // string
		ID:               id,
	}
}

func ZksInstanceNodeModelFromMap(m map[string]interface{}) *models.ZksInstanceNode {
	clusterInterface := m["cluster_interface"].(string)
	id := m["id"].(string)
	return &models.ZksInstanceNode{
		ClusterInterface: clusterInterface,
		ID:               id,
	}
}

func SetZksInstanceNodeResourceData(d *schema.ResourceData, m *models.ZksInstanceNode) {
	d.Set("cluster_interface", m.ClusterInterface)
	d.Set("id", m.ID)
}

func SetZksInstanceNodeSubResourceData(m []*models.ZksInstanceNode) (d []*map[string]interface{}) {
	for _, ZksInstanceNodeModel := range m {
		if ZksInstanceNodeModel != nil {
			properties := make(map[string]interface{})
			properties["cluster_interface"] = ZksInstanceNodeModel.ClusterInterface
			properties["id"] = ZksInstanceNodeModel.ID
			d = append(d, &properties)
		}
	}
	return
}

func ZksInstanceNodeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_interface": {
			Description: `Cluster interface of the node`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"id": {
			Description: `ID of the node`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func GetZksInstanceNodePropertyFields() (t []string) {
	return []string{
		"cluster_interface",
		"id",
	}
}
