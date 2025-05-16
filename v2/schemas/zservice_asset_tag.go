package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZserviceAssetTagModel(d *schema.ResourceData) *models.ZserviceAssetTag {
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

	return &models.ZserviceAssetTag{
		Tag: tag,
	}
}

func ZserviceAssetTagModelFromMap(m map[string]interface{}) *models.ZserviceAssetTag {
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

	return &models.ZserviceAssetTag{
		Tag: tag,
	}
}

func SetZserviceAssetTagResourceData(d *schema.ResourceData, m *models.ZserviceAssetTag) {
	d.Set("tag", m.Tag)
}

func SetZserviceAssetTagSubResourceData(m []*models.ZserviceAssetTag) (d []*map[string]interface{}) {
	for _, ZserviceAssetTagModel := range m {
		if ZserviceAssetTagModel != nil {
			properties := make(map[string]interface{})
			properties["tag"] = ZserviceAssetTagModel.Tag
			d = append(d, &properties)
		}
	}
	return
}

func ZserviceAssetTagSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"tag": {
			Description: ``,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

func GetZserviceAssetTagPropertyFields() (t []string) {
	return []string{
		"tag",
	}
}
