package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GitRepoListResponseModel(d *schema.ResourceData) *models.GitRepoListResponse {
	var gitrepos []*models.GitRepoGetResponse // []*GitRepoGetResponse
	gitreposInterface, gitreposIsSet := d.GetOk("gitrepos")
	if gitreposIsSet {
		var items []interface{}
		if listItems, isList := gitreposInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = gitreposInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := GitRepoGetResponseModelFromMap(v.(map[string]interface{}))
			gitrepos = append(gitrepos, m)
		}
	}
	var stateSummary *models.GitRepoListStatusSummary // GitRepoListStatusSummary
	stateSummaryInterface, stateSummaryIsSet := d.GetOk("state_summary")
	if stateSummaryIsSet && stateSummaryInterface != nil {
		stateSummaryMap := stateSummaryInterface.([]interface{})
		if len(stateSummaryMap) > 0 {
			stateSummary = GitRepoListStatusSummaryModelFromMap(stateSummaryMap[0].(map[string]interface{}))
		}
	}
	totalCountInt, _ := d.Get("total_count").(int)
	totalCount := int32(totalCountInt)
	return &models.GitRepoListResponse{
		Gitrepos:     gitrepos,
		StateSummary: stateSummary,
		TotalCount:   totalCount,
	}
}

func GitRepoListResponseModelFromMap(m map[string]interface{}) *models.GitRepoListResponse {
	var gitrepos []*models.GitRepoGetResponse // []*GitRepoGetResponse
	gitreposInterface, gitreposIsSet := m["gitrepos"]
	if gitreposIsSet {
		var items []interface{}
		if listItems, isList := gitreposInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = gitreposInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := GitRepoGetResponseModelFromMap(v.(map[string]interface{}))
			gitrepos = append(gitrepos, m)
		}
	}
	var stateSummary *models.GitRepoListStatusSummary // GitRepoListStatusSummary
	stateSummaryInterface, stateSummaryIsSet := m["state_summary"]
	if stateSummaryIsSet && stateSummaryInterface != nil {
		stateSummaryMap := stateSummaryInterface.([]interface{})
		if len(stateSummaryMap) > 0 {
			stateSummary = GitRepoListStatusSummaryModelFromMap(stateSummaryMap[0].(map[string]interface{}))
		}
	}
	//
	totalCount := int32(m["total_count"].(int)) // int32
	return &models.GitRepoListResponse{
		Gitrepos:     gitrepos,
		StateSummary: stateSummary,
		TotalCount:   totalCount,
	}
}

func SetGitRepoListResponseResourceData(d *schema.ResourceData, m *models.GitRepoListResponse) {
	d.Set("gitrepos", SetGitRepoGetResponseSubResourceData(m.Gitrepos))
	d.Set("state_summary", SetGitRepoListStatusSummarySubResourceData([]*models.GitRepoListStatusSummary{m.StateSummary}))
	d.Set("total_count", m.TotalCount)
}

func SetGitRepoListResponseSubResourceData(m []*models.GitRepoListResponse) (d []*map[string]interface{}) {
	for _, GitRepoListResponseModel := range m {
		if GitRepoListResponseModel != nil {
			properties := make(map[string]interface{})
			properties["gitrepos"] = SetGitRepoGetResponseSubResourceData(GitRepoListResponseModel.Gitrepos)
			properties["state_summary"] = SetGitRepoListStatusSummarySubResourceData([]*models.GitRepoListStatusSummary{GitRepoListResponseModel.StateSummary})
			properties["total_count"] = GitRepoListResponseModel.TotalCount
			d = append(d, &properties)
		}
	}
	return
}

func GitRepoListResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"gitrepos": {
			Description: `List of GitRepo entries`,
			Type:        schema.TypeList, //GoType: []*GitRepoGetResponse
			Elem: &schema.Resource{
				Schema: GitRepoGetResponseSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"state_summary": {
			Description: `Summary of GitRepo statuses`,
			Type:        schema.TypeList, //GoType: GitRepoListStatusSummary
			Elem: &schema.Resource{
				Schema: GitRepoListStatusSummarySchema(),
			},
			Optional: true,
		},

		"total_count": {
			Description: `Total count of GitRepos`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetGitRepoListResponsePropertyFields() (t []string) {
	return []string{
		"gitrepos",
		"state_summary",
		"total_count",
	}
}
