package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate VMManifest resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func VMManifestModel(d *schema.ResourceData) *models.VMManifest {
	acKind := d.Get("ac_kind").(string)
	acVersion := d.Get("ac_version").(string)
	appType := d.Get("app_type").(*models.AppType) // AppType
	var configuration *models.UserDataTemplate     // UserDataTemplate
	configurationInterface, configurationIsSet := d.GetOk("configuration")
	if configurationIsSet {
		configurationMap := configurationInterface.([]interface{})[0].(map[string]interface{})
		configuration = UserDataTemplateModelFromMap(configurationMap)
	}
	var containerDetail *models.ContainerDetail // ContainerDetail
	containerDetailInterface, containerDetailIsSet := d.GetOk("container_detail")
	if containerDetailIsSet {
		containerDetailMap := containerDetailInterface.([]interface{})[0].(map[string]interface{})
		containerDetail = ContainerDetailModelFromMap(containerDetailMap)
	}
	cPUPinningEnabled := d.Get("cpu_pinning_enabled").(bool)
	deploymentType := d.Get("deployment_type").(*models.DeploymentType) // DeploymentType
	var desc *models.Details                                            // Details
	descInterface, descIsSet := d.GetOk("desc")
	if descIsSet {
		descMap := descInterface.([]interface{})[0].(map[string]interface{})
		desc = DetailsModelFromMap(descMap)
	}
	description := d.Get("description").(string)
	displayName := d.Get("display_name").(string)
	enablevnc := d.Get("enablevnc").(bool)
	images := d.Get("images").([]*models.VMManifestImage)   // []*VMManifestImage
	interfaces := d.Get("interfaces").([]*models.Interface) // []*Interface
	var module *models.ModuleDetail                         // ModuleDetail
	moduleInterface, moduleIsSet := d.GetOk("module")
	if moduleIsSet {
		moduleMap := moduleInterface.([]interface{})[0].(map[string]interface{})
		module = ModuleDetailModelFromMap(moduleMap)
	}
	name := d.Get("name").(string)
	var owner *models.Author // Author
	ownerInterface, ownerIsSet := d.GetOk("owner")
	if ownerIsSet {
		ownerMap := ownerInterface.([]interface{})[0].(map[string]interface{})
		owner = AuthorModelFromMap(ownerMap)
	}
	permissions := d.Get("permissions").([]models.Permission) // []Permission
	resources := d.Get("resources").([]*models.Resource)      // []*Resource
	vmmode := d.Get("vmmode").(string)
	return &models.VMManifest{
		AcKind:            &acKind,    // string true false false
		AcVersion:         &acVersion, // string true false false
		AppType:           appType,
		Configuration:     configuration,
		ContainerDetail:   containerDetail,
		CPUPinningEnabled: cPUPinningEnabled,
		DeploymentType:    deploymentType,
		Desc:              desc,
		Description:       description,
		DisplayName:       displayName,
		Enablevnc:         enablevnc,
		Images:            images,
		Interfaces:        interfaces,
		Module:            module,
		Name:              name,
		Owner:             owner,
		Permissions:       permissions,
		Resources:         resources,
		Vmmode:            &vmmode, // string true false false
	}
}

func VMManifestModelFromMap(m map[string]interface{}) *models.VMManifest {
	acKind := m["ac_kind"].(string)
	acVersion := m["ac_version"].(string)
	appType := m["app_type"].(*models.AppType) // AppType
	var configuration *models.UserDataTemplate // UserDataTemplate
	configurationInterface, configurationIsSet := m["configuration"]
	if configurationIsSet {
		configurationMap := configurationInterface.([]interface{})[0].(map[string]interface{})
		configuration = UserDataTemplateModelFromMap(configurationMap)
	}
	//
	var containerDetail *models.ContainerDetail // ContainerDetail
	containerDetailInterface, containerDetailIsSet := m["container_detail"]
	if containerDetailIsSet {
		containerDetailMap := containerDetailInterface.([]interface{})[0].(map[string]interface{})
		containerDetail = ContainerDetailModelFromMap(containerDetailMap)
	}
	//
	cPUPinningEnabled := m["cpu_pinning_enabled"].(bool)
	deploymentType := m["deployment_type"].(*models.DeploymentType) // DeploymentType
	var desc *models.Details                                        // Details
	descInterface, descIsSet := m["desc"]
	if descIsSet {
		descMap := descInterface.([]interface{})[0].(map[string]interface{})
		desc = DetailsModelFromMap(descMap)
	}
	//
	description := m["description"].(string)
	displayName := m["display_name"].(string)
	enablevnc := m["enablevnc"].(bool)
	images := m["images"].([]*models.VMManifestImage)   // []*VMManifestImage
	interfaces := m["interfaces"].([]*models.Interface) // []*Interface
	var module *models.ModuleDetail                     // ModuleDetail
	moduleInterface, moduleIsSet := m["module"]
	if moduleIsSet {
		moduleMap := moduleInterface.([]interface{})[0].(map[string]interface{})
		module = ModuleDetailModelFromMap(moduleMap)
	}
	//
	name := m["name"].(string)
	var owner *models.Author // Author
	ownerInterface, ownerIsSet := m["owner"]
	if ownerIsSet {
		ownerMap := ownerInterface.([]interface{})[0].(map[string]interface{})
		owner = AuthorModelFromMap(ownerMap)
	}
	//
	permissions := m["permissions"].([]models.Permission) // []Permission
	resources := m["resources"].([]*models.Resource)      // []*Resource
	vmmode := m["vmmode"].(string)
	return &models.VMManifest{
		AcKind:            &acKind,
		AcVersion:         &acVersion,
		AppType:           appType,
		Configuration:     configuration,
		ContainerDetail:   containerDetail,
		CPUPinningEnabled: cPUPinningEnabled,
		DeploymentType:    deploymentType,
		Desc:              desc,
		Description:       description,
		DisplayName:       displayName,
		Enablevnc:         enablevnc,
		Images:            images,
		Interfaces:        interfaces,
		Module:            module,
		Name:              name,
		Owner:             owner,
		Permissions:       permissions,
		Resources:         resources,
		Vmmode:            &vmmode,
	}
}

// Update the underlying VMManifest resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVMManifestResourceData(d *schema.ResourceData, m *models.VMManifest) {
	d.Set("ac_kind", m.AcKind)
	d.Set("ac_version", m.AcVersion)
	d.Set("app_type", m.AppType)
	d.Set("configuration", SetUserDataTemplateSubResourceData([]*models.UserDataTemplate{m.Configuration}))
	d.Set("container_detail", SetContainerDetailSubResourceData([]*models.ContainerDetail{m.ContainerDetail}))
	d.Set("cpu_pinning_enabled", m.CPUPinningEnabled)
	d.Set("deployment_type", m.DeploymentType)
	d.Set("desc", SetDetailsSubResourceData([]*models.Details{m.Desc}))
	d.Set("description", m.Description)
	d.Set("display_name", m.DisplayName)
	d.Set("enablevnc", m.Enablevnc)
	d.Set("images", SetVMManifestImageSubResourceData(m.Images))
	d.Set("interfaces", SetInterfaceSubResourceData(m.Interfaces))
	d.Set("module", SetModuleDetailSubResourceData([]*models.ModuleDetail{m.Module}))
	d.Set("name", m.Name)
	d.Set("owner", SetAuthorSubResourceData([]*models.Author{m.Owner}))
	d.Set("permissions", m.Permissions)
	d.Set("resources", SetResourceSubResourceData(m.Resources))
	d.Set("vmmode", m.Vmmode)
}

// Iterate throught and update the VMManifest resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVMManifestSubResourceData(m []*models.VMManifest) (d []*map[string]interface{}) {
	for _, VMManifestModel := range m {
		if VMManifestModel != nil {
			properties := make(map[string]interface{})
			properties["ac_kind"] = VMManifestModel.AcKind
			properties["ac_version"] = VMManifestModel.AcVersion
			properties["app_type"] = VMManifestModel.AppType
			properties["configuration"] = SetUserDataTemplateSubResourceData([]*models.UserDataTemplate{VMManifestModel.Configuration})
			properties["container_detail"] = SetContainerDetailSubResourceData([]*models.ContainerDetail{VMManifestModel.ContainerDetail})
			properties["cpu_pinning_enabled"] = VMManifestModel.CPUPinningEnabled
			properties["deployment_type"] = VMManifestModel.DeploymentType
			properties["desc"] = SetDetailsSubResourceData([]*models.Details{VMManifestModel.Desc})
			properties["description"] = VMManifestModel.Description
			properties["display_name"] = VMManifestModel.DisplayName
			properties["enablevnc"] = VMManifestModel.Enablevnc
			properties["images"] = SetVMManifestImageSubResourceData(VMManifestModel.Images)
			properties["interfaces"] = SetInterfaceSubResourceData(VMManifestModel.Interfaces)
			properties["module"] = SetModuleDetailSubResourceData([]*models.ModuleDetail{VMManifestModel.Module})
			properties["name"] = VMManifestModel.Name
			properties["owner"] = SetAuthorSubResourceData([]*models.Author{VMManifestModel.Owner})
			properties["permissions"] = VMManifestModel.Permissions
			properties["resources"] = SetResourceSubResourceData(VMManifestModel.Resources)
			properties["vmmode"] = VMManifestModel.Vmmode
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VMManifest resource defined in the Terraform configuration
func VMManifestSchema() map[string]*schema.Schema {
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

		"configuration": {
			Description: `Template for Custom Configuration. Used for Cloud-Init`,
			Type:        schema.TypeList, //GoType: UserDataTemplate
			Elem: &schema.Resource{
				Schema: UserDataTemplateSchema(),
			},
			Optional: true,
		},

		"container_detail": {
			Description: `Create options direct the creation of the Docker container`,
			Type:        schema.TypeList, //GoType: ContainerDetail
			Elem: &schema.Resource{
				Schema: ContainerDetailSchema(),
			},
			Optional: true,
		},

		"cpu_pinning_enabled": {
			Description: `Enable CpuPinning`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"deployment_type": {
			Description: `type of deployment for the app, eg: azure, k3s, standalone`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"desc": {
			Description: `Description of the application`,
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

		"enablevnc": {
			Description: `UI map: AppEditPage:IdentityPane:VNC_Field, AppDetailsPage:IdentityPane:VNC_Field`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"images": {
			Description: `UI map: AppEditPage:DrivesPane, AppDetailsPage:DrivesPane`,
			Type:        schema.TypeList, //GoType: []*VMManifestImage
			Elem: &schema.Resource{
				Schema: VMManifestImageSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"interfaces": {
			Description: `UI map: AppEditPage:EnvironmentsPane, AppDetailsPage:EnvironmentsPane`,
			Type:        schema.TypeList, //GoType: []*Interface
			Elem: &schema.Resource{
				Schema: InterfaceSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"module": {
			Description: `Azure module specific details like module twin, environment variable, routes`,
			Type:        schema.TypeList, //GoType: ModuleDetail
			Elem: &schema.Resource{
				Schema: ModuleDetailSchema(),
			},
			Optional: true,
		},

		"name": {
			Description: `UI map: AppEditPage:IdentityPane:Name_Field, AppDetailsPage:IdentityPane:Name_Field`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"owner": {
			Description: `Owner of the application`,
			Type:        schema.TypeList, //GoType: Author
			Elem: &schema.Resource{
				Schema: AuthorSchema(),
			},
			Optional: true,
		},

		"permissions": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []Permission
			Elem: &schema.Resource{
				Schema: PermissionSchema(),
			},
			Optional: true,
		},

		"resources": {
			Description: `UI map: AppEditPage:ResourcesPane, AppDetailsPage:ResourcesPane`,
			Type:        schema.TypeList, //GoType: []*Resource
			Elem: &schema.Resource{
				Schema: ResourceSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"vmmode": {
			Description: `UI map: AppEditPage:IdentityPane:VM_Mode_Field, AppDetailsPage:IdentityPane:VM_Mode_Field`,
			Type:        schema.TypeString,
			Default:     "HV_HVM",
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the VMManifest resource
func GetVMManifestPropertyFields() (t []string) {
	return []string{
		"ac_kind",
		"ac_version",
		"app_type",
		"configuration",
		"container_detail",
		"cpu_pinning_enabled",
		"deployment_type",
		"desc",
		"description",
		"display_name",
		"enablevnc",
		"images",
		"interfaces",
		"module",
		"name",
		"owner",
		"permissions",
		"resources",
		"vmmode",
	}
}
