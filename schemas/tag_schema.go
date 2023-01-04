package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate Tag resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func TagModel(d *schema.ResourceData) *models.Tag {
	var attestationPolicy *models.PolicyConfig // PolicyConfig
	attestationPolicyInterface, attestationPolicyIsSet := d.GetOk("attestation_policy")
	if attestationPolicyIsSet {
		attestationPolicyMap := attestationPolicyInterface.([]interface{})[0].(map[string]interface{})
		attestationPolicy = PolicyConfigModelFromMap(attestationPolicyMap)
	}
	var deployment *models.Deployment // Deployment
	deploymentInterface, deploymentIsSet := d.GetOk("deployment")
	if deploymentIsSet {
		deploymentMap := deploymentInterface.([]interface{})[0].(map[string]interface{})
		deployment = DeploymentModelFromMap(deploymentMap)
	}
	description, _ := d.Get("description").(string)
	var edgeviewPolicy *models.PolicyConfig // PolicyConfig
	edgeviewPolicyInterface, edgeviewPolicyIsSet := d.GetOk("edgeview_policy")
	if edgeviewPolicyIsSet {
		edgeviewPolicyMap := edgeviewPolicyInterface.([]interface{})[0].(map[string]interface{})
		edgeviewPolicy = PolicyConfigModelFromMap(edgeviewPolicyMap)
	}
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	var networkPolicy *models.PolicyConfig // PolicyConfig
	networkPolicyInterface, networkPolicyIsSet := d.GetOk("network_policy")
	if networkPolicyIsSet {
		networkPolicyMap := networkPolicyInterface.([]interface{})[0].(map[string]interface{})
		networkPolicy = PolicyConfigModelFromMap(networkPolicyMap)
	}
	title, _ := d.Get("title").(string)
	var typeVar *models.TagType // TagType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewTagType(models.TagType(typeModel))
	}
	return &models.Tag{
		AttestationPolicy: attestationPolicy,
		Deployment:        deployment,
		Description:       description,
		EdgeviewPolicy:    edgeviewPolicy,
		ID:                id,
		Name:              &name, // string true false false
		NetworkPolicy:     networkPolicy,
		Title:             &title, // string true false false
		Type:              typeVar,
	}
}

func TagModelFromMap(m map[string]interface{}) *models.Tag {
	var attestationPolicy *models.PolicyConfig // PolicyConfig
	attestationPolicyInterface, attestationPolicyIsSet := m["attestation_policy"]
	if attestationPolicyIsSet {
		attestationPolicyMap := attestationPolicyInterface.([]interface{})[0].(map[string]interface{})
		attestationPolicy = PolicyConfigModelFromMap(attestationPolicyMap)
	}
	//
	var deployment *models.Deployment // Deployment
	deploymentInterface, deploymentIsSet := m["deployment"]
	if deploymentIsSet {
		deploymentMap := deploymentInterface.([]interface{})[0].(map[string]interface{})
		deployment = DeploymentModelFromMap(deploymentMap)
	}
	//
	description := m["description"].(string)
	var edgeviewPolicy *models.PolicyConfig // PolicyConfig
	edgeviewPolicyInterface, edgeviewPolicyIsSet := m["edgeview_policy"]
	if edgeviewPolicyIsSet {
		edgeviewPolicyMap := edgeviewPolicyInterface.([]interface{})[0].(map[string]interface{})
		edgeviewPolicy = PolicyConfigModelFromMap(edgeviewPolicyMap)
	}
	//
	id := m["id"].(string)
	name := m["name"].(string)
	var networkPolicy *models.PolicyConfig // PolicyConfig
	networkPolicyInterface, networkPolicyIsSet := m["network_policy"]
	if networkPolicyIsSet {
		networkPolicyMap := networkPolicyInterface.([]interface{})[0].(map[string]interface{})
		networkPolicy = PolicyConfigModelFromMap(networkPolicyMap)
	}
	//
	title := m["title"].(string)
	typeVar := m["type"].(*models.TagType) // TagType
	return &models.Tag{
		AttestationPolicy: attestationPolicy,
		Deployment:        deployment,
		Description:       description,
		EdgeviewPolicy:    edgeviewPolicy,
		ID:                id,
		Name:              &name,
		NetworkPolicy:     networkPolicy,
		Title:             &title,
		Type:              typeVar,
	}
}

// Update the underlying Tag resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetTagResourceData(d *schema.ResourceData, m *models.Tag) {
	d.Set("app_policy", SetPolicyConfigSubResourceData([]*models.PolicyConfig{m.AppPolicy}))
	d.Set("attestation_policy", SetPolicyConfigSubResourceData([]*models.PolicyConfig{m.AttestationPolicy}))
	d.Set("attr", m.Attr)
	d.Set("cloud_policy", SetPolicyConfigSubResourceData([]*models.PolicyConfig{m.CloudPolicy}))
	d.Set("deployment", SetDeploymentSubResourceData([]*models.Deployment{m.Deployment}))
	d.Set("description", m.Description)
	d.Set("edgeview_policy", SetPolicyConfigSubResourceData([]*models.PolicyConfig{m.EdgeviewPolicy}))
	d.Set("id", m.ID)
	d.Set("module_policy", SetPolicyConfigSubResourceData(m.ModulePolicy))
	d.Set("name", m.Name)
	d.Set("network_policy", SetPolicyConfigSubResourceData([]*models.PolicyConfig{m.NetworkPolicy}))
	d.Set("numdevices", m.Numdevices)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("title", m.Title)
	d.Set("type", m.Type)
}

// Iterate through and update the Tag resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetTagSubResourceData(m []*models.Tag) (d []*map[string]interface{}) {
	for _, TagModel := range m {
		if TagModel != nil {
			properties := make(map[string]interface{})
			properties["app_policy"] = SetPolicyConfigSubResourceData([]*models.PolicyConfig{TagModel.AppPolicy})
			properties["attestation_policy"] = SetPolicyConfigSubResourceData([]*models.PolicyConfig{TagModel.AttestationPolicy})
			properties["attr"] = TagModel.Attr
			properties["cloud_policy"] = SetPolicyConfigSubResourceData([]*models.PolicyConfig{TagModel.CloudPolicy})
			properties["deployment"] = SetDeploymentSubResourceData([]*models.Deployment{TagModel.Deployment})
			properties["description"] = TagModel.Description
			properties["edgeview_policy"] = SetPolicyConfigSubResourceData([]*models.PolicyConfig{TagModel.EdgeviewPolicy})
			properties["id"] = TagModel.ID
			properties["module_policy"] = SetPolicyConfigSubResourceData(TagModel.ModulePolicy)
			properties["name"] = TagModel.Name
			properties["network_policy"] = SetPolicyConfigSubResourceData([]*models.PolicyConfig{TagModel.NetworkPolicy})
			properties["numdevices"] = TagModel.Numdevices
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{TagModel.Revision})
			properties["title"] = TagModel.Title
			properties["type"] = TagModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the Tag resource defined in the Terraform configuration
func TagSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_policy": {
			Description: `Resource group wide policy for edge applications to be deployed on all edge nodes on this resource group`,
			Type:        schema.TypeList, //GoType: PolicyConfig
			Elem: &schema.Resource{
				Schema: PolicyConfigSchema(),
			},
			Computed: true,
		},

		"attestation_policy": {
			Description: `Attestation policy to enforce on all devices of this project`,
			Type:        schema.TypeList, //GoType: PolicyConfig
			Elem: &schema.Resource{
				Schema: PolicyConfigSchema(),
			},
			Optional: true,
		},

		"attr": {
			Description: `Resource group wide configuration for edge nodes`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Computed: true,
		},

		"cloud_policy": {
			Description: `Resource group wide policy for Azure IoTEdge configuration to be applied to all edge applications`,
			Type:        schema.TypeList, //GoType: PolicyConfig
			Elem: &schema.Resource{
				Schema: PolicyConfigSchema(),
			},
			Computed: true,
		},

		"deployment": {
			Description: `Deployment template containing different types of policies`,
			Type:        schema.TypeList, //GoType: Deployment
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
			Type:        schema.TypeList, //GoType: PolicyConfig
			Elem: &schema.Resource{
				Schema: PolicyConfigSchema(),
			},
			Optional: true,
		},

		"id": {
			Description: `System defined universally unique Id of the resource group.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"module_policy": {
			Description: `Resource group wide policy for Azure module configuration to be applied to all edge applications`,
			Type:        schema.TypeList, //GoType: []*PolicyConfig
			Elem: &schema.Resource{
				Schema: PolicyConfigSchema(),
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
			Type:        schema.TypeList, //GoType: PolicyConfig
			Elem: &schema.Resource{
				Schema: PolicyConfigSchema(),
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
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevisionSchema(),
			},
			Computed: true,
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

// Retrieve property field names for updating the Tag resource
func GetTagPropertyFields() (t []string) {
	return []string{
		"attestation_policy",
		"deployment",
		"description",
		"edgeview_policy",
		"id",
		"name",
		"network_policy",
		"title",
		"type",
	}
}
