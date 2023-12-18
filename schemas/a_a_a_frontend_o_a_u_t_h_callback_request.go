package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFrontendOAUTHCallbackRequestModel(d *schema.ResourceData) *models.AAAFrontendOAUTHCallbackRequest {
	code, _ := d.Get("code").(string)
	state, _ := d.Get("state").(string)
	return &models.AAAFrontendOAUTHCallbackRequest{
		Code:  code,
		State: state,
	}
}

func AAAFrontendOAUTHCallbackRequestModelFromMap(m map[string]interface{}) *models.AAAFrontendOAUTHCallbackRequest {
	code := m["code"].(string)
	state := m["state"].(string)
	return &models.AAAFrontendOAUTHCallbackRequest{
		Code:  code,
		State: state,
	}
}

func SetAAAFrontendOAUTHCallbackRequestResourceData(d *schema.ResourceData, m *models.AAAFrontendOAUTHCallbackRequest) {
	d.Set("code", m.Code)
	d.Set("state", m.State)
}

func SetAAAFrontendOAUTHCallbackRequestSubResourceData(m []*models.AAAFrontendOAUTHCallbackRequest) (d []*map[string]interface{}) {
	for _, AAAFrontendOAUTHCallbackRequestModel := range m {
		if AAAFrontendOAUTHCallbackRequestModel != nil {
			properties := make(map[string]interface{})
			properties["code"] = AAAFrontendOAUTHCallbackRequestModel.Code
			properties["state"] = AAAFrontendOAUTHCallbackRequestModel.State
			d = append(d, &properties)
		}
	}
	return
}

func AAAFrontendOAUTHCallbackRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"code": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAAAFrontendOAUTHCallbackRequestPropertyFields() (t []string) {
	return []string{
		"code",
		"state",
	}
}
