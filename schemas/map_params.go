package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate MapParams resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MapParamsModel(d *schema.ResourceData) *models.MapParams {
	appPortInt, _ := d.Get("app_port").(int)
	appPort := int64(appPortInt)
	return &models.MapParams{
		AppPort: appPort,
	}
}

func MapParamsModelFromMap(m map[string]interface{}) *models.MapParams {
	appPort := int64(m["app_port"].(int)) // int64 false false false
	return &models.MapParams{
		AppPort: appPort,
	}
}

// Update the underlying MapParams resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMapParamsResourceData(d *schema.ResourceData, m *models.MapParams) {
	d.Set("app_port", m.AppPort)
}

// Iterate through and update the MapParams resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMapParamsSubResourceData(m []*models.MapParams) (d []*map[string]interface{}) {
	for _, MapParamsModel := range m {
		if MapParamsModel != nil {
			properties := make(map[string]interface{})
			properties["app_port"] = MapParamsModel.AppPort
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the MapParams resource defined in the Terraform configuration
func MapParamsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_port": {
			Description: `Application Port value`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the MapParams resource
func GetMapParamsPropertyFields() (t []string) {
	return []string{
		"app_port",
	}
}
