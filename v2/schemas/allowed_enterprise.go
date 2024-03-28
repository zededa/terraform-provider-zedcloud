package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AllowedEnterpriseModel(d *schema.ResourceData) *models.AllowedEnterprise {
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	roleID, _ := d.Get("role_id").(string)
	return &models.AllowedEnterprise{
		ID:     id,
		Name:   name,
		RoleID: roleID,
	}
}

func AllowedEnterpriseModelFromMap(m map[string]interface{}) *models.AllowedEnterprise {
	id := m["id"].(string)
	name := m["name"].(string)
	roleID := m["role_id"].(string)
	return &models.AllowedEnterprise{
		ID:     id,
		Name:   name,
		RoleID: roleID,
	}
}

func SetAllowedEnterpriseResourceData(d *schema.ResourceData, m *models.AllowedEnterprise) {
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("role_id", m.RoleID)
}

func SetAllowedEnterpriseSubResourceData(m []*models.AllowedEnterprise) (d []*map[string]interface{}) {
	for _, AllowedEnterpriseModel := range m {
		if AllowedEnterpriseModel != nil {
			properties := make(map[string]interface{})
			properties["id"] = AllowedEnterpriseModel.ID
			properties["name"] = AllowedEnterpriseModel.Name
			properties["role_id"] = AllowedEnterpriseModel.RoleID
			d = append(d, &properties)
		}
	}
	return
}

func AllowedEnterpriseSchema() map[string]*schema.Schema {
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

		"role_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAllowedEnterprisePropertyFields() (t []string) {
	return []string{
		"id",
		"name",
		"role_id",
	}
}
