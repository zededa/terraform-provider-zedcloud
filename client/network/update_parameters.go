package network

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

// UpdateParams creates a new EdgeNetworkConfigurationUpdateEdgeNetworkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func UpdateParams() *EdgeNetworkConfigurationUpdateEdgeNetworkParams {
	return &EdgeNetworkConfigurationUpdateEdgeNetworkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNetworkConfigurationUpdateEdgeNetworkParamsWithTimeout creates a new EdgeNetworkConfigurationUpdateEdgeNetworkParams object
// with the ability to set a timeout on a request.
func NewEdgeNetworkConfigurationUpdateEdgeNetworkParamsWithTimeout(timeout time.Duration) *EdgeNetworkConfigurationUpdateEdgeNetworkParams {
	return &EdgeNetworkConfigurationUpdateEdgeNetworkParams{
		timeout: timeout,
	}
}

// NewEdgeNetworkConfigurationUpdateEdgeNetworkParamsWithContext creates a new EdgeNetworkConfigurationUpdateEdgeNetworkParams object
// with the ability to set a context for a request.
func NewEdgeNetworkConfigurationUpdateEdgeNetworkParamsWithContext(ctx context.Context) *EdgeNetworkConfigurationUpdateEdgeNetworkParams {
	return &EdgeNetworkConfigurationUpdateEdgeNetworkParams{
		Context: ctx,
	}
}

// NewEdgeNetworkConfigurationUpdateEdgeNetworkParamsWithHTTPClient creates a new EdgeNetworkConfigurationUpdateEdgeNetworkParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNetworkConfigurationUpdateEdgeNetworkParamsWithHTTPClient(client *http.Client) *EdgeNetworkConfigurationUpdateEdgeNetworkParams {
	return &EdgeNetworkConfigurationUpdateEdgeNetworkParams{
		HTTPClient: client,
	}
}

/*
EdgeNetworkConfigurationUpdateEdgeNetworkParams contains all the parameters to send to the API endpoint

	for the edge network configuration update edge network operation.

	Typically these are written to a http.Request.
*/
type EdgeNetworkConfigurationUpdateEdgeNetworkParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.Network

	/* ID.

	   System defined universally unique Id of the network
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge network configuration update edge network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNetworkConfigurationUpdateEdgeNetworkParams) WithDefaults() *EdgeNetworkConfigurationUpdateEdgeNetworkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge network configuration update edge network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNetworkConfigurationUpdateEdgeNetworkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge network configuration update edge network params
func (o *EdgeNetworkConfigurationUpdateEdgeNetworkParams) WithTimeout(timeout time.Duration) *EdgeNetworkConfigurationUpdateEdgeNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge network configuration update edge network params
func (o *EdgeNetworkConfigurationUpdateEdgeNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge network configuration update edge network params
func (o *EdgeNetworkConfigurationUpdateEdgeNetworkParams) WithContext(ctx context.Context) *EdgeNetworkConfigurationUpdateEdgeNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge network configuration update edge network params
func (o *EdgeNetworkConfigurationUpdateEdgeNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge network configuration update edge network params
func (o *EdgeNetworkConfigurationUpdateEdgeNetworkParams) WithHTTPClient(client *http.Client) *EdgeNetworkConfigurationUpdateEdgeNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge network configuration update edge network params
func (o *EdgeNetworkConfigurationUpdateEdgeNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge network configuration update edge network params
func (o *EdgeNetworkConfigurationUpdateEdgeNetworkParams) WithXRequestID(xRequestID *string) *EdgeNetworkConfigurationUpdateEdgeNetworkParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge network configuration update edge network params
func (o *EdgeNetworkConfigurationUpdateEdgeNetworkParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the edge network configuration update edge network params
func (o *EdgeNetworkConfigurationUpdateEdgeNetworkParams) WithBody(body *models.Network) *EdgeNetworkConfigurationUpdateEdgeNetworkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edge network configuration update edge network params
func (o *EdgeNetworkConfigurationUpdateEdgeNetworkParams) SetBody(body *models.Network) {
	o.Body = body
}

// WithID adds the id to the edge network configuration update edge network params
func (o *EdgeNetworkConfigurationUpdateEdgeNetworkParams) WithID(id string) *EdgeNetworkConfigurationUpdateEdgeNetworkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge network configuration update edge network params
func (o *EdgeNetworkConfigurationUpdateEdgeNetworkParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNetworkConfigurationUpdateEdgeNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
