package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFailureResponseGenerateTokenCauseModel(d *schema.ResourceData) *models.AAAFailureResponseGenerateTokenCause {
	aAAFailureResponseGenerateTokenCause, _ := d.Get("a_a_a_failure_response_generate_token_cause").(models.AAAFailureResponseGenerateTokenCause)
	return &aAAFailureResponseGenerateTokenCause
}

func AAAFailureResponseGenerateTokenCauseModelFromMap(m map[string]interface{}) *models.AAAFailureResponseGenerateTokenCause {
	aAAFailureResponseGenerateTokenCause := m["a_a_a_failure_response_generate_token_cause"].(models.AAAFailureResponseGenerateTokenCause)
	return &aAAFailureResponseGenerateTokenCause
}

func SetAAAFailureResponseGenerateTokenCauseResourceData(d *schema.ResourceData, m *models.AAAFailureResponseGenerateTokenCause) {
}

func SetAAAFailureResponseGenerateTokenCauseSubResourceData(m []*models.AAAFailureResponseGenerateTokenCause) (d []*map[string]interface{}) {
	for _, AAAFailureResponseGenerateTokenCauseModel := range m {
		if AAAFailureResponseGenerateTokenCauseModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFailureResponseGenerateTokenCauseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAAAFailureResponseGenerateTokenCausePropertyFields() (t []string) {
	return []string{}
}
