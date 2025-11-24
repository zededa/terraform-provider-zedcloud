package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ClusterMetadataModel(d *schema.ResourceData) *models.ClusterMetadata {
	clusterID, _ := d.Get("cluster_id").(string)
	clusterName, _ := d.Get("cluster_name").(string)
	kubeVersion, _ := d.Get("kube_version").(string)
	orchestratorClusterID, _ := d.Get("orchestrator_cluster_id").(string)
	return &models.ClusterMetadata{
		ClusterID:             &clusterID, // string
		ClusterName:           clusterName,
		KubeVersion:           kubeVersion,
		OrchestratorClusterID: &orchestratorClusterID, // string
	}
}

func ClusterMetadataModelFromMap(m map[string]interface{}) *models.ClusterMetadata {
	clusterID := m["cluster_id"].(string)
	clusterName := m["cluster_name"].(string)
	kubeVersion := m["kube_version"].(string)
	orchestratorClusterID := m["orchestrator_cluster_id"].(string)
	return &models.ClusterMetadata{
		ClusterID:             &clusterID,
		ClusterName:           clusterName,
		KubeVersion:           kubeVersion,
		OrchestratorClusterID: &orchestratorClusterID,
	}
}

func SetClusterMetadataResourceData(d *schema.ResourceData, m *models.ClusterMetadata) {
	d.Set("cluster_id", m.ClusterID)
	d.Set("cluster_name", m.ClusterName)
	d.Set("kube_version", m.KubeVersion)
	d.Set("orchestrator_cluster_id", m.OrchestratorClusterID)
}

func SetClusterMetadataSubResourceData(m []*models.ClusterMetadata) (d []*map[string]interface{}) {
	for _, ClusterMetadataModel := range m {
		if ClusterMetadataModel != nil {
			properties := make(map[string]interface{})
			properties["cluster_id"] = ClusterMetadataModel.ClusterID
			properties["cluster_name"] = ClusterMetadataModel.ClusterName
			properties["kube_version"] = ClusterMetadataModel.KubeVersion
			properties["orchestrator_cluster_id"] = ClusterMetadataModel.OrchestratorClusterID
			d = append(d, &properties)
		}
	}
	return
}

func ClusterMetadataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_id": {
			Description: `ID of the target cluster. Example: 'cluster-abc123'`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"cluster_name": {
			Description: `Name of the target cluster. Example: 'production-k8s'`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"kube_version": {
			Description: `Kubernetes version of the target cluster. Example: 'v1.28.0'`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"orchestrator_cluster_id": {
			Description: `Orchestrator cluster ID. Example: 'rancher-cluster-123'`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func GetClusterMetadataPropertyFields() (t []string) {
	return []string{
		"cluster_id",
		"cluster_name",
		"kube_version",
		"orchestrator_cluster_id",
	}
}
