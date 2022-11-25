package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate AppNamingScheme resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AppNamingSchemeModel(d *schema.ResourceData) *models.AppNamingScheme {
	appNamingScheme := d.Get("app_naming_scheme").(models.AppNamingScheme)
	return &appNamingScheme
}

func AppNamingSchemeModelFromMap(m map[string]interface{}) *models.AppNamingScheme {
	appNamingScheme := m["app_naming_scheme"].(models.AppNamingScheme)
	return &appNamingScheme
}

// Update the underlying AppNamingScheme resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppNamingSchemeResourceData(d *schema.ResourceData, m *models.AppNamingScheme) {
}

// Iterate throught and update the AppNamingScheme resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppNamingSchemeSubResourceData(m []*models.AppNamingScheme) (d []*map[string]interface{}) {
	for _, AppNamingSchemeModel := range m {
		if AppNamingSchemeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppNamingScheme resource defined in the Terraform configuration
func AppNamingSchemeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the AppNamingScheme resource
func GetAppNamingSchemePropertyFields() (t []string) {
	return []string{}
}
