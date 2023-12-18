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

// NewIdentityAccessManagementDeleteCredentialParams creates a new IdentityAccessManagementDeleteCredentialParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementDeleteCredentialParams() *IdentityAccessManagementDeleteCredentialParams {
	return &IdentityAccessManagementDeleteCredentialParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementDeleteCredentialParamsWithTimeout creates a new IdentityAccessManagementDeleteCredentialParams object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementDeleteCredentialParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementDeleteCredentialParams {
	return &IdentityAccessManagementDeleteCredentialParams{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementDeleteCredentialParamsWithContext creates a new IdentityAccessManagementDeleteCredentialParams object
// with the ability to set a context for a request.
func NewIdentityAccessManagementDeleteCredentialParamsWithContext(ctx context.Context) *IdentityAccessManagementDeleteCredentialParams {
	return &IdentityAccessManagementDeleteCredentialParams{
		Context: ctx,
	}
}

// NewIdentityAccessManagementDeleteCredentialParamsWithHTTPClient creates a new IdentityAccessManagementDeleteCredentialParams object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementDeleteCredentialParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementDeleteCredentialParams {
	return &IdentityAccessManagementDeleteCredentialParams{
		HTTPClient: client,
	}
}

/*
IdentityAccessManagementDeleteCredentialParams contains all the parameters to send to the API endpoint

	for the identity access management delete credential operation.

	Typically these are written to a http.Request.
*/
type IdentityAccessManagementDeleteCredentialParams struct {

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

// WithDefaults hydrates default values in the identity access management delete credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementDeleteCredentialParams) WithDefaults() *IdentityAccessManagementDeleteCredentialParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management delete credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementDeleteCredentialParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management delete credential params
func (o *IdentityAccessManagementDeleteCredentialParams) WithTimeout(timeout time.Duration) *IdentityAccessManagementDeleteCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management delete credential params
func (o *IdentityAccessManagementDeleteCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management delete credential params
func (o *IdentityAccessManagementDeleteCredentialParams) WithContext(ctx context.Context) *IdentityAccessManagementDeleteCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management delete credential params
func (o *IdentityAccessManagementDeleteCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management delete credential params
func (o *IdentityAccessManagementDeleteCredentialParams) WithHTTPClient(client *http.Client) *IdentityAccessManagementDeleteCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management delete credential params
func (o *IdentityAccessManagementDeleteCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management delete credential params
func (o *IdentityAccessManagementDeleteCredentialParams) WithXRequestID(xRequestID *string) *IdentityAccessManagementDeleteCredentialParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management delete credential params
func (o *IdentityAccessManagementDeleteCredentialParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the identity access management delete credential params
func (o *IdentityAccessManagementDeleteCredentialParams) WithID(id string) *IdentityAccessManagementDeleteCredentialParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the identity access management delete credential params
func (o *IdentityAccessManagementDeleteCredentialParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementDeleteCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
