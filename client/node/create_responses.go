package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// EdgeNodeConfigurationCreateEdgeNodeReader is a Reader for the EdgeNodeConfigurationCreateEdgeNode structure.
type EdgeNodeConfigurationCreateEdgeNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationCreateEdgeNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationCreateEdgeNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeNodeConfigurationCreateEdgeNodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEdgeNodeConfigurationCreateEdgeNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationCreateEdgeNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeNodeConfigurationCreateEdgeNodeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationCreateEdgeNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationCreateEdgeNodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationCreateEdgeNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationCreateEdgeNodeOK creates a EdgeNodeConfigurationCreateEdgeNodeOK with default headers values
func NewEdgeNodeConfigurationCreateEdgeNodeOK() *EdgeNodeConfigurationCreateEdgeNodeOK {
	return &EdgeNodeConfigurationCreateEdgeNodeOK{}
}

/*
EdgeNodeConfigurationCreateEdgeNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeConfigurationCreateEdgeNodeOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration create edge node o k response has a 2xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node configuration create edge node o k response has a 3xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration create edge node o k response has a 4xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration create edge node o k response has a 5xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration create edge node o k response a status code equal to that given
func (o *EdgeNodeConfigurationCreateEdgeNodeOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeConfigurationCreateEdgeNodeOK) Error() string {
	return fmt.Sprintf("[POST /v1/devices][%d] Create EdgeNode: OK  \n%+v", 200, spew.Sdump(o.Payload.Error))
}

func (o *EdgeNodeConfigurationCreateEdgeNodeOK) String() string {
	return fmt.Sprintf("[POST /v1/devices][%d] Create EdgeNode: OK  \n%+v", 200, spew.Sdump(o.Payload.Error))
}

func (o *EdgeNodeConfigurationCreateEdgeNodeOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationCreateEdgeNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationCreateEdgeNodeBadRequest creates a EdgeNodeConfigurationCreateEdgeNodeBadRequest with default headers values
func NewEdgeNodeConfigurationCreateEdgeNodeBadRequest() *EdgeNodeConfigurationCreateEdgeNodeBadRequest {
	return &EdgeNodeConfigurationCreateEdgeNodeBadRequest{}
}

/*
EdgeNodeConfigurationCreateEdgeNodeBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type EdgeNodeConfigurationCreateEdgeNodeBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration create edge node bad request response has a 2xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration create edge node bad request response has a 3xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration create edge node bad request response has a 4xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration create edge node bad request response has a 5xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration create edge node bad request response a status code equal to that given
func (o *EdgeNodeConfigurationCreateEdgeNodeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EdgeNodeConfigurationCreateEdgeNodeBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/devices][%d] Create EdgeNode: BadRequest\n%+v", 400, spew.Sdump(o.Payload.Error))
}

func (o *EdgeNodeConfigurationCreateEdgeNodeBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/devices][%d] Create EdgeNode: BadRequest\n%+v", 400, spew.Sdump(o.Payload.Error))
}

func (o *EdgeNodeConfigurationCreateEdgeNodeBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationCreateEdgeNodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationCreateEdgeNodeUnauthorized creates a EdgeNodeConfigurationCreateEdgeNodeUnauthorized with default headers values
func NewEdgeNodeConfigurationCreateEdgeNodeUnauthorized() *EdgeNodeConfigurationCreateEdgeNodeUnauthorized {
	return &EdgeNodeConfigurationCreateEdgeNodeUnauthorized{}
}

/*
EdgeNodeConfigurationCreateEdgeNodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationCreateEdgeNodeUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration create edge node unauthorized response has a 2xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration create edge node unauthorized response has a 3xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration create edge node unauthorized response has a 4xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration create edge node unauthorized response has a 5xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration create edge node unauthorized response a status code equal to that given
func (o *EdgeNodeConfigurationCreateEdgeNodeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeConfigurationCreateEdgeNodeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/devices][%d] Create EdgeNode: Unauthorized  \n%+v", 401, spew.Sdump(o.Payload.Error))
}

func (o *EdgeNodeConfigurationCreateEdgeNodeUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/devices][%d] Create EdgeNode: Unauthorized  \n%+v", 401, spew.Sdump(o.Payload.Error))
}

func (o *EdgeNodeConfigurationCreateEdgeNodeUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationCreateEdgeNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationCreateEdgeNodeForbidden creates a EdgeNodeConfigurationCreateEdgeNodeForbidden with default headers values
func NewEdgeNodeConfigurationCreateEdgeNodeForbidden() *EdgeNodeConfigurationCreateEdgeNodeForbidden {
	return &EdgeNodeConfigurationCreateEdgeNodeForbidden{}
}

/*
EdgeNodeConfigurationCreateEdgeNodeForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationCreateEdgeNodeForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration create edge node forbidden response has a 2xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration create edge node forbidden response has a 3xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration create edge node forbidden response has a 4xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration create edge node forbidden response has a 5xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration create edge node forbidden response a status code equal to that given
func (o *EdgeNodeConfigurationCreateEdgeNodeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeConfigurationCreateEdgeNodeForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/devices][%d] Create EdgeNode: Forbidden  \n%+v", 403, spew.Sdump(o.Payload.Error))
}

func (o *EdgeNodeConfigurationCreateEdgeNodeForbidden) String() string {
	return fmt.Sprintf("[POST /v1/devices][%d] Create EdgeNode: Forbidden  \n%+v", 403, spew.Sdump(o.Payload.Error))
}

func (o *EdgeNodeConfigurationCreateEdgeNodeForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationCreateEdgeNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationCreateEdgeNodeConflict creates a EdgeNodeConfigurationCreateEdgeNodeConflict with default headers values
func NewEdgeNodeConfigurationCreateEdgeNodeConflict() *EdgeNodeConfigurationCreateEdgeNodeConflict {
	return &EdgeNodeConfigurationCreateEdgeNodeConflict{}
}

/*
EdgeNodeConfigurationCreateEdgeNodeConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this edge node record will conflict with an already existing edge node record.
*/
type EdgeNodeConfigurationCreateEdgeNodeConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration create edge node conflict response has a 2xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration create edge node conflict response has a 3xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration create edge node conflict response has a 4xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration create edge node conflict response has a 5xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration create edge node conflict response a status code equal to that given
func (o *EdgeNodeConfigurationCreateEdgeNodeConflict) IsCode(code int) bool {
	return code == 409
}

func (o *EdgeNodeConfigurationCreateEdgeNodeConflict) Error() string {
	return fmt.Sprintf("[POST /v1/devices][%d] Create EdgeNode: Conflict  \n%+v", 409, spew.Sdump(o.Payload.Error))
}

func (o *EdgeNodeConfigurationCreateEdgeNodeConflict) String() string {
	return fmt.Sprintf("[POST /v1/devices][%d] Create EdgeNode: Conflict  \n%+v", 409, spew.Sdump(o.Payload.Error))
}

func (o *EdgeNodeConfigurationCreateEdgeNodeConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationCreateEdgeNodeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationCreateEdgeNodeInternalServerError creates a EdgeNodeConfigurationCreateEdgeNodeInternalServerError with default headers values
func NewEdgeNodeConfigurationCreateEdgeNodeInternalServerError() *EdgeNodeConfigurationCreateEdgeNodeInternalServerError {
	return &EdgeNodeConfigurationCreateEdgeNodeInternalServerError{}
}

/*
EdgeNodeConfigurationCreateEdgeNodeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationCreateEdgeNodeInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration create edge node internal server error response has a 2xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration create edge node internal server error response has a 3xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration create edge node internal server error response has a 4xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration create edge node internal server error response has a 5xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration create edge node internal server error response a status code equal to that given
func (o *EdgeNodeConfigurationCreateEdgeNodeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeConfigurationCreateEdgeNodeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/devices][%d] Create EdgeNode: Internal Server Error  \n%+v", 500, spew.Sdump(o.Payload.Error))
}

func (o *EdgeNodeConfigurationCreateEdgeNodeInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/devices][%d] Create EdgeNode: Internal Server Error  \n%+v", 500, spew.Sdump(o.Payload.Error))
}

func (o *EdgeNodeConfigurationCreateEdgeNodeInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationCreateEdgeNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationCreateEdgeNodeGatewayTimeout creates a EdgeNodeConfigurationCreateEdgeNodeGatewayTimeout with default headers values
func NewEdgeNodeConfigurationCreateEdgeNodeGatewayTimeout() *EdgeNodeConfigurationCreateEdgeNodeGatewayTimeout {
	return &EdgeNodeConfigurationCreateEdgeNodeGatewayTimeout{}
}

/*
EdgeNodeConfigurationCreateEdgeNodeGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationCreateEdgeNodeGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration create edge node gateway timeout response has a 2xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration create edge node gateway timeout response has a 3xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration create edge node gateway timeout response has a 4xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration create edge node gateway timeout response has a 5xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration create edge node gateway timeout response a status code equal to that given
func (o *EdgeNodeConfigurationCreateEdgeNodeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeConfigurationCreateEdgeNodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/devices][%d] Create EdgeNode: Gateway Timeout  \n%+v", 504, spew.Sdump(o.Payload.Error))
}

func (o *EdgeNodeConfigurationCreateEdgeNodeGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/devices][%d] Create EdgeNode: Gateway Timeout  \n%+v", 504, spew.Sdump(o.Payload.Error))
}

func (o *EdgeNodeConfigurationCreateEdgeNodeGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationCreateEdgeNodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationCreateEdgeNodeDefault creates a EdgeNodeConfigurationCreateEdgeNodeDefault with default headers values
func NewEdgeNodeConfigurationCreateEdgeNodeDefault(code int) *EdgeNodeConfigurationCreateEdgeNodeDefault {
	return &EdgeNodeConfigurationCreateEdgeNodeDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeConfigurationCreateEdgeNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationCreateEdgeNodeDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the edge node configuration create edge node default response
func (o *EdgeNodeConfigurationCreateEdgeNodeDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node configuration create edge node default response has a 2xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node configuration create edge node default response has a 3xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node configuration create edge node default response has a 4xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node configuration create edge node default response has a 5xx status code
func (o *EdgeNodeConfigurationCreateEdgeNodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node configuration create edge node default response a status code equal to that given
func (o *EdgeNodeConfigurationCreateEdgeNodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeConfigurationCreateEdgeNodeDefault) Error() string {
	return fmt.Sprintf("[POST /v1/devices][%d] Create EdgeNode: Default  \n%+v", o._statusCode, spew.Sdump(o.Payload))
}

func (o *EdgeNodeConfigurationCreateEdgeNodeDefault) String() string {
	return fmt.Sprintf("[POST /v1/devices][%d] Create EdgeNode: Default  \n%+v", o._statusCode, spew.Sdump(o.Payload))
}

func (o *EdgeNodeConfigurationCreateEdgeNodeDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationCreateEdgeNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
