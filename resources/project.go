package resources

import (
	"context"
	"errors"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider/client"
	project "github.com/zededa/terraform-provider/client/project"
	"github.com/zededa/terraform-provider/models"
	zschema "github.com/zededa/terraform-provider/schemas"
)

func ProjectResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateProject,
		ReadContext:   GetProject,
		/*
			DeleteContext: Project_Delete,
			CreateContext: Project_GetByID,
			UpdateContext: Project_Update,
		*/
		Schema: zschema.Tag(),
	}
}

func ProjectDataSource() *schema.Resource {
	return &schema.Resource{
		Schema: zschema.Tag(),
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
		project, diags = readProjectName(ctx, d, m)
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

	params := project.NewProjectGetByIDParams()

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

	resp, err := client.Project.ProjectGetByID(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return nil, diags
	}

	edgeNode := resp.GetPayload()

	return edgeNode, diags
}

// func Project_GetByName(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics

// 	model := zschema.ProjectModel(d)
// 	params := project.NewProjectGetByNameParams()
// 	params.SetBody(model)

// 	client := m.(*apiclient.Zedcloudapi)

// 	resp, err := client.Project.ProjectGetByName(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
// 		return diags
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	return diags
// }

// func Project_Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	d.Partial(true)

// 	params := project.NewProjectUpdateParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	params.SetBody(zschema.ProjectModel(d))
// 	// ProjectUpdateBody

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	// makes a bulk update for all properties that were changed
// 	client := m.(*apiclient.Zedcloudapi)
// 	resp, err := client.Project.ProjectUpdate(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	d.Partial(false)

// 	return diags
// }

// func Project_Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics

// 	params := project.NewProjectDeleteParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	client := m.(*apiclient.Zedcloudapi)

// 	resp, err := client.Project.ProjectDelete(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
// 		return diags
// 	}

// 	d.SetId("")
// 	return diags
// }
