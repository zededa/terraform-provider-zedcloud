package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func InstanceTransitionActionModel(d *schema.ResourceData) *models.InstanceTransitionAction {
	instanceTransitionAction, _ := d.Get("instance_transition_action").(models.InstanceTransitionAction)
	return &instanceTransitionAction
}

func InstanceTransitionActionModelFromMap(m map[string]interface{}) *models.InstanceTransitionAction {
	instanceTransitionAction := m["instance_transition_action"].(models.InstanceTransitionAction)
	return &instanceTransitionAction
}

// Update the underlying InstanceTransitionAction resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetInstanceTransitionActionResourceData(d *schema.ResourceData, m *models.InstanceTransitionAction) {
}

// Iterate through and update the InstanceTransitionAction resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetInstanceTransitionActionSubResourceData(m []*models.InstanceTransitionAction) (d []*map[string]interface{}) {
	for _, InstanceTransitionActionModel := range m {
		if InstanceTransitionActionModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the InstanceTransitionAction resource defined in the Terraform configuration
func InstanceTransitionActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the InstanceTransitionAction resource
func GetInstanceTransitionActionPropertyFields() (t []string) {
	return []string{}
}
