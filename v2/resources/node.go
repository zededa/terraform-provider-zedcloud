package resources

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/go-cty/cty"
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
		CustomizeDiff: customizeNodeDiff,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// customizeNodeDiff runs every plan-time validation for an edge node. The SDK
// takes a single CustomizeDiff, and helper/customdiff is not vendored, so the
// checks are chained here.
func customizeNodeDiff(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
	if err := validateBondMemberInterfaces(ctx, d, m); err != nil {
		return err
	}
	return validateClusteredNodeBaseImage(ctx, d, m)
}

// validateClusteredNodeBaseImage rejects a per-node base_image on a node that
// belongs to an edge-node cluster.
//
// Once a node is a cluster member the controller refuses a per-device base
// image outright -- devBaseImageApply answers
// "device: <name> is in cluster, device base image cannot be applied" with a
// 400 -- because EVE-OS has to be rolled out across the cluster one node at a
// time so workloads can migrate. Without this check the mistake only surfaces
// as that opaque 400 halfway through an apply, so fail at plan time and say
// where the setting actually belongs (CI-836).
func validateClusteredNodeBaseImage(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
	if !rawConfigHasAttr(d.GetRawConfig(), "base_image") {
		return nil
	}

	clusterID := clusteredNodeID(d)
	if clusterID == "" {
		return nil
	}

	name, _ := d.Get("name").(string)
	if name == "" {
		name = d.Id()
	}

	return fmt.Errorf(
		"edge node %q is a member of edge-node cluster %s, so base_image cannot be set on it: "+
			"the controller rejects a per-node base image for a cluster member. "+
			"Set base_image on the zedcloud_edgenode_cluster resource instead, which upgrades the "+
			"member nodes one at a time and migrates workloads between them",
		name, clusterID)
}

// clusteredNodeID returns the id of the edge-node cluster this node belongs to,
// or "" when it belongs to none. Membership is established by the controller
// when the cluster object names the node, so it is read from the merged diff
// rather than from the config.
func clusteredNodeID(d *schema.ResourceDiff) string {
	raw, isSet := d.GetOk("edge_node_cluster")
	if !isSet {
		return ""
	}
	items, isList := raw.([]interface{})
	if !isList || len(items) == 0 || items[0] == nil {
		return ""
	}
	entry, isMap := items[0].(map[string]interface{})
	if !isMap {
		return ""
	}
	id, _ := entry["id"].(string)
	return id
}

// rawConfigHasAttr reports whether the user's configuration actually sets the
// attribute. It deliberately reads the raw config and not the merged plan:
// after a cluster upgrade the controller puts base_image on the device object
// and the provider refreshes it into state, and that state value must not be
// mistaken for something the user asked for.
func rawConfigHasAttr(rawConfig cty.Value, attr string) bool {
	if rawConfig.IsNull() || !rawConfig.IsKnown() || !rawConfig.Type().IsObjectType() {
		return false
	}
	if !rawConfig.Type().HasAttribute(attr) {
		return false
	}
	v := rawConfig.GetAttr(attr)
	if v.IsNull() || !v.IsKnown() {
		return false
	}
	if v.CanIterateElements() {
		return v.LengthInt() > 0
	}
	return true
}

// validateBondMemberInterfaces rejects configurations that explicitly declare
// netname/netid on a bond member interface. The API refuses such interfaces
// with a generic 400, so fail at plan time with an actionable message instead.
// It inspects the raw config (not the merged plan) because netid is a Computed
// attribute whose stale state value is carried forward when absent from the
// config; that carried value is dropped when building the request payload and
// must not trigger this error (ENG-2782).
func validateBondMemberInterfaces(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
	return checkBondMemberInterfaces(d.GetRawConfig())
}

func checkBondMemberInterfaces(rawConfig cty.Value) error {
	if rawConfig.IsNull() || !rawConfig.IsKnown() || !rawConfig.Type().IsObjectType() || !rawConfig.Type().HasAttribute("interfaces") {
		return nil
	}
	interfaces := rawConfig.GetAttr("interfaces")
	if interfaces.IsNull() || !interfaces.IsKnown() || !interfaces.CanIterateElements() {
		return nil
	}
	for it := interfaces.ElementIterator(); it.Next(); {
		_, iface := it.Element()
		if iface.IsNull() || !iface.IsKnown() {
			continue
		}
		if rawConfigString(iface, "intf_usage") != string(models.AdapterUsageADAPTERUSAGEBONDMEMBER) {
			continue
		}
		if rawConfigString(iface, "netname") != "" || rawConfigString(iface, "netid") != "" {
			return fmt.Errorf(
				"interface %q: netname/netid must not be set when intf_usage is %s, a bond member interface cannot carry a network identity",
				rawConfigString(iface, "intfname"),
				models.AdapterUsageADAPTERUSAGEBONDMEMBER,
			)
		}
	}
	return nil
}

func rawConfigString(v cty.Value, attr string) string {
	if !v.Type().IsObjectType() || !v.Type().HasAttribute(attr) {
		return ""
	}
	attrVal := v.GetAttr(attr)
	if attrVal.IsNull() || !attrVal.IsKnown() || attrVal.Type() != cty.String {
		return ""
	}
	return attrVal.AsString()
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
		if diags := rejectBaseImageOnClusterMember(createdNode); diags.HasError() {
			return diags
		}
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
	// Use the revision we just fetched from the server as the optimistic-lock
	// token instead of the (possibly stale) value cached in Terraform state.
	// The device version can advance out-of-band between Terraform's last
	// refresh and this update; sending the stale revision makes the API reject
	// the update with HTTP 409 "Version mismatch".
	newNode.Revision = existingNode.Revision
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
			newImageName := newNode.BaseImage[0].ImageName
			existingImageName := existingNode.BaseImage[0].ImageName
			existingActivate := existingNode.BaseImage[0].Activate

			if newImageName != nil && existingImageName != nil {
				if *newImageName != *existingImageName {
					hasBaseImageChange = true
				} else if existingActivate != nil && !*existingActivate {
					hasBaseImageChange = true
				}
			}
		}

		if hasBaseImageChange {
			if diags := rejectBaseImageOnClusterMember(existingNode); diags.HasError() {
				return diags
			}
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
		if isStatusNotFound(err) {
			d.SetId("")
			return diags
		}
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

// rejectBaseImageOnClusterMember is the apply-time half of
// validateClusteredNodeBaseImage. The plan-time check reads the config, which
// covers the normal path; this one reads what the API just told us, and catches
// the cases the plan cannot see -- a node that joined a cluster between the
// last refresh and this apply, or a targeted apply that skipped the refresh.
//
// Without it the request still fails, just later and with the controller's
// opaque 400 instead of an explanation.
func rejectBaseImageOnClusterMember(node *models.Node) diag.Diagnostics {
	var diags diag.Diagnostics

	if node == nil || node.EdgeNodeCluster == nil || node.EdgeNodeCluster.ID == "" {
		return diags
	}

	name := ""
	if node.Name != nil {
		name = *node.Name
	}

	return append(diags, diag.Errorf(
		"edge node %q is a member of edge-node cluster %s, so base_image cannot be applied to it: "+
			"the controller rejects a per-node base image for a cluster member. "+
			"Set base_image on the zedcloud_edgenode_cluster resource instead, which upgrades the "+
			"member nodes one at a time and migrates workloads between them",
		name, node.EdgeNodeCluster.ID)...)
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
