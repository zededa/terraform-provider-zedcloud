package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func HelmChartDetailModel(d *schema.ResourceData) *models.HelmChartDetail {
	annotations := map[string]string{}
	annotationsInterface, annotationsIsSet := d.GetOk("annotations")
	if annotationsIsSet {
		annotationsMap := annotationsInterface.(map[string]interface{})
		for k, v := range annotationsMap {
			if v == nil {
				continue
			}
			annotations[k] = v.(string)
		}
	}

	aPIVersion, _ := d.Get("api_version").(string)
	appVersion, _ := d.Get("app_version").(string)
	var dependencies []*models.HelmChartDependency // []*HelmChartDependency
	dependenciesInterface, dependenciesIsSet := d.GetOk("dependencies")
	if dependenciesIsSet {
		var items []interface{}
		if listItems, isList := dependenciesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = dependenciesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := HelmChartDependencyModelFromMap(v.(map[string]interface{}))
			dependencies = append(dependencies, m)
		}
	}
	description, _ := d.Get("description").(string)
	home, _ := d.Get("home").(string)
	icon, _ := d.Get("icon").(string)
	var keywords []string
	keywordsInterface, keywordsIsSet := d.GetOk("keywords")
	if keywordsIsSet {
		var items []interface{}
		if listItems, isList := keywordsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = keywordsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			keywords = append(keywords, v.(string))
		}
	}
	var maintainers []*models.HelmChartMaintainer // []*HelmChartMaintainer
	maintainersInterface, maintainersIsSet := d.GetOk("maintainers")
	if maintainersIsSet {
		var items []interface{}
		if listItems, isList := maintainersInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = maintainersInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := HelmChartMaintainerModelFromMap(v.(map[string]interface{}))
			maintainers = append(maintainers, m)
		}
	}
	name, _ := d.Get("name").(string)
	var sources []string
	sourcesInterface, sourcesIsSet := d.GetOk("sources")
	if sourcesIsSet {
		var items []interface{}
		if listItems, isList := sourcesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = sourcesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			sources = append(sources, v.(string))
		}
	}
	version, _ := d.Get("version").(string)
	return &models.HelmChartDetail{
		Annotations:  annotations,
		APIVersion:   aPIVersion,
		AppVersion:   appVersion,
		Dependencies: dependencies,
		Description:  description,
		Home:         home,
		Icon:         icon,
		Keywords:     keywords,
		Maintainers:  maintainers,
		Name:         name,
		Sources:      sources,
		Version:      version,
	}
}

func HelmChartDetailModelFromMap(m map[string]interface{}) *models.HelmChartDetail {
	annotations := map[string]string{}
	annotationsInterface, annotationsIsSet := m["annotations"]
	if annotationsIsSet {
		annotationsMap := annotationsInterface.(map[string]interface{})
		for k, v := range annotationsMap {
			if v == nil {
				continue
			}
			annotations[k] = v.(string)
		}
	}

	aPIVersion := m["api_version"].(string)
	appVersion := m["app_version"].(string)
	var dependencies []*models.HelmChartDependency // []*HelmChartDependency
	dependenciesInterface, dependenciesIsSet := m["dependencies"]
	if dependenciesIsSet {
		var items []interface{}
		if listItems, isList := dependenciesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = dependenciesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := HelmChartDependencyModelFromMap(v.(map[string]interface{}))
			dependencies = append(dependencies, m)
		}
	}
	description := m["description"].(string)
	home := m["home"].(string)
	icon := m["icon"].(string)
	var keywords []string
	keywordsInterface, keywordsIsSet := m["keywords"]
	if keywordsIsSet {
		var items []interface{}
		if listItems, isList := keywordsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = keywordsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			keywords = append(keywords, v.(string))
		}
	}
	var maintainers []*models.HelmChartMaintainer // []*HelmChartMaintainer
	maintainersInterface, maintainersIsSet := m["maintainers"]
	if maintainersIsSet {
		var items []interface{}
		if listItems, isList := maintainersInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = maintainersInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := HelmChartMaintainerModelFromMap(v.(map[string]interface{}))
			maintainers = append(maintainers, m)
		}
	}
	name := m["name"].(string)
	var sources []string
	sourcesInterface, sourcesIsSet := m["sources"]
	if sourcesIsSet {
		var items []interface{}
		if listItems, isList := sourcesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = sourcesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			sources = append(sources, v.(string))
		}
	}
	version := m["version"].(string)
	return &models.HelmChartDetail{
		Annotations:  annotations,
		APIVersion:   aPIVersion,
		AppVersion:   appVersion,
		Dependencies: dependencies,
		Description:  description,
		Home:         home,
		Icon:         icon,
		Keywords:     keywords,
		Maintainers:  maintainers,
		Name:         name,
		Sources:      sources,
		Version:      version,
	}
}

func SetHelmChartDetailResourceData(d *schema.ResourceData, m *models.HelmChartDetail) {
	d.Set("annotations", m.Annotations)
	d.Set("api_version", m.APIVersion)
	d.Set("app_version", m.AppVersion)
	d.Set("dependencies", SetHelmChartDependencySubResourceData(m.Dependencies))
	d.Set("description", m.Description)
	d.Set("home", m.Home)
	d.Set("icon", m.Icon)
	d.Set("keywords", m.Keywords)
	d.Set("maintainers", SetHelmChartMaintainerSubResourceData(m.Maintainers))
	d.Set("name", m.Name)
	d.Set("sources", m.Sources)
	d.Set("version", m.Version)
}

func SetHelmChartDetailSubResourceData(m []*models.HelmChartDetail) (d []*map[string]interface{}) {
	for _, HelmChartDetailModel := range m {
		if HelmChartDetailModel != nil {
			properties := make(map[string]interface{})
			properties["annotations"] = HelmChartDetailModel.Annotations
			properties["api_version"] = HelmChartDetailModel.APIVersion
			properties["app_version"] = HelmChartDetailModel.AppVersion
			properties["dependencies"] = SetHelmChartDependencySubResourceData(HelmChartDetailModel.Dependencies)
			properties["description"] = HelmChartDetailModel.Description
			properties["home"] = HelmChartDetailModel.Home
			properties["icon"] = HelmChartDetailModel.Icon
			properties["keywords"] = HelmChartDetailModel.Keywords
			properties["maintainers"] = SetHelmChartMaintainerSubResourceData(HelmChartDetailModel.Maintainers)
			properties["name"] = HelmChartDetailModel.Name
			properties["sources"] = HelmChartDetailModel.Sources
			properties["version"] = HelmChartDetailModel.Version
			d = append(d, &properties)
		}
	}
	return
}

func HelmChartDetailSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"annotations": {
			Description: `Chart annotations including category, images, and licenses`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"api_version": {
			Description: `Helm chart API version`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"app_version": {
			Description: `Version of the application packaged in the chart`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"dependencies": {
			Description: `List of chart dependencies`,
			Type:        schema.TypeList, //GoType: []*HelmChartDependency
			Elem: &schema.Resource{
				Schema: HelmChartDependencySchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"description": {
			Description: `Description of the Helm chart`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"home": {
			Description: `Home page URL for the chart`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"icon": {
			Description: `Icon URL for the chart`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"keywords": {
			Description: `Keywords associated with the chart`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"maintainers": {
			Description: `List of chart maintainers`,
			Type:        schema.TypeList, //GoType: []*HelmChartMaintainer
			Elem: &schema.Resource{
				Schema: HelmChartMaintainerSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"name": {
			Description: `Name of the Helm chart`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"sources": {
			Description: `List of source code URLs for the chart`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"version": {
			Description: `Version of the Helm chart`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetHelmChartDetailPropertyFields() (t []string) {
	return []string{
		"annotations",
		"api_version",
		"app_version",
		"dependencies",
		"description",
		"home",
		"icon",
		"keywords",
		"maintainers",
		"name",
		"sources",
		"version",
	}
}
