package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// Function to perform the following actions:
// (1) Translate AppNamingSchemeV2 resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AppNamingSchemeV2Model(d *schema.ResourceData) *models.AppNamingSchemeV2 {
	appNamingSchemeV2, _ := d.Get("app_naming_scheme_v2").(models.AppNamingSchemeV2)
	return &appNamingSchemeV2
}

func AppNamingSchemeV2ModelFromMap(m map[string]interface{}) *models.AppNamingSchemeV2 {
	appNamingSchemeV2 := m["app_naming_scheme_v2"].(models.AppNamingSchemeV2)
	return &appNamingSchemeV2
}

// Update the underlying AppNamingSchemeV2 resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppNamingSchemeV2ResourceData(d *schema.ResourceData, m *models.AppNamingSchemeV2) {
}

// Iterate through and update the AppNamingSchemeV2 resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppNamingSchemeV2SubResourceData(m []*models.AppNamingSchemeV2) (d []*map[string]interface{}) {
	for _, AppNamingSchemeV2Model := range m {
		if AppNamingSchemeV2Model != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppNamingSchemeV2 resource defined in the Terraform configuration
func AppNamingSchemeV2Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the AppNamingSchemeV2 resource
func GetAppNamingSchemeV2PropertyFields() (t []string) {
	return []string{}
}
