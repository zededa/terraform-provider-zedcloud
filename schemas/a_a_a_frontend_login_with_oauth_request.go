package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFrontendLoginWithOauthRequestModel(d *schema.ResourceData) *models.AAAFrontendLoginWithOauthRequest {
	enterpriseName, _ := d.Get("enterprise_name").(string)
	usernameAtRealm, _ := d.Get("username_at_realm").(string)
	return &models.AAAFrontendLoginWithOauthRequest{
		EnterpriseName:  enterpriseName,
		UsernameAtRealm: usernameAtRealm,
	}
}

func AAAFrontendLoginWithOauthRequestModelFromMap(m map[string]interface{}) *models.AAAFrontendLoginWithOauthRequest {
	enterpriseName := m["enterprise_name"].(string)
	usernameAtRealm := m["username_at_realm"].(string)
	return &models.AAAFrontendLoginWithOauthRequest{
		EnterpriseName:  enterpriseName,
		UsernameAtRealm: usernameAtRealm,
	}
}

func SetAAAFrontendLoginWithOauthRequestResourceData(d *schema.ResourceData, m *models.AAAFrontendLoginWithOauthRequest) {
	d.Set("enterprise_name", m.EnterpriseName)
	d.Set("username_at_realm", m.UsernameAtRealm)
}

func SetAAAFrontendLoginWithOauthRequestSubResourceData(m []*models.AAAFrontendLoginWithOauthRequest) (d []*map[string]interface{}) {
	for _, AAAFrontendLoginWithOauthRequestModel := range m {
		if AAAFrontendLoginWithOauthRequestModel != nil {
			properties := make(map[string]interface{})
			properties["enterprise_name"] = AAAFrontendLoginWithOauthRequestModel.EnterpriseName
			properties["username_at_realm"] = AAAFrontendLoginWithOauthRequestModel.UsernameAtRealm
			d = append(d, &properties)
		}
	}
	return
}

func AAAFrontendLoginWithOauthRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enterprise_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"username_at_realm": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAAAFrontendLoginWithOauthRequestPropertyFields() (t []string) {
	return []string{
		"enterprise_name",
		"username_at_realm",
	}
}
