package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/go-openapi/strfmt"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZksInstanceMetricsModel(d *schema.ResourceData) *models.ZksInstanceMetrics {
	architecture, _ := d.Get("architecture").(string)
	var capacity *models.ClusterCapacity // ClusterCapacity
	capacityInterface, capacityIsSet := d.GetOk("capacity")
	if capacityIsSet && capacityInterface != nil {
		capacityMap := capacityInterface.([]interface{})
		if len(capacityMap) > 0 {
			capacity = ClusterCapacityModelFromMap(capacityMap[0].(map[string]interface{}))
		}
	}
	created, _ := d.Get("created").(strfmt.DateTime)
	deploymentsInt, _ := d.Get("deployments").(int)
	deployments := int32(deploymentsInt)
	id, _ := d.Get("id").(string)
	kubeVersion, _ := d.Get("kube_version").(string)
	nOfNodesInt, _ := d.Get("n_of_nodes").(int)
	nOfNodes := int32(nOfNodesInt)
	name, _ := d.Get("name").(string)
	provider, _ := d.Get("provider").(string)
	totalResourcesInt, _ := d.Get("total_resources").(int)
	totalResources := int32(totalResourcesInt)
	return &models.ZksInstanceMetrics{
		Architecture:   architecture,
		Capacity:       capacity,
		Created:        created,
		Deployments:    deployments,
		ID:             id,
		KubeVersion:    kubeVersion,
		NOfNodes:       nOfNodes,
		Name:           name,
		Provider:       provider,
		TotalResources: totalResources,
	}
}

func ZksInstanceMetricsModelFromMap(m map[string]interface{}) *models.ZksInstanceMetrics {
	architecture := m["architecture"].(string)
	var capacity *models.ClusterCapacity // ClusterCapacity
	capacityInterface, capacityIsSet := m["capacity"]
	if capacityIsSet && capacityInterface != nil {
		capacityMap := capacityInterface.([]interface{})
		if len(capacityMap) > 0 {
			capacity = ClusterCapacityModelFromMap(capacityMap[0].(map[string]interface{}))
		}
	}
	//
	created := m["created"].(strfmt.DateTime)
	deployments := int32(m["deployments"].(int)) // int32
	id := m["id"].(string)
	kubeVersion := m["kube_version"].(string)
	nOfNodes := int32(m["n_of_nodes"].(int)) // int32
	name := m["name"].(string)
	provider := m["provider"].(string)
	totalResources := int32(m["total_resources"].(int)) // int32
	return &models.ZksInstanceMetrics{
		Architecture:   architecture,
		Capacity:       capacity,
		Created:        created,
		Deployments:    deployments,
		ID:             id,
		KubeVersion:    kubeVersion,
		NOfNodes:       nOfNodes,
		Name:           name,
		Provider:       provider,
		TotalResources: totalResources,
	}
}

func SetZksInstanceMetricsResourceData(d *schema.ResourceData, m *models.ZksInstanceMetrics) {
	d.Set("architecture", m.Architecture)
	d.Set("capacity", SetClusterCapacitySubResourceData([]*models.ClusterCapacity{m.Capacity}))
	d.Set("created", m.Created)
	d.Set("deployments", m.Deployments)
	d.Set("id", m.ID)
	d.Set("kube_version", m.KubeVersion)
	d.Set("n_of_nodes", m.NOfNodes)
	d.Set("name", m.Name)
	d.Set("provider", m.Provider)
	d.Set("total_resources", m.TotalResources)
}

func SetZksInstanceMetricsSubResourceData(m []*models.ZksInstanceMetrics) (d []*map[string]interface{}) {
	for _, ZksInstanceMetricsModel := range m {
		if ZksInstanceMetricsModel != nil {
			properties := make(map[string]interface{})
			properties["architecture"] = ZksInstanceMetricsModel.Architecture
			properties["capacity"] = SetClusterCapacitySubResourceData([]*models.ClusterCapacity{ZksInstanceMetricsModel.Capacity})
			properties["created"] = ZksInstanceMetricsModel.Created.String()
			properties["deployments"] = ZksInstanceMetricsModel.Deployments
			properties["id"] = ZksInstanceMetricsModel.ID
			properties["kube_version"] = ZksInstanceMetricsModel.KubeVersion
			properties["n_of_nodes"] = ZksInstanceMetricsModel.NOfNodes
			properties["name"] = ZksInstanceMetricsModel.Name
			properties["provider"] = ZksInstanceMetricsModel.Provider
			properties["total_resources"] = ZksInstanceMetricsModel.TotalResources
			d = append(d, &properties)
		}
	}
	return
}

func ZksInstanceMetricsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"architecture": {
			Description: `Architecture of the zks instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"capacity": {
			Description: `Capacity information of the zks instance`,
			Type:        schema.TypeList, //GoType: ClusterCapacity
			Elem: &schema.Resource{
				Schema: ClusterCapacitySchema(),
			},
			Optional: true,
		},

		"created": {
			Description:  `Creation time of the zks instance`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"deployments": {
			Description: `Number of deployments in the zks instance`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"id": {
			Description: `ZKS instance ID for which metrics are provided`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"kube_version": {
			Description: `Kubernetes version of the zks instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"n_of_nodes": {
			Description: `Number of nodes in the zks instance`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"name": {
			Description: `ZKS instance name for which metrics are provided`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"provider": {
			Description: `Provider of the zks instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"total_resources": {
			Description: `Total resources in the zks instance`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetZksInstanceMetricsPropertyFields() (t []string) {
	return []string{
		"architecture",
		"capacity",
		"created",
		"deployments",
		"id",
		"kube_version",
		"n_of_nodes",
		"name",
		"provider",
		"total_resources",
	}
}
