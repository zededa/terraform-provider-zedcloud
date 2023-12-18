package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAALoginModeResponseModel(d *schema.ResourceData) *models.AAALoginModeResponse {
	var mode *models.AAALoginModeResponseMode // AAALoginModeResponseMode
	modeInterface, modeIsSet := d.GetOk("mode")
	if modeIsSet {
		modeModel := modeInterface.(string)
		mode = models.NewAAALoginModeResponseMode(models.AAALoginModeResponseMode(modeModel))
	}
	username, _ := d.Get("username").(string)
	return &models.AAALoginModeResponse{
		Mode:     mode,
		Username: username,
	}
}

func AAALoginModeResponseModelFromMap(m map[string]interface{}) *models.AAALoginModeResponse {
	var mode *models.AAALoginModeResponseMode // AAALoginModeResponseMode
	modeInterface, modeIsSet := m["mode"]
	if modeIsSet {
		modeModel := modeInterface.(string)
		mode = models.NewAAALoginModeResponseMode(models.AAALoginModeResponseMode(modeModel))
	}
	username := m["username"].(string)
	return &models.AAALoginModeResponse{
		Mode:     mode,
		Username: username,
	}
}

func SetAAALoginModeResponseResourceData(d *schema.ResourceData, m *models.AAALoginModeResponse) {
	d.Set("mode", m.Mode)
	d.Set("username", m.Username)
}

func SetAAALoginModeResponseSubResourceData(m []*models.AAALoginModeResponse) (d []*map[string]interface{}) {
	for _, AAALoginModeResponseModel := range m {
		if AAALoginModeResponseModel != nil {
			properties := make(map[string]interface{})
			properties["mode"] = AAALoginModeResponseModel.Mode
			properties["username"] = AAALoginModeResponseModel.Username
			d = append(d, &properties)
		}
	}
	return
}

func AAALoginModeResponseSchema() map[string]*schema.Schema {
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

func GetAAALoginModeResponsePropertyFields() (t []string) {
	return []string{
		"mode",
		"username",
	}
}
