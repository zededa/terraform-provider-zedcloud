package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// Function to perform the following actions:
// (1) Translate EnrollmentType resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EnrollmentTypeModel(d *schema.ResourceData) *models.EnrollmentType {
	enrollmentType, _ := d.Get("enrollment_type").(models.EnrollmentType)
	return &enrollmentType
}

func EnrollmentTypeModelFromMap(m map[string]interface{}) *models.EnrollmentType {
	enrollmentType := m["enrollment_type"].(models.EnrollmentType)
	return &enrollmentType
}

// Update the underlying EnrollmentType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEnrollmentTypeResourceData(d *schema.ResourceData, m *models.EnrollmentType) {
}

// Iterate through and update the EnrollmentType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEnrollmentTypeSubResourceData(m []*models.EnrollmentType) (d []*map[string]interface{}) {
	for _, EnrollmentTypeModel := range m {
		if EnrollmentTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the EnrollmentType resource defined in the Terraform configuration
func EnrollmentTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the EnrollmentType resource
func GetEnrollmentTypePropertyFields() (t []string) {
	return []string{}
}
