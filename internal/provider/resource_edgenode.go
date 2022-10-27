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
	deviceUrlExtension   = "devices"
	resourceTypeEdgeNode = "EdgeNode"
)

func getEdgeNodeUrl(name, id, urlType string) string {
	return zedcloudapi.UrlForObjectRequest(deviceUrlExtension, name, id, urlType)
}

func newEdgeNodeResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: createEdgeNodeResource,
		ReadContext:   readEdgeNodeResource,
		// UpdateContext: updateEdgeNodeResource,
		// DeleteContext: deleteEdgeNodeResource,
		Schema: zschemas.EdgeNodeSchema,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}

}

func createEdgeNodeResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	id := resourcedata.GetStr(d, "id") // FIXME: empty for create?
	edgeNodeName := resourcedata.GetStr(d, "name")

	log.Printf("[INFO] Creating EdgeNode: %s", edgeNodeName)

	localState, err := getEdgeNodeState(d, true)
	if err != nil {
		return diag.FromErr(zerrors.New(id, resourceTypeEdgeNode, edgeNodeName, "create", err))
	}

	client, ok := meta.(*client.Client)
	if !ok {
		return diag.Errorf("expect meta to be of type client.Client{} but is %T", meta)
	}
	// FIXME: why do we set a field to the client instance which is request scoped?
	client.XRequestIdPrefix = "TF-edgenode-create"

	response := &models.ZsrvResponse{}
	_, err = client.SendReq("POST", deviceUrlExtension, localState, response)
	if err != nil {
		return diag.FromErr(zerrors.New(id, resourceTypeEdgeNode, edgeNodeName, "create", err))
	}

	// we must update local state with the id of the newly created object
	newEdgeNodeID := response.ObjectID
	d.SetId(newEdgeNodeID)

	log.Printf("EdgeNode %s (ID: %s) Successfully created\n", response.ObjectName, newEdgeNodeID)

	// fetch EdgeNode state from zedcloud api with the new object's id to validate that is has been created
	remoteState, err := fetchEdgeNodeState(edgeNodeName, newEdgeNodeID, meta)
	if err != nil {
		// no object with this id exist in the api's state
		var notFoundErr *zerrors.ObjectNotFound
		if errors.As(err, &notFoundErr) {
			// we need to remove it from local state
			state.RemoveLocal(d, resourceTypeEdgeNode, id, edgeNodeName)
			return diag.FromErr(zerrors.New(id, resourceTypeEdgeNode, edgeNodeName, "find", err))
		}
		// api error
		return diag.FromErr(zerrors.New(id, resourceTypeEdgeNode, edgeNodeName, "fetch", err))
	}

	// update base image version in local state
	newBaseOsVersion := getEdgeNodeBaseOs(remoteState, resourcedata.GetStr(d, "eve_image_version"))
	localState.BaseImage = newBaseOsVersion

	// set the new image version in remote state
	if newBaseOsVersion != nil {
		log.Printf(
			"[TRACE]: Update BaseOsImage. ImageName: %s, Activate: %t",
			*remoteState.BaseImage[0].ImageName,
			remoteState.BaseImage[0].Activate,
		)
		if err := setEdgeNodeBaseOs(client, localState); err != nil {
			return diag.FromErr(zerrors.New(id, resourceTypeEdgeNode, edgeNodeName, "update base os for", err))
		}
	}

	// // Activating / De-Activating the device requires activate / deactivate call
	// err = edgeNodeUpdateAdminState(client, d, id)
	// if err != nil {
	// 	return diag.Errorf("%s Failed to update Admin state. Err: %s",
	// 		errMsgPrefix, err.Error())
	// }

	// // Get Edge node config and publish the latest version. This is mainly to
	// // published the computed fields. Object rev. changes for every update
	// err = getEdgeNodeAndPublishData(client, d, name, id)
	// if err != nil {
	// 	log.Printf("***[ERROR]- Failed to get EdgeNode: %s (ID: %s) after creating it. Err: %s", name, id, err.Error())
	// }
	return diags
}

func getEdgeNodeState(d *schema.ResourceData, create bool) (*models.DeviceConfig, error) {
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

	if eveImageVersion := resourcedata.GetStr(d, "eve_image_version"); eveImageVersion == "" {
		return nil, fmt.Errorf("eve_image_version must be specified.")
	}

	sysInterfaces, err := getSystemInterfaces(d)
	if err != nil {
		return nil, err
	}

	var modelID *string
	var obkey string
	projectID := resourcedata.GetStrPtrOrNil(d, "project_id")

	// FIXME: move this to create remove crate flag?
	if create {
		// model id and project id will be set only during create.
		modelID = resourcedata.GetStrPtrOrNil(d, "model_id")
		if modelID == nil || *modelID == "" {
			return nil, fmt.Errorf("model_id must be specified for the EdgeNode.")
		}

		if projectID == nil || *projectID == "" {
			return nil, fmt.Errorf("project_id must be specified for the EdgeNode.")
		}
		obkey = resourcedata.GetStr(d, "onboard_key")
	}

	// FIXME: this needs to go to update only!
	// else {
	// 	// Change of Project for a device is NOT supported with Update operation.
	// 	// Check if the configured project is different from the config we got
	// 	// from Cloud. If different - error out.
	// 	// If project_id is not configured, ignore the check. We cannot verify
	// 	// the case of device created with project_id set to non-default project,
	// 	// but changed to remove set the project_id to empty, though this also is
	// 	// not supported. Such an update request would be rejected by ZEDCloud.
	// 	if projectID != nil && state.ProjectID != nil &&
	// 		*state.ProjectID != *projectID {
	// 		// Update. Project cannot be changed
	// 		return fmt.Errorf("project_id cannot be changed after EdgeNode is "+
	// 			"created. Current: %s, New: %s", *state.ProjectID, *projectID)
	// 	}
	// }

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
		Obkey:       obkey,
		ProjectID:   projectID,
		Serialno:    resourcedata.GetStr(d, "serialno"),
		Tags:        resourcedata.GetStrMap(d, "tags"),
	}, nil
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
			return configItems, fmt.Errorf("Value (%+v) for Config Item (%s) must be a valid string.", val, key)
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

// BaseImage is supposed to have only one entry. if there are multiple, reset the BaseOS config
func getEdgeNodeBaseOs(remoteState *models.DeviceConfig, eveImageVersion string) []*models.BaseOSImage {
	hasSingleBaseImage := remoteState.BaseImage != nil && len(remoteState.BaseImage) == 1
	matchesEveVersion := *remoteState.BaseImage[0].ImageName == eveImageVersion
	if hasSingleBaseImage && matchesEveVersion && remoteState.BaseImage[0].Activate {
		// no update not required
		return nil
	}
	return []*models.BaseOSImage{
		{
			ImageName: &eveImageVersion,
			Version:   &eveImageVersion,
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

func readEdgeNodeResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readEdgeNodeDataSource(ctx, d, meta)
}

// func adminStatePtr(strVal string) *models.AdminState {
// 	val := models.AdminState(strVal)
// 	return &val
// }

// func edgeNodeUpdateAdminState(client *zedcloudapi.Client, d *schema.ResourceData, id string) error {
// 	cfg, err := getEdgeNodeConfig(client, "", id)
// 	if err != nil {
// 		return fmt.Errorf("Failed to find Edge Node to set Admin state. err: %s",
// 			err.Error())
// 	}
// 	var action string
// 	action, err = setAdminState(cfg, d)
// 	if err != nil {
// 		return fmt.Errorf("Failed to set Admin state in Config. err: %s", err.Error())
// 	}
// 	if action == "" {
// 		return nil
// 	}
// 	_, err = edgeNodeSendPutReq(client, cfg, action)
// 	if err != nil {
// 		return fmt.Errorf("Failed to set Admin state. Err: %s", err.Error())
// 	}
// 	return nil
// }

// func adapterUsagePtr(strVal string) *models.AdapterUsage {
// 	val := models.AdapterUsage(strVal)
// 	return &val
// }

// func setAdminState(cfg *models.DeviceConfig, d *schema.ResourceData) (string, error) {
// 	action, err := checkAdminStateValue(d)
// 	if err != nil {
// 		return "", err
// 	}
// 	strVal := resourcedata.GetStr(d, "adminstate_config")
// 	adminstate := models.AdminState(strVal)
// 	if cfg.AdminState != nil {
// 		if adminstate == *cfg.AdminState {
// 			// Admin State same as configured value - no action needed.
// 			return "", nil
// 		}
// 		// If adminstate_config is ACTIVE and the device is already in
// 		// Registered state. No action needed.
// 		if strVal == "ADMIN_STATE_ACTIVE" &&
// 			*cfg.AdminState == models.AdminStateADMINSTATEREGISTERED {
// 			return "", nil
// 		}
// 	}
// 	cfg.AdminState = &adminstate
// 	return action, nil
// }

// Update the Resource Group
// func updateEdgeNodeResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics

// 	client, ok := meta.(*client.Client)
// 	if !ok {
// 		return diag.Errorf("expect meta to be of type client.Client{} but is %T", meta)
// 	}
// 	id := resourcedata.GetStr(d, "id")
// 	name := resourcedata.GetStr(d, "name")
// 	errMsgPrefix := getErrMsgPrefix(name, id, "Edge Node", "Update")
// 	cfg, err := getEdgeNodeConfig(client, name, id)
// 	if err != nil {
// 		return diag.Errorf("%s Failed to find Edge Node. err: %s",
// 			errMsgPrefix, err.Error())
// 	}
// 	err = checkInvalidAttrForUpdate(d, *cfg.Name, cfg.ID)
// 	if err != nil {
// 		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
// 	}
// 	if cfg.ID != id {
// 		return diag.Errorf("%s ID of edge node changed. Expected: %s, actual: %s. "+
// 			"Object in zedcontrol is not same as expected by Terraform.",
// 			errMsgPrefix, id, cfg.ID)
// 	}
// 	err = rdDeviceConfig(cfg, d, false)
// 	if err != nil {
// 		return diag.Errorf("%s Error: %s", errMsgPrefix, err.Error())
// 	}
// 	_, err = edgeNodeSendPutReq(client, cfg, "update")
// 	if err != nil {
// 		return diag.Errorf("%s Update request failed. Err: %s",
// 			errMsgPrefix, err.Error())
// 	}
// 	err = edgeNodeUpdateBaseOs(client, cfg, resourcedata.GetStr(d, "eve_image_version"))
// 	if err != nil {
// 		return diag.Errorf("%s %s", errMsgPrefix, err.Error())
// 	}
// 	err = edgeNodeUpdateAdminState(client, d, id)
// 	if err != nil {
// 		return diag.Errorf("%s %s", errMsgPrefix, err.Error())
// 	}
// 	err = getEdgeNodeAndPublishData(client, d, name, id)
// 	if err != nil {
// 		return diag.Errorf("%s", err.Error())
// 	}
// 	log.Printf("[INFO] EdgeNode %s (ID: %s) Update Successful.", name, cfg.ID)
// 	return diags
// }

// // Delete the Resource Group
// func deleteEdgeNodeResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics

// 	client, ok := meta.(*client.Client)
// 	if !ok {
// 		return diag.Errorf("expect meta to be of type client.Client{} but is %T", meta)
// 	}
// 	name := resourcedata.GetStr(d, "name")
// 	id := resourcedata.GetStr(d, "id")
// 	errMsgPrefix := getErrMsgPrefix(name, id, "Edge Node", "Delete")
// 	client.XRequestIdPrefix = "TF-edgenode-delete"
// 	urlExtension := getEdgeNodeUrl(name, id, "delete")
// 	rspData := &models.ZsrvResponse{}
// 	_, err := client.SendReq("DELETE", urlExtension, nil, rspData)
// 	if err != nil {
// 		return diag.Errorf("%s. Err: %s", errMsgPrefix, err.Error())
// 	}
// 	log.Printf("[INFO] EdgeNode Delete Successful.")
// 	return diags
// }
