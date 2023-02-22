package resources

import (
	"context"
	"errors"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider/client"
	config "github.com/zededa/terraform-provider/client/image_configuration"
	"github.com/zededa/terraform-provider/models"
	zschema "github.com/zededa/terraform-provider/schemas"
)

func ImageResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateImage,
		ReadContext:   ReadImage,
		UpdateContext: UpdateImage,
		DeleteContext: DeleteImage,
		Schema:        zschema.Image(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func ImageDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: ReadImage,
		Schema:      zschema.Image(),
	}
}

func CreateImage(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.ImageModel(d)
	params := config.CreateImageParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.Image.Create(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
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
	if errs := readImageByName(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func ReadImage(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if _, isSet := d.GetOk("name"); isSet {
		return readImageByName(ctx, d, m)
	}
	return readImageByID(ctx, d, m)
}

func readImageByID(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.GetByIDParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	id, isSet := d.GetOk("id")
	if isSet {
		id, _ := id.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.Image.GetByID(params, nil)

	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	image := resp.GetPayload()
	zschema.SetImageResourceData(d, image)
	d.SetId(image.ID)

	return diags
}

func readImageByName(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.GetByNameParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	nameVal, nameIsSet := d.GetOk("name")
	if nameIsSet {
		name, _ := nameVal.(string)
		params.Name = name
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: name")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.Image.GetByName(params, nil)

	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	image := resp.GetPayload()
	zschema.SetImageResourceData(d, image)
	d.SetId(image.ID)

	return diags
}

func UpdateImage(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.UpdateParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.ImageModel(d))

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
	resp, err := client.Image.Update(params, nil)
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
	if errs := readImageByName(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func DeleteImage(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	resp, err := client.Image.Delete(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}

func uplink(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.UplinkParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.ImageModel(d))

	nameVal, nameIsSet := d.GetOk("name")
	if nameIsSet {
		params.Name = nameVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: name")...)
		return diags
	}

	// makes a bulk update for all properties that were changed
	client := m.(*api_client.ZedcloudAPI)
	resp, accepted, err := client.Image.Uplink(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	if accepted != nil {
		// FIXME: what to do?
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		return diags
	}

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	// if errs := GetDevice(ctx, d, m); err != nil {
	// 	return append(diags, errs...)
	// }

	return diags
}

// func ImageConfiguration_MarkEveImageLatest(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	d.Partial(true)

// 	params := config.NewImageConfigurationMarkEveImageLatestParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	params.SetBody(zschema.ImageConfigurationModel(d))
// 	// models.ImageConfig

// 	// makes a bulk update for all properties that were changed
// 	client := m.(*apiclient.Zedcloudapi)
// 	resp, err := client.ImageConfiguration.ImageConfigurationMarkEveImageLatest(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	d.Partial(false)

// 	return diags
// }

// func ImageConfiguration_MarkEveImageLatest2(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	d.Partial(true)

// 	params := config.NewImageConfigurationMarkEveImageLatest2Params()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	params.SetBody(zschema.ImageConfigurationModel(d))
// 	// ImageConfigurationMarkEveImageLatest2Body

// 	imageArchVal, imageArchIsSet := d.GetOk("image_arch")
// 	if imageArchIsSet {
// 		params.ImageArch = imageArchVal.(string)
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: imageArch")...)
// 		return diags
// 	}

// 	// makes a bulk update for all properties that were changed
// 	client := m.(*apiclient.Zedcloudapi)
// 	resp, err := client.ImageConfiguration.ImageConfigurationMarkEveImageLatest2(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	d.Partial(false)

// 	return diags
// }
// func ImageConfiguration_UploadImageChunked(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	d.Partial(true)

// 	params := config.NewImageConfigurationUploadImageChunkedParams()

// 	contentRangeVal, contentRangeIsSet := d.GetOk("content_range")
// 	if contentRangeIsSet {
// 		params.ContentRange = contentRangeVal.(string)
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: Content-Range")...)
// 		return diags
// 	}

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	params.SetBody(zschema.ImageConfigurationModel(d))
// 	// strfmt.Base64

// 	nameVal, nameIsSet := d.GetOk("name")
// 	if nameIsSet {
// 		params.Name = nameVal.(string)
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: name")...)
// 		return diags
// 	}

// 	// makes a bulk update for all properties that were changed
// 	client := m.(*apiclient.Zedcloudapi)
// 	resp, err := client.ImageConfiguration.ImageConfigurationUploadImageChunked(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	d.Partial(false)

// 	return diags
// }

// func ImageConfiguration_UploadImageFile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	d.Partial(true)

// 	params := config.NewImageConfigurationUploadImageFileParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	imageFileVal, imageFileIsSet := d.GetOk("image_file")
// 	if imageFileIsSet {
// 		params.ImageFile = imageFileVal.(io.ReadCloser)
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: imageFile")...)
// 		return diags
// 	}

// 	nameVal, nameIsSet := d.GetOk("name")
// 	if nameIsSet {
// 		params.Name = nameVal.(string)
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: name")...)
// 		return diags
// 	}

// 	// makes a bulk update for all properties that were changed
// 	client := m.(*apiclient.Zedcloudapi)
// 	resp, err := client.ImageConfiguration.ImageConfigurationUploadImageFile(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	d.Partial(false)

// 	return diags
// }

// func ImageConfiguration_DeleteLatestImage(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics

// 	params := config.NewImageConfigurationDeleteLatestImageParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = &id
// 	}

// 	client := m.(*apiclient.Zedcloudapi)

// 	resp, err := client.ImageConfiguration.ImageConfigurationDeleteLatestImage(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
// 		return diags
// 	}

// 	d.SetId("")
// 	return diags
// }
