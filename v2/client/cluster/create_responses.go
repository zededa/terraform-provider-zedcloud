// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// CreateReader is a Reader for the Create structure.
type CreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewCreatePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewCreateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateOK creates a CreateOK with default headers values
func NewCreateOK() *CreateOK {
	return &CreateOK{}
}

/*
CreateOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration create cluster o k response has a 2xx status code
func (o *CreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node cluster configuration create cluster o k response has a 3xx status code
func (o *CreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration create cluster o k response has a 4xx status code
func (o *CreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node cluster configuration create cluster o k response has a 5xx status code
func (o *CreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node cluster configuration create cluster o k response a status code equal to that given
func (o *CreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge node cluster configuration create cluster o k response
func (o *CreateOK) Code() int {
	return 200
}

func (o *CreateOK) Error() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] edgeNodeConfigurationCreateOK  %+v", 200, o.Payload)
}

func (o *CreateOK) String() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] edgeNodeConfigurationCreateOK  %+v", 200, o.Payload)
}

func (o *CreateOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *CreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBadRequest creates a CreateBadRequest with default headers values
func NewCreateBadRequest() *CreateBadRequest {
	return &CreateBadRequest{}
}

/*
CreateBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type CreateBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration create cluster bad request response has a 2xx status code
func (o *CreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration create cluster bad request response has a 3xx status code
func (o *CreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration create cluster bad request response has a 4xx status code
func (o *CreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node cluster configuration create cluster bad request response has a 5xx status code
func (o *CreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node cluster configuration create cluster bad request response a status code equal to that given
func (o *CreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the edge node cluster configuration create cluster bad request response
func (o *CreateBadRequest) Code() int {
	return 400
}

func (o *CreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] edgeNodeConfigurationCreateBadRequest  %+v", 400, o.Payload)
}

func (o *CreateBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] edgeNodeConfigurationCreateBadRequest  %+v", 400, o.Payload)
}

func (o *CreateBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *CreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUnauthorized creates a CreateUnauthorized with default headers values
func NewCreateUnauthorized() *CreateUnauthorized {
	return &CreateUnauthorized{}
}

/*
CreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type CreateUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration create cluster unauthorized response has a 2xx status code
func (o *CreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration create cluster unauthorized response has a 3xx status code
func (o *CreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration create cluster unauthorized response has a 4xx status code
func (o *CreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node cluster configuration create cluster unauthorized response has a 5xx status code
func (o *CreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node cluster configuration create cluster unauthorized response a status code equal to that given
func (o *CreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge node cluster configuration create cluster unauthorized response
func (o *CreateUnauthorized) Code() int {
	return 401
}

func (o *CreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] edgeNodeConfigurationCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] edgeNodeConfigurationCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *CreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateForbidden creates a CreateForbidden with default headers values
func NewCreateForbidden() *CreateForbidden {
	return &CreateForbidden{}
}

/*
CreateForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type CreateForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration create cluster forbidden response has a 2xx status code
func (o *CreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration create cluster forbidden response has a 3xx status code
func (o *CreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration create cluster forbidden response has a 4xx status code
func (o *CreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node cluster configuration create cluster forbidden response has a 5xx status code
func (o *CreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node cluster configuration create cluster forbidden response a status code equal to that given
func (o *CreateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge node cluster configuration create cluster forbidden response
func (o *CreateForbidden) Code() int {
	return 403
}

func (o *CreateForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] edgeNodeConfigurationCreateForbidden  %+v", 403, o.Payload)
}

func (o *CreateForbidden) String() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] edgeNodeConfigurationCreateForbidden  %+v", 403, o.Payload)
}

func (o *CreateForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *CreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNotFound creates a CreateNotFound with default headers values
func NewCreateNotFound() *CreateNotFound {
	return &CreateNotFound{}
}

/*
CreateNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type CreateNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration create cluster not found response has a 2xx status code
func (o *CreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration create cluster not found response has a 3xx status code
func (o *CreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration create cluster not found response has a 4xx status code
func (o *CreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node cluster configuration create cluster not found response has a 5xx status code
func (o *CreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node cluster configuration create cluster not found response a status code equal to that given
func (o *CreateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge node cluster configuration create cluster not found response
func (o *CreateNotFound) Code() int {
	return 404
}

func (o *CreateNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] edgeNodeConfigurationCreateNotFound  %+v", 404, o.Payload)
}

func (o *CreateNotFound) String() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] edgeNodeConfigurationCreateNotFound  %+v", 404, o.Payload)
}

func (o *CreateNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *CreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConflict creates a CreateConflict with default headers values
func NewCreateConflict() *CreateConflict {
	return &CreateConflict{}
}

/*
CreateConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this edge-node cluster record will conflict with an already existing edge-node cluster record.
*/
type CreateConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration create cluster conflict response has a 2xx status code
func (o *CreateConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration create cluster conflict response has a 3xx status code
func (o *CreateConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration create cluster conflict response has a 4xx status code
func (o *CreateConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node cluster configuration create cluster conflict response has a 5xx status code
func (o *CreateConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node cluster configuration create cluster conflict response a status code equal to that given
func (o *CreateConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the edge node cluster configuration create cluster conflict response
func (o *CreateConflict) Code() int {
	return 409
}

func (o *CreateConflict) Error() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] edgeNodeConfigurationCreateConflict  %+v", 409, o.Payload)
}

func (o *CreateConflict) String() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] edgeNodeConfigurationCreateConflict  %+v", 409, o.Payload)
}

func (o *CreateConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *CreateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePreconditionFailed creates a CreatePreconditionFailed with default headers values
func NewCreatePreconditionFailed() *CreatePreconditionFailed {
	return &CreatePreconditionFailed{}
}

/*
CreatePreconditionFailed describes a response with status code 412, with default header values.

Precondition failed. Some of preconditions haven't been met to start request processing.
*/
type CreatePreconditionFailed struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration create cluster precondition failed response has a 2xx status code
func (o *CreatePreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration create cluster precondition failed response has a 3xx status code
func (o *CreatePreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration create cluster precondition failed response has a 4xx status code
func (o *CreatePreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node cluster configuration create cluster precondition failed response has a 5xx status code
func (o *CreatePreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node cluster configuration create cluster precondition failed response a status code equal to that given
func (o *CreatePreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the edge node cluster configuration create cluster precondition failed response
func (o *CreatePreconditionFailed) Code() int {
	return 412
}

func (o *CreatePreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] edgeNodeConfigurationCreatePreconditionFailed  %+v", 412, o.Payload)
}

func (o *CreatePreconditionFailed) String() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] edgeNodeConfigurationCreatePreconditionFailed  %+v", 412, o.Payload)
}

func (o *CreatePreconditionFailed) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *CreatePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternalServerError creates a CreateInternalServerError with default headers values
func NewCreateInternalServerError() *CreateInternalServerError {
	return &CreateInternalServerError{}
}

/*
CreateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type CreateInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration create cluster internal server error response has a 2xx status code
func (o *CreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration create cluster internal server error response has a 3xx status code
func (o *CreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration create cluster internal server error response has a 4xx status code
func (o *CreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node cluster configuration create cluster internal server error response has a 5xx status code
func (o *CreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node cluster configuration create cluster internal server error response a status code equal to that given
func (o *CreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge node cluster configuration create cluster internal server error response
func (o *CreateInternalServerError) Code() int {
	return 500
}

func (o *CreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] edgeNodeConfigurationCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] edgeNodeConfigurationCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *CreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGatewayTimeout creates a CreateGatewayTimeout with default headers values
func NewCreateGatewayTimeout() *CreateGatewayTimeout {
	return &CreateGatewayTimeout{}
}

/*
CreateGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type CreateGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration create cluster gateway timeout response has a 2xx status code
func (o *CreateGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration create cluster gateway timeout response has a 3xx status code
func (o *CreateGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration create cluster gateway timeout response has a 4xx status code
func (o *CreateGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node cluster configuration create cluster gateway timeout response has a 5xx status code
func (o *CreateGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node cluster configuration create cluster gateway timeout response a status code equal to that given
func (o *CreateGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge node cluster configuration create cluster gateway timeout response
func (o *CreateGatewayTimeout) Code() int {
	return 504
}

func (o *CreateGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] edgeNodeConfigurationCreateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *CreateGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] edgeNodeConfigurationCreateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *CreateGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *CreateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDefault creates a CreateDefault with default headers values
func NewCreateDefault(code int) *CreateDefault {
	return &CreateDefault{
		_statusCode: code,
	}
}

/*
CreateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge node cluster configuration create cluster default response has a 2xx status code
func (o *CreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node cluster configuration create cluster default response has a 3xx status code
func (o *CreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node cluster configuration create cluster default response has a 4xx status code
func (o *CreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node cluster configuration create cluster default response has a 5xx status code
func (o *CreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node cluster configuration create cluster default response a status code equal to that given
func (o *CreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge node cluster configuration create cluster default response
func (o *CreateDefault) Code() int {
	return o._statusCode
}

func (o *CreateDefault) Error() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] _Create default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDefault) String() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] _Create default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *CreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
