package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFrontendLogoutRequestModel(d *schema.ResourceData) *models.AAAFrontendLogoutRequest {
	aAAFrontendLogoutRequest, _ := d.Get("a_a_a_frontend_logout_request").(models.AAAFrontendLogoutRequest)
	return &aAAFrontendLogoutRequest
}

func AAAFrontendLogoutRequestModelFromMap(m map[string]interface{}) *models.AAAFrontendLogoutRequest {
	aAAFrontendLogoutRequest := m["a_a_a_frontend_logout_request"].(models.AAAFrontendLogoutRequest)
	return &aAAFrontendLogoutRequest
}

func SetAAAFrontendLogoutRequestResourceData(d *schema.ResourceData, m *models.AAAFrontendLogoutRequest) {
}

func SetAAAFrontendLogoutRequestSubResourceData(m []*models.AAAFrontendLogoutRequest) (d []*map[string]interface{}) {
	for _, AAAFrontendLogoutRequestModel := range m {
		if AAAFrontendLogoutRequestModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFrontendLogoutRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAAAFrontendLogoutRequestPropertyFields() (t []string) {
	return []string{}
}
