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

// PublishBaseOSParams creates a new EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func PublishBaseOSParams() *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2ParamsWithTimeout creates a new EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params object
// with the ability to set a timeout on a request.
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2ParamsWithTimeout(timeout time.Duration) *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params{
		timeout: timeout,
	}
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2ParamsWithContext creates a new EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params object
// with the ability to set a context for a request.
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2ParamsWithContext(ctx context.Context) *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params{
		Context: ctx,
	}
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2ParamsWithHTTPClient creates a new EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2ParamsWithHTTPClient(client *http.Client) *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params{
		HTTPClient: client,
	}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params contains all the parameters to send to the API endpoint

	for the edge node configuration update edge node base o s2 operation.

	Typically these are written to a http.Request.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params struct {

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

// WithDefaults hydrates default values in the edge node configuration update edge node base o s2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params) WithDefaults() *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge node configuration update edge node base o s2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge node configuration update edge node base o s2 params
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params) WithTimeout(timeout time.Duration) *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge node configuration update edge node base o s2 params
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge node configuration update edge node base o s2 params
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params) WithContext(ctx context.Context) *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge node configuration update edge node base o s2 params
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge node configuration update edge node base o s2 params
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params) WithHTTPClient(client *http.Client) *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge node configuration update edge node base o s2 params
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge node configuration update edge node base o s2 params
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params) WithXRequestID(xRequestID *string) *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge node configuration update edge node base o s2 params
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the edge node configuration update edge node base o s2 params
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params) WithID(id string) *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge node configuration update edge node base o s2 params
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
