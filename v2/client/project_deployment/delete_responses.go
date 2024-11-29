// Code generated by go-swagger; DO NOT EDIT.

package project_deployment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// DeleteReader is a Reader for the Delete structure.
type DeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteOK creates a ResourceGroupDeleteResourceGroupV2OK with default headers values
func NewDeleteOK() *DeleteOK {
	return &DeleteOK{}
}

/*
DeleteOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this resource group delete resource group v2 o k response has a 2xx status code
func (o *DeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resource group delete resource group v2 o k response has a 3xx status code
func (o *DeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group delete resource group v2 o k response has a 4xx status code
func (o *DeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource group delete resource group v2 o k response has a 5xx status code
func (o *DeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group delete resource group v2 o k response a status code equal to that given
func (o *DeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the resource group delete resource group v2 o k response
func (o *DeleteOK) Code() int {
	return 200
}

func (o *DeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments/id/{id}][%d] projectDeploymentDeleteOK %s", 200, payload)
}

func (o *DeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments/id/{id}][%d] projectDeploymentDeleteOK %s", 200, payload)
}

func (o *DeleteOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUnauthorized creates a DeleteUnauthorized with default headers values
func NewDeleteUnauthorized() *DeleteUnauthorized {
	return &DeleteUnauthorized{}
}

/*
DeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type DeleteUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this resource group delete resource group v2 unauthorized response has a 2xx status code
func (o *DeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group delete resource group v2 unauthorized response has a 3xx status code
func (o *DeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group delete resource group v2 unauthorized response has a 4xx status code
func (o *DeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this resource group delete resource group v2 unauthorized response has a 5xx status code
func (o *DeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group delete resource group v2 unauthorized response a status code equal to that given
func (o *DeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the resource group delete resource group v2 unauthorized response
func (o *DeleteUnauthorized) Code() int {
	return 401
}

func (o *DeleteUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments/id/{id}][%d] projectDeploymentDeleteUnauthorized %s", 401, payload)
}

func (o *DeleteUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments/id/{id}][%d] projectDeploymentDeleteUnauthorized %s", 401, payload)
}

func (o *DeleteUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteForbidden creates a DeleteForbidden with default headers values
func NewDeleteForbidden() *DeleteForbidden {
	return &DeleteForbidden{}
}

/*
DeleteForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type DeleteForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this resource group delete resource group v2 forbidden response has a 2xx status code
func (o *DeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group delete resource group v2 forbidden response has a 3xx status code
func (o *DeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group delete resource group v2 forbidden response has a 4xx status code
func (o *DeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this resource group delete resource group v2 forbidden response has a 5xx status code
func (o *DeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this resource group delete resource group v2 forbidden response a status code equal to that given
func (o *DeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the resource group delete resource group v2 forbidden response
func (o *DeleteForbidden) Code() int {
	return 403
}

func (o *DeleteForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments/id/{id}][%d] projectDeploymentDeleteForbidden %s", 403, payload)
}

func (o *DeleteForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments/id/{id}][%d] projectDeploymentDeleteForbidden %s", 403, payload)
}

func (o *DeleteForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInternalServerError creates a DeleteInternalServerError with default headers values
func NewDeleteInternalServerError() *DeleteInternalServerError {
	return &DeleteInternalServerError{}
}

/*
DeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type DeleteInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this resource group delete resource group v2 internal server error response has a 2xx status code
func (o *DeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group delete resource group v2 internal server error response has a 3xx status code
func (o *DeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group delete resource group v2 internal server error response has a 4xx status code
func (o *DeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource group delete resource group v2 internal server error response has a 5xx status code
func (o *DeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this resource group delete resource group v2 internal server error response a status code equal to that given
func (o *DeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the resource group delete resource group v2 internal server error response
func (o *DeleteInternalServerError) Code() int {
	return 500
}

func (o *DeleteInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments/id/{id}][%d] projectDeploymentDeleteInternalServerError %s", 500, payload)
}

func (o *DeleteInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments/id/{id}][%d] projectDeploymentDeleteInternalServerError %s", 500, payload)
}

func (o *DeleteInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGatewayTimeout creates a DeleteGatewayTimeout with default headers values
func NewDeleteGatewayTimeout() *DeleteGatewayTimeout {
	return &DeleteGatewayTimeout{}
}

/*
DeleteGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type DeleteGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this resource group delete resource group v2 gateway timeout response has a 2xx status code
func (o *DeleteGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource group delete resource group v2 gateway timeout response has a 3xx status code
func (o *DeleteGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource group delete resource group v2 gateway timeout response has a 4xx status code
func (o *DeleteGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource group delete resource group v2 gateway timeout response has a 5xx status code
func (o *DeleteGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this resource group delete resource group v2 gateway timeout response a status code equal to that given
func (o *DeleteGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the resource group delete resource group v2 gateway timeout response
func (o *DeleteGatewayTimeout) Code() int {
	return 504
}

func (o *DeleteGatewayTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments/id/{id}][%d] projectDeploymentDeleteGatewayTimeout %s", 504, payload)
}

func (o *DeleteGatewayTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments/id/{id}][%d] projectDeploymentDeleteGatewayTimeout %s", 504, payload)
}

func (o *DeleteGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDefault creates a DeleteDefault with default headers values
func NewDeleteDefault(code int) *DeleteDefault {
	return &DeleteDefault{
		_statusCode: code,
	}
}

/*
DeleteDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this resource group delete resource group v2 default response has a 2xx status code
func (o *DeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this resource group delete resource group v2 default response has a 3xx status code
func (o *DeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this resource group delete resource group v2 default response has a 4xx status code
func (o *DeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this resource group delete resource group v2 default response has a 5xx status code
func (o *DeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this resource group delete resource group v2 default response a status code equal to that given
func (o *DeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the resource group delete resource group v2 default response
func (o *DeleteDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments/id/{id}][%d] projectDeploymentDelete default %s", o._statusCode, payload)
}

func (o *DeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v2/projects/id/{projectId}/deployments/id/{id}][%d] projectDeploymentDelete default %s", o._statusCode, payload)
}

func (o *DeleteDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *DeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
