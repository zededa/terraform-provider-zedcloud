package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AppInterfaceModel(d *schema.ResourceData) *models.AppInterface {
	accessVlanIDInt, _ := d.Get("access_vlan_id").(int)
	accessVlanID := int64(accessVlanIDInt)
	var acls []*models.AppACE // []*AppACE
	aclsInterface, aclsIsSet := d.GetOk("acls")
	if aclsIsSet {
		var items []interface{}
		if listItems, isList := aclsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = aclsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AppACEModelFromMap(v.(map[string]interface{}))
			acls = append(acls, m)
		}
	}
	defaultNetInstance, _ := d.Get("default_net_instance").(bool)
	directattach, _ := d.Get("directattach").(bool)
	var eidregister *models.EIDRegister // EIDRegister
	eidregisterInterface, eidregisterIsSet := d.GetOk("eidregister")
	if eidregisterIsSet && eidregisterInterface != nil {
		eidregisterMap := eidregisterInterface.([]interface{})
		if len(eidregisterMap) > 0 {
			eidregister = EIDRegisterModelFromMap(eidregisterMap[0].(map[string]interface{}))
		}
	}
	intfname, _ := d.Get("intfname").(string)
	intforderInt, _ := d.Get("intforder").(int)
	intforder := int64(intforderInt)
	var io *models.PhyAdapter // PhyAdapter
	ioInterface, ioIsSet := d.GetOk("io")
	if ioIsSet && ioInterface != nil {
		ioMap := ioInterface.([]interface{})
		if len(ioMap) > 0 {
			io = PhyAdapterModelFromMap(ioMap[0].(map[string]interface{}))
		}
	}
	ipaddr, _ := d.Get("ipaddr").(string)
	macaddr, _ := d.Get("macaddr").(string)
	netinstid, _ := d.Get("netinstid").(string)
	netinstname, _ := d.Get("netinstname").(string)
	netinsttag := map[string]string{}
	netinsttagInterface, netinsttagIsSet := d.GetOk("netinsttag")
	if netinsttagIsSet {
		netinsttagMap := netinsttagInterface.(map[string]interface{})
		for k, v := range netinsttagMap {
			if v == nil {
				continue
			}
			netinsttag[k] = v.(string)
		}
	}

	netname, _ := d.Get("netname").(string)
	privateip, _ := d.Get("privateip").(bool)
	return &models.AppInterface{
		AccessVlanID:       accessVlanID,
		Acls:               acls,
		DefaultNetInstance: defaultNetInstance,
		Directattach:       &directattach, // bool true false false
		Eidregister:        eidregister,
		Intfname:           &intfname,  // string true false false
		Intforder:          &intforder, // int64 true false false
		Io:                 io,
		Ipaddr:             &ipaddr,  // string true false false
		Macaddr:            &macaddr, // string true false false
		Netinstid:          netinstid,
		Netinstname:        &netinstname, // string true false false
		Netinsttag:         netinsttag,
		Netname:            netname,
		Privateip:          &privateip, // string true false false
	}
}

func AppInterfaceModelFromMap(m map[string]interface{}) *models.AppInterface {
	accessVlanID := int64(m["access_vlan_id"].(int)) // int64
	var acls []*models.AppACE                        // []*AppACE
	aclsInterface, aclsIsSet := m["acls"]
	if aclsIsSet {
		var items []interface{}
		if listItems, isList := aclsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = aclsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AppACEModelFromMap(v.(map[string]interface{}))
			acls = append(acls, m)
		}
	}
	defaultNetInstance := m["default_net_instance"].(bool)
	directattach := m["directattach"].(bool)
	var eidregister *models.EIDRegister // EIDRegister
	eidregisterInterface, eidregisterIsSet := m["eidregister"]
	if eidregisterIsSet && eidregisterInterface != nil {
		eidregisterMap := eidregisterInterface.([]interface{})
		if len(eidregisterMap) > 0 {
			eidregister = EIDRegisterModelFromMap(eidregisterMap[0].(map[string]interface{}))
		}
	}
	//
	intfname := m["intfname"].(string)
	intforder := int64(m["intforder"].(int)) // int64
	var io *models.PhyAdapter                // PhyAdapter
	ioInterface, ioIsSet := m["io"]
	if ioIsSet && ioInterface != nil {
		ioMap := ioInterface.([]interface{})
		if len(ioMap) > 0 {
			io = PhyAdapterModelFromMap(ioMap[0].(map[string]interface{}))
		}
	}
	//
	ipaddr := m["ipaddr"].(string)
	macaddr := m["macaddr"].(string)
	netinstid := m["netinstid"].(string)
	netinstname := m["netinstname"].(string)
	netinsttag := map[string]string{}
	netinsttagInterface, netinsttagIsSet := m["netinsttag"]
	if netinsttagIsSet {
		netinsttagMap := netinsttagInterface.(map[string]interface{})
		for k, v := range netinsttagMap {
			if v == nil {
				continue
			}
			netinsttag[k] = v.(string)
		}
	}

	netname := m["netname"].(string)
	privateip := m["privateip"].(bool)
	return &models.AppInterface{
		AccessVlanID:       accessVlanID,
		Acls:               acls,
		DefaultNetInstance: defaultNetInstance,
		Directattach:       &directattach,
		Eidregister:        eidregister,
		Intfname:           &intfname,
		Intforder:          &intforder,
		Io:                 io,
		Ipaddr:             &ipaddr,
		Macaddr:            &macaddr,
		Netinstid:          netinstid,
		Netinstname:        &netinstname,
		Netinsttag:         netinsttag,
		Netname:            netname,
		Privateip:          &privateip,
	}
}

// Update the underlying AppInterface resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppInterfaceResourceData(d *schema.ResourceData, m *models.AppInterface) {
	d.Set("access_vlan_id", m.AccessVlanID)
	d.Set("acls", SetAppACESubResourceData(m.Acls))
	d.Set("default_net_instance", m.DefaultNetInstance)
	d.Set("directattach", m.Directattach)
	d.Set("eidregister", SetEIDRegisterSubResourceData([]*models.EIDRegister{m.Eidregister}))
	d.Set("intfname", m.Intfname)
	d.Set("intforder", m.Intforder)
	d.Set("io", SetPhyAdapterSubResourceData([]*models.PhyAdapter{m.Io}))
	d.Set("ipaddr", m.Ipaddr)
	d.Set("macaddr", m.Macaddr)
	d.Set("netinstid", m.Netinstid)
	d.Set("netinstname", m.Netinstname)
	d.Set("netinsttag", m.Netinsttag)
	d.Set("netname", m.Netname)
	d.Set("privateip", m.Privateip)
}

// Iterate through and update the AppInterface resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppInterfaceSubResourceData(m []*models.AppInterface) (d []*map[string]interface{}) {
	for _, AppInterfaceModel := range m {
		if AppInterfaceModel != nil {
			properties := make(map[string]interface{})
			properties["access_vlan_id"] = AppInterfaceModel.AccessVlanID
			properties["acls"] = SetAppACESubResourceData(AppInterfaceModel.Acls)
			properties["default_net_instance"] = AppInterfaceModel.DefaultNetInstance
			properties["directattach"] = AppInterfaceModel.Directattach
			properties["eidregister"] = SetEIDRegisterSubResourceData([]*models.EIDRegister{AppInterfaceModel.Eidregister})
			properties["intfname"] = AppInterfaceModel.Intfname
			properties["intforder"] = AppInterfaceModel.Intforder
			properties["io"] = SetPhyAdapterSubResourceData([]*models.PhyAdapter{AppInterfaceModel.Io})
			properties["ipaddr"] = AppInterfaceModel.Ipaddr
			properties["macaddr"] = AppInterfaceModel.Macaddr
			properties["netinstid"] = AppInterfaceModel.Netinstid
			properties["netinstname"] = AppInterfaceModel.Netinstname
			properties["netinsttag"] = AppInterfaceModel.Netinsttag
			properties["netname"] = AppInterfaceModel.Netname
			properties["privateip"] = AppInterfaceModel.Privateip
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppInterface resource defined in the Terraform configuration
func AppInterfaceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access_vlan_id": {
			Description: `access port VLAN ID, vlan id of zero will be treated as trunk port and vlan id 1 is implicitly used by linux bridges`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"acls": {
			Description: `app Acls`,
			Type:        schema.TypeList, //GoType: []*AppACE
			Elem: &schema.Resource{
				Schema: AppACESchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"default_net_instance": {
			Description: `default instance flag`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"directattach": {
			Description: `direct attach flag`,
			Type:        schema.TypeBool,
			Default:     false,
			Optional:    true,
		},

		"eidregister": {
			Description: `EID register details`,
			Type:        schema.TypeList, //GoType: EIDRegister
			Elem: &schema.Resource{
				Schema: EIDRegisterSchema(),
			},
			Optional: true,
		},

		"intfname": {
			Description: `intf Name`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"intforder": {
			Description: `intforder`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"io": {
			Description: `Physical Adapter to be matched for interface assignment. Applicable only when "direct attach" flag is true`,
			Type:        schema.TypeList, //GoType: PhyAdapter
			Elem: &schema.Resource{
				Schema: PhyAdapterSchema(),
			},
			Optional: true,
		},

		"ipaddr": {
			Description: `IP address`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"macaddr": {
			Description: `MAC address`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"netinstid": {
			Description: `Network Instance id to be matched for interface assignment.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"netinstname": {
			Description: `Network Instance name to be matched for interface assignment. Applicable only when "direct attach" flag is false`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"netinsttag": {
			Description: `Network Instance tag to be matched for interface assignment. Applicable only when "direct attach" flag is false`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"netname": {
			Description: `network name: will be deprecated in future, use netinstname`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"privateip": {
			Description: `Private IP flag`,
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the AppInterface resource
func GetAppInterfacePropertyFields() (t []string) {
	return []string{
		"access_vlan_id",
		"acls",
		"default_net_instance",
		"directattach",
		"eidregister",
		"intfname",
		"intforder",
		"io",
		"ipaddr",
		"macaddr",
		"netinstid",
		"netinstname",
		"netinsttag",
		"netname",
		"privateip",
	}
}
