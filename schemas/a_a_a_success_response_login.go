package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider/models"
)

func AAASuccessResponseLoginModel(d *schema.ResourceData) *models.AAASuccessResponseLogin {
	var aPIToken *models.Token64 // Token64
	apiTokenInterface, apiTokenIsSet := d.GetOk("api_token")
	if apiTokenIsSet && apiTokenInterface != nil {
		apiTokenMap := apiTokenInterface.([]interface{})
		if len(apiTokenMap) > 0 {
			aPIToken = Token64ModelFromMap(apiTokenMap[0].(map[string]interface{}))
		}
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
	var enterpriseID *models.Identifier64 // Identifier64
	enterpriseIdInterface, enterpriseIdIsSet := d.GetOk("enterprise_id")
	if enterpriseIdIsSet && enterpriseIdInterface != nil {
		enterpriseIdMap := enterpriseIdInterface.([]interface{})
		if len(enterpriseIdMap) > 0 {
			enterpriseID = Identifier64ModelFromMap(enterpriseIdMap[0].(map[string]interface{}))
		}
	}
	isJWTValid, _ := d.Get("is_j_w_t_valid").(bool)
	var loginToken *models.Token64 // Token64
	loginTokenInterface, loginTokenIsSet := d.GetOk("login_token")
	if loginTokenIsSet && loginTokenInterface != nil {
		loginTokenMap := loginTokenInterface.([]interface{})
		if len(loginTokenMap) > 0 {
			loginToken = Token64ModelFromMap(loginTokenMap[0].(map[string]interface{}))
		}
	}
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
	var realmID *models.Identifier64 // Identifier64
	realmIdInterface, realmIdIsSet := d.GetOk("realm_id")
	if realmIdIsSet && realmIdInterface != nil {
		realmIdMap := realmIdInterface.([]interface{})
		if len(realmIdMap) > 0 {
			realmID = Identifier64ModelFromMap(realmIdMap[0].(map[string]interface{}))
		}
	}
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
	var token *models.Token64 // Token64
	tokenInterface, tokenIsSet := d.GetOk("token")
	if tokenIsSet && tokenInterface != nil {
		tokenMap := tokenInterface.([]interface{})
		if len(tokenMap) > 0 {
			token = Token64ModelFromMap(tokenMap[0].(map[string]interface{}))
		}
	}
	var userID *models.Identifier64 // Identifier64
	userIdInterface, userIdIsSet := d.GetOk("user_id")
	if userIdIsSet && userIdInterface != nil {
		userIdMap := userIdInterface.([]interface{})
		if len(userIdMap) > 0 {
			userID = Identifier64ModelFromMap(userIdMap[0].(map[string]interface{}))
		}
	}
	return &models.AAASuccessResponseLogin{
		APIToken:     aPIToken,
		DetailedUser: detailedUser,
		Enterprise:   enterprise,
		EnterpriseID: enterpriseID,
		IsJWTValid:   isJWTValid,
		LoginToken:   loginToken,
		PasswordExpiryNotificationPeriodInSeconds: passwordExpiryNotificationPeriodInSeconds,
		PasswordExpiryTime:                        passwordExpiryTime,
		Policies:                                  policies,
		Realm:                                     realm,
		RealmID:                                   realmID,
		Role:                                      role,
		SimpleUser:                                simpleUser,
		Token:                                     token,
		UserID:                                    userID,
	}
}

func AAASuccessResponseLoginModelFromMap(m map[string]interface{}) *models.AAASuccessResponseLogin {
	var aPIToken *models.Token64 // Token64
	apiTokenInterface, apiTokenIsSet := m["api_token"]
	if apiTokenIsSet && apiTokenInterface != nil {
		apiTokenMap := apiTokenInterface.([]interface{})
		if len(apiTokenMap) > 0 {
			aPIToken = Token64ModelFromMap(apiTokenMap[0].(map[string]interface{}))
		}
	}
	//
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
	var enterpriseID *models.Identifier64 // Identifier64
	enterpriseIdInterface, enterpriseIdIsSet := m["enterprise_id"]
	if enterpriseIdIsSet && enterpriseIdInterface != nil {
		enterpriseIdMap := enterpriseIdInterface.([]interface{})
		if len(enterpriseIdMap) > 0 {
			enterpriseID = Identifier64ModelFromMap(enterpriseIdMap[0].(map[string]interface{}))
		}
	}
	//
	isJWTValid := m["is_j_w_t_valid"].(bool)
	var loginToken *models.Token64 // Token64
	loginTokenInterface, loginTokenIsSet := m["login_token"]
	if loginTokenIsSet && loginTokenInterface != nil {
		loginTokenMap := loginTokenInterface.([]interface{})
		if len(loginTokenMap) > 0 {
			loginToken = Token64ModelFromMap(loginTokenMap[0].(map[string]interface{}))
		}
	}
	//
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
	var realmID *models.Identifier64 // Identifier64
	realmIdInterface, realmIdIsSet := m["realm_id"]
	if realmIdIsSet && realmIdInterface != nil {
		realmIdMap := realmIdInterface.([]interface{})
		if len(realmIdMap) > 0 {
			realmID = Identifier64ModelFromMap(realmIdMap[0].(map[string]interface{}))
		}
	}
	//
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
	var token *models.Token64 // Token64
	tokenInterface, tokenIsSet := m["token"]
	if tokenIsSet && tokenInterface != nil {
		tokenMap := tokenInterface.([]interface{})
		if len(tokenMap) > 0 {
			token = Token64ModelFromMap(tokenMap[0].(map[string]interface{}))
		}
	}
	//
	var userID *models.Identifier64 // Identifier64
	userIdInterface, userIdIsSet := m["user_id"]
	if userIdIsSet && userIdInterface != nil {
		userIdMap := userIdInterface.([]interface{})
		if len(userIdMap) > 0 {
			userID = Identifier64ModelFromMap(userIdMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.AAASuccessResponseLogin{
		APIToken:     aPIToken,
		DetailedUser: detailedUser,
		Enterprise:   enterprise,
		EnterpriseID: enterpriseID,
		IsJWTValid:   isJWTValid,
		LoginToken:   loginToken,
		PasswordExpiryNotificationPeriodInSeconds: passwordExpiryNotificationPeriodInSeconds,
		PasswordExpiryTime:                        passwordExpiryTime,
		Policies:                                  policies,
		Realm:                                     realm,
		RealmID:                                   realmID,
		Role:                                      role,
		SimpleUser:                                simpleUser,
		Token:                                     token,
		UserID:                                    userID,
	}
}

func SetAAASuccessResponseLoginResourceData(d *schema.ResourceData, m *models.AAASuccessResponseLogin) {
	d.Set("api_token", SetToken64SubResourceData([]*models.Token64{m.APIToken}))
	d.Set("detailed_user", SetDetailedUserSubResourceData([]*models.DetailedUser{m.DetailedUser}))
	d.Set("enterprise", SetEnterpriseSubResourceData([]*models.Enterprise{m.Enterprise}))
	d.Set("enterprise_id", SetIdentifier64SubResourceData([]*models.Identifier64{m.EnterpriseID}))
	d.Set("is_j_w_t_valid", m.IsJWTValid)
	d.Set("login_token", SetToken64SubResourceData([]*models.Token64{m.LoginToken}))
	d.Set("password_expiry_notification_period_in_seconds", m.PasswordExpiryNotificationPeriodInSeconds)
	d.Set("password_expiry_time", m.PasswordExpiryTime)
	d.Set("policies", SetPolicyConfigSubResourceData(m.Policies))
	d.Set("realm", SetRealmSubResourceData([]*models.Realm{m.Realm}))
	d.Set("realm_id", SetIdentifier64SubResourceData([]*models.Identifier64{m.RealmID}))
	d.Set("role", SetRoleSubResourceData([]*models.Role{m.Role}))
	d.Set("simple_user", SetSimpleUserSubResourceData([]*models.SimpleUser{m.SimpleUser}))
	d.Set("token", SetToken64SubResourceData([]*models.Token64{m.Token}))
	d.Set("user_id", SetIdentifier64SubResourceData([]*models.Identifier64{m.UserID}))
}

func SetAAASuccessResponseLoginSubResourceData(m []*models.AAASuccessResponseLogin) (d []*map[string]interface{}) {
	for _, AAASuccessResponseLoginModel := range m {
		if AAASuccessResponseLoginModel != nil {
			properties := make(map[string]interface{})
			properties["api_token"] = SetToken64SubResourceData([]*models.Token64{AAASuccessResponseLoginModel.APIToken})
			properties["detailed_user"] = SetDetailedUserSubResourceData([]*models.DetailedUser{AAASuccessResponseLoginModel.DetailedUser})
			properties["enterprise"] = SetEnterpriseSubResourceData([]*models.Enterprise{AAASuccessResponseLoginModel.Enterprise})
			properties["enterprise_id"] = SetIdentifier64SubResourceData([]*models.Identifier64{AAASuccessResponseLoginModel.EnterpriseID})
			properties["is_j_w_t_valid"] = AAASuccessResponseLoginModel.IsJWTValid
			properties["login_token"] = SetToken64SubResourceData([]*models.Token64{AAASuccessResponseLoginModel.LoginToken})
			properties["password_expiry_notification_period_in_seconds"] = AAASuccessResponseLoginModel.PasswordExpiryNotificationPeriodInSeconds
			properties["password_expiry_time"] = AAASuccessResponseLoginModel.PasswordExpiryTime.String()
			properties["policies"] = SetPolicyConfigSubResourceData(AAASuccessResponseLoginModel.Policies)
			properties["realm"] = SetRealmSubResourceData([]*models.Realm{AAASuccessResponseLoginModel.Realm})
			properties["realm_id"] = SetIdentifier64SubResourceData([]*models.Identifier64{AAASuccessResponseLoginModel.RealmID})
			properties["role"] = SetRoleSubResourceData([]*models.Role{AAASuccessResponseLoginModel.Role})
			properties["simple_user"] = SetSimpleUserSubResourceData([]*models.SimpleUser{AAASuccessResponseLoginModel.SimpleUser})
			properties["token"] = SetToken64SubResourceData([]*models.Token64{AAASuccessResponseLoginModel.Token})
			properties["user_id"] = SetIdentifier64SubResourceData([]*models.Identifier64{AAASuccessResponseLoginModel.UserID})
			d = append(d, &properties)
		}
	}
	return
}

func AAASuccessResponseLoginSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"api_token": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Token64
			Elem: &schema.Resource{
				Schema: Token64Schema(),
			},
			Optional: true,
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

		"enterprise_id": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Identifier64
			Elem: &schema.Resource{
				Schema: Identifier64Schema(),
			},
			Optional: true,
		},

		"is_j_w_t_valid": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"login_token": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Token64
			Elem: &schema.Resource{
				Schema: Token64Schema(),
			},
			Optional: true,
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

		"realm_id": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Identifier64
			Elem: &schema.Resource{
				Schema: Identifier64Schema(),
			},
			Optional: true,
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
			Type:        schema.TypeList, //GoType: Identifier64
			Elem: &schema.Resource{
				Schema: Identifier64Schema(),
			},
			Optional: true,
		},
	}
}

func GetAAASuccessResponseLoginPropertyFields() (t []string) {
	return []string{
		"api_token",
		"detailed_user",
		"enterprise",
		"enterprise_id",
		"is_j_w_t_valid",
		"login_token",
		"password_expiry_notification_period_in_seconds",
		"password_expiry_time",
		"policies",
		"realm",
		"realm_id",
		"role",
		"simple_user",
		"token",
		"user_id",
	}
}
