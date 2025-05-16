package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AssetGroupValueListModel(d *schema.ResourceData) *models.AssetGroupValueList {
	var values []string
	valuesInterface, valuesIsSet := d.GetOk("values")
	if valuesIsSet {
		var items []interface{}
		if listItems, isList := valuesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = valuesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			values = append(values, v.(string))
		}
	}
	return &models.AssetGroupValueList{
		Values: values,
	}
}

func AssetGroupValueListModelFromMap(m map[string]interface{}) *models.AssetGroupValueList {
	var values []string
	valuesInterface, valuesIsSet := m["values"]
	if valuesIsSet {
		var items []interface{}
		if listItems, isList := valuesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = valuesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			values = append(values, v.(string))
		}
	}
	return &models.AssetGroupValueList{
		Values: values,
	}
}

func SetAssetGroupValueListResourceData(d *schema.ResourceData, m *models.AssetGroupValueList) {
	d.Set("values", m.Values)
}

func SetAssetGroupValueListSubResourceData(m []*models.AssetGroupValueList) (d []*map[string]interface{}) {
	for _, AssetGroupValueListModel := range m {
		if AssetGroupValueListModel != nil {
			properties := make(map[string]interface{})
			properties["values"] = AssetGroupValueListModel.Values
			d = append(d, &properties)
		}
	}
	return
}

func AssetGroupValueListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"values": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

func GetAssetGroupValueListPropertyFields() (t []string) {
	return []string{
		"values",
	}
}
