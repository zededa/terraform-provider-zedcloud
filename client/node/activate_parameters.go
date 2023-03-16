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

// ActivationParams creates a new EdgeNodeConfigurationActivateEdgeNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func ActivationParams() *EdgeNodeConfigurationActivateEdgeNodeParams {
	return &EdgeNodeConfigurationActivateEdgeNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNodeConfigurationActivateEdgeNodeParamsWithTimeout creates a new EdgeNodeConfigurationActivateEdgeNodeParams object
// with the ability to set a timeout on a request.
func NewEdgeNodeConfigurationActivateEdgeNodeParamsWithTimeout(timeout time.Duration) *EdgeNodeConfigurationActivateEdgeNodeParams {
	return &EdgeNodeConfigurationActivateEdgeNodeParams{
		timeout: timeout,
	}
}

// NewEdgeNodeConfigurationActivateEdgeNodeParamsWithContext creates a new EdgeNodeConfigurationActivateEdgeNodeParams object
// with the ability to set a context for a request.
func NewEdgeNodeConfigurationActivateEdgeNodeParamsWithContext(ctx context.Context) *EdgeNodeConfigurationActivateEdgeNodeParams {
	return &EdgeNodeConfigurationActivateEdgeNodeParams{
		Context: ctx,
	}
}

// NewEdgeNodeConfigurationActivateEdgeNodeParamsWithHTTPClient creates a new EdgeNodeConfigurationActivateEdgeNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNodeConfigurationActivateEdgeNodeParamsWithHTTPClient(client *http.Client) *EdgeNodeConfigurationActivateEdgeNodeParams {
	return &EdgeNodeConfigurationActivateEdgeNodeParams{
		HTTPClient: client,
	}
}

/*
EdgeNodeConfigurationActivateEdgeNodeParams contains all the parameters to send to the API endpoint

	for the edge node configuration activate edge node operation.

	Typically these are written to a http.Request.
*/
type EdgeNodeConfigurationActivateEdgeNodeParams struct {

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

// WithDefaults hydrates default values in the edge node configuration activate edge node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationActivateEdgeNodeParams) WithDefaults() *EdgeNodeConfigurationActivateEdgeNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge node configuration activate edge node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationActivateEdgeNodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge node configuration activate edge node params
func (o *EdgeNodeConfigurationActivateEdgeNodeParams) WithTimeout(timeout time.Duration) *EdgeNodeConfigurationActivateEdgeNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge node configuration activate edge node params
func (o *EdgeNodeConfigurationActivateEdgeNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge node configuration activate edge node params
func (o *EdgeNodeConfigurationActivateEdgeNodeParams) WithContext(ctx context.Context) *EdgeNodeConfigurationActivateEdgeNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge node configuration activate edge node params
func (o *EdgeNodeConfigurationActivateEdgeNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge node configuration activate edge node params
func (o *EdgeNodeConfigurationActivateEdgeNodeParams) WithHTTPClient(client *http.Client) *EdgeNodeConfigurationActivateEdgeNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge node configuration activate edge node params
func (o *EdgeNodeConfigurationActivateEdgeNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge node configuration activate edge node params
func (o *EdgeNodeConfigurationActivateEdgeNodeParams) WithXRequestID(xRequestID *string) *EdgeNodeConfigurationActivateEdgeNodeParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge node configuration activate edge node params
func (o *EdgeNodeConfigurationActivateEdgeNodeParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the edge node configuration activate edge node params
func (o *EdgeNodeConfigurationActivateEdgeNodeParams) WithID(id string) *EdgeNodeConfigurationActivateEdgeNodeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge node configuration activate edge node params
func (o *EdgeNodeConfigurationActivateEdgeNodeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNodeConfigurationActivateEdgeNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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