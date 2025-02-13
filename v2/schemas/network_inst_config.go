package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func NetworkInstConfigModel(d *schema.ResourceData) *models.NetworkInstConfig {
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
	return &models.NetworkInstConfig{
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

func NetworkInstConfigModelFromMap(m map[string]interface{}) *models.NetworkInstConfig {
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
	return &models.NetworkInstConfig{
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

func SetNetworkInstConfigResourceData(d *schema.ResourceData, m *models.NetworkInstConfig) {
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

func SetNetworkInstConfigSubResourceData(m []*models.NetworkInstConfig) (d []*map[string]interface{}) {
	for _, NetworkInstConfigModel := range m {
		if NetworkInstConfigModel != nil {
			properties := make(map[string]interface{})
			properties["device_default"] = NetworkInstConfigModel.DeviceDefault
			properties["dns_list"] = SetStaticDNSListSubResourceData(NetworkInstConfigModel.DNSList)
			properties["ip"] = SetDhcpServerConfigSubResourceData([]*models.DhcpServerConfig{NetworkInstConfigModel.IP})
			properties["kind"] = NetworkInstConfigModel.Kind
			properties["opaque"] = SetNetInstOpaqueConfigSubResourceData([]*models.NetInstOpaqueConfig{NetworkInstConfigModel.Opaque})
			properties["port"] = NetworkInstConfigModel.Port
			properties["port_tags"] = NetworkInstConfigModel.PortTags
			properties["propagate_connected_routes"] = NetworkInstConfigModel.PropagateConnectedRoutes
			properties["static_routes"] = SetStaticIPRouteSubResourceData(NetworkInstConfigModel.StaticRoutes)
			properties["tags"] = NetworkInstConfigModel.Tags
			properties["type"] = NetworkInstConfigModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func NetworkInstConfig() map[string]*schema.Schema {
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
				Schema: StaticDNSList(),
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

func GetNetworkInstConfigPropertyFields() (t []string) {
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
