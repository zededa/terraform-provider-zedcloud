package resources

import (
	"context"
	"errors"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/node"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

// Note, an edge-node and a device-config are the same thing. Due is inconcistency in the API
// definition both terms are used interchangeably.

func NodeResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateNode,
		ReadContext:   ReadNode,
		UpdateContext: UpdateNode,
		DeleteContext: DeleteNode,
		Schema:        zschema.Node(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func NodeDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: ReadNode,
		Schema:      zschema.Node(),
	}
}

func CreateNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.NodeModel(d)
	params := config.CreateParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.Node.Create(params, nil)
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
			if err.ErrorCode != nil && *err.ErrorCode == models.ErrorCodeSuccess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if diags.HasError() {
			return diags
		}
	}

	// due to api design, we need to fetch the newly created edge-node/edgeNode-config
	edgeNode, diags := readNodeByName(ctx, d, m)
	if diags.HasError() {
		return diags
	}
	// publish the api response to local state and the d instance
	zschema.SetNodeResourceData(d, edgeNode)
	d.SetId(edgeNode.ID)

	// to set base-image the api requires separate requests
	if diags := setBaseImage(ctx, d, m, params.Body.BaseImage, edgeNode.BaseImage); diags.HasError() {
		return diags
	}

	// to set admin-state the api requires separate requests
	if diags := setAdminState(ctx, d, m, params.Body.AdminState, edgeNode.AdminState); diags.HasError() {
		return diags
	}

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if diags := ReadNode(ctx, d, m); diags.HasError() {
		return diags
	}

	return diags
}

func ReadNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var edgeNode *models.Node
	var diags diag.Diagnostics

	if _, isSet := d.GetOk("name"); isSet {
		edgeNode, diags = readNodeByName(ctx, d, m)
	} else if _, isSet := d.GetOk("id"); isSet {
		edgeNode, diags = readNodeByID(ctx, d, m)
	}

	if diags.HasError() {
		return diags
	}

	zschema.SetNodeResourceData(d, edgeNode)
	d.SetId(edgeNode.ID)

	return diags
}

func UpdateNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.UpdateParams()
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
	params.SetBody(zschema.NodeModel(d))

	// makes a bulk update for all properties that were changed
	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.Node.Update(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.ErrorCode != nil && *err.ErrorCode == models.ErrorCodeSuccess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if diags.HasError() {
			return diags
		}
	}

	// due to api design, we need to fetch the newly created edge-node/device-config
	edgeNode, diags := readNodeByName(ctx, d, m)
	if diags.HasError() {
		return diags
	}
	// publish the api response to local state and the d instance
	zschema.SetNodeResourceData(d, edgeNode)
	d.SetId(edgeNode.ID)

	// to set base-image the api requires separate requests
	if diags := setBaseImage(ctx, d, m, params.Body.BaseImage, edgeNode.BaseImage); len(diags) > 0 {
		return diags
	}

	// to set admin-state the api requires separate requests
	if diags := setAdminState(ctx, d, m, params.Body.AdminState, edgeNode.AdminState); len(diags) > 0 {
		return diags
	}

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := ReadNode(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func DeleteNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.DeleteParams()

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

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.Node.Delete(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}

func activateNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.ActivationParams()

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

	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.Node.Activate(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.ErrorCode != nil && *err.ErrorCode == models.ErrorCodeSuccess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if diags.HasError() {
			return diags
		}
	}

	return diags
}

func deactivateNode(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.DeactivationParams()

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

	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.Node.Deactivate(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.ErrorCode != nil && *err.ErrorCode == models.ErrorCodeSuccess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if diags.HasError() {
			return diags
		}
	}

	return diags
}

func applyBaseOS(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.UpdateBaseOSParams()

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
	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.Node.UpdateBaseOS(params, nil)
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

	params := config.PublishBaseOSParams()

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
	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.Node.PublishBaseOSParams(params, nil)
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

	if diags := publishBaseOS(ctx, d, m); len(diags) != 0 {
		return diags
	}
	if diags := applyBaseOS(ctx, d, m); len(diags) != 0 {
		return diags
	}

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

	if *localAdminState == models.ADMINSTATE_ACTIVE {
		// do not activate if already registered
		if *remoteAdminState == models.ADMINSTATE_REGISTERED {
			return nil
		}
		if diags := activateNode(ctx, d, m); len(diags) != 0 {
			return diags
		}
	}
	if *localAdminState == models.ADMINSTATE_INACTIVE {
		if diags := deactivateNode(ctx, d, m); len(diags) != 0 {
			return diags
		}
	}

	return nil
}

func readNodeByID(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.Node, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := config.GetByIDParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	id, isSet := d.GetOk("id")
	if isSet {
		params.ID = id.(string)
	} else {
		return nil, append(diags, diag.Errorf("missing client parameter: id")...)
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.Node.GetByID(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return nil, append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	edgeNode := resp.GetPayload()

	return edgeNode, diags
}

func readNodeByName(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.Node, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := config.GetByNameParams()

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

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.Node.GetByName(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return nil, append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	edgeNode := resp.GetPayload()

	return edgeNode, diags
}
