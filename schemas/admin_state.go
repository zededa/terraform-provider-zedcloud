package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate AdminState resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AdminStateModel(d *schema.ResourceData) *models.AdminState {
	adminState, _ := d.Get("admin_state").(models.AdminState)
	return &adminState
}

func AdminStateModelFromMap(m map[string]interface{}) *models.AdminState {
	adminState := m["admin_state"].(models.AdminState)
	return &adminState
}

// Update the underlying AdminState resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAdminStateResourceData(d *schema.ResourceData, m *models.AdminState) {
}

// Iterate through and update the AdminState resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAdminStateSubResourceData(m []*models.AdminState) (d []*map[string]interface{}) {
	for _, AdminStateModel := range m {
		if AdminStateModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AdminState resource defined in the Terraform configuration
func AdminStateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the AdminState resource
func GetAdminStatePropertyFields() (t []string) {
	return []string{}
}
