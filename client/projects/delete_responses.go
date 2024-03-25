package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// ProjectsDeleteReader is a Reader for the ProjectsDelete structure.
type ProjectsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewProjectsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewProjectsDeleteGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewProjectsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectsDeleteOK creates a ProjectsDeleteOK with default headers values
func NewProjectsDeleteOK() *ProjectsDeleteOK {
	return &ProjectsDeleteOK{}
}

/*
ProjectsDeleteOK describes a response with status code 200, with default header values.

A successful response.
*/
type ProjectsDeleteOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this projects delete o k response has a 2xx status code
func (o *ProjectsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects delete o k response has a 3xx status code
func (o *ProjectsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects delete o k response has a 4xx status code
func (o *ProjectsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects delete o k response has a 5xx status code
func (o *ProjectsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects delete o k response a status code equal to that given
func (o *ProjectsDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/projects/id/{id}][%d] projectsDeleteOK  %+v", 200, o.Payload)
}

func (o *ProjectsDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /v1/projects/id/{id}][%d] projectsDeleteOK  %+v", 200, o.Payload)
}

func (o *ProjectsDeleteOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ProjectsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDeleteUnauthorized creates a ProjectsDeleteUnauthorized with default headers values
func NewProjectsDeleteUnauthorized() *ProjectsDeleteUnauthorized {
	return &ProjectsDeleteUnauthorized{}
}

/*
ProjectsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ProjectsDeleteUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this projects delete unauthorized response has a 2xx status code
func (o *ProjectsDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects delete unauthorized response has a 3xx status code
func (o *ProjectsDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects delete unauthorized response has a 4xx status code
func (o *ProjectsDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects delete unauthorized response has a 5xx status code
func (o *ProjectsDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects delete unauthorized response a status code equal to that given
func (o *ProjectsDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/projects/id/{id}][%d] projectsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/projects/id/{id}][%d] projectsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsDeleteUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ProjectsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDeleteForbidden creates a ProjectsDeleteForbidden with default headers values
func NewProjectsDeleteForbidden() *ProjectsDeleteForbidden {
	return &ProjectsDeleteForbidden{}
}

/*
ProjectsDeleteForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the request or does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ProjectsDeleteForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this projects delete forbidden response has a 2xx status code
func (o *ProjectsDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects delete forbidden response has a 3xx status code
func (o *ProjectsDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects delete forbidden response has a 4xx status code
func (o *ProjectsDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects delete forbidden response has a 5xx status code
func (o *ProjectsDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects delete forbidden response a status code equal to that given
func (o *ProjectsDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectsDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/projects/id/{id}][%d] projectsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/projects/id/{id}][%d] projectsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsDeleteForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ProjectsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDeleteNotFound creates a ProjectsDeleteNotFound with default headers values
func NewProjectsDeleteNotFound() *ProjectsDeleteNotFound {
	return &ProjectsDeleteNotFound{}
}

/*
ProjectsDeleteNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type ProjectsDeleteNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this projects delete not found response has a 2xx status code
func (o *ProjectsDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects delete not found response has a 3xx status code
func (o *ProjectsDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects delete not found response has a 4xx status code
func (o *ProjectsDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects delete not found response has a 5xx status code
func (o *ProjectsDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects delete not found response a status code equal to that given
func (o *ProjectsDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/projects/id/{id}][%d] projectsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/projects/id/{id}][%d] projectsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsDeleteNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ProjectsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDeleteInternalServerError creates a ProjectsDeleteInternalServerError with default headers values
func NewProjectsDeleteInternalServerError() *ProjectsDeleteInternalServerError {
	return &ProjectsDeleteInternalServerError{}
}

/*
ProjectsDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ProjectsDeleteInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this projects delete internal server error response has a 2xx status code
func (o *ProjectsDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects delete internal server error response has a 3xx status code
func (o *ProjectsDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects delete internal server error response has a 4xx status code
func (o *ProjectsDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects delete internal server error response has a 5xx status code
func (o *ProjectsDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects delete internal server error response a status code equal to that given
func (o *ProjectsDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectsDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/projects/id/{id}][%d] projectsDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *ProjectsDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/projects/id/{id}][%d] projectsDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *ProjectsDeleteInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ProjectsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDeleteGatewayTimeout creates a ProjectsDeleteGatewayTimeout with default headers values
func NewProjectsDeleteGatewayTimeout() *ProjectsDeleteGatewayTimeout {
	return &ProjectsDeleteGatewayTimeout{}
}

/*
ProjectsDeleteGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ProjectsDeleteGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this projects delete gateway timeout response has a 2xx status code
func (o *ProjectsDeleteGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects delete gateway timeout response has a 3xx status code
func (o *ProjectsDeleteGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects delete gateway timeout response has a 4xx status code
func (o *ProjectsDeleteGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects delete gateway timeout response has a 5xx status code
func (o *ProjectsDeleteGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this projects delete gateway timeout response a status code equal to that given
func (o *ProjectsDeleteGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *ProjectsDeleteGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/projects/id/{id}][%d] projectsDeleteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ProjectsDeleteGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /v1/projects/id/{id}][%d] projectsDeleteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ProjectsDeleteGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ProjectsDeleteGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDeleteDefault creates a ProjectsDeleteDefault with default headers values
func NewProjectsDeleteDefault(code int) *ProjectsDeleteDefault {
	return &ProjectsDeleteDefault{
		_statusCode: code,
	}
}

/*
ProjectsDeleteDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ProjectsDeleteDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the projects delete default response
func (o *ProjectsDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this projects delete default response has a 2xx status code
func (o *ProjectsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this projects delete default response has a 3xx status code
func (o *ProjectsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this projects delete default response has a 4xx status code
func (o *ProjectsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this projects delete default response has a 5xx status code
func (o *ProjectsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this projects delete default response a status code equal to that given
func (o *ProjectsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ProjectsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/projects/id/{id}][%d] Projects_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectsDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/projects/id/{id}][%d] Projects_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectsDeleteDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ProjectsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
