package resources

import (
	"context"
	"errors"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/client"
	config "github.com/zededa/terraform-provider-zedcloud/client/image"
	"github.com/zededa/terraform-provider-zedcloud/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/schemas"
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

	// due to api design, we need to fetch the newly created image
	image, diags := readImageByID(ctx, d, m)
	if diags.HasError() {
		return diags
	}

	// publish the api response to local state and the d instance
	zschema.SetImageResourceData(d, image)
	d.SetId(image.ID)

	// uplink the image to set status to IMAGE_STATUS_READY
	if diags := uplinkImage(ctx, d, m, image); diags.HasError() {
		return diags
	}

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if diags := ReadImage(ctx, d, m); diags.HasError() {
		return diags
	}

	return diags
}

func ReadImage(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var image *models.Image

	if _, isSet := d.GetOk("name"); isSet {
		image, diags = readImageByName(ctx, d, m)
	} else {
		image, diags = readImageByID(ctx, d, m)
	}

	if diags.HasError() {
		return diags
	}

	zschema.SetImageResourceData(d, image)
	d.SetId(image.ID)

	return diags
}

func readImageByID(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.Image, diag.Diagnostics) {
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
		return nil, diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.Image.GetByID(params, nil)

	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return nil, diags
	}

	return resp.GetPayload(), diags
}

func readImageByName(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.Image, diag.Diagnostics) {
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
		return nil, diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.Image.GetByName(params, nil)

	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return nil, diags
	}

	return resp.GetPayload(), diags
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
	if errs := ReadImage(ctx, d, m); err != nil {
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

func uplinkImage(ctx context.Context, d *schema.ResourceData, m interface{}, image *models.Image) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.UplinkParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(image)
	params.ID = image.ID

	client := m.(*api_client.ZedcloudAPI)

	resp, accepted, err := client.Image.Uplink(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	var responseData *models.ZsrvResponse
	if resp != nil {
		responseData = resp.GetPayload()
	} else if accepted != nil {
		log.Println("[WARN] uplink image accepted: The API gateway accepted the request for uplinking but the uplinking process has not been completed. Please check ImageStatus and ImageError fields to track the status of uplinking process and any error messages.")
		responseData = accepted.GetPayload()
	}

	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.ErrorCode != nil && (*err.ErrorCode == models.ErrorCodeSuccess || *err.ErrorCode == models.ErrorCodeAccepted) {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if diags.HasError() {
			return diags
		}
	}

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
