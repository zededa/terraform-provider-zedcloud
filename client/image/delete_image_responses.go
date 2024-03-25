package image

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/models"
)

// ImageConfigurationDeleteImageReader is a Reader for the ImageConfigurationDeleteImage structure.
type ImageConfigurationDeleteImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageConfigurationDeleteImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageConfigurationDeleteImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewImageConfigurationDeleteImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImageConfigurationDeleteImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImageConfigurationDeleteImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewImageConfigurationDeleteImageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageConfigurationDeleteImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewImageConfigurationDeleteImageGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewImageConfigurationDeleteImageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImageConfigurationDeleteImageOK creates a ImageConfigurationDeleteImageOK with default headers values
func NewImageConfigurationDeleteImageOK() *ImageConfigurationDeleteImageOK {
	return &ImageConfigurationDeleteImageOK{}
}

/*
ImageConfigurationDeleteImageOK describes a response with status code 200, with default header values.

A successful response.
*/
type ImageConfigurationDeleteImageOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration delete image o k response has a 2xx status code
func (o *ImageConfigurationDeleteImageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image configuration delete image o k response has a 3xx status code
func (o *ImageConfigurationDeleteImageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration delete image o k response has a 4xx status code
func (o *ImageConfigurationDeleteImageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration delete image o k response has a 5xx status code
func (o *ImageConfigurationDeleteImageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration delete image o k response a status code equal to that given
func (o *ImageConfigurationDeleteImageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the image configuration delete image o k response
func (o *ImageConfigurationDeleteImageOK) Code() int {
	return 200
}

func (o *ImageConfigurationDeleteImageOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] imageConfigurationDeleteImageOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationDeleteImageOK) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] imageConfigurationDeleteImageOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationDeleteImageOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationDeleteImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationDeleteImageUnauthorized creates a ImageConfigurationDeleteImageUnauthorized with default headers values
func NewImageConfigurationDeleteImageUnauthorized() *ImageConfigurationDeleteImageUnauthorized {
	return &ImageConfigurationDeleteImageUnauthorized{}
}

/*
ImageConfigurationDeleteImageUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ImageConfigurationDeleteImageUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration delete image unauthorized response has a 2xx status code
func (o *ImageConfigurationDeleteImageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration delete image unauthorized response has a 3xx status code
func (o *ImageConfigurationDeleteImageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration delete image unauthorized response has a 4xx status code
func (o *ImageConfigurationDeleteImageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration delete image unauthorized response has a 5xx status code
func (o *ImageConfigurationDeleteImageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration delete image unauthorized response a status code equal to that given
func (o *ImageConfigurationDeleteImageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the image configuration delete image unauthorized response
func (o *ImageConfigurationDeleteImageUnauthorized) Code() int {
	return 401
}

func (o *ImageConfigurationDeleteImageUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] imageConfigurationDeleteImageUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationDeleteImageUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] imageConfigurationDeleteImageUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationDeleteImageUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationDeleteImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationDeleteImageForbidden creates a ImageConfigurationDeleteImageForbidden with default headers values
func NewImageConfigurationDeleteImageForbidden() *ImageConfigurationDeleteImageForbidden {
	return &ImageConfigurationDeleteImageForbidden{}
}

/*
ImageConfigurationDeleteImageForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ImageConfigurationDeleteImageForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration delete image forbidden response has a 2xx status code
func (o *ImageConfigurationDeleteImageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration delete image forbidden response has a 3xx status code
func (o *ImageConfigurationDeleteImageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration delete image forbidden response has a 4xx status code
func (o *ImageConfigurationDeleteImageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration delete image forbidden response has a 5xx status code
func (o *ImageConfigurationDeleteImageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration delete image forbidden response a status code equal to that given
func (o *ImageConfigurationDeleteImageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the image configuration delete image forbidden response
func (o *ImageConfigurationDeleteImageForbidden) Code() int {
	return 403
}

func (o *ImageConfigurationDeleteImageForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] imageConfigurationDeleteImageForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationDeleteImageForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] imageConfigurationDeleteImageForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationDeleteImageForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationDeleteImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationDeleteImageNotFound creates a ImageConfigurationDeleteImageNotFound with default headers values
func NewImageConfigurationDeleteImageNotFound() *ImageConfigurationDeleteImageNotFound {
	return &ImageConfigurationDeleteImageNotFound{}
}

/*
ImageConfigurationDeleteImageNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type ImageConfigurationDeleteImageNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration delete image not found response has a 2xx status code
func (o *ImageConfigurationDeleteImageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration delete image not found response has a 3xx status code
func (o *ImageConfigurationDeleteImageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration delete image not found response has a 4xx status code
func (o *ImageConfigurationDeleteImageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration delete image not found response has a 5xx status code
func (o *ImageConfigurationDeleteImageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration delete image not found response a status code equal to that given
func (o *ImageConfigurationDeleteImageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the image configuration delete image not found response
func (o *ImageConfigurationDeleteImageNotFound) Code() int {
	return 404
}

func (o *ImageConfigurationDeleteImageNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] imageConfigurationDeleteImageNotFound  %+v", 404, o.Payload)
}

func (o *ImageConfigurationDeleteImageNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] imageConfigurationDeleteImageNotFound  %+v", 404, o.Payload)
}

func (o *ImageConfigurationDeleteImageNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationDeleteImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationDeleteImageConflict creates a ImageConfigurationDeleteImageConflict with default headers values
func NewImageConfigurationDeleteImageConflict() *ImageConfigurationDeleteImageConflict {
	return &ImageConfigurationDeleteImageConflict{}
}

/*
ImageConfigurationDeleteImageConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because there are edge application bundles using this edge application image
*/
type ImageConfigurationDeleteImageConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration delete image conflict response has a 2xx status code
func (o *ImageConfigurationDeleteImageConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration delete image conflict response has a 3xx status code
func (o *ImageConfigurationDeleteImageConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration delete image conflict response has a 4xx status code
func (o *ImageConfigurationDeleteImageConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration delete image conflict response has a 5xx status code
func (o *ImageConfigurationDeleteImageConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration delete image conflict response a status code equal to that given
func (o *ImageConfigurationDeleteImageConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the image configuration delete image conflict response
func (o *ImageConfigurationDeleteImageConflict) Code() int {
	return 409
}

func (o *ImageConfigurationDeleteImageConflict) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] imageConfigurationDeleteImageConflict  %+v", 409, o.Payload)
}

func (o *ImageConfigurationDeleteImageConflict) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] imageConfigurationDeleteImageConflict  %+v", 409, o.Payload)
}

func (o *ImageConfigurationDeleteImageConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationDeleteImageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationDeleteImageInternalServerError creates a ImageConfigurationDeleteImageInternalServerError with default headers values
func NewImageConfigurationDeleteImageInternalServerError() *ImageConfigurationDeleteImageInternalServerError {
	return &ImageConfigurationDeleteImageInternalServerError{}
}

/*
ImageConfigurationDeleteImageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ImageConfigurationDeleteImageInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration delete image internal server error response has a 2xx status code
func (o *ImageConfigurationDeleteImageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration delete image internal server error response has a 3xx status code
func (o *ImageConfigurationDeleteImageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration delete image internal server error response has a 4xx status code
func (o *ImageConfigurationDeleteImageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration delete image internal server error response has a 5xx status code
func (o *ImageConfigurationDeleteImageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration delete image internal server error response a status code equal to that given
func (o *ImageConfigurationDeleteImageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the image configuration delete image internal server error response
func (o *ImageConfigurationDeleteImageInternalServerError) Code() int {
	return 500
}

func (o *ImageConfigurationDeleteImageInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] imageConfigurationDeleteImageInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationDeleteImageInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] imageConfigurationDeleteImageInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationDeleteImageInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationDeleteImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationDeleteImageGatewayTimeout creates a ImageConfigurationDeleteImageGatewayTimeout with default headers values
func NewImageConfigurationDeleteImageGatewayTimeout() *ImageConfigurationDeleteImageGatewayTimeout {
	return &ImageConfigurationDeleteImageGatewayTimeout{}
}

/*
ImageConfigurationDeleteImageGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ImageConfigurationDeleteImageGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration delete image gateway timeout response has a 2xx status code
func (o *ImageConfigurationDeleteImageGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration delete image gateway timeout response has a 3xx status code
func (o *ImageConfigurationDeleteImageGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration delete image gateway timeout response has a 4xx status code
func (o *ImageConfigurationDeleteImageGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration delete image gateway timeout response has a 5xx status code
func (o *ImageConfigurationDeleteImageGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration delete image gateway timeout response a status code equal to that given
func (o *ImageConfigurationDeleteImageGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the image configuration delete image gateway timeout response
func (o *ImageConfigurationDeleteImageGatewayTimeout) Code() int {
	return 504
}

func (o *ImageConfigurationDeleteImageGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] imageConfigurationDeleteImageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationDeleteImageGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] imageConfigurationDeleteImageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationDeleteImageGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationDeleteImageGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationDeleteImageDefault creates a ImageConfigurationDeleteImageDefault with default headers values
func NewImageConfigurationDeleteImageDefault(code int) *ImageConfigurationDeleteImageDefault {
	return &ImageConfigurationDeleteImageDefault{
		_statusCode: code,
	}
}

/*
ImageConfigurationDeleteImageDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ImageConfigurationDeleteImageDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this image configuration delete image default response has a 2xx status code
func (o *ImageConfigurationDeleteImageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this image configuration delete image default response has a 3xx status code
func (o *ImageConfigurationDeleteImageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this image configuration delete image default response has a 4xx status code
func (o *ImageConfigurationDeleteImageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this image configuration delete image default response has a 5xx status code
func (o *ImageConfigurationDeleteImageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this image configuration delete image default response a status code equal to that given
func (o *ImageConfigurationDeleteImageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the image configuration delete image default response
func (o *ImageConfigurationDeleteImageDefault) Code() int {
	return o._statusCode
}

func (o *ImageConfigurationDeleteImageDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] ImageConfiguration_DeleteImage default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationDeleteImageDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/images/id/{id}][%d] ImageConfiguration_DeleteImage default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationDeleteImageDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ImageConfigurationDeleteImageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
