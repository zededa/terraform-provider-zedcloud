package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFailureResponseSessionDetailsCauseModel(d *schema.ResourceData) *models.AAAFailureResponseSessionDetailsCause {
	aAAFailureResponseSessionDetailsCause, _ := d.Get("a_a_a_failure_response_session_details_cause").(models.AAAFailureResponseSessionDetailsCause)
	return &aAAFailureResponseSessionDetailsCause
}

func AAAFailureResponseSessionDetailsCauseModelFromMap(m map[string]interface{}) *models.AAAFailureResponseSessionDetailsCause {
	aAAFailureResponseSessionDetailsCause := m["a_a_a_failure_response_session_details_cause"].(models.AAAFailureResponseSessionDetailsCause)
	return &aAAFailureResponseSessionDetailsCause
}

func SetAAAFailureResponseSessionDetailsCauseResourceData(d *schema.ResourceData, m *models.AAAFailureResponseSessionDetailsCause) {
}

func SetAAAFailureResponseSessionDetailsCauseSubResourceData(m []*models.AAAFailureResponseSessionDetailsCause) (d []*map[string]interface{}) {
	for _, AAAFailureResponseSessionDetailsCauseModel := range m {
		if AAAFailureResponseSessionDetailsCauseModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFailureResponseSessionDetailsCauseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAAAFailureResponseSessionDetailsCausePropertyFields() (t []string) {
	return []string{}
}
