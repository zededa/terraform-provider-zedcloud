package image_configuration

import (
	"fmt"
	"io"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// ImageConfigurationGetImageByNameReader is a Reader for the ImageConfigurationGetImageByName structure.
type ImageConfigurationGetImageByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageConfigurationGetImageByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageConfigurationGetImageByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewImageConfigurationGetImageByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImageConfigurationGetImageByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImageConfigurationGetImageByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageConfigurationGetImageByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewImageConfigurationGetImageByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewImageConfigurationGetImageByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImageConfigurationGetImageByNameOK creates a ImageConfigurationGetImageByNameOK with default headers values
func NewImageConfigurationGetImageByNameOK() *ImageConfigurationGetImageByNameOK {
	return &ImageConfigurationGetImageByNameOK{}
}

/*
ImageConfigurationGetImageByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type ImageConfigurationGetImageByNameOK struct {
	Payload *models.Image
}

// IsSuccess returns true when this image configuration get image by name o k response has a 2xx status code
func (o *ImageConfigurationGetImageByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image configuration get image by name o k response has a 3xx status code
func (o *ImageConfigurationGetImageByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration get image by name o k response has a 4xx status code
func (o *ImageConfigurationGetImageByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration get image by name o k response has a 5xx status code
func (o *ImageConfigurationGetImageByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration get image by name o k response a status code equal to that given
func (o *ImageConfigurationGetImageByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the image configuration get image by name o k response
func (o *ImageConfigurationGetImageByNameOK) Code() int {
	return 200
}

func (o *ImageConfigurationGetImageByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/name/{name}][%d] imageConfigurationGetImageByNameOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationGetImageByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/name/{name}][%d] imageConfigurationGetImageByNameOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationGetImageByNameOK) GetPayload() *models.Image {
	return o.Payload
}

func (o *ImageConfigurationGetImageByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Image)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetImageByNameUnauthorized creates a ImageConfigurationGetImageByNameUnauthorized with default headers values
func NewImageConfigurationGetImageByNameUnauthorized() *ImageConfigurationGetImageByNameUnauthorized {
	return &ImageConfigurationGetImageByNameUnauthorized{}
}

/*
ImageConfigurationGetImageByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ImageConfigurationGetImageByNameUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration get image by name unauthorized response has a 2xx status code
func (o *ImageConfigurationGetImageByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration get image by name unauthorized response has a 3xx status code
func (o *ImageConfigurationGetImageByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration get image by name unauthorized response has a 4xx status code
func (o *ImageConfigurationGetImageByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration get image by name unauthorized response has a 5xx status code
func (o *ImageConfigurationGetImageByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration get image by name unauthorized response a status code equal to that given
func (o *ImageConfigurationGetImageByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the image configuration get image by name unauthorized response
func (o *ImageConfigurationGetImageByNameUnauthorized) Code() int {
	return 401
}

func (o *ImageConfigurationGetImageByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/name/{name}][%d] imageConfigurationGetImageByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationGetImageByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/name/{name}][%d] imageConfigurationGetImageByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationGetImageByNameUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetImageByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetImageByNameForbidden creates a ImageConfigurationGetImageByNameForbidden with default headers values
func NewImageConfigurationGetImageByNameForbidden() *ImageConfigurationGetImageByNameForbidden {
	return &ImageConfigurationGetImageByNameForbidden{}
}

/*
ImageConfigurationGetImageByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ImageConfigurationGetImageByNameForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration get image by name forbidden response has a 2xx status code
func (o *ImageConfigurationGetImageByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration get image by name forbidden response has a 3xx status code
func (o *ImageConfigurationGetImageByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration get image by name forbidden response has a 4xx status code
func (o *ImageConfigurationGetImageByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration get image by name forbidden response has a 5xx status code
func (o *ImageConfigurationGetImageByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration get image by name forbidden response a status code equal to that given
func (o *ImageConfigurationGetImageByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the image configuration get image by name forbidden response
func (o *ImageConfigurationGetImageByNameForbidden) Code() int {
	return 403
}

func (o *ImageConfigurationGetImageByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/name/{name}][%d] imageConfigurationGetImageByNameForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationGetImageByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/name/{name}][%d] imageConfigurationGetImageByNameForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationGetImageByNameForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetImageByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetImageByNameNotFound creates a ImageConfigurationGetImageByNameNotFound with default headers values
func NewImageConfigurationGetImageByNameNotFound() *ImageConfigurationGetImageByNameNotFound {
	return &ImageConfigurationGetImageByNameNotFound{}
}

/*
ImageConfigurationGetImageByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type ImageConfigurationGetImageByNameNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration get image by name not found response has a 2xx status code
func (o *ImageConfigurationGetImageByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration get image by name not found response has a 3xx status code
func (o *ImageConfigurationGetImageByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration get image by name not found response has a 4xx status code
func (o *ImageConfigurationGetImageByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration get image by name not found response has a 5xx status code
func (o *ImageConfigurationGetImageByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration get image by name not found response a status code equal to that given
func (o *ImageConfigurationGetImageByNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the image configuration get image by name not found response
func (o *ImageConfigurationGetImageByNameNotFound) Code() int {
	return 404
}

func (o *ImageConfigurationGetImageByNameNotFound) Error() string {
	return fmt.Sprintf("GET /v1/apps/images/name/{name} => Not Found: %d\n%+v", 404, spew.Sdump(o.Payload.Error))
}

func (o *ImageConfigurationGetImageByNameNotFound) String() string {
	return fmt.Sprintf("GET /v1/apps/images/name/{name} => Not Found: %d\n%+v", 404, spew.Sdump(o.Payload.Error))
}

func (o *ImageConfigurationGetImageByNameNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetImageByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetImageByNameInternalServerError creates a ImageConfigurationGetImageByNameInternalServerError with default headers values
func NewImageConfigurationGetImageByNameInternalServerError() *ImageConfigurationGetImageByNameInternalServerError {
	return &ImageConfigurationGetImageByNameInternalServerError{}
}

/*
ImageConfigurationGetImageByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ImageConfigurationGetImageByNameInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration get image by name internal server error response has a 2xx status code
func (o *ImageConfigurationGetImageByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration get image by name internal server error response has a 3xx status code
func (o *ImageConfigurationGetImageByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration get image by name internal server error response has a 4xx status code
func (o *ImageConfigurationGetImageByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration get image by name internal server error response has a 5xx status code
func (o *ImageConfigurationGetImageByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration get image by name internal server error response a status code equal to that given
func (o *ImageConfigurationGetImageByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the image configuration get image by name internal server error response
func (o *ImageConfigurationGetImageByNameInternalServerError) Code() int {
	return 500
}

func (o *ImageConfigurationGetImageByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/name/{name}][%d] imageConfigurationGetImageByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationGetImageByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/name/{name}][%d] imageConfigurationGetImageByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationGetImageByNameInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetImageByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetImageByNameGatewayTimeout creates a ImageConfigurationGetImageByNameGatewayTimeout with default headers values
func NewImageConfigurationGetImageByNameGatewayTimeout() *ImageConfigurationGetImageByNameGatewayTimeout {
	return &ImageConfigurationGetImageByNameGatewayTimeout{}
}

/*
ImageConfigurationGetImageByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ImageConfigurationGetImageByNameGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration get image by name gateway timeout response has a 2xx status code
func (o *ImageConfigurationGetImageByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration get image by name gateway timeout response has a 3xx status code
func (o *ImageConfigurationGetImageByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration get image by name gateway timeout response has a 4xx status code
func (o *ImageConfigurationGetImageByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration get image by name gateway timeout response has a 5xx status code
func (o *ImageConfigurationGetImageByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration get image by name gateway timeout response a status code equal to that given
func (o *ImageConfigurationGetImageByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the image configuration get image by name gateway timeout response
func (o *ImageConfigurationGetImageByNameGatewayTimeout) Code() int {
	return 504
}

func (o *ImageConfigurationGetImageByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/name/{name}][%d] imageConfigurationGetImageByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationGetImageByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/name/{name}][%d] imageConfigurationGetImageByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationGetImageByNameGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetImageByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetImageByNameDefault creates a ImageConfigurationGetImageByNameDefault with default headers values
func NewImageConfigurationGetImageByNameDefault(code int) *ImageConfigurationGetImageByNameDefault {
	return &ImageConfigurationGetImageByNameDefault{
		_statusCode: code,
	}
}

/*
ImageConfigurationGetImageByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ImageConfigurationGetImageByNameDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this image configuration get image by name default response has a 2xx status code
func (o *ImageConfigurationGetImageByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this image configuration get image by name default response has a 3xx status code
func (o *ImageConfigurationGetImageByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this image configuration get image by name default response has a 4xx status code
func (o *ImageConfigurationGetImageByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this image configuration get image by name default response has a 5xx status code
func (o *ImageConfigurationGetImageByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this image configuration get image by name default response a status code equal to that given
func (o *ImageConfigurationGetImageByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the image configuration get image by name default response
func (o *ImageConfigurationGetImageByNameDefault) Code() int {
	return o._statusCode
}

func (o *ImageConfigurationGetImageByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/name/{name}][%d] ImageConfiguration_GetImageByName default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationGetImageByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/name/{name}][%d] ImageConfiguration_GetImageByName default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationGetImageByNameDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ImageConfigurationGetImageByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
