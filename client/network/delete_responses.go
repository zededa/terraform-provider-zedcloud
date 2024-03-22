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

// EdgeNetworkConfigurationDeleteEdgeNetworkReader is a Reader for the EdgeNetworkConfigurationDeleteEdgeNetwork structure.
type EdgeNetworkConfigurationDeleteEdgeNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNetworkConfigurationDeleteEdgeNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNetworkConfigurationDeleteEdgeNetworkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNetworkConfigurationDeleteEdgeNetworkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNetworkConfigurationDeleteEdgeNetworkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNetworkConfigurationDeleteEdgeNetworkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNetworkConfigurationDeleteEdgeNetworkGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNetworkConfigurationDeleteEdgeNetworkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNetworkConfigurationDeleteEdgeNetworkOK creates a EdgeNetworkConfigurationDeleteEdgeNetworkOK with default headers values
func NewEdgeNetworkConfigurationDeleteEdgeNetworkOK() *EdgeNetworkConfigurationDeleteEdgeNetworkOK {
	return &EdgeNetworkConfigurationDeleteEdgeNetworkOK{}
}

/*
EdgeNetworkConfigurationDeleteEdgeNetworkOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNetworkConfigurationDeleteEdgeNetworkOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration delete edge network o k response has a 2xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge network configuration delete edge network o k response has a 3xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration delete edge network o k response has a 4xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network configuration delete edge network o k response has a 5xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network configuration delete edge network o k response a status code equal to that given
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge network configuration delete edge network o k response
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkOK) Code() int {
	return 200
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/networks/id/{id}][%d] edgeNetworkConfigurationDeleteEdgeNetworkOK  %+v", 200, o.Payload)
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkOK) String() string {
	return fmt.Sprintf("[DELETE /v1/networks/id/{id}][%d] edgeNetworkConfigurationDeleteEdgeNetworkOK  %+v", 200, o.Payload)
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationDeleteEdgeNetworkUnauthorized creates a EdgeNetworkConfigurationDeleteEdgeNetworkUnauthorized with default headers values
func NewEdgeNetworkConfigurationDeleteEdgeNetworkUnauthorized() *EdgeNetworkConfigurationDeleteEdgeNetworkUnauthorized {
	return &EdgeNetworkConfigurationDeleteEdgeNetworkUnauthorized{}
}

/*
EdgeNetworkConfigurationDeleteEdgeNetworkUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNetworkConfigurationDeleteEdgeNetworkUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration delete edge network unauthorized response has a 2xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration delete edge network unauthorized response has a 3xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration delete edge network unauthorized response has a 4xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network configuration delete edge network unauthorized response has a 5xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network configuration delete edge network unauthorized response a status code equal to that given
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge network configuration delete edge network unauthorized response
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkUnauthorized) Code() int {
	return 401
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/networks/id/{id}][%d] edgeNetworkConfigurationDeleteEdgeNetworkUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/networks/id/{id}][%d] edgeNetworkConfigurationDeleteEdgeNetworkUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationDeleteEdgeNetworkForbidden creates a EdgeNetworkConfigurationDeleteEdgeNetworkForbidden with default headers values
func NewEdgeNetworkConfigurationDeleteEdgeNetworkForbidden() *EdgeNetworkConfigurationDeleteEdgeNetworkForbidden {
	return &EdgeNetworkConfigurationDeleteEdgeNetworkForbidden{}
}

/*
EdgeNetworkConfigurationDeleteEdgeNetworkForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNetworkConfigurationDeleteEdgeNetworkForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration delete edge network forbidden response has a 2xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration delete edge network forbidden response has a 3xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration delete edge network forbidden response has a 4xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network configuration delete edge network forbidden response has a 5xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network configuration delete edge network forbidden response a status code equal to that given
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge network configuration delete edge network forbidden response
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkForbidden) Code() int {
	return 403
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/networks/id/{id}][%d] edgeNetworkConfigurationDeleteEdgeNetworkForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/networks/id/{id}][%d] edgeNetworkConfigurationDeleteEdgeNetworkForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationDeleteEdgeNetworkNotFound creates a EdgeNetworkConfigurationDeleteEdgeNetworkNotFound with default headers values
func NewEdgeNetworkConfigurationDeleteEdgeNetworkNotFound() *EdgeNetworkConfigurationDeleteEdgeNetworkNotFound {
	return &EdgeNetworkConfigurationDeleteEdgeNetworkNotFound{}
}

/*
EdgeNetworkConfigurationDeleteEdgeNetworkNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNetworkConfigurationDeleteEdgeNetworkNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration delete edge network not found response has a 2xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration delete edge network not found response has a 3xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration delete edge network not found response has a 4xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network configuration delete edge network not found response has a 5xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network configuration delete edge network not found response a status code equal to that given
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge network configuration delete edge network not found response
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkNotFound) Code() int {
	return 404
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/networks/id/{id}][%d] edgeNetworkConfigurationDeleteEdgeNetworkNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/networks/id/{id}][%d] edgeNetworkConfigurationDeleteEdgeNetworkNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationDeleteEdgeNetworkInternalServerError creates a EdgeNetworkConfigurationDeleteEdgeNetworkInternalServerError with default headers values
func NewEdgeNetworkConfigurationDeleteEdgeNetworkInternalServerError() *EdgeNetworkConfigurationDeleteEdgeNetworkInternalServerError {
	return &EdgeNetworkConfigurationDeleteEdgeNetworkInternalServerError{}
}

/*
EdgeNetworkConfigurationDeleteEdgeNetworkInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNetworkConfigurationDeleteEdgeNetworkInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration delete edge network internal server error response has a 2xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration delete edge network internal server error response has a 3xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration delete edge network internal server error response has a 4xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network configuration delete edge network internal server error response has a 5xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge network configuration delete edge network internal server error response a status code equal to that given
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge network configuration delete edge network internal server error response
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkInternalServerError) Code() int {
	return 500
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/networks/id/{id}][%d] edgeNetworkConfigurationDeleteEdgeNetworkInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/networks/id/{id}][%d] edgeNetworkConfigurationDeleteEdgeNetworkInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationDeleteEdgeNetworkGatewayTimeout creates a EdgeNetworkConfigurationDeleteEdgeNetworkGatewayTimeout with default headers values
func NewEdgeNetworkConfigurationDeleteEdgeNetworkGatewayTimeout() *EdgeNetworkConfigurationDeleteEdgeNetworkGatewayTimeout {
	return &EdgeNetworkConfigurationDeleteEdgeNetworkGatewayTimeout{}
}

/*
EdgeNetworkConfigurationDeleteEdgeNetworkGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNetworkConfigurationDeleteEdgeNetworkGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration delete edge network gateway timeout response has a 2xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration delete edge network gateway timeout response has a 3xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration delete edge network gateway timeout response has a 4xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network configuration delete edge network gateway timeout response has a 5xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge network configuration delete edge network gateway timeout response a status code equal to that given
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge network configuration delete edge network gateway timeout response
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/networks/id/{id}][%d] edgeNetworkConfigurationDeleteEdgeNetworkGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /v1/networks/id/{id}][%d] edgeNetworkConfigurationDeleteEdgeNetworkGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationDeleteEdgeNetworkDefault creates a EdgeNetworkConfigurationDeleteEdgeNetworkDefault with default headers values
func NewEdgeNetworkConfigurationDeleteEdgeNetworkDefault(code int) *EdgeNetworkConfigurationDeleteEdgeNetworkDefault {
	return &EdgeNetworkConfigurationDeleteEdgeNetworkDefault{
		_statusCode: code,
	}
}

/*
EdgeNetworkConfigurationDeleteEdgeNetworkDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNetworkConfigurationDeleteEdgeNetworkDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge network configuration delete edge network default response has a 2xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge network configuration delete edge network default response has a 3xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge network configuration delete edge network default response has a 4xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge network configuration delete edge network default response has a 5xx status code
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge network configuration delete edge network default response a status code equal to that given
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge network configuration delete edge network default response
func (o *EdgeNetworkConfigurationDeleteEdgeNetworkDefault) Code() int {
	return o._statusCode
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/networks/id/{id}][%d] EdgeNetworkConfiguration_DeleteEdgeNetwork default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/networks/id/{id}][%d] EdgeNetworkConfiguration_DeleteEdgeNetwork default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNetworkConfigurationDeleteEdgeNetworkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
