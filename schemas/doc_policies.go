package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func DocPoliciesModel(d *schema.ResourceData) *models.DocPolicies {
	var list []*models.DocPolicySummary // []*DocPolicySummary
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
			m := DocPolicySummaryModelFromMap(v.(map[string]interface{}))
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
	return &models.DocPolicies{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
	}
}

func DocPoliciesModelFromMap(m map[string]interface{}) *models.DocPolicies {
	var list []*models.DocPolicySummary // []*DocPolicySummary
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
			m := DocPolicySummaryModelFromMap(v.(map[string]interface{}))
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
	return &models.DocPolicies{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
	}
}

func SetDocPoliciesResourceData(d *schema.ResourceData, m *models.DocPolicies) {
	d.Set("list", SetDocPolicySummarySubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_state", SetSummarySubResourceData([]*models.Summary{m.SummaryByState}))
}

func SetDocPoliciesSubResourceData(m []*models.DocPolicies) (d []*map[string]interface{}) {
	for _, DocPoliciesModel := range m {
		if DocPoliciesModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetDocPolicySummarySubResourceData(DocPoliciesModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{DocPoliciesModel.Next})
			properties["summary_by_state"] = SetSummarySubResourceData([]*models.Summary{DocPoliciesModel.SummaryByState})
			d = append(d, &properties)
		}
	}
	return
}

func DocPoliciesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `List of docpolicy`,
			Type:        schema.TypeList, //GoType: []*DocPolicySummary
			Elem: &schema.Resource{
				Schema: DocPolicySummarySchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"next": {
			Description: `Page details of the filtered records`,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},

		"summary_by_state": {
			Description: `Summary of filtered docpolicy`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},
	}
}

func GetDocPoliciesPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_state",
	}
}
