package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZksInstanceBulkUpdateDataModel(d *schema.ResourceData) *models.ZksInstanceBulkUpdateData {
	tags := map[string]string{}
	tagsInterface, tagsIsSet := d.GetOk("tags")
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	return &models.ZksInstanceBulkUpdateData{
		Tags: tags,
	}
}

func ZksInstanceBulkUpdateDataModelFromMap(m map[string]interface{}) *models.ZksInstanceBulkUpdateData {
	tags := map[string]string{}
	tagsInterface, tagsIsSet := m["tags"]
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	return &models.ZksInstanceBulkUpdateData{
		Tags: tags,
	}
}

func SetZksInstanceBulkUpdateDataResourceData(d *schema.ResourceData, m *models.ZksInstanceBulkUpdateData) {
	d.Set("tags", m.Tags)
}

func SetZksInstanceBulkUpdateDataSubResourceData(m []*models.ZksInstanceBulkUpdateData) (d []*map[string]interface{}) {
	for _, ZksInstanceBulkUpdateDataModel := range m {
		if ZksInstanceBulkUpdateDataModel != nil {
			properties := make(map[string]interface{})
			properties["tags"] = ZksInstanceBulkUpdateDataModel.Tags
			d = append(d, &properties)
		}
	}
	return
}

func ZksInstanceBulkUpdateDataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"tags": {
			Description: `New tags for the ZKS instance`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

func GetZksInstanceBulkUpdateDataPropertyFields() (t []string) {
	return []string{
		"tags",
	}
}
