// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

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

// NewIdentityAccessManagementDeleteRealmParams creates a new IdentityAccessManagementDeleteRealmParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementDeleteRealmParams() *IdentityAccessManagementDeleteRealmParams {
	return &IdentityAccessManagementDeleteRealmParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementDeleteRealmParamsWithTimeout creates a new IdentityAccessManagementDeleteRealmParams object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementDeleteRealmParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementDeleteRealmParams {
	return &IdentityAccessManagementDeleteRealmParams{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementDeleteRealmParamsWithContext creates a new IdentityAccessManagementDeleteRealmParams object
// with the ability to set a context for a request.
func NewIdentityAccessManagementDeleteRealmParamsWithContext(ctx context.Context) *IdentityAccessManagementDeleteRealmParams {
	return &IdentityAccessManagementDeleteRealmParams{
		Context: ctx,
	}
}

// NewIdentityAccessManagementDeleteRealmParamsWithHTTPClient creates a new IdentityAccessManagementDeleteRealmParams object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementDeleteRealmParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementDeleteRealmParams {
	return &IdentityAccessManagementDeleteRealmParams{
		HTTPClient: client,
	}
}

/*
IdentityAccessManagementDeleteRealmParams contains all the parameters to send to the API endpoint

	for the identity access management delete realm operation.

	Typically these are written to a http.Request.
*/
type IdentityAccessManagementDeleteRealmParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the identity access management delete realm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementDeleteRealmParams) WithDefaults() *IdentityAccessManagementDeleteRealmParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management delete realm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementDeleteRealmParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management delete realm params
func (o *IdentityAccessManagementDeleteRealmParams) WithTimeout(timeout time.Duration) *IdentityAccessManagementDeleteRealmParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management delete realm params
func (o *IdentityAccessManagementDeleteRealmParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management delete realm params
func (o *IdentityAccessManagementDeleteRealmParams) WithContext(ctx context.Context) *IdentityAccessManagementDeleteRealmParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management delete realm params
func (o *IdentityAccessManagementDeleteRealmParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management delete realm params
func (o *IdentityAccessManagementDeleteRealmParams) WithHTTPClient(client *http.Client) *IdentityAccessManagementDeleteRealmParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management delete realm params
func (o *IdentityAccessManagementDeleteRealmParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management delete realm params
func (o *IdentityAccessManagementDeleteRealmParams) WithXRequestID(xRequestID *string) *IdentityAccessManagementDeleteRealmParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management delete realm params
func (o *IdentityAccessManagementDeleteRealmParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the identity access management delete realm params
func (o *IdentityAccessManagementDeleteRealmParams) WithID(id string) *IdentityAccessManagementDeleteRealmParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the identity access management delete realm params
func (o *IdentityAccessManagementDeleteRealmParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementDeleteRealmParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
