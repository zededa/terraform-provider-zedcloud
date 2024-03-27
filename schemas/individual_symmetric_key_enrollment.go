package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate IndividualSymmetricKeyEnrollment resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func IndividualSymmetricKeyEnrollmentModel(d *schema.ResourceData) *models.IndividualSymmetricKeyEnrollment {
	registrationID, _ := d.Get("registration_id").(string)
	return &models.IndividualSymmetricKeyEnrollment{
		RegistrationID: registrationID,
	}
}

func IndividualSymmetricKeyEnrollmentModelFromMap(m map[string]interface{}) *models.IndividualSymmetricKeyEnrollment {
	registrationID := m["registration_id"].(string)
	return &models.IndividualSymmetricKeyEnrollment{
		RegistrationID: registrationID,
	}
}

// Update the underlying IndividualSymmetricKeyEnrollment resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetIndividualSymmetricKeyEnrollmentResourceData(d *schema.ResourceData, m *models.IndividualSymmetricKeyEnrollment) {
	d.Set("registration_id", m.RegistrationID)
}

// Iterate through and update the IndividualSymmetricKeyEnrollment resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetIndividualSymmetricKeyEnrollmentSubResourceData(m []*models.IndividualSymmetricKeyEnrollment) (d []*map[string]interface{}) {
	for _, IndividualSymmetricKeyEnrollmentModel := range m {
		if IndividualSymmetricKeyEnrollmentModel != nil {
			properties := make(map[string]interface{})
			properties["registration_id"] = IndividualSymmetricKeyEnrollmentModel.RegistrationID
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the IndividualSymmetricKeyEnrollment resource defined in the Terraform configuration
func IndividualSymmetricKeyEnrollmentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"registration_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the IndividualSymmetricKeyEnrollment resource
func GetIndividualSymmetricKeyEnrollmentPropertyFields() (t []string) {
	return []string{
		"registration_id",
	}
}
