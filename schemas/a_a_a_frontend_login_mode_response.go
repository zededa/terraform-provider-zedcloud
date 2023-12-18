package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFrontendLoginModeResponseModel(d *schema.ResourceData) *models.AAAFrontendLoginModeResponse {
	var mode *models.AAAFrontendLoginModeResponseMode // AAAFrontendLoginModeResponseMode
	modeInterface, modeIsSet := d.GetOk("mode")
	if modeIsSet {
		modeModel := modeInterface.(string)
		mode = models.NewAAAFrontendLoginModeResponseMode(models.AAAFrontendLoginModeResponseMode(modeModel))
	}
	username, _ := d.Get("username").(string)
	return &models.AAAFrontendLoginModeResponse{
		Mode:     mode,
		Username: username,
	}
}

func AAAFrontendLoginModeResponseModelFromMap(m map[string]interface{}) *models.AAAFrontendLoginModeResponse {
	var mode *models.AAAFrontendLoginModeResponseMode // AAAFrontendLoginModeResponseMode
	modeInterface, modeIsSet := m["mode"]
	if modeIsSet {
		modeModel := modeInterface.(string)
		mode = models.NewAAAFrontendLoginModeResponseMode(models.AAAFrontendLoginModeResponseMode(modeModel))
	}
	username := m["username"].(string)
	return &models.AAAFrontendLoginModeResponse{
		Mode:     mode,
		Username: username,
	}
}

func SetAAAFrontendLoginModeResponseResourceData(d *schema.ResourceData, m *models.AAAFrontendLoginModeResponse) {
	d.Set("mode", m.Mode)
	d.Set("username", m.Username)
}

func SetAAAFrontendLoginModeResponseSubResourceData(m []*models.AAAFrontendLoginModeResponse) (d []*map[string]interface{}) {
	for _, AAAFrontendLoginModeResponseModel := range m {
		if AAAFrontendLoginModeResponseModel != nil {
			properties := make(map[string]interface{})
			properties["mode"] = AAAFrontendLoginModeResponseModel.Mode
			properties["username"] = AAAFrontendLoginModeResponseModel.Username
			d = append(d, &properties)
		}
	}
	return
}

func AAAFrontendLoginModeResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"mode": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"username": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAAAFrontendLoginModeResponsePropertyFields() (t []string) {
	return []string{
		"mode",
		"username",
	}
}
