package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AppTypeModel(d *schema.ResourceData) *models.AppType {
	appType, _ := d.Get("app_type").(models.AppType)
	return &appType
}

func AppTypeModelFromMap(m map[string]interface{}) *models.AppType {
	appType := m["app_type"].(models.AppType)
	return &appType
}

// Update the underlying AppType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppTypeResourceData(d *schema.ResourceData, m *models.AppType) {
}

// Iterate through and update the AppType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppTypeSubResourceData(m []*models.AppType) (d []*map[string]interface{}) {
	for _, AppTypeModel := range m {
		if AppTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppType resource defined in the Terraform configuration
func AppTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the AppType resource
func GetAppTypePropertyFields() (t []string) {
	return []string{}
}
