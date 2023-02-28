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

// GetByNameParams creates a new ImageConfigurationGetImageByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func GetByNameParams() *ImageConfigurationGetImageByNameParams {
	return &ImageConfigurationGetImageByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageConfigurationGetImageByNameParamsWithTimeout creates a new ImageConfigurationGetImageByNameParams object
// with the ability to set a timeout on a request.
func NewImageConfigurationGetImageByNameParamsWithTimeout(timeout time.Duration) *ImageConfigurationGetImageByNameParams {
	return &ImageConfigurationGetImageByNameParams{
		timeout: timeout,
	}
}

// NewImageConfigurationGetImageByNameParamsWithContext creates a new ImageConfigurationGetImageByNameParams object
// with the ability to set a context for a request.
func NewImageConfigurationGetImageByNameParamsWithContext(ctx context.Context) *ImageConfigurationGetImageByNameParams {
	return &ImageConfigurationGetImageByNameParams{
		Context: ctx,
	}
}

// NewImageConfigurationGetImageByNameParamsWithHTTPClient creates a new ImageConfigurationGetImageByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageConfigurationGetImageByNameParamsWithHTTPClient(client *http.Client) *ImageConfigurationGetImageByNameParams {
	return &ImageConfigurationGetImageByNameParams{
		HTTPClient: client,
	}
}

/*
ImageConfigurationGetImageByNameParams contains all the parameters to send to the API endpoint

	for the image configuration get image by name operation.

	Typically these are written to a http.Request.
*/
type ImageConfigurationGetImageByNameParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* Name.

	   User defined name of the image, unique across the enterprise. Once image is created, name can’t be changed.
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image configuration get image by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationGetImageByNameParams) WithDefaults() *ImageConfigurationGetImageByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image configuration get image by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationGetImageByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the image configuration get image by name params
func (o *ImageConfigurationGetImageByNameParams) WithTimeout(timeout time.Duration) *ImageConfigurationGetImageByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image configuration get image by name params
func (o *ImageConfigurationGetImageByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image configuration get image by name params
func (o *ImageConfigurationGetImageByNameParams) WithContext(ctx context.Context) *ImageConfigurationGetImageByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image configuration get image by name params
func (o *ImageConfigurationGetImageByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image configuration get image by name params
func (o *ImageConfigurationGetImageByNameParams) WithHTTPClient(client *http.Client) *ImageConfigurationGetImageByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image configuration get image by name params
func (o *ImageConfigurationGetImageByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the image configuration get image by name params
func (o *ImageConfigurationGetImageByNameParams) WithXRequestID(xRequestID *string) *ImageConfigurationGetImageByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the image configuration get image by name params
func (o *ImageConfigurationGetImageByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the image configuration get image by name params
func (o *ImageConfigurationGetImageByNameParams) WithName(name string) *ImageConfigurationGetImageByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the image configuration get image by name params
func (o *ImageConfigurationGetImageByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ImageConfigurationGetImageByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
