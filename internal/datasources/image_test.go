// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package datasources

import (
	"reflect"
	"testing"

	"github.com/go-test/deep"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	"github.com/zededa/zedcloud-api/swagger_models"
)

var ouputImageEmpty = map[string]interface{}{
	"id":               "",
	"image_error":      "",
	"image_local":      "",
	"image_status":     "",
	"image_version":    "",
	"origin_type":      "",
	"revision":         []interface{}{},
	"datastore_id":     "",
	"description":      "",
	"image_arch":       "",
	"image_format":     "",
	"image_rel_url":    "",
	"image_sha_256":    "",
	"image_size_bytes": "",
	"image_type":       "",
	"name":             "",
	"title":            "",
}

var inputImageFullCfg = &swagger_models.ImageConfig{
	DatastoreID:    strPtr("test datastore id"),
	Description:    "test description",
	ID:             "sample id",
	ImageArch:      modelArchTypePtr("test image arch"),
	ImageError:     "sample error",
	ImageFormat:    configFormatPtr("RAW"),
	ImageLocal:     "Sample local",
	ImageRelURL:    "sample rel url",
	ImageSha256:    "Sample SHA",
	ImageSizeBytes: "0",
	ImageStatus:    imageStatusPtr("Ready"),
	ImageType:      imageTypePtr("IMAGE_TYPE_APPLICATION"),
	ImageVersion:   "Sample Version",
	Name:           strPtr("Sample Name"),
	OriginType:     originPtr("ORIGIN_IMPORTED"),
	Revision: &swagger_models.ObjectRevision{
		CreatedAt: "2020-12-15T06:21:24Z",
		CreatedBy: strPtr("sample@example.com"),
		Curr:      strPtr("1"),
		Prev:      "0",
		UpdatedAt: "2020-12-15T06:21:24Z",
		UpdatedBy: strPtr("user@example.com"),
	},
	Title: strPtr("Sample Title"),
}

var outputImageFullCfg = map[string]interface{}{
	"datastore_id":     "test datastore id",
	"description":      "test description",
	"id":               "sample id",
	"image_arch":       "test image arch",
	"image_error":      "sample error",
	"image_format":     "RAW",
	"image_local":      "Sample local",
	"image_rel_url":    "sample rel url",
	"image_sha_256":    "Sample SHA",
	"image_size_bytes": "0",
	"image_status":     "Ready",
	"image_type":       "IMAGE_TYPE_APPLICATION",
	"image_version":    "Sample Version",
	"name":             "Sample Name",
	"origin_type":      "ORIGIN_IMPORTED",
	"revision": []interface{}{map[string]interface{}{
		"created_at": "2020-12-15T06:21:24Z",
		"created_by": "sample@example.com",
		"curr":       "1",
		"prev":       "0",
		"updated_at": "2020-12-15T06:21:24Z",
		"updated_by": "user@example.com",
	}},
	"title": "Sample Title",
}

func TestFlattenImageConfig(t *testing.T) {
	cases := []struct {
		description         string
		input               *swagger_models.ImageConfig
		expected            map[string]interface{}
		expectAllSchemaKeys bool
	}{
		{
			input:       nil,
			expected:    map[string]interface{}{},
			description: "nil Struct",
		},
		{
			input:       &swagger_models.ImageConfig{},
			expected:    ouputImageEmpty,
			description: "Empty Struct",
		},
		{
			input:       inputImageFullCfg,
			expected:    outputImageFullCfg,
			description: "Full Cfg Struct",
		},
	}

	for _, c := range cases {
		out := flattenImageConfig(c.input)
		err := verifyFlattenOutput(zschemas.ImageSchema, out, c.expectAllSchemaKeys)
		if err != nil {
			t.Fatalf("Test Failed: %s\n Errors in flatten output. Err: %s",
				c.description, err.Error())
		}
		if diff := deep.Equal(out, c.expected); diff != nil {
			t.Fatalf("Test Failed: %s\n"+
				"Error matching output and expected.\n"+
				"Output: %#v\n"+
				"Expected: %#v\n"+
				"Diff: %#v", c.description, out, c.expected, diff)
		}
		if !reflect.DeepEqual(out, c.expected) {
			t.Fatalf("Error matching output and expected: %#v vs %#v", out, c.expected)
		}
	}
}
