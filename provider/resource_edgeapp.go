// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/schemas"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
	"io/ioutil"
	"log"
)

var edgeAppUrlExtension = "apps"

var EdgeAppResourceSchema = &schema.Resource{
	CreateContext: createEdgeAppResource,
	ReadContext:   readEdgeApp,
	UpdateContext: updateEdgeAppResource,
	DeleteContext: deleteEdgeAppResource,
	Schema:        schemas.EdgeAppSchema,
	Importer: &schema.ResourceImporter{
		StateContext: schema.ImportStatePassthroughContext,
	},
}

// The schema for a resource group data source
func getEdgeAppResourceSchema() *schema.Resource {
	return EdgeAppResourceSchema
}

func getEdgeAppUrl(name, id, urlType string) string {
	return zedcloudapi.UrlForObjectRequest(edgeAppUrlExtension, name, id, urlType)
}

func rdUpdateAppCfg(cfg *swagger_models.App, d *schema.ResourceData) error {
	cfg.Title = rdEntryStrPtrOrNil(d, "title")
	cfg.Description = rdEntryStr(d, "description")
	cfg.Cpus = rdEntryInt64(d, "cpus")
	cfg.Drives = rdEntryInt64(d, "drives")
	manifest_file := rdEntryStr(d, "manifest_file")
	if manifest_file == "" {
		return fmt.Errorf("manifest_file needs to be specified to create EdgeApp")
	}
	manifestStr, err := ioutil.ReadFile(manifest_file)
	if err != nil {
		return fmt.Errorf("Failed to read manifest file %s. err: %v",
			manifest_file, err)
	}
	manifest := &swagger_models.VMManifest{}
	err = json.Unmarshal([]byte(manifestStr), manifest)
	if err != nil {
		return fmt.Errorf("Failed to UnMarshaling Manifest file. Err: %v", err)
	}
	cfg.ManifestJSON = manifest
	cfg.Networks = rdEntryInt64(d, "networks")
	cfg.Storage = rdEntryInt64(d, "storage")
	cfg.UserDefinedVersion = rdEntryStr(d, "user_defined_version")
	return nil
}

// Create the Resource Group
func createEdgeAppResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	name := rdEntryStr(d, "name")
	id := rdEntryStr(d, "id")
	errMsgPrefix := fmt.Sprintf("[ERROR] App Instance %s (id: %s) Create Failed.",
		name, id)
	if client == nil {
		return diag.Errorf("%s nil Client", errMsgPrefix)
	}
	cfg := &swagger_models.App{
		Name: &name,
	}
	err := rdUpdateAppCfg(cfg, d)
	if err != nil {
		return diag.Errorf("%s %s", errMsgPrefix, err.Error())
	}
	log.Printf("[INFO] Creating EdgeApp: %s", name)
	client.XRequestIdPrefix = "TF-edgeapp-create"
	rspData := &swagger_models.ZsrvResponse{}
	_, err = client.SendReq("POST", edgeAppUrlExtension, cfg, rspData)
	if err != nil {
		return diag.Errorf("%s Err: %s", errMsgPrefix, err.Error())
	}
	log.Printf("Edge App %s (ID: %s) Successfully created\n",
		rspData.ObjectName, rspData.ObjectID)
	d.SetId(rspData.ObjectID)
	return diags
}

// Update the Resource Group
func updateEdgeAppResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	name := rdEntryStr(d, "name")
	id := rdEntryStr(d, "id")
	errMsgPrefix := getErrMsgPrefix("Edge App", name, id, "Update")
	if client == nil {
		return diag.Errorf("%s nil Client", errMsgPrefix)
	}
	cfg, err := getEdgeApp(client, name, id)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	err = rdUpdateAppCfg(cfg, d)
	if err != nil {
		return diag.Errorf("%s %s", errMsgPrefix, err.Error())
	}
	log.Printf("[INFO] Updating EdgeApp: %s (ID: %s)", name, cfg.ID)
	client.XRequestIdPrefix = "TF-edgeapp-update"
	urlExtension := getEdgeAppUrl(name, id, "update")
	rspData := &swagger_models.ZsrvResponse{}
	_, err = client.SendReq("PUT", urlExtension, cfg, rspData)
	if err != nil {
		return diag.Errorf("%s Request Failed. err: %s", errMsgPrefix, err.Error())
	}
	return diags
}

// Delete the Resource Group
func deleteEdgeAppResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := (meta.(Client)).Client
	name := rdEntryStr(d, "name")
	id := rdEntryStr(d, "id")
	errMsgPrefix := getErrMsgPrefix("Edge App", name, id, "Delete")
	cfg, err := getEdgeApp(client, name, id)
	if err != nil {
		return diag.Errorf("%s Failed to get EdgeApp. err: %s", errMsgPrefix, err.Error())
	}
	if cfg == nil {
		log.Printf("%s Unexpected Error. nil config", errMsgPrefix)
		return diags
	}
	client.XRequestIdPrefix = "TF-edgeapp-delete"
	urlExtension := getEdgeAppUrl(name, id, "delete")
	rspData := &swagger_models.ZsrvResponse{}
	_, err = client.SendReq("DELETE", urlExtension, nil, rspData)
	if err != nil {
		return diag.Errorf("%s. Request Failed. err: %s", errMsgPrefix, err.Error())
	}
	log.Printf("[INFO] Edge App %s(id:%s) Delete Successful.", name, cfg.ID)
	return diags
}
