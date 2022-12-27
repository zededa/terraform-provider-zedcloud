package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate MetricsDetail resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MetricsDetailModel(d *schema.ResourceData) *models.MetricsDetail {
	queries := d.Get("queries").(map[string]string) // map[string]string
	results := d.Get("results").(map[string]string) // map[string]string
	return &models.MetricsDetail{
		Queries: queries,
		Results: results,
	}
}

func MetricsDetailModelFromMap(m map[string]interface{}) *models.MetricsDetail {
	queries := m["queries"].(map[string]string)
	results := m["results"].(map[string]string)
	return &models.MetricsDetail{
		Queries: queries,
		Results: results,
	}
}

// Update the underlying MetricsDetail resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMetricsDetailResourceData(d *schema.ResourceData, m *models.MetricsDetail) {
	d.Set("queries", m.Queries)
	d.Set("results", m.Results)
}

// Iterate throught and update the MetricsDetail resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMetricsDetailSubResourceData(m []*models.MetricsDetail) (d []*map[string]interface{}) {
	for _, MetricsDetailModel := range m {
		if MetricsDetailModel != nil {
			properties := make(map[string]interface{})
			properties["queries"] = MetricsDetailModel.Queries
			properties["results"] = MetricsDetailModel.Results
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the MetricsDetail resource defined in the Terraform configuration
func MetricsDetailSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"queries": {
			Description: `Mapping of queries variable keys and value`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"results": {
			Description: `Mapping of results variable keys and value`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the MetricsDetail resource
func GetMetricsDetailPropertyFields() (t []string) {
	return []string{
		"queries",
		"results",
	}
}
