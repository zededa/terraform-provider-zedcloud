package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFailureResponseQueryAllSessionDetailsCauseModel(d *schema.ResourceData) *models.AAAFailureResponseQueryAllSessionDetailsCause {
	aAAFailureResponseQueryAllSessionDetailsCause, _ := d.Get("a_a_a_failure_response_query_all_session_details_cause").(models.AAAFailureResponseQueryAllSessionDetailsCause)
	return &aAAFailureResponseQueryAllSessionDetailsCause
}

func AAAFailureResponseQueryAllSessionDetailsCauseModelFromMap(m map[string]interface{}) *models.AAAFailureResponseQueryAllSessionDetailsCause {
	aAAFailureResponseQueryAllSessionDetailsCause := m["a_a_a_failure_response_query_all_session_details_cause"].(models.AAAFailureResponseQueryAllSessionDetailsCause)
	return &aAAFailureResponseQueryAllSessionDetailsCause
}

func SetAAAFailureResponseQueryAllSessionDetailsCauseResourceData(d *schema.ResourceData, m *models.AAAFailureResponseQueryAllSessionDetailsCause) {
}

func SetAAAFailureResponseQueryAllSessionDetailsCauseSubResourceData(m []*models.AAAFailureResponseQueryAllSessionDetailsCause) (d []*map[string]interface{}) {
	for _, AAAFailureResponseQueryAllSessionDetailsCauseModel := range m {
		if AAAFailureResponseQueryAllSessionDetailsCauseModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFailureResponseQueryAllSessionDetailsCauseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAAAFailureResponseQueryAllSessionDetailsCausePropertyFields() (t []string) {
	return []string{}
}
