// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/zededa/terraform-provider-zedcloud/internal/resourcedata"
	"github.com/zededa/terraform-provider-zedcloud/internal/test"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	"github.com/zededa/terraform-provider-zedcloud/utils"
	models "github.com/zededa/zedcloud-api/swagger_models"
)

func copyMap(originalMap map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	for k, v := range originalMap {
		newMap[k] = v
	}
	return newMap
}

var rdEdgeNodeEmptyOutput = map[string]interface{}{
	"adminstate":        "",
	"adminstate_config": "",
	"asset_id":          "",
	"client_ip":         "",
	"cluster_id":        "",
	"config_items":      map[string]interface{}{},
	"cpu":               0,
	"description":       "",
	"dev_location":      []interface{}{},
	"eve_image_version": "",
	"id":                "",
	"interface":         []interface{}{},
	"memory":            0,
	"model_id":          "",
	"name":              "",
	"project_id":        "",
	"reset_counter":     0,
	"reset_time":        "",
	"revision":          []interface{}{},
	"serialno":          "",
	"storage":           0,
	"tags":              map[string]interface{}{},
	"thread":            0,
	"title":             "",
	"utype":             "",
}

var rdEdgeNodeFullCfg = map[string]interface{}{
	"name":              "sample-EdgeNode",
	"id":                "SAMPLE-EdgeNode-ID",
	"description":       "SAMPLE EdgeNode DESCRIPTION",
	"title":             "sample-title 2",
	"adminstate":        "",
	"adminstate_config": "ADMIN_STATE_ACTIVE",
	"asset_id":          "Sample Asset ID",
	"client_ip":         "10.10.1.1",
	"cluster_id":        "sample-cluster-id",
	"config_items":      map[string]interface{}{},
	"dev_location": []interface{}{
		map[string]interface{}{
			"city":        "sample-city",
			"country":     "sample-country",
			"freeloc":     "sample-freeloc",
			"hostname":    "sample-host",
			"loc":         "sample-loc",
			"org":         "sample-org",
			"postal":      "sample-postal",
			"region":      "sample-region",
			"underlay_ip": "10.10.1.1",
		},
	},
	"eve_image_version": "6.8.2-kvm-amd64",
	"interface": []interface{}{
		map[string]interface{}{
			"cost":       1,
			"intfname":   "intf2",
			"intf_usage": "ADAPTER_USAGE_DISABLED",
			"ipaddr":     "10.10.1.6",
			"macaddr":    "0.1.3",
			"netname":    "sample-net2",
			"tags": map[string]interface{}{
				"tag12": "value12",
				"tag22": "value22",
			},
		},
		map[string]interface{}{
			"cost":       10,
			"intfname":   "intf1",
			"intf_usage": "ADAPTER_USAGE_MANAGEMENT",
			"ipaddr":     "10.10.1.5",
			"macaddr":    "0.1.2",
			"netname":    "sample-net",
			"tags": map[string]interface{}{
				"tag1": "value1",
				"tag2": "value2",
			},
		},
	},
	"memory":        "",
	"model_id":      "Test Model",
	"project_id":    "sample-project",
	"reset_counter": 10,
	"reset_time":    "sample-time",
	"revision": []interface{}{map[string]interface{}{
		"created_at": "2020-12-15T06:21:24Z",
		"created_by": "sample@example.com",
		"curr":       "1",
		"updated_at": "2020-12-15T06:21:24Z",
		"updated_by": "user@example.com",
	}},
	"serialno": "sample-serial-no",
	"storage":  100,
	"tags": map[string]interface{}{
		"edtag1": "edval1",
		"edtag2": "edval2",
	},
	"thread": 2,
	"utype":  "AMD64",
}

// efo - Expected Flattened Output
//
//	var efoEdgeNodeFullCfg = map[string]interface{}{
//		"name":        rdEdgeNodeFullCfg["name"],
//		"id":          rdEdgeNodeFullCfg["id"],
//		"description": rdEdgeNodeFullCfg["description"],
//		"title":       rdEdgeNodeFullCfg["title"],
//		"adminstate":  "",
//		// adminstate config is not set in rdDeviceConfig - so ignoring this field.
//		"adminstate_config": "",
//		"asset_id":          rdEdgeNodeFullCfg["asset_id"],
//		"client_ip":         rdEdgeNodeFullCfg["client_ip"],
//		"cluster_id":        rdEdgeNodeFullCfg["cluster_id"],
//		"config_items":      rdEdgeNodeFullCfg["config_items"],
//		"cpu":               0,
//		"dev_location":      rdEdgeNodeFullCfg["dev_location"],
//		"eve_image_version": rdEdgeNodeFullCfg["eve_image_version"],
//		"interface":         rdEdgeNodeFullCfg["interface"],
//		"memory":            0,
//		"model_id":          rdEdgeNodeFullCfg["model_id"],
//		"project_id":        rdEdgeNodeFullCfg["project_id"],
//		"reset_counter":     0,
//		"reset_time":        "",
//		"revision":          []interface{}{},
//		"serialno":          rdEdgeNodeFullCfg["serialno"],
//		"storage":           0,
//		"tags":              rdEdgeNodeFullCfg["tags"],
//		"thread":            0,
//		"utype":             "",
//	}
var efoEdgeNodeFullCfg = map[string]interface{}{
	"name":        rdEdgeNodeFullCfg["name"],
	"id":          rdEdgeNodeFullCfg["id"],
	"description": rdEdgeNodeFullCfg["description"],
	"title":       rdEdgeNodeFullCfg["title"],
	"adminstate":  "",
	// adminstate config is not set in rdDeviceConfig - so ignoring this field.
	"adminstate_config": "",
	"asset_id":          rdEdgeNodeFullCfg["asset_id"],
	"client_ip":         rdEdgeNodeFullCfg["client_ip"],
	"cluster_id":        rdEdgeNodeFullCfg["cluster_id"],
	"config_items":      rdEdgeNodeFullCfg["config_items"],
	"cpu":               0,
	"dev_location":      rdEdgeNodeFullCfg["dev_location"],
	"eve_image_version": rdEdgeNodeFullCfg["eve_image_version"],
	"interface":         rdEdgeNodeFullCfg["interface"],
	"memory":            0,
	"model_id":          rdEdgeNodeFullCfg["model_id"],
	"project_id":        rdEdgeNodeFullCfg["project_id"],
	"reset_counter":     0,
	"reset_time":        "",
	"revision":          []interface{}{},
	"serialno":          rdEdgeNodeFullCfg["serialno"],
	"storage":           0,
	"tags":              rdEdgeNodeFullCfg["tags"],
	"thread":            0,
	"utype":             "",
}

var rdEdgeNodeNilProjectId = map[string]interface{}{
	"name":              "sample-EdgeNode",
	"id":                "SAMPLE-EdgeNode-ID",
	"adminstate":        "",
	"adminstate_config": "ADMIN_STATE_ACTIVE",
	"eve_image_version": "6.8.2-kvm-amd64",
	"interface":         []interface{}{},
	"model_id":          "Test Model",
}

func rdEdgeNodeNilProjectIdOutput() map[string]interface{} {
	newMap := copyMap(rdEdgeNodeEmptyOutput)
	newMap["eve_image_version"] = rdEdgeNodeNilProjectId["eve_image_version"]
	newMap["id"] = rdEdgeNodeNilProjectId["id"]
	newMap["interface"] = rdEdgeNodeNilProjectId["interface"]
	newMap["name"] = rdEdgeNodeNilProjectId["name"]
	return newMap
}

var rdEdgeNodeEmptyProjectId = map[string]interface{}{
	"name":              "sample-EdgeNode",
	"id":                "SAMPLE-EdgeNode-ID",
	"adminstate":        "",
	"adminstate_config": "ADMIN_STATE_ACTIVE",
	"eve_image_version": "6.8.2-kvm-amd64",
	"interface":         []interface{}{},
	"model_id":          "Test Model",
	"project_id":        "",
}

func rdEdgeNodeUpdateProjectIdToNilOutput() map[string]interface{} {
	newMap := copyMap(rdEdgeNodeNilProjectIdOutput())
	newMap["project_id"] = "SampleProjectID"
	return newMap
}

type setupForFlatteningTests struct {
	remoteStateFromTestdata *schema.ResourceData
	expectedFlattenedState  map[string]interface{}
}

func setupFlatteningTest(t *testing.T) setupForFlatteningTests {
	jsonInput, err := ioutil.ReadFile("./testdata/edge_node/input_flatten_complete.json")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	var mockAPIResponseData map[string]interface{}
	if err := json.Unmarshal(jsonInput, &mockAPIResponseData); err != nil {
		t.Log(err)
		t.FailNow()
	}
	// provide the test data (api-reponse) to be marshaled into the schema and transformed into resource
	// data that represent the remote state
	remoteStateFromTestdata := schema.TestResourceDataRaw(t, zschemas.EdgeNodeSchema, mockAPIResponseData)

	jsonOutput, err := ioutil.ReadFile("./testdata/edge_node/output_flatten_complete.json")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	var expectedFlattenedState map[string]interface{}
	if err := json.Unmarshal(jsonOutput, &expectedFlattenedState); err != nil {
		t.Log(err)
		t.FailNow()
	}

	return setupForFlatteningTests{
		remoteStateFromTestdata: remoteStateFromTestdata,
		expectedFlattenedState:  test.MapItemsFloatToInt(expectedFlattenedState),
	}
}

// apiResponseJSON, err := json.MarshalIndent(efoEdgeNodeFullCfg, "", "    ")
//
//	if err != nil {
//		t.Log(err)
//		t.FailNow()
//	}
//
// fmt.Println(string(apiResponseJSON))
func TestEdgeNodeFlattening(t *testing.T) {
	t.Run("flatten local state for edgenode creation", func(t *testing.T) {
		setup := setupFlatteningTest(t)

		localState, err := getStateForEdgeNodeCreation(setup.remoteStateFromTestdata)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		// we need to set the ID we expect to be returned in the api repsonse
		localState.ID = efoEdgeNodeFullCfg["id"].(string)
		// FIXME: why don't we set the base image in local state/create request but check for its existence in d?
		localEVEVersion := resourcedata.GetStr(setup.remoteStateFromTestdata, "eve_image_version")
		localState.BaseImage = []*models.BaseOSImage{
			{
				ImageName: &localEVEVersion,
				Version:   &localEVEVersion,
				Activate:  true,
			},
		}

		flattenedState, err := flattenEdgeNodeState(localState)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		if err := utils.CheckIfAllKeysExist(zschemas.EdgeNodeSchema, flattenedState); err != nil {
			t.Log(err)
			t.FailNow()
		}

		assert.Equal(t, setup.expectedFlattenedState, flattenedState)
	})
}

// In each test case, call rdXXX to get the appropriate config struct,
//
//	feed it to flattenXXX, verify output of flattenXXX is same as input
// func TestRDEdgeNodeConfig(t *testing.T) {
// 	testCases := []struct {
// 		input                   map[string]interface{}
// 		update                  bool
// 		update_project_id       bool
// 		description             string
// 		expectError             bool
// 		expectedFlattenedOutput map[string]interface{}
// 		expectAllSchemaKeys     bool
// 	}{
// 		{
// 			input:                   map[string]interface{}{},
// 			description:             "Empty RD",
// 			expectError:             true,
// 			expectedFlattenedOutput: rdEdgeNodeEmptyOutput,
// 		},
// 		{
// 			input:                   rdEdgeNodeFullCfg,
// 			description:             "Full Cfg Struct",
// 			expectError:             false,
// 			expectedFlattenedOutput: efoEdgeNodeFullCfg,
// 		},
// 		{
// 			input:                   rdEdgeNodeNilProjectId,
// 			description:             "Nil Project Create",
// 			expectError:             true,
// 			expectedFlattenedOutput: rdEdgeNodeNilProjectIdOutput(),
// 		},
// 		{
// 			input:                   rdEdgeNodeNilProjectId,
// 			description:             "Update non-nil project id to Nil",
// 			expectError:             false,
// 			update:                  true,
// 			update_project_id:       true,
// 			expectedFlattenedOutput: rdEdgeNodeUpdateProjectIdToNilOutput(),
// 		},
// 		{
// 			input:                   rdEdgeNodeEmptyProjectId,
// 			description:             "Empty Project Update",
// 			expectError:             false,
// 			update:                  true,
// 			expectedFlattenedOutput: rdEdgeNodeNilProjectIdOutput(),
// 		},
// 	}

// 	for _, tt := range testCases {
// 		// provide the test data (api-reponse) to be marshaled into the schema
// 		// and transformed into resource data that represent the remote state
// 		remoteState := schema.TestResourceDataRaw(t, zschemas.EdgeNodeSchema, tt.input)

// 		localState := &models.DeviceConfig{}
// 		name, ok := tt.input["name"].(string)
// 		if !ok {
// 			name = ""
// 		}

// 		localState.Name = &name
// 		localState.ID, ok = tt.input["id"].(string)

// 		localState.BaseImage = cfgBaseosForEveVersionStr(resourcedata.GetStr(remoteState, "eve_image_version"))
// 		if tt.update_project_id {
// 			project_id := "SampleProjectID"
// 			localState.ProjectID = &project_id
// 		}

// 		err := getEdgeNodeState(localState, remoteState, !tt.update)
// 		if err != nil {
// 			if !tt.expectError {
// 				t.Fatalf("Test Failed: %s\n"+
// 					"Unexpected error from updateEdgeNodeCfgFromResourceData.\n"+
// 					"Error: %+v\n",
// 					tt.description,
// 					err.Error())
// 			}
// 			// No point in continuing cfg is invalid
// 			continue
// 		}
// 		if tt.expectError {
// 			t.Fatalf("Test Failed: %s. Expecting Error, but did not get one", tt.description)
// 		}

// 		out := flattenEdgeNodeState(localState)
// 		if err := verifyFlattenOutput(zschemas.EdgeNodeSchema, out, tt.expectAllSchemaKeys);err != nil {
// 			t.Fatalf("Test Failed: %s\n Errors in flatten output. Err: %s", tt.description, err.Error())
// 		}

// 		if diff := deep.Equal(out, tt.expectedFlattenedOutput); diff != nil {
// 			t.Fatalf("Test Failed: %s\n"+
// 				"Error matching Flattened output and input.\n"+
// 				"Output: %#v\n"+
// 				"expectedFlattenedOutput : %#v\n"+
// 				"Diff: %#v",
// 				tt.description,
// 				out,
// 				tt.expectedFlattenedOutput,
// 				diff)
// 		}
// 	}
// }
