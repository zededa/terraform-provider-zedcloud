package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFrontendLoginModeResponseModeModel(d *schema.ResourceData) *models.AAAFrontendLoginModeResponseMode {
	aAAFrontendLoginModeResponseMode, _ := d.Get("a_a_a_frontend_login_mode_response_mode").(models.AAAFrontendLoginModeResponseMode)
	return &aAAFrontendLoginModeResponseMode
}

func AAAFrontendLoginModeResponseModeModelFromMap(m map[string]interface{}) *models.AAAFrontendLoginModeResponseMode {
	aAAFrontendLoginModeResponseMode := m["a_a_a_frontend_login_mode_response_mode"].(models.AAAFrontendLoginModeResponseMode)
	return &aAAFrontendLoginModeResponseMode
}

func SetAAAFrontendLoginModeResponseModeResourceData(d *schema.ResourceData, m *models.AAAFrontendLoginModeResponseMode) {
}

func SetAAAFrontendLoginModeResponseModeSubResourceData(m []*models.AAAFrontendLoginModeResponseMode) (d []*map[string]interface{}) {
	for _, AAAFrontendLoginModeResponseModeModel := range m {
		if AAAFrontendLoginModeResponseModeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFrontendLoginModeResponseModeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAAAFrontendLoginModeResponseModePropertyFields() (t []string) {
	return []string{}
}
