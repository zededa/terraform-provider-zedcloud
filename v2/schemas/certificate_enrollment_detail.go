package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func CertificateEnrollmentDetailModel(d *schema.ResourceData) *models.CertificateEnrollmentDetail {
	var groupCertificateEnrollment *models.GroupCertificateEnrollment // GroupCertificateEnrollment
	groupCertificateEnrollmentInterface, groupCertificateEnrollmentIsSet := d.GetOk("group_certificate_enrollment")
	if groupCertificateEnrollmentIsSet && groupCertificateEnrollmentInterface != nil {
		groupCertificateEnrollmentMap := groupCertificateEnrollmentInterface.([]interface{})
		if len(groupCertificateEnrollmentMap) > 0 {
			groupCertificateEnrollment = GroupCertificateEnrollmentModelFromMap(groupCertificateEnrollmentMap[0].(map[string]interface{}))
		}
	}
	individualCertificateEnrollment, _ := d.Get("individual_certificate_enrollment").(models.IndividualCertificateEnrollment) // IndividualCertificateEnrollment
	var typeVar *models.EnrollmentType                                                                                        // EnrollmentType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewEnrollmentType(models.EnrollmentType(typeModel))
	}
	return &models.CertificateEnrollmentDetail{
		GroupCertificateEnrollment:      groupCertificateEnrollment,
		IndividualCertificateEnrollment: individualCertificateEnrollment,
		Type:                            typeVar,
	}
}

func CertificateEnrollmentDetailModelFromMap(m map[string]interface{}) *models.CertificateEnrollmentDetail {
	var groupCertificateEnrollment *models.GroupCertificateEnrollment // GroupCertificateEnrollment
	groupCertificateEnrollmentInterface, groupCertificateEnrollmentIsSet := m["group_certificate_enrollment"]
	if groupCertificateEnrollmentIsSet && groupCertificateEnrollmentInterface != nil {
		groupCertificateEnrollmentMap := groupCertificateEnrollmentInterface.([]interface{})
		if len(groupCertificateEnrollmentMap) > 0 {
			groupCertificateEnrollment = GroupCertificateEnrollmentModelFromMap(groupCertificateEnrollmentMap[0].(map[string]interface{}))
		}
	}
	//
	individualCertificateEnrollment := m["individual_certificate_enrollment"].(models.IndividualCertificateEnrollment) // IndividualCertificateEnrollment
	var typeVar *models.EnrollmentType                                                                                 // EnrollmentType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewEnrollmentType(models.EnrollmentType(typeModel))
	}
	return &models.CertificateEnrollmentDetail{
		GroupCertificateEnrollment:      groupCertificateEnrollment,
		IndividualCertificateEnrollment: individualCertificateEnrollment,
		Type:                            typeVar,
	}
}

func SetCertificateEnrollmentDetailResourceData(d *schema.ResourceData, m *models.CertificateEnrollmentDetail) {
	d.Set("group_certificate_enrollment", SetGroupCertificateEnrollmentSubResourceData([]*models.GroupCertificateEnrollment{m.GroupCertificateEnrollment}))
	d.Set("individual_certificate_enrollment", m.IndividualCertificateEnrollment)
	d.Set("type", m.Type)
}

func SetCertificateEnrollmentDetailSubResourceData(m []*models.CertificateEnrollmentDetail) (d []*map[string]interface{}) {
	for _, CertificateEnrollmentDetailModel := range m {
		if CertificateEnrollmentDetailModel != nil {
			properties := make(map[string]interface{})
			properties["group_certificate_enrollment"] = SetGroupCertificateEnrollmentSubResourceData([]*models.GroupCertificateEnrollment{CertificateEnrollmentDetailModel.GroupCertificateEnrollment})
			properties["individual_certificate_enrollment"] = CertificateEnrollmentDetailModel.IndividualCertificateEnrollment
			properties["type"] = CertificateEnrollmentDetailModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func CertificateEnrollmentDetail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"group_certificate_enrollment": {
			Description: ``,
			Type:        schema.TypeList, //GoType: GroupCertificateEnrollment
			Elem: &schema.Resource{
				Schema: GroupCertificateEnrollmentSchema(),
			},
			Optional: true,
		},

		"individual_certificate_enrollment": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetCertificateEnrollmentDetailPropertyFields() (t []string) {
	return []string{
		"group_certificate_enrollment",
		"individual_certificate_enrollment",
		"type",
	}
}
