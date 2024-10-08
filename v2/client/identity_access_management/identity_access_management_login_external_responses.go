// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// IdentityAccessManagementLoginExternalReader is a Reader for the IdentityAccessManagementLoginExternal structure.
type IdentityAccessManagementLoginExternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementLoginExternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementLoginExternalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementLoginExternalUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementLoginExternalForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementLoginExternalInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementLoginExternalGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementLoginExternalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementLoginExternalOK creates a IdentityAccessManagementLoginExternalOK with default headers values
func NewIdentityAccessManagementLoginExternalOK() *IdentityAccessManagementLoginExternalOK {
	return &IdentityAccessManagementLoginExternalOK{}
}

/*
IdentityAccessManagementLoginExternalOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementLoginExternalOK struct {
	Payload *models.AAAFrontendLoginResponse
}

// IsSuccess returns true when this identity access management login external o k response has a 2xx status code
func (o *IdentityAccessManagementLoginExternalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management login external o k response has a 3xx status code
func (o *IdentityAccessManagementLoginExternalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management login external o k response has a 4xx status code
func (o *IdentityAccessManagementLoginExternalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management login external o k response has a 5xx status code
func (o *IdentityAccessManagementLoginExternalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management login external o k response a status code equal to that given
func (o *IdentityAccessManagementLoginExternalOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the identity access management login external o k response
func (o *IdentityAccessManagementLoginExternalOK) Code() int {
	return 200
}

func (o *IdentityAccessManagementLoginExternalOK) Error() string {
	return fmt.Sprintf("[POST /v1/login/external][%d] identityAccessManagementLoginExternalOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalOK) String() string {
	return fmt.Sprintf("[POST /v1/login/external][%d] identityAccessManagementLoginExternalOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalOK) GetPayload() *models.AAAFrontendLoginResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginExternalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AAAFrontendLoginResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginExternalUnauthorized creates a IdentityAccessManagementLoginExternalUnauthorized with default headers values
func NewIdentityAccessManagementLoginExternalUnauthorized() *IdentityAccessManagementLoginExternalUnauthorized {
	return &IdentityAccessManagementLoginExternalUnauthorized{}
}

/*
IdentityAccessManagementLoginExternalUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementLoginExternalUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management login external unauthorized response has a 2xx status code
func (o *IdentityAccessManagementLoginExternalUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management login external unauthorized response has a 3xx status code
func (o *IdentityAccessManagementLoginExternalUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management login external unauthorized response has a 4xx status code
func (o *IdentityAccessManagementLoginExternalUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management login external unauthorized response has a 5xx status code
func (o *IdentityAccessManagementLoginExternalUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management login external unauthorized response a status code equal to that given
func (o *IdentityAccessManagementLoginExternalUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the identity access management login external unauthorized response
func (o *IdentityAccessManagementLoginExternalUnauthorized) Code() int {
	return 401
}

func (o *IdentityAccessManagementLoginExternalUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/login/external][%d] identityAccessManagementLoginExternalUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/login/external][%d] identityAccessManagementLoginExternalUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginExternalUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginExternalForbidden creates a IdentityAccessManagementLoginExternalForbidden with default headers values
func NewIdentityAccessManagementLoginExternalForbidden() *IdentityAccessManagementLoginExternalForbidden {
	return &IdentityAccessManagementLoginExternalForbidden{}
}

/*
IdentityAccessManagementLoginExternalForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementLoginExternalForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management login external forbidden response has a 2xx status code
func (o *IdentityAccessManagementLoginExternalForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management login external forbidden response has a 3xx status code
func (o *IdentityAccessManagementLoginExternalForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management login external forbidden response has a 4xx status code
func (o *IdentityAccessManagementLoginExternalForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management login external forbidden response has a 5xx status code
func (o *IdentityAccessManagementLoginExternalForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management login external forbidden response a status code equal to that given
func (o *IdentityAccessManagementLoginExternalForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the identity access management login external forbidden response
func (o *IdentityAccessManagementLoginExternalForbidden) Code() int {
	return 403
}

func (o *IdentityAccessManagementLoginExternalForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/login/external][%d] identityAccessManagementLoginExternalForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalForbidden) String() string {
	return fmt.Sprintf("[POST /v1/login/external][%d] identityAccessManagementLoginExternalForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginExternalForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginExternalInternalServerError creates a IdentityAccessManagementLoginExternalInternalServerError with default headers values
func NewIdentityAccessManagementLoginExternalInternalServerError() *IdentityAccessManagementLoginExternalInternalServerError {
	return &IdentityAccessManagementLoginExternalInternalServerError{}
}

/*
IdentityAccessManagementLoginExternalInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementLoginExternalInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management login external internal server error response has a 2xx status code
func (o *IdentityAccessManagementLoginExternalInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management login external internal server error response has a 3xx status code
func (o *IdentityAccessManagementLoginExternalInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management login external internal server error response has a 4xx status code
func (o *IdentityAccessManagementLoginExternalInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management login external internal server error response has a 5xx status code
func (o *IdentityAccessManagementLoginExternalInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management login external internal server error response a status code equal to that given
func (o *IdentityAccessManagementLoginExternalInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the identity access management login external internal server error response
func (o *IdentityAccessManagementLoginExternalInternalServerError) Code() int {
	return 500
}

func (o *IdentityAccessManagementLoginExternalInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/login/external][%d] identityAccessManagementLoginExternalInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/login/external][%d] identityAccessManagementLoginExternalInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginExternalInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginExternalGatewayTimeout creates a IdentityAccessManagementLoginExternalGatewayTimeout with default headers values
func NewIdentityAccessManagementLoginExternalGatewayTimeout() *IdentityAccessManagementLoginExternalGatewayTimeout {
	return &IdentityAccessManagementLoginExternalGatewayTimeout{}
}

/*
IdentityAccessManagementLoginExternalGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementLoginExternalGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management login external gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementLoginExternalGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management login external gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementLoginExternalGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management login external gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementLoginExternalGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management login external gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementLoginExternalGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management login external gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementLoginExternalGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the identity access management login external gateway timeout response
func (o *IdentityAccessManagementLoginExternalGatewayTimeout) Code() int {
	return 504
}

func (o *IdentityAccessManagementLoginExternalGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/login/external][%d] identityAccessManagementLoginExternalGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/login/external][%d] identityAccessManagementLoginExternalGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginExternalGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginExternalDefault creates a IdentityAccessManagementLoginExternalDefault with default headers values
func NewIdentityAccessManagementLoginExternalDefault(code int) *IdentityAccessManagementLoginExternalDefault {
	return &IdentityAccessManagementLoginExternalDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementLoginExternalDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementLoginExternalDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this identity access management login external default response has a 2xx status code
func (o *IdentityAccessManagementLoginExternalDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management login external default response has a 3xx status code
func (o *IdentityAccessManagementLoginExternalDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management login external default response has a 4xx status code
func (o *IdentityAccessManagementLoginExternalDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management login external default response has a 5xx status code
func (o *IdentityAccessManagementLoginExternalDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management login external default response a status code equal to that given
func (o *IdentityAccessManagementLoginExternalDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the identity access management login external default response
func (o *IdentityAccessManagementLoginExternalDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementLoginExternalDefault) Error() string {
	return fmt.Sprintf("[POST /v1/login/external][%d] IdentityAccessManagement_LoginExternal default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalDefault) String() string {
	return fmt.Sprintf("[POST /v1/login/external][%d] IdentityAccessManagement_LoginExternal default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementLoginExternalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
