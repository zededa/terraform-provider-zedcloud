package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func LocalOperatorConsolePolicyModel(d *schema.ResourceData) *models.LocalOperatorConsolePolicy {
	id, _ := d.Get("id").(string)
	locURL, _ := d.Get("loc_url").(string)
	return &models.LocalOperatorConsolePolicy{
		ID:     id,
		LocURL: &locURL, // string true false false
	}
}

func LocalOperatorConsolePolicyModelFromMap(m map[string]interface{}) *models.LocalOperatorConsolePolicy {
	id := m["id"].(string)
	locURL := m["loc_url"].(string)
	return &models.LocalOperatorConsolePolicy{
		ID:     id,
		LocURL: &locURL,
	}
}

func SetLocalOperatorConsolePolicyResourceData(d *schema.ResourceData, m *models.LocalOperatorConsolePolicy) {
	d.Set("id", m.ID)
	d.Set("loc_url", m.LocURL)
}

func SetLocalOperatorConsolePolicySubResourceData(m []*models.LocalOperatorConsolePolicy) (d []*map[string]interface{}) {
	for _, LocalOperatorConsolePolicyModel := range m {
		if LocalOperatorConsolePolicyModel != nil {
			properties := make(map[string]interface{})
			properties["id"] = LocalOperatorConsolePolicyModel.ID
			properties["loc_url"] = LocalOperatorConsolePolicyModel.LocURL
			d = append(d, &properties)
		}
	}
	return
}

func LocalOperatorConsolePolicy() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: `unique policy id`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"loc_url": {
			Description: `Local operator console URL`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func GetLocalOperatorConsolePolicyPropertyFields() (t []string) {
	return []string{
		"id",
		"loc_url",
	}
}
