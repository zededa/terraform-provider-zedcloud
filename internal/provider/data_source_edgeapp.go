// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/internal/client"
	"github.com/zededa/terraform-provider-zedcloud/internal/resourcedata"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	"github.com/zededa/terraform-provider-zedcloud/utils"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
)

var EdgeAppDataSourceSchema = &schema.Resource{
	ReadContext: readDataSourceEdgeApp,
	Schema:      zschemas.EdgeAppSchema,
	Description: "Schema for data source zedcloud_edgeapp. Must specify id or name",
}

// The schema for a resource group data source
func getEdgeAppDataSourceSchema() *schema.Resource {
	return EdgeAppDataSourceSchema
}

func getEdgeApp(client *zedcloudapi.Client, name, id string) (*swagger_models.App, error, *http.Response) {
	rspData := &swagger_models.App{}
	httpResp, err := client.GetObj(edgeAppUrlExtension, name, id, false, rspData)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] FAILED to get EdgeApp %s ( id: %s). Err: %s",
			name, id, err.Error()), httpResp
	}
	return rspData, nil, httpResp
}

func flattenAppResources(cfgList []*swagger_models.Resource) []interface{} {
	entryList := make([]interface{}, 0)
	for _, cfg := range cfgList {
		entryList = append(entryList, map[string]interface{}{
			"name":  cfg.Name,
			"value": cfg.Value,
		})
	}
	return entryList
}

func flattenAppAuthor(cfg *swagger_models.Author) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"company": cfg.Company,
		"email":   cfg.Email,
		"user":    cfg.User,
		"website": cfg.Website,
	}}
}

func flattenAppModule(cfg *swagger_models.ModuleDetail) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"environment": utils.FlattenStringMap(cfg.Environment),
		"module_type": utils.PtrValStr(cfg.ModuleType),
		"routes":      utils.FlattenStringMap(cfg.Routes),
		"twin_detail": cfg.TwinDetail,
	}}
}

func flattenAppInterfaceAclMatches(cfgList []*swagger_models.Match) []interface{} {
	entryList := make([]interface{}, 0)
	for _, cfg := range cfgList {
		entryList = append(entryList, map[string]interface{}{
			"type":  cfg.Type,
			"value": cfg.Value,
		})
	}
	return entryList
}

func flattenAppInterfaceAclActionMapParams(
	cfg *swagger_models.MapParams) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"port": cfg.AppPort,
	}}
}

func flattenAppInterfaceAclActionLimitParams(
	cfg *swagger_models.LimitParams) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"limitburst": cfg.Limitburst,
		"limitrate":  cfg.Limitrate,
		"limitunit":  cfg.Limitunit,
	}}
}
func flattenAppInterfaceAclActions(cfgList []*swagger_models.ACLAction) []interface{} {
	entryList := make([]interface{}, 0)
	for _, cfg := range cfgList {
		entryList = append(entryList, map[string]interface{}{
			"drop":        cfg.Drop,
			"limit":       cfg.Limit,
			"limit_param": flattenAppInterfaceAclActionLimitParams(cfg.LimitValue),
			"portmap":     cfg.Portmap,
			"mapparam":    flattenAppInterfaceAclActionMapParams(cfg.Portmapto),
		})
	}
	return entryList
}

func flattenAppInterfaceAcls(cfgList []*swagger_models.ACL) []interface{} {
	entryList := make([]interface{}, 0)
	for _, cfg := range cfgList {
		entryList = append(entryList, map[string]interface{}{
			"action": flattenAppInterfaceAclActions(cfg.Actions),
			"match":  flattenAppInterfaceAclMatches(cfg.Matches),
			"name":   cfg.Name,
		})
	}
	return entryList
}

func flattenAppInterfaces(cfgList []*swagger_models.Interface) []interface{} {
	entryList := make([]interface{}, 0)
	for _, cfg := range cfgList {
		entryList = append(entryList, map[string]interface{}{
			"acl":          flattenAppInterfaceAcls(cfg.Acls),
			"directattach": cfg.Directattach,
			"name":         cfg.Name,
			"optional":     cfg.Optional,
			"privateip":    cfg.Privateip,
			"type":         cfg.Type,
		})
	}
	return entryList
}

func flattenAppImageParams(cfgList []*swagger_models.Param) []interface{} {
	entryList := make([]interface{}, 0)
	for _, cfg := range cfgList {
		entryList = append(entryList, map[string]interface{}{
			"name":  cfg.Name,
			"value": cfg.Value,
		})
	}
	return entryList
}

func appImageMaxSize(cfg *swagger_models.VMManifestImage) int {
	maxsizeStr := cfg.Maxsize
	if maxsizeStr == "" {
		maxsizeStr = "0"
	}
	maxsize, err := strconv.Atoi(maxsizeStr)
	if err != nil {
		log.Printf("[ERROR] Error in flattening AppImage %s - Invalid maxsize (%s)",
			cfg.Imagename, cfg.Maxsize)
	}
	return maxsize
}

func flattenAppImages(cfgList []*swagger_models.VMManifestImage) []interface{} {
	entryList := make([]interface{}, 0)
	for _, cfg := range cfgList {

		entryList = append(entryList, map[string]interface{}{
			"cleartext":   cfg.Cleartext,
			"drvtype":     cfg.Drvtype,
			"ignorepurge": cfg.Ignorepurge,
			"imagename":   cfg.Imagename,
			"maxsize":     appImageMaxSize(cfg),
			"mountpath":   cfg.Mountpath,
			"param":       flattenAppImageParams(cfg.Params),
			"preserve":    cfg.Preserve,
			"readonly":    cfg.Readonly,
			"target":      cfg.Target,
			"volumelabel": cfg.Volumelabel,
		})
	}
	return entryList
}

func flattenDescDetails(cfg *swagger_models.Details) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"agreement_list": utils.FlattenStringMap(cfg.AgreementList),
		"app_category":   utils.PtrValStr(cfg.AppCategory),
		"category":       utils.PtrValStr(cfg.Category),
		"license_list":   utils.FlattenStringMap(cfg.LicenseList),
		"logo":           utils.FlattenStringMap(cfg.Logo),
		"os":             cfg.Os,
		"support":        cfg.Support,
	}}
}

func flattenContainerDetail(cfg *swagger_models.ContainerDetail) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"container_create_option": cfg.ContainerCreateOption,
	}}
}

func flattenUserDataTemplate(cfg *swagger_models.UserDataTemplate) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"custom_config": flattenCustomConfig(cfg.CustomConfig),
	}}
}

func flattenVmManifest(cfg *swagger_models.VMManifest) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"apptype":          utils.PtrValStr(cfg.AppType),
		"configuration":    flattenUserDataTemplate(cfg.Configuration),
		"container_detail": flattenContainerDetail(cfg.ContainerDetail),
		"deployment_type":  utils.PtrValStr(cfg.DeploymentType),
		"desc_detail":      flattenDescDetails(cfg.Desc),
		"description":      cfg.Description,
		"display_name":     cfg.DisplayName,
		"enablevnc":        cfg.Enablevnc,
		"image":            flattenAppImages(cfg.Images),
		"interface":        flattenAppInterfaces(cfg.Interfaces),
		"module":           flattenAppModule(cfg.Module),
		"name":             cfg.Name,
		"owner":            flattenAppAuthor(cfg.Owner),
		"resource":         flattenAppResources(cfg.Resources),
		"vmmode":           utils.PtrValStr(cfg.Vmmode),
	}}
}

func flattenEdgeAppConfig(cfg *swagger_models.App,
	manifestConfigured bool) map[string]interface{} {
	data := map[string]interface{}{
		"id":            cfg.ID,
		"origin_type":   utils.PtrValStr(cfg.OriginType),
		"parent_detail": utils.FlattenObjectParentDetail(cfg.ParentDetail),
		"revision":      utils.FlattenObjectRevision(cfg.Revision),
	}
	data["name"] = utils.PtrValStr(cfg.Name)
	data["title"] = utils.PtrValStr(cfg.Title)
	data["description"] = cfg.Description
	data["cpus"] = int(cfg.Cpus)
	data["drives"] = int(cfg.Drives)
	if manifestConfigured {
		data["manifest"] = flattenVmManifest(cfg.ManifestJSON)
	}
	data["memory"] = int(cfg.Memory)
	data["networks"] = int(cfg.Networks)
	data["storage"] = int(cfg.Storage)
	data["user_defined_version"] = cfg.UserDefinedVersion
	utils.CheckIfAllKeysExist(zschemas.EdgeAppSchema, data)
	return data
}

// Read the Resource Group
func readEdgeApp(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	id := resourcedata.GetStr(d, "id")
	name := resourcedata.GetStr(d, "name")

	client, ok := meta.(*client.Client)
	if !ok {
		return diag.Errorf("expect meta to be of type client.Client{} but is %T", meta)
	}

	log.Printf("PROVIDER:  readEdgeAppResource - id: %s, name: %s", id, name)
	if (id == "") && (name == "") {
		return diag.Errorf("The arguments \"id\" or \"name\" are required, but no definition was found.")
	}
	cfg, err, httpResp := getEdgeApp(client, name, id)
	if err != nil {
		if httpResp != nil && zedcloudapi.IsObjectNotFound(httpResp) {
			log.Printf("PROVIDER:  REMOVED Network %s (id: %s) from State", id, name)
			schema.RemoveFromState(d, nil)
			return diags
		}
		return diag.Errorf("[ERROR] Edge App %s (id: %s) not found. Err: %s",
			name, id, err.Error())
	}

	// Take the Config and convert it to terraform object
	// Special case for manifest / manifest_file
	manifestConfigured, _ := rdEntryByKey(d, "manifest")
	flattenedData := flattenEdgeAppConfig(cfg, manifestConfigured)
	marshalData(d, flattenedData)
	return diags
}

func readDataSourceEdgeApp(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readEdgeApp(ctx, d, meta)
}
