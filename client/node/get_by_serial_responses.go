package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/models"
)

// EdgeNodeConfigurationGetEdgeNodeBySerialReader is a Reader for the EdgeNodeConfigurationGetEdgeNodeBySerial structure.
type EdgeNodeConfigurationGetEdgeNodeBySerialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationGetEdgeNodeBySerialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeConfigurationGetEdgeNodeBySerialUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationGetEdgeNodeBySerialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeConfigurationGetEdgeNodeBySerialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationGetEdgeNodeBySerialInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationGetEdgeNodeBySerialGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationGetEdgeNodeBySerialDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationGetEdgeNodeBySerialOK creates a EdgeNodeConfigurationGetEdgeNodeBySerialOK with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeBySerialOK() *EdgeNodeConfigurationGetEdgeNodeBySerialOK {
	return &EdgeNodeConfigurationGetEdgeNodeBySerialOK{}
}

/*
EdgeNodeConfigurationGetEdgeNodeBySerialOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeConfigurationGetEdgeNodeBySerialOK struct {
	Payload *models.Node
}

// IsSuccess returns true when this edge node configuration get edge node by serial o k response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node configuration get edge node by serial o k response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node by serial o k response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration get edge node by serial o k response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration get edge node by serial o k response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/serial/{serialno}][%d] edgeNodeConfigurationGetEdgeNodeBySerialOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialOK) String() string {
	return fmt.Sprintf("[GET /v1/devices/serial/{serialno}][%d] edgeNodeConfigurationGetEdgeNodeBySerialOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialOK) GetPayload() *models.Node {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Node)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeBySerialUnauthorized creates a EdgeNodeConfigurationGetEdgeNodeBySerialUnauthorized with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeBySerialUnauthorized() *EdgeNodeConfigurationGetEdgeNodeBySerialUnauthorized {
	return &EdgeNodeConfigurationGetEdgeNodeBySerialUnauthorized{}
}

/*
EdgeNodeConfigurationGetEdgeNodeBySerialUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationGetEdgeNodeBySerialUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get edge node by serial unauthorized response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get edge node by serial unauthorized response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node by serial unauthorized response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration get edge node by serial unauthorized response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration get edge node by serial unauthorized response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/serial/{serialno}][%d] edgeNodeConfigurationGetEdgeNodeBySerialUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/devices/serial/{serialno}][%d] edgeNodeConfigurationGetEdgeNodeBySerialUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeBySerialForbidden creates a EdgeNodeConfigurationGetEdgeNodeBySerialForbidden with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeBySerialForbidden() *EdgeNodeConfigurationGetEdgeNodeBySerialForbidden {
	return &EdgeNodeConfigurationGetEdgeNodeBySerialForbidden{}
}

/*
EdgeNodeConfigurationGetEdgeNodeBySerialForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationGetEdgeNodeBySerialForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get edge node by serial forbidden response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get edge node by serial forbidden response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node by serial forbidden response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration get edge node by serial forbidden response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration get edge node by serial forbidden response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/serial/{serialno}][%d] edgeNodeConfigurationGetEdgeNodeBySerialForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialForbidden) String() string {
	return fmt.Sprintf("[GET /v1/devices/serial/{serialno}][%d] edgeNodeConfigurationGetEdgeNodeBySerialForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeBySerialNotFound creates a EdgeNodeConfigurationGetEdgeNodeBySerialNotFound with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeBySerialNotFound() *EdgeNodeConfigurationGetEdgeNodeBySerialNotFound {
	return &EdgeNodeConfigurationGetEdgeNodeBySerialNotFound{}
}

/*
EdgeNodeConfigurationGetEdgeNodeBySerialNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeConfigurationGetEdgeNodeBySerialNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get edge node by serial not found response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get edge node by serial not found response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node by serial not found response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration get edge node by serial not found response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration get edge node by serial not found response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/serial/{serialno}][%d] edgeNodeConfigurationGetEdgeNodeBySerialNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialNotFound) String() string {
	return fmt.Sprintf("[GET /v1/devices/serial/{serialno}][%d] edgeNodeConfigurationGetEdgeNodeBySerialNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeBySerialInternalServerError creates a EdgeNodeConfigurationGetEdgeNodeBySerialInternalServerError with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeBySerialInternalServerError() *EdgeNodeConfigurationGetEdgeNodeBySerialInternalServerError {
	return &EdgeNodeConfigurationGetEdgeNodeBySerialInternalServerError{}
}

/*
EdgeNodeConfigurationGetEdgeNodeBySerialInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationGetEdgeNodeBySerialInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get edge node by serial internal server error response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get edge node by serial internal server error response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node by serial internal server error response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration get edge node by serial internal server error response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration get edge node by serial internal server error response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/serial/{serialno}][%d] edgeNodeConfigurationGetEdgeNodeBySerialInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/devices/serial/{serialno}][%d] edgeNodeConfigurationGetEdgeNodeBySerialInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeBySerialGatewayTimeout creates a EdgeNodeConfigurationGetEdgeNodeBySerialGatewayTimeout with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeBySerialGatewayTimeout() *EdgeNodeConfigurationGetEdgeNodeBySerialGatewayTimeout {
	return &EdgeNodeConfigurationGetEdgeNodeBySerialGatewayTimeout{}
}

/*
EdgeNodeConfigurationGetEdgeNodeBySerialGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationGetEdgeNodeBySerialGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get edge node by serial gateway timeout response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get edge node by serial gateway timeout response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node by serial gateway timeout response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration get edge node by serial gateway timeout response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration get edge node by serial gateway timeout response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/serial/{serialno}][%d] edgeNodeConfigurationGetEdgeNodeBySerialGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/devices/serial/{serialno}][%d] edgeNodeConfigurationGetEdgeNodeBySerialGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeBySerialDefault creates a EdgeNodeConfigurationGetEdgeNodeBySerialDefault with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeBySerialDefault(code int) *EdgeNodeConfigurationGetEdgeNodeBySerialDefault {
	return &EdgeNodeConfigurationGetEdgeNodeBySerialDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeConfigurationGetEdgeNodeBySerialDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationGetEdgeNodeBySerialDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the edge node configuration get edge node by serial default response
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node configuration get edge node by serial default response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node configuration get edge node by serial default response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node configuration get edge node by serial default response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node configuration get edge node by serial default response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node configuration get edge node by serial default response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeBySerialDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialDefault) Error() string {
	return fmt.Sprintf("[GET /v1/devices/serial/{serialno}][%d] EdgeNodeConfiguration_GetEdgeNodeBySerial default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialDefault) String() string {
	return fmt.Sprintf("[GET /v1/devices/serial/{serialno}][%d] EdgeNodeConfiguration_GetEdgeNodeBySerial default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeBySerialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
