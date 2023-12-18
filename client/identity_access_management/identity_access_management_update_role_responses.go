// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/zededa/terraform-provider/models"
)

// IdentityAccessManagementUpdateRoleReader is a Reader for the IdentityAccessManagementUpdateRole structure.
type IdentityAccessManagementUpdateRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementUpdateRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementUpdateRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementUpdateRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementUpdateRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIdentityAccessManagementUpdateRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewIdentityAccessManagementUpdateRoleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementUpdateRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementUpdateRoleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementUpdateRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementUpdateRoleOK creates a IdentityAccessManagementUpdateRoleOK with default headers values
func NewIdentityAccessManagementUpdateRoleOK() *IdentityAccessManagementUpdateRoleOK {
	return &IdentityAccessManagementUpdateRoleOK{}
}

/*
IdentityAccessManagementUpdateRoleOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementUpdateRoleOK struct {
	Payload *models.CrudResponse
}

// IsSuccess returns true when this identity access management update role o k response has a 2xx status code
func (o *IdentityAccessManagementUpdateRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management update role o k response has a 3xx status code
func (o *IdentityAccessManagementUpdateRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management update role o k response has a 4xx status code
func (o *IdentityAccessManagementUpdateRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management update role o k response has a 5xx status code
func (o *IdentityAccessManagementUpdateRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management update role o k response a status code equal to that given
func (o *IdentityAccessManagementUpdateRoleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the identity access management update role o k response
func (o *IdentityAccessManagementUpdateRoleOK) Code() int {
	return 200
}

func (o *IdentityAccessManagementUpdateRoleOK) Error() string {
	return fmt.Sprintf("[PUT /v1/roles/id/{id}][%d] identityAccessManagementUpdateRoleOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementUpdateRoleOK) String() string {
	return fmt.Sprintf("[PUT /v1/roles/id/{id}][%d] identityAccessManagementUpdateRoleOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementUpdateRoleOK) GetPayload() *models.CrudResponse {
	return o.Payload
}

func (o *IdentityAccessManagementUpdateRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementUpdateRoleUnauthorized creates a IdentityAccessManagementUpdateRoleUnauthorized with default headers values
func NewIdentityAccessManagementUpdateRoleUnauthorized() *IdentityAccessManagementUpdateRoleUnauthorized {
	return &IdentityAccessManagementUpdateRoleUnauthorized{}
}

/*
IdentityAccessManagementUpdateRoleUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementUpdateRoleUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management update role unauthorized response has a 2xx status code
func (o *IdentityAccessManagementUpdateRoleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management update role unauthorized response has a 3xx status code
func (o *IdentityAccessManagementUpdateRoleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management update role unauthorized response has a 4xx status code
func (o *IdentityAccessManagementUpdateRoleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management update role unauthorized response has a 5xx status code
func (o *IdentityAccessManagementUpdateRoleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management update role unauthorized response a status code equal to that given
func (o *IdentityAccessManagementUpdateRoleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the identity access management update role unauthorized response
func (o *IdentityAccessManagementUpdateRoleUnauthorized) Code() int {
	return 401
}

func (o *IdentityAccessManagementUpdateRoleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/roles/id/{id}][%d] identityAccessManagementUpdateRoleUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementUpdateRoleUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/roles/id/{id}][%d] identityAccessManagementUpdateRoleUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementUpdateRoleUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementUpdateRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementUpdateRoleForbidden creates a IdentityAccessManagementUpdateRoleForbidden with default headers values
func NewIdentityAccessManagementUpdateRoleForbidden() *IdentityAccessManagementUpdateRoleForbidden {
	return &IdentityAccessManagementUpdateRoleForbidden{}
}

/*
IdentityAccessManagementUpdateRoleForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementUpdateRoleForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management update role forbidden response has a 2xx status code
func (o *IdentityAccessManagementUpdateRoleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management update role forbidden response has a 3xx status code
func (o *IdentityAccessManagementUpdateRoleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management update role forbidden response has a 4xx status code
func (o *IdentityAccessManagementUpdateRoleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management update role forbidden response has a 5xx status code
func (o *IdentityAccessManagementUpdateRoleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management update role forbidden response a status code equal to that given
func (o *IdentityAccessManagementUpdateRoleForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the identity access management update role forbidden response
func (o *IdentityAccessManagementUpdateRoleForbidden) Code() int {
	return 403
}

func (o *IdentityAccessManagementUpdateRoleForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/roles/id/{id}][%d] identityAccessManagementUpdateRoleForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementUpdateRoleForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/roles/id/{id}][%d] identityAccessManagementUpdateRoleForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementUpdateRoleForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementUpdateRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementUpdateRoleNotFound creates a IdentityAccessManagementUpdateRoleNotFound with default headers values
func NewIdentityAccessManagementUpdateRoleNotFound() *IdentityAccessManagementUpdateRoleNotFound {
	return &IdentityAccessManagementUpdateRoleNotFound{}
}

/*
IdentityAccessManagementUpdateRoleNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type IdentityAccessManagementUpdateRoleNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management update role not found response has a 2xx status code
func (o *IdentityAccessManagementUpdateRoleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management update role not found response has a 3xx status code
func (o *IdentityAccessManagementUpdateRoleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management update role not found response has a 4xx status code
func (o *IdentityAccessManagementUpdateRoleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management update role not found response has a 5xx status code
func (o *IdentityAccessManagementUpdateRoleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management update role not found response a status code equal to that given
func (o *IdentityAccessManagementUpdateRoleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the identity access management update role not found response
func (o *IdentityAccessManagementUpdateRoleNotFound) Code() int {
	return 404
}

func (o *IdentityAccessManagementUpdateRoleNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/roles/id/{id}][%d] identityAccessManagementUpdateRoleNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementUpdateRoleNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/roles/id/{id}][%d] identityAccessManagementUpdateRoleNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementUpdateRoleNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementUpdateRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementUpdateRoleConflict creates a IdentityAccessManagementUpdateRoleConflict with default headers values
func NewIdentityAccessManagementUpdateRoleConflict() *IdentityAccessManagementUpdateRoleConflict {
	return &IdentityAccessManagementUpdateRoleConflict{}
}

/*
IdentityAccessManagementUpdateRoleConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing IAM role record.
*/
type IdentityAccessManagementUpdateRoleConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management update role conflict response has a 2xx status code
func (o *IdentityAccessManagementUpdateRoleConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management update role conflict response has a 3xx status code
func (o *IdentityAccessManagementUpdateRoleConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management update role conflict response has a 4xx status code
func (o *IdentityAccessManagementUpdateRoleConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management update role conflict response has a 5xx status code
func (o *IdentityAccessManagementUpdateRoleConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management update role conflict response a status code equal to that given
func (o *IdentityAccessManagementUpdateRoleConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the identity access management update role conflict response
func (o *IdentityAccessManagementUpdateRoleConflict) Code() int {
	return 409
}

func (o *IdentityAccessManagementUpdateRoleConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/roles/id/{id}][%d] identityAccessManagementUpdateRoleConflict  %+v", 409, o.Payload)
}

func (o *IdentityAccessManagementUpdateRoleConflict) String() string {
	return fmt.Sprintf("[PUT /v1/roles/id/{id}][%d] identityAccessManagementUpdateRoleConflict  %+v", 409, o.Payload)
}

func (o *IdentityAccessManagementUpdateRoleConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementUpdateRoleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementUpdateRoleInternalServerError creates a IdentityAccessManagementUpdateRoleInternalServerError with default headers values
func NewIdentityAccessManagementUpdateRoleInternalServerError() *IdentityAccessManagementUpdateRoleInternalServerError {
	return &IdentityAccessManagementUpdateRoleInternalServerError{}
}

/*
IdentityAccessManagementUpdateRoleInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementUpdateRoleInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management update role internal server error response has a 2xx status code
func (o *IdentityAccessManagementUpdateRoleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management update role internal server error response has a 3xx status code
func (o *IdentityAccessManagementUpdateRoleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management update role internal server error response has a 4xx status code
func (o *IdentityAccessManagementUpdateRoleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management update role internal server error response has a 5xx status code
func (o *IdentityAccessManagementUpdateRoleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management update role internal server error response a status code equal to that given
func (o *IdentityAccessManagementUpdateRoleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the identity access management update role internal server error response
func (o *IdentityAccessManagementUpdateRoleInternalServerError) Code() int {
	return 500
}

func (o *IdentityAccessManagementUpdateRoleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/roles/id/{id}][%d] identityAccessManagementUpdateRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementUpdateRoleInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/roles/id/{id}][%d] identityAccessManagementUpdateRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementUpdateRoleInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementUpdateRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementUpdateRoleGatewayTimeout creates a IdentityAccessManagementUpdateRoleGatewayTimeout with default headers values
func NewIdentityAccessManagementUpdateRoleGatewayTimeout() *IdentityAccessManagementUpdateRoleGatewayTimeout {
	return &IdentityAccessManagementUpdateRoleGatewayTimeout{}
}

/*
IdentityAccessManagementUpdateRoleGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementUpdateRoleGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management update role gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementUpdateRoleGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management update role gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementUpdateRoleGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management update role gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementUpdateRoleGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management update role gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementUpdateRoleGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management update role gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementUpdateRoleGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the identity access management update role gateway timeout response
func (o *IdentityAccessManagementUpdateRoleGatewayTimeout) Code() int {
	return 504
}

func (o *IdentityAccessManagementUpdateRoleGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/roles/id/{id}][%d] identityAccessManagementUpdateRoleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementUpdateRoleGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/roles/id/{id}][%d] identityAccessManagementUpdateRoleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementUpdateRoleGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementUpdateRoleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementUpdateRoleDefault creates a IdentityAccessManagementUpdateRoleDefault with default headers values
func NewIdentityAccessManagementUpdateRoleDefault(code int) *IdentityAccessManagementUpdateRoleDefault {
	return &IdentityAccessManagementUpdateRoleDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementUpdateRoleDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementUpdateRoleDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this identity access management update role default response has a 2xx status code
func (o *IdentityAccessManagementUpdateRoleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management update role default response has a 3xx status code
func (o *IdentityAccessManagementUpdateRoleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management update role default response has a 4xx status code
func (o *IdentityAccessManagementUpdateRoleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management update role default response has a 5xx status code
func (o *IdentityAccessManagementUpdateRoleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management update role default response a status code equal to that given
func (o *IdentityAccessManagementUpdateRoleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the identity access management update role default response
func (o *IdentityAccessManagementUpdateRoleDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementUpdateRoleDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/roles/id/{id}][%d] IdentityAccessManagement_UpdateRole default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementUpdateRoleDefault) String() string {
	return fmt.Sprintf("[PUT /v1/roles/id/{id}][%d] IdentityAccessManagement_UpdateRole default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementUpdateRoleDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementUpdateRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
IdentityAccessManagementUpdateRoleBody Role detail
//
// Role meta data
// Example: {"description":"","id":"AAGFABAEqnH4je5PHZTXSmHOs-XC","name":"SysRoot","revision":{"createdAt":"2020-07-16T18:19:56Z","createdBy":"SYSTEM_ROOT","curr":"1","prev":"","updatedAt":"1970-01-01T00:00:01Z","updatedBy":"SYSTEM_ROOT"},"scopes":[{"accessApp":"PermissionAccessCreateReadUpdateDelete","accessDevice":"PermissionAccessCreateReadUpdateDelete","accessEnterprise":"PermissionAccessCreateReadUpdateDelete","accessStorage":"PermissionAccessCreateReadUpdateDelete","accessUser":"PermissionAccessCreateReadUpdateDelete","enterpriseFilter":["srAll"],"projectFilter":["srAll"]}],"state":"ROLE_STATE_ACTIVE","title":"SysRoot","type":"USER_ROLE_CLUSTER"}
swagger:model IdentityAccessManagementUpdateRoleBody
*/
type IdentityAccessManagementUpdateRoleBody struct {

	// Detailed description of the role
	// Max Length: 256
	Description string `json:"description,omitempty"`

	// User defined name of the role. Name cannot be changed once created
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Name *string `json:"name"`

	// Map of project tags filter
	ProjectTags string `json:"projectTags,omitempty"`

	// System defined info
	// Read Only: true
	Revision *models.ObjectRevision `json:"revision,omitempty"`

	// Scopes/Permissions associated with the role
	// Required: true
	Scopes []*models.Scope `json:"scopes"`

	// State of the role
	State *models.RoleState `json:"state,omitempty"`

	// User defined title of the role. Title can be changed anytime
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+
	Title *string `json:"title"`

	// Type of the role
	// Required: true
	Type *models.UserRole `json:"type"`
}

// Validate validates this identity access management update role body
func (o *IdentityAccessManagementUpdateRoleBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScopes(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IdentityAccessManagementUpdateRoleBody) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(o.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("body"+"."+"description", "body", o.Description, 256); err != nil {
		return err
	}

	return nil
}

func (o *IdentityAccessManagementUpdateRoleBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	if err := validate.MinLength("body"+"."+"name", "body", *o.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("body"+"."+"name", "body", *o.Name, 256); err != nil {
		return err
	}

	if err := validate.Pattern("body"+"."+"name", "body", *o.Name, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (o *IdentityAccessManagementUpdateRoleBody) validateRevision(formats strfmt.Registry) error {
	if swag.IsZero(o.Revision) { // not required
		return nil
	}

	if o.Revision != nil {
		if err := o.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "revision")
			}
			return err
		}
	}

	return nil
}

func (o *IdentityAccessManagementUpdateRoleBody) validateScopes(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"scopes", "body", o.Scopes); err != nil {
		return err
	}

	for i := 0; i < len(o.Scopes); i++ {
		if swag.IsZero(o.Scopes[i]) { // not required
			continue
		}

		if o.Scopes[i] != nil {
			if err := o.Scopes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "scopes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "scopes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *IdentityAccessManagementUpdateRoleBody) validateState(formats strfmt.Registry) error {
	if swag.IsZero(o.State) { // not required
		return nil
	}

	if o.State != nil {
		if err := o.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "state")
			}
			return err
		}
	}

	return nil
}

func (o *IdentityAccessManagementUpdateRoleBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"title", "body", o.Title); err != nil {
		return err
	}

	if err := validate.MinLength("body"+"."+"title", "body", *o.Title, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("body"+"."+"title", "body", *o.Title, 256); err != nil {
		return err
	}

	if err := validate.Pattern("body"+"."+"title", "body", *o.Title, `[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+`); err != nil {
		return err
	}

	return nil
}

func (o *IdentityAccessManagementUpdateRoleBody) validateType(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	if err := validate.Required("body"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	if o.Type != nil {
		if err := o.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this identity access management update role body based on the context it is used
func (o *IdentityAccessManagementUpdateRoleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateScopes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IdentityAccessManagementUpdateRoleBody) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

	if o.Revision != nil {

		if swag.IsZero(o.Revision) { // not required
			return nil
		}

		if err := o.Revision.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "revision")
			}
			return err
		}
	}

	return nil
}

func (o *IdentityAccessManagementUpdateRoleBody) contextValidateScopes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Scopes); i++ {

		if o.Scopes[i] != nil {

			if swag.IsZero(o.Scopes[i]) { // not required
				return nil
			}

			if err := o.Scopes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "scopes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "scopes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *IdentityAccessManagementUpdateRoleBody) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if o.State != nil {

		if swag.IsZero(o.State) { // not required
			return nil
		}

		if err := o.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "state")
			}
			return err
		}
	}

	return nil
}

func (o *IdentityAccessManagementUpdateRoleBody) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if o.Type != nil {

		if err := o.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *IdentityAccessManagementUpdateRoleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IdentityAccessManagementUpdateRoleBody) UnmarshalBinary(b []byte) error {
	var res IdentityAccessManagementUpdateRoleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
