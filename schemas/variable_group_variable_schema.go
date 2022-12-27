package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate VariableGroupVariable resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func VariableGroupVariableModel(d *schema.ResourceData) *models.VariableGroupVariable {
	defaultVar := d.Get("default").(string)
	encode := d.Get("encode").(*models.VariableFileEncoding)   // VariableFileEncoding
	format := d.Get("format").(*models.VariableVariableFormat) // VariableVariableFormat
	label := d.Get("label").(string)
	maxLength := d.Get("max_length").(string)
	name := d.Get("name").(string)
	options := d.Get("options").([]*models.VariableOptionVal) // []*VariableOptionVal
	processInput := d.Get("process_input").(string)
	required := d.Get("required").(bool)
	typeVar := d.Get("type").(string)
	value := d.Get("value").(string)
	return &models.VariableGroupVariable{
		Default:      defaultVar,
		Encode:       encode,
		Format:       format,
		Label:        &label, // string true false false
		MaxLength:    maxLength,
		Name:         &name, // string true false false
		Options:      options,
		ProcessInput: processInput,
		Required:     &required, // bool true false false
		Type:         typeVar,
		Value:        value,
	}
}

func VariableGroupVariableModelFromMap(m map[string]interface{}) *models.VariableGroupVariable {
	defaultVar := m["default"].(string)
	encode := m["encode"].(*models.VariableFileEncoding)   // VariableFileEncoding
	format := m["format"].(*models.VariableVariableFormat) // VariableVariableFormat
	label := m["label"].(string)
	maxLength := m["max_length"].(string)
	name := m["name"].(string)
	options := m["options"].([]*models.VariableOptionVal) // []*VariableOptionVal
	processInput := m["process_input"].(string)
	required := m["required"].(bool)
	typeVar := m["type"].(string)
	value := m["value"].(string)
	return &models.VariableGroupVariable{
		Default:      defaultVar,
		Encode:       encode,
		Format:       format,
		Label:        &label,
		MaxLength:    maxLength,
		Name:         &name,
		Options:      options,
		ProcessInput: processInput,
		Required:     &required,
		Type:         typeVar,
		Value:        value,
	}
}

// Update the underlying VariableGroupVariable resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVariableGroupVariableResourceData(d *schema.ResourceData, m *models.VariableGroupVariable) {
	d.Set("default", m.Default)
	d.Set("encode", m.Encode)
	d.Set("format", m.Format)
	d.Set("label", m.Label)
	d.Set("max_length", m.MaxLength)
	d.Set("name", m.Name)
	d.Set("options", SetVariableOptionValSubResourceData(m.Options))
	d.Set("process_input", m.ProcessInput)
	d.Set("required", m.Required)
	d.Set("type", m.Type)
	d.Set("value", m.Value)
}

// Iterate throught and update the VariableGroupVariable resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVariableGroupVariableSubResourceData(m []*models.VariableGroupVariable) (d []*map[string]interface{}) {
	for _, VariableGroupVariableModel := range m {
		if VariableGroupVariableModel != nil {
			properties := make(map[string]interface{})
			properties["default"] = VariableGroupVariableModel.Default
			properties["encode"] = VariableGroupVariableModel.Encode
			properties["format"] = VariableGroupVariableModel.Format
			properties["label"] = VariableGroupVariableModel.Label
			properties["max_length"] = VariableGroupVariableModel.MaxLength
			properties["name"] = VariableGroupVariableModel.Name
			properties["options"] = SetVariableOptionValSubResourceData(VariableGroupVariableModel.Options)
			properties["process_input"] = VariableGroupVariableModel.ProcessInput
			properties["required"] = VariableGroupVariableModel.Required
			properties["type"] = VariableGroupVariableModel.Type
			properties["value"] = VariableGroupVariableModel.Value
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VariableGroupVariable resource defined in the Terraform configuration
func VariableGroupVariableSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"default": {
			Description: `Default value of the variable. (Optional. Default: <Default value based on type>)`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"encode": {
			Description: `Encoding of file content. Applicable if format is VARIABLE_FORMAT_FILE`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"format": {
			Description: `Format of the user variable. (Required)`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Required: true,
		},

		"label": {
			Description: `Label for the variable (Required)`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"max_length": {
			Description: `Max length of the value of the variable(Optional. Default: 1024)`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Description: `Name of the Variable (Required)`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"options": {
			Description: `Key-Value pair of options. Applicable if format is VARIABLE_FORMAT_DROPDOWN`,
			Type:        schema.TypeList, //GoType: []*VariableOptionVal
			Elem: &schema.Resource{
				Schema: VariableOptionValSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"process_input": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"required": {
			Description: `This variable MUST be specified when creating an App Instance. (Optional. Default: False)`,
			Type:        schema.TypeBool,
			Required:    true,
		},

		"type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"value": {
			Description: `User-specified value of the variable.(Required if required is true. Optional otherwise)`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the VariableGroupVariable resource
func GetVariableGroupVariablePropertyFields() (t []string) {
	return []string{
		"default",
		"encode",
		"format",
		"label",
		"max_length",
		"name",
		"options",
		"process_input",
		"required",
		"type",
		"value",
	}
}
