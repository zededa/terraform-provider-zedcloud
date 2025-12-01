package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func PrivateRepoSpecInfoModel(d *schema.ResourceData) *models.PrivateRepoSpecInfo {
	var auth *models.PrivateRepoAuthInfo // PrivateRepoAuthInfo
	authInterface, authIsSet := d.GetOk("auth")
	if authIsSet && authInterface != nil {
		authMap := authInterface.([]interface{})
		if len(authMap) > 0 {
			auth = PrivateRepoAuthInfoModelFromMap(authMap[0].(map[string]interface{}))
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
	return &models.PrivateRepoSpecInfo{
		Auth:        auth,
		RepoDetails: repoDetails,
	}
}

func PrivateRepoSpecInfoModelFromMap(m map[string]interface{}) *models.PrivateRepoSpecInfo {
	var auth *models.PrivateRepoAuthInfo // PrivateRepoAuthInfo
	authInterface, authIsSet := m["auth"]
	if authIsSet && authInterface != nil {
		authMap := authInterface.([]interface{})
		if len(authMap) > 0 {
			auth = PrivateRepoAuthInfoModelFromMap(authMap[0].(map[string]interface{}))
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
	return &models.PrivateRepoSpecInfo{
		Auth:        auth,
		RepoDetails: repoDetails,
	}
}

func SetPrivateRepoSpecInfoResourceData(d *schema.ResourceData, m *models.PrivateRepoSpecInfo) {
	d.Set("auth", SetPrivateRepoAuthInfoSubResourceData([]*models.PrivateRepoAuthInfo{m.Auth}))
	d.Set("repo_details", SetPrivateRepoRepositoryDetailsSubResourceData([]*models.PrivateRepoRepositoryDetails{m.RepoDetails}))
}

func SetPrivateRepoSpecInfoSubResourceData(m []*models.PrivateRepoSpecInfo) (d []*map[string]interface{}) {
	for _, PrivateRepoSpecInfoModel := range m {
		if PrivateRepoSpecInfoModel != nil {
			properties := make(map[string]interface{})
			properties["auth"] = SetPrivateRepoAuthInfoSubResourceData([]*models.PrivateRepoAuthInfo{PrivateRepoSpecInfoModel.Auth})
			properties["repo_details"] = SetPrivateRepoRepositoryDetailsSubResourceData([]*models.PrivateRepoRepositoryDetails{PrivateRepoSpecInfoModel.RepoDetails})
			d = append(d, &properties)
		}
	}
	return
}

func PrivateRepoSpecInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"auth": {
			Description: `Enhanced authentication configuration`,
			Type:        schema.TypeList, //GoType: PrivateRepoAuthInfo
			Elem: &schema.Resource{
				Schema: PrivateRepoAuthInfoSchema(),
			},
			Optional: true,
		},

		"repo_details": {
			Description: `Repository details and configuration`,
			Type:        schema.TypeList, //GoType: PrivateRepoRepositoryDetails
			Elem: &schema.Resource{
				Schema: PrivateRepoRepositoryDetailsSchema(),
			},
			Optional: true,
		},
	}
}

func GetPrivateRepoSpecInfoPropertyFields() (t []string) {
	return []string{
		"auth",
		"repo_details",
	}
}
