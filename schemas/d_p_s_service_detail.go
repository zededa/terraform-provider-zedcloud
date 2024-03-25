package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func DPSServiceDetailModel(d *schema.ResourceData) *models.DPSServiceDetail {
	var enrollment *models.EnrollmentDetail // EnrollmentDetail
	enrollmentInterface, enrollmentIsSet := d.GetOk("enrollment")
	if enrollmentIsSet && enrollmentInterface != nil {
		enrollmentMap := enrollmentInterface.([]interface{})
		if len(enrollmentMap) > 0 {
			enrollment = EnrollmentDetailModelFromMap(enrollmentMap[0].(map[string]interface{}))
		}
	}
	var serviceDetail *models.AzureResourceAndServiceDetail // AzureResourceAndServiceDetail
	serviceDetailInterface, serviceDetailIsSet := d.GetOk("service_detail")
	if serviceDetailIsSet && serviceDetailInterface != nil {
		serviceDetailMap := serviceDetailInterface.([]interface{})
		if len(serviceDetailMap) > 0 {
			serviceDetail = AzureResourceAndServiceDetailModelFromMap(serviceDetailMap[0].(map[string]interface{}))
		}
	}
	return &models.DPSServiceDetail{
		Enrollment:    enrollment,
		ServiceDetail: serviceDetail,
	}
}

func DPSServiceDetailModelFromMap(m map[string]interface{}) *models.DPSServiceDetail {
	var enrollment *models.EnrollmentDetail // EnrollmentDetail
	enrollmentInterface, enrollmentIsSet := m["enrollment"]
	if enrollmentIsSet && enrollmentInterface != nil {
		enrollmentMap := enrollmentInterface.([]interface{})
		if len(enrollmentMap) > 0 {
			enrollment = EnrollmentDetailModelFromMap(enrollmentMap[0].(map[string]interface{}))
		}
	}
	//
	var serviceDetail *models.AzureResourceAndServiceDetail // AzureResourceAndServiceDetail
	serviceDetailInterface, serviceDetailIsSet := m["service_detail"]
	if serviceDetailIsSet && serviceDetailInterface != nil {
		serviceDetailMap := serviceDetailInterface.([]interface{})
		if len(serviceDetailMap) > 0 {
			serviceDetail = AzureResourceAndServiceDetailModelFromMap(serviceDetailMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.DPSServiceDetail{
		Enrollment:    enrollment,
		ServiceDetail: serviceDetail,
	}
}

func SetDPSServiceDetailResourceData(d *schema.ResourceData, m *models.DPSServiceDetail) {
	d.Set("enrollment", SetEnrollmentDetailSubResourceData([]*models.EnrollmentDetail{m.Enrollment}))
	d.Set("service_detail", SetAzureResourceAndServiceDetailSubResourceData([]*models.AzureResourceAndServiceDetail{m.ServiceDetail}))
}

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

func DPSServiceDetail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enrollment": {
			Description: ``,
			Type:        schema.TypeList, //GoType: EnrollmentDetail
			Elem: &schema.Resource{
				Schema: EnrollmentDetail(),
			},
			Optional: true,
		},

		"service_detail": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AzureResourceAndServiceDetail
			Elem: &schema.Resource{
				Schema: AzureResourceAndServiceDetail(),
			},
			Optional: true,
		},
	}
}

func GetDPSServiceDetailPropertyFields() (t []string) {
	return []string{
		"enrollment",
		"service_detail",
	}
}
