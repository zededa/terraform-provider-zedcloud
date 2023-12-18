package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func OpaqueObjectCategoryModel(d *schema.ResourceData) *models.OpaqueObjectCategory {
	opaqueObjectCategory, _ := d.Get("opaque_object_category").(models.OpaqueObjectCategory)
	return &opaqueObjectCategory
}

func OpaqueObjectCategoryModelFromMap(m map[string]interface{}) *models.OpaqueObjectCategory {
	opaqueObjectCategory := m["opaque_object_category"].(models.OpaqueObjectCategory)
	return &opaqueObjectCategory
}

func SetOpaqueObjectCategoryResourceData(d *schema.ResourceData, m *models.OpaqueObjectCategory) {
}

func SetOpaqueObjectCategorySubResourceData(m []*models.OpaqueObjectCategory) (d []*map[string]interface{}) {
	for _, OpaqueObjectCategoryModel := range m {
		if OpaqueObjectCategoryModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func OpaqueObjectCategorySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetOpaqueObjectCategoryPropertyFields() (t []string) {
	return []string{}
}
