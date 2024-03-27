package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate CPUSummary resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func CPUSummaryModel(d *schema.ResourceData) *models.CPUSummary {
	utilization, _ := d.Get("utilization").(float64)
	return &models.CPUSummary{
		Utilization: utilization,
	}
}

func CPUSummaryModelFromMap(m map[string]interface{}) *models.CPUSummary {
	utilization := m["utilization"].(float64)
	return &models.CPUSummary{
		Utilization: utilization,
	}
}

// Update the underlying CPUSummary resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetCPUSummaryResourceData(d *schema.ResourceData, m *models.CPUSummary) {
	d.Set("utilization", m.Utilization)
}

// Iterate through and update the CPUSummary resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetCPUSummarySubResourceData(m []*models.CPUSummary) (d []*map[string]interface{}) {
	for _, CPUSummaryModel := range m {
		if CPUSummaryModel != nil {
			properties := make(map[string]interface{})
			properties["utilization"] = CPUSummaryModel.Utilization
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the CPUSummary resource defined in the Terraform configuration
func CPUSummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"utilization": {
			Description: `CPU Utilization`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the CPUSummary resource
func GetCPUSummaryPropertyFields() (t []string) {
	return []string{
		"utilization",
	}
}
