package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate EDConfigItem resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EDConfigItemModel(d *schema.ResourceData) *models.EDConfigItem {
	boolValue := d.Get("bool_value").(bool)
	floatValue := d.Get("float_value").(float32)
	key := d.Get("key").(string)
	stringValue := d.Get("string_value").(string)
	uint32Value := int64(d.Get("uint32_value").(int))
	uint64Value := d.Get("uint64_value").(string)
	valueType := d.Get("value_type").(string)
	return &models.EDConfigItem{
		BoolValue:   boolValue,
		FloatValue:  floatValue,
		Key:         key,
		StringValue: stringValue,
		Uint32Value: uint32Value,
		Uint64Value: uint64Value,
		ValueType:   valueType,
	}
}

func EDConfigItemModelFromMap(m map[string]interface{}) *models.EDConfigItem {
	boolValue := m["bool_value"].(bool)
	floatValue := m["float_value"].(float32)
	key := m["key"].(string)
	stringValue := m["string_value"].(string)
	uint32Value := int64(m["uint32_value"].(int)) // int64 false false false
	uint64Value := m["uint64_value"].(string)
	valueType := m["value_type"].(string)
	return &models.EDConfigItem{
		BoolValue:   boolValue,
		FloatValue:  floatValue,
		Key:         key,
		StringValue: stringValue,
		Uint32Value: uint32Value,
		Uint64Value: uint64Value,
		ValueType:   valueType,
	}
}

// Update the underlying EDConfigItem resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEDConfigItemResourceData(d *schema.ResourceData, m *models.EDConfigItem) {
	d.Set("bool_value", m.BoolValue)
	d.Set("float_value", m.FloatValue)
	d.Set("key", m.Key)
	d.Set("string_value", m.StringValue)
	d.Set("uint32_value", m.Uint32Value)
	d.Set("uint64_value", m.Uint64Value)
	d.Set("value_type", m.ValueType)
}

// Iterate throught and update the EDConfigItem resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEDConfigItemSubResourceData(m []*models.EDConfigItem) (d []*map[string]interface{}) {
	for _, EDConfigItemModel := range m {
		if EDConfigItemModel != nil {
			properties := make(map[string]interface{})
			properties["bool_value"] = EDConfigItemModel.BoolValue
			properties["float_value"] = EDConfigItemModel.FloatValue
			properties["key"] = EDConfigItemModel.Key
			properties["string_value"] = EDConfigItemModel.StringValue
			properties["uint32_value"] = EDConfigItemModel.Uint32Value
			properties["uint64_value"] = EDConfigItemModel.Uint64Value
			properties["value_type"] = EDConfigItemModel.ValueType
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the EDConfigItem resource defined in the Terraform configuration
func EDConfigItemSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"bool_value": {
			Description: `boolean value`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"float_value": {
			Description: `float value`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"key": {
			Description: `key`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"string_value": {
			Description: `string value`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"uint32_value": {
			Description: `uint32 value`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"uint64_value": {
			Description: `uint64 value in string format`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"value_type": {
			Description: `value type`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the EDConfigItem resource
func GetEDConfigItemPropertyFields() (t []string) {
	return []string{
		"bool_value",
		"float_value",
		"key",
		"string_value",
		"uint32_value",
		"uint64_value",
		"value_type",
	}
}
