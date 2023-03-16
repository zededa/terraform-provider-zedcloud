package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AppACEMatchModel(d *schema.ResourceData) *models.AppACEMatch {
	typeVar, _ := d.Get("type").(string)
	value, _ := d.Get("value").(string)
	return &models.AppACEMatch{
		Type:  &typeVar, // string true false false
		Value: &value,   // string true false false
	}
}

func AppACEMatchModelFromMap(m map[string]interface{}) *models.AppACEMatch {
	typeVar := m["type"].(string)
	value := m["value"].(string)
	return &models.AppACEMatch{
		Type:  &typeVar,
		Value: &value,
	}
}

// Update the underlying AppACEMatch resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppACEMatchResourceData(d *schema.ResourceData, m *models.AppACEMatch) {
	d.Set("type", m.Type)
	d.Set("value", m.Value)
}

// Iterate through and update the AppACEMatch resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppACEMatchSubResourceData(m []*models.AppACEMatch) (d []*map[string]interface{}) {
	for _, AppACEMatchModel := range m {
		if AppACEMatchModel != nil {
			properties := make(map[string]interface{})
			properties["type"] = AppACEMatchModel.Type
			properties["value"] = AppACEMatchModel.Value
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppACEMatch resource defined in the Terraform configuration
func AppACEMatchSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type": {
			Description: `Type`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"value": {
			Description: `Value`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

// Retrieve property field names for updating the AppACEMatch resource
func GetAppACEMatchPropertyFields() (t []string) {
	return []string{
		"type",
		"value",
	}
}
