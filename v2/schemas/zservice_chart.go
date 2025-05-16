package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZserviceChartModel(d *schema.ResourceData) *models.ZserviceChart {
	chartID, _ := d.Get("chart_id").(string)
	chartName, _ := d.Get("chart_name").(string)
	chartVersion, _ := d.Get("chart_version").(string)
	repositoryURL, _ := d.Get("repository_url").(string)
	values, _ := d.Get("values").(string)
	return &models.ZserviceChart{
		ChartID:       chartID,
		ChartName:     chartName,
		ChartVersion:  chartVersion,
		RepositoryURL: repositoryURL,
		Values:        values,
	}
}

func ZserviceChartModelFromMap(m map[string]interface{}) *models.ZserviceChart {
	chartID := m["chart_id"].(string)
	chartName := m["chart_name"].(string)
	chartVersion := m["chart_version"].(string)
	repositoryURL := m["repository_url"].(string)
	values := m["values"].(string)
	return &models.ZserviceChart{
		ChartID:       chartID,
		ChartName:     chartName,
		ChartVersion:  chartVersion,
		RepositoryURL: repositoryURL,
		Values:        values,
	}
}

func SetZserviceChartResourceData(d *schema.ResourceData, m *models.ZserviceChart) {
	d.Set("chart_id", m.ChartID)
	d.Set("chart_name", m.ChartName)
	d.Set("chart_version", m.ChartVersion)
	d.Set("repository_url", m.RepositoryURL)
	d.Set("values", m.Values)
}

func SetZserviceChartSubResourceData(m []*models.ZserviceChart) (d []*map[string]interface{}) {
	for _, ZserviceChartModel := range m {
		if ZserviceChartModel != nil {
			properties := make(map[string]interface{})
			properties["chart_id"] = ZserviceChartModel.ChartID
			properties["chart_name"] = ZserviceChartModel.ChartName
			properties["chart_version"] = ZserviceChartModel.ChartVersion
			properties["repository_url"] = ZserviceChartModel.RepositoryURL
			properties["values"] = ZserviceChartModel.Values
			d = append(d, &properties)
		}
	}
	return
}

func ZserviceChartSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"chart_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"chart_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"chart_version": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"repository_url": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"values": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetZserviceChartPropertyFields() (t []string) {
	return []string{
		"chart_id",
		"chart_name",
		"chart_version",
		"repository_url",
		"values",
	}
}
