package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate TPMEnrollmentDetail resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func TPMEnrollmentDetailModel(d *schema.ResourceData) *models.TPMEnrollmentDetail {
	typeVarModel, _ := d.Get("type").(models.EnrollmentType) // EnrollmentType
	typeVar := &typeVarModel
	if !ok {
		typeVar = nil
	}
	return &models.TPMEnrollmentDetail{
		Type: typeVar,
	}
}

func TPMEnrollmentDetailModelFromMap(m map[string]interface{}) *models.TPMEnrollmentDetail {
	typeVar := m["type"].(*models.EnrollmentType) // EnrollmentType
	return &models.TPMEnrollmentDetail{
		Type: typeVar,
	}
}

// Update the underlying TPMEnrollmentDetail resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetTPMEnrollmentDetailResourceData(d *schema.ResourceData, m *models.TPMEnrollmentDetail) {
	d.Set("type", m.Type)
}

// Iterate through and update the TPMEnrollmentDetail resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetTPMEnrollmentDetailSubResourceData(m []*models.TPMEnrollmentDetail) (d []*map[string]interface{}) {
	for _, TPMEnrollmentDetailModel := range m {
		if TPMEnrollmentDetailModel != nil {
			properties := make(map[string]interface{})
			properties["type"] = TPMEnrollmentDetailModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the TPMEnrollmentDetail resource defined in the Terraform configuration
func TPMEnrollmentDetailSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the TPMEnrollmentDetail resource
func GetTPMEnrollmentDetailPropertyFields() (t []string) {
	return []string{
		"type",
	}
}
