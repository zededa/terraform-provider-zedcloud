// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// IdentityAccessManagementGetEnterpriseSelfReader is a Reader for the IdentityAccessManagementGetEnterpriseSelf structure.
type IdentityAccessManagementGetEnterpriseSelfReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementGetEnterpriseSelfReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementGetEnterpriseSelfOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementGetEnterpriseSelfUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementGetEnterpriseSelfForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIdentityAccessManagementGetEnterpriseSelfNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementGetEnterpriseSelfInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementGetEnterpriseSelfGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementGetEnterpriseSelfDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementGetEnterpriseSelfOK creates a IdentityAccessManagementGetEnterpriseSelfOK with default headers values
func NewIdentityAccessManagementGetEnterpriseSelfOK() *IdentityAccessManagementGetEnterpriseSelfOK {
	return &IdentityAccessManagementGetEnterpriseSelfOK{}
}

/*
IdentityAccessManagementGetEnterpriseSelfOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementGetEnterpriseSelfOK struct {
	Payload *models.CrudResponseRead
}

// IsSuccess returns true when this identity access management get enterprise self o k response has a 2xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management get enterprise self o k response has a 3xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get enterprise self o k response has a 4xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get enterprise self o k response has a 5xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get enterprise self o k response a status code equal to that given
func (o *IdentityAccessManagementGetEnterpriseSelfOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the identity access management get enterprise self o k response
func (o *IdentityAccessManagementGetEnterpriseSelfOK) Code() int {
	return 200
}

func (o *IdentityAccessManagementGetEnterpriseSelfOK) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises/self][%d] identityAccessManagementGetEnterpriseSelfOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseSelfOK) String() string {
	return fmt.Sprintf("[GET /v1/enterprises/self][%d] identityAccessManagementGetEnterpriseSelfOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseSelfOK) GetPayload() *models.CrudResponseRead {
	return o.Payload
}

func (o *IdentityAccessManagementGetEnterpriseSelfOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CrudResponseRead)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetEnterpriseSelfUnauthorized creates a IdentityAccessManagementGetEnterpriseSelfUnauthorized with default headers values
func NewIdentityAccessManagementGetEnterpriseSelfUnauthorized() *IdentityAccessManagementGetEnterpriseSelfUnauthorized {
	return &IdentityAccessManagementGetEnterpriseSelfUnauthorized{}
}

/*
IdentityAccessManagementGetEnterpriseSelfUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementGetEnterpriseSelfUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get enterprise self unauthorized response has a 2xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get enterprise self unauthorized response has a 3xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get enterprise self unauthorized response has a 4xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get enterprise self unauthorized response has a 5xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get enterprise self unauthorized response a status code equal to that given
func (o *IdentityAccessManagementGetEnterpriseSelfUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the identity access management get enterprise self unauthorized response
func (o *IdentityAccessManagementGetEnterpriseSelfUnauthorized) Code() int {
	return 401
}

func (o *IdentityAccessManagementGetEnterpriseSelfUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises/self][%d] identityAccessManagementGetEnterpriseSelfUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseSelfUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/enterprises/self][%d] identityAccessManagementGetEnterpriseSelfUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseSelfUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetEnterpriseSelfUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetEnterpriseSelfForbidden creates a IdentityAccessManagementGetEnterpriseSelfForbidden with default headers values
func NewIdentityAccessManagementGetEnterpriseSelfForbidden() *IdentityAccessManagementGetEnterpriseSelfForbidden {
	return &IdentityAccessManagementGetEnterpriseSelfForbidden{}
}

/*
IdentityAccessManagementGetEnterpriseSelfForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementGetEnterpriseSelfForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get enterprise self forbidden response has a 2xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get enterprise self forbidden response has a 3xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get enterprise self forbidden response has a 4xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get enterprise self forbidden response has a 5xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get enterprise self forbidden response a status code equal to that given
func (o *IdentityAccessManagementGetEnterpriseSelfForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the identity access management get enterprise self forbidden response
func (o *IdentityAccessManagementGetEnterpriseSelfForbidden) Code() int {
	return 403
}

func (o *IdentityAccessManagementGetEnterpriseSelfForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises/self][%d] identityAccessManagementGetEnterpriseSelfForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseSelfForbidden) String() string {
	return fmt.Sprintf("[GET /v1/enterprises/self][%d] identityAccessManagementGetEnterpriseSelfForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseSelfForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetEnterpriseSelfForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetEnterpriseSelfNotFound creates a IdentityAccessManagementGetEnterpriseSelfNotFound with default headers values
func NewIdentityAccessManagementGetEnterpriseSelfNotFound() *IdentityAccessManagementGetEnterpriseSelfNotFound {
	return &IdentityAccessManagementGetEnterpriseSelfNotFound{}
}

/*
IdentityAccessManagementGetEnterpriseSelfNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type IdentityAccessManagementGetEnterpriseSelfNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get enterprise self not found response has a 2xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get enterprise self not found response has a 3xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get enterprise self not found response has a 4xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get enterprise self not found response has a 5xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get enterprise self not found response a status code equal to that given
func (o *IdentityAccessManagementGetEnterpriseSelfNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the identity access management get enterprise self not found response
func (o *IdentityAccessManagementGetEnterpriseSelfNotFound) Code() int {
	return 404
}

func (o *IdentityAccessManagementGetEnterpriseSelfNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises/self][%d] identityAccessManagementGetEnterpriseSelfNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseSelfNotFound) String() string {
	return fmt.Sprintf("[GET /v1/enterprises/self][%d] identityAccessManagementGetEnterpriseSelfNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseSelfNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetEnterpriseSelfNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetEnterpriseSelfInternalServerError creates a IdentityAccessManagementGetEnterpriseSelfInternalServerError with default headers values
func NewIdentityAccessManagementGetEnterpriseSelfInternalServerError() *IdentityAccessManagementGetEnterpriseSelfInternalServerError {
	return &IdentityAccessManagementGetEnterpriseSelfInternalServerError{}
}

/*
IdentityAccessManagementGetEnterpriseSelfInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementGetEnterpriseSelfInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get enterprise self internal server error response has a 2xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get enterprise self internal server error response has a 3xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get enterprise self internal server error response has a 4xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get enterprise self internal server error response has a 5xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management get enterprise self internal server error response a status code equal to that given
func (o *IdentityAccessManagementGetEnterpriseSelfInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the identity access management get enterprise self internal server error response
func (o *IdentityAccessManagementGetEnterpriseSelfInternalServerError) Code() int {
	return 500
}

func (o *IdentityAccessManagementGetEnterpriseSelfInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises/self][%d] identityAccessManagementGetEnterpriseSelfInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseSelfInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/enterprises/self][%d] identityAccessManagementGetEnterpriseSelfInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseSelfInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetEnterpriseSelfInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetEnterpriseSelfGatewayTimeout creates a IdentityAccessManagementGetEnterpriseSelfGatewayTimeout with default headers values
func NewIdentityAccessManagementGetEnterpriseSelfGatewayTimeout() *IdentityAccessManagementGetEnterpriseSelfGatewayTimeout {
	return &IdentityAccessManagementGetEnterpriseSelfGatewayTimeout{}
}

/*
IdentityAccessManagementGetEnterpriseSelfGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementGetEnterpriseSelfGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get enterprise self gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get enterprise self gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get enterprise self gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get enterprise self gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management get enterprise self gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementGetEnterpriseSelfGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the identity access management get enterprise self gateway timeout response
func (o *IdentityAccessManagementGetEnterpriseSelfGatewayTimeout) Code() int {
	return 504
}

func (o *IdentityAccessManagementGetEnterpriseSelfGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises/self][%d] identityAccessManagementGetEnterpriseSelfGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseSelfGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/enterprises/self][%d] identityAccessManagementGetEnterpriseSelfGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseSelfGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetEnterpriseSelfGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetEnterpriseSelfDefault creates a IdentityAccessManagementGetEnterpriseSelfDefault with default headers values
func NewIdentityAccessManagementGetEnterpriseSelfDefault(code int) *IdentityAccessManagementGetEnterpriseSelfDefault {
	return &IdentityAccessManagementGetEnterpriseSelfDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementGetEnterpriseSelfDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementGetEnterpriseSelfDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this identity access management get enterprise self default response has a 2xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management get enterprise self default response has a 3xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management get enterprise self default response has a 4xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management get enterprise self default response has a 5xx status code
func (o *IdentityAccessManagementGetEnterpriseSelfDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management get enterprise self default response a status code equal to that given
func (o *IdentityAccessManagementGetEnterpriseSelfDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the identity access management get enterprise self default response
func (o *IdentityAccessManagementGetEnterpriseSelfDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementGetEnterpriseSelfDefault) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises/self][%d] IdentityAccessManagement_GetEnterpriseSelf default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseSelfDefault) String() string {
	return fmt.Sprintf("[GET /v1/enterprises/self][%d] IdentityAccessManagement_GetEnterpriseSelf default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseSelfDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementGetEnterpriseSelfDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
