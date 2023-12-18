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

// IdentityAccessManagementDeleteDocPolicyReader is a Reader for the IdentityAccessManagementDeleteDocPolicy structure.
type IdentityAccessManagementDeleteDocPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementDeleteDocPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementDeleteDocPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementDeleteDocPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementDeleteDocPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIdentityAccessManagementDeleteDocPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementDeleteDocPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementDeleteDocPolicyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementDeleteDocPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementDeleteDocPolicyOK creates a IdentityAccessManagementDeleteDocPolicyOK with default headers values
func NewIdentityAccessManagementDeleteDocPolicyOK() *IdentityAccessManagementDeleteDocPolicyOK {
	return &IdentityAccessManagementDeleteDocPolicyOK{}
}

/*
IdentityAccessManagementDeleteDocPolicyOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementDeleteDocPolicyOK struct {
	Payload *models.CrudResponse
}

// IsSuccess returns true when this identity access management delete doc policy o k response has a 2xx status code
func (o *IdentityAccessManagementDeleteDocPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management delete doc policy o k response has a 3xx status code
func (o *IdentityAccessManagementDeleteDocPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete doc policy o k response has a 4xx status code
func (o *IdentityAccessManagementDeleteDocPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management delete doc policy o k response has a 5xx status code
func (o *IdentityAccessManagementDeleteDocPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management delete doc policy o k response a status code equal to that given
func (o *IdentityAccessManagementDeleteDocPolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the identity access management delete doc policy o k response
func (o *IdentityAccessManagementDeleteDocPolicyOK) Code() int {
	return 200
}

func (o *IdentityAccessManagementDeleteDocPolicyOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/cloud/policies/id/{id}][%d] identityAccessManagementDeleteDocPolicyOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementDeleteDocPolicyOK) String() string {
	return fmt.Sprintf("[DELETE /v1/cloud/policies/id/{id}][%d] identityAccessManagementDeleteDocPolicyOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementDeleteDocPolicyOK) GetPayload() *models.CrudResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteDocPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteDocPolicyUnauthorized creates a IdentityAccessManagementDeleteDocPolicyUnauthorized with default headers values
func NewIdentityAccessManagementDeleteDocPolicyUnauthorized() *IdentityAccessManagementDeleteDocPolicyUnauthorized {
	return &IdentityAccessManagementDeleteDocPolicyUnauthorized{}
}

/*
IdentityAccessManagementDeleteDocPolicyUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementDeleteDocPolicyUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete doc policy unauthorized response has a 2xx status code
func (o *IdentityAccessManagementDeleteDocPolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete doc policy unauthorized response has a 3xx status code
func (o *IdentityAccessManagementDeleteDocPolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete doc policy unauthorized response has a 4xx status code
func (o *IdentityAccessManagementDeleteDocPolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management delete doc policy unauthorized response has a 5xx status code
func (o *IdentityAccessManagementDeleteDocPolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management delete doc policy unauthorized response a status code equal to that given
func (o *IdentityAccessManagementDeleteDocPolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the identity access management delete doc policy unauthorized response
func (o *IdentityAccessManagementDeleteDocPolicyUnauthorized) Code() int {
	return 401
}

func (o *IdentityAccessManagementDeleteDocPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/cloud/policies/id/{id}][%d] identityAccessManagementDeleteDocPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementDeleteDocPolicyUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/cloud/policies/id/{id}][%d] identityAccessManagementDeleteDocPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementDeleteDocPolicyUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteDocPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteDocPolicyForbidden creates a IdentityAccessManagementDeleteDocPolicyForbidden with default headers values
func NewIdentityAccessManagementDeleteDocPolicyForbidden() *IdentityAccessManagementDeleteDocPolicyForbidden {
	return &IdentityAccessManagementDeleteDocPolicyForbidden{}
}

/*
IdentityAccessManagementDeleteDocPolicyForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementDeleteDocPolicyForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete doc policy forbidden response has a 2xx status code
func (o *IdentityAccessManagementDeleteDocPolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete doc policy forbidden response has a 3xx status code
func (o *IdentityAccessManagementDeleteDocPolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete doc policy forbidden response has a 4xx status code
func (o *IdentityAccessManagementDeleteDocPolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management delete doc policy forbidden response has a 5xx status code
func (o *IdentityAccessManagementDeleteDocPolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management delete doc policy forbidden response a status code equal to that given
func (o *IdentityAccessManagementDeleteDocPolicyForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the identity access management delete doc policy forbidden response
func (o *IdentityAccessManagementDeleteDocPolicyForbidden) Code() int {
	return 403
}

func (o *IdentityAccessManagementDeleteDocPolicyForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/cloud/policies/id/{id}][%d] identityAccessManagementDeleteDocPolicyForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementDeleteDocPolicyForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/cloud/policies/id/{id}][%d] identityAccessManagementDeleteDocPolicyForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementDeleteDocPolicyForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteDocPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteDocPolicyNotFound creates a IdentityAccessManagementDeleteDocPolicyNotFound with default headers values
func NewIdentityAccessManagementDeleteDocPolicyNotFound() *IdentityAccessManagementDeleteDocPolicyNotFound {
	return &IdentityAccessManagementDeleteDocPolicyNotFound{}
}

/*
IdentityAccessManagementDeleteDocPolicyNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type IdentityAccessManagementDeleteDocPolicyNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete doc policy not found response has a 2xx status code
func (o *IdentityAccessManagementDeleteDocPolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete doc policy not found response has a 3xx status code
func (o *IdentityAccessManagementDeleteDocPolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete doc policy not found response has a 4xx status code
func (o *IdentityAccessManagementDeleteDocPolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management delete doc policy not found response has a 5xx status code
func (o *IdentityAccessManagementDeleteDocPolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management delete doc policy not found response a status code equal to that given
func (o *IdentityAccessManagementDeleteDocPolicyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the identity access management delete doc policy not found response
func (o *IdentityAccessManagementDeleteDocPolicyNotFound) Code() int {
	return 404
}

func (o *IdentityAccessManagementDeleteDocPolicyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/cloud/policies/id/{id}][%d] identityAccessManagementDeleteDocPolicyNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementDeleteDocPolicyNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/cloud/policies/id/{id}][%d] identityAccessManagementDeleteDocPolicyNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementDeleteDocPolicyNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteDocPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteDocPolicyInternalServerError creates a IdentityAccessManagementDeleteDocPolicyInternalServerError with default headers values
func NewIdentityAccessManagementDeleteDocPolicyInternalServerError() *IdentityAccessManagementDeleteDocPolicyInternalServerError {
	return &IdentityAccessManagementDeleteDocPolicyInternalServerError{}
}

/*
IdentityAccessManagementDeleteDocPolicyInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementDeleteDocPolicyInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete doc policy internal server error response has a 2xx status code
func (o *IdentityAccessManagementDeleteDocPolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete doc policy internal server error response has a 3xx status code
func (o *IdentityAccessManagementDeleteDocPolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete doc policy internal server error response has a 4xx status code
func (o *IdentityAccessManagementDeleteDocPolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management delete doc policy internal server error response has a 5xx status code
func (o *IdentityAccessManagementDeleteDocPolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management delete doc policy internal server error response a status code equal to that given
func (o *IdentityAccessManagementDeleteDocPolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the identity access management delete doc policy internal server error response
func (o *IdentityAccessManagementDeleteDocPolicyInternalServerError) Code() int {
	return 500
}

func (o *IdentityAccessManagementDeleteDocPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/cloud/policies/id/{id}][%d] identityAccessManagementDeleteDocPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementDeleteDocPolicyInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/cloud/policies/id/{id}][%d] identityAccessManagementDeleteDocPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementDeleteDocPolicyInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteDocPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteDocPolicyGatewayTimeout creates a IdentityAccessManagementDeleteDocPolicyGatewayTimeout with default headers values
func NewIdentityAccessManagementDeleteDocPolicyGatewayTimeout() *IdentityAccessManagementDeleteDocPolicyGatewayTimeout {
	return &IdentityAccessManagementDeleteDocPolicyGatewayTimeout{}
}

/*
IdentityAccessManagementDeleteDocPolicyGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementDeleteDocPolicyGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete doc policy gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementDeleteDocPolicyGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete doc policy gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementDeleteDocPolicyGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete doc policy gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementDeleteDocPolicyGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management delete doc policy gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementDeleteDocPolicyGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management delete doc policy gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementDeleteDocPolicyGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the identity access management delete doc policy gateway timeout response
func (o *IdentityAccessManagementDeleteDocPolicyGatewayTimeout) Code() int {
	return 504
}

func (o *IdentityAccessManagementDeleteDocPolicyGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/cloud/policies/id/{id}][%d] identityAccessManagementDeleteDocPolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementDeleteDocPolicyGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /v1/cloud/policies/id/{id}][%d] identityAccessManagementDeleteDocPolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementDeleteDocPolicyGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteDocPolicyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteDocPolicyDefault creates a IdentityAccessManagementDeleteDocPolicyDefault with default headers values
func NewIdentityAccessManagementDeleteDocPolicyDefault(code int) *IdentityAccessManagementDeleteDocPolicyDefault {
	return &IdentityAccessManagementDeleteDocPolicyDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementDeleteDocPolicyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementDeleteDocPolicyDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this identity access management delete doc policy default response has a 2xx status code
func (o *IdentityAccessManagementDeleteDocPolicyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management delete doc policy default response has a 3xx status code
func (o *IdentityAccessManagementDeleteDocPolicyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management delete doc policy default response has a 4xx status code
func (o *IdentityAccessManagementDeleteDocPolicyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management delete doc policy default response has a 5xx status code
func (o *IdentityAccessManagementDeleteDocPolicyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management delete doc policy default response a status code equal to that given
func (o *IdentityAccessManagementDeleteDocPolicyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the identity access management delete doc policy default response
func (o *IdentityAccessManagementDeleteDocPolicyDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementDeleteDocPolicyDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/cloud/policies/id/{id}][%d] IdentityAccessManagement_DeleteDocPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementDeleteDocPolicyDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/cloud/policies/id/{id}][%d] IdentityAccessManagement_DeleteDocPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementDeleteDocPolicyDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteDocPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
