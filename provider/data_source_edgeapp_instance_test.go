// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"github.com/go-test/deep"
	"github.com/zededa/zedcloud-api/swagger_models"
	"testing"
)

var ouputAppInstEmpty = map[string]interface{}{
	"bundleversion":         "",
	"cluster_id":            "",
	"id":                    "",
	"is_secret_updated":     false,
	"project_id":            "",
	"purge":                 []interface{}{},
	"refresh":               []interface{}{},
	"restart":               []interface{}{},
	"revision":              []interface{}{},
	"user_defined_version":  "",
	"vminfo":                []interface{}{},
	"activate":              false,
	"app_id":                "",
	"app_policy_id":         "",
	"app_type":              "",
	"collect_stats_ip_addr": "",
	"custom_config":         []interface{}{},
	"description":           "",
	"deployment_type":       "",
	"device_id":             "",
	"drive":                 []interface{}{},
	"interface":             []interface{}{},
	"logs":                  []interface{}{},
	"name":                  "",
	"remote_console":        false,
	"tags":                  map[string]interface{}{},
	"title":                 "",
}

func appInstFullCfg() *swagger_models.AppInstance {
	deploymentTypePtr, _ := deploymentTypePtr("DEPLOYMENT_TYPE_STAND_ALONE")
	return &swagger_models.AppInstance{
		Name:               strPtr("SampleName"),
		ID:                 "Sample ID",
		Description:        "Sample Description",
		Title:              strPtr("Sample Title"),
		Activate:           boolPtr(true),
		AppID:              strPtr("Sample App ID"),
		AppPolicyID:        "Sample App Policy ID",
		AppType:            appTypePtr("APP_TYPE_VM"),
		Bundleversion:      "Sample Bundle Version",
		ClusterID:          "Sample Cluster ID",
		CollectStatsIPAddr: "10.1.1.10",
		CryptoKey:          "Sample Crypto Key",
		CustomConfig: &swagger_models.CustomConfig{
			Add:                true,
			AllowStorageResize: true,
			FieldDelimiter:     ",",
			Name:               "Sample Custom Config",
			Override:           true,
			Template:           "Sample Template",
			VariableGroups: []*swagger_models.CustomConfigVariableGroup{
				&swagger_models.CustomConfigVariableGroup{
					Condition: &swagger_models.VariableGroupCondition{
						Name:     "SampleVGCondition",
						Operator: variableGroupConditionOperatorPtr("Sample Operator"),
						Value:    "Sample Operator Value",
					},
					Name:     "Sample VG NAME",
					Required: true,
					Variables: []*swagger_models.VariableGroupVariable{
						&swagger_models.VariableGroupVariable{
							Default:   "Sample Default Value",
							Encode:    variableFileEncodingPtr("Sample Encoding"),
							Format:    variableVariableFormatPtr("Sample Format"),
							Label:     strPtr("Sample label"),
							MaxLength: "10",
							Name:      strPtr("Sample VGV Name"),
							Options: []*swagger_models.VariableOptionVal{
								&swagger_models.VariableOptionVal{
									Label: "Sample Label",
									Value: "Sample Value",
								},
							},
							ProcessInput: "Sample Input",
							Required:     true,
							Value:        "Sample Value",
						},
					},
				},
			},
		},
		DeploymentType: deploymentTypePtr,
		DeviceID:       strPtr("Sample Device ID"),
		Drives: []*swagger_models.Drive{
			&swagger_models.Drive{
				Cleartext:   true,
				Drvtype:     strPtr("CDROM"),
				Ignorepurge: true,
				Imagename:   strPtr("SampleImageName"),
				Imvolname:   "Sample IMVOLNAME",
				Maxsize:     nil,
				Mountpath:   "Sample Mount Path",
				Mvolname:    "Sample MVolname",
				Preserve:    true,
				Readonly:    true,
				Target:      strPtr("Sample Target"),
				Volumelabel: "SampleVolLabel",
			},
		},
		EncryptedSecrets: map[string]string{
			"key1": "secret1",
			"key2": "secret2",
		},
		Interfaces: []*swagger_models.AppInterface{
			&swagger_models.AppInterface{
				AccessVlanID: 64,
				Acls: []*swagger_models.AppACE{
					&swagger_models.AppACE{
						Actions: []*swagger_models.AppACEAction{
							&swagger_models.AppACEAction{
								Drop:       true,
								Limit:      true,
								Limitburst: int64Ptr(64),
								Limitrate:  int64Ptr(65),
								Limitunit:  strPtr("unit"),
								Mapparams: &swagger_models.AppMapParams{
									Port: int64Ptr(100),
								},
								Portmap: true,
							},
						},
						ID: int32Ptr(124),
						Matches: []*swagger_models.AppACEMatch{
							&swagger_models.AppACEMatch{
								Type:  strPtr("matchType"),
								Value: strPtr("matchValue"),
							},
						},
						Name: strPtr("Sample ACL Name"),
					},
				},
				DefaultNetInstance: true,
				Directattach:       true,
				Intfname:           strPtr("eth10"),
				Io: &swagger_models.PhyAdapter{
					Name: "Sample Phy Adapter",
					Tags: map[string]string{
						"iotag1": "ioval1",
						"iotag2": "ioval2",
					},
					Type: ioTypePtr("IoType"),
				},
				Ipaddr:      strPtr("10.1.1.220"),
				Macaddr:     strPtr("a.1.2"),
				Netinstname: strPtr("Sample Network Instance"),
				Netinsttag: map[string]string{
					"nettag1": "nettagval1",
					"nettag2": "nettagval2",
				},
			},
		},
		IsSecretUpdated: true,
		Logs: &swagger_models.AppInstanceLogs{
			Access: boolPtr(true),
		},
		ProjectID: strPtr("sample-project-id"),
		Purge: &swagger_models.ZedCloudOpsCmd{
			Counter: 10,
			OpsTime: "Sample Purge Ops Time",
		},
		Refresh: &swagger_models.ZedCloudOpsCmd{
			Counter: 11,
			OpsTime: "Sample Refresh Ops Time",
		},
		RemoteConsole: true,
		Restart: &swagger_models.ZedCloudOpsCmd{
			Counter: 12,
			OpsTime: "Sample Restart Ops Time",
		},
		Revision: &swagger_models.ObjectRevision{
			CreatedAt: "2020-12-15T06:21:24Z",
			CreatedBy: strPtr("sample@example.com"),
			Curr:      strPtr("1"),
			Prev:      "0",
			UpdatedAt: "2020-12-15T06:21:24Z",
			UpdatedBy: strPtr("user@example.com"),
		},
		Tags: map[string]string{
			"tag1": "val1",
			"tag2": "val2",
		},
		UserDefinedVersion: "Sample UserDefined Version",
		Vminfo: &swagger_models.VM{
			Cpus:       int64Ptr(16),
			Memory:     int64Ptr(2048),
			Mode:       hvModePtr("sampleCMMode"),
			Vnc:        true,
			VncDisplay: 5,
		},
	}
}

var outputAppInstFullCfg = map[string]interface{}{
	"name":                  "SampleName",
	"id":                    "Sample ID",
	"description":           "Sample Description",
	"title":                 "Sample Title",
	"activate":              true,
	"app_id":                "Sample App ID",
	"app_policy_id":         "Sample App Policy ID",
	"app_type":              "APP_TYPE_VM",
	"bundleversion":         "Sample Bundle Version",
	"cluster_id":            "Sample Cluster ID",
	"collect_stats_ip_addr": "10.1.1.10",
	"custom_config": []interface{}{
		map[string]interface{}{
			"add":                  true,
			"allow_storage_resize": true,
			"field_delimiter":      ",",
			"name":                 "Sample Custom Config",
			"override":             true,
			"template":             "Sample Template",
			"variable_group": []interface{}{
				map[string]interface{}{
					"condition": []interface{}{
						map[string]interface{}{
							"name":     "SampleVGCondition",
							"operator": "Sample Operator",
							"value":    "Sample Operator Value",
						},
					},
					"name":     "Sample VG NAME",
					"required": true,
					"variable": []interface{}{
						map[string]interface{}{
							"default":    "Sample Default Value",
							"encode":     "Sample Encoding",
							"format":     "Sample Format",
							"label":      "Sample label",
							"max_length": "10",
							"name":       "Sample VGV Name",
							"option": []interface{}{
								map[string]interface{}{
									"label": "Sample Label",
									"value": "Sample Value",
								},
							},
							"required": true,
							"value":    "Sample Value",
						},
					},
				},
			},
		},
	},
	"deployment_type": "DEPLOYMENT_TYPE_STAND_ALONE",
	"device_id":       "Sample Device ID",
	"drive": []interface{}{
		map[string]interface{}{
			"cleartext":   true,
			"drvtype":     "CDROM",
			"ignorepurge": true,
			"imagename":   "SampleImageName",
			"imvolname":   "Sample IMVOLNAME",
			"maxsize":     0,
			"mountpath":   "Sample Mount Path",
			"mvolname":    "Sample MVolname",
			"preserve":    true,
			"readonly":    true,
			"target":      "Sample Target",
			"volumelabel": "SampleVolLabel",
		},
	},
	"interface": []interface{}{
		map[string]interface{}{
			"access_vlan_id": 64,
			"acl": []interface{}{
				map[string]interface{}{
					"action": []interface{}{
						map[string]interface{}{
							"drop":       true,
							"limit":      true,
							"limitburst": 64,
							"limitrate":  65,
							"limitunit":  "unit",
							"mapparams": []interface{}{
								map[string]interface{}{
									"port": 100,
								},
							},
							"portmap": true,
						},
					},
					"id": 124,
					"match": []interface{}{
						map[string]interface{}{
							"type":  "matchType",
							"value": "matchValue",
						},
					},
					"name": "Sample ACL Name",
				},
			},
			"default_net_instance": true,
			"directattach":         true,
			"intfname":             "eth10",
			"io": []interface{}{
				map[string]interface{}{
					"name": "Sample Phy Adapter",
					"tags": map[string]interface{}{
						"iotag1": "ioval1",
						"iotag2": "ioval2",
					},
					"type": "IoType",
				},
			},
			"ipaddr":      "10.1.1.220",
			"macaddr":     "a.1.2",
			"netinstname": "Sample Network Instance",
			"netinsttag": map[string]interface{}{
				"nettag1": "nettagval1",
				"nettag2": "nettagval2",
			},
		},
	},
	"is_secret_updated": true,
	"logs": []interface{}{
		map[string]interface{}{
			"access": true,
		},
	},
	"project_id": "sample-project-id",
	"purge": []interface{}{
		map[string]interface{}{
			"counter":  10,
			"ops_time": "Sample Purge Ops Time",
		},
	},
	"refresh": []interface{}{
		map[string]interface{}{
			"counter":  11,
			"ops_time": "Sample Refresh Ops Time",
		},
	},
	"remote_console": true,
	"restart": []interface{}{
		map[string]interface{}{
			"counter":  12,
			"ops_time": "Sample Restart Ops Time",
		},
	},
	"revision": []interface{}{
		map[string]interface{}{
			"created_at": "2020-12-15T06:21:24Z",
			"created_by": "sample@example.com",
			"curr":       "1",
			"prev":       "0",
			"updated_at": "2020-12-15T06:21:24Z",
			"updated_by": "user@example.com",
		},
	},
	"tags": map[string]interface{}{
		"tag1": "val1",
		"tag2": "val2",
	},
	"user_defined_version": "Sample UserDefined Version",
	"vminfo": []interface{}{
		map[string]interface{}{
			"cpus":        16,
			"memory":      2048,
			"mode":        "sampleCMMode",
			"vnc":         true,
			"vnc_display": 5,
		},
	},
}

func TestFlattenAppInstance(t *testing.T) {
	cases := []struct {
		input       *swagger_models.AppInstance
		expected    map[string]interface{}
		description string
	}{
		{
			input:       nil,
			expected:    map[string]interface{}{},
			description: "nil Struct",
		},
		{
			input:       &swagger_models.AppInstance{},
			expected:    ouputAppInstEmpty,
			description: "Empty Struct",
		},
		{
			input:       appInstFullCfg(),
			expected:    outputAppInstFullCfg,
			description: "Full Cfg",
		},
	}

	for _, c := range cases {
		out := flattenAppInstance(c.input)
		if diff := deep.Equal(out, c.expected); diff != nil {
			t.Fatalf("Test Failed: %s\n"+
				"Error matching Flattened output and input.\n"+
				"Output: %#v\n"+
				"Expected Output : %#v\n"+
				"Diff Actual vs Expected: %#v", c.description, out,
				c.expected, diff)
		}
	}
}
