package resources

import (
	"context"
	"errors"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/client"
	project "github.com/zededa/terraform-provider-zedcloud/client/projects"
	"github.com/zededa/terraform-provider-zedcloud/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/schemas"
)

func ProjectResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateProject,
		ReadContext:   GetProject,
		UpdateContext: UpdateProject,
		DeleteContext: DeleteProject,
		Schema:        zschema.Project(),
	}
}

func ProjectDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: GetProject,
		Schema:      zschema.Project(),
	}
}

func CreateProject(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.TagModel(d)
	params := project.CreateParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.Project.Create(params, nil)
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
	if errs := GetProject(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func GetProject(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var project *models.Tag
	var diags diag.Diagnostics

	if _, isSet := d.GetOk("name"); isSet {
		project, diags = readProjectByName(ctx, d, m)
	} else if _, isSet := d.GetOk("id"); isSet {
		project, diags = readProjectByID(ctx, d, m)
	}

	if diags.HasError() {
		return diags
	}

	zschema.SetTagResourceData(d, project)
	d.SetId(project.ID)

	return diags
}

func readProjectByID(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.Tag, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := project.GetByIDParams()

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

	resp, err := client.Project.GetByID(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return nil, diags
	}

	return resp.GetPayload(), diags
}

func readProjectByName(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.Tag, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := project.GetByNameParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	name, isSet := d.GetOk("name")
	if isSet {
		params.Name = name.(string)
	} else {
		return nil, append(diags, diag.Errorf("missing client parameter: name")...)
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.Project.GetByName(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return nil, diags
	}

	return resp.GetPayload(), diags
}

func UpdateProject(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := project.UpdateParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.TagModel(d))

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.Project.Update(params, nil)
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

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := GetProject(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func DeleteProject(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := project.DeleteParams()

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

	resp, err := client.Project.Delete(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
