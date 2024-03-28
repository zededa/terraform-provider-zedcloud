package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func TPMEnrollmentDetailModel(d *schema.ResourceData) *models.TPMEnrollmentDetail {
	var typeVar *models.EnrollmentType // EnrollmentType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewEnrollmentType(models.EnrollmentType(typeModel))
	}
	return &models.TPMEnrollmentDetail{
		Type: typeVar,
	}
}

func TPMEnrollmentDetailModelFromMap(m map[string]interface{}) *models.TPMEnrollmentDetail {
	var typeVar *models.EnrollmentType // EnrollmentType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewEnrollmentType(models.EnrollmentType(typeModel))
	}
	return &models.TPMEnrollmentDetail{
		Type: typeVar,
	}
}

func SetTPMEnrollmentDetailResourceData(d *schema.ResourceData, m *models.TPMEnrollmentDetail) {
	d.Set("type", m.Type)
}

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

func TPMEnrollmentDetail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetTPMEnrollmentDetailPropertyFields() (t []string) {
	return []string{
		"type",
	}
}
