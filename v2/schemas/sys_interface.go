package schemas

import (
	"reflect"
	"slices"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SysInterfaceModel(d *schema.ResourceData) *models.SysInterface {
	var adapterSpecificNet *models.Network // NetConfig
	adapterSpecificNetInterface, adapterSpecificNetIsSet := d.GetOk("adapter_specific_net")
	if adapterSpecificNetIsSet && adapterSpecificNetInterface != nil {
		adapterSpecificNetMap := adapterSpecificNetInterface.([]interface{})
		if len(adapterSpecificNetMap) > 0 {
			adapterSpecificNet = NetworkModelFromMap(adapterSpecificNetMap[0].(map[string]interface{}))
		}
	}
	costInt, _ := d.Get("cost").(int)
	cost := int64(costInt)
	var intfUsage *models.AdapterUsage // AdapterUsage
	intfUsageInterface, intfUsageIsSet := d.GetOk("intf_usage")
	if intfUsageIsSet {
		intfUsageModel := intfUsageInterface.(string)
		if intfUsageModel != "" {
			intfUsage = models.NewAdapterUsage(models.AdapterUsage(intfUsageModel))
		}
	}
	intfname, _ := d.Get("intfname").(string)
	ipaddr, _ := d.Get("ipaddr").(string)
	macaddr, _ := d.Get("macaddr").(string)
	var netDhcp *models.NetworkDHCPType // NetworkDHCPType
	netDhcpInterface, netDhcpIsSet := d.GetOk("net_dhcp")
	if netDhcpIsSet {
		netDhcpModel := netDhcpInterface.(string)
		if netDhcpModel != "" {
			netDhcp = models.NewNetworkDHCPType(models.NetworkDHCPType(netDhcpModel))
		}
	}
	netid, _ := d.Get("netid").(string)
	netname, _ := d.Get("netname").(string)
	var sharedLabels []string
	sharedLabelsInterface, sharedLabelsIsSet := d.GetOk("sharedLabels")
	if sharedLabelsIsSet {
		var items []interface{}
		if listItems, isList := sharedLabelsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = sharedLabelsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			sharedLabels = append(sharedLabels, v.(string))
		}
	}
	tags := map[string]string{}
	tagsInterface, tagsIsSet := d.GetOk("tags")
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	ztype, _ := d.Get("ztype").(string)
	return &models.SysInterface{
		AdapterSpecificNet: adapterSpecificNet,
		Cost:               cost,
		IntfUsage:          intfUsage,
		Intfname:           intfname,
		Ipaddr:             ipaddr,
		Macaddr:            macaddr,
		NetDhcp:            netDhcp,
		Netid:              netid,
		Netname:            netname,
		SharedLabels:       sharedLabels,
		Tags:               tags,
		Ztype:              ztype,
	}
}

func SysInterfaceModelFromMap(m map[string]interface{}) *models.SysInterface {
	var adapterSpecificNet *models.Network // NetConfig
	adapterSpecificNetInterface, adapterSpecificNetIsSet := m["adapter_specific_net"]
	if adapterSpecificNetIsSet && adapterSpecificNetInterface != nil {
		adapterSpecificNetMap := adapterSpecificNetInterface.([]interface{})
		if len(adapterSpecificNetMap) > 0 {
			adapterSpecificNet = NetworkModelFromMap(adapterSpecificNetMap[0].(map[string]interface{}))
		}
	}
	//
	cost := int64(m["cost"].(int))     // int64
	var intfUsage *models.AdapterUsage // AdapterUsage
	intfUsageInterface, intfUsageIsSet := m["intf_usage"]
	if intfUsageIsSet {
		intfUsageModel := intfUsageInterface.(string)
		if intfUsageModel != "" {
			intfUsage = models.NewAdapterUsage(models.AdapterUsage(intfUsageModel))
		}
	}
	intfname := m["intfname"].(string)
	ipaddr := m["ipaddr"].(string)
	macaddr := m["macaddr"].(string)
	var netDhcp *models.NetworkDHCPType // NetworkDHCPType
	netDhcpInterface, netDhcpIsSet := m["net_dhcp"]
	if netDhcpIsSet {
		netDhcpModel := netDhcpInterface.(string)
		if netDhcpModel != "" {
			netDhcp = models.NewNetworkDHCPType(models.NetworkDHCPType(netDhcpModel))
		}
	}
	netid := m["netid"].(string)
	netname := m["netname"].(string)
	var sharedLabels []string
	sharedLabelsInterface, sharedLabelsIsSet := m["sharedLabels"]
	if sharedLabelsIsSet {
		var items []interface{}
		if listItems, isList := sharedLabelsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = sharedLabelsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			sharedLabels = append(sharedLabels, v.(string))
		}
	}
	tags := map[string]string{}
	tagsInterface, tagsIsSet := m["tags"]
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	ztype := m["ztype"].(string)
	return &models.SysInterface{
		AdapterSpecificNet: adapterSpecificNet,
		Cost:               cost,
		IntfUsage:          intfUsage,
		Intfname:           intfname,
		Ipaddr:             ipaddr,
		Macaddr:            macaddr,
		NetDhcp:            netDhcp,
		Netid:              netid,
		Netname:            netname,
		SharedLabels:       sharedLabels,
		Tags:               tags,
		Ztype:              ztype,
	}
}

func SetSysInterfaceResourceData(d *schema.ResourceData, m *models.SysInterface) {
	d.Set("adapter_specific_net", SetNetworkSubResourceData([]*models.Network{m.AdapterSpecificNet}))
	d.Set("cost", m.Cost)
	d.Set("intf_usage", m.IntfUsage)
	d.Set("intfname", m.Intfname)
	d.Set("ipaddr", m.Ipaddr)
	d.Set("macaddr", m.Macaddr)
	d.Set("net_dhcp", m.NetDhcp)
	d.Set("netid", m.Netid)
	d.Set("netname", m.Netname)
	d.Set("shared_labels", m.SharedLabels)
	d.Set("tags", m.Tags)
	d.Set("ztype", m.Ztype)
}

func SetSysInterfaceSubResourceData(m []*models.SysInterface) (d []*map[string]interface{}) {
	for _, SysInterfaceModel := range m {
		if SysInterfaceModel != nil {
			properties := make(map[string]interface{})
			properties["adapter_specific_net"] = SetNetworkSubResourceData([]*models.Network{SysInterfaceModel.AdapterSpecificNet})
			properties["cost"] = SysInterfaceModel.Cost
			properties["intf_usage"] = SysInterfaceModel.IntfUsage
			properties["intfname"] = SysInterfaceModel.Intfname
			properties["ipaddr"] = SysInterfaceModel.Ipaddr
			properties["macaddr"] = SysInterfaceModel.Macaddr
			properties["net_dhcp"] = SysInterfaceModel.NetDhcp
			properties["netid"] = SysInterfaceModel.Netid
			properties["netname"] = SysInterfaceModel.Netname
			properties["shared_labels"] = SysInterfaceModel.SharedLabels
			properties["tags"] = SysInterfaceModel.Tags
			properties["ztype"] = SysInterfaceModel.Ztype
			d = append(d, &properties)
		}
	}
	return
}

func SysInterfaceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"adapter_specific_net": {
			Description: `device adapter specific network configuration`,
			Type:        schema.TypeList, //GoType: NetConfig
			Elem: &schema.Resource{
				Schema: Network(),
			},
			Optional: true,
		},

		"cost": {
			Description: `cost of using this interface. Default is 0.`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"intf_usage": {
			Description: `Adapter Udage`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"intfname": {
			Description: `name of interface in the manifest to which this network or adapter maps to`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"ipaddr": {
			Description: `IP address: we will be needing this in cae of static network`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"macaddr": {
			Description: `mac address needs to be over-written in some cases`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"net_dhcp": {
			Description: `Network DHCP type`,
			Type:        schema.TypeString,
			Default:     models.NetworkDHCPTypeNETWORKDHCPTYPEUNSPECIFIED,
			Optional:    true,
		},

		"netid": {
			Description: `network identifier has to have value if the netname is not empty`,
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
		},

		"netname": {
			Description: `network name: if attaching a network use netname`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"shared_labels": {
			Description: `A set of user-defined shared labels attached to the adapter`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"tags": {
			Description: `Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"ztype": {
			Description: `Z Type`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetSysInterfacePropertyFields() (t []string) {
	return []string{
		"adapter_specific_net",
		"cost",
		"intf_usage",
		"intfname",
		"ipaddr",
		"macaddr",
		"net_dhcp",
		"netid",
		"netname",
		"shared_labels",
		"tags",
		"ztype",
	}
}

func CompareSystemInterfaceList(a, b []*models.SysInterface) bool {
	if len(a) != len(b) {
		return false
	}

	slices.SortFunc(a, func(i, j *models.SysInterface) int {
		if n := strings.Compare(strings.ToLower(i.Intfname), strings.ToLower(j.Intfname)); n != 0 {
			return n
		}
		return strings.Compare(strings.ToLower(string(*i.IntfUsage)), strings.ToLower(string(*j.IntfUsage)))
	})

	slices.SortFunc(b, func(i, j *models.SysInterface) int {
		if n := strings.Compare(strings.ToLower(i.Intfname), strings.ToLower(j.Intfname)); n != 0 {
			return n
		}
		return strings.Compare(strings.ToLower(string(*i.IntfUsage)), strings.ToLower(string(*j.IntfUsage)))
	})

	equal := slices.EqualFunc(a, b, func(x, y *models.SysInterface) bool {
		if x.Cost != y.Cost {
			return false
		}
		if x.IntfUsage != nil && y.IntfUsage != nil {
			if *x.IntfUsage != *y.IntfUsage {
				return false
			}
		} else if x.IntfUsage != y.IntfUsage { // one of them is nil
			return false
		}
		if x.Intfname != y.Intfname {
			return false
		}
		if x.Ipaddr != y.Ipaddr {
			return false
		}
		if x.Macaddr != y.Macaddr {
			return false
		}
		if x.Netname != y.Netname {
			return false
		}
		if x.Netid != y.Netid {
			return false
		}
		if x.NetDhcp != nil && y.NetDhcp != nil {
			if *x.NetDhcp != *y.NetDhcp {
				return false
			}
		} else if x.NetDhcp != y.NetDhcp { // one of them is nil
			return false
		}
		if x.Ztype != y.Ztype {
			return false
		}
		// Compare Tags
		if (x.Tags == nil && len(y.Tags) == 0) || (y.Tags == nil && len(x.Tags) == 0) {
			// treat nil and empty map as equal
		} else if !reflect.DeepEqual(x.Tags, y.Tags) {
			return false
		}

		if len(x.SharedLabels) != len(y.SharedLabels) {
			return false
		}
		slices.Sort(x.SharedLabels)
		slices.Sort(y.SharedLabels)
		if !reflect.DeepEqual(x.SharedLabels, y.SharedLabels) {
			return false
		}
		return true
	})
	return equal
}
