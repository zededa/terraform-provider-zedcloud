package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	kubernetes_secrets "github.com/zededa/terraform-provider-zedcloud/v2/client/kubernetes_secret"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

/*
KubernetesSecrets kubernetes secrets API
*/

func KubernetesSecretResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateSecret,
		UpdateContext: schema.NoopContext,
		ReadContext:   GetSecret,
		DeleteContext: schema.NoopContext,
		Schema:        zschema.SecretRequestSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func KubernetesSecretDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: GetSecret,
		Schema:      zschema.SecretRequestSchema(),
	}
}

func CreateSecret(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.SecretRequestModel(d)
	params := kubernetes_secrets.NewCreateSecretParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, _, err := client.KubernetesSecret.CreateSecret(params, nil)
	if err != nil {
		log.Printf("[TRACE] KubernetesSecretsCreateSecret error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("ßßßßßKubernetesSecretsCreateSecret error: %s", err)...)
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
	if errs := GetSecret(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	return diags
}

func GetSecret(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := kubernetes_secrets.NewGetSecretParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	secretId := d.Id()
	if secretId == "" {
		diags = append(diags, diag.Errorf("missing secret id")...)
		return diags
	}
	params.SecretID = secretId

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.KubernetesSecret.GetSecret(params, nil)
	if err != nil {
		log.Printf("[TRACE] KubernetesSecretsGetSecret error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("KubernetesSecretsGetSecret error: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	// Only update metadata, preserve the data field since API doesn't return secret data
	if respModel != nil && respModel.Metadata != nil {
		d.Set("metadata", zschema.SetSecretMetadataSubResourceData([]*models.SecretMetadata{respModel.Metadata}))
	}

	return diags
}
