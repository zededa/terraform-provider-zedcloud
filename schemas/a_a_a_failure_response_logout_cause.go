package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFailureResponseLogoutCauseModel(d *schema.ResourceData) *models.AAAFailureResponseLogoutCause {
	aAAFailureResponseLogoutCause, _ := d.Get("a_a_a_failure_response_logout_cause").(models.AAAFailureResponseLogoutCause)
	return &aAAFailureResponseLogoutCause
}

func AAAFailureResponseLogoutCauseModelFromMap(m map[string]interface{}) *models.AAAFailureResponseLogoutCause {
	aAAFailureResponseLogoutCause := m["a_a_a_failure_response_logout_cause"].(models.AAAFailureResponseLogoutCause)
	return &aAAFailureResponseLogoutCause
}

func SetAAAFailureResponseLogoutCauseResourceData(d *schema.ResourceData, m *models.AAAFailureResponseLogoutCause) {
}

func SetAAAFailureResponseLogoutCauseSubResourceData(m []*models.AAAFailureResponseLogoutCause) (d []*map[string]interface{}) {
	for _, AAAFailureResponseLogoutCauseModel := range m {
		if AAAFailureResponseLogoutCauseModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFailureResponseLogoutCauseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAAAFailureResponseLogoutCausePropertyFields() (t []string) {
	return []string{}
}
