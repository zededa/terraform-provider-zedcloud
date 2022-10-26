// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package datasources

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	zedcloudapi "github.com/zededa/zedcloud-api"
	models "github.com/zededa/zedcloud-api/swagger_models"
)

var VolumeInstanceDataSourceSchema = &schema.Resource{
	ReadContext: readDataSourceVolumeInstance,
	Schema:      zschemas.VolumeInstanceSchema,
	Description: "Schema for data source zedcloud_volume_instance. Must specify id or name",
}

// The schema for a resource group data source
func getVolumeInstanceDataSourceSchema() *schema.Resource {
	return VolumeInstanceDataSourceSchema
}

func getVolumeInstance(client *zedcloudapi.Client, name, id string) (*models.VolInstConfig, error, *http.Response) {
	rspData := &models.VolInstConfig{}
	httpResp, err := client.GetObj(volumeInstanceUrlExtension, name, id, false, rspData)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] FAILED to get VolumeInstance %s ( id: %s). Err: %w", name, id, err), httpResp
	}
	return rspData, nil, httpResp
}

func flattenVolInstConfig(cfg *models.VolInstConfig) map[string]interface{} {
	if cfg == nil {
		return map[string]interface{}{}
	}

	// Publish Compute Only Fields by Default.
	data := map[string]interface{}{
		"id":         cfg.ID,
		"implicit":   cfg.Implicit,
		"project_id": cfg.ProjectID,
		"purge":      flattenZedCloudOpsCmd(cfg.Purge),
		"revision":   flattenObjectRevision(cfg.Revision),
	}
	data["accessmode"] = strPtrVal(cfg.Accessmode)
	data["cleartext"] = cfg.Cleartext
	data["content_tree_id"] = cfg.ContentTreeID
	data["description"] = cfg.Description
	data["device_id"] = cfg.DeviceID
	data["image"] = cfg.Image
	data["label"] = cfg.Label
	data["multiattach"] = cfg.Multiattach
	data["name"] = cfg.Name
	data["size_bytes"] = cfg.SizeBytes
	data["title"] = cfg.Title
	data["type"] = strPtrVal(cfg.Type)
	checkIfAllKeysExist(zschemas.VolumeInstanceSchema, data)
	return data
}

func getVolumeInstanceAndPublishData(client *zedcloudapi.Client, d *schema.ResourceData, name, id string) error {
	cfg, err, httpRsp := getVolumeInstance(client, name, id)
	if err != nil {
		return fmt.Errorf("could not get volume instance %s (id=%s): %w", name, id, err)
	}

	if zedcloudapi.IsObjectNotFound(httpRsp) {
		log.Printf("could not get volume instance %s (id=%s): not found", name, id)
		return schema.RemoveFromState(d, nil)
	}

	if cfg.Implicit {
		return fmt.Errorf("volume instance %s (id=%s) is implicit and cannot be used as resources or data sources", name, id)
	}

	setLocalState(d, flattenVolInstConfig(cfg))
	return nil
}

// Read the Resource Group
func readVolumeInstance(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	id := getStr(d, "id")
	name := getStr(d, "name")

	if client == nil {
		return diag.Errorf("nil Client.")
	}
	log.Printf("PROVIDER:  readVolumeInstanceResource - id: %s, name: %s", id, name)
	if (id == "") && (name == "") {
		return diag.Errorf("The arguments \"id\" or \"name\" are required, but no definition was found.")
	}
	err := getVolumeInstanceAndPublishData(client, d, name, id)
	if err != nil {
		return diag.Errorf("%s", err.Error())
	}
	return diags
}

func readDataSourceVolumeInstance(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readVolumeInstance(ctx, d, meta)
}
