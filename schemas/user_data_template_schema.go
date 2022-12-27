package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate UserDataTemplate resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func UserDataTemplateModel(d *schema.ResourceData) *models.UserDataTemplate {
	var customConfig *models.CustomConfig // CustomConfig
	customConfigInterface, customConfigIsSet := d.GetOk("custom_config")
	if customConfigIsSet {
		customConfigMap := customConfigInterface.([]interface{})[0].(map[string]interface{})
		customConfig = CustomConfigModelFromMap(customConfigMap)
	}
	return &models.UserDataTemplate{
		CustomConfig: customConfig,
	}
}

func UserDataTemplateModelFromMap(m map[string]interface{}) *models.UserDataTemplate {
	var customConfig *models.CustomConfig // CustomConfig
	customConfigInterface, customConfigIsSet := m["custom_config"]
	if customConfigIsSet {
		customConfigMap := customConfigInterface.([]interface{})[0].(map[string]interface{})
		customConfig = CustomConfigModelFromMap(customConfigMap)
	}
	//
	return &models.UserDataTemplate{
		CustomConfig: customConfig,
	}
}

// Update the underlying UserDataTemplate resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetUserDataTemplateResourceData(d *schema.ResourceData, m *models.UserDataTemplate) {
	d.Set("custom_config", SetCustomConfigSubResourceData([]*models.CustomConfig{m.CustomConfig}))
}

// Iterate throught and update the UserDataTemplate resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetUserDataTemplateSubResourceData(m []*models.UserDataTemplate) (d []*map[string]interface{}) {
	for _, UserDataTemplateModel := range m {
		if UserDataTemplateModel != nil {
			properties := make(map[string]interface{})
			properties["custom_config"] = SetCustomConfigSubResourceData([]*models.CustomConfig{UserDataTemplateModel.CustomConfig})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the UserDataTemplate resource defined in the Terraform configuration
func UserDataTemplateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"custom_config": {
			Description: ``,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the UserDataTemplate resource
func GetUserDataTemplatePropertyFields() (t []string) {
	return []string{
		"custom_config",
	}
}
