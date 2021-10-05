// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/schemas"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
	"log"
)

var appInstUrlExtension = "apps/instances"

var AppInstResourceSchema = &schema.Resource{
	CreateContext: createAppInstResource,
	ReadContext:   readResourceAppInst,
	UpdateContext: updateAppInstResource,
	DeleteContext: deleteAppInstResource,
	Schema:        schemas.AppInstSchema,
	Importer: &schema.ResourceImporter{
		StateContext: schema.ImportStatePassthroughContext,
	},
}

// The schema for a resource group data source
func getAppInstResourceSchema() *schema.Resource {
	return AppInstResourceSchema
}

func getAppInstUrl(name, id, urlType string) string {
	return zedcloudapi.UrlForObjectRequest(appInstUrlExtension, name, id, urlType)
}

func rdAppInstVariableGroupVariable(
	d map[string]interface{}) (*swagger_models.VariableGroupVariable, error) {
	variable := swagger_models.VariableGroupVariable{}
	variable.Default = rdEntryStr(d, "default")
	encode := swagger_models.VariableFileEncoding(rdEntryStr(d, "encode"))
	variable.Encode = &encode
	format := swagger_models.VariableVariableFormat(rdEntryStr(d, "format"))
	variable.Format = &format
	variable.Label = rdEntryStrPtrOrNil(d, "label")
	variable.MaxLength = rdEntryStr(d, "max_length")
	variable.Name = rdEntryStrPtrOrNil(d, "name")

	variable.Options = make([]*swagger_models.VariableOptionVal, 0)
	options := rdEntryList(d, "option")
	for _, option := range options {
		optionVal := swagger_models.VariableOptionVal{}
		optionVal.Label = rdEntryStr(option, "label")
		optionVal.Value = rdEntryStr(option, "value")
		variable.Options = append(variable.Options, &optionVal)
	}

	variable.ProcessInput = rdEntryStr(d, "process_input")
	variable.Required = rdEntryBool(d, "required")
	variable.Type = rdEntryStr(d, "type")
	variable.Value = rdEntryStr(d, "value")
	return &variable, nil
}

func rdAppInstCustomVariableConditionEntry(d map[string]interface{}) (interface{}, error) {
	vgc := swagger_models.VariableGroupCondition{}
	vgc.Name = rdEntryStr(d, "name")
	operator := swagger_models.VariableGroupConditionOperator(rdEntryStr(d, "operator"))
	vgc.Operator = &operator
	vgc.Value = rdEntryStr(d, "value")
	return &vgc, nil
}

func rdAppInstCustomVariableCondition(
	d map[string]interface{}) (*swagger_models.VariableGroupCondition, error) {
	val, err := rdEntryStructPtr(d, "condition", rdAppInstCustomVariableConditionEntry)
	if val == nil || err != nil {
		return nil, err
	}
	return val.(*swagger_models.VariableGroupCondition), nil
}

func rdAppInstCustomVariableGroup(
	d map[string]interface{}) (*swagger_models.CustomConfigVariableGroup, error) {

	var err error
	ccvg := swagger_models.CustomConfigVariableGroup{}
	ccvg.Condition, err = rdAppInstCustomVariableCondition(d)
	if err != nil {
		return nil, err
	}
	ccvg.Name = rdEntryStr(d, "name")
	ccvg.Required = rdEntryBool(d, "required")

	ccvg.Variables = make([]*swagger_models.VariableGroupVariable, 0)
	variables := rdEntryList(d, "variable")
	for _, v := range variables {
		vgv, err := rdAppInstVariableGroupVariable(v.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		ccvg.Variables = append(ccvg.Variables, vgv)
	}
	return &ccvg, nil
}

func rdAppInstCustomConfig(d map[string]interface{}) (interface{}, error) {
	customConfig := swagger_models.CustomConfig{}
	customConfig.Add = rdEntryBool(d, "add")
	customConfig.AllowStorageResize = rdEntryBool(d, "allow_storage_resize")
	customConfig.FieldDelimiter = rdEntryStr(d, "field_delimiter")
	customConfig.Name = rdEntryStr(d, "name")
	customConfig.Override = rdEntryBool(d, "override")
	customConfig.Template = rdEntryStr(d, "template")
	customConfig.VariableGroups = make([]*swagger_models.CustomConfigVariableGroup, 0)
	variableGroups := rdEntryList(d, "variable_group")
	for _, g := range variableGroups {
		vg, err := rdAppInstCustomVariableGroup(g.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		customConfig.VariableGroups = append(customConfig.VariableGroups, vg)
	}
	return &customConfig, nil
}

func rdAppInstVmInfoEntry(d map[string]interface{}) (interface{}, error) {
	mode := swagger_models.HvMode(rdEntryStr(d, "mode"))
	if mode == "" {
		return nil, fmt.Errorf("Vm Mode cannot be empty string")
	}
	return &swagger_models.VM{
		Cpus:       rdEntryInt64PtrOrNil(d, "cpus"),
		Memory:     rdEntryInt64PtrOrNil(d, "memory"),
		Mode:       &mode,
		Vnc:        rdEntryBool(d, "vnc"),
		VncDisplay: rdEntryInt64(d, "vnc_display"),
	}, nil
}

func rdAppInstVmInfo(d *schema.ResourceData) (*swagger_models.VM, error) {
	val, err := rdEntryStructPtr(d, "vminfo", rdAppInstVmInfoEntry)
	if val == nil || err != nil {
		return nil, err
	}
	return val.(*swagger_models.VM), nil
}

func rdAppInstIntfAclMatches(d map[string]interface{}) (
	[]*swagger_models.AppACEMatch, error) {
	val := rdEntryList(d, "match")
	matchList := make([]*swagger_models.AppACEMatch, 0)
	for _, dataEntry := range val {
		matchData := dataEntry.(map[string]interface{})
		matchEntry := &swagger_models.AppACEMatch{}
		matchEntry.Type = rdEntryStrPtrOrNil(matchData, "type")
		matchEntry.Value = rdEntryStrPtrOrNil(matchData, "value")
		matchList = append(matchList, matchEntry)
	}
	return matchList, nil
}

func rdAppInstIntfIO(d map[string]interface{}) (
	*swagger_models.PhyAdapter, error) {

	entryFunc := func(d map[string]interface{}) (interface{}, error) {
		iotype := swagger_models.IoType(rdEntryStr(d, "type"))
		return &swagger_models.PhyAdapter{
			Name: rdEntryStr(d, "name"),
			Tags: rdEntryStrMap(d, "tags"),
			Type: &iotype,
		}, nil
	}
	val, err := rdEntryStructPtr(d, "io", entryFunc)
	if val == nil || err != nil {
		return nil, err
	}
	return val.(*swagger_models.PhyAdapter), nil
}

func rdAppInstIntfAclActionMapParams(d map[string]interface{}) (
	*swagger_models.AppMapParams, error) {

	entryFunc := func(d map[string]interface{}) (interface{}, error) {
		return &swagger_models.AppMapParams{
			Port: rdEntryInt64PtrOrNil(d, "port"),
		}, nil
	}

	val, err := rdEntryStructPtr(d, "mapparams", entryFunc)
	if val == nil || err != nil {
		return nil, err
	}
	return val.(*swagger_models.AppMapParams), nil
}

func rdAppInstIntfAclActionEntry(d map[string]interface{}) (*swagger_models.AppACEAction, error) {
	actionEntry := &swagger_models.AppACEAction{}
	actionEntry.Drop = rdEntryBool(d, "drop")
	actionEntry.Limit = rdEntryBool(d, "limit")
	actionEntry.Limitburst = rdEntryInt64PtrOrNil(d, "limitburst")
	actionEntry.Limitrate = rdEntryInt64PtrOrNil(d, "limitrate")
	actionEntry.Limitunit = rdEntryStrPtrOrNil(d, "limitunit")
	var err error
	actionEntry.Mapparams, err = rdAppInstIntfAclActionMapParams(d)
	if err != nil {
		return nil, err
	}
	actionEntry.Portmap = rdEntryBool(d, "portmap")
	return actionEntry, nil
}

func rdAppInstIntfAclActions(d map[string]interface{}) ([]*swagger_models.AppACEAction, error) {
	actionList := make([]*swagger_models.AppACEAction, 0)
	actions := rdEntryList(d, "action")
	for _, val := range actions {
		actionEntry, err := rdAppInstIntfAclActionEntry(val.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		actionList = append(actionList, actionEntry)
	}
	return actionList, nil
}

func rdAppInstIntfAclEntry(d map[string]interface{}) (*swagger_models.AppACE, error) {
	var err error
	aclEntry := &swagger_models.AppACE{}
	aclEntry.Actions, err = rdAppInstIntfAclActions(d)
	if err != nil {
		return nil, err
	}
	aclEntry.ID = rdEntryInt32PtrOrNil(d, "id")
	aclEntry.Matches, err = rdAppInstIntfAclMatches(d)
	if err != nil {
		return nil, err
	}
	aclEntry.Name = rdEntryStrPtrOrNil(d, "name")
	return aclEntry, nil
}

func rdAppInstIntfAcls(d map[string]interface{}) ([]*swagger_models.AppACE, error) {
	aclList := make([]*swagger_models.AppACE, 0)
	acls := rdEntryList(d, "acl")
	for _, val := range acls {
		aclEntry, err := rdAppInstIntfAclEntry(val.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		aclList = append(aclList, aclEntry)
	}
	return aclList, nil
}

func rdAppInstInterfaceEntry(d map[string]interface{}) (
	*swagger_models.AppInterface, error) {
	intf := &swagger_models.AppInterface{}
	var err error
	intf.AccessVlanID = rdEntryInt64(d, "access_vlan_id")
	intf.Acls, err = rdAppInstIntfAcls(d)
	if err != nil {
		return nil, err
	}
	intf.DefaultNetInstance = rdEntryBool(d, "default_net_instance")
	intf.Intfname = rdEntryStrPtrOrNil(d, "intfname")
	intf.Io, err = rdAppInstIntfIO(d)
	if err != nil {
		return nil, err
	}
	intf.Ipaddr = rdEntryStrPtrOrNil(d, "ipaddr")
	intf.Macaddr = rdEntryStrPtrOrNil(d, "macaddr")
	intf.Netinstname = rdEntryStrPtrOrNil(d, "netinstname")
	intf.Netinsttag = rdEntryStrMap(d, "netinsttag")
	return intf, nil
}

func rdAppInstInterfaces(d *schema.ResourceData) ([]*swagger_models.AppInterface, error) {
	intfList := make([]*swagger_models.AppInterface, 0)
	interfaces := rdEntryList(d, "interface")
	if len(interfaces) == 0 {
		return intfList, nil
	}
	for _, val := range interfaces {
		intfData := val.(map[string]interface{})
		intfEntry, err := rdAppInstInterfaceEntry(intfData)
		if err != nil {
			return intfList, err
		}
		intfList = append(intfList, intfEntry)
	}
	return intfList, nil
}

func rdAppInstDriveEntry(d map[string]interface{}) *swagger_models.Drive {
	return &swagger_models.Drive{
		Cleartext:   rdEntryBool(d, "cleartext"),
		Drvtype:     rdEntryStrPtrOrNil(d, "drvtype"),
		Ignorepurge: rdEntryBool(d, "ignorepurge"),
		Imagename:   rdEntryStrPtrOrNil(d, "imagename"),
		Maxsize:     rdEntryUint64PtrOrNil(d, "maxsize"),
		Mountpath:   rdEntryStr(d, "mountpath"),
		Preserve:    rdEntryBool(d, "preserve"),
		Readonly:    rdEntryBool(d, "Readonly"),
		Target:      rdEntryStrPtrOrNil(d, "Target"),
		Volumelabel: rdEntryStr(d, "volumelabel"),
	}
}

func rdAppInstDrives(d *schema.ResourceData) []*swagger_models.Drive {
	entryList := make([]*swagger_models.Drive, 0)
	dataList := rdEntryList(d, "drive")
	if len(dataList) == 0 {
		return entryList
	}
	for _, val := range dataList {
		entry := val.(map[string]interface{})
		entryList = append(entryList, rdAppInstDriveEntry(entry))
	}
	return entryList
}

func rdAppInstLogs(d map[string]interface{}) (interface{}, error) {
	return &swagger_models.AppInstanceLogs{
		Access: rdEntryBoolPtrOrNil(d, "access"),
	}, nil
}

func rdUpdateAppInstCfg(cfg *swagger_models.AppInstance, d *schema.ResourceData) error {
	cfg.Activate = rdEntryBoolPtrOrNil(d, "activate")
	cfg.AppID = rdEntryStrPtrOrNil(d, "app_id")
	cfg.AppPolicyID = rdEntryStr(d, "app_policy_id")
	appType := swagger_models.AppType(rdEntryStr(d, "app_type"))
	cfg.AppType = &appType
	cfg.Bundleversion = rdEntryStr(d, "bundleversion")
	cfg.ClusterID = rdEntryStr(d, "cluster_id")
	cfg.CollectStatsIPAddr = rdEntryStr(d, "collect_stats_ip_addr")
	cfg.CryptoKey = rdEntryStr(d, "crypto_key")
	val, err := rdEntryStructPtr(d, "custom_config", rdAppInstCustomConfig)
	if err != nil {
		return err
	}
	cfg.CustomConfig = nil
	if val != nil {
		cfg.CustomConfig = val.(*swagger_models.CustomConfig)
	}
	deploymentType := swagger_models.DeploymentType(
		rdEntryStr(d, "deployment_type"))
	cfg.DeploymentType = &deploymentType
	cfg.Description = rdEntryStr(d, "description")
	cfg.DeviceID = rdEntryStrPtrOrNil(d, "device_id")
	cfg.Drives = rdAppInstDrives(d)
	cfg.EncryptedSecrets = rdEntryStrMap(d, "encrypted_secrets")
	cfg.Interfaces, err = rdAppInstInterfaces(d)
	if err != nil {
		return err
	}
	val, err = rdEntryStructPtr(d, "logs", rdAppInstLogs)
	if err != nil {
		return err
	}
	cfg.Logs = nil
	if val != nil {
		cfg.Logs = val.(*swagger_models.AppInstanceLogs)
	}
	cfg.ProjectID = rdEntryStrPtrOrNil(d, "project_id")
	cfg.RemoteConsole = rdEntryBool(d, "remote_console")
	cfg.Tags = rdEntryStrMap(d, "tags")
	cfg.Title = rdEntryStrPtrOrNil(d, "title")
	cfg.UserDefinedVersion = rdEntryStr(d, "user_defined_version")
	cfg.Vminfo, err = rdAppInstVmInfo(d)
	if err != nil {
		return err
	}
	return nil
}

// Create the Resource Group
func createAppInstResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	name := rdEntryStr(d, "name")
	id := rdEntryStr(d, "id")
	errMsgPrefix := getErrMsgPrefix(name, id, "App Instance", "Create")
	if client == nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, "nil Client")
	}
	cfg := &swagger_models.AppInstance{
		Name: &name,
	}
	err := rdUpdateAppInstCfg(cfg, d)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	log.Printf("[INFO] Creating AppInst: %s", name)
	client.XRequestIdPrefix = "TF-appinst-create"
	rspData := &swagger_models.ZsrvResponse{}
	_, err = client.SendReq("POST", appInstUrlExtension, cfg, rspData)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	id = rspData.ObjectID
	log.Printf("App Instance %s (ID: %s) Successfully created\n",
		rspData.ObjectName, id)
	d.SetId(id)
	err = getAppInstAndPublishData(client, d, name, id, true)
	if err != nil {
		log.Printf("***[ERROR]- Failed to get App Instance: %s (ID: %s) after "+
			"creating it. Err: %s", name, id, err.Error())
	}
	return diags
}

// Update the Resource Group
func updateAppInstResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	name := rdEntryStr(d, "name")
	id := rdEntryStr(d, "id")
	errMsgPrefix := getErrMsgPrefix(name, id, "App Instance", "Update")
	if client == nil {
		return diag.Errorf("%s nil Client", errMsgPrefix)
	}
	cfg, err := getAppInstance(client, name, id)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	if cfg == nil {
		return diag.Errorf("%s Failed to find App Instance", errMsgPrefix)
	}
	err = checkInvalidAttrForUpdate(d, *cfg.Name, cfg.ID)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	log.Printf("[INFO] Found AppInst: %s (ID: %s)", name, cfg.ID)
	err = rdUpdateAppInstCfg(cfg, d)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	log.Printf("[INFO] Updating AppInst: %s", name)
	client.XRequestIdPrefix = "TF-appinst-update"
	urlExtension := getAppInstUrl(name, id, "update")
	rspData := &swagger_models.ZsrvResponse{}
	_, err = client.SendReq("PUT", urlExtension, cfg, rspData)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	log.Printf("App Instance %s (id: %s) update SUCCESS", name, cfg.ID)
	err = getAppInstAndPublishData(client, d, name, id, true)
	if err != nil {
		log.Printf("***[ERROR]- Failed to get App Instance: %s (ID: %s) after "+
			"updating it. Err: %s", name, id, err.Error())
	}
	return diags
}

// Delete the Resource Group
func deleteAppInstResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := (meta.(Client)).Client
	name := rdEntryStr(d, "name")
	id := rdEntryStr(d, "id")
	errMsgPrefix := getErrMsgPrefix(name, id, "App Instance", "Delete")
	cfg, err := getAppInstance(client, name, id)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	client.XRequestIdPrefix = "TF-appinst-delete"
	urlExtension := getAppInstUrl(name, id, "delete")
	rspData := &swagger_models.ZsrvResponse{}
	_, err = client.SendReq("DELETE", urlExtension, nil, rspData)
	if err != nil {
		return diag.Errorf("%s. Err: %s", errMsgPrefix, err.Error())
	}
	log.Printf("[INFO] App Instance %s (ID: %s) Delete Successful.",
		*cfg.Name, cfg.ID)
	return diags
}

func readResourceAppInst(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readAppInst(ctx, d, meta, true)
}
