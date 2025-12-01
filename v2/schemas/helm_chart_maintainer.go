package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func HelmChartMaintainerModel(d *schema.ResourceData) *models.HelmChartMaintainer {
	name, _ := d.Get("name").(string)
	url, _ := d.Get("url").(string)
	return &models.HelmChartMaintainer{
		Name: name,
		URL:  url,
	}
}

func HelmChartMaintainerModelFromMap(m map[string]interface{}) *models.HelmChartMaintainer {
	name := m["name"].(string)
	url := m["url"].(string)
	return &models.HelmChartMaintainer{
		Name: name,
		URL:  url,
	}
}

func SetHelmChartMaintainerResourceData(d *schema.ResourceData, m *models.HelmChartMaintainer) {
	d.Set("name", m.Name)
	d.Set("url", m.URL)
}

func SetHelmChartMaintainerSubResourceData(m []*models.HelmChartMaintainer) (d []*map[string]interface{}) {
	for _, HelmChartMaintainerModel := range m {
		if HelmChartMaintainerModel != nil {
			properties := make(map[string]interface{})
			properties["name"] = HelmChartMaintainerModel.Name
			properties["url"] = HelmChartMaintainerModel.URL
			d = append(d, &properties)
		}
	}
	return
}

func HelmChartMaintainerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Description: `Name of the maintainer. Example: 'Zededa Team'`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"url": {
			Description: `URL or contact information for the maintainer. Example: 'https://github.com/bitnami/charts'`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetHelmChartMaintainerPropertyFields() (t []string) {
	return []string{
		"name",
		"url",
	}
}
