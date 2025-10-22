package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func VlanAdapterModel(d *schema.ResourceData) *models.VlanAdapter {
	var interfaceVar *models.SysInterface // SysInterface
	interfaceInterface, interfaceIsSet := d.GetOk("interface")
	if interfaceIsSet && interfaceInterface != nil {
		interfaceMap := interfaceInterface.([]interface{})
		if len(interfaceMap) > 0 {
			interfaceVar = SysInterfaceModelFromMap(interfaceMap[0].(map[string]interface{}))
		}
	}
	interfaceName, _ := d.Get("interface_name").(string)
	logicalLabel, _ := d.Get("logical_label").(string)
	lowerLayerName, _ := d.Get("lower_layer_name").(string)
	vlanIDInt, _ := d.Get("vlan_id").(int)
	vlanID := int64(vlanIDInt)
	return &models.VlanAdapter{
		Interface:      interfaceVar,
		InterfaceName:  interfaceName,
		LogicalLabel:   logicalLabel,
		LowerLayerName: lowerLayerName,
		VlanID:         vlanID,
	}
}

func VlanAdapterModelFromMap(m map[string]interface{}) *models.VlanAdapter {
	var interfaceVar *models.SysInterface // SysInterface
	interfaceInterface, interfaceIsSet := m["interface"]
	if interfaceIsSet && interfaceInterface != nil {
		interfaceMap := interfaceInterface.([]interface{})
		if len(interfaceMap) > 0 {
			interfaceVar = SysInterfaceModelFromMap(interfaceMap[0].(map[string]interface{}))
		}
	}
	//
	interfaceName := m["interface_name"].(string)
	logicalLabel := m["logical_label"].(string)
	lowerLayerName := m["lower_layer_name"].(string)
	vlanID := int64(m["vlan_id"].(int)) // int64
	return &models.VlanAdapter{
		Interface:      interfaceVar,
		InterfaceName:  interfaceName,
		LogicalLabel:   logicalLabel,
		LowerLayerName: lowerLayerName,
		VlanID:         vlanID,
	}
}

func SetVlanAdapterResourceData(d *schema.ResourceData, m *models.VlanAdapter) {
	d.Set("interface", SetSysInterfaceSubResourceData([]*models.SysInterface{m.Interface}))
	d.Set("interface_name", m.InterfaceName)
	d.Set("logical_label", m.LogicalLabel)
	d.Set("lower_layer_name", m.LowerLayerName)
	d.Set("vlan_id", m.VlanID)
}

func SetVlanAdapterSubResourceData(m []*models.VlanAdapter) (d []*map[string]interface{}) {
	for _, VlanAdapterModel := range m {
		if VlanAdapterModel != nil {
			properties := make(map[string]interface{})
			properties["interface"] = SetSysInterfaceSubResourceData([]*models.SysInterface{VlanAdapterModel.Interface})
			properties["interface_name"] = VlanAdapterModel.InterfaceName
			properties["logical_label"] = VlanAdapterModel.LogicalLabel
			properties["lower_layer_name"] = VlanAdapterModel.LowerLayerName
			properties["vlan_id"] = VlanAdapterModel.VlanID
			d = append(d, &properties)
		}
	}
	return
}

func VlanAdapterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"interface": {
			Description: `System Sub-Interface`,
			Type:        schema.TypeList, //GoType: SysInterface
			Elem: &schema.Resource{
				Schema: SysInterfaceSchema(),
			},
			Optional: true,
		},

		"interface_name": {
			Description: `A physical name of the VLAN sub-interface`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"logical_label": {
			Description: `Logical name of this VLAN adapter`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"lower_layer_name": {
			Description: `Logical name of the lower layer adapter (bond or physicalIO)`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"vlan_id": {
			Description: `Valid values are from 1 to 4094`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetVlanAdapterPropertyFields() (t []string) {
	return []string{
		"interface",
		"interface_name",
		"logical_label",
		"lower_layer_name",
		"vlan_id",
	}
}
