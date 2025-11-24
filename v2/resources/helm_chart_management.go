package resources

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/client/helm_chart_management"

	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

/*
HelmChartManagement helm chart management API
*/

// helmChartManagementSchema returns the schema for helm chart management resource
// note: the schema is different from the API schema.
func helmChartManagementSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"file": {
			Description: `File name of the chart (required)`,
			Type:        schema.TypeString,
			Required:    true,
		},
		"name": {
			Description: `Name of the chart (required)`,
			Type:        schema.TypeString,
			Required:    true,
		},
		"version": {
			Description: `Version of the chart (required for update)`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func HelmChartManagementResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateHelmChart,
		ReadContext:   schema.NoopContext,
		UpdateContext: UpdateHelmChart,
		DeleteContext: DeleteHelmChart,
		Schema:        helmChartManagementSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func HelmChartManagementDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: schema.NoopContext,
		Schema:      helmChartManagementSchema(),
	}
}

func CreateHelmChart(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	chartFile := d.Get("file").(string)
	if chartFile == "" {
		return diag.Errorf("file is required")
	}
	file, err := os.Open(chartFile)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to open chart file: %w", err))
	}
	defer file.Close()

	chartFileName := d.Get("name").(string)
	if chartFileName == "" {
		return diag.Errorf("name is required")
	}

	params := helm_chart_management.NewCreateHelmChartParams()
	params.Chart = runtime.NamedReader(chartFileName, file)

	client := m.(*api_client.ZedcloudAPI)

	resp, _, err := client.HelmChartManagement.CreateHelmChart(params, nil)
	if err != nil {
		log.Printf("[TRACE] HelmChartManagementCreateHelmChart error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("HelmChartManagementCreateHelmChart error: %s", err)...)
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

	return diags
}

func UpdateHelmChart(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	params := helm_chart_management.NewUpdateHelmChartParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	chartName, chartNameIsSet := d.GetOk("name")
	if chartNameIsSet {
		params.ChartName = chartName.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: name")...)
		return diags
	}

	chartFile, chartIsSet := d.GetOk("file")
	if chartIsSet {
		chartFilePath := chartFile.(string)

		file, err := os.Open(chartFilePath)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to open chart file: %w", err))
		}
		defer file.Close()

		params.Chart = runtime.NamedReader(params.ChartName, file)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: file")...)
		return diags
	}

	chartVersion, chartVersionIsSet := d.GetOk("version")
	if chartVersionIsSet {
		params.ChartVersion = chartVersion.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: version")...)
		return diags
	}

	// makes a bulk update for all properties that were changed
	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.HelmChartManagement.UpdateHelmChart(params, nil)
	if err != nil {
		log.Printf("[TRACE] HelmChartManagementUpdateHelmChart error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("HelmChartManagementUpdateHelmChart error: %s", err)...)
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

	d.Partial(false)

	return diags
}

func DeleteHelmChart(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := helm_chart_management.NewDeleteHelmChartParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	chartNameVal, chartNameIsSet := d.GetOk("name")
	if chartNameIsSet {
		params.ChartName = chartNameVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: name")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	_, err := client.HelmChartManagement.DeleteHelmChart(params, nil)
	if err != nil {
		log.Printf("[TRACE] HelmChartManagementDeleteHelmChart error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("HelmChartManagementDeleteHelmChart error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
