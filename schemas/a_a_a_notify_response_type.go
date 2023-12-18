package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAANotifyResponseTypeModel(d *schema.ResourceData) *models.AAANotifyResponseType {
	aAANotifyResponseType, _ := d.Get("a_a_a_notify_response_type").(models.AAANotifyResponseType)
	return &aAANotifyResponseType
}

func AAANotifyResponseTypeModelFromMap(m map[string]interface{}) *models.AAANotifyResponseType {
	aAANotifyResponseType := m["a_a_a_notify_response_type"].(models.AAANotifyResponseType)
	return &aAANotifyResponseType
}

func SetAAANotifyResponseTypeResourceData(d *schema.ResourceData, m *models.AAANotifyResponseType) {
}

func SetAAANotifyResponseTypeSubResourceData(m []*models.AAANotifyResponseType) (d []*map[string]interface{}) {
	for _, AAANotifyResponseTypeModel := range m {
		if AAANotifyResponseTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AAANotifyResponseTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAAANotifyResponseTypePropertyFields() (t []string) {
	return []string{}
}
