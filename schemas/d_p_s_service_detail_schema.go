package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DPSServiceDetail resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DPSServiceDetailModel(d *schema.ResourceData) *models.DPSServiceDetail {
	var enrollment *models.EnrollmentDetail // EnrollmentDetail
	enrollmentInterface, enrollmentIsSet := d.GetOk("enrollment")
	if enrollmentIsSet {
		enrollmentMap := enrollmentInterface.([]interface{})[0].(map[string]interface{})
		enrollment = EnrollmentDetailModelFromMap(enrollmentMap)
	}
	var serviceDetail *models.AzureResourceAndServiceDetail // AzureResourceAndServiceDetail
	serviceDetailInterface, serviceDetailIsSet := d.GetOk("service_detail")
	if serviceDetailIsSet {
		serviceDetailMap := serviceDetailInterface.([]interface{})[0].(map[string]interface{})
		serviceDetail = AzureResourceAndServiceDetailModelFromMap(serviceDetailMap)
	}
	return &models.DPSServiceDetail{
		Enrollment:    enrollment,
		ServiceDetail: serviceDetail,
	}
}

func DPSServiceDetailModelFromMap(m map[string]interface{}) *models.DPSServiceDetail {
	var enrollment *models.EnrollmentDetail // EnrollmentDetail
	enrollmentInterface, enrollmentIsSet := m["enrollment"]
	if enrollmentIsSet {
		enrollmentMap := enrollmentInterface.([]interface{})[0].(map[string]interface{})
		enrollment = EnrollmentDetailModelFromMap(enrollmentMap)
	}
	//
	var serviceDetail *models.AzureResourceAndServiceDetail // AzureResourceAndServiceDetail
	serviceDetailInterface, serviceDetailIsSet := m["service_detail"]
	if serviceDetailIsSet {
		serviceDetailMap := serviceDetailInterface.([]interface{})[0].(map[string]interface{})
		serviceDetail = AzureResourceAndServiceDetailModelFromMap(serviceDetailMap)
	}
	//
	return &models.DPSServiceDetail{
		Enrollment:    enrollment,
		ServiceDetail: serviceDetail,
	}
}

// Update the underlying DPSServiceDetail resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDPSServiceDetailResourceData(d *schema.ResourceData, m *models.DPSServiceDetail) {
	d.Set("enrollment", SetEnrollmentDetailSubResourceData([]*models.EnrollmentDetail{m.Enrollment}))
	d.Set("service_detail", SetAzureResourceAndServiceDetailSubResourceData([]*models.AzureResourceAndServiceDetail{m.ServiceDetail}))
}

// Iterate throught and update the DPSServiceDetail resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDPSServiceDetailSubResourceData(m []*models.DPSServiceDetail) (d []*map[string]interface{}) {
	for _, DPSServiceDetailModel := range m {
		if DPSServiceDetailModel != nil {
			properties := make(map[string]interface{})
			properties["enrollment"] = SetEnrollmentDetailSubResourceData([]*models.EnrollmentDetail{DPSServiceDetailModel.Enrollment})
			properties["service_detail"] = SetAzureResourceAndServiceDetailSubResourceData([]*models.AzureResourceAndServiceDetail{DPSServiceDetailModel.ServiceDetail})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DPSServiceDetail resource defined in the Terraform configuration
func DPSServiceDetailSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enrollment": {
			Description: ``,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"service_detail": {
			Description: ``,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the DPSServiceDetail resource
func GetDPSServiceDetailPropertyFields() (t []string) {
	return []string{
		"enrollment",
		"service_detail",
	}
}
