package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GroupCertificateEnrollmentModel(d *schema.ResourceData) *models.GroupCertificateEnrollment {
	groupName, _ := d.Get("group_name").(string)
	return &models.GroupCertificateEnrollment{
		GroupName: groupName,
	}
}

func GroupCertificateEnrollmentModelFromMap(m map[string]interface{}) *models.GroupCertificateEnrollment {
	groupName := m["group_name"].(string)
	return &models.GroupCertificateEnrollment{
		GroupName: groupName,
	}
}

func SetGroupCertificateEnrollmentResourceData(d *schema.ResourceData, m *models.GroupCertificateEnrollment) {
	d.Set("group_name", m.GroupName)
}

func SetGroupCertificateEnrollmentSubResourceData(m []*models.GroupCertificateEnrollment) (d []*map[string]interface{}) {
	for _, GroupCertificateEnrollmentModel := range m {
		if GroupCertificateEnrollmentModel != nil {
			properties := make(map[string]interface{})
			properties["group_name"] = GroupCertificateEnrollmentModel.GroupName
			d = append(d, &properties)
		}
	}
	return
}

func GroupCertificateEnrollmentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"group_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetGroupCertificateEnrollmentPropertyFields() (t []string) {
	return []string{
		"group_name",
	}
}
