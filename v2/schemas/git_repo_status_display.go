package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GitRepoStatusDisplayModel(d *schema.ResourceData) *models.GitRepoStatusDisplay {
	readyBundleDeployments, _ := d.Get("ready_bundle_deployments").(string)
	state, _ := d.Get("state").(string)
	return &models.GitRepoStatusDisplay{
		ReadyBundleDeployments: readyBundleDeployments,
		State:                  state,
	}
}

func GitRepoStatusDisplayModelFromMap(m map[string]interface{}) *models.GitRepoStatusDisplay {
	readyBundleDeployments := m["ready_bundle_deployments"].(string)
	state := m["state"].(string)
	return &models.GitRepoStatusDisplay{
		ReadyBundleDeployments: readyBundleDeployments,
		State:                  state,
	}
}

func SetGitRepoStatusDisplayResourceData(d *schema.ResourceData, m *models.GitRepoStatusDisplay) {
	d.Set("ready_bundle_deployments", m.ReadyBundleDeployments)
	d.Set("state", m.State)
}

func SetGitRepoStatusDisplaySubResourceData(m []*models.GitRepoStatusDisplay) (d []*map[string]interface{}) {
	for _, GitRepoStatusDisplayModel := range m {
		if GitRepoStatusDisplayModel != nil {
			properties := make(map[string]interface{})
			properties["ready_bundle_deployments"] = GitRepoStatusDisplayModel.ReadyBundleDeployments
			properties["state"] = GitRepoStatusDisplayModel.State
			d = append(d, &properties)
		}
	}
	return
}

func GitRepoStatusDisplaySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ready_bundle_deployments": {
			Description: `Ready bundle deployments count`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"state": {
			Description: `Current state of the GitRepo`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetGitRepoStatusDisplayPropertyFields() (t []string) {
	return []string{
		"ready_bundle_deployments",
		"state",
	}
}
