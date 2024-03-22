package volume_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/models"
)

// VolumeInstanceConfigurationGetVolumeInstanceReader is a Reader for the VolumeInstanceConfigurationGetVolumeInstance structure.
type VolumeInstanceConfigurationGetVolumeInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeInstanceConfigurationGetVolumeInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeInstanceConfigurationGetVolumeInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewVolumeInstanceConfigurationGetVolumeInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewVolumeInstanceConfigurationGetVolumeInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewVolumeInstanceConfigurationGetVolumeInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewVolumeInstanceConfigurationGetVolumeInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewVolumeInstanceConfigurationGetVolumeInstanceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewVolumeInstanceConfigurationGetVolumeInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeInstanceConfigurationGetVolumeInstanceOK creates a VolumeInstanceConfigurationGetVolumeInstanceOK with default headers values
func NewVolumeInstanceConfigurationGetVolumeInstanceOK() *VolumeInstanceConfigurationGetVolumeInstanceOK {
	return &VolumeInstanceConfigurationGetVolumeInstanceOK{}
}

/*
VolumeInstanceConfigurationGetVolumeInstanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type VolumeInstanceConfigurationGetVolumeInstanceOK struct {
	Payload *models.VolumeInstance
}

// IsSuccess returns true when this volume instance configuration get volume instance o k response has a 2xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volume instance configuration get volume instance o k response has a 3xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration get volume instance o k response has a 4xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance configuration get volume instance o k response has a 5xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration get volume instance o k response a status code equal to that given
func (o *VolumeInstanceConfigurationGetVolumeInstanceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the volume instance configuration get volume instance o k response
func (o *VolumeInstanceConfigurationGetVolumeInstanceOK) Code() int {
	return 200
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceOK) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationGetVolumeInstanceOK  %+v", 200, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceOK) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationGetVolumeInstanceOK  %+v", 200, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceOK) GetPayload() *models.VolumeInstance {
	return o.Payload
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationGetVolumeInstanceUnauthorized creates a VolumeInstanceConfigurationGetVolumeInstanceUnauthorized with default headers values
func NewVolumeInstanceConfigurationGetVolumeInstanceUnauthorized() *VolumeInstanceConfigurationGetVolumeInstanceUnauthorized {
	return &VolumeInstanceConfigurationGetVolumeInstanceUnauthorized{}
}

/*
VolumeInstanceConfigurationGetVolumeInstanceUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type VolumeInstanceConfigurationGetVolumeInstanceUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration get volume instance unauthorized response has a 2xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration get volume instance unauthorized response has a 3xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration get volume instance unauthorized response has a 4xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance configuration get volume instance unauthorized response has a 5xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration get volume instance unauthorized response a status code equal to that given
func (o *VolumeInstanceConfigurationGetVolumeInstanceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the volume instance configuration get volume instance unauthorized response
func (o *VolumeInstanceConfigurationGetVolumeInstanceUnauthorized) Code() int {
	return 401
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationGetVolumeInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationGetVolumeInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationGetVolumeInstanceForbidden creates a VolumeInstanceConfigurationGetVolumeInstanceForbidden with default headers values
func NewVolumeInstanceConfigurationGetVolumeInstanceForbidden() *VolumeInstanceConfigurationGetVolumeInstanceForbidden {
	return &VolumeInstanceConfigurationGetVolumeInstanceForbidden{}
}

/*
VolumeInstanceConfigurationGetVolumeInstanceForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type VolumeInstanceConfigurationGetVolumeInstanceForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration get volume instance forbidden response has a 2xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration get volume instance forbidden response has a 3xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration get volume instance forbidden response has a 4xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance configuration get volume instance forbidden response has a 5xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration get volume instance forbidden response a status code equal to that given
func (o *VolumeInstanceConfigurationGetVolumeInstanceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the volume instance configuration get volume instance forbidden response
func (o *VolumeInstanceConfigurationGetVolumeInstanceForbidden) Code() int {
	return 403
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationGetVolumeInstanceForbidden  %+v", 403, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceForbidden) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationGetVolumeInstanceForbidden  %+v", 403, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationGetVolumeInstanceNotFound creates a VolumeInstanceConfigurationGetVolumeInstanceNotFound with default headers values
func NewVolumeInstanceConfigurationGetVolumeInstanceNotFound() *GetByIDNotFound {
	return &GetByIDNotFound{}
}

/*
GetByIDNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetByIDNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration get volume instance not found response has a 2xx status code
func (o *GetByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration get volume instance not found response has a 3xx status code
func (o *GetByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration get volume instance not found response has a 4xx status code
func (o *GetByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance configuration get volume instance not found response has a 5xx status code
func (o *GetByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration get volume instance not found response a status code equal to that given
func (o *GetByIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the volume instance configuration get volume instance not found response
func (o *GetByIDNotFound) Code() int {
	return 404
}

func (o *GetByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationGetVolumeInstanceNotFound  %+v", 404, o.Payload)
}

func (o *GetByIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationGetVolumeInstanceNotFound  %+v", 404, o.Payload)
}

func (o *GetByIDNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationGetVolumeInstanceInternalServerError creates a VolumeInstanceConfigurationGetVolumeInstanceInternalServerError with default headers values
func NewVolumeInstanceConfigurationGetVolumeInstanceInternalServerError() *VolumeInstanceConfigurationGetVolumeInstanceInternalServerError {
	return &VolumeInstanceConfigurationGetVolumeInstanceInternalServerError{}
}

/*
VolumeInstanceConfigurationGetVolumeInstanceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type VolumeInstanceConfigurationGetVolumeInstanceInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration get volume instance internal server error response has a 2xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration get volume instance internal server error response has a 3xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration get volume instance internal server error response has a 4xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance configuration get volume instance internal server error response has a 5xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this volume instance configuration get volume instance internal server error response a status code equal to that given
func (o *VolumeInstanceConfigurationGetVolumeInstanceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the volume instance configuration get volume instance internal server error response
func (o *VolumeInstanceConfigurationGetVolumeInstanceInternalServerError) Code() int {
	return 500
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationGetVolumeInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationGetVolumeInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationGetVolumeInstanceGatewayTimeout creates a VolumeInstanceConfigurationGetVolumeInstanceGatewayTimeout with default headers values
func NewVolumeInstanceConfigurationGetVolumeInstanceGatewayTimeout() *VolumeInstanceConfigurationGetVolumeInstanceGatewayTimeout {
	return &VolumeInstanceConfigurationGetVolumeInstanceGatewayTimeout{}
}

/*
VolumeInstanceConfigurationGetVolumeInstanceGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type VolumeInstanceConfigurationGetVolumeInstanceGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration get volume instance gateway timeout response has a 2xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration get volume instance gateway timeout response has a 3xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration get volume instance gateway timeout response has a 4xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance configuration get volume instance gateway timeout response has a 5xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this volume instance configuration get volume instance gateway timeout response a status code equal to that given
func (o *VolumeInstanceConfigurationGetVolumeInstanceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the volume instance configuration get volume instance gateway timeout response
func (o *VolumeInstanceConfigurationGetVolumeInstanceGatewayTimeout) Code() int {
	return 504
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationGetVolumeInstanceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationGetVolumeInstanceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationGetVolumeInstanceDefault creates a VolumeInstanceConfigurationGetVolumeInstanceDefault with default headers values
func NewVolumeInstanceConfigurationGetVolumeInstanceDefault(code int) *VolumeInstanceConfigurationGetVolumeInstanceDefault {
	return &VolumeInstanceConfigurationGetVolumeInstanceDefault{
		_statusCode: code,
	}
}

/*
VolumeInstanceConfigurationGetVolumeInstanceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type VolumeInstanceConfigurationGetVolumeInstanceDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this volume instance configuration get volume instance default response has a 2xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this volume instance configuration get volume instance default response has a 3xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this volume instance configuration get volume instance default response has a 4xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this volume instance configuration get volume instance default response has a 5xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this volume instance configuration get volume instance default response a status code equal to that given
func (o *VolumeInstanceConfigurationGetVolumeInstanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the volume instance configuration get volume instance default response
func (o *VolumeInstanceConfigurationGetVolumeInstanceDefault) Code() int {
	return o._statusCode
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceDefault) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}][%d] VolumeInstanceConfiguration_GetVolumeInstance default  %+v", o._statusCode, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceDefault) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/id/{id}][%d] VolumeInstanceConfiguration_GetVolumeInstance default  %+v", o._statusCode, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
