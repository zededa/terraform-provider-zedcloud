package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate CertificateEnrollmentDetail resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func CertificateEnrollmentDetailModel(d *schema.ResourceData) *models.CertificateEnrollmentDetail {
	certificateEnrollmentDetail, _ := d.Get("certificate_enrollment_detail").(models.CertificateEnrollmentDetail)
	return &certificateEnrollmentDetail
}

func CertificateEnrollmentDetailModelFromMap(m map[string]interface{}) *models.CertificateEnrollmentDetail {
	certificateEnrollmentDetail := m["certificate_enrollment_detail"].(models.CertificateEnrollmentDetail)
	return &certificateEnrollmentDetail
}

// Update the underlying CertificateEnrollmentDetail resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetCertificateEnrollmentDetailResourceData(d *schema.ResourceData, m *models.CertificateEnrollmentDetail) {
}

// Iterate through and update the CertificateEnrollmentDetail resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetCertificateEnrollmentDetailSubResourceData(m []*models.CertificateEnrollmentDetail) (d []*map[string]interface{}) {
	for _, CertificateEnrollmentDetailModel := range m {
		if CertificateEnrollmentDetailModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the CertificateEnrollmentDetail resource defined in the Terraform configuration
func CertificateEnrollmentDetailSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the CertificateEnrollmentDetail resource
func GetCertificateEnrollmentDetailPropertyFields() (t []string) {
	return []string{}
}
