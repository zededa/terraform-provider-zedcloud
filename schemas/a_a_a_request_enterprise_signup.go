package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAARequestEnterpriseSignupModel(d *schema.ResourceData) *models.AAARequestEnterpriseSignup {
	var adminUser *models.AdminUserSignup // AdminUserSignup
	adminUserInterface, adminUserIsSet := d.GetOk("admin_user")
	if adminUserIsSet && adminUserInterface != nil {
		adminUserMap := adminUserInterface.([]interface{})
		if len(adminUserMap) > 0 {
			adminUser = AdminUserSignupModelFromMap(adminUserMap[0].(map[string]interface{}))
		}
	}
	employerName, _ := d.Get("employer_name").(string)
	var enterprise *models.Enterprise // Enterprise
	enterpriseInterface, enterpriseIsSet := d.GetOk("enterprise")
	if enterpriseIsSet && enterpriseInterface != nil {
		enterpriseMap := enterpriseInterface.([]interface{})
		if len(enterpriseMap) > 0 {
			enterprise = EnterpriseModelFromMap(enterpriseMap[0].(map[string]interface{}))
		}
	}
	var profileType *models.AuthProfileType // AuthProfileType
	profileTypeInterface, profileTypeIsSet := d.GetOk("profile_type")
	if profileTypeIsSet {
		profileTypeModel := profileTypeInterface.(string)
		profileType = models.NewAuthProfileType(models.AuthProfileType(profileTypeModel))
	}
	var realmList []string
	realmListInterface, realmListIsSet := d.GetOk("realmList")
	if realmListIsSet {
		var items []interface{}
		if listItems, isList := realmListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = realmListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			realmList = append(realmList, v.(string))
		}
	}
	title, _ := d.Get("title").(string)
	token, _ := d.Get("token").(string)
	return &models.AAARequestEnterpriseSignup{
		AdminUser:    adminUser,
		EmployerName: employerName,
		Enterprise:   enterprise,
		ProfileType:  profileType,
		RealmList:    realmList,
		Title:        title,
		Token:        token,
	}
}

func AAARequestEnterpriseSignupModelFromMap(m map[string]interface{}) *models.AAARequestEnterpriseSignup {
	var adminUser *models.AdminUserSignup // AdminUserSignup
	adminUserInterface, adminUserIsSet := m["admin_user"]
	if adminUserIsSet && adminUserInterface != nil {
		adminUserMap := adminUserInterface.([]interface{})
		if len(adminUserMap) > 0 {
			adminUser = AdminUserSignupModelFromMap(adminUserMap[0].(map[string]interface{}))
		}
	}
	//
	employerName := m["employer_name"].(string)
	var enterprise *models.Enterprise // Enterprise
	enterpriseInterface, enterpriseIsSet := m["enterprise"]
	if enterpriseIsSet && enterpriseInterface != nil {
		enterpriseMap := enterpriseInterface.([]interface{})
		if len(enterpriseMap) > 0 {
			enterprise = EnterpriseModelFromMap(enterpriseMap[0].(map[string]interface{}))
		}
	}
	//
	var profileType *models.AuthProfileType // AuthProfileType
	profileTypeInterface, profileTypeIsSet := m["profile_type"]
	if profileTypeIsSet {
		profileTypeModel := profileTypeInterface.(string)
		profileType = models.NewAuthProfileType(models.AuthProfileType(profileTypeModel))
	}
	var realmList []string
	realmListInterface, realmListIsSet := m["realmList"]
	if realmListIsSet {
		var items []interface{}
		if listItems, isList := realmListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = realmListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			realmList = append(realmList, v.(string))
		}
	}
	title := m["title"].(string)
	token := m["token"].(string)
	return &models.AAARequestEnterpriseSignup{
		AdminUser:    adminUser,
		EmployerName: employerName,
		Enterprise:   enterprise,
		ProfileType:  profileType,
		RealmList:    realmList,
		Title:        title,
		Token:        token,
	}
}

func SetAAARequestEnterpriseSignupResourceData(d *schema.ResourceData, m *models.AAARequestEnterpriseSignup) {
	d.Set("admin_user", SetAdminUserSignupSubResourceData([]*models.AdminUserSignup{m.AdminUser}))
	d.Set("employer_name", m.EmployerName)
	d.Set("enterprise", SetEnterpriseSubResourceData([]*models.Enterprise{m.Enterprise}))
	d.Set("profile_type", m.ProfileType)
	d.Set("realm_list", m.RealmList)
	d.Set("title", m.Title)
	d.Set("token", m.Token)
}

func SetAAARequestEnterpriseSignupSubResourceData(m []*models.AAARequestEnterpriseSignup) (d []*map[string]interface{}) {
	for _, AAARequestEnterpriseSignupModel := range m {
		if AAARequestEnterpriseSignupModel != nil {
			properties := make(map[string]interface{})
			properties["admin_user"] = SetAdminUserSignupSubResourceData([]*models.AdminUserSignup{AAARequestEnterpriseSignupModel.AdminUser})
			properties["employer_name"] = AAARequestEnterpriseSignupModel.EmployerName
			properties["enterprise"] = SetEnterpriseSubResourceData([]*models.Enterprise{AAARequestEnterpriseSignupModel.Enterprise})
			properties["profile_type"] = AAARequestEnterpriseSignupModel.ProfileType
			properties["realm_list"] = AAARequestEnterpriseSignupModel.RealmList
			properties["title"] = AAARequestEnterpriseSignupModel.Title
			properties["token"] = AAARequestEnterpriseSignupModel.Token
			d = append(d, &properties)
		}
	}
	return
}

func AAARequestEnterpriseSignupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"admin_user": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AdminUserSignup
			Elem: &schema.Resource{
				Schema: AdminUserSignupSchema(),
			},
			Optional: true,
		},

		"employer_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"enterprise": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Enterprise
			Elem: &schema.Resource{
				Schema: EnterpriseSchema(),
			},
			Optional: true,
		},

		"profile_type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"realm_list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"title": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"token": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAAARequestEnterpriseSignupPropertyFields() (t []string) {
	return []string{
		"admin_user",
		"employer_name",
		"enterprise",
		"profile_type",
		"realm_list",
		"title",
		"token",
	}
}
