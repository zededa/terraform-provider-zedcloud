package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AssetGroupAssetTagListModel(d *schema.ResourceData) *models.AssetGroupAssetTagList {
	var values []*models.ZserviceAssetTag // []*ZserviceAssetTag
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
			m := ZserviceAssetTagModelFromMap(v.(map[string]interface{}))
			values = append(values, m)
		}
	}
	return &models.AssetGroupAssetTagList{
		Values: values,
	}
}

func AssetGroupAssetTagListModelFromMap(m map[string]interface{}) *models.AssetGroupAssetTagList {
	var values []*models.ZserviceAssetTag // []*ZserviceAssetTag
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
			m := ZserviceAssetTagModelFromMap(v.(map[string]interface{}))
			values = append(values, m)
		}
	}
	return &models.AssetGroupAssetTagList{
		Values: values,
	}
}

func SetAssetGroupAssetTagListResourceData(d *schema.ResourceData, m *models.AssetGroupAssetTagList) {
	d.Set("values", SetZserviceAssetTagSubResourceData(m.Values))
}

func SetAssetGroupAssetTagListSubResourceData(m []*models.AssetGroupAssetTagList) (d []*map[string]interface{}) {
	for _, AssetGroupAssetTagListModel := range m {
		if AssetGroupAssetTagListModel != nil {
			properties := make(map[string]interface{})
			properties["values"] = SetZserviceAssetTagSubResourceData(AssetGroupAssetTagListModel.Values)
			d = append(d, &properties)
		}
	}
	return
}

func AssetGroupAssetTagListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"values": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*ZserviceAssetTag
			Elem: &schema.Resource{
				Schema: ZserviceAssetTagSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

func GetAssetGroupAssetTagListPropertyFields() (t []string) {
	return []string{
		"values",
	}
}
