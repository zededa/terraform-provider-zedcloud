package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func PrivateRepoSpecModel(d *schema.ResourceData) *models.PrivateRepoSpec {
	var auth *models.PrivateRepoAuthConfig // PrivateRepoAuthConfig
	authInterface, authIsSet := d.GetOk("auth")
	if authIsSet && authInterface != nil {
		authMap := authInterface.([]interface{})
		if len(authMap) > 0 {
			auth = PrivateRepoAuthConfigModelFromMap(authMap[0].(map[string]interface{}))
		}
	}
	var repoDetails *models.PrivateRepoRepositoryDetails // PrivateRepoRepositoryDetails
	repoDetailsInterface, repoDetailsIsSet := d.GetOk("repo_details")
	if repoDetailsIsSet && repoDetailsInterface != nil {
		repoDetailsMap := repoDetailsInterface.([]interface{})
		if len(repoDetailsMap) > 0 {
			repoDetails = PrivateRepoRepositoryDetailsModelFromMap(repoDetailsMap[0].(map[string]interface{}))
		}
	}
	return &models.PrivateRepoSpec{
		Auth:        auth,
		RepoDetails: repoDetails,
	}
}

func PrivateRepoSpecModelFromMap(m map[string]interface{}) *models.PrivateRepoSpec {
	var auth *models.PrivateRepoAuthConfig // PrivateRepoAuthConfig
	authInterface, authIsSet := m["auth"]
	if authIsSet && authInterface != nil {
		authMap := authInterface.([]interface{})
		if len(authMap) > 0 {
			auth = PrivateRepoAuthConfigModelFromMap(authMap[0].(map[string]interface{}))
		}
	}
	//
	var repoDetails *models.PrivateRepoRepositoryDetails // PrivateRepoRepositoryDetails
	repoDetailsInterface, repoDetailsIsSet := m["repo_details"]
	if repoDetailsIsSet && repoDetailsInterface != nil {
		repoDetailsMap := repoDetailsInterface.([]interface{})
		if len(repoDetailsMap) > 0 {
			repoDetails = PrivateRepoRepositoryDetailsModelFromMap(repoDetailsMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.PrivateRepoSpec{
		Auth:        auth,
		RepoDetails: repoDetails,
	}
}

func SetPrivateRepoSpecResourceData(d *schema.ResourceData, m *models.PrivateRepoSpec) {
	d.Set("auth", SetPrivateRepoAuthConfigSubResourceData([]*models.PrivateRepoAuthConfig{m.Auth}))
	d.Set("repo_details", SetPrivateRepoRepositoryDetailsSubResourceData([]*models.PrivateRepoRepositoryDetails{m.RepoDetails}))
}

func SetPrivateRepoSpecSubResourceData(m []*models.PrivateRepoSpec) (d []*map[string]interface{}) {
	for _, PrivateRepoSpecModel := range m {
		if PrivateRepoSpecModel != nil {
			properties := make(map[string]interface{})
			properties["auth"] = SetPrivateRepoAuthConfigSubResourceData([]*models.PrivateRepoAuthConfig{PrivateRepoSpecModel.Auth})
			properties["repo_details"] = SetPrivateRepoRepositoryDetailsSubResourceData([]*models.PrivateRepoRepositoryDetails{PrivateRepoSpecModel.RepoDetails})
			d = append(d, &properties)
		}
	}
	return
}

func PrivateRepoSpecSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"auth": {
			Description: `Authentication configuration`,
			Type:        schema.TypeList, //GoType: PrivateRepoAuthConfig
			Elem: &schema.Resource{
				Schema: PrivateRepoAuthConfigSchema(),
			},
			Optional: true,
		},

		"repo_details": {
			Description: `Repository details and configuration (required)`,
			Type:        schema.TypeList, //GoType: PrivateRepoRepositoryDetails
			Elem: &schema.Resource{
				Schema: PrivateRepoRepositoryDetailsSchema(),
			},
			Required: true,
		},
	}
}

func GetPrivateRepoSpecPropertyFields() (t []string) {
	return []string{
		"auth",
		"repo_details",
	}
}
