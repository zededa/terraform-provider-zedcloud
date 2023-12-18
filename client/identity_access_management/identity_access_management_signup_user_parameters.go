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

// NewIdentityAccessManagementSignupUserParams creates a new IdentityAccessManagementSignupUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementSignupUserParams() *IdentityAccessManagementSignupUserParams {
	return &IdentityAccessManagementSignupUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementSignupUserParamsWithTimeout creates a new IdentityAccessManagementSignupUserParams object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementSignupUserParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementSignupUserParams {
	return &IdentityAccessManagementSignupUserParams{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementSignupUserParamsWithContext creates a new IdentityAccessManagementSignupUserParams object
// with the ability to set a context for a request.
func NewIdentityAccessManagementSignupUserParamsWithContext(ctx context.Context) *IdentityAccessManagementSignupUserParams {
	return &IdentityAccessManagementSignupUserParams{
		Context: ctx,
	}
}

// NewIdentityAccessManagementSignupUserParamsWithHTTPClient creates a new IdentityAccessManagementSignupUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementSignupUserParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementSignupUserParams {
	return &IdentityAccessManagementSignupUserParams{
		HTTPClient: client,
	}
}

/*
IdentityAccessManagementSignupUserParams contains all the parameters to send to the API endpoint

	for the identity access management signup user operation.

	Typically these are written to a http.Request.
*/
type IdentityAccessManagementSignupUserParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.AAARequestAdminUserSignup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the identity access management signup user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementSignupUserParams) WithDefaults() *IdentityAccessManagementSignupUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management signup user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementSignupUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management signup user params
func (o *IdentityAccessManagementSignupUserParams) WithTimeout(timeout time.Duration) *IdentityAccessManagementSignupUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management signup user params
func (o *IdentityAccessManagementSignupUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management signup user params
func (o *IdentityAccessManagementSignupUserParams) WithContext(ctx context.Context) *IdentityAccessManagementSignupUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management signup user params
func (o *IdentityAccessManagementSignupUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management signup user params
func (o *IdentityAccessManagementSignupUserParams) WithHTTPClient(client *http.Client) *IdentityAccessManagementSignupUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management signup user params
func (o *IdentityAccessManagementSignupUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management signup user params
func (o *IdentityAccessManagementSignupUserParams) WithXRequestID(xRequestID *string) *IdentityAccessManagementSignupUserParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management signup user params
func (o *IdentityAccessManagementSignupUserParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the identity access management signup user params
func (o *IdentityAccessManagementSignupUserParams) WithBody(body *models.AAARequestAdminUserSignup) *IdentityAccessManagementSignupUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the identity access management signup user params
func (o *IdentityAccessManagementSignupUserParams) SetBody(body *models.AAARequestAdminUserSignup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementSignupUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
