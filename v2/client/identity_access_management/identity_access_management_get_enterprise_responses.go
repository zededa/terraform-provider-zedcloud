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

// IdentityAccessManagementGetEnterpriseReader is a Reader for the IdentityAccessManagementGetEnterprise structure.
type IdentityAccessManagementGetEnterpriseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementGetEnterpriseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementGetEnterpriseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementGetEnterpriseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementGetEnterpriseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIdentityAccessManagementGetEnterpriseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementGetEnterpriseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementGetEnterpriseGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementGetEnterpriseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementGetEnterpriseOK creates a IdentityAccessManagementGetEnterpriseOK with default headers values
func NewIdentityAccessManagementGetEnterpriseOK() *IdentityAccessManagementGetEnterpriseOK {
	return &IdentityAccessManagementGetEnterpriseOK{}
}

/*
IdentityAccessManagementGetEnterpriseOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementGetEnterpriseOK struct {
	Payload *models.Enterprise
}

// IsSuccess returns true when this identity access management get enterprise o k response has a 2xx status code
func (o *IdentityAccessManagementGetEnterpriseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management get enterprise o k response has a 3xx status code
func (o *IdentityAccessManagementGetEnterpriseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get enterprise o k response has a 4xx status code
func (o *IdentityAccessManagementGetEnterpriseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get enterprise o k response has a 5xx status code
func (o *IdentityAccessManagementGetEnterpriseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get enterprise o k response a status code equal to that given
func (o *IdentityAccessManagementGetEnterpriseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the identity access management get enterprise o k response
func (o *IdentityAccessManagementGetEnterpriseOK) Code() int {
	return 200
}

func (o *IdentityAccessManagementGetEnterpriseOK) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises/id/{id}][%d] identityAccessManagementGetEnterpriseOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseOK) String() string {
	return fmt.Sprintf("[GET /v1/enterprises/id/{id}][%d] identityAccessManagementGetEnterpriseOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseOK) GetPayload() *models.Enterprise {
	return o.Payload
}

func (o *IdentityAccessManagementGetEnterpriseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Enterprise)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetEnterpriseUnauthorized creates a IdentityAccessManagementGetEnterpriseUnauthorized with default headers values
func NewIdentityAccessManagementGetEnterpriseUnauthorized() *IdentityAccessManagementGetEnterpriseUnauthorized {
	return &IdentityAccessManagementGetEnterpriseUnauthorized{}
}

/*
IdentityAccessManagementGetEnterpriseUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementGetEnterpriseUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get enterprise unauthorized response has a 2xx status code
func (o *IdentityAccessManagementGetEnterpriseUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get enterprise unauthorized response has a 3xx status code
func (o *IdentityAccessManagementGetEnterpriseUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get enterprise unauthorized response has a 4xx status code
func (o *IdentityAccessManagementGetEnterpriseUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get enterprise unauthorized response has a 5xx status code
func (o *IdentityAccessManagementGetEnterpriseUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get enterprise unauthorized response a status code equal to that given
func (o *IdentityAccessManagementGetEnterpriseUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the identity access management get enterprise unauthorized response
func (o *IdentityAccessManagementGetEnterpriseUnauthorized) Code() int {
	return 401
}

func (o *IdentityAccessManagementGetEnterpriseUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises/id/{id}][%d] identityAccessManagementGetEnterpriseUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/enterprises/id/{id}][%d] identityAccessManagementGetEnterpriseUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetEnterpriseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetEnterpriseForbidden creates a IdentityAccessManagementGetEnterpriseForbidden with default headers values
func NewIdentityAccessManagementGetEnterpriseForbidden() *IdentityAccessManagementGetEnterpriseForbidden {
	return &IdentityAccessManagementGetEnterpriseForbidden{}
}

/*
IdentityAccessManagementGetEnterpriseForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementGetEnterpriseForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get enterprise forbidden response has a 2xx status code
func (o *IdentityAccessManagementGetEnterpriseForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get enterprise forbidden response has a 3xx status code
func (o *IdentityAccessManagementGetEnterpriseForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get enterprise forbidden response has a 4xx status code
func (o *IdentityAccessManagementGetEnterpriseForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get enterprise forbidden response has a 5xx status code
func (o *IdentityAccessManagementGetEnterpriseForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get enterprise forbidden response a status code equal to that given
func (o *IdentityAccessManagementGetEnterpriseForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the identity access management get enterprise forbidden response
func (o *IdentityAccessManagementGetEnterpriseForbidden) Code() int {
	return 403
}

func (o *IdentityAccessManagementGetEnterpriseForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises/id/{id}][%d] identityAccessManagementGetEnterpriseForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseForbidden) String() string {
	return fmt.Sprintf("[GET /v1/enterprises/id/{id}][%d] identityAccessManagementGetEnterpriseForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetEnterpriseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetEnterpriseNotFound creates a IdentityAccessManagementGetEnterpriseNotFound with default headers values
func NewIdentityAccessManagementGetEnterpriseNotFound() *IdentityAccessManagementGetEnterpriseNotFound {
	return &IdentityAccessManagementGetEnterpriseNotFound{}
}

/*
IdentityAccessManagementGetEnterpriseNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type IdentityAccessManagementGetEnterpriseNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get enterprise not found response has a 2xx status code
func (o *IdentityAccessManagementGetEnterpriseNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get enterprise not found response has a 3xx status code
func (o *IdentityAccessManagementGetEnterpriseNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get enterprise not found response has a 4xx status code
func (o *IdentityAccessManagementGetEnterpriseNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get enterprise not found response has a 5xx status code
func (o *IdentityAccessManagementGetEnterpriseNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get enterprise not found response a status code equal to that given
func (o *IdentityAccessManagementGetEnterpriseNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the identity access management get enterprise not found response
func (o *IdentityAccessManagementGetEnterpriseNotFound) Code() int {
	return 404
}

func (o *IdentityAccessManagementGetEnterpriseNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises/id/{id}][%d] identityAccessManagementGetEnterpriseNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseNotFound) String() string {
	return fmt.Sprintf("[GET /v1/enterprises/id/{id}][%d] identityAccessManagementGetEnterpriseNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetEnterpriseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetEnterpriseInternalServerError creates a IdentityAccessManagementGetEnterpriseInternalServerError with default headers values
func NewIdentityAccessManagementGetEnterpriseInternalServerError() *IdentityAccessManagementGetEnterpriseInternalServerError {
	return &IdentityAccessManagementGetEnterpriseInternalServerError{}
}

/*
IdentityAccessManagementGetEnterpriseInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementGetEnterpriseInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get enterprise internal server error response has a 2xx status code
func (o *IdentityAccessManagementGetEnterpriseInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get enterprise internal server error response has a 3xx status code
func (o *IdentityAccessManagementGetEnterpriseInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get enterprise internal server error response has a 4xx status code
func (o *IdentityAccessManagementGetEnterpriseInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get enterprise internal server error response has a 5xx status code
func (o *IdentityAccessManagementGetEnterpriseInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management get enterprise internal server error response a status code equal to that given
func (o *IdentityAccessManagementGetEnterpriseInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the identity access management get enterprise internal server error response
func (o *IdentityAccessManagementGetEnterpriseInternalServerError) Code() int {
	return 500
}

func (o *IdentityAccessManagementGetEnterpriseInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises/id/{id}][%d] identityAccessManagementGetEnterpriseInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/enterprises/id/{id}][%d] identityAccessManagementGetEnterpriseInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetEnterpriseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetEnterpriseGatewayTimeout creates a IdentityAccessManagementGetEnterpriseGatewayTimeout with default headers values
func NewIdentityAccessManagementGetEnterpriseGatewayTimeout() *IdentityAccessManagementGetEnterpriseGatewayTimeout {
	return &IdentityAccessManagementGetEnterpriseGatewayTimeout{}
}

/*
IdentityAccessManagementGetEnterpriseGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementGetEnterpriseGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get enterprise gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementGetEnterpriseGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get enterprise gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementGetEnterpriseGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get enterprise gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementGetEnterpriseGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get enterprise gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementGetEnterpriseGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management get enterprise gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementGetEnterpriseGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the identity access management get enterprise gateway timeout response
func (o *IdentityAccessManagementGetEnterpriseGatewayTimeout) Code() int {
	return 504
}

func (o *IdentityAccessManagementGetEnterpriseGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises/id/{id}][%d] identityAccessManagementGetEnterpriseGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/enterprises/id/{id}][%d] identityAccessManagementGetEnterpriseGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetEnterpriseGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetEnterpriseDefault creates a IdentityAccessManagementGetEnterpriseDefault with default headers values
func NewIdentityAccessManagementGetEnterpriseDefault(code int) *IdentityAccessManagementGetEnterpriseDefault {
	return &IdentityAccessManagementGetEnterpriseDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementGetEnterpriseDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementGetEnterpriseDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this identity access management get enterprise default response has a 2xx status code
func (o *IdentityAccessManagementGetEnterpriseDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management get enterprise default response has a 3xx status code
func (o *IdentityAccessManagementGetEnterpriseDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management get enterprise default response has a 4xx status code
func (o *IdentityAccessManagementGetEnterpriseDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management get enterprise default response has a 5xx status code
func (o *IdentityAccessManagementGetEnterpriseDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management get enterprise default response a status code equal to that given
func (o *IdentityAccessManagementGetEnterpriseDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the identity access management get enterprise default response
func (o *IdentityAccessManagementGetEnterpriseDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementGetEnterpriseDefault) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises/id/{id}][%d] IdentityAccessManagement_GetEnterprise default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseDefault) String() string {
	return fmt.Sprintf("[GET /v1/enterprises/id/{id}][%d] IdentityAccessManagement_GetEnterprise default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementGetEnterpriseDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementGetEnterpriseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
