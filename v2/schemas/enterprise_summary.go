package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func EnterpriseSummaryModel(d *schema.ResourceData) *models.EnterpriseSummary {
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	return &models.EnterpriseSummary{
		ID:   id,
		Name: name,
	}
}

func EnterpriseSummaryModelFromMap(m map[string]interface{}) *models.EnterpriseSummary {
	id := m["id"].(string)
	name := m["name"].(string)
	return &models.EnterpriseSummary{
		ID:   id,
		Name: name,
	}
}

func SetEnterpriseSummaryResourceData(d *schema.ResourceData, m *models.EnterpriseSummary) {
	d.Set("id", m.ID)
	d.Set("name", m.Name)
}

func SetEnterpriseSummarySubResourceData(m []*models.EnterpriseSummary) (d []*map[string]interface{}) {
	for _, EnterpriseSummaryModel := range m {
		if EnterpriseSummaryModel != nil {
			properties := make(map[string]interface{})
			properties["id"] = EnterpriseSummaryModel.ID
			properties["name"] = EnterpriseSummaryModel.Name
			d = append(d, &properties)
		}
	}
	return
}

func EnterpriseSummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: ``,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetEnterpriseSummaryPropertyFields() (t []string) {
	return []string{
		"id",
		"name",
	}
}
