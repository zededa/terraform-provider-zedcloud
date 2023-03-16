package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func VMManifestModel(d *schema.ResourceData) *models.VMManifest {
	acVersion, _ := d.Get("ac_version").(string)
	var appType *models.AppType // AppType
	appTypeInterface, appTypeIsSet := d.GetOk("app_type")
	if appTypeIsSet {
		appTypeModel := appTypeInterface.(string)
		appType = models.NewAppType(models.AppType(appTypeModel))
	}
	var configuration *models.UserDataTemplate // UserDataTemplate
	configurationInterface, configurationIsSet := d.GetOk("configuration")
	if configurationIsSet && configurationInterface != nil {
		configurationMap := configurationInterface.([]interface{})
		if len(configurationMap) > 0 {
			configuration = UserTemplateModelFromMap(configurationMap[0].(map[string]interface{}))
		}
	}
	var containerDetail *models.ContainerDetail // ContainerDetail
	containerDetailInterface, containerDetailIsSet := d.GetOk("container_detail")
	if containerDetailIsSet && containerDetailInterface != nil {
		containerDetailMap := containerDetailInterface.([]interface{})
		if len(containerDetailMap) > 0 {
			containerDetail = ContainerDetailModelFromMap(containerDetailMap[0].(map[string]interface{}))
		}
	}
	cPUPinningEnabled, _ := d.Get("cpu_pinning_enabled").(bool)
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
	enablevnc, _ := d.Get("enablevnc").(bool)
	var images []*models.VMManifestImage // []*VMManifestImage
	imagesInterface, imagesIsSet := d.GetOk("images")
	if imagesIsSet {
		var items []interface{}
		if listItems, isList := imagesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = imagesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := VMManifestImageModelFromMap(v.(map[string]interface{}))
			images = append(images, m)
		}
	}
	var interfaces []*models.Interface // []*Interface
	interfacesInterface, interfacesIsSet := d.GetOk("interfaces")
	if interfacesIsSet {
		var items []interface{}
		if listItems, isList := interfacesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = interfacesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := InterfaceModelFromMap(v.(map[string]interface{}))
			interfaces = append(interfaces, m)
		}
	}
	var module *models.ModuleDetail // ModuleDetail
	moduleInterface, moduleIsSet := d.GetOk("module")
	if moduleIsSet && moduleInterface != nil {
		moduleMap := moduleInterface.([]interface{})
		if len(moduleMap) > 0 {
			module = ModuleDetailModelFromMap(moduleMap[0].(map[string]interface{}))
		}
	}
	name, _ := d.Get("name").(string)
	var owner *models.Author // Author
	ownerInterface, ownerIsSet := d.GetOk("owner")
	if ownerIsSet && ownerInterface != nil {
		ownerMap := ownerInterface.([]interface{})
		if len(ownerMap) > 0 {
			owner = AuthorModelFromMap(ownerMap[0].(map[string]interface{}))
		}
	}
	var permissions []models.Permission // []Permission
	permissionsInterface, permissionsIsSet := d.GetOk("permissions")
	if permissionsIsSet {
		var items []interface{}
		if listItems, isList := permissionsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = permissionsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := PermissionModelFromMap(v.(map[string]interface{}))
			permissions = append(permissions, *m)
		}
	}
	var resources []*models.Resource // []*Resource
	resourcesInterface, resourcesIsSet := d.GetOk("resources")
	if resourcesIsSet {
		var items []interface{}
		if listItems, isList := resourcesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = resourcesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ResourceModelFromMap(v.(map[string]interface{}))
			resources = append(resources, m)
		}
	}
	vmmode, _ := d.Get("vmmode").(string)
	return &models.VMManifest{
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
	acVersion := m["ac_version"].(string)
	var appType *models.AppType // AppType
	appTypeInterface, appTypeIsSet := m["app_type"]
	if appTypeIsSet {
		appTypeModel := appTypeInterface.(string)
		appType = models.NewAppType(models.AppType(appTypeModel))
	}
	var configuration *models.UserDataTemplate // UserDataTemplate
	configurationInterface, configurationIsSet := m["configuration"]
	if configurationIsSet && configurationInterface != nil {
		configurationMap := configurationInterface.([]interface{})
		if len(configurationMap) > 0 {
			configuration = UserTemplateModelFromMap(configurationMap[0].(map[string]interface{}))
		}
	}
	//
	var containerDetail *models.ContainerDetail // ContainerDetail
	containerDetailInterface, containerDetailIsSet := m["container_detail"]
	if containerDetailIsSet && containerDetailInterface != nil {
		containerDetailMap := containerDetailInterface.([]interface{})
		if len(containerDetailMap) > 0 {
			containerDetail = ContainerDetailModelFromMap(containerDetailMap[0].(map[string]interface{}))
		}
	}
	//
	cPUPinningEnabled := m["cpu_pinning_enabled"].(bool)
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
	enablevnc := m["enablevnc"].(bool)
	var images []*models.VMManifestImage // []*VMManifestImage
	imagesInterface, imagesIsSet := m["images"]
	if imagesIsSet {
		var items []interface{}
		if listItems, isList := imagesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = imagesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := VMManifestImageModelFromMap(v.(map[string]interface{}))
			images = append(images, m)
		}
	}
	var interfaces []*models.Interface // []*Interface
	interfacesInterface, interfacesIsSet := m["interfaces"]
	if interfacesIsSet {
		var items []interface{}
		if listItems, isList := interfacesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = interfacesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := InterfaceModelFromMap(v.(map[string]interface{}))
			interfaces = append(interfaces, m)
		}
	}
	var module *models.ModuleDetail // ModuleDetail
	moduleInterface, moduleIsSet := m["module"]
	if moduleIsSet && moduleInterface != nil {
		moduleMap := moduleInterface.([]interface{})
		if len(moduleMap) > 0 {
			module = ModuleDetailModelFromMap(moduleMap[0].(map[string]interface{}))
		}
	}
	//
	name := m["name"].(string)
	var owner *models.Author // Author
	ownerInterface, ownerIsSet := m["owner"]
	if ownerIsSet && ownerInterface != nil {
		ownerMap := ownerInterface.([]interface{})
		if len(ownerMap) > 0 {
			owner = AuthorModelFromMap(ownerMap[0].(map[string]interface{}))
		}
	}
	//
	var permissions []models.Permission // []Permission
	permissionsInterface, permissionsIsSet := m["permissions"]
	if permissionsIsSet {
		var items []interface{}
		if listItems, isList := permissionsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = permissionsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := PermissionModelFromMap(v.(map[string]interface{}))
			permissions = append(permissions, *m)
		}
	}
	var resources []*models.Resource // []*Resource
	resourcesInterface, resourcesIsSet := m["resources"]
	if resourcesIsSet {
		var items []interface{}
		if listItems, isList := resourcesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = resourcesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ResourceModelFromMap(v.(map[string]interface{}))
			resources = append(resources, m)
		}
	}
	vmmode := m["vmmode"].(string)
	return &models.VMManifest{
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

func SetVMManifestResourceData(d *schema.ResourceData, m *models.VMManifest) {
	d.Set("ac_version", m.AcVersion)
	d.Set("app_type", m.AppType)
	d.Set("configuration", SetUserTemplateSubResourceData([]*models.UserDataTemplate{m.Configuration}))
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

func SetVMManifestSubResourceData(m []*models.VMManifest) (d []*map[string]interface{}) {
	for _, VMManifestModel := range m {
		if VMManifestModel != nil {
			properties := make(map[string]interface{})
			properties["ac_version"] = VMManifestModel.AcVersion
			properties["app_type"] = VMManifestModel.AppType
			properties["configuration"] = SetUserTemplateSubResourceData([]*models.UserDataTemplate{VMManifestModel.Configuration})
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
				Schema: UserTemplate(),
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
			Computed: true,
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
