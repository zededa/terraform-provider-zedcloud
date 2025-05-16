package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZserviceAssetGroupReadROModel(d *schema.ResourceData) *models.ZserviceAssetGroupReadRO {
	assetCountInt, _ := d.Get("asset_count").(int)
	assetCount := int64(assetCountInt)
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
	return &models.ZserviceAssetGroupReadRO{
		AssetCount:  assetCount,
		AssetIds:    assetIds,
		AssetTags:   assetTags,
		AssetType:   assetType,
		Description: description,
		ID:          id,
		Name:        name,
	}
}

func ZserviceAssetGroupReadROModelFromMap(m map[string]interface{}) *models.ZserviceAssetGroupReadRO {
	assetCount := int64(m["asset_count"].(int)) // int64
	var assetIds *models.ZserviceAssetIDs       // ZserviceAssetIDs
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
	return &models.ZserviceAssetGroupReadRO{
		AssetCount:  assetCount,
		AssetIds:    assetIds,
		AssetTags:   assetTags,
		AssetType:   assetType,
		Description: description,
		ID:          id,
		Name:        name,
	}
}

func SetZserviceAssetGroupReadROResourceData(d *schema.ResourceData, m *models.ZserviceAssetGroupReadRO) {
	d.Set("asset_count", m.AssetCount)
	d.Set("asset_ids", SetZserviceAssetIDsSubResourceData([]*models.ZserviceAssetIDs{m.AssetIds}))
	d.Set("asset_tags", SetZserviceAssetTagsSubResourceData([]*models.ZserviceAssetTags{m.AssetTags}))
	d.Set("asset_type", m.AssetType)
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
}

func SetZserviceAssetGroupReadROSubResourceData(m []*models.ZserviceAssetGroupReadRO) (d []*map[string]interface{}) {
	for _, ZserviceAssetGroupReadROModel := range m {
		if ZserviceAssetGroupReadROModel != nil {
			properties := make(map[string]interface{})
			properties["asset_count"] = ZserviceAssetGroupReadROModel.AssetCount
			properties["asset_ids"] = SetZserviceAssetIDsSubResourceData([]*models.ZserviceAssetIDs{ZserviceAssetGroupReadROModel.AssetIds})
			properties["asset_tags"] = SetZserviceAssetTagsSubResourceData([]*models.ZserviceAssetTags{ZserviceAssetGroupReadROModel.AssetTags})
			properties["asset_type"] = ZserviceAssetGroupReadROModel.AssetType
			properties["description"] = ZserviceAssetGroupReadROModel.Description
			properties["id"] = ZserviceAssetGroupReadROModel.ID
			properties["name"] = ZserviceAssetGroupReadROModel.Name
			d = append(d, &properties)
		}
	}
	return
}

func ZserviceAssetGroupReadROSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"asset_count": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

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

func GetZserviceAssetGroupReadROPropertyFields() (t []string) {
	return []string{
		"asset_count",
		"asset_ids",
		"asset_tags",
		"asset_type",
		"description",
		"id",
		"name",
	}
}
