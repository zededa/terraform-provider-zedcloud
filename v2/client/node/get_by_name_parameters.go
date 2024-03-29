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

// GetByNameParams creates a new EdgeNodeConfigurationGetEdgeNodeByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func GetByNameParams() *EdgeNodeConfigurationGetEdgeNodeByNameParams {
	return &EdgeNodeConfigurationGetEdgeNodeByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNodeConfigurationGetEdgeNodeByNameParamsWithTimeout creates a new EdgeNodeConfigurationGetEdgeNodeByNameParams object
// with the ability to set a timeout on a request.
func NewEdgeNodeConfigurationGetEdgeNodeByNameParamsWithTimeout(timeout time.Duration) *EdgeNodeConfigurationGetEdgeNodeByNameParams {
	return &EdgeNodeConfigurationGetEdgeNodeByNameParams{
		timeout: timeout,
	}
}

// NewEdgeNodeConfigurationGetEdgeNodeByNameParamsWithContext creates a new EdgeNodeConfigurationGetEdgeNodeByNameParams object
// with the ability to set a context for a request.
func NewEdgeNodeConfigurationGetEdgeNodeByNameParamsWithContext(ctx context.Context) *EdgeNodeConfigurationGetEdgeNodeByNameParams {
	return &EdgeNodeConfigurationGetEdgeNodeByNameParams{
		Context: ctx,
	}
}

// NewEdgeNodeConfigurationGetEdgeNodeByNameParamsWithHTTPClient creates a new EdgeNodeConfigurationGetEdgeNodeByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNodeConfigurationGetEdgeNodeByNameParamsWithHTTPClient(client *http.Client) *EdgeNodeConfigurationGetEdgeNodeByNameParams {
	return &EdgeNodeConfigurationGetEdgeNodeByNameParams{
		HTTPClient: client,
	}
}

/*
EdgeNodeConfigurationGetEdgeNodeByNameParams contains all the parameters to send to the API endpoint

	for the edge node configuration get edge node by name operation.

	Typically these are written to a http.Request.
*/
type EdgeNodeConfigurationGetEdgeNodeByNameParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* Name.

	   user defined device name for a device
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge node configuration get edge node by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationGetEdgeNodeByNameParams) WithDefaults() *EdgeNodeConfigurationGetEdgeNodeByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge node configuration get edge node by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationGetEdgeNodeByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge node configuration get edge node by name params
func (o *EdgeNodeConfigurationGetEdgeNodeByNameParams) WithTimeout(timeout time.Duration) *EdgeNodeConfigurationGetEdgeNodeByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge node configuration get edge node by name params
func (o *EdgeNodeConfigurationGetEdgeNodeByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge node configuration get edge node by name params
func (o *EdgeNodeConfigurationGetEdgeNodeByNameParams) WithContext(ctx context.Context) *EdgeNodeConfigurationGetEdgeNodeByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge node configuration get edge node by name params
func (o *EdgeNodeConfigurationGetEdgeNodeByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge node configuration get edge node by name params
func (o *EdgeNodeConfigurationGetEdgeNodeByNameParams) WithHTTPClient(client *http.Client) *EdgeNodeConfigurationGetEdgeNodeByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge node configuration get edge node by name params
func (o *EdgeNodeConfigurationGetEdgeNodeByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge node configuration get edge node by name params
func (o *EdgeNodeConfigurationGetEdgeNodeByNameParams) WithXRequestID(xRequestID *string) *EdgeNodeConfigurationGetEdgeNodeByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge node configuration get edge node by name params
func (o *EdgeNodeConfigurationGetEdgeNodeByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the edge node configuration get edge node by name params
func (o *EdgeNodeConfigurationGetEdgeNodeByNameParams) WithName(name string) *EdgeNodeConfigurationGetEdgeNodeByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the edge node configuration get edge node by name params
func (o *EdgeNodeConfigurationGetEdgeNodeByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNodeConfigurationGetEdgeNodeByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
