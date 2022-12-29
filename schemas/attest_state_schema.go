package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate AttestState resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AttestStateModel(d *schema.ResourceData) *models.AttestState {
	attestState, _ := d.Get("attest_state").(models.AttestState)
	return &attestState
}

func AttestStateModelFromMap(m map[string]interface{}) *models.AttestState {
	attestState := m["attest_state"].(models.AttestState)
	return &attestState
}

// Update the underlying AttestState resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAttestStateResourceData(d *schema.ResourceData, m *models.AttestState) {
}

// Iterate through and update the AttestState resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAttestStateSubResourceData(m []*models.AttestState) (d []*map[string]interface{}) {
	for _, AttestStateModel := range m {
		if AttestStateModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AttestState resource defined in the Terraform configuration
func AttestStateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the AttestState resource
func GetAttestStatePropertyFields() (t []string) {
	return []string{}
}
