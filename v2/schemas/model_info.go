package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ModelInfoModel(d *schema.ResourceData) *models.ModelInfo {
	brandName, _ := d.Get("brand_name").(string)
	modelName, _ := d.Get("model_name").(string)
	return &models.ModelInfo{
		BrandName: brandName,
		ModelName: modelName,
	}
}

func ModelInfoModelFromMap(m map[string]interface{}) *models.ModelInfo {
	brandName := m["brand_name"].(string)
	modelName := m["model_name"].(string)
	return &models.ModelInfo{
		BrandName: brandName,
		ModelName: modelName,
	}
}

func SetModelInfoResourceData(d *schema.ResourceData, m *models.ModelInfo) {
	d.Set("brand_id", m.BrandID)
	d.Set("brand_name", m.BrandName)
	d.Set("model_id", m.ModelID)
	d.Set("model_name", m.ModelName)
}

func SetModelInfoSubResourceData(m []*models.ModelInfo) (d []*map[string]interface{}) {
	for _, ModelInfoModel := range m {
		if ModelInfoModel != nil {
			properties := make(map[string]interface{})
			properties["brand_id"] = ModelInfoModel.BrandID
			properties["brand_name"] = ModelInfoModel.BrandName
			properties["model_id"] = ModelInfoModel.ModelID
			properties["model_name"] = ModelInfoModel.ModelName
			d = append(d, &properties)
		}
	}
	return
}

func ModelInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"brand_id": {
			Description: `system generated unique id for a brand`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"brand_name": {
			Description: `brand name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"model_id": {
			Description: `system generated unique id for a model`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"model_name": {
			Description: `model name`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetModelInfoPropertyFields() (t []string) {
	return []string{
		"brand_name",
		"model_name",
	}
}
