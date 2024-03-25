package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func FlowlogMetricModel(d *schema.ResourceData) *models.FlowlogMetric {
	flowlogMetric, _ := d.Get("flowlog_metric").(models.FlowlogMetric)
	return &flowlogMetric
}

func FlowlogMetricModelFromMap(m map[string]interface{}) *models.FlowlogMetric {
	flowlogMetric := m["flowlog_metric"].(models.FlowlogMetric)
	return &flowlogMetric
}

// Update the underlying FlowlogMetric resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetFlowlogMetricResourceData(d *schema.ResourceData, m *models.FlowlogMetric) {
}

// Iterate through and update the FlowlogMetric resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetFlowlogMetricSubResourceData(m []*models.FlowlogMetric) (d []*map[string]interface{}) {
	for _, FlowlogMetricModel := range m {
		if FlowlogMetricModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the FlowlogMetric resource defined in the Terraform configuration
func FlowlogMetricSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the FlowlogMetric resource
func GetFlowlogMetricPropertyFields() (t []string) {
	return []string{}
}
