package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AppFilterModel(d *schema.ResourceData) *models.AppFilter {
	var appCategory *models.AppCategory // AppCategory
	appCategoryInterface, appCategoryIsSet := d.GetOk("app_category")
	if appCategoryIsSet {
		appCategoryModel := appCategoryInterface.(string)
		appCategory = models.NewAppCategory(models.AppCategory(appCategoryModel))
	}
	var appType *models.AppType // AppType
	appTypeInterface, appTypeIsSet := d.GetOk("app_type")
	if appTypeIsSet {
		appTypeModel := appTypeInterface.(string)
		appType = models.NewAppType(models.AppType(appTypeModel))
	}
	var categories []*models.AppCategory // []*AppCategory
	categoriesInterface, categoriesIsSet := d.GetOk("categories")
	if categoriesIsSet {
		var items []interface{}
		if listItems, isList := categoriesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = categoriesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AppCategoryModelFromMap(v.(map[string]interface{}))
			categories = append(categories, m)
		}
	}
	category, _ := d.Get("category").(string)
	var deploymentType *models.DeploymentType // DeploymentType
	deploymentTypeInterface, deploymentTypeIsSet := d.GetOk("deployment_type")
	if deploymentTypeIsSet {
		deploymentTypeModel := deploymentTypeInterface.(string)
		deploymentType = models.NewDeploymentType(models.DeploymentType(deploymentTypeModel))
	}
	namePattern, _ := d.Get("name_pattern").(string)
	var originType *models.Origin // Origin
	originTypeInterface, originTypeIsSet := d.GetOk("origin_type")
	if originTypeIsSet {
		originTypeModel := originTypeInterface.(string)
		originType = models.NewOrigin(models.Origin(originTypeModel))
	}
	return &models.AppFilter{
		AppCategory:    appCategory,
		AppType:        appType,
		Categories:     categories,
		Category:       category,
		DeploymentType: deploymentType,
		NamePattern:    namePattern,
		OriginType:     originType,
	}
}

func AppFilterModelFromMap(m map[string]interface{}) *models.AppFilter {
	var appCategory *models.AppCategory // AppCategory
	appCategoryInterface, appCategoryIsSet := m["app_category"]
	if appCategoryIsSet {
		appCategoryModel := appCategoryInterface.(string)
		appCategory = models.NewAppCategory(models.AppCategory(appCategoryModel))
	}
	var appType *models.AppType // AppType
	appTypeInterface, appTypeIsSet := m["app_type"]
	if appTypeIsSet {
		appTypeModel := appTypeInterface.(string)
		appType = models.NewAppType(models.AppType(appTypeModel))
	}
	var categories []*models.AppCategory // []*AppCategory
	categoriesInterface, categoriesIsSet := m["categories"]
	if categoriesIsSet {
		var items []interface{}
		if listItems, isList := categoriesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = categoriesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AppCategoryModelFromMap(v.(map[string]interface{}))
			categories = append(categories, m)
		}
	}
	category := m["category"].(string)
	var deploymentType *models.DeploymentType // DeploymentType
	deploymentTypeInterface, deploymentTypeIsSet := m["deployment_type"]
	if deploymentTypeIsSet {
		deploymentTypeModel := deploymentTypeInterface.(string)
		deploymentType = models.NewDeploymentType(models.DeploymentType(deploymentTypeModel))
	}
	namePattern := m["name_pattern"].(string)
	var originType *models.Origin // Origin
	originTypeInterface, originTypeIsSet := m["origin_type"]
	if originTypeIsSet {
		originTypeModel := originTypeInterface.(string)
		originType = models.NewOrigin(models.Origin(originTypeModel))
	}
	return &models.AppFilter{
		AppCategory:    appCategory,
		AppType:        appType,
		Categories:     categories,
		Category:       category,
		DeploymentType: deploymentType,
		NamePattern:    namePattern,
		OriginType:     originType,
	}
}

// Update the underlying AppFilter resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppFilterResourceData(d *schema.ResourceData, m *models.AppFilter) {
	d.Set("app_category", m.AppCategory)
	d.Set("app_type", m.AppType)
	d.Set("categories", SetAppCategorySubResourceData(m.Categories))
	d.Set("category", m.Category)
	d.Set("deployment_type", m.DeploymentType)
	d.Set("name_pattern", m.NamePattern)
	d.Set("origin_type", m.OriginType)
}

// Iterate through and update the AppFilter resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppFilterSubResourceData(m []*models.AppFilter) (d []*map[string]interface{}) {
	for _, AppFilterModel := range m {
		if AppFilterModel != nil {
			properties := make(map[string]interface{})
			properties["app_category"] = AppFilterModel.AppCategory
			properties["app_type"] = AppFilterModel.AppType
			properties["categories"] = SetAppCategorySubResourceData(AppFilterModel.Categories)
			properties["category"] = AppFilterModel.Category
			properties["deployment_type"] = AppFilterModel.DeploymentType
			properties["name_pattern"] = AppFilterModel.NamePattern
			properties["origin_type"] = AppFilterModel.OriginType
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppFilter resource defined in the Terraform configuration
func AppFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_category": {
			Description: `category type of the bundle`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"app_type": {
			Description: `app type, eg: vm, container, module`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"categories": {
			Description: `category types of the bundle`,
			Type:        schema.TypeList, //GoType: []*AppCategory
			Elem: &schema.Resource{
				Schema: AppCategorySchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"category": {
			Description: `category type of the bundle`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"deployment_type": {
			Description: `type of deployment for the app, eg: azure, k3s, standalone`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name_pattern": {
			Description: `Search * namePattern * in name field to filter records`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"origin_type": {
			Description: `origin of bundle`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the AppFilter resource
func GetAppFilterPropertyFields() (t []string) {
	return []string{
		"app_category",
		"app_type",
		"categories",
		"category",
		"deployment_type",
		"name_pattern",
		"origin_type",
	}
}
