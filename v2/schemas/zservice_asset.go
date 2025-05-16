package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZserviceAssetModel(d *schema.ResourceData) *models.ZserviceAsset {
	groupID, _ := d.Get("group_id").(string)
	var targetSelector *models.ZserviceAssetGroup // ZserviceAssetGroup
	targetSelectorInterface, targetSelectorIsSet := d.GetOk("target_selector")
	if targetSelectorIsSet && targetSelectorInterface != nil {
		targetSelectorMap := targetSelectorInterface.([]interface{})
		if len(targetSelectorMap) > 0 {
			targetSelector = ZserviceAssetGroupModelFromMap(targetSelectorMap[0].(map[string]interface{}))
		}
	}
	return &models.ZserviceAsset{
		GroupID:        groupID,
		TargetSelector: targetSelector,
	}
}

func ZserviceAssetModelFromMap(m map[string]interface{}) *models.ZserviceAsset {
	groupID := m["group_id"].(string)
	var targetSelector *models.ZserviceAssetGroup // ZserviceAssetGroup
	targetSelectorInterface, targetSelectorIsSet := m["target_selector"]
	if targetSelectorIsSet && targetSelectorInterface != nil {
		targetSelectorMap := targetSelectorInterface.([]interface{})
		if len(targetSelectorMap) > 0 {
			targetSelector = ZserviceAssetGroupModelFromMap(targetSelectorMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.ZserviceAsset{
		GroupID:        groupID,
		TargetSelector: targetSelector,
	}
}

func SetZserviceAssetResourceData(d *schema.ResourceData, m *models.ZserviceAsset) {
	d.Set("group_id", m.GroupID)
	d.Set("target_selector", SetZserviceAssetGroupSubResourceData([]*models.ZserviceAssetGroup{m.TargetSelector}))
}

func SetZserviceAssetSubResourceData(m []*models.ZserviceAsset) (d []*map[string]interface{}) {
	for _, ZserviceAssetModel := range m {
		if ZserviceAssetModel != nil {
			properties := make(map[string]interface{})
			properties["group_id"] = ZserviceAssetModel.GroupID
			properties["target_selector"] = SetZserviceAssetGroupSubResourceData([]*models.ZserviceAssetGroup{ZserviceAssetModel.TargetSelector})
			d = append(d, &properties)
		}
	}
	return
}

func ZserviceAssetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"group_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"target_selector": {
			Description: ``,
			Type:        schema.TypeList, //GoType: ZserviceAssetGroup
			Elem: &schema.Resource{
				Schema: ZserviceAssetGroupSchema(),
			},
			Optional: true,
		},
	}
}

func GetZserviceAssetPropertyFields() (t []string) {
	return []string{
		"group_id",
		"target_selector",
	}
}
