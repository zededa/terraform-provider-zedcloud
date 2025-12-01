package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ClusterCapacityModel(d *schema.ResourceData) *models.ClusterCapacity {
	var cpus *models.ResourceCapacityUsage // ResourceCapacityUsage
	cpusInterface, cpusIsSet := d.GetOk("cpus")
	if cpusIsSet && cpusInterface != nil {
		cpusMap := cpusInterface.([]interface{})
		if len(cpusMap) > 0 {
			cpus = ResourceCapacityUsageModelFromMap(cpusMap[0].(map[string]interface{}))
		}
	}
	var memory *models.ResourceCapacityUsage // ResourceCapacityUsage
	memoryInterface, memoryIsSet := d.GetOk("memory")
	if memoryIsSet && memoryInterface != nil {
		memoryMap := memoryInterface.([]interface{})
		if len(memoryMap) > 0 {
			memory = ResourceCapacityUsageModelFromMap(memoryMap[0].(map[string]interface{}))
		}
	}
	var pods *models.PodsCapacity // PodsCapacity
	podsInterface, podsIsSet := d.GetOk("pods")
	if podsIsSet && podsInterface != nil {
		podsMap := podsInterface.([]interface{})
		if len(podsMap) > 0 {
			pods = PodsCapacityModelFromMap(podsMap[0].(map[string]interface{}))
		}
	}
	return &models.ClusterCapacity{
		Cpus:   cpus,
		Memory: memory,
		Pods:   pods,
	}
}

func ClusterCapacityModelFromMap(m map[string]interface{}) *models.ClusterCapacity {
	var cpus *models.ResourceCapacityUsage // ResourceCapacityUsage
	cpusInterface, cpusIsSet := m["cpus"]
	if cpusIsSet && cpusInterface != nil {
		cpusMap := cpusInterface.([]interface{})
		if len(cpusMap) > 0 {
			cpus = ResourceCapacityUsageModelFromMap(cpusMap[0].(map[string]interface{}))
		}
	}
	//
	var memory *models.ResourceCapacityUsage // ResourceCapacityUsage
	memoryInterface, memoryIsSet := m["memory"]
	if memoryIsSet && memoryInterface != nil {
		memoryMap := memoryInterface.([]interface{})
		if len(memoryMap) > 0 {
			memory = ResourceCapacityUsageModelFromMap(memoryMap[0].(map[string]interface{}))
		}
	}
	//
	var pods *models.PodsCapacity // PodsCapacity
	podsInterface, podsIsSet := m["pods"]
	if podsIsSet && podsInterface != nil {
		podsMap := podsInterface.([]interface{})
		if len(podsMap) > 0 {
			pods = PodsCapacityModelFromMap(podsMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.ClusterCapacity{
		Cpus:   cpus,
		Memory: memory,
		Pods:   pods,
	}
}

func SetClusterCapacityResourceData(d *schema.ResourceData, m *models.ClusterCapacity) {
	d.Set("cpus", SetResourceCapacityUsageSubResourceData([]*models.ResourceCapacityUsage{m.Cpus}))
	d.Set("memory", SetResourceCapacityUsageSubResourceData([]*models.ResourceCapacityUsage{m.Memory}))
	d.Set("pods", SetPodsCapacitySubResourceData([]*models.PodsCapacity{m.Pods}))
}

func SetClusterCapacitySubResourceData(m []*models.ClusterCapacity) (d []*map[string]interface{}) {
	for _, ClusterCapacityModel := range m {
		if ClusterCapacityModel != nil {
			properties := make(map[string]interface{})
			properties["cpus"] = SetResourceCapacityUsageSubResourceData([]*models.ResourceCapacityUsage{ClusterCapacityModel.Cpus})
			properties["memory"] = SetResourceCapacityUsageSubResourceData([]*models.ResourceCapacityUsage{ClusterCapacityModel.Memory})
			properties["pods"] = SetPodsCapacitySubResourceData([]*models.PodsCapacity{ClusterCapacityModel.Pods})
			d = append(d, &properties)
		}
	}
	return
}

func ClusterCapacitySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cpus": {
			Description: `CPU capacity of the cluster`,
			Type:        schema.TypeList, //GoType: ResourceCapacityUsage
			Elem: &schema.Resource{
				Schema: ResourceCapacityUsageSchema(),
			},
			Optional: true,
		},

		"memory": {
			Description: `Memory capacity of the cluster`,
			Type:        schema.TypeList, //GoType: ResourceCapacityUsage
			Elem: &schema.Resource{
				Schema: ResourceCapacityUsageSchema(),
			},
			Optional: true,
		},

		"pods": {
			Description: `Pods capacity of the cluster`,
			Type:        schema.TypeList, //GoType: PodsCapacity
			Elem: &schema.Resource{
				Schema: PodsCapacitySchema(),
			},
			Optional: true,
		},
	}
}

func GetClusterCapacityPropertyFields() (t []string) {
	return []string{
		"cpus",
		"memory",
		"pods",
	}
}
