package image

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/zededa/terraform-provider/models"
)

// UplinkParams creates a new ImageConfigurationUplinkImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func UplinkParams() *ImageConfigurationUplinkImageParams {
	return &ImageConfigurationUplinkImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageConfigurationUplinkImageParamsWithTimeout creates a new ImageConfigurationUplinkImageParams object
// with the ability to set a timeout on a request.
func NewImageConfigurationUplinkImageParamsWithTimeout(timeout time.Duration) *ImageConfigurationUplinkImageParams {
	return &ImageConfigurationUplinkImageParams{
		timeout: timeout,
	}
}

// NewImageConfigurationUplinkImageParamsWithContext creates a new ImageConfigurationUplinkImageParams object
// with the ability to set a context for a request.
func NewImageConfigurationUplinkImageParamsWithContext(ctx context.Context) *ImageConfigurationUplinkImageParams {
	return &ImageConfigurationUplinkImageParams{
		Context: ctx,
	}
}

// NewImageConfigurationUplinkImageParamsWithHTTPClient creates a new ImageConfigurationUplinkImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageConfigurationUplinkImageParamsWithHTTPClient(client *http.Client) *ImageConfigurationUplinkImageParams {
	return &ImageConfigurationUplinkImageParams{
		HTTPClient: client,
	}
}

/*
ImageConfigurationUplinkImageParams contains all the parameters to send to the API endpoint

	for the image configuration uplink image operation.

	Typically these are written to a http.Request.
*/
type ImageConfigurationUplinkImageParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.Image

	/* ID.

	   User defined id of the image, unique across the enterprise. Once image is created, ID can’t be changed.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image configuration uplink image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationUplinkImageParams) WithDefaults() *ImageConfigurationUplinkImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image configuration uplink image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationUplinkImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the image configuration uplink image params
func (o *ImageConfigurationUplinkImageParams) WithTimeout(timeout time.Duration) *ImageConfigurationUplinkImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image configuration uplink image params
func (o *ImageConfigurationUplinkImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image configuration uplink image params
func (o *ImageConfigurationUplinkImageParams) WithContext(ctx context.Context) *ImageConfigurationUplinkImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image configuration uplink image params
func (o *ImageConfigurationUplinkImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image configuration uplink image params
func (o *ImageConfigurationUplinkImageParams) WithHTTPClient(client *http.Client) *ImageConfigurationUplinkImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image configuration uplink image params
func (o *ImageConfigurationUplinkImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the image configuration uplink image params
func (o *ImageConfigurationUplinkImageParams) WithXRequestID(xRequestID *string) *ImageConfigurationUplinkImageParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the image configuration uplink image params
func (o *ImageConfigurationUplinkImageParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the image configuration uplink image params
func (o *ImageConfigurationUplinkImageParams) WithBody(body *models.Image) *ImageConfigurationUplinkImageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the image configuration uplink image params
func (o *ImageConfigurationUplinkImageParams) SetBody(body *models.Image) {
	o.Body = body
}

// WithID adds the id to the image configuration uplink image params
func (o *ImageConfigurationUplinkImageParams) WithID(id string) *ImageConfigurationUplinkImageParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the image configuration uplink image params
func (o *ImageConfigurationUplinkImageParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ImageConfigurationUplinkImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
