package application_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceReader is a Reader for the EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstance structure.
type EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceOK creates a EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceOK with default headers values
func NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceOK() *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceOK {
	return &EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceOK{}
}

/*
EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration refresh edge application instance o k response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge application instance configuration refresh edge application instance o k response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration refresh edge application instance o k response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance configuration refresh edge application instance o k response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration refresh edge application instance o k response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge application instance configuration refresh edge application instance o k response
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceOK) Code() int {
	return 200
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceOK) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/refresh][%d] edgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceOK) String() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/refresh][%d] edgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceUnauthorized creates a EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceUnauthorized with default headers values
func NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceUnauthorized() *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceUnauthorized {
	return &EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceUnauthorized{}
}

/*
EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration refresh edge application instance unauthorized response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration refresh edge application instance unauthorized response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration refresh edge application instance unauthorized response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance configuration refresh edge application instance unauthorized response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration refresh edge application instance unauthorized response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge application instance configuration refresh edge application instance unauthorized response
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceUnauthorized) Code() int {
	return 401
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/refresh][%d] edgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/refresh][%d] edgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceForbidden creates a EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceForbidden with default headers values
func NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceForbidden() *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceForbidden {
	return &EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceForbidden{}
}

/*
EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration refresh edge application instance forbidden response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration refresh edge application instance forbidden response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration refresh edge application instance forbidden response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance configuration refresh edge application instance forbidden response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration refresh edge application instance forbidden response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge application instance configuration refresh edge application instance forbidden response
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceForbidden) Code() int {
	return 403
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/refresh][%d] edgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/refresh][%d] edgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceNotFound creates a EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceNotFound with default headers values
func NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceNotFound() *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceNotFound {
	return &EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceNotFound{}
}

/*
EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration refresh edge application instance not found response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration refresh edge application instance not found response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration refresh edge application instance not found response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance configuration refresh edge application instance not found response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration refresh edge application instance not found response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge application instance configuration refresh edge application instance not found response
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceNotFound) Code() int {
	return 404
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/refresh][%d] edgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceNotFound  %+v", 404, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/refresh][%d] edgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceNotFound  %+v", 404, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceConflict creates a EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceConflict with default headers values
func NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceConflict() *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceConflict {
	return &EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceConflict{}
}

/*
EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge network record.
*/
type EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration refresh edge application instance conflict response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration refresh edge application instance conflict response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration refresh edge application instance conflict response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance configuration refresh edge application instance conflict response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration refresh edge application instance conflict response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the edge application instance configuration refresh edge application instance conflict response
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceConflict) Code() int {
	return 409
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/refresh][%d] edgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceConflict  %+v", 409, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceConflict) String() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/refresh][%d] edgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceConflict  %+v", 409, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceInternalServerError creates a EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceInternalServerError with default headers values
func NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceInternalServerError() *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceInternalServerError {
	return &EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceInternalServerError{}
}

/*
EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration refresh edge application instance internal server error response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration refresh edge application instance internal server error response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration refresh edge application instance internal server error response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance configuration refresh edge application instance internal server error response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application instance configuration refresh edge application instance internal server error response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge application instance configuration refresh edge application instance internal server error response
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceInternalServerError) Code() int {
	return 500
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/refresh][%d] edgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/refresh][%d] edgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceGatewayTimeout creates a EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceGatewayTimeout with default headers values
func NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceGatewayTimeout() *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceGatewayTimeout {
	return &EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceGatewayTimeout{}
}

/*
EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration refresh edge application instance gateway timeout response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration refresh edge application instance gateway timeout response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration refresh edge application instance gateway timeout response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance configuration refresh edge application instance gateway timeout response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application instance configuration refresh edge application instance gateway timeout response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge application instance configuration refresh edge application instance gateway timeout response
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/refresh][%d] edgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/refresh][%d] edgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceDefault creates a EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceDefault with default headers values
func NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceDefault(code int) *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceDefault {
	return &EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceDefault{
		_statusCode: code,
	}
}

/*
EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge application instance configuration refresh edge application instance default response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge application instance configuration refresh edge application instance default response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge application instance configuration refresh edge application instance default response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge application instance configuration refresh edge application instance default response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge application instance configuration refresh edge application instance default response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge application instance configuration refresh edge application instance default response
func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceDefault) Code() int {
	return o._statusCode
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/refresh][%d] EdgeApplicationInstanceConfiguration_RefreshEdgeApplicationInstance default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceDefault) String() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/refresh][%d] EdgeApplicationInstanceConfiguration_RefreshEdgeApplicationInstance default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}