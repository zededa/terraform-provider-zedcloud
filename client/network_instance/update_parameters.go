package network_instance

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
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// NewEdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams creates a new EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams() *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams {
	return &EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParamsWithTimeout creates a new EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams object
// with the ability to set a timeout on a request.
func NewEdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParamsWithTimeout(timeout time.Duration) *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams {
	return &EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams{
		timeout: timeout,
	}
}

// NewEdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParamsWithContext creates a new EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams object
// with the ability to set a context for a request.
func NewEdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParamsWithContext(ctx context.Context) *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams {
	return &EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams{
		Context: ctx,
	}
}

// NewEdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParamsWithHTTPClient creates a new EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParamsWithHTTPClient(client *http.Client) *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams {
	return &EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams{
		HTTPClient: client,
	}
}

/*
EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams contains all the parameters to send to the API endpoint

	for the edge network instance configuration update edge network instance operation.

	Typically these are written to a http.Request.
*/
type EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.NetworkInstance

	/* ID.

	   System defined universally unique Id of the network instance
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge network instance configuration update edge network instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams) WithDefaults() *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge network instance configuration update edge network instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge network instance configuration update edge network instance params
func (o *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams) WithTimeout(timeout time.Duration) *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge network instance configuration update edge network instance params
func (o *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge network instance configuration update edge network instance params
func (o *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams) WithContext(ctx context.Context) *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge network instance configuration update edge network instance params
func (o *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge network instance configuration update edge network instance params
func (o *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams) WithHTTPClient(client *http.Client) *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge network instance configuration update edge network instance params
func (o *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge network instance configuration update edge network instance params
func (o *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams) WithXRequestID(xRequestID *string) *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge network instance configuration update edge network instance params
func (o *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the edge network instance configuration update edge network instance params
func (o *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams) WithBody(body *models.NetworkInstance) *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edge network instance configuration update edge network instance params
func (o *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams) SetBody(body *models.NetworkInstance) {
	o.Body = body
}

// WithID adds the id to the edge network instance configuration update edge network instance params
func (o *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams) WithID(id string) *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge network instance configuration update edge network instance params
func (o *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNetworkInstanceConfigurationUpdateEdgeNetworkInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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
