package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AppInstStatusListMsgModel(d *schema.ResourceData) *models.AppInstStatusListMsg {
	var list []*models.AppInstStatusSummaryMsg // []*AppInstStatusSummaryMsg
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
			m := AppInstStatusSummaryMsgModelFromMap(v.(map[string]interface{}))
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
	return &models.AppInstStatusListMsg{
		List:             list,
		Next:             next,
		SummaryByAppType: summaryByAppType,
		SummaryByState:   summaryByState,
	}
}

func AppInstStatusListMsgModelFromMap(m map[string]interface{}) *models.AppInstStatusListMsg {
	var list []*models.AppInstStatusSummaryMsg // []*AppInstStatusSummaryMsg
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
			m := AppInstStatusSummaryMsgModelFromMap(v.(map[string]interface{}))
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
	return &models.AppInstStatusListMsg{
		List:             list,
		Next:             next,
		SummaryByAppType: summaryByAppType,
		SummaryByState:   summaryByState,
	}
}

// Update the underlying AppInstStatusListMsg resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppInstStatusListMsgResourceData(d *schema.ResourceData, m *models.AppInstStatusListMsg) {
	d.Set("list", SetAppInstStatusSummaryMsgSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_app_type", SetSummarySubResourceData([]*models.Summary{m.SummaryByAppType}))
	d.Set("summary_by_state", SetSummarySubResourceData([]*models.Summary{m.SummaryByState}))
}

// Iterate through and update the AppInstStatusListMsg resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppInstStatusListMsgSubResourceData(m []*models.AppInstStatusListMsg) (d []*map[string]interface{}) {
	for _, AppInstStatusListMsgModel := range m {
		if AppInstStatusListMsgModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetAppInstStatusSummaryMsgSubResourceData(AppInstStatusListMsgModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{AppInstStatusListMsgModel.Next})
			properties["summary_by_app_type"] = SetSummarySubResourceData([]*models.Summary{AppInstStatusListMsgModel.SummaryByAppType})
			properties["summary_by_state"] = SetSummarySubResourceData([]*models.Summary{AppInstStatusListMsgModel.SummaryByState})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppInstStatusListMsg resource defined in the Terraform configuration
func AppInstStatusListMsgSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `app instance status summary`,
			Type:        schema.TypeList, //GoType: []*AppInstStatusSummaryMsg
			Elem: &schema.Resource{
				Schema: AppInstStatusSummaryMsgSchema(),
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
	}
}

// Retrieve property field names for updating the AppInstStatusListMsg resource
func GetAppInstStatusListMsgPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_app_type",
		"summary_by_state",
	}
}
