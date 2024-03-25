package volume_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// VolumeInstanceConfigurationCreateVolumeInstanceReader is a Reader for the VolumeInstanceConfigurationCreateVolumeInstance structure.
type VolumeInstanceConfigurationCreateVolumeInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeInstanceConfigurationCreateVolumeInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeInstanceConfigurationCreateVolumeInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVolumeInstanceConfigurationCreateVolumeInstanceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewVolumeInstanceConfigurationCreateVolumeInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewVolumeInstanceConfigurationCreateVolumeInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewVolumeInstanceConfigurationCreateVolumeInstanceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewVolumeInstanceConfigurationCreateVolumeInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewVolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewVolumeInstanceConfigurationCreateVolumeInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeInstanceConfigurationCreateVolumeInstanceOK creates a VolumeInstanceConfigurationCreateVolumeInstanceOK with default headers values
func NewVolumeInstanceConfigurationCreateVolumeInstanceOK() *VolumeInstanceConfigurationCreateVolumeInstanceOK {
	return &VolumeInstanceConfigurationCreateVolumeInstanceOK{}
}

/*
VolumeInstanceConfigurationCreateVolumeInstanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type VolumeInstanceConfigurationCreateVolumeInstanceOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration create volume instance o k response has a 2xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volume instance configuration create volume instance o k response has a 3xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration create volume instance o k response has a 4xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance configuration create volume instance o k response has a 5xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration create volume instance o k response a status code equal to that given
func (o *VolumeInstanceConfigurationCreateVolumeInstanceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the volume instance configuration create volume instance o k response
func (o *VolumeInstanceConfigurationCreateVolumeInstanceOK) Code() int {
	return 200
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceOK) Error() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceOK  %+v", 200, o.Payload)
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceOK) String() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceOK  %+v", 200, o.Payload)
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationCreateVolumeInstanceBadRequest creates a VolumeInstanceConfigurationCreateVolumeInstanceBadRequest with default headers values
func NewVolumeInstanceConfigurationCreateVolumeInstanceBadRequest() *VolumeInstanceConfigurationCreateVolumeInstanceBadRequest {
	return &VolumeInstanceConfigurationCreateVolumeInstanceBadRequest{}
}

/*
VolumeInstanceConfigurationCreateVolumeInstanceBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type VolumeInstanceConfigurationCreateVolumeInstanceBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration create volume instance bad request response has a 2xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration create volume instance bad request response has a 3xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration create volume instance bad request response has a 4xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance configuration create volume instance bad request response has a 5xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration create volume instance bad request response a status code equal to that given
func (o *VolumeInstanceConfigurationCreateVolumeInstanceBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the volume instance configuration create volume instance bad request response
func (o *VolumeInstanceConfigurationCreateVolumeInstanceBadRequest) Code() int {
	return 400
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceBadRequest  %+v", 400, o.Payload)
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceBadRequest  %+v", 400, o.Payload)
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationCreateVolumeInstanceUnauthorized creates a VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized with default headers values
func NewVolumeInstanceConfigurationCreateVolumeInstanceUnauthorized() *VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized {
	return &VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized{}
}

/*
VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration create volume instance unauthorized response has a 2xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration create volume instance unauthorized response has a 3xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration create volume instance unauthorized response has a 4xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance configuration create volume instance unauthorized response has a 5xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration create volume instance unauthorized response a status code equal to that given
func (o *VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the volume instance configuration create volume instance unauthorized response
func (o *VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized) Code() int {
	return 401
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationCreateVolumeInstanceForbidden creates a VolumeInstanceConfigurationCreateVolumeInstanceForbidden with default headers values
func NewVolumeInstanceConfigurationCreateVolumeInstanceForbidden() *VolumeInstanceConfigurationCreateVolumeInstanceForbidden {
	return &VolumeInstanceConfigurationCreateVolumeInstanceForbidden{}
}

/*
VolumeInstanceConfigurationCreateVolumeInstanceForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type VolumeInstanceConfigurationCreateVolumeInstanceForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration create volume instance forbidden response has a 2xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration create volume instance forbidden response has a 3xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration create volume instance forbidden response has a 4xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance configuration create volume instance forbidden response has a 5xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration create volume instance forbidden response a status code equal to that given
func (o *VolumeInstanceConfigurationCreateVolumeInstanceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the volume instance configuration create volume instance forbidden response
func (o *VolumeInstanceConfigurationCreateVolumeInstanceForbidden) Code() int {
	return 403
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceForbidden  %+v", 403, o.Payload)
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceForbidden) String() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceForbidden  %+v", 403, o.Payload)
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationCreateVolumeInstanceConflict creates a VolumeInstanceConfigurationCreateVolumeInstanceConflict with default headers values
func NewVolumeInstanceConfigurationCreateVolumeInstanceConflict() *VolumeInstanceConfigurationCreateVolumeInstanceConflict {
	return &VolumeInstanceConfigurationCreateVolumeInstanceConflict{}
}

/*
VolumeInstanceConfigurationCreateVolumeInstanceConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this edge volume instance record will conflict with an already existing edge volume instance record.
*/
type VolumeInstanceConfigurationCreateVolumeInstanceConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration create volume instance conflict response has a 2xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration create volume instance conflict response has a 3xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration create volume instance conflict response has a 4xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance configuration create volume instance conflict response has a 5xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration create volume instance conflict response a status code equal to that given
func (o *VolumeInstanceConfigurationCreateVolumeInstanceConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the volume instance configuration create volume instance conflict response
func (o *VolumeInstanceConfigurationCreateVolumeInstanceConflict) Code() int {
	return 409
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceConflict) Error() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceConflict  %+v", 409, o.Payload)
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceConflict) String() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceConflict  %+v", 409, o.Payload)
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationCreateVolumeInstanceInternalServerError creates a VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError with default headers values
func NewVolumeInstanceConfigurationCreateVolumeInstanceInternalServerError() *VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError {
	return &VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError{}
}

/*
VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration create volume instance internal server error response has a 2xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration create volume instance internal server error response has a 3xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration create volume instance internal server error response has a 4xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance configuration create volume instance internal server error response has a 5xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this volume instance configuration create volume instance internal server error response a status code equal to that given
func (o *VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the volume instance configuration create volume instance internal server error response
func (o *VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError) Code() int {
	return 500
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout creates a VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout with default headers values
func NewVolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout() *VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout {
	return &VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout{}
}

/*
VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration create volume instance gateway timeout response has a 2xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration create volume instance gateway timeout response has a 3xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration create volume instance gateway timeout response has a 4xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance configuration create volume instance gateway timeout response has a 5xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this volume instance configuration create volume instance gateway timeout response a status code equal to that given
func (o *VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the volume instance configuration create volume instance gateway timeout response
func (o *VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout) Code() int {
	return 504
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] volumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationCreateVolumeInstanceDefault creates a VolumeInstanceConfigurationCreateVolumeInstanceDefault with default headers values
func NewVolumeInstanceConfigurationCreateVolumeInstanceDefault(code int) *VolumeInstanceConfigurationCreateVolumeInstanceDefault {
	return &VolumeInstanceConfigurationCreateVolumeInstanceDefault{
		_statusCode: code,
	}
}

/*
VolumeInstanceConfigurationCreateVolumeInstanceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type VolumeInstanceConfigurationCreateVolumeInstanceDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this volume instance configuration create volume instance default response has a 2xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this volume instance configuration create volume instance default response has a 3xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this volume instance configuration create volume instance default response has a 4xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this volume instance configuration create volume instance default response has a 5xx status code
func (o *VolumeInstanceConfigurationCreateVolumeInstanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this volume instance configuration create volume instance default response a status code equal to that given
func (o *VolumeInstanceConfigurationCreateVolumeInstanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the volume instance configuration create volume instance default response
func (o *VolumeInstanceConfigurationCreateVolumeInstanceDefault) Code() int {
	return o._statusCode
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] VolumeInstanceConfiguration_CreateVolumeInstance default  %+v", o._statusCode, o.Payload)
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceDefault) String() string {
	return fmt.Sprintf("[POST /v1/volumes/instances][%d] VolumeInstanceConfiguration_CreateVolumeInstance default  %+v", o._statusCode, o.Payload)
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *VolumeInstanceConfigurationCreateVolumeInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
