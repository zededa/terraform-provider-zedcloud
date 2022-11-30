// Code generated by go-swagger; DO NOT EDIT.

package device_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// GetDeviceStatusByNameReader is a Reader for the GetDeviceStatusByName structure.
type GetDeviceStatusByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceStatusByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceStatusByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDeviceStatusByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDeviceStatusByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeviceStatusByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDeviceStatusByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetDeviceStatusByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeviceStatusByNameOK creates a GetDeviceStatusByNameOK with default headers values
func NewGetDeviceStatusByNameOK() *GetDeviceStatusByNameOK {
	return &GetDeviceStatusByNameOK{}
}

/*
GetDeviceStatusByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetDeviceStatusByNameOK struct {
	Payload *models.DeviceStatus
}

// IsSuccess returns true when this get device status by name o k response has a 2xx status code
func (o *GetDeviceStatusByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device status by name o k response has a 3xx status code
func (o *GetDeviceStatusByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device status by name o k response has a 4xx status code
func (o *GetDeviceStatusByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device status by name o k response has a 5xx status code
func (o *GetDeviceStatusByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device status by name o k response a status code equal to that given
func (o *GetDeviceStatusByNameOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDeviceStatusByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status][%d] getDeviceStatusByNameOK  %+v", 200, o.Payload)
}

func (o *GetDeviceStatusByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status][%d] getDeviceStatusByNameOK  %+v", 200, o.Payload)
}

func (o *GetDeviceStatusByNameOK) GetPayload() *models.DeviceStatus {
	return o.Payload
}

func (o *GetDeviceStatusByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceStatusByNameUnauthorized creates a GetDeviceStatusByNameUnauthorized with default headers values
func NewGetDeviceStatusByNameUnauthorized() *GetDeviceStatusByNameUnauthorized {
	return &GetDeviceStatusByNameUnauthorized{}
}

/*
GetDeviceStatusByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetDeviceStatusByNameUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get device status by name unauthorized response has a 2xx status code
func (o *GetDeviceStatusByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get device status by name unauthorized response has a 3xx status code
func (o *GetDeviceStatusByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device status by name unauthorized response has a 4xx status code
func (o *GetDeviceStatusByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get device status by name unauthorized response has a 5xx status code
func (o *GetDeviceStatusByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get device status by name unauthorized response a status code equal to that given
func (o *GetDeviceStatusByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetDeviceStatusByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status][%d] getDeviceStatusByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDeviceStatusByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status][%d] getDeviceStatusByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDeviceStatusByNameUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetDeviceStatusByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceStatusByNameForbidden creates a GetDeviceStatusByNameForbidden with default headers values
func NewGetDeviceStatusByNameForbidden() *GetDeviceStatusByNameForbidden {
	return &GetDeviceStatusByNameForbidden{}
}

/*
GetDeviceStatusByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetDeviceStatusByNameForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get device status by name forbidden response has a 2xx status code
func (o *GetDeviceStatusByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get device status by name forbidden response has a 3xx status code
func (o *GetDeviceStatusByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device status by name forbidden response has a 4xx status code
func (o *GetDeviceStatusByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get device status by name forbidden response has a 5xx status code
func (o *GetDeviceStatusByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get device status by name forbidden response a status code equal to that given
func (o *GetDeviceStatusByNameForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetDeviceStatusByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status][%d] getDeviceStatusByNameForbidden  %+v", 403, o.Payload)
}

func (o *GetDeviceStatusByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status][%d] getDeviceStatusByNameForbidden  %+v", 403, o.Payload)
}

func (o *GetDeviceStatusByNameForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetDeviceStatusByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceStatusByNameNotFound creates a GetDeviceStatusByNameNotFound with default headers values
func NewGetDeviceStatusByNameNotFound() *GetDeviceStatusByNameNotFound {
	return &GetDeviceStatusByNameNotFound{}
}

/*
GetDeviceStatusByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetDeviceStatusByNameNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get device status by name not found response has a 2xx status code
func (o *GetDeviceStatusByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get device status by name not found response has a 3xx status code
func (o *GetDeviceStatusByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device status by name not found response has a 4xx status code
func (o *GetDeviceStatusByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get device status by name not found response has a 5xx status code
func (o *GetDeviceStatusByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get device status by name not found response a status code equal to that given
func (o *GetDeviceStatusByNameNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetDeviceStatusByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status][%d] getDeviceStatusByNameNotFound  %+v", 404, o.Payload)
}

func (o *GetDeviceStatusByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status][%d] getDeviceStatusByNameNotFound  %+v", 404, o.Payload)
}

func (o *GetDeviceStatusByNameNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetDeviceStatusByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceStatusByNameInternalServerError creates a GetDeviceStatusByNameInternalServerError with default headers values
func NewGetDeviceStatusByNameInternalServerError() *GetDeviceStatusByNameInternalServerError {
	return &GetDeviceStatusByNameInternalServerError{}
}

/*
GetDeviceStatusByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetDeviceStatusByNameInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get device status by name internal server error response has a 2xx status code
func (o *GetDeviceStatusByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get device status by name internal server error response has a 3xx status code
func (o *GetDeviceStatusByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device status by name internal server error response has a 4xx status code
func (o *GetDeviceStatusByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device status by name internal server error response has a 5xx status code
func (o *GetDeviceStatusByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get device status by name internal server error response a status code equal to that given
func (o *GetDeviceStatusByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetDeviceStatusByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status][%d] getDeviceStatusByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDeviceStatusByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status][%d] getDeviceStatusByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDeviceStatusByNameInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetDeviceStatusByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceStatusByNameGatewayTimeout creates a GetDeviceStatusByNameGatewayTimeout with default headers values
func NewGetDeviceStatusByNameGatewayTimeout() *GetDeviceStatusByNameGatewayTimeout {
	return &GetDeviceStatusByNameGatewayTimeout{}
}

/*
GetDeviceStatusByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetDeviceStatusByNameGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get device status by name gateway timeout response has a 2xx status code
func (o *GetDeviceStatusByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get device status by name gateway timeout response has a 3xx status code
func (o *GetDeviceStatusByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device status by name gateway timeout response has a 4xx status code
func (o *GetDeviceStatusByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device status by name gateway timeout response has a 5xx status code
func (o *GetDeviceStatusByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get device status by name gateway timeout response a status code equal to that given
func (o *GetDeviceStatusByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetDeviceStatusByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status][%d] getDeviceStatusByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetDeviceStatusByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{name}/status][%d] getDeviceStatusByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetDeviceStatusByNameGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetDeviceStatusByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
