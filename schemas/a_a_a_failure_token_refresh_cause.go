package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFailureTokenRefreshCauseModel(d *schema.ResourceData) *models.AAAFailureTokenRefreshCause {
	aAAFailureTokenRefreshCause, _ := d.Get("a_a_a_failure_token_refresh_cause").(models.AAAFailureTokenRefreshCause)
	return &aAAFailureTokenRefreshCause
}

func AAAFailureTokenRefreshCauseModelFromMap(m map[string]interface{}) *models.AAAFailureTokenRefreshCause {
	aAAFailureTokenRefreshCause := m["a_a_a_failure_token_refresh_cause"].(models.AAAFailureTokenRefreshCause)
	return &aAAFailureTokenRefreshCause
}

func SetAAAFailureTokenRefreshCauseResourceData(d *schema.ResourceData, m *models.AAAFailureTokenRefreshCause) {
}

func SetAAAFailureTokenRefreshCauseSubResourceData(m []*models.AAAFailureTokenRefreshCause) (d []*map[string]interface{}) {
	for _, AAAFailureTokenRefreshCauseModel := range m {
		if AAAFailureTokenRefreshCauseModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFailureTokenRefreshCauseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAAAFailureTokenRefreshCausePropertyFields() (t []string) {
	return []string{}
}
