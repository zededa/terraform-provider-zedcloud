package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate AppPolicy resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AppPolicyModel(d *schema.ResourceData) *models.AppPolicy {
	apps := d.Get("apps").([]*models.AppConfig) // []*AppConfig
	return &models.AppPolicy{
		Apps: apps,
	}
}

func AppPolicyModelFromMap(m map[string]interface{}) *models.AppPolicy {
	apps := m["apps"].([]*models.AppConfig) // []*AppConfig
	return &models.AppPolicy{
		Apps: apps,
	}
}

// Update the underlying AppPolicy resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppPolicyResourceData(d *schema.ResourceData, m *models.AppPolicy) {
	d.Set("apps", SetAppConfigSubResourceData(m.Apps))
}

// Iterate throught and update the AppPolicy resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppPolicySubResourceData(m []*models.AppPolicy) (d []*map[string]interface{}) {
	for _, AppPolicyModel := range m {
		if AppPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["apps"] = SetAppConfigSubResourceData(AppPolicyModel.Apps)
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppPolicy resource defined in the Terraform configuration
func AppPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"apps": {
			Description: `list of app details that will be provisioned on all the devices of the project to which this policy is attached`,
			Type:        schema.TypeList, //GoType: []*AppConfig
			Elem: &schema.Resource{
				Schema: AppConfigSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},
	}
}

// Retrieve property field names for updating the AppPolicy resource
func GetAppPolicyPropertyFields() (t []string) {
	return []string{
		"apps",
	}
}
