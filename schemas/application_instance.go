package schemas

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func ApplicationInstanceModel(d *schema.ResourceData) *models.AppInstance {
	activate, _ := d.Get("activate").(bool)
	appID, _ := d.Get("app_id").(string)
	appPolicyID, _ := d.Get("app_policy_id").(string)
	var appType *models.AppType // AppType
	appTypeInterface, appTypeIsSet := d.GetOk("app_type")
	if appTypeIsSet {
		appTypeModel := appTypeInterface.(string)
		appType = models.NewAppType(models.AppType(appTypeModel))
		if err := appType.Validate(nil); err != nil {
			log.Fatalf("invalid value for field app_type (%v): %v", *appType, err)
		}
	}
	bundleversion, _ := d.Get("bundleversion").(string)
	clusterID, _ := d.Get("cluster_id").(string)
	collectStatsIPAddr, _ := d.Get("collect_stats_ip_addr").(string)
	cryptoKey, _ := d.Get("crypto_key").(string)
	var customConfig *models.CustomConfig // CustomConfig
	customConfigInterface, customConfigIsSet := d.GetOk("custom_config")
	if customConfigIsSet && customConfigInterface != nil {
		customConfigMap := customConfigInterface.([]interface{})
		if len(customConfigMap) > 0 {
			customConfig = CustomConfigModelFromMap(customConfigMap[0].(map[string]interface{}))
		}
	}
	var deploymentType *models.DeploymentType // DeploymentType
	deploymentTypeInterface, deploymentTypeIsSet := d.GetOk("deployment_type")
	if deploymentTypeIsSet {
		deploymentTypeModel := deploymentTypeInterface.(string)
		deploymentType = models.NewDeploymentType(models.DeploymentType(deploymentTypeModel))
	}
	description, _ := d.Get("description").(string)
	deviceID, _ := d.Get("device_id").(string)
	var drives []*models.Drive // []*Drive
	drivesInterface, drivesIsSet := d.GetOk("drives")
	if drivesIsSet {
		var items []interface{}
		if listItems, isList := drivesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = drivesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := DriveModelFromMap(v.(map[string]interface{}))
			drives = append(drives, m)
		}
	}
	encryptedSecrets := map[string]string{}
	encryptedSecretsInterface, encryptedSecretsIsSet := d.GetOk("encryptedSecrets")
	if encryptedSecretsIsSet {
		encryptedSecretsMap := encryptedSecretsInterface.(map[string]interface{})
		for k, v := range encryptedSecretsMap {
			if v == nil {
				continue
			}
			encryptedSecrets[k] = v.(string)
		}
	}

	id, _ := d.Get("id").(string)
	var interfaces []*models.AppInterface // []*AppInterface
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
			m := AppInterfaceModelFromMap(v.(map[string]interface{}))
			interfaces = append(interfaces, m)
		}
	}
	isSecretUpdated, _ := d.Get("is_secret_updated").(bool)
	var logs *models.AppInstanceLogs // AppInstanceLogs
	logsInterface, logsIsSet := d.GetOk("logs")
	if logsIsSet && logsInterface != nil {
		logsMap := logsInterface.([]interface{})
		if len(logsMap) > 0 {
			logs = AppInstanceLogsModelFromMap(logsMap[0].(map[string]interface{}))
		}
	}
	var manifestInfo *models.ManifestInfo // ManifestInfo
	manifestInfoInterface, manifestInfoIsSet := d.GetOk("manifest_info")
	if manifestInfoIsSet && manifestInfoInterface != nil {
		manifestInfoMap := manifestInfoInterface.([]interface{})
		if len(manifestInfoMap) > 0 {
			manifestInfo = ManifestInfoModelFromMap(manifestInfoMap[0].(map[string]interface{}))
		}
	}
	name, _ := d.Get("name").(string)
	projectID, _ := d.Get("project_id").(string)
	var purge *models.ZedCloudOpsCmd // ZedCloudOpsCmd
	purgeInterface, purgeIsSet := d.GetOk("purge")
	if purgeIsSet && purgeInterface != nil {
		purgeMap := purgeInterface.([]interface{})
		if len(purgeMap) > 0 {
			purge = ZedCloudOpsCmdModelFromMap(purgeMap[0].(map[string]interface{}))
		}
	}
	var refresh *models.ZedCloudOpsCmd // ZedCloudOpsCmd
	refreshInterface, refreshIsSet := d.GetOk("refresh")
	if refreshIsSet && refreshInterface != nil {
		refreshMap := refreshInterface.([]interface{})
		if len(refreshMap) > 0 {
			refresh = ZedCloudOpsCmdModelFromMap(refreshMap[0].(map[string]interface{}))
		}
	}
	remoteConsole, _ := d.Get("remote_console").(bool)
	var restart *models.ZedCloudOpsCmd // ZedCloudOpsCmd
	restartInterface, restartIsSet := d.GetOk("restart")
	if restartIsSet && restartInterface != nil {
		restartMap := restartInterface.([]interface{})
		if len(restartMap) > 0 {
			restart = ZedCloudOpsCmdModelFromMap(restartMap[0].(map[string]interface{}))
		}
	}
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := d.GetOk("revision")
	if revisionIsSet && revisionInterface != nil {
		revisionMap := revisionInterface.([]interface{})
		if len(revisionMap) > 0 {
			revision = ObjectRevisionModelFromMap(revisionMap[0].(map[string]interface{}))
		}
	}
	startDelayInSecondsInt, _ := d.Get("start_delay_in_seconds").(int)
	startDelayInSeconds := int32(startDelayInSecondsInt)
	tags := map[string]string{}
	tagsInterface, tagsIsSet := d.GetOk("tags")
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	title, _ := d.Get("title").(string)
	userData, _ := d.Get("user_data").(string)
	userDefinedVersion, _ := d.Get("user_defined_version").(string)
	var vminfo *models.VM // VM
	vminfoInterface, vminfoIsSet := d.GetOk("vminfo")
	if vminfoIsSet && vminfoInterface != nil {
		vminfoMap := vminfoInterface.([]interface{})
		if len(vminfoMap) > 0 {
			vminfo = VMModelFromMap(vminfoMap[0].(map[string]interface{}))
		}
	}
	return &models.AppInstance{
		Activate:            &activate, // string true false false
		AppID:               &appID,    // string true false false
		AppPolicyID:         appPolicyID,
		AppType:             appType,
		Bundleversion:       bundleversion,
		ClusterID:           clusterID,
		CollectStatsIPAddr:  collectStatsIPAddr,
		CryptoKey:           cryptoKey,
		CustomConfig:        customConfig,
		DeploymentType:      deploymentType,
		Description:         description,
		DeviceID:            &deviceID, // string true false false
		Drives:              drives,
		EncryptedSecrets:    encryptedSecrets,
		ID:                  id,
		Interfaces:          interfaces,
		IsSecretUpdated:     isSecretUpdated,
		Logs:                logs,
		ManifestInfo:        manifestInfo,
		Name:                &name,      // string true false false
		ProjectID:           &projectID, // string true false false
		Purge:               purge,
		Refresh:             refresh,
		RemoteConsole:       remoteConsole,
		Restart:             restart,
		Revision:            revision,
		StartDelayInSeconds: startDelayInSeconds,
		Tags:                tags,
		Title:               &title, // string true false false
		UserData:            userData,
		UserDefinedVersion:  userDefinedVersion,
		Vminfo:              vminfo,
	}
}

func ApplicationInstanceModelFromMap(m map[string]interface{}) *models.AppInstance {
	activate := m["activate"].(bool)
	appID := m["app_id"].(string)
	appPolicyID := m["app_policy_id"].(string)
	var appType *models.AppType // AppType
	appTypeInterface, appTypeIsSet := m["app_type"]
	if appTypeIsSet {
		appTypeModel := appTypeInterface.(string)
		appType = models.NewAppType(models.AppType(appTypeModel))
		if err := appType.Validate(nil); err != nil {
			log.Fatalf("invalid value for field app_type (%v): %v", *appType, err)
		}
	}
	bundleversion := m["bundleversion"].(string)
	clusterID := m["cluster_id"].(string)
	collectStatsIPAddr := m["collect_stats_ip_addr"].(string)
	cryptoKey := m["crypto_key"].(string)
	var customConfig *models.CustomConfig // CustomConfig
	customConfigInterface, customConfigIsSet := m["custom_config"]
	if customConfigIsSet && customConfigInterface != nil {
		customConfigMap := customConfigInterface.([]interface{})
		if len(customConfigMap) > 0 {
			customConfig = CustomConfigModelFromMap(customConfigMap[0].(map[string]interface{}))
		}
	}
	//
	var deploymentType *models.DeploymentType // DeploymentType
	deploymentTypeInterface, deploymentTypeIsSet := m["deployment_type"]
	if deploymentTypeIsSet {
		deploymentTypeModel := deploymentTypeInterface.(string)
		deploymentType = models.NewDeploymentType(models.DeploymentType(deploymentTypeModel))
	}
	description := m["description"].(string)
	deviceID := m["device_id"].(string)
	var drives []*models.Drive // []*Drive
	drivesInterface, drivesIsSet := m["drives"]
	if drivesIsSet {
		var items []interface{}
		if listItems, isList := drivesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = drivesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := DriveModelFromMap(v.(map[string]interface{}))
			drives = append(drives, m)
		}
	}
	encryptedSecrets := map[string]string{}
	encryptedSecretsInterface, encryptedSecretsIsSet := m["encrypted_secrets"]
	if encryptedSecretsIsSet {
		encryptedSecretsMap := encryptedSecretsInterface.(map[string]interface{})
		for k, v := range encryptedSecretsMap {
			if v == nil {
				continue
			}
			encryptedSecrets[k] = v.(string)
		}
	}

	id := m["id"].(string)
	var interfaces []*models.AppInterface // []*AppInterface
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
			m := AppInterfaceModelFromMap(v.(map[string]interface{}))
			interfaces = append(interfaces, m)
		}
	}
	isSecretUpdated := m["is_secret_updated"].(bool)
	var logs *models.AppInstanceLogs // AppInstanceLogs
	logsInterface, logsIsSet := m["logs"]
	if logsIsSet && logsInterface != nil {
		logsMap := logsInterface.([]interface{})
		if len(logsMap) > 0 {
			logs = AppInstanceLogsModelFromMap(logsMap[0].(map[string]interface{}))
		}
	}
	//
	var manifestInfo *models.ManifestInfo // ManifestInfo
	manifestInfoInterface, manifestInfoIsSet := m["manifest_info"]
	if manifestInfoIsSet && manifestInfoInterface != nil {
		manifestInfoMap := manifestInfoInterface.([]interface{})
		if len(manifestInfoMap) > 0 {
			manifestInfo = ManifestInfoModelFromMap(manifestInfoMap[0].(map[string]interface{}))
		}
	}
	//
	name := m["name"].(string)
	projectID := m["project_id"].(string)
	var purge *models.ZedCloudOpsCmd // ZedCloudOpsCmd
	purgeInterface, purgeIsSet := m["purge"]
	if purgeIsSet && purgeInterface != nil {
		purgeMap := purgeInterface.([]interface{})
		if len(purgeMap) > 0 {
			purge = ZedCloudOpsCmdModelFromMap(purgeMap[0].(map[string]interface{}))
		}
	}
	//
	var refresh *models.ZedCloudOpsCmd // ZedCloudOpsCmd
	refreshInterface, refreshIsSet := m["refresh"]
	if refreshIsSet && refreshInterface != nil {
		refreshMap := refreshInterface.([]interface{})
		if len(refreshMap) > 0 {
			refresh = ZedCloudOpsCmdModelFromMap(refreshMap[0].(map[string]interface{}))
		}
	}
	//
	remoteConsole := m["remote_console"].(bool)
	var restart *models.ZedCloudOpsCmd // ZedCloudOpsCmd
	restartInterface, restartIsSet := m["restart"]
	if restartIsSet && restartInterface != nil {
		restartMap := restartInterface.([]interface{})
		if len(restartMap) > 0 {
			restart = ZedCloudOpsCmdModelFromMap(restartMap[0].(map[string]interface{}))
		}
	}
	//
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := m["revision"]
	if revisionIsSet && revisionInterface != nil {
		revisionMap := revisionInterface.([]interface{})
		if len(revisionMap) > 0 {
			revision = ObjectRevisionModelFromMap(revisionMap[0].(map[string]interface{}))
		}
	}
	//
	startDelayInSeconds := int32(m["start_delay_in_seconds"].(int)) // int64
	tags := map[string]string{}
	tagsInterface, tagsIsSet := m["tags"]
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	title := m["title"].(string)
	userData := m["user_data"].(string)
	userDefinedVersion := m["user_defined_version"].(string)
	var vminfo *models.VM // VM
	vminfoInterface, vminfoIsSet := m["vminfo"]
	if vminfoIsSet && vminfoInterface != nil {
		vminfoMap := vminfoInterface.([]interface{})
		if len(vminfoMap) > 0 {
			vminfo = VMModelFromMap(vminfoMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.AppInstance{
		Activate:            &activate,
		AppID:               &appID,
		AppPolicyID:         appPolicyID,
		AppType:             appType,
		Bundleversion:       bundleversion,
		ClusterID:           clusterID,
		CollectStatsIPAddr:  collectStatsIPAddr,
		CryptoKey:           cryptoKey,
		CustomConfig:        customConfig,
		DeploymentType:      deploymentType,
		Description:         description,
		DeviceID:            &deviceID,
		Drives:              drives,
		EncryptedSecrets:    encryptedSecrets,
		ID:                  id,
		Interfaces:          interfaces,
		IsSecretUpdated:     isSecretUpdated,
		Logs:                logs,
		ManifestInfo:        manifestInfo,
		Name:                &name,
		ProjectID:           &projectID,
		Purge:               purge,
		Refresh:             refresh,
		RemoteConsole:       remoteConsole,
		Restart:             restart,
		Revision:            revision,
		StartDelayInSeconds: startDelayInSeconds,
		Tags:                tags,
		Title:               &title,
		UserData:            userData,
		UserDefinedVersion:  userDefinedVersion,
		Vminfo:              vminfo,
	}
}

func SetApplicationInstanceResourceData(d *schema.ResourceData, m *models.AppInstance) {
	d.Set("activate", m.Activate)
	d.Set("app_id", m.AppID)
	d.Set("app_policy_id", m.AppPolicyID)
	d.Set("app_type", m.AppType)
	d.Set("bundleversion", m.Bundleversion)
	d.Set("cluster_id", m.ClusterID)
	d.Set("collect_stats_ip_addr", m.CollectStatsIPAddr)
	d.Set("crypto_key", m.CryptoKey)
	d.Set("custom_config", SetCustomConfigSubResourceData([]*models.CustomConfig{m.CustomConfig}))
	d.Set("deployment_type", m.DeploymentType)
	d.Set("description", m.Description)
	d.Set("device_id", m.DeviceID)
	d.Set("drives", SetDriveSubResourceData(m.Drives))
	d.Set("encrypted_secrets", m.EncryptedSecrets)
	d.Set("id", m.ID)
	d.Set("interfaces", SetAppInterfaceSubResourceData(m.Interfaces))
	d.Set("is_secret_updated", m.IsSecretUpdated)
	d.Set("logs", SetAppInstanceLogsSubResourceData([]*models.AppInstanceLogs{m.Logs}))
	d.Set("manifest_info", SetManifestInfoSubResourceData([]*models.ManifestInfo{m.ManifestInfo}))
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	d.Set("purge", SetZedCloudOpsCmdSubResourceData([]*models.ZedCloudOpsCmd{m.Purge}))
	d.Set("refresh", SetZedCloudOpsCmdSubResourceData([]*models.ZedCloudOpsCmd{m.Refresh}))
	d.Set("remote_console", m.RemoteConsole)
	d.Set("restart", SetZedCloudOpsCmdSubResourceData([]*models.ZedCloudOpsCmd{m.Restart}))
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("start_delay_in_seconds", m.StartDelayInSeconds)
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
	d.Set("user_data", m.UserData)
	d.Set("user_defined_version", m.UserDefinedVersion)
	d.Set("vminfo", SetVMSubResourceData([]*models.VM{m.Vminfo}))
}

func SetApplicationInstanceSubResourceData(m []*models.AppInstance) (d []*map[string]interface{}) {
	for _, AppInstanceModel := range m {
		if AppInstanceModel != nil {
			properties := make(map[string]interface{})
			properties["activate"] = AppInstanceModel.Activate
			properties["app_id"] = AppInstanceModel.AppID
			properties["app_policy_id"] = AppInstanceModel.AppPolicyID
			properties["app_type"] = AppInstanceModel.AppType
			properties["bundleversion"] = AppInstanceModel.Bundleversion
			properties["cluster_id"] = AppInstanceModel.ClusterID
			properties["collect_stats_ip_addr"] = AppInstanceModel.CollectStatsIPAddr
			properties["crypto_key"] = AppInstanceModel.CryptoKey
			properties["custom_config"] = SetCustomConfigSubResourceData([]*models.CustomConfig{AppInstanceModel.CustomConfig})
			properties["deployment_type"] = AppInstanceModel.DeploymentType
			properties["description"] = AppInstanceModel.Description
			properties["device_id"] = AppInstanceModel.DeviceID
			properties["drives"] = SetDriveSubResourceData(AppInstanceModel.Drives)
			properties["encrypted_secrets"] = AppInstanceModel.EncryptedSecrets
			properties["id"] = AppInstanceModel.ID
			properties["interfaces"] = SetAppInterfaceSubResourceData(AppInstanceModel.Interfaces)
			properties["is_secret_updated"] = AppInstanceModel.IsSecretUpdated
			properties["logs"] = SetAppInstanceLogsSubResourceData([]*models.AppInstanceLogs{AppInstanceModel.Logs})
			properties["manifest_info"] = SetManifestInfoSubResourceData([]*models.ManifestInfo{AppInstanceModel.ManifestInfo})
			properties["name"] = AppInstanceModel.Name
			properties["project_id"] = AppInstanceModel.ProjectID
			properties["purge"] = SetZedCloudOpsCmdSubResourceData([]*models.ZedCloudOpsCmd{AppInstanceModel.Purge})
			properties["refresh"] = SetZedCloudOpsCmdSubResourceData([]*models.ZedCloudOpsCmd{AppInstanceModel.Refresh})
			properties["remote_console"] = AppInstanceModel.RemoteConsole
			properties["restart"] = SetZedCloudOpsCmdSubResourceData([]*models.ZedCloudOpsCmd{AppInstanceModel.Restart})
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{AppInstanceModel.Revision})
			properties["start_delay_in_seconds"] = AppInstanceModel.StartDelayInSeconds
			properties["tags"] = AppInstanceModel.Tags
			properties["title"] = AppInstanceModel.Title
			properties["user_data"] = AppInstanceModel.UserData
			properties["user_defined_version"] = AppInstanceModel.UserDefinedVersion
			properties["vminfo"] = SetVMSubResourceData([]*models.VM{AppInstanceModel.Vminfo})
			d = append(d, &properties)
		}
	}
	return
}

func ApplicationInstance() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"activate": {
			Description: `app instance activation flag`,
			Type:        schema.TypeBool,
			Required:    true,
		},

		"app_id": {
			Description: `User defined name of the edge app, unique across the enterprise. Once app name is created, name can’t be changed`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"app_policy_id": {
			Description: `app policy id`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"app_type": {
			Description: `type of bundle`,
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "APP_TYPE_UNSPECIFIED",
		},

		"bundleversion": {
			Description: `version of bundle app is referring to`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"cluster_id": {
			Description: `System defined universally unique clusterInstance ID, unique across the enterprise.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"collect_stats_ip_addr": {
			Description: `holds the static Ip of the app instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"crypto_key": {
			Description: `Crypto Key for decrypting user secret information`,
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Sensitive:   true,
		},

		"custom_config": {
			Description: `Application initialization script template in cloud-config format and user specified values`,
			Type:        schema.TypeList, //GoType: CustomConfig
			Elem: &schema.Resource{
				Schema: CustomConfig(),
			},
			Optional: true,
		},

		"deployment_type": {
			Description: `type of deployment for the app, eg: azure, k3s, standalone`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"description": {
			Description: `Detailed description of the app instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"device_id": {
			Description: `User defined name of the device name, unique across the enterprise. Once device name is defined, name can’t be changed`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"drives": {
			Description: `drive details`,
			Type:        schema.TypeList, //GoType: []*Drive
			Elem: &schema.Resource{
				Schema: DriveSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"encrypted_secrets": {
			Description: `user encrypted secrets map`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional:  true,
			Computed:  true,
			Sensitive: true,
		},

		"id": {
			Description: `System defined universally unique Id of the app instance`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"interfaces": {
			Description: `application interfaces`,
			Type:        schema.TypeList, //GoType: []*AppInterface
			Elem: &schema.Resource{
				Schema: AppInterface(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional:         true,
			DiffSuppressFunc: diffSuppressInterfaceListOrder("interfaces"),
		},

		"is_secret_updated": {
			Description: `This field tells whether user secrets has updated or not, especially the cusotom config`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"logs": {
			Description: `App Instance logs`,
			Type:        schema.TypeList, //GoType: AppInstanceLogs
			Elem: &schema.Resource{
				Schema: AppInstanceLogs(),
			},
			Optional: true,
		},

		"manifest_info": {
			Description: `app manifest Info`,
			Type:        schema.TypeList, //GoType: ManifestInfo
			Elem: &schema.Resource{
				Schema: ManifestInfoSchema(),
			},
			Optional: true,
		},

		"name": {
			Description: `User defined name of the app instance, unique across the enterprise. Once app instance is created, name can’t be changed`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"project_id": {
			Description: `project name which the given app instance belong to`,
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
		},

		"purge": {
			Description: `Purge counter: ZedCloudOpsCmd`,
			Type:        schema.TypeList, //GoType: ZedCloudOpsCmd
			Elem: &schema.Resource{
				Schema: ZedCloudOpsCmdSchema(),
			},
			Optional: true,
			Computed: true,
		},

		"refresh": {
			Description: `Refresh counter: ZedCloudOpsCmd`,
			Type:        schema.TypeList, //GoType: ZedCloudOpsCmd
			Elem: &schema.Resource{
				Schema: ZedCloudOpsCmdSchema(),
			},
			Optional: true,
			Computed: true,
		},

		"remote_console": {
			Description: `Remote console flag`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"restart": {
			Description: `Restart counter: ZedCloudOpsCmd`,
			Type:        schema.TypeList, //GoType: ZedCloudOpsCmd
			Elem: &schema.Resource{
				Schema: ZedCloudOpsCmdSchema(),
			},
			Optional: true,
			Computed: true,
		},

		"revision": {
			Description: `app instance object revision details`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Optional: true,
			Computed: true,
		},

		"start_delay_in_seconds": {
			Description: `start delay is the time in seconds EVE should wait after boot before starting the application instance`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"tags": {
			Description: `Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"title": {
			Description: `User defined title of the app instance. Title can be changed at any time`,
			Type:        schema.TypeString,
			// Required:    true,
			Optional: true,
		},

		"user_data": {
			Description: `Deprecated: Application initiazation script in cloud-config format to be used by cloud-init`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"user_defined_version": {
			Description: `Instance version tells which edge app does this instance is running`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"vminfo": {
			Description: `virtual machine info`,
			Type:        schema.TypeList, //GoType: VM
			Elem: &schema.Resource{
				Schema: VMSchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the AppInstance resource
func GetApplicationInstancePropertyFields() (t []string) {
	return []string{
		"activate",
		"app_id",
		"app_policy_id",
		"app_type",
		"bundleversion",
		"cluster_id",
		"collect_stats_ip_addr",
		"crypto_key",
		"custom_config",
		"deployment_type",
		"description",
		"device_id",
		"drives",
		"encrypted_secrets",
		"id",
		"interfaces",
		"is_secret_updated",
		"logs",
		"manifest_info",
		"name",
		"project_id",
		"purge",
		"refresh",
		"remote_console",
		"restart",
		"revision",
		"start_delay_in_seconds",
		"tags",
		"title",
		"user_data",
		"user_defined_version",
		"vminfo",
	}
}
