package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AppInstConfigStatusListModel(d *schema.ResourceData) *models.AppInstConfigStatusList {
	var list []*models.AppInstConfigStatus // []*AppInstConfigStatus
	listInterface, listIsSet := d.GetOk("list")
	if listIsSet {
		var items []interface{}
		if listItems, isList := listInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = listInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AppInstConfigStatusModelFromMap(v.(map[string]interface{}))
			list = append(list, m)
		}
	}
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := d.GetOk("next")
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	var summaryByAppType *models.Summary // Summary
	summaryByAppTypeInterface, summaryByAppTypeIsSet := d.GetOk("summary_by_app_type")
	if summaryByAppTypeIsSet && summaryByAppTypeInterface != nil {
		summaryByAppTypeMap := summaryByAppTypeInterface.([]interface{})
		if len(summaryByAppTypeMap) > 0 {
			summaryByAppType = SummaryModelFromMap(summaryByAppTypeMap[0].(map[string]interface{}))
		}
	}
	var summaryByState *models.Summary // Summary
	summaryByStateInterface, summaryByStateIsSet := d.GetOk("summary_by_state")
	if summaryByStateIsSet && summaryByStateInterface != nil {
		summaryByStateMap := summaryByStateInterface.([]interface{})
		if len(summaryByStateMap) > 0 {
			summaryByState = SummaryModelFromMap(summaryByStateMap[0].(map[string]interface{}))
		}
	}
	totalCountInt, _ := d.Get("total_count").(int)
	totalCount := int64(totalCountInt)
	return &models.AppInstConfigStatusList{
		List:             list,
		Next:             next,
		SummaryByAppType: summaryByAppType,
		SummaryByState:   summaryByState,
		TotalCount:       totalCount,
	}
}

func AppInstConfigStatusListModelFromMap(m map[string]interface{}) *models.AppInstConfigStatusList {
	var list []*models.AppInstConfigStatus // []*AppInstConfigStatus
	listInterface, listIsSet := m["list"]
	if listIsSet {
		var items []interface{}
		if listItems, isList := listInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = listInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AppInstConfigStatusModelFromMap(v.(map[string]interface{}))
			list = append(list, m)
		}
	}
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := m["next"]
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	//
	var summaryByAppType *models.Summary // Summary
	summaryByAppTypeInterface, summaryByAppTypeIsSet := m["summary_by_app_type"]
	if summaryByAppTypeIsSet && summaryByAppTypeInterface != nil {
		summaryByAppTypeMap := summaryByAppTypeInterface.([]interface{})
		if len(summaryByAppTypeMap) > 0 {
			summaryByAppType = SummaryModelFromMap(summaryByAppTypeMap[0].(map[string]interface{}))
		}
	}
	//
	var summaryByState *models.Summary // Summary
	summaryByStateInterface, summaryByStateIsSet := m["summary_by_state"]
	if summaryByStateIsSet && summaryByStateInterface != nil {
		summaryByStateMap := summaryByStateInterface.([]interface{})
		if len(summaryByStateMap) > 0 {
			summaryByState = SummaryModelFromMap(summaryByStateMap[0].(map[string]interface{}))
		}
	}
	//
	totalCount := int64(m["total_count"].(int)) // int64
	return &models.AppInstConfigStatusList{
		List:             list,
		Next:             next,
		SummaryByAppType: summaryByAppType,
		SummaryByState:   summaryByState,
		TotalCount:       totalCount,
	}
}

// Update the underlying AppInstConfigStatusList resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppInstConfigStatusListResourceData(d *schema.ResourceData, m *models.AppInstConfigStatusList) {
	d.Set("list", SetAppInstConfigStatusSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_app_type", SetSummarySubResourceData([]*models.Summary{m.SummaryByAppType}))
	d.Set("summary_by_state", SetSummarySubResourceData([]*models.Summary{m.SummaryByState}))
	d.Set("total_count", m.TotalCount)
}

// Iterate through and update the AppInstConfigStatusList resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppInstConfigStatusListSubResourceData(m []*models.AppInstConfigStatusList) (d []*map[string]interface{}) {
	for _, AppInstConfigStatusListModel := range m {
		if AppInstConfigStatusListModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetAppInstConfigStatusSubResourceData(AppInstConfigStatusListModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{AppInstConfigStatusListModel.Next})
			properties["summary_by_app_type"] = SetSummarySubResourceData([]*models.Summary{AppInstConfigStatusListModel.SummaryByAppType})
			properties["summary_by_state"] = SetSummarySubResourceData([]*models.Summary{AppInstConfigStatusListModel.SummaryByState})
			properties["total_count"] = AppInstConfigStatusListModel.TotalCount
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppInstConfigStatusList resource defined in the Terraform configuration
func AppInstConfigStatusListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `app instance status summary`,
			Type:        schema.TypeList, //GoType: []*AppInstConfigStatus
			Elem: &schema.Resource{
				Schema: AppInstConfigStatusSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"next": {
			Description: `cursor next`,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},

		"summary_by_app_type": {
			Description: `app instance summary by application type`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"summary_by_state": {
			Description: `app instance status by state`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"total_count": {
			Description: `total number of records`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the AppInstConfigStatusList resource
func GetAppInstConfigStatusListPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_app_type",
		"summary_by_state",
		"total_count",
	}
}
