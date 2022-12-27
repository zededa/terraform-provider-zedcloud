package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeviceStatusListMsg resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceStatusListMsgModel(d *schema.ResourceData) *models.DeviceStatusListMsg {
	list := d.Get("list").([]*models.DeviceStatusSummaryMsg) // []*DeviceStatusSummaryMsg
	var next *models.Cursor                                  // Cursor
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
	totalEdgeviewActive := int64(d.Get("total_edgeview_active").(int))
	return &models.DeviceStatusListMsg{
		List:                      list,
		Next:                      next,
		SummaryByAppInstanceCount: summaryByAppInstanceCount,
		SummaryByEVEDistribution:  summaryByEVEDistribution,
		SummaryByState:            summaryByState,
		TotalCount:                totalCount,
		TotalEdgeviewActive:       totalEdgeviewActive,
	}
}

func DeviceStatusListMsgModelFromMap(m map[string]interface{}) *models.DeviceStatusListMsg {
	list := m["list"].([]*models.DeviceStatusSummaryMsg) // []*DeviceStatusSummaryMsg
	var next *models.Cursor                              // Cursor
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
	totalCount := int64(m["total_count"].(int))                    // int64 false false false
	totalEdgeviewActive := int64(m["total_edgeview_active"].(int)) // int64 false false false
	return &models.DeviceStatusListMsg{
		List:                      list,
		Next:                      next,
		SummaryByAppInstanceCount: summaryByAppInstanceCount,
		SummaryByEVEDistribution:  summaryByEVEDistribution,
		SummaryByState:            summaryByState,
		TotalCount:                totalCount,
		TotalEdgeviewActive:       totalEdgeviewActive,
	}
}

// Update the underlying DeviceStatusListMsg resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceStatusListMsgResourceData(d *schema.ResourceData, m *models.DeviceStatusListMsg) {
	d.Set("list", SetDeviceStatusSummaryMsgSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_app_instance_count", SetSummarySubResourceData([]*models.Summary{m.SummaryByAppInstanceCount}))
	d.Set("summary_by_e_v_e_distribution", SetSummarySubResourceData([]*models.Summary{m.SummaryByEVEDistribution}))
	d.Set("summary_by_state", SetSummarySubResourceData([]*models.Summary{m.SummaryByState}))
	d.Set("total_count", m.TotalCount)
	d.Set("total_edgeview_active", m.TotalEdgeviewActive)
}

// Iterate throught and update the DeviceStatusListMsg resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceStatusListMsgSubResourceData(m []*models.DeviceStatusListMsg) (d []*map[string]interface{}) {
	for _, DeviceStatusListMsgModel := range m {
		if DeviceStatusListMsgModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetDeviceStatusSummaryMsgSubResourceData(DeviceStatusListMsgModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{DeviceStatusListMsgModel.Next})
			properties["summary_by_app_instance_count"] = SetSummarySubResourceData([]*models.Summary{DeviceStatusListMsgModel.SummaryByAppInstanceCount})
			properties["summary_by_e_v_e_distribution"] = SetSummarySubResourceData([]*models.Summary{DeviceStatusListMsgModel.SummaryByEVEDistribution})
			properties["summary_by_state"] = SetSummarySubResourceData([]*models.Summary{DeviceStatusListMsgModel.SummaryByState})
			properties["total_count"] = DeviceStatusListMsgModel.TotalCount
			properties["total_edgeview_active"] = DeviceStatusListMsgModel.TotalEdgeviewActive
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceStatusListMsg resource defined in the Terraform configuration
func DeviceStatusListMsgSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*DeviceStatusSummaryMsg
			Elem: &schema.Resource{
				Schema: DeviceStatusSummaryMsgSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"next": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},

		"summary_by_app_instance_count": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"summary_by_e_v_e_distribution": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"summary_by_state": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"total_count": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"total_edgeview_active": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the DeviceStatusListMsg resource
func GetDeviceStatusListMsgPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_app_instance_count",
		"summary_by_e_v_e_distribution",
		"summary_by_state",
		"total_count",
		"total_edgeview_active",
	}
}
