package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GitRepoDataModel(d *schema.ResourceData) *models.GitRepoData {
	var auth *models.GitRepoAuthConfig // GitRepoAuthConfig
	authInterface, authIsSet := d.GetOk("auth")
	if authIsSet && authInterface != nil {
		authMap := authInterface.([]interface{})
		if len(authMap) > 0 {
			auth = GitRepoAuthConfigModelFromMap(authMap[0].(map[string]interface{}))
		}
	}
	projectID, _ := d.Get("project_id").(string)
	var repoDetails *models.GitRepoRepositoryDetails // GitRepoRepositoryDetails
	repoDetailsInterface, repoDetailsIsSet := d.GetOk("repo_details")
	if repoDetailsIsSet && repoDetailsInterface != nil {
		repoDetailsMap := repoDetailsInterface.([]interface{})
		if len(repoDetailsMap) > 0 {
			repoDetails = GitRepoRepositoryDetailsModelFromMap(repoDetailsMap[0].(map[string]interface{}))
		}
	}
	var target *models.GitRepoTargetConfig // GitRepoTargetConfig
	targetInterface, targetIsSet := d.GetOk("target")
	if targetIsSet && targetInterface != nil {
		targetMap := targetInterface.([]interface{})
		if len(targetMap) > 0 {
			target = GitRepoTargetConfigModelFromMap(targetMap[0].(map[string]interface{}))
		}
	}
	return &models.GitRepoData{
		Auth:        auth,
		ProjectID:   &projectID, // string
		RepoDetails: repoDetails,
		Target:      target,
	}
}

func GitRepoDataModelFromMap(m map[string]interface{}) *models.GitRepoData {
	var auth *models.GitRepoAuthConfig // GitRepoAuthConfig
	authInterface, authIsSet := m["auth"]
	if authIsSet && authInterface != nil {
		authMap := authInterface.([]interface{})
		if len(authMap) > 0 {
			auth = GitRepoAuthConfigModelFromMap(authMap[0].(map[string]interface{}))
		}
	}
	//
	projectID := m["project_id"].(string)
	var repoDetails *models.GitRepoRepositoryDetails // GitRepoRepositoryDetails
	repoDetailsInterface, repoDetailsIsSet := m["repo_details"]
	if repoDetailsIsSet && repoDetailsInterface != nil {
		repoDetailsMap := repoDetailsInterface.([]interface{})
		if len(repoDetailsMap) > 0 {
			repoDetails = GitRepoRepositoryDetailsModelFromMap(repoDetailsMap[0].(map[string]interface{}))
		}
	}
	//
	var target *models.GitRepoTargetConfig // GitRepoTargetConfig
	targetInterface, targetIsSet := m["target"]
	if targetIsSet && targetInterface != nil {
		targetMap := targetInterface.([]interface{})
		if len(targetMap) > 0 {
			target = GitRepoTargetConfigModelFromMap(targetMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.GitRepoData{
		Auth:        auth,
		ProjectID:   &projectID,
		RepoDetails: repoDetails,
		Target:      target,
	}
}

func SetGitRepoDataResourceData(d *schema.ResourceData, m *models.GitRepoData) {
	d.Set("auth", SetGitRepoAuthConfigSubResourceData([]*models.GitRepoAuthConfig{m.Auth}))
	d.Set("project_id", m.ProjectID)
	d.Set("repo_details", SetGitRepoRepositoryDetailsSubResourceData([]*models.GitRepoRepositoryDetails{m.RepoDetails}))
	d.Set("target", SetGitRepoTargetConfigSubResourceData([]*models.GitRepoTargetConfig{m.Target}))
}

func SetGitRepoDataSubResourceData(m []*models.GitRepoData) (d []*map[string]interface{}) {
	for _, GitRepoDataModel := range m {
		if GitRepoDataModel != nil {
			properties := make(map[string]interface{})
			properties["auth"] = SetGitRepoAuthConfigSubResourceData([]*models.GitRepoAuthConfig{GitRepoDataModel.Auth})
			properties["project_id"] = GitRepoDataModel.ProjectID
			properties["repo_details"] = SetGitRepoRepositoryDetailsSubResourceData([]*models.GitRepoRepositoryDetails{GitRepoDataModel.RepoDetails})
			properties["target"] = SetGitRepoTargetConfigSubResourceData([]*models.GitRepoTargetConfig{GitRepoDataModel.Target})
			d = append(d, &properties)
		}
	}
	return
}

func GitRepoDataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"auth": {
			Description: `Authentication configuration`,
			Type:        schema.TypeList, //GoType: GitRepoAuthConfig
			Elem: &schema.Resource{
				Schema: GitRepoAuthConfigSchema(),
			},
			Optional: true,
		},

		"project_id": {
			Description: `Project ID where the GitRepo belongs (required)`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"repo_details": {
			Description: `Repository details and configuration (required)`,
			Type:        schema.TypeList, //GoType: GitRepoRepositoryDetails
			Elem: &schema.Resource{
				Schema: GitRepoRepositoryDetailsSchema(),
			},
			Required: true,
		},

		"target": {
			Description: `Target deployment configuration (required)`,
			Type:        schema.TypeList, //GoType: GitRepoTargetConfig
			Elem: &schema.Resource{
				Schema: GitRepoTargetConfigSchema(),
			},
			Required: true,
		},
	}
}

func GetGitRepoDataPropertyFields() (t []string) {
	return []string{
		"auth",
		"project_id",
		"repo_details",
		"target",
	}
}
