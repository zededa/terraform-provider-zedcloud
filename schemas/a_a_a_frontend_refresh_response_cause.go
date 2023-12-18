package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFrontendRefreshResponseCauseModel(d *schema.ResourceData) *models.AAAFrontendRefreshResponseCause {
	aAAFrontendRefreshResponseCause, _ := d.Get("a_a_a_frontend_refresh_response_cause").(models.AAAFrontendRefreshResponseCause)
	return &aAAFrontendRefreshResponseCause
}

func AAAFrontendRefreshResponseCauseModelFromMap(m map[string]interface{}) *models.AAAFrontendRefreshResponseCause {
	aAAFrontendRefreshResponseCause := m["a_a_a_frontend_refresh_response_cause"].(models.AAAFrontendRefreshResponseCause)
	return &aAAFrontendRefreshResponseCause
}

func SetAAAFrontendRefreshResponseCauseResourceData(d *schema.ResourceData, m *models.AAAFrontendRefreshResponseCause) {
}

func SetAAAFrontendRefreshResponseCauseSubResourceData(m []*models.AAAFrontendRefreshResponseCause) (d []*map[string]interface{}) {
	for _, AAAFrontendRefreshResponseCauseModel := range m {
		if AAAFrontendRefreshResponseCauseModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFrontendRefreshResponseCauseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAAAFrontendRefreshResponseCausePropertyFields() (t []string) {
	return []string{}
}
