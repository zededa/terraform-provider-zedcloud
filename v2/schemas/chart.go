package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ChartModel(d *schema.ResourceData) *models.Chart {
	customValues, _ := d.Get("custom_values").(any) // any

	name, _ := d.Get("name").(string)
	repoIdentifier, _ := d.Get("repo_identifier").(string)
	version, _ := d.Get("version").(string)
	return &models.Chart{
		CustomValues:   customValues,
		Name:           &name, // string
		RepoIdentifier: repoIdentifier,
		Version:        &version, // string
	}
}

func ChartModelFromMap(m map[string]interface{}) *models.Chart {
	customValues := m["custom_values"].(any)

	name := m["name"].(string)
	repoIdentifier := m["repo_identifier"].(string)
	version := m["version"].(string)
	return &models.Chart{
		CustomValues:   customValues,
		Name:           &name,
		RepoIdentifier: repoIdentifier,
		Version:        &version,
	}
}

func SetChartResourceData(d *schema.ResourceData, m *models.Chart) {
	d.Set("custom_values", m.CustomValues)
	d.Set("name", m.Name)
	d.Set("repo_identifier", m.RepoIdentifier)
	d.Set("version", m.Version)
}

func SetChartSubResourceData(m []*models.Chart) (d []*map[string]interface{}) {
	for _, ChartModel := range m {
		if ChartModel != nil {
			properties := make(map[string]interface{})
			properties["custom_values"] = ChartModel.CustomValues
			properties["name"] = ChartModel.Name
			properties["repo_identifier"] = ChartModel.RepoIdentifier
			properties["version"] = ChartModel.Version
			d = append(d, &properties)
		}
	}
	return
}

func ChartSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"custom_values": {
			Description: `Custom values for the Helm chart. Example: {'replicaCount': 3, 'service': {'type': 'LoadBalancer', 'port': 80}}`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Description: `Name of the Helm chart. Example: 'nginx'`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"repo_identifier": {
			Description: `Repository identifier for the Helm chart. Example: 'bitnami'`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"version": {
			Description: `Version of the Helm chart. Example: '15.4.0'`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func GetChartPropertyFields() (t []string) {
	return []string{
		"custom_values",
		"name",
		"repo_identifier",
		"version",
	}
}
