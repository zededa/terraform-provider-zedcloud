package resources

import (
	"context"
	"errors"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/client"
	config "github.com/zededa/terraform-provider-zedcloud/client/application_instance"
	"github.com/zededa/terraform-provider-zedcloud/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/schemas"
)

// Patch envelope reference update only updates the reference between patch envelope
// and app instance. Only POST api is supported for that.
// Hence both CreateContext and UpdateContext call the same function CreatePatchEnvelopeReference
// that does POST API call with updated params.
// Since there is no GET api support for getting the reference, ReadContext is set to noop.
// DeletePatchEnvelopeReference will remove any existing reference created for the
// app instance when the resource block is removed.
func PatchEnvelopeReferenceUpdate() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreatePatchEnvelopeReference,
		UpdateContext: CreatePatchEnvelopeReference,
		ReadContext:   schema.NoopContext,
		DeleteContext: DeletePatchEnvelopeReference,
		Schema:        zschema.PatchReferenceUpdateConfigSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func CreatePatchEnvelopeReference(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.PatchReferenceUpdateConfigModel(d)
	params := config.NewEdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.ApplicationInstance.EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstance(params, nil)
	log.Printf("[TRACE] CreatePatchEnvelopeReference response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	responseData := resp.GetPayload()
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

	d.SetId(responseData.ObjectID)

	return diags
}

func DeletePatchEnvelopeReference(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.PatchReferenceUpdateConfigModel(d)
	model.PatchenvelopeID = ""
	params := config.NewEdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.ApplicationInstance.EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstance(params, nil)
	log.Printf("[TRACE] DeletePatchEnvelopeReference response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	responseData := resp.GetPayload()
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

	d.SetId("")

	return diags
}
