package node

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

// NewEdgeNodeConfigurationGetEdgeNodeOnboardingParams creates a new EdgeNodeConfigurationGetEdgeNodeOnboardingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeNodeConfigurationGetEdgeNodeOnboardingParams() *EdgeNodeConfigurationGetEdgeNodeOnboardingParams {
	return &EdgeNodeConfigurationGetEdgeNodeOnboardingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNodeConfigurationGetEdgeNodeOnboardingParamsWithTimeout creates a new EdgeNodeConfigurationGetEdgeNodeOnboardingParams object
// with the ability to set a timeout on a request.
func NewEdgeNodeConfigurationGetEdgeNodeOnboardingParamsWithTimeout(timeout time.Duration) *EdgeNodeConfigurationGetEdgeNodeOnboardingParams {
	return &EdgeNodeConfigurationGetEdgeNodeOnboardingParams{
		timeout: timeout,
	}
}

// NewEdgeNodeConfigurationGetEdgeNodeOnboardingParamsWithContext creates a new EdgeNodeConfigurationGetEdgeNodeOnboardingParams object
// with the ability to set a context for a request.
func NewEdgeNodeConfigurationGetEdgeNodeOnboardingParamsWithContext(ctx context.Context) *EdgeNodeConfigurationGetEdgeNodeOnboardingParams {
	return &EdgeNodeConfigurationGetEdgeNodeOnboardingParams{
		Context: ctx,
	}
}

// NewEdgeNodeConfigurationGetEdgeNodeOnboardingParamsWithHTTPClient creates a new EdgeNodeConfigurationGetEdgeNodeOnboardingParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNodeConfigurationGetEdgeNodeOnboardingParamsWithHTTPClient(client *http.Client) *EdgeNodeConfigurationGetEdgeNodeOnboardingParams {
	return &EdgeNodeConfigurationGetEdgeNodeOnboardingParams{
		HTTPClient: client,
	}
}

/*
EdgeNodeConfigurationGetEdgeNodeOnboardingParams contains all the parameters to send to the API endpoint

	for the edge node configuration get edge node onboarding operation.

	Typically these are written to a http.Request.
*/
type EdgeNodeConfigurationGetEdgeNodeOnboardingParams struct {

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

// WithDefaults hydrates default values in the edge node configuration get edge node onboarding params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingParams) WithDefaults() *EdgeNodeConfigurationGetEdgeNodeOnboardingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge node configuration get edge node onboarding params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge node configuration get edge node onboarding params
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingParams) WithTimeout(timeout time.Duration) *EdgeNodeConfigurationGetEdgeNodeOnboardingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge node configuration get edge node onboarding params
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge node configuration get edge node onboarding params
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingParams) WithContext(ctx context.Context) *EdgeNodeConfigurationGetEdgeNodeOnboardingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge node configuration get edge node onboarding params
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge node configuration get edge node onboarding params
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingParams) WithHTTPClient(client *http.Client) *EdgeNodeConfigurationGetEdgeNodeOnboardingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge node configuration get edge node onboarding params
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge node configuration get edge node onboarding params
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingParams) WithXRequestID(xRequestID *string) *EdgeNodeConfigurationGetEdgeNodeOnboardingParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge node configuration get edge node onboarding params
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the edge node configuration get edge node onboarding params
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingParams) WithID(id string) *EdgeNodeConfigurationGetEdgeNodeOnboardingParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge node configuration get edge node onboarding params
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNodeConfigurationGetEdgeNodeOnboardingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
