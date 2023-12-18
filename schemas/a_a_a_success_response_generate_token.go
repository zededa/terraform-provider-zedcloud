package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAASuccessResponseGenerateTokenModel(d *schema.ResourceData) *models.AAASuccessResponseGenerateToken {
	var login *models.AAASuccessResponseLogin // AAASuccessResponseLogin
	loginInterface, loginIsSet := d.GetOk("login")
	if loginIsSet && loginInterface != nil {
		loginMap := loginInterface.([]interface{})
		if len(loginMap) > 0 {
			login = AAASuccessResponseLoginModelFromMap(loginMap[0].(map[string]interface{}))
		}
	}
	var sessionDetails *models.SessionDetails // SessionDetails
	sessionDetailsInterface, sessionDetailsIsSet := d.GetOk("session_details")
	if sessionDetailsIsSet && sessionDetailsInterface != nil {
		sessionDetailsMap := sessionDetailsInterface.([]interface{})
		if len(sessionDetailsMap) > 0 {
			sessionDetails = SessionDetailsModelFromMap(sessionDetailsMap[0].(map[string]interface{}))
		}
	}
	return &models.AAASuccessResponseGenerateToken{
		Login:          login,
		SessionDetails: sessionDetails,
	}
}

func AAASuccessResponseGenerateTokenModelFromMap(m map[string]interface{}) *models.AAASuccessResponseGenerateToken {
	var login *models.AAASuccessResponseLogin // AAASuccessResponseLogin
	loginInterface, loginIsSet := m["login"]
	if loginIsSet && loginInterface != nil {
		loginMap := loginInterface.([]interface{})
		if len(loginMap) > 0 {
			login = AAASuccessResponseLoginModelFromMap(loginMap[0].(map[string]interface{}))
		}
	}
	//
	var sessionDetails *models.SessionDetails // SessionDetails
	sessionDetailsInterface, sessionDetailsIsSet := m["session_details"]
	if sessionDetailsIsSet && sessionDetailsInterface != nil {
		sessionDetailsMap := sessionDetailsInterface.([]interface{})
		if len(sessionDetailsMap) > 0 {
			sessionDetails = SessionDetailsModelFromMap(sessionDetailsMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.AAASuccessResponseGenerateToken{
		Login:          login,
		SessionDetails: sessionDetails,
	}
}

func SetAAASuccessResponseGenerateTokenResourceData(d *schema.ResourceData, m *models.AAASuccessResponseGenerateToken) {
	d.Set("login", SetAAASuccessResponseLoginSubResourceData([]*models.AAASuccessResponseLogin{m.Login}))
	d.Set("session_details", SetSessionDetailsSubResourceData([]*models.SessionDetails{m.SessionDetails}))
}

func SetAAASuccessResponseGenerateTokenSubResourceData(m []*models.AAASuccessResponseGenerateToken) (d []*map[string]interface{}) {
	for _, AAASuccessResponseGenerateTokenModel := range m {
		if AAASuccessResponseGenerateTokenModel != nil {
			properties := make(map[string]interface{})
			properties["login"] = SetAAASuccessResponseLoginSubResourceData([]*models.AAASuccessResponseLogin{AAASuccessResponseGenerateTokenModel.Login})
			properties["session_details"] = SetSessionDetailsSubResourceData([]*models.SessionDetails{AAASuccessResponseGenerateTokenModel.SessionDetails})
			d = append(d, &properties)
		}
	}
	return
}

func AAASuccessResponseGenerateTokenSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"login": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AAASuccessResponseLogin
			Elem: &schema.Resource{
				Schema: AAASuccessResponseLoginSchema(),
			},
			Optional: true,
		},

		"session_details": {
			Description: ``,
			Type:        schema.TypeList, //GoType: SessionDetails
			Elem: &schema.Resource{
				Schema: SessionDetailsSchema(),
			},
			Optional: true,
		},
	}
}

func GetAAASuccessResponseGenerateTokenPropertyFields() (t []string) {
	return []string{
		"login",
		"session_details",
	}
}
