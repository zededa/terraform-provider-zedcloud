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

var rdImageEmptyOutput = map[string]interface{}{
	"datastore_id":     "",
	"description":      "",
	"id":               "",
	"image_arch":       "",
	"image_error":      "",
	"image_format":     "",
	"image_local":      "",
	"image_rel_url":    "",
	"image_sha_256":    "",
	"image_size_bytes": "",
	"image_status":     "",
	"image_type":       "",
	"image_version":    "",
	"name":             "",
	"origin_type":      "",
	"revision":         []interface{}{},
	"title":            "",
}

var rdImageFullCfg = map[string]interface{}{
	"datastore_id":     "SAMPLE-DS-ID",
	"description":      "SAMPLE IMAGE DESCRIPTION",
	"id":               "SAMPLE-IMAGE-ID",
	"image_arch":       "AMD64",
	"image_error":      "SUCCESS",
	"image_format":     "QCOW2",
	"image_local":      "sampleLocalLoc",
	"image_rel_url":    "sampleRelUrl",
	"image_sha_256":    "sampleSHA",
	"image_size_bytes": "123435",
	"image_status":     "READY",
	"image_type":       "IMAGE_TYPE_APPLICATION",
	"image_version":    "1.2.3",
	"name":             "sample-image",
	"origin_type":      "ORIGIN_LOCAL",
	"revision": []interface{}{map[string]interface{}{
		"created_at": "2020-12-15T06:21:24Z",
		"created_by": "sample@example.com",
		"curr":       "1",
		"updated_at": "2020-12-15T06:21:24Z",
		"updated_by": "user@example.com",
	}},
	"title": "Sample Title",
}

// efo - Expected Flattened Output
var efoImageFullCfg = map[string]interface{}{
	"datastore_id":     "SAMPLE-DS-ID",
	"description":      "SAMPLE IMAGE DESCRIPTION",
	"id":               "SAMPLE-IMAGE-ID",
	"image_arch":       "AMD64",
	"image_error":      "", // Computed fields will not be filled in cfg
	"image_format":     "QCOW2",
	"image_local":      "", // Computed field
	"image_rel_url":    "sampleRelUrl",
	"image_sha_256":    "sampleSHA",
	"image_size_bytes": "123435",
	"image_status":     "", // Computed field
	"image_type":       "IMAGE_TYPE_APPLICATION",
	"image_version":    "", // Computed field
	"name":             "sample-image",
	"origin_type":      "",              // Computed field
	"revision":         []interface{}{}, // Computed field
	"title":            "Sample Title",
}

// In each test case, call rdXXX to get the appropriate config struct,
//  feed it to flattenXXX, verify output of flattenXXX is same as input
func TestRDTImageConfig(t *testing.T) {
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
			expectedFlattenedOutput: rdImageEmptyOutput,
		},
		{
			input:                   rdImageFullCfg,
			description:             "Full Cfg Struct",
			expectError:             false,
			expectedFlattenedOutput: efoImageFullCfg,
		},
	}

	for _, c := range cases {
		rd := schema.TestResourceDataRaw(t, zschemas.ImageSchema, c.input)
		cfg := &swagger_models.ImageConfig{}
		name, ok := c.input["name"].(string)
		if !ok {
			name = ""
		}
		cfg.Name = &name
		cfg.ID, ok = c.input["id"].(string)
		err := updateImageCfgFromResourceData(cfg, rd)
		if err != nil {
			if !c.expectError {
				t.Fatalf("Test Failed: %s\n"+
					"Unexpected error from updateImageCfgFromResourceData.\n"+
					"Error: %+v\n", c.description, err.Error())
			}
		} else {
			if c.expectError {
				t.Fatalf("Test Failed: %s. Expecting Error, but did not get one",
					c.description)
			}
		}
		out := flattenImageConfig(cfg, false)
		err = verifyFlattenOutput(zschemas.ImageSchema, out, c.expectAllSchemaKeys)
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
