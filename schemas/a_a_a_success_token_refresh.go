package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAASuccessTokenRefreshModel(d *schema.ResourceData) *models.AAASuccessTokenRefresh {
	var token *models.Token64 // Token64
	tokenInterface, tokenIsSet := d.GetOk("token")
	if tokenIsSet && tokenInterface != nil {
		tokenMap := tokenInterface.([]interface{})
		if len(tokenMap) > 0 {
			token = Token64ModelFromMap(tokenMap[0].(map[string]interface{}))
		}
	}
	return &models.AAASuccessTokenRefresh{
		Token: token,
	}
}

func AAASuccessTokenRefreshModelFromMap(m map[string]interface{}) *models.AAASuccessTokenRefresh {
	var token *models.Token64 // Token64
	tokenInterface, tokenIsSet := m["token"]
	if tokenIsSet && tokenInterface != nil {
		tokenMap := tokenInterface.([]interface{})
		if len(tokenMap) > 0 {
			token = Token64ModelFromMap(tokenMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.AAASuccessTokenRefresh{
		Token: token,
	}
}

func SetAAASuccessTokenRefreshResourceData(d *schema.ResourceData, m *models.AAASuccessTokenRefresh) {
	d.Set("token", SetToken64SubResourceData([]*models.Token64{m.Token}))
}

func SetAAASuccessTokenRefreshSubResourceData(m []*models.AAASuccessTokenRefresh) (d []*map[string]interface{}) {
	for _, AAASuccessTokenRefreshModel := range m {
		if AAASuccessTokenRefreshModel != nil {
			properties := make(map[string]interface{})
			properties["token"] = SetToken64SubResourceData([]*models.Token64{AAASuccessTokenRefreshModel.Token})
			d = append(d, &properties)
		}
	}
	return
}

func AAASuccessTokenRefreshSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"token": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Token64
			Elem: &schema.Resource{
				Schema: Token64Schema(),
			},
			Optional: true,
		},
	}
}

func GetAAASuccessTokenRefreshPropertyFields() (t []string) {
	return []string{
		"token",
	}
}
