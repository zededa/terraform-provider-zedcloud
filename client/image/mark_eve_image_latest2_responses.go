package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/zededa/terraform-provider-zedcloud/models"
)

// ImageConfigurationMarkEveImageLatest2Reader is a Reader for the ImageConfigurationMarkEveImageLatest2 structure.
type ImageConfigurationMarkEveImageLatest2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageConfigurationMarkEveImageLatest2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageConfigurationMarkEveImageLatest2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImageConfigurationMarkEveImageLatest2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImageConfigurationMarkEveImageLatest2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImageConfigurationMarkEveImageLatest2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImageConfigurationMarkEveImageLatest2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageConfigurationMarkEveImageLatest2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewImageConfigurationMarkEveImageLatest2GatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewImageConfigurationMarkEveImageLatest2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImageConfigurationMarkEveImageLatest2OK creates a ImageConfigurationMarkEveImageLatest2OK with default headers values
func NewImageConfigurationMarkEveImageLatest2OK() *ImageConfigurationMarkEveImageLatest2OK {
	return &ImageConfigurationMarkEveImageLatest2OK{}
}

/*
ImageConfigurationMarkEveImageLatest2OK describes a response with status code 200, with default header values.

A successful response.
*/
type ImageConfigurationMarkEveImageLatest2OK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration mark eve image latest2 o k response has a 2xx status code
func (o *ImageConfigurationMarkEveImageLatest2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image configuration mark eve image latest2 o k response has a 3xx status code
func (o *ImageConfigurationMarkEveImageLatest2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration mark eve image latest2 o k response has a 4xx status code
func (o *ImageConfigurationMarkEveImageLatest2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration mark eve image latest2 o k response has a 5xx status code
func (o *ImageConfigurationMarkEveImageLatest2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration mark eve image latest2 o k response a status code equal to that given
func (o *ImageConfigurationMarkEveImageLatest2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the image configuration mark eve image latest2 o k response
func (o *ImageConfigurationMarkEveImageLatest2OK) Code() int {
	return 200
}

func (o *ImageConfigurationMarkEveImageLatest2OK) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationMarkEveImageLatest2OK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationMarkEveImageLatest2OK) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationMarkEveImageLatest2OK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationMarkEveImageLatest2OK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationMarkEveImageLatest2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationMarkEveImageLatest2BadRequest creates a ImageConfigurationMarkEveImageLatest2BadRequest with default headers values
func NewImageConfigurationMarkEveImageLatest2BadRequest() *ImageConfigurationMarkEveImageLatest2BadRequest {
	return &ImageConfigurationMarkEveImageLatest2BadRequest{}
}

/*
ImageConfigurationMarkEveImageLatest2BadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type ImageConfigurationMarkEveImageLatest2BadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration mark eve image latest2 bad request response has a 2xx status code
func (o *ImageConfigurationMarkEveImageLatest2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration mark eve image latest2 bad request response has a 3xx status code
func (o *ImageConfigurationMarkEveImageLatest2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration mark eve image latest2 bad request response has a 4xx status code
func (o *ImageConfigurationMarkEveImageLatest2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration mark eve image latest2 bad request response has a 5xx status code
func (o *ImageConfigurationMarkEveImageLatest2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration mark eve image latest2 bad request response a status code equal to that given
func (o *ImageConfigurationMarkEveImageLatest2BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the image configuration mark eve image latest2 bad request response
func (o *ImageConfigurationMarkEveImageLatest2BadRequest) Code() int {
	return 400
}

func (o *ImageConfigurationMarkEveImageLatest2BadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationMarkEveImageLatest2BadRequest  %+v", 400, o.Payload)
}

func (o *ImageConfigurationMarkEveImageLatest2BadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationMarkEveImageLatest2BadRequest  %+v", 400, o.Payload)
}

func (o *ImageConfigurationMarkEveImageLatest2BadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationMarkEveImageLatest2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationMarkEveImageLatest2Unauthorized creates a ImageConfigurationMarkEveImageLatest2Unauthorized with default headers values
func NewImageConfigurationMarkEveImageLatest2Unauthorized() *ImageConfigurationMarkEveImageLatest2Unauthorized {
	return &ImageConfigurationMarkEveImageLatest2Unauthorized{}
}

/*
ImageConfigurationMarkEveImageLatest2Unauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ImageConfigurationMarkEveImageLatest2Unauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration mark eve image latest2 unauthorized response has a 2xx status code
func (o *ImageConfigurationMarkEveImageLatest2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration mark eve image latest2 unauthorized response has a 3xx status code
func (o *ImageConfigurationMarkEveImageLatest2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration mark eve image latest2 unauthorized response has a 4xx status code
func (o *ImageConfigurationMarkEveImageLatest2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration mark eve image latest2 unauthorized response has a 5xx status code
func (o *ImageConfigurationMarkEveImageLatest2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration mark eve image latest2 unauthorized response a status code equal to that given
func (o *ImageConfigurationMarkEveImageLatest2Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the image configuration mark eve image latest2 unauthorized response
func (o *ImageConfigurationMarkEveImageLatest2Unauthorized) Code() int {
	return 401
}

func (o *ImageConfigurationMarkEveImageLatest2Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationMarkEveImageLatest2Unauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationMarkEveImageLatest2Unauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationMarkEveImageLatest2Unauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationMarkEveImageLatest2Unauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationMarkEveImageLatest2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationMarkEveImageLatest2Forbidden creates a ImageConfigurationMarkEveImageLatest2Forbidden with default headers values
func NewImageConfigurationMarkEveImageLatest2Forbidden() *ImageConfigurationMarkEveImageLatest2Forbidden {
	return &ImageConfigurationMarkEveImageLatest2Forbidden{}
}

/*
ImageConfigurationMarkEveImageLatest2Forbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ImageConfigurationMarkEveImageLatest2Forbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration mark eve image latest2 forbidden response has a 2xx status code
func (o *ImageConfigurationMarkEveImageLatest2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration mark eve image latest2 forbidden response has a 3xx status code
func (o *ImageConfigurationMarkEveImageLatest2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration mark eve image latest2 forbidden response has a 4xx status code
func (o *ImageConfigurationMarkEveImageLatest2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration mark eve image latest2 forbidden response has a 5xx status code
func (o *ImageConfigurationMarkEveImageLatest2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration mark eve image latest2 forbidden response a status code equal to that given
func (o *ImageConfigurationMarkEveImageLatest2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the image configuration mark eve image latest2 forbidden response
func (o *ImageConfigurationMarkEveImageLatest2Forbidden) Code() int {
	return 403
}

func (o *ImageConfigurationMarkEveImageLatest2Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationMarkEveImageLatest2Forbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationMarkEveImageLatest2Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationMarkEveImageLatest2Forbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationMarkEveImageLatest2Forbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationMarkEveImageLatest2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationMarkEveImageLatest2NotFound creates a ImageConfigurationMarkEveImageLatest2NotFound with default headers values
func NewImageConfigurationMarkEveImageLatest2NotFound() *ImageConfigurationMarkEveImageLatest2NotFound {
	return &ImageConfigurationMarkEveImageLatest2NotFound{}
}

/*
ImageConfigurationMarkEveImageLatest2NotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type ImageConfigurationMarkEveImageLatest2NotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration mark eve image latest2 not found response has a 2xx status code
func (o *ImageConfigurationMarkEveImageLatest2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration mark eve image latest2 not found response has a 3xx status code
func (o *ImageConfigurationMarkEveImageLatest2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration mark eve image latest2 not found response has a 4xx status code
func (o *ImageConfigurationMarkEveImageLatest2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration mark eve image latest2 not found response has a 5xx status code
func (o *ImageConfigurationMarkEveImageLatest2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration mark eve image latest2 not found response a status code equal to that given
func (o *ImageConfigurationMarkEveImageLatest2NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the image configuration mark eve image latest2 not found response
func (o *ImageConfigurationMarkEveImageLatest2NotFound) Code() int {
	return 404
}

func (o *ImageConfigurationMarkEveImageLatest2NotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationMarkEveImageLatest2NotFound  %+v", 404, o.Payload)
}

func (o *ImageConfigurationMarkEveImageLatest2NotFound) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationMarkEveImageLatest2NotFound  %+v", 404, o.Payload)
}

func (o *ImageConfigurationMarkEveImageLatest2NotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationMarkEveImageLatest2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationMarkEveImageLatest2InternalServerError creates a ImageConfigurationMarkEveImageLatest2InternalServerError with default headers values
func NewImageConfigurationMarkEveImageLatest2InternalServerError() *ImageConfigurationMarkEveImageLatest2InternalServerError {
	return &ImageConfigurationMarkEveImageLatest2InternalServerError{}
}

/*
ImageConfigurationMarkEveImageLatest2InternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ImageConfigurationMarkEveImageLatest2InternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration mark eve image latest2 internal server error response has a 2xx status code
func (o *ImageConfigurationMarkEveImageLatest2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration mark eve image latest2 internal server error response has a 3xx status code
func (o *ImageConfigurationMarkEveImageLatest2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration mark eve image latest2 internal server error response has a 4xx status code
func (o *ImageConfigurationMarkEveImageLatest2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration mark eve image latest2 internal server error response has a 5xx status code
func (o *ImageConfigurationMarkEveImageLatest2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration mark eve image latest2 internal server error response a status code equal to that given
func (o *ImageConfigurationMarkEveImageLatest2InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the image configuration mark eve image latest2 internal server error response
func (o *ImageConfigurationMarkEveImageLatest2InternalServerError) Code() int {
	return 500
}

func (o *ImageConfigurationMarkEveImageLatest2InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationMarkEveImageLatest2InternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationMarkEveImageLatest2InternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationMarkEveImageLatest2InternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationMarkEveImageLatest2InternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationMarkEveImageLatest2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationMarkEveImageLatest2GatewayTimeout creates a ImageConfigurationMarkEveImageLatest2GatewayTimeout with default headers values
func NewImageConfigurationMarkEveImageLatest2GatewayTimeout() *ImageConfigurationMarkEveImageLatest2GatewayTimeout {
	return &ImageConfigurationMarkEveImageLatest2GatewayTimeout{}
}

/*
ImageConfigurationMarkEveImageLatest2GatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ImageConfigurationMarkEveImageLatest2GatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration mark eve image latest2 gateway timeout response has a 2xx status code
func (o *ImageConfigurationMarkEveImageLatest2GatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration mark eve image latest2 gateway timeout response has a 3xx status code
func (o *ImageConfigurationMarkEveImageLatest2GatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration mark eve image latest2 gateway timeout response has a 4xx status code
func (o *ImageConfigurationMarkEveImageLatest2GatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration mark eve image latest2 gateway timeout response has a 5xx status code
func (o *ImageConfigurationMarkEveImageLatest2GatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration mark eve image latest2 gateway timeout response a status code equal to that given
func (o *ImageConfigurationMarkEveImageLatest2GatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the image configuration mark eve image latest2 gateway timeout response
func (o *ImageConfigurationMarkEveImageLatest2GatewayTimeout) Code() int {
	return 504
}

func (o *ImageConfigurationMarkEveImageLatest2GatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationMarkEveImageLatest2GatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationMarkEveImageLatest2GatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationMarkEveImageLatest2GatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationMarkEveImageLatest2GatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationMarkEveImageLatest2GatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationMarkEveImageLatest2Default creates a ImageConfigurationMarkEveImageLatest2Default with default headers values
func NewImageConfigurationMarkEveImageLatest2Default(code int) *ImageConfigurationMarkEveImageLatest2Default {
	return &ImageConfigurationMarkEveImageLatest2Default{
		_statusCode: code,
	}
}

/*
ImageConfigurationMarkEveImageLatest2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ImageConfigurationMarkEveImageLatest2Default struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this image configuration mark eve image latest2 default response has a 2xx status code
func (o *ImageConfigurationMarkEveImageLatest2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this image configuration mark eve image latest2 default response has a 3xx status code
func (o *ImageConfigurationMarkEveImageLatest2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this image configuration mark eve image latest2 default response has a 4xx status code
func (o *ImageConfigurationMarkEveImageLatest2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this image configuration mark eve image latest2 default response has a 5xx status code
func (o *ImageConfigurationMarkEveImageLatest2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this image configuration mark eve image latest2 default response a status code equal to that given
func (o *ImageConfigurationMarkEveImageLatest2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the image configuration mark eve image latest2 default response
func (o *ImageConfigurationMarkEveImageLatest2Default) Code() int {
	return o._statusCode
}

func (o *ImageConfigurationMarkEveImageLatest2Default) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] ImageConfiguration_MarkEveImageLatest2 default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationMarkEveImageLatest2Default) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] ImageConfiguration_MarkEveImageLatest2 default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationMarkEveImageLatest2Default) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ImageConfigurationMarkEveImageLatest2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ImageConfigurationMarkEveImageLatest2Body Image metadata detail
//
// Image metadata for edge gateway Base OS or for eedge applications.
// Example: {"description":"My test image in QCOW2 format for Edge computing","dsId":"7927f6e3-484d-4105-a98e-868b21c1cb61","id":"d1125b0f-633d-459c-99c6-f47e7467cebc","imageArch":"AMD64","imageError":"Image uplinked successfully...","imageFormat":3,"imageLocal":"","imageRelUrl":"edge/computing/AMD64","imageSha256":"ADC5BB9DD39F83DD74C276B0BA119FB27925A5CDEA343FE1F2C8433F28AB345B","imageSizeBytes":142016512,"imageStatus":4,"imageType":2,"imageVersion":"","name":"my-test-image","originType":2,"revision":{"createdAt":{"seconds":1592068270},"createdBy":"admin@my-company.com","curr":"1","updatedAt":{"seconds":1592068271},"updatedBy":"admin@my-company.com"},"title":"My Test Image for Edge Computing"}
swagger:model ImageConfigurationMarkEveImageLatest2Body
*/
type ImageConfigurationMarkEveImageLatest2Body struct {

	// Datastore Id where image binary is located.
	// Required: true
	// Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}
	DatastoreID *string `json:"datastoreId"`

	// Detailed description of the image.
	// Max Length: 256
	Description string `json:"description,omitempty"`

	// System defined universally unique Id of the image.
	// Read Only: true
	// Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}
	ID string `json:"id,omitempty"`

	// Image upload/uplink detailed error/status message
	// Read Only: true
	ImageError string `json:"imageError,omitempty"`

	// Image binary format.
	// Required: true
	ImageFormat *models.ConfigFormat `json:"imageFormat"`

	// Internal image location.
	// Read Only: true
	ImageLocal string `json:"imageLocal,omitempty"`

	// Image relative path w.r.t. Datastore
	ImageRelURL string `json:"imageRelUrl,omitempty"`

	// Image checksum in SHA256 format
	ImageSha256 string `json:"imageSha256,omitempty"`

	// Image size in KBytes.
	ImageSizeBytes string `json:"imageSizeBytes,omitempty"`

	// Image status
	// Read Only: true
	ImageStatus *models.ImageStatus `json:"imageStatus,omitempty"`

	// Image type
	// Required: true
	ImageType *models.ImageType `json:"imageType"`

	// system defined info
	ImageVersion string `json:"imageVersion,omitempty"`

	// User defined name of the image, unique across the enterprise. Once image is created, name can’t be changed.
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Name *string `json:"name"`

	// Origin type of image.
	// Read Only: true
	OriginType *models.Origin `json:"originType,omitempty"`

	// project access list of the image
	ProjectAccessList []string `json:"projectAccessList"`

	// system defined info
	// Read Only: true
	Revision *models.ObjectRevision `json:"revision,omitempty"`

	// User defined title of the image. Title can be changed at any time.
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+
	Title *string `json:"title"`
}

// Validate validates this image configuration mark eve image latest2 body
func (o *ImageConfigurationMarkEveImageLatest2Body) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDatastoreID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateImageFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateImageStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateImageType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOriginType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ImageConfigurationMarkEveImageLatest2Body) validateDatastoreID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"datastoreId", "body", o.DatastoreID); err != nil {
		return err
	}

	if err := validate.Pattern("body"+"."+"datastoreId", "body", *o.DatastoreID, `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`); err != nil {
		return err
	}

	return nil
}

func (o *ImageConfigurationMarkEveImageLatest2Body) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(o.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("body"+"."+"description", "body", o.Description, 256); err != nil {
		return err
	}

	return nil
}

func (o *ImageConfigurationMarkEveImageLatest2Body) validateID(formats strfmt.Registry) error {
	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"id", "body", o.ID, `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`); err != nil {
		return err
	}

	return nil
}

func (o *ImageConfigurationMarkEveImageLatest2Body) validateImageFormat(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"imageFormat", "body", o.ImageFormat); err != nil {
		return err
	}

	if err := validate.Required("body"+"."+"imageFormat", "body", o.ImageFormat); err != nil {
		return err
	}

	if o.ImageFormat != nil {
		if err := o.ImageFormat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "imageFormat")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "imageFormat")
			}
			return err
		}
	}

	return nil
}

func (o *ImageConfigurationMarkEveImageLatest2Body) validateImageStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.ImageStatus) { // not required
		return nil
	}

	if o.ImageStatus != nil {
		if err := o.ImageStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "imageStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "imageStatus")
			}
			return err
		}
	}

	return nil
}

func (o *ImageConfigurationMarkEveImageLatest2Body) validateImageType(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"imageType", "body", o.ImageType); err != nil {
		return err
	}

	if err := validate.Required("body"+"."+"imageType", "body", o.ImageType); err != nil {
		return err
	}

	if o.ImageType != nil {
		if err := o.ImageType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "imageType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "imageType")
			}
			return err
		}
	}

	return nil
}

func (o *ImageConfigurationMarkEveImageLatest2Body) validateName(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	if err := validate.MinLength("body"+"."+"name", "body", *o.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("body"+"."+"name", "body", *o.Name, 256); err != nil {
		return err
	}

	if err := validate.Pattern("body"+"."+"name", "body", *o.Name, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (o *ImageConfigurationMarkEveImageLatest2Body) validateOriginType(formats strfmt.Registry) error {
	if swag.IsZero(o.OriginType) { // not required
		return nil
	}

	if o.OriginType != nil {
		if err := o.OriginType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "originType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "originType")
			}
			return err
		}
	}

	return nil
}

func (o *ImageConfigurationMarkEveImageLatest2Body) validateRevision(formats strfmt.Registry) error {
	if swag.IsZero(o.Revision) { // not required
		return nil
	}

	if o.Revision != nil {
		if err := o.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "revision")
			}
			return err
		}
	}

	return nil
}

func (o *ImageConfigurationMarkEveImageLatest2Body) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"title", "body", o.Title); err != nil {
		return err
	}

	if err := validate.MinLength("body"+"."+"title", "body", *o.Title, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("body"+"."+"title", "body", *o.Title, 256); err != nil {
		return err
	}

	if err := validate.Pattern("body"+"."+"title", "body", *o.Title, `[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this image configuration mark eve image latest2 body based on the context it is used
func (o *ImageConfigurationMarkEveImageLatest2Body) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateImageError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateImageFormat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateImageLocal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateImageStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateImageType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOriginType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ImageConfigurationMarkEveImageLatest2Body) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "body"+"."+"id", "body", string(o.ID)); err != nil {
		return err
	}

	return nil
}

func (o *ImageConfigurationMarkEveImageLatest2Body) contextValidateImageError(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "body"+"."+"imageError", "body", string(o.ImageError)); err != nil {
		return err
	}

	return nil
}

func (o *ImageConfigurationMarkEveImageLatest2Body) contextValidateImageFormat(ctx context.Context, formats strfmt.Registry) error {

	if o.ImageFormat != nil {
		if err := o.ImageFormat.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "imageFormat")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "imageFormat")
			}
			return err
		}
	}

	return nil
}

func (o *ImageConfigurationMarkEveImageLatest2Body) contextValidateImageLocal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "body"+"."+"imageLocal", "body", string(o.ImageLocal)); err != nil {
		return err
	}

	return nil
}

func (o *ImageConfigurationMarkEveImageLatest2Body) contextValidateImageStatus(ctx context.Context, formats strfmt.Registry) error {

	if o.ImageStatus != nil {
		if err := o.ImageStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "imageStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "imageStatus")
			}
			return err
		}
	}

	return nil
}

func (o *ImageConfigurationMarkEveImageLatest2Body) contextValidateImageType(ctx context.Context, formats strfmt.Registry) error {

	if o.ImageType != nil {
		if err := o.ImageType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "imageType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "imageType")
			}
			return err
		}
	}

	return nil
}

func (o *ImageConfigurationMarkEveImageLatest2Body) contextValidateOriginType(ctx context.Context, formats strfmt.Registry) error {

	if o.OriginType != nil {
		if err := o.OriginType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "originType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "originType")
			}
			return err
		}
	}

	return nil
}

func (o *ImageConfigurationMarkEveImageLatest2Body) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

	if o.Revision != nil {
		if err := o.Revision.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "revision")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ImageConfigurationMarkEveImageLatest2Body) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageConfigurationMarkEveImageLatest2Body) UnmarshalBinary(b []byte) error {
	var res ImageConfigurationMarkEveImageLatest2Body
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
