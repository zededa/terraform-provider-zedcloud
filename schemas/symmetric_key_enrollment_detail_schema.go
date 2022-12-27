package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate SymmetricKeyEnrollmentDetail resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func SymmetricKeyEnrollmentDetailModel(d *schema.ResourceData) *models.SymmetricKeyEnrollmentDetail {
	var groupSymmetricKeyEnrollment *models.GroupSymmetricKeyEnrollment // GroupSymmetricKeyEnrollment
	groupSymmetricKeyEnrollmentInterface, groupSymmetricKeyEnrollmentIsSet := d.GetOk("group_symmetric_key_enrollment")
	if groupSymmetricKeyEnrollmentIsSet {
		groupSymmetricKeyEnrollmentMap := groupSymmetricKeyEnrollmentInterface.([]interface{})[0].(map[string]interface{})
		groupSymmetricKeyEnrollment = GroupSymmetricKeyEnrollmentModelFromMap(groupSymmetricKeyEnrollmentMap)
	}
	var individualSymmetricKeyEnrollment *models.IndividualSymmetricKeyEnrollment // IndividualSymmetricKeyEnrollment
	individualSymmetricKeyEnrollmentInterface, individualSymmetricKeyEnrollmentIsSet := d.GetOk("individual_symmetric_key_enrollment")
	if individualSymmetricKeyEnrollmentIsSet {
		individualSymmetricKeyEnrollmentMap := individualSymmetricKeyEnrollmentInterface.([]interface{})[0].(map[string]interface{})
		individualSymmetricKeyEnrollment = IndividualSymmetricKeyEnrollmentModelFromMap(individualSymmetricKeyEnrollmentMap)
	}
	typeVar := d.Get("type").(*models.EnrollmentType) // EnrollmentType
	return &models.SymmetricKeyEnrollmentDetail{
		GroupSymmetricKeyEnrollment:      groupSymmetricKeyEnrollment,
		IndividualSymmetricKeyEnrollment: individualSymmetricKeyEnrollment,
		Type:                             typeVar,
	}
}

func SymmetricKeyEnrollmentDetailModelFromMap(m map[string]interface{}) *models.SymmetricKeyEnrollmentDetail {
	var groupSymmetricKeyEnrollment *models.GroupSymmetricKeyEnrollment // GroupSymmetricKeyEnrollment
	groupSymmetricKeyEnrollmentInterface, groupSymmetricKeyEnrollmentIsSet := m["group_symmetric_key_enrollment"]
	if groupSymmetricKeyEnrollmentIsSet {
		groupSymmetricKeyEnrollmentMap := groupSymmetricKeyEnrollmentInterface.([]interface{})[0].(map[string]interface{})
		groupSymmetricKeyEnrollment = GroupSymmetricKeyEnrollmentModelFromMap(groupSymmetricKeyEnrollmentMap)
	}
	//
	var individualSymmetricKeyEnrollment *models.IndividualSymmetricKeyEnrollment // IndividualSymmetricKeyEnrollment
	individualSymmetricKeyEnrollmentInterface, individualSymmetricKeyEnrollmentIsSet := m["individual_symmetric_key_enrollment"]
	if individualSymmetricKeyEnrollmentIsSet {
		individualSymmetricKeyEnrollmentMap := individualSymmetricKeyEnrollmentInterface.([]interface{})[0].(map[string]interface{})
		individualSymmetricKeyEnrollment = IndividualSymmetricKeyEnrollmentModelFromMap(individualSymmetricKeyEnrollmentMap)
	}
	//
	typeVar := m["type"].(*models.EnrollmentType) // EnrollmentType
	return &models.SymmetricKeyEnrollmentDetail{
		GroupSymmetricKeyEnrollment:      groupSymmetricKeyEnrollment,
		IndividualSymmetricKeyEnrollment: individualSymmetricKeyEnrollment,
		Type:                             typeVar,
	}
}

// Update the underlying SymmetricKeyEnrollmentDetail resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetSymmetricKeyEnrollmentDetailResourceData(d *schema.ResourceData, m *models.SymmetricKeyEnrollmentDetail) {
	d.Set("group_symmetric_key_enrollment", SetGroupSymmetricKeyEnrollmentSubResourceData([]*models.GroupSymmetricKeyEnrollment{m.GroupSymmetricKeyEnrollment}))
	d.Set("individual_symmetric_key_enrollment", SetIndividualSymmetricKeyEnrollmentSubResourceData([]*models.IndividualSymmetricKeyEnrollment{m.IndividualSymmetricKeyEnrollment}))
	d.Set("type", m.Type)
}

// Iterate throught and update the SymmetricKeyEnrollmentDetail resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetSymmetricKeyEnrollmentDetailSubResourceData(m []*models.SymmetricKeyEnrollmentDetail) (d []*map[string]interface{}) {
	for _, SymmetricKeyEnrollmentDetailModel := range m {
		if SymmetricKeyEnrollmentDetailModel != nil {
			properties := make(map[string]interface{})
			properties["group_symmetric_key_enrollment"] = SetGroupSymmetricKeyEnrollmentSubResourceData([]*models.GroupSymmetricKeyEnrollment{SymmetricKeyEnrollmentDetailModel.GroupSymmetricKeyEnrollment})
			properties["individual_symmetric_key_enrollment"] = SetIndividualSymmetricKeyEnrollmentSubResourceData([]*models.IndividualSymmetricKeyEnrollment{SymmetricKeyEnrollmentDetailModel.IndividualSymmetricKeyEnrollment})
			properties["type"] = SymmetricKeyEnrollmentDetailModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the SymmetricKeyEnrollmentDetail resource defined in the Terraform configuration
func SymmetricKeyEnrollmentDetailSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"group_symmetric_key_enrollment": {
			Description: ``,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"individual_symmetric_key_enrollment": {
			Description: ``,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"type": {
			Description: ``,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the SymmetricKeyEnrollmentDetail resource
func GetSymmetricKeyEnrollmentDetailPropertyFields() (t []string) {
	return []string{
		"group_symmetric_key_enrollment",
		"individual_symmetric_key_enrollment",
		"type",
	}
}
