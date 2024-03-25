package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GroupSymmetricKeyEnrollmentModel(d *schema.ResourceData) *models.GroupSymmetricKeyEnrollment {
	groupName, _ := d.Get("group_name").(string)
	return &models.GroupSymmetricKeyEnrollment{
		GroupName: groupName,
	}
}

func GroupSymmetricKeyEnrollmentModelFromMap(m map[string]interface{}) *models.GroupSymmetricKeyEnrollment {
	groupName := m["group_name"].(string)
	return &models.GroupSymmetricKeyEnrollment{
		GroupName: groupName,
	}
}

func SetGroupSymmetricKeyEnrollmentResourceData(d *schema.ResourceData, m *models.GroupSymmetricKeyEnrollment) {
	d.Set("group_name", m.GroupName)
}

func SetGroupSymmetricKeyEnrollmentSubResourceData(m []*models.GroupSymmetricKeyEnrollment) (d []*map[string]interface{}) {
	for _, GroupSymmetricKeyEnrollmentModel := range m {
		if GroupSymmetricKeyEnrollmentModel != nil {
			properties := make(map[string]interface{})
			properties["group_name"] = GroupSymmetricKeyEnrollmentModel.GroupName
			d = append(d, &properties)
		}
	}
	return
}

func GroupSymmetricKeyEnrollmentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"group_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetGroupSymmetricKeyEnrollmentPropertyFields() (t []string) {
	return []string{
		"group_name",
	}
}
