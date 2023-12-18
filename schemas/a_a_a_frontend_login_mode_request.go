package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFrontendLoginModeRequestModel(d *schema.ResourceData) *models.AAAFrontendLoginModeRequest {
	username, _ := d.Get("username").(string)
	return &models.AAAFrontendLoginModeRequest{
		Username: username,
	}
}

func AAAFrontendLoginModeRequestModelFromMap(m map[string]interface{}) *models.AAAFrontendLoginModeRequest {
	username := m["username"].(string)
	return &models.AAAFrontendLoginModeRequest{
		Username: username,
	}
}

func SetAAAFrontendLoginModeRequestResourceData(d *schema.ResourceData, m *models.AAAFrontendLoginModeRequest) {
	d.Set("username", m.Username)
}

func SetAAAFrontendLoginModeRequestSubResourceData(m []*models.AAAFrontendLoginModeRequest) (d []*map[string]interface{}) {
	for _, AAAFrontendLoginModeRequestModel := range m {
		if AAAFrontendLoginModeRequestModel != nil {
			properties := make(map[string]interface{})
			properties["username"] = AAAFrontendLoginModeRequestModel.Username
			d = append(d, &properties)
		}
	}
	return
}

func AAAFrontendLoginModeRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"username": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAAAFrontendLoginModeRequestPropertyFields() (t []string) {
	return []string{
		"username",
	}
}
