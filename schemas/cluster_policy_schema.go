package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ClusterPolicy resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ClusterPolicyModel(d *schema.ResourceData) *models.ClusterPolicy {
	appPolicyID, _ := d.Get("app_policy_id").(string)
	var clusterConfig *models.ClusterConfig // ClusterConfig
	clusterConfigInterface, clusterConfigIsSet := d.GetOk("cluster_config")
	if clusterConfigIsSet {
		clusterConfigMap := clusterConfigInterface.([]interface{})[0].(map[string]interface{})
		clusterConfig = ClusterConfigModelFromMap(clusterConfigMap)
	}
	networkPolicyID, _ := d.Get("network_policy_id").(string)
	typeVarModel, ok := d.Get("type").(models.ClusterType) // ClusterType
	typeVar := &typeVarModel
	if !ok {
		typeVar = nil
	}
	return &models.ClusterPolicy{
		AppPolicyID:     &appPolicyID, // string true false false
		ClusterConfig:   clusterConfig,
		NetworkPolicyID: &networkPolicyID, // string true false false
		Type:            typeVar,
	}
}

func ClusterPolicyModelFromMap(m map[string]interface{}) *models.ClusterPolicy {
	appPolicyID := m["app_policy_id"].(string)
	var clusterConfig *models.ClusterConfig // ClusterConfig
	clusterConfigInterface, clusterConfigIsSet := m["cluster_config"]
	if clusterConfigIsSet {
		clusterConfigMap := clusterConfigInterface.([]interface{})[0].(map[string]interface{})
		clusterConfig = ClusterConfigModelFromMap(clusterConfigMap)
	}
	//
	networkPolicyID := m["network_policy_id"].(string)
	typeVar := m["type"].(*models.ClusterType) // ClusterType
	return &models.ClusterPolicy{
		AppPolicyID:     &appPolicyID,
		ClusterConfig:   clusterConfig,
		NetworkPolicyID: &networkPolicyID,
		Type:            typeVar,
	}
}

// Update the underlying ClusterPolicy resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetClusterPolicyResourceData(d *schema.ResourceData, m *models.ClusterPolicy) {
	d.Set("app_policy_id", m.AppPolicyID)
	d.Set("cluster_config", SetClusterConfigSubResourceData([]*models.ClusterConfig{m.ClusterConfig}))
	d.Set("network_policy_id", m.NetworkPolicyID)
	d.Set("type", m.Type)
}

// Iterate through and update the ClusterPolicy resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetClusterPolicySubResourceData(m []*models.ClusterPolicy) (d []*map[string]interface{}) {
	for _, ClusterPolicyModel := range m {
		if ClusterPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["app_policy_id"] = ClusterPolicyModel.AppPolicyID
			properties["cluster_config"] = SetClusterConfigSubResourceData([]*models.ClusterConfig{ClusterPolicyModel.ClusterConfig})
			properties["network_policy_id"] = ClusterPolicyModel.NetworkPolicyID
			properties["type"] = ClusterPolicyModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ClusterPolicy resource defined in the Terraform configuration
func ClusterPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_policy_id": {
			Description: `UUID of the app policy linked to this cluster policy`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"cluster_config": {
			Description: `Cluster Policy Parameters`,
			Type:        schema.TypeList, //GoType: ClusterConfig
			Elem: &schema.Resource{
				Schema: ClusterConfigSchema(),
			},
			Optional: true,
		},

		"network_policy_id": {
			Description: `UUID of the network policy linked to this cluster policy`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"type": {
			Description: `Type of cluster`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

// Retrieve property field names for updating the ClusterPolicy resource
func GetClusterPolicyPropertyFields() (t []string) {
	return []string{
		"app_policy_id",
		"cluster_config",
		"network_policy_id",
		"type",
	}
}
