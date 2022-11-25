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

// GetHardwareModelByNameReader is a Reader for the GetHardwareModelByName structure.
type GetHardwareModelByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHardwareModelByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHardwareModelByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetHardwareModelByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetHardwareModelByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetHardwareModelByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetHardwareModelByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetHardwareModelByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHardwareModelByNameOK creates a GetHardwareModelByNameOK with default headers values
func NewGetHardwareModelByNameOK() *GetHardwareModelByNameOK {
	return &GetHardwareModelByNameOK{}
}

/*
GetHardwareModelByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetHardwareModelByNameOK struct {
	Payload *models.SysModel
}

// IsSuccess returns true when this get hardware model by name o k response has a 2xx status code
func (o *GetHardwareModelByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get hardware model by name o k response has a 3xx status code
func (o *GetHardwareModelByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get hardware model by name o k response has a 4xx status code
func (o *GetHardwareModelByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get hardware model by name o k response has a 5xx status code
func (o *GetHardwareModelByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get hardware model by name o k response a status code equal to that given
func (o *GetHardwareModelByNameOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetHardwareModelByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] getHardwareModelByNameOK  %+v", 200, o.Payload)
}

func (o *GetHardwareModelByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] getHardwareModelByNameOK  %+v", 200, o.Payload)
}

func (o *GetHardwareModelByNameOK) GetPayload() *models.SysModel {
	return o.Payload
}

func (o *GetHardwareModelByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SysModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHardwareModelByNameUnauthorized creates a GetHardwareModelByNameUnauthorized with default headers values
func NewGetHardwareModelByNameUnauthorized() *GetHardwareModelByNameUnauthorized {
	return &GetHardwareModelByNameUnauthorized{}
}

/*
GetHardwareModelByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetHardwareModelByNameUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get hardware model by name unauthorized response has a 2xx status code
func (o *GetHardwareModelByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get hardware model by name unauthorized response has a 3xx status code
func (o *GetHardwareModelByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get hardware model by name unauthorized response has a 4xx status code
func (o *GetHardwareModelByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get hardware model by name unauthorized response has a 5xx status code
func (o *GetHardwareModelByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get hardware model by name unauthorized response a status code equal to that given
func (o *GetHardwareModelByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetHardwareModelByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] getHardwareModelByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *GetHardwareModelByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] getHardwareModelByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *GetHardwareModelByNameUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetHardwareModelByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHardwareModelByNameForbidden creates a GetHardwareModelByNameForbidden with default headers values
func NewGetHardwareModelByNameForbidden() *GetHardwareModelByNameForbidden {
	return &GetHardwareModelByNameForbidden{}
}

/*
GetHardwareModelByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetHardwareModelByNameForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get hardware model by name forbidden response has a 2xx status code
func (o *GetHardwareModelByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get hardware model by name forbidden response has a 3xx status code
func (o *GetHardwareModelByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get hardware model by name forbidden response has a 4xx status code
func (o *GetHardwareModelByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get hardware model by name forbidden response has a 5xx status code
func (o *GetHardwareModelByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get hardware model by name forbidden response a status code equal to that given
func (o *GetHardwareModelByNameForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetHardwareModelByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] getHardwareModelByNameForbidden  %+v", 403, o.Payload)
}

func (o *GetHardwareModelByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] getHardwareModelByNameForbidden  %+v", 403, o.Payload)
}

func (o *GetHardwareModelByNameForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetHardwareModelByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHardwareModelByNameNotFound creates a GetHardwareModelByNameNotFound with default headers values
func NewGetHardwareModelByNameNotFound() *GetHardwareModelByNameNotFound {
	return &GetHardwareModelByNameNotFound{}
}

/*
GetHardwareModelByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetHardwareModelByNameNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get hardware model by name not found response has a 2xx status code
func (o *GetHardwareModelByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get hardware model by name not found response has a 3xx status code
func (o *GetHardwareModelByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get hardware model by name not found response has a 4xx status code
func (o *GetHardwareModelByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get hardware model by name not found response has a 5xx status code
func (o *GetHardwareModelByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get hardware model by name not found response a status code equal to that given
func (o *GetHardwareModelByNameNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetHardwareModelByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] getHardwareModelByNameNotFound  %+v", 404, o.Payload)
}

func (o *GetHardwareModelByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] getHardwareModelByNameNotFound  %+v", 404, o.Payload)
}

func (o *GetHardwareModelByNameNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetHardwareModelByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHardwareModelByNameInternalServerError creates a GetHardwareModelByNameInternalServerError with default headers values
func NewGetHardwareModelByNameInternalServerError() *GetHardwareModelByNameInternalServerError {
	return &GetHardwareModelByNameInternalServerError{}
}

/*
GetHardwareModelByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetHardwareModelByNameInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get hardware model by name internal server error response has a 2xx status code
func (o *GetHardwareModelByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get hardware model by name internal server error response has a 3xx status code
func (o *GetHardwareModelByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get hardware model by name internal server error response has a 4xx status code
func (o *GetHardwareModelByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get hardware model by name internal server error response has a 5xx status code
func (o *GetHardwareModelByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get hardware model by name internal server error response a status code equal to that given
func (o *GetHardwareModelByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetHardwareModelByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] getHardwareModelByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHardwareModelByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] getHardwareModelByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHardwareModelByNameInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetHardwareModelByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHardwareModelByNameGatewayTimeout creates a GetHardwareModelByNameGatewayTimeout with default headers values
func NewGetHardwareModelByNameGatewayTimeout() *GetHardwareModelByNameGatewayTimeout {
	return &GetHardwareModelByNameGatewayTimeout{}
}

/*
GetHardwareModelByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetHardwareModelByNameGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get hardware model by name gateway timeout response has a 2xx status code
func (o *GetHardwareModelByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get hardware model by name gateway timeout response has a 3xx status code
func (o *GetHardwareModelByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get hardware model by name gateway timeout response has a 4xx status code
func (o *GetHardwareModelByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get hardware model by name gateway timeout response has a 5xx status code
func (o *GetHardwareModelByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get hardware model by name gateway timeout response a status code equal to that given
func (o *GetHardwareModelByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetHardwareModelByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] getHardwareModelByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetHardwareModelByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] getHardwareModelByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetHardwareModelByNameGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetHardwareModelByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
