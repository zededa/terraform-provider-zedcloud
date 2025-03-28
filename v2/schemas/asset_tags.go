package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AssetTagsModel(d *schema.ResourceData) *models.AssetTags {
	var assetTag []*models.AssetTag // []*AssetTag
	assetTagInterface, assetTagIsSet := d.GetOk("asset_tag")
	if assetTagIsSet {
		var items []interface{}
		if listItems, isList := assetTagInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = assetTagInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AssetTagModelFromMap(v.(map[string]interface{}))
			assetTag = append(assetTag, m)
		}
	}
	return &models.AssetTags{
		AssetTag: assetTag,
	}
}

func AssetTagsModelFromMap(m map[string]interface{}) *models.AssetTags {
	var assetTag []*models.AssetTag // []*AssetTag
	assetTagInterface, assetTagIsSet := m["asset_tag"]
	if assetTagIsSet {
		var items []interface{}
		if listItems, isList := assetTagInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = assetTagInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AssetTagModelFromMap(v.(map[string]interface{}))
			assetTag = append(assetTag, m)
		}
	}
	return &models.AssetTags{
		AssetTag: assetTag,
	}
}

func SetAssetTagsResourceData(d *schema.ResourceData, m *models.AssetTags) {
	d.Set("asset_tag", SetAssetTagSubResourceData(m.AssetTag))
}

func SetAssetTagsSubResourceData(m []*models.AssetTags) (d []*map[string]interface{}) {
	for _, AssetTagsModel := range m {
		if AssetTagsModel != nil {
			properties := make(map[string]interface{})
			properties["asset_tag"] = SetAssetTagSubResourceData(AssetTagsModel.AssetTag)
			d = append(d, &properties)
		}
	}
	return
}

func AssetTagsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"asset_tag": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*AssetTag
			Elem: &schema.Resource{
				Schema: AssetTagSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

func GetAssetTagsPropertyFields() (t []string) {
	return []string{
		"asset_tag",
	}
}
