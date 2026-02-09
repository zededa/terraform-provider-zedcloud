package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/client/z_k_s_cluster_instance"

	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

/*
ZksClusterInstanceImports z k s cluster instances API
*/

func ZksInstanceImportResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: ImportZKSInstance,
		ReadContext:   GetZKSInstance,
		UpdateContext: UpdateZKSInstance,
		DeleteContext: DeleteZKSInstance,
		Schema:        zschema.ZksInstanceSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func ImportZKSInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.ZksInstanceModel(d)
	params := z_k_s_cluster_instance.NewImportZKSInstanceParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, _, err := client.ZksClusterInstance.ImportZKSInstance(params, nil)
	if err != nil {
		log.Printf("[TRACE] ImportZKSInstance error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("ImportZKSInstance error: %s", err)...)
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
	if errs := GetZKSInstance(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	return diags
}
