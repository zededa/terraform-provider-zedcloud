package image_configuration

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// GetByIDParams creates a new ImageConfigurationGetImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func GetByIDParams() *ImageConfigurationGetImageParams {
	return &ImageConfigurationGetImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageConfigurationGetImageParamsWithTimeout creates a new ImageConfigurationGetImageParams object
// with the ability to set a timeout on a request.
func NewImageConfigurationGetImageParamsWithTimeout(timeout time.Duration) *ImageConfigurationGetImageParams {
	return &ImageConfigurationGetImageParams{
		timeout: timeout,
	}
}

// NewImageConfigurationGetImageParamsWithContext creates a new ImageConfigurationGetImageParams object
// with the ability to set a context for a request.
func NewImageConfigurationGetImageParamsWithContext(ctx context.Context) *ImageConfigurationGetImageParams {
	return &ImageConfigurationGetImageParams{
		Context: ctx,
	}
}

// NewImageConfigurationGetImageParamsWithHTTPClient creates a new ImageConfigurationGetImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageConfigurationGetImageParamsWithHTTPClient(client *http.Client) *ImageConfigurationGetImageParams {
	return &ImageConfigurationGetImageParams{
		HTTPClient: client,
	}
}

/*
ImageConfigurationGetImageParams contains all the parameters to send to the API endpoint

	for the image configuration get image operation.

	Typically these are written to a http.Request.
*/
type ImageConfigurationGetImageParams struct {

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

// WithDefaults hydrates default values in the image configuration get image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationGetImageParams) WithDefaults() *ImageConfigurationGetImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image configuration get image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationGetImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the image configuration get image params
func (o *ImageConfigurationGetImageParams) WithTimeout(timeout time.Duration) *ImageConfigurationGetImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image configuration get image params
func (o *ImageConfigurationGetImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image configuration get image params
func (o *ImageConfigurationGetImageParams) WithContext(ctx context.Context) *ImageConfigurationGetImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image configuration get image params
func (o *ImageConfigurationGetImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image configuration get image params
func (o *ImageConfigurationGetImageParams) WithHTTPClient(client *http.Client) *ImageConfigurationGetImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image configuration get image params
func (o *ImageConfigurationGetImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the image configuration get image params
func (o *ImageConfigurationGetImageParams) WithXRequestID(xRequestID *string) *ImageConfigurationGetImageParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the image configuration get image params
func (o *ImageConfigurationGetImageParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the image configuration get image params
func (o *ImageConfigurationGetImageParams) WithID(id string) *ImageConfigurationGetImageParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the image configuration get image params
func (o *ImageConfigurationGetImageParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ImageConfigurationGetImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
