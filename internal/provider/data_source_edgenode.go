// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/internal/client"
	zerrors "github.com/zededa/terraform-provider-zedcloud/internal/errors"
	"github.com/zededa/terraform-provider-zedcloud/internal/resourcedata"
	"github.com/zededa/terraform-provider-zedcloud/internal/state"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	"github.com/zededa/terraform-provider-zedcloud/utils"
	zedcloudAPI "github.com/zededa/zedcloud-api"
	models "github.com/zededa/zedcloud-api/swagger_models"
)

func newEdgeNodeDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readEdgeNodeDataSource,
		Schema:      zschemas.EdgeNodeSchema,
		Description: "Schema for data source zedcloud_edgenode. Must specify id or name",
	}
}

func readEdgeNodeDataSource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	// required fields
	id := resourcedata.GetStr(d, "id")
	name := resourcedata.GetStr(d, "name")
	if id == "" && name == "" {
		return diag.Errorf("missing required fields \"id\" or \"name\" in resource data")
	}

	apiClient, ok := meta.(*client.Client)
	if !ok {
		return diag.FromErr(fmt.Errorf("expect meta to be of type client.Client{} but is %T", meta))
	}

	// fetch the object from zedcloud api
	remoteState, err := fetchEdgeNodeState(apiClient, name, id)
	if err != nil {
		// no object with this id exist in the api's state
		var notFoundErr *zerrors.ObjectNotFound
		if errors.As(err, &notFoundErr) {
			// we need to remove it from local state
			state.RemoveLocal(d, resourceTypeEdgeNode, id, name)
			return diag.FromErr(fmt.Errorf(
				"[ERROR] could not find %s %s (id=%s): %w",
				resourceTypeEdgeNode,
				name,
				id,
				err,
			))
		}

		// api error
		return diag.FromErr(fmt.Errorf(
			"[ERROR] could not fetch %s %s (id=%s): %w",
			resourceTypeEdgeNode,
			name,
			id,
			err,
		))
	}

	// flatten the api response into normalized terraform format
	flattenedRemoteState, err := flattenEdgeNodeState(remoteState)
	if err != nil {
		return diag.FromErr(fmt.Errorf(
			"[ERROR] could not flatten api state for %s %s (id=%s): %w",
			resourceTypeEdgeNode,
			name,
			id,
			err,
		))
	}

	// we always need to set the state even in a read,
	// the object could have been changed outside of terraform
	state.SetLocal(d, flattenedRemoteState)

	return diags
}

func fetchEdgeNodeState(apiClient *client.Client, name, id string) (*models.DeviceConfig, error) {
	responseData := &models.DeviceConfig{}
	resp, err := apiClient.GetObj(deviceUrlExtension, name, id, false, responseData)
	if err != nil {
		return nil, err
	}

	if zedcloudAPI.IsObjectNotFound(resp) {
		return nil, &zerrors.ObjectNotFound{id}
	}

	return responseData, nil
}

func flattenEdgeNodeState(state *models.DeviceConfig) (map[string]interface{}, error) {
	if state == nil {
		return map[string]interface{}{}, nil
	}
	eveImageVersion := ""
	for _, image := range state.BaseImage {
		if image.Activate {
			eveImageVersion = *image.ImageName
			break
		}
	}

	flattened := map[string]interface{}{
		"adminstate":    utils.PtrValStr(state.AdminState),
		"cluster_id":    state.ClusterID,
		"cpu":           int(state.CPU),
		"id":            state.ID,
		"memory":        int(state.Memory),
		"reset_counter": int(state.ResetCounter),
		"reset_time":    state.ResetTime,
		"revision":      utils.FlattenObjectRevision(state.Revision),
		"storage":       int(state.Storage),
		"thread":        int(state.Thread),
		"utype":         utils.PtrValStr(state.Utype),
	}
	flattened["adminstate_config"] = utils.PtrValStr(state.AdminState)
	flattened["asset_id"] = state.AssetID
	flattened["client_ip"] = state.ClientIP
	flattened["config_items"] = flattenEDConfigItems(state.ConfigItem)
	flattened["description"] = state.Description
	flattened["dev_location"] = flattenGeoLocation(state.DevLocation)
	flattened["eve_image_version"] = eveImageVersion
	flattened["interface"] = flattenSysInterfaces(state.Interfaces)
	flattened["model_id"] = utils.PtrValStr(state.ModelID)
	flattened["name"] = utils.PtrValStr(state.Name)
	flattened["project_id"] = utils.PtrValStr(state.ProjectID)
	flattened["serialno"] = state.Serialno
	flattened["tags"] = utils.FlattenStringMap(state.Tags)
	flattened["title"] = utils.PtrValStr(state.Title)

	if err := utils.CheckIfAllKeysExist(zschemas.EdgeNodeSchema, flattened); err != nil {
		return nil, err
	}

	return flattened, nil
}

func flattenGeoLocation(entry interface{}) []interface{} {
	loc, ok := entry.(*models.GeoLocation)
	if !ok {
		return []interface{}{}
	}
	if loc == nil {
		return []interface{}{}
	}
	return []interface{}{
		map[string]interface{}{
			"city":        loc.City,
			"country":     loc.Country,
			"freeloc":     loc.Freeloc,
			"hostname":    loc.Hostname,
			"loc":         loc.Loc,
			"org":         loc.Org,
			"postal":      loc.Postal,
			"region":      loc.Region,
			"underlay_ip": loc.UnderlayIP,
		}}
}

func flattenSysInterfaces(interfaces []*models.SysInterface) []interface{} {
	flattened := make([]interface{}, 0)
	for _, sysInterface := range interfaces {
		flattened = append(flattened, map[string]interface{}{
			"cost":       int(sysInterface.Cost),
			"intfname":   sysInterface.Intfname,
			"intf_usage": utils.PtrValStr(sysInterface.IntfUsage),
			"ipaddr":     sysInterface.Ipaddr,
			"macaddr":    sysInterface.Macaddr,
			"netname":    sysInterface.Netname,
			"tags":       utils.FlattenStringMap(sysInterface.Tags),
		})
	}
	return flattened
}

func flattenEDConfigItems(items []*models.EDConfigItem) map[string]interface{} {
	flattened := make(map[string]interface{})
	for _, item := range items {
		flattened[item.Key] = item.StringValue
	}
	return flattened
}
