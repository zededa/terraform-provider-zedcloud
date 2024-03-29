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

// NewIdentityAccessManagementGetUserSessionSelfParams creates a new IdentityAccessManagementGetUserSessionSelfParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementGetUserSessionSelfParams() *IdentityAccessManagementGetUserSessionSelfParams {
	return &IdentityAccessManagementGetUserSessionSelfParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementGetUserSessionSelfParamsWithTimeout creates a new IdentityAccessManagementGetUserSessionSelfParams object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementGetUserSessionSelfParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementGetUserSessionSelfParams {
	return &IdentityAccessManagementGetUserSessionSelfParams{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementGetUserSessionSelfParamsWithContext creates a new IdentityAccessManagementGetUserSessionSelfParams object
// with the ability to set a context for a request.
func NewIdentityAccessManagementGetUserSessionSelfParamsWithContext(ctx context.Context) *IdentityAccessManagementGetUserSessionSelfParams {
	return &IdentityAccessManagementGetUserSessionSelfParams{
		Context: ctx,
	}
}

// NewIdentityAccessManagementGetUserSessionSelfParamsWithHTTPClient creates a new IdentityAccessManagementGetUserSessionSelfParams object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementGetUserSessionSelfParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementGetUserSessionSelfParams {
	return &IdentityAccessManagementGetUserSessionSelfParams{
		HTTPClient: client,
	}
}

/*
IdentityAccessManagementGetUserSessionSelfParams contains all the parameters to send to the API endpoint

	for the identity access management get user session self operation.

	Typically these are written to a http.Request.
*/
type IdentityAccessManagementGetUserSessionSelfParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the identity access management get user session self params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementGetUserSessionSelfParams) WithDefaults() *IdentityAccessManagementGetUserSessionSelfParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management get user session self params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementGetUserSessionSelfParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management get user session self params
func (o *IdentityAccessManagementGetUserSessionSelfParams) WithTimeout(timeout time.Duration) *IdentityAccessManagementGetUserSessionSelfParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management get user session self params
func (o *IdentityAccessManagementGetUserSessionSelfParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management get user session self params
func (o *IdentityAccessManagementGetUserSessionSelfParams) WithContext(ctx context.Context) *IdentityAccessManagementGetUserSessionSelfParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management get user session self params
func (o *IdentityAccessManagementGetUserSessionSelfParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management get user session self params
func (o *IdentityAccessManagementGetUserSessionSelfParams) WithHTTPClient(client *http.Client) *IdentityAccessManagementGetUserSessionSelfParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management get user session self params
func (o *IdentityAccessManagementGetUserSessionSelfParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management get user session self params
func (o *IdentityAccessManagementGetUserSessionSelfParams) WithXRequestID(xRequestID *string) *IdentityAccessManagementGetUserSessionSelfParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management get user session self params
func (o *IdentityAccessManagementGetUserSessionSelfParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementGetUserSessionSelfParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
