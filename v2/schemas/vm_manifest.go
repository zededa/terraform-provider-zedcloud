package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func VMManifestModel(d *schema.ResourceData) *models.VMManifest {
	acKind, _ := d.Get("ac_kind").(string)
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
	enableOemWinLicenseKey, _ := d.Get("enable_oem_win_license_key").(bool)
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
	dockerComposeTarImageName, _ := d.Get("docker_compose_tar_image_name").(string)
	dockerComposeYamlText, _ := d.Get("docker_compose_yaml_text").(string)
	persistentRuntimeSizeBytes, _ := d.Get("persistent_runtime_size_bytes").(string)
	var runtimeProtocolVersion *models.DockerRuntimeProtocolVersion // DockerRuntimeProtocolVersion
	runtimeProtocolVersionInterface, runtimeProtocolVersionIsSet := d.GetOk("runtime_protocol_version")
	if runtimeProtocolVersionIsSet {
		runtimeProtocolVersionModel := runtimeProtocolVersionInterface.(string)
		runtimeProtocolVersion = models.NewDockerRuntimeProtocolVersion(models.DockerRuntimeProtocolVersion(runtimeProtocolVersionModel))
	}
	var runtimeVersion *models.DockerRuntimeVersion // DockerRuntimeVersion
	runtimeVersionInterface, runtimeVersionIsSet := d.GetOk("runtime_version")
	if runtimeVersionIsSet {
		runtimeVersionModel := runtimeVersionInterface.(string)
		runtimeVersion = models.NewDockerRuntimeVersion(models.DockerRuntimeVersion(runtimeVersionModel))
	}
	return &models.VMManifest{
		AcKind:                     &acKind,    // string
		AcVersion:                  &acVersion, // string
		AppType:                    appType,
		Configuration:              configuration,
		ContainerDetail:            containerDetail,
		CPUPinningEnabled:          cPUPinningEnabled,
		DeploymentType:             deploymentType,
		Desc:                       desc,
		Description:                description,
		DisplayName:                displayName,
		EnableOemWinLicenseKey:     enableOemWinLicenseKey,
		Enablevnc:                  enablevnc,
		Images:                     images,
		Interfaces:                 interfaces,
		Module:                     module,
		Name:                       name,
		Owner:                      owner,
		Permissions:                permissions,
		Resources:                  resources,
		Vmmode:                     &vmmode, // string
		DockerComposeTarImageName:  dockerComposeTarImageName,
		DockerComposeYamlText:      dockerComposeYamlText,
		PersistentRuntimeSizeBytes: persistentRuntimeSizeBytes,
		RuntimeProtocolVersion:     runtimeProtocolVersion,
		RuntimeVersion:             runtimeVersion,
	}
}

func VMManifestModelFromMap(m map[string]interface{}) *models.VMManifest {
	acKind := m["ac_kind"].(string)
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
	enableOemWinLicenseKey := m["enable_oem_win_license_key"].(bool)
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
	dockerComposeTarImageName := m["docker_compose_tar_image_name"].(string)
	dockerComposeYamlText := m["docker_compose_yaml_text"].(string)
	persistentRuntimeSizeBytes := m["persistent_runtime_size_bytes"].(string)
	var runtimeProtocolVersion *models.DockerRuntimeProtocolVersion // DockerRuntimeProtocolVersion
	runtimeProtocolVersionInterface, runtimeProtocolVersionIsSet := m["runtime_protocol_version"]
	if runtimeProtocolVersionIsSet {
		runtimeProtocolVersionModel := runtimeProtocolVersionInterface.(string)
		runtimeProtocolVersion = models.NewDockerRuntimeProtocolVersion(models.DockerRuntimeProtocolVersion(runtimeProtocolVersionModel))
	}
	var runtimeVersion *models.DockerRuntimeVersion // DockerRuntimeVersion
	runtimeVersionInterface, runtimeVersionIsSet := m["runtime_version"]
	if runtimeVersionIsSet {
		runtimeVersionModel := runtimeVersionInterface.(string)
		runtimeVersion = models.NewDockerRuntimeVersion(models.DockerRuntimeVersion(runtimeVersionModel))
	}
	return &models.VMManifest{
		AcKind:                     &acKind,
		AcVersion:                  &acVersion,
		AppType:                    appType,
		Configuration:              configuration,
		ContainerDetail:            containerDetail,
		CPUPinningEnabled:          cPUPinningEnabled,
		DeploymentType:             deploymentType,
		Desc:                       desc,
		Description:                description,
		DisplayName:                displayName,
		EnableOemWinLicenseKey:     enableOemWinLicenseKey,
		Enablevnc:                  enablevnc,
		Images:                     images,
		Interfaces:                 interfaces,
		Module:                     module,
		Name:                       name,
		Owner:                      owner,
		Permissions:                permissions,
		Resources:                  resources,
		Vmmode:                     &vmmode,
		DockerComposeTarImageName:  dockerComposeTarImageName,
		DockerComposeYamlText:      dockerComposeYamlText,
		PersistentRuntimeSizeBytes: persistentRuntimeSizeBytes,
		RuntimeProtocolVersion:     runtimeProtocolVersion,
		RuntimeVersion:             runtimeVersion,
	}
}

func SetVMManifestResourceData(d *schema.ResourceData, m *models.VMManifest) {
	d.Set("ac_kind", m.AcKind)
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
	d.Set("docker_compose_tar_image_name", m.DockerComposeTarImageName)
	d.Set("docker_compose_yaml_text", m.DockerComposeYamlText)
	d.Set("persistent_runtime_size_bytes", m.PersistentRuntimeSizeBytes)
	d.Set("runtime_protocol_version", m.RuntimeProtocolVersion)
	d.Set("runtime_version", m.RuntimeVersion)
}

func SetVMManifestSubResourceData(m []*models.VMManifest) (d []*map[string]interface{}) {
	for _, VMManifestModel := range m {
		if VMManifestModel != nil {
			properties := make(map[string]interface{})
			properties["ac_kind"] = VMManifestModel.AcKind
			properties["ac_version"] = VMManifestModel.AcVersion
			properties["app_type"] = VMManifestModel.AppType
			properties["configuration"] = SetUserTemplateSubResourceData([]*models.UserDataTemplate{VMManifestModel.Configuration})
			properties["container_detail"] = SetContainerDetailSubResourceData([]*models.ContainerDetail{VMManifestModel.ContainerDetail})
			properties["cpu_pinning_enabled"] = VMManifestModel.CPUPinningEnabled
			properties["deployment_type"] = VMManifestModel.DeploymentType
			properties["desc"] = SetDetailsSubResourceData([]*models.Details{VMManifestModel.Desc})
			properties["description"] = VMManifestModel.Description
			properties["display_name"] = VMManifestModel.DisplayName
			properties["enable_oem_win_license_key"] = VMManifestModel.EnableOemWinLicenseKey
			properties["enablevnc"] = VMManifestModel.Enablevnc
			properties["images"] = SetVMManifestImageSubResourceData(VMManifestModel.Images)
			properties["interfaces"] = SetInterfaceSubResourceData(VMManifestModel.Interfaces)
			properties["module"] = SetModuleDetailSubResourceData([]*models.ModuleDetail{VMManifestModel.Module})
			properties["name"] = VMManifestModel.Name
			properties["owner"] = SetAuthorSubResourceData([]*models.Author{VMManifestModel.Owner})
			properties["permissions"] = VMManifestModel.Permissions
			properties["resources"] = SetResourceSubResourceData(VMManifestModel.Resources)
			properties["vmmode"] = VMManifestModel.Vmmode
			properties["docker_compose_tar_image_name"] = VMManifestModel.DockerComposeTarImageName
			properties["docker_compose_yaml_text"] = VMManifestModel.DockerComposeYamlText
			properties["persistent_runtime_size_bytes"] = VMManifestModel.PersistentRuntimeSizeBytes
			properties["runtime_protocol_version"] = VMManifestModel.RuntimeProtocolVersion
			properties["runtime_version"] = VMManifestModel.RuntimeVersion
			d = append(d, &properties)
		}
	}
	return
}

func VMManifest() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ac_kind": {
			Description: `UI map: N/A - not exposed to users`,
			Type:        schema.TypeString,
			// Default:     "VMManifest",
			Required: true,
		},

		"ac_version": {
			Description: `UI map: N/A - not exposed to users`,
			Type:        schema.TypeString,
			// Default:     "1.2.0",
			Required: true,
		},

		"app_type": {
			Description: `App (bundle) type. The correct values are: "APP_TYPE_UNSPECIFIED","APP_TYPE_VM",` +
				`"APP_TYPE_VM_RUNTIME","APP_TYPE_CONTAINER","APP_TYPE_MODULE", "APP_TYPE_DOCKER_COMPOSE".`,
			Type:     schema.TypeString,
			Optional: true,
		},

		"configuration": {
			Description: `Template for Custom Configuration. Used for Cloud-Init`,
			Type:        schema.TypeList, // GoType: UserDataTemplate
			Elem: &schema.Resource{
				Schema: UserTemplate(),
			},
			Optional: true,
		},

		"container_detail": {
			Description: `Create options direct the creation of the Docker container`,
			Type:        schema.TypeList, // GoType: ContainerDetail
			Elem: &schema.Resource{
				Schema: ContainerDetail(),
			},
			Optional: true,
		},

		"cpu_pinning_enabled": {
			Description: `Enable CpuPinning`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"deployment_type": {
			Description: `Deployment type for the app. The correct values are: ` +
				`"DEPLOYMENT_TYPE_UNSPECIFIED","DEPLOYMENT_TYPE_STAND_ALONE","DEPLOYMENT_TYPE_AZURE",` +
				`"DEPLOYMENT_TYPE_K3S","DEPLOYMENT_TYPE_AWS","DEPLOYMENT_TYPE_K3S_AZURE","DEPLOYMENT_TYPE_K3S_AWS",` +
				`"DEPLOYMENT_TYPE_VMWARE_VCE".`,
			Type:     schema.TypeString,
			Optional: true,
		},

		"desc": {
			Description: `Description of the application`,
			Type:        schema.TypeList, // GoType: Details
			Elem: &schema.Resource{
				Schema: Details(),
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

		"enable_oem_win_license_key": {
			Description: `UI map: AppEditPage:IdentityPane:ENABLEVMCONFIG_Field, AppDetailsPage:IdentityPane:ENABLEVMCONFIG_Field`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"enablevnc": {
			Description: `UI map: AppEditPage:IdentityPane:VNC_Field, AppDetailsPage:IdentityPane:VNC_Field`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"images": {
			Description: `UI map: AppEditPage:DrivesPane, AppDetailsPage:DrivesPane`,
			Type:        schema.TypeList, // GoType: []*VMManifestImage
			Elem: &schema.Resource{
				Schema: VMManifestImage(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"interfaces": {
			Description: `UI map: AppEditPage:EnvironmentsPane, AppDetailsPage:EnvironmentsPane`,
			Type:        schema.TypeList, // GoType: []*Interface
			Elem: &schema.Resource{
				Schema: Interface(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"module": {
			Description: `Azure module specific details like module twin, environment variable, routes`,
			Type:        schema.TypeList, // GoType: ModuleDetail
			Elem: &schema.Resource{
				Schema: ModuleDetail(),
			},
			Optional: true,
		},

		"name": {
			Description: `UI map: AppEditPage:IdentityPane:Name_Field, AppDetailsPage:IdentityPane:Name_Field`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"owner": {
			Description: `Owner of the application`,
			Type:        schema.TypeList, // GoType: Author
			Elem: &schema.Resource{
				Schema: Author(),
			},
			Optional: true,
		},

		"permissions": {
			Description: ``,
			Type:        schema.TypeList, // GoType: []Permission
			Elem: &schema.Resource{
				Schema: Permission(),
			},
			Optional: true,
		},

		"resources": {
			Description: `UI map: AppEditPage:ResourcesPane, AppDetailsPage:ResourcesPane`,
			Type:        schema.TypeList, // GoType: []*Resource
			Elem: &schema.Resource{
				Schema: Resource(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"vmmode": {
			Description: `UI map: AppEditPage:IdentityPane:VM_Mode_Field, AppDetailsPage:IdentityPane:VM_Mode_Field`,
			Type:        schema.TypeString,
			//Default:     "HV_HVM",
			Optional: true,
		},

		"docker_compose_tar_image_name": {
			Description: `Docker compose tar image name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"docker_compose_yaml_text": {
			Description: `Docker compose base64 encoded plain text`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"persistent_runtime_size_bytes": {
			Description: `Size of persistent blank storage for runtime in bytes`,
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "0",
		},

		"runtime_protocol_version": {
			Description: `Indicates the internal communication protocol to pass configuration between Zedcloud and docker-compose runtime`,
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "RuntimeProtocolVersion_Unknown",
		},

		"runtime_version": {
			Description: `Indicates the version of container orchestration software`,
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "RuntimeVersion_Unknown",
		},
	}
}

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
		"enable_oem_win_license_key",
		"enablevnc",
		"images",
		"interfaces",
		"module",
		"name",
		"owner",
		"permissions",
		"resources",
		"vmmode",
		"docker_compose_tar_image_name",
		"docker_compose_yaml_text",
		"persistent_runtime_size_bytes",
		"runtime_protocol_version",
		"runtime_version",
	}
}
