package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func HelmChartInfoResponseModel(d *schema.ResourceData) *models.HelmChartInfoResponse {
	appReadme, _ := d.Get("app_readme").(string)
	var chart *models.HelmChartDetail // HelmChartDetail
	chartInterface, chartIsSet := d.GetOk("chart")
	if chartIsSet && chartInterface != nil {
		chartMap := chartInterface.([]interface{})
		if len(chartMap) > 0 {
			chart = HelmChartDetailModelFromMap(chartMap[0].(map[string]interface{}))
		}
	}
	questions, _ := d.Get("questions").(any) // any

	readme, _ := d.Get("readme").(string)
	repoIdentifier, _ := d.Get("repo_identifier").(string)
	values, _ := d.Get("values").(any) // any

	return &models.HelmChartInfoResponse{
		AppReadme:      appReadme,
		Chart:          chart,
		Questions:      questions,
		Readme:         readme,
		RepoIdentifier: repoIdentifier,
		Values:         values,
	}
}

func HelmChartInfoResponseModelFromMap(m map[string]interface{}) *models.HelmChartInfoResponse {
	appReadme := m["app_readme"].(string)
	var chart *models.HelmChartDetail // HelmChartDetail
	chartInterface, chartIsSet := m["chart"]
	if chartIsSet && chartInterface != nil {
		chartMap := chartInterface.([]interface{})
		if len(chartMap) > 0 {
			chart = HelmChartDetailModelFromMap(chartMap[0].(map[string]interface{}))
		}
	}
	//
	questions := m["questions"].(any)

	readme := m["readme"].(string)
	repoIdentifier := m["repo_identifier"].(string)
	values := m["values"].(any)

	return &models.HelmChartInfoResponse{
		AppReadme:      appReadme,
		Chart:          chart,
		Questions:      questions,
		Readme:         readme,
		RepoIdentifier: repoIdentifier,
		Values:         values,
	}
}

func SetHelmChartInfoResponseResourceData(d *schema.ResourceData, m *models.HelmChartInfoResponse) {
	d.Set("app_readme", m.AppReadme)
	d.Set("chart", SetHelmChartDetailSubResourceData([]*models.HelmChartDetail{m.Chart}))
	d.Set("questions", m.Questions)
	d.Set("readme", m.Readme)
	d.Set("repo_identifier", m.RepoIdentifier)
	d.Set("values", m.Values)
}

func SetHelmChartInfoResponseSubResourceData(m []*models.HelmChartInfoResponse) (d []*map[string]interface{}) {
	for _, HelmChartInfoResponseModel := range m {
		if HelmChartInfoResponseModel != nil {
			properties := make(map[string]interface{})
			properties["app_readme"] = HelmChartInfoResponseModel.AppReadme
			properties["chart"] = SetHelmChartDetailSubResourceData([]*models.HelmChartDetail{HelmChartInfoResponseModel.Chart})
			properties["questions"] = HelmChartInfoResponseModel.Questions
			properties["readme"] = HelmChartInfoResponseModel.Readme
			properties["repo_identifier"] = HelmChartInfoResponseModel.RepoIdentifier
			properties["values"] = HelmChartInfoResponseModel.Values
			d = append(d, &properties)
		}
	}
	return
}

func HelmChartInfoResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_readme": {
			Description: `Application-specific README content`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"chart": {
			Description: `Detailed chart metadata and information`,
			Type:        schema.TypeList, //GoType: HelmChartDetail
			Elem: &schema.Resource{
				Schema: HelmChartDetailSchema(),
			},
			Optional: true,
		},

		"questions": {
			Description: `Configuration questions for the Helm chart`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"readme": {
			Description: `README content for the Helm chart`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"repo_identifier": {
			Description: `Identifier for the Helm chart repository`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"values": {
			Description: `Default values for the Helm chart`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetHelmChartInfoResponsePropertyFields() (t []string) {
	return []string{
		"app_readme",
		"chart",
		"questions",
		"readme",
		"repo_identifier",
		"values",
	}
}
