package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func VariableGroupVariableModel(d *schema.ResourceData) *models.VariableGroupVariable {
	defaultVar, _ := d.Get("default").(string)
	var encode *models.VariableFileEncoding // VariableFileEncoding
	encodeInterface, encodeIsSet := d.GetOk("encode")
	if encodeIsSet {
		encodeModel := encodeInterface.(string)
		encode = models.NewVariableFileEncoding(models.VariableFileEncoding(encodeModel))
	}
	var format *models.VariableVariableFormat // VariableVariableFormat
	formatInterface, formatIsSet := d.GetOk("format")
	if formatIsSet {
		formatModel := formatInterface.(string)
		format = models.NewVariableVariableFormat(models.VariableVariableFormat(formatModel))
	}
	label, _ := d.Get("label").(string)
	maxLength, _ := d.Get("max_length").(string)
	name, _ := d.Get("name").(string)
	var options []*models.VariableOptionVal // []*VariableOptionVal
	optionsInterface, optionsIsSet := d.GetOk("options")
	if optionsIsSet {
		var items []interface{}
		if listItems, isList := optionsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = optionsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := VariableOptionValModelFromMap(v.(map[string]interface{}))
			options = append(options, m)
		}
	}
	processInput, _ := d.Get("process_input").(string)
	required, _ := d.Get("required").(bool)
	typeVar, _ := d.Get("type").(string)
	value, _ := d.Get("value").(string)
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
	var encode *models.VariableFileEncoding // VariableFileEncoding
	encodeInterface, encodeIsSet := m["encode"]
	if encodeIsSet {
		encodeModel := encodeInterface.(string)
		encode = models.NewVariableFileEncoding(models.VariableFileEncoding(encodeModel))
	}
	var format *models.VariableVariableFormat // VariableVariableFormat
	formatInterface, formatIsSet := m["format"]
	if formatIsSet {
		formatModel := formatInterface.(string)
		format = models.NewVariableVariableFormat(models.VariableVariableFormat(formatModel))
	}
	label := m["label"].(string)
	maxLength := m["max_length"].(string)
	name := m["name"].(string)
	var options []*models.VariableOptionVal // []*VariableOptionVal
	optionsInterface, optionsIsSet := m["options"]
	if optionsIsSet {
		var items []interface{}
		if listItems, isList := optionsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = optionsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := VariableOptionValModelFromMap(v.(map[string]interface{}))
			options = append(options, m)
		}
	}
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
func VariableGroupVariable() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"default": {
			Description: `Default value of the variable. (Optional. Default: <Default value based on type>)`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"encode": {
			Description: `Encoding of file content. Applicable if format is VARIABLE_FORMAT_FILE`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"format": {
			Description: `Format of the user variable. (Required)`,
			Type:        schema.TypeString,
			Required:    true,
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
