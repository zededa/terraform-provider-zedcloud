package resources

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

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
The api design does not support updates on a deployment. If you want to add changes to a deployment, you need to create a new version.
When you create a new version of a deployment, it makes sense to show that it depends on the previous version.
It will help create a proper dependency graph for the Terraform plan.`,
		CreateContext: CreateDeployment,
		DeleteContext: DeleteDeployment,
		ReadContext:   GetDeploymentByID,
		UpdateContext: UpdateDeployment,
		Schema:        zschema.Deployment(),
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

// UpdateDeployment updates a deployment in ZedCloud.
func UpdateDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Update not supported for deployment, you need to create a new version instead",
	})
	return diags
}

// getLatestProjectDeployment returns the latest deployment ID for a given project.
func getLatestProjectDeployment(ctx context.Context, client *api_client.ZedcloudAPI, projectID string) (string, error) {
	params := deployment.NewGetListbyIdParams()
	params.SetProjectID(projectID)

	resp, err := client.Deployment.GetListByProjectID(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return "", err
	}

	responseData := resp.GetPayload()
	var version int
	var latestDeploymentID string
	for _, dep := range responseData.List {
		deplVersion, err := strconv.Atoi(*dep.Revision.Curr)
		if err != nil {
			return "", err
		}

		if deplVersion > version {
			version = deplVersion
			latestDeploymentID = dep.ID
		}
	}
	return latestDeploymentID, nil
}
