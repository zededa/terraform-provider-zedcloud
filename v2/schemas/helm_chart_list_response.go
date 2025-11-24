package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func HelmChartListResponseModel(d *schema.ResourceData) *models.HelmChartListResponse {
	aPIVersion, _ := d.Get("api_version").(string)
	generated, _ := d.Get("generated").(string)
	repoIdentifier, _ := d.Get("repo_identifier").(string)
	repoName, _ := d.Get("repo_name").(string)
	return &models.HelmChartListResponse{
		APIVersion:     aPIVersion,
		Generated:      generated,
		RepoIdentifier: repoIdentifier,
		RepoName:       repoName,
	}
}

func HelmChartListResponseModelFromMap(m map[string]interface{}) *models.HelmChartListResponse {
	aPIVersion := m["api_version"].(string)
	generated := m["generated"].(string)
	repoIdentifier := m["repo_identifier"].(string)
	repoName := m["repo_name"].(string)
	return &models.HelmChartListResponse{
		APIVersion:     aPIVersion,
		Generated:      generated,
		RepoIdentifier: repoIdentifier,
		RepoName:       repoName,
	}
}

func SetHelmChartListResponseResourceData(d *schema.ResourceData, m *models.HelmChartListResponse) {
	d.Set("api_version", m.APIVersion)
	d.Set("generated", m.Generated)
	d.Set("repo_identifier", m.RepoIdentifier)
	d.Set("repo_name", m.RepoName)
}

func SetHelmChartListResponseSubResourceData(m []*models.HelmChartListResponse) (d []*map[string]interface{}) {
	for _, HelmChartListResponseModel := range m {
		if HelmChartListResponseModel != nil {
			properties := make(map[string]interface{})
			properties["api_version"] = HelmChartListResponseModel.APIVersion
			properties["generated"] = HelmChartListResponseModel.Generated
			properties["repo_identifier"] = HelmChartListResponseModel.RepoIdentifier
			properties["repo_name"] = HelmChartListResponseModel.RepoName
			d = append(d, &properties)
		}
	}
	return
}

func HelmChartListResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"api_version": {
			Description: `API version of the chart index`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"generated": {
			Description: `Timestamp when the chart index was generated`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"repo_identifier": {
			Description: `Identifier for the Helm chart repository`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"repo_name": {
			Description: `Name of the Helm chart repository`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetHelmChartListResponsePropertyFields() (t []string) {
	return []string{
		"api_version",
		"generated",
		"repo_identifier",
		"repo_name",
	}
}
