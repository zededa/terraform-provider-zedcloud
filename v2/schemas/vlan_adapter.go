package schemas

import (
	"fmt"
	"slices"
	"strings"

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

// CompareVlanAdapterLists compares two VlanAdapter lists for equality, ignoring order
func CompareVlanAdapterLists(a, b []*models.VlanAdapter) (bool, string) {
	if len(a) != len(b) {
		return false, fmt.Sprintf("vlan_adapters length mismatch: %d vs %d", len(a), len(b))
	}

	// Sort both slices by LogicalLabel
	slices.SortFunc(a, func(i, j *models.VlanAdapter) int {
		return strings.Compare(strings.ToLower(i.LogicalLabel), strings.ToLower(j.LogicalLabel))
	})

	slices.SortFunc(b, func(i, j *models.VlanAdapter) int {
		return strings.Compare(strings.ToLower(i.LogicalLabel), strings.ToLower(j.LogicalLabel))
	})

	reason := ""
	equal := slices.EqualFunc(a, b, func(x, y *models.VlanAdapter) bool {
		if x.InterfaceName != y.InterfaceName {
			reason = fmt.Sprintf("InterfaceName mismatch: %s vs %s", x.InterfaceName, y.InterfaceName)
			return false
		}
		if x.LogicalLabel != y.LogicalLabel {
			reason = fmt.Sprintf("LogicalLabel mismatch: %s vs %s", x.LogicalLabel, y.LogicalLabel)
			return false
		}
		if x.LowerLayerName != y.LowerLayerName {
			reason = fmt.Sprintf("LowerLayerName mismatch: %s vs %s", x.LowerLayerName, y.LowerLayerName)
			return false
		}
		if x.VlanID != y.VlanID {
			reason = fmt.Sprintf("VlanID mismatch: %d vs %d", x.VlanID, y.VlanID)
			return false
		}

		// Compare nested Interface (SysInterface)
		if x.Interface != nil && y.Interface != nil {
			interfaceMatch, interfaceReason := CompareSystemInterfaceList([]*models.SysInterface{x.Interface}, []*models.SysInterface{y.Interface})
			if !interfaceMatch {
				reason = fmt.Sprintf("Interface mismatch: %s", interfaceReason)
				return false
			}
		} else if x.Interface != y.Interface {
			reason = fmt.Sprintf("Interface mismatch: one is nil, other is not")
			return false
		}

		return true
	})
	return equal, reason
}
