package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ObjectTag resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ObjectTagModel(d *schema.ResourceData) *models.ObjectTag {
	description := d.Get("description").(string)
	key := d.Get("key").(string)
	value := d.Get("value").(string)
	return &models.ObjectTag{
		Description: description,
		Key:         key,
		Value:       value,
	}
}

func ObjectTagModelFromMap(m map[string]interface{}) *models.ObjectTag {
	description := m["description"].(string)
	key := m["key"].(string)
	value := m["value"].(string)
	return &models.ObjectTag{
		Description: description,
		Key:         key,
		Value:       value,
	}
}

// Update the underlying ObjectTag resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetObjectTagResourceData(d *schema.ResourceData, m *models.ObjectTag) {
	d.Set("description", m.Description)
	d.Set("key", m.Key)
	d.Set("value", m.Value)
}

// Iterate throught and update the ObjectTag resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetObjectTagSubResourceData(m []*models.ObjectTag) (d []*map[string]interface{}) {
	for _, ObjectTagModel := range m {
		if ObjectTagModel != nil {
			properties := make(map[string]interface{})
			properties["description"] = ObjectTagModel.Description
			properties["key"] = ObjectTagModel.Key
			properties["value"] = ObjectTagModel.Value
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ObjectTag resource defined in the Terraform configuration
func ObjectTagSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Description: `Description of the tag key`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"key": {
			Description: `object tag key identifier.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"value": {
			Description: `value associated for the object tag key`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the ObjectTag resource
func GetObjectTagPropertyFields() (t []string) {
	return []string{
		"description",
		"key",
		"value",
	}
}
