package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZksInstanceK3sBulkUpgradeModel(d *schema.ResourceData) *models.ZksInstanceK3sBulkUpgrade {
	var data *models.ZksInstanceK3sBulkUpgradeData // ZksInstanceK3sBulkUpgradeData
	dataInterface, dataIsSet := d.GetOk("data")
	if dataIsSet && dataInterface != nil {
		dataMap := dataInterface.([]interface{})
		if len(dataMap) > 0 {
			data = ZksInstanceK3sBulkUpgradeDataModelFromMap(dataMap[0].(map[string]interface{}))
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
	return &models.ZksInstanceK3sBulkUpgrade{
		Data: data,
		Ids:  ids,
	}
}

func ZksInstanceK3sBulkUpgradeModelFromMap(m map[string]interface{}) *models.ZksInstanceK3sBulkUpgrade {
	var data *models.ZksInstanceK3sBulkUpgradeData // ZksInstanceK3sBulkUpgradeData
	dataInterface, dataIsSet := m["data"]
	if dataIsSet && dataInterface != nil {
		dataMap := dataInterface.([]interface{})
		if len(dataMap) > 0 {
			data = ZksInstanceK3sBulkUpgradeDataModelFromMap(dataMap[0].(map[string]interface{}))
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
	return &models.ZksInstanceK3sBulkUpgrade{
		Data: data,
		Ids:  ids,
	}
}

func SetZksInstanceK3sBulkUpgradeResourceData(d *schema.ResourceData, m *models.ZksInstanceK3sBulkUpgrade) {
	d.Set("data", SetZksInstanceK3sBulkUpgradeDataSubResourceData([]*models.ZksInstanceK3sBulkUpgradeData{m.Data}))
	d.Set("ids", m.Ids)
}

func SetZksInstanceK3sBulkUpgradeSubResourceData(m []*models.ZksInstanceK3sBulkUpgrade) (d []*map[string]interface{}) {
	for _, ZksInstanceK3sBulkUpgradeModel := range m {
		if ZksInstanceK3sBulkUpgradeModel != nil {
			properties := make(map[string]interface{})
			properties["data"] = SetZksInstanceK3sBulkUpgradeDataSubResourceData([]*models.ZksInstanceK3sBulkUpgradeData{ZksInstanceK3sBulkUpgradeModel.Data})
			properties["ids"] = ZksInstanceK3sBulkUpgradeModel.Ids
			d = append(d, &properties)
		}
	}
	return
}

func ZksInstanceK3sBulkUpgradeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"data": {
			Description: `Data for bulk upgrade`,
			Type:        schema.TypeList, //GoType: ZksInstanceK3sBulkUpgradeData
			Elem: &schema.Resource{
				Schema: ZksInstanceK3sBulkUpgradeDataSchema(),
			},
			Optional: true,
		},

		"ids": {
			Description: `List of ZKS instance IDs to be upgraded`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

func GetZksInstanceK3sBulkUpgradePropertyFields() (t []string) {
	return []string{
		"data",
		"ids",
	}
}
