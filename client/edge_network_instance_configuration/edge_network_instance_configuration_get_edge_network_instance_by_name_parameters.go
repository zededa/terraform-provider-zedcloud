// Code generated by go-swagger; DO NOT EDIT.

package edge_network_instance_configuration

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

// GetByNameParams creates a new EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func GetByNameParams() *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams {
	return &EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParamsWithTimeout creates a new EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams object
// with the ability to set a timeout on a request.
func NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParamsWithTimeout(timeout time.Duration) *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams {
	return &EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams{
		timeout: timeout,
	}
}

// NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParamsWithContext creates a new EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams object
// with the ability to set a context for a request.
func NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParamsWithContext(ctx context.Context) *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams {
	return &EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams{
		Context: ctx,
	}
}

// NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParamsWithHTTPClient creates a new EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParamsWithHTTPClient(client *http.Client) *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams {
	return &EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams{
		HTTPClient: client,
	}
}

/*
EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams contains all the parameters to send to the API endpoint

	for the edge network instance configuration get edge network instance by name operation.

	Typically these are written to a http.Request.
*/
type EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* Name.

	   User defined name of the network instance, unique across the enterprise. Once object is created, name can’t be changed
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge network instance configuration get edge network instance by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams) WithDefaults() *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge network instance configuration get edge network instance by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge network instance configuration get edge network instance by name params
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams) WithTimeout(timeout time.Duration) *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge network instance configuration get edge network instance by name params
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge network instance configuration get edge network instance by name params
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams) WithContext(ctx context.Context) *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge network instance configuration get edge network instance by name params
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge network instance configuration get edge network instance by name params
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams) WithHTTPClient(client *http.Client) *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge network instance configuration get edge network instance by name params
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge network instance configuration get edge network instance by name params
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams) WithXRequestID(xRequestID *string) *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge network instance configuration get edge network instance by name params
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the edge network instance configuration get edge network instance by name params
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams) WithName(name string) *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the edge network instance configuration get edge network instance by name params
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNetworkInstanceConfigurationGetEdgeNetworkInstanceByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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