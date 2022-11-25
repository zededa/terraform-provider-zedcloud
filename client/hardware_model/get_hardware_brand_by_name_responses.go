// Code generated by go-swagger; DO NOT EDIT.

package hardware_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// GetHardwareBrandByNameReader is a Reader for the GetHardwareBrandByName structure.
type GetHardwareBrandByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHardwareBrandByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHardwareBrandByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetHardwareBrandByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetHardwareBrandByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetHardwareBrandByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetHardwareBrandByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetHardwareBrandByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHardwareBrandByNameOK creates a GetHardwareBrandByNameOK with default headers values
func NewGetHardwareBrandByNameOK() *GetHardwareBrandByNameOK {
	return &GetHardwareBrandByNameOK{}
}

/*
GetHardwareBrandByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetHardwareBrandByNameOK struct {
	Payload *models.SysBrand
}

// IsSuccess returns true when this get hardware brand by name o k response has a 2xx status code
func (o *GetHardwareBrandByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get hardware brand by name o k response has a 3xx status code
func (o *GetHardwareBrandByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get hardware brand by name o k response has a 4xx status code
func (o *GetHardwareBrandByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get hardware brand by name o k response has a 5xx status code
func (o *GetHardwareBrandByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get hardware brand by name o k response a status code equal to that given
func (o *GetHardwareBrandByNameOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetHardwareBrandByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] getHardwareBrandByNameOK  %+v", 200, o.Payload)
}

func (o *GetHardwareBrandByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] getHardwareBrandByNameOK  %+v", 200, o.Payload)
}

func (o *GetHardwareBrandByNameOK) GetPayload() *models.SysBrand {
	return o.Payload
}

func (o *GetHardwareBrandByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SysBrand)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHardwareBrandByNameUnauthorized creates a GetHardwareBrandByNameUnauthorized with default headers values
func NewGetHardwareBrandByNameUnauthorized() *GetHardwareBrandByNameUnauthorized {
	return &GetHardwareBrandByNameUnauthorized{}
}

/*
GetHardwareBrandByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetHardwareBrandByNameUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get hardware brand by name unauthorized response has a 2xx status code
func (o *GetHardwareBrandByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get hardware brand by name unauthorized response has a 3xx status code
func (o *GetHardwareBrandByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get hardware brand by name unauthorized response has a 4xx status code
func (o *GetHardwareBrandByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get hardware brand by name unauthorized response has a 5xx status code
func (o *GetHardwareBrandByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get hardware brand by name unauthorized response a status code equal to that given
func (o *GetHardwareBrandByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetHardwareBrandByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] getHardwareBrandByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *GetHardwareBrandByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] getHardwareBrandByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *GetHardwareBrandByNameUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetHardwareBrandByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHardwareBrandByNameForbidden creates a GetHardwareBrandByNameForbidden with default headers values
func NewGetHardwareBrandByNameForbidden() *GetHardwareBrandByNameForbidden {
	return &GetHardwareBrandByNameForbidden{}
}

/*
GetHardwareBrandByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetHardwareBrandByNameForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get hardware brand by name forbidden response has a 2xx status code
func (o *GetHardwareBrandByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get hardware brand by name forbidden response has a 3xx status code
func (o *GetHardwareBrandByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get hardware brand by name forbidden response has a 4xx status code
func (o *GetHardwareBrandByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get hardware brand by name forbidden response has a 5xx status code
func (o *GetHardwareBrandByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get hardware brand by name forbidden response a status code equal to that given
func (o *GetHardwareBrandByNameForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetHardwareBrandByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] getHardwareBrandByNameForbidden  %+v", 403, o.Payload)
}

func (o *GetHardwareBrandByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] getHardwareBrandByNameForbidden  %+v", 403, o.Payload)
}

func (o *GetHardwareBrandByNameForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetHardwareBrandByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHardwareBrandByNameNotFound creates a GetHardwareBrandByNameNotFound with default headers values
func NewGetHardwareBrandByNameNotFound() *GetHardwareBrandByNameNotFound {
	return &GetHardwareBrandByNameNotFound{}
}

/*
GetHardwareBrandByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetHardwareBrandByNameNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get hardware brand by name not found response has a 2xx status code
func (o *GetHardwareBrandByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get hardware brand by name not found response has a 3xx status code
func (o *GetHardwareBrandByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get hardware brand by name not found response has a 4xx status code
func (o *GetHardwareBrandByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get hardware brand by name not found response has a 5xx status code
func (o *GetHardwareBrandByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get hardware brand by name not found response a status code equal to that given
func (o *GetHardwareBrandByNameNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetHardwareBrandByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] getHardwareBrandByNameNotFound  %+v", 404, o.Payload)
}

func (o *GetHardwareBrandByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] getHardwareBrandByNameNotFound  %+v", 404, o.Payload)
}

func (o *GetHardwareBrandByNameNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetHardwareBrandByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHardwareBrandByNameInternalServerError creates a GetHardwareBrandByNameInternalServerError with default headers values
func NewGetHardwareBrandByNameInternalServerError() *GetHardwareBrandByNameInternalServerError {
	return &GetHardwareBrandByNameInternalServerError{}
}

/*
GetHardwareBrandByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetHardwareBrandByNameInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get hardware brand by name internal server error response has a 2xx status code
func (o *GetHardwareBrandByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get hardware brand by name internal server error response has a 3xx status code
func (o *GetHardwareBrandByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get hardware brand by name internal server error response has a 4xx status code
func (o *GetHardwareBrandByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get hardware brand by name internal server error response has a 5xx status code
func (o *GetHardwareBrandByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get hardware brand by name internal server error response a status code equal to that given
func (o *GetHardwareBrandByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetHardwareBrandByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] getHardwareBrandByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHardwareBrandByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] getHardwareBrandByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHardwareBrandByNameInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetHardwareBrandByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHardwareBrandByNameGatewayTimeout creates a GetHardwareBrandByNameGatewayTimeout with default headers values
func NewGetHardwareBrandByNameGatewayTimeout() *GetHardwareBrandByNameGatewayTimeout {
	return &GetHardwareBrandByNameGatewayTimeout{}
}

/*
GetHardwareBrandByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetHardwareBrandByNameGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get hardware brand by name gateway timeout response has a 2xx status code
func (o *GetHardwareBrandByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get hardware brand by name gateway timeout response has a 3xx status code
func (o *GetHardwareBrandByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get hardware brand by name gateway timeout response has a 4xx status code
func (o *GetHardwareBrandByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get hardware brand by name gateway timeout response has a 5xx status code
func (o *GetHardwareBrandByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get hardware brand by name gateway timeout response a status code equal to that given
func (o *GetHardwareBrandByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetHardwareBrandByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] getHardwareBrandByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetHardwareBrandByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] getHardwareBrandByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetHardwareBrandByNameGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetHardwareBrandByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
