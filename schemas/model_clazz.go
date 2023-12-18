package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func ModelClazzModel(d *schema.ResourceData) *models.ModelClazz {
	modelClazz, _ := d.Get("model_clazz").(models.ModelClazz)
	return &modelClazz
}

func ModelClazzModelFromMap(m map[string]interface{}) *models.ModelClazz {
	modelClazz := m["model_clazz"].(models.ModelClazz)
	return &modelClazz
}

func SetModelClazzResourceData(d *schema.ResourceData, m *models.ModelClazz) {
}

func SetModelClazzSubResourceData(m []*models.ModelClazz) (d []*map[string]interface{}) {
	for _, ModelClazzModel := range m {
		if ModelClazzModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func ModelClazzSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetModelClazzPropertyFields() (t []string) {
	return []string{}
}
