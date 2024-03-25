package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func FlowlogCategoryTypeModel(d *schema.ResourceData) *models.FlowlogCategoryType {
	flowlogCategoryType, _ := d.Get("flowlog_category_type").(models.FlowlogCategoryType)
	return &flowlogCategoryType
}

func FlowlogCategoryTypeModelFromMap(m map[string]interface{}) *models.FlowlogCategoryType {
	flowlogCategoryType := m["flowlog_category_type"].(models.FlowlogCategoryType)
	return &flowlogCategoryType
}

// Update the underlying FlowlogCategoryType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetFlowlogCategoryTypeResourceData(d *schema.ResourceData, m *models.FlowlogCategoryType) {
}

// Iterate through and update the FlowlogCategoryType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetFlowlogCategoryTypeSubResourceData(m []*models.FlowlogCategoryType) (d []*map[string]interface{}) {
	for _, FlowlogCategoryTypeModel := range m {
		if FlowlogCategoryTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the FlowlogCategoryType resource defined in the Terraform configuration
func FlowlogCategoryTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the FlowlogCategoryType resource
func GetFlowlogCategoryTypePropertyFields() (t []string) {
	return []string{}
}
