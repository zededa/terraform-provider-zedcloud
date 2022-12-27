package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeviceConfigList resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceConfigListModel(d *schema.ResourceData) *models.DeviceConfigList {
	list := d.Get("list").([]*models.DeviceConfigSummary) // []*DeviceConfigSummary
	var next *models.Cursor                               // Cursor
	nextInterface, nextIsSet := d.GetOk("next")
	if nextIsSet {
		nextMap := nextInterface.([]interface{})[0].(map[string]interface{})
		next = CursorModelFromMap(nextMap)
	}
	var summaryByState *models.Summary // Summary
	summaryByStateInterface, summaryByStateIsSet := d.GetOk("summary_by_state")
	if summaryByStateIsSet {
		summaryByStateMap := summaryByStateInterface.([]interface{})[0].(map[string]interface{})
		summaryByState = SummaryModelFromMap(summaryByStateMap)
	}
	var summaryByTagDistribution *models.Summary // Summary
	summaryByTagDistributionInterface, summaryByTagDistributionIsSet := d.GetOk("summary_by_tag_distribution")
	if summaryByTagDistributionIsSet {
		summaryByTagDistributionMap := summaryByTagDistributionInterface.([]interface{})[0].(map[string]interface{})
		summaryByTagDistribution = SummaryModelFromMap(summaryByTagDistributionMap)
	}
	return &models.DeviceConfigList{
		List:                     list,
		Next:                     next,
		SummaryByState:           summaryByState,
		SummaryByTagDistribution: summaryByTagDistribution,
	}
}

func DeviceConfigListModelFromMap(m map[string]interface{}) *models.DeviceConfigList {
	list := m["list"].([]*models.DeviceConfigSummary) // []*DeviceConfigSummary
	var next *models.Cursor                           // Cursor
	nextInterface, nextIsSet := m["next"]
	if nextIsSet {
		nextMap := nextInterface.([]interface{})[0].(map[string]interface{})
		next = CursorModelFromMap(nextMap)
	}
	//
	var summaryByState *models.Summary // Summary
	summaryByStateInterface, summaryByStateIsSet := m["summary_by_state"]
	if summaryByStateIsSet {
		summaryByStateMap := summaryByStateInterface.([]interface{})[0].(map[string]interface{})
		summaryByState = SummaryModelFromMap(summaryByStateMap)
	}
	//
	var summaryByTagDistribution *models.Summary // Summary
	summaryByTagDistributionInterface, summaryByTagDistributionIsSet := m["summary_by_tag_distribution"]
	if summaryByTagDistributionIsSet {
		summaryByTagDistributionMap := summaryByTagDistributionInterface.([]interface{})[0].(map[string]interface{})
		summaryByTagDistribution = SummaryModelFromMap(summaryByTagDistributionMap)
	}
	//
	return &models.DeviceConfigList{
		List:                     list,
		Next:                     next,
		SummaryByState:           summaryByState,
		SummaryByTagDistribution: summaryByTagDistribution,
	}
}

// Update the underlying DeviceConfigList resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceConfigListResourceData(d *schema.ResourceData, m *models.DeviceConfigList) {
	d.Set("list", SetDeviceConfigSummarySubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_state", SetSummarySubResourceData([]*models.Summary{m.SummaryByState}))
	d.Set("summary_by_tag_distribution", SetSummarySubResourceData([]*models.Summary{m.SummaryByTagDistribution}))
}

// Iterate throught and update the DeviceConfigList resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceConfigListSubResourceData(m []*models.DeviceConfigList) (d []*map[string]interface{}) {
	for _, DeviceConfigListModel := range m {
		if DeviceConfigListModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetDeviceConfigSummarySubResourceData(DeviceConfigListModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{DeviceConfigListModel.Next})
			properties["summary_by_state"] = SetSummarySubResourceData([]*models.Summary{DeviceConfigListModel.SummaryByState})
			properties["summary_by_tag_distribution"] = SetSummarySubResourceData([]*models.Summary{DeviceConfigListModel.SummaryByTagDistribution})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceConfigList resource defined in the Terraform configuration
func DeviceConfigListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `device config list`,
			Type:        schema.TypeList, //GoType: []*DeviceConfigSummary
			Elem: &schema.Resource{
				Schema: DeviceConfigSummarySchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},

		"next": {
			Description: `filter next`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Required: true,
		},

		"summary_by_state": {
			Description: `Summary by state`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Required: true,
		},

		"summary_by_tag_distribution": {
			Description: `Summary by project distribution`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the DeviceConfigList resource
func GetDeviceConfigListPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_state",
		"summary_by_tag_distribution",
	}
}
