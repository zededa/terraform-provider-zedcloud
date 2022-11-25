// Code generated by go-swagger; DO NOT EDIT.

package resource_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// GetResourceGroupReader is a Reader for the GetResourceGroup structure.
type GetResourceGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourceGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetResourceGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetResourceGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResourceGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetResourceGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetResourceGroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResourceGroupOK creates a GetResourceGroupOK with default headers values
func NewGetResourceGroupOK() *GetResourceGroupOK {
	return &GetResourceGroupOK{}
}

/*
GetResourceGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetResourceGroupOK struct {
	Payload *models.Tag
}

// IsSuccess returns true when this get resource group o k response has a 2xx status code
func (o *GetResourceGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get resource group o k response has a 3xx status code
func (o *GetResourceGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource group o k response has a 4xx status code
func (o *GetResourceGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resource group o k response has a 5xx status code
func (o *GetResourceGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource group o k response a status code equal to that given
func (o *GetResourceGroupOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetResourceGroupOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] getResourceGroupOK  %+v", 200, o.Payload)
}

func (o *GetResourceGroupOK) String() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] getResourceGroupOK  %+v", 200, o.Payload)
}

func (o *GetResourceGroupOK) GetPayload() *models.Tag {
	return o.Payload
}

func (o *GetResourceGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceGroupUnauthorized creates a GetResourceGroupUnauthorized with default headers values
func NewGetResourceGroupUnauthorized() *GetResourceGroupUnauthorized {
	return &GetResourceGroupUnauthorized{}
}

/*
GetResourceGroupUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetResourceGroupUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get resource group unauthorized response has a 2xx status code
func (o *GetResourceGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource group unauthorized response has a 3xx status code
func (o *GetResourceGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource group unauthorized response has a 4xx status code
func (o *GetResourceGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resource group unauthorized response has a 5xx status code
func (o *GetResourceGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource group unauthorized response a status code equal to that given
func (o *GetResourceGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetResourceGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] getResourceGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetResourceGroupUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] getResourceGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetResourceGroupUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceGroupForbidden creates a GetResourceGroupForbidden with default headers values
func NewGetResourceGroupForbidden() *GetResourceGroupForbidden {
	return &GetResourceGroupForbidden{}
}

/*
GetResourceGroupForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetResourceGroupForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get resource group forbidden response has a 2xx status code
func (o *GetResourceGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource group forbidden response has a 3xx status code
func (o *GetResourceGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource group forbidden response has a 4xx status code
func (o *GetResourceGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resource group forbidden response has a 5xx status code
func (o *GetResourceGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource group forbidden response a status code equal to that given
func (o *GetResourceGroupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetResourceGroupForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] getResourceGroupForbidden  %+v", 403, o.Payload)
}

func (o *GetResourceGroupForbidden) String() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] getResourceGroupForbidden  %+v", 403, o.Payload)
}

func (o *GetResourceGroupForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceGroupNotFound creates a GetResourceGroupNotFound with default headers values
func NewGetResourceGroupNotFound() *GetResourceGroupNotFound {
	return &GetResourceGroupNotFound{}
}

/*
GetResourceGroupNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetResourceGroupNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get resource group not found response has a 2xx status code
func (o *GetResourceGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource group not found response has a 3xx status code
func (o *GetResourceGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource group not found response has a 4xx status code
func (o *GetResourceGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resource group not found response has a 5xx status code
func (o *GetResourceGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource group not found response a status code equal to that given
func (o *GetResourceGroupNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetResourceGroupNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] getResourceGroupNotFound  %+v", 404, o.Payload)
}

func (o *GetResourceGroupNotFound) String() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] getResourceGroupNotFound  %+v", 404, o.Payload)
}

func (o *GetResourceGroupNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceGroupInternalServerError creates a GetResourceGroupInternalServerError with default headers values
func NewGetResourceGroupInternalServerError() *GetResourceGroupInternalServerError {
	return &GetResourceGroupInternalServerError{}
}

/*
GetResourceGroupInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetResourceGroupInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get resource group internal server error response has a 2xx status code
func (o *GetResourceGroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource group internal server error response has a 3xx status code
func (o *GetResourceGroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource group internal server error response has a 4xx status code
func (o *GetResourceGroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resource group internal server error response has a 5xx status code
func (o *GetResourceGroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get resource group internal server error response a status code equal to that given
func (o *GetResourceGroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetResourceGroupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] getResourceGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetResourceGroupInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] getResourceGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetResourceGroupInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceGroupGatewayTimeout creates a GetResourceGroupGatewayTimeout with default headers values
func NewGetResourceGroupGatewayTimeout() *GetResourceGroupGatewayTimeout {
	return &GetResourceGroupGatewayTimeout{}
}

/*
GetResourceGroupGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetResourceGroupGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get resource group gateway timeout response has a 2xx status code
func (o *GetResourceGroupGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource group gateway timeout response has a 3xx status code
func (o *GetResourceGroupGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource group gateway timeout response has a 4xx status code
func (o *GetResourceGroupGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resource group gateway timeout response has a 5xx status code
func (o *GetResourceGroupGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get resource group gateway timeout response a status code equal to that given
func (o *GetResourceGroupGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetResourceGroupGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] getResourceGroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetResourceGroupGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}][%d] getResourceGroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetResourceGroupGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceGroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
