package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func AppMapParamsModel(d *schema.ResourceData) *models.AppMapParams {
	portInt, _ := d.Get("port").(int)
	port := int64(portInt)
	return &models.AppMapParams{
		Port: &port, // int64 true false false
	}
}

func AppMapParamsModelFromMap(m map[string]interface{}) *models.AppMapParams {
	port := int64(m["port"].(int)) // int64
	return &models.AppMapParams{
		Port: &port,
	}
}

// Update the underlying AppMapParams resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppMapParamsResourceData(d *schema.ResourceData, m *models.AppMapParams) {
	d.Set("port", m.Port)
}

// Iterate through and update the AppMapParams resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppMapParamsSubResourceData(m []*models.AppMapParams) (d []*map[string]interface{}) {
	for _, AppMapParamsModel := range m {
		if AppMapParamsModel != nil {
			properties := make(map[string]interface{})
			properties["port"] = AppMapParamsModel.Port
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppMapParams resource defined in the Terraform configuration
func AppMapParamsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"port": {
			Description: `Application port`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the AppMapParams resource
func GetAppMapParamsPropertyFields() (t []string) {
	return []string{
		"port",
	}
}
