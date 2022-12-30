package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate NetworkInstConfig resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func NetworkInstConfigModel(d *schema.ResourceData) *models.NetworkInstConfig {
	deviceDefault, _ := d.Get("device_default").(string)
	dNSList, _ := d.Get("dns_list").([]*models.StaticDNSList) // []*StaticDNSList
	var ip *models.DhcpServerConfig                           // DhcpServerConfig
	ipInterface, ipIsSet := d.GetOk("ip")
	if ipIsSet {
		ipMap := ipInterface.([]interface{})[0].(map[string]interface{})
		ip = DhcpServerConfigModelFromMap(ipMap)
	}
	var kind *models.NetworkInstanceKind // NetworkInstanceKind
	kindInterface, kindIsSet := d.GetOk("kind")
	if kindIsSet {
		kindModel := kindInterface.(models.NetworkInstanceKind)
		kind = &kindModel
	}
	var opaque *models.NetInstOpaqueConfig // NetInstOpaqueConfig
	opaqueInterface, opaqueIsSet := d.GetOk("opaque")
	if opaqueIsSet {
		opaqueMap := opaqueInterface.([]interface{})[0].(map[string]interface{})
		opaque = NetInstOpaqueConfigModelFromMap(opaqueMap)
	}
	port, _ := d.Get("port").(string)
	portTags, _ := d.Get("port_tags").(map[string]string) // map[string]string
	tags, _ := d.Get("tags").(map[string]string)          // map[string]string
	var typeVar *models.NetworkInstanceDhcpType           // NetworkInstanceDhcpType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(models.NetworkInstanceDhcpType)
		typeVar = &typeModel
	}
	return &models.NetworkInstConfig{
		DeviceDefault: deviceDefault,
		DNSList:       dNSList,
		IP:            ip,
		Kind:          kind,
		Opaque:        opaque,
		Port:          port,
		PortTags:      portTags,
		Tags:          tags,
		Type:          typeVar,
	}
}

func NetworkInstConfigModelFromMap(m map[string]interface{}) *models.NetworkInstConfig {
	deviceDefault := m["device_default"].(string)
	dNSList := m["dns_list"].([]*models.StaticDNSList) // []*StaticDNSList
	var ip *models.DhcpServerConfig                    // DhcpServerConfig
	ipInterface, ipIsSet := m["ip"]
	if ipIsSet {
		ipMap := ipInterface.([]interface{})[0].(map[string]interface{})
		ip = DhcpServerConfigModelFromMap(ipMap)
	}
	//
	kind := m["kind"].(*models.NetworkInstanceKind) // NetworkInstanceKind
	var opaque *models.NetInstOpaqueConfig          // NetInstOpaqueConfig
	opaqueInterface, opaqueIsSet := m["opaque"]
	if opaqueIsSet {
		opaqueMap := opaqueInterface.([]interface{})[0].(map[string]interface{})
		opaque = NetInstOpaqueConfigModelFromMap(opaqueMap)
	}
	//
	port := m["port"].(string)
	portTags := m["port_tags"].(map[string]string)
	tags := m["tags"].(map[string]string)
	typeVar := m["type"].(*models.NetworkInstanceDhcpType) // NetworkInstanceDhcpType
	return &models.NetworkInstConfig{
		DeviceDefault: deviceDefault,
		DNSList:       dNSList,
		IP:            ip,
		Kind:          kind,
		Opaque:        opaque,
		Port:          port,
		PortTags:      portTags,
		Tags:          tags,
		Type:          typeVar,
	}
}

// Update the underlying NetworkInstConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetworkInstConfigResourceData(d *schema.ResourceData, m *models.NetworkInstConfig) {
	d.Set("device_default", m.DeviceDefault)
	d.Set("dns_list", SetStaticDNSListSubResourceData(m.DNSList))
	d.Set("ip", SetDhcpServerConfigSubResourceData([]*models.DhcpServerConfig{m.IP}))
	d.Set("kind", m.Kind)
	d.Set("opaque", SetNetInstOpaqueConfigSubResourceData([]*models.NetInstOpaqueConfig{m.Opaque}))
	d.Set("port", m.Port)
	d.Set("port_tags", m.PortTags)
	d.Set("tags", m.Tags)
	d.Set("type", m.Type)
}

// Iterate through and update the NetworkInstConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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
			properties["tags"] = NetworkInstConfigModel.Tags
			properties["type"] = NetworkInstConfigModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetworkInstConfig resource defined in the Terraform configuration
func NetworkInstConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"device_default": {
			Description: `flag to indicate if this is the default network instance for the device`,
			Type:        schema.TypeString,
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
				Schema: DhcpServerConfigSchema(),
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

// Retrieve property field names for updating the NetworkInstConfig resource
func GetNetworkInstConfigPropertyFields() (t []string) {
	return []string{
		"device_default",
		"dns_list",
		"ip",
		"kind",
		"opaque",
		"port",
		"port_tags",
		"tags",
		"type",
	}
}
