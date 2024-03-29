package application

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

// DeleteParams creates a new EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func DeleteParams() *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams {
	return &EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleParamsWithTimeout creates a new EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams object
// with the ability to set a timeout on a request.
func NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleParamsWithTimeout(timeout time.Duration) *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams {
	return &EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams{
		timeout: timeout,
	}
}

// NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleParamsWithContext creates a new EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams object
// with the ability to set a context for a request.
func NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleParamsWithContext(ctx context.Context) *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams {
	return &EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams{
		Context: ctx,
	}
}

// NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleParamsWithHTTPClient creates a new EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleParamsWithHTTPClient(client *http.Client) *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams {
	return &EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams{
		HTTPClient: client,
	}
}

/*
EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams contains all the parameters to send to the API endpoint

	for the edge application configuration delete edge application bundle operation.

	Typically these are written to a http.Request.
*/
type EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   System defined universally unique Id of the edge application
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge application configuration delete edge application bundle params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams) WithDefaults() *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge application configuration delete edge application bundle params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge application configuration delete edge application bundle params
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams) WithTimeout(timeout time.Duration) *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge application configuration delete edge application bundle params
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge application configuration delete edge application bundle params
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams) WithContext(ctx context.Context) *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge application configuration delete edge application bundle params
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge application configuration delete edge application bundle params
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams) WithHTTPClient(client *http.Client) *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge application configuration delete edge application bundle params
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge application configuration delete edge application bundle params
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams) WithXRequestID(xRequestID *string) *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge application configuration delete edge application bundle params
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the edge application configuration delete edge application bundle params
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams) WithID(id string) *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge application configuration delete edge application bundle params
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
