package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
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

	newNode := zschema.NodeModel(d)
	params := config.CreateParams()
	params.SetBody(newNode)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.Node.Create(params, nil)
	if err != nil {
		log.Printf("[TRACE] edge node create error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("edge node create error: %s", err)...)
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

	// Fetch the newly created node
	createdNode, diags := readNodeByName(ctx, d, m)
	if diags.HasError() {
		return diags
	}

	// Set ID so subsequent operations can use it
	d.SetId(createdNode.ID)

	// Check if base image needs to be published and applied
	if len(newNode.BaseImage) > 0 {
		// For creation, always publish and apply if base image is specified
		if diags := publishBaseOS(ctx, d, client, newNode); len(diags) > 0 {
			return diags
		}
		if diags := applyBaseOS(ctx, d, client, newNode); len(diags) > 0 {
			return diags
		}
	}

	// Check if admin state needs to be set
	if newNode.AdminState != nil {
		if diags := setAdminState(ctx, d, m, newNode.AdminState, createdNode.AdminState); len(diags) > 0 {
			return diags
		}
	}

	// Fetch latest node data and update state
	finalNode, diags := readNodeByID(ctx, d, m)
	if diags.HasError() {
		return diags
	}

	zschema.SetNodeResourceData(d, finalNode)
	d.SetId(finalNode.ID)

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

	// Fetch existing node data
	existingNode, diags := readNodeByID(ctx, d, m)
	if diags.HasError() {
		return diags
	}

	// Prepare update parameters
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
	newNode := zschema.NodeModel(d)
	params.SetBody(newNode)

	// Call node update
	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.Node.Update(params, nil)
	if err != nil {
		log.Printf("[TRACE] edge node update error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("edge node update error: %s", err)...)
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

	// Check for base image changes
	if len(newNode.BaseImage) > 0 {
		hasBaseImageChange := false

		if len(existingNode.BaseImage) == 0 {
			hasBaseImageChange = true
		} else if len(newNode.BaseImage) > 0 && len(existingNode.BaseImage) > 0 {
			newVersion := newNode.BaseImage[0].Version
			existingImageName := existingNode.BaseImage[0].ImageName
			existingActivate := existingNode.BaseImage[0].Activate

			if newVersion != nil && existingImageName != nil {
				if *newVersion != *existingImageName {
					hasBaseImageChange = true
				} else if existingActivate != nil && !*existingActivate {
					hasBaseImageChange = true
				}
			}
		}

		if hasBaseImageChange {
			if diags := publishBaseOS(ctx, d, client, newNode); len(diags) > 0 {
				return diags
			}
			if diags := applyBaseOS(ctx, d, client, newNode); len(diags) > 0 {
				return diags
			}
		}
	}

	// Check for admin state changes
	if newNode.AdminState != nil && existingNode.AdminState != nil {
		if *newNode.AdminState != *existingNode.AdminState {
			if diags := setAdminState(ctx, d, m, newNode.AdminState, existingNode.AdminState); len(diags) > 0 {
				return diags
			}
		}
	} else if newNode.AdminState != nil && existingNode.AdminState == nil {
		if diags := setAdminState(ctx, d, m, newNode.AdminState, nil); len(diags) > 0 {
			return diags
		}
	}

	// Fetch latest node data and update state
	updatedNode, diags := readNodeByID(ctx, d, m)
	if diags.HasError() {
		return diags
	}

	zschema.SetNodeResourceData(d, updatedNode)
	d.SetId(updatedNode.ID)

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

	_, err := client.Node.Delete(params, nil)
	if err != nil {
		log.Printf("[TRACE] edge node delete error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("edge node delete error: %s", err)...)
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

func applyBaseOS(ctx context.Context, d *schema.ResourceData, client *api_client.ZedcloudAPI, edgeNode *models.Node) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.UpdateBaseOSParams()
	params.ID = edgeNode.ID
	params.SetBody(edgeNode)
	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	resp, err := client.Node.UpdateBaseOS(params, nil)
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

func publishBaseOS(ctx context.Context,
	d *schema.ResourceData,
	client *api_client.ZedcloudAPI,
	edgeNode *models.Node) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.PublishBaseOSParams()
	params.ID = edgeNode.ID
	params.SetBody(edgeNode)
	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	resp, err := client.Node.PublishBaseOS(params, nil)
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

// to set the base-os-image, the api requires several requests.
// 1. set the image config with a create or udpate request
// 2. publish the image config
// 3. apply the image config
func setBaseImage(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
	edgeNode *models.Node,
	remoteImages []*models.BaseOSImage,
) diag.Diagnostics {
	localImages := edgeNode.BaseImage
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

	client := m.(*api_client.ZedcloudAPI)
	if diags := publishBaseOS(ctx, d, client, edgeNode); len(diags) != 0 {
		return diags
	}
	if diags := applyBaseOS(ctx, d, client, edgeNode); len(diags) != 0 {
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
	if err != nil {
		log.Printf("[TRACE] edge node read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return nil, diags
		}

		diags = append(diags, diag.Errorf("edge node read error: %s", err)...)
		return nil, diags
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
	if err != nil {
		log.Printf("[TRACE] edge node read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return nil, diags
		}

		diags = append(diags, diag.Errorf("edge node read error: %s", err)...)
		return nil, diags
	}

	edgeNode := resp.GetPayload()

	return edgeNode, diags
}
