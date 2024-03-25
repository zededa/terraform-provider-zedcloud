package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate CmdArg resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func CmdArgModel(d *schema.ResourceData) *models.CmdArg {
	key, _ := d.Get("key").(string)
	value, _ := d.Get("value").(string)
	return &models.CmdArg{
		Key:   key,
		Value: value,
	}
}

func CmdArgModelFromMap(m map[string]interface{}) *models.CmdArg {
	key := m["key"].(string)
	value := m["value"].(string)
	return &models.CmdArg{
		Key:   key,
		Value: value,
	}
}

// Update the underlying CmdArg resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetCmdArgResourceData(d *schema.ResourceData, m *models.CmdArg) {
	d.Set("key", m.Key)
	d.Set("value", m.Value)
}

// Iterate through and update the CmdArg resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetCmdArgSubResourceData(m []*models.CmdArg) (d []*map[string]interface{}) {
	for _, CmdArgModel := range m {
		if CmdArgModel != nil {
			properties := make(map[string]interface{})
			properties["key"] = CmdArgModel.Key
			properties["value"] = CmdArgModel.Value
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the CmdArg resource defined in the Terraform configuration
func CmdArgSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"key": {
			Description: `Command line argument: key`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"value": {
			Description: `Command line argument: value`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the CmdArg resource
func GetCmdArgPropertyFields() (t []string) {
	return []string{
		"key",
		"value",
	}
}
