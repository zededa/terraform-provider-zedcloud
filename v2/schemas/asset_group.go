package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AssetGroupModel(d *schema.ResourceData) *models.AssetGroup {
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
	return &models.AssetGroup{
		AssetIds:    assetIds,
		AssetTags:   assetTags,
		Description: description,
		ID:          id,
		Name:        &name, // string
		ProjectID:   projectID,
		Title:       title,
	}
}

func AssetGroupModelFromMap(m map[string]interface{}) *models.AssetGroup {
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
	return &models.AssetGroup{
		AssetIds:    assetIds,
		AssetTags:   assetTags,
		Description: description,
		ID:          id,
		Name:        &name,
		ProjectID:   projectID,
		Title:       title,
	}
}

func SetAssetGroupResourceData(d *schema.ResourceData, m *models.AssetGroup) {
	d.Set("asset_ids", SetAssetIDsSubResourceData([]*models.AssetIDs{m.AssetIds}))
	d.Set("asset_tags", SetAssetTagsSubResourceData([]*models.AssetTags{m.AssetTags}))
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	d.Set("title", m.Title)
}

func SetAssetGroupSubResourceData(m []*models.AssetGroup) (d []*map[string]interface{}) {
	for _, AssetGroupModel := range m {
		if AssetGroupModel != nil {
			properties := make(map[string]interface{})
			properties["asset_ids"] = SetAssetIDsSubResourceData([]*models.AssetIDs{AssetGroupModel.AssetIds})
			properties["asset_tags"] = SetAssetTagsSubResourceData([]*models.AssetTags{AssetGroupModel.AssetTags})
			properties["description"] = AssetGroupModel.Description
			properties["id"] = AssetGroupModel.ID
			properties["name"] = AssetGroupModel.Name
			properties["project_id"] = AssetGroupModel.ProjectID
			properties["title"] = AssetGroupModel.Title
			d = append(d, &properties)
		}
	}
	return
}

func AssetGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
			Required:    true,
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

func GetAssetGroupPropertyFields() (t []string) {
	return []string{
		"asset_ids",
		"asset_tags",
		"description",
		"id",
		"name",
		"project_id",
		"title",
	}
}
