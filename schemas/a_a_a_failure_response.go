package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFailureResponseModel(d *schema.ResourceData) *models.AAAFailureResponse {
	var credential *models.AAAFailureResponseCredentialChange // AAAFailureResponseCredentialChange
	credentialInterface, credentialIsSet := d.GetOk("credential")
	if credentialIsSet && credentialInterface != nil {
		credentialMap := credentialInterface.([]interface{})
		if len(credentialMap) > 0 {
			credential = AAAFailureResponseCredentialChangeModelFromMap(credentialMap[0].(map[string]interface{}))
		}
	}
	var details *models.AAAFailureResponseSessionDetails // AAAFailureResponseSessionDetails
	detailsInterface, detailsIsSet := d.GetOk("details")
	if detailsIsSet && detailsInterface != nil {
		detailsMap := detailsInterface.([]interface{})
		if len(detailsMap) > 0 {
			details = AAAFailureResponseSessionDetailsModelFromMap(detailsMap[0].(map[string]interface{}))
		}
	}
	var generateToken *models.AAAFailureResponseGenerateToken // AAAFailureResponseGenerateToken
	generateTokenInterface, generateTokenIsSet := d.GetOk("generate_token")
	if generateTokenIsSet && generateTokenInterface != nil {
		generateTokenMap := generateTokenInterface.([]interface{})
		if len(generateTokenMap) > 0 {
			generateToken = AAAFailureResponseGenerateTokenModelFromMap(generateTokenMap[0].(map[string]interface{}))
		}
	}
	var login *models.AAAFailureResponseLogin // AAAFailureResponseLogin
	loginInterface, loginIsSet := d.GetOk("login")
	if loginIsSet && loginInterface != nil {
		loginMap := loginInterface.([]interface{})
		if len(loginMap) > 0 {
			login = AAAFailureResponseLoginModelFromMap(loginMap[0].(map[string]interface{}))
		}
	}
	var logout *models.AAAFailureResponseLogout // AAAFailureResponseLogout
	logoutInterface, logoutIsSet := d.GetOk("logout")
	if logoutIsSet && logoutInterface != nil {
		logoutMap := logoutInterface.([]interface{})
		if len(logoutMap) > 0 {
			logout = AAAFailureResponseLogoutModelFromMap(logoutMap[0].(map[string]interface{}))
		}
	}
	var querySessionDetails *models.AAAFailureResponseQueryAllSessionDetails // AAAFailureResponseQueryAllSessionDetails
	querySessionDetailsInterface, querySessionDetailsIsSet := d.GetOk("query_session_details")
	if querySessionDetailsIsSet && querySessionDetailsInterface != nil {
		querySessionDetailsMap := querySessionDetailsInterface.([]interface{})
		if len(querySessionDetailsMap) > 0 {
			querySessionDetails = AAAFailureResponseQueryAllSessionDetailsModelFromMap(querySessionDetailsMap[0].(map[string]interface{}))
		}
	}
	var refresh *models.AAAFailureTokenRefresh // AAAFailureTokenRefresh
	refreshInterface, refreshIsSet := d.GetOk("refresh")
	if refreshIsSet && refreshInterface != nil {
		refreshMap := refreshInterface.([]interface{})
		if len(refreshMap) > 0 {
			refresh = AAAFailureTokenRefreshModelFromMap(refreshMap[0].(map[string]interface{}))
		}
	}
	var typeVar *models.AAAFailureResponseType // AAAFailureResponseType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewAAAFailureResponseType(models.AAAFailureResponseType(typeModel))
	}
	return &models.AAAFailureResponse{
		Credential:          credential,
		Details:             details,
		GenerateToken:       generateToken,
		Login:               login,
		Logout:              logout,
		QuerySessionDetails: querySessionDetails,
		Refresh:             refresh,
		Type:                typeVar,
	}
}

func AAAFailureResponseModelFromMap(m map[string]interface{}) *models.AAAFailureResponse {
	var credential *models.AAAFailureResponseCredentialChange // AAAFailureResponseCredentialChange
	credentialInterface, credentialIsSet := m["credential"]
	if credentialIsSet && credentialInterface != nil {
		credentialMap := credentialInterface.([]interface{})
		if len(credentialMap) > 0 {
			credential = AAAFailureResponseCredentialChangeModelFromMap(credentialMap[0].(map[string]interface{}))
		}
	}
	//
	var details *models.AAAFailureResponseSessionDetails // AAAFailureResponseSessionDetails
	detailsInterface, detailsIsSet := m["details"]
	if detailsIsSet && detailsInterface != nil {
		detailsMap := detailsInterface.([]interface{})
		if len(detailsMap) > 0 {
			details = AAAFailureResponseSessionDetailsModelFromMap(detailsMap[0].(map[string]interface{}))
		}
	}
	//
	var generateToken *models.AAAFailureResponseGenerateToken // AAAFailureResponseGenerateToken
	generateTokenInterface, generateTokenIsSet := m["generate_token"]
	if generateTokenIsSet && generateTokenInterface != nil {
		generateTokenMap := generateTokenInterface.([]interface{})
		if len(generateTokenMap) > 0 {
			generateToken = AAAFailureResponseGenerateTokenModelFromMap(generateTokenMap[0].(map[string]interface{}))
		}
	}
	//
	var login *models.AAAFailureResponseLogin // AAAFailureResponseLogin
	loginInterface, loginIsSet := m["login"]
	if loginIsSet && loginInterface != nil {
		loginMap := loginInterface.([]interface{})
		if len(loginMap) > 0 {
			login = AAAFailureResponseLoginModelFromMap(loginMap[0].(map[string]interface{}))
		}
	}
	//
	var logout *models.AAAFailureResponseLogout // AAAFailureResponseLogout
	logoutInterface, logoutIsSet := m["logout"]
	if logoutIsSet && logoutInterface != nil {
		logoutMap := logoutInterface.([]interface{})
		if len(logoutMap) > 0 {
			logout = AAAFailureResponseLogoutModelFromMap(logoutMap[0].(map[string]interface{}))
		}
	}
	//
	var querySessionDetails *models.AAAFailureResponseQueryAllSessionDetails // AAAFailureResponseQueryAllSessionDetails
	querySessionDetailsInterface, querySessionDetailsIsSet := m["query_session_details"]
	if querySessionDetailsIsSet && querySessionDetailsInterface != nil {
		querySessionDetailsMap := querySessionDetailsInterface.([]interface{})
		if len(querySessionDetailsMap) > 0 {
			querySessionDetails = AAAFailureResponseQueryAllSessionDetailsModelFromMap(querySessionDetailsMap[0].(map[string]interface{}))
		}
	}
	//
	var refresh *models.AAAFailureTokenRefresh // AAAFailureTokenRefresh
	refreshInterface, refreshIsSet := m["refresh"]
	if refreshIsSet && refreshInterface != nil {
		refreshMap := refreshInterface.([]interface{})
		if len(refreshMap) > 0 {
			refresh = AAAFailureTokenRefreshModelFromMap(refreshMap[0].(map[string]interface{}))
		}
	}
	//
	var typeVar *models.AAAFailureResponseType // AAAFailureResponseType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewAAAFailureResponseType(models.AAAFailureResponseType(typeModel))
	}
	return &models.AAAFailureResponse{
		Credential:          credential,
		Details:             details,
		GenerateToken:       generateToken,
		Login:               login,
		Logout:              logout,
		QuerySessionDetails: querySessionDetails,
		Refresh:             refresh,
		Type:                typeVar,
	}
}

func SetAAAFailureResponseResourceData(d *schema.ResourceData, m *models.AAAFailureResponse) {
	d.Set("credential", SetAAAFailureResponseCredentialChangeSubResourceData([]*models.AAAFailureResponseCredentialChange{m.Credential}))
	d.Set("details", SetAAAFailureResponseSessionDetailsSubResourceData([]*models.AAAFailureResponseSessionDetails{m.Details}))
	d.Set("generate_token", SetAAAFailureResponseGenerateTokenSubResourceData([]*models.AAAFailureResponseGenerateToken{m.GenerateToken}))
	d.Set("login", SetAAAFailureResponseLoginSubResourceData([]*models.AAAFailureResponseLogin{m.Login}))
	d.Set("logout", SetAAAFailureResponseLogoutSubResourceData([]*models.AAAFailureResponseLogout{m.Logout}))
	d.Set("query_session_details", SetAAAFailureResponseQueryAllSessionDetailsSubResourceData([]*models.AAAFailureResponseQueryAllSessionDetails{m.QuerySessionDetails}))
	d.Set("refresh", SetAAAFailureTokenRefreshSubResourceData([]*models.AAAFailureTokenRefresh{m.Refresh}))
	d.Set("type", m.Type)
}

func SetAAAFailureResponseSubResourceData(m []*models.AAAFailureResponse) (d []*map[string]interface{}) {
	for _, AAAFailureResponseModel := range m {
		if AAAFailureResponseModel != nil {
			properties := make(map[string]interface{})
			properties["credential"] = SetAAAFailureResponseCredentialChangeSubResourceData([]*models.AAAFailureResponseCredentialChange{AAAFailureResponseModel.Credential})
			properties["details"] = SetAAAFailureResponseSessionDetailsSubResourceData([]*models.AAAFailureResponseSessionDetails{AAAFailureResponseModel.Details})
			properties["generate_token"] = SetAAAFailureResponseGenerateTokenSubResourceData([]*models.AAAFailureResponseGenerateToken{AAAFailureResponseModel.GenerateToken})
			properties["login"] = SetAAAFailureResponseLoginSubResourceData([]*models.AAAFailureResponseLogin{AAAFailureResponseModel.Login})
			properties["logout"] = SetAAAFailureResponseLogoutSubResourceData([]*models.AAAFailureResponseLogout{AAAFailureResponseModel.Logout})
			properties["query_session_details"] = SetAAAFailureResponseQueryAllSessionDetailsSubResourceData([]*models.AAAFailureResponseQueryAllSessionDetails{AAAFailureResponseModel.QuerySessionDetails})
			properties["refresh"] = SetAAAFailureTokenRefreshSubResourceData([]*models.AAAFailureTokenRefresh{AAAFailureResponseModel.Refresh})
			properties["type"] = AAAFailureResponseModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func AAAFailureResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"credential": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AAAFailureResponseCredentialChange
			Elem: &schema.Resource{
				Schema: AAAFailureResponseCredentialChangeSchema(),
			},
			Optional: true,
		},

		"details": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AAAFailureResponseSessionDetails
			Elem: &schema.Resource{
				Schema: AAAFailureResponseSessionDetailsSchema(),
			},
			Optional: true,
		},

		"generate_token": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AAAFailureResponseGenerateToken
			Elem: &schema.Resource{
				Schema: AAAFailureResponseGenerateTokenSchema(),
			},
			Optional: true,
		},

		"login": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AAAFailureResponseLogin
			Elem: &schema.Resource{
				Schema: AAAFailureResponseLoginSchema(),
			},
			Optional: true,
		},

		"logout": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AAAFailureResponseLogout
			Elem: &schema.Resource{
				Schema: AAAFailureResponseLogoutSchema(),
			},
			Optional: true,
		},

		"query_session_details": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AAAFailureResponseQueryAllSessionDetails
			Elem: &schema.Resource{
				Schema: AAAFailureResponseQueryAllSessionDetailsSchema(),
			},
			Optional: true,
		},

		"refresh": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AAAFailureTokenRefresh
			Elem: &schema.Resource{
				Schema: AAAFailureTokenRefreshSchema(),
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

func GetAAAFailureResponsePropertyFields() (t []string) {
	return []string{
		"credential",
		"details",
		"generate_token",
		"login",
		"logout",
		"query_session_details",
		"refresh",
		"type",
	}
}
