package datastore

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

// CreateParams creates a new DatastoreConfigurationCreateDatastoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func CreateParams() *DatastoreConfigurationCreateDatastoreParams {
	return &DatastoreConfigurationCreateDatastoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDatastoreConfigurationCreateDatastoreParamsWithTimeout creates a new DatastoreConfigurationCreateDatastoreParams object
// with the ability to set a timeout on a request.
func NewDatastoreConfigurationCreateDatastoreParamsWithTimeout(timeout time.Duration) *DatastoreConfigurationCreateDatastoreParams {
	return &DatastoreConfigurationCreateDatastoreParams{
		timeout: timeout,
	}
}

// NewDatastoreConfigurationCreateDatastoreParamsWithContext creates a new DatastoreConfigurationCreateDatastoreParams object
// with the ability to set a context for a request.
func NewDatastoreConfigurationCreateDatastoreParamsWithContext(ctx context.Context) *DatastoreConfigurationCreateDatastoreParams {
	return &DatastoreConfigurationCreateDatastoreParams{
		Context: ctx,
	}
}

// NewDatastoreConfigurationCreateDatastoreParamsWithHTTPClient creates a new DatastoreConfigurationCreateDatastoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewDatastoreConfigurationCreateDatastoreParamsWithHTTPClient(client *http.Client) *DatastoreConfigurationCreateDatastoreParams {
	return &DatastoreConfigurationCreateDatastoreParams{
		HTTPClient: client,
	}
}

/*
DatastoreConfigurationCreateDatastoreParams contains all the parameters to send to the API endpoint

	for the datastore configuration create datastore operation.

	Typically these are written to a http.Request.
*/
type DatastoreConfigurationCreateDatastoreParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.Datastore

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the datastore configuration create datastore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DatastoreConfigurationCreateDatastoreParams) WithDefaults() *DatastoreConfigurationCreateDatastoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the datastore configuration create datastore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DatastoreConfigurationCreateDatastoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the datastore configuration create datastore params
func (o *DatastoreConfigurationCreateDatastoreParams) WithTimeout(timeout time.Duration) *DatastoreConfigurationCreateDatastoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the datastore configuration create datastore params
func (o *DatastoreConfigurationCreateDatastoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the datastore configuration create datastore params
func (o *DatastoreConfigurationCreateDatastoreParams) WithContext(ctx context.Context) *DatastoreConfigurationCreateDatastoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the datastore configuration create datastore params
func (o *DatastoreConfigurationCreateDatastoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the datastore configuration create datastore params
func (o *DatastoreConfigurationCreateDatastoreParams) WithHTTPClient(client *http.Client) *DatastoreConfigurationCreateDatastoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the datastore configuration create datastore params
func (o *DatastoreConfigurationCreateDatastoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the datastore configuration create datastore params
func (o *DatastoreConfigurationCreateDatastoreParams) WithXRequestID(xRequestID *string) *DatastoreConfigurationCreateDatastoreParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the datastore configuration create datastore params
func (o *DatastoreConfigurationCreateDatastoreParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the datastore configuration create datastore params
func (o *DatastoreConfigurationCreateDatastoreParams) WithBody(body *models.Datastore) *DatastoreConfigurationCreateDatastoreParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the datastore configuration create datastore params
func (o *DatastoreConfigurationCreateDatastoreParams) SetBody(body *models.Datastore) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DatastoreConfigurationCreateDatastoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
