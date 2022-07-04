// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"github.com/go-test/deep"
	"github.com/zededa/zedcloud-api/swagger_models"
	"testing"
)

var ouputNetInstEmpty = map[string]interface{}{
	"name":              "",
	"id":                "",
	"title":             "",
	"cluster_id":        "",
	"description":       "",
	"device_default":    false,
	"device_id":         "",
	"dns_list":          []interface{}{},
	"ip":                []interface{}{},
	"kind":              "",
	"network_policy_id": "",
	"opaque":            []interface{}{},
	"port":              "",
	"port_tags":         map[string]interface{}{},
	"project_id":        "",
	"revision":          []interface{}{},
	"tags":              map[string]interface{}{},
	"type":              "",
}

func netInstFullCfg() *swagger_models.NetInstConfig {
	kind := swagger_models.NetworkInstanceKind("NETWORK_INSTANCE_KIND_LOCAL")
	dhcpType := swagger_models.NetworkInstanceDhcpType("NETWORK_INSTANCE_DHCP_TYPE_V4")
	opaqueType := swagger_models.OpaqueConfigType("OPAQUE_CONFIG_TYPE_VPN")
	return &swagger_models.NetInstConfig{
		Name:          strPtr("SampleName"),
		ID:            "Sample ID",
		Description:   "Sample Description",
		Title:         strPtr("Sample Title"),
		ClusterID:     "Sample Cluster ID",
		DeviceDefault: true,
		DeviceID:      strPtr("sample device ID"),
		DNSList: []*swagger_models.StaticDNSList{
			&swagger_models.StaticDNSList{
				Addrs:    []string{"10.10.1.1", "10.10.2.2"},
				Hostname: "www.example.com",
			},
			&swagger_models.StaticDNSList{
				Addrs:    []string{"10.1.1.2", "10.2.2.3"},
				Hostname: "www.example2.com",
			},
		},
		IP: &swagger_models.DhcpServerConfig{
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
		Kind:            &kind,
		NetworkPolicyID: "sample policy id",
		Opaque: &swagger_models.NetInstOpaqueConfig{
			Oconfig: "Sample OConfig",
			Type:    &opaqueType,
		},
		Port: strPtr("uplink"),
		PortTags: map[string]string{
			"ptag1": "ptagval1",
			"ptag2": "ptagval2",
		},
		ProjectID: "sample-project-id",
		Revision: &swagger_models.ObjectRevision{
			CreatedAt: "2020-12-15T06:21:24Z",
			CreatedBy: strPtr("sample@example.com"),
			Curr:      strPtr("1"),
			Prev:      "0",
			UpdatedAt: "2020-12-15T06:21:24Z",
			UpdatedBy: strPtr("user@example.com"),
		},
		Tags: map[string]string{
			"tag1": "val1",
			"tag2": "val2",
		},
		Type: &dhcpType,
	}
}

var outputNetInstFullCfg = map[string]interface{}{
	"name":           "SampleName",
	"id":             "Sample ID",
	"description":    "Sample Description",
	"title":          "Sample Title",
	"cluster_id":     "Sample Cluster ID",
	"device_default": true,
	"device_id":      "sample device ID",
	"dns_list": []interface{}{
		map[string]interface{}{
			"addrs":    []interface{}{"10.10.1.1", "10.10.2.2"},
			"hostname": "www.example.com",
		},
		map[string]interface{}{
			"addrs":    []interface{}{"10.1.1.2", "10.2.2.3"},
			"hostname": "www.example2.com",
		},
	},
	"ip": []interface{}{
		map[string]interface{}{
			"dhcp_range": []interface{}{
				map[string]interface{}{
					"start": "10.10.1.1",
					"end":   "10.10.1.10",
				},
			},
			"dns":     []interface{}{"10.10.1.1", "10.10.1.2"},
			"domain":  "example.com",
			"gateway": "10.10.1.1",
			"mask":    "255.255.255.0",
			"ntp":     "10.10.1.1",
			"subnet":  "10.1.1.0/16",
		},
	},
	"kind":              "NETWORK_INSTANCE_KIND_LOCAL",
	"network_policy_id": "sample policy id",
	"opaque": []interface{}{
		map[string]interface{}{
			"oconfig": "Sample OConfig",
			"type":    "OPAQUE_CONFIG_TYPE_VPN",
		},
	},
	"port": "uplink",
	"port_tags": map[string]interface{}{
		"ptag1": "ptagval1",
		"ptag2": "ptagval2",
	},
	"project_id": "sample-project-id",
	"revision": []interface{}{map[string]interface{}{
		"created_at": "2020-12-15T06:21:24Z",
		"created_by": "sample@example.com",
		"curr":       "1",
		"prev":       "0",
		"updated_at": "2020-12-15T06:21:24Z",
		"updated_by": "user@example.com",
	}},
	"tags": map[string]interface{}{
		"tag1": "val1",
		"tag2": "val2",
	},
	"type": "NETWORK_INSTANCE_DHCP_TYPE_V4",
}

func TestFlattenNetInstConfig(t *testing.T) {
	cases := []struct {
		input       *swagger_models.NetInstConfig
		expected    map[string]interface{}
		description string
	}{
		{
			input:       nil,
			expected:    map[string]interface{}{},
			description: "nil Struct",
		},
		{
			input:       &swagger_models.NetInstConfig{},
			expected:    ouputNetInstEmpty,
			description: "Empty Struct",
		},
		{
			input:       netInstFullCfg(),
			expected:    outputNetInstFullCfg,
			description: "Full Cfg",
		},
	}

	for _, c := range cases {
		out := flattenNetInstConfig(c.input)
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
