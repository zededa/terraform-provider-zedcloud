package resources

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	apiclient "github.com/zededa/terraform-provider/client"
	"github.com/zededa/terraform-provider/client/edge_node_configuration"
	"github.com/zededa/terraform-provider/models"
	zschema "github.com/zededa/terraform-provider/schemas"
)

// Note, an edge-node and a device-config are the same thing. Due is inconcistency in the API
// definition both terms are used interchangeably in the resulting generated code,.

func EdgeNodeConfiguration() *schema.Resource {
	return &schema.Resource{
		ReadContext:   GetEdgeNode,
		CreateContext: CreateEdgeNode,
		UpdateContext: UpdateEdgeNode,
		DeleteContext: DeleteEdgeNode,
		Schema:        zschema.DeviceConfigSchema(),
	}
}

func DataResourceEdgeNodeConfiguration() *schema.Resource {
	return &schema.Resource{
		ReadContext: GetEdgeNode,
		Schema:      zschema.DeviceConfigSchema(),
	}
}

func GetEdgeNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fmt.Println("------GET---------------------------------------")
	deviceConfig, diags := getEdgeNode(ctx, d, m)
	if len(diags) > 0 {
		return diags
	}
	zschema.SetDeviceConfigResourceData(d, deviceConfig)
	d.SetId(deviceConfig.ID)
	fmt.Println("------END GET---------------------------------------")

	return diags
}

func UpdateEdgeNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	fmt.Println("------POST---------------------------------------")
	// we need to fetch the edge-node/device-config and add/change the fields according to config
	deviceConfig, diags := getEdgeNode(ctx, d, m)
	if len(diags) > 0 {
		return diags
	}
	// publish the api response to local state and the d instance
	zschema.SetDeviceConfigResourceData(d, deviceConfig)

	params := edge_node_configuration.NewEdgeNodeConfigurationUpdateEdgeNodeParams()
	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}
	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		params.ID = idVal.(string)
	} else {
		return diag.Errorf("missing client parameter: id")
	}
	params.SetBody(zschema.DeviceConfigModel(d))

	if err := os.WriteFile("/tmp/req-update", []byte("==========REQ=============\n"+spew.Sdump(params)), 0644); err != nil {
		fmt.Println(err)
	}

	// makes a bulk update for all properties that were changed
	client := m.(*apiclient.Zedcloudapi)
	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationUpdateEdgeNode(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.Ec != nil && *err.Ec == models.ZsrvErrorCodeZMsgSuccess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if len(diags) > 0 {
			return diags
		}
	}

	// due to api design, we need to fetch the newly created edge-node/device-config
	deviceConfig, diags = getEdgeNode(ctx, d, m)
	if len(diags) > 0 {
		return diags
	}
	// publish the api response to local state and the d instance
	zschema.SetDeviceConfigResourceData(d, deviceConfig)
	d.SetId(deviceConfig.ID)

	// to set base-image the api requires separate requests
	if diags := setBaseImage(ctx, d, m, params.Body.BaseImage, deviceConfig.BaseImage); len(diags) > 0 {
		return diags
	}

	// to set admin-state the api requires separate requests
	if diags := setAdminState(ctx, d, m, params.Body.AdminState, deviceConfig.AdminState); len(diags) > 0 {
		return diags
	}

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := GetEdgeNode(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	fmt.Println("------END POST---------------------------------------")
	return diags
}

func CreateEdgeNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.DeviceConfigModel(d)
	params := edge_node_configuration.NewEdgeNodeConfigurationCreateEdgeNodeParams()
	params.SetBody(model)

	fmt.Println("------CREATE---------------------------------------")
	if err := os.WriteFile("/tmp/req", []byte("==========REQ=============\n"+spew.Sdump(params)), 0644); err != nil {
		fmt.Println(err)
	}

	client := m.(*apiclient.Zedcloudapi)

	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationCreateEdgeNode(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.Ec != nil && *err.Ec == models.ZsrvErrorCodeZMsgSuccess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if len(diags) > 0 {
			return diags
		}
	}

	// due to api design, we need to fetch the newly created edge-node/device-config
	deviceConfig, diags := getEdgeNode(ctx, d, m)
	if len(diags) > 0 {
		return diags
	}
	// publish the api response to local state and the d instance
	zschema.SetDeviceConfigResourceData(d, deviceConfig)
	d.SetId(deviceConfig.ID)

	// to set base-image the api requires separate requests
	if diags := setBaseImage(ctx, d, m, params.Body.BaseImage, deviceConfig.BaseImage); len(diags) > 0 {
		return diags
	}

	// to set admin-state the api requires separate requests
	if diags := setAdminState(ctx, d, m, params.Body.AdminState, deviceConfig.AdminState); len(diags) > 0 {
		return diags
	}

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := GetEdgeNode(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	fmt.Println("------END CREATE---------------------------------------")
	return diags
}

func DeleteEdgeNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fmt.Println("------DELETE---------------------------------------")
	var diags diag.Diagnostics

	params := edge_node_configuration.NewEdgeNodeConfigurationDeleteEdgeNodeParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*apiclient.Zedcloudapi)

	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationDeleteEdgeNode(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	d.SetId("")
	fmt.Println("------DONE DELETE---------------------------------------")
	return diags
}

func activateEdgeNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := edge_node_configuration.NewEdgeNodeConfigurationActivateEdgeNodeParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*apiclient.Zedcloudapi)
	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationActivateEdgeNode(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.Ec != nil && *err.Ec == models.ZsrvErrorCodeZMsgSuccess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if len(diags) > 0 {
			return diags
		}
	}

	return diags
}

func deactivateEdgeNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := edge_node_configuration.NewEdgeNodeConfigurationDeActivateEdgeNodeParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*apiclient.Zedcloudapi)
	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationDeActivateEdgeNode(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.Ec != nil && *err.Ec == models.ZsrvErrorCodeZMsgSuccess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if len(diags) > 0 {
			return diags
		}
	}

	return diags
}

func applyBaseOS(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := edge_node_configuration.NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	// makes a bulk update for all properties that were changed
	client := m.(*apiclient.Zedcloudapi)
	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationUpdateEdgeNodeBaseOS(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		return diags
	}

	return diags
}

func publishBaseOS(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := edge_node_configuration.NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	// makes a bulk update for all properties that were changed
	client := m.(*apiclient.Zedcloudapi)
	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationUpdateEdgeNodeBaseOS2(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		return diags
	}

	return diags
}

// to set the base-os-image, the api requires several requests.
// 1. set the image config with a create or udpate request
// 2. publish the image config
// 3. apply the image config
func setBaseImage(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
	localImages, remoteImages []*models.BaseOSImage,
) diag.Diagnostics {

	fmt.Println("++++++++++++++++++++++++set image+++++++++++++++++++++++++++++++++++")

	if len(localImages) == 0 {
		return nil
	}
	if len(localImages) > 1 {
		return diag.Errorf("expect exactly one base image definition but got: %d", len(localImages))
	}
	if localImages[0] == nil {
		return diag.FromErr(errors.New("expect non-nil base image definition"))
	}
	if localImages[0].Version == nil {
		return diag.FromErr(errors.New("expect version to be set in base image definition"))
	}

	// image in tf-config file
	localImageVersion := *(localImages[0].Version)

	// image in api response
	var remoteImageName string
	var remoteImageIsActive bool
	if len(remoteImages) == 1 && remoteImages[0] != nil {
		if remoteImages[0].ImageName != nil {
			remoteImageName = *(remoteImages[0].ImageName)
		}
		if remoteImages[0].Activate != nil {
			remoteImageIsActive = *(remoteImages[0].Activate)
		}
	}

	// do not update if local config euqals api config
	// for some reason, we compare version against name
	if localImageVersion == remoteImageName && remoteImageIsActive {
		return nil
	}

	fmt.Println("++++++++++++++++++++++++ image data +++++++++++++++++++++++++++++++++++")
	spew.Dump(remoteImages[0])

	if diags := publishBaseOS(ctx, d, m); len(diags) != 0 {
		return diags
	}
	if diags := applyBaseOS(ctx, d, m); len(diags) != 0 {
		return diags
	}

	fmt.Println("++++++++++++++++++++++++ done set image +++++++++++++++++++++++++++++++++++")
	return nil
}

func setAdminState(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
	localAdminState, remoteAdminState *models.AdminState,
) diag.Diagnostics {

	// no config
	if localAdminState == nil {
		return nil
	}

	// config same as api state
	if remoteAdminState != nil {
		if *localAdminState == *remoteAdminState {
			return nil
		}
	}

	// states differ
	fmt.Println("====================================")
	fmt.Println(*localAdminState)
	fmt.Println(*remoteAdminState)
	fmt.Println("====================================")

	if *localAdminState == models.ADMINSTATE_ACTIVE {
		// do not activate if already registered
		if *remoteAdminState == models.ADMINSTATE_REGISTERED {
			return nil
		}
		if diags := activateEdgeNode(ctx, d, m); len(diags) != 0 {
			return diags
		}
	}
	if *localAdminState == models.ADMINSTATE_INACTIVE {
		if diags := deactivateEdgeNode(ctx, d, m); len(diags) != 0 {
			return diags
		}
	}

	return nil
}

func getEdgeNode(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.DeviceConfig, diag.Diagnostics) {
	var diags diag.Diagnostics
	fmt.Println("------get---------------------------------------")

	params := edge_node_configuration.NewEdgeNodeConfigurationGetEdgeNodeByNameParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	nameVal, nameIsSet := d.GetOk("name")
	if nameIsSet {
		params.Name = nameVal.(string)
	} else {
		return nil, append(diags, diag.Errorf("missing client parameter: name")...)
	}

	client := m.(*apiclient.Zedcloudapi)

	resp, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationGetEdgeNodeByName(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return nil, append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	deviceConfig := resp.GetPayload()
	if err := os.WriteFile("/tmp/req", []byte("==========RESP=============\n"+spew.Sdump(deviceConfig)), 0644); err != nil {
		fmt.Println(err)
	}
	fmt.Println("------END get---------------------------------------")

	return deviceConfig, diags
}
