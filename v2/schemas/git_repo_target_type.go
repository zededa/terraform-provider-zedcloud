package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GitRepoTargetTypeModel(d *schema.ResourceData) *models.GitRepoTargetType {
	gitRepoTargetType, _ := d.Get("git_repo_target_type").(models.GitRepoTargetType)
	return &gitRepoTargetType
}

func GitRepoTargetTypeModelFromMap(m map[string]interface{}) *models.GitRepoTargetType {
	gitRepoTargetType := m["git_repo_target_type"].(models.GitRepoTargetType)
	return &gitRepoTargetType
}

func SetGitRepoTargetTypeResourceData(d *schema.ResourceData, m *models.GitRepoTargetType) {
}

func SetGitRepoTargetTypeSubResourceData(m []*models.GitRepoTargetType) (d []*map[string]interface{}) {
	for _, GitRepoTargetTypeModel := range m {
		if GitRepoTargetTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func GitRepoTargetTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetGitRepoTargetTypePropertyFields() (t []string) {
	return []string{}
}
