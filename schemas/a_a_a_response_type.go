package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAResponseTypeModel(d *schema.ResourceData) *models.AAAResponseType {
	aAAResponseType, _ := d.Get("a_a_a_response_type").(models.AAAResponseType)
	return &aAAResponseType
}

func AAAResponseTypeModelFromMap(m map[string]interface{}) *models.AAAResponseType {
	aAAResponseType := m["a_a_a_response_type"].(models.AAAResponseType)
	return &aAAResponseType
}

func SetAAAResponseTypeResourceData(d *schema.ResourceData, m *models.AAAResponseType) {
}

func SetAAAResponseTypeSubResourceData(m []*models.AAAResponseType) (d []*map[string]interface{}) {
	for _, AAAResponseTypeModel := range m {
		if AAAResponseTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AAAResponseTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAAAResponseTypePropertyFields() (t []string) {
	return []string{}
}
