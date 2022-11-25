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

// QueryResourceGroupsReader is a Reader for the QueryResourceGroups structure.
type QueryResourceGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryResourceGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryResourceGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryResourceGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQueryResourceGroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryResourceGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryResourceGroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewQueryResourceGroupsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryResourceGroupsOK creates a QueryResourceGroupsOK with default headers values
func NewQueryResourceGroupsOK() *QueryResourceGroupsOK {
	return &QueryResourceGroupsOK{}
}

/*
QueryResourceGroupsOK describes a response with status code 200, with default header values.

A successful response.
*/
type QueryResourceGroupsOK struct {
	Payload *models.Tags
}

// IsSuccess returns true when this query resource groups o k response has a 2xx status code
func (o *QueryResourceGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query resource groups o k response has a 3xx status code
func (o *QueryResourceGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query resource groups o k response has a 4xx status code
func (o *QueryResourceGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query resource groups o k response has a 5xx status code
func (o *QueryResourceGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query resource groups o k response a status code equal to that given
func (o *QueryResourceGroupsOK) IsCode(code int) bool {
	return code == 200
}

func (o *QueryResourceGroupsOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects][%d] queryResourceGroupsOK  %+v", 200, o.Payload)
}

func (o *QueryResourceGroupsOK) String() string {
	return fmt.Sprintf("[GET /v1/projects][%d] queryResourceGroupsOK  %+v", 200, o.Payload)
}

func (o *QueryResourceGroupsOK) GetPayload() *models.Tags {
	return o.Payload
}

func (o *QueryResourceGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tags)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryResourceGroupsBadRequest creates a QueryResourceGroupsBadRequest with default headers values
func NewQueryResourceGroupsBadRequest() *QueryResourceGroupsBadRequest {
	return &QueryResourceGroupsBadRequest{}
}

/*
QueryResourceGroupsBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type QueryResourceGroupsBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this query resource groups bad request response has a 2xx status code
func (o *QueryResourceGroupsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query resource groups bad request response has a 3xx status code
func (o *QueryResourceGroupsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query resource groups bad request response has a 4xx status code
func (o *QueryResourceGroupsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query resource groups bad request response has a 5xx status code
func (o *QueryResourceGroupsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query resource groups bad request response a status code equal to that given
func (o *QueryResourceGroupsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *QueryResourceGroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/projects][%d] queryResourceGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *QueryResourceGroupsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/projects][%d] queryResourceGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *QueryResourceGroupsBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *QueryResourceGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryResourceGroupsUnauthorized creates a QueryResourceGroupsUnauthorized with default headers values
func NewQueryResourceGroupsUnauthorized() *QueryResourceGroupsUnauthorized {
	return &QueryResourceGroupsUnauthorized{}
}

/*
QueryResourceGroupsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type QueryResourceGroupsUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this query resource groups unauthorized response has a 2xx status code
func (o *QueryResourceGroupsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query resource groups unauthorized response has a 3xx status code
func (o *QueryResourceGroupsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query resource groups unauthorized response has a 4xx status code
func (o *QueryResourceGroupsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this query resource groups unauthorized response has a 5xx status code
func (o *QueryResourceGroupsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this query resource groups unauthorized response a status code equal to that given
func (o *QueryResourceGroupsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *QueryResourceGroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/projects][%d] queryResourceGroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *QueryResourceGroupsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/projects][%d] queryResourceGroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *QueryResourceGroupsUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *QueryResourceGroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryResourceGroupsForbidden creates a QueryResourceGroupsForbidden with default headers values
func NewQueryResourceGroupsForbidden() *QueryResourceGroupsForbidden {
	return &QueryResourceGroupsForbidden{}
}

/*
QueryResourceGroupsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type QueryResourceGroupsForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this query resource groups forbidden response has a 2xx status code
func (o *QueryResourceGroupsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query resource groups forbidden response has a 3xx status code
func (o *QueryResourceGroupsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query resource groups forbidden response has a 4xx status code
func (o *QueryResourceGroupsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query resource groups forbidden response has a 5xx status code
func (o *QueryResourceGroupsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query resource groups forbidden response a status code equal to that given
func (o *QueryResourceGroupsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *QueryResourceGroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/projects][%d] queryResourceGroupsForbidden  %+v", 403, o.Payload)
}

func (o *QueryResourceGroupsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/projects][%d] queryResourceGroupsForbidden  %+v", 403, o.Payload)
}

func (o *QueryResourceGroupsForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *QueryResourceGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryResourceGroupsInternalServerError creates a QueryResourceGroupsInternalServerError with default headers values
func NewQueryResourceGroupsInternalServerError() *QueryResourceGroupsInternalServerError {
	return &QueryResourceGroupsInternalServerError{}
}

/*
QueryResourceGroupsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type QueryResourceGroupsInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this query resource groups internal server error response has a 2xx status code
func (o *QueryResourceGroupsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query resource groups internal server error response has a 3xx status code
func (o *QueryResourceGroupsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query resource groups internal server error response has a 4xx status code
func (o *QueryResourceGroupsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query resource groups internal server error response has a 5xx status code
func (o *QueryResourceGroupsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query resource groups internal server error response a status code equal to that given
func (o *QueryResourceGroupsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *QueryResourceGroupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/projects][%d] queryResourceGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryResourceGroupsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/projects][%d] queryResourceGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryResourceGroupsInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *QueryResourceGroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryResourceGroupsGatewayTimeout creates a QueryResourceGroupsGatewayTimeout with default headers values
func NewQueryResourceGroupsGatewayTimeout() *QueryResourceGroupsGatewayTimeout {
	return &QueryResourceGroupsGatewayTimeout{}
}

/*
QueryResourceGroupsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type QueryResourceGroupsGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this query resource groups gateway timeout response has a 2xx status code
func (o *QueryResourceGroupsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query resource groups gateway timeout response has a 3xx status code
func (o *QueryResourceGroupsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query resource groups gateway timeout response has a 4xx status code
func (o *QueryResourceGroupsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this query resource groups gateway timeout response has a 5xx status code
func (o *QueryResourceGroupsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this query resource groups gateway timeout response a status code equal to that given
func (o *QueryResourceGroupsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *QueryResourceGroupsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/projects][%d] queryResourceGroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *QueryResourceGroupsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/projects][%d] queryResourceGroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *QueryResourceGroupsGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *QueryResourceGroupsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
