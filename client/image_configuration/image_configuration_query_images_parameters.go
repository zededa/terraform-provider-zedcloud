package image_configuration

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

// NewImageConfigurationQueryImagesParams creates a new ImageConfigurationQueryImagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImageConfigurationQueryImagesParams() *ImageConfigurationQueryImagesParams {
	return &ImageConfigurationQueryImagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageConfigurationQueryImagesParamsWithTimeout creates a new ImageConfigurationQueryImagesParams object
// with the ability to set a timeout on a request.
func NewImageConfigurationQueryImagesParamsWithTimeout(timeout time.Duration) *ImageConfigurationQueryImagesParams {
	return &ImageConfigurationQueryImagesParams{
		timeout: timeout,
	}
}

// NewImageConfigurationQueryImagesParamsWithContext creates a new ImageConfigurationQueryImagesParams object
// with the ability to set a context for a request.
func NewImageConfigurationQueryImagesParamsWithContext(ctx context.Context) *ImageConfigurationQueryImagesParams {
	return &ImageConfigurationQueryImagesParams{
		Context: ctx,
	}
}

// NewImageConfigurationQueryImagesParamsWithHTTPClient creates a new ImageConfigurationQueryImagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageConfigurationQueryImagesParamsWithHTTPClient(client *http.Client) *ImageConfigurationQueryImagesParams {
	return &ImageConfigurationQueryImagesParams{
		HTTPClient: client,
	}
}

/*
ImageConfigurationQueryImagesParams contains all the parameters to send to the API endpoint

	for the image configuration query images operation.

	Typically these are written to a http.Request.
*/
type ImageConfigurationQueryImagesParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* FilterDatastoreID.

	   Datastore id to be matched.
	*/
	FilterDatastoreID *string

	/* FilterImageArch.

	   Image architecture to be matched.

	   Default: "UNSPECIFIED"
	*/
	FilterImageArch *string

	/* FilterImageStatus.

	    Image status to be matched.

	- IMAGE_STATUS_CREATED: Image metadata is created
	- IMAGE_STATUS_UPLOADING: Image binary is being uploaded to Datstore
	- IMAGE_STATUS_READY: Image is ready for download
	- IMAGE_STATUS_INUSE: Image is being used by edge applications
	- IMAGE_STATUS_FAILED: Image binary upload has failed
	- IMAGE_STATUS_UPLINKING: Image metadata is being uplinked with Datstore binary

	    Default: "IMAGE_STATUS_UNSPECIFIED"
	*/
	FilterImageStatus *string

	/* FilterImageType.

	    Image type to ne matched.

	- IMAGE_TYPE_EVE: Base OS images for edge gateway
	- IMAGE_TYPE_APPLICATION: Edge application images
	- IMAGE_TYPE_EVEPRIVATE: Private Base OS images for edge gateway

	    Default: "IMAGE_TYPE_UNSPECIFIED"
	*/
	FilterImageType *string

	/* FilterNamePattern.

	   Image name pattern to be matched.
	*/
	FilterNamePattern *string

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

	/* Summary.

	   Only summary of the records required
	*/
	Summary *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image configuration query images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationQueryImagesParams) WithDefaults() *ImageConfigurationQueryImagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image configuration query images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationQueryImagesParams) SetDefaults() {
	var (
		filterImageArchDefault = string("UNSPECIFIED")

		filterImageStatusDefault = string("IMAGE_STATUS_UNSPECIFIED")

		filterImageTypeDefault = string("IMAGE_TYPE_UNSPECIFIED")
	)

	val := ImageConfigurationQueryImagesParams{
		FilterImageArch:   &filterImageArchDefault,
		FilterImageStatus: &filterImageStatusDefault,
		FilterImageType:   &filterImageTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) WithTimeout(timeout time.Duration) *ImageConfigurationQueryImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) WithContext(ctx context.Context) *ImageConfigurationQueryImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) WithHTTPClient(client *http.Client) *ImageConfigurationQueryImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) WithXRequestID(xRequestID *string) *ImageConfigurationQueryImagesParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithFilterDatastoreID adds the filterDatastoreID to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) WithFilterDatastoreID(filterDatastoreID *string) *ImageConfigurationQueryImagesParams {
	o.SetFilterDatastoreID(filterDatastoreID)
	return o
}

// SetFilterDatastoreID adds the filterDatastoreId to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) SetFilterDatastoreID(filterDatastoreID *string) {
	o.FilterDatastoreID = filterDatastoreID
}

// WithFilterImageArch adds the filterImageArch to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) WithFilterImageArch(filterImageArch *string) *ImageConfigurationQueryImagesParams {
	o.SetFilterImageArch(filterImageArch)
	return o
}

// SetFilterImageArch adds the filterImageArch to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) SetFilterImageArch(filterImageArch *string) {
	o.FilterImageArch = filterImageArch
}

// WithFilterImageStatus adds the filterImageStatus to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) WithFilterImageStatus(filterImageStatus *string) *ImageConfigurationQueryImagesParams {
	o.SetFilterImageStatus(filterImageStatus)
	return o
}

// SetFilterImageStatus adds the filterImageStatus to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) SetFilterImageStatus(filterImageStatus *string) {
	o.FilterImageStatus = filterImageStatus
}

// WithFilterImageType adds the filterImageType to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) WithFilterImageType(filterImageType *string) *ImageConfigurationQueryImagesParams {
	o.SetFilterImageType(filterImageType)
	return o
}

// SetFilterImageType adds the filterImageType to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) SetFilterImageType(filterImageType *string) {
	o.FilterImageType = filterImageType
}

// WithFilterNamePattern adds the filterNamePattern to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) WithFilterNamePattern(filterNamePattern *string) *ImageConfigurationQueryImagesParams {
	o.SetFilterNamePattern(filterNamePattern)
	return o
}

// SetFilterNamePattern adds the filterNamePattern to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) SetFilterNamePattern(filterNamePattern *string) {
	o.FilterNamePattern = filterNamePattern
}

// WithNextOrderBy adds the nextOrderBy to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) WithNextOrderBy(nextOrderBy []string) *ImageConfigurationQueryImagesParams {
	o.SetNextOrderBy(nextOrderBy)
	return o
}

// SetNextOrderBy adds the nextOrderBy to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) SetNextOrderBy(nextOrderBy []string) {
	o.NextOrderBy = nextOrderBy
}

// WithNextPageNum adds the nextPageNum to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) WithNextPageNum(nextPageNum *int64) *ImageConfigurationQueryImagesParams {
	o.SetNextPageNum(nextPageNum)
	return o
}

// SetNextPageNum adds the nextPageNum to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) SetNextPageNum(nextPageNum *int64) {
	o.NextPageNum = nextPageNum
}

// WithNextPageSize adds the nextPageSize to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) WithNextPageSize(nextPageSize *int64) *ImageConfigurationQueryImagesParams {
	o.SetNextPageSize(nextPageSize)
	return o
}

// SetNextPageSize adds the nextPageSize to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) SetNextPageSize(nextPageSize *int64) {
	o.NextPageSize = nextPageSize
}

// WithNextPageToken adds the nextPageToken to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) WithNextPageToken(nextPageToken *string) *ImageConfigurationQueryImagesParams {
	o.SetNextPageToken(nextPageToken)
	return o
}

// SetNextPageToken adds the nextPageToken to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) SetNextPageToken(nextPageToken *string) {
	o.NextPageToken = nextPageToken
}

// WithNextTotalPages adds the nextTotalPages to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) WithNextTotalPages(nextTotalPages *int64) *ImageConfigurationQueryImagesParams {
	o.SetNextTotalPages(nextTotalPages)
	return o
}

// SetNextTotalPages adds the nextTotalPages to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) SetNextTotalPages(nextTotalPages *int64) {
	o.NextTotalPages = nextTotalPages
}

// WithSummary adds the summary to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) WithSummary(summary *bool) *ImageConfigurationQueryImagesParams {
	o.SetSummary(summary)
	return o
}

// SetSummary adds the summary to the image configuration query images params
func (o *ImageConfigurationQueryImagesParams) SetSummary(summary *bool) {
	o.Summary = summary
}

// WriteToRequest writes these params to a swagger request
func (o *ImageConfigurationQueryImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FilterDatastoreID != nil {

		// query param filter.datastoreId
		var qrFilterDatastoreID string

		if o.FilterDatastoreID != nil {
			qrFilterDatastoreID = *o.FilterDatastoreID
		}
		qFilterDatastoreID := qrFilterDatastoreID
		if qFilterDatastoreID != "" {

			if err := r.SetQueryParam("filter.datastoreId", qFilterDatastoreID); err != nil {
				return err
			}
		}
	}

	if o.FilterImageArch != nil {

		// query param filter.imageArch
		var qrFilterImageArch string

		if o.FilterImageArch != nil {
			qrFilterImageArch = *o.FilterImageArch
		}
		qFilterImageArch := qrFilterImageArch
		if qFilterImageArch != "" {

			if err := r.SetQueryParam("filter.imageArch", qFilterImageArch); err != nil {
				return err
			}
		}
	}

	if o.FilterImageStatus != nil {

		// query param filter.imageStatus
		var qrFilterImageStatus string

		if o.FilterImageStatus != nil {
			qrFilterImageStatus = *o.FilterImageStatus
		}
		qFilterImageStatus := qrFilterImageStatus
		if qFilterImageStatus != "" {

			if err := r.SetQueryParam("filter.imageStatus", qFilterImageStatus); err != nil {
				return err
			}
		}
	}

	if o.FilterImageType != nil {

		// query param filter.imageType
		var qrFilterImageType string

		if o.FilterImageType != nil {
			qrFilterImageType = *o.FilterImageType
		}
		qFilterImageType := qrFilterImageType
		if qFilterImageType != "" {

			if err := r.SetQueryParam("filter.imageType", qFilterImageType); err != nil {
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

// bindParamImageConfigurationQueryImages binds the parameter next.orderBy
func (o *ImageConfigurationQueryImagesParams) bindParamNextOrderBy(formats strfmt.Registry) []string {
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
