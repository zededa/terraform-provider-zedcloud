package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFrontendLogoutResponseCauseModel(d *schema.ResourceData) *models.AAAFrontendLogoutResponseCause {
	aAAFrontendLogoutResponseCause, _ := d.Get("a_a_a_frontend_logout_response_cause").(models.AAAFrontendLogoutResponseCause)
	return &aAAFrontendLogoutResponseCause
}

func AAAFrontendLogoutResponseCauseModelFromMap(m map[string]interface{}) *models.AAAFrontendLogoutResponseCause {
	aAAFrontendLogoutResponseCause := m["a_a_a_frontend_logout_response_cause"].(models.AAAFrontendLogoutResponseCause)
	return &aAAFrontendLogoutResponseCause
}

func SetAAAFrontendLogoutResponseCauseResourceData(d *schema.ResourceData, m *models.AAAFrontendLogoutResponseCause) {
}

func SetAAAFrontendLogoutResponseCauseSubResourceData(m []*models.AAAFrontendLogoutResponseCause) (d []*map[string]interface{}) {
	for _, AAAFrontendLogoutResponseCauseModel := range m {
		if AAAFrontendLogoutResponseCauseModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFrontendLogoutResponseCauseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAAAFrontendLogoutResponseCausePropertyFields() (t []string) {
	return []string{}
}
