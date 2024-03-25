package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func ProxyServerModel(d *schema.ResourceData) *models.Server {
	portInt, _ := d.Get("port").(int)
	port := int64(portInt)
	var proto *models.Proto // NetworkProxyProto
	protoInterface, protoIsSet := d.GetOk("proto")
	if protoIsSet {
		protoModel := protoInterface.(string)
		proto = models.NewNetworkProxyProto(models.Proto(protoModel))
	}
	server, _ := d.Get("server").(string)
	return &models.Server{
		Port:   port,
		Proto:  proto,
		Server: server,
	}
}

func ProxyServerModelFromMap(m map[string]interface{}) *models.Server {
	port := int64(m["port"].(int)) // int64
	var proto *models.Proto        // NetworkProxyProto
	protoInterface, protoIsSet := m["proto"]
	if protoIsSet {
		protoModel := protoInterface.(string)
		proto = models.NewNetworkProxyProto(models.Proto(protoModel))
	}
	server := m["server"].(string)
	return &models.Server{
		Port:   port,
		Proto:  proto,
		Server: server,
	}
}

func SetProxyServerResourceData(d *schema.ResourceData, m *models.Server) {
	d.Set("port", m.Port)
	d.Set("proto", m.Proto)
	d.Set("server", m.Server)
}

func SetProxyServerSubResourceData(m []*models.Server) (d []*map[string]interface{}) {
	for _, NetProxyServerModel := range m {
		if NetProxyServerModel != nil {
			properties := make(map[string]interface{})
			properties["port"] = NetProxyServerModel.Port
			properties["proto"] = NetProxyServerModel.Proto
			properties["server"] = NetProxyServerModel.Server
			d = append(d, &properties)
		}
	}
	return
}

func ProxyServer() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"port": {
			Description: `Net Proxy Port`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"proto": {
			Description: `Net Proxy proto`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"server": {
			Description: `Net Proxy Server`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the NetProxyServer resource
func GetNetProxyServerPropertyFields() (t []string) {
	return []string{
		"port",
		"proto",
		"server",
	}
}

func CompareProxyServer(a, b []*models.Server) bool {
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
			if oldList.Port != newList.Port {
				continue
			}
			if *oldList.Proto != *newList.Proto {
				continue
			}
			if oldList.Server != newList.Server {
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
			if oldList.Port != newList.Port {
				continue
			}
			if *oldList.Proto != *newList.Proto {
				continue
			}
			if oldList.Server != newList.Server {
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
