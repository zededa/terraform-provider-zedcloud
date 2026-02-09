package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZksInstanceBulkUpdateModel(d *schema.ResourceData) *models.ZksInstanceBulkUpdate {
	var data *models.ZksInstanceBulkUpdateData // ZksInstanceBulkUpdateData
	dataInterface, dataIsSet := d.GetOk("data")
	if dataIsSet && dataInterface != nil {
		dataMap := dataInterface.([]interface{})
		if len(dataMap) > 0 {
			data = ZksInstanceBulkUpdateDataModelFromMap(dataMap[0].(map[string]interface{}))
		}
	}
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
	return &models.ZksInstanceBulkUpdate{
		Data: data,
		Ids:  ids,
	}
}

func ZksInstanceBulkUpdateModelFromMap(m map[string]interface{}) *models.ZksInstanceBulkUpdate {
	var data *models.ZksInstanceBulkUpdateData // ZksInstanceBulkUpdateData
	dataInterface, dataIsSet := m["data"]
	if dataIsSet && dataInterface != nil {
		dataMap := dataInterface.([]interface{})
		if len(dataMap) > 0 {
			data = ZksInstanceBulkUpdateDataModelFromMap(dataMap[0].(map[string]interface{}))
		}
	}
	//
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
	return &models.ZksInstanceBulkUpdate{
		Data: data,
		Ids:  ids,
	}
}

func SetZksInstanceBulkUpdateResourceData(d *schema.ResourceData, m *models.ZksInstanceBulkUpdate) {
	d.Set("data", SetZksInstanceBulkUpdateDataSubResourceData([]*models.ZksInstanceBulkUpdateData{m.Data}))
	d.Set("ids", m.Ids)
}

func SetZksInstanceBulkUpdateSubResourceData(m []*models.ZksInstanceBulkUpdate) (d []*map[string]interface{}) {
	for _, ZksInstanceBulkUpdateModel := range m {
		if ZksInstanceBulkUpdateModel != nil {
			properties := make(map[string]interface{})
			properties["data"] = SetZksInstanceBulkUpdateDataSubResourceData([]*models.ZksInstanceBulkUpdateData{ZksInstanceBulkUpdateModel.Data})
			properties["ids"] = ZksInstanceBulkUpdateModel.Ids
			d = append(d, &properties)
		}
	}
	return
}

func ZksInstanceBulkUpdateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"data": {
			Description: `Data for bulk update`,
			Type:        schema.TypeList, //GoType: ZksInstanceBulkUpdateData
			Elem: &schema.Resource{
				Schema: ZksInstanceBulkUpdateDataSchema(),
			},
			Optional: true,
		},

		"ids": {
			Description: `List of ZKS instance IDs to be updated`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

func GetZksInstanceBulkUpdatePropertyFields() (t []string) {
	return []string{
		"data",
		"ids",
	}
}
