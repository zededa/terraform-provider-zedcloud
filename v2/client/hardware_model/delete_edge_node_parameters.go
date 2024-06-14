// Code generated by go-swagger; DO NOT EDIT.

package hardware_model

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

// NewHardwareModelDeleteEdgeNodeParams creates a new HardwareModelDeleteEdgeNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHardwareModelDeleteEdgeNodeParams() *HardwareModelDeleteEdgeNodeParams {
	return &HardwareModelDeleteEdgeNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHardwareModelDeleteEdgeNodeParamsWithTimeout creates a new HardwareModelDeleteEdgeNodeParams object
// with the ability to set a timeout on a request.
func NewHardwareModelDeleteEdgeNodeParamsWithTimeout(timeout time.Duration) *HardwareModelDeleteEdgeNodeParams {
	return &HardwareModelDeleteEdgeNodeParams{
		timeout: timeout,
	}
}

// NewHardwareModelDeleteEdgeNodeParamsWithContext creates a new HardwareModelDeleteEdgeNodeParams object
// with the ability to set a context for a request.
func NewHardwareModelDeleteEdgeNodeParamsWithContext(ctx context.Context) *HardwareModelDeleteEdgeNodeParams {
	return &HardwareModelDeleteEdgeNodeParams{
		Context: ctx,
	}
}

// NewHardwareModelDeleteEdgeNodeParamsWithHTTPClient creates a new HardwareModelDeleteEdgeNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewHardwareModelDeleteEdgeNodeParamsWithHTTPClient(client *http.Client) *HardwareModelDeleteEdgeNodeParams {
	return &HardwareModelDeleteEdgeNodeParams{
		HTTPClient: client,
	}
}

/*
HardwareModelDeleteEdgeNodeParams contains all the parameters to send to the API endpoint

	for the hardware model delete edge node operation.

	Typically these are written to a http.Request.
*/
type HardwareModelDeleteEdgeNodeParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   System defined universally unique Id of the pcr template
	*/
	ID string

	/* ModelID.

	   Device model identifier
	*/
	ModelID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the hardware model delete edge node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HardwareModelDeleteEdgeNodeParams) WithDefaults() *HardwareModelDeleteEdgeNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hardware model delete edge node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HardwareModelDeleteEdgeNodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hardware model delete edge node params
func (o *HardwareModelDeleteEdgeNodeParams) WithTimeout(timeout time.Duration) *HardwareModelDeleteEdgeNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hardware model delete edge node params
func (o *HardwareModelDeleteEdgeNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hardware model delete edge node params
func (o *HardwareModelDeleteEdgeNodeParams) WithContext(ctx context.Context) *HardwareModelDeleteEdgeNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hardware model delete edge node params
func (o *HardwareModelDeleteEdgeNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hardware model delete edge node params
func (o *HardwareModelDeleteEdgeNodeParams) WithHTTPClient(client *http.Client) *HardwareModelDeleteEdgeNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hardware model delete edge node params
func (o *HardwareModelDeleteEdgeNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the hardware model delete edge node params
func (o *HardwareModelDeleteEdgeNodeParams) WithXRequestID(xRequestID *string) *HardwareModelDeleteEdgeNodeParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the hardware model delete edge node params
func (o *HardwareModelDeleteEdgeNodeParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the hardware model delete edge node params
func (o *HardwareModelDeleteEdgeNodeParams) WithID(id string) *HardwareModelDeleteEdgeNodeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the hardware model delete edge node params
func (o *HardwareModelDeleteEdgeNodeParams) SetID(id string) {
	o.ID = id
}

// WithModelID adds the modelID to the hardware model delete edge node params
func (o *HardwareModelDeleteEdgeNodeParams) WithModelID(modelID *string) *HardwareModelDeleteEdgeNodeParams {
	o.SetModelID(modelID)
	return o
}

// SetModelID adds the modelId to the hardware model delete edge node params
func (o *HardwareModelDeleteEdgeNodeParams) SetModelID(modelID *string) {
	o.ModelID = modelID
}

// WriteToRequest writes these params to a swagger request
func (o *HardwareModelDeleteEdgeNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ModelID != nil {

		// query param modelId
		var qrModelID string

		if o.ModelID != nil {
			qrModelID = *o.ModelID
		}
		qModelID := qrModelID
		if qModelID != "" {

			if err := r.SetQueryParam("modelId", qModelID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}