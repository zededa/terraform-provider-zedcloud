package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func FlowlogActionModel(d *schema.ResourceData) *models.FlowlogAction {
	flowlogAction, _ := d.Get("flowlog_action").(models.FlowlogAction)
	return &flowlogAction
}

func FlowlogActionModelFromMap(m map[string]interface{}) *models.FlowlogAction {
	flowlogAction := m["flowlog_action"].(models.FlowlogAction)
	return &flowlogAction
}

// Update the underlying FlowlogAction resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetFlowlogActionResourceData(d *schema.ResourceData, m *models.FlowlogAction) {
}

// Iterate through and update the FlowlogAction resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetFlowlogActionSubResourceData(m []*models.FlowlogAction) (d []*map[string]interface{}) {
	for _, FlowlogActionModel := range m {
		if FlowlogActionModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the FlowlogAction resource defined in the Terraform configuration
func FlowlogActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the FlowlogAction resource
func GetFlowlogActionPropertyFields() (t []string) {
	return []string{}
}
