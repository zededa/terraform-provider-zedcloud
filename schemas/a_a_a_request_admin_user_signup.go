package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAARequestAdminUserSignupModel(d *schema.ResourceData) *models.AAARequestAdminUserSignup {
	email, _ := d.Get("email").(string)
	enterpriseName, _ := d.Get("enterprise_name").(string)
	firstName, _ := d.Get("first_name").(string)
	fullName, _ := d.Get("full_name").(string)
	name, _ := d.Get("name").(string)
	password, _ := d.Get("password").(string)
	var typeVar *models.CredentialType // CredentialType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewCredentialType(models.CredentialType(typeModel))
	}
	return &models.AAARequestAdminUserSignup{
		Email:          email,
		EnterpriseName: enterpriseName,
		FirstName:      firstName,
		FullName:       fullName,
		Name:           name,
		Password:       password,
		Type:           typeVar,
	}
}

func AAARequestAdminUserSignupModelFromMap(m map[string]interface{}) *models.AAARequestAdminUserSignup {
	email := m["email"].(string)
	enterpriseName := m["enterprise_name"].(string)
	firstName := m["first_name"].(string)
	fullName := m["full_name"].(string)
	name := m["name"].(string)
	password := m["password"].(string)
	var typeVar *models.CredentialType // CredentialType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewCredentialType(models.CredentialType(typeModel))
	}
	return &models.AAARequestAdminUserSignup{
		Email:          email,
		EnterpriseName: enterpriseName,
		FirstName:      firstName,
		FullName:       fullName,
		Name:           name,
		Password:       password,
		Type:           typeVar,
	}
}

func SetAAARequestAdminUserSignupResourceData(d *schema.ResourceData, m *models.AAARequestAdminUserSignup) {
	d.Set("email", m.Email)
	d.Set("enterprise_name", m.EnterpriseName)
	d.Set("first_name", m.FirstName)
	d.Set("full_name", m.FullName)
	d.Set("name", m.Name)
	d.Set("password", m.Password)
	d.Set("type", m.Type)
}

func SetAAARequestAdminUserSignupSubResourceData(m []*models.AAARequestAdminUserSignup) (d []*map[string]interface{}) {
	for _, AAARequestAdminUserSignupModel := range m {
		if AAARequestAdminUserSignupModel != nil {
			properties := make(map[string]interface{})
			properties["email"] = AAARequestAdminUserSignupModel.Email
			properties["enterprise_name"] = AAARequestAdminUserSignupModel.EnterpriseName
			properties["first_name"] = AAARequestAdminUserSignupModel.FirstName
			properties["full_name"] = AAARequestAdminUserSignupModel.FullName
			properties["name"] = AAARequestAdminUserSignupModel.Name
			properties["password"] = AAARequestAdminUserSignupModel.Password
			properties["type"] = AAARequestAdminUserSignupModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func AAARequestAdminUserSignupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"email": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"enterprise_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"first_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"full_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"password": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAAARequestAdminUserSignupPropertyFields() (t []string) {
	return []string{
		"email",
		"enterprise_name",
		"first_name",
		"full_name",
		"name",
		"password",
		"type",
	}
}
