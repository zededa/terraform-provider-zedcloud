package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func ClusterPolicyModel(d *schema.ResourceData) *models.ClusterPolicy {
	appPolicyID, _ := d.Get("app_policy_id").(string)
	var clusterConfig *models.ClusterConfig // ClusterConfig
	clusterConfigInterface, clusterConfigIsSet := d.GetOk("cluster_config")
	if clusterConfigIsSet && clusterConfigInterface != nil {
		clusterConfigMap := clusterConfigInterface.([]interface{})
		if len(clusterConfigMap) > 0 {
			clusterConfig = ClusterConfigModelFromMap(clusterConfigMap[0].(map[string]interface{}))
		}
	}
	networkPolicyID, _ := d.Get("network_policy_id").(string)
	var typeVar *models.ClusterType // ClusterType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewClusterType(models.ClusterType(typeModel))
	}
	return &models.ClusterPolicy{
		AppPolicyID:     &appPolicyID,
		ClusterConfig:   clusterConfig,
		NetworkPolicyID: &networkPolicyID,
		Type:            typeVar,
	}
}

func ClusterPolicyModelFromMap(m map[string]interface{}) *models.ClusterPolicy {
	appPolicyID := m["app_policy_id"].(string)
	var clusterConfig *models.ClusterConfig // ClusterConfig
	clusterConfigInterface, clusterConfigIsSet := m["cluster_config"]
	if clusterConfigIsSet && clusterConfigInterface != nil {
		clusterConfigMap := clusterConfigInterface.([]interface{})
		if len(clusterConfigMap) > 0 {
			clusterConfig = ClusterConfigModelFromMap(clusterConfigMap[0].(map[string]interface{}))
		}
	}
	//
	networkPolicyID := m["network_policy_id"].(string)
	var typeVar *models.ClusterType // ClusterType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewClusterType(models.ClusterType(typeModel))
	}
	return &models.ClusterPolicy{
		AppPolicyID:     &appPolicyID,
		ClusterConfig:   clusterConfig,
		NetworkPolicyID: &networkPolicyID,
		Type:            typeVar,
	}
}

func SetClusterPolicyResourceData(d *schema.ResourceData, m *models.ClusterPolicy) {
	d.Set("app_policy_id", m.AppPolicyID)
	d.Set("cluster_config", SetClusterConfigSubResourceData([]*models.ClusterConfig{m.ClusterConfig}))
	d.Set("network_policy_id", m.NetworkPolicyID)
	d.Set("type", m.Type)
}

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

func ClusterPolicy() map[string]*schema.Schema {
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
				Schema: ClusterConfig(),
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

func GetClusterPolicyPropertyFields() (t []string) {
	return []string{
		"app_policy_id",
		"cluster_config",
		"network_policy_id",
		"type",
	}
}
