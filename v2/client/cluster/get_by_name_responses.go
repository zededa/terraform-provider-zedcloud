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

// GetByNameReader is a Reader for the GetByName structure.
type GetByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetByNameConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewGetByNamePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetByNameOK creates a GetByNameOK with default headers values
func NewGetByNameOK() *GetByNameOK {
	return &GetByNameOK{}
}

/*
GetByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetByNameOK struct {
	Payload *models.ClusterConfigSummary
}

// IsSuccess returns true when this edge node cluster configuration get cluster by name o k response has a 2xx status code
func (o *GetByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node cluster configuration get cluster by name o k response has a 3xx status code
func (o *GetByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration get cluster by name o k response has a 4xx status code
func (o *GetByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node cluster configuration get cluster by name o k response has a 5xx status code
func (o *GetByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node cluster configuration get cluster by name o k response a status code equal to that given
func (o *GetByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge node cluster configuration get cluster by name o k response
func (o *GetByNameOK) Code() int {
	return 200
}

func (o *GetByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/name/{name}][%d] edgeNodeConfigurationGetByNameOK  %+v", 200, o.Payload)
}

func (o *GetByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/cluster/name/{name}][%d] edgeNodeConfigurationGetByNameOK  %+v", 200, o.Payload)
}

func (o *GetByNameOK) GetPayload() *models.ClusterConfigSummary {
	return o.Payload
}

func (o *GetByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterConfigSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetByNameUnauthorized creates a GetByNameUnauthorized with default headers values
func NewGetByNameUnauthorized() *GetByNameUnauthorized {
	return &GetByNameUnauthorized{}
}

/*
GetByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetByNameUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration get cluster by name unauthorized response has a 2xx status code
func (o *GetByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration get cluster by name unauthorized response has a 3xx status code
func (o *GetByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration get cluster by name unauthorized response has a 4xx status code
func (o *GetByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node cluster configuration get cluster by name unauthorized response has a 5xx status code
func (o *GetByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node cluster configuration get cluster by name unauthorized response a status code equal to that given
func (o *GetByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge node cluster configuration get cluster by name unauthorized response
func (o *GetByNameUnauthorized) Code() int {
	return 401
}

func (o *GetByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/name/{name}][%d] edgeNodeConfigurationGetByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *GetByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/cluster/name/{name}][%d] edgeNodeConfigurationGetByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *GetByNameUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetByNameForbidden creates a GetByNameForbidden with default headers values
func NewGetByNameForbidden() *GetByNameForbidden {
	return &GetByNameForbidden{}
}

/*
GetByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type GetByNameForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration get cluster by name forbidden response has a 2xx status code
func (o *GetByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration get cluster by name forbidden response has a 3xx status code
func (o *GetByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration get cluster by name forbidden response has a 4xx status code
func (o *GetByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node cluster configuration get cluster by name forbidden response has a 5xx status code
func (o *GetByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node cluster configuration get cluster by name forbidden response a status code equal to that given
func (o *GetByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge node cluster configuration get cluster by name forbidden response
func (o *GetByNameForbidden) Code() int {
	return 403
}

func (o *GetByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/name/{name}][%d] edgeNodeConfigurationGetByNameForbidden  %+v", 403, o.Payload)
}

func (o *GetByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/cluster/name/{name}][%d] edgeNodeConfigurationGetByNameForbidden  %+v", 403, o.Payload)
}

func (o *GetByNameForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetByNameNotFound creates a GetByNameNotFound with default headers values
func NewGetByNameNotFound() *GetByNameNotFound {
	return &GetByNameNotFound{}
}

/*
GetByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetByNameNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration get cluster by name not found response has a 2xx status code
func (o *GetByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration get cluster by name not found response has a 3xx status code
func (o *GetByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration get cluster by name not found response has a 4xx status code
func (o *GetByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node cluster configuration get cluster by name not found response has a 5xx status code
func (o *GetByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node cluster configuration get cluster by name not found response a status code equal to that given
func (o *GetByNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge node cluster configuration get cluster by name not found response
func (o *GetByNameNotFound) Code() int {
	return 404
}

func (o *GetByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/name/{name}][%d] edgeNodeConfigurationGetByNameNotFound  %+v", 404, o.Payload)
}

func (o *GetByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/cluster/name/{name}][%d] edgeNodeConfigurationGetByNameNotFound  %+v", 404, o.Payload)
}

func (o *GetByNameNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetByNameConflict creates a GetByNameConflict with default headers values
func NewGetByNameConflict() *GetByNameConflict {
	return &GetByNameConflict{}
}

/*
GetByNameConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this edge-node cluster record will conflict with an already existing edge-node cluster record.
*/
type GetByNameConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration get cluster by name conflict response has a 2xx status code
func (o *GetByNameConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration get cluster by name conflict response has a 3xx status code
func (o *GetByNameConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration get cluster by name conflict response has a 4xx status code
func (o *GetByNameConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node cluster configuration get cluster by name conflict response has a 5xx status code
func (o *GetByNameConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node cluster configuration get cluster by name conflict response a status code equal to that given
func (o *GetByNameConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the edge node cluster configuration get cluster by name conflict response
func (o *GetByNameConflict) Code() int {
	return 409
}

func (o *GetByNameConflict) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/name/{name}][%d] edgeNodeConfigurationGetByNameConflict  %+v", 409, o.Payload)
}

func (o *GetByNameConflict) String() string {
	return fmt.Sprintf("[GET /v1/cluster/name/{name}][%d] edgeNodeConfigurationGetByNameConflict  %+v", 409, o.Payload)
}

func (o *GetByNameConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetByNameConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetByNamePreconditionFailed creates a GetByNamePreconditionFailed with default headers values
func NewGetByNamePreconditionFailed() *GetByNamePreconditionFailed {
	return &GetByNamePreconditionFailed{}
}

/*
GetByNamePreconditionFailed describes a response with status code 412, with default header values.

Precondition failed. Some of preconditions haven't been met to start request processing.
*/
type GetByNamePreconditionFailed struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration get cluster by name precondition failed response has a 2xx status code
func (o *GetByNamePreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration get cluster by name precondition failed response has a 3xx status code
func (o *GetByNamePreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration get cluster by name precondition failed response has a 4xx status code
func (o *GetByNamePreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node cluster configuration get cluster by name precondition failed response has a 5xx status code
func (o *GetByNamePreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node cluster configuration get cluster by name precondition failed response a status code equal to that given
func (o *GetByNamePreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the edge node cluster configuration get cluster by name precondition failed response
func (o *GetByNamePreconditionFailed) Code() int {
	return 412
}

func (o *GetByNamePreconditionFailed) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/name/{name}][%d] edgeNodeConfigurationGetByNamePreconditionFailed  %+v", 412, o.Payload)
}

func (o *GetByNamePreconditionFailed) String() string {
	return fmt.Sprintf("[GET /v1/cluster/name/{name}][%d] edgeNodeConfigurationGetByNamePreconditionFailed  %+v", 412, o.Payload)
}

func (o *GetByNamePreconditionFailed) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetByNamePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetByNameInternalServerError creates a GetByNameInternalServerError with default headers values
func NewGetByNameInternalServerError() *GetByNameInternalServerError {
	return &GetByNameInternalServerError{}
}

/*
GetByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetByNameInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration get cluster by name internal server error response has a 2xx status code
func (o *GetByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration get cluster by name internal server error response has a 3xx status code
func (o *GetByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration get cluster by name internal server error response has a 4xx status code
func (o *GetByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node cluster configuration get cluster by name internal server error response has a 5xx status code
func (o *GetByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node cluster configuration get cluster by name internal server error response a status code equal to that given
func (o *GetByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge node cluster configuration get cluster by name internal server error response
func (o *GetByNameInternalServerError) Code() int {
	return 500
}

func (o *GetByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/name/{name}][%d] edgeNodeConfigurationGetByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *GetByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/cluster/name/{name}][%d] edgeNodeConfigurationGetByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *GetByNameInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetByNameGatewayTimeout creates a GetByNameGatewayTimeout with default headers values
func NewGetByNameGatewayTimeout() *GetByNameGatewayTimeout {
	return &GetByNameGatewayTimeout{}
}

/*
GetByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetByNameGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node cluster configuration get cluster by name gateway timeout response has a 2xx status code
func (o *GetByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node cluster configuration get cluster by name gateway timeout response has a 3xx status code
func (o *GetByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node cluster configuration get cluster by name gateway timeout response has a 4xx status code
func (o *GetByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node cluster configuration get cluster by name gateway timeout response has a 5xx status code
func (o *GetByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node cluster configuration get cluster by name gateway timeout response a status code equal to that given
func (o *GetByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge node cluster configuration get cluster by name gateway timeout response
func (o *GetByNameGatewayTimeout) Code() int {
	return 504
}

func (o *GetByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/name/{name}][%d] edgeNodeConfigurationGetByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/cluster/name/{name}][%d] edgeNodeConfigurationGetByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetByNameGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetByNameDefault creates a GetByNameDefault with default headers values
func NewGetByNameDefault(code int) *GetByNameDefault {
	return &GetByNameDefault{
		_statusCode: code,
	}
}

/*
GetByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetByNameDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge node cluster configuration get cluster by name default response has a 2xx status code
func (o *GetByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node cluster configuration get cluster by name default response has a 3xx status code
func (o *GetByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node cluster configuration get cluster by name default response has a 4xx status code
func (o *GetByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node cluster configuration get cluster by name default response has a 5xx status code
func (o *GetByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node cluster configuration get cluster by name default response a status code equal to that given
func (o *GetByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge node cluster configuration get cluster by name default response
func (o *GetByNameDefault) Code() int {
	return o._statusCode
}

func (o *GetByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/cluster/name/{name}][%d] _GetByName default  %+v", o._statusCode, o.Payload)
}

func (o *GetByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/cluster/name/{name}][%d] _GetByName default  %+v", o._statusCode, o.Payload)
}

func (o *GetByNameDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *GetByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
