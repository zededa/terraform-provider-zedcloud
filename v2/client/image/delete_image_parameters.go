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

// DeleteParams creates a new ImageConfigurationDeleteImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func DeleteParams() *ImageConfigurationDeleteImageParams {
	return &ImageConfigurationDeleteImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageConfigurationDeleteImageParamsWithTimeout creates a new ImageConfigurationDeleteImageParams object
// with the ability to set a timeout on a request.
func NewImageConfigurationDeleteImageParamsWithTimeout(timeout time.Duration) *ImageConfigurationDeleteImageParams {
	return &ImageConfigurationDeleteImageParams{
		timeout: timeout,
	}
}

// NewImageConfigurationDeleteImageParamsWithContext creates a new ImageConfigurationDeleteImageParams object
// with the ability to set a context for a request.
func NewImageConfigurationDeleteImageParamsWithContext(ctx context.Context) *ImageConfigurationDeleteImageParams {
	return &ImageConfigurationDeleteImageParams{
		Context: ctx,
	}
}

// NewImageConfigurationDeleteImageParamsWithHTTPClient creates a new ImageConfigurationDeleteImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageConfigurationDeleteImageParamsWithHTTPClient(client *http.Client) *ImageConfigurationDeleteImageParams {
	return &ImageConfigurationDeleteImageParams{
		HTTPClient: client,
	}
}

/*
ImageConfigurationDeleteImageParams contains all the parameters to send to the API endpoint

	for the image configuration delete image operation.

	Typically these are written to a http.Request.
*/
type ImageConfigurationDeleteImageParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   System defined universally unique Id of the image
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image configuration delete image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationDeleteImageParams) WithDefaults() *ImageConfigurationDeleteImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image configuration delete image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationDeleteImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the image configuration delete image params
func (o *ImageConfigurationDeleteImageParams) WithTimeout(timeout time.Duration) *ImageConfigurationDeleteImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image configuration delete image params
func (o *ImageConfigurationDeleteImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image configuration delete image params
func (o *ImageConfigurationDeleteImageParams) WithContext(ctx context.Context) *ImageConfigurationDeleteImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image configuration delete image params
func (o *ImageConfigurationDeleteImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image configuration delete image params
func (o *ImageConfigurationDeleteImageParams) WithHTTPClient(client *http.Client) *ImageConfigurationDeleteImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image configuration delete image params
func (o *ImageConfigurationDeleteImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the image configuration delete image params
func (o *ImageConfigurationDeleteImageParams) WithXRequestID(xRequestID *string) *ImageConfigurationDeleteImageParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the image configuration delete image params
func (o *ImageConfigurationDeleteImageParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the image configuration delete image params
func (o *ImageConfigurationDeleteImageParams) WithID(id string) *ImageConfigurationDeleteImageParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the image configuration delete image params
func (o *ImageConfigurationDeleteImageParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ImageConfigurationDeleteImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
