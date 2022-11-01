// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/zededa/terraform-provider-zedcloud/internal/resourcedata"
	"github.com/zededa/terraform-provider-zedcloud/internal/test"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	"github.com/zededa/terraform-provider-zedcloud/utils"
	models "github.com/zededa/zedcloud-api/swagger_models"
)

type setupForFlatteningTests struct {
	stateFromTestdata      *schema.ResourceData
	expectedFlattenedState map[string]interface{}
}

func setupFlatteningTest(t *testing.T, input, output string) setupForFlatteningTests {
	jsonInput, err := ioutil.ReadFile(filepath.Join("testdata", "edge_node", input))
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
	stateFromTestdata := schema.TestResourceDataRaw(t, zschemas.EdgeNodeSchema, mockAPIResponseData)

	jsonOutput, err := ioutil.ReadFile(filepath.Join("testdata", "edge_node", output))
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
		stateFromTestdata:      stateFromTestdata,
		expectedFlattenedState: test.FloatToInt(expectedFlattenedState),
	}
}

func TestEdgeNodeFlattening(t *testing.T) {
	t.Run("flatten local empty state", func(t *testing.T) {
		t.Parallel()

		input := "flatten/empty/input.json"
		output := "flatten/empty/output.json"
		setup := setupFlatteningTest(t, input, output)

		_, err := getStateForEdgeNodeCreation(setup.stateFromTestdata)
		if !assert.Error(t, err) {
			t.Log("expect due to missing eve_image_version")
			t.FailNow()
		}
	})
	t.Run("flatten local state for edge node creation", func(t *testing.T) {
		t.Parallel()

		input := "flatten/create/input_flatten_complete.json"
		output := "flatten/create/output_flatten_complete.json"
		setup := setupFlatteningTest(t, input, output)

		localState, err := getStateForEdgeNodeCreation(setup.stateFromTestdata)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		// create request doesn't submit ID
		localState.ID = efoEdgeNodeFullCfg["id"].(string)

		// FIXME: why don't we set the base image in local state/create request but check for its existence in d?
		localEVEVersion := resourcedata.GetStr(setup.stateFromTestdata, "eve_image_version")
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
	t.Run("flatten local state for edge node creation: missing eve_image_version", func(t *testing.T) {
		t.Parallel()

		input := "flatten/create/input_missing_eve_image.json"
		output := "noop.json"
		setup := setupFlatteningTest(t, input, output)

		_, err := getStateForEdgeNodeCreation(setup.stateFromTestdata)
		if !assert.Error(t, err) {
			t.Log("expect due to missing eve_image_version")
			t.FailNow()
		}
	})
	t.Run("flatten local state for edge node creation: missing model_id", func(t *testing.T) {
		t.Parallel()

		input := "flatten/create/input_missing_model_id.json"
		output := "noop.json"
		setup := setupFlatteningTest(t, input, output)

		_, err := getStateForEdgeNodeCreation(setup.stateFromTestdata)
		if !assert.Error(t, err) {
			t.Log("expect due to missing model_id")
			t.FailNow()
		}
	})
	t.Run("flatten local state for edge node creation: missing project_id", func(t *testing.T) {
		t.Parallel()

		input := "flatten/create/input_missing_project_id.json"
		output := "noop.json"
		setup := setupFlatteningTest(t, input, output)

		_, err := getStateForEdgeNodeCreation(setup.stateFromTestdata)
		if !assert.Error(t, err) {
			t.Log("expect due to missing project_id")
			t.FailNow()
		}
	})
	t.Run("flatten local state for edge node update: success (no model_id)", func(t *testing.T) {
		t.Parallel()

		input := "flatten/update/input_flatten_complete.json"
		output := "flatten/update/output_flatten_complete.json"
		setup := setupFlatteningTest(t, input, output)

		remoteState := &models.DeviceConfig{
			ID: efoEdgeNodeFullCfg["id"].(string),
		}
		stateForUpdatingRemote, err := getStateForEdgeNodeUpdate(remoteState, setup.stateFromTestdata)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		// create request doesn't submit ID
		stateForUpdatingRemote.ID = efoEdgeNodeFullCfg["id"].(string)

		// FIXME: why don't we set the base image in local state/create request but check for its existence in d?
		localEVEVersion := resourcedata.GetStr(setup.stateFromTestdata, "eve_image_version")
		stateForUpdatingRemote.BaseImage = []*models.BaseOSImage{
			{
				ImageName: &localEVEVersion,
				Version:   &localEVEVersion,
				Activate:  true,
			},
		}

		flattenedState, err := flattenEdgeNodeState(stateForUpdatingRemote)
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
	t.Run("flatten local state for edge node update: missing eve_image_version", func(t *testing.T) {
		t.Parallel()

		input := "flatten/update/input_missing_eve_image.json"
		output := "noop.json"
		setup := setupFlatteningTest(t, input, output)

		remoteState := &models.DeviceConfig{
			ID: efoEdgeNodeFullCfg["id"].(string),
		}
		_, err := getStateForEdgeNodeUpdate(remoteState, setup.stateFromTestdata)
		if !assert.Error(t, err) {
			t.Log("expect due to missing eve_image_version")
			t.FailNow()
		}
	})
	t.Run("flatten local state for edge node update: change project_id", func(t *testing.T) {
		t.Parallel()

		input := "flatten/update/input_flatten_complete.json"
		output := "flatten/update/output_flatten_complete.json"
		setup := setupFlatteningTest(t, input, output)

		remoteState := &models.DeviceConfig{
			ID:        efoEdgeNodeFullCfg["id"].(string),
			ProjectID: utils.StrPtr("updated_id"),
		}

		_, err := getStateForEdgeNodeUpdate(remoteState, setup.stateFromTestdata)
		if !assert.Error(t, err) {
			t.Log("expect due to changed missing project_id")
			t.FailNow()
		}
	})
	t.Run("flatten local state for edge node update: set project_id to nil", func(t *testing.T) {
		t.Parallel()

		input := "flatten/update/input_flatten_complete.json"
		output := "flatten/update/output_project_id_nil.json"
		setup := setupFlatteningTest(t, input, output)

		remoteState := &models.DeviceConfig{
			ID: efoEdgeNodeFullCfg["id"].(string),
		}

		// set project_id to nil
		setup.stateFromTestdata.Set("project_id", nil)

		stateForUpdatingRemote, err := getStateForEdgeNodeUpdate(remoteState, setup.stateFromTestdata)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		// create request doesn't submit ID
		stateForUpdatingRemote.ID = efoEdgeNodeFullCfg["id"].(string)

		// FIXME: why don't we set the base image in local state/create request but check for its existence in d?
		localEVEVersion := resourcedata.GetStr(setup.stateFromTestdata, "eve_image_version")
		stateForUpdatingRemote.BaseImage = []*models.BaseOSImage{
			{
				ImageName: &localEVEVersion,
				Version:   &localEVEVersion,
				Activate:  true,
			},
		}

		flattenedState, err := flattenEdgeNodeState(stateForUpdatingRemote)
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
