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

// NewIdentityAccessManagementLoginExternalParams creates a new IdentityAccessManagementLoginExternalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementLoginExternalParams() *IdentityAccessManagementLoginExternalParams {
	return &IdentityAccessManagementLoginExternalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementLoginExternalParamsWithTimeout creates a new IdentityAccessManagementLoginExternalParams object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementLoginExternalParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementLoginExternalParams {
	return &IdentityAccessManagementLoginExternalParams{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementLoginExternalParamsWithContext creates a new IdentityAccessManagementLoginExternalParams object
// with the ability to set a context for a request.
func NewIdentityAccessManagementLoginExternalParamsWithContext(ctx context.Context) *IdentityAccessManagementLoginExternalParams {
	return &IdentityAccessManagementLoginExternalParams{
		Context: ctx,
	}
}

// NewIdentityAccessManagementLoginExternalParamsWithHTTPClient creates a new IdentityAccessManagementLoginExternalParams object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementLoginExternalParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementLoginExternalParams {
	return &IdentityAccessManagementLoginExternalParams{
		HTTPClient: client,
	}
}

/*
IdentityAccessManagementLoginExternalParams contains all the parameters to send to the API endpoint

	for the identity access management login external operation.

	Typically these are written to a http.Request.
*/
type IdentityAccessManagementLoginExternalParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.AAAFrontendLoginWithOauthRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the identity access management login external params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementLoginExternalParams) WithDefaults() *IdentityAccessManagementLoginExternalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management login external params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementLoginExternalParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management login external params
func (o *IdentityAccessManagementLoginExternalParams) WithTimeout(timeout time.Duration) *IdentityAccessManagementLoginExternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management login external params
func (o *IdentityAccessManagementLoginExternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management login external params
func (o *IdentityAccessManagementLoginExternalParams) WithContext(ctx context.Context) *IdentityAccessManagementLoginExternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management login external params
func (o *IdentityAccessManagementLoginExternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management login external params
func (o *IdentityAccessManagementLoginExternalParams) WithHTTPClient(client *http.Client) *IdentityAccessManagementLoginExternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management login external params
func (o *IdentityAccessManagementLoginExternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management login external params
func (o *IdentityAccessManagementLoginExternalParams) WithXRequestID(xRequestID *string) *IdentityAccessManagementLoginExternalParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management login external params
func (o *IdentityAccessManagementLoginExternalParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the identity access management login external params
func (o *IdentityAccessManagementLoginExternalParams) WithBody(body *models.AAAFrontendLoginWithOauthRequest) *IdentityAccessManagementLoginExternalParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the identity access management login external params
func (o *IdentityAccessManagementLoginExternalParams) SetBody(body *models.AAAFrontendLoginWithOauthRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementLoginExternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
