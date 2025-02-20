package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func DeploymentModel(d *schema.ResourceData) *models.Deployment {
	var appInstPolicies []*models.AppInstPolicy // []*AppInstPolicy
	appInstPoliciesInterface, appInstPoliciesIsSet := d.GetOk("app_inst_policies")
	if appInstPoliciesIsSet {
		var items []interface{}
		if listItems, isList := appInstPoliciesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = appInstPoliciesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AppInstPolicyModelFromMap(v.(map[string]interface{}))
			appInstPolicies = append(appInstPolicies, m)
		}
	}
	var clusterPolicy *models.ClusterInstPolicy // ClusterInstPolicy
	clusterPolicyInterface, clusterPolicyIsSet := d.GetOk("cluster_policy")
	if clusterPolicyIsSet && clusterPolicyInterface != nil {
		clusterPolicyMap := clusterPolicyInterface.([]interface{})
		if len(clusterPolicyMap) > 0 {
			clusterPolicy = ClusterInstPolicyModelFromMap(clusterPolicyMap[0].(map[string]interface{}))
		}
	}
	deploymentTag, _ := d.Get("deployment_tag").(string)
	var devicePolicies []*models.DevicePolicy // []*DevicePolicy
	devicePoliciesInterface, devicePoliciesIsSet := d.GetOk("device_policies")
	if devicePoliciesIsSet {
		var items []interface{}
		if listItems, isList := devicePoliciesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = devicePoliciesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := DevicePolicyModelFromMap(v.(map[string]interface{}))
			devicePolicies = append(devicePolicies, m)
		}
	}
	var edgeviewPolicy *models.EdgeViewPolicy // EdgeViewPolicy
	edgeviewPolicyInterface, edgeviewPolicyIsSet := d.GetOk("edgeview_policy")
	if edgeviewPolicyIsSet && edgeviewPolicyInterface != nil {
		edgeviewPolicyMap := edgeviewPolicyInterface.([]interface{})
		if len(edgeviewPolicyMap) > 0 {
			edgeviewPolicy = EdgeViewPolicyModelFromMap(edgeviewPolicyMap[0].(map[string]interface{}))
		}
	}
	id, _ := d.Get("id").(string)
	var integrationPolicy *models.IntegrationPolicy // IntegrationPolicy
	integrationPolicyInterface, integrationPolicyIsSet := d.GetOk("integration_policy")
	if integrationPolicyIsSet && integrationPolicyInterface != nil {
		integrationPolicyMap := integrationPolicyInterface.([]interface{})
		if len(integrationPolicyMap) > 0 {
			integrationPolicy = IntegrationPolicyModelFromMap(integrationPolicyMap[0].(map[string]interface{}))
		}
	}
	name, _ := d.Get("name").(string)
	var networkInstPolicies []*models.NetworkInstPolicy // []*NetworkInstPolicy
	networkInstPoliciesInterface, networkInstPoliciesIsSet := d.GetOk("network_inst_policies")
	if networkInstPoliciesIsSet {
		var items []interface{}
		if listItems, isList := networkInstPoliciesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = networkInstPoliciesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := NetworkInstPolicyModelFromMap(v.(map[string]interface{}))
			networkInstPolicies = append(networkInstPolicies, m)
		}
	}
	projectID, _ := d.Get("project_id").(string)
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := d.GetOk("revision")
	if revisionIsSet && revisionInterface != nil {
		revisionMap := revisionInterface.([]interface{})
		if len(revisionMap) > 0 {
			revision = ObjectRevisionModelFromMap(revisionMap[0].(map[string]interface{}))
		}
	}
	title, _ := d.Get("title").(string)
	var volumeInstPolicies []*models.VolumeInstPolicy // []*VolumeInstPolicy
	volumeInstPoliciesInterface, volumeInstPoliciesIsSet := d.GetOk("volume_inst_policies")
	if volumeInstPoliciesIsSet {
		var items []interface{}
		if listItems, isList := volumeInstPoliciesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = volumeInstPoliciesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := VolumeInstPolicyModelFromMap(v.(map[string]interface{}))
			volumeInstPolicies = append(volumeInstPolicies, m)
		}
	}
	return &models.Deployment{
		AppInstPolicies:     appInstPolicies,
		ClusterPolicy:       clusterPolicy,
		DeploymentTag:       deploymentTag,
		DevicePolicies:      devicePolicies,
		EdgeviewPolicy:      edgeviewPolicy,
		ID:                  id,
		IntegrationPolicy:   integrationPolicy,
		Name:                name,
		NetworkInstPolicies: networkInstPolicies,
		ProjectID:           projectID,
		Revision:            revision,
		Title:               title,
		VolumeInstPolicies:  volumeInstPolicies,
	}
}

func DeploymentModelFromMap(m map[string]interface{}) *models.Deployment {
	var appInstPolicies []*models.AppInstPolicy // []*AppInstPolicy
	appInstPoliciesInterface, appInstPoliciesIsSet := m["app_inst_policies"]
	if appInstPoliciesIsSet {
		var items []interface{}
		if listItems, isList := appInstPoliciesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = appInstPoliciesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AppInstPolicyModelFromMap(v.(map[string]interface{}))
			appInstPolicies = append(appInstPolicies, m)
		}
	}
	var clusterPolicy *models.ClusterInstPolicy // ClusterInstPolicy
	clusterPolicyInterface, clusterPolicyIsSet := m["cluster_policy"]
	if clusterPolicyIsSet && clusterPolicyInterface != nil {
		clusterPolicyMap := clusterPolicyInterface.([]interface{})
		if len(clusterPolicyMap) > 0 {
			clusterPolicy = ClusterInstPolicyModelFromMap(clusterPolicyMap[0].(map[string]interface{}))
		}
	}
	//
	deploymentTag := m["deployment_tag"].(string)
	var devicePolicies []*models.DevicePolicy // []*DevicePolicy
	devicePoliciesInterface, devicePoliciesIsSet := m["device_policies"]
	if devicePoliciesIsSet {
		var items []interface{}
		if listItems, isList := devicePoliciesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = devicePoliciesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := DevicePolicyModelFromMap(v.(map[string]interface{}))
			devicePolicies = append(devicePolicies, m)
		}
	}
	var edgeviewPolicy *models.EdgeViewPolicy // EdgeViewPolicy
	edgeviewPolicyInterface, edgeviewPolicyIsSet := m["edgeview_policy"]
	if edgeviewPolicyIsSet && edgeviewPolicyInterface != nil {
		edgeviewPolicyMap := edgeviewPolicyInterface.([]interface{})
		if len(edgeviewPolicyMap) > 0 {
			edgeviewPolicy = EdgeViewPolicyModelFromMap(edgeviewPolicyMap[0].(map[string]interface{}))
		}
	}
	//
	id := m["id"].(string)
	var integrationPolicy *models.IntegrationPolicy // IntegrationPolicy
	integrationPolicyInterface, integrationPolicyIsSet := m["integration_policy"]
	if integrationPolicyIsSet && integrationPolicyInterface != nil {
		integrationPolicyMap := integrationPolicyInterface.([]interface{})
		if len(integrationPolicyMap) > 0 {
			integrationPolicy = IntegrationPolicyModelFromMap(integrationPolicyMap[0].(map[string]interface{}))
		}
	}
	//
	name := m["name"].(string)
	var networkInstPolicies []*models.NetworkInstPolicy // []*NetworkInstPolicy
	networkInstPoliciesInterface, networkInstPoliciesIsSet := m["network_inst_policies"]
	if networkInstPoliciesIsSet {
		var items []interface{}
		if listItems, isList := networkInstPoliciesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = networkInstPoliciesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := NetworkInstPolicyModelFromMap(v.(map[string]interface{}))
			networkInstPolicies = append(networkInstPolicies, m)
		}
	}
	projectID := m["project_id"].(string)
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := m["revision"]
	if revisionIsSet && revisionInterface != nil {
		revisionMap := revisionInterface.([]interface{})
		if len(revisionMap) > 0 {
			revision = ObjectRevisionModelFromMap(revisionMap[0].(map[string]interface{}))
		}
	}
	//
	title := m["title"].(string)
	var volumeInstPolicies []*models.VolumeInstPolicy // []*VolumeInstPolicy
	volumeInstPoliciesInterface, volumeInstPoliciesIsSet := m["volume_inst_policies"]
	if volumeInstPoliciesIsSet {
		var items []interface{}
		if listItems, isList := volumeInstPoliciesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = volumeInstPoliciesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := VolumeInstPolicyModelFromMap(v.(map[string]interface{}))
			volumeInstPolicies = append(volumeInstPolicies, m)
		}
	}
	return &models.Deployment{
		AppInstPolicies:     appInstPolicies,
		ClusterPolicy:       clusterPolicy,
		DeploymentTag:       deploymentTag,
		DevicePolicies:      devicePolicies,
		EdgeviewPolicy:      edgeviewPolicy,
		ID:                  id,
		IntegrationPolicy:   integrationPolicy,
		Name:                name,
		NetworkInstPolicies: networkInstPolicies,
		ProjectID:           projectID,
		Revision:            revision,
		Title:               title,
		VolumeInstPolicies:  volumeInstPolicies,
	}
}

func SetDeploymentResourceData(d *schema.ResourceData, m *models.Deployment) {
	d.Set("app_inst_policies", SetAppInstPolicySubResourceData(m.AppInstPolicies))
	d.Set("cluster_policy", SetClusterInstPolicySubResourceData([]*models.ClusterInstPolicy{m.ClusterPolicy}))
	d.Set("deployment_tag", m.DeploymentTag)
	d.Set("device_policies", SetDevicePolicySubResourceData(m.DevicePolicies))
	d.Set("edgeview_policy", SetEdgeViewPolicySubResourceData([]*models.EdgeViewPolicy{m.EdgeviewPolicy}))
	d.Set("id", m.ID)
	d.Set("integration_policy", SetIntegrationPolicySubResourceData([]*models.IntegrationPolicy{m.IntegrationPolicy}))
	d.Set("name", m.Name)
	d.Set("network_inst_policies", SetNetworkInstPolicySubResourceData(m.NetworkInstPolicies))
	d.Set("project_id", m.ProjectID)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("title", m.Title)
	d.Set("volume_inst_policies", SetVolumeInstPolicySubResourceData(m.VolumeInstPolicies))
}

func SetDeploymentSubResourceData(m []*models.Deployment) (d []*map[string]interface{}) {
	for _, DeploymentModel := range m {
		if DeploymentModel != nil {
			properties := make(map[string]interface{})
			properties["app_inst_policies"] = SetAppInstPolicySubResourceData(DeploymentModel.AppInstPolicies)
			properties["cluster_policy"] = SetClusterInstPolicySubResourceData([]*models.ClusterInstPolicy{DeploymentModel.ClusterPolicy})
			properties["deployment_tag"] = DeploymentModel.DeploymentTag
			properties["device_policies"] = SetDevicePolicySubResourceData(DeploymentModel.DevicePolicies)
			properties["edgeview_policy"] = SetEdgeViewPolicySubResourceData([]*models.EdgeViewPolicy{DeploymentModel.EdgeviewPolicy})
			properties["id"] = DeploymentModel.ID
			properties["integration_policy"] = SetIntegrationPolicySubResourceData([]*models.IntegrationPolicy{DeploymentModel.IntegrationPolicy})
			properties["name"] = DeploymentModel.Name
			properties["network_inst_policies"] = SetNetworkInstPolicySubResourceData(DeploymentModel.NetworkInstPolicies)
			properties["project_id"] = DeploymentModel.ProjectID
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{DeploymentModel.Revision})
			properties["title"] = DeploymentModel.Title
			properties["volume_inst_policies"] = SetVolumeInstPolicySubResourceData(DeploymentModel.VolumeInstPolicies)
			d = append(d, &properties)
		}
	}
	return
}

func Deployment() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_inst_policies": {
			Description: `list of app instance policies`,
			Type:        schema.TypeList, //GoType: []*AppInstPolicy
			Elem: &schema.Resource{
				Schema: AppInstPolicy(),
			},
			Optional: true,
		},

		"cluster_policy": {
			Description: `cluster policy details`,
			Type:        schema.TypeList, //GoType: ClusterInstPolicy
			Elem: &schema.Resource{
				Schema: ClusterInstPolicy(),
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
				Schema: DevicePolicy(),
			},
			Optional: true,
		},

		"edgeview_policy": {
			Description: `edge view policy details`,
			Type:        schema.TypeList, //GoType: EdgeViewPolicy
			Elem: &schema.Resource{
				Schema: EdgeViewPolicy(),
			},
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
				Schema: IntegrationPolicy(),
			},
			Optional: true,
		},

		"name": {
			Description: `user defined name for the deployment`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"network_inst_policies": {
			Description: `list of network instance policies`,
			Type:        schema.TypeList, //GoType: []*NetworkInstPolicy
			Elem: &schema.Resource{
				Schema: NetworkInstPolicySchema(),
			},
			Optional: true,
		},

		"project_id": {
			Description: `project id`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"revision": {
			Description: `object revision`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Computed: true,
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
			Optional: true,
		},
	}
}

func GetDeploymentPropertyFields() (t []string) {
	return []string{
		"app_inst_policies",
		"cluster_policy",
		"deployment_tag",
		"device_policies",
		"edgeview_policy",
		"id",
		"integration_policy",
		"name",
		"network_inst_policies",
		"project_id",
		"revision",
		"title",
		"volume_inst_policies",
	}
}
