package resources

import (
	"context"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	z_deployments "github.com/zededa/terraform-provider-zedcloud/v2/client/z_deployments"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

/*
ZDeployments z deployments API
*/

func ZDeployments() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateZDeployment,
		DeleteContext: DeleteZDeployment,
		ReadContext:   ReadZDeployment,
		UpdateContext: UpdateZDeployment,
		Schema:        zschema.ZserviceDeploymentConfigSchema(),
	}
}

func DataResourceZDeployments() *schema.Resource {
	return &schema.Resource{
		ReadContext: ZDeployments_ListDeployments,
		Schema:      zschema.ZserviceDeploymentConfigSchema(),
	}
}

func ReadZDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := z_deployments.NewZDeploymentsGetDeploymentParams()

	deploymentIdVal, deploymentIdIsSet := d.GetOk("deployment_id")
	if deploymentIdIsSet {
		params.DeploymentID = deploymentIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: deploymentId")...)
		return diags
	}

	clients := m.(*ProviderClients)
	client := clients.ZServicesClient

	resp, err := client.Zdeployments.ZDeploymentsGetDeployment(params, nil)
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
	zschema.SetZserviceDeploymentReadROResourceData(d, respModel)

	return diags
}

func ZDeployments_ListDeployments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := z_deployments.NewZDeploymentsListDeploymentsParams()

	clients := m.(*ProviderClients)
	client := clients.ZServicesClient

	resp, err := client.Zdeployments.ZDeploymentsListDeployments(params, nil)
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
	if len(respModel.Deployments) == 0 {
		return append(diags, diag.Errorf("no devices found")...)
	}
	// limit output to a single result
	result := respModel.Deployments[0]

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	d.SetId(result.DeploymentID)
	return diags
}

func CreateZDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.ZserviceDeploymentConfigModel(d)
	params := z_deployments.NewZDeploymentsCreateDploymentParams()
	params.SetBody(model)

	clients := m.(*ProviderClients)
	client := clients.ZServicesClient

	resp, err := client.Zdeployments.ZDeploymentsCreateDployment(params, nil)
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

	// note, we need to set the ID in any case, even GetByName endpoint seems to require the ID
	// but doesn't return any error if it's not set.
	d.SetId(responseData.DeploymentID)
	return diags
}

func UpdateZDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	params := z_deployments.NewZDeploymentsUpdateDeploymentParams()

	params.SetBody(zschema.ZserviceDeploymentConfigModel(d))
	// ZDeploymentsUpdateDeploymentBody

	deploymentIdVal, deploymentIdIsSet := d.GetOk("deployment_id")
	if deploymentIdIsSet {
		params.DeploymentID = deploymentIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: deploymentId")...)
		return diags
	}

	// makes a bulk update for all properties that were changed
	clients := m.(*ProviderClients)
	client := clients.ZServicesClient
	_, err := client.Zdeployments.ZDeploymentsUpdateDeployment(params, nil)
	if err != nil {
		log.Printf("[TRACE] ZDeployments.ZDeploymentsUpdateDeployment error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("ZDeployments.ZDeploymentsUpdateDeployment error: %s", err)...)
		return diags
	}

	// responseData := resp.GetPayload()

	d.Partial(false)

	return diags
}

func DeleteZDeployment(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := z_deployments.NewZDeploymentsDeleteDeploymentParams()

	deploymentIdVal, deploymentIdIsSet := d.GetOk("deployment_id")
	if deploymentIdIsSet {
		params.DeploymentID = deploymentIdVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: deploymentId")...)
		return diags
	}

	clients := m.(*ProviderClients)
	client := clients.ZServicesClient

	_, err := client.Zdeployments.ZDeploymentsDeleteDeployment(params, nil)
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
