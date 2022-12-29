package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate PCRValue resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PCRValueModel(d *schema.ResourceData) *models.PCRValue {
	indexInt, _ := d.Get("index").(int)
	index := int64(indexInt)
	typeVarModel, ok := d.Get("type").(models.PCRType) // PCRType
	typeVar := &typeVarModel
	if !ok {
		typeVar = nil
	}
	value, _ := d.Get("value").(string)
	return &models.PCRValue{
		Index: &index, // int64 true false false
		Type:  typeVar,
		Value: &value, // string true false false
	}
}

func PCRValueModelFromMap(m map[string]interface{}) *models.PCRValue {
	index := int64(m["index"].(int))       // int64 true false false
	typeVar := m["type"].(*models.PCRType) // PCRType
	value := m["value"].(string)
	return &models.PCRValue{
		Index: &index,
		Type:  typeVar,
		Value: &value,
	}
}

// Update the underlying PCRValue resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPCRValueResourceData(d *schema.ResourceData, m *models.PCRValue) {
	d.Set("index", m.Index)
	d.Set("type", m.Type)
	d.Set("value", m.Value)
}

// Iterate through and update the PCRValue resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPCRValueSubResourceData(m []*models.PCRValue) (d []*map[string]interface{}) {
	for _, PCRValueModel := range m {
		if PCRValueModel != nil {
			properties := make(map[string]interface{})
			properties["index"] = PCRValueModel.Index
			properties["type"] = PCRValueModel.Type
			properties["value"] = PCRValueModel.Value
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the PCRValue resource defined in the Terraform configuration
func PCRValueSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Description: `Current index for the PCR item in the list. First element has the index 0`,
			Type:        schema.TypeInt,
			Required:    true,
		},

		"type": {
			Description: `Type of value for the PCR item. Could be one of the following values: 0 (PCR_TYPE_UNSPECIFIED), 1 (PCR_TYPE_HASH) or 2 (PCR_TYPE_EVENT_LOG)`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"value": {
			Description: `Actual value for the PCR item.`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

// Retrieve property field names for updating the PCRValue resource
func GetPCRValuePropertyFields() (t []string) {
	return []string{
		"index",
		"type",
		"value",
	}
}
