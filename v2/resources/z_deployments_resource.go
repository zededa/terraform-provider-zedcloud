package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	z_deployments "github.com/zededa/terraform-provider-zedcloud/v2/client/zdeployments"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

/*
ZDeployments z deployments API
*/

func ZDeployments() *schema.Resource {
	return &schema.Resource{
		CreateContext: ZDeployments_CreateDployment,
		DeleteContext: ZDeployments_DeleteDeployment,
		ReadContext:   ZDeployments_GetDeployment,
		UpdateContext: ZDeployments_UpdateDeployment,

		Schema: zschema.ZDeploymentsSchema(),
	}
}

func DataResourceZDeployments() *schema.Resource {
	return &schema.Resource{
		ReadContext: ZDeployments_ListDeployments,
		Schema:      zschema.ZDeploymentsSchema(),
	}
}

func ZDeployments_GetDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := z_deployments.NewZDeploymentsGetDeploymentParams()

	deploymentIdVal, deploymentIdIsSet := d.GetOk("deployment_id")
	if deploymentIdIsSet {
		params.DeploymentID = deploymentIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: deploymentId")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.ZDeployments.ZDeploymentsGetDeployment(params, nil)
	if err != nil {
		log.Printf("[TRACE] ZDeployments.ZDeploymentsGetDeployment error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("ZDeployments.ZDeploymentsGetDeployment error: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	zschema.SetZDeploymentsResourceData(d, respModel)

	return diags
}

func ZDeployments_ListDeployments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := z_deployments.NewZDeploymentsListDeploymentsParams()

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.ZDeployments.ZDeploymentsListDeployments(params, nil)
	if err != nil {
		log.Printf("[TRACE] ZDeployments.ZDeploymentsListDeployments error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("ZDeployments.ZDeploymentsListDeployments error: %s", err)...)
		return diags
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
	return GetDevice(ctx, d, m)
}

func ZDeployments_CreateDployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.ZDeploymentsModel(d)
	params := z_deployments.NewZDeploymentsCreateDploymentParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.ZDeployments.ZDeploymentsCreateDployment(params, nil)
	if err != nil {
		log.Printf("[TRACE] ZDeployments.ZDeploymentsCreateDployment error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("ZDeployments.ZDeploymentsCreateDployment error: %s", err)...)
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
	if errs := GetDevice(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	return diags
}

func ZDeployments_UpdateDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	params := z_deployments.NewZDeploymentsUpdateDeploymentParams()

	params.SetBody(zschema.ZDeploymentsModel(d))
	// ZDeploymentsUpdateDeploymentBody

	deploymentIdVal, deploymentIdIsSet := d.GetOk("deployment_id")
	if deploymentIdIsSet {
		params.DeploymentID = deploymentIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: deploymentId")...)
		return diags
	}

	// makes a bulk update for all properties that were changed
	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.ZDeployments.ZDeploymentsUpdateDeployment(params, nil)
	if err != nil {
		log.Printf("[TRACE] ZDeployments.ZDeploymentsUpdateDeployment error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("ZDeployments.ZDeploymentsUpdateDeployment error: %s", err)...)
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
	if errs := GetDevice(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	d.Partial(false)

	return diags
}

func ZDeployments_DeleteDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := z_deployments.NewZDeploymentsDeleteDeploymentParams()

	deploymentIdVal, deploymentIdIsSet := d.GetOk("deployment_id")
	if deploymentIdIsSet {
		params.DeploymentID = deploymentIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: deploymentId")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	_, err := client.ZDeployments.ZDeploymentsDeleteDeployment(params, nil)
	if err != nil {
		log.Printf("[TRACE] ZDeployments.ZDeploymentsDeleteDeployment error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("ZDeployments.ZDeploymentsDeleteDeployment error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
