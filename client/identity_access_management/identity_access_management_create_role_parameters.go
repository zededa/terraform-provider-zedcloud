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

	"github.com/zededa/terraform-provider/models"
)

// NewIdentityAccessManagementCreateRoleParams creates a new IdentityAccessManagementCreateRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementCreateRoleParams() *IdentityAccessManagementCreateRoleParams {
	return &IdentityAccessManagementCreateRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementCreateRoleParamsWithTimeout creates a new IdentityAccessManagementCreateRoleParams object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementCreateRoleParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementCreateRoleParams {
	return &IdentityAccessManagementCreateRoleParams{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementCreateRoleParamsWithContext creates a new IdentityAccessManagementCreateRoleParams object
// with the ability to set a context for a request.
func NewIdentityAccessManagementCreateRoleParamsWithContext(ctx context.Context) *IdentityAccessManagementCreateRoleParams {
	return &IdentityAccessManagementCreateRoleParams{
		Context: ctx,
	}
}

// NewIdentityAccessManagementCreateRoleParamsWithHTTPClient creates a new IdentityAccessManagementCreateRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementCreateRoleParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementCreateRoleParams {
	return &IdentityAccessManagementCreateRoleParams{
		HTTPClient: client,
	}
}

/*
IdentityAccessManagementCreateRoleParams contains all the parameters to send to the API endpoint

	for the identity access management create role operation.

	Typically these are written to a http.Request.
*/
type IdentityAccessManagementCreateRoleParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.Role

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the identity access management create role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementCreateRoleParams) WithDefaults() *IdentityAccessManagementCreateRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management create role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementCreateRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management create role params
func (o *IdentityAccessManagementCreateRoleParams) WithTimeout(timeout time.Duration) *IdentityAccessManagementCreateRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management create role params
func (o *IdentityAccessManagementCreateRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management create role params
func (o *IdentityAccessManagementCreateRoleParams) WithContext(ctx context.Context) *IdentityAccessManagementCreateRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management create role params
func (o *IdentityAccessManagementCreateRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management create role params
func (o *IdentityAccessManagementCreateRoleParams) WithHTTPClient(client *http.Client) *IdentityAccessManagementCreateRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management create role params
func (o *IdentityAccessManagementCreateRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management create role params
func (o *IdentityAccessManagementCreateRoleParams) WithXRequestID(xRequestID *string) *IdentityAccessManagementCreateRoleParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management create role params
func (o *IdentityAccessManagementCreateRoleParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the identity access management create role params
func (o *IdentityAccessManagementCreateRoleParams) WithBody(body *models.Role) *IdentityAccessManagementCreateRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the identity access management create role params
func (o *IdentityAccessManagementCreateRoleParams) SetBody(body *models.Role) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementCreateRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
