package image

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/models"
)

// CreateImageParams creates a new ImageConfigurationCreateImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func CreateImageParams() *ImageConfigurationCreateImageParams {
	return &ImageConfigurationCreateImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageConfigurationCreateImageParamsWithTimeout creates a new ImageConfigurationCreateImageParams object
// with the ability to set a timeout on a request.
func NewImageConfigurationCreateImageParamsWithTimeout(timeout time.Duration) *ImageConfigurationCreateImageParams {
	return &ImageConfigurationCreateImageParams{
		timeout: timeout,
	}
}

// NewImageConfigurationCreateImageParamsWithContext creates a new ImageConfigurationCreateImageParams object
// with the ability to set a context for a request.
func NewImageConfigurationCreateImageParamsWithContext(ctx context.Context) *ImageConfigurationCreateImageParams {
	return &ImageConfigurationCreateImageParams{
		Context: ctx,
	}
}

// NewImageConfigurationCreateImageParamsWithHTTPClient creates a new ImageConfigurationCreateImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageConfigurationCreateImageParamsWithHTTPClient(client *http.Client) *ImageConfigurationCreateImageParams {
	return &ImageConfigurationCreateImageParams{
		HTTPClient: client,
	}
}

/*
ImageConfigurationCreateImageParams contains all the parameters to send to the API endpoint

	for the image configuration create image operation.

	Typically these are written to a http.Request.
*/
type ImageConfigurationCreateImageParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.Image

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image configuration create image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationCreateImageParams) WithDefaults() *ImageConfigurationCreateImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image configuration create image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationCreateImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the image configuration create image params
func (o *ImageConfigurationCreateImageParams) WithTimeout(timeout time.Duration) *ImageConfigurationCreateImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image configuration create image params
func (o *ImageConfigurationCreateImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image configuration create image params
func (o *ImageConfigurationCreateImageParams) WithContext(ctx context.Context) *ImageConfigurationCreateImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image configuration create image params
func (o *ImageConfigurationCreateImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image configuration create image params
func (o *ImageConfigurationCreateImageParams) WithHTTPClient(client *http.Client) *ImageConfigurationCreateImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image configuration create image params
func (o *ImageConfigurationCreateImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the image configuration create image params
func (o *ImageConfigurationCreateImageParams) WithXRequestID(xRequestID *string) *ImageConfigurationCreateImageParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the image configuration create image params
func (o *ImageConfigurationCreateImageParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the image configuration create image params
func (o *ImageConfigurationCreateImageParams) WithBody(body *models.Image) *ImageConfigurationCreateImageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the image configuration create image params
func (o *ImageConfigurationCreateImageParams) SetBody(body *models.Image) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ImageConfigurationCreateImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
