package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func RunStateModel(d *schema.ResourceData) *models.RunState {
	runState, _ := d.Get("run_state").(models.RunState)
	return &runState
}

func RunStateModelFromMap(m map[string]interface{}) *models.RunState {
	runState := m["run_state"].(models.RunState)
	return &runState
}

// Update the underlying RunState resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetRunStateResourceData(d *schema.ResourceData, m *models.RunState) {
}

// Iterate through and update the RunState resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetRunStateSubResourceData(m []*models.RunState) (d []*map[string]interface{}) {
	for _, RunStateModel := range m {
		if RunStateModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the RunState resource defined in the Terraform configuration
func RunStateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the RunState resource
func GetRunStatePropertyFields() (t []string) {
	return []string{}
}
