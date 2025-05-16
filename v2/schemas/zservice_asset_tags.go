package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZserviceAssetTagsModel(d *schema.ResourceData) *models.ZserviceAssetTags {
	var assetTag []*models.ZserviceAssetTag // []*ZserviceAssetTag
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
			m := ZserviceAssetTagModelFromMap(v.(map[string]interface{}))
			assetTag = append(assetTag, m)
		}
	}
	return &models.ZserviceAssetTags{
		AssetTag: assetTag,
	}
}

func ZserviceAssetTagsModelFromMap(m map[string]interface{}) *models.ZserviceAssetTags {
	var assetTag []*models.ZserviceAssetTag // []*ZserviceAssetTag
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
			m := ZserviceAssetTagModelFromMap(v.(map[string]interface{}))
			assetTag = append(assetTag, m)
		}
	}
	return &models.ZserviceAssetTags{
		AssetTag: assetTag,
	}
}

func SetZserviceAssetTagsResourceData(d *schema.ResourceData, m *models.ZserviceAssetTags) {
	d.Set("asset_tag", SetZserviceAssetTagSubResourceData(m.AssetTag))
}

func SetZserviceAssetTagsSubResourceData(m []*models.ZserviceAssetTags) (d []*map[string]interface{}) {
	for _, ZserviceAssetTagsModel := range m {
		if ZserviceAssetTagsModel != nil {
			properties := make(map[string]interface{})
			properties["asset_tag"] = SetZserviceAssetTagSubResourceData(ZserviceAssetTagsModel.AssetTag)
			d = append(d, &properties)
		}
	}
	return
}

func ZserviceAssetTagsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"asset_tag": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*ZserviceAssetTag
			Elem: &schema.Resource{
				Schema: ZserviceAssetTagSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

func GetZserviceAssetTagsPropertyFields() (t []string) {
	return []string{
		"asset_tag",
	}
}
