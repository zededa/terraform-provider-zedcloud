package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func VMManifestSummaryModel(d *schema.ResourceData) *models.VMManifestSummary {
	acKind, _ := d.Get("ac_kind").(string)
	acVersion, _ := d.Get("ac_version").(string)
	var appType *models.AppType // AppType
	appTypeInterface, appTypeIsSet := d.GetOk("app_type")
	if appTypeIsSet {
		appTypeModel := appTypeInterface.(string)
		appType = models.NewAppType(models.AppType(appTypeModel))
	}
	var deploymentType *models.DeploymentType // DeploymentType
	deploymentTypeInterface, deploymentTypeIsSet := d.GetOk("deployment_type")
	if deploymentTypeIsSet {
		deploymentTypeModel := deploymentTypeInterface.(string)
		deploymentType = models.NewDeploymentType(models.DeploymentType(deploymentTypeModel))
	}
	var desc *models.Details // Details
	descInterface, descIsSet := d.GetOk("desc")
	if descIsSet && descInterface != nil {
		descMap := descInterface.([]interface{})
		if len(descMap) > 0 {
			desc = DetailsModelFromMap(descMap[0].(map[string]interface{}))
		}
	}
	description, _ := d.Get("description").(string)
	displayName, _ := d.Get("display_name").(string)
	var module *models.ModuleSummary // ModuleSummary
	moduleInterface, moduleIsSet := d.GetOk("module")
	if moduleIsSet && moduleInterface != nil {
		moduleMap := moduleInterface.([]interface{})
		if len(moduleMap) > 0 {
			module = ModuleSummaryModelFromMap(moduleMap[0].(map[string]interface{}))
		}
	}
	name, _ := d.Get("name").(string)
	return &models.VMManifestSummary{
		AcKind:         &acKind,    // string true false false
		AcVersion:      &acVersion, // string true false false
		AppType:        appType,
		DeploymentType: deploymentType,
		Desc:           desc,
		Description:    description,
		DisplayName:    displayName,
		Module:         module,
		Name:           name,
	}
}

func VMManifestSummaryModelFromMap(m map[string]interface{}) *models.VMManifestSummary {
	acKind := m["ac_kind"].(string)
	acVersion := m["ac_version"].(string)
	var appType *models.AppType // AppType
	appTypeInterface, appTypeIsSet := m["app_type"]
	if appTypeIsSet {
		appTypeModel := appTypeInterface.(string)
		appType = models.NewAppType(models.AppType(appTypeModel))
	}
	var deploymentType *models.DeploymentType // DeploymentType
	deploymentTypeInterface, deploymentTypeIsSet := m["deployment_type"]
	if deploymentTypeIsSet {
		deploymentTypeModel := deploymentTypeInterface.(string)
		deploymentType = models.NewDeploymentType(models.DeploymentType(deploymentTypeModel))
	}
	var desc *models.Details // Details
	descInterface, descIsSet := m["desc"]
	if descIsSet && descInterface != nil {
		descMap := descInterface.([]interface{})
		if len(descMap) > 0 {
			desc = DetailsModelFromMap(descMap[0].(map[string]interface{}))
		}
	}
	//
	description := m["description"].(string)
	displayName := m["display_name"].(string)
	var module *models.ModuleSummary // ModuleSummary
	moduleInterface, moduleIsSet := m["module"]
	if moduleIsSet && moduleInterface != nil {
		moduleMap := moduleInterface.([]interface{})
		if len(moduleMap) > 0 {
			module = ModuleSummaryModelFromMap(moduleMap[0].(map[string]interface{}))
		}
	}
	//
	name := m["name"].(string)
	return &models.VMManifestSummary{
		AcKind:         &acKind,
		AcVersion:      &acVersion,
		AppType:        appType,
		DeploymentType: deploymentType,
		Desc:           desc,
		Description:    description,
		DisplayName:    displayName,
		Module:         module,
		Name:           name,
	}
}

// Update the underlying VMManifestSummary resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVMManifestSummaryResourceData(d *schema.ResourceData, m *models.VMManifestSummary) {
	d.Set("ac_kind", m.AcKind)
	d.Set("ac_version", m.AcVersion)
	d.Set("app_type", m.AppType)
	d.Set("deployment_type", m.DeploymentType)
	d.Set("desc", SetDetailsSubResourceData([]*models.Details{m.Desc}))
	d.Set("description", m.Description)
	d.Set("display_name", m.DisplayName)
	d.Set("module", SetModuleSummarySubResourceData([]*models.ModuleSummary{m.Module}))
	d.Set("name", m.Name)
}

// Iterate through and update the VMManifestSummary resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVMManifestSummarySubResourceData(m []*models.VMManifestSummary) (d []*map[string]interface{}) {
	for _, VMManifestSummaryModel := range m {
		if VMManifestSummaryModel != nil {
			properties := make(map[string]interface{})
			properties["ac_kind"] = VMManifestSummaryModel.AcKind
			properties["ac_version"] = VMManifestSummaryModel.AcVersion
			properties["app_type"] = VMManifestSummaryModel.AppType
			properties["deployment_type"] = VMManifestSummaryModel.DeploymentType
			properties["desc"] = SetDetailsSubResourceData([]*models.Details{VMManifestSummaryModel.Desc})
			properties["description"] = VMManifestSummaryModel.Description
			properties["display_name"] = VMManifestSummaryModel.DisplayName
			properties["module"] = SetModuleSummarySubResourceData([]*models.ModuleSummary{VMManifestSummaryModel.Module})
			properties["name"] = VMManifestSummaryModel.Name
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VMManifestSummary resource defined in the Terraform configuration
func VMManifestSummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ac_kind": {
			Description: `UI map: N/A - not exposed to users`,
			Type:        schema.TypeString,
			Default:     "VMManifest",
			Optional:    true,
		},

		"ac_version": {
			Description: `UI map: N/A - not exposed to users`,
			Type:        schema.TypeString,
			Default:     "1.2.0",
			Optional:    true,
		},

		"app_type": {
			Description: `bundle type, eg: vm, container, module`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"deployment_type": {
			Description: `type of deployment for the app, eg: azure, k3s, standalone`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"desc": {
			Description: `Details of the Edge App`,
			Type:        schema.TypeList, //GoType: Details
			Elem: &schema.Resource{
				Schema: DetailsSchema(),
			},
			Optional: true,
		},

		"description": {
			Description: `UI map: AppDetailsPage:IdentityPane:DescriptionField, AppMarketplacePage:AppCard:DescriptionField`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"display_name": {
			Description: `UI map: AppEditPage:IdentityPane:Title_Field, AppDetailsPage:IdentityPane:Title_Field`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"module": {
			Description: `Azure module specific details like module twin, environment variable, routes`,
			Type:        schema.TypeList, //GoType: ModuleSummary
			Elem: &schema.Resource{
				Schema: ModuleSummarySchema(),
			},
			Optional: true,
		},

		"name": {
			Description: `UI map: AppEditPage:IdentityPane:Name_Field, AppDetailsPage:IdentityPane:Name_Field`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the VMManifestSummary resource
func GetVMManifestSummaryPropertyFields() (t []string) {
	return []string{
		"ac_kind",
		"ac_version",
		"app_type",
		"deployment_type",
		"desc",
		"description",
		"display_name",
		"module",
		"name",
	}
}
