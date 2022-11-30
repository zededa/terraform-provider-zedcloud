// Code generated by go-swagger; DO NOT EDIT.

package device_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetDeviceAttestationParams creates a new GetDeviceAttestationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceAttestationParams() *GetDeviceAttestationParams {
	return &GetDeviceAttestationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceAttestationParamsWithTimeout creates a new GetDeviceAttestationParams object
// with the ability to set a timeout on a request.
func NewGetDeviceAttestationParamsWithTimeout(timeout time.Duration) *GetDeviceAttestationParams {
	return &GetDeviceAttestationParams{
		timeout: timeout,
	}
}

// NewGetDeviceAttestationParamsWithContext creates a new GetDeviceAttestationParams object
// with the ability to set a context for a request.
func NewGetDeviceAttestationParamsWithContext(ctx context.Context) *GetDeviceAttestationParams {
	return &GetDeviceAttestationParams{
		Context: ctx,
	}
}

// NewGetDeviceAttestationParamsWithHTTPClient creates a new GetDeviceAttestationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceAttestationParamsWithHTTPClient(client *http.Client) *GetDeviceAttestationParams {
	return &GetDeviceAttestationParams{
		HTTPClient: client,
	}
}

/*
GetDeviceAttestationParams contains all the parameters to send to the API endpoint

	for the get device attestation operation.

	Typically these are written to a http.Request.
*/
type GetDeviceAttestationParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   system generated unique id for a device
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device attestation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceAttestationParams) WithDefaults() *GetDeviceAttestationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device attestation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceAttestationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device attestation params
func (o *GetDeviceAttestationParams) WithTimeout(timeout time.Duration) *GetDeviceAttestationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device attestation params
func (o *GetDeviceAttestationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device attestation params
func (o *GetDeviceAttestationParams) WithContext(ctx context.Context) *GetDeviceAttestationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device attestation params
func (o *GetDeviceAttestationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device attestation params
func (o *GetDeviceAttestationParams) WithHTTPClient(client *http.Client) *GetDeviceAttestationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device attestation params
func (o *GetDeviceAttestationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get device attestation params
func (o *GetDeviceAttestationParams) WithXRequestID(xRequestID *string) *GetDeviceAttestationParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get device attestation params
func (o *GetDeviceAttestationParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the get device attestation params
func (o *GetDeviceAttestationParams) WithID(id string) *GetDeviceAttestationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get device attestation params
func (o *GetDeviceAttestationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceAttestationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
