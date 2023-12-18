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

// IdentityAccessManagementGetDocPolicyReader is a Reader for the IdentityAccessManagementGetDocPolicy structure.
type IdentityAccessManagementGetDocPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementGetDocPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementGetDocPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementGetDocPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementGetDocPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIdentityAccessManagementGetDocPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementGetDocPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementGetDocPolicyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementGetDocPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementGetDocPolicyOK creates a IdentityAccessManagementGetDocPolicyOK with default headers values
func NewIdentityAccessManagementGetDocPolicyOK() *IdentityAccessManagementGetDocPolicyOK {
	return &IdentityAccessManagementGetDocPolicyOK{}
}

/*
IdentityAccessManagementGetDocPolicyOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementGetDocPolicyOK struct {
	Payload *models.CrudResponseRead
}

// IsSuccess returns true when this identity access management get doc policy o k response has a 2xx status code
func (o *IdentityAccessManagementGetDocPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management get doc policy o k response has a 3xx status code
func (o *IdentityAccessManagementGetDocPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get doc policy o k response has a 4xx status code
func (o *IdentityAccessManagementGetDocPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get doc policy o k response has a 5xx status code
func (o *IdentityAccessManagementGetDocPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get doc policy o k response a status code equal to that given
func (o *IdentityAccessManagementGetDocPolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the identity access management get doc policy o k response
func (o *IdentityAccessManagementGetDocPolicyOK) Code() int {
	return 200
}

func (o *IdentityAccessManagementGetDocPolicyOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{id}][%d] identityAccessManagementGetDocPolicyOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementGetDocPolicyOK) String() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{id}][%d] identityAccessManagementGetDocPolicyOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementGetDocPolicyOK) GetPayload() *models.CrudResponseRead {
	return o.Payload
}

func (o *IdentityAccessManagementGetDocPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CrudResponseRead)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetDocPolicyUnauthorized creates a IdentityAccessManagementGetDocPolicyUnauthorized with default headers values
func NewIdentityAccessManagementGetDocPolicyUnauthorized() *IdentityAccessManagementGetDocPolicyUnauthorized {
	return &IdentityAccessManagementGetDocPolicyUnauthorized{}
}

/*
IdentityAccessManagementGetDocPolicyUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementGetDocPolicyUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get doc policy unauthorized response has a 2xx status code
func (o *IdentityAccessManagementGetDocPolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get doc policy unauthorized response has a 3xx status code
func (o *IdentityAccessManagementGetDocPolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get doc policy unauthorized response has a 4xx status code
func (o *IdentityAccessManagementGetDocPolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get doc policy unauthorized response has a 5xx status code
func (o *IdentityAccessManagementGetDocPolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get doc policy unauthorized response a status code equal to that given
func (o *IdentityAccessManagementGetDocPolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the identity access management get doc policy unauthorized response
func (o *IdentityAccessManagementGetDocPolicyUnauthorized) Code() int {
	return 401
}

func (o *IdentityAccessManagementGetDocPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{id}][%d] identityAccessManagementGetDocPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementGetDocPolicyUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{id}][%d] identityAccessManagementGetDocPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementGetDocPolicyUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetDocPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetDocPolicyForbidden creates a IdentityAccessManagementGetDocPolicyForbidden with default headers values
func NewIdentityAccessManagementGetDocPolicyForbidden() *IdentityAccessManagementGetDocPolicyForbidden {
	return &IdentityAccessManagementGetDocPolicyForbidden{}
}

/*
IdentityAccessManagementGetDocPolicyForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementGetDocPolicyForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get doc policy forbidden response has a 2xx status code
func (o *IdentityAccessManagementGetDocPolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get doc policy forbidden response has a 3xx status code
func (o *IdentityAccessManagementGetDocPolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get doc policy forbidden response has a 4xx status code
func (o *IdentityAccessManagementGetDocPolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get doc policy forbidden response has a 5xx status code
func (o *IdentityAccessManagementGetDocPolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get doc policy forbidden response a status code equal to that given
func (o *IdentityAccessManagementGetDocPolicyForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the identity access management get doc policy forbidden response
func (o *IdentityAccessManagementGetDocPolicyForbidden) Code() int {
	return 403
}

func (o *IdentityAccessManagementGetDocPolicyForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{id}][%d] identityAccessManagementGetDocPolicyForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementGetDocPolicyForbidden) String() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{id}][%d] identityAccessManagementGetDocPolicyForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementGetDocPolicyForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetDocPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetDocPolicyNotFound creates a IdentityAccessManagementGetDocPolicyNotFound with default headers values
func NewIdentityAccessManagementGetDocPolicyNotFound() *IdentityAccessManagementGetDocPolicyNotFound {
	return &IdentityAccessManagementGetDocPolicyNotFound{}
}

/*
IdentityAccessManagementGetDocPolicyNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type IdentityAccessManagementGetDocPolicyNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get doc policy not found response has a 2xx status code
func (o *IdentityAccessManagementGetDocPolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get doc policy not found response has a 3xx status code
func (o *IdentityAccessManagementGetDocPolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get doc policy not found response has a 4xx status code
func (o *IdentityAccessManagementGetDocPolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get doc policy not found response has a 5xx status code
func (o *IdentityAccessManagementGetDocPolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get doc policy not found response a status code equal to that given
func (o *IdentityAccessManagementGetDocPolicyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the identity access management get doc policy not found response
func (o *IdentityAccessManagementGetDocPolicyNotFound) Code() int {
	return 404
}

func (o *IdentityAccessManagementGetDocPolicyNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{id}][%d] identityAccessManagementGetDocPolicyNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementGetDocPolicyNotFound) String() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{id}][%d] identityAccessManagementGetDocPolicyNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementGetDocPolicyNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetDocPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetDocPolicyInternalServerError creates a IdentityAccessManagementGetDocPolicyInternalServerError with default headers values
func NewIdentityAccessManagementGetDocPolicyInternalServerError() *IdentityAccessManagementGetDocPolicyInternalServerError {
	return &IdentityAccessManagementGetDocPolicyInternalServerError{}
}

/*
IdentityAccessManagementGetDocPolicyInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementGetDocPolicyInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get doc policy internal server error response has a 2xx status code
func (o *IdentityAccessManagementGetDocPolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get doc policy internal server error response has a 3xx status code
func (o *IdentityAccessManagementGetDocPolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get doc policy internal server error response has a 4xx status code
func (o *IdentityAccessManagementGetDocPolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get doc policy internal server error response has a 5xx status code
func (o *IdentityAccessManagementGetDocPolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management get doc policy internal server error response a status code equal to that given
func (o *IdentityAccessManagementGetDocPolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the identity access management get doc policy internal server error response
func (o *IdentityAccessManagementGetDocPolicyInternalServerError) Code() int {
	return 500
}

func (o *IdentityAccessManagementGetDocPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{id}][%d] identityAccessManagementGetDocPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementGetDocPolicyInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{id}][%d] identityAccessManagementGetDocPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementGetDocPolicyInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetDocPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetDocPolicyGatewayTimeout creates a IdentityAccessManagementGetDocPolicyGatewayTimeout with default headers values
func NewIdentityAccessManagementGetDocPolicyGatewayTimeout() *IdentityAccessManagementGetDocPolicyGatewayTimeout {
	return &IdentityAccessManagementGetDocPolicyGatewayTimeout{}
}

/*
IdentityAccessManagementGetDocPolicyGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementGetDocPolicyGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get doc policy gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementGetDocPolicyGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get doc policy gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementGetDocPolicyGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get doc policy gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementGetDocPolicyGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get doc policy gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementGetDocPolicyGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management get doc policy gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementGetDocPolicyGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the identity access management get doc policy gateway timeout response
func (o *IdentityAccessManagementGetDocPolicyGatewayTimeout) Code() int {
	return 504
}

func (o *IdentityAccessManagementGetDocPolicyGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{id}][%d] identityAccessManagementGetDocPolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementGetDocPolicyGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{id}][%d] identityAccessManagementGetDocPolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementGetDocPolicyGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetDocPolicyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetDocPolicyDefault creates a IdentityAccessManagementGetDocPolicyDefault with default headers values
func NewIdentityAccessManagementGetDocPolicyDefault(code int) *IdentityAccessManagementGetDocPolicyDefault {
	return &IdentityAccessManagementGetDocPolicyDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementGetDocPolicyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementGetDocPolicyDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this identity access management get doc policy default response has a 2xx status code
func (o *IdentityAccessManagementGetDocPolicyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management get doc policy default response has a 3xx status code
func (o *IdentityAccessManagementGetDocPolicyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management get doc policy default response has a 4xx status code
func (o *IdentityAccessManagementGetDocPolicyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management get doc policy default response has a 5xx status code
func (o *IdentityAccessManagementGetDocPolicyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management get doc policy default response a status code equal to that given
func (o *IdentityAccessManagementGetDocPolicyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the identity access management get doc policy default response
func (o *IdentityAccessManagementGetDocPolicyDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementGetDocPolicyDefault) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{id}][%d] IdentityAccessManagement_GetDocPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementGetDocPolicyDefault) String() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/id/{id}][%d] IdentityAccessManagement_GetDocPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementGetDocPolicyDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementGetDocPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
