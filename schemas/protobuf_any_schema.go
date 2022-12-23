package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ProtobufAny resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ProtobufAnyModel(d *schema.ResourceData) *models.ProtobufAny {
	typeURL := d.Get("type_url").(string)
	value := d.Get("value").(strfmt.Base64)
	return &models.ProtobufAny{
		TypeURL: typeURL,
		Value:   value,
	}
}

func ProtobufAnyModelFromMap(m map[string]interface{}) *models.ProtobufAny {
	typeURL := m["type_url"].(string)
	value := m["value"].(strfmt.Base64)
	return &models.ProtobufAny{
		TypeURL: typeURL,
		Value:   value,
	}
}

// Update the underlying ProtobufAny resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetProtobufAnyResourceData(d *schema.ResourceData, m *models.ProtobufAny) {
	d.Set("type_url", m.TypeURL)
	d.Set("value", m.Value)
}

// Iterate throught and update the ProtobufAny resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetProtobufAnySubResourceData(m []*models.ProtobufAny) (d []*map[string]interface{}) {
	for _, ProtobufAnyModel := range m {
		if ProtobufAnyModel != nil {
			properties := make(map[string]interface{})
			properties["type_url"] = ProtobufAnyModel.TypeURL
			properties["value"] = ProtobufAnyModel.Value
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ProtobufAny resource defined in the Terraform configuration
func ProtobufAnySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type_url": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"value": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},
	}
}

// Retrieve property field names for updating the ProtobufAny resource
func GetProtobufAnyPropertyFields() (t []string) {
	return []string{
		"type_url",
		"value",
	}
}
