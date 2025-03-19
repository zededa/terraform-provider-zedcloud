package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	asset_group_service "github.com/zededa/terraform-provider-zedcloud/v2/client/asset_group_service"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

/*
AssetGroupService asset group service API
*/

func AssetGroupService() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateAssetGroup,
		DeleteContext: DeleteAssetGroup,
		ReadContext:   ReadAssetGroup,
		UpdateContext: UpdateAssetGroup,
		Schema:        zschema.AssetGroupSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func AssetGroupDataResource() *schema.Resource {
	return &schema.Resource{
		Schema:      zschema.AssetGroupSchema(),
		ReadContext: ReadAssetGroup,
	}
}

func CreateAssetGroup(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.AssetGroupModel(d)
	params := asset_group_service.NewAssetGroupServiceCreateAssetGroupParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.AssetGroup.AssetGroupServiceCreateAssetGroup(params, nil)
	if err != nil {
		log.Printf("[TRACE] AssetGroupService.AssetGroupServiceCreateAssetGroup error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("AssetGroupService.AssetGroupServiceCreateAssetGroup error: %s", err)...)
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
	if diags := ReadAssetGroup(ctx, d, m); diags.HasError() {
		return diags
	}
	return diags
}

func ReadAssetGroup(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if _, isSet := d.GetOk("name"); isSet {
		return GetAssetGroupByName(ctx, d, m)
	}
	return GetAssetGroupByID(ctx, d, m)
}

func GetAssetGroupByID(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := asset_group_service.NewAssetGroupServiceGetAssetGroupParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.AssetGroup.AssetGroupServiceGetAssetGroup(params, nil)
	if err != nil {
		log.Printf("[TRACE] AssetGroupService.AssetGroupServiceGetAssetGroup error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("AssetGroupService.AssetGroupServiceGetAssetGroup error: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	zschema.SetAssetGroupReadResourceData(d, respModel)

	return diags
}

func GetAssetGroupByName(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := asset_group_service.NewAssetGroupServiceGetAssetGroupByNameParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	nameVal, nameIsSet := d.GetOk("name")
	if nameIsSet {
		params.Name = nameVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: name")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.AssetGroup.AssetGroupServiceGetAssetGroupByName(params, nil)
	if err != nil {
		log.Printf("[TRACE] AssetGroupService.AssetGroupServiceGetAssetGroupByName error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("AssetGroupService.AssetGroupServiceGetAssetGroupByName error: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	zschema.SetAssetGroupReadResourceData(d, respModel)

	return diags
}

func UpdateAssetGroup(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	params := asset_group_service.NewAssetGroupServiceUpdateAssetGroupParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.AssetGroupModel(d))
	// AssetGroupServiceUpdateAssetGroupBody

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	// makes a bulk update for all properties that were changed
	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.AssetGroup.AssetGroupServiceUpdateAssetGroup(params, nil)
	if err != nil {
		log.Printf("[TRACE] AssetGroupService.AssetGroupServiceUpdateAssetGroup error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("AssetGroupService.AssetGroupServiceUpdateAssetGroup error: %s", err)...)
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

	// note, we need to set the ID in any case, even GetByName endpoint seems to require items
	// but doesn't return any error if it's not set.
	d.SetId(responseData.ObjectID)

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := ReadAssetGroup(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	return diags
}

func DeleteAssetGroup(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := asset_group_service.NewAssetGroupServiceDeleteAssetGroupParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	_, err := client.AssetGroup.AssetGroupServiceDeleteAssetGroup(params, nil)
	if err != nil {
		log.Printf("[TRACE] AssetGroupService.AssetGroupServiceDeleteAssetGroup error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("AssetGroupService.AssetGroupServiceDeleteAssetGroup error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
