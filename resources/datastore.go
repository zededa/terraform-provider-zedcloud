package resources

import (
	"context"
	"errors"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider/client"
	config "github.com/zededa/terraform-provider/client/datastore"
	"github.com/zededa/terraform-provider/models"
	zschema "github.com/zededa/terraform-provider/schemas"
)

func DatastoreResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateDatastore,
		ReadContext:   ReadDatastore,
		UpdateContext: UpdateDatastore,
		DeleteContext: DeleteDatastore,
		Schema:        zschema.Datastore(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func DatastoreDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: ReadDatastore,
		Schema:      zschema.Datastore(),
	}
}

func CreateDatastore(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.DatastoreModel(d)
	params := config.CreateParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	success, created, err := client.Datastore.Create(params, nil)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	var responseData *models.ZsrvResponse
	if success != nil {
		log.Printf("[TRACE] response: %v", success)
		responseData = success.GetPayload()
	} else {
		log.Printf("[TRACE] response: %v", created)
		responseData = created.GetPayload()
	}

	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.ErrorCode != nil &&
				(*err.ErrorCode == models.ErrorCodeSuccess ||
					*err.ErrorCode == models.ErrorCodeAccepted) {
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
	if errs := readDatastoreByName(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func ReadDatastore(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if _, isSet := d.GetOk("name"); isSet {
		return readDatastoreByName(ctx, d, m)
	}
	return readDatastoreByID(ctx, d, m)
}

func readDatastoreByID(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	resp, err := client.Datastore.GetByID(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	datastore := resp.GetPayload()
	zschema.SetDatastoreResourceData(d, datastore)
	d.SetId(datastore.ID)

	return diags
}

func readDatastoreByName(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	resp, err := client.Datastore.GetByName(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	datastore := resp.GetPayload()
	zschema.SetDatastoreResourceData(d, datastore)
	d.SetId(datastore.ID)

	return diags
}

func UpdateDatastore(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.UpdateParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.DatastoreModel(d))

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.Datastore.Update(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
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
	if errs := readDatastoreByName(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func DeleteDatastore(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	resp, err := client.Datastore.Delete(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
