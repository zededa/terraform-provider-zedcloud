package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	kubernetes_deployments "github.com/zededa/terraform-provider-zedcloud/v2/client/kubernetes_deployment"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

/*
KubernetesDeployments kubernetes deployments API
*/

func KubernetesDeploymentResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateK3sDeployment,
		ReadContext:   GetK3sDeployment,
		UpdateContext: UpdateK3sDeployment,
		DeleteContext: DeleteK3sDeployment,
		Schema:        zschema.DeploymentRequestSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func KubernetesDeploymentDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: GetK3sDeployment,
		Schema:      zschema.DeploymentRequestSchema(),
	}
}

func CreateK3sDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.DeploymentRequestModel(d)
	params := kubernetes_deployments.NewCreateDeploymentParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, _, err := client.KubernetesDeployment.CreateDeployment(params, nil)
	if err != nil {
		log.Printf("[TRACE] KubernetesDeploymentsCreateDeployment error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("KubernetesDeploymentsCreateDeployment error: %s", err)...)
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
	if errs := GetK3sDeployment(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	return diags
}

func UpdateK3sDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	params := kubernetes_deployments.NewUpdateDeploymentParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.DeploymentRequestModel(d))
	// models.DeploymentRequest

	deploymentIdVal, deploymentIdIsSet := d.GetOk("id")
	if deploymentIdIsSet {
		params.DeploymentID = deploymentIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	// makes a bulk update for all properties that were changed
	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.KubernetesDeployment.UpdateDeployment(params, nil)
	if err != nil {
		log.Printf("[TRACE] KubernetesDeploymentsUpdateDeployment error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("KubernetesDeploymentsUpdateDeployment error: %s", err)...)
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
	if errs := GetK3sDeployment(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	d.Partial(false)

	return diags
}

func GetK3sDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := kubernetes_deployments.NewGetDeploymentParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	deploymentId := d.Id()
	if deploymentId == "" {
		diags = append(diags, diag.Errorf("missing deployment id")...)
		return diags
	}
	params.DeploymentID = deploymentId

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.KubernetesDeployment.GetDeployment(params, nil)
	if err != nil {
		log.Printf("[TRACE] KubernetesDeploymentsGetDeployment error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("KubernetesDeploymentsGetDeployment error: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	zschema.SetDeploymentRequestResourceData(d,
		zschema.ConvertDeploymentGetResponseToDeploymentRequestModel(respModel))

	return diags
}

func DeleteK3sDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := kubernetes_deployments.NewDeleteDeploymentParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	deploymentId := d.Id()
	if deploymentId == "" {
		diags = append(diags, diag.Errorf("missing deployment id")...)
		return diags
	}
	params.DeploymentID = deploymentId

	client := m.(*api_client.ZedcloudAPI)

	_, err := client.KubernetesDeployment.DeleteDeployment(params, nil)
	if err != nil {
		log.Printf("[TRACE] KubernetesDeploymentsDeleteDeployment error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("KubernetesDeploymentsDeleteDeployment error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
