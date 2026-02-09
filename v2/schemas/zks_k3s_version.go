package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZksK3sVersionModel(d *schema.ResourceData) *models.ZksK3sVersion {
	id, _ := d.Get("id").(string)
	typeVar, _ := d.Get("type").(string)
	version, _ := d.Get("version").(string)
	return &models.ZksK3sVersion{
		ID:      id,
		Type:    typeVar,
		Version: version,
	}
}

func ZksK3sVersionModelFromMap(m map[string]interface{}) *models.ZksK3sVersion {
	id := m["id"].(string)
	typeVar := m["type"].(string)
	version := m["version"].(string)
	return &models.ZksK3sVersion{
		ID:      id,
		Type:    typeVar,
		Version: version,
	}
}

func SetZksK3sVersionResourceData(d *schema.ResourceData, m *models.ZksK3sVersion) {
	d.Set("id", m.ID)
	d.Set("type", m.Type)
	d.Set("version", m.Version)
}

func SetZksK3sVersionSubResourceData(m []*models.ZksK3sVersion) (d []*map[string]interface{}) {
	for _, ZksK3sVersionModel := range m {
		if ZksK3sVersionModel != nil {
			properties := make(map[string]interface{})
			properties["id"] = ZksK3sVersionModel.ID
			properties["type"] = ZksK3sVersionModel.Type
			properties["version"] = ZksK3sVersionModel.Version
			d = append(d, &properties)
		}
	}
	return
}

func ZksK3sVersionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: `Unique identifier for the K3s version`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"type": {
			Description: `Type of the K3s version (e.g., release, beta)`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"version": {
			Description: `K3s version string`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetZksK3sVersionPropertyFields() (t []string) {
	return []string{
		"id",
		"type",
		"version",
	}
}
