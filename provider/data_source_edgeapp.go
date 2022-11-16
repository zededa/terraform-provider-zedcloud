// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
	"log"
	"net/http"
	"strconv"
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

func getEdgeApp(client *zedcloudapi.Client,
	name, id string) (*swagger_models.App, error, *http.Response) {
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
		"environment": flattenStringMap(cfg.Environment),
		"module_type": ptrValStr(cfg.ModuleType),
		"routes":      flattenStringMap(cfg.Routes),
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
		"agreement_list": flattenStringMap(cfg.AgreementList),
		"app_category":   ptrValStr(cfg.AppCategory),
		"category":       ptrValStr(cfg.Category),
		"license_list":   flattenStringMap(cfg.LicenseList),
		"logo":           flattenStringMap(cfg.Logo),
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
		"ac_version":       ptrValStr(cfg.AcVersion),
		"apptype":          ptrValStr(cfg.AppType),
		"configuration":    flattenUserDataTemplate(cfg.Configuration),
		"container_detail": flattenContainerDetail(cfg.ContainerDetail),
		"deployment_type":  ptrValStr(cfg.DeploymentType),
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
		"vmmode":           ptrValStr(cfg.Vmmode),
	}}
}

func flattenEdgeAppConfig(cfg *swagger_models.App,
	manifestConfigured bool) map[string]interface{} {
	data := map[string]interface{}{
		"id":            cfg.ID,
		"origin_type":   ptrValStr(cfg.OriginType),
		"parent_detail": flattenObjectParentDetail(cfg.ParentDetail),
		"revision":      flattenObjectRevision(cfg.Revision),
	}
	data["name"] = ptrValStr(cfg.Name)
	data["title"] = ptrValStr(cfg.Title)
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
	flattenedDataCheckKeys(zschemas.EdgeAppSchema, data)
	return data
}

// Read the Resource Group
func readEdgeApp(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	id := rdEntryStr(d, "id")
	name := rdEntryStr(d, "name")

	if client == nil {
		return diag.Errorf("nil Client.")
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
