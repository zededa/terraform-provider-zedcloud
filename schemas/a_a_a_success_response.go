package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAASuccessResponseModel(d *schema.ResourceData) *models.AAASuccessResponse {
	var credential *models.AAASuccessResponseCredentialChange // AAASuccessResponseCredentialChange
	credentialInterface, credentialIsSet := d.GetOk("credential")
	if credentialIsSet && credentialInterface != nil {
		credentialMap := credentialInterface.([]interface{})
		if len(credentialMap) > 0 {
			credential = AAASuccessResponseCredentialChangeModelFromMap(credentialMap[0].(map[string]interface{}))
		}
	}
	var details *models.AAASuccessSessionDetailsResponse // AAASuccessSessionDetailsResponse
	detailsInterface, detailsIsSet := d.GetOk("details")
	if detailsIsSet && detailsInterface != nil {
		detailsMap := detailsInterface.([]interface{})
		if len(detailsMap) > 0 {
			details = AAASuccessSessionDetailsResponseModelFromMap(detailsMap[0].(map[string]interface{}))
		}
	}
	var generateToken *models.AAASuccessResponseGenerateToken // AAASuccessResponseGenerateToken
	generateTokenInterface, generateTokenIsSet := d.GetOk("generate_token")
	if generateTokenIsSet && generateTokenInterface != nil {
		generateTokenMap := generateTokenInterface.([]interface{})
		if len(generateTokenMap) > 0 {
			generateToken = AAASuccessResponseGenerateTokenModelFromMap(generateTokenMap[0].(map[string]interface{}))
		}
	}
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
	var querySessionDetails *models.AAASuccessResponseQueryAllSessionDetails // AAASuccessResponseQueryAllSessionDetails
	querySessionDetailsInterface, querySessionDetailsIsSet := d.GetOk("query_session_details")
	if querySessionDetailsIsSet && querySessionDetailsInterface != nil {
		querySessionDetailsMap := querySessionDetailsInterface.([]interface{})
		if len(querySessionDetailsMap) > 0 {
			querySessionDetails = AAASuccessResponseQueryAllSessionDetailsModelFromMap(querySessionDetailsMap[0].(map[string]interface{}))
		}
	}
	var refresh *models.AAASuccessTokenRefresh // AAASuccessTokenRefresh
	refreshInterface, refreshIsSet := d.GetOk("refresh")
	if refreshIsSet && refreshInterface != nil {
		refreshMap := refreshInterface.([]interface{})
		if len(refreshMap) > 0 {
			refresh = AAASuccessTokenRefreshModelFromMap(refreshMap[0].(map[string]interface{}))
		}
	}
	var typeVar *models.AAASuccessResponseType // AAASuccessResponseType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewAAASuccessResponseType(models.AAASuccessResponseType(typeModel))
	}
	return &models.AAASuccessResponse{
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

func AAASuccessResponseModelFromMap(m map[string]interface{}) *models.AAASuccessResponse {
	var credential *models.AAASuccessResponseCredentialChange // AAASuccessResponseCredentialChange
	credentialInterface, credentialIsSet := m["credential"]
	if credentialIsSet && credentialInterface != nil {
		credentialMap := credentialInterface.([]interface{})
		if len(credentialMap) > 0 {
			credential = AAASuccessResponseCredentialChangeModelFromMap(credentialMap[0].(map[string]interface{}))
		}
	}
	//
	var details *models.AAASuccessSessionDetailsResponse // AAASuccessSessionDetailsResponse
	detailsInterface, detailsIsSet := m["details"]
	if detailsIsSet && detailsInterface != nil {
		detailsMap := detailsInterface.([]interface{})
		if len(detailsMap) > 0 {
			details = AAASuccessSessionDetailsResponseModelFromMap(detailsMap[0].(map[string]interface{}))
		}
	}
	//
	var generateToken *models.AAASuccessResponseGenerateToken // AAASuccessResponseGenerateToken
	generateTokenInterface, generateTokenIsSet := m["generate_token"]
	if generateTokenIsSet && generateTokenInterface != nil {
		generateTokenMap := generateTokenInterface.([]interface{})
		if len(generateTokenMap) > 0 {
			generateToken = AAASuccessResponseGenerateTokenModelFromMap(generateTokenMap[0].(map[string]interface{}))
		}
	}
	//
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
	var querySessionDetails *models.AAASuccessResponseQueryAllSessionDetails // AAASuccessResponseQueryAllSessionDetails
	querySessionDetailsInterface, querySessionDetailsIsSet := m["query_session_details"]
	if querySessionDetailsIsSet && querySessionDetailsInterface != nil {
		querySessionDetailsMap := querySessionDetailsInterface.([]interface{})
		if len(querySessionDetailsMap) > 0 {
			querySessionDetails = AAASuccessResponseQueryAllSessionDetailsModelFromMap(querySessionDetailsMap[0].(map[string]interface{}))
		}
	}
	//
	var refresh *models.AAASuccessTokenRefresh // AAASuccessTokenRefresh
	refreshInterface, refreshIsSet := m["refresh"]
	if refreshIsSet && refreshInterface != nil {
		refreshMap := refreshInterface.([]interface{})
		if len(refreshMap) > 0 {
			refresh = AAASuccessTokenRefreshModelFromMap(refreshMap[0].(map[string]interface{}))
		}
	}
	//
	var typeVar *models.AAASuccessResponseType // AAASuccessResponseType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewAAASuccessResponseType(models.AAASuccessResponseType(typeModel))
	}
	return &models.AAASuccessResponse{
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

func SetAAASuccessResponseResourceData(d *schema.ResourceData, m *models.AAASuccessResponse) {
	d.Set("credential", SetAAASuccessResponseCredentialChangeSubResourceData([]*models.AAASuccessResponseCredentialChange{m.Credential}))
	d.Set("details", SetAAASuccessSessionDetailsResponseSubResourceData([]*models.AAASuccessSessionDetailsResponse{m.Details}))
	d.Set("generate_token", SetAAASuccessResponseGenerateTokenSubResourceData([]*models.AAASuccessResponseGenerateToken{m.GenerateToken}))
	d.Set("login", SetAAASuccessResponseLoginSubResourceData([]*models.AAASuccessResponseLogin{m.Login}))
	d.Set("logout", SetAAASuccessResponseLogoutSubResourceData([]*models.AAASuccessResponseLogout{m.Logout}))
	d.Set("query_session_details", SetAAASuccessResponseQueryAllSessionDetailsSubResourceData([]*models.AAASuccessResponseQueryAllSessionDetails{m.QuerySessionDetails}))
	d.Set("refresh", SetAAASuccessTokenRefreshSubResourceData([]*models.AAASuccessTokenRefresh{m.Refresh}))
	d.Set("type", m.Type)
}

func SetAAASuccessResponseSubResourceData(m []*models.AAASuccessResponse) (d []*map[string]interface{}) {
	for _, AAASuccessResponseModel := range m {
		if AAASuccessResponseModel != nil {
			properties := make(map[string]interface{})
			properties["credential"] = SetAAASuccessResponseCredentialChangeSubResourceData([]*models.AAASuccessResponseCredentialChange{AAASuccessResponseModel.Credential})
			properties["details"] = SetAAASuccessSessionDetailsResponseSubResourceData([]*models.AAASuccessSessionDetailsResponse{AAASuccessResponseModel.Details})
			properties["generate_token"] = SetAAASuccessResponseGenerateTokenSubResourceData([]*models.AAASuccessResponseGenerateToken{AAASuccessResponseModel.GenerateToken})
			properties["login"] = SetAAASuccessResponseLoginSubResourceData([]*models.AAASuccessResponseLogin{AAASuccessResponseModel.Login})
			properties["logout"] = SetAAASuccessResponseLogoutSubResourceData([]*models.AAASuccessResponseLogout{AAASuccessResponseModel.Logout})
			properties["query_session_details"] = SetAAASuccessResponseQueryAllSessionDetailsSubResourceData([]*models.AAASuccessResponseQueryAllSessionDetails{AAASuccessResponseModel.QuerySessionDetails})
			properties["refresh"] = SetAAASuccessTokenRefreshSubResourceData([]*models.AAASuccessTokenRefresh{AAASuccessResponseModel.Refresh})
			properties["type"] = AAASuccessResponseModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func AAASuccessResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"credential": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AAASuccessResponseCredentialChange
			Elem: &schema.Resource{
				Schema: AAASuccessResponseCredentialChangeSchema(),
			},
			Optional: true,
		},

		"details": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AAASuccessSessionDetailsResponse
			Elem: &schema.Resource{
				Schema: AAASuccessSessionDetailsResponseSchema(),
			},
			Optional: true,
		},

		"generate_token": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AAASuccessResponseGenerateToken
			Elem: &schema.Resource{
				Schema: AAASuccessResponseGenerateTokenSchema(),
			},
			Optional: true,
		},

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

		"query_session_details": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AAASuccessResponseQueryAllSessionDetails
			Elem: &schema.Resource{
				Schema: AAASuccessResponseQueryAllSessionDetailsSchema(),
			},
			Optional: true,
		},

		"refresh": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AAASuccessTokenRefresh
			Elem: &schema.Resource{
				Schema: AAASuccessTokenRefreshSchema(),
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

func GetAAASuccessResponsePropertyFields() (t []string) {
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
