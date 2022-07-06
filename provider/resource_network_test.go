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

var rdNetworkEmptyOutput = map[string]interface{}{
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

var rdNetworkFullCfg = map[string]interface{}{
	"name":        "SampleName",
	"id":          "Sample ID",
	"description": "Sample Description",
	"title":       "Sample Title",
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
	"enterprise_default": true,
	"ip": []interface{}{
		map[string]interface{}{
			"dhcp": "NETWORK_DHCP_TYPE_CLIENT",
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
	"kind":       "NETWORK_KIND_V4",
	"project_id": "sample-project-id",
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
		"updated_at": "2020-12-15T06:21:24Z",
		"updated_by": "user@example.com",
	}},
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
					"identity":   "Sample Identity",
					"key_scheme": "NETWORK_WIFIKEY_SCHEME_WPAPSK",
					"priority":   10,
					"secret": []interface{}{
						map[string]interface{}{
							"wifi_passwd": "sample wifi password",
						},
					},
					"wifi_ssid": "Sample Wifi SSID",
				},
			},
		},
	},
}

// efo - Expected Flattened Output
// Computed fields will how up as empty - Kets will be there, but values will
//  be empty
var efoNetworkFullCfg = map[string]interface{}{
	"name":               rdNetworkFullCfg["name"],
	"id":                 rdNetworkFullCfg["id"],
	"description":        rdNetworkFullCfg["description"],
	"title":              rdNetworkFullCfg["title"],
	"dns_list":           rdNetworkFullCfg["dns_list"],
	"enterprise_default": rdNetworkFullCfg["enterprise_default"],
	"ip":                 rdNetworkFullCfg["ip"],
	"kind":               rdNetworkFullCfg["kind"],
	"project_id":         rdNetworkFullCfg["project_id"],
	"proxy":              rdNetworkFullCfg["proxy"],
	"revision":           []interface{}{},
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

var rdNetworkFirstLevelStructKeysAbsent = map[string]interface{}{
	"name":               "SampleName",
	"id":                 "Sample ID",
	"description":        "Sample Description",
	"title":              "Sample Title",
	"enterprise_default": true,
	"kind":               "NETWORK_KIND_V4",
	"project_id":         "sample-project-id",
}

// efo - Expected Flattened Output
// Computed fields will how up as empty - Kets will be there, but values will
//  be empty
var efoNetworkFirstLevelStructKeysAbsent = map[string]interface{}{
	"name":               rdNetworkFirstLevelStructKeysAbsent["name"],
	"id":                 rdNetworkFirstLevelStructKeysAbsent["id"],
	"description":        rdNetworkFirstLevelStructKeysAbsent["description"],
	"title":              rdNetworkFirstLevelStructKeysAbsent["title"],
	"dns_list":           []interface{}{},
	"enterprise_default": rdNetworkFirstLevelStructKeysAbsent["enterprise_default"],
	"ip":                 []interface{}{},
	"kind":               rdNetworkFirstLevelStructKeysAbsent["kind"],
	"project_id":         rdNetworkFirstLevelStructKeysAbsent["project_id"],
	"proxy":              []interface{}{},
	"revision":           []interface{}{},
	"wireless":           []interface{}{},
}

var rdNetworkFirstLevelStructKeysEmpty = map[string]interface{}{
	"name":               "SampleName",
	"id":                 "Sample ID",
	"description":        "Sample Description",
	"dns_list":           []interface{}{},
	"title":              "Sample Title",
	"enterprise_default": true,
	"ip":                 []interface{}{},
	"kind":               "NETWORK_KIND_V4",
	"project_id":         "sample-project-id",
	"revision":           []interface{}{},
	"wireless":           []interface{}{},
}

// efo - Expected Flattened Output
// Computed fields will how up as empty - Kets will be there, but values will
//  be empty
var efoNetworkFirstLevelStructKeysEmpty = map[string]interface{}{
	"name":               rdNetworkFirstLevelStructKeysAbsent["name"],
	"id":                 rdNetworkFirstLevelStructKeysAbsent["id"],
	"description":        rdNetworkFirstLevelStructKeysAbsent["description"],
	"title":              rdNetworkFirstLevelStructKeysAbsent["title"],
	"dns_list":           []interface{}{},
	"enterprise_default": rdNetworkFirstLevelStructKeysAbsent["enterprise_default"],
	"ip":                 []interface{}{},
	"kind":               rdNetworkFirstLevelStructKeysAbsent["kind"],
	"project_id":         rdNetworkFirstLevelStructKeysAbsent["project_id"],
	"proxy":              []interface{}{},
	"revision":           []interface{}{},
	"wireless":           []interface{}{},
}

var rdNetwork2ndLevelStructsEmpty = map[string]interface{}{
	"name":        "SampleName",
	"id":          "Sample ID",
	"description": "Sample Description",
	"title":       "Sample Title",
	"dns_list": []interface{}{
		map[string]interface{}{
			"addrs":    []interface{}{},
			"hostname": "www.example.com",
		},
		map[string]interface{}{
			"addrs":    []interface{}{},
			"hostname": "www.example2.com",
		},
	},
	"enterprise_default": true,
	"ip": []interface{}{
		map[string]interface{}{
			"dhcp":       "NETWORK_DHCP_TYPE_CLIENT",
			"dhcp_range": []interface{}{},
			"dns":        []interface{}{},
			"domain":     "example.com",
			"gateway":    "10.10.1.1",
			"mask":       "255.255.255.0",
			"ntp":        "10.10.1.1",
			"subnet":     "10.1.1.0/16",
		},
	},
	"kind":       "NETWORK_KIND_V4",
	"project_id": "sample-project-id",
	"proxy": []interface{}{
		map[string]interface{}{
			"exceptions":          "exception1",
			"network_proxy":       true,
			"network_proxy_certs": []interface{}{},
			"network_proxy_url":   "www.example.com",
			"pacfile":             "testpacfile",
			"proxy":               []interface{}{},
		},
	},
	"revision": []interface{}{map[string]interface{}{
		"created_at": "2020-12-15T06:21:24Z",
		"created_by": "sample@example.com",
		"curr":       "1",
		"updated_at": "2020-12-15T06:21:24Z",
		"updated_by": "user@example.com",
	}},
	"wireless": []interface{}{
		map[string]interface{}{
			// In reality, we can have either Cellular cfg or wifi cfg depending
			// on type - not both. But for unit test, lets populate both structures.
			"cellular_cfg": []interface{}{},
			"type":         "NETWORK_WIRELESS_TYPE_WIFI",
			"wifi_cfg":     []interface{}{},
		},
	},
}
var efoNetwork2ndLevelStructsEmpty = map[string]interface{}{
	"name":               rdNetworkFullCfg["name"],
	"id":                 rdNetworkFullCfg["id"],
	"description":        rdNetworkFullCfg["description"],
	"title":              rdNetworkFullCfg["title"],
	"dns_list":           rdNetworkFullCfg["dns_list"],
	"enterprise_default": rdNetworkFullCfg["enterprise_default"],
	"ip":                 rdNetworkFullCfg["ip"],
	"kind":               rdNetworkFullCfg["kind"],
	"project_id":         rdNetworkFullCfg["project_id"],
	"proxy":              rdNetworkFullCfg["proxy"],
	"revision":           []interface{}{},
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

// In each test case, call rdXXX to get the appropriate config struct,
//  feed it to flattenXXX, verify output of flattenXXX is same as input
func TestRDNetworkConfig(t *testing.T) {
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
			expectedFlattenedOutput: rdNetworkEmptyOutput,
		},
		{
			input:                   rdNetworkFullCfg,
			description:             "Full Cfg Struct",
			expectError:             false,
			expectedFlattenedOutput: efoNetworkFullCfg,
		},
		{
			input:                   rdNetworkFirstLevelStructKeysAbsent,
			description:             "1st level Struct keys absent",
			expectError:             false,
			expectedFlattenedOutput: efoNetworkFirstLevelStructKeysAbsent,
		},
		{
			input:                   rdNetworkFirstLevelStructKeysEmpty,
			description:             "1st level Struct keys absent",
			expectError:             false,
			expectedFlattenedOutput: efoNetworkFirstLevelStructKeysEmpty,
		},
		/*
			{
				input:                   rdNetwork2ndLevelStructsEmpty,
				description:             "2nd level keys Empty",
				expectError:             false,
				expectedFlattenedOutput: efoNetwork2ndLevelStructsEmpty,
			},
		*/
	}

	for _, c := range cases {
		rd := schema.TestResourceDataRaw(t, zschemas.NetworkSchema, c.input)
		cfg := &swagger_models.NetConfig{}
		name, ok := c.input["name"].(string)
		if !ok {
			name = ""
		}
		cfg.Name = &name
		cfg.ID, ok = c.input["id"].(string)
		err := rdToNetConfig(cfg, rd)
		if err != nil {
			if !c.expectError {
				t.Fatalf("Test Failed: %s\n"+
					"Unexpected error from rdToNetConfig.\n"+
					"Error: %+v\n", c.description, err.Error())
			}
		} else {
			if c.expectError {
				t.Fatalf("Test Failed: %s. Expecting Error, but did not get one",
					c.description)
			}
		}
		out := flattenNetConfig(cfg)
		err = verifyFlattenOutput(zschemas.NetworkSchema, out, c.expectAllSchemaKeys)
		if err != nil {
			t.Fatalf("Test Failed: %s\n Errors in flatten output. Err: %s",
				c.description, err.Error())
		}
		if diff := deep.Equal(out, c.expectedFlattenedOutput); diff != nil {
			t.Fatalf("Test Failed: %s\n"+
				"Error matching Flattened output and input.\n"+
				"Output: %#v\n"+
				"Input : %#v\n"+
				"Diff ( expected vs actual): %#v", c.description, out,
				c.expectedFlattenedOutput, diff)
		}
	}
}
