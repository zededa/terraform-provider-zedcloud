package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate CustomConfig resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func CustomConfigModel(d *schema.ResourceData) *models.CustomConfig {
	add := d.Get("add").(bool)
	allowStorageResize := d.Get("allow_storage_resize").(bool)
	fieldDelimiter := d.Get("field_delimiter").(string)
	name := d.Get("name").(string)
	override := d.Get("override").(bool)
	template := d.Get("template").(string)
	variableGroups := d.Get("variable_groups").([]*models.CustomConfigVariableGroup) // []*CustomConfigVariableGroup
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

func CustomConfigModelFromMap(m map[string]interface{}) *models.CustomConfig {
	add := m["add"].(bool)
	allowStorageResize := m["allow_storage_resize"].(bool)
	fieldDelimiter := m["field_delimiter"].(string)
	name := m["name"].(string)
	override := m["override"].(bool)
	template := m["template"].(string)
	variableGroups := m["variable_groups"].([]*models.CustomConfigVariableGroup) // []*CustomConfigVariableGroup
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

// Update the underlying CustomConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetCustomConfigResourceData(d *schema.ResourceData, m *models.CustomConfig) {
	d.Set("add", m.Add)
	d.Set("allow_storage_resize", m.AllowStorageResize)
	d.Set("field_delimiter", m.FieldDelimiter)
	d.Set("name", m.Name)
	d.Set("override", m.Override)
	d.Set("template", m.Template)
	d.Set("variable_groups", SetCustomConfigVariableGroupSubResourceData(m.VariableGroups))
}

// Iterate throught and update the CustomConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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
func CustomConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"add": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"allow_storage_resize": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"field_delimiter": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"override": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"template": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"variable_groups": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*CustomConfigVariableGroup
			Elem: &schema.Resource{
				Schema: CustomConfigVariableGroupSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
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
