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

// IdentityAccessManagementCreateCredentialReader is a Reader for the IdentityAccessManagementCreateCredential structure.
type IdentityAccessManagementCreateCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementCreateCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementCreateCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIdentityAccessManagementCreateCredentialBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewIdentityAccessManagementCreateCredentialUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementCreateCredentialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewIdentityAccessManagementCreateCredentialConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementCreateCredentialInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementCreateCredentialGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementCreateCredentialDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementCreateCredentialOK creates a IdentityAccessManagementCreateCredentialOK with default headers values
func NewIdentityAccessManagementCreateCredentialOK() *IdentityAccessManagementCreateCredentialOK {
	return &IdentityAccessManagementCreateCredentialOK{}
}

/*
IdentityAccessManagementCreateCredentialOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementCreateCredentialOK struct {
	Payload *models.CrudResponse
}

// IsSuccess returns true when this identity access management create credential o k response has a 2xx status code
func (o *IdentityAccessManagementCreateCredentialOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management create credential o k response has a 3xx status code
func (o *IdentityAccessManagementCreateCredentialOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create credential o k response has a 4xx status code
func (o *IdentityAccessManagementCreateCredentialOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management create credential o k response has a 5xx status code
func (o *IdentityAccessManagementCreateCredentialOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management create credential o k response a status code equal to that given
func (o *IdentityAccessManagementCreateCredentialOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the identity access management create credential o k response
func (o *IdentityAccessManagementCreateCredentialOK) Code() int {
	return 200
}

func (o *IdentityAccessManagementCreateCredentialOK) Error() string {
	return fmt.Sprintf("[POST /v1/credentials][%d] identityAccessManagementCreateCredentialOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementCreateCredentialOK) String() string {
	return fmt.Sprintf("[POST /v1/credentials][%d] identityAccessManagementCreateCredentialOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementCreateCredentialOK) GetPayload() *models.CrudResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateCredentialBadRequest creates a IdentityAccessManagementCreateCredentialBadRequest with default headers values
func NewIdentityAccessManagementCreateCredentialBadRequest() *IdentityAccessManagementCreateCredentialBadRequest {
	return &IdentityAccessManagementCreateCredentialBadRequest{}
}

/*
IdentityAccessManagementCreateCredentialBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type IdentityAccessManagementCreateCredentialBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create credential bad request response has a 2xx status code
func (o *IdentityAccessManagementCreateCredentialBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create credential bad request response has a 3xx status code
func (o *IdentityAccessManagementCreateCredentialBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create credential bad request response has a 4xx status code
func (o *IdentityAccessManagementCreateCredentialBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management create credential bad request response has a 5xx status code
func (o *IdentityAccessManagementCreateCredentialBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management create credential bad request response a status code equal to that given
func (o *IdentityAccessManagementCreateCredentialBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the identity access management create credential bad request response
func (o *IdentityAccessManagementCreateCredentialBadRequest) Code() int {
	return 400
}

func (o *IdentityAccessManagementCreateCredentialBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/credentials][%d] identityAccessManagementCreateCredentialBadRequest  %+v", 400, o.Payload)
}

func (o *IdentityAccessManagementCreateCredentialBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/credentials][%d] identityAccessManagementCreateCredentialBadRequest  %+v", 400, o.Payload)
}

func (o *IdentityAccessManagementCreateCredentialBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateCredentialBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateCredentialUnauthorized creates a IdentityAccessManagementCreateCredentialUnauthorized with default headers values
func NewIdentityAccessManagementCreateCredentialUnauthorized() *IdentityAccessManagementCreateCredentialUnauthorized {
	return &IdentityAccessManagementCreateCredentialUnauthorized{}
}

/*
IdentityAccessManagementCreateCredentialUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementCreateCredentialUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create credential unauthorized response has a 2xx status code
func (o *IdentityAccessManagementCreateCredentialUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create credential unauthorized response has a 3xx status code
func (o *IdentityAccessManagementCreateCredentialUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create credential unauthorized response has a 4xx status code
func (o *IdentityAccessManagementCreateCredentialUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management create credential unauthorized response has a 5xx status code
func (o *IdentityAccessManagementCreateCredentialUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management create credential unauthorized response a status code equal to that given
func (o *IdentityAccessManagementCreateCredentialUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the identity access management create credential unauthorized response
func (o *IdentityAccessManagementCreateCredentialUnauthorized) Code() int {
	return 401
}

func (o *IdentityAccessManagementCreateCredentialUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/credentials][%d] identityAccessManagementCreateCredentialUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementCreateCredentialUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/credentials][%d] identityAccessManagementCreateCredentialUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementCreateCredentialUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateCredentialUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateCredentialForbidden creates a IdentityAccessManagementCreateCredentialForbidden with default headers values
func NewIdentityAccessManagementCreateCredentialForbidden() *IdentityAccessManagementCreateCredentialForbidden {
	return &IdentityAccessManagementCreateCredentialForbidden{}
}

/*
IdentityAccessManagementCreateCredentialForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementCreateCredentialForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create credential forbidden response has a 2xx status code
func (o *IdentityAccessManagementCreateCredentialForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create credential forbidden response has a 3xx status code
func (o *IdentityAccessManagementCreateCredentialForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create credential forbidden response has a 4xx status code
func (o *IdentityAccessManagementCreateCredentialForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management create credential forbidden response has a 5xx status code
func (o *IdentityAccessManagementCreateCredentialForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management create credential forbidden response a status code equal to that given
func (o *IdentityAccessManagementCreateCredentialForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the identity access management create credential forbidden response
func (o *IdentityAccessManagementCreateCredentialForbidden) Code() int {
	return 403
}

func (o *IdentityAccessManagementCreateCredentialForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/credentials][%d] identityAccessManagementCreateCredentialForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementCreateCredentialForbidden) String() string {
	return fmt.Sprintf("[POST /v1/credentials][%d] identityAccessManagementCreateCredentialForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementCreateCredentialForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateCredentialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateCredentialConflict creates a IdentityAccessManagementCreateCredentialConflict with default headers values
func NewIdentityAccessManagementCreateCredentialConflict() *IdentityAccessManagementCreateCredentialConflict {
	return &IdentityAccessManagementCreateCredentialConflict{}
}

/*
IdentityAccessManagementCreateCredentialConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this IAM credential record will conflict with an already existing IAM credential record.
*/
type IdentityAccessManagementCreateCredentialConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create credential conflict response has a 2xx status code
func (o *IdentityAccessManagementCreateCredentialConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create credential conflict response has a 3xx status code
func (o *IdentityAccessManagementCreateCredentialConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create credential conflict response has a 4xx status code
func (o *IdentityAccessManagementCreateCredentialConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management create credential conflict response has a 5xx status code
func (o *IdentityAccessManagementCreateCredentialConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management create credential conflict response a status code equal to that given
func (o *IdentityAccessManagementCreateCredentialConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the identity access management create credential conflict response
func (o *IdentityAccessManagementCreateCredentialConflict) Code() int {
	return 409
}

func (o *IdentityAccessManagementCreateCredentialConflict) Error() string {
	return fmt.Sprintf("[POST /v1/credentials][%d] identityAccessManagementCreateCredentialConflict  %+v", 409, o.Payload)
}

func (o *IdentityAccessManagementCreateCredentialConflict) String() string {
	return fmt.Sprintf("[POST /v1/credentials][%d] identityAccessManagementCreateCredentialConflict  %+v", 409, o.Payload)
}

func (o *IdentityAccessManagementCreateCredentialConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateCredentialConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateCredentialInternalServerError creates a IdentityAccessManagementCreateCredentialInternalServerError with default headers values
func NewIdentityAccessManagementCreateCredentialInternalServerError() *IdentityAccessManagementCreateCredentialInternalServerError {
	return &IdentityAccessManagementCreateCredentialInternalServerError{}
}

/*
IdentityAccessManagementCreateCredentialInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementCreateCredentialInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create credential internal server error response has a 2xx status code
func (o *IdentityAccessManagementCreateCredentialInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create credential internal server error response has a 3xx status code
func (o *IdentityAccessManagementCreateCredentialInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create credential internal server error response has a 4xx status code
func (o *IdentityAccessManagementCreateCredentialInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management create credential internal server error response has a 5xx status code
func (o *IdentityAccessManagementCreateCredentialInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management create credential internal server error response a status code equal to that given
func (o *IdentityAccessManagementCreateCredentialInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the identity access management create credential internal server error response
func (o *IdentityAccessManagementCreateCredentialInternalServerError) Code() int {
	return 500
}

func (o *IdentityAccessManagementCreateCredentialInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/credentials][%d] identityAccessManagementCreateCredentialInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementCreateCredentialInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/credentials][%d] identityAccessManagementCreateCredentialInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementCreateCredentialInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateCredentialInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateCredentialGatewayTimeout creates a IdentityAccessManagementCreateCredentialGatewayTimeout with default headers values
func NewIdentityAccessManagementCreateCredentialGatewayTimeout() *IdentityAccessManagementCreateCredentialGatewayTimeout {
	return &IdentityAccessManagementCreateCredentialGatewayTimeout{}
}

/*
IdentityAccessManagementCreateCredentialGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementCreateCredentialGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management create credential gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementCreateCredentialGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management create credential gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementCreateCredentialGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management create credential gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementCreateCredentialGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management create credential gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementCreateCredentialGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management create credential gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementCreateCredentialGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the identity access management create credential gateway timeout response
func (o *IdentityAccessManagementCreateCredentialGatewayTimeout) Code() int {
	return 504
}

func (o *IdentityAccessManagementCreateCredentialGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/credentials][%d] identityAccessManagementCreateCredentialGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementCreateCredentialGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/credentials][%d] identityAccessManagementCreateCredentialGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementCreateCredentialGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateCredentialGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateCredentialDefault creates a IdentityAccessManagementCreateCredentialDefault with default headers values
func NewIdentityAccessManagementCreateCredentialDefault(code int) *IdentityAccessManagementCreateCredentialDefault {
	return &IdentityAccessManagementCreateCredentialDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementCreateCredentialDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementCreateCredentialDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this identity access management create credential default response has a 2xx status code
func (o *IdentityAccessManagementCreateCredentialDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management create credential default response has a 3xx status code
func (o *IdentityAccessManagementCreateCredentialDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management create credential default response has a 4xx status code
func (o *IdentityAccessManagementCreateCredentialDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management create credential default response has a 5xx status code
func (o *IdentityAccessManagementCreateCredentialDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management create credential default response a status code equal to that given
func (o *IdentityAccessManagementCreateCredentialDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the identity access management create credential default response
func (o *IdentityAccessManagementCreateCredentialDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementCreateCredentialDefault) Error() string {
	return fmt.Sprintf("[POST /v1/credentials][%d] IdentityAccessManagement_CreateCredential default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementCreateCredentialDefault) String() string {
	return fmt.Sprintf("[POST /v1/credentials][%d] IdentityAccessManagement_CreateCredential default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementCreateCredentialDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementCreateCredentialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
