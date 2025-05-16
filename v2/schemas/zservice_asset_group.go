package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZserviceAssetGroupModel(d *schema.ResourceData) *models.ZserviceAssetGroup {
	var assetIds *models.AssetGroupValueList // AssetGroupValueList
	assetIdsInterface, assetIdsIsSet := d.GetOk("asset_ids")
	if assetIdsIsSet && assetIdsInterface != nil {
		assetIdsMap := assetIdsInterface.([]interface{})
		if len(assetIdsMap) > 0 {
			assetIds = AssetGroupValueListModelFromMap(assetIdsMap[0].(map[string]interface{}))
		}
	}
	var assetTags *models.AssetGroupAssetTagList // AssetGroupAssetTagList
	assetTagsInterface, assetTagsIsSet := d.GetOk("asset_tags")
	if assetTagsIsSet && assetTagsInterface != nil {
		assetTagsMap := assetTagsInterface.([]interface{})
		if len(assetTagsMap) > 0 {
			assetTags = AssetGroupAssetTagListModelFromMap(assetTagsMap[0].(map[string]interface{}))
		}
	}
	var labels *models.AssetGroupValueList // AssetGroupValueList
	labelsInterface, labelsIsSet := d.GetOk("labels")
	if labelsIsSet && labelsInterface != nil {
		labelsMap := labelsInterface.([]interface{})
		if len(labelsMap) > 0 {
			labels = AssetGroupValueListModelFromMap(labelsMap[0].(map[string]interface{}))
		}
	}
	return &models.ZserviceAssetGroup{
		AssetIds:  assetIds,
		AssetTags: assetTags,
		Labels:    labels,
	}
}

func ZserviceAssetGroupModelFromMap(m map[string]interface{}) *models.ZserviceAssetGroup {
	var assetIds *models.AssetGroupValueList // AssetGroupValueList
	assetIdsInterface, assetIdsIsSet := m["asset_ids"]
	if assetIdsIsSet && assetIdsInterface != nil {
		assetIdsMap := assetIdsInterface.([]interface{})
		if len(assetIdsMap) > 0 {
			assetIds = AssetGroupValueListModelFromMap(assetIdsMap[0].(map[string]interface{}))
		}
	}
	//
	var assetTags *models.AssetGroupAssetTagList // AssetGroupAssetTagList
	assetTagsInterface, assetTagsIsSet := m["asset_tags"]
	if assetTagsIsSet && assetTagsInterface != nil {
		assetTagsMap := assetTagsInterface.([]interface{})
		if len(assetTagsMap) > 0 {
			assetTags = AssetGroupAssetTagListModelFromMap(assetTagsMap[0].(map[string]interface{}))
		}
	}
	//
	var labels *models.AssetGroupValueList // AssetGroupValueList
	labelsInterface, labelsIsSet := m["labels"]
	if labelsIsSet && labelsInterface != nil {
		labelsMap := labelsInterface.([]interface{})
		if len(labelsMap) > 0 {
			labels = AssetGroupValueListModelFromMap(labelsMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.ZserviceAssetGroup{
		AssetIds:  assetIds,
		AssetTags: assetTags,
		Labels:    labels,
	}
}

func SetZserviceAssetGroupResourceData(d *schema.ResourceData, m *models.ZserviceAssetGroup) {
	d.Set("asset_ids", SetAssetGroupValueListSubResourceData([]*models.AssetGroupValueList{m.AssetIds}))
	d.Set("asset_tags", SetAssetGroupAssetTagListSubResourceData([]*models.AssetGroupAssetTagList{m.AssetTags}))
	d.Set("labels", SetAssetGroupValueListSubResourceData([]*models.AssetGroupValueList{m.Labels}))
}

func SetZserviceAssetGroupSubResourceData(m []*models.ZserviceAssetGroup) (d []*map[string]interface{}) {
	for _, ZserviceAssetGroupModel := range m {
		if ZserviceAssetGroupModel != nil {
			properties := make(map[string]interface{})
			properties["asset_ids"] = SetAssetGroupValueListSubResourceData([]*models.AssetGroupValueList{ZserviceAssetGroupModel.AssetIds})
			properties["asset_tags"] = SetAssetGroupAssetTagListSubResourceData([]*models.AssetGroupAssetTagList{ZserviceAssetGroupModel.AssetTags})
			properties["labels"] = SetAssetGroupValueListSubResourceData([]*models.AssetGroupValueList{ZserviceAssetGroupModel.Labels})
			d = append(d, &properties)
		}
	}
	return
}

func ZserviceAssetGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"asset_ids": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AssetGroupValueList
			Elem: &schema.Resource{
				Schema: AssetGroupValueListSchema(),
			},
			Optional: true,
		},

		"asset_tags": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AssetGroupAssetTagList
			Elem: &schema.Resource{
				Schema: AssetGroupAssetTagListSchema(),
			},
			Optional: true,
		},

		"labels": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AssetGroupValueList
			Elem: &schema.Resource{
				Schema: AssetGroupValueListSchema(),
			},
			Optional: true,
		},
	}
}

func GetZserviceAssetGroupPropertyFields() (t []string) {
	return []string{
		"asset_ids",
		"asset_tags",
		"labels",
	}
}
