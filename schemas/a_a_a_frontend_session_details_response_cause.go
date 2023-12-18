package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFrontendSessionDetailsResponseCauseModel(d *schema.ResourceData) *models.AAAFrontendSessionDetailsResponseCause {
	aAAFrontendSessionDetailsResponseCause, _ := d.Get("a_a_a_frontend_session_details_response_cause").(models.AAAFrontendSessionDetailsResponseCause)
	return &aAAFrontendSessionDetailsResponseCause
}

func AAAFrontendSessionDetailsResponseCauseModelFromMap(m map[string]interface{}) *models.AAAFrontendSessionDetailsResponseCause {
	aAAFrontendSessionDetailsResponseCause := m["a_a_a_frontend_session_details_response_cause"].(models.AAAFrontendSessionDetailsResponseCause)
	return &aAAFrontendSessionDetailsResponseCause
}

func SetAAAFrontendSessionDetailsResponseCauseResourceData(d *schema.ResourceData, m *models.AAAFrontendSessionDetailsResponseCause) {
}

func SetAAAFrontendSessionDetailsResponseCauseSubResourceData(m []*models.AAAFrontendSessionDetailsResponseCause) (d []*map[string]interface{}) {
	for _, AAAFrontendSessionDetailsResponseCauseModel := range m {
		if AAAFrontendSessionDetailsResponseCauseModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFrontendSessionDetailsResponseCauseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAAAFrontendSessionDetailsResponseCausePropertyFields() (t []string) {
	return []string{}
}
