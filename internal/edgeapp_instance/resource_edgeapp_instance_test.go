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

var rdAppInstEmptyOutput = map[string]interface{}{
	"activate":              false,
	"app_id":                "",
	"app_policy_id":         "",
	"app_type":              "",
	"bundleversion":         "",
	"cluster_id":            "",
	"collect_stats_ip_addr": "",
	"description":           "",
	"device_id":             "",
	"drive":                 []interface{}{},
	"id":                    "",
	"interface":             []interface{}{},
	"is_secret_updated":     false,
	"logs":                  []interface{}{},
	"name":                  "",
	"project_id":            "",
	"purge":                 []interface{}{},
	"refresh":               []interface{}{},
	"remote_console":        false,
	"restart":               []interface{}{},
	"revision":              []interface{}{},
	"tags":                  map[string]interface{}{},
	"title":                 "",
	"user_defined_version":  "",
	"vminfo":                []interface{}{},
}

var rdAppInstFullCfg = map[string]interface{}{
	"name":                  "sample-name",
	"id":                    "SAMPLE-ID",
	"description":           "SAMPLE DESCRIPTION",
	"title":                 "sample-title",
	"activate":              true,
	"app_id":                "sample-app-id",
	"app_policy_id":         "sample-app-policy-id",
	"app_type":              "APP_TYPE_VM",
	"cluster_id":            "Sample Cluster ID",
	"collect_stats_ip_addr": "10.10.1.1",
	"crypto_key":            "Sample Crypto Ckey",
	"custom_config": []interface{}{
		map[string]interface{}{
			"add":                  true,
			"allow_storage_resize": true,
			"field_delimiter":      "sample delimiter",
			"name":                 "sample name",
			"override":             true,
			"template":             "sample template",
			"variable_group": []interface{}{
				map[string]interface{}{
					"condition": []interface{}{
						map[string]interface{}{
							"name":     "sample name condition",
							"operator": "Sample Operator",
							"value":    "Sample Value",
						},
					},
					"name":     "sample name10",
					"required": true,
					"variable": []interface{}{
						map[string]interface{}{
							"default":    "Sample Default Value",
							"encode":     "FILE_ENCODING_BASE64",
							"format":     "VARIABLE_FORMAT_NUMBER",
							"label":      "sample label",
							"max_length": "45",
							"name":       "variable name",
							"option": []interface{}{
								map[string]interface{}{
									"label": "l1",
									"value": "v1",
								},
							},
							"required": true,
							"value":    "sample value",
						},
					},
				},
			},
		},
	},
	"device_id": "sample-device-id",
	"drive":     []interface{}{},
	"encrypted_secrets": map[string]interface{}{
		"secret1": "secret-value1",
		"secret2": "secret-value2",
	},
	"interface": []interface{}{
		map[string]interface{}{
			"access_vlan_id": 2,
			"acl": []interface{}{
				map[string]interface{}{
					"action": []interface{}{
						map[string]interface{}{
							"drop":       false,
							"limit":      true,
							"limitburst": 200,
							"limitrate":  100,
							"limitunit":  "sample-unit",
							"mapparams": []interface{}{
								map[string]interface{}{
									"port": 20,
								},
							},
							"portmap": true,
						},
					},
					"id": 10,
					"match": []interface{}{
						map[string]interface{}{
							"type":  "proto",
							"value": "ip",
						},
					},
					"name": "ace1",
				},
				/*
					map[string]interface{}{
						"action": []interface{}{
							map[string]interface{}{
								"drop": true,
							},
						},
						"id": 20,
						"match": []interface{}{
							map[string]interface{}{
								"type":  "proto",
								"value": "tcp",
							},
						},
						"name": "ace2",
					},
				*/
			},

			"default_net_instance": false,
			"directattach":         false,
			"intfname":             "eth1",
			"io": []interface{}{
				map[string]interface{}{
					"name": "eth1",
					"tags": map[string]interface{}{
						"phytag1": "phytag-val1",
						"phytag2": "phytag-val2",
						"phytag3": "phytag-val3",
					},
					"type": "IO_TYPE_ETH",
				},
			},
			"ipaddr":      "10.10.1.1",
			"macaddr":     "a.b.c",
			"netinstname": "sample-netinst",
			"netinsttag": map[string]interface{}{
				"nitag1": "nitag-val1",
				"nitag2": "nitag-val2",
				"nitag3": "nitag-val3",
			},
		},
		/*
			map[string]interface{}{
				"intfname": "usb2",
				"io": []interface{}{
					map[string]interface{}{
						"name": "usb2",
						"type": "IO_TYPE_USB",
					},
				},
			},
		*/
	},
	"is_secret_updated": false,
	"logs": []interface{}{
		map[string]interface{}{
			"access": false,
		},
	},
	"project_id": "sample-project-id",
	"purge": []interface{}{
		map[string]interface{}{
			"counter":  10,
			"ops_time": "Sample Time",
		},
	},
	"refresh": []interface{}{
		map[string]interface{}{
			"counter":  20,
			"ops_time": "Sample Time1",
		},
	},
	"remote_console": false,
	"restart": []interface{}{
		map[string]interface{}{
			"counter":  30,
			"ops_time": "Sample Time3",
		},
	},
	"revision": []interface{}{
		map[string]interface{}{
			"created_at": "2020-12-15T06:21:24Z",
			"created_by": "sample@example.com",
			"curr":       "1",
			"updated_at": "2020-12-15T06:21:24Z",
			"updated_by": "user@example.com",
		},
	},
	"tags": map[string]interface{}{
		"tag1": "value1",
		"tag2": "value2",
	},
	"user_defined_version": "10.1.1",
	"vminfo": []interface{}{
		map[string]interface{}{
			"cpus":        3,
			"memory":      100,
			"mode":        "HV_FML",
			"vnc":         true,
			"vnc_display": 1,
		},
	},
}

// efo - Expected Flattened Output
var efoAppInstFullCfg = map[string]interface{}{
	"name":                  rdAppInstFullCfg["name"],
	"id":                    rdAppInstFullCfg["id"],
	"description":           rdAppInstFullCfg["description"],
	"title":                 rdAppInstFullCfg["title"],
	"activate":              rdAppInstFullCfg["activate"],
	"app_id":                rdAppInstFullCfg["app_id"],
	"app_policy_id":         rdAppInstFullCfg["app_policy_id"],
	"app_type":              rdAppInstFullCfg["app_type"],
	"bundleversion":         "", // Computed
	"cluster_id":            rdAppInstFullCfg["cluster_id"],
	"collect_stats_ip_addr": rdAppInstFullCfg["collect_stats_ip_addr"],
	"device_id":             rdAppInstFullCfg["device_id"],
	"drive":                 rdAppInstFullCfg["drive"],
	"interface": []interface{}{
		map[string]interface{}{
			"access_vlan_id": 2,
			// Acls Can't be configured - so read only. CfgUpdate will not have them.
			"acl":                  []interface{}{},
			"default_net_instance": false,
			"directattach":         false,
			"intfname":             "eth1",
			"io": []interface{}{
				map[string]interface{}{
					"name": "eth1",
					"tags": map[string]interface{}{
						"phytag1": "phytag-val1",
						"phytag2": "phytag-val2",
						"phytag3": "phytag-val3",
					},
					"type": "IO_TYPE_ETH",
				},
			},
			"ipaddr":      "10.10.1.1",
			"macaddr":     "a.b.c",
			"netinstname": "sample-netinst",
			"netinsttag": map[string]interface{}{
				"nitag1": "nitag-val1",
				"nitag2": "nitag-val2",
				"nitag3": "nitag-val3",
			},
		},
	},
	"is_secret_updated":    false,
	"logs":                 rdAppInstFullCfg["logs"],
	"project_id":           rdAppInstFullCfg["project_id"],
	"purge":                []interface{}{}, // Computed
	"refresh":              []interface{}{}, // Computed
	"remote_console":       rdAppInstFullCfg["remote_console"],
	"restart":              []interface{}{}, // Computed
	"revision":             []interface{}{}, // Computed
	"tags":                 rdAppInstFullCfg["tags"],
	"user_defined_version": rdAppInstFullCfg["user_defined_version"],
	// vminfo is not set in the config.
	"vminfo": []interface{}{},
}

// In each test case, call rdXXX to get the appropriate config struct,
//  feed it to flattenXXX, verify output of flattenXXX is same as input
func TestRDAppInstConfig(t *testing.T) {
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
			expectedFlattenedOutput: rdAppInstEmptyOutput,
		},
		{
			input:                   rdAppInstFullCfg,
			description:             "Full Cfg Struct",
			expectError:             false,
			expectedFlattenedOutput: efoAppInstFullCfg,
		},
	}

	for _, c := range cases {
		rd := schema.TestResourceDataRaw(t, zschemas.AppInstSchema, c.input)
		cfg := &swagger_models.AppInstance{}
		name, ok := c.input["name"].(string)
		if !ok {
			name = ""
		}
		cfg.Name = &name
		cfg.ID, ok = c.input["id"].(string)
		err := rdUpdateAppInstCfg(cfg, rd)
		if err != nil {
			if !c.expectError {
				t.Fatalf("Test Failed: %s\n"+
					"Unexpected error from updateAppInstCfgFromResourceData.\n"+
					"Error: %+v\n", c.description, err.Error())
			}
		} else {
			if c.expectError {
				t.Fatalf("Test Failed: %s. Expecting Error, but did not get one",
					c.description)
			}
		}
		out := flattenAppInstance(cfg, true)
		err = verifyFlattenOutput(zschemas.AppInstSchema, out, c.expectAllSchemaKeys)
		if err != nil {
			t.Fatalf("Test Failed: %s\n Errors in flatten output. Err: %s\nout: %++v",
				c.description, err.Error(), out)
		}
		if diff := deep.Equal(out, c.expectedFlattenedOutput); diff != nil {
			t.Fatalf("Test Failed: %s\n"+
				"Error matching Flattened output and Expected output.\n"+
				"Output: %#v\n"+
				"Expected Output : %#v\n"+
				"Diff Actual vs. Expected: %#v", c.description, out, c.expectedFlattenedOutput, diff)
		}
	}
}
