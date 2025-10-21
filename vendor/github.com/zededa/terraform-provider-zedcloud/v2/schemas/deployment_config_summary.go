package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func DeploymentConfigSummaryModel(d *schema.ResourceData) *models.DeploymentConfigSummary {
	var appInstPolicies []*models.PolicyCommon // []*PolicyCommon
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
			m := PolicyCommonModelFromMap(v.(map[string]interface{}))
			appInstPolicies = append(appInstPolicies, m)
		}
	}
	var clusterPolicy *models.PolicyCommon // PolicyCommon
	clusterPolicyInterface, clusterPolicyIsSet := d.GetOk("cluster_policy")
	if clusterPolicyIsSet && clusterPolicyInterface != nil {
		clusterPolicyMap := clusterPolicyInterface.([]interface{})
		if len(clusterPolicyMap) > 0 {
			clusterPolicy = PolicyCommonModelFromMap(clusterPolicyMap[0].(map[string]interface{}))
		}
	}
	deploymentTag, _ := d.Get("deployment_tag").(string)
	var devicePolicies []*models.PolicyCommon // []*PolicyCommon
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
			m := PolicyCommonModelFromMap(v.(map[string]interface{}))
			devicePolicies = append(devicePolicies, m)
		}
	}
	var edgeViewPolicy *models.PolicyCommon // PolicyCommon
	edgeViewPolicyInterface, edgeViewPolicyIsSet := d.GetOk("edge_view_policy")
	if edgeViewPolicyIsSet && edgeViewPolicyInterface != nil {
		edgeViewPolicyMap := edgeViewPolicyInterface.([]interface{})
		if len(edgeViewPolicyMap) > 0 {
			edgeViewPolicy = PolicyCommonModelFromMap(edgeViewPolicyMap[0].(map[string]interface{}))
		}
	}
	id, _ := d.Get("id").(string)
	var integrationPolicy *models.PolicyCommon // PolicyCommon
	integrationPolicyInterface, integrationPolicyIsSet := d.GetOk("integration_policy")
	if integrationPolicyIsSet && integrationPolicyInterface != nil {
		integrationPolicyMap := integrationPolicyInterface.([]interface{})
		if len(integrationPolicyMap) > 0 {
			integrationPolicy = PolicyCommonModelFromMap(integrationPolicyMap[0].(map[string]interface{}))
		}
	}
	name, _ := d.Get("name").(string)
	var networkInstPolicies []*models.PolicyCommon // []*PolicyCommon
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
			m := PolicyCommonModelFromMap(v.(map[string]interface{}))
			networkInstPolicies = append(networkInstPolicies, m)
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
	title, _ := d.Get("title").(string)
	var volumeInstPolicies []*models.PolicyCommon // []*PolicyCommon
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
			m := PolicyCommonModelFromMap(v.(map[string]interface{}))
			volumeInstPolicies = append(volumeInstPolicies, m)
		}
	}
	return &models.DeploymentConfigSummary{
		AppInstPolicies:     appInstPolicies,
		ClusterPolicy:       clusterPolicy,
		DeploymentTag:       deploymentTag,
		DevicePolicies:      devicePolicies,
		EdgeViewPolicy:      edgeViewPolicy,
		ID:                  id,
		IntegrationPolicy:   integrationPolicy,
		Name:                name,
		NetworkInstPolicies: networkInstPolicies,
		Revision:            revision,
		Title:               title,
		VolumeInstPolicies:  volumeInstPolicies,
	}
}

func DeploymentConfigSummaryModelFromMap(m map[string]interface{}) *models.DeploymentConfigSummary {
	var appInstPolicies []*models.PolicyCommon // []*PolicyCommon
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
			m := PolicyCommonModelFromMap(v.(map[string]interface{}))
			appInstPolicies = append(appInstPolicies, m)
		}
	}
	var clusterPolicy *models.PolicyCommon // PolicyCommon
	clusterPolicyInterface, clusterPolicyIsSet := m["cluster_policy"]
	if clusterPolicyIsSet && clusterPolicyInterface != nil {
		clusterPolicyMap := clusterPolicyInterface.([]interface{})
		if len(clusterPolicyMap) > 0 {
			clusterPolicy = PolicyCommonModelFromMap(clusterPolicyMap[0].(map[string]interface{}))
		}
	}
	//
	deploymentTag := m["deployment_tag"].(string)
	var devicePolicies []*models.PolicyCommon // []*PolicyCommon
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
			m := PolicyCommonModelFromMap(v.(map[string]interface{}))
			devicePolicies = append(devicePolicies, m)
		}
	}
	var edgeViewPolicy *models.PolicyCommon // PolicyCommon
	edgeViewPolicyInterface, edgeViewPolicyIsSet := m["edge_view_policy"]
	if edgeViewPolicyIsSet && edgeViewPolicyInterface != nil {
		edgeViewPolicyMap := edgeViewPolicyInterface.([]interface{})
		if len(edgeViewPolicyMap) > 0 {
			edgeViewPolicy = PolicyCommonModelFromMap(edgeViewPolicyMap[0].(map[string]interface{}))
		}
	}
	//
	id := m["id"].(string)
	var integrationPolicy *models.PolicyCommon // PolicyCommon
	integrationPolicyInterface, integrationPolicyIsSet := m["integration_policy"]
	if integrationPolicyIsSet && integrationPolicyInterface != nil {
		integrationPolicyMap := integrationPolicyInterface.([]interface{})
		if len(integrationPolicyMap) > 0 {
			integrationPolicy = PolicyCommonModelFromMap(integrationPolicyMap[0].(map[string]interface{}))
		}
	}
	//
	name := m["name"].(string)
	var networkInstPolicies []*models.PolicyCommon // []*PolicyCommon
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
			m := PolicyCommonModelFromMap(v.(map[string]interface{}))
			networkInstPolicies = append(networkInstPolicies, m)
		}
	}
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
	var volumeInstPolicies []*models.PolicyCommon // []*PolicyCommon
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
			m := PolicyCommonModelFromMap(v.(map[string]interface{}))
			volumeInstPolicies = append(volumeInstPolicies, m)
		}
	}
	return &models.DeploymentConfigSummary{
		AppInstPolicies:     appInstPolicies,
		ClusterPolicy:       clusterPolicy,
		DeploymentTag:       deploymentTag,
		DevicePolicies:      devicePolicies,
		EdgeViewPolicy:      edgeViewPolicy,
		ID:                  id,
		IntegrationPolicy:   integrationPolicy,
		Name:                name,
		NetworkInstPolicies: networkInstPolicies,
		Revision:            revision,
		Title:               title,
		VolumeInstPolicies:  volumeInstPolicies,
	}
}

func SetDeploymentConfigSummaryResourceData(d *schema.ResourceData, m *models.DeploymentConfigSummary) {
	d.Set("app_inst_policies", SetPolicyCommonSubResourceData(m.AppInstPolicies))
	d.Set("cluster_policy", SetPolicyCommonSubResourceData([]*models.PolicyCommon{m.ClusterPolicy}))
	d.Set("deployment_tag", m.DeploymentTag)
	d.Set("device_policies", SetPolicyCommonSubResourceData(m.DevicePolicies))
	d.Set("edge_view_policy", SetPolicyCommonSubResourceData([]*models.PolicyCommon{m.EdgeViewPolicy}))
	d.Set("id", m.ID)
	d.Set("integration_policy", SetPolicyCommonSubResourceData([]*models.PolicyCommon{m.IntegrationPolicy}))
	d.Set("name", m.Name)
	d.Set("network_inst_policies", SetPolicyCommonSubResourceData(m.NetworkInstPolicies))
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("title", m.Title)
	d.Set("volume_inst_policies", SetPolicyCommonSubResourceData(m.VolumeInstPolicies))
}

func SetDeploymentConfigSummarySubResourceData(m []*models.DeploymentConfigSummary) (d []*map[string]interface{}) {
	for _, DeploymentConfigSummaryModel := range m {
		if DeploymentConfigSummaryModel != nil {
			properties := make(map[string]interface{})
			properties["app_inst_policies"] = SetPolicyCommonSubResourceData(DeploymentConfigSummaryModel.AppInstPolicies)
			properties["cluster_policy"] = SetPolicyCommonSubResourceData([]*models.PolicyCommon{DeploymentConfigSummaryModel.ClusterPolicy})
			properties["deployment_tag"] = DeploymentConfigSummaryModel.DeploymentTag
			properties["device_policies"] = SetPolicyCommonSubResourceData(DeploymentConfigSummaryModel.DevicePolicies)
			properties["edge_view_policy"] = SetPolicyCommonSubResourceData([]*models.PolicyCommon{DeploymentConfigSummaryModel.EdgeViewPolicy})
			properties["id"] = DeploymentConfigSummaryModel.ID
			properties["integration_policy"] = SetPolicyCommonSubResourceData([]*models.PolicyCommon{DeploymentConfigSummaryModel.IntegrationPolicy})
			properties["name"] = DeploymentConfigSummaryModel.Name
			properties["network_inst_policies"] = SetPolicyCommonSubResourceData(DeploymentConfigSummaryModel.NetworkInstPolicies)
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{DeploymentConfigSummaryModel.Revision})
			properties["title"] = DeploymentConfigSummaryModel.Title
			properties["volume_inst_policies"] = SetPolicyCommonSubResourceData(DeploymentConfigSummaryModel.VolumeInstPolicies)
			d = append(d, &properties)
		}
	}
	return
}

func DeploymentConfigSummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_inst_policies": {
			Description: `list of app instance policies`,
			Type:        schema.TypeList, //GoType: []*PolicyCommon
			Elem: &schema.Resource{
				Schema: PolicyCommon(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"cluster_policy": {
			Description: `cluster policy details`,
			Type:        schema.TypeList, //GoType: PolicyCommon
			Elem: &schema.Resource{
				Schema: PolicyCommon(),
			},
			Optional: true,
		},

		"deployment_tag": {
			Description: `user defined tag for an deployment`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"device_policies": {
			Description: `list of device policies`,
			Type:        schema.TypeList, //GoType: []*PolicyCommon
			Elem: &schema.Resource{
				Schema: PolicyCommon(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"edge_view_policy": {
			Description: `edge view policy details`,
			Type:        schema.TypeList, //GoType: PolicyCommon
			Elem: &schema.Resource{
				Schema: PolicyCommon(),
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
			Type:        schema.TypeList, //GoType: PolicyCommon
			Elem: &schema.Resource{
				Schema: PolicyCommon(),
			},
			Optional: true,
		},

		"name": {
			Description: `user defined deployment name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"network_inst_policies": {
			Description: `list of network instance policies`,
			Type:        schema.TypeList, //GoType: []*PolicyCommon
			Elem: &schema.Resource{
				Schema: PolicyCommon(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"revision": {
			Description: `object revision details`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Optional: true,
		},

		"title": {
			Description: `user defined deployment title`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"volume_inst_policies": {
			Description: `list of volume instance policies`,
			Type:        schema.TypeList, //GoType: []*PolicyCommon
			Elem: &schema.Resource{
				Schema: PolicyCommon(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

func GetDeploymentConfigSummaryPropertyFields() (t []string) {
	return []string{
		"app_inst_policies",
		"cluster_policy",
		"deployment_tag",
		"device_policies",
		"edge_view_policy",
		"id",
		"integration_policy",
		"name",
		"network_inst_policies",
		"revision",
		"title",
		"volume_inst_policies",
	}
}
