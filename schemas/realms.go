package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func RealmsModel(d *schema.ResourceData) *models.Realms {
	var list []*models.Realm // []*Realm
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
			m := RealmModelFromMap(v.(map[string]interface{}))
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
	return &models.Realms{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
	}
}

func RealmsModelFromMap(m map[string]interface{}) *models.Realms {
	var list []*models.Realm // []*Realm
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
			m := RealmModelFromMap(v.(map[string]interface{}))
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
	return &models.Realms{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
	}
}

func SetRealmsResourceData(d *schema.ResourceData, m *models.Realms) {
	d.Set("list", SetRealmSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_state", SetSummarySubResourceData([]*models.Summary{m.SummaryByState}))
}

func SetRealmsSubResourceData(m []*models.Realms) (d []*map[string]interface{}) {
	for _, RealmsModel := range m {
		if RealmsModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetRealmSubResourceData(RealmsModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{RealmsModel.Next})
			properties["summary_by_state"] = SetSummarySubResourceData([]*models.Summary{RealmsModel.SummaryByState})
			d = append(d, &properties)
		}
	}
	return
}

func RealmsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `List of realms`,
			Type:        schema.TypeList, //GoType: []*Realm
			Elem: &schema.Resource{
				Schema: RealmSchema(),
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
			Description: `Summary of filtered realms`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},
	}
}

func GetRealmsPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_state",
	}
}
