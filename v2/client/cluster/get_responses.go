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

// GetReader is a Reader for the Get structure.
type GetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewGetPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOK creates a GetOK with default headers values
func NewGetOK() *GetOK {
	return &GetOK{}
}

/*
GetOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetOK struct {
	Payload *models.ClusterConfigSummary
}

// IsSuccess returns true when this edge node cluster configuration get cluster o k response has a 2xx status code
func (o *GetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node cluster configuration get cluster o k response has a 3xx status code
func (o *GetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration get cluster o k response has a 4xx status code
func (o *GetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node cluster configuration get cluster o k response has a 5xx status code
func (o *GetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node cluster configuration get cluster o k response a status code equal to that given
func (o *GetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge node cluster configuration get cluster o k response
func (o *GetOK) Code() int {
	return 200
}

func (o *GetOK) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/id/{id}][%d] edgeNodeConfigurationGetOK  %+v", 200, o.Payload)
}

func (o *GetOK) String() string {
	return fmt.Sprintf("[GET /v1/cluster/id/{id}][%d] edgeNodeConfigurationGetOK  %+v", 200, o.Payload)
}

func (o *GetOK) GetPayload() *models.ClusterConfigSummary {
	return o.Payload
}

func (o *GetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterConfigSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUnauthorized creates a GetUnauthorized with default headers values
func NewGetUnauthorized() *GetUnauthorized {
	return &GetUnauthorized{}
}

/*
GetUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration get cluster unauthorized response has a 2xx status code
func (o *GetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration get cluster unauthorized response has a 3xx status code
func (o *GetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration get cluster unauthorized response has a 4xx status code
func (o *GetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node cluster configuration get cluster unauthorized response has a 5xx status code
func (o *GetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node cluster configuration get cluster unauthorized response a status code equal to that given
func (o *GetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge node cluster configuration get cluster unauthorized response
func (o *GetUnauthorized) Code() int {
	return 401
}

func (o *GetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/id/{id}][%d] edgeNodeConfigurationGetUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/cluster/id/{id}][%d] edgeNodeConfigurationGetUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetForbidden creates a GetForbidden with default headers values
func NewGetForbidden() *GetForbidden {
	return &GetForbidden{}
}

/*
GetForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type GetForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration get cluster forbidden response has a 2xx status code
func (o *GetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration get cluster forbidden response has a 3xx status code
func (o *GetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration get cluster forbidden response has a 4xx status code
func (o *GetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node cluster configuration get cluster forbidden response has a 5xx status code
func (o *GetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node cluster configuration get cluster forbidden response a status code equal to that given
func (o *GetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge node cluster configuration get cluster forbidden response
func (o *GetForbidden) Code() int {
	return 403
}

func (o *GetForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/id/{id}][%d] edgeNodeConfigurationGetForbidden  %+v", 403, o.Payload)
}

func (o *GetForbidden) String() string {
	return fmt.Sprintf("[GET /v1/cluster/id/{id}][%d] edgeNodeConfigurationGetForbidden  %+v", 403, o.Payload)
}

func (o *GetForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotFound creates a GetNotFound with default headers values
func NewGetNotFound() *GetNotFound {
	return &GetNotFound{}
}

/*
GetNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration get cluster not found response has a 2xx status code
func (o *GetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration get cluster not found response has a 3xx status code
func (o *GetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration get cluster not found response has a 4xx status code
func (o *GetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node cluster configuration get cluster not found response has a 5xx status code
func (o *GetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node cluster configuration get cluster not found response a status code equal to that given
func (o *GetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge node cluster configuration get cluster not found response
func (o *GetNotFound) Code() int {
	return 404
}

func (o *GetNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/id/{id}][%d] edgeNodeConfigurationGetNotFound  %+v", 404, o.Payload)
}

func (o *GetNotFound) String() string {
	return fmt.Sprintf("[GET /v1/cluster/id/{id}][%d] edgeNodeConfigurationGetNotFound  %+v", 404, o.Payload)
}

func (o *GetNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConflict creates a GetConflict with default headers values
func NewGetConflict() *GetConflict {
	return &GetConflict{}
}

/*
GetConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this edge-node cluster record will conflict with an already existing edge-node cluster record.
*/
type GetConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration get cluster conflict response has a 2xx status code
func (o *GetConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration get cluster conflict response has a 3xx status code
func (o *GetConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration get cluster conflict response has a 4xx status code
func (o *GetConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node cluster configuration get cluster conflict response has a 5xx status code
func (o *GetConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node cluster configuration get cluster conflict response a status code equal to that given
func (o *GetConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the edge node cluster configuration get cluster conflict response
func (o *GetConflict) Code() int {
	return 409
}

func (o *GetConflict) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/id/{id}][%d] edgeNodeConfigurationGetConflict  %+v", 409, o.Payload)
}

func (o *GetConflict) String() string {
	return fmt.Sprintf("[GET /v1/cluster/id/{id}][%d] edgeNodeConfigurationGetConflict  %+v", 409, o.Payload)
}

func (o *GetConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPreconditionFailed creates a GetPreconditionFailed with default headers values
func NewGetPreconditionFailed() *GetPreconditionFailed {
	return &GetPreconditionFailed{}
}

/*
GetPreconditionFailed describes a response with status code 412, with default header values.

Precondition failed. Some of preconditions haven't been met to start request processing.
*/
type GetPreconditionFailed struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration get cluster precondition failed response has a 2xx status code
func (o *GetPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration get cluster precondition failed response has a 3xx status code
func (o *GetPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration get cluster precondition failed response has a 4xx status code
func (o *GetPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node cluster configuration get cluster precondition failed response has a 5xx status code
func (o *GetPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node cluster configuration get cluster precondition failed response a status code equal to that given
func (o *GetPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the edge node cluster configuration get cluster precondition failed response
func (o *GetPreconditionFailed) Code() int {
	return 412
}

func (o *GetPreconditionFailed) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/id/{id}][%d] edgeNodeConfigurationGetPreconditionFailed  %+v", 412, o.Payload)
}

func (o *GetPreconditionFailed) String() string {
	return fmt.Sprintf("[GET /v1/cluster/id/{id}][%d] edgeNodeConfigurationGetPreconditionFailed  %+v", 412, o.Payload)
}

func (o *GetPreconditionFailed) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInternalServerError creates a GetInternalServerError with default headers values
func NewGetInternalServerError() *GetInternalServerError {
	return &GetInternalServerError{}
}

/*
GetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration get cluster internal server error response has a 2xx status code
func (o *GetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration get cluster internal server error response has a 3xx status code
func (o *GetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration get cluster internal server error response has a 4xx status code
func (o *GetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node cluster configuration get cluster internal server error response has a 5xx status code
func (o *GetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node cluster configuration get cluster internal server error response a status code equal to that given
func (o *GetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge node cluster configuration get cluster internal server error response
func (o *GetInternalServerError) Code() int {
	return 500
}

func (o *GetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/id/{id}][%d] edgeNodeConfigurationGetInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/cluster/id/{id}][%d] edgeNodeConfigurationGetInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGatewayTimeout creates a GetGatewayTimeout with default headers values
func NewGetGatewayTimeout() *GetGatewayTimeout {
	return &GetGatewayTimeout{}
}

/*
GetGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration get cluster gateway timeout response has a 2xx status code
func (o *GetGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration get cluster gateway timeout response has a 3xx status code
func (o *GetGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration get cluster gateway timeout response has a 4xx status code
func (o *GetGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node cluster configuration get cluster gateway timeout response has a 5xx status code
func (o *GetGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node cluster configuration get cluster gateway timeout response a status code equal to that given
func (o *GetGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge node cluster configuration get cluster gateway timeout response
func (o *GetGatewayTimeout) Code() int {
	return 504
}

func (o *GetGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/id/{id}][%d] edgeNodeConfigurationGetGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/cluster/id/{id}][%d] edgeNodeConfigurationGetGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDefault creates a GetDefault with default headers values
func NewGetDefault(code int) *GetDefault {
	return &GetDefault{
		_statusCode: code,
	}
}

/*
GetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge node cluster configuration get cluster default response has a 2xx status code
func (o *GetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node cluster configuration get cluster default response has a 3xx status code
func (o *GetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node cluster configuration get cluster default response has a 4xx status code
func (o *GetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node cluster configuration get cluster default response has a 5xx status code
func (o *GetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node cluster configuration get cluster default response a status code equal to that given
func (o *GetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge node cluster configuration get cluster default response
func (o *GetDefault) Code() int {
	return o._statusCode
}

func (o *GetDefault) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/id/{id}][%d] _Get default  %+v", o._statusCode, o.Payload)
}

func (o *GetDefault) String() string {
	return fmt.Sprintf("[GET /v1/cluster/id/{id}][%d] _Get default  %+v", o._statusCode, o.Payload)
}

func (o *GetDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *GetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
