package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate PolicyConfig resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PolicyConfigModel(d *schema.ResourceData) *models.PolicyConfig {
	var appPolicy *models.AppPolicy // AppPolicy
	appPolicyInterface, appPolicyIsSet := d.GetOk("app_policy")
	if appPolicyIsSet {
		appPolicyMap := appPolicyInterface.([]interface{})[0].(map[string]interface{})
		appPolicy = AppPolicyModelFromMap(appPolicyMap)
	}
	var attestationPolicy *models.AttestationPolicy // AttestationPolicy
	attestationPolicyInterface, attestationPolicyIsSet := d.GetOk("attestation_policy")
	if attestationPolicyIsSet {
		attestationPolicyMap := attestationPolicyInterface.([]interface{})[0].(map[string]interface{})
		attestationPolicy = AttestationPolicyModelFromMap(attestationPolicyMap)
	}
	attr, _ := d.Get("attr").(map[string]string) // map[string]string
	var azurePolicy *models.AzurePolicy          // AzurePolicy
	azurePolicyInterface, azurePolicyIsSet := d.GetOk("azure_policy")
	if azurePolicyIsSet {
		azurePolicyMap := azurePolicyInterface.([]interface{})[0].(map[string]interface{})
		azurePolicy = AzurePolicyModelFromMap(azurePolicyMap)
	}
	var clusterPolicy *models.ClusterPolicy // ClusterPolicy
	clusterPolicyInterface, clusterPolicyIsSet := d.GetOk("cluster_policy")
	if clusterPolicyIsSet {
		clusterPolicyMap := clusterPolicyInterface.([]interface{})[0].(map[string]interface{})
		clusterPolicy = ClusterPolicyModelFromMap(clusterPolicyMap)
	}
	description, _ := d.Get("description").(string)
	var edgeviewPolicy *models.EdgeviewPolicy // EdgeviewPolicy
	edgeviewPolicyInterface, edgeviewPolicyIsSet := d.GetOk("edgeview_policy")
	if edgeviewPolicyIsSet {
		edgeviewPolicyMap := edgeviewPolicyInterface.([]interface{})[0].(map[string]interface{})
		edgeviewPolicy = EdgeviewPolicyModelFromMap(edgeviewPolicyMap)
	}
	id, _ := d.Get("id").(string)
	var modulePolicy *models.ModulePolicy // ModulePolicy
	modulePolicyInterface, modulePolicyIsSet := d.GetOk("module_policy")
	if modulePolicyIsSet {
		modulePolicyMap := modulePolicyInterface.([]interface{})[0].(map[string]interface{})
		modulePolicy = ModulePolicyModelFromMap(modulePolicyMap)
	}
	name, _ := d.Get("name").(string)
	var networkPolicy *models.NetworkPolicy // NetworkPolicy
	networkPolicyInterface, networkPolicyIsSet := d.GetOk("network_policy")
	if networkPolicyIsSet {
		networkPolicyMap := networkPolicyInterface.([]interface{})[0].(map[string]interface{})
		networkPolicy = NetworkPolicyModelFromMap(networkPolicyMap)
	}
	statusMessage, _ := d.Get("status_message").(string)
	title, _ := d.Get("title").(string)
	var typeVar *models.PolicyType // PolicyType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(models.PolicyType)
		typeVar = &typeModel
	}
	return &models.PolicyConfig{
		AppPolicy:         appPolicy,
		AttestationPolicy: attestationPolicy,
		Attr:              attr,
		AzurePolicy:       azurePolicy,
		ClusterPolicy:     clusterPolicy,
		Description:       description,
		EdgeviewPolicy:    edgeviewPolicy,
		ID:                id,
		ModulePolicy:      modulePolicy,
		Name:              &name, // string true false false
		NetworkPolicy:     networkPolicy,
		StatusMessage:     statusMessage,
		Title:             &title, // string true false false
		Type:              typeVar,
	}
}

func PolicyConfigModelFromMap(m map[string]interface{}) *models.PolicyConfig {
	var appPolicy *models.AppPolicy // AppPolicy
	appPolicyInterface, appPolicyIsSet := m["app_policy"]
	if appPolicyIsSet {
		appPolicyMap := appPolicyInterface.([]interface{})[0].(map[string]interface{})
		appPolicy = AppPolicyModelFromMap(appPolicyMap)
	}
	//
	var attestationPolicy *models.AttestationPolicy // AttestationPolicy
	attestationPolicyInterface, attestationPolicyIsSet := m["attestation_policy"]
	if attestationPolicyIsSet {
		attestationPolicyMap := attestationPolicyInterface.([]interface{})[0].(map[string]interface{})
		attestationPolicy = AttestationPolicyModelFromMap(attestationPolicyMap)
	}
	//
	attr := m["attr"].(map[string]string)
	var azurePolicy *models.AzurePolicy // AzurePolicy
	azurePolicyInterface, azurePolicyIsSet := m["azure_policy"]
	if azurePolicyIsSet {
		azurePolicyMap := azurePolicyInterface.([]interface{})[0].(map[string]interface{})
		azurePolicy = AzurePolicyModelFromMap(azurePolicyMap)
	}
	//
	var clusterPolicy *models.ClusterPolicy // ClusterPolicy
	clusterPolicyInterface, clusterPolicyIsSet := m["cluster_policy"]
	if clusterPolicyIsSet {
		clusterPolicyMap := clusterPolicyInterface.([]interface{})[0].(map[string]interface{})
		clusterPolicy = ClusterPolicyModelFromMap(clusterPolicyMap)
	}
	//
	description := m["description"].(string)
	var edgeviewPolicy *models.EdgeviewPolicy // EdgeviewPolicy
	edgeviewPolicyInterface, edgeviewPolicyIsSet := m["edgeview_policy"]
	if edgeviewPolicyIsSet {
		edgeviewPolicyMap := edgeviewPolicyInterface.([]interface{})[0].(map[string]interface{})
		edgeviewPolicy = EdgeviewPolicyModelFromMap(edgeviewPolicyMap)
	}
	//
	id := m["id"].(string)
	var modulePolicy *models.ModulePolicy // ModulePolicy
	modulePolicyInterface, modulePolicyIsSet := m["module_policy"]
	if modulePolicyIsSet {
		modulePolicyMap := modulePolicyInterface.([]interface{})[0].(map[string]interface{})
		modulePolicy = ModulePolicyModelFromMap(modulePolicyMap)
	}
	//
	name := m["name"].(string)
	var networkPolicy *models.NetworkPolicy // NetworkPolicy
	networkPolicyInterface, networkPolicyIsSet := m["network_policy"]
	if networkPolicyIsSet {
		networkPolicyMap := networkPolicyInterface.([]interface{})[0].(map[string]interface{})
		networkPolicy = NetworkPolicyModelFromMap(networkPolicyMap)
	}
	//
	statusMessage := m["status_message"].(string)
	title := m["title"].(string)
	typeVar := m["type"].(*models.PolicyType) // PolicyType
	return &models.PolicyConfig{
		AppPolicy:         appPolicy,
		AttestationPolicy: attestationPolicy,
		Attr:              attr,
		AzurePolicy:       azurePolicy,
		ClusterPolicy:     clusterPolicy,
		Description:       description,
		EdgeviewPolicy:    edgeviewPolicy,
		ID:                id,
		ModulePolicy:      modulePolicy,
		Name:              &name,
		NetworkPolicy:     networkPolicy,
		StatusMessage:     statusMessage,
		Title:             &title,
		Type:              typeVar,
	}
}

// Update the underlying PolicyConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPolicyConfigResourceData(d *schema.ResourceData, m *models.PolicyConfig) {
	d.Set("app_policy", SetAppPolicySubResourceData([]*models.AppPolicy{m.AppPolicy}))
	d.Set("attestation_policy", SetAttestationPolicySubResourceData([]*models.AttestationPolicy{m.AttestationPolicy}))
	d.Set("attr", m.Attr)
	d.Set("azure_policy", SetAzurePolicySubResourceData([]*models.AzurePolicy{m.AzurePolicy}))
	d.Set("cluster_policy", SetClusterPolicySubResourceData([]*models.ClusterPolicy{m.ClusterPolicy}))
	d.Set("description", m.Description)
	d.Set("edgeview_policy", SetEdgeviewPolicySubResourceData([]*models.EdgeviewPolicy{m.EdgeviewPolicy}))
	d.Set("id", m.ID)
	d.Set("module_policy", SetModulePolicySubResourceData([]*models.ModulePolicy{m.ModulePolicy}))
	d.Set("name", m.Name)
	d.Set("network_policy", SetNetworkPolicySubResourceData([]*models.NetworkPolicy{m.NetworkPolicy}))
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("status", m.Status)
	d.Set("status_message", m.StatusMessage)
	d.Set("title", m.Title)
	d.Set("type", m.Type)
}

// Iterate through and update the PolicyConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPolicyConfigSubResourceData(m []*models.PolicyConfig) (d []*map[string]interface{}) {
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

// Schema mapping representing the PolicyConfig resource defined in the Terraform configuration
func PolicyConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_policy": {
			Description: `app policy, which is used in auto app instance deployment`,
			Type:        schema.TypeList, //GoType: AppPolicy
			Elem: &schema.Resource{
				Schema: AppPolicySchema(),
			},
			Optional: true,
		},

		"attestation_policy": {
			Description: `attestation policy to enforce on all devices in this project`,
			Type:        schema.TypeList, //GoType: AttestationPolicy
			Elem: &schema.Resource{
				Schema: AttestationPolicySchema(),
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
				Schema: AzurePolicySchema(),
			},
			Optional: true,
		},

		"cluster_policy": {
			Description: `cluster policy to bring up cluster on devices in this project`,
			Type:        schema.TypeList, //GoType: ClusterPolicy
			Elem: &schema.Resource{
				Schema: ClusterPolicySchema(),
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
				Schema: EdgeviewPolicySchema(),
			},
			Optional: true,
		},

		"id": {
			Description: `System defined universally unique Id of the policy request`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"module_policy": {
			Description: `module policy, which is used in auto module deployment`,
			Type:        schema.TypeList, //GoType: ModulePolicy
			Elem: &schema.Resource{
				Schema: ModulePolicySchema(),
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
				Schema: NetworkPolicySchema(),
			},
			Optional: true,
		},

		"revision": {
			Description: `system defined info`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevisionSchema(),
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
			Required:    true,
		},

		"type": {
			Description: `type of policy`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

// Retrieve property field names for updating the PolicyConfig resource
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
		"module_policy",
		"name",
		"network_policy",
		"status_message",
		"title",
		"type",
	}
}
