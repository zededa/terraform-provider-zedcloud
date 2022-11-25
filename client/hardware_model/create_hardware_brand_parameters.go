// Code generated by go-swagger; DO NOT EDIT.

package hardware_model

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

	"github.com/zededa/terraform-provider/models"
)

// NewCreateHardwareBrandParams creates a new CreateHardwareBrandParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateHardwareBrandParams() *CreateHardwareBrandParams {
	return &CreateHardwareBrandParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateHardwareBrandParamsWithTimeout creates a new CreateHardwareBrandParams object
// with the ability to set a timeout on a request.
func NewCreateHardwareBrandParamsWithTimeout(timeout time.Duration) *CreateHardwareBrandParams {
	return &CreateHardwareBrandParams{
		timeout: timeout,
	}
}

// NewCreateHardwareBrandParamsWithContext creates a new CreateHardwareBrandParams object
// with the ability to set a context for a request.
func NewCreateHardwareBrandParamsWithContext(ctx context.Context) *CreateHardwareBrandParams {
	return &CreateHardwareBrandParams{
		Context: ctx,
	}
}

// NewCreateHardwareBrandParamsWithHTTPClient creates a new CreateHardwareBrandParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateHardwareBrandParamsWithHTTPClient(client *http.Client) *CreateHardwareBrandParams {
	return &CreateHardwareBrandParams{
		HTTPClient: client,
	}
}

/*
CreateHardwareBrandParams contains all the parameters to send to the API endpoint

	for the create hardware brand operation.

	Typically these are written to a http.Request.
*/
type CreateHardwareBrandParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.SysBrand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create hardware brand params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateHardwareBrandParams) WithDefaults() *CreateHardwareBrandParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create hardware brand params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateHardwareBrandParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create hardware brand params
func (o *CreateHardwareBrandParams) WithTimeout(timeout time.Duration) *CreateHardwareBrandParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create hardware brand params
func (o *CreateHardwareBrandParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create hardware brand params
func (o *CreateHardwareBrandParams) WithContext(ctx context.Context) *CreateHardwareBrandParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create hardware brand params
func (o *CreateHardwareBrandParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create hardware brand params
func (o *CreateHardwareBrandParams) WithHTTPClient(client *http.Client) *CreateHardwareBrandParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create hardware brand params
func (o *CreateHardwareBrandParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create hardware brand params
func (o *CreateHardwareBrandParams) WithXRequestID(xRequestID *string) *CreateHardwareBrandParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create hardware brand params
func (o *CreateHardwareBrandParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the create hardware brand params
func (o *CreateHardwareBrandParams) WithBody(body *models.SysBrand) *CreateHardwareBrandParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create hardware brand params
func (o *CreateHardwareBrandParams) SetBody(body *models.SysBrand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateHardwareBrandParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
