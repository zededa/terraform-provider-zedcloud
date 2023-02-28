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

// NewImageConfigurationMarkEveImageLatest2Params creates a new ImageConfigurationMarkEveImageLatest2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImageConfigurationMarkEveImageLatest2Params() *ImageConfigurationMarkEveImageLatest2Params {
	return &ImageConfigurationMarkEveImageLatest2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageConfigurationMarkEveImageLatest2ParamsWithTimeout creates a new ImageConfigurationMarkEveImageLatest2Params object
// with the ability to set a timeout on a request.
func NewImageConfigurationMarkEveImageLatest2ParamsWithTimeout(timeout time.Duration) *ImageConfigurationMarkEveImageLatest2Params {
	return &ImageConfigurationMarkEveImageLatest2Params{
		timeout: timeout,
	}
}

// NewImageConfigurationMarkEveImageLatest2ParamsWithContext creates a new ImageConfigurationMarkEveImageLatest2Params object
// with the ability to set a context for a request.
func NewImageConfigurationMarkEveImageLatest2ParamsWithContext(ctx context.Context) *ImageConfigurationMarkEveImageLatest2Params {
	return &ImageConfigurationMarkEveImageLatest2Params{
		Context: ctx,
	}
}

// NewImageConfigurationMarkEveImageLatest2ParamsWithHTTPClient creates a new ImageConfigurationMarkEveImageLatest2Params object
// with the ability to set a custom HTTPClient for a request.
func NewImageConfigurationMarkEveImageLatest2ParamsWithHTTPClient(client *http.Client) *ImageConfigurationMarkEveImageLatest2Params {
	return &ImageConfigurationMarkEveImageLatest2Params{
		HTTPClient: client,
	}
}

/*
ImageConfigurationMarkEveImageLatest2Params contains all the parameters to send to the API endpoint

	for the image configuration mark eve image latest2 operation.

	Typically these are written to a http.Request.
*/
type ImageConfigurationMarkEveImageLatest2Params struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body ImageConfigurationMarkEveImageLatest2Body

	/* ImageArch.

	   Image Architecture.
	*/
	ImageArch string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image configuration mark eve image latest2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationMarkEveImageLatest2Params) WithDefaults() *ImageConfigurationMarkEveImageLatest2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image configuration mark eve image latest2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationMarkEveImageLatest2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the image configuration mark eve image latest2 params
func (o *ImageConfigurationMarkEveImageLatest2Params) WithTimeout(timeout time.Duration) *ImageConfigurationMarkEveImageLatest2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image configuration mark eve image latest2 params
func (o *ImageConfigurationMarkEveImageLatest2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image configuration mark eve image latest2 params
func (o *ImageConfigurationMarkEveImageLatest2Params) WithContext(ctx context.Context) *ImageConfigurationMarkEveImageLatest2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image configuration mark eve image latest2 params
func (o *ImageConfigurationMarkEveImageLatest2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image configuration mark eve image latest2 params
func (o *ImageConfigurationMarkEveImageLatest2Params) WithHTTPClient(client *http.Client) *ImageConfigurationMarkEveImageLatest2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image configuration mark eve image latest2 params
func (o *ImageConfigurationMarkEveImageLatest2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the image configuration mark eve image latest2 params
func (o *ImageConfigurationMarkEveImageLatest2Params) WithXRequestID(xRequestID *string) *ImageConfigurationMarkEveImageLatest2Params {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the image configuration mark eve image latest2 params
func (o *ImageConfigurationMarkEveImageLatest2Params) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the image configuration mark eve image latest2 params
func (o *ImageConfigurationMarkEveImageLatest2Params) WithBody(body ImageConfigurationMarkEveImageLatest2Body) *ImageConfigurationMarkEveImageLatest2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the image configuration mark eve image latest2 params
func (o *ImageConfigurationMarkEveImageLatest2Params) SetBody(body ImageConfigurationMarkEveImageLatest2Body) {
	o.Body = body
}

// WithImageArch adds the imageArch to the image configuration mark eve image latest2 params
func (o *ImageConfigurationMarkEveImageLatest2Params) WithImageArch(imageArch string) *ImageConfigurationMarkEveImageLatest2Params {
	o.SetImageArch(imageArch)
	return o
}

// SetImageArch adds the imageArch to the image configuration mark eve image latest2 params
func (o *ImageConfigurationMarkEveImageLatest2Params) SetImageArch(imageArch string) {
	o.ImageArch = imageArch
}

// WriteToRequest writes these params to a swagger request
func (o *ImageConfigurationMarkEveImageLatest2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param imageArch
	if err := r.SetPathParam("imageArch", o.ImageArch); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
