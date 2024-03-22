package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/models"
)

// EdgeNetworkConfigurationGetEdgeNetworkByNameReader is a Reader for the EdgeNetworkConfigurationGetEdgeNetworkByName structure.
type EdgeNetworkConfigurationGetEdgeNetworkByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNetworkConfigurationGetEdgeNetworkByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNetworkConfigurationGetEdgeNetworkByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNetworkConfigurationGetEdgeNetworkByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNetworkConfigurationGetEdgeNetworkByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNetworkConfigurationGetEdgeNetworkByNameOK creates a EdgeNetworkConfigurationGetEdgeNetworkByNameOK with default headers values
func NewEdgeNetworkConfigurationGetEdgeNetworkByNameOK() *EdgeNetworkConfigurationGetEdgeNetworkByNameOK {
	return &EdgeNetworkConfigurationGetEdgeNetworkByNameOK{}
}

/*
EdgeNetworkConfigurationGetEdgeNetworkByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNetworkConfigurationGetEdgeNetworkByNameOK struct {
	Payload *models.Network
}

// IsSuccess returns true when this edge network configuration get edge network by name o k response has a 2xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge network configuration get edge network by name o k response has a 3xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration get edge network by name o k response has a 4xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network configuration get edge network by name o k response has a 5xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network configuration get edge network by name o k response a status code equal to that given
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge network configuration get edge network by name o k response
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameOK) Code() int {
	return 200
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] edgeNetworkConfigurationGetEdgeNetworkByNameOK  %+v", 200, o.Payload)
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] edgeNetworkConfigurationGetEdgeNetworkByNameOK  %+v", 200, o.Payload)
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameOK) GetPayload() *models.Network {
	return o.Payload
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Network)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized creates a EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized with default headers values
func NewEdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized() *EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized {
	return &EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized{}
}

/*
EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration get edge network by name unauthorized response has a 2xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration get edge network by name unauthorized response has a 3xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration get edge network by name unauthorized response has a 4xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network configuration get edge network by name unauthorized response has a 5xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network configuration get edge network by name unauthorized response a status code equal to that given
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge network configuration get edge network by name unauthorized response
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized) Code() int {
	return 401
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] edgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] edgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationGetEdgeNetworkByNameForbidden creates a EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden with default headers values
func NewEdgeNetworkConfigurationGetEdgeNetworkByNameForbidden() *EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden {
	return &EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden{}
}

/*
EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration get edge network by name forbidden response has a 2xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration get edge network by name forbidden response has a 3xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration get edge network by name forbidden response has a 4xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network configuration get edge network by name forbidden response has a 5xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network configuration get edge network by name forbidden response a status code equal to that given
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge network configuration get edge network by name forbidden response
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden) Code() int {
	return 403
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] edgeNetworkConfigurationGetEdgeNetworkByNameForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] edgeNetworkConfigurationGetEdgeNetworkByNameForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationGetEdgeNetworkByNameNotFound creates a EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound with default headers values
func NewEdgeNetworkConfigurationGetEdgeNetworkByNameNotFound() *EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound {
	return &EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound{}
}

/*
EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration get edge network by name not found response has a 2xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration get edge network by name not found response has a 3xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration get edge network by name not found response has a 4xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network configuration get edge network by name not found response has a 5xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network configuration get edge network by name not found response a status code equal to that given
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge network configuration get edge network by name not found response
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound) Code() int {
	return 404
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] edgeNetworkConfigurationGetEdgeNetworkByNameNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] edgeNetworkConfigurationGetEdgeNetworkByNameNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError creates a EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError with default headers values
func NewEdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError() *EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError {
	return &EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError{}
}

/*
EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration get edge network by name internal server error response has a 2xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration get edge network by name internal server error response has a 3xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration get edge network by name internal server error response has a 4xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network configuration get edge network by name internal server error response has a 5xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge network configuration get edge network by name internal server error response a status code equal to that given
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge network configuration get edge network by name internal server error response
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError) Code() int {
	return 500
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] edgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] edgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout creates a EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout with default headers values
func NewEdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout() *EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout {
	return &EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout{}
}

/*
EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration get edge network by name gateway timeout response has a 2xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration get edge network by name gateway timeout response has a 3xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration get edge network by name gateway timeout response has a 4xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network configuration get edge network by name gateway timeout response has a 5xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge network configuration get edge network by name gateway timeout response a status code equal to that given
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge network configuration get edge network by name gateway timeout response
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] edgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] edgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationGetEdgeNetworkByNameDefault creates a EdgeNetworkConfigurationGetEdgeNetworkByNameDefault with default headers values
func NewEdgeNetworkConfigurationGetEdgeNetworkByNameDefault(code int) *EdgeNetworkConfigurationGetEdgeNetworkByNameDefault {
	return &EdgeNetworkConfigurationGetEdgeNetworkByNameDefault{
		_statusCode: code,
	}
}

/*
EdgeNetworkConfigurationGetEdgeNetworkByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNetworkConfigurationGetEdgeNetworkByNameDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge network configuration get edge network by name default response has a 2xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge network configuration get edge network by name default response has a 3xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge network configuration get edge network by name default response has a 4xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge network configuration get edge network by name default response has a 5xx status code
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge network configuration get edge network by name default response a status code equal to that given
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge network configuration get edge network by name default response
func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameDefault) Code() int {
	return o._statusCode
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] EdgeNetworkConfiguration_GetEdgeNetworkByName default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/networks/name/{name}][%d] EdgeNetworkConfiguration_GetEdgeNetworkByName default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNetworkConfigurationGetEdgeNetworkByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
