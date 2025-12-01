package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func TargetClustersModel(d *schema.ResourceData) *models.TargetClusters {
	var clusterMetadata *models.ClusterMetadata // ClusterMetadata
	clusterMetadataInterface, clusterMetadataIsSet := d.GetOk("cluster_metadata")
	if clusterMetadataIsSet && clusterMetadataInterface != nil {
		clusterMetadataMap := clusterMetadataInterface.([]interface{})
		if len(clusterMetadataMap) > 0 {
			clusterMetadata = ClusterMetadataModelFromMap(clusterMetadataMap[0].(map[string]interface{}))
		}
	}
	clusterTags := map[string]string{}
	clusterTagsInterface, clusterTagsIsSet := d.GetOk("clusterTags")
	if clusterTagsIsSet {
		clusterTagsMap := clusterTagsInterface.(map[string]interface{})
		for k, v := range clusterTagsMap {
			if v == nil {
				continue
			}
			clusterTags[k] = v.(string)
		}
	}

	isMultiCluster, _ := d.Get("is_multi_cluster").(bool)
	return &models.TargetClusters{
		ClusterMetadata: clusterMetadata,
		ClusterTags:     clusterTags,
		IsMultiCluster:  &isMultiCluster, // bool
	}
}

func TargetClustersModelFromMap(m map[string]interface{}) *models.TargetClusters {
	var clusterMetadata *models.ClusterMetadata // ClusterMetadata
	clusterMetadataInterface, clusterMetadataIsSet := m["cluster_metadata"]
	if clusterMetadataIsSet && clusterMetadataInterface != nil {
		clusterMetadataMap := clusterMetadataInterface.([]interface{})
		if len(clusterMetadataMap) > 0 {
			clusterMetadata = ClusterMetadataModelFromMap(clusterMetadataMap[0].(map[string]interface{}))
		}
	}
	//
	clusterTags := map[string]string{}
	clusterTagsInterface, clusterTagsIsSet := m["cluster_tags"]
	if clusterTagsIsSet {
		clusterTagsMap := clusterTagsInterface.(map[string]interface{})
		for k, v := range clusterTagsMap {
			if v == nil {
				continue
			}
			clusterTags[k] = v.(string)
		}
	}

	isMultiCluster := m["is_multi_cluster"].(bool)
	return &models.TargetClusters{
		ClusterMetadata: clusterMetadata,
		ClusterTags:     clusterTags,
		IsMultiCluster:  &isMultiCluster,
	}
}

func SetTargetClustersResourceData(d *schema.ResourceData, m *models.TargetClusters) {
	d.Set("cluster_metadata", SetClusterMetadataSubResourceData([]*models.ClusterMetadata{m.ClusterMetadata}))
	d.Set("cluster_tags", m.ClusterTags)
	d.Set("is_multi_cluster", m.IsMultiCluster)
}

func SetTargetClustersSubResourceData(m []*models.TargetClusters) (d []*map[string]interface{}) {
	for _, TargetClustersModel := range m {
		if TargetClustersModel != nil {
			properties := make(map[string]interface{})
			properties["cluster_metadata"] = SetClusterMetadataSubResourceData([]*models.ClusterMetadata{TargetClustersModel.ClusterMetadata})
			properties["cluster_tags"] = TargetClustersModel.ClusterTags
			properties["is_multi_cluster"] = TargetClustersModel.IsMultiCluster
			d = append(d, &properties)
		}
	}
	return
}

func TargetClustersSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_metadata": {
			Description: `Metadata for the target cluster when is_multi_cluster is false`,
			Type:        schema.TypeList, //GoType: ClusterMetadata
			Elem: &schema.Resource{
				Schema: ClusterMetadataSchema(),
			},
			Optional: true,
		},

		"cluster_tags": {
			Description: `Tags for the clusters where the deployment will be applied. Example: {'environment': 'production', 'region': 'us-west-2'}`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"is_multi_cluster": {
			Description: `Indicates if the deployment is for multiple clusters. Example: false`,
			Type:        schema.TypeBool,
			Required:    true,
		},
	}
}

func GetTargetClustersPropertyFields() (t []string) {
	return []string{
		"cluster_metadata",
		"cluster_tags",
		"is_multi_cluster",
	}
}
