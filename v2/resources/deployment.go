package resources

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	"github.com/zededa/terraform-provider-zedcloud/v2/client/deployment"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

// DeploymentResource returns the schema and methods for the Deployment resource.
func DeploymentResource() *schema.Resource {
	return &schema.Resource{
		Description: `Resource "zedcloud_deployment" manages deployments in ZedCloud. A deployment is a collection of policies applied to a device.
The controller versions deployments: every change creates a new version of the deployment with a new id, and only the latest
version of a project's deployment can be changed. Changing this resource in place does exactly that: the provider asks the
controller for the next version, waits until its policies are applied, and the resource then tracks the new version (its id changes).
Each version needs its own deployment_tag (the controller requires the tag to be unique per project), so a change must come with
a new deployment_tag. Edge nodes select a version by its tag: they keep running the version whose tag they carry until their
deployment_tag is moved to the new one, which allows a staged rollout. Older versions stay on the controller until the
resource is destroyed; destroy removes the current version and the older ones. Declaring a second zedcloud_deployment for a
project that already has one also creates a new version (kept for compatibility); in-place changes are the recommended way.`,
		CreateContext: CreateDeployment,
		DeleteContext: DeleteDeployment,
		ReadContext:   GetDeploymentByID,
		UpdateContext: UpdateDeployment,
		CustomizeDiff: customizeDeploymentDiff,
		Schema:        zschema.Deployment(),
		Timeouts: &schema.ResourceTimeout{
			Update: schema.DefaultTimeout(deploymentUpdateTimeout),
		},
		Importer: &schema.ResourceImporter{
			StateContext: deploymentImportStateContext,
		},
	}
}

// DeploymentDataSource returns the schema and read function for the
// zedcloud_deployment data source. It supports lookup by id or name; in
// both cases project_id is required. When multiple versions of a
// deployment share a name, the latest revision is returned.
func DeploymentDataSource() *schema.Resource {
	s := zschema.Deployment()
	s["id"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "system generated unique id for a deployment",
	}
	s["name"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "user defined name for the deployment",
	}
	return &schema.Resource{
		Description: `Data source "zedcloud_deployment" reads an existing deployment from ZedCloud. Look up by ` + "`id`" + ` or ` + "`name`" + ` within a ` + "`project_id`" + `.`,
		ReadContext: GetDeployment,
		Schema:      s,
	}
}

// GetDeployment is the read function for the zedcloud_deployment data
// source. Dispatches to id- or name-based lookup; project_id is required
// either way.
func GetDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	if _, ok := d.GetOk("project_id"); !ok {
		return append(diags, diag.Errorf("missing required parameter: project_id")...)
	}

	_, hasID := d.GetOk("id")
	_, hasName := d.GetOk("name")
	if !hasID && !hasName {
		return append(diags, diag.Errorf("one of `id` or `name` must be set")...)
	}

	if hasID {
		return GetDeploymentByID(ctx, d, m)
	}
	return getDeploymentByName(ctx, d, m)
}

// getDeploymentByName lists deployments in a project, filters by the
// configured name, picks the latest revision, then fetches the full
// deployment by id to populate state.
func getDeploymentByName(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*api_client.ZedcloudAPI)
	projectID := d.Get("project_id").(string)
	wantName := d.Get("name").(string)

	listParams := deployment.NewGetListbyIdParams()
	listParams.SetProjectID(projectID)

	listResp, err := client.Deployment.GetListByProjectID(listParams, nil)
	if err != nil {
		log.Printf("[TRACE] deployment list error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			return append(diags, ds...)
		}
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	var matchID string
	var matchVersion int
	for _, dep := range listResp.GetPayload().List {
		if dep.Name != wantName {
			continue
		}
		if dep.Revision == nil || dep.Revision.Curr == nil {
			continue
		}
		v, err := strconv.Atoi(*dep.Revision.Curr)
		if err != nil {
			continue
		}
		if matchID == "" || v > matchVersion {
			matchID = dep.ID
			matchVersion = v
		}
	}

	if matchID == "" {
		return append(diags, diag.Errorf("no deployment named %q found in project %s", wantName, projectID)...)
	}

	getParams := deployment.NewGetByIDParams()
	getParams.ID = matchID
	getParams.ProjectID = projectID
	resp, err := client.Deployment.GetByID(getParams, nil)
	if err != nil {
		log.Printf("[TRACE] deployment get by id error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			return append(diags, ds...)
		}
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	zschema.SetDeploymentResourceData(d, resp.GetPayload())
	d.SetId(matchID)
	return diags
}

// deploymentImportStateContext is a custom import function that accepts a
// composite ID in the form "project_id:deployment_id". Both parts are
// required because the GetByID API endpoint requires the project ID.
func deploymentImportStateContext(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	parts := strings.SplitN(d.Id(), ":", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return nil, fmt.Errorf(
			"invalid import ID %q: expected format \"project_id:deployment_id\"",
			d.Id(),
		)
	}
	projectID, deploymentID := parts[0], parts[1]
	if err := d.Set("project_id", projectID); err != nil {
		return nil, fmt.Errorf("failed to set project_id: %w", err)
	}
	d.SetId(deploymentID)
	return []*schema.ResourceData{d}, nil
}

// CreateDeployment creates a new deployment in ZedCloud.
func CreateDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*api_client.ZedcloudAPI)
	model := zschema.DeploymentModel(d)

	latestDeploymentID, err := getLatestProjectDeployment(ctx, client, model.ProjectID)
	if err != nil {
		log.Printf("[TRACE] deployment create error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}
		log.Printf("[ERROR] failed to get latest deployment for project: %s", err)
		diags = append(diags, diag.Errorf("failed to get latest deployment for project: %s", err)...)
		return diags
	}

	log.Printf("[TRACE] latest deployment ID: %s, Name: %s", latestDeploymentID, model.Name)
	var responseData *models.ZsrvResponse
	if latestDeploymentID == "" {
		params := deployment.NewCreateParams()
		params.SetBody(model)
		params.SetProjectID(model.ProjectID)
		resp, err := client.Deployment.Create(params, nil)
		if err != nil {
			log.Printf("[TRACE] deployment create error: %s", spew.Sdump(err))
			if ds, ok := ZsrvResponderToDiags(err); ok {
				diags = append(diags, ds...)
				return diags
			}
			diags = append(diags, diag.Errorf("unexpected: %s", err)...)
			return diags
		}

		responseData = resp.GetPayload()
	} else {
		params := deployment.NewCreateNewVersionParams()
		model.ID = latestDeploymentID
		params.SetBody(model)
		params.SetProjectID(model.ProjectID)
		params.SetID(latestDeploymentID)
		resp, err := client.Deployment.CreateNewVersion(params, nil)
		if err != nil {
			log.Printf("[TRACE] deployment create error: %s", spew.Sdump(err))
			if ds, ok := ZsrvResponderToDiags(err); ok {
				diags = append(diags, ds...)
				return diags
			}
			diags = append(diags, diag.Errorf("unexpected: %s", err)...)
			return diags
		}

		responseData = resp.GetPayload()
	}

	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.ErrorCode != nil && *err.ErrorCode == models.ErrorCodeAccepted {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if diags.HasError() {
			return diags
		}
	}
	// note, we need to set the ID before fetching the newly created project.
	d.SetId(responseData.ObjectID)

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := GetDeploymentByID(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}
	return diags
}

// DeleteDeployment deletes a deployment in ZedCloud.
func DeleteDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := deployment.NewDeleteParams()

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

	projectIdVal, projectIdIsSet := d.GetOk("project_id")
	if projectIdIsSet {
		params.ProjectID = projectIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: projectId")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	// Updates create a new version and leave the older ones on the
	// controller (nodes may still be using them). On destroy, remove the
	// versions this resource went through as well: while any exists, the
	// project cannot be deleted (409). Newer versions, created outside
	// Terraform, are left alone.
	var older []deploymentVersion
	if versions, err := listProjectDeploymentVersions(ctx, client, params.ProjectID); err == nil {
		mine := -1
		for _, v := range versions {
			if v.ID == params.ID {
				mine = v.Version
			}
		}
		for _, v := range versions {
			if mine >= 0 && v.Version < mine {
				older = append(older, v)
			}
		}
	} else {
		log.Printf("[WARN] deployment delete: listing versions of project %s: %s", params.ProjectID, err)
	}

	resp, err := client.Deployment.Delete(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		if isStatusNotFound(err) {
			d.SetId("")
			return diags
		}
		log.Printf("[TRACE] deployment delete error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	for _, v := range older { // newest first
		p := deployment.NewDeleteParams()
		p.ProjectID = params.ProjectID
		p.ID = v.ID
		if _, err := client.Deployment.Delete(p, nil); err != nil && !isStatusNotFound(err) {
			diags = append(diags, diag.Errorf("deleting older version %d (%s) of the deployment: %s", v.Version, v.ID, err)...)
		}
	}
	if diags.HasError() {
		return diags
	}

	d.SetId("")
	return diags
}

// GetDeploymentByID reads a deployment by ID.
func GetDeploymentByID(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := deployment.NewGetByIDParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	idVal, idIsSet := d.GetOk("id")
	if !idIsSet {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}
	id, _ := idVal.(string)
	params.ID = id

	projectIdVal, projectIdIsSet := d.GetOk("project_id")
	if !projectIdIsSet {
		diags = append(diags, diag.Errorf("missing client parameter: projectId")...)
		return diags
	}
	params.ProjectID = projectIdVal.(string)

	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.Deployment.GetByID(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		log.Printf("[TRACE] deployment get by id error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	respModel := resp.GetPayload()
	zschema.SetDeploymentResourceData(d, respModel)
	d.SetId(id)

	return diags
}

// UpdateDeployment changes a deployment in place. The controller has no
// in-place update: PUT /v2/projects/id/{projectId}/deployments/id/{id} creates
// the next version (new id, version+1) from the latest one, answers 202 with
// the new id and a job id, and applies the policies asynchronously. The
// resource moves to the new version once that job completes. If the job
// fails, the controller deletes the new version, so the resource keeps the old
// id and the error is returned.
func UpdateDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*api_client.ZedcloudAPI)
	model := zschema.DeploymentModel(d)
	oldID := d.Id()
	model.ID = oldID
	// Send the user's device policy names: the controller returns them
	// decorated ("<name>-pcr") and would decorate them again.
	for _, p := range model.DevicePolicies {
		if p != nil && p.MetaData != nil {
			p.MetaData.Name = zschema.StripDevicePolicyNameDecoration(p.MetaData.Name)
		}
	}

	if model.DeploymentTag == "" {
		return diag.Errorf("deployment_tag must be set to update a deployment: the controller rejects a new deployment version without it")
	}

	// The controller only accepts a new version from the latest one, and its
	// error for anything else carries no message, so check first.
	latestID, err := getLatestProjectDeployment(ctx, client, model.ProjectID)
	if err != nil {
		if ds, ok := ZsrvResponderToDiags(err); ok {
			return append(diags, ds...)
		}
		return append(diags, diag.Errorf("deployment update: listing deployment versions: %s", err)...)
	}
	if latestID != "" && latestID != oldID {
		return append(diags, notLatestDeploymentDiag(oldID, latestID, model.ProjectID))
	}

	params := deployment.NewCreateNewVersionParams()
	params.SetBody(model)
	params.SetProjectID(model.ProjectID)
	params.SetID(oldID)
	resp, err := client.Deployment.CreateNewVersion(params, nil)
	if err != nil {
		log.Printf("[TRACE] deployment update error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			return append(diags, explainDeploymentVersionError(ds, oldID, model.ProjectID)...)
		}
		return append(diags, diag.Errorf("deployment update: %s", err)...)
	}
	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, e := range responseData.Error {
			if e.ErrorCode != nil && (*e.ErrorCode == models.ErrorCodeAccepted || *e.ErrorCode == models.ErrorCodeSuccess) {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(e.Details))...)
		}
		if diags.HasError() {
			return explainDeploymentVersionError(diags, oldID, model.ProjectID)
		}
	}
	if responseData == nil || responseData.ObjectID == "" {
		return append(diags, diag.Errorf("deployment update: the controller did not return the id of the new deployment version")...)
	}

	// Wait for the policies of the new version to be applied. With no
	// policies in the request gilas sends no policy message, so no one would
	// ever complete the job.
	if responseData.JobID != "" && deploymentHasPolicies(model) {
		if _, err := deploymentPoller(client).wait(ctx, responseData.JobID, d.Timeout(schema.TimeoutUpdate)); err != nil {
			return append(diags, diag.Errorf("deployment update: applying the policies of the new version %s failed (the controller "+
				"removes a version whose policies fail, so %s is still the current version): %s", responseData.ObjectID, oldID, err)...)
		}
	}

	d.SetId(responseData.ObjectID)
	if errs := GetDeploymentByID(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}
	return diags
}

const (
	deploymentUpdateTimeout     = 5 * time.Minute
	deploymentNotLatestFragment = "Supported only in latest deployment version"
	// objectType of the job a deployment update creates (seine deploymentUpdate).
	deploymentJobObjectType = "OBJECT_TYPE_DEPLOYMENT"
)

// deploymentPoller waits for the job that applies a new deployment version's
// policies, with the provider's job poll settings.
func deploymentPoller(c *api_client.ZedcloudAPI) *jobPoller {
	s := c.Settings
	if s == nil {
		s = api_client.DefaultSettings()
	}
	return &jobPoller{
		client:      c.Jobs,
		objectType:  deploymentJobObjectType,
		intervalMin: s.JobPollIntervalMin,
		intervalMax: s.JobPollIntervalMax,
	}
}

func deploymentHasPolicies(dep *models.Deployment) bool {
	return len(dep.AppInstPolicies) > 0 || len(dep.NetworkInstPolicies) > 0 ||
		len(dep.VolumeInstPolicies) > 0 || len(dep.DevicePolicies) > 0 || dep.EdgeviewPolicy != nil
}

// explainDeploymentVersionError adds a hint to the controller's error when
// the deployment is not the latest version of the project's deployment (a
// race with a version created outside Terraform after the pre-check).
func explainDeploymentVersionError(diags diag.Diagnostics, id, projectID string) diag.Diagnostics {
	for i := range diags {
		if strings.Contains(diags[i].Summary+" "+diags[i].Detail, deploymentNotLatestFragment) {
			diags[i].Detail = strings.TrimSpace(diags[i].Detail + "\n\n" + notLatestDeploymentDiag(id, "", projectID).Detail)
		}
	}
	return diags
}

func notLatestDeploymentDiag(id, latestID, projectID string) diag.Diagnostic {
	latest := "<latest id>"
	if latestID != "" {
		latest = latestID
	}
	return diag.Diagnostic{
		Severity: diag.Error,
		Summary:  fmt.Sprintf("deployment %s is not the latest version in project %s", id, projectID),
		Detail: fmt.Sprintf("The controller only creates a new version from the latest one, and a newer version (%s) was created outside "+
			"Terraform. To manage the latest version, remove this resource from state and import it: "+
			"terraform state rm <address> && terraform import <address> %s:%s", latest, projectID, latest),
	}
}

// customizeDeploymentDiff validates PCR device policies at plan time: the
// controller rejects bad indices with a 400 (gilas validateDeploymentPCRPolicies).
func customizeDeploymentDiff(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
	if err := validateDeploymentTagChange(d); err != nil {
		return err
	}
	if !d.NewValueKnown("device_policies") {
		return nil
	}
	for i, p := range asList(d.Get("device_policies")) {
		pm, ok := p.(map[string]interface{})
		if !ok {
			continue
		}
		pcr := asList(pm["pcr_policy"])
		subType, _ := pm["policy_sub_type"].(string)
		if subType == string(models.DevicePolicyTypeDEVICEPOLICYTYPEPCR) && (len(pcr) == 0 || pcr[0] == nil) {
			return fmt.Errorf("device_policies.%d: policy_sub_type %s requires a pcr_policy block", i, subType)
		}
		if len(pcr) == 0 || pcr[0] == nil {
			continue
		}
		if subType != string(models.DevicePolicyTypeDEVICEPOLICYTYPEPCR) {
			return fmt.Errorf("device_policies.%d: pcr_policy requires policy_sub_type = %q", i, models.DevicePolicyTypeDEVICEPOLICYTYPEPCR)
		}
		indices := zschema.PCRIndicesFromInterface(pcr[0].(map[string]interface{})["pcr_indices"])
		if err := zschema.ValidatePCRIndices(int64sToInts(indices)); err != nil {
			return fmt.Errorf("device_policies.%d.pcr_policy.pcr_indices: %w", i, err)
		}
	}
	return nil
}

// deploymentVersion is one version of a project's deployment.
type deploymentVersion struct {
	ID      string
	Version int
}

// listProjectDeploymentVersions returns the deployment versions in a project,
// newest first.
func listProjectDeploymentVersions(ctx context.Context, client *api_client.ZedcloudAPI, projectID string) ([]deploymentVersion, error) {
	params := deployment.NewGetListbyIdParams()
	params.SetProjectID(projectID)

	resp, err := client.Deployment.GetListByProjectID(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return nil, err
	}

	var out []deploymentVersion
	for _, dep := range resp.GetPayload().List {
		if dep == nil || dep.Revision == nil || dep.Revision.Curr == nil {
			continue
		}
		v, err := strconv.Atoi(*dep.Revision.Curr)
		if err != nil {
			return nil, err
		}
		out = append(out, deploymentVersion{ID: dep.ID, Version: v})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Version > out[j].Version })
	return out, nil
}

// getLatestProjectDeployment returns the latest deployment ID for a given project.
func getLatestProjectDeployment(ctx context.Context, client *api_client.ZedcloudAPI, projectID string) (string, error) {
	versions, err := listProjectDeploymentVersions(ctx, client, projectID)
	if err != nil || len(versions) == 0 {
		return "", err
	}
	return versions[0].ID, nil
}

// deploymentUpdatableKeys are the attributes whose change makes Terraform
// update the deployment, i.e. create a new version.
var deploymentUpdatableKeys = []string{
	"app_inst_policies", "cluster_policy", "device_policies", "edgeview_policy",
	"integration_policy", "name", "network_inst_policies", "title", "volume_inst_policies",
}

// validateDeploymentTagChange fails the plan when an existing deployment is
// changed without a new deployment_tag. Every change creates a new version,
// and the controller requires each version in a project to have its own
// deployment tag (unique constraint on enterprise, deployment tag, project);
// without a new tag the update fails with a 409. Edge nodes select a version
// by its tag, so a new tag is also how nodes are moved to the new version.
func validateDeploymentTagChange(d *schema.ResourceDiff) error {
	if d.Id() == "" || d.HasChange("deployment_tag") {
		return nil
	}
	for _, k := range deploymentUpdatableKeys {
		// GetChangedKeysPrefix looks at the planned diff, after the
		// DiffSuppressFuncs; HasChange would also see suppressed differences
		// such as the controller's decorated device policy names.
		if changed := d.GetChangedKeysPrefix(k); len(changed) > 0 {
			return fmt.Errorf("changing %s creates a new version of the deployment, and the controller requires every version "+
				"in a project to have its own deployment_tag: set a new deployment_tag (for example %q). "+
				"Edge nodes keep running the version whose tag they carry until their deployment_tag is changed to the new one",
				k, suggestDeploymentTag(d))

		}
	}
	return nil
}

func suggestDeploymentTag(d *schema.ResourceDiff) string {
	tag, _ := d.Get("deployment_tag").(string)
	if tag == "" {
		return "depl:v2"
	}
	return tag + "-v2"
}
