package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate Permission resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PermissionModel(d *schema.ResourceData) *models.Permission {
	permission, _ := d.Get("permission").(models.Permission)
	return &permission
}

func PermissionModelFromMap(m map[string]interface{}) *models.Permission {
	permission := m["permission"].(models.Permission)
	return &permission
}

// Update the underlying Permission resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPermissionResourceData(d *schema.ResourceData, m *models.Permission) {
}

// Iterate through and update the Permission resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPermissionSubResourceData(m []*models.Permission) (d []*map[string]interface{}) {
	for _, PermissionModel := range m {
		if PermissionModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the Permission resource defined in the Terraform configuration
func PermissionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the Permission resource
func GetPermissionPropertyFields() (t []string) {
	return []string{}
}
