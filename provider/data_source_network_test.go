// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"github.com/go-test/deep"
	"github.com/zededa/zedcloud-api/swagger_models"
	"reflect"
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

var outputFullCfg = map[string]interface{}{
	"description": "Test Full Cfg",
	"dns_list": []interface{}{
		map[string]interface{}{
			"addrs": []string{
				"10.10.1.1",
				"10.10.2.2",
			},
			"hostname": "test Host",
		},
		map[string]interface{}{
			"addrs": []string{
				"10.30.1.1",
				"10.4.22.2",
			},
			"hostname": "test Host2",
		},
	},
	"enterprise_default": true,
	"id":                 "TEST ID",
	"ip": []interface{}{map[string]interface{}{
		"dhcp": "dhcp field",
		"dhcp_range": map[string]interface{}{
			"end":   "10.10.1.255",
			"start": "10.10.1.0",
		},
		"dns": []interface{}{
			"dns1",
			"dns2",
		},
		"domain":  "test domain",
		"gateway": "test gateway",
		"mask":    "test mask",
		"ntp":     "test ntp",
		"subnet":  "255.255.255.0",
	}},
	"kind":       "V4",
	"name":       "Test Full Cfg",
	"project_id": "TEST PROJECT ID",
	"proxy":      []interface{}{map[string]interface{}{}},
	"revision":   []interface{}{map[string]interface{}{}},
	"title":      "TEST Title",
	"wireless":   []interface{}{map[string]interface{}{}},
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
		/*
		   {
		       input: &swagger_models.NetConfig{},
		       expected: outputFullCfg,
		       description: "Full CfgEmpty Struct",
		   },
		*/
	}

	for _, c := range cases {
		out := flattenNetConfig(c.input)
		if diff := deep.Equal(out, c.expected); diff != nil {
			t.Fatalf("Error matching output and expected.\nOutput: %#v\nExpected: %#v\n"+
				"Diff: %#v", out, c.expected, diff)
		}
		if !reflect.DeepEqual(out, c.expected) {
			t.Fatalf("Error matching output and expected: %#v vs %#v", out, c.expected)
		}
	}
}
