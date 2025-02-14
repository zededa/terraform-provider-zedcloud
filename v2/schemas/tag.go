package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func TagModel(d *schema.ResourceData) *models.Tag {
	var appPolicy *models.Policy // AppPolicyConfig
	appPolicyInterface, appPolicyIsSet := d.GetOk("app_policy")
	if appPolicyIsSet && appPolicyInterface != nil {
		appPolicyMap := appPolicyInterface.([]interface{})
		if len(appPolicyMap) > 0 {
			appPolicy = PolicyConfigModelFromMap(appPolicyMap[0].(map[string]interface{}))
		}
	}
	var attestationPolicy *models.Policy // PolicyConfig
	attestationPolicyInterface, attestationPolicyIsSet := d.GetOk("attestation_policy")
	if attestationPolicyIsSet && attestationPolicyInterface != nil {
		attestationPolicyMap := attestationPolicyInterface.([]interface{})
		if len(attestationPolicyMap) > 0 {
			attestationPolicy = PolicyConfigModelFromMap(attestationPolicyMap[0].(map[string]interface{}))
		}
	}
	var configurationLockPolicy *models.Policy // Policy
	configurationLockPolicyInterface, configurationLockPolicyIsSet := d.GetOk("configuration_lock_policy")
	if configurationLockPolicyIsSet && configurationLockPolicyInterface != nil {
		configurationLockPolicyMap := configurationLockPolicyInterface.([]interface{})
		if len(configurationLockPolicyMap) > 0 {
			configurationLockPolicy = PolicyConfigModelFromMap(configurationLockPolicyMap[0].(map[string]interface{}))
		}
	}
	var deployment *models.Deployment // Deployment
	deploymentInterface, deploymentIsSet := d.GetOk("deployment")
	if deploymentIsSet && deploymentInterface != nil {
		deploymentMap := deploymentInterface.([]interface{})
		if len(deploymentMap) > 0 {
			deployment = DeploymentModelFromMap(deploymentMap[0].(map[string]interface{}))
		}
	}
	description, _ := d.Get("description").(string)
	var edgeviewPolicy *models.Policy // Policy
	edgeviewPolicyInterface, edgeviewPolicyIsSet := d.GetOk("edgeview_policy")
	if edgeviewPolicyIsSet && edgeviewPolicyInterface != nil {
		edgeviewPolicyMap := edgeviewPolicyInterface.([]interface{})
		if len(edgeviewPolicyMap) > 0 {
			edgeviewPolicy = PolicyConfigModelFromMap(edgeviewPolicyMap[0].(map[string]interface{}))
		}
	}
	id, _ := d.Get("id").(string)
	var localOperatorConsolePolicy *models.Policy // PolicyConfig
	localOperatorConsolePolicyInterface, localOperatorConsolePolicyIsSet := d.GetOk("local_operator_console_policy")
	if localOperatorConsolePolicyIsSet && localOperatorConsolePolicyInterface != nil {
		localOperatorConsolePolicyMap := localOperatorConsolePolicyInterface.([]interface{})
		if len(localOperatorConsolePolicyMap) > 0 {
			localOperatorConsolePolicy = PolicyConfigModelFromMap(localOperatorConsolePolicyMap[0].(map[string]interface{}))
		}
	}
	name, _ := d.Get("name").(string)
	var networkPolicy *models.Policy // PolicyConfig
	networkPolicyInterface, networkPolicyIsSet := d.GetOk("network_policy")
	if networkPolicyIsSet && networkPolicyInterface != nil {
		networkPolicyMap := networkPolicyInterface.([]interface{})
		if len(networkPolicyMap) > 0 {
			networkPolicy = PolicyConfigModelFromMap(networkPolicyMap[0].(map[string]interface{}))
		}
	}
	var tagLevelSettings *models.TagLevelSettings // TagLevelSettings
	tagLevelSettingsInterface, tagLevelSettingsIsSet := d.GetOk("tag_level_settings")
	if tagLevelSettingsIsSet && tagLevelSettingsInterface != nil {
		tagLevelSettingsMap := tagLevelSettingsInterface.([]interface{})
		if len(tagLevelSettingsMap) > 0 {
			tagLevelSettings = TagLevelSettingsModelFromMap(tagLevelSettingsMap[0].(map[string]interface{}))
		}
	}
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
	var typeVar *models.TagType // TagType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewTagType(models.TagType(typeModel))
	}
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := d.GetOk("revision")
	if revisionIsSet {
		revisionMap := revisionInterface.([]interface{})[0].(map[string]interface{})
		revision = ObjectRevisionModelFromMap(revisionMap)
	}
	return &models.Tag{
		AppPolicy:                  appPolicy,
		AttestationPolicy:          attestationPolicy,
		ConfigurationLockPolicy:    configurationLockPolicy,
		Deployment:                 deployment,
		Description:                description,
		EdgeviewPolicy:             edgeviewPolicy,
		ID:                         id,
		LocalOperatorConsolePolicy: localOperatorConsolePolicy,
		Name:                       &name, // string
		NetworkPolicy:              networkPolicy,
		Revision:                   revision,
		TagLevelSettings:           tagLevelSettings,
		Tags:                       tags,
		Title:                      &title, // string
		Type:                       typeVar,
	}
}

func TagModelFromMap(m map[string]interface{}) *models.Tag {
	var appPolicy *models.Policy // AppPolicyConfig
	appPolicyInterface, appPolicyIsSet := m["app_policy"]
	if appPolicyIsSet && appPolicyInterface != nil {
		appPolicyMap := appPolicyInterface.([]interface{})
		if len(appPolicyMap) > 0 {
			appPolicy = PolicyConfigModelFromMap(appPolicyMap[0].(map[string]interface{}))
		}
	}
	var attestationPolicy *models.Policy // PolicyConfig
	attestationPolicyInterface, attestationPolicyIsSet := m["attestation_policy"]
	if attestationPolicyIsSet && attestationPolicyInterface != nil {
		attestationPolicyMap := attestationPolicyInterface.([]interface{})
		if len(attestationPolicyMap) > 0 {
			attestationPolicy = PolicyConfigModelFromMap(attestationPolicyMap[0].(map[string]interface{}))
		}
	}
	//
	var configurationLockPolicy *models.Policy // Policy
	configurationLockPolicyInterface, configurationLockPolicyIsSet := m["configuration_lock_policy"]
	if configurationLockPolicyIsSet && configurationLockPolicyInterface != nil {
		configurationLockPolicyMap := configurationLockPolicyInterface.([]interface{})
		if len(configurationLockPolicyMap) > 0 {
			configurationLockPolicy = PolicyConfigModelFromMap(configurationLockPolicyMap[0].(map[string]interface{}))
		}
	}
	//
	var deployment *models.Deployment // Deployment
	deploymentInterface, deploymentIsSet := m["deployment"]
	if deploymentIsSet && deploymentInterface != nil {
		deploymentMap := deploymentInterface.([]interface{})
		if len(deploymentMap) > 0 {
			deployment = DeploymentModelFromMap(deploymentMap[0].(map[string]interface{}))
		}
	}
	//
	description := m["description"].(string)
	var edgeviewPolicy *models.Policy // PolicyConfig
	edgeviewPolicyInterface, edgeviewPolicyIsSet := m["edgeview_policy"]
	if edgeviewPolicyIsSet && edgeviewPolicyInterface != nil {
		edgeviewPolicyMap := edgeviewPolicyInterface.([]interface{})
		if len(edgeviewPolicyMap) > 0 {
			edgeviewPolicy = PolicyConfigModelFromMap(edgeviewPolicyMap[0].(map[string]interface{}))
		}
	}
	//
	id := m["id"].(string)
	var localOperatorConsolePolicy *models.Policy // PolicyConfig
	localOperatorConsolePolicyInterface, localOperatorConsolePolicyIsSet := m["local_operator_console_policy"]
	if localOperatorConsolePolicyIsSet && localOperatorConsolePolicyInterface != nil {
		localOperatorConsolePolicyMap := localOperatorConsolePolicyInterface.([]interface{})
		if len(localOperatorConsolePolicyMap) > 0 {
			localOperatorConsolePolicy = PolicyConfigModelFromMap(localOperatorConsolePolicyMap[0].(map[string]interface{}))
		}
	}
	//
	name := m["name"].(string)
	var networkPolicy *models.Policy // PolicyConfig
	networkPolicyInterface, networkPolicyIsSet := m["network_policy"]
	if networkPolicyIsSet && networkPolicyInterface != nil {
		networkPolicyMap := networkPolicyInterface.([]interface{})
		if len(networkPolicyMap) > 0 {
			networkPolicy = PolicyConfigModelFromMap(networkPolicyMap[0].(map[string]interface{}))
		}
	}
	//
	var tagLevelSettings *models.TagLevelSettings // TagLevelSettings
	tagLevelSettingsInterface, tagLevelSettingsIsSet := m["tag_level_settings"]
	if tagLevelSettingsIsSet && tagLevelSettingsInterface != nil {
		tagLevelSettingsMap := tagLevelSettingsInterface.([]interface{})
		if len(tagLevelSettingsMap) > 0 {
			tagLevelSettings = TagLevelSettingsModelFromMap(tagLevelSettingsMap[0].(map[string]interface{}))
		}
	}
	//
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
	var typeVar *models.TagType // TagType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewTagType(models.TagType(typeModel))
	}
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := m["revision"]
	if revisionIsSet && revisionInterface != nil {
		revisionMap := revisionInterface.([]interface{})
		if len(revisionMap) > 0 {
			revision = ObjectRevisionModelFromMap(revisionMap[0].(map[string]interface{}))
		}
	}
	return &models.Tag{
		AppPolicy:                  appPolicy,
		AttestationPolicy:          attestationPolicy,
		ConfigurationLockPolicy:    configurationLockPolicy,
		Deployment:                 deployment,
		Description:                description,
		EdgeviewPolicy:             edgeviewPolicy,
		ID:                         id,
		LocalOperatorConsolePolicy: localOperatorConsolePolicy,
		Name:                       &name,
		NetworkPolicy:              networkPolicy,
		Revision:                   revision,
		TagLevelSettings:           tagLevelSettings,
		Tags:                       tags,
		Title:                      &title,
		Type:                       typeVar,
	}
}

func SetTagResourceData(d *schema.ResourceData, m *models.Tag) {
	d.Set("app_policy", SetPolicyConfigSubResourceData([]*models.Policy{m.AppPolicy}))
	d.Set("attestation_policy", SetPolicyConfigSubResourceData([]*models.Policy{m.AttestationPolicy}))
	d.Set("attr", m.Attr)
	d.Set("cloud_policy", SetPolicyConfigSubResourceData([]*models.Policy{m.CloudPolicy}))
	d.Set("configuration_lock_policy", SetPolicyConfigSubResourceData([]*models.Policy{m.ConfigurationLockPolicy}))
	d.Set("deployment", SetDeploymentSubResourceData([]*models.Deployment{m.Deployment}))
	d.Set("description", m.Description)
	d.Set("edgeview_policy", SetPolicyConfigSubResourceData([]*models.Policy{m.EdgeviewPolicy}))
	d.Set("id", m.ID)
	d.Set("local_operator_console_policy", SetPolicyConfigSubResourceData([]*models.Policy{m.LocalOperatorConsolePolicy}))
	d.Set("module_policy", SetPolicyConfigSubResourceData(m.ModulePolicy))
	d.Set("name", m.Name)
	d.Set("network_policy", SetPolicyConfigSubResourceData([]*models.Policy{m.NetworkPolicy}))
	d.Set("numdevices", m.Numdevices)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("tag_level_settings", SetTagLevelSettingsSubResourceData([]*models.TagLevelSettings{m.TagLevelSettings}))
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
	d.Set("type", m.Type)
}

func SetTagSubResourceData(m []*models.Tag) (d []*map[string]interface{}) {
	for _, TagModel := range m {
		if TagModel != nil {
			properties := make(map[string]interface{})
			properties["app_policy"] = SetPolicyConfigSubResourceData([]*models.Policy{TagModel.AppPolicy})
			properties["attestation_policy"] = SetPolicyConfigSubResourceData([]*models.Policy{TagModel.AttestationPolicy})
			properties["attr"] = TagModel.Attr
			properties["cloud_policy"] = SetPolicyConfigSubResourceData([]*models.Policy{TagModel.CloudPolicy})
			properties["configuration_lock_policy"] = SetPolicyConfigSubResourceData([]*models.Policy{TagModel.ConfigurationLockPolicy})
			properties["deployment"] = SetDeploymentSubResourceData([]*models.Deployment{TagModel.Deployment})
			properties["description"] = TagModel.Description
			properties["edgeview_policy"] = SetPolicyConfigSubResourceData([]*models.Policy{TagModel.EdgeviewPolicy})
			properties["id"] = TagModel.ID
			properties["local_operator_console_policy"] = SetPolicyConfigSubResourceData([]*models.Policy{TagModel.LocalOperatorConsolePolicy})
			properties["module_policy"] = SetPolicyConfigSubResourceData(TagModel.ModulePolicy)
			properties["name"] = TagModel.Name
			properties["network_policy"] = SetPolicyConfigSubResourceData([]*models.Policy{TagModel.NetworkPolicy})
			properties["numdevices"] = TagModel.Numdevices
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{TagModel.Revision})
			properties["tag_level_settings"] = SetTagLevelSettingsSubResourceData([]*models.TagLevelSettings{TagModel.TagLevelSettings})
			properties["tags"] = TagModel.Tags
			properties["title"] = TagModel.Title
			properties["type"] = TagModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func Project() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_policy": {
			Description: `Resource group wide policy for edge applications to be deployed on all edge nodes on this resource group`,
			Type:        schema.TypeList, // GoType: PolicyConfig
			Elem: &schema.Resource{
				Schema: Policy(),
			},
			Optional: true,
		},

		"attestation_policy": {
			Description: `Attestation policy to enforce on all devices of this project`,
			Type:        schema.TypeList, // GoType: PolicyConfig
			Elem: &schema.Resource{
				Schema: Policy(),
			},
			Optional: true,
			Computed: true,
		},

		"attr": {
			Description: `Resource group wide configuration for edge nodes`,
			Type:        schema.TypeMap, // GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Computed: true,
		},

		"cloud_policy": {
			Description: `Resource group wide policy for Azure IoTEdge configuration to be applied to all edge applications`,
			Type:        schema.TypeList, // GoType: PolicyConfig
			Elem: &schema.Resource{
				Schema: Policy(),
			},
			Computed: true,
		},

		"configuration_lock_policy": {
			Description: `Configuration lock prevents users to send unintentional misconfigurations`,
			Type:        schema.TypeList, // GoType: Policy
			Elem: &schema.Resource{
				Schema: Policy(),
			},
			Optional: true,
		},

		"deployment": {
			Description: `Deployment template containing different types of policies`,
			Type:        schema.TypeList, // GoType: Deployment
			Elem: &schema.Resource{
				Schema: DeploymentSchema(),
			},
			Optional: true,
		},

		"description": {
			Description: `Detailed description of the resource group.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"edgeview_policy": {
			Description: `Edgeview policy on devices of this project`,
			Type:        schema.TypeList, // GoType: PolicyConfig
			Elem: &schema.Resource{
				Schema: Policy(),
			},
			Optional: true,
		},

		"id": {
			Description: `System defined universally unique Id of the resource group.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"local_operator_console_policy": {
			Description: `Local operator console policy on devices of this project`,
			Type:        schema.TypeList, // GoType: PolicyConfig
			Elem: &schema.Resource{
				Schema: Policy(),
			},
			Optional: true,
		},

		"module_policy": {
			Description: `Resource group wide policy for Azure module configuration to be applied to all edge applications`,
			Type:        schema.TypeList, // GoType: []*PolicyConfig
			Elem: &schema.Resource{
				Schema: Policy(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Computed: true,
		},

		"name": {
			Description: `User defined name of the resource group, unique across the enterprise. Once resource group is created, name can’t be changed.`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"network_policy": {
			Description: `Network policy to enforce on all devices of this project`,
			Type:        schema.TypeList, // GoType: PolicyConfig
			Elem: &schema.Resource{
				Schema: Policy(),
			},
			Optional: true,
		},

		"numdevices": {
			Description: `Number of edge nodes in this resource group`,
			Type:        schema.TypeInt,
			Computed:    true,
		},

		"revision": {
			Description: `system defined info`,
			Type:        schema.TypeList, // GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Computed: true,
		},

		"tag_level_settings": {
			Description: `tag level setting within a enterprise`,
			Type:        schema.TypeList, // GoType: TagLevelSettings
			Elem: &schema.Resource{
				Schema: TagLevelSettingsSchema(),
			},
			Optional: true,
			DiffSuppressFunc: diffSupressMapInterfaceNonConfigChangesV2("tag_level_settings",
				[]string{"flow_log_transmission", "interface_ordering"}),
		},

		"tags": {
			Description: `Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.`,
			Type:        schema.TypeMap, // GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"title": {
			Description: `User defined title of the resource group. Title can be changed at any time.`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"type": {
			Description: `Resource group type`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func GetTagPropertyFields() (t []string) {
	return []string{
		"app_policy",
		"attestation_policy",
		"configuration_lock_policy",
		"deployment",
		"description",
		"edgeview_policy",
		"id",
		"local_operator_console_policy",
		"name",
		"network_policy",
		"tag_level_settings",
		"tags",
		"title",
		"type",
	}
}
