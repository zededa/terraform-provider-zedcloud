package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func PatchStatusListMsgModel(d *schema.ResourceData) *models.PatchStatusListMsg {
	var list []*models.PatchEnvelopeStatus // []*PatchEnvelopeStatus
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
			m := PatchEnvelopeStatusModelFromMap(v.(map[string]interface{}))
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
	return &models.PatchStatusListMsg{
		List:       list,
		Next:       next,
		TotalCount: totalCount,
	}
}

func PatchStatusListMsgModelFromMap(m map[string]interface{}) *models.PatchStatusListMsg {
	var list []*models.PatchEnvelopeStatus // []*PatchEnvelopeStatus
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
			m := PatchEnvelopeStatusModelFromMap(v.(map[string]interface{}))
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
	return &models.PatchStatusListMsg{
		List:       list,
		Next:       next,
		TotalCount: totalCount,
	}
}

func SetPatchStatusListMsgResourceData(d *schema.ResourceData, m *models.PatchStatusListMsg) {
	d.Set("list", SetPatchEnvelopeStatusSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("total_count", m.TotalCount)
}

func SetPatchStatusListMsgSubResourceData(m []*models.PatchStatusListMsg) (d []*map[string]interface{}) {
	for _, PatchStatusListMsgModel := range m {
		if PatchStatusListMsgModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetPatchEnvelopeStatusSubResourceData(PatchStatusListMsgModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{PatchStatusListMsgModel.Next})
			properties["total_count"] = PatchStatusListMsgModel.TotalCount
			d = append(d, &properties)
		}
	}
	return
}

func PatchStatusListMsgSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*PatchEnvelopeStatus
			Elem: &schema.Resource{
				Schema: PatchEnvelopeStatusSchema(),
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

		"total_count": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetPatchStatusListMsgPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"total_count",
	}
}
