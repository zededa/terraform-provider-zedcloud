// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-test/deep"
	"github.com/zededa/zedcloud-api/swagger_models"
	"testing"
)

var ouputNetworkEmpty = map[string]interface{}{
	"description":        "",
	"dns_list":           []interface{}{},
	"enterprise_default": false,
	"id":                 "",
	"ip":                 []interface{}{},
	"kind":               "",
	"name":               "",
	"project_id":         "",
	"proxy":              []interface{}{},
	"revision":           []interface{}{},
	"title":              "",
	"wireless":           []interface{}{},
}

func netCfgFull() *swagger_models.NetConfig {
	return &swagger_models.NetConfig{
		Name:        strPtr("Test Full Cfg Name"),
		ID:          "Test Full Cfg ID",
		Description: "Test Full Cfg Description",
		Title:       strPtr("Test Full Cfg Title"),
		DNSList: []*swagger_models.StaticDNSList{
			&swagger_models.StaticDNSList{
				Addrs:    []string{"10.10.1.1", "10.10.2.2"},
				Hostname: "www.example.com",
			},
			&swagger_models.StaticDNSList{
				Addrs:    []string{"10.30.1.1", "10.4.22.2"},
				Hostname: "www.2.example.com",
			},
		},
		EnterpriseDefault: true,
		IP: &swagger_models.IPSpec{
			Dhcp: networkDHCPTypePtr("NETWORK_DHCP_TYPE_CLIENT"),
			DhcpRange: &swagger_models.DhcpIPRange{
				Start: "10.10.1.1",
				End:   "10.10.1.10",
			},
			DNS:     []string{"10.10.1.1", "10.10.1.2"},
			Domain:  "example.com",
			Gateway: "10.10.1.1",
			Mask:    "255.255.255.0",
			Ntp:     "10.10.1.1",
			Subnet:  "10.1.1.0/16",
		},
		Kind:      networkKindPtr("NETWORK_KIND_V4"),
		ProjectID: strPtr("Test Full Cfg project-id"),
		Proxy: &swagger_models.NetProxyConfig{
			Exceptions:   "exception1",
			NetworkProxy: true,
			NetworkProxyCerts: []strfmt.Base64{strfmt.Base64("cert1"),
				strfmt.Base64("cert2")},
			NetworkProxyURL: "www.example.com",
			Pacfile:         "testpacfile",
			Proxies: []*swagger_models.NetProxyServer{
				&swagger_models.NetProxyServer{
					Port:   30,
					Proto:  networkProxyProtoPtr("NETWORK_PROXY_PROTO_HTTPS"),
					Server: "example.com",
				},
				&swagger_models.NetProxyServer{
					Port:   40,
					Proto:  networkProxyProtoPtr("NETWORK_PROXY_PROTO_HTTP"),
					Server: "example.com",
				},
				&swagger_models.NetProxyServer{
					Port:   50,
					Proto:  networkProxyProtoPtr("NETWORK_PROXY_PROTO_OTHER"),
					Server: "example.com",
				},
			},
		},
		Revision: &swagger_models.ObjectRevision{
			CreatedAt: "2020-12-15T06:21:24Z",
			CreatedBy: strPtr("sample@example.com"),
			Curr:      strPtr("1"),
			Prev:      "0",
			UpdatedAt: "2020-12-15T06:21:24Z",
			UpdatedBy: strPtr("user@example.com"),
		},
		Wireless: &swagger_models.NetWirelessConfig{
			// In reality, we can have either Cellular cfg or wifi cfg depending
			// on type - not both. But for unit test, lets populate both structures.
			CellularCfg: &swagger_models.NetCellularConfig{
				APN: "Sample APN",
			},
			Type: networkWirelessTypePtr("NETWORK_WIRELESS_TYPE_WIFI"),
			WifiCfg: &swagger_models.NetWifiConfig{
				Identity:  "Sample Identity",
				KeyScheme: networkWiFiKeySchemePtr("NETWORK_WIFIKEY_SCHEME_WPAPSK"),
				Priority:  10,
				Secret: &swagger_models.NetWifiConfigSecrets{
					WiFiPasswd: "sample wifi password",
				},
				WifiSSID: "Sample Wifi SSID",
			},
		},
	}
}

var outputFullCfg = map[string]interface{}{
	"description": "Test Full Cfg Description",
	"dns_list": []interface{}{
		map[string]interface{}{
			"addrs": []interface{}{
				"10.10.1.1",
				"10.10.2.2",
			},
			"hostname": "www.example.com",
		},
		map[string]interface{}{
			"addrs": []interface{}{
				"10.30.1.1",
				"10.4.22.2",
			},
			"hostname": "www.2.example.com",
		},
	},
	"enterprise_default": true,
	"id":                 "Test Full Cfg ID",
	"ip": []interface{}{map[string]interface{}{
		"dhcp": "NETWORK_DHCP_TYPE_CLIENT",
		"dhcp_range": []interface{}{
			map[string]interface{}{
				"start": "10.10.1.1",
				"end":   "10.10.1.10",
			},
		},
		"dns": []interface{}{
			"10.10.1.1",
			"10.10.1.2",
		},
		"domain":  "example.com",
		"gateway": "10.10.1.1",
		"mask":    "255.255.255.0",
		"ntp":     "10.10.1.1",
		"subnet":  "10.1.1.0/16",
	}},
	"kind":       "NETWORK_KIND_V4",
	"name":       "Test Full Cfg Name",
	"project_id": "Test Full Cfg project-id",
	"proxy": []interface{}{
		map[string]interface{}{
			"exceptions":          "exception1",
			"network_proxy":       true,
			"network_proxy_certs": []interface{}{"cert1", "cert2"},
			"network_proxy_url":   "www.example.com",
			"pacfile":             "testpacfile",
			"proxy": []interface{}{
				map[string]interface{}{
					"port":   30,
					"proto":  "NETWORK_PROXY_PROTO_HTTPS",
					"server": "example.com",
				},
				map[string]interface{}{
					"port":   40,
					"proto":  "NETWORK_PROXY_PROTO_HTTP",
					"server": "example.com",
				},
				map[string]interface{}{
					"port":   50,
					"proto":  "NETWORK_PROXY_PROTO_OTHER",
					"server": "example.com",
				},
			},
		},
	},
	"revision": []interface{}{map[string]interface{}{
		"created_at": "2020-12-15T06:21:24Z",
		"created_by": "sample@example.com",
		"curr":       "1",
		"prev":       "0",
		"updated_at": "2020-12-15T06:21:24Z",
		"updated_by": "user@example.com",
	}},
	"title": "Test Full Cfg Title",
	"wireless": []interface{}{
		map[string]interface{}{
			// In reality, we can have either Cellular cfg or wifi cfg depending
			// on type - not both. But for unit test, lets populate both structures.
			"cellular_cfg": []interface{}{
				map[string]interface{}{
					"apn": "Sample APN",
				},
			},
			"type": "NETWORK_WIRELESS_TYPE_WIFI",
			"wifi_cfg": []interface{}{
				map[string]interface{}{
					"key_scheme": "NETWORK_WIFIKEY_SCHEME_WPAPSK",
					"priority":   10,
					"wifi_ssid":  "Sample Wifi SSID",
				},
			},
		},
	},
}

func TestFlattenNetworkConfig(t *testing.T) {
	cases := []struct {
		input       *swagger_models.NetConfig
		expected    map[string]interface{}
		description string
	}{
		{
			input:       nil,
			expected:    map[string]interface{}{},
			description: "nil Struct",
		},
		{
			input:       &swagger_models.NetConfig{},
			expected:    ouputNetworkEmpty,
			description: "Empty Struct",
		},
		{
			input:       netCfgFull(),
			expected:    outputFullCfg,
			description: "Full Cfg",
		},
	}

	for _, c := range cases {
		out := flattenNetConfig(c.input)
		if diff := deep.Equal(out, c.expected); diff != nil {
			t.Fatalf("Test Failed: %s\n"+
				"Error matching Flattened output and input.\n"+
				"Output: %#v\n"+
				"Expected Output : %#v\n"+
				"Diff Actual vs Expected: %#v", c.description, out,
				c.expected, diff)
		}
	}
}
