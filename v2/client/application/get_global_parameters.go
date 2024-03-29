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

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams creates a new EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams() *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParamsWithTimeout creates a new EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams object
// with the ability to set a timeout on a request.
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParamsWithTimeout(timeout time.Duration) *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams{
		timeout: timeout,
	}
}

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParamsWithContext creates a new EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams object
// with the ability to set a context for a request.
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParamsWithContext(ctx context.Context) *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams{
		Context: ctx,
	}
}

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParamsWithHTTPClient creates a new EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParamsWithHTTPClient(client *http.Client) *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams{
		HTTPClient: client,
	}
}

/*
EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams contains all the parameters to send to the API endpoint

	for the edge application configuration get global edge application bundle operation.

	Typically these are written to a http.Request.
*/
type EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams struct {

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

// WithDefaults hydrates default values in the edge application configuration get global edge application bundle params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams) WithDefaults() *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge application configuration get global edge application bundle params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge application configuration get global edge application bundle params
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams) WithTimeout(timeout time.Duration) *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge application configuration get global edge application bundle params
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge application configuration get global edge application bundle params
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams) WithContext(ctx context.Context) *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge application configuration get global edge application bundle params
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge application configuration get global edge application bundle params
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams) WithHTTPClient(client *http.Client) *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge application configuration get global edge application bundle params
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge application configuration get global edge application bundle params
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams) WithXRequestID(xRequestID *string) *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge application configuration get global edge application bundle params
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the edge application configuration get global edge application bundle params
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams) WithID(id string) *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge application configuration get global edge application bundle params
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
