package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFailureResponseTypeModel(d *schema.ResourceData) *models.AAAFailureResponseType {
	aAAFailureResponseType, _ := d.Get("a_a_a_failure_response_type").(models.AAAFailureResponseType)
	return &aAAFailureResponseType
}

func AAAFailureResponseTypeModelFromMap(m map[string]interface{}) *models.AAAFailureResponseType {
	aAAFailureResponseType := m["a_a_a_failure_response_type"].(models.AAAFailureResponseType)
	return &aAAFailureResponseType
}

func SetAAAFailureResponseTypeResourceData(d *schema.ResourceData, m *models.AAAFailureResponseType) {
}

func SetAAAFailureResponseTypeSubResourceData(m []*models.AAAFailureResponseType) (d []*map[string]interface{}) {
	for _, AAAFailureResponseTypeModel := range m {
		if AAAFailureResponseTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFailureResponseTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAAAFailureResponseTypePropertyFields() (t []string) {
	return []string{}
}
