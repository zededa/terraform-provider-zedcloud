package schemas

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// customParamKeyMaxLen and customUserInputKey mirror the controller's own
// rules for a custom user parameter (CustomParamKeyMaxLen and
// ValidateCustomParam in zedcloud's libs/zutils/validate.go): the key is three
// '_'-separated segments drawn from a restricted alphabet.
//
// Enforcing them here turns an opaque "HTTP status code: 400 / Invalid
// characters in the key" at apply time into a plan-time error that names the
// offending key. Found the hard way, running the acceptance suite against a
// controller for the first time once NFR-165 §3.4 made this map reach the API
// at all.
const (
	customParamKeyMaxLen = 1024
	customParamValMaxLen = 1024
)

var customUserInputKey = regexp.MustCompile(
	`^[a-zA-Z0-9\-.%@#:~!=]+_[a-zA-Z0-9\-.%@#:~!=]+_[a-zA-Z0-9\-.%@#:~!=]+$`)

func validateCustomUserInput(v interface{}, path cty.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	entries, ok := v.(map[string]interface{})
	if !ok {
		return diags
	}

	for key, raw := range entries {
		value, _ := raw.(string)

		switch {
		case key == "" || value == "":
			diags = append(diags, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "custom_user_input entry is empty",
				Detail:        "Neither the key nor the value of a custom user parameter may be empty.",
				AttributePath: path,
			})
		case len(key) > customParamKeyMaxLen:
			diags = append(diags, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "custom_user_input key is too long",
				Detail:        fmt.Sprintf("Key %q is %d characters; the maximum is %d.", key, len(key), customParamKeyMaxLen),
				AttributePath: path,
			})
		case !customUserInputKey.MatchString(key):
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "invalid custom_user_input key",
				Detail: fmt.Sprintf(
					"Key %q is not accepted by the controller. A key must be exactly three "+
						"'_'-separated segments, each made up of letters, digits or -.%%@#:~!=  "+
						"— for example \"acme_ui_theme\".", key),
				AttributePath: path,
			})
		case len(value) > customParamValMaxLen:
			diags = append(diags, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "custom_user_input value is too long",
				Detail:        fmt.Sprintf("The value for key %q is %d characters; the maximum is %d.", key, len(value), customParamValMaxLen),
				AttributePath: path,
			})
		}
	}

	return diags
}

// LastLoginTime and LastLogoutTime are deliberately absent from both model
// builders below. They are server-observed and read-only; sending them back
// would at best be ignored. See NFR-165 §3.6 — the previous code asserted
// d.Get("last_login_time") to strfmt.DateTime, which can never succeed for a
// TypeString attribute, so the zero value was sent on every request anyway.

func DetailedUserModel(d *schema.ResourceData) *models.DetailedUser {
	hubspotID, _ := d.Get("hubspot_id").(string)
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
	// NFR-165 §3.4: the schema key is snake_case. Reading "customUserInput"
	// here always missed, so custom user parameters were dropped on every
	// create and update and an imported user never converged.
	customUserInput := stringMap(d.Get("custom_user_input"))

	email, _ := d.Get("email").(string)
	firstName, _ := d.Get("first_name").(string)
	fullName, _ := d.Get("full_name").(string)
	id, _ := d.Get("id").(string)
	locale, _ := d.Get("locale").(string)
	notifyPref, _ := d.Get("notify_pref").(string)
	phone, _ := d.Get("phone").(string)
	roleID, _ := d.Get("role_id").(string)
	var state *models.UserState
	stateInterface, stateIsSet := d.GetOk("state")
	if stateIsSet {
		stateStr := stateInterface.(string)
		state = models.NewUserState(models.UserState(stateStr))
	}
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
		State:              state,   // UserState
		TimeZone:           timeZone,
		Type:               typeVar,
		Username:           &username, // string
	}
}

func DetailedUserModelFromMap(m map[string]interface{}) *models.DetailedUser {
	hubspotID := m["hubspot_id"].(string)
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
	customUserInput := stringMap(m["custom_user_input"])

	email := m["email"].(string)
	firstName := m["first_name"].(string)
	fullName := m["full_name"].(string)
	id := m["id"].(string)
	locale := m["locale"].(string)
	notifyPref := m["notify_pref"].(string)
	phone := m["phone"].(string)
	roleID := m["role_id"].(string)
	var state *models.UserState
	stateInterface, stateIsSet := m["state"]
	if stateIsSet {
		stateStr := stateInterface.(string)
		state = models.NewUserState(models.UserState(stateStr))
	}
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
		State:              state,
		TimeZone:           timeZone,
		Type:               typeVar,
		Username:           &username,
	}
}

func SetDetailedUserResourceData(d *schema.ResourceData, m *models.DetailedUser) {
	d.Set("hubspot_id", m.HubspotID)
	d.Set("last_login_time", m.LastLoginTime.String())
	d.Set("last_logout_time", m.LastLogoutTime.String())
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

		// NFR-165 §4.2: server-observed, never configurable. Previously
		// Optional with a blanket DiffSuppressFunc, which hid the fact that
		// the value could not round-trip at all.
		"last_login_time": {
			Description: `Last login time of the user`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"last_logout_time": {
			Description: `Last logout time of the user`,
			Type:        schema.TypeString,
			Computed:    true,
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
			Optional:         true,
			DiffSuppressFunc: diffSuppressChangedList("allowed_enterprises"),
		},

		"custom_user_input": {
			Description: `Custom user parameters. Each key must be exactly three ` +
				`'_'-separated segments, for example "acme_ui_theme".`,
			Type: schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional:         true,
			ValidateDiagFunc: validateCustomUserInput,
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

		// NFR-165 §3.7: the account identifier. The API has no rename, so a
		// change has to plan a replacement rather than an update that fails
		// at apply time.
		"username": {
			Description: `User defined name`,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
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
		"state",
		"time_zone",
		"type",
		"username",
	}
}
