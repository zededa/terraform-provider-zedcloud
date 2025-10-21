package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AppProfileFilterModel(d *schema.ResourceData) *models.AppProfileFilter {
	id, _ := d.Get("id").(string)
	namePattern, _ := d.Get("name_pattern").(string)
	return &models.AppProfileFilter{
		ID:          id,
		NamePattern: namePattern,
	}
}

func AppProfileFilterModelFromMap(m map[string]interface{}) *models.AppProfileFilter {
	id := m["id"].(string)
	namePattern := m["name_pattern"].(string)
	return &models.AppProfileFilter{
		ID:          id,
		NamePattern: namePattern,
	}
}

func SetAppProfileFilterResourceData(d *schema.ResourceData, m *models.AppProfileFilter) {
	d.Set("id", m.ID)
	d.Set("name_pattern", m.NamePattern)
}

func SetAppProfileFilterSubResourceData(m []*models.AppProfileFilter) (d []*map[string]interface{}) {
	for _, AppProfileFilterModel := range m {
		if AppProfileFilterModel != nil {
			properties := make(map[string]interface{})
			properties["id"] = AppProfileFilterModel.ID
			properties["name_pattern"] = AppProfileFilterModel.NamePattern
			d = append(d, &properties)
		}
	}
	return
}

func AppProfileFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: `app profile id`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name_pattern": {
			Description: `app profile name pattern`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAppProfileFilterPropertyFields() (t []string) {
	return []string{
		"id",
		"name_pattern",
	}
}
