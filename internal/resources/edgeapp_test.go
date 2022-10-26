// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package resources

import (
	"testing"

	"github.com/go-test/deep"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	"github.com/zededa/zedcloud-api/swagger_models"
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
	"name":          "sample-name",
	"id":            "SAMPLE-ID",
	"description":   "SAMPLE DESCRIPTION",
	"title":         "Sample Title",
	"cpus":          0,
	"drives":        0,
	"manifest_file": "../samples/TFTest-ubuntu-xenial-16.04.json",
	"memory":        0,
	"networks":      0,
	"origin_type":   "ORIGIN_LOCAL",
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

// In each test case, call rdXXX to get the appropriate config struct,
//
//	feed it to flattenXXX, verify output of flattenXXX is same as input
func TestRDEdgeAppConfig(t *testing.T) {
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
			expectedFlattenedOutput: rdEdgeAppEmptyOutput,
		},
		{
			input:                   rdEdgeAppMinCfg,
			description:             "Minimum Config RD",
			expectError:             false,
			expectedFlattenedOutput: rdEdgeAppEmptyOutput,
		},
		{
			input:                   rdEdgeAppFullCfg,
			description:             "Full Cfg Struct",
			expectError:             false,
			expectedFlattenedOutput: efoEdgeAppFullCfg,
		},
	}

	for _, c := range cases {
		rd := schema.TestResourceDataRaw(t, zschemas.EdgeAppSchema, c.input)
		cfg := &swagger_models.App{}
		name, ok := c.input["name"].(string)
		if !ok {
			name = ""
		}
		cfg.Name = &name
		cfg.ID, ok = c.input["id"].(string)
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
		out := flattenEdgeAppConfig(cfg, false)
		err = verifyFlattenOutput(zschemas.EdgeAppSchema, out, c.expectAllSchemaKeys)
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
