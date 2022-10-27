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

var EdgeAppInstanceDataSourceSchema = &schema.Resource{
	ReadContext: readDataSourceAppInst,
	Schema:      zschemas.AppInstSchema,
	Description: "Schema for data source zedcloud_edgeapp_instance. Must specify " +
		"id or name",
}

// The schema for a resource group data source
func getAppInstDataSourceSchema() *schema.Resource {
	return EdgeAppInstanceDataSourceSchema
}

func getAppInstance(client *zedcloudapi.Client,
	name, id string) (*swagger_models.AppInstance, error, *http.Response) {
	rspData := &swagger_models.AppInstance{}
	httpResp, err := client.GetObj(appInstUrlExtension, name, id, false, rspData)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] FAILED to get AppInst %s ( id: %s). Err: %s",
			name, id, err.Error()), httpResp
	}
	return rspData, nil, httpResp
}

func flattenVM(vm *swagger_models.VM) []interface{} {
	if vm == nil {
		return []interface{}{}
	}

	cpus := 0
	if vm.Cpus != nil {
		cpus = int(*vm.Cpus)
	}
	memory := 0
	if vm.Memory != nil {
		memory = int(*vm.Memory)
	}
	mode := ""
	if vm.Mode != nil {
		mode = string(*vm.Mode)
	}
	return []interface{}{map[string]interface{}{
		"cpus":        cpus,
		"memory":      memory,
		"mode":        mode,
		"vnc":         vm.Vnc,
		"vnc_display": int(vm.VncDisplay),
	}}
}

func flattenAppInstanceLogs(cfg *swagger_models.AppInstanceLogs) []interface{} {
	if cfg == nil || cfg.Access == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"access": *cfg.Access,
	}}
}

func flattenAppACEMatches(cfgList []*swagger_models.AppACEMatch) []interface{} {
	entryList := make([]interface{}, 0)
	for _, cfg := range cfgList {
		entry := map[string]interface{}{
			"type":  utils.PtrValStr(cfg.Type),
			"value": utils.PtrValStr(cfg.Value),
		}
		entryList = append(entryList, entry)
	}
	return entryList
}

func flattenAppMapParams(cfg *swagger_models.AppMapParams) []interface{} {
	if cfg == nil || cfg.Port == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"port": int(*cfg.Port),
	}}
}

func flattenAppACEActions(cfgList []*swagger_models.AppACEAction) []interface{} {
	entryList := make([]interface{}, 0)
	for _, cfg := range cfgList {
		limitburst := 0
		if cfg.Limitburst != nil {
			limitburst = int(*cfg.Limitburst)
		}
		limitrate := 0
		if cfg.Limitrate != nil {
			limitrate = int(*cfg.Limitrate)
		}
		entry := map[string]interface{}{
			"drop":       cfg.Drop,
			"limit":      cfg.Limit,
			"limitburst": limitburst,
			"limitrate":  limitrate,
			"limitunit":  utils.PtrValStr(cfg.Limitunit),
			"mapparams":  flattenAppMapParams(cfg.Mapparams),
			"portmap":    cfg.Portmap,
		}
		entryList = append(entryList, entry)
	}
	return entryList
}

func flattenAppACEs(cfgList []*swagger_models.AppACE) []interface{} {
	entryList := make([]interface{}, 0)
	for _, cfg := range cfgList {
		id := 0
		if cfg.ID != nil {
			id = int(*cfg.ID)
		}
		entry := map[string]interface{}{
			"action": flattenAppACEActions(cfg.Actions),
			"id":     id,
			"match":  flattenAppACEMatches(cfg.Matches),
			"name":   utils.PtrValStr(cfg.Name),
		}
		entryList = append(entryList, entry)
	}
	return entryList
}

func flattenPhyAdapter(cfg *swagger_models.PhyAdapter) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	ioType := ""
	if cfg.Type != nil {
		ioType = string(*cfg.Type)
	}
	return []interface{}{map[string]interface{}{
		"name": cfg.Name,
		"tags": utils.FlattenStringMap(cfg.Tags),
		"type": ioType,
	}}
}

func flattenAppInstInterfaces(cfgList []*swagger_models.AppInterface) []interface{} {
	entryList := make([]interface{}, 0)
	for _, cfg := range cfgList {
		entry := map[string]interface{}{
			"access_vlan_id":       int(cfg.AccessVlanID),
			"acl":                  flattenAppACEs(cfg.Acls),
			"default_net_instance": cfg.DefaultNetInstance,
			"directattach":         cfg.Directattach,
			"intfname":             utils.PtrValStr(cfg.Intfname),
			"io":                   flattenPhyAdapter(cfg.Io),
			"ipaddr":               utils.PtrValStr(cfg.Ipaddr),
			"macaddr":              utils.PtrValStr(cfg.Macaddr),
			"netinstname":          utils.PtrValStr(cfg.Netinstname),
			"netinsttag":           utils.FlattenStringMap(cfg.Netinsttag),
		}
		entryList = append(entryList, entry)
	}
	return entryList
}

func flattenDrives(cfgList []*swagger_models.Drive) []interface{} {
	entryList := make([]interface{}, 0)
	for _, cfg := range cfgList {
		if cfg.Imvolname != "" {
			continue
		}
		maxsize := 0
		if cfg.Maxsize != nil {
			maxsize = int(*cfg.Maxsize)
		}
		entry := map[string]interface{}{
			"cleartext":   cfg.Cleartext,
			"drvtype":     utils.PtrValStr(cfg.Drvtype),
			"ignorepurge": cfg.Ignorepurge,
			"imagename":   utils.PtrValStr(cfg.Imagename),
			"imvolname":   cfg.Imvolname,
			"maxsize":     maxsize,
			"mountpath":   cfg.Mountpath,
			"mvolname":    cfg.Mvolname,
			"preserve":    cfg.Preserve,
			"readonly":    cfg.Readonly,
			"target":      utils.PtrValStr(cfg.Target),
			"volumelabel": cfg.Volumelabel,
		}
		entryList = append(entryList, entry)
	}
	return entryList
}

func flattenCustomConfig(cfg *swagger_models.CustomConfig) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"add": cfg.Add,
	}}
}

func flattenAppInstance(cfg *swagger_models.AppInstance, resource bool) map[string]interface{} {
	if cfg == nil {
		return map[string]interface{}{}
	}
	// Do not publish sensitive fields - crypto_key, encrypted_secrets
	data := map[string]interface{}{
		"bundleversion":        cfg.Bundleversion,
		"cluster_id":           cfg.ClusterID,
		"id":                   cfg.ID,
		"is_secret_updated":    cfg.IsSecretUpdated,
		"project_id":           utils.PtrValStr(cfg.ProjectID),
		"purge":                flattenZedCloudOpsCmd(cfg.Purge),
		"refresh":              flattenZedCloudOpsCmd(cfg.Refresh),
		"restart":              flattenZedCloudOpsCmd(cfg.Restart),
		"revision":             utils.FlattenObjectRevision(cfg.Revision),
		"user_defined_version": cfg.UserDefinedVersion,
		"vminfo":               flattenVM(cfg.Vminfo),
	}
	data["activate"] = ptrValBool(cfg.Activate)
	data["app_id"] = utils.PtrValStr(cfg.AppID)
	data["app_policy_id"] = cfg.AppPolicyID
	data["app_type"] = utils.PtrValStr(cfg.AppType)
	data["collect_stats_ip_addr"] = cfg.CollectStatsIPAddr
	if !resource {
		// Do not flatten secrets for resources. ZEDCloud doesn't return
		// details of secrets. So cfg is the source of truth for secrets.
		data["custom_config"] = flattenCustomConfig(cfg.CustomConfig)
	}
	data["description"] = cfg.Description
	data["device_id"] = utils.PtrValStr(cfg.DeviceID)
	data["drive"] = flattenDrives(cfg.Drives)
	data["interface"] = flattenAppInstInterfaces(cfg.Interfaces)
	data["logs"] = flattenAppInstanceLogs(cfg.Logs)
	data["name"] = utils.PtrValStr(cfg.Name)
	data["remote_console"] = cfg.RemoteConsole
	data["tags"] = utils.FlattenStringMap(cfg.Tags)
	data["title"] = utils.PtrValStr(cfg.Title)
	utils.CheckIfAllKeysExist(zschemas.AppInstSchema, data)
	return data
}

func getAppInstAndPublishData(client *zedcloudapi.Client, d *schema.ResourceData,
	name, id string, resource bool) error {
	cfg, err, httpRsp := getAppInstance(client, name, id)
	if err != nil {
		err = fmt.Errorf("[ERROR] App Instance %s (id: %s) not found. Err: %s",
			name, id, err.Error())
		if httpRsp != nil && zedcloudapi.IsObjectNotFound(httpRsp) {
			log.Printf("PROVIDER:  REMOVED App Instance %s (id: %s) from State", id, name)
			schema.RemoveFromState(d, nil)
			return nil
		}
		return err
	}
	marshalData(d, flattenAppInstance(cfg, resource))
	return nil
}

// Read the Resource Group
func readAppInst(ctx context.Context, d *schema.ResourceData,
	meta interface{}, resource bool) diag.Diagnostics {
	var diags diag.Diagnostics
	client := (meta.(Client)).Client
	id := resourcedata.GetStr(d, "id")
	name := resourcedata.GetStr(d, "name")

	if client == nil {
		return diag.Errorf("nil Client.")
	}
	log.Printf("PROVIDER:  readAppInstResource - id: %s, name: %s", id, name)
	if (id == "") && (name == "") {
		return diag.Errorf("The arguments \"id\" or \"name\" are required, but no definition was found.")
	}
	err := getAppInstAndPublishData(client, d, name, id, resource)
	if err != nil {
		return diag.Errorf("%s", err.Error())
	}
	return diags
}

func readDataSourceAppInst(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readAppInst(ctx, d, meta, false)
}
