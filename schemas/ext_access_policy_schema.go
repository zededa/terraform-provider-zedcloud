package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ExtAccessPolicy resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ExtAccessPolicyModel(d *schema.ResourceData) *models.ExtAccessPolicy {
	allowExt := d.Get("allow_ext").(bool)
	return &models.ExtAccessPolicy{
		AllowExt: allowExt,
	}
}

func ExtAccessPolicyModelFromMap(m map[string]interface{}) *models.ExtAccessPolicy {
	allowExt := m["allow_ext"].(bool)
	return &models.ExtAccessPolicy{
		AllowExt: allowExt,
	}
}

// Update the underlying ExtAccessPolicy resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetExtAccessPolicyResourceData(d *schema.ResourceData, m *models.ExtAccessPolicy) {
	d.Set("allow_ext", m.AllowExt)
}

// Iterate throught and update the ExtAccessPolicy resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetExtAccessPolicySubResourceData(m []*models.ExtAccessPolicy) (d []*map[string]interface{}) {
	for _, ExtAccessPolicyModel := range m {
		if ExtAccessPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["allow_ext"] = ExtAccessPolicyModel.AllowExt
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ExtAccessPolicy resource defined in the Terraform configuration
func ExtAccessPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_ext": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the ExtAccessPolicy resource
func GetExtAccessPolicyPropertyFields() (t []string) {
	return []string{
		"allow_ext",
	}
}
