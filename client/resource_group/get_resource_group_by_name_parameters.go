// Code generated by go-swagger; DO NOT EDIT.

package resource_group

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

// NewGetResourceGroupByNameParams creates a new GetResourceGroupByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetResourceGroupByNameParams() *GetResourceGroupByNameParams {
	return &GetResourceGroupByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourceGroupByNameParamsWithTimeout creates a new GetResourceGroupByNameParams object
// with the ability to set a timeout on a request.
func NewGetResourceGroupByNameParamsWithTimeout(timeout time.Duration) *GetResourceGroupByNameParams {
	return &GetResourceGroupByNameParams{
		timeout: timeout,
	}
}

// NewGetResourceGroupByNameParamsWithContext creates a new GetResourceGroupByNameParams object
// with the ability to set a context for a request.
func NewGetResourceGroupByNameParamsWithContext(ctx context.Context) *GetResourceGroupByNameParams {
	return &GetResourceGroupByNameParams{
		Context: ctx,
	}
}

// NewGetResourceGroupByNameParamsWithHTTPClient creates a new GetResourceGroupByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetResourceGroupByNameParamsWithHTTPClient(client *http.Client) *GetResourceGroupByNameParams {
	return &GetResourceGroupByNameParams{
		HTTPClient: client,
	}
}

/*
GetResourceGroupByNameParams contains all the parameters to send to the API endpoint

	for the get resource group by name operation.

	Typically these are written to a http.Request.
*/
type GetResourceGroupByNameParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* Name.

	   User defined name of the resource group, unique across the enterprise. Once resource group is created, name can’t be changed.
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get resource group by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceGroupByNameParams) WithDefaults() *GetResourceGroupByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get resource group by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceGroupByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get resource group by name params
func (o *GetResourceGroupByNameParams) WithTimeout(timeout time.Duration) *GetResourceGroupByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource group by name params
func (o *GetResourceGroupByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource group by name params
func (o *GetResourceGroupByNameParams) WithContext(ctx context.Context) *GetResourceGroupByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource group by name params
func (o *GetResourceGroupByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource group by name params
func (o *GetResourceGroupByNameParams) WithHTTPClient(client *http.Client) *GetResourceGroupByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource group by name params
func (o *GetResourceGroupByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get resource group by name params
func (o *GetResourceGroupByNameParams) WithXRequestID(xRequestID *string) *GetResourceGroupByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get resource group by name params
func (o *GetResourceGroupByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the get resource group by name params
func (o *GetResourceGroupByNameParams) WithName(name string) *GetResourceGroupByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get resource group by name params
func (o *GetResourceGroupByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourceGroupByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
