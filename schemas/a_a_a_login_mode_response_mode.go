package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAALoginModeResponseModeModel(d *schema.ResourceData) *models.AAALoginModeResponseMode {
	aAALoginModeResponseMode, _ := d.Get("a_a_a_login_mode_response_mode").(models.AAALoginModeResponseMode)
	return &aAALoginModeResponseMode
}

func AAALoginModeResponseModeModelFromMap(m map[string]interface{}) *models.AAALoginModeResponseMode {
	aAALoginModeResponseMode := m["a_a_a_login_mode_response_mode"].(models.AAALoginModeResponseMode)
	return &aAALoginModeResponseMode
}

func SetAAALoginModeResponseModeResourceData(d *schema.ResourceData, m *models.AAALoginModeResponseMode) {
}

func SetAAALoginModeResponseModeSubResourceData(m []*models.AAALoginModeResponseMode) (d []*map[string]interface{}) {
	for _, AAALoginModeResponseModeModel := range m {
		if AAALoginModeResponseModeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AAALoginModeResponseModeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAAALoginModeResponseModePropertyFields() (t []string) {
	return []string{}
}
