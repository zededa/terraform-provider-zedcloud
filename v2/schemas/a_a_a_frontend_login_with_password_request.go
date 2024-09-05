package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AAAFrontendLoginWithPasswordRequestModel(d *schema.ResourceData) *models.AAAFrontendLoginWithPasswordRequest {
	enterpriseName, _ := d.Get("enterprise_name").(string)
	password, _ := d.Get("password").(string)
	realm, _ := d.Get("realm").(string)
	username, _ := d.Get("username").(string)
	usernameAtRealm, _ := d.Get("username_at_realm").(string)
	verboseDetailedUser, _ := d.Get("verbose_detailed_user").(bool)
	verboseEnterprise, _ := d.Get("verbose_enterprise").(bool)
	verbosePolicy, _ := d.Get("verbose_policy").(bool)
	verboseRealm, _ := d.Get("verbose_realm").(bool)
	verboseSimpleUser, _ := d.Get("verbose_simple_user").(bool)
	return &models.AAAFrontendLoginWithPasswordRequest{
		EnterpriseName:      enterpriseName,
		Password:            password,
		Realm:               realm,
		Username:            username,
		UsernameAtRealm:     usernameAtRealm,
		VerboseDetailedUser: verboseDetailedUser,
		VerboseEnterprise:   verboseEnterprise,
		VerbosePolicy:       verbosePolicy,
		VerboseRealm:        verboseRealm,
		VerboseSimpleUser:   verboseSimpleUser,
	}
}

func AAAFrontendLoginWithPasswordRequestModelFromMap(m map[string]interface{}) *models.AAAFrontendLoginWithPasswordRequest {
	enterpriseName := m["enterprise_name"].(string)
	password := m["password"].(string)
	realm := m["realm"].(string)
	username := m["username"].(string)
	usernameAtRealm := m["username_at_realm"].(string)
	verboseDetailedUser := m["verbose_detailed_user"].(bool)
	verboseEnterprise := m["verbose_enterprise"].(bool)
	verbosePolicy := m["verbose_policy"].(bool)
	verboseRealm := m["verbose_realm"].(bool)
	verboseSimpleUser := m["verbose_simple_user"].(bool)
	return &models.AAAFrontendLoginWithPasswordRequest{
		EnterpriseName:      enterpriseName,
		Password:            password,
		Realm:               realm,
		Username:            username,
		UsernameAtRealm:     usernameAtRealm,
		VerboseDetailedUser: verboseDetailedUser,
		VerboseEnterprise:   verboseEnterprise,
		VerbosePolicy:       verbosePolicy,
		VerboseRealm:        verboseRealm,
		VerboseSimpleUser:   verboseSimpleUser,
	}
}

func SetAAAFrontendLoginWithPasswordRequestResourceData(d *schema.ResourceData, m *models.AAAFrontendLoginWithPasswordRequest) {
	d.Set("enterprise_name", m.EnterpriseName)
	d.Set("password", m.Password)
	d.Set("realm", m.Realm)
	d.Set("username", m.Username)
	d.Set("username_at_realm", m.UsernameAtRealm)
	d.Set("verbose_detailed_user", m.VerboseDetailedUser)
	d.Set("verbose_enterprise", m.VerboseEnterprise)
	d.Set("verbose_policy", m.VerbosePolicy)
	d.Set("verbose_realm", m.VerboseRealm)
	d.Set("verbose_simple_user", m.VerboseSimpleUser)
}

func SetAAAFrontendLoginWithPasswordRequestSubResourceData(m []*models.AAAFrontendLoginWithPasswordRequest) (d []*map[string]interface{}) {
	for _, AAAFrontendLoginWithPasswordRequestModel := range m {
		if AAAFrontendLoginWithPasswordRequestModel != nil {
			properties := make(map[string]interface{})
			properties["enterprise_name"] = AAAFrontendLoginWithPasswordRequestModel.EnterpriseName
			properties["password"] = AAAFrontendLoginWithPasswordRequestModel.Password
			properties["realm"] = AAAFrontendLoginWithPasswordRequestModel.Realm
			properties["username"] = AAAFrontendLoginWithPasswordRequestModel.Username
			properties["username_at_realm"] = AAAFrontendLoginWithPasswordRequestModel.UsernameAtRealm
			properties["verbose_detailed_user"] = AAAFrontendLoginWithPasswordRequestModel.VerboseDetailedUser
			properties["verbose_enterprise"] = AAAFrontendLoginWithPasswordRequestModel.VerboseEnterprise
			properties["verbose_policy"] = AAAFrontendLoginWithPasswordRequestModel.VerbosePolicy
			properties["verbose_realm"] = AAAFrontendLoginWithPasswordRequestModel.VerboseRealm
			properties["verbose_simple_user"] = AAAFrontendLoginWithPasswordRequestModel.VerboseSimpleUser
			d = append(d, &properties)
		}
	}
	return
}

func AAAFrontendLoginWithPasswordRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enterprise_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"password": {
			Description: ``,
			Type:        schema.TypeString,
			Required:    true,
		},

		"realm": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"username": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"username_at_realm": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"verbose_detailed_user": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"verbose_enterprise": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"verbose_policy": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"verbose_realm": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"verbose_simple_user": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}

func GetAAAFrontendLoginWithPasswordRequestPropertyFields() (t []string) {
	return []string{
		"enterprise_name",
		"password",
		"realm",
		"username",
		"username_at_realm",
		"verbose_detailed_user",
		"verbose_enterprise",
		"verbose_policy",
		"verbose_realm",
		"verbose_simple_user",
	}
}
