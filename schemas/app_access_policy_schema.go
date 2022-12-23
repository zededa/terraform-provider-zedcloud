package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate AppAccessPolicy resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AppAccessPolicyModel(d *schema.ResourceData) *models.AppAccessPolicy {
	allowApp := d.Get("allow_app").(bool)
	return &models.AppAccessPolicy{
		AllowApp: allowApp,
	}
}

func AppAccessPolicyModelFromMap(m map[string]interface{}) *models.AppAccessPolicy {
	allowApp := m["allow_app"].(bool)
	return &models.AppAccessPolicy{
		AllowApp: allowApp,
	}
}

// Update the underlying AppAccessPolicy resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppAccessPolicyResourceData(d *schema.ResourceData, m *models.AppAccessPolicy) {
	d.Set("allow_app", m.AllowApp)
}

// Iterate throught and update the AppAccessPolicy resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppAccessPolicySubResourceData(m []*models.AppAccessPolicy) (d []*map[string]interface{}) {
	for _, AppAccessPolicyModel := range m {
		if AppAccessPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["allow_app"] = AppAccessPolicyModel.AllowApp
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppAccessPolicy resource defined in the Terraform configuration
func AppAccessPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_app": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the AppAccessPolicy resource
func GetAppAccessPolicyPropertyFields() (t []string) {
	return []string{
		"allow_app",
	}
}
