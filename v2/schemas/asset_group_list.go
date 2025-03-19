package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AssetGroupListModel(d *schema.ResourceData) *models.AssetGroupList {
	var list []*models.AssetGroup // []*AssetGroup
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
			m := AssetGroupModelFromMap(v.(map[string]interface{}))
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
	totalCountInt, _ := d.Get("total_count").(int)
	totalCount := int64(totalCountInt)
	return &models.AssetGroupList{
		List:       list,
		Next:       next,
		TotalCount: totalCount,
	}
}

func AssetGroupListModelFromMap(m map[string]interface{}) *models.AssetGroupList {
	var list []*models.AssetGroup // []*AssetGroup
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
			m := AssetGroupModelFromMap(v.(map[string]interface{}))
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
	totalCount := int64(m["total_count"].(int)) // int64
	return &models.AssetGroupList{
		List:       list,
		Next:       next,
		TotalCount: totalCount,
	}
}

func SetAssetGroupListResourceData(d *schema.ResourceData, m *models.AssetGroupList) {
	d.Set("list", SetAssetGroupSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("total_count", m.TotalCount)
}

func SetAssetGroupListSubResourceData(m []*models.AssetGroupList) (d []*map[string]interface{}) {
	for _, AssetGroupListModel := range m {
		if AssetGroupListModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetAssetGroupSubResourceData(AssetGroupListModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{AssetGroupListModel.Next})
			properties["total_count"] = AssetGroupListModel.TotalCount
			d = append(d, &properties)
		}
	}
	return
}

func AssetGroupListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `asset group summary list`,
			Type:        schema.TypeList, //GoType: []*AssetGroup
			Elem: &schema.Resource{
				Schema: AssetGroupSchema(),
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

		"total_count": {
			Description: `asset group total count`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetAssetGroupListPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"total_count",
	}
}
