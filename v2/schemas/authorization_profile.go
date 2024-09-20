package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AuthorizationProfileModel(d *schema.ResourceData) *models.AuthorizationProfile {
	active, _ := d.Get("active").(bool)
	defaultRoleID, _ := d.Get("default_role_id").(string)
	description, _ := d.Get("description").(string)
	disableAutoUserCreate, _ := d.Get("disable_auto_user_create").(bool)
	enterpriseID, _ := d.Get("enterprise_id").(string)
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	var oauthProfile *models.OAUTHProfile // OAUTHProfile
	oauthProfileInterface, oauthProfileIsSet := d.GetOk("oauth_profile")
	if oauthProfileIsSet && oauthProfileInterface != nil {
		oauthProfileMap := oauthProfileInterface.([]interface{})
		if len(oauthProfileMap) > 0 {
			oauthProfile = OAUTHProfileModelFromMap(oauthProfileMap[0].(map[string]interface{}))
		}
	}
	var passwordProfile *models.PasswordProfile // PasswordProfile
	passwordProfileInterface, passwordProfileIsSet := d.GetOk("password_profile")
	if passwordProfileIsSet && passwordProfileInterface != nil {
		passwordProfileMap := passwordProfileInterface.([]interface{})
		if len(passwordProfileMap) > 0 {
			passwordProfile = PasswordProfileModelFromMap(passwordProfileMap[0].(map[string]interface{}))
		}
	}
	var profileType *models.AuthProfileType // AuthProfileType
	profileTypeInterface, profileTypeIsSet := d.GetOk("profile_type")
	if profileTypeIsSet {
		profileTypeModel := profileTypeInterface.(string)
		profileType = models.NewAuthProfileType(models.AuthProfileType(profileTypeModel))
	}
	testOnly, _ := d.Get("test_only").(bool)
	title, _ := d.Get("title").(string)
	var typeVar *models.AuthType // AuthType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewAuthType(models.AuthType(typeModel))
	}
	return &models.AuthorizationProfile{
		Active:                active,
		DefaultRoleID:         &defaultRoleID, // string
		Description:           description,
		DisableAutoUserCreate: disableAutoUserCreate,
		EnterpriseID:          enterpriseID,
		ID:                    id,
		Name:                  &name, // string
		OauthProfile:          oauthProfile,
		PasswordProfile:       passwordProfile,
		ProfileType:           profileType,
		TestOnly:              testOnly,
		Title:                 &title, // string
		Type:                  typeVar,
	}
}

func AuthorizationProfileModelFromMap(m map[string]interface{}) *models.AuthorizationProfile {
	active := m["active"].(bool)
	defaultRoleID := m["default_role_id"].(string)
	description := m["description"].(string)
	disableAutoUserCreate := m["disable_auto_user_create"].(bool)
	enterpriseID := m["enterprise_id"].(string)
	id := m["id"].(string)
	name := m["name"].(string)
	var oauthProfile *models.OAUTHProfile // OAUTHProfile
	oauthProfileInterface, oauthProfileIsSet := m["oauth_profile"]
	if oauthProfileIsSet && oauthProfileInterface != nil {
		oauthProfileMap := oauthProfileInterface.([]interface{})
		if len(oauthProfileMap) > 0 {
			oauthProfile = OAUTHProfileModelFromMap(oauthProfileMap[0].(map[string]interface{}))
		}
	}
	//
	var passwordProfile *models.PasswordProfile // PasswordProfile
	passwordProfileInterface, passwordProfileIsSet := m["password_profile"]
	if passwordProfileIsSet && passwordProfileInterface != nil {
		passwordProfileMap := passwordProfileInterface.([]interface{})
		if len(passwordProfileMap) > 0 {
			passwordProfile = PasswordProfileModelFromMap(passwordProfileMap[0].(map[string]interface{}))
		}
	}
	//
	var profileType *models.AuthProfileType // AuthProfileType
	profileTypeInterface, profileTypeIsSet := m["profile_type"]
	if profileTypeIsSet {
		profileTypeModel := profileTypeInterface.(string)
		profileType = models.NewAuthProfileType(models.AuthProfileType(profileTypeModel))
	}
	testOnly := m["test_only"].(bool)
	title := m["title"].(string)
	var typeVar *models.AuthType // AuthType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewAuthType(models.AuthType(typeModel))
	}
	return &models.AuthorizationProfile{
		Active:                active,
		DefaultRoleID:         &defaultRoleID,
		Description:           description,
		DisableAutoUserCreate: disableAutoUserCreate,
		EnterpriseID:          enterpriseID,
		ID:                    id,
		Name:                  &name,
		OauthProfile:          oauthProfile,
		PasswordProfile:       passwordProfile,
		ProfileType:           profileType,
		TestOnly:              testOnly,
		Title:                 &title,
		Type:                  typeVar,
	}
}

func SetAuthorizationProfileResourceData(d *schema.ResourceData, m *models.AuthorizationProfile) {
	d.Set("active", m.Active)
	d.Set("default_role_id", m.DefaultRoleID)
	d.Set("description", m.Description)
	d.Set("disable_auto_user_create", m.DisableAutoUserCreate)
	d.Set("enterprise_id", m.EnterpriseID)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("oauth_profile", SetOAUTHProfileSubResourceData([]*models.OAUTHProfile{m.OauthProfile}))
	d.Set("password_profile", SetPasswordProfileSubResourceData([]*models.PasswordProfile{m.PasswordProfile}))
	d.Set("profile_type", m.ProfileType)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("test_only", m.TestOnly)
	d.Set("title", m.Title)
	d.Set("type", m.Type)
}

func SetAuthorizationProfileSubResourceData(m []*models.AuthorizationProfile) (d []*map[string]interface{}) {
	for _, AuthorizationProfileModel := range m {
		if AuthorizationProfileModel != nil {
			properties := make(map[string]interface{})
			properties["active"] = AuthorizationProfileModel.Active
			properties["default_role_id"] = AuthorizationProfileModel.DefaultRoleID
			properties["description"] = AuthorizationProfileModel.Description
			properties["disable_auto_user_create"] = AuthorizationProfileModel.DisableAutoUserCreate
			properties["enterprise_id"] = AuthorizationProfileModel.EnterpriseID
			properties["id"] = AuthorizationProfileModel.ID
			properties["name"] = AuthorizationProfileModel.Name
			properties["oauth_profile"] = SetOAUTHProfileSubResourceData([]*models.OAUTHProfile{AuthorizationProfileModel.OauthProfile})
			properties["password_profile"] = SetPasswordProfileSubResourceData([]*models.PasswordProfile{AuthorizationProfileModel.PasswordProfile})
			properties["profile_type"] = AuthorizationProfileModel.ProfileType
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{AuthorizationProfileModel.Revision})
			properties["test_only"] = AuthorizationProfileModel.TestOnly
			properties["title"] = AuthorizationProfileModel.Title
			properties["type"] = AuthorizationProfileModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func AuthorizationProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active": {
			Description: `Mark this profile as active. Only one profile can be active in a given enterprise`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"default_role_id": {
			Description: `Default Role ID to associate with the profile`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"description": {
			Description: `Detailed description of the profile`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"disable_auto_user_create": {
			Description: `Do not automatically create new users if this is set`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"enterprise_id": {
			Description: `Parent enterprise ID of the authorization profile`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `Unique system defined profile ID`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `User defined name of the profile. Profile name is unique within an enterprise. Name can't be changed once created`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"oauth_profile": {
			Description: `Oauth profile configuration details`,
			Type:        schema.TypeList, //GoType: OAUTHProfile
			Elem: &schema.Resource{
				Schema: OAUTHProfileSchema(),
			},
			Optional: true,
		},

		"password_profile": {
			Description: ``,
			Type:        schema.TypeList, //GoType: PasswordProfile
			Elem: &schema.Resource{
				Schema: PasswordProfileSchema(),
			},
			Optional: true,
		},

		"profile_type": {
			Description: `Authorization profile type`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"revision": {
			Description: `system defined info`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Computed: true,
		},

		"test_only": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"title": {
			Description: `User defined title for the profile. Title can be changed anytime`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"type": {
			Description: `Type of the profile`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAuthorizationProfilePropertyFields() (t []string) {
	return []string{
		"active",
		"default_role_id",
		"description",
		"disable_auto_user_create",
		"enterprise_id",
		"id",
		"name",
		"oauth_profile",
		"password_profile",
		"profile_type",
		"test_only",
		"title",
		"type",
	}
}
