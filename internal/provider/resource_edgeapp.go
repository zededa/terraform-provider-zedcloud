// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/internal/client"
	"github.com/zededa/terraform-provider-zedcloud/internal/resourcedata"
	"github.com/zededa/terraform-provider-zedcloud/schemas"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
)

var edgeAppUrlExtension = "apps"

var EdgeAppResourceSchema = &schema.Resource{
	CreateContext: createEdgeAppResource,
	ReadContext:   readResourceEdgeApp,
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

func rdEntryAppManifestFromFile(d *schema.ResourceData) (*swagger_models.VMManifest, error) {
	manifest_file := resourcedata.GetStr(d, "manifest_file")
	if manifest_file == "" {
		return nil, fmt.Errorf("One of manifest or manifest_file needs to be specified" +
			" to create EdgeApp")
	}
	manifestStr, err := ioutil.ReadFile(manifest_file)
	if err != nil {
		return nil, fmt.Errorf("Failed to read manifest file %s. err: %v",
			manifest_file, err)
	}
	manifest := &swagger_models.VMManifest{}
	err = json.Unmarshal([]byte(manifestStr), manifest)
	if err != nil {
		return nil, fmt.Errorf("Failed to UnMarshaling Manifest file. Err: %v", err)
	}
	return manifest, nil
}

func rdAppManifestResources(d map[string]interface{}) ([]*swagger_models.Resource, error) {
	entryFunc := func(d map[string]interface{}) (*swagger_models.Resource, error) {
		return &swagger_models.Resource{
			Name:  resourcedata.GetStr(d, "name"),
			Value: resourcedata.GetStr(d, "value"),
		}, nil
	}
	cfgList := make([]*swagger_models.Resource, 0)
	rdList := rdEntryList(d, "resource")
	if len(rdList) == 0 {
		return cfgList, nil
	}
	for _, val := range rdList {
		rdEntry := val.(map[string]interface{})
		cfgEntry, err := entryFunc(rdEntry)
		if err != nil {
			return cfgList, err
		}
		cfgList = append(cfgList, cfgEntry)
	}
	return cfgList, nil
}

func rdAppManifestAuthor(d map[string]interface{}) (*swagger_models.Author, error) {
	entryFunc := func(d map[string]interface{}) (interface{}, error) {
		return &swagger_models.Author{
			Company: resourcedata.GetStr(d, "company"),
			Email:   resourcedata.GetStr(d, "email"),
			User:    resourcedata.GetStr(d, "user"),
			Website: resourcedata.GetStr(d, "website"),
		}, nil
	}
	val, err := resourcedata.GetStructPtr(d, "owner", entryFunc)
	if val == nil || err != nil {
		return nil, err
	}
	return val.(*swagger_models.Author), nil
}

func appModuleTypePtr(strVal string) (*swagger_models.ModuleType, error) {
	val := swagger_models.ModuleType(strVal)
	if err := val.Validate(nil); err != nil {
		return nil, err
	}
	return &val, nil
}

func rdAppManifestModuleDetails(d map[string]interface{}) (
	*swagger_models.ModuleDetail, error) {
	entryFunc := func(d map[string]interface{}) (interface{}, error) {
		cfg := swagger_models.ModuleDetail{}
		var err error
		cfg.Environment = resourcedata.GetStrMap(d, "environment")
		cfg.ModuleType, err = appModuleTypePtr(resourcedata.GetStr(d, "module_type"))
		if err != nil {
			return nil, err
		}
		cfg.Routes = resourcedata.GetStrMap(d, "routes")
		cfg.TwinDetail = resourcedata.GetStr(d, "twin_detail")
		return &cfg, nil
	}
	val, err := resourcedata.GetStructPtr(d, "module", entryFunc)
	if val == nil || err != nil {
		return nil, err
	}
	return val.(*swagger_models.ModuleDetail), nil
}

func rdAppManifestInterfaceAclMatches(d map[string]interface{}) ([]*swagger_models.Match, error) {
	entryFunc := func(d map[string]interface{}) (*swagger_models.Match, error) {
		return &swagger_models.Match{
			Type:  resourcedata.GetStr(d, "type"),
			Value: resourcedata.GetStr(d, "value"),
		}, nil
	}
	cfgList := make([]*swagger_models.Match, 0)
	rdList := rdEntryList(d, "match")
	if len(rdList) == 0 {
		return cfgList, nil
	}
	for _, val := range rdList {
		rdEntry := val.(map[string]interface{})
		cfgEntry, err := entryFunc(rdEntry)
		if err != nil {
			return cfgList, err
		}
		cfgList = append(cfgList, cfgEntry)
	}
	return cfgList, nil
}

func rdAppManifestIntfAclActionMapParams(d map[string]interface{}) (
	*swagger_models.MapParams, error) {
	entryFunc := func(d map[string]interface{}) (interface{}, error) {
		return &swagger_models.MapParams{
			AppPort: rdEntryInt64(d, "port"),
		}, nil
	}
	val, err := resourcedata.GetStructPtr(d, "mapparam", entryFunc)
	if val == nil || err != nil {
		return nil, err
	}
	return val.(*swagger_models.MapParams), nil
}

func rdAppManifestIntfAclActionLimitParams(d map[string]interface{}) (
	*swagger_models.LimitParams, error) {
	entryFunc := func(d map[string]interface{}) (interface{}, error) {
		return &swagger_models.LimitParams{
			Limitburst: rdEntryInt64(d, "limitburst"),
			Limitrate:  rdEntryInt64(d, "limitrate"),
			Limitunit:  resourcedata.GetStr(d, "limitunit"),
		}, nil
	}
	val, err := resourcedata.GetStructPtr(d, "limit_param", entryFunc)
	if val == nil || err != nil {
		return nil, err
	}
	return val.(*swagger_models.LimitParams), nil
}

func rdAppManifestInterfaceAclActions(d map[string]interface{}) (
	[]*swagger_models.ACLAction, error) {
	entryFunc := func(d map[string]interface{}) (*swagger_models.ACLAction, error) {
		cfg := swagger_models.ACLAction{}
		var err error
		cfg.Drop = rdEntryBool(d, "drop")
		cfg.Limit = rdEntryBool(d, "limit")
		cfg.LimitValue, err = rdAppManifestIntfAclActionLimitParams(d)
		if err != nil {
			return nil, err
		}
		cfg.Portmap = rdEntryBool(d, "portmap")
		cfg.Portmapto, err = rdAppManifestIntfAclActionMapParams(d)
		return &cfg, nil
	}
	cfgList := make([]*swagger_models.ACLAction, 0)
	rdList := rdEntryList(d, "action")
	if len(rdList) == 0 {
		return cfgList, nil
	}
	for _, val := range rdList {
		rdEntry := val.(map[string]interface{})
		cfgEntry, err := entryFunc(rdEntry)
		if err != nil {
			return cfgList, err
		}
		cfgList = append(cfgList, cfgEntry)
	}
	return cfgList, nil
}

func rdAppManifestInterfaceAcls(d map[string]interface{}) ([]*swagger_models.ACL, error) {
	entryFunc := func(d map[string]interface{}) (*swagger_models.ACL, error) {
		cfg := swagger_models.ACL{}
		var err error
		cfg.Actions, err = rdAppManifestInterfaceAclActions(d)
		if err != nil {
			return nil, err
		}
		cfg.Matches, err = rdAppManifestInterfaceAclMatches(d)
		if err != nil {
			return nil, err
		}
		cfg.Name = resourcedata.GetStr(d, "name")
		return &cfg, nil
	}
	cfgList := make([]*swagger_models.ACL, 0)
	rdList := rdEntryList(d, "acl")
	if len(rdList) == 0 {
		return cfgList, nil
	}
	for _, val := range rdList {
		rdEntry := val.(map[string]interface{})
		cfgEntry, err := entryFunc(rdEntry)
		if err != nil {
			return cfgList, err
		}
		cfgList = append(cfgList, cfgEntry)
	}
	return cfgList, nil
}

func rdAppManifestInterfacesEntry(d map[string]interface{}) (
	*swagger_models.Interface, error) {
	cfg := &swagger_models.Interface{}
	var err error
	cfg.Acls, err = rdAppManifestInterfaceAcls(d)
	if err != nil {
		return nil, err
	}
	cfg.Directattach = rdEntryBool(d, "ignorepurge")
	cfg.Name = resourcedata.GetStr(d, "imagename")
	cfg.Optional = rdEntryBool(d, "preserve")
	cfg.Privateip = rdEntryBool(d, "readonly")
	cfg.Type = resourcedata.GetStr(d, "target")
	return cfg, nil
}

func rdAppManifestInterfaces(d map[string]interface{}) ([]*swagger_models.Interface, error) {
	cfgList := make([]*swagger_models.Interface, 0)
	rdList := rdEntryList(d, "interface")
	if len(rdList) == 0 {
		return cfgList, nil
	}
	for _, val := range rdList {
		rdEntry := val.(map[string]interface{})
		cfgEntry, err := rdAppManifestInterfacesEntry(rdEntry)
		if err != nil {
			return cfgList, err
		}
		cfgList = append(cfgList, cfgEntry)
	}
	return cfgList, nil
}

func rdAppManifestImageParams(d map[string]interface{}) ([]*swagger_models.Param, error) {
	entryFunc := func(d map[string]interface{}) (*swagger_models.Param, error) {
		return &swagger_models.Param{
			Name:  resourcedata.GetStr(d, "name"),
			Value: resourcedata.GetStr(d, "value"),
		}, nil
	}
	cfgList := make([]*swagger_models.Param, 0)
	rdList := rdEntryList(d, "param")
	if len(rdList) == 0 {
		return cfgList, nil
	}
	for _, val := range rdList {
		rdEntry := val.(map[string]interface{})
		cfgEntry, err := entryFunc(rdEntry)
		if err != nil {
			return cfgList, err
		}
		cfgList = append(cfgList, cfgEntry)
	}
	return cfgList, nil
}

func rdAppManifestImagesEntry(d map[string]interface{}) (
	*swagger_models.VMManifestImage, error) {
	cfg := swagger_models.VMManifestImage{}
	var err error
	cfg.Cleartext = rdEntryBool(d, "cleartext")
	cfg.Drvtype = resourcedata.GetStr(d, "drvtype")
	cfg.Ignorepurge = rdEntryBool(d, "ignorepurge")
	cfg.Imagename = resourcedata.GetStr(d, "imagename")
	cfg.Maxsize = strconv.FormatInt(rdEntryInt64(d, "maxsize"), 10)
	cfg.Mountpath = resourcedata.GetStr(d, "mountpath")
	cfg.Params, err = rdAppManifestImageParams(d)
	if err != nil {
		return nil, err
	}
	cfg.Preserve = rdEntryBool(d, "preserve")
	cfg.Readonly = rdEntryBool(d, "readonly")
	cfg.Target = resourcedata.GetStr(d, "target")
	cfg.Volumelabel = resourcedata.GetStr(d, "volumelabel")
	return &cfg, nil
}

func rdAppManifestImages(d map[string]interface{}) ([]*swagger_models.VMManifestImage, error) {
	cfgList := make([]*swagger_models.VMManifestImage, 0)
	rdList := rdEntryList(d, "image")
	if len(rdList) == 0 {
		return cfgList, nil
	}
	for _, val := range rdList {
		rdEntry := val.(map[string]interface{})
		cfgEntry, err := rdAppManifestImagesEntry(rdEntry)
		if err != nil {
			return cfgList, err
		}
		cfgList = append(cfgList, cfgEntry)
	}
	return cfgList, nil
}

func appCategoryPtr(strVal string) (*swagger_models.AppCategory, error) {
	val := swagger_models.AppCategory(strVal)
	if err := val.Validate(nil); err != nil {
		return nil, err
	}
	return &val, nil
}

func rdAppManifestDetailsEntry(d map[string]interface{}) (interface{}, error) {
	var err error
	cfg := swagger_models.Details{}
	cfg.AgreementList = resourcedata.GetStrMap(d, "agreement_list")
	cfg.AppCategory, err = appCategoryPtr(resourcedata.GetStr(d, "app_category"))
	if err != nil {
		return nil, err
	}
	cfg.Category = resourcedata.GetStrPtrOrNil(d, "category")
	cfg.LicenseList = resourcedata.GetStrMap(d, "license_list")
	cfg.Logo = resourcedata.GetStrMap(d, "logo")
	cfg.Os = resourcedata.GetStr(d, "os")
	cfg.Support = resourcedata.GetStr(d, "support")
	return &cfg, nil
}

func rdAppManifestDetails(d map[string]interface{}) (*swagger_models.Details, error) {
	val, err := resourcedata.GetStructPtr(d, "desc_detail", rdAppManifestDetailsEntry)
	if val == nil || err != nil {
		return nil, err
	}
	return val.(*swagger_models.Details), nil
}

func rdAppManifestContainerDetail(d map[string]interface{}) (
	*swagger_models.ContainerDetail, error) {
	entryFunc := func(d map[string]interface{}) (interface{}, error) {
		cfg := swagger_models.ContainerDetail{}
		cfg.ContainerCreateOption = resourcedata.GetStr(d, "container_create_option")
		return &cfg, nil
	}
	val, err := resourcedata.GetStructPtr(d, "container_detail", entryFunc)
	if val == nil || err != nil {
		return nil, err
	}
	return val.(*swagger_models.ContainerDetail), nil
}

func rdAppManifestUserDataTemplate(d map[string]interface{}) (
	*swagger_models.UserDataTemplate, error) {

	entryFunc := func(d map[string]interface{}) (interface{}, error) {
		cfg := swagger_models.UserDataTemplate{}
		val, err := resourcedata.GetStructPtr(d, "custom_config", rdAppInstCustomConfig)
		if val == nil || err != nil {
			return nil, err
		}
		cfg.CustomConfig = val.(*swagger_models.CustomConfig)
		return &cfg, nil
	}
	val, err := resourcedata.GetStructPtr(d, "configuration", entryFunc)
	if val == nil || err != nil {
		return nil, err
	}
	return val.(*swagger_models.UserDataTemplate), nil
}

func rdAppManifestEntry(d map[string]interface{}) (interface{}, error) {
	var err error
	cfg := swagger_models.VMManifest{}
	cfg.AppType = appTypePtr(resourcedata.GetStr(d, "apptype"))
	cfg.Configuration, err = rdAppManifestUserDataTemplate(d)
	if err != nil {
		return nil, err
	}
	cfg.ContainerDetail, err = rdAppManifestContainerDetail(d)
	if err != nil {
		return nil, err
	}
	cfg.DeploymentType, err = deploymentTypePtr(resourcedata.GetStr(d, "deployment_type"))
	if err != nil {
		return nil, err
	}
	cfg.Desc, err = rdAppManifestDetails(d)
	if err != nil {
		return nil, err
	}
	cfg.Description = resourcedata.GetStr(d, "description")
	cfg.DisplayName = resourcedata.GetStr(d, "display_name")
	cfg.Enablevnc = rdEntryBool(d, "enablevnc")
	cfg.Images, err = rdAppManifestImages(d)
	if err != nil {
		return nil, err
	}
	cfg.Interfaces, err = rdAppManifestInterfaces(d)
	if err != nil {
		return nil, err
	}
	cfg.Module, err = rdAppManifestModuleDetails(d)
	if err != nil {
		return nil, err
	}
	cfg.Name = resourcedata.GetStr(d, "name")
	cfg.Owner, err = rdAppManifestAuthor(d)
	if err != nil {
		return nil, err
	}
	cfg.Resources, err = rdAppManifestResources(d)
	if err != nil {
		return nil, err
	}
	cfg.Vmmode = resourcedata.GetStrPtrOrNil(d, "vmmode")
	return &cfg, nil
}

func rdAppManifest(rd *schema.ResourceData) (*swagger_models.VMManifest, error) {
	// Manifest can come through manifest or manifest_file.
	// Only one of them must be specified.
	ok, val := rdEntryByKey(rd, "manifest")
	if !ok {
		// Key manifest not present in Resource Data. Check for manifest_file
		return rdEntryAppManifestFromFile(rd)
	}
	ok, _ = rdEntryByKey(rd, "manifest_file")
	if ok {
		return nil, fmt.Errorf("Both manifest and manifest_file are specified. " +
			"Only one of them must be specified.")
	}
	val, err := resourcedata.GetStructPtr(rd, "manifest", rdAppManifestEntry)
	if val == nil || err != nil {
		return nil, err
	}
	return val.(*swagger_models.VMManifest), nil
}

func rdUpdateAppCfg(cfg *swagger_models.App, d *schema.ResourceData) error {
	cfg.Title = resourcedata.GetStrPtrOrNil(d, "title")
	cfg.Description = resourcedata.GetStr(d, "description")
	cfg.Cpus = rdEntryInt64(d, "cpus")
	cfg.Drives = rdEntryInt64(d, "drives")
	manifestJSON, err := rdAppManifest(d)
	if err != nil {
		return err
	}
	cfg.ManifestJSON = manifestJSON
	cfg.Networks = rdEntryInt64(d, "networks")
	cfg.Storage = rdEntryInt64(d, "storage")
	cfg.UserDefinedVersion = resourcedata.GetStr(d, "user_defined_version")
	return nil
}

// Create the Resource Group
func createEdgeAppResource(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client, ok := meta.(*client.Client)
	if !ok {
		return diag.Errorf("expect meta to be of type client.Client{} but is %T", meta)
	}

	name := resourcedata.GetStr(d, "name")
	id := resourcedata.GetStr(d, "id")
	errMsgPrefix := fmt.Sprintf("[ERROR] Edge App %s (id: %s) Create Failed.",
		name, id)
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
		log.Printf("[ERROR] Edge App Create Failed: %s, %s Err: %s",
			name, errMsgPrefix, err.Error())
		return diag.Errorf("%s Err: %s", errMsgPrefix, err.Error())
	}
	log.Printf("Edge App %s (ID: %s) Successfully created\n",
		rspData.ObjectName, rspData.ObjectID)
	d.SetId(rspData.ObjectID)
	return diags
}

// Update the Resource Group
func updateEdgeAppResource(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	name := resourcedata.GetStr(d, "name")
	id := resourcedata.GetStr(d, "id")
	errMsgPrefix := getErrMsgPrefix(name, id, "Edge App", "Update")
	if client == nil {
		return diag.Errorf("%s nil Client", errMsgPrefix)
	}
	cfg, err, _ := getEdgeApp(client, name, id)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	err = checkInvalidAttrForUpdate(d, *cfg.Name, cfg.ID)
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
func deleteEdgeAppResource(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := (meta.(Client)).Client
	name := resourcedata.GetStr(d, "name")
	id := resourcedata.GetStr(d, "id")
	errMsgPrefix := getErrMsgPrefix(name, id, "Edge App", "Delete")
	cfg, err, httpRsp := getEdgeApp(client, name, id)
	if err != nil {
		if httpRsp != nil && zedcloudapi.IsObjectNotFound(httpRsp) {
			log.Printf("PROVIDER:  EdgeApp %s (id: %s) Not Found", id, name)
			return diags
		}
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
func readResourceEdgeApp(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	return readEdgeApp(ctx, d, meta)
}
