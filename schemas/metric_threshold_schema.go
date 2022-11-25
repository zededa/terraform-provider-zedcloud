package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate MetricThreshold resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MetricThresholdModel(d *schema.ResourceData) *models.MetricThreshold {
	red := d.Get("red").(float64)
	yellow := d.Get("yellow").(float64)
	return &models.MetricThreshold{
		Red:    red,
		Yellow: yellow,
	}
}

func MetricThresholdModelFromMap(m map[string]interface{}) *models.MetricThreshold {
	red := m["red"].(float64)
	yellow := m["yellow"].(float64)
	return &models.MetricThreshold{
		Red:    red,
		Yellow: yellow,
	}
}

// Update the underlying MetricThreshold resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMetricThresholdResourceData(d *schema.ResourceData, m *models.MetricThreshold) {
	d.Set("red", m.Red)
	d.Set("yellow", m.Yellow)
}

// Iterate throught and update the MetricThreshold resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMetricThresholdSubResourceData(m []*models.MetricThreshold) (d []*map[string]interface{}) {
	for _, MetricThresholdModel := range m {
		if MetricThresholdModel != nil {
			properties := make(map[string]interface{})
			properties["red"] = MetricThresholdModel.Red
			properties["yellow"] = MetricThresholdModel.Yellow
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the MetricThreshold resource defined in the Terraform configuration
func MetricThresholdSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"red": {
			Description: ``,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"yellow": {
			Description: ``,
			Type:        schema.TypeFloat,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the MetricThreshold resource
func GetMetricThresholdPropertyFields() (t []string) {
	return []string{
		"red",
		"yellow",
	}
}
