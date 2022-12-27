package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeviceStatusConfigList resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceStatusConfigListModel(d *schema.ResourceData) *models.DeviceStatusConfigList {
	list := d.Get("list").([]*models.DeviceStatusConfig) // []*DeviceStatusConfig
	var next *models.Cursor                              // Cursor
	nextInterface, nextIsSet := d.GetOk("next")
	if nextIsSet {
		nextMap := nextInterface.([]interface{})[0].(map[string]interface{})
		next = CursorModelFromMap(nextMap)
	}
	var summaryByAppInstanceCount *models.Summary // Summary
	summaryByAppInstanceCountInterface, summaryByAppInstanceCountIsSet := d.GetOk("summary_by_app_instance_count")
	if summaryByAppInstanceCountIsSet {
		summaryByAppInstanceCountMap := summaryByAppInstanceCountInterface.([]interface{})[0].(map[string]interface{})
		summaryByAppInstanceCount = SummaryModelFromMap(summaryByAppInstanceCountMap)
	}
	var summaryByEVEDistribution *models.Summary // Summary
	summaryByEVEDistributionInterface, summaryByEVEDistributionIsSet := d.GetOk("summary_by_e_v_e_distribution")
	if summaryByEVEDistributionIsSet {
		summaryByEVEDistributionMap := summaryByEVEDistributionInterface.([]interface{})[0].(map[string]interface{})
		summaryByEVEDistribution = SummaryModelFromMap(summaryByEVEDistributionMap)
	}
	var summaryByState *models.Summary // Summary
	summaryByStateInterface, summaryByStateIsSet := d.GetOk("summary_by_state")
	if summaryByStateIsSet {
		summaryByStateMap := summaryByStateInterface.([]interface{})[0].(map[string]interface{})
		summaryByState = SummaryModelFromMap(summaryByStateMap)
	}
	totalCount := int64(d.Get("total_count").(int))
	totalEvActiveCount := int64(d.Get("total_ev_active_count").(int))
	return &models.DeviceStatusConfigList{
		List:                      list,
		Next:                      next,
		SummaryByAppInstanceCount: summaryByAppInstanceCount,
		SummaryByEVEDistribution:  summaryByEVEDistribution,
		SummaryByState:            summaryByState,
		TotalCount:                &totalCount, // int64 true false false
		TotalEvActiveCount:        totalEvActiveCount,
	}
}

func DeviceStatusConfigListModelFromMap(m map[string]interface{}) *models.DeviceStatusConfigList {
	list := m["list"].([]*models.DeviceStatusConfig) // []*DeviceStatusConfig
	var next *models.Cursor                          // Cursor
	nextInterface, nextIsSet := m["next"]
	if nextIsSet {
		nextMap := nextInterface.([]interface{})[0].(map[string]interface{})
		next = CursorModelFromMap(nextMap)
	}
	//
	var summaryByAppInstanceCount *models.Summary // Summary
	summaryByAppInstanceCountInterface, summaryByAppInstanceCountIsSet := m["summary_by_app_instance_count"]
	if summaryByAppInstanceCountIsSet {
		summaryByAppInstanceCountMap := summaryByAppInstanceCountInterface.([]interface{})[0].(map[string]interface{})
		summaryByAppInstanceCount = SummaryModelFromMap(summaryByAppInstanceCountMap)
	}
	//
	var summaryByEVEDistribution *models.Summary // Summary
	summaryByEVEDistributionInterface, summaryByEVEDistributionIsSet := m["summary_by_e_v_e_distribution"]
	if summaryByEVEDistributionIsSet {
		summaryByEVEDistributionMap := summaryByEVEDistributionInterface.([]interface{})[0].(map[string]interface{})
		summaryByEVEDistribution = SummaryModelFromMap(summaryByEVEDistributionMap)
	}
	//
	var summaryByState *models.Summary // Summary
	summaryByStateInterface, summaryByStateIsSet := m["summary_by_state"]
	if summaryByStateIsSet {
		summaryByStateMap := summaryByStateInterface.([]interface{})[0].(map[string]interface{})
		summaryByState = SummaryModelFromMap(summaryByStateMap)
	}
	//
	totalCount := int64(m["total_count"].(int))                   // int64 true false false
	totalEvActiveCount := int64(m["total_ev_active_count"].(int)) // int64 false false false
	return &models.DeviceStatusConfigList{
		List:                      list,
		Next:                      next,
		SummaryByAppInstanceCount: summaryByAppInstanceCount,
		SummaryByEVEDistribution:  summaryByEVEDistribution,
		SummaryByState:            summaryByState,
		TotalCount:                &totalCount,
		TotalEvActiveCount:        totalEvActiveCount,
	}
}

// Update the underlying DeviceStatusConfigList resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceStatusConfigListResourceData(d *schema.ResourceData, m *models.DeviceStatusConfigList) {
	d.Set("list", SetDeviceStatusConfigSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_app_instance_count", SetSummarySubResourceData([]*models.Summary{m.SummaryByAppInstanceCount}))
	d.Set("summary_by_e_v_e_distribution", SetSummarySubResourceData([]*models.Summary{m.SummaryByEVEDistribution}))
	d.Set("summary_by_state", SetSummarySubResourceData([]*models.Summary{m.SummaryByState}))
	d.Set("total_count", m.TotalCount)
	d.Set("total_ev_active_count", m.TotalEvActiveCount)
}

// Iterate throught and update the DeviceStatusConfigList resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceStatusConfigListSubResourceData(m []*models.DeviceStatusConfigList) (d []*map[string]interface{}) {
	for _, DeviceStatusConfigListModel := range m {
		if DeviceStatusConfigListModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetDeviceStatusConfigSubResourceData(DeviceStatusConfigListModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{DeviceStatusConfigListModel.Next})
			properties["summary_by_app_instance_count"] = SetSummarySubResourceData([]*models.Summary{DeviceStatusConfigListModel.SummaryByAppInstanceCount})
			properties["summary_by_e_v_e_distribution"] = SetSummarySubResourceData([]*models.Summary{DeviceStatusConfigListModel.SummaryByEVEDistribution})
			properties["summary_by_state"] = SetSummarySubResourceData([]*models.Summary{DeviceStatusConfigListModel.SummaryByState})
			properties["total_count"] = DeviceStatusConfigListModel.TotalCount
			properties["total_ev_active_count"] = DeviceStatusConfigListModel.TotalEvActiveCount
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceStatusConfigList resource defined in the Terraform configuration
func DeviceStatusConfigListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `List of device status config`,
			Type:        schema.TypeList, //GoType: []*DeviceStatusConfig
			Elem: &schema.Resource{
				Schema: DeviceStatusConfigSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},

		"next": {
			Description: `Page details of the filtered records`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"summary_by_app_instance_count": {
			Description: `Device status config summary by app instance count`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Required: true,
		},

		"summary_by_e_v_e_distribution": {
			Description: `Device status config summary by eve distribution`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Required: true,
		},

		"summary_by_state": {
			Description: `Device status config summary by state`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Required: true,
		},

		"total_count": {
			Description: `total count of devices`,
			Type:        schema.TypeInt,
			Required:    true,
		},

		"total_ev_active_count": {
			Description: `total count of edgeview active of devices`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the DeviceStatusConfigList resource
func GetDeviceStatusConfigListPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_app_instance_count",
		"summary_by_e_v_e_distribution",
		"summary_by_state",
		"total_count",
		"total_ev_active_count",
	}
}
