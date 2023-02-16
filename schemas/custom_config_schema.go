package schemas

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func CustomConfigModel(d *schema.ResourceData) *models.CustomConfig {
	add, _ := d.Get("add").(bool)
	allowStorageResize, _ := d.Get("allow_storage_resize").(bool)
	fieldDelimiter, _ := d.Get("field_delimiter").(string)
	name, _ := d.Get("name").(string)
	override, _ := d.Get("override").(bool)
	template, _ := d.Get("template").(string)
	var variableGroups []*models.CustomConfigVariableGroup // []*CustomConfigVariableGroup
	variableGroupsInterface, variableGroupsIsSet := d.GetOk("variable_groups")
	if variableGroupsIsSet {
		var items []interface{}
		if listItems, isList := variableGroupsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = variableGroupsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := CustomConfigVariableGroupModelFromMap(v.(map[string]interface{}))
			variableGroups = append(variableGroups, m)
		}
	}
	c := &models.CustomConfig{
		Add:                add,
		AllowStorageResize: allowStorageResize,
		FieldDelimiter:     fieldDelimiter,
		Name:               name,
		Override:           override,
		Template:           template,
		VariableGroups:     variableGroups,
	}
	spew.Dump(c)
	return c
}

func CustomConfigModelFromMap(m map[string]interface{}) *models.CustomConfig {
	add := m["add"].(bool)
	allowStorageResize := m["allow_storage_resize"].(bool)
	fieldDelimiter := m["field_delimiter"].(string)
	name := m["name"].(string)
	override := m["override"].(bool)
	template := m["template"].(string)
	var variableGroups []*models.CustomConfigVariableGroup // []*CustomConfigVariableGroup
	variableGroupsInterface, variableGroupsIsSet := m["variable_groups"]
	if variableGroupsIsSet {
		var items []interface{}
		if listItems, isList := variableGroupsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = variableGroupsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := CustomConfigVariableGroupModelFromMap(v.(map[string]interface{}))
			variableGroups = append(variableGroups, m)
		}
	}
	return &models.CustomConfig{
		Add:                add,
		AllowStorageResize: allowStorageResize,
		FieldDelimiter:     fieldDelimiter,
		Name:               name,
		Override:           override,
		Template:           template,
		VariableGroups:     variableGroups,
	}
}

func SetCustomConfigResourceData(d *schema.ResourceData, m *models.CustomConfig) {
	d.Set("add", m.Add)
	d.Set("allow_storage_resize", m.AllowStorageResize)
	d.Set("field_delimiter", m.FieldDelimiter)
	d.Set("name", m.Name)
	d.Set("override", m.Override)
	d.Set("template", m.Template)
	d.Set("variable_groups", SetCustomConfigVariableGroupSubResourceData(m.VariableGroups))
}

func SetCustomConfigSubResourceData(m []*models.CustomConfig) (d []*map[string]interface{}) {
	for _, CustomConfigModel := range m {
		if CustomConfigModel != nil {
			properties := make(map[string]interface{})
			properties["add"] = CustomConfigModel.Add
			properties["allow_storage_resize"] = CustomConfigModel.AllowStorageResize
			properties["field_delimiter"] = CustomConfigModel.FieldDelimiter
			properties["name"] = CustomConfigModel.Name
			properties["override"] = CustomConfigModel.Override
			properties["template"] = CustomConfigModel.Template
			properties["variable_groups"] = SetCustomConfigVariableGroupSubResourceData(CustomConfigModel.VariableGroups)
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the CustomConfig resource defined in the Terraform configuration
func CustomConfig() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"add": {
			Description: `Add the Custom Config to App Instance (Optional. Default: False)`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"allow_storage_resize": {
			Description: `Allow Appinstance storage to be resized after app instance is created. (Optional. Default: False)`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"field_delimiter": {
			Description: `Field delimiter used in specifying variables in template. (Required)`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Description: `Name of CustomConfig (Required)`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"override": {
			Description: `Override existing custom config from App Bundle Manifest (Optional. Default: False)`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"template": {
			Description: `base64 encrypted template string. (Optional)`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"variable_groups": {
			Description: `List of Variable groups. (Required)`,
			Type:        schema.TypeList, //GoType: []*CustomConfigVariableGroup
			Elem: &schema.Resource{
				Schema: VariableGroup(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the CustomConfig resource
func GetCustomConfigPropertyFields() (t []string) {
	return []string{
		"add",
		"allow_storage_resize",
		"field_delimiter",
		"name",
		"override",
		"template",
		"variable_groups",
	}
}
