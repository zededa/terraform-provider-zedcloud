package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	z_asset_groups "github.com/zededa/terraform-provider-zedcloud/v2/client/zasset_groups"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

/*
ZAssetGroups z asset groups API
*/

func ZAssetGroups() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateZAssetGroups,
		DeleteContext: DeleteZAssetGroup,
		ReadContext:   ReadZAssetGroup,
		UpdateContext: UpdateZAssetGroup,

		Schema: zschema.ZserviceAssetGroupSchema(),
	}
}

func DataResourceZAssetGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: ReadZAssetGroup,
		Schema:      zschema.ZserviceAssetGroupSchema(),
	}
}

func ReadZAssetGroup(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := z_asset_groups.NewZAssetGroupsGetAssetGroupParams()

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	clients := m.(*ProviderClients)
	client := clients.ZServicesClient

	resp, err := client.ZassetGroup.ZAssetGroupsGetAssetGroup(params, nil)
	if err != nil {
		log.Printf("[TRACE] ZassetGroup.ZAssetGroupsGetAssetGroup error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("ZAssetGroups.ZAssetGroupsGetAssetGroup error: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	zschema.SetZserviceAssetGroupReadROResourceData(d, respModel)

	return diags
}

func ListZAssetGroups(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := z_asset_groups.NewZAssetGroupsListAssetGroupsParams()

	clients := m.(*ProviderClients)
	client := clients.ZServicesClient
	resp, err := client.ZassetGroup.ZAssetGroupsListAssetGroups(params, nil)
	if err != nil {
		log.Printf("[TRACE] ZAssetGroups.ZAssetGroupsListAssetGroups error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("ZAssetGroups.ZAssetGroupsListAssetGroups error: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	if len(respModel.AssetGroups) == 0 {
		return append(diags, diag.Errorf("no devices found")...)
	}
	// limit output to a single result
	result := respModel.AssetGroups[0]

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	d.SetId(result.ID)
	return diags
}

func CreateZAssetGroups(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.ZserviceAssetGroupConfigModel(d)
	params := z_asset_groups.NewZAssetGroupsCreateAssetGroupsParams()
	params.SetBody(model)

	clients := m.(*ProviderClients)
	client := clients.ZServicesClient

	resp, err := client.ZassetGroup.ZAssetGroupsCreateAssetGroups(params, nil)
	if err != nil {
		log.Printf("[TRACE] ZassetGroup.ZAssetGroupsCreateAssetGroups error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("ZassetGroup.ZAssetGroupsCreateAssetGroups error: %s", err)...)
		return diags
	}

	responseData := resp.GetPayload()

	// note, we need to set the ID in any case, even GetByName endpoint seems to require the ID
	// but doesn't return any error if it's not set.
	d.SetId(responseData.AssetGroupID)

	return diags
}

func UpdateZAssetGroup(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	params := z_asset_groups.NewZAssetGroupsUpdateAssetGroupParams()

	model := zschema.ZserviceAssetGroupConfigModel(d)
	params.SetBody(model)
	// ZAssetGroupsUpdateAssetGroupBody

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	// makes a bulk update for all properties that were changed
	clients := m.(*ProviderClients)
	client := clients.ZServicesClient
	resp, err := client.ZassetGroup.ZAssetGroupsUpdateAssetGroup(params, nil)
	if err != nil {
		log.Printf("[TRACE] ZAssetGroups.ZAssetGroupsUpdateAssetGroup error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("ZAssetGroups.ZAssetGroupsUpdateAssetGroup error: %s", err)...)
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

func DeleteZAssetGroup(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := z_asset_groups.NewZAssetGroupsDeleteAssetGroupParams()

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	clients := m.(*ProviderClients)
	client := clients.ZServicesClient

	_, err := client.ZassetGroup.ZAssetGroupsDeleteAssetGroup(params, nil)
	if err != nil {
		log.Printf("[TRACE] ZAssetGroups.ZAssetGroupsDeleteAssetGroup error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("ZAssetGroups.ZAssetGroupsDeleteAssetGroup error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
