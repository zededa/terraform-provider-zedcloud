package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	private_helm_repository "github.com/zededa/terraform-provider-zedcloud/v2/client/private_helm_repository"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

/*
PrivateHelmRepositories private helm repositories API
*/

func PrivateHelmRepositoryResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreatePrivateRepo,
		ReadContext:   GetPrivateRepo,
		UpdateContext: UpdatePrivateRepo,
		DeleteContext: DeletePrivateRepo,
		Schema:        zschema.PrivateRepoRequestSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func PrivateHelmRepositoryDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: GetPrivateRepo,
		Schema:      zschema.PrivateRepoRequestSchema(),
	}
}

//func GetPrivateRepoCharts(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
//	var diags diag.Diagnostics
//
//	params := private_helm_repositories.NewPrivateHelmRepositoriesGetPrivateRepoChartsParams()
//
//	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
//	if xRequestIdIsSet {
//		params.XRequestID = xRequestIdVal.(*string)
//	}
//
//	privateRepoIdVal, privateRepoIdIsSet := d.GetOk("private_repo_id")
//	if privateRepoIdIsSet {
//		params.PrivateRepoID = privateRepoIdVal.(string)
//	} else {
//		diags = append(diags, diag.Errorf("missing client parameter: privateRepoId")...)
//		return diags
//	}
//
//	client := m.(*api_client.ZedcloudAPI)
//
//	resp, err := client.PrivateHelmRepositories.PrivateHelmRepositoriesGetPrivateRepoCharts(params, nil)
//	if err != nil {
//		log.Printf("[TRACE] PrivateHelmRepositories.PrivateHelmRepositoriesGetPrivateRepoCharts error: %s", spew.Sdump(err))
//		if ds, ok := ZsrvResponderToDiags(err); ok {
//			diags = append(diags, ds...)
//			return diags
//		}
//
//		diags = append(diags, diag.Errorf("PrivateHelmRepositories.PrivateHelmRepositoriesGetPrivateRepoCharts error: %s", err)...)
//		return diags
//	}
//
//	respModel := resp.GetPayload()
//	zschema.SetPrivateHelmRepositoriesResourceData(d, respModel)
//
//	return diags
//}

func CreatePrivateRepo(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.PrivateRepoRequestModel(d)
	params := private_helm_repository.NewCreatePrivateRepoParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, _, err := client.PrivateHelmRepository.CreatePrivateRepo(params, nil)
	if err != nil {
		log.Printf("[TRACE] PrivateHelmRepositoriesCreatePrivateRepo error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("PrivateHelmRepositoriesCreatePrivateRepo error: %s", err)...)
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

	// note, we need to set the ID in any case, even GetByName endpoint seems to require the ID
	// but doesn't return any error if it's not set.
	d.SetId(responseData.ObjectID)

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := GetPrivateRepo(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	return diags
}

func UpdatePrivateRepo(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	privateRepoModel := zschema.PrivateRepoRequestModel(d)
	body := private_helm_repository.UpdatePrivateRepoBody{
		Metadata: privateRepoModel.Metadata,
		Spec:     privateRepoModel.Spec,
	}

	params := private_helm_repository.NewUpdatePrivateRepoParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(body)
	// PrivateHelmRepositoriesUpdatePrivateRepoBody

	privateRepoID := d.Id()
	if privateRepoID == "" {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}
	params.ID = privateRepoID

	// makes a bulk update for all properties that were changed
	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.PrivateHelmRepository.UpdatePrivateRepo(params, nil)
	if err != nil {
		log.Printf("[TRACE] PrivateHelmRepositoriesUpdatePrivateRepo error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("PrivateHelmRepositoriesUpdatePrivateRepo error: %s", err)...)
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
	if errs := GetPrivateRepo(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	d.Partial(false)

	return diags
}

func GetPrivateRepo(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := private_helm_repository.NewGetPrivateRepoParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	privateRepoID := d.Id()
	if privateRepoID == "" {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}
	params.PrivateRepoID = privateRepoID

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.PrivateHelmRepository.GetPrivateRepo(params, nil)
	if err != nil {
		log.Printf("[TRACE] PrivateHelmRepositoriesGetPrivateRepo error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("PrivateHelmRepositoriesGetPrivateRepo error: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	zschema.SetPrivateRepoRequestResourceData(d,
		zschema.ConvertPrivateRepoResponseModelToPrivateRepoRequestModel(respModel))

	return diags
}

func DeletePrivateRepo(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := private_helm_repository.NewDeletePrivateRepoParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	privateRepoID := d.Id()
	if privateRepoID == "" {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}
	params.PrivateRepoID = privateRepoID

	client := m.(*api_client.ZedcloudAPI)

	_, err := client.PrivateHelmRepository.DeletePrivateRepo(params, nil)
	if err != nil {
		log.Printf("[TRACE] PrivateHelmRepositoriesDeletePrivateRepo error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("PrivateHelmRepositoriesDeletePrivateRepo error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
