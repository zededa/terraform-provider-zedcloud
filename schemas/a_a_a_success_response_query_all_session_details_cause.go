package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAASuccessResponseQueryAllSessionDetailsCauseModel(d *schema.ResourceData) *models.AAASuccessResponseQueryAllSessionDetailsCause {
	aAASuccessResponseQueryAllSessionDetailsCause, _ := d.Get("a_a_a_success_response_query_all_session_details_cause").(models.AAASuccessResponseQueryAllSessionDetailsCause)
	return &aAASuccessResponseQueryAllSessionDetailsCause
}

func AAASuccessResponseQueryAllSessionDetailsCauseModelFromMap(m map[string]interface{}) *models.AAASuccessResponseQueryAllSessionDetailsCause {
	aAASuccessResponseQueryAllSessionDetailsCause := m["a_a_a_success_response_query_all_session_details_cause"].(models.AAASuccessResponseQueryAllSessionDetailsCause)
	return &aAASuccessResponseQueryAllSessionDetailsCause
}

func SetAAASuccessResponseQueryAllSessionDetailsCauseResourceData(d *schema.ResourceData, m *models.AAASuccessResponseQueryAllSessionDetailsCause) {
}

func SetAAASuccessResponseQueryAllSessionDetailsCauseSubResourceData(m []*models.AAASuccessResponseQueryAllSessionDetailsCause) (d []*map[string]interface{}) {
	for _, AAASuccessResponseQueryAllSessionDetailsCauseModel := range m {
		if AAASuccessResponseQueryAllSessionDetailsCauseModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AAASuccessResponseQueryAllSessionDetailsCauseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAAASuccessResponseQueryAllSessionDetailsCausePropertyFields() (t []string) {
	return []string{}
}
