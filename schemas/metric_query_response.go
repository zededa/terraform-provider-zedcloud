package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate MetricQueryResponse resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MetricQueryResponseModel(d *schema.ResourceData) *models.MetricQueryResponse {
	list, _ := d.Get("list").([]*models.MetricQueryResponseItem) // []*MetricQueryResponseItem
	metricType, _ := d.Get("metric_type").(string)
	var threshold *models.MetricThreshold // MetricThreshold
	thresholdInterface, thresholdIsSet := d.GetOk("threshold")
	if thresholdIsSet {
		thresholdMap := thresholdInterface.([]interface{})[0].(map[string]interface{})
		threshold = MetricThresholdModelFromMap(thresholdMap)
	}
	xLabel, _ := d.Get("x_label").(string)
	yLabels, _ := d.Get("y_labels").([]string)
	return &models.MetricQueryResponse{
		List:       list,
		MetricType: metricType,
		Threshold:  threshold,
		XLabel:     xLabel,
		YLabels:    yLabels,
	}
}

func MetricQueryResponseModelFromMap(m map[string]interface{}) *models.MetricQueryResponse {
	list := m["list"].([]*models.MetricQueryResponseItem) // []*MetricQueryResponseItem
	metricType := m["metric_type"].(string)
	var threshold *models.MetricThreshold // MetricThreshold
	thresholdInterface, thresholdIsSet := m["threshold"]
	if thresholdIsSet {
		thresholdMap := thresholdInterface.([]interface{})[0].(map[string]interface{})
		threshold = MetricThresholdModelFromMap(thresholdMap)
	}
	//
	xLabel := m["x_label"].(string)
	yLabels := m["y_labels"].([]string)
	return &models.MetricQueryResponse{
		List:       list,
		MetricType: metricType,
		Threshold:  threshold,
		XLabel:     xLabel,
		YLabels:    yLabels,
	}
}

// Update the underlying MetricQueryResponse resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMetricQueryResponseResourceData(d *schema.ResourceData, m *models.MetricQueryResponse) {
	d.Set("list", SetMetricQueryResponseItemSubResourceData(m.List))
	d.Set("metric_type", m.MetricType)
	d.Set("threshold", SetMetricThresholdSubResourceData([]*models.MetricThreshold{m.Threshold}))
	d.Set("x_label", m.XLabel)
	d.Set("y_labels", m.YLabels)
}

// Iterate through and update the MetricQueryResponse resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMetricQueryResponseSubResourceData(m []*models.MetricQueryResponse) (d []*map[string]interface{}) {
	for _, MetricQueryResponseModel := range m {
		if MetricQueryResponseModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetMetricQueryResponseItemSubResourceData(MetricQueryResponseModel.List)
			properties["metric_type"] = MetricQueryResponseModel.MetricType
			properties["threshold"] = SetMetricThresholdSubResourceData([]*models.MetricThreshold{MetricQueryResponseModel.Threshold})
			properties["x_label"] = MetricQueryResponseModel.XLabel
			properties["y_labels"] = MetricQueryResponseModel.YLabels
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the MetricQueryResponse resource defined in the Terraform configuration
func MetricQueryResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*MetricQueryResponseItem
			Elem: &schema.Resource{
				Schema: MetricQueryResponseItemSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"metric_type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"threshold": {
			Description: ``,
			Type:        schema.TypeList, //GoType: MetricThreshold
			Elem: &schema.Resource{
				Schema: MetricThresholdSchema(),
			},
			Optional: true,
		},

		"x_label": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"y_labels": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the MetricQueryResponse resource
func GetMetricQueryResponsePropertyFields() (t []string) {
	return []string{
		"list",
		"metric_type",
		"threshold",
		"x_label",
		"y_labels",
	}
}
