package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func NodeActionModel(d *schema.ResourceData) *models.NodeAction {
	nodeAction, _ := d.Get("node_action").(models.NodeAction)
	return &nodeAction
}

func NodeActionModelFromMap(m map[string]interface{}) *models.NodeAction {
	nodeAction := m["node_action"].(models.NodeAction)
	return &nodeAction
}

func SetNodeActionResourceData(d *schema.ResourceData, m *models.NodeAction) {
}

func SetNodeActionSubResourceData(m []*models.NodeAction) (d []*map[string]interface{}) {
	for _, NodeActionModel := range m {
		if NodeActionModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func NodeActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetNodeActionPropertyFields() (t []string) {
	return []string{}
}
