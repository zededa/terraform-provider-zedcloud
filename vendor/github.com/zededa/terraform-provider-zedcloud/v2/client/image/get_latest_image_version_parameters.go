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

// NewImageConfigurationGetLatestImageVersionParams creates a new ImageConfigurationGetLatestImageVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImageConfigurationGetLatestImageVersionParams() *ImageConfigurationGetLatestImageVersionParams {
	return &ImageConfigurationGetLatestImageVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageConfigurationGetLatestImageVersionParamsWithTimeout creates a new ImageConfigurationGetLatestImageVersionParams object
// with the ability to set a timeout on a request.
func NewImageConfigurationGetLatestImageVersionParamsWithTimeout(timeout time.Duration) *ImageConfigurationGetLatestImageVersionParams {
	return &ImageConfigurationGetLatestImageVersionParams{
		timeout: timeout,
	}
}

// NewImageConfigurationGetLatestImageVersionParamsWithContext creates a new ImageConfigurationGetLatestImageVersionParams object
// with the ability to set a context for a request.
func NewImageConfigurationGetLatestImageVersionParamsWithContext(ctx context.Context) *ImageConfigurationGetLatestImageVersionParams {
	return &ImageConfigurationGetLatestImageVersionParams{
		Context: ctx,
	}
}

// NewImageConfigurationGetLatestImageVersionParamsWithHTTPClient creates a new ImageConfigurationGetLatestImageVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageConfigurationGetLatestImageVersionParamsWithHTTPClient(client *http.Client) *ImageConfigurationGetLatestImageVersionParams {
	return &ImageConfigurationGetLatestImageVersionParams{
		HTTPClient: client,
	}
}

/*
ImageConfigurationGetLatestImageVersionParams contains all the parameters to send to the API endpoint

	for the image configuration get latest image version operation.

	Typically these are written to a http.Request.
*/
type ImageConfigurationGetLatestImageVersionParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// EnterpriseID.
	EnterpriseID *string

	/* ImageArch.

	   Image architecture to be matched : 'AMD64' or 'ARM64'.
	*/
	ImageArch string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image configuration get latest image version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationGetLatestImageVersionParams) WithDefaults() *ImageConfigurationGetLatestImageVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image configuration get latest image version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationGetLatestImageVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the image configuration get latest image version params
func (o *ImageConfigurationGetLatestImageVersionParams) WithTimeout(timeout time.Duration) *ImageConfigurationGetLatestImageVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image configuration get latest image version params
func (o *ImageConfigurationGetLatestImageVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image configuration get latest image version params
func (o *ImageConfigurationGetLatestImageVersionParams) WithContext(ctx context.Context) *ImageConfigurationGetLatestImageVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image configuration get latest image version params
func (o *ImageConfigurationGetLatestImageVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image configuration get latest image version params
func (o *ImageConfigurationGetLatestImageVersionParams) WithHTTPClient(client *http.Client) *ImageConfigurationGetLatestImageVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image configuration get latest image version params
func (o *ImageConfigurationGetLatestImageVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the image configuration get latest image version params
func (o *ImageConfigurationGetLatestImageVersionParams) WithXRequestID(xRequestID *string) *ImageConfigurationGetLatestImageVersionParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the image configuration get latest image version params
func (o *ImageConfigurationGetLatestImageVersionParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithEnterpriseID adds the enterpriseID to the image configuration get latest image version params
func (o *ImageConfigurationGetLatestImageVersionParams) WithEnterpriseID(enterpriseID *string) *ImageConfigurationGetLatestImageVersionParams {
	o.SetEnterpriseID(enterpriseID)
	return o
}

// SetEnterpriseID adds the enterpriseId to the image configuration get latest image version params
func (o *ImageConfigurationGetLatestImageVersionParams) SetEnterpriseID(enterpriseID *string) {
	o.EnterpriseID = enterpriseID
}

// WithImageArch adds the imageArch to the image configuration get latest image version params
func (o *ImageConfigurationGetLatestImageVersionParams) WithImageArch(imageArch string) *ImageConfigurationGetLatestImageVersionParams {
	o.SetImageArch(imageArch)
	return o
}

// SetImageArch adds the imageArch to the image configuration get latest image version params
func (o *ImageConfigurationGetLatestImageVersionParams) SetImageArch(imageArch string) {
	o.ImageArch = imageArch
}

// WriteToRequest writes these params to a swagger request
func (o *ImageConfigurationGetLatestImageVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EnterpriseID != nil {

		// query param enterpriseId
		var qrEnterpriseID string

		if o.EnterpriseID != nil {
			qrEnterpriseID = *o.EnterpriseID
		}
		qEnterpriseID := qrEnterpriseID
		if qEnterpriseID != "" {

			if err := r.SetQueryParam("enterpriseId", qEnterpriseID); err != nil {
				return err
			}
		}
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
