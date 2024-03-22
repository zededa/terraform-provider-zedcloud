package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate MetricQueryResponseItem resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MetricQueryResponseItemModel(d *schema.ResourceData) *models.MetricQueryResponseItem {
	timestamp, _ := d.Get("timestamp").(strfmt.DateTime)
	values, _ := d.Get("values").([]float64)
	return &models.MetricQueryResponseItem{
		Timestamp: timestamp,
		Values:    values,
	}
}

func MetricQueryResponseItemModelFromMap(m map[string]interface{}) *models.MetricQueryResponseItem {
	timestamp := m["timestamp"].(strfmt.DateTime)
	values := m["values"].([]float64)
	return &models.MetricQueryResponseItem{
		Timestamp: timestamp,
		Values:    values,
	}
}

// Update the underlying MetricQueryResponseItem resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMetricQueryResponseItemResourceData(d *schema.ResourceData, m *models.MetricQueryResponseItem) {
	d.Set("timestamp", m.Timestamp)
	d.Set("values", m.Values)
}

// Iterate through and update the MetricQueryResponseItem resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMetricQueryResponseItemSubResourceData(m []*models.MetricQueryResponseItem) (d []*map[string]interface{}) {
	for _, MetricQueryResponseItemModel := range m {
		if MetricQueryResponseItemModel != nil {
			properties := make(map[string]interface{})
			properties["timestamp"] = MetricQueryResponseItemModel.Timestamp.String()
			properties["values"] = MetricQueryResponseItemModel.Values
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the MetricQueryResponseItem resource defined in the Terraform configuration
func MetricQueryResponseItemSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"timestamp": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"values": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []float64
			Elem: &schema.Schema{
				Type: schema.TypeFloat,
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the MetricQueryResponseItem resource
func GetMetricQueryResponseItemPropertyFields() (t []string) {
	return []string{
		"timestamp",
		"values",
	}
}
