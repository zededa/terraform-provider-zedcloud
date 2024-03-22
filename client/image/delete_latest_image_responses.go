package image

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/models"
)

// ImageConfigurationDeleteLatestImageReader is a Reader for the ImageConfigurationDeleteLatestImage structure.
type ImageConfigurationDeleteLatestImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageConfigurationDeleteLatestImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageConfigurationDeleteLatestImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewImageConfigurationDeleteLatestImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImageConfigurationDeleteLatestImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImageConfigurationDeleteLatestImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewImageConfigurationDeleteLatestImageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageConfigurationDeleteLatestImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewImageConfigurationDeleteLatestImageGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewImageConfigurationDeleteLatestImageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImageConfigurationDeleteLatestImageOK creates a ImageConfigurationDeleteLatestImageOK with default headers values
func NewImageConfigurationDeleteLatestImageOK() *ImageConfigurationDeleteLatestImageOK {
	return &ImageConfigurationDeleteLatestImageOK{}
}

/*
ImageConfigurationDeleteLatestImageOK describes a response with status code 200, with default header values.

A successful response.
*/
type ImageConfigurationDeleteLatestImageOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration delete latest image o k response has a 2xx status code
func (o *ImageConfigurationDeleteLatestImageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image configuration delete latest image o k response has a 3xx status code
func (o *ImageConfigurationDeleteLatestImageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration delete latest image o k response has a 4xx status code
func (o *ImageConfigurationDeleteLatestImageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration delete latest image o k response has a 5xx status code
func (o *ImageConfigurationDeleteLatestImageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration delete latest image o k response a status code equal to that given
func (o *ImageConfigurationDeleteLatestImageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the image configuration delete latest image o k response
func (o *ImageConfigurationDeleteLatestImageOK) Code() int {
	return 200
}

func (o *ImageConfigurationDeleteLatestImageOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/baseos/latest][%d] imageConfigurationDeleteLatestImageOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationDeleteLatestImageOK) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/baseos/latest][%d] imageConfigurationDeleteLatestImageOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationDeleteLatestImageOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationDeleteLatestImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationDeleteLatestImageUnauthorized creates a ImageConfigurationDeleteLatestImageUnauthorized with default headers values
func NewImageConfigurationDeleteLatestImageUnauthorized() *ImageConfigurationDeleteLatestImageUnauthorized {
	return &ImageConfigurationDeleteLatestImageUnauthorized{}
}

/*
ImageConfigurationDeleteLatestImageUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ImageConfigurationDeleteLatestImageUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration delete latest image unauthorized response has a 2xx status code
func (o *ImageConfigurationDeleteLatestImageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration delete latest image unauthorized response has a 3xx status code
func (o *ImageConfigurationDeleteLatestImageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration delete latest image unauthorized response has a 4xx status code
func (o *ImageConfigurationDeleteLatestImageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration delete latest image unauthorized response has a 5xx status code
func (o *ImageConfigurationDeleteLatestImageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration delete latest image unauthorized response a status code equal to that given
func (o *ImageConfigurationDeleteLatestImageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the image configuration delete latest image unauthorized response
func (o *ImageConfigurationDeleteLatestImageUnauthorized) Code() int {
	return 401
}

func (o *ImageConfigurationDeleteLatestImageUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/baseos/latest][%d] imageConfigurationDeleteLatestImageUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationDeleteLatestImageUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/baseos/latest][%d] imageConfigurationDeleteLatestImageUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationDeleteLatestImageUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationDeleteLatestImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationDeleteLatestImageForbidden creates a ImageConfigurationDeleteLatestImageForbidden with default headers values
func NewImageConfigurationDeleteLatestImageForbidden() *ImageConfigurationDeleteLatestImageForbidden {
	return &ImageConfigurationDeleteLatestImageForbidden{}
}

/*
ImageConfigurationDeleteLatestImageForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ImageConfigurationDeleteLatestImageForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration delete latest image forbidden response has a 2xx status code
func (o *ImageConfigurationDeleteLatestImageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration delete latest image forbidden response has a 3xx status code
func (o *ImageConfigurationDeleteLatestImageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration delete latest image forbidden response has a 4xx status code
func (o *ImageConfigurationDeleteLatestImageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration delete latest image forbidden response has a 5xx status code
func (o *ImageConfigurationDeleteLatestImageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration delete latest image forbidden response a status code equal to that given
func (o *ImageConfigurationDeleteLatestImageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the image configuration delete latest image forbidden response
func (o *ImageConfigurationDeleteLatestImageForbidden) Code() int {
	return 403
}

func (o *ImageConfigurationDeleteLatestImageForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/baseos/latest][%d] imageConfigurationDeleteLatestImageForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationDeleteLatestImageForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/baseos/latest][%d] imageConfigurationDeleteLatestImageForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationDeleteLatestImageForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationDeleteLatestImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationDeleteLatestImageNotFound creates a ImageConfigurationDeleteLatestImageNotFound with default headers values
func NewImageConfigurationDeleteLatestImageNotFound() *ImageConfigurationDeleteLatestImageNotFound {
	return &ImageConfigurationDeleteLatestImageNotFound{}
}

/*
ImageConfigurationDeleteLatestImageNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type ImageConfigurationDeleteLatestImageNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration delete latest image not found response has a 2xx status code
func (o *ImageConfigurationDeleteLatestImageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration delete latest image not found response has a 3xx status code
func (o *ImageConfigurationDeleteLatestImageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration delete latest image not found response has a 4xx status code
func (o *ImageConfigurationDeleteLatestImageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration delete latest image not found response has a 5xx status code
func (o *ImageConfigurationDeleteLatestImageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration delete latest image not found response a status code equal to that given
func (o *ImageConfigurationDeleteLatestImageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the image configuration delete latest image not found response
func (o *ImageConfigurationDeleteLatestImageNotFound) Code() int {
	return 404
}

func (o *ImageConfigurationDeleteLatestImageNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/baseos/latest][%d] imageConfigurationDeleteLatestImageNotFound  %+v", 404, o.Payload)
}

func (o *ImageConfigurationDeleteLatestImageNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/baseos/latest][%d] imageConfigurationDeleteLatestImageNotFound  %+v", 404, o.Payload)
}

func (o *ImageConfigurationDeleteLatestImageNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationDeleteLatestImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationDeleteLatestImageConflict creates a ImageConfigurationDeleteLatestImageConflict with default headers values
func NewImageConfigurationDeleteLatestImageConflict() *ImageConfigurationDeleteLatestImageConflict {
	return &ImageConfigurationDeleteLatestImageConflict{}
}

/*
ImageConfigurationDeleteLatestImageConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because there are edge application bundles using this edge application image
*/
type ImageConfigurationDeleteLatestImageConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration delete latest image conflict response has a 2xx status code
func (o *ImageConfigurationDeleteLatestImageConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration delete latest image conflict response has a 3xx status code
func (o *ImageConfigurationDeleteLatestImageConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration delete latest image conflict response has a 4xx status code
func (o *ImageConfigurationDeleteLatestImageConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration delete latest image conflict response has a 5xx status code
func (o *ImageConfigurationDeleteLatestImageConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration delete latest image conflict response a status code equal to that given
func (o *ImageConfigurationDeleteLatestImageConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the image configuration delete latest image conflict response
func (o *ImageConfigurationDeleteLatestImageConflict) Code() int {
	return 409
}

func (o *ImageConfigurationDeleteLatestImageConflict) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/baseos/latest][%d] imageConfigurationDeleteLatestImageConflict  %+v", 409, o.Payload)
}

func (o *ImageConfigurationDeleteLatestImageConflict) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/baseos/latest][%d] imageConfigurationDeleteLatestImageConflict  %+v", 409, o.Payload)
}

func (o *ImageConfigurationDeleteLatestImageConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationDeleteLatestImageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationDeleteLatestImageInternalServerError creates a ImageConfigurationDeleteLatestImageInternalServerError with default headers values
func NewImageConfigurationDeleteLatestImageInternalServerError() *ImageConfigurationDeleteLatestImageInternalServerError {
	return &ImageConfigurationDeleteLatestImageInternalServerError{}
}

/*
ImageConfigurationDeleteLatestImageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ImageConfigurationDeleteLatestImageInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration delete latest image internal server error response has a 2xx status code
func (o *ImageConfigurationDeleteLatestImageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration delete latest image internal server error response has a 3xx status code
func (o *ImageConfigurationDeleteLatestImageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration delete latest image internal server error response has a 4xx status code
func (o *ImageConfigurationDeleteLatestImageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration delete latest image internal server error response has a 5xx status code
func (o *ImageConfigurationDeleteLatestImageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration delete latest image internal server error response a status code equal to that given
func (o *ImageConfigurationDeleteLatestImageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the image configuration delete latest image internal server error response
func (o *ImageConfigurationDeleteLatestImageInternalServerError) Code() int {
	return 500
}

func (o *ImageConfigurationDeleteLatestImageInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/baseos/latest][%d] imageConfigurationDeleteLatestImageInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationDeleteLatestImageInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/baseos/latest][%d] imageConfigurationDeleteLatestImageInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationDeleteLatestImageInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationDeleteLatestImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationDeleteLatestImageGatewayTimeout creates a ImageConfigurationDeleteLatestImageGatewayTimeout with default headers values
func NewImageConfigurationDeleteLatestImageGatewayTimeout() *ImageConfigurationDeleteLatestImageGatewayTimeout {
	return &ImageConfigurationDeleteLatestImageGatewayTimeout{}
}

/*
ImageConfigurationDeleteLatestImageGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ImageConfigurationDeleteLatestImageGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration delete latest image gateway timeout response has a 2xx status code
func (o *ImageConfigurationDeleteLatestImageGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration delete latest image gateway timeout response has a 3xx status code
func (o *ImageConfigurationDeleteLatestImageGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration delete latest image gateway timeout response has a 4xx status code
func (o *ImageConfigurationDeleteLatestImageGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration delete latest image gateway timeout response has a 5xx status code
func (o *ImageConfigurationDeleteLatestImageGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration delete latest image gateway timeout response a status code equal to that given
func (o *ImageConfigurationDeleteLatestImageGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the image configuration delete latest image gateway timeout response
func (o *ImageConfigurationDeleteLatestImageGatewayTimeout) Code() int {
	return 504
}

func (o *ImageConfigurationDeleteLatestImageGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/baseos/latest][%d] imageConfigurationDeleteLatestImageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationDeleteLatestImageGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/baseos/latest][%d] imageConfigurationDeleteLatestImageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationDeleteLatestImageGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationDeleteLatestImageGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationDeleteLatestImageDefault creates a ImageConfigurationDeleteLatestImageDefault with default headers values
func NewImageConfigurationDeleteLatestImageDefault(code int) *ImageConfigurationDeleteLatestImageDefault {
	return &ImageConfigurationDeleteLatestImageDefault{
		_statusCode: code,
	}
}

/*
ImageConfigurationDeleteLatestImageDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ImageConfigurationDeleteLatestImageDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this image configuration delete latest image default response has a 2xx status code
func (o *ImageConfigurationDeleteLatestImageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this image configuration delete latest image default response has a 3xx status code
func (o *ImageConfigurationDeleteLatestImageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this image configuration delete latest image default response has a 4xx status code
func (o *ImageConfigurationDeleteLatestImageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this image configuration delete latest image default response has a 5xx status code
func (o *ImageConfigurationDeleteLatestImageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this image configuration delete latest image default response a status code equal to that given
func (o *ImageConfigurationDeleteLatestImageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the image configuration delete latest image default response
func (o *ImageConfigurationDeleteLatestImageDefault) Code() int {
	return o._statusCode
}

func (o *ImageConfigurationDeleteLatestImageDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/baseos/latest][%d] ImageConfiguration_DeleteLatestImage default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationDeleteLatestImageDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/baseos/latest][%d] ImageConfiguration_DeleteLatestImage default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationDeleteLatestImageDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ImageConfigurationDeleteLatestImageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
