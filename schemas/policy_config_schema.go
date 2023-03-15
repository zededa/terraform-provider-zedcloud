package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func PolicyConfigModel(d *schema.ResourceData) *models.Policy {
	var appPolicy *models.AppPolicy // AppPolicy
	appPolicyInterface, appPolicyIsSet := d.GetOk("app_policy")
	if appPolicyIsSet && appPolicyInterface != nil {
		appPolicyMap := appPolicyInterface.([]interface{})
		if len(appPolicyMap) > 0 {
			appPolicy = AppPolicyModelFromMap(appPolicyMap[0].(map[string]interface{}))
		}
	}
	var attestationPolicy *models.AttestationPolicy // AttestationPolicy
	attestationPolicyInterface, attestationPolicyIsSet := d.GetOk("attestation_policy")
	if attestationPolicyIsSet && attestationPolicyInterface != nil {
		attestationPolicyMap := attestationPolicyInterface.([]interface{})
		if len(attestationPolicyMap) > 0 {
			attestationPolicy = AttestationPolicyModelFromMap(attestationPolicyMap[0].(map[string]interface{}))
		}
	}
	attr := map[string]string{}
	attrInterface, attrIsSet := d.GetOk("attr")
	if attrIsSet {
		attrMap := attrInterface.(map[string]interface{})
		for k, v := range attrMap {
			if v == nil {
				continue
			}
			attr[k] = v.(string)
		}
	}

	var azurePolicy *models.AzurePolicy // AzurePolicy
	azurePolicyInterface, azurePolicyIsSet := d.GetOk("azure_policy")
	if azurePolicyIsSet && azurePolicyInterface != nil {
		azurePolicyMap := azurePolicyInterface.([]interface{})
		if len(azurePolicyMap) > 0 {
			azurePolicy = AzurePolicyModelFromMap(azurePolicyMap[0].(map[string]interface{}))
		}
	}
	var clusterPolicy *models.ClusterPolicy // ClusterPolicy
	clusterPolicyInterface, clusterPolicyIsSet := d.GetOk("cluster_policy")
	if clusterPolicyIsSet && clusterPolicyInterface != nil {
		clusterPolicyMap := clusterPolicyInterface.([]interface{})
		if len(clusterPolicyMap) > 0 {
			clusterPolicy = ClusterPolicyModelFromMap(clusterPolicyMap[0].(map[string]interface{}))
		}
	}
	description, _ := d.Get("description").(string)
	var edgeviewPolicy *models.EdgeviewPolicy // EdgeviewPolicy
	edgeviewPolicyInterface, edgeviewPolicyIsSet := d.GetOk("edgeview_policy")
	if edgeviewPolicyIsSet && edgeviewPolicyInterface != nil {
		edgeviewPolicyMap := edgeviewPolicyInterface.([]interface{})
		if len(edgeviewPolicyMap) > 0 {
			edgeviewPolicy = EdgeviewPolicyModelFromMap(edgeviewPolicyMap[0].(map[string]interface{}))
		}
	}
	id, _ := d.Get("id").(string)
	var localOperatorConsolePolicy *models.LocalOperatorConsolePolicy // LocalOperatorConsolePolicy
	localOperatorConsolePolicyInterface, localOperatorConsolePolicyIsSet := d.GetOk("local_operator_console_policy")
	if localOperatorConsolePolicyIsSet && localOperatorConsolePolicyInterface != nil {
		localOperatorConsolePolicyMap := localOperatorConsolePolicyInterface.([]interface{})
		if len(localOperatorConsolePolicyMap) > 0 {
			localOperatorConsolePolicy = LocalOperatorConsolePolicyModelFromMap(localOperatorConsolePolicyMap[0].(map[string]interface{}))
		}
	}
	var modulePolicy *models.ModulePolicy // ModulePolicy
	modulePolicyInterface, modulePolicyIsSet := d.GetOk("module_policy")
	if modulePolicyIsSet && modulePolicyInterface != nil {
		modulePolicyMap := modulePolicyInterface.([]interface{})
		if len(modulePolicyMap) > 0 {
			modulePolicy = ModulePolicyModelFromMap(modulePolicyMap[0].(map[string]interface{}))
		}
	}
	name, _ := d.Get("name").(string)
	var networkPolicy *models.NetworkPolicy // NetworkPolicy
	networkPolicyInterface, networkPolicyIsSet := d.GetOk("network_policy")
	if networkPolicyIsSet && networkPolicyInterface != nil {
		networkPolicyMap := networkPolicyInterface.([]interface{})
		if len(networkPolicyMap) > 0 {
			networkPolicy = NetworkPolicyModelFromMap(networkPolicyMap[0].(map[string]interface{}))
		}
	}
	statusMessage, _ := d.Get("status_message").(string)
	title, _ := d.Get("title").(string)
	var typeVar *models.PolicyType // PolicyType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewPolicyType(models.PolicyType(typeModel))
	}
	return &models.Policy{
		AppPolicy:                  appPolicy,
		AttestationPolicy:          attestationPolicy,
		Attr:                       attr,
		AzurePolicy:                azurePolicy,
		ClusterPolicy:              clusterPolicy,
		Description:                description,
		EdgeviewPolicy:             edgeviewPolicy,
		ID:                         id,
		LocalOperatorConsolePolicy: localOperatorConsolePolicy,
		ModulePolicy:               modulePolicy,
		Name:                       &name, // string true false false
		NetworkPolicy:              networkPolicy,
		StatusMessage:              statusMessage,
		Title:                      &title, // string true false false
		Type:                       typeVar,
	}
}

func PolicyConfigModelFromMap(m map[string]interface{}) *models.Policy {
	var appPolicy *models.AppPolicy // AppPolicy
	appPolicyInterface, appPolicyIsSet := m["app_policy"]
	if appPolicyIsSet && appPolicyInterface != nil {
		appPolicyMap := appPolicyInterface.([]interface{})
		if len(appPolicyMap) > 0 {
			appPolicy = AppPolicyModelFromMap(appPolicyMap[0].(map[string]interface{}))
		}
	}
	//
	var attestationPolicy *models.AttestationPolicy // AttestationPolicy
	attestationPolicyInterface, attestationPolicyIsSet := m["attestation_policy"]
	if attestationPolicyIsSet && attestationPolicyInterface != nil {
		attestationPolicyMap := attestationPolicyInterface.([]interface{})
		if len(attestationPolicyMap) > 0 {
			attestationPolicy = AttestationPolicyModelFromMap(attestationPolicyMap[0].(map[string]interface{}))
		}
	}
	//
	attr := map[string]string{}
	attrInterface, attrIsSet := m["attr"]
	if attrIsSet {
		attrMap := attrInterface.(map[string]interface{})
		for k, v := range attrMap {
			if v == nil {
				continue
			}
			attr[k] = v.(string)
		}
	}

	var azurePolicy *models.AzurePolicy // AzurePolicy
	azurePolicyInterface, azurePolicyIsSet := m["azure_policy"]
	if azurePolicyIsSet && azurePolicyInterface != nil {
		azurePolicyMap := azurePolicyInterface.([]interface{})
		if len(azurePolicyMap) > 0 {
			azurePolicy = AzurePolicyModelFromMap(azurePolicyMap[0].(map[string]interface{}))
		}
	}
	//
	var clusterPolicy *models.ClusterPolicy // ClusterPolicy
	clusterPolicyInterface, clusterPolicyIsSet := m["cluster_policy"]
	if clusterPolicyIsSet && clusterPolicyInterface != nil {
		clusterPolicyMap := clusterPolicyInterface.([]interface{})
		if len(clusterPolicyMap) > 0 {
			clusterPolicy = ClusterPolicyModelFromMap(clusterPolicyMap[0].(map[string]interface{}))
		}
	}
	//
	description := m["description"].(string)
	var edgeviewPolicy *models.EdgeviewPolicy // EdgeviewPolicy
	edgeviewPolicyInterface, edgeviewPolicyIsSet := m["edgeview_policy"]
	if edgeviewPolicyIsSet && edgeviewPolicyInterface != nil {
		edgeviewPolicyMap := edgeviewPolicyInterface.([]interface{})
		if len(edgeviewPolicyMap) > 0 {
			edgeviewPolicy = EdgeviewPolicyModelFromMap(edgeviewPolicyMap[0].(map[string]interface{}))
		}
	}
	//
	id := m["id"].(string)
	var localOperatorConsolePolicy *models.LocalOperatorConsolePolicy // LocalOperatorConsolePolicy
	localOperatorConsolePolicyInterface, localOperatorConsolePolicyIsSet := m["local_operator_console_policy"]
	if localOperatorConsolePolicyIsSet && localOperatorConsolePolicyInterface != nil {
		localOperatorConsolePolicyMap := localOperatorConsolePolicyInterface.([]interface{})
		if len(localOperatorConsolePolicyMap) > 0 {
			localOperatorConsolePolicy = LocalOperatorConsolePolicyModelFromMap(localOperatorConsolePolicyMap[0].(map[string]interface{}))
		}
	}
	//
	var modulePolicy *models.ModulePolicy // ModulePolicy
	modulePolicyInterface, modulePolicyIsSet := m["module_policy"]
	if modulePolicyIsSet && modulePolicyInterface != nil {
		modulePolicyMap := modulePolicyInterface.([]interface{})
		if len(modulePolicyMap) > 0 {
			modulePolicy = ModulePolicyModelFromMap(modulePolicyMap[0].(map[string]interface{}))
		}
	}
	//
	name := m["name"].(string)
	var networkPolicy *models.NetworkPolicy // NetworkPolicy
	networkPolicyInterface, networkPolicyIsSet := m["network_policy"]
	if networkPolicyIsSet && networkPolicyInterface != nil {
		networkPolicyMap := networkPolicyInterface.([]interface{})
		if len(networkPolicyMap) > 0 {
			networkPolicy = NetworkPolicyModelFromMap(networkPolicyMap[0].(map[string]interface{}))
		}
	}
	//
	statusMessage := m["status_message"].(string)
	title := m["title"].(string)
	var typeVar *models.PolicyType // PolicyType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewPolicyType(models.PolicyType(typeModel))
	}
	return &models.Policy{
		AppPolicy:                  appPolicy,
		AttestationPolicy:          attestationPolicy,
		Attr:                       attr,
		AzurePolicy:                azurePolicy,
		ClusterPolicy:              clusterPolicy,
		Description:                description,
		EdgeviewPolicy:             edgeviewPolicy,
		ID:                         id,
		LocalOperatorConsolePolicy: localOperatorConsolePolicy,
		ModulePolicy:               modulePolicy,
		Name:                       &name,
		NetworkPolicy:              networkPolicy,
		StatusMessage:              statusMessage,
		Title:                      &title,
		Type:                       typeVar,
	}
}

func SetPolicyConfigResourceData(d *schema.ResourceData, m *models.Policy) {
	d.Set("app_policy", SetAppPolicySubResourceData([]*models.AppPolicy{m.AppPolicy}))
	d.Set("attestation_policy", SetAttestationPolicySubResourceData([]*models.AttestationPolicy{m.AttestationPolicy}))
	d.Set("attr", m.Attr)
	d.Set("azure_policy", SetAzurePolicySubResourceData([]*models.AzurePolicy{m.AzurePolicy}))
	d.Set("cluster_policy", SetClusterPolicySubResourceData([]*models.ClusterPolicy{m.ClusterPolicy}))
	d.Set("description", m.Description)
	d.Set("edgeview_policy", SetEdgeviewPolicySubResourceData([]*models.EdgeviewPolicy{m.EdgeviewPolicy}))
	d.Set("id", m.ID)
	d.Set("local_operator_console_policy", SetLocalOperatorConsolePolicySubResourceData([]*models.LocalOperatorConsolePolicy{m.LocalOperatorConsolePolicy}))
	d.Set("module_policy", SetModulePolicySubResourceData([]*models.ModulePolicy{m.ModulePolicy}))
	d.Set("name", m.Name)
	d.Set("network_policy", SetNetworkPolicySubResourceData([]*models.NetworkPolicy{m.NetworkPolicy}))
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("status", m.Status)
	d.Set("status_message", m.StatusMessage)
	d.Set("title", m.Title)
	d.Set("type", m.Type)
}

func SetPolicyConfigSubResourceData(m []*models.Policy) (d []*map[string]interface{}) {
	for _, PolicyConfigModel := range m {
		if PolicyConfigModel != nil {
			properties := make(map[string]interface{})
			properties["app_policy"] = SetAppPolicySubResourceData([]*models.AppPolicy{PolicyConfigModel.AppPolicy})
			properties["attestation_policy"] = SetAttestationPolicySubResourceData([]*models.AttestationPolicy{PolicyConfigModel.AttestationPolicy})
			properties["attr"] = PolicyConfigModel.Attr
			properties["azure_policy"] = SetAzurePolicySubResourceData([]*models.AzurePolicy{PolicyConfigModel.AzurePolicy})
			properties["cluster_policy"] = SetClusterPolicySubResourceData([]*models.ClusterPolicy{PolicyConfigModel.ClusterPolicy})
			properties["description"] = PolicyConfigModel.Description
			properties["edgeview_policy"] = SetEdgeviewPolicySubResourceData([]*models.EdgeviewPolicy{PolicyConfigModel.EdgeviewPolicy})
			properties["id"] = PolicyConfigModel.ID
			properties["local_operator_console_policy"] = SetLocalOperatorConsolePolicySubResourceData([]*models.LocalOperatorConsolePolicy{PolicyConfigModel.LocalOperatorConsolePolicy})
			properties["module_policy"] = SetModulePolicySubResourceData([]*models.ModulePolicy{PolicyConfigModel.ModulePolicy})
			properties["name"] = PolicyConfigModel.Name
			properties["network_policy"] = SetNetworkPolicySubResourceData([]*models.NetworkPolicy{PolicyConfigModel.NetworkPolicy})
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{PolicyConfigModel.Revision})
			properties["status"] = PolicyConfigModel.Status
			properties["status_message"] = PolicyConfigModel.StatusMessage
			properties["title"] = PolicyConfigModel.Title
			properties["type"] = PolicyConfigModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func Policy() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_policy": {
			Description: `app policy, which is used in auto app instance deployment`,
			Type:        schema.TypeList, //GoType: AppPolicy
			Elem: &schema.Resource{
				Schema: AppPolicy(),
			},
			Optional: true,
		},

		"attestation_policy": {
			Description: `attestation policy to enforce on all devices in this project`,
			Type:        schema.TypeList, //GoType: AttestationPolicy
			Elem: &schema.Resource{
				Schema: AttestationPolicy(),
			},
			Optional: true,
		},

		"attr": {
			Description: `Mapping of policy  variable keys and policy variable values`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"azure_policy": {
			Description: `azure policy, which is used in configuring azure iot-edge.`,
			Type:        schema.TypeList, //GoType: AzurePolicy
			Elem: &schema.Resource{
				Schema: AzurePolicy(),
			},
			Optional: true,
		},

		"cluster_policy": {
			Description: `cluster policy to bring up cluster on devices in this project`,
			Type:        schema.TypeList, //GoType: ClusterPolicy
			Elem: &schema.Resource{
				Schema: ClusterPolicy(),
			},
			Optional: true,
		},

		"description": {
			Description: `Detailed description of the policy`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"edgeview_policy": {
			Description: `edgeview policy on devices of this project`,
			Type:        schema.TypeList, //GoType: EdgeviewPolicy
			Elem: &schema.Resource{
				Schema: EdgeviewPolicy(),
			},
			Optional: true,
		},

		"id": {
			Description: `System defined universally unique Id of the policy request`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"local_operator_console_policy": {
			Description: `local operator console policy to enforce on all devices in this project`,
			Type:        schema.TypeList, //GoType: LocalOperatorConsolePolicy
			Elem: &schema.Resource{
				Schema: LocalOperatorConsolePolicy(),
			},
			Optional: true,
		},

		"module_policy": {
			Description: `module policy, which is used in auto module deployment`,
			Type:        schema.TypeList, //GoType: ModulePolicy
			Elem: &schema.Resource{
				Schema: ModulePolicy(),
			},
			Optional: true,
		},

		"name": {
			Description: `User defined name of the policy request, unique across the enterprise. Once object is created, name can’t be changed`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"network_policy": {
			Description: `network policy to enforce on all devices in this project`,
			Type:        schema.TypeList, //GoType: NetworkPolicy
			Elem: &schema.Resource{
				Schema: NetworkPolicy(),
			},
			Optional: true,
		},

		"revision": {
			Description: `system defined info`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Computed: true,
		},

		"status": {
			Description: `status of the policy`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"status_message": {
			Description: `Detailed status message of the policy`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"title": {
			Description: `User defined title of the policy. Title can be changed at any time`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"type": {
			Description: `type of policy`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetPolicyConfigPropertyFields() (t []string) {
	return []string{
		"app_policy",
		"attestation_policy",
		"attr",
		"azure_policy",
		"cluster_policy",
		"description",
		"edgeview_policy",
		"id",
		"local_operator_console_policy",
		"module_policy",
		"name",
		"network_policy",
		"status_message",
		"title",
		"type",
	}
}
