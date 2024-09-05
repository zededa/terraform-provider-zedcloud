package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AAAFrontendLoginResponseModel(d *schema.ResourceData) *models.AAAFrontendLoginResponse {
	var aPIToken *models.Token64 // Token64
	apiTokenInterface, apiTokenIsSet := d.GetOk("api_token")
	if apiTokenIsSet && apiTokenInterface != nil {
		apiTokenMap := apiTokenInterface.([]interface{})
		if len(apiTokenMap) > 0 {
			aPIToken = Token64ModelFromMap(apiTokenMap[0].(map[string]interface{}))
		}
	}
	var cause *models.AAAFrontendLoginResponseCause // AAAFrontendLoginResponseCause
	causeInterface, causeIsSet := d.GetOk("cause")
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFrontendLoginResponseCause(models.AAAFrontendLoginResponseCause(causeModel))
	}
	var detailedUser *models.DetailedUser // DetailedUser
	detailedUserInterface, detailedUserIsSet := d.GetOk("detailed_user")
	if detailedUserIsSet && detailedUserInterface != nil {
		detailedUserMap := detailedUserInterface.([]interface{})
		if len(detailedUserMap) > 0 {
			detailedUser = DetailedUserModelFromMap(detailedUserMap[0].(map[string]interface{}))
		}
	}
	var enterprise *models.Enterprise // Enterprise
	enterpriseInterface, enterpriseIsSet := d.GetOk("enterprise")
	if enterpriseIsSet && enterpriseInterface != nil {
		enterpriseMap := enterpriseInterface.([]interface{})
		if len(enterpriseMap) > 0 {
			enterprise = EnterpriseModelFromMap(enterpriseMap[0].(map[string]interface{}))
		}
	}
	var loginToken *models.Token64 // Token64
	loginTokenInterface, loginTokenIsSet := d.GetOk("login_token")
	if loginTokenIsSet && loginTokenInterface != nil {
		loginTokenMap := loginTokenInterface.([]interface{})
		if len(loginTokenMap) > 0 {
			loginToken = Token64ModelFromMap(loginTokenMap[0].(map[string]interface{}))
		}
	}
	noOfLoginAttemptsLeftInt, _ := d.Get("no_of_login_attempts_left").(int)
	noOfLoginAttemptsLeft := int64(noOfLoginAttemptsLeftInt)
	passwordExpiryNotificationPeriodInSecondsInt, _ := d.Get("password_expiry_notification_period_in_seconds").(int)
	passwordExpiryNotificationPeriodInSeconds := int64(passwordExpiryNotificationPeriodInSecondsInt)
	passwordExpiryTime, _ := d.Get("password_expiry_time").(strfmt.DateTime)
	var policies []*models.Policy // []*Policy
	policiesInterface, policiesIsSet := d.GetOk("policies")
	if policiesIsSet {
		var items []interface{}
		if listItems, isList := policiesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = policiesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := PolicyConfigModelFromMap(v.(map[string]interface{}))
			policies = append(policies, m)
		}
	}
	var realm *models.Realm // Realm
	realmInterface, realmIsSet := d.GetOk("realm")
	if realmIsSet && realmInterface != nil {
		realmMap := realmInterface.([]interface{})
		if len(realmMap) > 0 {
			realm = RealmModelFromMap(realmMap[0].(map[string]interface{}))
		}
	}
	redirectURL, _ := d.Get("redirect_url").(string)
	var role *models.Role // Role
	roleInterface, roleIsSet := d.GetOk("role")
	if roleIsSet && roleInterface != nil {
		roleMap := roleInterface.([]interface{})
		if len(roleMap) > 0 {
			role = RoleModelFromMap(roleMap[0].(map[string]interface{}))
		}
	}
	var simpleUser *models.SimpleUser // SimpleUser
	simpleUserInterface, simpleUserIsSet := d.GetOk("simple_user")
	if simpleUserIsSet && simpleUserInterface != nil {
		simpleUserMap := simpleUserInterface.([]interface{})
		if len(simpleUserMap) > 0 {
			simpleUser = SimpleUserModelFromMap(simpleUserMap[0].(map[string]interface{}))
		}
	}
	var tempToken *models.Token64 // Token64
	tempTokenInterface, tempTokenIsSet := d.GetOk("temp_token")
	if tempTokenIsSet && tempTokenInterface != nil {
		tempTokenMap := tempTokenInterface.([]interface{})
		if len(tempTokenMap) > 0 {
			tempToken = Token64ModelFromMap(tempTokenMap[0].(map[string]interface{}))
		}
	}
	var token *models.Token64 // Token64
	tokenInterface, tokenIsSet := d.GetOk("token")
	if tokenIsSet && tokenInterface != nil {
		tokenMap := tokenInterface.([]interface{})
		if len(tokenMap) > 0 {
			token = Token64ModelFromMap(tokenMap[0].(map[string]interface{}))
		}
	}
	userID, _ := d.Get("user_id").(string)
	return &models.AAAFrontendLoginResponse{
		APIToken:              aPIToken,
		Cause:                 cause,
		DetailedUser:          detailedUser,
		Enterprise:            enterprise,
		LoginToken:            loginToken,
		NoOfLoginAttemptsLeft: noOfLoginAttemptsLeft,
		PasswordExpiryNotificationPeriodInSeconds: passwordExpiryNotificationPeriodInSeconds,
		PasswordExpiryTime:                        passwordExpiryTime,
		Policies:                                  policies,
		Realm:                                     realm,
		RedirectURL:                               redirectURL,
		Role:                                      role,
		SimpleUser:                                simpleUser,
		TempToken:                                 tempToken,
		Token:                                     token,
		UserID:                                    userID,
	}
}

func AAAFrontendLoginResponseModelFromMap(m map[string]interface{}) *models.AAAFrontendLoginResponse {
	var aPIToken *models.Token64 // Token64
	apiTokenInterface, apiTokenIsSet := m["api_token"]
	if apiTokenIsSet && apiTokenInterface != nil {
		apiTokenMap := apiTokenInterface.([]interface{})
		if len(apiTokenMap) > 0 {
			aPIToken = Token64ModelFromMap(apiTokenMap[0].(map[string]interface{}))
		}
	}
	//
	var cause *models.AAAFrontendLoginResponseCause // AAAFrontendLoginResponseCause
	causeInterface, causeIsSet := m["cause"]
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFrontendLoginResponseCause(models.AAAFrontendLoginResponseCause(causeModel))
	}
	var detailedUser *models.DetailedUser // DetailedUser
	detailedUserInterface, detailedUserIsSet := m["detailed_user"]
	if detailedUserIsSet && detailedUserInterface != nil {
		detailedUserMap := detailedUserInterface.([]interface{})
		if len(detailedUserMap) > 0 {
			detailedUser = DetailedUserModelFromMap(detailedUserMap[0].(map[string]interface{}))
		}
	}
	//
	var enterprise *models.Enterprise // Enterprise
	enterpriseInterface, enterpriseIsSet := m["enterprise"]
	if enterpriseIsSet && enterpriseInterface != nil {
		enterpriseMap := enterpriseInterface.([]interface{})
		if len(enterpriseMap) > 0 {
			enterprise = EnterpriseModelFromMap(enterpriseMap[0].(map[string]interface{}))
		}
	}
	//
	var loginToken *models.Token64 // Token64
	loginTokenInterface, loginTokenIsSet := m["login_token"]
	if loginTokenIsSet && loginTokenInterface != nil {
		loginTokenMap := loginTokenInterface.([]interface{})
		if len(loginTokenMap) > 0 {
			loginToken = Token64ModelFromMap(loginTokenMap[0].(map[string]interface{}))
		}
	}
	//
	noOfLoginAttemptsLeft := int64(m["no_of_login_attempts_left"].(int))                                          // int64
	passwordExpiryNotificationPeriodInSeconds := int64(m["password_expiry_notification_period_in_seconds"].(int)) // int64
	passwordExpiryTime := m["password_expiry_time"].(strfmt.DateTime)
	var policies []*models.Policy // []*Policy
	policiesInterface, policiesIsSet := m["policies"]
	if policiesIsSet {
		var items []interface{}
		if listItems, isList := policiesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = policiesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := PolicyConfigModelFromMap(v.(map[string]interface{}))
			policies = append(policies, m)
		}
	}
	var realm *models.Realm // Realm
	realmInterface, realmIsSet := m["realm"]
	if realmIsSet && realmInterface != nil {
		realmMap := realmInterface.([]interface{})
		if len(realmMap) > 0 {
			realm = RealmModelFromMap(realmMap[0].(map[string]interface{}))
		}
	}
	//
	redirectURL := m["redirect_url"].(string)
	var role *models.Role // Role
	roleInterface, roleIsSet := m["role"]
	if roleIsSet && roleInterface != nil {
		roleMap := roleInterface.([]interface{})
		if len(roleMap) > 0 {
			role = RoleModelFromMap(roleMap[0].(map[string]interface{}))
		}
	}
	//
	var simpleUser *models.SimpleUser // SimpleUser
	simpleUserInterface, simpleUserIsSet := m["simple_user"]
	if simpleUserIsSet && simpleUserInterface != nil {
		simpleUserMap := simpleUserInterface.([]interface{})
		if len(simpleUserMap) > 0 {
			simpleUser = SimpleUserModelFromMap(simpleUserMap[0].(map[string]interface{}))
		}
	}
	//
	var tempToken *models.Token64 // Token64
	tempTokenInterface, tempTokenIsSet := m["temp_token"]
	if tempTokenIsSet && tempTokenInterface != nil {
		tempTokenMap := tempTokenInterface.([]interface{})
		if len(tempTokenMap) > 0 {
			tempToken = Token64ModelFromMap(tempTokenMap[0].(map[string]interface{}))
		}
	}
	//
	var token *models.Token64 // Token64
	tokenInterface, tokenIsSet := m["token"]
	if tokenIsSet && tokenInterface != nil {
		tokenMap := tokenInterface.([]interface{})
		if len(tokenMap) > 0 {
			token = Token64ModelFromMap(tokenMap[0].(map[string]interface{}))
		}
	}
	//
	userID := m["user_id"].(string)
	return &models.AAAFrontendLoginResponse{
		APIToken:              aPIToken,
		Cause:                 cause,
		DetailedUser:          detailedUser,
		Enterprise:            enterprise,
		LoginToken:            loginToken,
		NoOfLoginAttemptsLeft: noOfLoginAttemptsLeft,
		PasswordExpiryNotificationPeriodInSeconds: passwordExpiryNotificationPeriodInSeconds,
		PasswordExpiryTime:                        passwordExpiryTime,
		Policies:                                  policies,
		Realm:                                     realm,
		RedirectURL:                               redirectURL,
		Role:                                      role,
		SimpleUser:                                simpleUser,
		TempToken:                                 tempToken,
		Token:                                     token,
		UserID:                                    userID,
	}
}

func SetAAAFrontendLoginResponseResourceData(d *schema.ResourceData, m *models.AAAFrontendLoginResponse) {
	d.Set("api_token", SetToken64SubResourceData([]*models.Token64{m.APIToken}))
	d.Set("cause", m.Cause)
	d.Set("detailed_user", SetDetailedUserSubResourceData([]*models.DetailedUser{m.DetailedUser}))
	d.Set("enterprise", SetEnterpriseSubResourceData([]*models.Enterprise{m.Enterprise}))
	d.Set("login_token", SetToken64SubResourceData([]*models.Token64{m.LoginToken}))
	d.Set("no_of_login_attempts_left", m.NoOfLoginAttemptsLeft)
	d.Set("password_expiry_notification_period_in_seconds", m.PasswordExpiryNotificationPeriodInSeconds)
	d.Set("password_expiry_time", m.PasswordExpiryTime)
	d.Set("policies", SetPolicyConfigSubResourceData(m.Policies))
	d.Set("realm", SetRealmSubResourceData([]*models.Realm{m.Realm}))
	d.Set("redirect_url", m.RedirectURL)
	d.Set("role", SetRoleSubResourceData([]*models.Role{m.Role}))
	d.Set("simple_user", SetSimpleUserSubResourceData([]*models.SimpleUser{m.SimpleUser}))
	d.Set("temp_token", SetToken64SubResourceData([]*models.Token64{m.TempToken}))
	d.Set("token", SetToken64SubResourceData([]*models.Token64{m.Token}))
	d.Set("user_id", m.UserID)
}

func SetAAAFrontendLoginResponseSubResourceData(m []*models.AAAFrontendLoginResponse) (d []*map[string]interface{}) {
	for _, AAAFrontendLoginResponseModel := range m {
		if AAAFrontendLoginResponseModel != nil {
			properties := make(map[string]interface{})
			properties["api_token"] = SetToken64SubResourceData([]*models.Token64{AAAFrontendLoginResponseModel.APIToken})
			properties["cause"] = AAAFrontendLoginResponseModel.Cause
			properties["detailed_user"] = SetDetailedUserSubResourceData([]*models.DetailedUser{AAAFrontendLoginResponseModel.DetailedUser})
			properties["enterprise"] = SetEnterpriseSubResourceData([]*models.Enterprise{AAAFrontendLoginResponseModel.Enterprise})
			properties["login_token"] = SetToken64SubResourceData([]*models.Token64{AAAFrontendLoginResponseModel.LoginToken})
			properties["no_of_login_attempts_left"] = AAAFrontendLoginResponseModel.NoOfLoginAttemptsLeft
			properties["password_expiry_notification_period_in_seconds"] = AAAFrontendLoginResponseModel.PasswordExpiryNotificationPeriodInSeconds
			properties["password_expiry_time"] = AAAFrontendLoginResponseModel.PasswordExpiryTime.String()
			properties["policies"] = SetPolicyConfigSubResourceData(AAAFrontendLoginResponseModel.Policies)
			properties["realm"] = SetRealmSubResourceData([]*models.Realm{AAAFrontendLoginResponseModel.Realm})
			properties["redirect_url"] = AAAFrontendLoginResponseModel.RedirectURL
			properties["role"] = SetRoleSubResourceData([]*models.Role{AAAFrontendLoginResponseModel.Role})
			properties["simple_user"] = SetSimpleUserSubResourceData([]*models.SimpleUser{AAAFrontendLoginResponseModel.SimpleUser})
			properties["temp_token"] = SetToken64SubResourceData([]*models.Token64{AAAFrontendLoginResponseModel.TempToken})
			properties["token"] = SetToken64SubResourceData([]*models.Token64{AAAFrontendLoginResponseModel.Token})
			properties["user_id"] = AAAFrontendLoginResponseModel.UserID
			d = append(d, &properties)
		}
	}
	return
}

func AAAFrontendLoginResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"api_token": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Token64
			Elem: &schema.Resource{
				Schema: Token64Schema(),
			},
			Optional: true,
		},

		"cause": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"detailed_user": {
			Description: ``,
			Type:        schema.TypeList, //GoType: DetailedUser
			Elem: &schema.Resource{
				Schema: DetailedUserSchema(),
			},
			Optional: true,
		},

		"enterprise": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Enterprise
			Elem: &schema.Resource{
				Schema: EnterpriseSchema(),
			},
			Optional: true,
		},

		"login_token": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Token64
			Elem: &schema.Resource{
				Schema: Token64Schema(),
			},
			Optional: true,
		},

		"no_of_login_attempts_left": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"password_expiry_notification_period_in_seconds": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"password_expiry_time": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"policies": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*Policy
			Elem: &schema.Resource{
				Schema: Policy(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"realm": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Realm
			Elem: &schema.Resource{
				Schema: RealmSchema(),
			},
			Optional: true,
		},

		"redirect_url": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"role": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Role
			Elem: &schema.Resource{
				Schema: RoleSchema(),
			},
			Optional: true,
		},

		"simple_user": {
			Description: ``,
			Type:        schema.TypeList, //GoType: SimpleUser
			Elem: &schema.Resource{
				Schema: SimpleUserSchema(),
			},
			Optional: true,
		},

		"temp_token": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Token64
			Elem: &schema.Resource{
				Schema: Token64Schema(),
			},
			Optional: true,
		},

		"token": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Token64
			Elem: &schema.Resource{
				Schema: Token64Schema(),
			},
			Optional: true,
		},

		"user_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAAAFrontendLoginResponsePropertyFields() (t []string) {
	return []string{
		"api_token",
		"cause",
		"detailed_user",
		"enterprise",
		"login_token",
		"no_of_login_attempts_left",
		"password_expiry_notification_period_in_seconds",
		"password_expiry_time",
		"policies",
		"realm",
		"redirect_url",
		"role",
		"simple_user",
		"temp_token",
		"token",
		"user_id",
	}
}
