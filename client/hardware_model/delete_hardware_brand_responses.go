package hardware_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/models"
)

// HardwareModelDeleteHardwareBrandReader is a Reader for the HardwareModelDeleteHardwareBrand structure.
type HardwareModelDeleteHardwareBrandReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareModelDeleteHardwareBrandReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHardwareModelDeleteHardwareBrandOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewHardwareModelDeleteHardwareBrandUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHardwareModelDeleteHardwareBrandForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHardwareModelDeleteHardwareBrandNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHardwareModelDeleteHardwareBrandInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewHardwareModelDeleteHardwareBrandGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHardwareModelDeleteHardwareBrandDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHardwareModelDeleteHardwareBrandOK creates a HardwareModelDeleteHardwareBrandOK with default headers values
func NewHardwareModelDeleteHardwareBrandOK() *HardwareModelDeleteHardwareBrandOK {
	return &HardwareModelDeleteHardwareBrandOK{}
}

/*
HardwareModelDeleteHardwareBrandOK describes a response with status code 200, with default header values.

A successful response.
*/
type HardwareModelDeleteHardwareBrandOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model delete hardware brand o k response has a 2xx status code
func (o *HardwareModelDeleteHardwareBrandOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hardware model delete hardware brand o k response has a 3xx status code
func (o *HardwareModelDeleteHardwareBrandOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model delete hardware brand o k response has a 4xx status code
func (o *HardwareModelDeleteHardwareBrandOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model delete hardware brand o k response has a 5xx status code
func (o *HardwareModelDeleteHardwareBrandOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model delete hardware brand o k response a status code equal to that given
func (o *HardwareModelDeleteHardwareBrandOK) IsCode(code int) bool {
	return code == 200
}

func (o *HardwareModelDeleteHardwareBrandOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/brands/id/{id}][%d] hardwareModelDeleteHardwareBrandOK  %+v", 200, o.Payload)
}

func (o *HardwareModelDeleteHardwareBrandOK) String() string {
	return fmt.Sprintf("[DELETE /v1/brands/id/{id}][%d] hardwareModelDeleteHardwareBrandOK  %+v", 200, o.Payload)
}

func (o *HardwareModelDeleteHardwareBrandOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelDeleteHardwareBrandOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelDeleteHardwareBrandUnauthorized creates a HardwareModelDeleteHardwareBrandUnauthorized with default headers values
func NewHardwareModelDeleteHardwareBrandUnauthorized() *HardwareModelDeleteHardwareBrandUnauthorized {
	return &HardwareModelDeleteHardwareBrandUnauthorized{}
}

/*
HardwareModelDeleteHardwareBrandUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type HardwareModelDeleteHardwareBrandUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model delete hardware brand unauthorized response has a 2xx status code
func (o *HardwareModelDeleteHardwareBrandUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model delete hardware brand unauthorized response has a 3xx status code
func (o *HardwareModelDeleteHardwareBrandUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model delete hardware brand unauthorized response has a 4xx status code
func (o *HardwareModelDeleteHardwareBrandUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model delete hardware brand unauthorized response has a 5xx status code
func (o *HardwareModelDeleteHardwareBrandUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model delete hardware brand unauthorized response a status code equal to that given
func (o *HardwareModelDeleteHardwareBrandUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *HardwareModelDeleteHardwareBrandUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/brands/id/{id}][%d] hardwareModelDeleteHardwareBrandUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelDeleteHardwareBrandUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/brands/id/{id}][%d] hardwareModelDeleteHardwareBrandUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelDeleteHardwareBrandUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelDeleteHardwareBrandUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelDeleteHardwareBrandForbidden creates a HardwareModelDeleteHardwareBrandForbidden with default headers values
func NewHardwareModelDeleteHardwareBrandForbidden() *HardwareModelDeleteHardwareBrandForbidden {
	return &HardwareModelDeleteHardwareBrandForbidden{}
}

/*
HardwareModelDeleteHardwareBrandForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type HardwareModelDeleteHardwareBrandForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model delete hardware brand forbidden response has a 2xx status code
func (o *HardwareModelDeleteHardwareBrandForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model delete hardware brand forbidden response has a 3xx status code
func (o *HardwareModelDeleteHardwareBrandForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model delete hardware brand forbidden response has a 4xx status code
func (o *HardwareModelDeleteHardwareBrandForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model delete hardware brand forbidden response has a 5xx status code
func (o *HardwareModelDeleteHardwareBrandForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model delete hardware brand forbidden response a status code equal to that given
func (o *HardwareModelDeleteHardwareBrandForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *HardwareModelDeleteHardwareBrandForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/brands/id/{id}][%d] hardwareModelDeleteHardwareBrandForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelDeleteHardwareBrandForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/brands/id/{id}][%d] hardwareModelDeleteHardwareBrandForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelDeleteHardwareBrandForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelDeleteHardwareBrandForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelDeleteHardwareBrandNotFound creates a HardwareModelDeleteHardwareBrandNotFound with default headers values
func NewHardwareModelDeleteHardwareBrandNotFound() *HardwareModelDeleteHardwareBrandNotFound {
	return &HardwareModelDeleteHardwareBrandNotFound{}
}

/*
HardwareModelDeleteHardwareBrandNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type HardwareModelDeleteHardwareBrandNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model delete hardware brand not found response has a 2xx status code
func (o *HardwareModelDeleteHardwareBrandNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model delete hardware brand not found response has a 3xx status code
func (o *HardwareModelDeleteHardwareBrandNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model delete hardware brand not found response has a 4xx status code
func (o *HardwareModelDeleteHardwareBrandNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model delete hardware brand not found response has a 5xx status code
func (o *HardwareModelDeleteHardwareBrandNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model delete hardware brand not found response a status code equal to that given
func (o *HardwareModelDeleteHardwareBrandNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *HardwareModelDeleteHardwareBrandNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/brands/id/{id}][%d] hardwareModelDeleteHardwareBrandNotFound  %+v", 404, o.Payload)
}

func (o *HardwareModelDeleteHardwareBrandNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/brands/id/{id}][%d] hardwareModelDeleteHardwareBrandNotFound  %+v", 404, o.Payload)
}

func (o *HardwareModelDeleteHardwareBrandNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelDeleteHardwareBrandNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelDeleteHardwareBrandInternalServerError creates a HardwareModelDeleteHardwareBrandInternalServerError with default headers values
func NewHardwareModelDeleteHardwareBrandInternalServerError() *HardwareModelDeleteHardwareBrandInternalServerError {
	return &HardwareModelDeleteHardwareBrandInternalServerError{}
}

/*
HardwareModelDeleteHardwareBrandInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type HardwareModelDeleteHardwareBrandInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model delete hardware brand internal server error response has a 2xx status code
func (o *HardwareModelDeleteHardwareBrandInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model delete hardware brand internal server error response has a 3xx status code
func (o *HardwareModelDeleteHardwareBrandInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model delete hardware brand internal server error response has a 4xx status code
func (o *HardwareModelDeleteHardwareBrandInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model delete hardware brand internal server error response has a 5xx status code
func (o *HardwareModelDeleteHardwareBrandInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model delete hardware brand internal server error response a status code equal to that given
func (o *HardwareModelDeleteHardwareBrandInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *HardwareModelDeleteHardwareBrandInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/brands/id/{id}][%d] hardwareModelDeleteHardwareBrandInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelDeleteHardwareBrandInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/brands/id/{id}][%d] hardwareModelDeleteHardwareBrandInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelDeleteHardwareBrandInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelDeleteHardwareBrandInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelDeleteHardwareBrandGatewayTimeout creates a HardwareModelDeleteHardwareBrandGatewayTimeout with default headers values
func NewHardwareModelDeleteHardwareBrandGatewayTimeout() *HardwareModelDeleteHardwareBrandGatewayTimeout {
	return &HardwareModelDeleteHardwareBrandGatewayTimeout{}
}

/*
HardwareModelDeleteHardwareBrandGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type HardwareModelDeleteHardwareBrandGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model delete hardware brand gateway timeout response has a 2xx status code
func (o *HardwareModelDeleteHardwareBrandGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model delete hardware brand gateway timeout response has a 3xx status code
func (o *HardwareModelDeleteHardwareBrandGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model delete hardware brand gateway timeout response has a 4xx status code
func (o *HardwareModelDeleteHardwareBrandGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model delete hardware brand gateway timeout response has a 5xx status code
func (o *HardwareModelDeleteHardwareBrandGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model delete hardware brand gateway timeout response a status code equal to that given
func (o *HardwareModelDeleteHardwareBrandGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *HardwareModelDeleteHardwareBrandGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/brands/id/{id}][%d] hardwareModelDeleteHardwareBrandGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelDeleteHardwareBrandGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /v1/brands/id/{id}][%d] hardwareModelDeleteHardwareBrandGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelDeleteHardwareBrandGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelDeleteHardwareBrandGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelDeleteHardwareBrandDefault creates a HardwareModelDeleteHardwareBrandDefault with default headers values
func NewHardwareModelDeleteHardwareBrandDefault(code int) *HardwareModelDeleteHardwareBrandDefault {
	return &HardwareModelDeleteHardwareBrandDefault{
		_statusCode: code,
	}
}

/*
HardwareModelDeleteHardwareBrandDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HardwareModelDeleteHardwareBrandDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the hardware model delete hardware brand default response
func (o *HardwareModelDeleteHardwareBrandDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this hardware model delete hardware brand default response has a 2xx status code
func (o *HardwareModelDeleteHardwareBrandDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this hardware model delete hardware brand default response has a 3xx status code
func (o *HardwareModelDeleteHardwareBrandDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this hardware model delete hardware brand default response has a 4xx status code
func (o *HardwareModelDeleteHardwareBrandDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this hardware model delete hardware brand default response has a 5xx status code
func (o *HardwareModelDeleteHardwareBrandDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this hardware model delete hardware brand default response a status code equal to that given
func (o *HardwareModelDeleteHardwareBrandDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *HardwareModelDeleteHardwareBrandDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/brands/id/{id}][%d] HardwareModel_DeleteHardwareBrand default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelDeleteHardwareBrandDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/brands/id/{id}][%d] HardwareModel_DeleteHardwareBrand default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelDeleteHardwareBrandDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *HardwareModelDeleteHardwareBrandDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
