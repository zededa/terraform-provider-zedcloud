package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func PrivateRepoRepositoryDetailsModel(d *schema.ResourceData) *models.PrivateRepoRepositoryDetails {
	branch, _ := d.Get("branch").(string)
	var typeVar *models.PrivateRepoType // PrivateRepoType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewPrivateRepoType(models.PrivateRepoType(typeModel))
	}
	url, _ := d.Get("url").(string)
	return &models.PrivateRepoRepositoryDetails{
		Branch: branch,
		Type:   typeVar,
		URL:    &url, // string
	}
}

func PrivateRepoRepositoryDetailsModelFromMap(m map[string]interface{}) *models.PrivateRepoRepositoryDetails {
	branch := m["branch"].(string)
	var typeVar *models.PrivateRepoType // PrivateRepoType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewPrivateRepoType(models.PrivateRepoType(typeModel))
	}
	url := m["url"].(string)
	return &models.PrivateRepoRepositoryDetails{
		Branch: branch,
		Type:   typeVar,
		URL:    &url,
	}
}

func SetPrivateRepoRepositoryDetailsResourceData(d *schema.ResourceData, m *models.PrivateRepoRepositoryDetails) {
	d.Set("branch", m.Branch)
	d.Set("type", m.Type)
	d.Set("url", m.URL)
}

func SetPrivateRepoRepositoryDetailsSubResourceData(m []*models.PrivateRepoRepositoryDetails) (d []*map[string]interface{}) {
	for _, PrivateRepoRepositoryDetailsModel := range m {
		if PrivateRepoRepositoryDetailsModel != nil {
			properties := make(map[string]interface{})
			properties["branch"] = PrivateRepoRepositoryDetailsModel.Branch
			properties["type"] = PrivateRepoRepositoryDetailsModel.Type
			properties["url"] = PrivateRepoRepositoryDetailsModel.URL
			d = append(d, &properties)
		}
	}
	return
}

func PrivateRepoRepositoryDetailsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"branch": {
			Description: `Git branch to track`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"type": {
			Description: `Type of the private repository (PRIVATE_REPO_TYPE_HELM_INDEX, PRIVATE_REPO_TYPE_GIT, PRIVATE_REPO_TYPE_UNSPECIFIED). This is a required parameter`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"url": {
			Description: `Private repository URL (required)`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func GetPrivateRepoRepositoryDetailsPropertyFields() (t []string) {
	return []string{
		"branch",
		"type",
		"url",
	}
}
