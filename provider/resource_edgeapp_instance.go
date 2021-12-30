// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
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
func hvModePtr(strVal string) *swagger_models.HvMode {
    val := swagger_models.HvMode(strVal)
    return &val
}

func variableFileEncodingPtr(strVal string) *swagger_models.VariableFileEncoding {
    val := swagger_models.VariableFileEncoding(strVal)
    return &val
}

func variableVariableFormatPtr(strVal string) *swagger_models.VariableVariableFormat {
    val := swagger_models.VariableVariableFormat(strVal)
    return &val
}

func rdAppInstVariableGroupVariable(
	d map[string]interface{}) (*swagger_models.VariableGroupVariable, error) {
    variable := swagger_models.VariableGroupVariable{
        Default: rdEntryStr(d, "default"),
        Encode: variableFileEncodingPtr(rdEntryStr(d, "encode")),
        Format: variableVariableFormatPtr(rdEntryStr(d, "format")),
        Label: rdEntryStrPtrOrNil(d, "label"),
        MaxLength: rdEntryStr(d, "max_length"),
        Name: rdEntryStrPtrOrNil(d, "name"),
        Options: make([]*swagger_models.VariableOptionVal, 0),
    }
	options := rdEntryList(d, "option")
	for _, option := range options {
		optionVal := swagger_models.VariableOptionVal{}
		optionVal.Label = rdEntryStr(option, "label")
		optionVal.Value = rdEntryStr(option, "value")
		variable.Options = append(variable.Options, &optionVal)
	}
	variable.Required = rdEntryBool(d, "required")
	variable.Value = rdEntryStr(d, "value")
	return &variable, nil
}

func variableGroupConditionOperatorPtr(strVal string) *swagger_models.VariableGroupConditionOperator {
    val := swagger_models.VariableGroupConditionOperator(strVal)
    return &val
}

func rdAppInstCustomVariableConditionEntry(d map[string]interface{}) (interface{}, error) {
	vgc := swagger_models.VariableGroupCondition{}
	vgc.Name = rdEntryStr(d, "name")
	vgc.Operator = variableGroupConditionOperatorPtr(rdEntryStr(d, "operator"))
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

func ioTypePtr(strVal string) *swagger_models.IoType {
    val := swagger_models.IoType(strVal)
    return &val
}

func rdAppInstIntfIO(d map[string]interface{}) (
	*swagger_models.PhyAdapter, error) {

	entryFunc := func(d map[string]interface{}) (interface{}, error) {
		return &swagger_models.PhyAdapter{
			Name: rdEntryStr(d, "name"),
			Tags: rdEntryStrMap(d, "tags"),
			Type: ioTypePtr(rdEntryStr(d, "type")),
		}, nil
	}
	val, err := rdEntryStructPtr(d, "io", entryFunc)
	if val == nil || err != nil {
		return nil, err
	}
	return val.(*swagger_models.PhyAdapter), nil
}

func rdAppInstInterfaceEntry(d map[string]interface{}) (
	*swagger_models.AppInterface, error) {
	intf := &swagger_models.AppInterface{}
	var err error
	intf.AccessVlanID = rdEntryInt64(d, "access_vlan_id")
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

func deploymentTypePtr(strVal string) *swagger_models.DeploymentType {
    val := swagger_models.DeploymentType(strVal)
    return &val
}

func appTypePtr(strVal string) *swagger_models.AppType {
    val := swagger_models.AppType(strVal)
    return &val
}

func rdUpdateAppInstCfg(cfg *swagger_models.AppInstance, d *schema.ResourceData) error {
	cfg.Activate = rdEntryBoolPtrOrNil(d, "activate")
	cfg.AppID = rdEntryStrPtrOrNil(d, "app_id")
	cfg.AppPolicyID = rdEntryStr(d, "app_policy_id")
	cfg.AppType = appTypePtr(rdEntryStr(d, "app_type"))
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
	cfg.DeploymentType = deploymentTypePtr(rdEntryStr(d, "deployment_type"))
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
