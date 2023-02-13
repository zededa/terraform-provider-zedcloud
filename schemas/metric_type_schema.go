package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate MetricType resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MetricTypeModel(d *schema.ResourceData) *models.MetricType {
	metricType, _ := d.Get("metric_type").(models.MetricType)
	return &metricType
}

func MetricTypeModelFromMap(m map[string]interface{}) *models.MetricType {
	metricType := m["metric_type"].(models.MetricType)
	return &metricType
}

// Update the underlying MetricType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMetricTypeResourceData(d *schema.ResourceData, m *models.MetricType) {
}

// Iterate through and update the MetricType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMetricTypeSubResourceData(m []*models.MetricType) (d []*map[string]interface{}) {
	for _, MetricTypeModel := range m {
		if MetricTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the MetricType resource defined in the Terraform configuration
func MetricTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the MetricType resource
func GetMetricTypePropertyFields() (t []string) {
	return []string{}
}
