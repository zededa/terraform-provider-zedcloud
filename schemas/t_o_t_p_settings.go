package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func TOTPSettingsModel(d *schema.ResourceData) *models.TOTPSettings {
	allowOverwriteByChildren, _ := d.Get("allow_overwrite_by_children").(bool)
	enforce, _ := d.Get("enforce").(bool)
	setForChildren, _ := d.Get("set_for_children").(bool)
	return &models.TOTPSettings{
		AllowOverwriteByChildren: allowOverwriteByChildren,
		Enforce:                  enforce,
		SetForChildren:           setForChildren,
	}
}

func TOTPSettingsModelFromMap(m map[string]interface{}) *models.TOTPSettings {
	allowOverwriteByChildren := m["allow_overwrite_by_children"].(bool)
	enforce := m["enforce"].(bool)
	setForChildren := m["set_for_children"].(bool)
	return &models.TOTPSettings{
		AllowOverwriteByChildren: allowOverwriteByChildren,
		Enforce:                  enforce,
		SetForChildren:           setForChildren,
	}
}

func SetTOTPSettingsResourceData(d *schema.ResourceData, m *models.TOTPSettings) {
	d.Set("allow_overwrite_by_children", m.AllowOverwriteByChildren)
	d.Set("enforce", m.Enforce)
	d.Set("set_for_children", m.SetForChildren)
}

func SetTOTPSettingsSubResourceData(m []*models.TOTPSettings) (d []*map[string]interface{}) {
	for _, TOTPSettingsModel := range m {
		if TOTPSettingsModel != nil {
			properties := make(map[string]interface{})
			properties["allow_overwrite_by_children"] = TOTPSettingsModel.AllowOverwriteByChildren
			properties["enforce"] = TOTPSettingsModel.Enforce
			properties["set_for_children"] = TOTPSettingsModel.SetForChildren
			d = append(d, &properties)
		}
	}
	return
}

func TOTPSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_overwrite_by_children": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"enforce": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"set_for_children": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}

func GetTOTPSettingsPropertyFields() (t []string) {
	return []string{
		"allow_overwrite_by_children",
		"enforce",
		"set_for_children",
	}
}
