// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/models"
)

// IdentityAccessManagementCreateRoleReader is a Reader for the IdentityAccessManagementCreateRole structure.
type IdentityAccessManagementCreateRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementCreateRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementCreateRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIdentityAccessManagementCreateRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewIdentityAccessManagementCreateRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementCreateRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewIdentityAccessManagementCreateRoleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementCreateRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementCreateRoleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementCreateRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementCreateRoleOK creates a IdentityAccessManagementCreateRoleOK with default headers values
func NewIdentityAccessManagementCreateRoleOK() *IdentityAccessManagementCreateRoleOK {
	return &IdentityAccessManagementCreateRoleOK{}
}

/*
IdentityAccessManagementCreateRoleOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementCreateRoleOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create role o k response has a 2xx status code
func (o *IdentityAccessManagementCreateRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management create role o k response has a 3xx status code
func (o *IdentityAccessManagementCreateRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create role o k response has a 4xx status code
func (o *IdentityAccessManagementCreateRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management create role o k response has a 5xx status code
func (o *IdentityAccessManagementCreateRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management create role o k response a status code equal to that given
func (o *IdentityAccessManagementCreateRoleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the identity access management create role o k response
func (o *IdentityAccessManagementCreateRoleOK) Code() int {
	return 200
}

func (o *IdentityAccessManagementCreateRoleOK) Error() string {
	return fmt.Sprintf("[POST /v1/roles][%d] identityAccessManagementCreateRoleOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementCreateRoleOK) String() string {
	return fmt.Sprintf("[POST /v1/roles][%d] identityAccessManagementCreateRoleOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementCreateRoleOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateRoleBadRequest creates a IdentityAccessManagementCreateRoleBadRequest with default headers values
func NewIdentityAccessManagementCreateRoleBadRequest() *IdentityAccessManagementCreateRoleBadRequest {
	return &IdentityAccessManagementCreateRoleBadRequest{}
}

/*
IdentityAccessManagementCreateRoleBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type IdentityAccessManagementCreateRoleBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create role bad request response has a 2xx status code
func (o *IdentityAccessManagementCreateRoleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create role bad request response has a 3xx status code
func (o *IdentityAccessManagementCreateRoleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create role bad request response has a 4xx status code
func (o *IdentityAccessManagementCreateRoleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management create role bad request response has a 5xx status code
func (o *IdentityAccessManagementCreateRoleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management create role bad request response a status code equal to that given
func (o *IdentityAccessManagementCreateRoleBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the identity access management create role bad request response
func (o *IdentityAccessManagementCreateRoleBadRequest) Code() int {
	return 400
}

func (o *IdentityAccessManagementCreateRoleBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/roles][%d] identityAccessManagementCreateRoleBadRequest  %+v", 400, o.Payload)
}

func (o *IdentityAccessManagementCreateRoleBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/roles][%d] identityAccessManagementCreateRoleBadRequest  %+v", 400, o.Payload)
}

func (o *IdentityAccessManagementCreateRoleBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateRoleUnauthorized creates a IdentityAccessManagementCreateRoleUnauthorized with default headers values
func NewIdentityAccessManagementCreateRoleUnauthorized() *IdentityAccessManagementCreateRoleUnauthorized {
	return &IdentityAccessManagementCreateRoleUnauthorized{}
}

/*
IdentityAccessManagementCreateRoleUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementCreateRoleUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create role unauthorized response has a 2xx status code
func (o *IdentityAccessManagementCreateRoleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create role unauthorized response has a 3xx status code
func (o *IdentityAccessManagementCreateRoleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create role unauthorized response has a 4xx status code
func (o *IdentityAccessManagementCreateRoleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management create role unauthorized response has a 5xx status code
func (o *IdentityAccessManagementCreateRoleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management create role unauthorized response a status code equal to that given
func (o *IdentityAccessManagementCreateRoleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the identity access management create role unauthorized response
func (o *IdentityAccessManagementCreateRoleUnauthorized) Code() int {
	return 401
}

func (o *IdentityAccessManagementCreateRoleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/roles][%d] identityAccessManagementCreateRoleUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementCreateRoleUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/roles][%d] identityAccessManagementCreateRoleUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementCreateRoleUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateRoleForbidden creates a IdentityAccessManagementCreateRoleForbidden with default headers values
func NewIdentityAccessManagementCreateRoleForbidden() *IdentityAccessManagementCreateRoleForbidden {
	return &IdentityAccessManagementCreateRoleForbidden{}
}

/*
IdentityAccessManagementCreateRoleForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementCreateRoleForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create role forbidden response has a 2xx status code
func (o *IdentityAccessManagementCreateRoleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create role forbidden response has a 3xx status code
func (o *IdentityAccessManagementCreateRoleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create role forbidden response has a 4xx status code
func (o *IdentityAccessManagementCreateRoleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management create role forbidden response has a 5xx status code
func (o *IdentityAccessManagementCreateRoleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management create role forbidden response a status code equal to that given
func (o *IdentityAccessManagementCreateRoleForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the identity access management create role forbidden response
func (o *IdentityAccessManagementCreateRoleForbidden) Code() int {
	return 403
}

func (o *IdentityAccessManagementCreateRoleForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/roles][%d] identityAccessManagementCreateRoleForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementCreateRoleForbidden) String() string {
	return fmt.Sprintf("[POST /v1/roles][%d] identityAccessManagementCreateRoleForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementCreateRoleForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateRoleConflict creates a IdentityAccessManagementCreateRoleConflict with default headers values
func NewIdentityAccessManagementCreateRoleConflict() *IdentityAccessManagementCreateRoleConflict {
	return &IdentityAccessManagementCreateRoleConflict{}
}

/*
IdentityAccessManagementCreateRoleConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this IAM role record will conflict with an already existing IAM role record.
*/
type IdentityAccessManagementCreateRoleConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create role conflict response has a 2xx status code
func (o *IdentityAccessManagementCreateRoleConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create role conflict response has a 3xx status code
func (o *IdentityAccessManagementCreateRoleConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create role conflict response has a 4xx status code
func (o *IdentityAccessManagementCreateRoleConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management create role conflict response has a 5xx status code
func (o *IdentityAccessManagementCreateRoleConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management create role conflict response a status code equal to that given
func (o *IdentityAccessManagementCreateRoleConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the identity access management create role conflict response
func (o *IdentityAccessManagementCreateRoleConflict) Code() int {
	return 409
}

func (o *IdentityAccessManagementCreateRoleConflict) Error() string {
	return fmt.Sprintf("[POST /v1/roles][%d] identityAccessManagementCreateRoleConflict  %+v", 409, o.Payload)
}

func (o *IdentityAccessManagementCreateRoleConflict) String() string {
	return fmt.Sprintf("[POST /v1/roles][%d] identityAccessManagementCreateRoleConflict  %+v", 409, o.Payload)
}

func (o *IdentityAccessManagementCreateRoleConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateRoleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateRoleInternalServerError creates a IdentityAccessManagementCreateRoleInternalServerError with default headers values
func NewIdentityAccessManagementCreateRoleInternalServerError() *IdentityAccessManagementCreateRoleInternalServerError {
	return &IdentityAccessManagementCreateRoleInternalServerError{}
}

/*
IdentityAccessManagementCreateRoleInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementCreateRoleInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create role internal server error response has a 2xx status code
func (o *IdentityAccessManagementCreateRoleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create role internal server error response has a 3xx status code
func (o *IdentityAccessManagementCreateRoleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create role internal server error response has a 4xx status code
func (o *IdentityAccessManagementCreateRoleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management create role internal server error response has a 5xx status code
func (o *IdentityAccessManagementCreateRoleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management create role internal server error response a status code equal to that given
func (o *IdentityAccessManagementCreateRoleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the identity access management create role internal server error response
func (o *IdentityAccessManagementCreateRoleInternalServerError) Code() int {
	return 500
}

func (o *IdentityAccessManagementCreateRoleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/roles][%d] identityAccessManagementCreateRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementCreateRoleInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/roles][%d] identityAccessManagementCreateRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementCreateRoleInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateRoleGatewayTimeout creates a IdentityAccessManagementCreateRoleGatewayTimeout with default headers values
func NewIdentityAccessManagementCreateRoleGatewayTimeout() *IdentityAccessManagementCreateRoleGatewayTimeout {
	return &IdentityAccessManagementCreateRoleGatewayTimeout{}
}

/*
IdentityAccessManagementCreateRoleGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementCreateRoleGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create role gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementCreateRoleGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create role gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementCreateRoleGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create role gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementCreateRoleGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management create role gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementCreateRoleGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management create role gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementCreateRoleGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the identity access management create role gateway timeout response
func (o *IdentityAccessManagementCreateRoleGatewayTimeout) Code() int {
	return 504
}

func (o *IdentityAccessManagementCreateRoleGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/roles][%d] identityAccessManagementCreateRoleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementCreateRoleGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/roles][%d] identityAccessManagementCreateRoleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementCreateRoleGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateRoleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateRoleDefault creates a IdentityAccessManagementCreateRoleDefault with default headers values
func NewIdentityAccessManagementCreateRoleDefault(code int) *IdentityAccessManagementCreateRoleDefault {
	return &IdentityAccessManagementCreateRoleDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementCreateRoleDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementCreateRoleDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this identity access management create role default response has a 2xx status code
func (o *IdentityAccessManagementCreateRoleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management create role default response has a 3xx status code
func (o *IdentityAccessManagementCreateRoleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management create role default response has a 4xx status code
func (o *IdentityAccessManagementCreateRoleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management create role default response has a 5xx status code
func (o *IdentityAccessManagementCreateRoleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management create role default response a status code equal to that given
func (o *IdentityAccessManagementCreateRoleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the identity access management create role default response
func (o *IdentityAccessManagementCreateRoleDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementCreateRoleDefault) Error() string {
	return fmt.Sprintf("[POST /v1/roles][%d] IdentityAccessManagement_CreateRole default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementCreateRoleDefault) String() string {
	return fmt.Sprintf("[POST /v1/roles][%d] IdentityAccessManagement_CreateRole default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementCreateRoleDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementCreateRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
