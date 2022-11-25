package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate Author resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AuthorModel(d *schema.ResourceData) *models.Author {
	company := d.Get("company").(string)
	email := d.Get("email").(string)
	group := d.Get("group").(string)
	user := d.Get("user").(string)
	website := d.Get("website").(string)
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

// Update the underlying Author resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAuthorResourceData(d *schema.ResourceData, m *models.Author) {
	d.Set("company", m.Company)
	d.Set("email", m.Email)
	d.Set("group", m.Group)
	d.Set("user", m.User)
	d.Set("website", m.Website)
}

// Iterate throught and update the Author resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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

// Schema mapping representing the Author resource defined in the Terraform configuration
func AuthorSchema() map[string]*schema.Schema {
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

// Retrieve property field names for updating the Author resource
func GetAuthorPropertyFields() (t []string) {
	return []string{
		"company",
		"email",
		"group",
		"user",
		"website",
	}
}
