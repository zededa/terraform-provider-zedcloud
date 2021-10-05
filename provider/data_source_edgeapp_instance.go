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
	name, id string) (*swagger_models.AppInstance, error) {
	rspData := &swagger_models.AppInstance{}
	err := client.GetObj(appInstUrlExtension, name, id, false, rspData)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] FAILED to get AppInst %s ( id: %s). Err: %s",
			name, id, err.Error())
	}
	return rspData, nil
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
			"type":  ptrValStr(cfg.Type),
			"value": ptrValStr(cfg.Value),
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
			"limitunit":  ptrValStr(cfg.Limitunit),
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
			"name":   ptrValStr(cfg.Name),
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
		"tags": flattenStringMap(cfg.Tags),
		"type": ioType,
	}}
}

func flattenAppInterfaces(cfgList []*swagger_models.AppInterface) []interface{} {
	entryList := make([]interface{}, 0)
	for _, cfg := range cfgList {
		entry := map[string]interface{}{
			"access_vlan_id":       int(cfg.AccessVlanID),
			"acl":                  flattenAppACEs(cfg.Acls),
			"default_net_instance": cfg.DefaultNetInstance,
			"directattach":         cfg.Directattach,
			"intfname":             ptrValStr(cfg.Intfname),
			"io":                   flattenPhyAdapter(cfg.Io),
			"ipaddr":               ptrValStr(cfg.Ipaddr),
			"macaddr":              ptrValStr(cfg.Macaddr),
			"netinstname":          ptrValStr(cfg.Netinstname),
			"netinsttag":           flattenStringMap(cfg.Netinsttag),
		}
		entryList = append(entryList, entry)
	}
	return entryList
}

func flattenDrives(cfgList []*swagger_models.Drive) []interface{} {
	entryList := make([]interface{}, 0)
	for _, cfg := range cfgList {
		maxsize := 0
		if cfg.Maxsize != nil {
			maxsize = int(*cfg.Maxsize)
		}
		entry := map[string]interface{}{
			"cleartext":   cfg.Cleartext,
			"drvtype":     ptrValStr(cfg.Drvtype),
			"ignorepurge": cfg.Ignorepurge,
			"imagename":   ptrValStr(cfg.Imagename),
			"imvolname":   cfg.Imvolname,
			"maxsize":     maxsize,
			"mountpath":   cfg.Mountpath,
			"mvolname":    cfg.Mvolname,
			"preserve":    cfg.Preserve,
			"readonly":    cfg.Readonly,
			"target":      ptrValStr(cfg.Target),
			"volumelabel": cfg.Volumelabel,
		}
		entryList = append(entryList, entry)
	}
	return entryList
}

func flattenVariableOptionVals(cfgList []*swagger_models.VariableOptionVal) []interface{} {
	entryList := make([]interface{}, 0)
	for _, cfg := range cfgList {
		entry := map[string]interface{}{
			"label": cfg.Label,
			"value": cfg.Value,
		}
		entryList = append(entryList, entry)
	}
	return entryList
}

func flattenVariableGroupVariables(cfgList []*swagger_models.VariableGroupVariable) []interface{} {
	entryList := make([]interface{}, 0)
	for _, cfg := range cfgList {
		encode := ""
		if cfg.Encode != nil {
			encode = string(*cfg.Encode)
		}
		format := ""
		if cfg.Format != nil {
			format = string(*cfg.Format)
		}
		entry := map[string]interface{}{
			"default":       cfg.Default,
			"encode":        encode,
			"format":        format,
			"label":         ptrValStr(cfg.Label),
			"max_length":    cfg.MaxLength,
			"name":          ptrValStr(cfg.Name),
			"option":        flattenVariableOptionVals(cfg.Options),
			"process_input": cfg.ProcessInput,
			"required":      cfg.Required,
			"type":          cfg.Type,
			"value":         cfg.Value,
		}
		entryList = append(entryList, entry)
	}
	return entryList
}

func flattenVariableGroupCondition(cfg *swagger_models.VariableGroupCondition) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	operator := ""
	if cfg.Operator != nil {
		operator = string(*cfg.Operator)
	}
	return []interface{}{map[string]interface{}{
		"name":     cfg.Name,
		"operator": operator,
		"value":    cfg.Value,
	}}
}

func flattenCustomConfigVariableGroups(cfgList []*swagger_models.CustomConfigVariableGroup) []interface{} {
	entryList := make([]interface{}, 0)
	for _, cfg := range cfgList {
		entry := map[string]interface{}{
			"condition": flattenVariableGroupCondition(cfg.Condition),
			"name":      cfg.Name,
			"required":  cfg.Required,
			"variable":  flattenVariableGroupVariables(cfg.Variables),
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
		"add":                  cfg.Add,
		"allow_storage_resize": cfg.AllowStorageResize,
		"field_delimiter":      cfg.FieldDelimiter,
		"name":                 cfg.Name,
		"override":             cfg.Override,
		"template":             cfg.Template,
		"variable_group":       flattenCustomConfigVariableGroups(cfg.VariableGroups),
	}}
}

func flattenAppInstance(cfg *swagger_models.AppInstance,
	computedOnly bool) map[string]interface{} {
	if cfg == nil {
		return map[string]interface{}{}
	}
	// Do not publish sensitive fields - crypto_key, encrypted_secrets
	data := map[string]interface{}{
		"bundleversion":        cfg.Bundleversion,
		"cluster_id":           cfg.ClusterID,
		"id":                   cfg.ID,
		"is_secret_updated":    cfg.IsSecretUpdated,
		"project_id":           ptrValStr(cfg.ProjectID),
		"purge":                flattenZedCloudOpsCmd(cfg.Purge),
		"refresh":              flattenZedCloudOpsCmd(cfg.Refresh),
		"restart":              flattenZedCloudOpsCmd(cfg.Restart),
		"revision":             flattenObjectRevision(cfg.Revision),
		"user_defined_version": cfg.UserDefinedVersion,
		"vminfo":               flattenVM(cfg.Vminfo),
	}
	if !computedOnly {
		data["activate"] = ptrValBool(cfg.Activate)
		data["app_id"] = ptrValStr(cfg.AppID)
		data["app_policy_id"] = cfg.AppPolicyID
		data["app_type"] = ptrValStr(cfg.AppType)
		data["collect_stats_ip_addr"] = cfg.CollectStatsIPAddr
		data["custom_config"] = flattenCustomConfig(cfg.CustomConfig)
		data["description"] = cfg.Description
		data["deployment_type"] = ptrValStr(cfg.DeploymentType)
		data["device_id"] = ptrValStr(cfg.DeviceID)
		data["drive"] = flattenDrives(cfg.Drives)
		data["interface"] = flattenAppInterfaces(cfg.Interfaces)
		data["logs"] = flattenAppInstanceLogs(cfg.Logs)
		data["name"] = ptrValStr(cfg.Name)
		data["remote_console"] = cfg.RemoteConsole
		data["tags"] = flattenStringMap(cfg.Tags)
		data["title"] = ptrValStr(cfg.Title)
	}
	flattenedDataCheckKeys(zschemas.AppInstSchema, data, computedOnly)
	return data
}

func getAppInstAndPublishData(client *zedcloudapi.Client, d *schema.ResourceData,
	name, id string, resource bool) error {
	cfg, err := getAppInstance(client, name, id)
	if err != nil {
		return fmt.Errorf("[ERROR] App Instance %s (id: %s) not found. Err: %s",
			name, id, err.Error())
	}
	marshalData(d, flattenAppInstance(cfg, resource))
	return nil
}

// Read the Resource Group
func readAppInst(ctx context.Context, d *schema.ResourceData, meta interface{},
	resource bool) diag.Diagnostics {
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	id := rdEntryStr(d, "id")
	name := rdEntryStr(d, "name")

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
