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

// UpdateBaseOSParams creates a new EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func UpdateBaseOSParams() *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSParamsWithTimeout creates a new EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams object
// with the ability to set a timeout on a request.
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSParamsWithTimeout(timeout time.Duration) *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams{
		timeout: timeout,
	}
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSParamsWithContext creates a new EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams object
// with the ability to set a context for a request.
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSParamsWithContext(ctx context.Context) *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams{
		Context: ctx,
	}
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSParamsWithHTTPClient creates a new EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOSParamsWithHTTPClient(client *http.Client) *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams{
		HTTPClient: client,
	}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams contains all the parameters to send to the API endpoint

	for the edge node configuration update edge node base o s operation.

	Typically these are written to a http.Request.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams struct {

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

// WithDefaults hydrates default values in the edge node configuration update edge node base o s params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams) WithDefaults() *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge node configuration update edge node base o s params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge node configuration update edge node base o s params
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams) WithTimeout(timeout time.Duration) *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge node configuration update edge node base o s params
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge node configuration update edge node base o s params
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams) WithContext(ctx context.Context) *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge node configuration update edge node base o s params
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge node configuration update edge node base o s params
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams) WithHTTPClient(client *http.Client) *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge node configuration update edge node base o s params
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge node configuration update edge node base o s params
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams) WithXRequestID(xRequestID *string) *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge node configuration update edge node base o s params
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the edge node configuration update edge node base o s params
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams) WithID(id string) *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge node configuration update edge node base o s params
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOSParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
