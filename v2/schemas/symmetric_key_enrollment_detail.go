package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SymmetricKeyEnrollmentDetailModel(d *schema.ResourceData) *models.SymmetricKeyEnrollmentDetail {
	var groupSymmetricKeyEnrollment *models.GroupSymmetricKeyEnrollment // GroupSymmetricKeyEnrollment
	groupSymmetricKeyEnrollmentInterface, groupSymmetricKeyEnrollmentIsSet := d.GetOk("group_symmetric_key_enrollment")
	if groupSymmetricKeyEnrollmentIsSet && groupSymmetricKeyEnrollmentInterface != nil {
		groupSymmetricKeyEnrollmentMap := groupSymmetricKeyEnrollmentInterface.([]interface{})
		if len(groupSymmetricKeyEnrollmentMap) > 0 {
			groupSymmetricKeyEnrollment = GroupSymmetricKeyEnrollmentModelFromMap(groupSymmetricKeyEnrollmentMap[0].(map[string]interface{}))
		}
	}
	var individualSymmetricKeyEnrollment *models.IndividualSymmetricKeyEnrollment // IndividualSymmetricKeyEnrollment
	individualSymmetricKeyEnrollmentInterface, individualSymmetricKeyEnrollmentIsSet := d.GetOk("individual_symmetric_key_enrollment")
	if individualSymmetricKeyEnrollmentIsSet && individualSymmetricKeyEnrollmentInterface != nil {
		individualSymmetricKeyEnrollmentMap := individualSymmetricKeyEnrollmentInterface.([]interface{})
		if len(individualSymmetricKeyEnrollmentMap) > 0 {
			individualSymmetricKeyEnrollment = IndividualSymmetricKeyEnrollmentModelFromMap(individualSymmetricKeyEnrollmentMap[0].(map[string]interface{}))
		}
	}
	var typeVar *models.EnrollmentType // EnrollmentType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewEnrollmentType(models.EnrollmentType(typeModel))
	}
	return &models.SymmetricKeyEnrollmentDetail{
		GroupSymmetricKeyEnrollment:      groupSymmetricKeyEnrollment,
		IndividualSymmetricKeyEnrollment: individualSymmetricKeyEnrollment,
		Type:                             typeVar,
	}
}

func SymmetricKeyEnrollmentDetailModelFromMap(m map[string]interface{}) *models.SymmetricKeyEnrollmentDetail {
	var groupSymmetricKeyEnrollment *models.GroupSymmetricKeyEnrollment // GroupSymmetricKeyEnrollment
	groupSymmetricKeyEnrollmentInterface, groupSymmetricKeyEnrollmentIsSet := m["group_symmetric_key_enrollment"]
	if groupSymmetricKeyEnrollmentIsSet && groupSymmetricKeyEnrollmentInterface != nil {
		groupSymmetricKeyEnrollmentMap := groupSymmetricKeyEnrollmentInterface.([]interface{})
		if len(groupSymmetricKeyEnrollmentMap) > 0 {
			groupSymmetricKeyEnrollment = GroupSymmetricKeyEnrollmentModelFromMap(groupSymmetricKeyEnrollmentMap[0].(map[string]interface{}))
		}
	}
	//
	var individualSymmetricKeyEnrollment *models.IndividualSymmetricKeyEnrollment // IndividualSymmetricKeyEnrollment
	individualSymmetricKeyEnrollmentInterface, individualSymmetricKeyEnrollmentIsSet := m["individual_symmetric_key_enrollment"]
	if individualSymmetricKeyEnrollmentIsSet && individualSymmetricKeyEnrollmentInterface != nil {
		individualSymmetricKeyEnrollmentMap := individualSymmetricKeyEnrollmentInterface.([]interface{})
		if len(individualSymmetricKeyEnrollmentMap) > 0 {
			individualSymmetricKeyEnrollment = IndividualSymmetricKeyEnrollmentModelFromMap(individualSymmetricKeyEnrollmentMap[0].(map[string]interface{}))
		}
	}
	//
	var typeVar *models.EnrollmentType // EnrollmentType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewEnrollmentType(models.EnrollmentType(typeModel))
	}
	return &models.SymmetricKeyEnrollmentDetail{
		GroupSymmetricKeyEnrollment:      groupSymmetricKeyEnrollment,
		IndividualSymmetricKeyEnrollment: individualSymmetricKeyEnrollment,
		Type:                             typeVar,
	}
}

func SetSymmetricKeyEnrollmentDetailResourceData(d *schema.ResourceData, m *models.SymmetricKeyEnrollmentDetail) {
	d.Set("group_symmetric_key_enrollment", SetGroupSymmetricKeyEnrollmentSubResourceData([]*models.GroupSymmetricKeyEnrollment{m.GroupSymmetricKeyEnrollment}))
	d.Set("individual_symmetric_key_enrollment", SetIndividualSymmetricKeyEnrollmentSubResourceData([]*models.IndividualSymmetricKeyEnrollment{m.IndividualSymmetricKeyEnrollment}))
	d.Set("type", m.Type)
}

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

func SymmetricKeyEnrollmentDetail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"group_symmetric_key_enrollment": {
			Description: ``,
			Type:        schema.TypeList, //GoType: GroupSymmetricKeyEnrollment
			Elem: &schema.Resource{
				Schema: GroupSymmetricKeyEnrollmentSchema(),
			},
			Optional: true,
		},

		"individual_symmetric_key_enrollment": {
			Description: ``,
			Type:        schema.TypeList, //GoType: IndividualSymmetricKeyEnrollment
			Elem: &schema.Resource{
				Schema: IndividualSymmetricKeyEnrollmentSchema(),
			},
			Optional: true,
		},

		"type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetSymmetricKeyEnrollmentDetailPropertyFields() (t []string) {
	return []string{
		"group_symmetric_key_enrollment",
		"individual_symmetric_key_enrollment",
		"type",
	}
}
