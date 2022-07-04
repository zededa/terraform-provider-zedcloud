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

var rdVolumeInstanceEmptyOutput = map[string]interface{}{
	"accessmode":      "",
	"cleartext":       false,
	"content_tree_id": "",
	"description":     "",
	"device_id":       "",
	"id":              "",
	"image":           "",
	"implicit":        false,
	"label":           "",
	"multiattach":     false,
	"name":            "",
	"project_id":      "",
	"purge":           []interface{}{},
	"revision":        []interface{}{},
	"size_bytes":      "",
	"title":           "",
	"type":            "",
}

var rdVolumeInstanceFullCfg = map[string]interface{}{
	"name":            "SampleName",
	"id":              "Sample ID",
	"description":     "Sample Description",
	"title":           "Sample Title",
	"accessmode":      "VOLUME_INSTANCE_ACCESS_MODE_READONLY",
	"cleartext":       true,
	"content_tree_id": "sample content tree id",
	"cluster_id":      "sample cluster id",
	"device_id":       "sample device ID",
	"image":           "sample image",
	"label":           "sample label",
	"multiattach":     true,
	"project_id":      "sample-project-id",
	"purge":           []interface{}{},
	"revision": []interface{}{map[string]interface{}{
		"created_at": "2020-12-15T06:21:24Z",
		"created_by": "sample@example.com",
		"curr":       "1",
		"updated_at": "2020-12-15T06:21:24Z",
		"updated_by": "user@example.com",
	}},
	"size_bytes": "100",
	"type":       "VOLUME_INSTANCE_TYPE_BLOCKSTORAGE",
}

// efo - Expected Flattened Output
// Computed fields will how up as empty - Kets will be there, but values will
//  be empty
var efoVolumeInstanceFullCfg = map[string]interface{}{
	"accessmode":      rdVolumeInstanceFullCfg["accessmode"],
	"cleartext":       rdVolumeInstanceFullCfg["cleartext"],
	"content_tree_id": rdVolumeInstanceFullCfg["content_tree_id"],
	"description":     rdVolumeInstanceFullCfg["description"],
	"device_id":       rdVolumeInstanceFullCfg["device_id"],
	"id":              rdVolumeInstanceFullCfg["id"],
	"image":           rdVolumeInstanceFullCfg["image"],
	"implicit":        false,
	"label":           rdVolumeInstanceFullCfg["label"],
	"multiattach":     rdVolumeInstanceFullCfg["multiattach"],
	"name":            rdVolumeInstanceFullCfg["name"],
	"project_id":      rdVolumeInstanceFullCfg["project_id"],
	"purge":           []interface{}{},
	"revision":        []interface{}{},
	"size_bytes":      rdVolumeInstanceFullCfg["size_bytes"],
	"title":           rdVolumeInstanceFullCfg["title"],
	"type":            rdVolumeInstanceFullCfg["type"],
}

// In each test case, call rdXXX to get the appropriate config struct,
//  feed it to flattenXXX, verify output of flattenXXX is same as input
func TestRDVolumeInstanceConfig(t *testing.T) {
	cases := []struct {
		description             string
		input                   map[string]interface{}
		expectedFlattenedOutput map[string]interface{}
		expectError             bool
		expectAllSchemaKeys     bool
	}{
		{
			description:             "Empty RD",
			input:                   map[string]interface{}{},
			expectedFlattenedOutput: rdVolumeInstanceEmptyOutput,
			expectError:             false,
		},
		{
			input:                   rdVolumeInstanceFullCfg,
			description:             "Full Cfg Struct",
			expectError:             false,
			expectedFlattenedOutput: efoVolumeInstanceFullCfg,
			expectAllSchemaKeys:     true,
		},
	}

	for _, c := range cases {
		rd := schema.TestResourceDataRaw(t, zschemas.VolumeInstanceSchema, c.input)
		cfg := &swagger_models.VolInstConfig{}
		name, ok := c.input["name"].(string)
		if !ok {
			name = ""
		}
		cfg.Name = name
		cfg.ID = rdEntryStr(c.input, "id")
		cfg.ProjectID = rdEntryStr(c.input, "project_id")
		err := rdUpdateVolumeInstanceCfg(cfg, rd)
		if err != nil {
			if !c.expectError {
				t.Fatalf("Test Failed: %s\n"+
					"Unexpected error from rdUpdateVolumeInstanceCfg.\n"+
					"Error: %+v\n", c.description, err.Error())
			}
		} else {
			if c.expectError {
				t.Fatalf("Test Failed: %s. Expecting Error, but did not get one",
					c.description)
			}
		}
		out := flattenVolInstConfig(cfg)
		err = verifyFlattenOutput(zschemas.VolumeInstanceSchema, out, c.expectAllSchemaKeys)
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
