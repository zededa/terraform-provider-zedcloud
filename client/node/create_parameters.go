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

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// CreateParams creates a new EdgeNodeConfigurationCreateEdgeNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func CreateParams() *EdgeNodeConfigurationCreateEdgeNodeParams {
	return &EdgeNodeConfigurationCreateEdgeNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNodeConfigurationCreateEdgeNodeParamsWithTimeout creates a new EdgeNodeConfigurationCreateEdgeNodeParams object
// with the ability to set a timeout on a request.
func NewEdgeNodeConfigurationCreateEdgeNodeParamsWithTimeout(timeout time.Duration) *EdgeNodeConfigurationCreateEdgeNodeParams {
	return &EdgeNodeConfigurationCreateEdgeNodeParams{
		timeout: timeout,
	}
}

// NewEdgeNodeConfigurationCreateEdgeNodeParamsWithContext creates a new EdgeNodeConfigurationCreateEdgeNodeParams object
// with the ability to set a context for a request.
func NewEdgeNodeConfigurationCreateEdgeNodeParamsWithContext(ctx context.Context) *EdgeNodeConfigurationCreateEdgeNodeParams {
	return &EdgeNodeConfigurationCreateEdgeNodeParams{
		Context: ctx,
	}
}

// NewEdgeNodeConfigurationCreateEdgeNodeParamsWithHTTPClient creates a new EdgeNodeConfigurationCreateEdgeNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNodeConfigurationCreateEdgeNodeParamsWithHTTPClient(client *http.Client) *EdgeNodeConfigurationCreateEdgeNodeParams {
	return &EdgeNodeConfigurationCreateEdgeNodeParams{
		HTTPClient: client,
	}
}

/*
EdgeNodeConfigurationCreateEdgeNodeParams contains all the parameters to send to the API endpoint

	for the edge node configuration create edge node operation.

	Typically these are written to a http.Request.
*/
type EdgeNodeConfigurationCreateEdgeNodeParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.Node

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge node configuration create edge node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationCreateEdgeNodeParams) WithDefaults() *EdgeNodeConfigurationCreateEdgeNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge node configuration create edge node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationCreateEdgeNodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge node configuration create edge node params
func (o *EdgeNodeConfigurationCreateEdgeNodeParams) WithTimeout(timeout time.Duration) *EdgeNodeConfigurationCreateEdgeNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge node configuration create edge node params
func (o *EdgeNodeConfigurationCreateEdgeNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge node configuration create edge node params
func (o *EdgeNodeConfigurationCreateEdgeNodeParams) WithContext(ctx context.Context) *EdgeNodeConfigurationCreateEdgeNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge node configuration create edge node params
func (o *EdgeNodeConfigurationCreateEdgeNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge node configuration create edge node params
func (o *EdgeNodeConfigurationCreateEdgeNodeParams) WithHTTPClient(client *http.Client) *EdgeNodeConfigurationCreateEdgeNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge node configuration create edge node params
func (o *EdgeNodeConfigurationCreateEdgeNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge node configuration create edge node params
func (o *EdgeNodeConfigurationCreateEdgeNodeParams) WithXRequestID(xRequestID *string) *EdgeNodeConfigurationCreateEdgeNodeParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge node configuration create edge node params
func (o *EdgeNodeConfigurationCreateEdgeNodeParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the edge node configuration create edge node params
func (o *EdgeNodeConfigurationCreateEdgeNodeParams) WithBody(body *models.Node) *EdgeNodeConfigurationCreateEdgeNodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edge node configuration create edge node params
func (o *EdgeNodeConfigurationCreateEdgeNodeParams) SetBody(body *models.Node) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNodeConfigurationCreateEdgeNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
