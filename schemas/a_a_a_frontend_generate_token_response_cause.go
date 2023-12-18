package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFrontendGenerateTokenResponseCauseModel(d *schema.ResourceData) *models.AAAFrontendGenerateTokenResponseCause {
	aAAFrontendGenerateTokenResponseCause, _ := d.Get("a_a_a_frontend_generate_token_response_cause").(models.AAAFrontendGenerateTokenResponseCause)
	return &aAAFrontendGenerateTokenResponseCause
}

func AAAFrontendGenerateTokenResponseCauseModelFromMap(m map[string]interface{}) *models.AAAFrontendGenerateTokenResponseCause {
	aAAFrontendGenerateTokenResponseCause := m["a_a_a_frontend_generate_token_response_cause"].(models.AAAFrontendGenerateTokenResponseCause)
	return &aAAFrontendGenerateTokenResponseCause
}

func SetAAAFrontendGenerateTokenResponseCauseResourceData(d *schema.ResourceData, m *models.AAAFrontendGenerateTokenResponseCause) {
}

func SetAAAFrontendGenerateTokenResponseCauseSubResourceData(m []*models.AAAFrontendGenerateTokenResponseCause) (d []*map[string]interface{}) {
	for _, AAAFrontendGenerateTokenResponseCauseModel := range m {
		if AAAFrontendGenerateTokenResponseCauseModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFrontendGenerateTokenResponseCauseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAAAFrontendGenerateTokenResponseCausePropertyFields() (t []string) {
	return []string{}
}
