package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ClusterNodeModel(d *schema.ResourceData) *models.ClusterNode {
	clusterInterface, _ := d.Get("cluster_interface").(string)
	clusterPrefix, _ := d.Get("cluster_prefix").(string)
	id, _ := d.Get("id").(string)
	var nodeType *models.ClusterNodeType // ClusterNodeType
	nodeTypeInterface, nodeTypeIsSet := d.GetOk("node_type")
	if nodeTypeIsSet {
		nodeTypeModel := nodeTypeInterface.(string)
		nodeType = models.NewClusterNodeType(models.ClusterNodeType(nodeTypeModel))
	}
	resourceLabels := map[string]string{}
	resourceLabelsInterface, resourceLabelsIsSet := d.GetOk("resourceLabels")
	if resourceLabelsIsSet {
		resourceLabelsMap := resourceLabelsInterface.(map[string]interface{})
		for k, v := range resourceLabelsMap {
			if v == nil {
				continue
			}
			resourceLabels[k] = v.(string)
		}
	}

	return &models.ClusterNode{
		ClusterInterface: &clusterInterface, // string
		ClusterPrefix:    clusterPrefix,
		ID:               id,
		NodeType:         nodeType,
		ResourceLabels:   resourceLabels,
	}
}

func ClusterNodeModelFromMap(m map[string]interface{}) *models.ClusterNode {
	clusterInterface := m["cluster_interface"].(string)
	clusterPrefix := m["cluster_prefix"].(string)
	id := m["id"].(string)
	var nodeType *models.ClusterNodeType // ClusterNodeType
	nodeTypeInterface, nodeTypeIsSet := m["node_type"]
	if nodeTypeIsSet {
		nodeTypeModel := nodeTypeInterface.(string)
		nodeType = models.NewClusterNodeType(models.ClusterNodeType(nodeTypeModel))
	}
	resourceLabels := map[string]string{}
	resourceLabelsInterface, resourceLabelsIsSet := m["resource_labels"]
	if resourceLabelsIsSet {
		resourceLabelsMap := resourceLabelsInterface.(map[string]interface{})
		for k, v := range resourceLabelsMap {
			if v == nil {
				continue
			}
			resourceLabels[k] = v.(string)
		}
	}

	return &models.ClusterNode{
		ClusterInterface: &clusterInterface,
		ClusterPrefix:    clusterPrefix,
		ID:               id,
		NodeType:         nodeType,
		ResourceLabels:   resourceLabels,
	}
}

func SetClusterNodeResourceData(d *schema.ResourceData, m *models.ClusterNode) {
	d.Set("cluster_interface", m.ClusterInterface)
	d.Set("cluster_prefix", m.ClusterPrefix)
	d.Set("id", m.ID)
	d.Set("node_type", m.NodeType)
	d.Set("resource_labels", m.ResourceLabels)
}

func SetClusterNodeSubResourceData(m []*models.ClusterNode) (d []*map[string]interface{}) {
	for _, ClusterNodeModel := range m {
		if ClusterNodeModel != nil {
			properties := make(map[string]interface{})
			properties["cluster_interface"] = ClusterNodeModel.ClusterInterface
			properties["cluster_prefix"] = ClusterNodeModel.ClusterPrefix
			properties["id"] = ClusterNodeModel.ID
			properties["node_type"] = ClusterNodeModel.NodeType
			properties["resource_labels"] = ClusterNodeModel.ResourceLabels
			d = append(d, &properties)
		}
	}
	return
}

func ClusterNodeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_interface": {
			Description: `User defined name of the cluster interface, ex.: eth0, eth1, wlan0, etc.`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"cluster_prefix": {
			Description: `A cluster prefix. The system will assign a value.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"id": {
			Description: `Id of the node to be included in a cluster`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"node_type": {
			Description: `Type of node. Currently it supports only: EDGE_NODE_CLUSTER_NODE_TYPE_SERVER`,
			Type:        schema.TypeString,
			Optional:    true,
			Default:	"EDGE_NODE_CLUSTER_NODE_TYPE_SERVER",
		},

		"resource_labels": {
			Description: `Resource labels are name/value pairs that enable you to categorize resources. Label names are case insensitive with max_length 512 and min_length 3. Label values are case sensitive with max_length 256 and min_length 3.`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

func GetClusterNodePropertyFields() (t []string) {
	return []string{
		"cluster_interface",
		"cluster_prefix",
		"id",
		"node_type",
		"resource_labels",
	}
}
