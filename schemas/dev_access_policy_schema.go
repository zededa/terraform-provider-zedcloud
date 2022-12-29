package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DevAccessPolicy resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DevAccessPolicyModel(d *schema.ResourceData) *models.DevAccessPolicy {
	allowDev, _ := d.Get("allow_dev").(bool)
	return &models.DevAccessPolicy{
		AllowDev: allowDev,
	}
}

func DevAccessPolicyModelFromMap(m map[string]interface{}) *models.DevAccessPolicy {
	allowDev := m["allow_dev"].(bool)
	return &models.DevAccessPolicy{
		AllowDev: allowDev,
	}
}

// Update the underlying DevAccessPolicy resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDevAccessPolicyResourceData(d *schema.ResourceData, m *models.DevAccessPolicy) {
	d.Set("allow_dev", m.AllowDev)
}

// Iterate through and update the DevAccessPolicy resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDevAccessPolicySubResourceData(m []*models.DevAccessPolicy) (d []*map[string]interface{}) {
	for _, DevAccessPolicyModel := range m {
		if DevAccessPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["allow_dev"] = DevAccessPolicyModel.AllowDev
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DevAccessPolicy resource defined in the Terraform configuration
func DevAccessPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_dev": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the DevAccessPolicy resource
func GetDevAccessPolicyPropertyFields() (t []string) {
	return []string{
		"allow_dev",
	}
}
