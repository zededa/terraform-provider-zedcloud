package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func SimcardStateModel(d *schema.ResourceData) *models.SimcardState {
	simcardState, _ := d.Get("simcard_state").(models.SimcardState)
	return &simcardState
}

func SimcardStateModelFromMap(m map[string]interface{}) *models.SimcardState {
	simcardState := m["simcard_state"].(models.SimcardState)
	return &simcardState
}

// Update the underlying SimcardState resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetSimcardStateResourceData(d *schema.ResourceData, m *models.SimcardState) {
}

// Iterate through and update the SimcardState resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetSimcardStateSubResourceData(m []*models.SimcardState) (d []*map[string]interface{}) {
	for _, SimcardStateModel := range m {
		if SimcardStateModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the SimcardState resource defined in the Terraform configuration
func SimcardStateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the SimcardState resource
func GetSimcardStatePropertyFields() (t []string) {
	return []string{}
}
