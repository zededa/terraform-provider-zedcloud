// Code generated by go-swagger; DO NOT EDIT.

package asset_group_service

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

// NewAssetGroupServiceCreateAssetGroupParams creates a new AssetGroupServiceCreateAssetGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssetGroupServiceCreateAssetGroupParams() *AssetGroupServiceCreateAssetGroupParams {
	return &AssetGroupServiceCreateAssetGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssetGroupServiceCreateAssetGroupParamsWithTimeout creates a new AssetGroupServiceCreateAssetGroupParams object
// with the ability to set a timeout on a request.
func NewAssetGroupServiceCreateAssetGroupParamsWithTimeout(timeout time.Duration) *AssetGroupServiceCreateAssetGroupParams {
	return &AssetGroupServiceCreateAssetGroupParams{
		timeout: timeout,
	}
}

// NewAssetGroupServiceCreateAssetGroupParamsWithContext creates a new AssetGroupServiceCreateAssetGroupParams object
// with the ability to set a context for a request.
func NewAssetGroupServiceCreateAssetGroupParamsWithContext(ctx context.Context) *AssetGroupServiceCreateAssetGroupParams {
	return &AssetGroupServiceCreateAssetGroupParams{
		Context: ctx,
	}
}

// NewAssetGroupServiceCreateAssetGroupParamsWithHTTPClient creates a new AssetGroupServiceCreateAssetGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssetGroupServiceCreateAssetGroupParamsWithHTTPClient(client *http.Client) *AssetGroupServiceCreateAssetGroupParams {
	return &AssetGroupServiceCreateAssetGroupParams{
		HTTPClient: client,
	}
}

/*
AssetGroupServiceCreateAssetGroupParams contains all the parameters to send to the API endpoint

	for the asset group service create asset group operation.

	Typically these are written to a http.Request.
*/
type AssetGroupServiceCreateAssetGroupParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.AssetGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the asset group service create asset group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssetGroupServiceCreateAssetGroupParams) WithDefaults() *AssetGroupServiceCreateAssetGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the asset group service create asset group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssetGroupServiceCreateAssetGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the asset group service create asset group params
func (o *AssetGroupServiceCreateAssetGroupParams) WithTimeout(timeout time.Duration) *AssetGroupServiceCreateAssetGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the asset group service create asset group params
func (o *AssetGroupServiceCreateAssetGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the asset group service create asset group params
func (o *AssetGroupServiceCreateAssetGroupParams) WithContext(ctx context.Context) *AssetGroupServiceCreateAssetGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the asset group service create asset group params
func (o *AssetGroupServiceCreateAssetGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the asset group service create asset group params
func (o *AssetGroupServiceCreateAssetGroupParams) WithHTTPClient(client *http.Client) *AssetGroupServiceCreateAssetGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the asset group service create asset group params
func (o *AssetGroupServiceCreateAssetGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the asset group service create asset group params
func (o *AssetGroupServiceCreateAssetGroupParams) WithXRequestID(xRequestID *string) *AssetGroupServiceCreateAssetGroupParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the asset group service create asset group params
func (o *AssetGroupServiceCreateAssetGroupParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the asset group service create asset group params
func (o *AssetGroupServiceCreateAssetGroupParams) WithBody(body *models.AssetGroup) *AssetGroupServiceCreateAssetGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the asset group service create asset group params
func (o *AssetGroupServiceCreateAssetGroupParams) SetBody(body *models.AssetGroup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AssetGroupServiceCreateAssetGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
