package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate EnrollmentMechanism resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EnrollmentMechanismModel(d *schema.ResourceData) *models.EnrollmentMechanism {
	enrollmentMechanism, _ := d.Get("enrollment_mechanism").(models.EnrollmentMechanism)
	return &enrollmentMechanism
}

func EnrollmentMechanismModelFromMap(m map[string]interface{}) *models.EnrollmentMechanism {
	enrollmentMechanism := m["enrollment_mechanism"].(models.EnrollmentMechanism)
	return &enrollmentMechanism
}

// Update the underlying EnrollmentMechanism resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEnrollmentMechanismResourceData(d *schema.ResourceData, m *models.EnrollmentMechanism) {
}

// Iterate through and update the EnrollmentMechanism resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEnrollmentMechanismSubResourceData(m []*models.EnrollmentMechanism) (d []*map[string]interface{}) {
	for _, EnrollmentMechanismModel := range m {
		if EnrollmentMechanismModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the EnrollmentMechanism resource defined in the Terraform configuration
func EnrollmentMechanismSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the EnrollmentMechanism resource
func GetEnrollmentMechanismPropertyFields() (t []string) {
	return []string{}
}
