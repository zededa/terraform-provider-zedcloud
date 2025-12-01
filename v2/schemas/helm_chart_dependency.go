package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func HelmChartDependencyModel(d *schema.ResourceData) *models.HelmChartDependency {
	name, _ := d.Get("name").(string)
	repository, _ := d.Get("repository").(string)
	var tags []string
	tagsInterface, tagsIsSet := d.GetOk("tags")
	if tagsIsSet {
		var items []interface{}
		if listItems, isList := tagsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = tagsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			tags = append(tags, v.(string))
		}
	}
	version, _ := d.Get("version").(string)
	return &models.HelmChartDependency{
		Name:       name,
		Repository: repository,
		Tags:       tags,
		Version:    version,
	}
}

func HelmChartDependencyModelFromMap(m map[string]interface{}) *models.HelmChartDependency {
	name := m["name"].(string)
	repository := m["repository"].(string)
	var tags []string
	tagsInterface, tagsIsSet := m["tags"]
	if tagsIsSet {
		var items []interface{}
		if listItems, isList := tagsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = tagsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			tags = append(tags, v.(string))
		}
	}
	version := m["version"].(string)
	return &models.HelmChartDependency{
		Name:       name,
		Repository: repository,
		Tags:       tags,
		Version:    version,
	}
}

func SetHelmChartDependencyResourceData(d *schema.ResourceData, m *models.HelmChartDependency) {
	d.Set("name", m.Name)
	d.Set("repository", m.Repository)
	d.Set("tags", m.Tags)
	d.Set("version", m.Version)
}

func SetHelmChartDependencySubResourceData(m []*models.HelmChartDependency) (d []*map[string]interface{}) {
	for _, HelmChartDependencyModel := range m {
		if HelmChartDependencyModel != nil {
			properties := make(map[string]interface{})
			properties["name"] = HelmChartDependencyModel.Name
			properties["repository"] = HelmChartDependencyModel.Repository
			properties["tags"] = HelmChartDependencyModel.Tags
			properties["version"] = HelmChartDependencyModel.Version
			d = append(d, &properties)
		}
	}
	return
}

func HelmChartDependencySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Description: `Name of the dependency chart. Example: 'postgresql'`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"repository": {
			Description: `Repository URL for the dependency. Example: 'https://charts.bitnami.com/bitnami'`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"tags": {
			Description: `Tags associated with the dependency. Example: ['database', 'postgresql']`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"version": {
			Description: `Version constraint for the dependency. Example: '^12.0.0'`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetHelmChartDependencyPropertyFields() (t []string) {
	return []string{
		"name",
		"repository",
		"tags",
		"version",
	}
}
