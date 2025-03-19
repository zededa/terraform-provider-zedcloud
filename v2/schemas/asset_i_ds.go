package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AssetIDsModel(d *schema.ResourceData) *models.AssetIDs {
	var ids []string
	idsInterface, idsIsSet := d.GetOk("ids")
	if idsIsSet {
		var items []interface{}
		if listItems, isList := idsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = idsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			ids = append(ids, v.(string))
		}
	}
	return &models.AssetIDs{
		Ids: ids,
	}
}

func AssetIDsModelFromMap(m map[string]interface{}) *models.AssetIDs {
	var ids []string
	idsInterface, idsIsSet := m["ids"]
	if idsIsSet {
		var items []interface{}
		if listItems, isList := idsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = idsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			ids = append(ids, v.(string))
		}
	}
	return &models.AssetIDs{
		Ids: ids,
	}
}

func SetAssetIDsResourceData(d *schema.ResourceData, m *models.AssetIDs) {
	d.Set("ids", m.Ids)
}

func SetAssetIDsSubResourceData(m []*models.AssetIDs) (d []*map[string]interface{}) {
	for _, AssetIDsModel := range m {
		if AssetIDsModel != nil {
			properties := make(map[string]interface{})
			properties["ids"] = AssetIDsModel.Ids
			d = append(d, &properties)
		}
	}
	return
}

func AssetIDsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ids": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

func GetAssetIDsPropertyFields() (t []string) {
	return []string{
		"ids",
	}
}
