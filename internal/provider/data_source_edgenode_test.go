// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

// import (
// 	"testing"

// 	"github.com/go-test/deep"
// 	"github.com/zededa/terraform-provider-zedcloud/utils"
// 	"github.com/zededa/zedcloud-api/swagger_models"
// )

// var ouputEdgeNodeEmpty = map[string]interface{}{
// 	"adminstate":        "",
// 	"adminstate_config": "",
// 	"asset_id":          "",
// 	"client_ip":         "",
// 	"cluster_id":        "",
// 	"config_items":      map[string]interface{}{},
// 	"cpu":               0,
// 	"description":       "",
// 	"dev_location":      []interface{}{},
// 	"eve_image_version": "",
// 	"id":                "",
// 	"interface":         []interface{}{},
// 	"memory":            0,
// 	"model_id":          "",
// 	"name":              "",
// 	"project_id":        "",
// 	"reset_counter":     0,
// 	"reset_time":        "",
// 	"revision":          []interface{}{},
// 	"serialno":          "",
// 	"storage":           0,
// 	"tags":              map[string]interface{}{},
// 	"title":             "",
// 	"thread":            0,
// 	"utype":             "",
// }

// func edgeNodeFullCfg() *swagger_models.DeviceConfig {
// 	return &swagger_models.DeviceConfig{
// 		Name:        utils.StrPtr("SampleName"),
// 		ID:          "Sample ID",
// 		Description: "Sample Description",
// 		Title:       utils.StrPtr("Sample Title"),
// 		AdminState:  adminStatePtr("RUNNING"),
// 		AssetID:     "Sample AssetID",
// 		BaseImage: []*swagger_models.BaseOSImage{
// 			&swagger_models.BaseOSImage{
// 				ImageName: utils.StrPtr("Sample eve image ver"),
// 				Version:   utils.StrPtr("Sample eve image ver"),
// 				Activate:  true,
// 			},
// 		},
// 		ClientIP:  "10.1.1.2",
// 		ClusterID: "Sample Cluster ID",
// 		ConfigItem: []*swagger_models.EDConfigItem{
// 			&swagger_models.EDConfigItem{
// 				Key:         "item1",
// 				StringValue: "val1",
// 			},
// 			&swagger_models.EDConfigItem{
// 				Key:         "item2",
// 				StringValue: "val2",
// 			},
// 		},
// 		DevLocation: &swagger_models.GeoLocation{
// 			City:       "Sample City",
// 			Country:    "Sample COuntry",
// 			Freeloc:    "Sample Freeloc",
// 			Hostname:   "Sample Hostname",
// 			Loc:        "Sample Location - lat/long ",
// 			Org:        "sample org",
// 			Postal:     "sample postal pin",
// 			Region:     "sample region",
// 			UnderlayIP: "10.1.1.1",
// 		},
// 		Interfaces: []*swagger_models.SysInterface{
// 			&swagger_models.SysInterface{
// 				Cost:      10,
// 				IntfUsage: adapterUsagePtr("ADAPTER_USAGE_APP_DIRECT"),
// 				Intfname:  "eth0",
// 				Ipaddr:    "10.1.1.1",
// 				Macaddr:   "a.b.c",
// 				Netname:   "sampleNetName",
// 				Tags: map[string]string{
// 					"tag1": "val1",
// 					"tag2": "val2",
// 				},
// 			},
// 			&swagger_models.SysInterface{
// 				Cost:      20,
// 				IntfUsage: adapterUsagePtr("ADAPTER_USAGE_MANAGEMENT"),
// 				Intfname:  "eth1",
// 				Ipaddr:    "10.1.1.2",
// 				Macaddr:   "a.b.d",
// 				Netname:   "sampleNetName2",
// 				Tags: map[string]string{
// 					"intftag12": "intfval12",
// 					"intftag22": "intfval22",
// 				},
// 			},
// 		},
// 		Memory:       1024,
// 		ModelID:      utils.StrPtr("Sample Model ID"),
// 		ProjectID:    utils.StrPtr("sample-project-id"),
// 		ResetCounter: 10,
// 		ResetTime:    "Sample Reset Time",
// 		Revision: &swagger_models.ObjectRevision{
// 			CreatedAt: "2020-12-15T06:21:24Z",
// 			CreatedBy: utils.StrPtr("sample@example.com"),
// 			Curr:      utils.StrPtr("1"),
// 			Prev:      "0",
// 			UpdatedAt: "2020-12-15T06:21:24Z",
// 			UpdatedBy: utils.StrPtr("user@example.com"),
// 		},
// 		Serialno: "Sample Serial Num",
// 		Storage:  2048,
// 		Tags: map[string]string{
// 			"tag1": "val1",
// 			"tag2": "val2",
// 		},
// 		Thread: 8,
// 		Utype:  modelArchTypePtr("Sample UType"),
// 	}
// }

// var outputEdgeNodeFullCfg = map[string]interface{}{
// 	"name":              "SampleName",
// 	"id":                "Sample ID",
// 	"description":       "Sample Description",
// 	"title":             "Sample Title",
// 	"adminstate":        "RUNNING",
// 	"adminstate_config": "RUNNING",
// 	"asset_id":          "Sample AssetID",
// 	"client_ip":         "10.1.1.2",
// 	"cluster_id":        "Sample Cluster ID",
// 	"config_items": map[string]interface{}{
// 		"item1": "val1",
// 		"item2": "val2",
// 	},
// 	"cpu": 0,
// 	"dev_location": []interface{}{
// 		map[string]interface{}{
// 			"city":        "Sample City",
// 			"country":     "Sample COuntry",
// 			"freeloc":     "Sample Freeloc",
// 			"hostname":    "Sample Hostname",
// 			"loc":         "Sample Location - lat/long ",
// 			"org":         "sample org",
// 			"postal":      "sample postal pin",
// 			"region":      "sample region",
// 			"underlay_ip": "10.1.1.1",
// 		},
// 	},
// 	"eve_image_version": "Sample eve image ver",
// 	"interface": []interface{}{
// 		map[string]interface{}{
// 			"cost":       10,
// 			"intf_usage": "ADAPTER_USAGE_APP_DIRECT",
// 			"intfname":   "eth0",
// 			"ipaddr":     "10.1.1.1",
// 			"macaddr":    "a.b.c",
// 			"netname":    "sampleNetName",
// 			"tags": map[string]interface{}{
// 				"tag1": "val1",
// 				"tag2": "val2",
// 			},
// 		},
// 		map[string]interface{}{
// 			"cost":       20,
// 			"intf_usage": "ADAPTER_USAGE_MANAGEMENT",
// 			"intfname":   "eth1",
// 			"ipaddr":     "10.1.1.2",
// 			"macaddr":    "a.b.d",
// 			"netname":    "sampleNetName2",
// 			"tags": map[string]interface{}{
// 				"intftag12": "intfval12",
// 				"intftag22": "intfval22",
// 			},
// 		},
// 	},
// 	"memory":        1024,
// 	"model_id":      "Sample Model ID",
// 	"project_id":    "sample-project-id",
// 	"reset_counter": 10,
// 	"reset_time":    "Sample Reset Time",
// 	"revision": []interface{}{map[string]interface{}{
// 		"created_at": "2020-12-15T06:21:24Z",
// 		"created_by": "sample@example.com",
// 		"curr":       "1",
// 		"prev":       "0",
// 		"updated_at": "2020-12-15T06:21:24Z",
// 		"updated_by": "user@example.com",
// 	}},
// 	"serialno": "Sample Serial Num",
// 	"storage":  2048,
// 	"tags": map[string]interface{}{
// 		"tag1": "val1",
// 		"tag2": "val2",
// 	},
// 	"thread": 8,
// 	"utype":  "Sample UType",
// }

// func TestFlattenDeviceConfig(t *testing.T) {
// 	cases := []struct {
// 		input       *swagger_models.DeviceConfig
// 		expected    map[string]interface{}
// 		description string
// 	}{
// 		{
// 			input:       nil,
// 			expected:    map[string]interface{}{},
// 			description: "nil Struct",
// 		},
// 		{
// 			input:       &swagger_models.DeviceConfig{},
// 			expected:    ouputEdgeNodeEmpty,
// 			description: "Empty Struct",
// 		},
// 		{
// 			input:       edgeNodeFullCfg(),
// 			expected:    outputEdgeNodeFullCfg,
// 			description: "Full Cfg",
// 		},
// 	}

// 	for _, c := range cases {
// 		out := flattenEdgeNodeState(c.input)
// 		if diff := deep.Equal(out, c.expected); diff != nil {
// 			t.Fatalf("Test Failed: %s\n"+
// 				"Error matching Flattened output and input.\n"+
// 				"Output: %#v\n"+
// 				"Expected Output : %#v\n"+
// 				"Diff Actual vs Expected: %#v", c.description, out,
// 				c.expected, diff)
// 		}
// 	}
// }
