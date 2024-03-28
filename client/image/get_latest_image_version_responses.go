package image

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/models"
)

// ImageConfigurationGetLatestImageVersionReader is a Reader for the ImageConfigurationGetLatestImageVersion structure.
type ImageConfigurationGetLatestImageVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageConfigurationGetLatestImageVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageConfigurationGetLatestImageVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewImageConfigurationGetLatestImageVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImageConfigurationGetLatestImageVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImageConfigurationGetLatestImageVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageConfigurationGetLatestImageVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewImageConfigurationGetLatestImageVersionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewImageConfigurationGetLatestImageVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImageConfigurationGetLatestImageVersionOK creates a ImageConfigurationGetLatestImageVersionOK with default headers values
func NewImageConfigurationGetLatestImageVersionOK() *ImageConfigurationGetLatestImageVersionOK {
	return &ImageConfigurationGetLatestImageVersionOK{}
}

/*
ImageConfigurationGetLatestImageVersionOK describes a response with status code 200, with default header values.

A successful response.
*/
type ImageConfigurationGetLatestImageVersionOK struct {
	Payload *models.Image
}

// IsSuccess returns true when this image configuration get latest image version o k response has a 2xx status code
func (o *ImageConfigurationGetLatestImageVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image configuration get latest image version o k response has a 3xx status code
func (o *ImageConfigurationGetLatestImageVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration get latest image version o k response has a 4xx status code
func (o *ImageConfigurationGetLatestImageVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration get latest image version o k response has a 5xx status code
func (o *ImageConfigurationGetLatestImageVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration get latest image version o k response a status code equal to that given
func (o *ImageConfigurationGetLatestImageVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the image configuration get latest image version o k response
func (o *ImageConfigurationGetLatestImageVersionOK) Code() int {
	return 200
}

func (o *ImageConfigurationGetLatestImageVersionOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationGetLatestImageVersionOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationGetLatestImageVersionOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationGetLatestImageVersionOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationGetLatestImageVersionOK) GetPayload() *models.Image {
	return o.Payload
}

func (o *ImageConfigurationGetLatestImageVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Image)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetLatestImageVersionUnauthorized creates a ImageConfigurationGetLatestImageVersionUnauthorized with default headers values
func NewImageConfigurationGetLatestImageVersionUnauthorized() *ImageConfigurationGetLatestImageVersionUnauthorized {
	return &ImageConfigurationGetLatestImageVersionUnauthorized{}
}

/*
ImageConfigurationGetLatestImageVersionUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ImageConfigurationGetLatestImageVersionUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration get latest image version unauthorized response has a 2xx status code
func (o *ImageConfigurationGetLatestImageVersionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration get latest image version unauthorized response has a 3xx status code
func (o *ImageConfigurationGetLatestImageVersionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration get latest image version unauthorized response has a 4xx status code
func (o *ImageConfigurationGetLatestImageVersionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration get latest image version unauthorized response has a 5xx status code
func (o *ImageConfigurationGetLatestImageVersionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration get latest image version unauthorized response a status code equal to that given
func (o *ImageConfigurationGetLatestImageVersionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the image configuration get latest image version unauthorized response
func (o *ImageConfigurationGetLatestImageVersionUnauthorized) Code() int {
	return 401
}

func (o *ImageConfigurationGetLatestImageVersionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationGetLatestImageVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationGetLatestImageVersionUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationGetLatestImageVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationGetLatestImageVersionUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetLatestImageVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetLatestImageVersionForbidden creates a ImageConfigurationGetLatestImageVersionForbidden with default headers values
func NewImageConfigurationGetLatestImageVersionForbidden() *ImageConfigurationGetLatestImageVersionForbidden {
	return &ImageConfigurationGetLatestImageVersionForbidden{}
}

/*
ImageConfigurationGetLatestImageVersionForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ImageConfigurationGetLatestImageVersionForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration get latest image version forbidden response has a 2xx status code
func (o *ImageConfigurationGetLatestImageVersionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration get latest image version forbidden response has a 3xx status code
func (o *ImageConfigurationGetLatestImageVersionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration get latest image version forbidden response has a 4xx status code
func (o *ImageConfigurationGetLatestImageVersionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration get latest image version forbidden response has a 5xx status code
func (o *ImageConfigurationGetLatestImageVersionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration get latest image version forbidden response a status code equal to that given
func (o *ImageConfigurationGetLatestImageVersionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the image configuration get latest image version forbidden response
func (o *ImageConfigurationGetLatestImageVersionForbidden) Code() int {
	return 403
}

func (o *ImageConfigurationGetLatestImageVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationGetLatestImageVersionForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationGetLatestImageVersionForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationGetLatestImageVersionForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationGetLatestImageVersionForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetLatestImageVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetLatestImageVersionNotFound creates a ImageConfigurationGetLatestImageVersionNotFound with default headers values
func NewImageConfigurationGetLatestImageVersionNotFound() *ImageConfigurationGetLatestImageVersionNotFound {
	return &ImageConfigurationGetLatestImageVersionNotFound{}
}

/*
ImageConfigurationGetLatestImageVersionNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type ImageConfigurationGetLatestImageVersionNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration get latest image version not found response has a 2xx status code
func (o *ImageConfigurationGetLatestImageVersionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration get latest image version not found response has a 3xx status code
func (o *ImageConfigurationGetLatestImageVersionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration get latest image version not found response has a 4xx status code
func (o *ImageConfigurationGetLatestImageVersionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration get latest image version not found response has a 5xx status code
func (o *ImageConfigurationGetLatestImageVersionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration get latest image version not found response a status code equal to that given
func (o *ImageConfigurationGetLatestImageVersionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the image configuration get latest image version not found response
func (o *ImageConfigurationGetLatestImageVersionNotFound) Code() int {
	return 404
}

func (o *ImageConfigurationGetLatestImageVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationGetLatestImageVersionNotFound  %+v", 404, o.Payload)
}

func (o *ImageConfigurationGetLatestImageVersionNotFound) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationGetLatestImageVersionNotFound  %+v", 404, o.Payload)
}

func (o *ImageConfigurationGetLatestImageVersionNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetLatestImageVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetLatestImageVersionInternalServerError creates a ImageConfigurationGetLatestImageVersionInternalServerError with default headers values
func NewImageConfigurationGetLatestImageVersionInternalServerError() *ImageConfigurationGetLatestImageVersionInternalServerError {
	return &ImageConfigurationGetLatestImageVersionInternalServerError{}
}

/*
ImageConfigurationGetLatestImageVersionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ImageConfigurationGetLatestImageVersionInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration get latest image version internal server error response has a 2xx status code
func (o *ImageConfigurationGetLatestImageVersionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration get latest image version internal server error response has a 3xx status code
func (o *ImageConfigurationGetLatestImageVersionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration get latest image version internal server error response has a 4xx status code
func (o *ImageConfigurationGetLatestImageVersionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration get latest image version internal server error response has a 5xx status code
func (o *ImageConfigurationGetLatestImageVersionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration get latest image version internal server error response a status code equal to that given
func (o *ImageConfigurationGetLatestImageVersionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the image configuration get latest image version internal server error response
func (o *ImageConfigurationGetLatestImageVersionInternalServerError) Code() int {
	return 500
}

func (o *ImageConfigurationGetLatestImageVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationGetLatestImageVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationGetLatestImageVersionInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationGetLatestImageVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationGetLatestImageVersionInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetLatestImageVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetLatestImageVersionGatewayTimeout creates a ImageConfigurationGetLatestImageVersionGatewayTimeout with default headers values
func NewImageConfigurationGetLatestImageVersionGatewayTimeout() *ImageConfigurationGetLatestImageVersionGatewayTimeout {
	return &ImageConfigurationGetLatestImageVersionGatewayTimeout{}
}

/*
ImageConfigurationGetLatestImageVersionGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ImageConfigurationGetLatestImageVersionGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration get latest image version gateway timeout response has a 2xx status code
func (o *ImageConfigurationGetLatestImageVersionGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration get latest image version gateway timeout response has a 3xx status code
func (o *ImageConfigurationGetLatestImageVersionGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration get latest image version gateway timeout response has a 4xx status code
func (o *ImageConfigurationGetLatestImageVersionGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration get latest image version gateway timeout response has a 5xx status code
func (o *ImageConfigurationGetLatestImageVersionGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration get latest image version gateway timeout response a status code equal to that given
func (o *ImageConfigurationGetLatestImageVersionGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the image configuration get latest image version gateway timeout response
func (o *ImageConfigurationGetLatestImageVersionGatewayTimeout) Code() int {
	return 504
}

func (o *ImageConfigurationGetLatestImageVersionGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationGetLatestImageVersionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationGetLatestImageVersionGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationGetLatestImageVersionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationGetLatestImageVersionGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetLatestImageVersionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetLatestImageVersionDefault creates a ImageConfigurationGetLatestImageVersionDefault with default headers values
func NewImageConfigurationGetLatestImageVersionDefault(code int) *ImageConfigurationGetLatestImageVersionDefault {
	return &ImageConfigurationGetLatestImageVersionDefault{
		_statusCode: code,
	}
}

/*
ImageConfigurationGetLatestImageVersionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ImageConfigurationGetLatestImageVersionDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this image configuration get latest image version default response has a 2xx status code
func (o *ImageConfigurationGetLatestImageVersionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this image configuration get latest image version default response has a 3xx status code
func (o *ImageConfigurationGetLatestImageVersionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this image configuration get latest image version default response has a 4xx status code
func (o *ImageConfigurationGetLatestImageVersionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this image configuration get latest image version default response has a 5xx status code
func (o *ImageConfigurationGetLatestImageVersionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this image configuration get latest image version default response a status code equal to that given
func (o *ImageConfigurationGetLatestImageVersionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the image configuration get latest image version default response
func (o *ImageConfigurationGetLatestImageVersionDefault) Code() int {
	return o._statusCode
}

func (o *ImageConfigurationGetLatestImageVersionDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] ImageConfiguration_GetLatestImageVersion default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationGetLatestImageVersionDefault) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] ImageConfiguration_GetLatestImageVersion default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationGetLatestImageVersionDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ImageConfigurationGetLatestImageVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
