package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func EnterpriseTypeModel(d *schema.ResourceData) *models.EnterpriseType {
	enterpriseType, _ := d.Get("enterprise_type").(models.EnterpriseType)
	return &enterpriseType
}

func EnterpriseTypeModelFromMap(m map[string]interface{}) *models.EnterpriseType {
	enterpriseType := m["enterprise_type"].(models.EnterpriseType)
	return &enterpriseType
}

func SetEnterpriseTypeResourceData(d *schema.ResourceData, m *models.EnterpriseType) {
}

func SetEnterpriseTypeSubResourceData(m []*models.EnterpriseType) (d []*map[string]interface{}) {
	for _, EnterpriseTypeModel := range m {
		if EnterpriseTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func EnterpriseTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetEnterpriseTypePropertyFields() (t []string) {
	return []string{}
}
