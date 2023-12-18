package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFailureResponseCredentialChangeCauseModel(d *schema.ResourceData) *models.AAAFailureResponseCredentialChangeCause {
	aAAFailureResponseCredentialChangeCause, _ := d.Get("a_a_a_failure_response_credential_change_cause").(models.AAAFailureResponseCredentialChangeCause)
	return &aAAFailureResponseCredentialChangeCause
}

func AAAFailureResponseCredentialChangeCauseModelFromMap(m map[string]interface{}) *models.AAAFailureResponseCredentialChangeCause {
	aAAFailureResponseCredentialChangeCause := m["a_a_a_failure_response_credential_change_cause"].(models.AAAFailureResponseCredentialChangeCause)
	return &aAAFailureResponseCredentialChangeCause
}

func SetAAAFailureResponseCredentialChangeCauseResourceData(d *schema.ResourceData, m *models.AAAFailureResponseCredentialChangeCause) {
}

func SetAAAFailureResponseCredentialChangeCauseSubResourceData(m []*models.AAAFailureResponseCredentialChangeCause) (d []*map[string]interface{}) {
	for _, AAAFailureResponseCredentialChangeCauseModel := range m {
		if AAAFailureResponseCredentialChangeCauseModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFailureResponseCredentialChangeCauseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAAAFailureResponseCredentialChangeCausePropertyFields() (t []string) {
	return []string{}
}
