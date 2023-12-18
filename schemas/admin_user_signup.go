package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AdminUserSignupModel(d *schema.ResourceData) *models.AdminUserSignup {
	email, _ := d.Get("email").(string)
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
	return &models.AdminUserSignup{
		Email:     email,
		FirstName: firstName,
		FullName:  fullName,
		Name:      name,
		Password:  password,
		Type:      typeVar,
	}
}

func AdminUserSignupModelFromMap(m map[string]interface{}) *models.AdminUserSignup {
	email := m["email"].(string)
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
	return &models.AdminUserSignup{
		Email:     email,
		FirstName: firstName,
		FullName:  fullName,
		Name:      name,
		Password:  password,
		Type:      typeVar,
	}
}

func SetAdminUserSignupResourceData(d *schema.ResourceData, m *models.AdminUserSignup) {
	d.Set("email", m.Email)
	d.Set("first_name", m.FirstName)
	d.Set("full_name", m.FullName)
	d.Set("name", m.Name)
	d.Set("password", m.Password)
	d.Set("type", m.Type)
}

func SetAdminUserSignupSubResourceData(m []*models.AdminUserSignup) (d []*map[string]interface{}) {
	for _, AdminUserSignupModel := range m {
		if AdminUserSignupModel != nil {
			properties := make(map[string]interface{})
			properties["email"] = AdminUserSignupModel.Email
			properties["first_name"] = AdminUserSignupModel.FirstName
			properties["full_name"] = AdminUserSignupModel.FullName
			properties["name"] = AdminUserSignupModel.Name
			properties["password"] = AdminUserSignupModel.Password
			properties["type"] = AdminUserSignupModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func AdminUserSignupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"email": {
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

func GetAdminUserSignupPropertyFields() (t []string) {
	return []string{
		"email",
		"first_name",
		"full_name",
		"name",
		"password",
		"type",
	}
}
