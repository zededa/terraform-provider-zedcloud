package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func DetailedUserModel(d *schema.ResourceData) *models.DetailedUser {
	hubspotID, _ := d.Get("hubspot_id").(string)
	lastLoginTime, _ := d.Get("last_login_time").(strfmt.DateTime)
	lastLogoutTime, _ := d.Get("last_logout_time").(strfmt.DateTime)
	sfdcID, _ := d.Get("sfdc_id").(string)
	var allowedEnterprises []*models.AllowedEnterprise // []*AllowedEnterprise
	allowedEnterprisesInterface, allowedEnterprisesIsSet := d.GetOk("allowed_enterprises")
	if allowedEnterprisesIsSet {
		var items []interface{}
		if listItems, isList := allowedEnterprisesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = allowedEnterprisesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AllowedEnterpriseModelFromMap(v.(map[string]interface{}))
			allowedEnterprises = append(allowedEnterprises, m)
		}
	}
	customUserInput := map[string]string{}
	customUserInputInterface, customUserInputIsSet := d.GetOk("customUserInput")
	if customUserInputIsSet {
		customUserInputMap := customUserInputInterface.(map[string]interface{})
		for k, v := range customUserInputMap {
			if v == nil {
				continue
			}
			customUserInput[k] = v.(string)
		}
	}

	email, _ := d.Get("email").(string)
	firstName, _ := d.Get("first_name").(string)
	fullName, _ := d.Get("full_name").(string)
	id, _ := d.Get("id").(string)
	locale, _ := d.Get("locale").(string)
	notifyPref, _ := d.Get("notify_pref").(string)
	phone, _ := d.Get("phone").(string)
	roleID, _ := d.Get("role_id").(string)
	timeZone, _ := d.Get("time_zone").(string)
	var typeVar *models.AuthType // AuthType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewAuthType(models.AuthType(typeModel))
	}
	username, _ := d.Get("username").(string)
	return &models.DetailedUser{
		HubspotID:          hubspotID,
		LastLoginTime:      lastLoginTime,
		LastLogoutTime:     lastLogoutTime,
		SfdcID:             sfdcID,
		AllowedEnterprises: allowedEnterprises,
		CustomUserInput:    customUserInput,
		Email:              &email, // string
		FirstName:          firstName,
		FullName:           fullName,
		ID:                 id,
		Locale:             locale,
		NotifyPref:         notifyPref,
		Phone:              phone,
		RoleID:             &roleID, // string
		TimeZone:           timeZone,
		Type:               typeVar,
		Username:           &username, // string
	}
}

func DetailedUserModelFromMap(m map[string]interface{}) *models.DetailedUser {
	hubspotID := m["hubspot_id"].(string)
	lastLoginTime := m["last_login_time"].(strfmt.DateTime)
	lastLogoutTime := m["last_logout_time"].(strfmt.DateTime)
	sfdcID := m["sfdc_id"].(string)
	var allowedEnterprises []*models.AllowedEnterprise // []*AllowedEnterprise
	allowedEnterprisesInterface, allowedEnterprisesIsSet := m["allowed_enterprises"]
	if allowedEnterprisesIsSet {
		var items []interface{}
		if listItems, isList := allowedEnterprisesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = allowedEnterprisesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AllowedEnterpriseModelFromMap(v.(map[string]interface{}))
			allowedEnterprises = append(allowedEnterprises, m)
		}
	}
	customUserInput := map[string]string{}
	customUserInputInterface, customUserInputIsSet := m["custom_user_input"]
	if customUserInputIsSet {
		customUserInputMap := customUserInputInterface.(map[string]interface{})
		for k, v := range customUserInputMap {
			if v == nil {
				continue
			}
			customUserInput[k] = v.(string)
		}
	}

	email := m["email"].(string)
	firstName := m["first_name"].(string)
	fullName := m["full_name"].(string)
	id := m["id"].(string)
	locale := m["locale"].(string)
	notifyPref := m["notify_pref"].(string)
	phone := m["phone"].(string)
	roleID := m["role_id"].(string)
	timeZone := m["time_zone"].(string)
	var typeVar *models.AuthType // AuthType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewAuthType(models.AuthType(typeModel))
	}
	username := m["username"].(string)
	return &models.DetailedUser{
		HubspotID:          hubspotID,
		LastLoginTime:      lastLoginTime,
		LastLogoutTime:     lastLogoutTime,
		SfdcID:             sfdcID,
		AllowedEnterprises: allowedEnterprises,
		CustomUserInput:    customUserInput,
		Email:              &email,
		FirstName:          firstName,
		FullName:           fullName,
		ID:                 id,
		Locale:             locale,
		NotifyPref:         notifyPref,
		Phone:              phone,
		RoleID:             &roleID,
		TimeZone:           timeZone,
		Type:               typeVar,
		Username:           &username,
	}
}

func SetDetailedUserResourceData(d *schema.ResourceData, m *models.DetailedUser) {
	d.Set("hubspot_id", m.HubspotID)
	d.Set("last_login_time", m.LastLoginTime)
	d.Set("last_logout_time", m.LastLogoutTime)
	d.Set("sfdc_id", m.SfdcID)
	d.Set("allowed_enterprises", SetAllowedEnterpriseSubResourceData(m.AllowedEnterprises))
	d.Set("custom_user_input", m.CustomUserInput)
	d.Set("email", m.Email)
	d.Set("email_state", m.EmailState)
	d.Set("enterprise_id", m.EnterpriseID)
	d.Set("first_name", m.FirstName)
	d.Set("full_name", m.FullName)
	d.Set("id", m.ID)
	d.Set("locale", m.Locale)
	d.Set("notify_pref", m.NotifyPref)
	d.Set("phone", m.Phone)
	d.Set("phone_state", m.PhoneState)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("role_id", m.RoleID)
	d.Set("state", m.State)
	d.Set("time_zone", m.TimeZone)
	d.Set("totp_enabled", m.TotpEnabled)
	d.Set("type", m.Type)
	d.Set("username", m.Username)
}

func SetDetailedUserSubResourceData(m []*models.DetailedUser) (d []*map[string]interface{}) {
	for _, DetailedUserModel := range m {
		if DetailedUserModel != nil {
			properties := make(map[string]interface{})
			properties["hubspot_id"] = DetailedUserModel.HubspotID
			properties["last_login_time"] = DetailedUserModel.LastLoginTime.String()
			properties["last_logout_time"] = DetailedUserModel.LastLogoutTime.String()
			properties["sfdc_id"] = DetailedUserModel.SfdcID
			properties["allowed_enterprises"] = SetAllowedEnterpriseSubResourceData(DetailedUserModel.AllowedEnterprises)
			properties["custom_user_input"] = DetailedUserModel.CustomUserInput
			properties["email"] = DetailedUserModel.Email
			properties["email_state"] = DetailedUserModel.EmailState
			properties["enterprise_id"] = DetailedUserModel.EnterpriseID
			properties["first_name"] = DetailedUserModel.FirstName
			properties["full_name"] = DetailedUserModel.FullName
			properties["id"] = DetailedUserModel.ID
			properties["locale"] = DetailedUserModel.Locale
			properties["notify_pref"] = DetailedUserModel.NotifyPref
			properties["phone"] = DetailedUserModel.Phone
			properties["phone_state"] = DetailedUserModel.PhoneState
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{DetailedUserModel.Revision})
			properties["role_id"] = DetailedUserModel.RoleID
			properties["state"] = DetailedUserModel.State
			properties["time_zone"] = DetailedUserModel.TimeZone
			properties["totp_enabled"] = DetailedUserModel.TotpEnabled
			properties["type"] = DetailedUserModel.Type
			properties["username"] = DetailedUserModel.Username
			d = append(d, &properties)
		}
	}
	return
}

func DetailedUserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"hubspot_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"last_login_time": {
			Description:  `Last login time of the user`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"last_logout_time": {
			Description:  `Last logout time of the user`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"sfdc_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"allowed_enterprises": {
			Description: `Permitted list of enterprises with their associated roles`,
			Type:        schema.TypeList, //GoType: []*AllowedEnterprise
			Elem: &schema.Resource{
				Schema: AllowedEnterpriseSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"custom_user_input": {
			Description: `Custom user parameters`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"email": {
			Description: `Email of the user`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"email_state": {
			Description: `Email state`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"enterprise_id": {
			Description: `Origin enterprise of the user`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"first_name": {
			Description: `First name of the user`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"full_name": {
			Description: `Full name of the user`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `Unique system defined user ID`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"locale": {
			Description: `Locale of the user`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"notify_pref": {
			Description: `Notification preference of the user`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"phone": {
			Description: `Phone number of the user`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"phone_state": {
			Description: `Phone state`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"revision": {
			Description: `system defined info`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Computed: true,
		},

		"role_id": {
			Description: `Role associated with the user`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"state": {
			Description: `User state`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"time_zone": {
			Description: `Preferred time zone of the user`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"totp_enabled": {
			Description: `Is TOTP enrolment enabled`,
			Type:        schema.TypeBool,
			Computed:    true,
		},

		"type": {
			Description: `Type of the user`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"username": {
			Description: `User defined name`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func GetDetailedUserPropertyFields() (t []string) {
	return []string{
		"hubspot_id",
		"last_login_time",
		"last_logout_time",
		"sfdc_id",
		"allowed_enterprises",
		"custom_user_input",
		"email",
		"first_name",
		"full_name",
		"id",
		"locale",
		"notify_pref",
		"phone",
		"role_id",
		"time_zone",
		"type",
		"username",
	}
}
