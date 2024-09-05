package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func TOTPSettingsModel(d *schema.ResourceData) *models.TOTPSettings {
	enforce, _ := d.Get("enforce").(bool)
	enforceInChildren, _ := d.Get("enforce_in_children").(bool)
	enforcedByParent, _ := d.Get("enforced_by_parent").(bool)
	return &models.TOTPSettings{
		Enforce:           enforce,
		EnforceInChildren: enforceInChildren,
		EnforcedByParent:  enforcedByParent,
	}
}

func TOTPSettingsModelFromMap(m map[string]interface{}) *models.TOTPSettings {
	enforce := m["enforce"].(bool)
	enforceInChildren := m["enforce_in_children"].(bool)
	enforcedByParent := m["enforced_by_parent"].(bool)
	return &models.TOTPSettings{
		Enforce:           enforce,
		EnforceInChildren: enforceInChildren,
		EnforcedByParent:  enforcedByParent,
	}
}

func SetTOTPSettingsResourceData(d *schema.ResourceData, m *models.TOTPSettings) {
	d.Set("enforce", m.Enforce)
	d.Set("enforce_in_children", m.EnforceInChildren)
	d.Set("enforced_by_parent", m.EnforcedByParent)
}

func SetTOTPSettingsSubResourceData(m []*models.TOTPSettings) (d []*map[string]interface{}) {
	for _, TOTPSettingsModel := range m {
		if TOTPSettingsModel != nil {
			properties := make(map[string]interface{})
			properties["enforce"] = TOTPSettingsModel.Enforce
			properties["enforce_in_children"] = TOTPSettingsModel.EnforceInChildren
			properties["enforced_by_parent"] = TOTPSettingsModel.EnforcedByParent
			d = append(d, &properties)
		}
	}
	return
}

func TOTPSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enforce": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"enforce_in_children": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"enforced_by_parent": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}

func GetTOTPSettingsPropertyFields() (t []string) {
	return []string{
		"enforce",
		"enforce_in_children",
		"enforced_by_parent",
	}
}
