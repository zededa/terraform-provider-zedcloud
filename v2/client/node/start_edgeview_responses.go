package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// EdgeNodeConfigurationStartEdgeviewEdgeNodeReader is a Reader for the EdgeNodeConfigurationStartEdgeviewEdgeNode structure.
type EdgeNodeConfigurationStartEdgeviewEdgeNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationStartEdgeviewEdgeNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeConfigurationStartEdgeviewEdgeNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationStartEdgeviewEdgeNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeConfigurationStartEdgeviewEdgeNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeNodeConfigurationStartEdgeviewEdgeNodeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationStartEdgeviewEdgeNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationStartEdgeviewEdgeNodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationStartEdgeviewEdgeNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationStartEdgeviewEdgeNodeOK creates a EdgeNodeConfigurationStartEdgeviewEdgeNodeOK with default headers values
func NewEdgeNodeConfigurationStartEdgeviewEdgeNodeOK() *EdgeNodeConfigurationStartEdgeviewEdgeNodeOK {
	return &EdgeNodeConfigurationStartEdgeviewEdgeNodeOK{}
}

/*
EdgeNodeConfigurationStartEdgeviewEdgeNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeConfigurationStartEdgeviewEdgeNodeOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration start edgeview edge node o k response has a 2xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node configuration start edgeview edge node o k response has a 3xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration start edgeview edge node o k response has a 4xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration start edgeview edge node o k response has a 5xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration start edgeview edge node o k response a status code equal to that given
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeOK) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/enable][%d] edgeNodeConfigurationStartEdgeviewEdgeNodeOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeOK) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/enable][%d] edgeNodeConfigurationStartEdgeviewEdgeNodeOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStartEdgeviewEdgeNodeUnauthorized creates a EdgeNodeConfigurationStartEdgeviewEdgeNodeUnauthorized with default headers values
func NewEdgeNodeConfigurationStartEdgeviewEdgeNodeUnauthorized() *EdgeNodeConfigurationStartEdgeviewEdgeNodeUnauthorized {
	return &EdgeNodeConfigurationStartEdgeviewEdgeNodeUnauthorized{}
}

/*
EdgeNodeConfigurationStartEdgeviewEdgeNodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationStartEdgeviewEdgeNodeUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration start edgeview edge node unauthorized response has a 2xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration start edgeview edge node unauthorized response has a 3xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration start edgeview edge node unauthorized response has a 4xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration start edgeview edge node unauthorized response has a 5xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration start edgeview edge node unauthorized response a status code equal to that given
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/enable][%d] edgeNodeConfigurationStartEdgeviewEdgeNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/enable][%d] edgeNodeConfigurationStartEdgeviewEdgeNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStartEdgeviewEdgeNodeForbidden creates a EdgeNodeConfigurationStartEdgeviewEdgeNodeForbidden with default headers values
func NewEdgeNodeConfigurationStartEdgeviewEdgeNodeForbidden() *EdgeNodeConfigurationStartEdgeviewEdgeNodeForbidden {
	return &EdgeNodeConfigurationStartEdgeviewEdgeNodeForbidden{}
}

/*
EdgeNodeConfigurationStartEdgeviewEdgeNodeForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationStartEdgeviewEdgeNodeForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration start edgeview edge node forbidden response has a 2xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration start edgeview edge node forbidden response has a 3xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration start edgeview edge node forbidden response has a 4xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration start edgeview edge node forbidden response has a 5xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration start edgeview edge node forbidden response a status code equal to that given
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/enable][%d] edgeNodeConfigurationStartEdgeviewEdgeNodeForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/enable][%d] edgeNodeConfigurationStartEdgeviewEdgeNodeForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStartEdgeviewEdgeNodeNotFound creates a EdgeNodeConfigurationStartEdgeviewEdgeNodeNotFound with default headers values
func NewEdgeNodeConfigurationStartEdgeviewEdgeNodeNotFound() *EdgeNodeConfigurationStartEdgeviewEdgeNodeNotFound {
	return &EdgeNodeConfigurationStartEdgeviewEdgeNodeNotFound{}
}

/*
EdgeNodeConfigurationStartEdgeviewEdgeNodeNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeConfigurationStartEdgeviewEdgeNodeNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration start edgeview edge node not found response has a 2xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration start edgeview edge node not found response has a 3xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration start edgeview edge node not found response has a 4xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration start edgeview edge node not found response has a 5xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration start edgeview edge node not found response a status code equal to that given
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/enable][%d] edgeNodeConfigurationStartEdgeviewEdgeNodeNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/enable][%d] edgeNodeConfigurationStartEdgeviewEdgeNodeNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStartEdgeviewEdgeNodeConflict creates a EdgeNodeConfigurationStartEdgeviewEdgeNodeConflict with default headers values
func NewEdgeNodeConfigurationStartEdgeviewEdgeNodeConflict() *EdgeNodeConfigurationStartEdgeviewEdgeNodeConflict {
	return &EdgeNodeConfigurationStartEdgeviewEdgeNodeConflict{}
}

/*
EdgeNodeConfigurationStartEdgeviewEdgeNodeConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge node record.
*/
type EdgeNodeConfigurationStartEdgeviewEdgeNodeConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration start edgeview edge node conflict response has a 2xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration start edgeview edge node conflict response has a 3xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration start edgeview edge node conflict response has a 4xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration start edgeview edge node conflict response has a 5xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration start edgeview edge node conflict response a status code equal to that given
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeConflict) IsCode(code int) bool {
	return code == 409
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/enable][%d] edgeNodeConfigurationStartEdgeviewEdgeNodeConflict  %+v", 409, o.Payload)
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeConflict) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/enable][%d] edgeNodeConfigurationStartEdgeviewEdgeNodeConflict  %+v", 409, o.Payload)
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStartEdgeviewEdgeNodeInternalServerError creates a EdgeNodeConfigurationStartEdgeviewEdgeNodeInternalServerError with default headers values
func NewEdgeNodeConfigurationStartEdgeviewEdgeNodeInternalServerError() *EdgeNodeConfigurationStartEdgeviewEdgeNodeInternalServerError {
	return &EdgeNodeConfigurationStartEdgeviewEdgeNodeInternalServerError{}
}

/*
EdgeNodeConfigurationStartEdgeviewEdgeNodeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationStartEdgeviewEdgeNodeInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration start edgeview edge node internal server error response has a 2xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration start edgeview edge node internal server error response has a 3xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration start edgeview edge node internal server error response has a 4xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration start edgeview edge node internal server error response has a 5xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration start edgeview edge node internal server error response a status code equal to that given
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/enable][%d] edgeNodeConfigurationStartEdgeviewEdgeNodeInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/enable][%d] edgeNodeConfigurationStartEdgeviewEdgeNodeInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStartEdgeviewEdgeNodeGatewayTimeout creates a EdgeNodeConfigurationStartEdgeviewEdgeNodeGatewayTimeout with default headers values
func NewEdgeNodeConfigurationStartEdgeviewEdgeNodeGatewayTimeout() *EdgeNodeConfigurationStartEdgeviewEdgeNodeGatewayTimeout {
	return &EdgeNodeConfigurationStartEdgeviewEdgeNodeGatewayTimeout{}
}

/*
EdgeNodeConfigurationStartEdgeviewEdgeNodeGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationStartEdgeviewEdgeNodeGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration start edgeview edge node gateway timeout response has a 2xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration start edgeview edge node gateway timeout response has a 3xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration start edgeview edge node gateway timeout response has a 4xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration start edgeview edge node gateway timeout response has a 5xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration start edgeview edge node gateway timeout response a status code equal to that given
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/enable][%d] edgeNodeConfigurationStartEdgeviewEdgeNodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/enable][%d] edgeNodeConfigurationStartEdgeviewEdgeNodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStartEdgeviewEdgeNodeDefault creates a EdgeNodeConfigurationStartEdgeviewEdgeNodeDefault with default headers values
func NewEdgeNodeConfigurationStartEdgeviewEdgeNodeDefault(code int) *EdgeNodeConfigurationStartEdgeviewEdgeNodeDefault {
	return &EdgeNodeConfigurationStartEdgeviewEdgeNodeDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeConfigurationStartEdgeviewEdgeNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationStartEdgeviewEdgeNodeDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the edge node configuration start edgeview edge node default response
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node configuration start edgeview edge node default response has a 2xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node configuration start edgeview edge node default response has a 3xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node configuration start edgeview edge node default response has a 4xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node configuration start edgeview edge node default response has a 5xx status code
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node configuration start edgeview edge node default response a status code equal to that given
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/enable][%d] EdgeNodeConfiguration_StartEdgeviewEdgeNode default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeDefault) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/enable][%d] EdgeNodeConfiguration_StartEdgeviewEdgeNode default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
EdgeNodeConfigurationStartEdgeviewEdgeNodeBody Device debug knob configuration
//
// Device debug knob configuration request payload holds the device debug mode properties
swagger:model EdgeNodeConfigurationStartEdgeviewEdgeNodeBody
*/
type EdgeNodeConfigurationStartEdgeviewEdgeNodeBody struct {

	// debug knob flag
	DebugKnob bool `json:"debugKnob,omitempty"`

	// debug knob expiry status flag
	Expired bool `json:"expired,omitempty"`

	// debug expiry time in minutes
	Expiry string `json:"expiry,omitempty"`
}

// Validate validates this edge node configuration start edgeview edge node body
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this edge node configuration start edgeview edge node body based on context it is used
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *EdgeNodeConfigurationStartEdgeviewEdgeNodeBody) UnmarshalBinary(b []byte) error {
	var res EdgeNodeConfigurationStartEdgeviewEdgeNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}