// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/internal/resourcedata"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	"github.com/zededa/terraform-provider-zedcloud/utils"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
)

var VolumeInstanceDataSourceSchema = &schema.Resource{
	ReadContext: readDataSourceVolumeInstance,
	Schema:      zschemas.VolumeInstanceSchema,
	Description: "Schema for data source zedcloud_volume_instance. Must specify " +
		"id or name",
}

// The schema for a resource group data source
func getVolumeInstanceDataSourceSchema() *schema.Resource {
	return VolumeInstanceDataSourceSchema
}

func getVolumeInstance(client *zedcloudapi.Client,
	name, id string) (*swagger_models.VolInstConfig, error, *http.Response) {
	rspData := &swagger_models.VolInstConfig{}
	httpResp, err := client.GetObj(volumeInstanceUrlExtension, name, id, false, rspData)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] FAILED to get VolumeInstance %s ( id: %s). Err: %s",
			name, id, err.Error()), httpResp
	}
	return rspData, nil, httpResp
}

func flattenVolInstConfig(cfg *swagger_models.VolInstConfig) map[string]interface{} {
	if cfg == nil {
		return map[string]interface{}{}
	}

	// Publish Compute Only Fields by Default.
	data := map[string]interface{}{
		"id":         cfg.ID,
		"implicit":   cfg.Implicit,
		"project_id": cfg.ProjectID,
		"purge":      flattenZedCloudOpsCmd(cfg.Purge),
		"revision":   utils.FlattenObjectRevision(cfg.Revision),
	}
	data["accessmode"] = utils.PtrValStr(cfg.Accessmode)
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
	data["type"] = utils.PtrValStr(cfg.Type)
	utils.CheckIfAllKeysExist(zschemas.VolumeInstanceSchema, data)
	return data
}

func getVolumeInstanceAndPublishData(client *zedcloudapi.Client, d *schema.ResourceData,
	name, id string) error {
	cfg, err, httpRsp := getVolumeInstance(client, name, id)
	if err != nil {
		err = fmt.Errorf("[ERROR] Volume Instance %s (id: %s) not found. Err: %s",
			name, id, err.Error())
		if httpRsp != nil && zedcloudapi.IsObjectNotFound(httpRsp) {
			log.Printf("PROVIDER:  REMOVED Volume Instance %s (id: %s) from State", id, name)
			schema.RemoveFromState(d, nil)
			return nil
		}
		return err
	}
	if cfg.Implicit {
		return fmt.Errorf("[ERROR] VolumeInstance %s (id: %s) is an Implicit Volume. "+
			"Implicit Volumes cannot be used as resources or data sources.",
			name, id)
	}
	marshalData(d, flattenVolInstConfig(cfg))
	return nil
}

// Read the Resource Group
func readVolumeInstance(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	id := resourcedata.GetStr(d, "id")
	name := resourcedata.GetStr(d, "name")

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
