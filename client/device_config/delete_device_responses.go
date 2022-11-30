// Code generated by go-swagger; DO NOT EDIT.

package device_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// DeleteDeviceReader is a Reader for the DeleteDevice structure.
type DeleteDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteDeviceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDeviceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDeviceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDeviceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteDeviceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDeviceOK creates a DeleteDeviceOK with default headers values
func NewDeleteDeviceOK() *DeleteDeviceOK {
	return &DeleteDeviceOK{}
}

/*
DeleteDeviceOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteDeviceOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this delete device o k response has a 2xx status code
func (o *DeleteDeviceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete device o k response has a 3xx status code
func (o *DeleteDeviceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete device o k response has a 4xx status code
func (o *DeleteDeviceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete device o k response has a 5xx status code
func (o *DeleteDeviceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete device o k response a status code equal to that given
func (o *DeleteDeviceOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteDeviceOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] deleteDeviceOK  %+v", 200, o.Payload)
}

func (o *DeleteDeviceOK) String() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] deleteDeviceOK  %+v", 200, o.Payload)
}

func (o *DeleteDeviceOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceUnauthorized creates a DeleteDeviceUnauthorized with default headers values
func NewDeleteDeviceUnauthorized() *DeleteDeviceUnauthorized {
	return &DeleteDeviceUnauthorized{}
}

/*
DeleteDeviceUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type DeleteDeviceUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this delete device unauthorized response has a 2xx status code
func (o *DeleteDeviceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete device unauthorized response has a 3xx status code
func (o *DeleteDeviceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete device unauthorized response has a 4xx status code
func (o *DeleteDeviceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete device unauthorized response has a 5xx status code
func (o *DeleteDeviceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete device unauthorized response a status code equal to that given
func (o *DeleteDeviceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteDeviceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] deleteDeviceUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteDeviceUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] deleteDeviceUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteDeviceUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteDeviceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceForbidden creates a DeleteDeviceForbidden with default headers values
func NewDeleteDeviceForbidden() *DeleteDeviceForbidden {
	return &DeleteDeviceForbidden{}
}

/*
DeleteDeviceForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type DeleteDeviceForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this delete device forbidden response has a 2xx status code
func (o *DeleteDeviceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete device forbidden response has a 3xx status code
func (o *DeleteDeviceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete device forbidden response has a 4xx status code
func (o *DeleteDeviceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete device forbidden response has a 5xx status code
func (o *DeleteDeviceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete device forbidden response a status code equal to that given
func (o *DeleteDeviceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteDeviceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] deleteDeviceForbidden  %+v", 403, o.Payload)
}

func (o *DeleteDeviceForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] deleteDeviceForbidden  %+v", 403, o.Payload)
}

func (o *DeleteDeviceForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteDeviceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceNotFound creates a DeleteDeviceNotFound with default headers values
func NewDeleteDeviceNotFound() *DeleteDeviceNotFound {
	return &DeleteDeviceNotFound{}
}

/*
DeleteDeviceNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type DeleteDeviceNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this delete device not found response has a 2xx status code
func (o *DeleteDeviceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete device not found response has a 3xx status code
func (o *DeleteDeviceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete device not found response has a 4xx status code
func (o *DeleteDeviceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete device not found response has a 5xx status code
func (o *DeleteDeviceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete device not found response a status code equal to that given
func (o *DeleteDeviceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteDeviceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] deleteDeviceNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDeviceNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] deleteDeviceNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDeviceNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteDeviceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceInternalServerError creates a DeleteDeviceInternalServerError with default headers values
func NewDeleteDeviceInternalServerError() *DeleteDeviceInternalServerError {
	return &DeleteDeviceInternalServerError{}
}

/*
DeleteDeviceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type DeleteDeviceInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this delete device internal server error response has a 2xx status code
func (o *DeleteDeviceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete device internal server error response has a 3xx status code
func (o *DeleteDeviceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete device internal server error response has a 4xx status code
func (o *DeleteDeviceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete device internal server error response has a 5xx status code
func (o *DeleteDeviceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete device internal server error response a status code equal to that given
func (o *DeleteDeviceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteDeviceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] deleteDeviceInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteDeviceInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] deleteDeviceInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteDeviceInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteDeviceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceGatewayTimeout creates a DeleteDeviceGatewayTimeout with default headers values
func NewDeleteDeviceGatewayTimeout() *DeleteDeviceGatewayTimeout {
	return &DeleteDeviceGatewayTimeout{}
}

/*
DeleteDeviceGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type DeleteDeviceGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this delete device gateway timeout response has a 2xx status code
func (o *DeleteDeviceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete device gateway timeout response has a 3xx status code
func (o *DeleteDeviceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete device gateway timeout response has a 4xx status code
func (o *DeleteDeviceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete device gateway timeout response has a 5xx status code
func (o *DeleteDeviceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete device gateway timeout response a status code equal to that given
func (o *DeleteDeviceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteDeviceGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] deleteDeviceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteDeviceGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /v1/devices/id/{id}][%d] deleteDeviceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteDeviceGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteDeviceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
