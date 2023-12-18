package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFrontendLoginResponseCauseModel(d *schema.ResourceData) *models.AAAFrontendLoginResponseCause {
	aAAFrontendLoginResponseCause, _ := d.Get("a_a_a_frontend_login_response_cause").(models.AAAFrontendLoginResponseCause)
	return &aAAFrontendLoginResponseCause
}

func AAAFrontendLoginResponseCauseModelFromMap(m map[string]interface{}) *models.AAAFrontendLoginResponseCause {
	aAAFrontendLoginResponseCause := m["a_a_a_frontend_login_response_cause"].(models.AAAFrontendLoginResponseCause)
	return &aAAFrontendLoginResponseCause
}

func SetAAAFrontendLoginResponseCauseResourceData(d *schema.ResourceData, m *models.AAAFrontendLoginResponseCause) {
}

func SetAAAFrontendLoginResponseCauseSubResourceData(m []*models.AAAFrontendLoginResponseCause) (d []*map[string]interface{}) {
	for _, AAAFrontendLoginResponseCauseModel := range m {
		if AAAFrontendLoginResponseCauseModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFrontendLoginResponseCauseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAAAFrontendLoginResponseCausePropertyFields() (t []string) {
	return []string{}
}
