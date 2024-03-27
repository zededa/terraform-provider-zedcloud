package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func FlowlogDirectionModel(d *schema.ResourceData) *models.FlowlogDirection {
	flowlogDirection, _ := d.Get("flowlog_direction").(models.FlowlogDirection)
	return &flowlogDirection
}

func FlowlogDirectionModelFromMap(m map[string]interface{}) *models.FlowlogDirection {
	flowlogDirection := m["flowlog_direction"].(models.FlowlogDirection)
	return &flowlogDirection
}

// Update the underlying FlowlogDirection resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetFlowlogDirectionResourceData(d *schema.ResourceData, m *models.FlowlogDirection) {
}

// Iterate through and update the FlowlogDirection resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetFlowlogDirectionSubResourceData(m []*models.FlowlogDirection) (d []*map[string]interface{}) {
	for _, FlowlogDirectionModel := range m {
		if FlowlogDirectionModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the FlowlogDirection resource defined in the Terraform configuration
func FlowlogDirectionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the FlowlogDirection resource
func GetFlowlogDirectionPropertyFields() (t []string) {
	return []string{}
}
