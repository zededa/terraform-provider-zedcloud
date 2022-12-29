package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate PCRType resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PCRTypeModel(d *schema.ResourceData) *models.PCRType {
	pCRType, _ := d.Get("p_c_r_type").(models.PCRType)
	return &pCRType
}

func PCRTypeModelFromMap(m map[string]interface{}) *models.PCRType {
	pCRType := m["p_c_r_type"].(models.PCRType)
	return &pCRType
}

// Update the underlying PCRType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPCRTypeResourceData(d *schema.ResourceData, m *models.PCRType) {
}

// Iterate through and update the PCRType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPCRTypeSubResourceData(m []*models.PCRType) (d []*map[string]interface{}) {
	for _, PCRTypeModel := range m {
		if PCRTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the PCRType resource defined in the Terraform configuration
func PCRTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the PCRType resource
func GetPCRTypePropertyFields() (t []string) {
	return []string{}
}
