package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func NetworkProxyModel(d *schema.ResourceData) *models.Proxy {
	exceptions, _ := d.Get("exceptions").(string)
	networkProxy, _ := d.Get("network_proxy").(bool)
	var networkProxyCerts []string
	networkProxyCertsInterface, networkProxyCertsIsSet := d.GetOk("networkProxyCerts")
	if networkProxyCertsIsSet {
		networkProxyCertsSlice := networkProxyCertsInterface.([]interface{})
		for _, i := range networkProxyCertsSlice {
			networkProxyCertsSlice = append(networkProxyCertsSlice, i.(string))
		}
	}
	networkProxyURL, _ := d.Get("network_proxy_url").(string)
	pacfile, _ := d.Get("pacfile").(string)
	var proxies []*models.Server // []*NetProxyServer
	proxiesInterface, proxiesIsSet := d.GetOk("proxies")
	if proxiesIsSet {
		var items []interface{}
		if listItems, isList := proxiesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = proxiesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := NetworkProxyServerModelFromMap(v.(map[string]interface{}))
			proxies = append(proxies, m)
		}
	}
	return &models.Proxy{
		Exceptions:        exceptions,
		NetworkProxy:      networkProxy,
		NetworkProxyCerts: networkProxyCerts,
		NetworkProxyURL:   networkProxyURL,
		Pacfile:           pacfile,
		Proxies:           proxies,
	}
}

func NetworkProxyModelFromMap(m map[string]interface{}) *models.Proxy {
	exceptions := m["exceptions"].(string)
	networkProxy := m["network_proxy"].(bool)
	var networkProxyCerts []string
	networkProxyCertsInterface, networkProxyCertsIsSet := m["networkProxyCerts"]
	if networkProxyCertsIsSet {
		networkProxyCertsSlice := networkProxyCertsInterface.([]interface{})
		for _, i := range networkProxyCertsSlice {
			networkProxyCertsSlice = append(networkProxyCertsSlice, i.(string))
		}
	}
	networkProxyURL := m["network_proxy_url"].(string)
	pacfile := m["pacfile"].(string)
	var proxies []*models.Server // []*NetProxyServer
	proxiesInterface, proxiesIsSet := m["proxies"]
	if proxiesIsSet {
		var items []interface{}
		if listItems, isList := proxiesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = proxiesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := NetworkProxyServerModelFromMap(v.(map[string]interface{}))
			proxies = append(proxies, m)
		}
	}
	return &models.Proxy{
		Exceptions:        exceptions,
		NetworkProxy:      networkProxy,
		NetworkProxyCerts: networkProxyCerts,
		NetworkProxyURL:   networkProxyURL,
		Pacfile:           pacfile,
		Proxies:           proxies,
	}
}

func SetNetworkProxyResourceData(d *schema.ResourceData, m *models.Proxy) {
	d.Set("exceptions", m.Exceptions)
	d.Set("network_proxy", m.NetworkProxy)
	d.Set("network_proxy_certs", m.NetworkProxyCerts)
	d.Set("network_proxy_url", m.NetworkProxyURL)
	d.Set("pacfile", m.Pacfile)
	d.Set("proxies", SetNetworkProxyServerSubResourceData(m.Proxies))
}

func SetNetProxyConfigSubResourceData(m []*models.Proxy) (d []*map[string]interface{}) {
	for _, NetProxyConfigModel := range m {
		if NetProxyConfigModel != nil {
			properties := make(map[string]interface{})
			properties["exceptions"] = NetProxyConfigModel.Exceptions
			properties["network_proxy"] = NetProxyConfigModel.NetworkProxy
			properties["network_proxy_certs"] = NetProxyConfigModel.NetworkProxyCerts
			properties["network_proxy_url"] = NetProxyConfigModel.NetworkProxyURL
			properties["pacfile"] = NetProxyConfigModel.Pacfile
			properties["proxies"] = SetNetworkProxyServerSubResourceData(NetProxyConfigModel.Proxies)
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetworkProxy resource defined in the Terraform configuration
func NetworkProxy() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"exceptions": {
			Description: `Proxy exceptions`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"network_proxy": {
			Description: "Enable WPAD (Web Proxy Auto Discovery) protocol to discover and download PAC file.",
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"network_proxy_certs": {
			Description: `Network Proxy Certificates`,
			Type:        schema.TypeList, //GoType: []strfmt.Base64
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"network_proxy_url": {
			Description: "URL for wpad.dat file to be downloaded. Used when network_proxy is set to False.",
			Type:        schema.TypeString,
			Optional:    true,
		},

		"pacfile": {
			Description: "Proxy configuration in a pacfile. Used when network_proxy is set to False.",
			Type:        schema.TypeString,
			Optional:    true,
		},

		"proxies": {
			Description: "Net Proxy: protocol level proxies. Used when network_proxy is set to False.",
			Type:        schema.TypeList, //GoType: []*NetProxyServer
			Elem: &schema.Resource{
				Schema: NetworkProxyServer(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the NetProxyConfig resource
func NetworkProxyPropertyFields() (t []string) {
	return []string{
		"exceptions",
		"network_proxy",
		"network_proxy_certs",
		"network_proxy_url",
		"pacfile",
		"proxies",
	}
}
