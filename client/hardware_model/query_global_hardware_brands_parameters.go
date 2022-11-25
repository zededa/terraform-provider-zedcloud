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

// NewQueryGlobalHardwareBrandsParams creates a new QueryGlobalHardwareBrandsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryGlobalHardwareBrandsParams() *QueryGlobalHardwareBrandsParams {
	return &QueryGlobalHardwareBrandsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryGlobalHardwareBrandsParamsWithTimeout creates a new QueryGlobalHardwareBrandsParams object
// with the ability to set a timeout on a request.
func NewQueryGlobalHardwareBrandsParamsWithTimeout(timeout time.Duration) *QueryGlobalHardwareBrandsParams {
	return &QueryGlobalHardwareBrandsParams{
		timeout: timeout,
	}
}

// NewQueryGlobalHardwareBrandsParamsWithContext creates a new QueryGlobalHardwareBrandsParams object
// with the ability to set a context for a request.
func NewQueryGlobalHardwareBrandsParamsWithContext(ctx context.Context) *QueryGlobalHardwareBrandsParams {
	return &QueryGlobalHardwareBrandsParams{
		Context: ctx,
	}
}

// NewQueryGlobalHardwareBrandsParamsWithHTTPClient creates a new QueryGlobalHardwareBrandsParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryGlobalHardwareBrandsParamsWithHTTPClient(client *http.Client) *QueryGlobalHardwareBrandsParams {
	return &QueryGlobalHardwareBrandsParams{
		HTTPClient: client,
	}
}

/*
QueryGlobalHardwareBrandsParams contains all the parameters to send to the API endpoint

	for the query global hardware brands operation.

	Typically these are written to a http.Request.
*/
type QueryGlobalHardwareBrandsParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* EnterpriseID.

	   deprecated field
	*/
	EnterpriseID *string

	/* FilterNamePattern.

	   Brand name pattern to be matched.
	*/
	FilterNamePattern *string

	/* FilterOriginType.

	   origin of object - ORIGIN_UNSPECIFIED: default options, which says no Operation/Invalid Operation - ORIGIN_IMPORTED: Object imported from global enterprise. - ORIGIN_LOCAL: Objectl created locally. - ORIGIN_GLOBAL: Object created in global store,to use this type user should have root previlage.

	   Default: "ORIGIN_UNSPECIFIED"
	*/
	FilterOriginType *string

	/* NextOrderBy.

	   OrderBy helps in sorting the list response
	*/
	NextOrderBy *string

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

	/* Summary.

	   Only summary of the records required

	   Format: boolean
	*/
	Summary *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query global hardware brands params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryGlobalHardwareBrandsParams) WithDefaults() *QueryGlobalHardwareBrandsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query global hardware brands params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryGlobalHardwareBrandsParams) SetDefaults() {
	var (
		filterOriginTypeDefault = string("ORIGIN_UNSPECIFIED")
	)

	val := QueryGlobalHardwareBrandsParams{
		FilterOriginType: &filterOriginTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) WithTimeout(timeout time.Duration) *QueryGlobalHardwareBrandsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) WithContext(ctx context.Context) *QueryGlobalHardwareBrandsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) WithHTTPClient(client *http.Client) *QueryGlobalHardwareBrandsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) WithXRequestID(xRequestID *string) *QueryGlobalHardwareBrandsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithEnterpriseID adds the enterpriseID to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) WithEnterpriseID(enterpriseID *string) *QueryGlobalHardwareBrandsParams {
	o.SetEnterpriseID(enterpriseID)
	return o
}

// SetEnterpriseID adds the enterpriseId to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) SetEnterpriseID(enterpriseID *string) {
	o.EnterpriseID = enterpriseID
}

// WithFilterNamePattern adds the filterNamePattern to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) WithFilterNamePattern(filterNamePattern *string) *QueryGlobalHardwareBrandsParams {
	o.SetFilterNamePattern(filterNamePattern)
	return o
}

// SetFilterNamePattern adds the filterNamePattern to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) SetFilterNamePattern(filterNamePattern *string) {
	o.FilterNamePattern = filterNamePattern
}

// WithFilterOriginType adds the filterOriginType to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) WithFilterOriginType(filterOriginType *string) *QueryGlobalHardwareBrandsParams {
	o.SetFilterOriginType(filterOriginType)
	return o
}

// SetFilterOriginType adds the filterOriginType to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) SetFilterOriginType(filterOriginType *string) {
	o.FilterOriginType = filterOriginType
}

// WithNextOrderBy adds the nextOrderBy to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) WithNextOrderBy(nextOrderBy *string) *QueryGlobalHardwareBrandsParams {
	o.SetNextOrderBy(nextOrderBy)
	return o
}

// SetNextOrderBy adds the nextOrderBy to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) SetNextOrderBy(nextOrderBy *string) {
	o.NextOrderBy = nextOrderBy
}

// WithNextPageNum adds the nextPageNum to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) WithNextPageNum(nextPageNum *int64) *QueryGlobalHardwareBrandsParams {
	o.SetNextPageNum(nextPageNum)
	return o
}

// SetNextPageNum adds the nextPageNum to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) SetNextPageNum(nextPageNum *int64) {
	o.NextPageNum = nextPageNum
}

// WithNextPageSize adds the nextPageSize to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) WithNextPageSize(nextPageSize *int64) *QueryGlobalHardwareBrandsParams {
	o.SetNextPageSize(nextPageSize)
	return o
}

// SetNextPageSize adds the nextPageSize to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) SetNextPageSize(nextPageSize *int64) {
	o.NextPageSize = nextPageSize
}

// WithNextPageToken adds the nextPageToken to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) WithNextPageToken(nextPageToken *string) *QueryGlobalHardwareBrandsParams {
	o.SetNextPageToken(nextPageToken)
	return o
}

// SetNextPageToken adds the nextPageToken to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) SetNextPageToken(nextPageToken *string) {
	o.NextPageToken = nextPageToken
}

// WithNextTotalPages adds the nextTotalPages to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) WithNextTotalPages(nextTotalPages *int64) *QueryGlobalHardwareBrandsParams {
	o.SetNextTotalPages(nextTotalPages)
	return o
}

// SetNextTotalPages adds the nextTotalPages to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) SetNextTotalPages(nextTotalPages *int64) {
	o.NextTotalPages = nextTotalPages
}

// WithSummary adds the summary to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) WithSummary(summary *bool) *QueryGlobalHardwareBrandsParams {
	o.SetSummary(summary)
	return o
}

// SetSummary adds the summary to the query global hardware brands params
func (o *QueryGlobalHardwareBrandsParams) SetSummary(summary *bool) {
	o.Summary = summary
}

// WriteToRequest writes these params to a swagger request
func (o *QueryGlobalHardwareBrandsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EnterpriseID != nil {

		// query param enterpriseId
		var qrEnterpriseID string

		if o.EnterpriseID != nil {
			qrEnterpriseID = *o.EnterpriseID
		}
		qEnterpriseID := qrEnterpriseID
		if qEnterpriseID != "" {

			if err := r.SetQueryParam("enterpriseId", qEnterpriseID); err != nil {
				return err
			}
		}
	}

	if o.FilterNamePattern != nil {

		// query param filter.namePattern
		var qrFilterNamePattern string

		if o.FilterNamePattern != nil {
			qrFilterNamePattern = *o.FilterNamePattern
		}
		qFilterNamePattern := qrFilterNamePattern
		if qFilterNamePattern != "" {

			if err := r.SetQueryParam("filter.namePattern", qFilterNamePattern); err != nil {
				return err
			}
		}
	}

	if o.FilterOriginType != nil {

		// query param filter.originType
		var qrFilterOriginType string

		if o.FilterOriginType != nil {
			qrFilterOriginType = *o.FilterOriginType
		}
		qFilterOriginType := qrFilterOriginType
		if qFilterOriginType != "" {

			if err := r.SetQueryParam("filter.originType", qFilterOriginType); err != nil {
				return err
			}
		}
	}

	if o.NextOrderBy != nil {

		// query param next.orderBy
		var qrNextOrderBy string

		if o.NextOrderBy != nil {
			qrNextOrderBy = *o.NextOrderBy
		}
		qNextOrderBy := qrNextOrderBy
		if qNextOrderBy != "" {

			if err := r.SetQueryParam("next.orderBy", qNextOrderBy); err != nil {
				return err
			}
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
