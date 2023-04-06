package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func ClusterConfigModel(d *schema.ResourceData) *models.ClusterConfig {
	minNodesRequiredInt, _ := d.Get("min_nodes_required").(int)
	minNodesRequired := int64(minNodesRequiredInt)
	return &models.ClusterConfig{
		MinNodesRequired: minNodesRequired,
	}
}

func ClusterConfigModelFromMap(m map[string]interface{}) *models.ClusterConfig {
	minNodesRequired := int64(m["min_nodes_required"].(int)) // int64
	return &models.ClusterConfig{
		MinNodesRequired: minNodesRequired,
	}
}

func SetClusterConfigResourceData(d *schema.ResourceData, m *models.ClusterConfig) {
	d.Set("min_nodes_required", m.MinNodesRequired)
}

func SetClusterConfigSubResourceData(m []*models.ClusterConfig) (d []*map[string]interface{}) {
	for _, ClusterConfigModel := range m {
		if ClusterConfigModel != nil {
			properties := make(map[string]interface{})
			properties["min_nodes_required"] = ClusterConfigModel.MinNodesRequired
			d = append(d, &properties)
		}
	}
	return
}

func ClusterConfig() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"min_nodes_required": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetClusterConfigPropertyFields() (t []string) {
	return []string{
		"min_nodes_required",
	}
}
