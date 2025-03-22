package schemas

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AssetGroupReadModel(d *schema.ResourceData) *models.AssetGroupRead {
	assetCountInt, _ := d.Get("asset_count").(int)
	assetCount := int64(assetCountInt)
	var assetIds *models.AssetIDs // AssetIDs
	assetIdsInterface, assetIdsIsSet := d.GetOk("asset_ids")
	if assetIdsIsSet && assetIdsInterface != nil {
		assetIdsMap := assetIdsInterface.([]interface{})
		if len(assetIdsMap) > 0 {
			assetIds = AssetIDsModelFromMap(assetIdsMap[0].(map[string]interface{}))
		}
	}
	var assetTags *models.AssetTags // AssetTags
	assetTagsInterface, assetTagsIsSet := d.GetOk("asset_tags")
	if assetTagsIsSet && assetTagsInterface != nil {
		assetTagsMap := assetTagsInterface.([]interface{})
		if len(assetTagsMap) > 0 {
			assetTags = AssetTagsModelFromMap(assetTagsMap[0].(map[string]interface{}))
		}
	}
	description, _ := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	projectID, _ := d.Get("project_id").(string)
	title, _ := d.Get("title").(string)
	return &models.AssetGroupRead{
		AssetCount:  assetCount,
		AssetIds:    assetIds,
		AssetTags:   assetTags,
		Description: description,
		ID:          id,
		Name:        name,
		ProjectID:   projectID,
		Title:       title,
	}
}

func AssetGroupReadModelFromMap(m map[string]interface{}) *models.AssetGroupRead {
	assetCount := int64(0)
	if m["asset_count"] != nil {
		assetCount = int64(m["asset_count"].(int))
	}
	var assetIds *models.AssetIDs // AssetIDs
	assetIdsInterface, assetIdsIsSet := m["asset_ids"]
	if assetIdsIsSet && assetIdsInterface != nil {
		assetIdsMap := assetIdsInterface.([]interface{})
		if len(assetIdsMap) > 0 {
			assetIds = AssetIDsModelFromMap(assetIdsMap[0].(map[string]interface{}))
		}
	}
	//
	var assetTags *models.AssetTags // AssetTags
	assetTagsInterface, assetTagsIsSet := m["asset_tags"]
	if assetTagsIsSet && assetTagsInterface != nil {
		assetTagsMap := assetTagsInterface.([]interface{})
		if len(assetTagsMap) > 0 {
			assetTags = AssetTagsModelFromMap(assetTagsMap[0].(map[string]interface{}))
		}
	}
	//
	description := m["description"].(string)
	id := m["id"].(string)
	name := m["name"].(string)
	projectID := m["project_id"].(string)
	title := m["title"].(string)
	return &models.AssetGroupRead{
		AssetCount:  assetCount,
		AssetIds:    assetIds,
		AssetTags:   assetTags,
		Description: description,
		ID:          id,
		Name:        name,
		ProjectID:   projectID,
		Title:       title,
	}
}

func SetAssetGroupReadResourceData(d *schema.ResourceData, m *models.AssetGroupRead) {
	if m.AssetCount != 0 { // Ensure AssetCount is valid before setting
		if err := d.Set("asset_count", m.AssetCount); err != nil {
			log.Printf("[ERROR] Failed to set asset_count: %s", err)
		}
	}
	if m.AssetIds != nil {
		if err := d.Set("asset_ids", SetAssetIDsSubResourceData([]*models.AssetIDs{m.AssetIds})); err != nil {
			log.Printf("[ERROR] Failed to set asset_ids: %s", err)
		}
	}
	if m.AssetTags != nil {
		if err := d.Set("asset_tags", SetAssetTagsSubResourceData([]*models.AssetTags{m.AssetTags})); err != nil {
			log.Printf("[ERROR] Failed to set asset_tags: %s", err)
		}
	}
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	d.Set("title", m.Title)
}

func SetAssetGroupReadSubResourceData(m []*models.AssetGroupRead) (d []*map[string]interface{}) {
	for _, AssetGroupReadModel := range m {
		if AssetGroupReadModel != nil {
			properties := make(map[string]interface{})
			properties["asset_count"] = AssetGroupReadModel.AssetCount
			properties["asset_ids"] = SetAssetIDsSubResourceData([]*models.AssetIDs{AssetGroupReadModel.AssetIds})
			properties["asset_tags"] = SetAssetTagsSubResourceData([]*models.AssetTags{AssetGroupReadModel.AssetTags})
			properties["description"] = AssetGroupReadModel.Description
			properties["id"] = AssetGroupReadModel.ID
			properties["name"] = AssetGroupReadModel.Name
			properties["project_id"] = AssetGroupReadModel.ProjectID
			properties["title"] = AssetGroupReadModel.Title
			d = append(d, &properties)
		}
	}
	return
}

func AssetGroupReadSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"asset_count": {
			Description: `Number of assets in the group`,
			Type:        schema.TypeInt,
			Computed:    true,
		},

		"asset_ids": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AssetIDs
			Elem: &schema.Resource{
				Schema: AssetIDsSchema(),
			},
			Optional: true,
		},

		"asset_tags": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AssetTags
			Elem: &schema.Resource{
				Schema: AssetTagsSchema(),
			},
			Optional: true,
		},

		"description": {
			Description: `Detailed description of the asset group.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `unique Id of the asset group.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `User defined name of the asset group, unique across the enterprise. Once asset group is created, name can’t be changed.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_id": {
			Description: `project id`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"title": {
			Description: `User defined title of the asset group. Title can be changed at any time.`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAssetGroupReadPropertyFields() (t []string) {
	return []string{
		"asset_count",
		"asset_ids",
		"asset_tags",
		"description",
		"id",
		"name",
		"project_id",
		"title",
	}
}
