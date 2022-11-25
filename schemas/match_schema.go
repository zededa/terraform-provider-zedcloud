package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate Match resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MatchModel(d *schema.ResourceData) *models.Match {
	typeVar := d.Get("type").(string)
	value := d.Get("value").(string)
	return &models.Match{
		Type:  typeVar,
		Value: value,
	}
}

func MatchModelFromMap(m map[string]interface{}) *models.Match {
	typeVar := m["type"].(string)
	value := m["value"].(string)
	return &models.Match{
		Type:  typeVar,
		Value: value,
	}
}

// Update the underlying Match resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMatchResourceData(d *schema.ResourceData, m *models.Match) {
	d.Set("type", m.Type)
	d.Set("value", m.Value)
}

// Iterate throught and update the Match resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMatchSubResourceData(m []*models.Match) (d []*map[string]interface{}) {
	for _, MatchModel := range m {
		if MatchModel != nil {
			properties := make(map[string]interface{})
			properties["type"] = MatchModel.Type
			properties["value"] = MatchModel.Value
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the Match resource defined in the Terraform configuration
func MatchSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"value": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the Match resource
func GetMatchPropertyFields() (t []string) {
	return []string{
		"type",
		"value",
	}
}
