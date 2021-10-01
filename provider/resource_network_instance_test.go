// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"github.com/go-test/deep"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	"github.com/zededa/zedcloud-api/swagger_models"
	"testing"
)

var rdNetworkInstanceEmptyOutput = map[string]interface{}{
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

var rdNetworkInstanceFullCfg = map[string]interface{}{
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
		"updated_at": "2020-12-15T06:21:24Z",
		"updated_by": "user@example.com",
	}},
	"tags": map[string]interface{}{
		"tag1": "val1",
		"tag2": "val2",
	},
	"type": "NETWORK_INSTANCE_DHCP_TYPE_V4",
}

// efo - Expected Flattened Output
// Computed fields will how up as empty - Kets will be there, but values will
//  be empty
var efoNetworkInstanceFullCfg = map[string]interface{}{
	"name":              rdNetworkInstanceFullCfg["name"],
	"id":                rdNetworkInstanceFullCfg["id"],
	"description":       rdNetworkInstanceFullCfg["description"],
	"title":             rdNetworkInstanceFullCfg["title"],
	"cluster_id":        rdNetworkInstanceFullCfg["cluster_id"],
	"device_default":    rdNetworkInstanceFullCfg["device_default"],
	"device_id":         rdNetworkInstanceFullCfg["device_id"],
	"dns_list":          rdNetworkInstanceFullCfg["dns_list"],
	"ip":                rdNetworkInstanceFullCfg["ip"],
	"kind":              rdNetworkInstanceFullCfg["kind"],
	"network_policy_id": rdNetworkInstanceFullCfg["network_policy_id"],
	"opaque":            rdNetworkInstanceFullCfg["opaque"],
	"port":              rdNetworkInstanceFullCfg["port"],
	"port_tags":         rdNetworkInstanceFullCfg["port_tags"],
	"project_id":        rdNetworkInstanceFullCfg["project_id"],
	"revision":          []interface{}{},
	"tags":              rdNetworkInstanceFullCfg["tags"],
	"type":              rdNetworkInstanceFullCfg["type"],
}

// In each test case, call rdXXX to get the appropriate config struct,
//  feed it to flattenXXX, verify output of flattenXXX is same as input
func TestRDNetworkInstanceConfig(t *testing.T) {
	cases := []struct {
		input                   map[string]interface{}
		description             string
		expectError             bool
		expectedFlattenedOutput map[string]interface{}
		expectAllSchemaKeys     bool
	}{
		{
			input:                   map[string]interface{}{},
			description:             "Empty RD",
			expectError:             false,
			expectedFlattenedOutput: rdNetworkInstanceEmptyOutput,
		},
		{
			input:                   rdNetworkInstanceFullCfg,
			description:             "Full Cfg Struct",
			expectError:             false,
			expectedFlattenedOutput: efoNetworkInstanceFullCfg,
		},
	}

	for _, c := range cases {
		rd := schema.TestResourceDataRaw(t, zschemas.NetworkInstanceSchema, c.input)
		cfg := &swagger_models.NetInstConfig{}
		name, ok := c.input["name"].(string)
		if !ok {
			name = ""
		}
		cfg.Name = &name
		cfg.ID, ok = c.input["id"].(string)
		cfg.ClusterID, ok = c.input["cluster_id"].(string)
		err := rdUpdateNetInstConfig(rd, cfg)
		if err != nil {
			if !c.expectError {
				t.Fatalf("Test Failed: %s\n"+
					"Unexpected error from rdUpdateNetInstConfig.\n"+
					"Error: %+v\n", c.description, err.Error())
			}
		} else {
			if c.expectError {
				t.Fatalf("Test Failed: %s. Expecting Error, but did not get one",
					c.description)
			}
		}
		out := flattenNetInstConfig(cfg, false)
		err = verifyFlattenOutput(zschemas.NetworkInstanceSchema, out, c.expectAllSchemaKeys)
		if err != nil {
			t.Fatalf("Test Failed: %s\n Errors in flatten output. Err: %s",
				c.description, err.Error())
		}
		if diff := deep.Equal(out, c.expectedFlattenedOutput); diff != nil {
			t.Fatalf("Test Failed: %s\n"+
				"Error matching Flattened output and input.\n"+
				"Output: %#v\n"+
				"Expected Output : %#v\n"+
				"Diff Actual vs Expected: %#v", c.description, out, c.expectedFlattenedOutput, diff)
		}
	}
}
