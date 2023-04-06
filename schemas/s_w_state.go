package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate SWState resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func SWStateModel(d *schema.ResourceData) *models.SWState {
	sWState, _ := d.Get("s_w_state").(models.SWState)
	return &sWState
}

func SWStateModelFromMap(m map[string]interface{}) *models.SWState {
	sWState := m["s_w_state"].(models.SWState)
	return &sWState
}

// Update the underlying SWState resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetSWStateResourceData(d *schema.ResourceData, m *models.SWState) {
}

// Iterate through and update the SWState resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetSWStateSubResourceData(m []*models.SWState) (d []*map[string]interface{}) {
	for _, SWStateModel := range m {
		if SWStateModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the SWState resource defined in the Terraform configuration
func SWStateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the SWState resource
func GetSWStatePropertyFields() (t []string) {
	return []string{}
}
