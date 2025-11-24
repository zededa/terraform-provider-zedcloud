package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	kubernetes_git_ops "github.com/zededa/terraform-provider-zedcloud/v2/client/kubernetes_git_ops"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

/*
KubernetesGitOps kubernetes git ops API
*/

func KubernetesGitOpsResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateGitRepo,
		ReadContext:   GetGitRepo,
		UpdateContext: UpdateGitRepo,
		DeleteContext: DeleteGitRepo,
		Schema:        zschema.GitRepoRequestSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func KubernetesGitOpsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: GetGitRepo,
		Schema:      zschema.GitRepoRequestSchema(),
	}
}

func CreateGitRepo(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.GitRepoRequestModel(d)
	params := kubernetes_git_ops.NewCreateGitRepoParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, _, err := client.KubernetesGitOps.CreateGitRepo(params, nil)
	if err != nil {
		log.Printf("[TRACE] KubernetesGitOpsCreateGitRepo error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("KubernetesGitOpsCreateGitRepo error: %s", err)...)
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
	if errs := GetGitRepo(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	return diags
}

func UpdateGitRepo(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	gitRepoModel := zschema.GitRepoRequestModel(d)
	body := kubernetes_git_ops.UpdateGitRepoBody{
		Metadata: gitRepoModel.Metadata,
		Data:     gitRepoModel.Data,
	}

	params := kubernetes_git_ops.NewUpdateGitRepoParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(body)
	// KubernetesGitOpsUpdateGitRepoBody

	gitrepoIdVal, gitrepoIdIsSet := d.GetOk("id")
	if gitrepoIdIsSet {
		params.GitrepoID = gitrepoIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	// makes a bulk update for all properties that were changed
	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.KubernetesGitOps.UpdateGitRepo(params, nil)
	if err != nil {
		log.Printf("[TRACE] KubernetesGitOpsUpdateGitRepo error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("KubernetesGitOpsUpdateGitRepo error: %s", err)...)
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
	if errs := GetGitRepo(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	d.Partial(false)

	return diags
}
func GetGitRepo(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := kubernetes_git_ops.NewGetGitRepoParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	gitrepoIdVal, gitrepoIdIsSet := d.GetOk("id")
	if gitrepoIdIsSet {
		params.GitrepoID = gitrepoIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.KubernetesGitOps.GetGitRepo(params, nil)
	if err != nil {
		log.Printf("[TRACE] KubernetesGitOpsGetGitRepo error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("KubernetesGitOpsGetGitRepo error: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	zschema.SetGitRepoRequestResourceData(d, zschema.ConvertGitRepoResponseModelToRequestModel(respModel))

	return diags
}

func DeleteGitRepo(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := kubernetes_git_ops.NewDeleteGitRepoParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	gitrepoIdVal, gitrepoIdIsSet := d.GetOk("id")
	if gitrepoIdIsSet {
		params.GitrepoID = gitrepoIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	_, err := client.KubernetesGitOps.DeleteGitRepo(params, nil)
	if err != nil {
		log.Printf("[TRACE] KubernetesGitOpsDeleteGitRepo error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("KubernetesGitOpsDeleteGitRepo error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
