package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AssetTagModel(d *schema.ResourceData) *models.AssetTag {
	tag := map[string]string{}
	tagInterface, tagIsSet := d.GetOk("tag")
	if tagIsSet {
		tagMap := tagInterface.(map[string]interface{})
		for k, v := range tagMap {
			if v == nil {
				continue
			}
			tag[k] = v.(string)
		}
	}

	return &models.AssetTag{
		Tag: tag,
	}
}

func AssetTagModelFromMap(m map[string]interface{}) *models.AssetTag {
	tag := map[string]string{}
	tagInterface, tagIsSet := m["tag"]
	if tagIsSet {
		tagMap := tagInterface.(map[string]interface{})
		for k, v := range tagMap {
			if v == nil {
				continue
			}
			tag[k] = v.(string)
		}
	}

	return &models.AssetTag{
		Tag: tag,
	}
}

func SetAssetTagResourceData(d *schema.ResourceData, m *models.AssetTag) {
	d.Set("tag", m.Tag)
}

func SetAssetTagSubResourceData(m []*models.AssetTag) (d []*map[string]interface{}) {
	for _, AssetTagModel := range m {
		if AssetTagModel != nil {
			properties := make(map[string]interface{})
			properties["tag"] = AssetTagModel.Tag
			d = append(d, &properties)
		}
	}
	return
}

func AssetTagSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"tag": {
			Description: `Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

func GetAssetTagPropertyFields() (t []string) {
	return []string{
		"tag",
	}
}
