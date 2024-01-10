package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func PatchEnvelopeListModel(d *schema.ResourceData) *models.PatchEnvelopeList {
	var list []*models.PatchEnvelope // []*PatchEnvelope
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
			m := PatchEnvelopeModelFromMap(v.(map[string]interface{}))
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
	var summaryByAction *models.Summary // Summary
	summaryByActionInterface, summaryByActionIsSet := d.GetOk("summary_by_action")
	if summaryByActionIsSet && summaryByActionInterface != nil {
		summaryByActionMap := summaryByActionInterface.([]interface{})
		if len(summaryByActionMap) > 0 {
			summaryByAction = SummaryModelFromMap(summaryByActionMap[0].(map[string]interface{}))
		}
	}
	totalCountInt, _ := d.Get("total_count").(int)
	totalCount := int64(totalCountInt)
	return &models.PatchEnvelopeList{
		List:            list,
		Next:            next,
		SummaryByAction: summaryByAction,
		TotalCount:      totalCount,
	}
}

func PatchEnvelopeListModelFromMap(m map[string]interface{}) *models.PatchEnvelopeList {
	var list []*models.PatchEnvelope // []*PatchEnvelope
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
			m := PatchEnvelopeModelFromMap(v.(map[string]interface{}))
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
	var summaryByAction *models.Summary // Summary
	summaryByActionInterface, summaryByActionIsSet := m["summary_by_action"]
	if summaryByActionIsSet && summaryByActionInterface != nil {
		summaryByActionMap := summaryByActionInterface.([]interface{})
		if len(summaryByActionMap) > 0 {
			summaryByAction = SummaryModelFromMap(summaryByActionMap[0].(map[string]interface{}))
		}
	}
	//
	totalCount := int64(m["total_count"].(int)) // int64
	return &models.PatchEnvelopeList{
		List:            list,
		Next:            next,
		SummaryByAction: summaryByAction,
		TotalCount:      totalCount,
	}
}

func SetPatchEnvelopeListResourceData(d *schema.ResourceData, m *models.PatchEnvelopeList) {
	d.Set("list", SetPatchEnvelopeSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_action", SetSummarySubResourceData([]*models.Summary{m.SummaryByAction}))
	d.Set("total_count", m.TotalCount)
}

func SetPatchEnvelopeListSubResourceData(m []*models.PatchEnvelopeList) (d []*map[string]interface{}) {
	for _, PatchEnvelopeListModel := range m {
		if PatchEnvelopeListModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetPatchEnvelopeSubResourceData(PatchEnvelopeListModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{PatchEnvelopeListModel.Next})
			properties["summary_by_action"] = SetSummarySubResourceData([]*models.Summary{PatchEnvelopeListModel.SummaryByAction})
			properties["total_count"] = PatchEnvelopeListModel.TotalCount
			d = append(d, &properties)
		}
	}
	return
}

func PatchEnvelopeListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `patch envelope summary list`,
			Type:        schema.TypeList, //GoType: []*PatchEnvelope
			Elem: &schema.Resource{
				Schema: PatchEnvelopeSchema(),
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

		"summary_by_action": {
			Description: `patch envelopes summary by action`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"total_count": {
			Description: `patch envelope total count`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetPatchEnvelopeListPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_action",
		"total_count",
	}
}
