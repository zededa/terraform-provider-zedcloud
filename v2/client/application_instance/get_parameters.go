package application_instance

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

// GetByIDParams creates a new EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func GetByIDParams() *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams {
	return &EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParamsWithTimeout creates a new EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams object
// with the ability to set a timeout on a request.
func NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParamsWithTimeout(timeout time.Duration) *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams {
	return &EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams{
		timeout: timeout,
	}
}

// NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParamsWithContext creates a new EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams object
// with the ability to set a context for a request.
func NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParamsWithContext(ctx context.Context) *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams {
	return &EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams{
		Context: ctx,
	}
}

// NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParamsWithHTTPClient creates a new EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParamsWithHTTPClient(client *http.Client) *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams {
	return &EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams{
		HTTPClient: client,
	}
}

/*
EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams contains all the parameters to send to the API endpoint

	for the edge application instance configuration get edge application instance operation.

	Typically these are written to a http.Request.
*/
type EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   System defined universally unique Id of the app instance
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge application instance configuration get edge application instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams) WithDefaults() *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge application instance configuration get edge application instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge application instance configuration get edge application instance params
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams) WithTimeout(timeout time.Duration) *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge application instance configuration get edge application instance params
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge application instance configuration get edge application instance params
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams) WithContext(ctx context.Context) *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge application instance configuration get edge application instance params
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge application instance configuration get edge application instance params
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams) WithHTTPClient(client *http.Client) *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge application instance configuration get edge application instance params
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge application instance configuration get edge application instance params
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams) WithXRequestID(xRequestID *string) *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge application instance configuration get edge application instance params
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the edge application instance configuration get edge application instance params
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams) WithID(id string) *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge application instance configuration get edge application instance params
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
