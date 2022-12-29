package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate Deployment resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeploymentModel(d *schema.ResourceData) *models.Deployment {
	appInstPolicies, _ := d.Get("app_inst_policies").([]*models.AppInstPolicy) // []*AppInstPolicy
	var clusterPolicy *models.ClusterInstPolicy                                // ClusterInstPolicy
	clusterPolicyInterface, clusterPolicyIsSet := d.GetOk("cluster_policy")
	if clusterPolicyIsSet {
		clusterPolicyMap := clusterPolicyInterface.([]interface{})[0].(map[string]interface{})
		clusterPolicy = ClusterInstPolicyModelFromMap(clusterPolicyMap)
	}
	deploymentTag, _ := d.Get("deployment_tag").(string)
	devicePolicies, _ := d.Get("device_policies").([]*models.DevicePolicy) // []*DevicePolicy
	id, _ := d.Get("id").(string)
	var integrationPolicy *models.IntegrationPolicy // IntegrationPolicy
	integrationPolicyInterface, integrationPolicyIsSet := d.GetOk("integration_policy")
	if integrationPolicyIsSet {
		integrationPolicyMap := integrationPolicyInterface.([]interface{})[0].(map[string]interface{})
		integrationPolicy = IntegrationPolicyModelFromMap(integrationPolicyMap)
	}
	name, _ := d.Get("name").(string)
	networkInstPolicies, _ := d.Get("network_inst_policies").([]*models.NetworkInstPolicy) // []*NetworkInstPolicy
	var revision *models.ObjectRevision                                                    // ObjectRevision
	revisionInterface, revisionIsSet := d.GetOk("revision")
	if revisionIsSet {
		revisionMap := revisionInterface.([]interface{})[0].(map[string]interface{})
		revision = ObjectRevisionModelFromMap(revisionMap)
	}
	title, _ := d.Get("title").(string)
	volumeInstPolicies, _ := d.Get("volume_inst_policies").([]*models.VolumeInstPolicy) // []*VolumeInstPolicy
	return &models.Deployment{
		AppInstPolicies:     appInstPolicies,
		ClusterPolicy:       clusterPolicy,
		DeploymentTag:       deploymentTag,
		DevicePolicies:      devicePolicies,
		ID:                  id,
		IntegrationPolicy:   integrationPolicy,
		Name:                name,
		NetworkInstPolicies: networkInstPolicies,
		Revision:            revision,
		Title:               title,
		VolumeInstPolicies:  volumeInstPolicies,
	}
}

func DeploymentModelFromMap(m map[string]interface{}) *models.Deployment {
	appInstPolicies := m["app_inst_policies"].([]*models.AppInstPolicy) // []*AppInstPolicy
	var clusterPolicy *models.ClusterInstPolicy                         // ClusterInstPolicy
	clusterPolicyInterface, clusterPolicyIsSet := m["cluster_policy"]
	if clusterPolicyIsSet {
		clusterPolicyMap := clusterPolicyInterface.([]interface{})[0].(map[string]interface{})
		clusterPolicy = ClusterInstPolicyModelFromMap(clusterPolicyMap)
	}
	//
	deploymentTag := m["deployment_tag"].(string)
	devicePolicies := m["device_policies"].([]*models.DevicePolicy) // []*DevicePolicy
	id := m["id"].(string)
	var integrationPolicy *models.IntegrationPolicy // IntegrationPolicy
	integrationPolicyInterface, integrationPolicyIsSet := m["integration_policy"]
	if integrationPolicyIsSet {
		integrationPolicyMap := integrationPolicyInterface.([]interface{})[0].(map[string]interface{})
		integrationPolicy = IntegrationPolicyModelFromMap(integrationPolicyMap)
	}
	//
	name := m["name"].(string)
	networkInstPolicies := m["network_inst_policies"].([]*models.NetworkInstPolicy) // []*NetworkInstPolicy
	var revision *models.ObjectRevision                                             // ObjectRevision
	revisionInterface, revisionIsSet := m["revision"]
	if revisionIsSet {
		revisionMap := revisionInterface.([]interface{})[0].(map[string]interface{})
		revision = ObjectRevisionModelFromMap(revisionMap)
	}
	//
	title := m["title"].(string)
	volumeInstPolicies := m["volume_inst_policies"].([]*models.VolumeInstPolicy) // []*VolumeInstPolicy
	return &models.Deployment{
		AppInstPolicies:     appInstPolicies,
		ClusterPolicy:       clusterPolicy,
		DeploymentTag:       deploymentTag,
		DevicePolicies:      devicePolicies,
		ID:                  id,
		IntegrationPolicy:   integrationPolicy,
		Name:                name,
		NetworkInstPolicies: networkInstPolicies,
		Revision:            revision,
		Title:               title,
		VolumeInstPolicies:  volumeInstPolicies,
	}
}

// Update the underlying Deployment resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeploymentResourceData(d *schema.ResourceData, m *models.Deployment) {
	d.Set("app_inst_policies", SetAppInstPolicySubResourceData(m.AppInstPolicies))
	d.Set("cluster_policy", SetClusterInstPolicySubResourceData([]*models.ClusterInstPolicy{m.ClusterPolicy}))
	d.Set("deployment_tag", m.DeploymentTag)
	d.Set("device_policies", SetDevicePolicySubResourceData(m.DevicePolicies))
	d.Set("id", m.ID)
	d.Set("integration_policy", SetIntegrationPolicySubResourceData([]*models.IntegrationPolicy{m.IntegrationPolicy}))
	d.Set("name", m.Name)
	d.Set("network_inst_policies", SetNetworkInstPolicySubResourceData(m.NetworkInstPolicies))
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("title", m.Title)
	d.Set("volume_inst_policies", SetVolumeInstPolicySubResourceData(m.VolumeInstPolicies))
}

// Iterate through and update the Deployment resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeploymentSubResourceData(m []*models.Deployment) (d []*map[string]interface{}) {
	for _, DeploymentModel := range m {
		if DeploymentModel != nil {
			properties := make(map[string]interface{})
			properties["app_inst_policies"] = SetAppInstPolicySubResourceData(DeploymentModel.AppInstPolicies)
			properties["cluster_policy"] = SetClusterInstPolicySubResourceData([]*models.ClusterInstPolicy{DeploymentModel.ClusterPolicy})
			properties["deployment_tag"] = DeploymentModel.DeploymentTag
			properties["device_policies"] = SetDevicePolicySubResourceData(DeploymentModel.DevicePolicies)
			properties["id"] = DeploymentModel.ID
			properties["integration_policy"] = SetIntegrationPolicySubResourceData([]*models.IntegrationPolicy{DeploymentModel.IntegrationPolicy})
			properties["name"] = DeploymentModel.Name
			properties["network_inst_policies"] = SetNetworkInstPolicySubResourceData(DeploymentModel.NetworkInstPolicies)
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{DeploymentModel.Revision})
			properties["title"] = DeploymentModel.Title
			properties["volume_inst_policies"] = SetVolumeInstPolicySubResourceData(DeploymentModel.VolumeInstPolicies)
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the Deployment resource defined in the Terraform configuration
func DeploymentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_inst_policies": {
			Description: `list of app instance policies`,
			Type:        schema.TypeList, //GoType: []*AppInstPolicy
			Elem: &schema.Resource{
				Schema: AppInstPolicySchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"cluster_policy": {
			Description: `cluster policy details`,
			Type:        schema.TypeList, //GoType: ClusterInstPolicy
			Elem: &schema.Resource{
				Schema: ClusterInstPolicySchema(),
			},
			Optional: true,
		},

		"deployment_tag": {
			Description: `user defined tag for the deployment, which is used while targeting set of devices`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"device_policies": {
			Description: `list of device policies`,
			Type:        schema.TypeList, //GoType: []*DevicePolicy
			Elem: &schema.Resource{
				Schema: DevicePolicySchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"id": {
			Description: `system generated unique id for an deployment`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"integration_policy": {
			Description: `integration policy details`,
			Type:        schema.TypeList, //GoType: IntegrationPolicy
			Elem: &schema.Resource{
				Schema: IntegrationPolicySchema(),
			},
			Optional: true,
		},

		"name": {
			Description: `user defined name for the deployment`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"network_inst_policies": {
			Description: `list of network instance policies`,
			Type:        schema.TypeList, //GoType: []*NetworkInstPolicy
			Elem: &schema.Resource{
				Schema: NetworkInstPolicySchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"revision": {
			Description: `object revision`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevisionSchema(),
			},
			Optional: true,
		},

		"title": {
			Description: `user defined title for the deployment`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"volume_inst_policies": {
			Description: `list of volume instamce policies`,
			Type:        schema.TypeList, //GoType: []*VolumeInstPolicy
			Elem: &schema.Resource{
				Schema: VolumeInstPolicySchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the Deployment resource
func GetDeploymentPropertyFields() (t []string) {
	return []string{
		"app_inst_policies",
		"cluster_policy",
		"deployment_tag",
		"device_policies",
		"id",
		"integration_policy",
		"name",
		"network_inst_policies",
		"revision",
		"title",
		"volume_inst_policies",
	}
}
