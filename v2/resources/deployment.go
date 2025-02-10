package resources

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	"github.com/zededa/terraform-provider-zedcloud/v2/client/deployment"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

func DeploymentResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateDeployment,
		DeleteContext: DeleteDeployment,
		ReadContext:   GetDeploymentByID,
		UpdateContext: UpdateDeployment,
		Schema:        zschema.Deployment(),
	}
}

func ProjectDeploymentDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: GetDeploymentByID,
		Schema:      zschema.Deployment(),
	}
}

func CreateDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*api_client.ZedcloudAPI)
	model := zschema.DeploymentModel(d)

	latestDeploymentID, err := getLatestProjectDeployment(ctx, client, model.ProjectID)
	if err != nil {
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
		log.Printf("[TRACE] response: %v", resp)
		if err != nil {
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
		log.Printf("[TRACE] response: %v", resp)
		if err != nil {
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
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}

func GetDeploymentByID(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := deployment.NewGetByIDParams()

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

	resp, err := client.Deployment.GetByID(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	respModel := resp.GetPayload()
	zschema.SetDeploymentResourceData(d, respModel)

	return diags
}

func UpdateDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "update not supported for deployment",
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
