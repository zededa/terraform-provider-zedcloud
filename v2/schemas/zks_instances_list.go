package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZksInstancesListModel(d *schema.ResourceData) *models.ZksInstancesList {
	var list []*models.ZksInstanceSummary // []*ZksInstanceSummary
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
			m := ZksInstanceSummaryModelFromMap(v.(map[string]interface{}))
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
	return &models.ZksInstancesList{
		List: list,
		Next: next,
	}
}

func ZksInstancesListModelFromMap(m map[string]interface{}) *models.ZksInstancesList {
	var list []*models.ZksInstanceSummary // []*ZksInstanceSummary
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
			m := ZksInstanceSummaryModelFromMap(v.(map[string]interface{}))
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
	return &models.ZksInstancesList{
		List: list,
		Next: next,
	}
}

func SetZksInstancesListResourceData(d *schema.ResourceData, m *models.ZksInstancesList) {
	d.Set("list", SetZksInstanceSummarySubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
}

func SetZksInstancesListSubResourceData(m []*models.ZksInstancesList) (d []*map[string]interface{}) {
	for _, ZksInstancesListModel := range m {
		if ZksInstancesListModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetZksInstanceSummarySubResourceData(ZksInstancesListModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{ZksInstancesListModel.Next})
			d = append(d, &properties)
		}
	}
	return
}

func ZksInstancesListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `List of ZKS instances`,
			Type:        schema.TypeList, //GoType: []*ZksInstanceSummary
			Elem: &schema.Resource{
				Schema: ZksInstanceSummarySchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"next": {
			Description: `next cursor`,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},
	}
}

func GetZksInstancesListPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
	}
}
