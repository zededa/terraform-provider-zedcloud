package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFailureResponseLoginCauseModel(d *schema.ResourceData) *models.AAAFailureResponseLoginCause {
	aAAFailureResponseLoginCause, _ := d.Get("a_a_a_failure_response_login_cause").(models.AAAFailureResponseLoginCause)
	return &aAAFailureResponseLoginCause
}

func AAAFailureResponseLoginCauseModelFromMap(m map[string]interface{}) *models.AAAFailureResponseLoginCause {
	aAAFailureResponseLoginCause := m["a_a_a_failure_response_login_cause"].(models.AAAFailureResponseLoginCause)
	return &aAAFailureResponseLoginCause
}

func SetAAAFailureResponseLoginCauseResourceData(d *schema.ResourceData, m *models.AAAFailureResponseLoginCause) {
}

func SetAAAFailureResponseLoginCauseSubResourceData(m []*models.AAAFailureResponseLoginCause) (d []*map[string]interface{}) {
	for _, AAAFailureResponseLoginCauseModel := range m {
		if AAAFailureResponseLoginCauseModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFailureResponseLoginCauseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAAAFailureResponseLoginCausePropertyFields() (t []string) {
	return []string{}
}
