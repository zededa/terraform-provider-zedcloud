package image

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewImageConfigurationDeleteLatestImageParams creates a new ImageConfigurationDeleteLatestImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImageConfigurationDeleteLatestImageParams() *ImageConfigurationDeleteLatestImageParams {
	return &ImageConfigurationDeleteLatestImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageConfigurationDeleteLatestImageParamsWithTimeout creates a new ImageConfigurationDeleteLatestImageParams object
// with the ability to set a timeout on a request.
func NewImageConfigurationDeleteLatestImageParamsWithTimeout(timeout time.Duration) *ImageConfigurationDeleteLatestImageParams {
	return &ImageConfigurationDeleteLatestImageParams{
		timeout: timeout,
	}
}

// NewImageConfigurationDeleteLatestImageParamsWithContext creates a new ImageConfigurationDeleteLatestImageParams object
// with the ability to set a context for a request.
func NewImageConfigurationDeleteLatestImageParamsWithContext(ctx context.Context) *ImageConfigurationDeleteLatestImageParams {
	return &ImageConfigurationDeleteLatestImageParams{
		Context: ctx,
	}
}

// NewImageConfigurationDeleteLatestImageParamsWithHTTPClient creates a new ImageConfigurationDeleteLatestImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageConfigurationDeleteLatestImageParamsWithHTTPClient(client *http.Client) *ImageConfigurationDeleteLatestImageParams {
	return &ImageConfigurationDeleteLatestImageParams{
		HTTPClient: client,
	}
}

/*
ImageConfigurationDeleteLatestImageParams contains all the parameters to send to the API endpoint

	for the image configuration delete latest image operation.

	Typically these are written to a http.Request.
*/
type ImageConfigurationDeleteLatestImageParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   System defined universally unique Id of the image
	*/
	ID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image configuration delete latest image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationDeleteLatestImageParams) WithDefaults() *ImageConfigurationDeleteLatestImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image configuration delete latest image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationDeleteLatestImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the image configuration delete latest image params
func (o *ImageConfigurationDeleteLatestImageParams) WithTimeout(timeout time.Duration) *ImageConfigurationDeleteLatestImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image configuration delete latest image params
func (o *ImageConfigurationDeleteLatestImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image configuration delete latest image params
func (o *ImageConfigurationDeleteLatestImageParams) WithContext(ctx context.Context) *ImageConfigurationDeleteLatestImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image configuration delete latest image params
func (o *ImageConfigurationDeleteLatestImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image configuration delete latest image params
func (o *ImageConfigurationDeleteLatestImageParams) WithHTTPClient(client *http.Client) *ImageConfigurationDeleteLatestImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image configuration delete latest image params
func (o *ImageConfigurationDeleteLatestImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the image configuration delete latest image params
func (o *ImageConfigurationDeleteLatestImageParams) WithXRequestID(xRequestID *string) *ImageConfigurationDeleteLatestImageParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the image configuration delete latest image params
func (o *ImageConfigurationDeleteLatestImageParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the image configuration delete latest image params
func (o *ImageConfigurationDeleteLatestImageParams) WithID(id *string) *ImageConfigurationDeleteLatestImageParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the image configuration delete latest image params
func (o *ImageConfigurationDeleteLatestImageParams) SetID(id *string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ImageConfigurationDeleteLatestImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
