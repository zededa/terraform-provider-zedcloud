package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate Summary resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func SummaryModel(d *schema.ResourceData) *models.Summary {
	description := d.Get("description").(string)
	total := int64(d.Get("total").(int))
	values := d.Get("values").(map[string]int64) // map[string]int64
	return &models.Summary{
		Description: description,
		Total:       total,
		Values:      values,
	}
}

func SummaryModelFromMap(m map[string]interface{}) *models.Summary {
	description := m["description"].(string)
	total := int64(m["total"].(int)) // int64 false false false
	values := m["values"].(map[string]int64)
	return &models.Summary{
		Description: description,
		Total:       total,
		Values:      values,
	}
}

// Update the underlying Summary resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetSummaryResourceData(d *schema.ResourceData, m *models.Summary) {
	d.Set("description", m.Description)
	d.Set("total", m.Total)
	d.Set("values", m.Values)
}

// Iterate throught and update the Summary resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetSummarySubResourceData(m []*models.Summary) (d []*map[string]interface{}) {
	for _, SummaryModel := range m {
		if SummaryModel != nil {
			properties := make(map[string]interface{})
			properties["description"] = SummaryModel.Description
			properties["total"] = SummaryModel.Total
			properties["values"] = SummaryModel.Values
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the Summary resource defined in the Terraform configuration
func SummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Description: `Summary description`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"total": {
			Description: `Total`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"values": {
			Description: `Values: Map for storing <string, uint32>`,
			Type:        schema.TypeMap, //GoType: map[string]int64
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the Summary resource
func GetSummaryPropertyFields() (t []string) {
	return []string{
		"description",
		"total",
		"values",
	}
}
