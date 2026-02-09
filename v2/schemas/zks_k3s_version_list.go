package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZksK3sVersionListModel(d *schema.ResourceData) *models.ZksK3sVersionList {
	var list []*models.ZksK3sVersion // []*ZksK3sVersion
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
			m := ZksK3sVersionModelFromMap(v.(map[string]interface{}))
			list = append(list, m)
		}
	}
	return &models.ZksK3sVersionList{
		List: list,
	}
}

func ZksK3sVersionListModelFromMap(m map[string]interface{}) *models.ZksK3sVersionList {
	var list []*models.ZksK3sVersion // []*ZksK3sVersion
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
			m := ZksK3sVersionModelFromMap(v.(map[string]interface{}))
			list = append(list, m)
		}
	}
	return &models.ZksK3sVersionList{
		List: list,
	}
}

func SetZksK3sVersionListResourceData(d *schema.ResourceData, m *models.ZksK3sVersionList) {
	d.Set("list", SetZksK3sVersionSubResourceData(m.List))
}

func SetZksK3sVersionListSubResourceData(m []*models.ZksK3sVersionList) (d []*map[string]interface{}) {
	for _, ZksK3sVersionListModel := range m {
		if ZksK3sVersionListModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetZksK3sVersionSubResourceData(ZksK3sVersionListModel.List)
			d = append(d, &properties)
		}
	}
	return
}

func ZksK3sVersionListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `List of available K3s versions`,
			Type:        schema.TypeList, //GoType: []*ZksK3sVersion
			Elem: &schema.Resource{
				Schema: ZksK3sVersionSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

func GetZksK3sVersionListPropertyFields() (t []string) {
	return []string{
		"list",
	}
}
