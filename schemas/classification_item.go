package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func ClassificationItemModel(d *schema.ResourceData) *models.ClassificationItem {
	destNode, _ := d.Get("dest_node").(string)
	levelInt, _ := d.Get("level").(int)
	level := int64(levelInt)
	metricValue, _ := d.Get("metric_value").(string)
	sourceNode, _ := d.Get("source_node").(string)
	return &models.ClassificationItem{
		DestNode:    destNode,
		Level:       level,
		MetricValue: metricValue,
		SourceNode:  sourceNode,
	}
}

func ClassificationItemModelFromMap(m map[string]interface{}) *models.ClassificationItem {
	destNode := m["dest_node"].(string)
	level := int64(m["level"].(int)) // int64
	metricValue := m["metric_value"].(string)
	sourceNode := m["source_node"].(string)
	return &models.ClassificationItem{
		DestNode:    destNode,
		Level:       level,
		MetricValue: metricValue,
		SourceNode:  sourceNode,
	}
}

// Update the underlying ClassificationItem resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetClassificationItemResourceData(d *schema.ResourceData, m *models.ClassificationItem) {
	d.Set("dest_node", m.DestNode)
	d.Set("level", m.Level)
	d.Set("metric_value", m.MetricValue)
	d.Set("source_node", m.SourceNode)
}

// Iterate through and update the ClassificationItem resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetClassificationItemSubResourceData(m []*models.ClassificationItem) (d []*map[string]interface{}) {
	for _, ClassificationItemModel := range m {
		if ClassificationItemModel != nil {
			properties := make(map[string]interface{})
			properties["dest_node"] = ClassificationItemModel.DestNode
			properties["level"] = ClassificationItemModel.Level
			properties["metric_value"] = ClassificationItemModel.MetricValue
			properties["source_node"] = ClassificationItemModel.SourceNode
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ClassificationItem resource defined in the Terraform configuration
func ClassificationItemSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"dest_node": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"level": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"metric_value": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"source_node": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the ClassificationItem resource
func GetClassificationItemPropertyFields() (t []string) {
	return []string{
		"dest_node",
		"level",
		"metric_value",
		"source_node",
	}
}
