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

var rdEdgeAppEmptyOutput = map[string]interface{}{
	"name":                 "",
	"id":                   "",
	"title":                "",
	"description":          "",
	"cpus":                 0,
	"drives":               0,
	"memory":               0,
	"networks":             0,
	"origin_type":          "",
	"parent_detail":        []interface{}{},
	"revision":             []interface{}{},
	"storage":              0,
	"user_defined_version": "",
}
var rdEdgeAppMinCfg = map[string]interface{}{
	"name":                 "",
	"id":                   "",
	"title":                "",
	"description":          "",
	"cpus":                 0,
	"drives":               0,
	"manifest_file":        "../samples/TFTest-ubuntu-xenial-16.04.json",
	"memory":               0,
	"networks":             0,
	"origin_type":          "",
	"parent_detail":        []interface{}{},
	"revision":             []interface{}{},
	"storage":              0,
	"user_defined_version": "",
}

var rdEdgeAppFullCfg = map[string]interface{}{
	"name":        "sample-name",
	"id":          "SAMPLE-ID",
	"description": "SAMPLE DESCRIPTION",
	"title":       "Sample Title",
	"cpus":        0,
	"drives":      0,
	"memory":      0,
	"networks":    0,
	"origin_type": "ORIGIN_LOCAL",
	"parentDetail": []interface{}{
		map[string]interface{}{
			"idOfParentObject":      "",
			"referenceExists":       false,
			"updateAvailable":       false,
			"versionOfParentObject": 0,
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
	"storage":              0,
	"user_defined_version": "",
}

var rdEdgeAppManifest = map[string]interface{}{
	"apptype": "APP_TYPE_VM",
	"configuration": []interface{}{
		map[string]interface{}{
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
		},
	},
	"container_detail": []interface{}{
		map[string]interface{}{
			"container_create_option": "Create Option String",
		},
	},
	"deployment_type": "DEPLOYMENT_TYPE_STAND_ALONE",
	"desc_detail": []interface{}{
		map[string]interface{}{
			"agreement_list": map[string]interface{}{
				"agreement1": "Agreement1 Value",
				"agreement2": "Agreement2 Value",
			},
			"app_category": "APP_CATEGORY_OPERATING_SYSTEM",
			"category":     "Type of Edge App",
			"license_list": map[string]interface{}{
				"License1": "License1 Value",
				"License2": "License2 Value",
			},
			"logo": map[string]interface{}{
				"logo1": "Logo Value 1",
				"logo2": "Logo Value 2",
			},
			"os":      "Linux",
			"support": "Sample Support",
		},
	},
	"description":  "App Description",
	"display_name": "Sample Display Name",
	"enablevnc":    true,
	"image": []interface{}{
		map[string]interface{}{
			"cleartext":   true,
			"drvtype":     "CDROM",
			"ignorepurge": true,
			"imagename":   "Sample Image Name",
			"maxsize":     "12345",
			"mountpath":   "Sample Mount Path",
			"param": []interface{}{
				map[string]interface{}{
					"name":  "Sample Param Name",
					"value": "Sample Param Value",
				},
			},
			"preserve":    true,
			"readonly":    true,
			"target":      true,
			"volumelabel": "Sample Vol label",
		},
	},
	"interface": []interface{}{
		map[string]interface{}{
			"acl": []interface{}{
				map[string]interface{}{
					"action": []interface{}{
						map[string]interface{}{
							"drop":  true,
							"limit": true,
							"limit_param": []interface{}{
								map[string]interface{}{
									"limitburst": 10,
									"limitrate":  20,
									"limitunit":  30,
								},
							},
							"portmap": true,
							"mapparam": []interface{}{
								map[string]interface{}{
									"port": 10,
								},
							},
						},
					},
					"match": []interface{}{
						map[string]interface{}{
							"type":  "Match Type",
							"value": "Match Value",
						},
					},
					"name": "Sample Name",
				},
			},
			"directattach": true,
			"name":         "Intf name",
			"optional":     true,
			"privateip":    true,
			"type":         "Eth",
		},
	},
	"module": []interface{}{
		map[string]interface{}{
			"environment": map[string]interface{}{
				"env1": "env1 value",
				"env2": "env2 value",
			},
			"module_type": "MODULE_TYPE_SYSTEM_DEFINED",
			"routes": map[string]interface{}{
				"route1": "route1 value",
				"route2": "route2 value",
			},
			"twin_detail": "Sample Twin Detail",
		},
	},
	"name": "Sample App Name",
	"owner": []interface{}{
		map[string]interface{}{
			"company": "Sample Company Name",
			"email":   "Sample email",
			"user":    "Sample User",
			"website": "Sample website",
		},
	},
	"resource": []interface{}{
		map[string]interface{}{
			"name":  "Sample Resource Name",
			"value": "Sample Resource Value",
		},
	},
	"vmmode": "Sample VM mode",
}

// efo - Expected Flattened Output
var efoEdgeAppFullCfg = map[string]interface{}{
	"name":                 rdEdgeAppFullCfg["name"],
	"id":                   rdEdgeAppFullCfg["id"],
	"title":                rdEdgeAppFullCfg["title"],
	"description":          rdEdgeAppFullCfg["description"],
	"cpus":                 rdEdgeAppFullCfg["cpus"],
	"drives":               rdEdgeAppFullCfg["drives"],
	"memory":               rdEdgeAppFullCfg["memory"],
	"networks":             rdEdgeAppFullCfg["networks"],
	"origin_type":          "",              // Computed field
	"parent_detail":        []interface{}{}, // Computed field
	"revision":             []interface{}{}, // Computed field
	"storage":              rdEdgeAppFullCfg["storage"],
	"user_defined_version": rdEdgeAppFullCfg["user_defined_version"],
}

func rdEdgeAppFullCfgInput(manifestFile bool) map[string]interface{} {
	if manifestFile {
		delete(rdEdgeAppFullCfg, "manifest")
		rdEdgeAppFullCfg["manifest_file"] = "../samples/TFTest-ubuntu-xenial-16.04.json"
		if _, ok := rdEdgeAppFullCfg["manifest"]; ok {
			panic("Unexpected Manfest key")
		}
		if _, ok := rdEdgeAppFullCfg["manifest_file"]; !ok {
			panic("manifest_file not found")
		}
	} else {
		delete(rdEdgeAppFullCfg, "manifest_file")
		rdEdgeAppFullCfg["manifest"] = []interface{}{rdEdgeAppManifest}
		if _, ok := rdEdgeAppFullCfg["manifest"]; !ok {
			panic("Manfest key Not Found")
		}
	}
	return rdEdgeAppFullCfg
}

func efoEdgeAppFullCfgOutput(manifestFile bool) map[string]interface{} {
	if manifestFile {
		if _, ok := efoEdgeAppFullCfg["manifest"]; ok {
			panic("manifest present in efoEdgeAppFullCfg")
		}
		delete(efoEdgeAppFullCfg, "manifest")
		if _, ok := efoEdgeAppFullCfg["manifest"]; ok {
			panic("Unexpected Manfest key")
		}
	} else {
		efoEdgeAppFullCfg["manifest"] = []interface{}{rdEdgeAppManifest}
		if _, ok := efoEdgeAppFullCfg["manifest"]; !ok {
			panic("Manfest key Not Found")
		}
	}
	return efoEdgeAppFullCfg
}

// In each test case, call rdXXX to get the appropriate config struct,
//  feed it to flattenXXX, verify output of flattenXXX is same as input
func TestRDEdgeAppConfig(t *testing.T) {
	cases := []struct {
		input                   map[string]interface{}
		inputFunc               func(bool) map[string]interface{}
		description             string
		expectError             bool
		efoFunc                 func(bool) map[string]interface{}
		expectedFlattenedOutput map[string]interface{}
		expectAllSchemaKeys     bool
		manifestFile            bool
	}{
		{
			input:                   map[string]interface{}{},
			description:             "Empty RD",
			expectError:             true,
			expectedFlattenedOutput: rdEdgeAppEmptyOutput,
		},
		{
			input:                   rdEdgeAppMinCfg,
			description:             "Minimum Config RD",
			expectError:             false,
			expectedFlattenedOutput: rdEdgeAppEmptyOutput,
		},
		{
			inputFunc:    rdEdgeAppFullCfgInput,
			description:  "Full Cfg Struct with Manifest File",
			expectError:  false,
			efoFunc:      efoEdgeAppFullCfgOutput,
			manifestFile: true,
		},
		{
			inputFunc:   rdEdgeAppFullCfgInput,
			description: "Full Cfg Struct with Manifest",
			expectError: false,
			efoFunc:     efoEdgeAppFullCfgOutput,
		},
	}

	for _, c := range cases {
		input := c.input
		if c.inputFunc != nil {
			input = c.inputFunc(c.manifestFile)
		}
		efo := c.expectedFlattenedOutput
		if c.efoFunc != nil {
			efo = c.efoFunc(c.manifestFile)
		}
		rd := schema.TestResourceDataRaw(t, zschemas.EdgeAppSchema, input)
		cfg := &swagger_models.App{}
		name, ok := input["name"].(string)
		if !ok {
			name = ""
		}
		cfg.Name = &name
		cfg.ID, ok = input["id"].(string)
		err := rdUpdateAppCfg(cfg, rd)
		if err != nil {
			if !c.expectError {
				t.Fatalf("Test Failed: %s\n"+
					"Unexpected error from updateEdgeAppCfgFromResourceData.\n"+
					"Error: %+v\n", c.description, err.Error())
			}
		} else {
			if c.expectError {
				t.Fatalf("Test Failed: %s. Expecting Error, but did not get one",
					c.description)
			}
		}
		out := flattenEdgeAppConfig(cfg, false, false)
		err = verifyFlattenOutput(zschemas.EdgeAppSchema, out, c.expectAllSchemaKeys)
		if err != nil {
			t.Fatalf("Test Failed: %s\n Errors in flatten output. Err: %s",
				c.description, err.Error())
		}
		if diff := deep.Equal(out, efo); diff != nil {
			t.Fatalf("Test Failed: %s\n"+
				"Error matching Flattened output and input.\n"+
				"Output: %#v\n"+
				"Expected Output : %#v\n"+
				"Diff: %#v", c.description, out, efo, diff)
		}
	}
}
