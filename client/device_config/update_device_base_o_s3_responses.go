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

// UpdateDeviceBaseOS3Reader is a Reader for the UpdateDeviceBaseOS3 structure.
type UpdateDeviceBaseOS3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceBaseOS3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceBaseOS3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateDeviceBaseOS3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateDeviceBaseOS3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateDeviceBaseOS3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateDeviceBaseOS3Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateDeviceBaseOS3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewUpdateDeviceBaseOS3GatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateDeviceBaseOS3OK creates a UpdateDeviceBaseOS3OK with default headers values
func NewUpdateDeviceBaseOS3OK() *UpdateDeviceBaseOS3OK {
	return &UpdateDeviceBaseOS3OK{}
}

/*
UpdateDeviceBaseOS3OK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateDeviceBaseOS3OK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this update device base o s3 o k response has a 2xx status code
func (o *UpdateDeviceBaseOS3OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update device base o s3 o k response has a 3xx status code
func (o *UpdateDeviceBaseOS3OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update device base o s3 o k response has a 4xx status code
func (o *UpdateDeviceBaseOS3OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update device base o s3 o k response has a 5xx status code
func (o *UpdateDeviceBaseOS3OK) IsServerError() bool {
	return false
}

// IsCode returns true when this update device base o s3 o k response a status code equal to that given
func (o *UpdateDeviceBaseOS3OK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateDeviceBaseOS3OK) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] updateDeviceBaseOS3OK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceBaseOS3OK) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] updateDeviceBaseOS3OK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceBaseOS3OK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateDeviceBaseOS3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDeviceBaseOS3Unauthorized creates a UpdateDeviceBaseOS3Unauthorized with default headers values
func NewUpdateDeviceBaseOS3Unauthorized() *UpdateDeviceBaseOS3Unauthorized {
	return &UpdateDeviceBaseOS3Unauthorized{}
}

/*
UpdateDeviceBaseOS3Unauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type UpdateDeviceBaseOS3Unauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this update device base o s3 unauthorized response has a 2xx status code
func (o *UpdateDeviceBaseOS3Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update device base o s3 unauthorized response has a 3xx status code
func (o *UpdateDeviceBaseOS3Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update device base o s3 unauthorized response has a 4xx status code
func (o *UpdateDeviceBaseOS3Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update device base o s3 unauthorized response has a 5xx status code
func (o *UpdateDeviceBaseOS3Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update device base o s3 unauthorized response a status code equal to that given
func (o *UpdateDeviceBaseOS3Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateDeviceBaseOS3Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] updateDeviceBaseOS3Unauthorized  %+v", 401, o.Payload)
}

func (o *UpdateDeviceBaseOS3Unauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] updateDeviceBaseOS3Unauthorized  %+v", 401, o.Payload)
}

func (o *UpdateDeviceBaseOS3Unauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateDeviceBaseOS3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDeviceBaseOS3Forbidden creates a UpdateDeviceBaseOS3Forbidden with default headers values
func NewUpdateDeviceBaseOS3Forbidden() *UpdateDeviceBaseOS3Forbidden {
	return &UpdateDeviceBaseOS3Forbidden{}
}

/*
UpdateDeviceBaseOS3Forbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type UpdateDeviceBaseOS3Forbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this update device base o s3 forbidden response has a 2xx status code
func (o *UpdateDeviceBaseOS3Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update device base o s3 forbidden response has a 3xx status code
func (o *UpdateDeviceBaseOS3Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update device base o s3 forbidden response has a 4xx status code
func (o *UpdateDeviceBaseOS3Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update device base o s3 forbidden response has a 5xx status code
func (o *UpdateDeviceBaseOS3Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update device base o s3 forbidden response a status code equal to that given
func (o *UpdateDeviceBaseOS3Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateDeviceBaseOS3Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] updateDeviceBaseOS3Forbidden  %+v", 403, o.Payload)
}

func (o *UpdateDeviceBaseOS3Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] updateDeviceBaseOS3Forbidden  %+v", 403, o.Payload)
}

func (o *UpdateDeviceBaseOS3Forbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateDeviceBaseOS3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDeviceBaseOS3NotFound creates a UpdateDeviceBaseOS3NotFound with default headers values
func NewUpdateDeviceBaseOS3NotFound() *UpdateDeviceBaseOS3NotFound {
	return &UpdateDeviceBaseOS3NotFound{}
}

/*
UpdateDeviceBaseOS3NotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type UpdateDeviceBaseOS3NotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this update device base o s3 not found response has a 2xx status code
func (o *UpdateDeviceBaseOS3NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update device base o s3 not found response has a 3xx status code
func (o *UpdateDeviceBaseOS3NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update device base o s3 not found response has a 4xx status code
func (o *UpdateDeviceBaseOS3NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update device base o s3 not found response has a 5xx status code
func (o *UpdateDeviceBaseOS3NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update device base o s3 not found response a status code equal to that given
func (o *UpdateDeviceBaseOS3NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateDeviceBaseOS3NotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] updateDeviceBaseOS3NotFound  %+v", 404, o.Payload)
}

func (o *UpdateDeviceBaseOS3NotFound) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] updateDeviceBaseOS3NotFound  %+v", 404, o.Payload)
}

func (o *UpdateDeviceBaseOS3NotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateDeviceBaseOS3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDeviceBaseOS3Conflict creates a UpdateDeviceBaseOS3Conflict with default headers values
func NewUpdateDeviceBaseOS3Conflict() *UpdateDeviceBaseOS3Conflict {
	return &UpdateDeviceBaseOS3Conflict{}
}

/*
UpdateDeviceBaseOS3Conflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing device record.
*/
type UpdateDeviceBaseOS3Conflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this update device base o s3 conflict response has a 2xx status code
func (o *UpdateDeviceBaseOS3Conflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update device base o s3 conflict response has a 3xx status code
func (o *UpdateDeviceBaseOS3Conflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update device base o s3 conflict response has a 4xx status code
func (o *UpdateDeviceBaseOS3Conflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update device base o s3 conflict response has a 5xx status code
func (o *UpdateDeviceBaseOS3Conflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update device base o s3 conflict response a status code equal to that given
func (o *UpdateDeviceBaseOS3Conflict) IsCode(code int) bool {
	return code == 409
}

func (o *UpdateDeviceBaseOS3Conflict) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] updateDeviceBaseOS3Conflict  %+v", 409, o.Payload)
}

func (o *UpdateDeviceBaseOS3Conflict) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] updateDeviceBaseOS3Conflict  %+v", 409, o.Payload)
}

func (o *UpdateDeviceBaseOS3Conflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateDeviceBaseOS3Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDeviceBaseOS3InternalServerError creates a UpdateDeviceBaseOS3InternalServerError with default headers values
func NewUpdateDeviceBaseOS3InternalServerError() *UpdateDeviceBaseOS3InternalServerError {
	return &UpdateDeviceBaseOS3InternalServerError{}
}

/*
UpdateDeviceBaseOS3InternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type UpdateDeviceBaseOS3InternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this update device base o s3 internal server error response has a 2xx status code
func (o *UpdateDeviceBaseOS3InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update device base o s3 internal server error response has a 3xx status code
func (o *UpdateDeviceBaseOS3InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update device base o s3 internal server error response has a 4xx status code
func (o *UpdateDeviceBaseOS3InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update device base o s3 internal server error response has a 5xx status code
func (o *UpdateDeviceBaseOS3InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update device base o s3 internal server error response a status code equal to that given
func (o *UpdateDeviceBaseOS3InternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateDeviceBaseOS3InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] updateDeviceBaseOS3InternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateDeviceBaseOS3InternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] updateDeviceBaseOS3InternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateDeviceBaseOS3InternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateDeviceBaseOS3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDeviceBaseOS3GatewayTimeout creates a UpdateDeviceBaseOS3GatewayTimeout with default headers values
func NewUpdateDeviceBaseOS3GatewayTimeout() *UpdateDeviceBaseOS3GatewayTimeout {
	return &UpdateDeviceBaseOS3GatewayTimeout{}
}

/*
UpdateDeviceBaseOS3GatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type UpdateDeviceBaseOS3GatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this update device base o s3 gateway timeout response has a 2xx status code
func (o *UpdateDeviceBaseOS3GatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update device base o s3 gateway timeout response has a 3xx status code
func (o *UpdateDeviceBaseOS3GatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update device base o s3 gateway timeout response has a 4xx status code
func (o *UpdateDeviceBaseOS3GatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this update device base o s3 gateway timeout response has a 5xx status code
func (o *UpdateDeviceBaseOS3GatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this update device base o s3 gateway timeout response a status code equal to that given
func (o *UpdateDeviceBaseOS3GatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *UpdateDeviceBaseOS3GatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] updateDeviceBaseOS3GatewayTimeout  %+v", 504, o.Payload)
}

func (o *UpdateDeviceBaseOS3GatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/unpublish][%d] updateDeviceBaseOS3GatewayTimeout  %+v", 504, o.Payload)
}

func (o *UpdateDeviceBaseOS3GatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateDeviceBaseOS3GatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
