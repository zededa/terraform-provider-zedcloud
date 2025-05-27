package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZserviceAssetIDsModel(d *schema.ResourceData) *models.ZserviceAssetIDs {
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
	return &models.ZserviceAssetIDs{
		Ids: ids,
	}
}

func ZserviceAssetIDsModelFromMap(m map[string]interface{}) *models.ZserviceAssetIDs {
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
	return &models.ZserviceAssetIDs{
		Ids: ids,
	}
}

func SetZserviceAssetIDsResourceData(d *schema.ResourceData, m *models.ZserviceAssetIDs) {
	d.Set("ids", m.Ids)
}

func SetZserviceAssetIDsSubResourceData(m []*models.ZserviceAssetIDs) (d []*map[string]interface{}) {
	for _, ZserviceAssetIDsModel := range m {
		if ZserviceAssetIDsModel != nil {
			properties := make(map[string]interface{})
			properties["ids"] = ZserviceAssetIDsModel.Ids
			d = append(d, &properties)
		}
	}
	return
}

func ZserviceAssetIDsSchema() map[string]*schema.Schema {
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

func GetZserviceAssetIDsPropertyFields() (t []string) {
	return []string{
		"ids",
	}
}
