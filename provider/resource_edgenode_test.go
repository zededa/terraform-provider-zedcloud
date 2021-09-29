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

var rdEdgeNodeEmptyOutput = map[string]interface{}{
	"adminstate":        "",
	"asset_id":          "",
	"client_ip":         "",
	"cluster_id":        "",
	"config_items":      map[string]interface{}{},
	"cpu":               0,
	"description":       "",
	"dev_location":      []interface{}{},
	"eve_image_version": "",
	"id":                "",
	"interface":         []interface{}{},
	"memory":            0,
	"model_id":          "",
	"name":              "",
	"project_id":        "",
	"reset_counter":     0,
	"reset_time":        "",
	"revision":          []interface{}{},
	"serialno":          "",
	"storage":           0,
	"tags":              map[string]interface{}{},
	"thread":            0,
	"title":             "",
	"utype":             "",
}

var rdEdgeNodeFullCfg = map[string]interface{}{
	"name":         "sample-EdgeNode",
	"id":           "SAMPLE-EdgeNode-ID",
	"description":  "SAMPLE EdgeNode DESCRIPTION",
	"title":        "sample-title 2",
	"asset_id":     "Sample Asset ID",
	"client_ip":    "10.10.1.1",
	"cluster_id":   "sample-cluster-id",
	"config_items": map[string]interface{}{},
	"dev_location": []interface{}{
		map[string]interface{}{
			"city":        "sample-city",
			"country":     "sample-country",
			"freeloc":     "sample-freeloc",
			"hostname":    "sample-host",
			"loc":         "sample-loc",
			"org":         "sample-org",
			"postal":      "sample-postal",
			"region":      "sample-region",
			"underlay_ip": "10.10.1.1",
		},
	},
	"eve_image_version": "6.8.2-kvm-amd64",
	"interface": []interface{}{
		map[string]interface{}{
			"cost":       10,
			"intfname":   "intf1",
			"intf_usage": "ADAPTER_USAGE_MANAGEMENT",
			"ipaddr":     "10.10.1.5",
			"macaddr":    "0.1.2",
			"netname":    "sample-net",
			"tags": map[string]interface{}{
				"tag1": "value1",
				"tag2": "value2",
			},
		},
		/* Since interfaces are of type schema.Set, need a custom diff function
		        to compare multiple interfaces
				map[string]interface{}{
					"cost":       1,
					"intfname":   "intf2",
					"intf_usage": "ADAPTER_USAGE_DISABLED",
					"ipaddr":     "10.10.1.6",
					"macaddr":    "0.1.3",
					"netname":    "sample-net2",
					"tags": map[string]interface{}{
						"tag12": "value12",
						"tag22": "value22",
					},
				},
		*/
	},
	"memory":        "",
	"model_id":      "",
	"project_id":    "sample-project",
	"reset_counter": 10,
	"reset_time":    "sample-time",
	"revision": []interface{}{map[string]interface{}{
		"created_at": "2020-12-15T06:21:24Z",
		"created_by": "sample@example.com",
		"curr":       "1",
		"updated_at": "2020-12-15T06:21:24Z",
		"updated_by": "user@example.com",
	}},
	"serialno": "sample-serial-no",
	"storage":  100,
	"tags": map[string]interface{}{
		"edtag1": "edval1",
		"edtag2": "edval2",
	},
	"thread": 2,
	"utype":  "AMD64",
}

// efo - Expected Flattened Output
var efoEdgeNodeFullCfg = map[string]interface{}{
	"name":              rdEdgeNodeFullCfg["name"],
	"id":                rdEdgeNodeFullCfg["id"],
	"description":       rdEdgeNodeFullCfg["description"],
	"title":             rdEdgeNodeFullCfg["title"],
	"adminstate":        "",
	"asset_id":          rdEdgeNodeFullCfg["asset_id"],
	"client_ip":         rdEdgeNodeFullCfg["client_ip"],
	"cluster_id":        rdEdgeNodeFullCfg["cluster_id"],
	"config_items":      rdEdgeNodeFullCfg["config_items"],
	"cpu":               0,
	"dev_location":      rdEdgeNodeFullCfg["dev_location"],
	"eve_image_version": rdEdgeNodeFullCfg["eve_image_version"],
	"interface":         rdEdgeNodeFullCfg["interface"],
	"memory":            0,
	"model_id":          "",
	"project_id":        "",
	"reset_counter":     0,
	"reset_time":        "",
	"revision":          []interface{}{},
	"serialno":          "",
	"storage":           0,
	"tags":              rdEdgeNodeFullCfg["tags"],
	"thread":            0,
	"utype":             "",
}

// In each test case, call rdXXX to get the appropriate config struct,
//  feed it to flattenXXX, verify output of flattenXXX is same as input
func TestRDEdgeNodeConfig(t *testing.T) {
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
			expectError:             true,
			expectedFlattenedOutput: rdEdgeNodeEmptyOutput,
		},
		{
			input:                   rdEdgeNodeFullCfg,
			description:             "Full Cfg Struct",
			expectError:             false,
			expectedFlattenedOutput: efoEdgeNodeFullCfg,
		},
	}

	for _, c := range cases {
		rd := schema.TestResourceDataRaw(t, zschemas.EdgeNodeSchema, c.input)
		cfg := &swagger_models.DeviceConfig{}
		name, ok := c.input["name"].(string)
		if !ok {
			name = ""
		}
		cfg.Name = &name
		cfg.ID, ok = c.input["id"].(string)
		cfg.BaseImage = cfgBaseosForEveVersionStr(rdEntryStr(rd, "eve_image_version"))
		err := rdDeviceConfig(cfg, rd)
		if err != nil {
			if !c.expectError {
				t.Fatalf("Test Failed: %s\n"+
					"Unexpected error from updateEdgeNodeCfgFromResourceData.\n"+
					"Error: %+v\n", c.description, err.Error())
			}
		} else {
			if c.expectError {
				t.Fatalf("Test Failed: %s. Expecting Error, but did not get one",
					c.description)
			}
		}
		out := flattenDeviceConfig(cfg)
		err = verifyFlattenOutput(zschemas.EdgeNodeSchema, out, c.expectAllSchemaKeys)
		if err != nil {
			t.Fatalf("Test Failed: %s\n Errors in flatten output. Err: %s",
				c.description, err.Error())
		}
		if diff := deep.Equal(out, c.expectedFlattenedOutput); diff != nil {
			t.Fatalf("Test Failed: %s\n"+
				"Error matching Flattened output and input.\n"+
				"Output: %#v\n"+
				"Input : %#v\n"+
				"Diff: %#v", c.description, out, c.expectedFlattenedOutput, diff)
		}
	}
}
