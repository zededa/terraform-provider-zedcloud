package application_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/models"
)

// EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceReader is a Reader for the EdgeApplicationInstanceConfigurationGetEdgeApplicationInstance structure.
type EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceOK creates a EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceOK with default headers values
func NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceOK() *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceOK {
	return &EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceOK{}
}

/*
EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceOK struct {
	Payload *models.AppInstance
}

// IsSuccess returns true when this edge application instance configuration get edge application instance o k response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge application instance configuration get edge application instance o k response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration get edge application instance o k response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance configuration get edge application instance o k response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration get edge application instance o k response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge application instance configuration get edge application instance o k response
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceOK) Code() int {
	return 200
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceOK) GetPayload() *models.AppInstance {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceUnauthorized creates a EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceUnauthorized with default headers values
func NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceUnauthorized() *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceUnauthorized {
	return &EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceUnauthorized{}
}

/*
EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration get edge application instance unauthorized response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration get edge application instance unauthorized response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration get edge application instance unauthorized response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance configuration get edge application instance unauthorized response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration get edge application instance unauthorized response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge application instance configuration get edge application instance unauthorized response
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceUnauthorized) Code() int {
	return 401
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceForbidden creates a EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceForbidden with default headers values
func NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceForbidden() *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceForbidden {
	return &EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceForbidden{}
}

/*
EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration get edge application instance forbidden response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration get edge application instance forbidden response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration get edge application instance forbidden response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance configuration get edge application instance forbidden response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration get edge application instance forbidden response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge application instance configuration get edge application instance forbidden response
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceForbidden) Code() int {
	return 403
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceNotFound creates a EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceNotFound with default headers values
func NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceNotFound() *GetApplicationInstanceNotFound {
	return &GetApplicationInstanceNotFound{}
}

/*
GetApplicationInstanceNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetApplicationInstanceNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration get edge application instance not found response has a 2xx status code
func (o *GetApplicationInstanceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration get edge application instance not found response has a 3xx status code
func (o *GetApplicationInstanceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration get edge application instance not found response has a 4xx status code
func (o *GetApplicationInstanceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance configuration get edge application instance not found response has a 5xx status code
func (o *GetApplicationInstanceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration get edge application instance not found response a status code equal to that given
func (o *GetApplicationInstanceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge application instance configuration get edge application instance not found response
func (o *GetApplicationInstanceNotFound) Code() int {
	return 404
}

func (o *GetApplicationInstanceNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceNotFound  %+v", 404, o.Payload)
}

func (o *GetApplicationInstanceNotFound) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceNotFound  %+v", 404, o.Payload)
}

func (o *GetApplicationInstanceNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetApplicationInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceInternalServerError creates a EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceInternalServerError with default headers values
func NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceInternalServerError() *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceInternalServerError {
	return &EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceInternalServerError{}
}

/*
EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration get edge application instance internal server error response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration get edge application instance internal server error response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration get edge application instance internal server error response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance configuration get edge application instance internal server error response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application instance configuration get edge application instance internal server error response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge application instance configuration get edge application instance internal server error response
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceInternalServerError) Code() int {
	return 500
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceGatewayTimeout creates a EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceGatewayTimeout with default headers values
func NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceGatewayTimeout() *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceGatewayTimeout {
	return &EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceGatewayTimeout{}
}

/*
EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration get edge application instance gateway timeout response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration get edge application instance gateway timeout response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration get edge application instance gateway timeout response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance configuration get edge application instance gateway timeout response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application instance configuration get edge application instance gateway timeout response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge application instance configuration get edge application instance gateway timeout response
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceDefault creates a EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceDefault with default headers values
func NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceDefault(code int) *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceDefault {
	return &EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceDefault{
		_statusCode: code,
	}
}

/*
EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge application instance configuration get edge application instance default response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge application instance configuration get edge application instance default response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge application instance configuration get edge application instance default response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge application instance configuration get edge application instance default response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge application instance configuration get edge application instance default response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge application instance configuration get edge application instance default response
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceDefault) Code() int {
	return o._statusCode
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}][%d] EdgeApplicationInstanceConfiguration_GetEdgeApplicationInstance default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceDefault) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}][%d] EdgeApplicationInstanceConfiguration_GetEdgeApplicationInstance default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
