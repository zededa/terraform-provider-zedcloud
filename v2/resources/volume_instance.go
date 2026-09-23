package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/volume_instance"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

func VolumeInstanceResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateVolumeInstance,
		ReadContext:   ReadVolumeInstance,
		UpdateContext: UpdateVolumeInstance,
		DeleteContext: DeleteVolumeInstance,
		Schema:        zschema.VolumeInstance(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func VolumeInstanceDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: ReadVolumeInstance,
		Schema:      zschema.VolumeInstance(),
	}
}

func CreateVolumeInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.VolumeInstanceModel(d)
	params := config.CreateParams()
	params.SetBody(model)

	_, isDevIDSet := d.GetOk("device_id")
	_, isEdgeNodeClusterSet := d.GetOk("edge_node_cluster")
	if !isDevIDSet && !isEdgeNodeClusterSet {
		diags = append(diags, diag.Errorf("either device_id or edge_node_cluster has to be specified")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.VolumeInstance.Create(params, nil)
	if err != nil {
		log.Printf("[TRACE] volume instance create error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("volume instance create error: %s", err)...)
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
	if errs := readVolumeInstanceByName(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func ReadVolumeInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if _, isSet := d.GetOk("name"); isSet {
		return readVolumeInstanceByName(ctx, d, m)
	}
	return readVolumeInstanceByID(ctx, d, m)
}

func readVolumeInstanceByID(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.GetByIDParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	id, isSet := d.GetOk("id")
	if isSet {
		params.ID = id.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.VolumeInstance.GetByID(params, nil)
	if err != nil {
		log.Printf("[TRACE] volume instance error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("volume instance read error: %s", err)...)
		return diags
	}

	volume := resp.GetPayload()
	zschema.SetVolumeInstanceResourceData(d, volume)
	d.SetId(volume.ID)

	return diags
}

func readVolumeInstanceByName(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.GetByNameParams()

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

	resp, err := client.VolumeInstance.GetByName(params, nil)
	if err != nil {
		log.Printf("[TRACE] volume instance error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("volume instance read error: %s", err)...)
		return diags
	}

	volume := resp.GetPayload()
	zschema.SetVolumeInstanceResourceData(d, volume)
	d.SetId(volume.ID)

	return diags
}

func UpdateVolumeInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.UpdateParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	model := zschema.VolumeInstanceModel(d)

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	// CI-752 / CI-790: use the revision the server holds right now as the
	// optimistic-lock token, not the one cached in Terraform state.
	//
	// The controller sets `exists.Version` from the request's `revision.curr`
	// and then updates `WHERE Version = <that>`; no row matches a stale value
	// and the write comes back HTTP 409 VersionConflict. A cluster-scoped
	// volume's version moves without Terraform doing anything -- allocating an
	// app instance propagates the designated node onto its explicit volumes --
	// so the state's revision goes stale on its own. Same fix as UpdateNode
	// (CI-790, #218), which this resource did not get.
	if rev, ds := currentVolumeInstanceRevision(d, m); ds.HasError() {
		return append(diags, ds...)
	} else if rev != nil {
		model.Revision = rev
	}

	params.SetBody(model)

	resp, err := client.VolumeInstance.Update(params, nil)
	if err != nil {
		log.Printf("[TRACE] volume instance update error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("volume instance update error: %s", err)...)
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
	if errs := readVolumeInstanceByName(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

// currentVolumeInstanceRevision fetches the volume instance's revision as the
// server holds it now.
//
// A nil revision with no error means "could not determine" -- the update then
// proceeds with whatever state carried, which is exactly the old behaviour. A
// read failure here must not be more fatal than the update it is protecting.
func currentVolumeInstanceRevision(d *schema.ResourceData, m interface{}) (*models.ObjectRevision, diag.Diagnostics) {
	var diags diag.Diagnostics

	id, isSet := d.GetOk("id")
	if !isSet {
		return nil, diags
	}

	params := config.GetByIDParams()
	params.ID = id.(string)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.VolumeInstance.GetByID(params, nil)
	if err != nil {
		log.Printf("[TRACE] volume instance revision read error: %s", spew.Sdump(err))
		return nil, diags
	}

	if payload := resp.GetPayload(); payload != nil {
		return payload.Revision, diags
	}
	return nil, diags
}

func DeleteVolumeInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.DeleteParams()

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

	_, err := client.VolumeInstance.Delete(params, nil)
	if err != nil {
		if isStatusNotFound(err) {
			d.SetId("")
			return diags
		}
		log.Printf("[TRACE] volume instance delete error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("volume instance delete error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}

func GetVolumeInstanceByID(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.GetByIDParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		params.ID = idVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.VolumeInstance.GetByID(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	volume := resp.GetPayload()
	zschema.SetVolumeInstanceResourceData(d, volume)
	d.SetId(volume.ID)

	return diags
}
