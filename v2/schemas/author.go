package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AuthorModel(d *schema.ResourceData) *models.Author {
	company, _ := d.Get("company").(string)
	email, _ := d.Get("email").(string)
	group, _ := d.Get("group").(string)
	user, _ := d.Get("user").(string)
	website, _ := d.Get("website").(string)
	return &models.Author{
		Company: company,
		Email:   email,
		Group:   group,
		User:    user,
		Website: website,
	}
}

func AuthorModelFromMap(m map[string]interface{}) *models.Author {
	company := m["company"].(string)
	email := m["email"].(string)
	group := m["group"].(string)
	user := m["user"].(string)
	website := m["website"].(string)
	return &models.Author{
		Company: company,
		Email:   email,
		Group:   group,
		User:    user,
		Website: website,
	}
}

func SetAuthorResourceData(d *schema.ResourceData, m *models.Author) {
	d.Set("company", m.Company)
	d.Set("email", m.Email)
	d.Set("group", m.Group)
	d.Set("user", m.User)
	d.Set("website", m.Website)
}

func SetAuthorSubResourceData(m []*models.Author) (d []*map[string]interface{}) {
	for _, AuthorModel := range m {
		if AuthorModel != nil {
			properties := make(map[string]interface{})
			properties["company"] = AuthorModel.Company
			properties["email"] = AuthorModel.Email
			properties["group"] = AuthorModel.Group
			properties["user"] = AuthorModel.User
			properties["website"] = AuthorModel.Website
			d = append(d, &properties)
		}
	}
	return
}

func Author() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"company": {
			Description: `UI map: AppEditPage:IdentityPane:Category_Field, AppDetailsPage:IdentityPane:Category_Field`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"email": {
			Description: `UI map: AppEditPage:DeveloperPane:Email_Field, AppDetailsPage:DeveloperPane:Email_Field`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"group": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"user": {
			Description: `UI map: AppEditPage:DeveloperPane:Name_Field, AppDetailsPage:DeveloperPane:Name_Field`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"website": {
			Description: `UI map: AppEditPage:DeveloperPane:Website_Field, AppDetailsPage:DeveloperPane:Website_Field`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAuthorPropertyFields() (t []string) {
	return []string{
		"company",
		"email",
		"group",
		"user",
		"website",
	}
}
