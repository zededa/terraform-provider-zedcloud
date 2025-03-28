// Code generated by go-swagger; DO NOT EDIT.

package app_profile_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// AppProfileServiceDeleteAppProfileReader is a Reader for the AppProfileServiceDeleteAppProfile structure.
type AppProfileServiceDeleteAppProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppProfileServiceDeleteAppProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppProfileServiceDeleteAppProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAppProfileServiceDeleteAppProfileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAppProfileServiceDeleteAppProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAppProfileServiceDeleteAppProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAppProfileServiceDeleteAppProfileConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAppProfileServiceDeleteAppProfileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewAppProfileServiceDeleteAppProfileGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAppProfileServiceDeleteAppProfileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAppProfileServiceDeleteAppProfileOK creates a AppProfileServiceDeleteAppProfileOK with default headers values
func NewAppProfileServiceDeleteAppProfileOK() *AppProfileServiceDeleteAppProfileOK {
	return &AppProfileServiceDeleteAppProfileOK{}
}

/*
AppProfileServiceDeleteAppProfileOK describes a response with status code 200, with default header values.

A successful response.
*/
type AppProfileServiceDeleteAppProfileOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this app profile service delete app profile o k response has a 2xx status code
func (o *AppProfileServiceDeleteAppProfileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this app profile service delete app profile o k response has a 3xx status code
func (o *AppProfileServiceDeleteAppProfileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app profile service delete app profile o k response has a 4xx status code
func (o *AppProfileServiceDeleteAppProfileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this app profile service delete app profile o k response has a 5xx status code
func (o *AppProfileServiceDeleteAppProfileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this app profile service delete app profile o k response a status code equal to that given
func (o *AppProfileServiceDeleteAppProfileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the app profile service delete app profile o k response
func (o *AppProfileServiceDeleteAppProfileOK) Code() int {
	return 200
}

func (o *AppProfileServiceDeleteAppProfileOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/appprofiles/id/{id}][%d] appProfileServiceDeleteAppProfileOK %s", 200, payload)
}

func (o *AppProfileServiceDeleteAppProfileOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/appprofiles/id/{id}][%d] appProfileServiceDeleteAppProfileOK %s", 200, payload)
}

func (o *AppProfileServiceDeleteAppProfileOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AppProfileServiceDeleteAppProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppProfileServiceDeleteAppProfileUnauthorized creates a AppProfileServiceDeleteAppProfileUnauthorized with default headers values
func NewAppProfileServiceDeleteAppProfileUnauthorized() *AppProfileServiceDeleteAppProfileUnauthorized {
	return &AppProfileServiceDeleteAppProfileUnauthorized{}
}

/*
AppProfileServiceDeleteAppProfileUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type AppProfileServiceDeleteAppProfileUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this app profile service delete app profile unauthorized response has a 2xx status code
func (o *AppProfileServiceDeleteAppProfileUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app profile service delete app profile unauthorized response has a 3xx status code
func (o *AppProfileServiceDeleteAppProfileUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app profile service delete app profile unauthorized response has a 4xx status code
func (o *AppProfileServiceDeleteAppProfileUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this app profile service delete app profile unauthorized response has a 5xx status code
func (o *AppProfileServiceDeleteAppProfileUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this app profile service delete app profile unauthorized response a status code equal to that given
func (o *AppProfileServiceDeleteAppProfileUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the app profile service delete app profile unauthorized response
func (o *AppProfileServiceDeleteAppProfileUnauthorized) Code() int {
	return 401
}

func (o *AppProfileServiceDeleteAppProfileUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/appprofiles/id/{id}][%d] appProfileServiceDeleteAppProfileUnauthorized %s", 401, payload)
}

func (o *AppProfileServiceDeleteAppProfileUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/appprofiles/id/{id}][%d] appProfileServiceDeleteAppProfileUnauthorized %s", 401, payload)
}

func (o *AppProfileServiceDeleteAppProfileUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AppProfileServiceDeleteAppProfileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppProfileServiceDeleteAppProfileForbidden creates a AppProfileServiceDeleteAppProfileForbidden with default headers values
func NewAppProfileServiceDeleteAppProfileForbidden() *AppProfileServiceDeleteAppProfileForbidden {
	return &AppProfileServiceDeleteAppProfileForbidden{}
}

/*
AppProfileServiceDeleteAppProfileForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type AppProfileServiceDeleteAppProfileForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this app profile service delete app profile forbidden response has a 2xx status code
func (o *AppProfileServiceDeleteAppProfileForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app profile service delete app profile forbidden response has a 3xx status code
func (o *AppProfileServiceDeleteAppProfileForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app profile service delete app profile forbidden response has a 4xx status code
func (o *AppProfileServiceDeleteAppProfileForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this app profile service delete app profile forbidden response has a 5xx status code
func (o *AppProfileServiceDeleteAppProfileForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this app profile service delete app profile forbidden response a status code equal to that given
func (o *AppProfileServiceDeleteAppProfileForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the app profile service delete app profile forbidden response
func (o *AppProfileServiceDeleteAppProfileForbidden) Code() int {
	return 403
}

func (o *AppProfileServiceDeleteAppProfileForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/appprofiles/id/{id}][%d] appProfileServiceDeleteAppProfileForbidden %s", 403, payload)
}

func (o *AppProfileServiceDeleteAppProfileForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/appprofiles/id/{id}][%d] appProfileServiceDeleteAppProfileForbidden %s", 403, payload)
}

func (o *AppProfileServiceDeleteAppProfileForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AppProfileServiceDeleteAppProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppProfileServiceDeleteAppProfileNotFound creates a AppProfileServiceDeleteAppProfileNotFound with default headers values
func NewAppProfileServiceDeleteAppProfileNotFound() *AppProfileServiceDeleteAppProfileNotFound {
	return &AppProfileServiceDeleteAppProfileNotFound{}
}

/*
AppProfileServiceDeleteAppProfileNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type AppProfileServiceDeleteAppProfileNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this app profile service delete app profile not found response has a 2xx status code
func (o *AppProfileServiceDeleteAppProfileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app profile service delete app profile not found response has a 3xx status code
func (o *AppProfileServiceDeleteAppProfileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app profile service delete app profile not found response has a 4xx status code
func (o *AppProfileServiceDeleteAppProfileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this app profile service delete app profile not found response has a 5xx status code
func (o *AppProfileServiceDeleteAppProfileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this app profile service delete app profile not found response a status code equal to that given
func (o *AppProfileServiceDeleteAppProfileNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the app profile service delete app profile not found response
func (o *AppProfileServiceDeleteAppProfileNotFound) Code() int {
	return 404
}

func (o *AppProfileServiceDeleteAppProfileNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/appprofiles/id/{id}][%d] appProfileServiceDeleteAppProfileNotFound %s", 404, payload)
}

func (o *AppProfileServiceDeleteAppProfileNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/appprofiles/id/{id}][%d] appProfileServiceDeleteAppProfileNotFound %s", 404, payload)
}

func (o *AppProfileServiceDeleteAppProfileNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AppProfileServiceDeleteAppProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppProfileServiceDeleteAppProfileConflict creates a AppProfileServiceDeleteAppProfileConflict with default headers values
func NewAppProfileServiceDeleteAppProfileConflict() *AppProfileServiceDeleteAppProfileConflict {
	return &AppProfileServiceDeleteAppProfileConflict{}
}

/*
AppProfileServiceDeleteAppProfileConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because there are instances of this app profile deployed in edge node(s)
*/
type AppProfileServiceDeleteAppProfileConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this app profile service delete app profile conflict response has a 2xx status code
func (o *AppProfileServiceDeleteAppProfileConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app profile service delete app profile conflict response has a 3xx status code
func (o *AppProfileServiceDeleteAppProfileConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app profile service delete app profile conflict response has a 4xx status code
func (o *AppProfileServiceDeleteAppProfileConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this app profile service delete app profile conflict response has a 5xx status code
func (o *AppProfileServiceDeleteAppProfileConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this app profile service delete app profile conflict response a status code equal to that given
func (o *AppProfileServiceDeleteAppProfileConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the app profile service delete app profile conflict response
func (o *AppProfileServiceDeleteAppProfileConflict) Code() int {
	return 409
}

func (o *AppProfileServiceDeleteAppProfileConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/appprofiles/id/{id}][%d] appProfileServiceDeleteAppProfileConflict %s", 409, payload)
}

func (o *AppProfileServiceDeleteAppProfileConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/appprofiles/id/{id}][%d] appProfileServiceDeleteAppProfileConflict %s", 409, payload)
}

func (o *AppProfileServiceDeleteAppProfileConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AppProfileServiceDeleteAppProfileConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppProfileServiceDeleteAppProfileInternalServerError creates a AppProfileServiceDeleteAppProfileInternalServerError with default headers values
func NewAppProfileServiceDeleteAppProfileInternalServerError() *AppProfileServiceDeleteAppProfileInternalServerError {
	return &AppProfileServiceDeleteAppProfileInternalServerError{}
}

/*
AppProfileServiceDeleteAppProfileInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type AppProfileServiceDeleteAppProfileInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this app profile service delete app profile internal server error response has a 2xx status code
func (o *AppProfileServiceDeleteAppProfileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app profile service delete app profile internal server error response has a 3xx status code
func (o *AppProfileServiceDeleteAppProfileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app profile service delete app profile internal server error response has a 4xx status code
func (o *AppProfileServiceDeleteAppProfileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this app profile service delete app profile internal server error response has a 5xx status code
func (o *AppProfileServiceDeleteAppProfileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this app profile service delete app profile internal server error response a status code equal to that given
func (o *AppProfileServiceDeleteAppProfileInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the app profile service delete app profile internal server error response
func (o *AppProfileServiceDeleteAppProfileInternalServerError) Code() int {
	return 500
}

func (o *AppProfileServiceDeleteAppProfileInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/appprofiles/id/{id}][%d] appProfileServiceDeleteAppProfileInternalServerError %s", 500, payload)
}

func (o *AppProfileServiceDeleteAppProfileInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/appprofiles/id/{id}][%d] appProfileServiceDeleteAppProfileInternalServerError %s", 500, payload)
}

func (o *AppProfileServiceDeleteAppProfileInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AppProfileServiceDeleteAppProfileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppProfileServiceDeleteAppProfileGatewayTimeout creates a AppProfileServiceDeleteAppProfileGatewayTimeout with default headers values
func NewAppProfileServiceDeleteAppProfileGatewayTimeout() *AppProfileServiceDeleteAppProfileGatewayTimeout {
	return &AppProfileServiceDeleteAppProfileGatewayTimeout{}
}

/*
AppProfileServiceDeleteAppProfileGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type AppProfileServiceDeleteAppProfileGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this app profile service delete app profile gateway timeout response has a 2xx status code
func (o *AppProfileServiceDeleteAppProfileGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app profile service delete app profile gateway timeout response has a 3xx status code
func (o *AppProfileServiceDeleteAppProfileGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app profile service delete app profile gateway timeout response has a 4xx status code
func (o *AppProfileServiceDeleteAppProfileGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this app profile service delete app profile gateway timeout response has a 5xx status code
func (o *AppProfileServiceDeleteAppProfileGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this app profile service delete app profile gateway timeout response a status code equal to that given
func (o *AppProfileServiceDeleteAppProfileGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the app profile service delete app profile gateway timeout response
func (o *AppProfileServiceDeleteAppProfileGatewayTimeout) Code() int {
	return 504
}

func (o *AppProfileServiceDeleteAppProfileGatewayTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/appprofiles/id/{id}][%d] appProfileServiceDeleteAppProfileGatewayTimeout %s", 504, payload)
}

func (o *AppProfileServiceDeleteAppProfileGatewayTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/appprofiles/id/{id}][%d] appProfileServiceDeleteAppProfileGatewayTimeout %s", 504, payload)
}

func (o *AppProfileServiceDeleteAppProfileGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AppProfileServiceDeleteAppProfileGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppProfileServiceDeleteAppProfileDefault creates a AppProfileServiceDeleteAppProfileDefault with default headers values
func NewAppProfileServiceDeleteAppProfileDefault(code int) *AppProfileServiceDeleteAppProfileDefault {
	return &AppProfileServiceDeleteAppProfileDefault{
		_statusCode: code,
	}
}

/*
AppProfileServiceDeleteAppProfileDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AppProfileServiceDeleteAppProfileDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this app profile service delete app profile default response has a 2xx status code
func (o *AppProfileServiceDeleteAppProfileDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this app profile service delete app profile default response has a 3xx status code
func (o *AppProfileServiceDeleteAppProfileDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this app profile service delete app profile default response has a 4xx status code
func (o *AppProfileServiceDeleteAppProfileDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this app profile service delete app profile default response has a 5xx status code
func (o *AppProfileServiceDeleteAppProfileDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this app profile service delete app profile default response a status code equal to that given
func (o *AppProfileServiceDeleteAppProfileDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the app profile service delete app profile default response
func (o *AppProfileServiceDeleteAppProfileDefault) Code() int {
	return o._statusCode
}

func (o *AppProfileServiceDeleteAppProfileDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/appprofiles/id/{id}][%d] AppProfileService_DeleteAppProfile default %s", o._statusCode, payload)
}

func (o *AppProfileServiceDeleteAppProfileDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/appprofiles/id/{id}][%d] AppProfileService_DeleteAppProfile default %s", o._statusCode, payload)
}

func (o *AppProfileServiceDeleteAppProfileDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *AppProfileServiceDeleteAppProfileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
