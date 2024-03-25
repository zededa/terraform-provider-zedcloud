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

// IdentityAccessManagementCreateUserReader is a Reader for the IdentityAccessManagementCreateUser structure.
type IdentityAccessManagementCreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementCreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementCreateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIdentityAccessManagementCreateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewIdentityAccessManagementCreateUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementCreateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewIdentityAccessManagementCreateUserConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementCreateUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementCreateUserGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementCreateUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementCreateUserOK creates a IdentityAccessManagementCreateUserOK with default headers values
func NewIdentityAccessManagementCreateUserOK() *IdentityAccessManagementCreateUserOK {
	return &IdentityAccessManagementCreateUserOK{}
}

/*
IdentityAccessManagementCreateUserOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementCreateUserOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create user o k response has a 2xx status code
func (o *IdentityAccessManagementCreateUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management create user o k response has a 3xx status code
func (o *IdentityAccessManagementCreateUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create user o k response has a 4xx status code
func (o *IdentityAccessManagementCreateUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management create user o k response has a 5xx status code
func (o *IdentityAccessManagementCreateUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management create user o k response a status code equal to that given
func (o *IdentityAccessManagementCreateUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the identity access management create user o k response
func (o *IdentityAccessManagementCreateUserOK) Code() int {
	return 200
}

func (o *IdentityAccessManagementCreateUserOK) Error() string {
	return fmt.Sprintf("[POST /v1/users][%d] identityAccessManagementCreateUserOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementCreateUserOK) String() string {
	return fmt.Sprintf("[POST /v1/users][%d] identityAccessManagementCreateUserOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementCreateUserOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateUserBadRequest creates a IdentityAccessManagementCreateUserBadRequest with default headers values
func NewIdentityAccessManagementCreateUserBadRequest() *IdentityAccessManagementCreateUserBadRequest {
	return &IdentityAccessManagementCreateUserBadRequest{}
}

/*
IdentityAccessManagementCreateUserBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type IdentityAccessManagementCreateUserBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create user bad request response has a 2xx status code
func (o *IdentityAccessManagementCreateUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create user bad request response has a 3xx status code
func (o *IdentityAccessManagementCreateUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create user bad request response has a 4xx status code
func (o *IdentityAccessManagementCreateUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management create user bad request response has a 5xx status code
func (o *IdentityAccessManagementCreateUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management create user bad request response a status code equal to that given
func (o *IdentityAccessManagementCreateUserBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the identity access management create user bad request response
func (o *IdentityAccessManagementCreateUserBadRequest) Code() int {
	return 400
}

func (o *IdentityAccessManagementCreateUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/users][%d] identityAccessManagementCreateUserBadRequest  %+v", 400, o.Payload)
}

func (o *IdentityAccessManagementCreateUserBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/users][%d] identityAccessManagementCreateUserBadRequest  %+v", 400, o.Payload)
}

func (o *IdentityAccessManagementCreateUserBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateUserUnauthorized creates a IdentityAccessManagementCreateUserUnauthorized with default headers values
func NewIdentityAccessManagementCreateUserUnauthorized() *IdentityAccessManagementCreateUserUnauthorized {
	return &IdentityAccessManagementCreateUserUnauthorized{}
}

/*
IdentityAccessManagementCreateUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementCreateUserUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create user unauthorized response has a 2xx status code
func (o *IdentityAccessManagementCreateUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create user unauthorized response has a 3xx status code
func (o *IdentityAccessManagementCreateUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create user unauthorized response has a 4xx status code
func (o *IdentityAccessManagementCreateUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management create user unauthorized response has a 5xx status code
func (o *IdentityAccessManagementCreateUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management create user unauthorized response a status code equal to that given
func (o *IdentityAccessManagementCreateUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the identity access management create user unauthorized response
func (o *IdentityAccessManagementCreateUserUnauthorized) Code() int {
	return 401
}

func (o *IdentityAccessManagementCreateUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/users][%d] identityAccessManagementCreateUserUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementCreateUserUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/users][%d] identityAccessManagementCreateUserUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementCreateUserUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateUserForbidden creates a IdentityAccessManagementCreateUserForbidden with default headers values
func NewIdentityAccessManagementCreateUserForbidden() *IdentityAccessManagementCreateUserForbidden {
	return &IdentityAccessManagementCreateUserForbidden{}
}

/*
IdentityAccessManagementCreateUserForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementCreateUserForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create user forbidden response has a 2xx status code
func (o *IdentityAccessManagementCreateUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create user forbidden response has a 3xx status code
func (o *IdentityAccessManagementCreateUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create user forbidden response has a 4xx status code
func (o *IdentityAccessManagementCreateUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management create user forbidden response has a 5xx status code
func (o *IdentityAccessManagementCreateUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management create user forbidden response a status code equal to that given
func (o *IdentityAccessManagementCreateUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the identity access management create user forbidden response
func (o *IdentityAccessManagementCreateUserForbidden) Code() int {
	return 403
}

func (o *IdentityAccessManagementCreateUserForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/users][%d] identityAccessManagementCreateUserForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementCreateUserForbidden) String() string {
	return fmt.Sprintf("[POST /v1/users][%d] identityAccessManagementCreateUserForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementCreateUserForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateUserConflict creates a IdentityAccessManagementCreateUserConflict with default headers values
func NewIdentityAccessManagementCreateUserConflict() *IdentityAccessManagementCreateUserConflict {
	return &IdentityAccessManagementCreateUserConflict{}
}

/*
IdentityAccessManagementCreateUserConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this IAM user record will conflict with an already existing IAM user record.
*/
type IdentityAccessManagementCreateUserConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create user conflict response has a 2xx status code
func (o *IdentityAccessManagementCreateUserConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create user conflict response has a 3xx status code
func (o *IdentityAccessManagementCreateUserConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create user conflict response has a 4xx status code
func (o *IdentityAccessManagementCreateUserConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management create user conflict response has a 5xx status code
func (o *IdentityAccessManagementCreateUserConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management create user conflict response a status code equal to that given
func (o *IdentityAccessManagementCreateUserConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the identity access management create user conflict response
func (o *IdentityAccessManagementCreateUserConflict) Code() int {
	return 409
}

func (o *IdentityAccessManagementCreateUserConflict) Error() string {
	return fmt.Sprintf("[POST /v1/users][%d] identityAccessManagementCreateUserConflict  %+v", 409, o.Payload)
}

func (o *IdentityAccessManagementCreateUserConflict) String() string {
	return fmt.Sprintf("[POST /v1/users][%d] identityAccessManagementCreateUserConflict  %+v", 409, o.Payload)
}

func (o *IdentityAccessManagementCreateUserConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateUserConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateUserInternalServerError creates a IdentityAccessManagementCreateUserInternalServerError with default headers values
func NewIdentityAccessManagementCreateUserInternalServerError() *IdentityAccessManagementCreateUserInternalServerError {
	return &IdentityAccessManagementCreateUserInternalServerError{}
}

/*
IdentityAccessManagementCreateUserInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementCreateUserInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create user internal server error response has a 2xx status code
func (o *IdentityAccessManagementCreateUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create user internal server error response has a 3xx status code
func (o *IdentityAccessManagementCreateUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create user internal server error response has a 4xx status code
func (o *IdentityAccessManagementCreateUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management create user internal server error response has a 5xx status code
func (o *IdentityAccessManagementCreateUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management create user internal server error response a status code equal to that given
func (o *IdentityAccessManagementCreateUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the identity access management create user internal server error response
func (o *IdentityAccessManagementCreateUserInternalServerError) Code() int {
	return 500
}

func (o *IdentityAccessManagementCreateUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/users][%d] identityAccessManagementCreateUserInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementCreateUserInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/users][%d] identityAccessManagementCreateUserInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementCreateUserInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateUserGatewayTimeout creates a IdentityAccessManagementCreateUserGatewayTimeout with default headers values
func NewIdentityAccessManagementCreateUserGatewayTimeout() *IdentityAccessManagementCreateUserGatewayTimeout {
	return &IdentityAccessManagementCreateUserGatewayTimeout{}
}

/*
IdentityAccessManagementCreateUserGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementCreateUserGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create user gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementCreateUserGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create user gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementCreateUserGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create user gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementCreateUserGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management create user gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementCreateUserGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management create user gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementCreateUserGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the identity access management create user gateway timeout response
func (o *IdentityAccessManagementCreateUserGatewayTimeout) Code() int {
	return 504
}

func (o *IdentityAccessManagementCreateUserGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/users][%d] identityAccessManagementCreateUserGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementCreateUserGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/users][%d] identityAccessManagementCreateUserGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementCreateUserGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateUserGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateUserDefault creates a IdentityAccessManagementCreateUserDefault with default headers values
func NewIdentityAccessManagementCreateUserDefault(code int) *IdentityAccessManagementCreateUserDefault {
	return &IdentityAccessManagementCreateUserDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementCreateUserDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementCreateUserDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this identity access management create user default response has a 2xx status code
func (o *IdentityAccessManagementCreateUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management create user default response has a 3xx status code
func (o *IdentityAccessManagementCreateUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management create user default response has a 4xx status code
func (o *IdentityAccessManagementCreateUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management create user default response has a 5xx status code
func (o *IdentityAccessManagementCreateUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management create user default response a status code equal to that given
func (o *IdentityAccessManagementCreateUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the identity access management create user default response
func (o *IdentityAccessManagementCreateUserDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementCreateUserDefault) Error() string {
	return fmt.Sprintf("[POST /v1/users][%d] IdentityAccessManagement_CreateUser default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementCreateUserDefault) String() string {
	return fmt.Sprintf("[POST /v1/users][%d] IdentityAccessManagement_CreateUser default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementCreateUserDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementCreateUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
