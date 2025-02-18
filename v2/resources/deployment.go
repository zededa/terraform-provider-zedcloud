package resources

import (
	"context"
	"errors"
	"log"
	"strconv"

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
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
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
