package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAANotifyResponseModel(d *schema.ResourceData) *models.AAANotifyResponse {
	var login *models.AAASuccessResponseLogin // AAASuccessResponseLogin
	loginInterface, loginIsSet := d.GetOk("login")
	if loginIsSet && loginInterface != nil {
		loginMap := loginInterface.([]interface{})
		if len(loginMap) > 0 {
			login = AAASuccessResponseLoginModelFromMap(loginMap[0].(map[string]interface{}))
		}
	}
	var logout *models.AAASuccessResponseLogout // AAASuccessResponseLogout
	logoutInterface, logoutIsSet := d.GetOk("logout")
	if logoutIsSet && logoutInterface != nil {
		logoutMap := logoutInterface.([]interface{})
		if len(logoutMap) > 0 {
			logout = AAASuccessResponseLogoutModelFromMap(logoutMap[0].(map[string]interface{}))
		}
	}
	var typeVar *models.AAANotifyResponseType // AAANotifyResponseType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewAAANotifyResponseType(models.AAANotifyResponseType(typeModel))
	}
	return &models.AAANotifyResponse{
		Login:  login,
		Logout: logout,
		Type:   typeVar,
	}
}

func AAANotifyResponseModelFromMap(m map[string]interface{}) *models.AAANotifyResponse {
	var login *models.AAASuccessResponseLogin // AAASuccessResponseLogin
	loginInterface, loginIsSet := m["login"]
	if loginIsSet && loginInterface != nil {
		loginMap := loginInterface.([]interface{})
		if len(loginMap) > 0 {
			login = AAASuccessResponseLoginModelFromMap(loginMap[0].(map[string]interface{}))
		}
	}
	//
	var logout *models.AAASuccessResponseLogout // AAASuccessResponseLogout
	logoutInterface, logoutIsSet := m["logout"]
	if logoutIsSet && logoutInterface != nil {
		logoutMap := logoutInterface.([]interface{})
		if len(logoutMap) > 0 {
			logout = AAASuccessResponseLogoutModelFromMap(logoutMap[0].(map[string]interface{}))
		}
	}
	//
	var typeVar *models.AAANotifyResponseType // AAANotifyResponseType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewAAANotifyResponseType(models.AAANotifyResponseType(typeModel))
	}
	return &models.AAANotifyResponse{
		Login:  login,
		Logout: logout,
		Type:   typeVar,
	}
}

func SetAAANotifyResponseResourceData(d *schema.ResourceData, m *models.AAANotifyResponse) {
	d.Set("login", SetAAASuccessResponseLoginSubResourceData([]*models.AAASuccessResponseLogin{m.Login}))
	d.Set("logout", SetAAASuccessResponseLogoutSubResourceData([]*models.AAASuccessResponseLogout{m.Logout}))
	d.Set("type", m.Type)
}

func SetAAANotifyResponseSubResourceData(m []*models.AAANotifyResponse) (d []*map[string]interface{}) {
	for _, AAANotifyResponseModel := range m {
		if AAANotifyResponseModel != nil {
			properties := make(map[string]interface{})
			properties["login"] = SetAAASuccessResponseLoginSubResourceData([]*models.AAASuccessResponseLogin{AAANotifyResponseModel.Login})
			properties["logout"] = SetAAASuccessResponseLogoutSubResourceData([]*models.AAASuccessResponseLogout{AAANotifyResponseModel.Logout})
			properties["type"] = AAANotifyResponseModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func AAANotifyResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"login": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AAASuccessResponseLogin
			Elem: &schema.Resource{
				Schema: AAASuccessResponseLoginSchema(),
			},
			Optional: true,
		},

		"logout": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AAASuccessResponseLogout
			Elem: &schema.Resource{
				Schema: AAASuccessResponseLogoutSchema(),
			},
			Optional: true,
		},

		"type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAAANotifyResponsePropertyFields() (t []string) {
	return []string{
		"login",
		"logout",
		"type",
	}
}
