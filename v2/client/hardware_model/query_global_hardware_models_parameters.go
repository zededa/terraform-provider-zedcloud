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
	"github.com/go-openapi/swag"
)

// NewHardwareModelQueryGlobalHardwareModelsParams creates a new HardwareModelQueryGlobalHardwareModelsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHardwareModelQueryGlobalHardwareModelsParams() *HardwareModelQueryGlobalHardwareModelsParams {
	return &HardwareModelQueryGlobalHardwareModelsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHardwareModelQueryGlobalHardwareModelsParamsWithTimeout creates a new HardwareModelQueryGlobalHardwareModelsParams object
// with the ability to set a timeout on a request.
func NewHardwareModelQueryGlobalHardwareModelsParamsWithTimeout(timeout time.Duration) *HardwareModelQueryGlobalHardwareModelsParams {
	return &HardwareModelQueryGlobalHardwareModelsParams{
		timeout: timeout,
	}
}

// NewHardwareModelQueryGlobalHardwareModelsParamsWithContext creates a new HardwareModelQueryGlobalHardwareModelsParams object
// with the ability to set a context for a request.
func NewHardwareModelQueryGlobalHardwareModelsParamsWithContext(ctx context.Context) *HardwareModelQueryGlobalHardwareModelsParams {
	return &HardwareModelQueryGlobalHardwareModelsParams{
		Context: ctx,
	}
}

// NewHardwareModelQueryGlobalHardwareModelsParamsWithHTTPClient creates a new HardwareModelQueryGlobalHardwareModelsParams object
// with the ability to set a custom HTTPClient for a request.
func NewHardwareModelQueryGlobalHardwareModelsParamsWithHTTPClient(client *http.Client) *HardwareModelQueryGlobalHardwareModelsParams {
	return &HardwareModelQueryGlobalHardwareModelsParams{
		HTTPClient: client,
	}
}

/*
HardwareModelQueryGlobalHardwareModelsParams contains all the parameters to send to the API endpoint

	for the hardware model query global hardware models operation.

	Typically these are written to a http.Request.
*/
type HardwareModelQueryGlobalHardwareModelsParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* BrandID.

	   System defined universally unique Id of the brand.
	*/
	BrandID []string

	/* NamePattern.

	   Model name pattern to be matched.
	*/
	NamePattern *string

	/* NextOrderBy.

	   OrderBy helps in sorting the list response
	*/
	NextOrderBy []string

	/* NextPageNum.

	   Page Number

	   Format: int64
	*/
	NextPageNum *int64

	/* NextPageSize.

	   Defines the page size

	   Format: int64
	*/
	NextPageSize *int64

	/* NextPageToken.

	   Page Token
	*/
	NextPageToken *string

	/* NextTotalPages.

	   Total number of pages to be fetched.

	   Format: int64
	*/
	NextTotalPages *int64

	/* OriginType.

	     origin of object

	 - ORIGIN_UNSPECIFIED: default options, which says no Operation/Invalid Operation
	 - ORIGIN_IMPORTED: Object imported from global enterprise.
	 - ORIGIN_LOCAL: Objectl created locally.
	 - ORIGIN_GLOBAL: Object created in global store,
	to use this type user should have root previlage.

	     Default: "ORIGIN_UNSPECIFIED"
	*/
	OriginType *string

	/* Summary.

	   Only summary of the records required
	*/
	Summary *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the hardware model query global hardware models params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HardwareModelQueryGlobalHardwareModelsParams) WithDefaults() *HardwareModelQueryGlobalHardwareModelsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hardware model query global hardware models params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HardwareModelQueryGlobalHardwareModelsParams) SetDefaults() {
	var (
		originTypeDefault = string("ORIGIN_UNSPECIFIED")
	)

	val := HardwareModelQueryGlobalHardwareModelsParams{
		OriginType: &originTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) WithTimeout(timeout time.Duration) *HardwareModelQueryGlobalHardwareModelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) WithContext(ctx context.Context) *HardwareModelQueryGlobalHardwareModelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) WithHTTPClient(client *http.Client) *HardwareModelQueryGlobalHardwareModelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) WithXRequestID(xRequestID *string) *HardwareModelQueryGlobalHardwareModelsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBrandID adds the brandID to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) WithBrandID(brandID []string) *HardwareModelQueryGlobalHardwareModelsParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) SetBrandID(brandID []string) {
	o.BrandID = brandID
}

// WithNamePattern adds the namePattern to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) WithNamePattern(namePattern *string) *HardwareModelQueryGlobalHardwareModelsParams {
	o.SetNamePattern(namePattern)
	return o
}

// SetNamePattern adds the namePattern to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) SetNamePattern(namePattern *string) {
	o.NamePattern = namePattern
}

// WithNextOrderBy adds the nextOrderBy to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) WithNextOrderBy(nextOrderBy []string) *HardwareModelQueryGlobalHardwareModelsParams {
	o.SetNextOrderBy(nextOrderBy)
	return o
}

// SetNextOrderBy adds the nextOrderBy to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) SetNextOrderBy(nextOrderBy []string) {
	o.NextOrderBy = nextOrderBy
}

// WithNextPageNum adds the nextPageNum to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) WithNextPageNum(nextPageNum *int64) *HardwareModelQueryGlobalHardwareModelsParams {
	o.SetNextPageNum(nextPageNum)
	return o
}

// SetNextPageNum adds the nextPageNum to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) SetNextPageNum(nextPageNum *int64) {
	o.NextPageNum = nextPageNum
}

// WithNextPageSize adds the nextPageSize to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) WithNextPageSize(nextPageSize *int64) *HardwareModelQueryGlobalHardwareModelsParams {
	o.SetNextPageSize(nextPageSize)
	return o
}

// SetNextPageSize adds the nextPageSize to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) SetNextPageSize(nextPageSize *int64) {
	o.NextPageSize = nextPageSize
}

// WithNextPageToken adds the nextPageToken to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) WithNextPageToken(nextPageToken *string) *HardwareModelQueryGlobalHardwareModelsParams {
	o.SetNextPageToken(nextPageToken)
	return o
}

// SetNextPageToken adds the nextPageToken to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) SetNextPageToken(nextPageToken *string) {
	o.NextPageToken = nextPageToken
}

// WithNextTotalPages adds the nextTotalPages to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) WithNextTotalPages(nextTotalPages *int64) *HardwareModelQueryGlobalHardwareModelsParams {
	o.SetNextTotalPages(nextTotalPages)
	return o
}

// SetNextTotalPages adds the nextTotalPages to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) SetNextTotalPages(nextTotalPages *int64) {
	o.NextTotalPages = nextTotalPages
}

// WithOriginType adds the originType to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) WithOriginType(originType *string) *HardwareModelQueryGlobalHardwareModelsParams {
	o.SetOriginType(originType)
	return o
}

// SetOriginType adds the originType to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) SetOriginType(originType *string) {
	o.OriginType = originType
}

// WithSummary adds the summary to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) WithSummary(summary *bool) *HardwareModelQueryGlobalHardwareModelsParams {
	o.SetSummary(summary)
	return o
}

// SetSummary adds the summary to the hardware model query global hardware models params
func (o *HardwareModelQueryGlobalHardwareModelsParams) SetSummary(summary *bool) {
	o.Summary = summary
}

// WriteToRequest writes these params to a swagger request
func (o *HardwareModelQueryGlobalHardwareModelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.BrandID != nil {

		// binding items for brandId
		joinedBrandID := o.bindParamBrandID(reg)

		// query array param brandId
		if err := r.SetQueryParam("brandId", joinedBrandID...); err != nil {
			return err
		}
	}

	if o.NamePattern != nil {

		// query param namePattern
		var qrNamePattern string

		if o.NamePattern != nil {
			qrNamePattern = *o.NamePattern
		}
		qNamePattern := qrNamePattern
		if qNamePattern != "" {

			if err := r.SetQueryParam("namePattern", qNamePattern); err != nil {
				return err
			}
		}
	}

	if o.NextOrderBy != nil {

		// binding items for next.orderBy
		joinedNextOrderBy := o.bindParamNextOrderBy(reg)

		// query array param next.orderBy
		if err := r.SetQueryParam("next.orderBy", joinedNextOrderBy...); err != nil {
			return err
		}
	}

	if o.NextPageNum != nil {

		// query param next.pageNum
		var qrNextPageNum int64

		if o.NextPageNum != nil {
			qrNextPageNum = *o.NextPageNum
		}
		qNextPageNum := swag.FormatInt64(qrNextPageNum)
		if qNextPageNum != "" {

			if err := r.SetQueryParam("next.pageNum", qNextPageNum); err != nil {
				return err
			}
		}
	}

	if o.NextPageSize != nil {

		// query param next.pageSize
		var qrNextPageSize int64

		if o.NextPageSize != nil {
			qrNextPageSize = *o.NextPageSize
		}
		qNextPageSize := swag.FormatInt64(qrNextPageSize)
		if qNextPageSize != "" {

			if err := r.SetQueryParam("next.pageSize", qNextPageSize); err != nil {
				return err
			}
		}
	}

	if o.NextPageToken != nil {

		// query param next.pageToken
		var qrNextPageToken string

		if o.NextPageToken != nil {
			qrNextPageToken = *o.NextPageToken
		}
		qNextPageToken := qrNextPageToken
		if qNextPageToken != "" {

			if err := r.SetQueryParam("next.pageToken", qNextPageToken); err != nil {
				return err
			}
		}
	}

	if o.NextTotalPages != nil {

		// query param next.totalPages
		var qrNextTotalPages int64

		if o.NextTotalPages != nil {
			qrNextTotalPages = *o.NextTotalPages
		}
		qNextTotalPages := swag.FormatInt64(qrNextTotalPages)
		if qNextTotalPages != "" {

			if err := r.SetQueryParam("next.totalPages", qNextTotalPages); err != nil {
				return err
			}
		}
	}

	if o.OriginType != nil {

		// query param originType
		var qrOriginType string

		if o.OriginType != nil {
			qrOriginType = *o.OriginType
		}
		qOriginType := qrOriginType
		if qOriginType != "" {

			if err := r.SetQueryParam("originType", qOriginType); err != nil {
				return err
			}
		}
	}

	if o.Summary != nil {

		// query param summary
		var qrSummary bool

		if o.Summary != nil {
			qrSummary = *o.Summary
		}
		qSummary := swag.FormatBool(qrSummary)
		if qSummary != "" {

			if err := r.SetQueryParam("summary", qSummary); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamHardwareModelQueryGlobalHardwareModels binds the parameter brandId
func (o *HardwareModelQueryGlobalHardwareModelsParams) bindParamBrandID(formats strfmt.Registry) []string {
	brandIDIR := o.BrandID

	var brandIDIC []string
	for _, brandIDIIR := range brandIDIR { // explode []string

		brandIDIIV := brandIDIIR // string as string
		brandIDIC = append(brandIDIC, brandIDIIV)
	}

	// items.CollectionFormat: "multi"
	brandIDIS := swag.JoinByFormat(brandIDIC, "multi")

	return brandIDIS
}

// bindParamHardwareModelQueryGlobalHardwareModels binds the parameter next.orderBy
func (o *HardwareModelQueryGlobalHardwareModelsParams) bindParamNextOrderBy(formats strfmt.Registry) []string {
	nextOrderByIR := o.NextOrderBy

	var nextOrderByIC []string
	for _, nextOrderByIIR := range nextOrderByIR { // explode []string

		nextOrderByIIV := nextOrderByIIR // string as string
		nextOrderByIC = append(nextOrderByIC, nextOrderByIIV)
	}

	// items.CollectionFormat: "multi"
	nextOrderByIS := swag.JoinByFormat(nextOrderByIC, "multi")

	return nextOrderByIS
}
