// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var netWifiConfigSecretsSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"wifi_passwd": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "SENSITIVE: Wifi Password. " +
				"This field will not be published by terraform import",
		},
	},
}

var netWifiConfigSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"identity": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "SENSITIVE. WPA2 Enterprise user identity/username. Use Value from Vault. " +
				"This field will not be published by terraform import",
		},
		"key_scheme": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Key management scheme. Valid Values: NETWORK_WIFIKEY_SCHEME_WPAPSK, " +
				"NETWORK_WIFIKEY_SCHEME_WPAEAPPSK",
		},
		"priority": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Priority of connection, default is 0",
		},
		"secret": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Description: "SENSITIVE. List of wifi secrets (passwords). " +
				"This field will not be published by terraform import command.",
			Elem: netWifiConfigSecretsSchema,
		},
		"wifi_ssid": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "WIFI SSID",
		},
	},
}

var netCellularConfigSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"apn": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "APN string",
		},
	},
}

var netWirelessConfigSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"cellular_cfg": &schema.Schema{
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Cellular config",
			Elem:        netCellularConfigSchema,
		},
		"type": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Description: "Type of Wireless Network. Valid values: NETWORK_WIRELESS_TYPE_WIFI, " +
				"NETWORK_WIRELESS_TYPE_CELLULAR",
		},
		"wifi_cfg": &schema.Schema{
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Wifi, can be multiple APs on a single wlan, e.g. one for 2.5Ghz, other 5Ghz SSIDs",
			Elem:        netWifiConfigSchema,
		},
	},
}

var netProxyServerSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"port": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Net Proxy Port",
		},
		"proto": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Description: "Net Proxy proto. Valid Values: NETWORK_PROXY_PROTO_HTTP, " +
				"NETWORK_PROXY_PROTO_HTTPS, NETWORK_PROXY_PROTO_SOCKS, NETWORK_PROXY_PROTO_FTP, " +
				"NETWORK_PROXY_PROTO_OTHER",
		},
		"server": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Net Proxy Server",
		},
	},
}

var netProxyConfigSchema = &schema.Resource{
	Description: "netProxyConfig block is used to configure Network Proxy settings for " +
		"the network. The following is a brief description of how to use the attributes " +
		"in the block:\n" +
		"1) If the proxy server requires certificates, set network_proxy_certs to carry the " +
		"   certificated" +
		"2) To have the EdgeNode auto discover pacfile, set network_proxy to True \n" +
		"3) To Configure EdgeNode to download the pacfile from a specific URL:\n" +
		"       a) set network_proxy = false\n" +
		"       b) set network_proxy_url to the URL of the pac file\n" +
		"4) To configure EdgeNode with the contents of the pacfile, set 'pacfile' " +
		"       to content of the pacfile.\n" +
		"5) To configure Proxy setting manually instead of using a pacfile, use the " +
		"   'proxy' and 'exceptions' blocks",
	Schema: map[string]*schema.Schema{
		"exceptions": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Proxy exceptions",
		},
		"network_proxy": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Enable WPAD (Web Proxy Auto Discovery) Protocol to discover and download PAC file.",
		},
		"network_proxy_certs": &schema.Schema{
			Type:        schema.TypeList,
			Optional:    true,
			Description: "List of Network Proxy Certificates",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"network_proxy_url": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "URL for wpad.dat file to be downloaded. Used when " +
				"network_proxy is set to False.",
		},
		"pacfile": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Proxy configuration in a pacfile. Used when network_proxy is set to False.",
		},
		"proxy": &schema.Schema{
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Net Proxy: protocol level proxies. Used when network_proxy is set to False.",
			Elem:        netProxyServerSchema,
		},
	},
}

var ipSpecResourceSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"dhcp": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Description: "DHCP Type (Static / Client etc.). Valid Values: NETWORK_DHCP_TYPE_STATIC, " +
				"NETWORK_DHCP_TYPE_PASSTHROUGH, NETWORK_DHCP_TYPE_CLIENT",
		},
		"dhcp_range": &schema.Schema{
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Range of IP addresses to be used for DHCP",
			Elem:        dhcpIpRangeSchema,
		},
		"dns": &schema.Schema{
			Type:        schema.TypeList,
			Optional:    true,
			Description: "List of IP Addresses of DNS servers",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"domain": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Network domain",
		},
		"gateway": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: "IP Address of Network Gateway",
		},
		"mask": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Subnet Mask",
		},
		"ntp": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: "IP Address of NTP Server",
		},
		"subnet": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Subnet address",
		},
	},
}

var StaticDNSEntryResourceSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"addrs": &schema.Schema{
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Set of IP addresses for the specified hostname",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"hostname": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: "DNS Host name",
		},
	},
}

var NetworkSchema = map[string]*schema.Schema{
	// Keep the following common fields at the top of schema definitions for all
	//  objects.
	"name":        name,
	"id":          id,
	"description": description,
	"title":       title,

	// Rest of the fields must be in the alphabetical order of keys
	"dns_list": &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "List of Static DNS entries",
		Elem:        StaticDNSEntryResourceSchema,
	},
	"enterprise_default": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Flag to make this network default for device",
		Default:     false,
	},
	"ip": &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "IP configuration for the Network",
		Elem:        ipSpecResourceSchema,
	},
	"kind": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Kind of network. Valid Values: NETWORK_KIND_V4, NETWORK_KIND_V6",
	},
	"project_id": projectID,
	"proxy": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Proxy configuration for the network",
		Elem:        netProxyConfigSchema,
	},
	"revision": revision,
	"wireless": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Wireless (Wifi / Cellular) Configuration for the Network",
		Elem:        netWirelessConfigSchema,
	},
}
