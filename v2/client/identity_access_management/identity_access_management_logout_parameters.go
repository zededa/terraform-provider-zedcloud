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

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// NewIdentityAccessManagementLogoutParams creates a new IdentityAccessManagementLogoutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementLogoutParams() *IdentityAccessManagementLogoutParams {
	return &IdentityAccessManagementLogoutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementLogoutParamsWithTimeout creates a new IdentityAccessManagementLogoutParams object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementLogoutParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementLogoutParams {
	return &IdentityAccessManagementLogoutParams{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementLogoutParamsWithContext creates a new IdentityAccessManagementLogoutParams object
// with the ability to set a context for a request.
func NewIdentityAccessManagementLogoutParamsWithContext(ctx context.Context) *IdentityAccessManagementLogoutParams {
	return &IdentityAccessManagementLogoutParams{
		Context: ctx,
	}
}

// NewIdentityAccessManagementLogoutParamsWithHTTPClient creates a new IdentityAccessManagementLogoutParams object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementLogoutParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementLogoutParams {
	return &IdentityAccessManagementLogoutParams{
		HTTPClient: client,
	}
}

/*
IdentityAccessManagementLogoutParams contains all the parameters to send to the API endpoint

	for the identity access management logout operation.

	Typically these are written to a http.Request.
*/
type IdentityAccessManagementLogoutParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body models.AAAFrontendLogoutRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the identity access management logout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementLogoutParams) WithDefaults() *IdentityAccessManagementLogoutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management logout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementLogoutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management logout params
func (o *IdentityAccessManagementLogoutParams) WithTimeout(timeout time.Duration) *IdentityAccessManagementLogoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management logout params
func (o *IdentityAccessManagementLogoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management logout params
func (o *IdentityAccessManagementLogoutParams) WithContext(ctx context.Context) *IdentityAccessManagementLogoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management logout params
func (o *IdentityAccessManagementLogoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management logout params
func (o *IdentityAccessManagementLogoutParams) WithHTTPClient(client *http.Client) *IdentityAccessManagementLogoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management logout params
func (o *IdentityAccessManagementLogoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management logout params
func (o *IdentityAccessManagementLogoutParams) WithXRequestID(xRequestID *string) *IdentityAccessManagementLogoutParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management logout params
func (o *IdentityAccessManagementLogoutParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the identity access management logout params
func (o *IdentityAccessManagementLogoutParams) WithBody(body models.AAAFrontendLogoutRequest) *IdentityAccessManagementLogoutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the identity access management logout params
func (o *IdentityAccessManagementLogoutParams) SetBody(body models.AAAFrontendLogoutRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementLogoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
