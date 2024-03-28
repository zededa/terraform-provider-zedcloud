package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// Function to perform the following actions:
// (1) Translate SysModelState resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func SysModelStateModel(d *schema.ResourceData) *models.SysModelState {
	sysModelState, _ := d.Get("sys_model_state").(models.SysModelState)
	return &sysModelState
}

func SysModelStateModelFromMap(m map[string]interface{}) *models.SysModelState {
	sysModelState := m["sys_model_state"].(models.SysModelState)
	return &sysModelState
}

// Update the underlying SysModelState resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetSysModelStateResourceData(d *schema.ResourceData, m *models.SysModelState) {
}

// Iterate through and update the SysModelState resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetSysModelStateSubResourceData(m []*models.SysModelState) (d []*map[string]interface{}) {
	for _, SysModelStateModel := range m {
		if SysModelStateModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the SysModelState resource defined in the Terraform configuration
func SysModelStateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the SysModelState resource
func GetSysModelStatePropertyFields() (t []string) {
	return []string{}
}
