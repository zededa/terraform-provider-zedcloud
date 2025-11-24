package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ResourceCapacityUsageModel(d *schema.ResourceData) *models.ResourceCapacityUsage {
	var reserved *models.ResourceUsage // ResourceUsage
	reservedInterface, reservedIsSet := d.GetOk("reserved")
	if reservedIsSet && reservedInterface != nil {
		reservedMap := reservedInterface.([]interface{})
		if len(reservedMap) > 0 {
			reserved = ResourceUsageModelFromMap(reservedMap[0].(map[string]interface{}))
		}
	}
	var used *models.ResourceUsage // ResourceUsage
	usedInterface, usedIsSet := d.GetOk("used")
	if usedIsSet && usedInterface != nil {
		usedMap := usedInterface.([]interface{})
		if len(usedMap) > 0 {
			used = ResourceUsageModelFromMap(usedMap[0].(map[string]interface{}))
		}
	}
	return &models.ResourceCapacityUsage{
		Reserved: reserved,
		Used:     used,
	}
}

func ResourceCapacityUsageModelFromMap(m map[string]interface{}) *models.ResourceCapacityUsage {
	var reserved *models.ResourceUsage // ResourceUsage
	reservedInterface, reservedIsSet := m["reserved"]
	if reservedIsSet && reservedInterface != nil {
		reservedMap := reservedInterface.([]interface{})
		if len(reservedMap) > 0 {
			reserved = ResourceUsageModelFromMap(reservedMap[0].(map[string]interface{}))
		}
	}
	//
	var used *models.ResourceUsage // ResourceUsage
	usedInterface, usedIsSet := m["used"]
	if usedIsSet && usedInterface != nil {
		usedMap := usedInterface.([]interface{})
		if len(usedMap) > 0 {
			used = ResourceUsageModelFromMap(usedMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.ResourceCapacityUsage{
		Reserved: reserved,
		Used:     used,
	}
}

func SetResourceCapacityUsageResourceData(d *schema.ResourceData, m *models.ResourceCapacityUsage) {
	d.Set("reserved", SetResourceUsageSubResourceData([]*models.ResourceUsage{m.Reserved}))
	d.Set("used", SetResourceUsageSubResourceData([]*models.ResourceUsage{m.Used}))
}

func SetResourceCapacityUsageSubResourceData(m []*models.ResourceCapacityUsage) (d []*map[string]interface{}) {
	for _, ResourceCapacityUsageModel := range m {
		if ResourceCapacityUsageModel != nil {
			properties := make(map[string]interface{})
			properties["reserved"] = SetResourceUsageSubResourceData([]*models.ResourceUsage{ResourceCapacityUsageModel.Reserved})
			properties["used"] = SetResourceUsageSubResourceData([]*models.ResourceUsage{ResourceCapacityUsageModel.Used})
			d = append(d, &properties)
		}
	}
	return
}

func ResourceCapacityUsageSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"reserved": {
			Description: `Reserved resource usage`,
			Type:        schema.TypeList, //GoType: ResourceUsage
			Elem: &schema.Resource{
				Schema: ResourceUsageSchema(),
			},
			Optional: true,
		},

		"used": {
			Description: `Used resource usage`,
			Type:        schema.TypeList, //GoType: ResourceUsage
			Elem: &schema.Resource{
				Schema: ResourceUsageSchema(),
			},
			Optional: true,
		},
	}
}

func GetResourceCapacityUsagePropertyFields() (t []string) {
	return []string{
		"reserved",
		"used",
	}
}
