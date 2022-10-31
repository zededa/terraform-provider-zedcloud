// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/internal/client"

	zerrors "github.com/zededa/terraform-provider-zedcloud/internal/errors"
	"github.com/zededa/terraform-provider-zedcloud/internal/resourcedata"
	"github.com/zededa/terraform-provider-zedcloud/internal/state"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	zedcloudapi "github.com/zededa/zedcloud-api"
	models "github.com/zededa/zedcloud-api/swagger_models"
)

const (
	deviceUrlExtension = "devices"
	objectTypeEdgeNode = "EdgeNode"
)

func getEdgeNodeUrl(name, edgeNodeID, urlType string) string {
	return zedcloudapi.UrlForObjectRequest(deviceUrlExtension, name, edgeNodeID, urlType)
}

func newEdgeNodeResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: createEdgeNodeResource,
		ReadContext:   readEdgeNodeResource,
		UpdateContext: updateEdgeNodeResource,
		DeleteContext: deleteEdgeNodeResource,
		Schema:        zschemas.EdgeNodeSchema,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}

}

func createEdgeNodeResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	edgeNodeID := ""
	edgeNodeName := resourcedata.GetStr(d, "name")

	log.Printf("[INFO] creating EdgeNode: %s", edgeNodeName)

	localState, err := getEdgeNodeStateForCreation(d)
	if err != nil {
		return diag.FromErr(zerrors.New(edgeNodeID, objectTypeEdgeNode, edgeNodeName, "create", err))
	}

	apiClient, ok := meta.(*client.Client)
	if !ok {
		return diag.Errorf("expect meta to be of type client.Client{} but is %T", meta)
	}
	// FIXME: why do we set a field to the client instance which is request scoped?
	apiClient.XRequestIdPrefix = "TF-edgenode-create"

	response := &models.ZsrvResponse{}
	_, err = apiClient.SendReq("POST", deviceUrlExtension, localState, response)
	if err != nil {
		return diag.FromErr(zerrors.New(edgeNodeID, objectTypeEdgeNode, edgeNodeName, "create", err))
	}

	// we must update local state with the id of the newly created object
	edgeNodeName = response.ObjectName
	edgeNodeID = response.ObjectID
	d.SetId(edgeNodeID)

	log.Printf("[INFO] EdgeNode %s created (id=%s) successfully\n", edgeNodeName, edgeNodeID)

	// fetch EdgeNode state from zedcloud api with the new object's id to validate that is has been created
	remoteState, err := fetchEdgeNodeState(apiClient, edgeNodeName, edgeNodeID)
	if err != nil {
		// no object with this id exist in the api's state
		var notFoundErr *zerrors.ObjectNotFound
		if errors.As(err, &notFoundErr) {
			// we need to remove it from local state
			state.RemoveLocal(d, objectTypeEdgeNode, edgeNodeID, edgeNodeName)
			return diag.FromErr(zerrors.New(edgeNodeID, objectTypeEdgeNode, edgeNodeName, "find", err))
		}
		// api error
		return diag.FromErr(zerrors.New(edgeNodeID, objectTypeEdgeNode, edgeNodeName, "fetch", err))
	}

	// update base image version in local state
	newBaseOsVersion := getEdgeNodeBaseOs(remoteState, resourcedata.GetStr(d, "eve_image_version"))

	// set the new image version in remote state
	if newBaseOsVersion != nil {
		localState.BaseImage = newBaseOsVersion
		log.Printf(
			"[TRACE]: update the base os image to: %s, set activate to: %t",
			*remoteState.BaseImage[0].ImageName,
			remoteState.BaseImage[0].Activate,
		)
		if err := setEdgeNodeBaseOs(apiClient, localState); err != nil {
			return diag.FromErr(zerrors.New(edgeNodeID, objectTypeEdgeNode, edgeNodeName, "update base os for", err))
		}
	}

	// activating / de-activating the device requires activate / deactivate call
	if err := edgeNodeUpdateAdminState(apiClient, d, edgeNodeID, edgeNodeName); err != nil {
		return diag.FromErr(zerrors.New(edgeNodeID, objectTypeEdgeNode, edgeNodeName, "update admin state for", err))
	}

	// sync local state with the updated remote state. this fetches the remotely computed fields and ensures, the update
	// was performed successfully
	if err := getEdgeNodeDataSource(ctx, d, meta); err != nil {
		return diag.FromErr(zerrors.New(edgeNodeID, objectTypeEdgeNode, edgeNodeName, "fetch after creating of", err))
	}
	return diags
}

func getEdgeNodeStateForCreation(d *schema.ResourceData) (*models.DeviceConfig, error) {
	// FIXME: why does the original code not set it in the local state?
	if eveImageVersion := resourcedata.GetStr(d, "eve_image_version"); eveImageVersion == "" {
		return nil, fmt.Errorf("eve_image_version must be specified.")
	}

	modelID := resourcedata.GetStrPtrOrNil(d, "model_id")
	if modelID == nil || *modelID == "" {
		return nil, fmt.Errorf("model_id must be specified for the EdgeNode.")
	}

	projectID := resourcedata.GetStrPtrOrNil(d, "project_id")
	if projectID == nil || *projectID == "" {
		return nil, fmt.Errorf("project_id must be specified for the EdgeNode.")
	}

	_, err := getAdminStateValue(d)
	if err != nil {
		return nil, err
	}

	configItems, err := getConfigItems(d)
	if err != nil {
		return nil, err
	}

	deviceLocation, err := getDeviceLocation(d)
	if err != nil {
		return nil, err
	}

	sysInterfaces, err := getSystemInterfaces(d)
	if err != nil {
		return nil, err
	}

	return &models.DeviceConfig{
		Name:        resourcedata.GetStrPtrOrNil(d, "name"),
		Description: resourcedata.GetStr(d, "description"),
		Title:       resourcedata.GetStrPtrOrNil(d, "title"),
		AssetID:     resourcedata.GetStr(d, "asset_id"),
		ClientIP:    resourcedata.GetStr(d, "client_ip"),
		ClusterID:   resourcedata.GetStr(d, "cluster_id"),
		ConfigItem:  configItems,
		DevLocation: deviceLocation,
		Interfaces:  sysInterfaces,
		ModelID:     modelID,
		Obkey:       resourcedata.GetStr(d, "onboard_key"),
		ProjectID:   projectID,
		Serialno:    resourcedata.GetStr(d, "serialno"),
		Tags:        resourcedata.GetStrMap(d, "tags"),
	}, nil
}

func getEdgeNodeStateForUpdate(remoteState *models.DeviceConfig, d *schema.ResourceData) (*models.DeviceConfig, error) {
	if eveImageVersion := resourcedata.GetStr(d, "eve_image_version"); eveImageVersion == "" {
		return nil, fmt.Errorf("eve_image_version must be specified.")
	}

	_, err := getAdminStateValue(d)
	if err != nil {
		return nil, err
	}

	configItems, err := getConfigItems(d)
	if err != nil {
		return nil, err
	}

	deviceLocation, err := getDeviceLocation(d)
	if err != nil {
		return nil, err
	}

	sysInterfaces, err := getSystemInterfaces(d)
	if err != nil {
		return nil, err
	}

	projectID := resourcedata.GetStrPtrOrNil(d, "project_id")
	// Change of Project for a device is NOT supported with Update operation.
	// Check if the configured project is different from the config we got
	// from Cloud. If different - error out.
	// If project_id is not configured, ignore the check. We cannot verify
	// the case of device created with project_id set to non-default project,
	// but changed to remove set the project_id to empty, though this also is
	// not supported. Such an update request would be rejected by ZEDCloud.
	if projectID != nil && remoteState.ProjectID != nil && *remoteState.ProjectID != *projectID {
		return nil, fmt.Errorf("project_id cannot be changed after EdgeNode was created. remote: %s, local: %s", *remoteState.ProjectID, *projectID)
	}

	remoteState.Name = resourcedata.GetStrPtrOrNil(d, "name")
	remoteState.Description = resourcedata.GetStr(d, "description")
	remoteState.Title = resourcedata.GetStrPtrOrNil(d, "title")
	remoteState.AssetID = resourcedata.GetStr(d, "asset_id")
	remoteState.ClientIP = resourcedata.GetStr(d, "client_ip")
	remoteState.ClusterID = resourcedata.GetStr(d, "cluster_id")
	remoteState.ConfigItem = configItems
	remoteState.DevLocation = deviceLocation
	remoteState.Interfaces = sysInterfaces
	remoteState.ProjectID = projectID
	remoteState.Serialno = resourcedata.GetStr(d, "serialno")
	remoteState.Tags = resourcedata.GetStrMap(d, "tags")
	return remoteState, nil
}

func getAdminStateValue(d *schema.ResourceData) (string, error) {
	strVal := resourcedata.GetStr(d, "adminstate_config")
	switch strVal {
	case "ADMIN_STATE_ACTIVE":
		return "activate", nil
	case "ADMIN_STATE_INACTIVE":
		return "deactivate", nil
	default:
		return "", fmt.Errorf(
			"adminstate_config empty, must be one of ADMIN_STATE_ACTIVE, ADMIN_STATE_INACTIVE",
		)
	}
}

func getConfigItems(d *schema.ResourceData) ([]*models.EDConfigItem, error) {
	configItems := make([]*models.EDConfigItem, 0)
	items, ok := d.GetOk("config_items")
	if !ok {
		return configItems, nil
	}
	itemsData, ok := items.(map[string]interface{})
	if !ok {
		return configItems, nil
	}
	for key, val := range itemsData {
		strVal, ok := val.(string)
		if !ok {
			return configItems, fmt.Errorf("Value (%+v) for config item (%s) must be a valid string.", val, key)
		}
		configItems = append(configItems, &models.EDConfigItem{
			Key:         key,
			StringValue: strVal,
		})
	}
	return configItems, nil
}

func getDeviceLocation(d *schema.ResourceData) (*models.GeoLocation, error) {
	locationData := resourcedata.GetList(d, "dev_location")
	if len(locationData) == 0 {
		return nil, nil
	}

	if len(locationData) > 1 {
		return nil, fmt.Errorf("expect exactly 1 location info but have %d", len(locationData))
	}

	// FIXME: can we cast to map[string]string{}?
	entry, ok := locationData[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("wrong type of device location: %T", locationData[0])
	}

	location := models.GeoLocation{}
	if val, ok := entry["city"]; ok {
		location.City = val.(string)
	}
	if val, ok := entry["country"]; ok {
		location.Country = val.(string)
	}
	if val, ok := entry["freeloc"]; ok {
		location.Freeloc = val.(string)
	}
	if val, ok := entry["hostname"]; ok {
		location.Hostname = val.(string)
	}
	if val, ok := entry["loc"]; ok {
		location.Loc = val.(string)
	}
	if val, ok := entry["org"]; ok {
		location.Org = val.(string)
	}
	if val, ok := entry["postal"]; ok {
		location.Postal = val.(string)
	}
	if val, ok := entry["region"]; ok {
		location.Region = val.(string)
	}
	if val, ok := entry["underlay_ip"]; ok {
		location.UnderlayIP = val.(string)
	}
	return &location, nil
}

func getSystemInterfaces(d *schema.ResourceData) ([]*models.SysInterface, error) {
	interfaceData := resourcedata.GetList(d, "interface")
	interfaces := make([]*models.SysInterface, 0)
	for _, val := range interfaceData {
		entry := val.(map[string]interface{})
		intf := models.SysInterface{}
		if val, ok := entry["cost"]; ok {
			intf.Cost = int64(val.(int))
		}
		if val, ok := entry["intf_usage"]; ok {
			usage := models.AdapterUsage(val.(string))
			intf.IntfUsage = &usage
		}
		if val, ok := entry["intfname"]; ok {
			intf.Intfname = val.(string)
		} else {
			return interfaces, fmt.Errorf("intfname must be specified for Interfaces")
		}
		if val, ok := entry["ipaddr"]; ok {
			intf.Ipaddr = val.(string)
		}
		if val, ok := entry["macaddr"]; ok {
			intf.Macaddr = val.(string)
		}
		if val, ok := entry["netname"]; ok {
			intf.Netname = val.(string)
		}
		intf.Tags = resourcedata.GetStrMap(entry, "tags")
		interfaces = append(interfaces, &intf)
	}
	return interfaces, nil
}

// BaseImage is supposed to have only one entry. If there are multiple, reset the BaseOS config.
func getEdgeNodeBaseOs(remoteState *models.DeviceConfig, localEVEVersion string) []*models.BaseOSImage {
	hasSingleBaseImage := remoteState.BaseImage != nil && len(remoteState.BaseImage) == 1
	matchesEveVersion := *remoteState.BaseImage[0].ImageName == localEVEVersion
	if hasSingleBaseImage && matchesEveVersion && remoteState.BaseImage[0].Activate {
		// no update not required
		return nil
	}
	return []*models.BaseOSImage{
		{
			ImageName: &localEVEVersion,
			Version:   &localEVEVersion,
			Activate:  true,
		},
	}
}

func setEdgeNodeBaseOs(client *client.Client, localState *models.DeviceConfig) error {
	// BaseOs update is special case - publish config followed by apply
	_, err := edgeNodeSendPutReq(client, localState, "publish")
	if err != nil {
		return err
	}
	rsp, err := edgeNodeSendPutReq(client, localState, "apply")
	// ignore status code HTTP_STATUS_CONFLICT
	if err != nil && rsp.StatusCode != http.StatusConflict {
		return err
	}
	return nil
}

func edgeNodeSendPutReq(client *client.Client, localState *models.DeviceConfig, reqType string) (*http.Response, error) {
	client.XRequestIdPrefix = "TF-edgenode-" + reqType
	urlExtension := getEdgeNodeUrl(*localState.Name, localState.ID, reqType)
	rspData := &models.ZsrvResponse{}
	return client.SendReq("PUT", urlExtension, localState, rspData)
}

func edgeNodeUpdateAdminState(apiClient *client.Client, d *schema.ResourceData, edgeNodeID, edgeNodeName string) error {
	remoteState, err := fetchEdgeNodeState(apiClient, "", edgeNodeID)
	if err != nil {
		// no object with this id exist in the api's state
		var notFoundErr *zerrors.ObjectNotFound
		if errors.As(err, &notFoundErr) {
			// we need to remove it from local state
			state.RemoveLocal(d, objectTypeEdgeNode, edgeNodeID, edgeNodeName)
			return zerrors.New(edgeNodeID, objectTypeEdgeNode, edgeNodeName, "find", err)
		}
		// api error
		return zerrors.New(edgeNodeID, objectTypeEdgeNode, edgeNodeName, "fetch", err)
	}

	var action string
	action, err = setAdminState(remoteState, d)
	if err != nil {
		return zerrors.New(edgeNodeID, objectTypeEdgeNode, edgeNodeName, "set local admin state for", err)
	}
	if action == "" {
		return nil
	}

	_, err = edgeNodeSendPutReq(apiClient, remoteState, action)
	if err != nil {
		return zerrors.New(edgeNodeID, objectTypeEdgeNode, edgeNodeName, "set remote admin state for", err)
	}
	return nil
}

func readEdgeNodeResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readEdgeNodeDataSource(ctx, d, meta)
}

func adminStatePtr(strVal string) *models.AdminState {
	val := models.AdminState(strVal)
	return &val
}

func adapterUsagePtr(strVal string) *models.AdapterUsage {
	val := models.AdapterUsage(strVal)
	return &val
}

func setAdminState(remoteState *models.DeviceConfig, d *schema.ResourceData) (string, error) {
	action, err := getAdminStateValue(d)
	if err != nil {
		return "", err
	}

	strVal := resourcedata.GetStr(d, "adminstate_config")
	localAdminState := models.AdminState(strVal)

	remoteStateNotNil := remoteState.AdminState != nil
	localEqualsRemote := localAdminState == *remoteState.AdminState

	// admin state same as configured value - no action needed
	if remoteStateNotNil && localEqualsRemote {
		return "", nil
	}

	if remoteStateNotNil {
		// adminstate_config is ACTIVE and the device is already in registered state
		if strVal == "ADMIN_STATE_ACTIVE" && *remoteState.AdminState == models.AdminStateADMINSTATEREGISTERED {
			return "", nil
		}
	}

	remoteState.AdminState = &localAdminState
	return action, nil
}

func updateEdgeNodeResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	edgeNodeID := resourcedata.GetStr(d, "id")
	edgeNodeName := resourcedata.GetStr(d, "name")

	apiClient, ok := meta.(*client.Client)
	if !ok {
		return diag.Errorf("expect meta to be of type client.Client{} but is %T", meta)
	}

	// fetch EdgeNode state from zedcloud api with the new object's edgeNodeID to validate that is has been created
	remoteState, err := fetchEdgeNodeState(apiClient, edgeNodeName, edgeNodeID)
	if err != nil {
		// no object with this edgeNodeID exist in the api's state
		var notFoundErr *zerrors.ObjectNotFound
		if errors.As(err, &notFoundErr) {
			// we need to remove it from local state
			state.RemoveLocal(d, objectTypeEdgeNode, edgeNodeID, edgeNodeName)
			return diag.FromErr(zerrors.New(edgeNodeID, objectTypeEdgeNode, edgeNodeName, "find", err))
		}
		// api error
		return diag.FromErr(zerrors.New(edgeNodeID, objectTypeEdgeNode, edgeNodeName, "fetch", err))
	}

	// we cannot update if ID or name have changed
	err = state.CheckInvalidAttrForUpdate(d, *remoteState.Name, remoteState.ID)
	if err != nil {
		return diag.FromErr(zerrors.New(
			edgeNodeID,
			objectTypeEdgeNode,
			edgeNodeName,
			"update, invalid attribute for",
			err,
		))
	}

	// we create the config from remote state and the updates in the local state, that we want to use to
	// update the remote state
	remoteStateForUpdate, err := getEdgeNodeStateForUpdate(remoteState, d)
	if err != nil {
		return diag.FromErr(zerrors.New(edgeNodeID, objectTypeEdgeNode, edgeNodeName, "update", err))
	}

	// update the remote state
	if _, err := edgeNodeSendPutReq(apiClient, remoteStateForUpdate, "update"); err != nil {
		return diag.FromErr(zerrors.New(edgeNodeID, objectTypeEdgeNode, edgeNodeName, "update", err))
	}

	// FIXME: why don't we update image version in the first update request?
	// update base image version in local state
	newBaseOsVersion := getEdgeNodeBaseOs(remoteState, resourcedata.GetStr(d, "eve_image_version"))

	// set the new image version in remote state
	if newBaseOsVersion != nil {
		remoteStateForUpdate.BaseImage = newBaseOsVersion
		log.Printf(
			"[TRACE]: update the base os image to: %s, set activate to: %t",
			*remoteState.BaseImage[0].ImageName,
			remoteState.BaseImage[0].Activate,
		)
		if err := setEdgeNodeBaseOs(apiClient, remoteStateForUpdate); err != nil {
			return diag.FromErr(zerrors.New(edgeNodeID, objectTypeEdgeNode, edgeNodeName, "update base os for", err))
		}
	}

	// activating / de-activating the device requires activate / deactivate call
	if err := edgeNodeUpdateAdminState(apiClient, d, edgeNodeID, edgeNodeName); err != nil {
		return diag.FromErr(zerrors.New(edgeNodeID, objectTypeEdgeNode, edgeNodeName, "update admin state for", err))
	}

	// sync local state with the updated remote state. this fetches the remotely computed fields and ensures, the update
	// was performed successfully
	if err := getEdgeNodeDataSource(ctx, d, meta); err != nil {
		return diag.FromErr(zerrors.New(edgeNodeID, objectTypeEdgeNode, edgeNodeName, "fetch after creating of", err))
	}

	log.Printf("[INFO] %s %s (id=%s) update successful", objectTypeEdgeNode, edgeNodeName, remoteStateForUpdate.ID)
	return diags
}

// // Delete the Resource Group
func deleteEdgeNodeResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	edgeNodeName := resourcedata.GetStr(d, "name")
	edgeNodeID := resourcedata.GetStr(d, "id")

	apiClient, ok := meta.(*client.Client)
	if !ok {
		return diag.Errorf("expect meta to be of type client.Client{} but is %T", meta)
	}
	// FIXME: do not set request scoped field on object
	apiClient.XRequestIdPrefix = "TF-edgenode-delete"

	urlExtension := getEdgeNodeUrl(edgeNodeName, edgeNodeID, "delete")

	rspData := &models.ZsrvResponse{}
	_, err := apiClient.SendReq("DELETE", urlExtension, nil, rspData)
	if err != nil {
		return diag.FromErr(zerrors.New(edgeNodeID, objectTypeEdgeNode, edgeNodeName, "delete", err))
	}

	log.Printf("[INFO] EdgeNode delete successfully")
	return diags
}
