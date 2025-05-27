package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZserviceAssetGroupConfigModel(d *schema.ResourceData) *models.ZserviceAssetGroupConfig {
	var assetIds *models.ZserviceAssetIDs // ZserviceAssetIDs
	assetIdsInterface, assetIdsIsSet := d.GetOk("asset_ids")
	if assetIdsIsSet && assetIdsInterface != nil {
		assetIdsMap := assetIdsInterface.([]interface{})
		if len(assetIdsMap) > 0 {
			assetIds = ZserviceAssetIDsModelFromMap(assetIdsMap[0].(map[string]interface{}))
		}
	}
	var assetTags *models.ZserviceAssetTags // ZserviceAssetTags
	assetTagsInterface, assetTagsIsSet := d.GetOk("asset_tags")
	if assetTagsIsSet && assetTagsInterface != nil {
		assetTagsMap := assetTagsInterface.([]interface{})
		if len(assetTagsMap) > 0 {
			assetTags = ZserviceAssetTagsModelFromMap(assetTagsMap[0].(map[string]interface{}))
		}
	}
	var assetType *models.ZserviceAssetType // ZserviceAssetType
	assetTypeInterface, assetTypeIsSet := d.GetOk("asset_type")
	if assetTypeIsSet {
		assetTypeModel := assetTypeInterface.(string)
		assetType = models.NewZserviceAssetType(models.ZserviceAssetType(assetTypeModel))
	}
	description, _ := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	return &models.ZserviceAssetGroupConfig{
		AssetIds:    assetIds,
		AssetTags:   assetTags,
		AssetType:   assetType,
		Description: description,
		ID:          id,
		Name:        name,
	}
}

func ZserviceAssetGroupConfigModelFromMap(m map[string]interface{}) *models.ZserviceAssetGroupConfig {
	var assetIds *models.ZserviceAssetIDs // ZserviceAssetIDs
	assetIdsInterface, assetIdsIsSet := m["asset_ids"]
	if assetIdsIsSet && assetIdsInterface != nil {
		assetIdsMap := assetIdsInterface.([]interface{})
		if len(assetIdsMap) > 0 {
			assetIds = ZserviceAssetIDsModelFromMap(assetIdsMap[0].(map[string]interface{}))
		}
	}
	//
	var assetTags *models.ZserviceAssetTags // ZserviceAssetTags
	assetTagsInterface, assetTagsIsSet := m["asset_tags"]
	if assetTagsIsSet && assetTagsInterface != nil {
		assetTagsMap := assetTagsInterface.([]interface{})
		if len(assetTagsMap) > 0 {
			assetTags = ZserviceAssetTagsModelFromMap(assetTagsMap[0].(map[string]interface{}))
		}
	}
	//
	var assetType *models.ZserviceAssetType // ZserviceAssetType
	assetTypeInterface, assetTypeIsSet := m["asset_type"]
	if assetTypeIsSet {
		assetTypeModel := assetTypeInterface.(string)
		assetType = models.NewZserviceAssetType(models.ZserviceAssetType(assetTypeModel))
	}
	description := m["description"].(string)
	id := m["id"].(string)
	name := m["name"].(string)
	return &models.ZserviceAssetGroupConfig{
		AssetIds:    assetIds,
		AssetTags:   assetTags,
		AssetType:   assetType,
		Description: description,
		ID:          id,
		Name:        name,
	}
}

func SetZserviceAssetGroupConfigResourceData(d *schema.ResourceData, m *models.ZserviceAssetGroupConfig) {
	d.Set("asset_ids", SetZserviceAssetIDsSubResourceData([]*models.ZserviceAssetIDs{m.AssetIds}))
	d.Set("asset_tags", SetZserviceAssetTagsSubResourceData([]*models.ZserviceAssetTags{m.AssetTags}))
	d.Set("asset_type", m.AssetType)
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
}

func SetZserviceAssetGroupConfigSubResourceData(m []*models.ZserviceAssetGroupConfig) (d []*map[string]interface{}) {
	for _, ZserviceAssetGroupConfigModel := range m {
		if ZserviceAssetGroupConfigModel != nil {
			properties := make(map[string]interface{})
			properties["asset_ids"] = SetZserviceAssetIDsSubResourceData([]*models.ZserviceAssetIDs{ZserviceAssetGroupConfigModel.AssetIds})
			properties["asset_tags"] = SetZserviceAssetTagsSubResourceData([]*models.ZserviceAssetTags{ZserviceAssetGroupConfigModel.AssetTags})
			properties["asset_type"] = ZserviceAssetGroupConfigModel.AssetType
			properties["description"] = ZserviceAssetGroupConfigModel.Description
			properties["id"] = ZserviceAssetGroupConfigModel.ID
			properties["name"] = ZserviceAssetGroupConfigModel.Name
			d = append(d, &properties)
		}
	}
	return
}

func ZserviceAssetGroupConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"asset_ids": {
			Description: ``,
			Type:        schema.TypeList, //GoType: ZserviceAssetIDs
			Elem: &schema.Resource{
				Schema: ZserviceAssetIDsSchema(),
			},
			Optional: true,
		},

		"asset_tags": {
			Description: ``,
			Type:        schema.TypeList, //GoType: ZserviceAssetTags
			Elem: &schema.Resource{
				Schema: ZserviceAssetTagsSchema(),
			},
			Optional: true,
		},

		"asset_type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"description": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: ``,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetZserviceAssetGroupConfigPropertyFields() (t []string) {
	return []string{
		"asset_ids",
		"asset_tags",
		"asset_type",
		"description",
		"id",
		"name",
	}
}
