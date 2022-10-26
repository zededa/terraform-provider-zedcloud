// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/schemas"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
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
		Default:   getStr(d, "default"),
		Encode:    variableFileEncodingPtr(getStr(d, "encode")),
		Format:    variableVariableFormatPtr(getStr(d, "format")),
		Label:     getStrPtrOrNil(d, "label"),
		MaxLength: getStr(d, "max_length"),
		Name:      getStrPtrOrNil(d, "name"),
		Options:   make([]*swagger_models.VariableOptionVal, 0),
	}
	options := getList(d, "option")
	for _, option := range options {
		optionVal := swagger_models.VariableOptionVal{}
		optionVal.Label = getStr(option, "label")
		optionVal.Value = getStr(option, "value")
		variable.Options = append(variable.Options, &optionVal)
	}
	variable.Required = getBool(d, "required")
	variable.Value = getStr(d, "value")
	return &variable, nil
}

func variableGroupConditionOperatorPtr(strVal string) *swagger_models.VariableGroupConditionOperator {
	val := swagger_models.VariableGroupConditionOperator(strVal)
	return &val
}

func rdAppInstCustomVariableConditionEntry(d map[string]interface{}) (interface{}, error) {
	vgc := swagger_models.VariableGroupCondition{}
	vgc.Name = getStr(d, "name")
	vgc.Operator = variableGroupConditionOperatorPtr(getStr(d, "operator"))
	vgc.Value = getStr(d, "value")
	return &vgc, nil
}

func rdAppInstCustomVariableCondition(
	d map[string]interface{}) (*swagger_models.VariableGroupCondition, error) {
	val, err := getStructPtr(d, "condition", rdAppInstCustomVariableConditionEntry)
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
	ccvg.Name = getStr(d, "name")
	ccvg.Required = getBool(d, "required")

	ccvg.Variables = make([]*swagger_models.VariableGroupVariable, 0)
	variables := getList(d, "variable")
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
	customConfig.Add = getBool(d, "add")
	customConfig.AllowStorageResize = getBool(d, "allow_storage_resize")
	customConfig.FieldDelimiter = getStr(d, "field_delimiter")
	customConfig.Name = getStr(d, "name")
	customConfig.Override = getBool(d, "override")
	customConfig.Template = getStr(d, "template")
	customConfig.VariableGroups = make([]*swagger_models.CustomConfigVariableGroup, 0)
	variableGroups := getList(d, "variable_group")
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
			Name: getStr(d, "name"),
			Tags: getStrMap(d, "tags"),
			Type: ioTypePtr(getStr(d, "type")),
		}, nil
	}
	val, err := getStructPtr(d, "io", entryFunc)
	if val == nil || err != nil {
		return nil, err
	}
	return val.(*swagger_models.PhyAdapter), nil
}

func rdAppInstInterfaceEntry(d map[string]interface{}) (
	*swagger_models.AppInterface, error) {
	intf := &swagger_models.AppInterface{}
	var err error
	intf.AccessVlanID = getInt64(d, "access_vlan_id")
	intf.DefaultNetInstance = getBool(d, "default_net_instance")
	intf.Intfname = getStrPtrOrNil(d, "intfname")
	intf.Io, err = rdAppInstIntfIO(d)
	if err != nil {
		return nil, err
	}
	intf.Ipaddr = getStrPtrOrNil(d, "ipaddr")
	intf.Macaddr = getStrPtrOrNil(d, "macaddr")
	intf.Netinstname = getStrPtrOrNil(d, "netinstname")
	intf.Netinsttag = getStrMap(d, "netinsttag")
	return intf, nil
}

func rdAppInstInterfaces(d *schema.ResourceData) ([]*swagger_models.AppInterface, error) {
	intfList := make([]*swagger_models.AppInterface, 0)
	interfaces := getList(d, "interface")
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
		Cleartext:   getBool(d, "cleartext"),
		Drvtype:     getStrPtrOrNil(d, "drvtype"),
		Ignorepurge: getBool(d, "ignorepurge"),
		Imagename:   getStrPtrOrNil(d, "imagename"),
		Maxsize:     getUint64PtrOrNil(d, "maxsize"),
		Mountpath:   getStr(d, "mountpath"),
		Preserve:    getBool(d, "preserve"),
		Readonly:    getBool(d, "Readonly"),
		Target:      getStrPtrOrNil(d, "Target"),
		Volumelabel: getStr(d, "volumelabel"),
	}
}

func rdAppInstDrives(d *schema.ResourceData) []*swagger_models.Drive {
	entryList := make([]*swagger_models.Drive, 0)
	dataList := getList(d, "drive")
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
		Access: getBoolPtrOrNil(d, "access"),
	}, nil
}

func deploymentTypePtr(strVal string) (*swagger_models.DeploymentType, error) {
	if strVal == "" {
		strVal = "DEPLOYMENT_TYPE_UNSPECIFIED"
	}
	val := swagger_models.DeploymentType(strVal)
	if err := val.Validate(nil); err != nil {
		return nil, err
	}
	return &val, nil
}

func appTypePtr(strVal string) *swagger_models.AppType {
	val := swagger_models.AppType(strVal)
	return &val
}

func rdUpdateAppInstCfg(cfg *swagger_models.AppInstance, d *schema.ResourceData) error {
	cfg.Activate = getBoolPtrOrNil(d, "activate")
	cfg.AppID = getStrPtrOrNil(d, "app_id")
	cfg.AppPolicyID = getStr(d, "app_policy_id")
	cfg.AppType = appTypePtr(getStr(d, "app_type"))
	cfg.Bundleversion = getStr(d, "bundleversion")
	cfg.ClusterID = getStr(d, "cluster_id")
	cfg.CollectStatsIPAddr = getStr(d, "collect_stats_ip_addr")
	cfg.CryptoKey = getStr(d, "crypto_key")
	val, err := getStructPtr(d, "custom_config", rdAppInstCustomConfig)
	if err != nil {
		return err
	}
	cfg.CustomConfig = nil
	if val != nil {
		cfg.CustomConfig = val.(*swagger_models.CustomConfig)
	}
	cfg.Description = getStr(d, "description")
	cfg.DeviceID = getStrPtrOrNil(d, "device_id")
	cfg.Drives = rdAppInstDrives(d)
	cfg.EncryptedSecrets = getStrMap(d, "encrypted_secrets")
	cfg.Interfaces, err = rdAppInstInterfaces(d)
	if err != nil {
		return err
	}
	val, err = getStructPtr(d, "logs", rdAppInstLogs)
	if err != nil {
		return err
	}
	cfg.Logs = nil
	if val != nil {
		cfg.Logs = val.(*swagger_models.AppInstanceLogs)
	}
	cfg.ProjectID = getStrPtrOrNil(d, "project_id")
	cfg.RemoteConsole = getBool(d, "remote_console")
	cfg.Tags = getStrMap(d, "tags")
	cfg.Title = getStrPtrOrNil(d, "title")
	cfg.UserDefinedVersion = getStr(d, "user_defined_version")
	return nil
}

// Create the Resource Group
func createAppInstResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	name := getStr(d, "name")
	id := getStr(d, "id")
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
	name := getStr(d, "name")
	id := getStr(d, "id")
	errMsgPrefix := getErrMsgPrefix(name, id, "App Instance", "Update")
	if client == nil {
		return diag.Errorf("%s nil Client", errMsgPrefix)
	}
	cfg, err, _ := getAppInstance(client, name, id)
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
	name := getStr(d, "name")
	id := getStr(d, "id")
	errMsgPrefix := getErrMsgPrefix(name, id, "App Instance", "Delete")
	cfg, err, httpRsp := getAppInstance(client, name, id)
	if err != nil {
		if httpRsp != nil && zedcloudapi.IsObjectNotFound(httpRsp) {
			log.Printf("%s Not Found", errMsgPrefix)
			return diags
		}
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
