package schemas

import (
	"fmt"
	"slices"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func BondAdapterModel(d *schema.ResourceData) *models.BondAdapter {
	var arp *models.ArpMonitor // ArpMonitor
	arpInterface, arpIsSet := d.GetOk("arp")
	if arpIsSet && arpInterface != nil {
		arpMap := arpInterface.([]interface{})
		if len(arpMap) > 0 {
			arp = ArpMonitorModelFromMap(arpMap[0].(map[string]interface{}))
		}
	}
	var bondMode *models.BondMode // BondMode
	bondModeInterface, bondModeIsSet := d.GetOk("bond_mode")
	if bondModeIsSet {
		bondModeModel := bondModeInterface.(string)
		bondMode = models.NewBondMode(models.BondMode(bondModeModel))
	}
	var interfaceVar *models.SysInterface // SysInterface
	interfaceInterface, interfaceIsSet := d.GetOk("interface")
	if interfaceIsSet && interfaceInterface != nil {
		interfaceMap := interfaceInterface.([]interface{})
		if len(interfaceMap) > 0 {
			interfaceVar = SysInterfaceModelFromMap(interfaceMap[0].(map[string]interface{}))
		}
	}
	interfaceName, _ := d.Get("interface_name").(string)
	lacpRate := models.LacpRateLACPRATEUNSPECIFIED.Pointer()
	lacpRateInterface, lacpRateIsSet := d.GetOk("lacp_rate")
	if lacpRateIsSet {
		lacpRateModel := lacpRateInterface.(string)
		if lacpRateModel != "" {
			lacpRate = models.NewLacpRate(models.LacpRate(lacpRateModel))
		}
	}
	logicalLabel, _ := d.Get("logical_label").(string)
	var lowerLayerNames []string
	lowerLayerNamesInterface, lowerLayerNamesIsSet := d.GetOk("lower_layer_names")
	if lowerLayerNamesIsSet {
		var items []interface{}
		if listItems, isList := lowerLayerNamesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = lowerLayerNamesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			lowerLayerNames = append(lowerLayerNames, v.(string))
		}
	}
	var mii *models.MIIMonitor // MIIMonitor
	miiInterface, miiIsSet := d.GetOk("mii")
	if miiIsSet && miiInterface != nil {
		miiMap := miiInterface.([]interface{})
		if len(miiMap) > 0 {
			mii = MIIMonitorModelFromMap(miiMap[0].(map[string]interface{}))
		}
	}
	return &models.BondAdapter{
		Arp:             arp,
		BondMode:        bondMode,
		Interface:       interfaceVar,
		InterfaceName:   interfaceName,
		LacpRate:        lacpRate,
		LogicalLabel:    logicalLabel,
		LowerLayerNames: lowerLayerNames,
		Mii:             mii,
	}
}

func BondAdapterModelFromMap(m map[string]interface{}) *models.BondAdapter {
	var arp *models.ArpMonitor // ArpMonitor
	arpInterface, arpIsSet := m["arp"]
	if arpIsSet && arpInterface != nil {
		arpMap := arpInterface.([]interface{})
		if len(arpMap) > 0 {
			arp = ArpMonitorModelFromMap(arpMap[0].(map[string]interface{}))
		}
	}
	//
	var bondMode *models.BondMode // BondMode
	bondModeInterface, bondModeIsSet := m["bond_mode"]
	if bondModeIsSet {
		bondModeModel := bondModeInterface.(string)
		bondMode = models.NewBondMode(models.BondMode(bondModeModel))
	}
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
	lacpRate := models.LacpRateLACPRATEUNSPECIFIED.Pointer()
	lacpRateInterface, lacpRateIsSet := m["lacp_rate"]
	if lacpRateIsSet {
		lacpRateModel := lacpRateInterface.(string)
		if lacpRateModel != "" {
			lacpRate = models.NewLacpRate(models.LacpRate(lacpRateModel))
		}
	}
	logicalLabel := m["logical_label"].(string)
	var lowerLayerNames []string
	lowerLayerNamesInterface, lowerLayerNamesIsSet := m["lower_layer_names"]
	if lowerLayerNamesIsSet {
		var items []interface{}
		if listItems, isList := lowerLayerNamesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = lowerLayerNamesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			lowerLayerNames = append(lowerLayerNames, v.(string))
		}
	}
	var mii *models.MIIMonitor // MIIMonitor
	miiInterface, miiIsSet := m["mii"]
	if miiIsSet && miiInterface != nil {
		miiMap := miiInterface.([]interface{})
		if len(miiMap) > 0 {
			mii = MIIMonitorModelFromMap(miiMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.BondAdapter{
		Arp:             arp,
		BondMode:        bondMode,
		Interface:       interfaceVar,
		InterfaceName:   interfaceName,
		LacpRate:        lacpRate,
		LogicalLabel:    logicalLabel,
		LowerLayerNames: lowerLayerNames,
		Mii:             mii,
	}
}

func SetBondAdapterResourceData(d *schema.ResourceData, m *models.BondAdapter) {
	d.Set("arp", SetArpMonitorSubResourceData([]*models.ArpMonitor{m.Arp}))
	d.Set("bond_mode", m.BondMode)
	d.Set("interface", SetSysInterfaceSubResourceData([]*models.SysInterface{m.Interface}))
	d.Set("interface_name", m.InterfaceName)
	d.Set("lacp_rate", m.LacpRate)
	d.Set("logical_label", m.LogicalLabel)
	d.Set("lower_layer_names", m.LowerLayerNames)
	d.Set("mii", SetMIIMonitorSubResourceData([]*models.MIIMonitor{m.Mii}))
}

func SetBondAdapterSubResourceData(m []*models.BondAdapter) (d []*map[string]interface{}) {
	for _, BondAdapterModel := range m {
		if BondAdapterModel != nil {
			properties := make(map[string]interface{})
			properties["arp"] = SetArpMonitorSubResourceData([]*models.ArpMonitor{BondAdapterModel.Arp})
			properties["bond_mode"] = BondAdapterModel.BondMode
			properties["interface"] = SetSysInterfaceSubResourceData([]*models.SysInterface{BondAdapterModel.Interface})
			properties["interface_name"] = BondAdapterModel.InterfaceName
			properties["lacp_rate"] = BondAdapterModel.LacpRate
			properties["logical_label"] = BondAdapterModel.LogicalLabel
			properties["lower_layer_names"] = BondAdapterModel.LowerLayerNames
			properties["mii"] = SetMIIMonitorSubResourceData([]*models.MIIMonitor{BondAdapterModel.Mii})
			d = append(d, &properties)
		}
	}
	return
}

func BondAdapterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"arp": {
			Description: `ARP monitor`,
			Type:        schema.TypeList, //GoType: ArpMonitor
			Elem: &schema.Resource{
				Schema: ArpMonitorSchema(),
			},
			Optional: true,
		},

		"bond_mode": {
			Description: `Bonding mode`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"interface": {
			Description: `System Sub-Interface`,
			Type:        schema.TypeList, //GoType: SysInterface
			Elem: &schema.Resource{
				Schema: SysInterfaceSchema(),
			},
			Optional: true,
		},

		"interface_name": {
			Description: `A physical name of the bond interface`,
			Type:        schema.TypeString,
			// Computed because the backend auto-assigns a name (e.g. "bond0", "bond1").
			Optional: true,
			Computed: true,
		},

		"lacp_rate": {
			Description: `LACPDU transmission rate in 802.3ad mode`,
			Type:        schema.TypeString,
			Optional:    true,
			// The backend always returns LACP_RATE_UNSPECIFIED when lacpRate is not set,
			// so use it as the default to avoid a perpetual plan diff.
			Default: string(models.LacpRateLACPRATEUNSPECIFIED),
		},

		"logical_label": {
			Description: `Logical name of this bond adapter`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"lower_layer_names": {
			Description: `Logical names of aggregated PhysicalIOs`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"mii": {
			Description: `MII monitor`,
			Type:        schema.TypeList, //GoType: MIIMonitor
			Elem: &schema.Resource{
				Schema: MIIMonitorSchema(),
			},
			Optional: true,
		},
	}
}

// CompareBondAdapterLists compares two BondAdapter lists for equality, ignoring order
func CompareBondAdapterLists(a, b []*models.BondAdapter) (bool, string) {
	if len(a) != len(b) {
		return false, fmt.Sprintf("bond_adapters length mismatch: %d vs %d", len(a), len(b))
	}

	slices.SortFunc(a, func(i, j *models.BondAdapter) int {
		return strings.Compare(strings.ToLower(i.LogicalLabel), strings.ToLower(j.LogicalLabel))
	})
	slices.SortFunc(b, func(i, j *models.BondAdapter) int {
		return strings.Compare(strings.ToLower(i.LogicalLabel), strings.ToLower(j.LogicalLabel))
	})

	reason := ""
	equal := slices.EqualFunc(a, b, func(x, y *models.BondAdapter) bool {
		if x.LogicalLabel != y.LogicalLabel {
			reason = fmt.Sprintf("LogicalLabel mismatch: %s vs %s", x.LogicalLabel, y.LogicalLabel)
			return false
		}
		// InterfaceName is auto-assigned by the backend (e.g. "bond0", "bond1") and
		// cannot be predicted from Terraform config, so it is excluded from comparison.
		if x.BondMode != nil && y.BondMode != nil {
			if *x.BondMode != *y.BondMode {
				reason = fmt.Sprintf("BondMode mismatch: %s vs %s", *x.BondMode, *y.BondMode)
				return false
			}
		} else if x.BondMode != y.BondMode {
			reason = "BondMode mismatch: one is nil, other is not"
			return false
		}
		// Treat nil as LACP_RATE_UNSPECIFIED: the schema Default always sends
		// LACP_RATE_UNSPECIFIED to the API, so the backend returns a non-nil
		// pointer even when the field was not set in the Terraform config.
		// YAML snapshots that omit lacprate must still compare equal.
		xLacpRate := models.LacpRateLACPRATEUNSPECIFIED
		if x.LacpRate != nil {
			xLacpRate = *x.LacpRate
		}
		yLacpRate := models.LacpRateLACPRATEUNSPECIFIED
		if y.LacpRate != nil {
			yLacpRate = *y.LacpRate
		}
		if xLacpRate != yLacpRate {
			reason = fmt.Sprintf("LacpRate mismatch: %s vs %s", xLacpRate, yLacpRate)
			return false
		}
		xNames := slices.Clone(x.LowerLayerNames)
		yNames := slices.Clone(y.LowerLayerNames)
		slices.Sort(xNames)
		slices.Sort(yNames)
		if !slices.Equal(xNames, yNames) {
			reason = fmt.Sprintf("LowerLayerNames mismatch: %v vs %v", x.LowerLayerNames, y.LowerLayerNames)
			return false
		}
		if x.Arp != nil && y.Arp != nil {
			if x.Arp.Interval != y.Arp.Interval {
				reason = fmt.Sprintf("Arp.Interval mismatch: %d vs %d", x.Arp.Interval, y.Arp.Interval)
				return false
			}
		} else if x.Arp != y.Arp {
			reason = "Arp mismatch: one is nil, other is not"
			return false
		}
		return true
	})
	return equal, reason
}

func GetBondAdapterPropertyFields() (t []string) {
	return []string{
		"arp",
		"bond_mode",
		"interface",
		"interface_name",
		"lacp_rate",
		"logical_label",
		"lower_layer_names",
		"mii",
	}
}
