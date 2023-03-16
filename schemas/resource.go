package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate Resource resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ResourceModel(d *schema.ResourceData) *models.Resource {
	name, _ := d.Get("name").(string)
	value, _ := d.Get("value").(string)
	return &models.Resource{
		Name:  name,
		Value: value,
	}
}

func ResourceModelFromMap(m map[string]interface{}) *models.Resource {
	name := m["name"].(string)
	value := m["value"].(string)
	return &models.Resource{
		Name:  name,
		Value: value,
	}
}

// Update the underlying Resource resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetResourceResourceData(d *schema.ResourceData, m *models.Resource) {
	d.Set("name", m.Name)
	d.Set("value", m.Value)
}

// Iterate through and update the Resource resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetResourceSubResourceData(m []*models.Resource) (d []*map[string]interface{}) {
	for _, ResourceModel := range m {
		if ResourceModel != nil {
			properties := make(map[string]interface{})
			properties["name"] = ResourceModel.Name
			properties["value"] = ResourceModel.Value
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the Resource resource defined in the Terraform configuration
func ResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Description: `Name of the Resource (Required)`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"value": {
			Description: `Value of Resource (Required)`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the Resource resource
func GetResourcePropertyFields() (t []string) {
	return []string{
		"name",
		"value",
	}
}
