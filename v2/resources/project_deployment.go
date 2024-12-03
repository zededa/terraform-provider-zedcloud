package resources

import (
	"context"
	"errors"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	"github.com/zededa/terraform-provider-zedcloud/v2/client/project_deployment"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

func ProjectDeploymentResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateProjectDeployment,
		DeleteContext: DeleteProjectDeploymentAll,
		ReadContext:   GetProjectDeploymentByID,
		UpdateContext: UpdateProjectDeployment,
		Schema:        zschema.Project(),
	}
}

func ProjectDeploymentDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: GetProject,
		Schema:      zschema.Project(),
	}
}

func GetProjectDeploymentByID(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := project_deployment.NewGetByIDParams()

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

	resp, err := client.ProjectDeployment.GetProjectDeploymentByID(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	respModel := resp.GetPayload()
	zschema.SetDeploymentResourceData(d, respModel)

	return diags
}

func GetProjectDeploymentListbyID(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := project_deployment.NewGetListbyIdParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	projectIdVal, projectIdIsSet := d.GetOk("project_id")
	if projectIdIsSet {
		params.ProjectID = projectIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: projectId")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.ProjectDeployment.GetProjectDeploymentListByID(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	respModel := resp.GetPayload()
	if len(respModel.List) == 0 {
		return append(diags, diag.Errorf("no devices found")...)
	}
	// limit output to a single result
	result := respModel.List[0]

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	d.SetId(result.ID)
	return diags
}

func CreateProjectDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.TagModel(d)
	params := project_deployment.NewCreateParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.ProjectDeployment.CreateProjectDeployment(params, nil)
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
	// note, we need to set the ID before fetching the newly created project.
	d.SetId(responseData.ObjectID)

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := GetProject(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	return diags
}

func UpdateProjectDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	params := project_deployment.NewUpdateParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.DeploymentModel(d))
	// ResourceGroupUpdateResourceGroupV2Body

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

	// makes a bulk update for all properties that were changed
	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.ProjectDeployment.UpdateProjectDeployment(params, nil)
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

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := GetProject(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	d.Partial(false)

	return diags
}

func DeleteProjectDeploymentAll(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := project_deployment.NewDeleteAllParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	projectIdVal, projectIdIsSet := d.GetOk("id")
	if projectIdIsSet {
		params.ProjectID = projectIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: projectId")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.ProjectDeployment.DeleteProjectDeploymentAll(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}

func DeleteProjectDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := project_deployment.NewDeleteParams()

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

	resp, err := client.ProjectDeployment.DeleteProjectDeployment(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
