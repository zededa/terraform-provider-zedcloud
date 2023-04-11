package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
	"golang.org/x/exp/slices"
)

func NetworkModel(d *schema.ResourceData) *models.Network {
	description, _ := d.Get("description").(string)
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
	enterpriseDefault, _ := d.Get("enterprise_default").(bool)
	id, _ := d.Get("id").(string)
	var ip *models.IPSpec // IPSpec
	ipInterface, ipIsSet := d.GetOk("ip")
	if ipIsSet && ipInterface != nil {
		ipMap := ipInterface.([]interface{})
		if len(ipMap) > 0 {
			ip = IPSpecModelFromMap(ipMap[0].(map[string]interface{}))
		}
	}
	var kind *models.NetworkKind // NetworkKind
	kindInterface, kindIsSet := d.GetOk("kind")
	if kindIsSet {
		kindModel := kindInterface.(string)
		kind = models.NewNetworkKind(models.NetworkKind(kindModel))
	}
	name, _ := d.Get("name").(string)
	projectID, _ := d.Get("project_id").(string)
	var proxy *models.Proxy // NetProxyConfig
	proxyInterface, proxyIsSet := d.GetOk("proxy")
	if proxyIsSet && proxyInterface != nil {
		proxyMap := proxyInterface.([]interface{})
		if len(proxyMap) > 0 {
			proxy = NetworkProxyModelFromMap(proxyMap[0].(map[string]interface{}))
		}
	}
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := d.GetOk("revision")
	if revisionIsSet && revisionInterface != nil {
		revisionMap := revisionInterface.([]interface{})
		if len(revisionMap) > 0 {
			revision = ObjectRevisionModelFromMap(revisionMap[0].(map[string]interface{}))
		}
	}
	title, _ := d.Get("title").(string)
	var wireless *models.Wireless // NetWirelessConfig
	wirelessInterface, wirelessIsSet := d.GetOk("wireless")
	if wirelessIsSet && wirelessInterface != nil {
		wirelessMap := wirelessInterface.([]interface{})
		if len(wirelessMap) > 0 {
			wireless = NetworkWirelessModelFromMap(wirelessMap[0].(map[string]interface{}))
		}
	}
	return &models.Network{
		Description:       description,
		DNSList:           dNSList,
		EnterpriseDefault: enterpriseDefault,
		ID:                id,
		IP:                ip,
		Kind:              kind,
		Name:              &name,
		ProjectID:         &projectID,
		Proxy:             proxy,
		Revision:          revision,
		Title:             &title,
		Wireless:          wireless,
	}
}

func NetworkModelFromMap(m map[string]interface{}) *models.Network {
	description := m["description"].(string)
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
	enterpriseDefault := m["enterprise_default"].(bool)
	id := m["id"].(string)
	var ip *models.IPSpec // IPSpec
	ipInterface, ipIsSet := m["ip"]
	if ipIsSet && ipInterface != nil {
		ipMap := ipInterface.([]interface{})
		if len(ipMap) > 0 {
			ip = IPSpecModelFromMap(ipMap[0].(map[string]interface{}))
		}
	}
	//
	var kind *models.NetworkKind // NetworkKind
	kindInterface, kindIsSet := m["kind"]
	if kindIsSet {
		kindModel := kindInterface.(string)
		kind = models.NewNetworkKind(models.NetworkKind(kindModel))
	}
	name := m["name"].(string)
	projectID := m["project_id"].(string)
	var proxy *models.Proxy // NetProxyConfig
	proxyInterface, proxyIsSet := m["proxy"]
	if proxyIsSet && proxyInterface != nil {
		proxyMap := proxyInterface.([]interface{})
		if len(proxyMap) > 0 {
			proxy = NetworkProxyModelFromMap(proxyMap[0].(map[string]interface{}))
		}
	}
	//
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := m["revision"]
	if revisionIsSet && revisionInterface != nil {
		revisionMap := revisionInterface.([]interface{})
		if len(revisionMap) > 0 {
			revision = ObjectRevisionModelFromMap(revisionMap[0].(map[string]interface{}))
		}
	}
	//
	title := m["title"].(string)
	var wireless *models.Wireless // NetWirelessConfig
	wirelessInterface, wirelessIsSet := m["wireless"]
	if wirelessIsSet && wirelessInterface != nil {
		wirelessMap := wirelessInterface.([]interface{})
		if len(wirelessMap) > 0 {
			wireless = NetworkWirelessModelFromMap(wirelessMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.Network{
		Description:       description,
		DNSList:           dNSList,
		EnterpriseDefault: enterpriseDefault,
		ID:                id,
		IP:                ip,
		Kind:              kind,
		Name:              &name,
		ProjectID:         &projectID,
		Proxy:             proxy,
		Revision:          revision,
		Title:             &title,
		Wireless:          wireless,
	}
}

func SetNetworkResourceData(d *schema.ResourceData, m *models.Network) {
	d.Set("description", m.Description)
	d.Set("dns_list", SetStaticDNSListSubResourceData(m.DNSList))
	d.Set("enterprise_default", m.EnterpriseDefault)
	d.Set("id", m.ID)
	d.Set("ip", SetIPSpecSubResourceData([]*models.IPSpec{m.IP}))
	d.Set("kind", m.Kind)
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	d.Set("proxy", SetNetProxyConfigSubResourceData([]*models.Proxy{m.Proxy}))
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("title", m.Title)
	d.Set("wireless", SetNetWirelessConfigSubResourceData([]*models.Wireless{m.Wireless}))
}

func SetNetworkSubResourceData(m []*models.Network) (d []*map[string]interface{}) {
	for _, NetConfigModel := range m {
		if NetConfigModel != nil {
			properties := make(map[string]interface{})
			properties["description"] = NetConfigModel.Description
			properties["dns_list"] = SetStaticDNSListSubResourceData(NetConfigModel.DNSList)
			properties["enterprise_default"] = NetConfigModel.EnterpriseDefault
			properties["id"] = NetConfigModel.ID
			properties["ip"] = SetIPSpecSubResourceData([]*models.IPSpec{NetConfigModel.IP})
			properties["kind"] = NetConfigModel.Kind
			properties["name"] = NetConfigModel.Name
			properties["project_id"] = NetConfigModel.ProjectID
			properties["proxy"] = SetNetProxyConfigSubResourceData([]*models.Proxy{NetConfigModel.Proxy})
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{NetConfigModel.Revision})
			properties["title"] = NetConfigModel.Title
			properties["wireless"] = SetNetWirelessConfigSubResourceData([]*models.Wireless{NetConfigModel.Wireless})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the Network resource defined in the Terraform configuration.
func Network() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Description: `Detailed description of the network`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"dns_list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*StaticDNSList
			Elem: &schema.Resource{
				Schema: StaticDNSList(),
			},
			DiffSuppressFunc: diffSuppressDNSListOrder("dns_list"),
			Optional:         true,
			Computed:         true,
		},

		"enterprise_default": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"id": {
			Description: `System defined universally unique Id of the network`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"ip": {
			Description: "IP configuration for the network",
			Type:        schema.TypeList, //GoType: IPSpec
			Elem: &schema.Resource{
				Schema: IPSpec(),
			},
			Required: true,
		},

		"kind": {
			Description: `Kind of network:
NETWORK_KIND_V4
NETWORK_KIND_V6`,
			Type:     schema.TypeString,
			Required: true,
		},

		"name": {
			Description: `User defined name of the network, unique across the enterprise. Once object is created, name can’t be changed`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"project_id": {
			Description: ``,
			Type:        schema.TypeString,
			Required:    true,
		},

		"proxy": {
			Description: `proxy block is used to configure network proxy settings. The following is a brief description of how to use the attributes in the block:
1) If the proxy server requires certificates, set network_proxy_certs to carry the certificates
2) To have the EdgeNode auto discover pacfile, set network_proxy to True
3) To configure EdgeNode to download the pacfile from a specific URL:
    a) set network_proxy = false
	b) set network_proxy_url to the URL of the pac file
4) To configure EdgeNode with the contents of the pacfile, set 'pacfile' to content of the pacfile.
5) To configure Proxy setting manually instead of using a pacfile, use the 'proxy' and 'exceptions' blocks`,
			Type: schema.TypeList, //GoType: NetProxyConfig
			Elem: &schema.Resource{
				Schema: NetworkProxy(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional:         true,
			Computed:         true,
			DiffSuppressFunc: diffSuppressProxyListOrder("proxy"),
		},

		"revision": {
			Description: `system defined info`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Optional: true,
			Computed: true,
		},

		"title": {
			Description: `User defined title of the network. Title can be changed at any time`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"wireless": {
			Description: ``,
			Type:        schema.TypeList, //GoType: NetWirelessConfig
			Elem: &schema.Resource{
				Schema: NetworkWireless(),
			},
			Optional: true,
			Computed: true,
		},
	}
}

// Retrieve property field names for updating the Network resource
func NetworkPropertyFields() (t []string) {
	return []string{
		"description",
		"dns_list",
		"enterprise_default",
		"id",
		"ip",
		"kind",
		"name",
		"project_id",
		"proxy",
		"revision",
		"title",
		"wireless",
	}
}

func CompareProxyLists(a, b []*models.Proxy) bool {
	if len(a) != len(b) {
		return false
	}
	// is each element of the new list in the old list?
	for _, newList := range b {
		if newList == nil {
			continue
		}

		found := false
		for _, oldList := range a {
			if oldList == nil {
				continue
			}
			if oldList.Exceptions != newList.Exceptions {
				continue
			}
			if oldList.NetworkProxy != newList.NetworkProxy {
				continue
			}
			if oldList.NetworkProxyURL != newList.NetworkProxyURL {
				continue
			}
			if oldList.Pacfile != newList.Pacfile {
				continue
			}
			slices.Sort(oldList.NetworkProxyCerts)
			slices.Sort(newList.NetworkProxyCerts)
			if !Equal(oldList.NetworkProxyCerts, newList.NetworkProxyCerts) {
				continue
			}
			if !CompareProxyServer(oldList.Proxies, newList.Proxies) {
				continue
			}
			found = true
			break
		}
		if !found {
			return false
		}
	}

	// is each element of the old list also in the new list?
	for _, oldList := range a {
		if oldList == nil {
			continue
		}

		found := false
		for _, newList := range b {
			if newList == nil {
				continue
			}
			if oldList.Exceptions != newList.Exceptions {
				continue
			}
			if oldList.NetworkProxy != newList.NetworkProxy {
				continue
			}
			if oldList.NetworkProxyURL != newList.NetworkProxyURL {
				continue
			}
			if oldList.Pacfile != newList.Pacfile {
				continue
			}
			slices.Sort(oldList.NetworkProxyCerts)
			slices.Sort(newList.NetworkProxyCerts)
			if !Equal(oldList.NetworkProxyCerts, newList.NetworkProxyCerts) {
				continue
			}
			if !CompareProxyServer(oldList.Proxies, newList.Proxies) {
				continue
			}
			found = true
			break
		}
		if !found {
			return false
		}
	}
	return true
}
