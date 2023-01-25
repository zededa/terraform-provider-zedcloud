// Code generated by go-swagger; DO NOT EDIT.

package edge_network_instance_status

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

// NewEdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams creates a new EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams() *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams {
	return &EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParamsWithTimeout creates a new EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams object
// with the ability to set a timeout on a request.
func NewEdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParamsWithTimeout(timeout time.Duration) *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams {
	return &EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams{
		timeout: timeout,
	}
}

// NewEdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParamsWithContext creates a new EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams object
// with the ability to set a context for a request.
func NewEdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParamsWithContext(ctx context.Context) *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams {
	return &EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams{
		Context: ctx,
	}
}

// NewEdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParamsWithHTTPClient creates a new EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParamsWithHTTPClient(client *http.Client) *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams {
	return &EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams{
		HTTPClient: client,
	}
}

/*
EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams contains all the parameters to send to the API endpoint

	for the edge network instance status get edge network instance status by name operation.

	Typically these are written to a http.Request.
*/
type EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge network instance status get edge network instance status by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams) WithDefaults() *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge network instance status get edge network instance status by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge network instance status get edge network instance status by name params
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams) WithTimeout(timeout time.Duration) *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge network instance status get edge network instance status by name params
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge network instance status get edge network instance status by name params
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams) WithContext(ctx context.Context) *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge network instance status get edge network instance status by name params
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge network instance status get edge network instance status by name params
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams) WithHTTPClient(client *http.Client) *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge network instance status get edge network instance status by name params
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge network instance status get edge network instance status by name params
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams) WithXRequestID(xRequestID *string) *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge network instance status get edge network instance status by name params
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the edge network instance status get edge network instance status by name params
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams) WithName(name string) *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the edge network instance status get edge network instance status by name params
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNetworkInstanceStatusGetEdgeNetworkInstanceStatusByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
