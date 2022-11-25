package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate GroupSymmetricKeyEnrollment resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func GroupSymmetricKeyEnrollmentModel(d *schema.ResourceData) *models.GroupSymmetricKeyEnrollment {
	groupName := d.Get("group_name").(string)
	return &models.GroupSymmetricKeyEnrollment{
		GroupName: groupName,
	}
}

func GroupSymmetricKeyEnrollmentModelFromMap(m map[string]interface{}) *models.GroupSymmetricKeyEnrollment {
	groupName := m["group_name"].(string)
	return &models.GroupSymmetricKeyEnrollment{
		GroupName: groupName,
	}
}

// Update the underlying GroupSymmetricKeyEnrollment resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetGroupSymmetricKeyEnrollmentResourceData(d *schema.ResourceData, m *models.GroupSymmetricKeyEnrollment) {
	d.Set("group_name", m.GroupName)
}

// Iterate throught and update the GroupSymmetricKeyEnrollment resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetGroupSymmetricKeyEnrollmentSubResourceData(m []*models.GroupSymmetricKeyEnrollment) (d []*map[string]interface{}) {
	for _, GroupSymmetricKeyEnrollmentModel := range m {
		if GroupSymmetricKeyEnrollmentModel != nil {
			properties := make(map[string]interface{})
			properties["group_name"] = GroupSymmetricKeyEnrollmentModel.GroupName
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the GroupSymmetricKeyEnrollment resource defined in the Terraform configuration
func GroupSymmetricKeyEnrollmentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"group_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the GroupSymmetricKeyEnrollment resource
func GetGroupSymmetricKeyEnrollmentPropertyFields() (t []string) {
	return []string{
		"group_name",
	}
}
