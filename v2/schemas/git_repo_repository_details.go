package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GitRepoRepositoryDetailsModel(d *schema.ResourceData) *models.GitRepoRepositoryDetails {
	branch, _ := d.Get("branch").(string)
	commit, _ := d.Get("commit").(string)
	var pathList []string
	pathListInterface, pathListIsSet := d.GetOk("path_list")
	if pathListIsSet {
		var items []interface{}
		if listItems, isList := pathListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = pathListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			pathList = append(pathList, v.(string))
		}
	}
	url, _ := d.Get("url").(string)
	return &models.GitRepoRepositoryDetails{
		Branch:   branch,
		Commit:   commit,
		PathList: pathList,
		URL:      &url, // string
	}
}

func GitRepoRepositoryDetailsModelFromMap(m map[string]interface{}) *models.GitRepoRepositoryDetails {
	branch := m["branch"].(string)
	commit := m["commit"].(string)
	var pathList []string
	pathListInterface, pathListIsSet := m["path_list"]
	if pathListIsSet {
		var items []interface{}
		if listItems, isList := pathListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = pathListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			pathList = append(pathList, v.(string))
		}
	}
	url := m["url"].(string)
	return &models.GitRepoRepositoryDetails{
		Branch:   branch,
		Commit:   commit,
		PathList: pathList,
		URL:      &url,
	}
}

func SetGitRepoRepositoryDetailsResourceData(d *schema.ResourceData, m *models.GitRepoRepositoryDetails) {
	d.Set("branch", m.Branch)
	d.Set("commit", m.Commit)
	d.Set("path_list", m.PathList)
	d.Set("url", m.URL)
}

func SetGitRepoRepositoryDetailsSubResourceData(m []*models.GitRepoRepositoryDetails) (d []*map[string]interface{}) {
	for _, GitRepoRepositoryDetailsModel := range m {
		if GitRepoRepositoryDetailsModel != nil {
			properties := make(map[string]interface{})
			properties["branch"] = GitRepoRepositoryDetailsModel.Branch
			properties["commit"] = GitRepoRepositoryDetailsModel.Commit
			properties["path_list"] = GitRepoRepositoryDetailsModel.PathList
			properties["url"] = GitRepoRepositoryDetailsModel.URL
			d = append(d, &properties)
		}
	}
	return
}

func GitRepoRepositoryDetailsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"branch": {
			Description: `Git branch to track`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"commit": {
			Description: `Specific git commit hash`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"path_list": {
			Description: `List of paths within the repository to monitor`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"url": {
			Description: `Git repository URL (required)`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func GetGitRepoRepositoryDetailsPropertyFields() (t []string) {
	return []string{
		"branch",
		"commit",
		"path_list",
		"url",
	}
}
