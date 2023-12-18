package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAASuccessResponseTypeModel(d *schema.ResourceData) *models.AAASuccessResponseType {
	aAASuccessResponseType, _ := d.Get("a_a_a_success_response_type").(models.AAASuccessResponseType)
	return &aAASuccessResponseType
}

func AAASuccessResponseTypeModelFromMap(m map[string]interface{}) *models.AAASuccessResponseType {
	aAASuccessResponseType := m["a_a_a_success_response_type"].(models.AAASuccessResponseType)
	return &aAASuccessResponseType
}

func SetAAASuccessResponseTypeResourceData(d *schema.ResourceData, m *models.AAASuccessResponseType) {
}

func SetAAASuccessResponseTypeSubResourceData(m []*models.AAASuccessResponseType) (d []*map[string]interface{}) {
	for _, AAASuccessResponseTypeModel := range m {
		if AAASuccessResponseTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AAASuccessResponseTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAAASuccessResponseTypePropertyFields() (t []string) {
	return []string{}
}
