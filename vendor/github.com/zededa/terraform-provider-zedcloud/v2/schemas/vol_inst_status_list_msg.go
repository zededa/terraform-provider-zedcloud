package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func VolInstStatusListMsgModel(d *schema.ResourceData) *models.VolInstStatusListMsg {
	var list []*models.VolInstStatusSummaryMsg // []*VolInstStatusSummaryMsg
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
			m := VolInstStatusSummaryMsgModelFromMap(v.(map[string]interface{}))
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
	var summaryByState *models.Summary // Summary
	summaryByStateInterface, summaryByStateIsSet := d.GetOk("summary_by_state")
	if summaryByStateIsSet && summaryByStateInterface != nil {
		summaryByStateMap := summaryByStateInterface.([]interface{})
		if len(summaryByStateMap) > 0 {
			summaryByState = SummaryModelFromMap(summaryByStateMap[0].(map[string]interface{}))
		}
	}
	var summaryByType *models.Summary // Summary
	summaryByTypeInterface, summaryByTypeIsSet := d.GetOk("summary_by_type")
	if summaryByTypeIsSet && summaryByTypeInterface != nil {
		summaryByTypeMap := summaryByTypeInterface.([]interface{})
		if len(summaryByTypeMap) > 0 {
			summaryByType = SummaryModelFromMap(summaryByTypeMap[0].(map[string]interface{}))
		}
	}
	totalCountInt, _ := d.Get("total_count").(int)
	totalCount := int64(totalCountInt)
	return &models.VolInstStatusListMsg{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
		SummaryByType:  summaryByType,
		TotalCount:     totalCount,
	}
}

func VolInstStatusListMsgModelFromMap(m map[string]interface{}) *models.VolInstStatusListMsg {
	var list []*models.VolInstStatusSummaryMsg // []*VolInstStatusSummaryMsg
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
			m := VolInstStatusSummaryMsgModelFromMap(v.(map[string]interface{}))
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
	var summaryByState *models.Summary // Summary
	summaryByStateInterface, summaryByStateIsSet := m["summary_by_state"]
	if summaryByStateIsSet && summaryByStateInterface != nil {
		summaryByStateMap := summaryByStateInterface.([]interface{})
		if len(summaryByStateMap) > 0 {
			summaryByState = SummaryModelFromMap(summaryByStateMap[0].(map[string]interface{}))
		}
	}
	//
	var summaryByType *models.Summary // Summary
	summaryByTypeInterface, summaryByTypeIsSet := m["summary_by_type"]
	if summaryByTypeIsSet && summaryByTypeInterface != nil {
		summaryByTypeMap := summaryByTypeInterface.([]interface{})
		if len(summaryByTypeMap) > 0 {
			summaryByType = SummaryModelFromMap(summaryByTypeMap[0].(map[string]interface{}))
		}
	}
	//
	totalCount := int64(m["total_count"].(int)) // int64
	return &models.VolInstStatusListMsg{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
		SummaryByType:  summaryByType,
		TotalCount:     totalCount,
	}
}

// Update the underlying VolInstStatusListMsg resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVolInstStatusListMsgResourceData(d *schema.ResourceData, m *models.VolInstStatusListMsg) {
	d.Set("list", SetVolInstStatusSummaryMsgSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_state", SetSummarySubResourceData([]*models.Summary{m.SummaryByState}))
	d.Set("summary_by_type", SetSummarySubResourceData([]*models.Summary{m.SummaryByType}))
	d.Set("total_count", m.TotalCount)
}

// Iterate through and update the VolInstStatusListMsg resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVolInstStatusListMsgSubResourceData(m []*models.VolInstStatusListMsg) (d []*map[string]interface{}) {
	for _, VolInstStatusListMsgModel := range m {
		if VolInstStatusListMsgModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetVolInstStatusSummaryMsgSubResourceData(VolInstStatusListMsgModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{VolInstStatusListMsgModel.Next})
			properties["summary_by_state"] = SetSummarySubResourceData([]*models.Summary{VolInstStatusListMsgModel.SummaryByState})
			properties["summary_by_type"] = SetSummarySubResourceData([]*models.Summary{VolInstStatusListMsgModel.SummaryByType})
			properties["total_count"] = VolInstStatusListMsgModel.TotalCount
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VolInstStatusListMsg resource defined in the Terraform configuration
func VolInstStatusListMsgSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*VolInstStatusSummaryMsg
			Elem: &schema.Resource{
				Schema: VolInstStatusSummaryMsgSchema(),
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

		"summary_by_state": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"summary_by_type": {
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
	}
}

// Retrieve property field names for updating the VolInstStatusListMsg resource
func GetVolInstStatusListMsgPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_state",
		"summary_by_type",
		"total_count",
	}
}
