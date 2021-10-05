// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"github.com/go-test/deep"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	"github.com/zededa/zedcloud-api/swagger_models"
	"reflect"
	"testing"
)

var ouputVolInstEmpty = map[string]interface{}{
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

func createObjectRevision() *swagger_models.ObjectRevision {
	return &swagger_models.ObjectRevision{
		CreatedAt: "2020-12-15T06:21:24Z",
		CreatedBy: strPtr("sample@example.com"),
		Curr:      strPtr("2"),
		Prev:      "1",
		UpdatedAt: "2020-12-15T06:21:24Z",
		UpdatedBy: strPtr("user@example.com"),
	}
}

var accessMode = swagger_models.VolumeInstanceAccessModeVOLUMEINSTANCEACCESSMODEREADWRITE
var volType = swagger_models.VolumeInstanceTypeVOLUMEINSTANCETYPEEMPTYDIR
var inputVolInstFullCfg = &swagger_models.VolInstConfig{
	Accessmode:    &accessMode,
	Cleartext:     false,
	ContentTreeID: "test content_tree_id",
	Description:   "test description",
	DeviceID:      "test device_id",
	ID:            "test id",
	Image:         "test image",
	Implicit:      true,
	Label:         "test label",
	Multiattach:   true,
	Name:          "test-name",
	ProjectID:     "test project_id",
	Purge:         nil,
	Revision:      createObjectRevision(),
	SizeBytes:     "0",
	Title:         "test title",
	Type:          &volType,
}

var outputVolInstFullCfg = map[string]interface{}{
	"accessmode":      "VOLUME_INSTANCE_ACCESS_MODE_READWRITE",
	"cleartext":       false,
	"content_tree_id": "test content_tree_id",
	"description":     "test description",
	"device_id":       "test device_id",
	"id":              "test id",
	"image":           "test image",
	"implicit":        true,
	"label":           "test label",
	"multiattach":     true,
	"name":            "test-name",
	"project_id":      "test project_id",
	"purge":           []interface{}{},
	"revision": []interface{}{map[string]interface{}{
		"created_at": "2020-12-15T06:21:24Z",
		"created_by": "sample@example.com",
		"curr":       "2",
		"prev":       "1",
		"updated_at": "2020-12-15T06:21:24Z",
		"updated_by": "user@example.com",
	}},
	"size_bytes": "0",
	"title":      "test title",
	"type":       "VOLUME_INSTANCE_TYPE_EMPTYDIR",
}

func TestFlattenVolInstConfig(t *testing.T) {
	cases := []struct {
		description         string
		input               *swagger_models.VolInstConfig
		expected            map[string]interface{}
		expectAllSchemaKeys bool
	}{
		{
			input:       nil,
			expected:    map[string]interface{}{},
			description: "nil Struct",
		},
		{
			input:       &swagger_models.VolInstConfig{},
			expected:    ouputVolInstEmpty,
			description: "Empty Struct",
		},
		{
			input:       inputVolInstFullCfg,
			expected:    outputVolInstFullCfg,
			description: "Full Cfg Struct",
		},
	}

	for _, c := range cases {
		// The routine itself has verification for computed fields.
		// so no need to verify the output for the "computedOnly" call
		flattenVolInstConfig(c.input, true) // computedOnly
		out := flattenVolInstConfig(c.input, false)
		err := verifyFlattenOutput(zschemas.VolumeInstanceSchema, out, c.expectAllSchemaKeys)
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
