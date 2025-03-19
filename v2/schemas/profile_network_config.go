package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ProfileNetworkConfigModel(d *schema.ResourceData) *models.ProfileNetworkConfig {
	deviceDefault, _ := d.Get("device_default").(bool)
	var dNSList []*models.StaticDNSList // []*StaticDNSList
	dnsListInterface, dnsListIsSet := d.GetOk("dns_list")
	if dnsListIsSet {
		var items []interface{}
		if listItems, isList := dnsListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = dnsListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := StaticDNSListModelFromMap(v.(map[string]interface{}))
			dNSList = append(dNSList, m)
		}
	}
	var ip *models.DhcpServerConfig // DhcpServerConfig
	ipInterface, ipIsSet := d.GetOk("ip")
	if ipIsSet && ipInterface != nil {
		ipMap := ipInterface.([]interface{})
		if len(ipMap) > 0 {
			ip = DhcpServerConfigModelFromMap(ipMap[0].(map[string]interface{}))
		}
	}
	var kind *models.NetworkInstanceKind // NetworkInstanceKind
	kindInterface, kindIsSet := d.GetOk("kind")
	if kindIsSet {
		kindModel := kindInterface.(string)
		kind = models.NewNetworkInstanceKind(models.NetworkInstanceKind(kindModel))
	}
	var opaque *models.NetInstOpaqueConfig // NetInstOpaqueConfig
	opaqueInterface, opaqueIsSet := d.GetOk("opaque")
	if opaqueIsSet && opaqueInterface != nil {
		opaqueMap := opaqueInterface.([]interface{})
		if len(opaqueMap) > 0 {
			opaque = NetInstOpaqueConfigModelFromMap(opaqueMap[0].(map[string]interface{}))
		}
	}
	port, _ := d.Get("port").(string)
	portTags := map[string]string{}
	portTagsInterface, portTagsIsSet := d.GetOk("portTags")
	if portTagsIsSet {
		portTagsMap := portTagsInterface.(map[string]interface{})
		for k, v := range portTagsMap {
			if v == nil {
				continue
			}
			portTags[k] = v.(string)
		}
	}

	propagateConnectedRoutes, _ := d.Get("propagate_connected_routes").(bool)
	var staticRoutes []*models.StaticIPRoute // []*StaticIPRoute
	staticRoutesInterface, staticRoutesIsSet := d.GetOk("static_routes")
	if staticRoutesIsSet {
		var items []interface{}
		if listItems, isList := staticRoutesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = staticRoutesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := StaticIPRouteModelFromMap(v.(map[string]interface{}))
			staticRoutes = append(staticRoutes, m)
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

	var typeVar *models.NetworkInstanceDhcpType // NetworkInstanceDhcpType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewNetworkInstanceDhcpType(models.NetworkInstanceDhcpType(typeModel))
	}
	return &models.ProfileNetworkConfig{
		DeviceDefault:            deviceDefault,
		DNSList:                  dNSList,
		IP:                       ip,
		Kind:                     kind,
		Opaque:                   opaque,
		Port:                     port,
		PortTags:                 portTags,
		PropagateConnectedRoutes: propagateConnectedRoutes,
		StaticRoutes:             staticRoutes,
		Tags:                     tags,
		Type:                     typeVar,
	}
}

func ProfileNetworkConfigModelFromMap(m map[string]interface{}) *models.ProfileNetworkConfig {
	deviceDefault := m["device_default"].(bool)
	var dNSList []*models.StaticDNSList // []*StaticDNSList
	dnsListInterface, dnsListIsSet := m["dns_list"]
	if dnsListIsSet {
		var items []interface{}
		if listItems, isList := dnsListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = dnsListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := StaticDNSListModelFromMap(v.(map[string]interface{}))
			dNSList = append(dNSList, m)
		}
	}
	var ip *models.DhcpServerConfig // DhcpServerConfig
	ipInterface, ipIsSet := m["ip"]
	if ipIsSet && ipInterface != nil {
		ipMap := ipInterface.([]interface{})
		if len(ipMap) > 0 {
			ip = DhcpServerConfigModelFromMap(ipMap[0].(map[string]interface{}))
		}
	}
	//
	var kind *models.NetworkInstanceKind // NetworkInstanceKind
	kindInterface, kindIsSet := m["kind"]
	if kindIsSet {
		kindModel := kindInterface.(string)
		kind = models.NewNetworkInstanceKind(models.NetworkInstanceKind(kindModel))
	}
	var opaque *models.NetInstOpaqueConfig // NetInstOpaqueConfig
	opaqueInterface, opaqueIsSet := m["opaque"]
	if opaqueIsSet && opaqueInterface != nil {
		opaqueMap := opaqueInterface.([]interface{})
		if len(opaqueMap) > 0 {
			opaque = NetInstOpaqueConfigModelFromMap(opaqueMap[0].(map[string]interface{}))
		}
	}
	//
	port := m["port"].(string)
	portTags := map[string]string{}
	portTagsInterface, portTagsIsSet := m["port_tags"]
	if portTagsIsSet {
		portTagsMap := portTagsInterface.(map[string]interface{})
		for k, v := range portTagsMap {
			if v == nil {
				continue
			}
			portTags[k] = v.(string)
		}
	}

	propagateConnectedRoutes := m["propagate_connected_routes"].(bool)
	var staticRoutes []*models.StaticIPRoute // []*StaticIPRoute
	staticRoutesInterface, staticRoutesIsSet := m["static_routes"]
	if staticRoutesIsSet {
		var items []interface{}
		if listItems, isList := staticRoutesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = staticRoutesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := StaticIPRouteModelFromMap(v.(map[string]interface{}))
			staticRoutes = append(staticRoutes, m)
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

	var typeVar *models.NetworkInstanceDhcpType // NetworkInstanceDhcpType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewNetworkInstanceDhcpType(models.NetworkInstanceDhcpType(typeModel))
	}
	return &models.ProfileNetworkConfig{
		DeviceDefault:            deviceDefault,
		DNSList:                  dNSList,
		IP:                       ip,
		Kind:                     kind,
		Opaque:                   opaque,
		Port:                     port,
		PortTags:                 portTags,
		PropagateConnectedRoutes: propagateConnectedRoutes,
		StaticRoutes:             staticRoutes,
		Tags:                     tags,
		Type:                     typeVar,
	}
}

func SetProfileNetworkConfigResourceData(d *schema.ResourceData, m *models.ProfileNetworkConfig) {
	d.Set("device_default", m.DeviceDefault)
	d.Set("dns_list", SetStaticDNSListSubResourceData(m.DNSList))
	d.Set("ip", SetDhcpServerConfigSubResourceData([]*models.DhcpServerConfig{m.IP}))
	d.Set("kind", m.Kind)
	d.Set("opaque", SetNetInstOpaqueConfigSubResourceData([]*models.NetInstOpaqueConfig{m.Opaque}))
	d.Set("port", m.Port)
	d.Set("port_tags", m.PortTags)
	d.Set("propagate_connected_routes", m.PropagateConnectedRoutes)
	d.Set("static_routes", SetStaticIPRouteSubResourceData(m.StaticRoutes))
	d.Set("tags", m.Tags)
	d.Set("type", m.Type)
}

func SetProfileNetworkConfigSubResourceData(m []*models.ProfileNetworkConfig) (d []*map[string]interface{}) {
	for _, ProfileNetworkConfigModel := range m {
		if ProfileNetworkConfigModel != nil {
			properties := make(map[string]interface{})
			properties["device_default"] = ProfileNetworkConfigModel.DeviceDefault
			properties["dns_list"] = SetStaticDNSListSubResourceData(ProfileNetworkConfigModel.DNSList)
			properties["ip"] = SetDhcpServerConfigSubResourceData([]*models.DhcpServerConfig{ProfileNetworkConfigModel.IP})
			properties["kind"] = ProfileNetworkConfigModel.Kind
			properties["opaque"] = SetNetInstOpaqueConfigSubResourceData([]*models.NetInstOpaqueConfig{ProfileNetworkConfigModel.Opaque})
			properties["port"] = ProfileNetworkConfigModel.Port
			properties["port_tags"] = ProfileNetworkConfigModel.PortTags
			properties["propagate_connected_routes"] = ProfileNetworkConfigModel.PropagateConnectedRoutes
			properties["static_routes"] = SetStaticIPRouteSubResourceData(ProfileNetworkConfigModel.StaticRoutes)
			properties["tags"] = ProfileNetworkConfigModel.Tags
			properties["type"] = ProfileNetworkConfigModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func ProfileNetworkConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"device_default": {
			Description: `flag to indicate if this is the default network instance for the device`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"dns_list": {
			Description: `List of Static DNS entries`,
			Type:        schema.TypeList, //GoType: []*StaticDNSList
			Elem: &schema.Resource{
				Schema: StaticDNSListSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"ip": {
			Description: `Dhcp Server Configuration`,
			Type:        schema.TypeList, //GoType: DhcpServerConfig
			Elem: &schema.Resource{
				Schema: DhcpServerConfig(),
			},
			Optional: true,
		},

		"kind": {
			Description: `Kind of Network Instance ( Local, Switch etc )`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"opaque": {
			Description: `Service specific Config`,
			Type:        schema.TypeList, //GoType: NetInstOpaqueConfig
			Elem: &schema.Resource{
				Schema: NetInstOpaqueConfigSchema(),
			},
			Optional: true,
		},

		"port": {
			Description: `name of port mapping in the model`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"port_tags": {
			Description: `Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"propagate_connected_routes": {
			Description: `Automatically propagate connected routes`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"static_routes": {
			Description: `List of Static IP routes`,
			Type:        schema.TypeList, //GoType: []*StaticIPRoute
			Elem: &schema.Resource{
				Schema: StaticIPRouteSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
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

		"type": {
			Description: `Type of DHCP for this Network Instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetProfileNetworkConfigPropertyFields() (t []string) {
	return []string{
		"device_default",
		"dns_list",
		"ip",
		"kind",
		"opaque",
		"port",
		"port_tags",
		"propagate_connected_routes",
		"static_routes",
		"tags",
		"type",
	}
}
