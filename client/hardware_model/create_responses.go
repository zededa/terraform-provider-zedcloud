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

// HardwareModelCreateHardwareModelReader is a Reader for the HardwareModelCreateHardwareModel structure.
type HardwareModelCreateHardwareModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareModelCreateHardwareModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHardwareModelCreateHardwareModelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHardwareModelCreateHardwareModelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewHardwareModelCreateHardwareModelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHardwareModelCreateHardwareModelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewHardwareModelCreateHardwareModelConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHardwareModelCreateHardwareModelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewHardwareModelCreateHardwareModelGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHardwareModelCreateHardwareModelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHardwareModelCreateHardwareModelOK creates a HardwareModelCreateHardwareModelOK with default headers values
func NewHardwareModelCreateHardwareModelOK() *HardwareModelCreateHardwareModelOK {
	return &HardwareModelCreateHardwareModelOK{}
}

/*
HardwareModelCreateHardwareModelOK describes a response with status code 200, with default header values.

A successful response.
*/
type HardwareModelCreateHardwareModelOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model create hardware model o k response has a 2xx status code
func (o *HardwareModelCreateHardwareModelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hardware model create hardware model o k response has a 3xx status code
func (o *HardwareModelCreateHardwareModelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model create hardware model o k response has a 4xx status code
func (o *HardwareModelCreateHardwareModelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model create hardware model o k response has a 5xx status code
func (o *HardwareModelCreateHardwareModelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model create hardware model o k response a status code equal to that given
func (o *HardwareModelCreateHardwareModelOK) IsCode(code int) bool {
	return code == 200
}

func (o *HardwareModelCreateHardwareModelOK) Error() string {
	return fmt.Sprintf("[POST /v1/sysmodels][%d] hardwareModelCreateHardwareModelOK  %+v", 200, o.Payload)
}

func (o *HardwareModelCreateHardwareModelOK) String() string {
	return fmt.Sprintf("[POST /v1/sysmodels][%d] hardwareModelCreateHardwareModelOK  %+v", 200, o.Payload)
}

func (o *HardwareModelCreateHardwareModelOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelCreateHardwareModelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelCreateHardwareModelBadRequest creates a HardwareModelCreateHardwareModelBadRequest with default headers values
func NewHardwareModelCreateHardwareModelBadRequest() *HardwareModelCreateHardwareModelBadRequest {
	return &HardwareModelCreateHardwareModelBadRequest{}
}

/*
HardwareModelCreateHardwareModelBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type HardwareModelCreateHardwareModelBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model create hardware model bad request response has a 2xx status code
func (o *HardwareModelCreateHardwareModelBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model create hardware model bad request response has a 3xx status code
func (o *HardwareModelCreateHardwareModelBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model create hardware model bad request response has a 4xx status code
func (o *HardwareModelCreateHardwareModelBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model create hardware model bad request response has a 5xx status code
func (o *HardwareModelCreateHardwareModelBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model create hardware model bad request response a status code equal to that given
func (o *HardwareModelCreateHardwareModelBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *HardwareModelCreateHardwareModelBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/sysmodels][%d] hardwareModelCreateHardwareModelBadRequest  %+v", 400, o.Payload)
}

func (o *HardwareModelCreateHardwareModelBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/sysmodels][%d] hardwareModelCreateHardwareModelBadRequest  %+v", 400, o.Payload)
}

func (o *HardwareModelCreateHardwareModelBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelCreateHardwareModelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelCreateHardwareModelUnauthorized creates a HardwareModelCreateHardwareModelUnauthorized with default headers values
func NewHardwareModelCreateHardwareModelUnauthorized() *HardwareModelCreateHardwareModelUnauthorized {
	return &HardwareModelCreateHardwareModelUnauthorized{}
}

/*
HardwareModelCreateHardwareModelUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type HardwareModelCreateHardwareModelUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model create hardware model unauthorized response has a 2xx status code
func (o *HardwareModelCreateHardwareModelUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model create hardware model unauthorized response has a 3xx status code
func (o *HardwareModelCreateHardwareModelUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model create hardware model unauthorized response has a 4xx status code
func (o *HardwareModelCreateHardwareModelUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model create hardware model unauthorized response has a 5xx status code
func (o *HardwareModelCreateHardwareModelUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model create hardware model unauthorized response a status code equal to that given
func (o *HardwareModelCreateHardwareModelUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *HardwareModelCreateHardwareModelUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/sysmodels][%d] hardwareModelCreateHardwareModelUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelCreateHardwareModelUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/sysmodels][%d] hardwareModelCreateHardwareModelUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelCreateHardwareModelUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelCreateHardwareModelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelCreateHardwareModelForbidden creates a HardwareModelCreateHardwareModelForbidden with default headers values
func NewHardwareModelCreateHardwareModelForbidden() *HardwareModelCreateHardwareModelForbidden {
	return &HardwareModelCreateHardwareModelForbidden{}
}

/*
HardwareModelCreateHardwareModelForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type HardwareModelCreateHardwareModelForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model create hardware model forbidden response has a 2xx status code
func (o *HardwareModelCreateHardwareModelForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model create hardware model forbidden response has a 3xx status code
func (o *HardwareModelCreateHardwareModelForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model create hardware model forbidden response has a 4xx status code
func (o *HardwareModelCreateHardwareModelForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model create hardware model forbidden response has a 5xx status code
func (o *HardwareModelCreateHardwareModelForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model create hardware model forbidden response a status code equal to that given
func (o *HardwareModelCreateHardwareModelForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *HardwareModelCreateHardwareModelForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/sysmodels][%d] hardwareModelCreateHardwareModelForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelCreateHardwareModelForbidden) String() string {
	return fmt.Sprintf("[POST /v1/sysmodels][%d] hardwareModelCreateHardwareModelForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelCreateHardwareModelForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelCreateHardwareModelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelCreateHardwareModelConflict creates a HardwareModelCreateHardwareModelConflict with default headers values
func NewHardwareModelCreateHardwareModelConflict() *HardwareModelCreateHardwareModelConflict {
	return &HardwareModelCreateHardwareModelConflict{}
}

/*
HardwareModelCreateHardwareModelConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this hardware model record will conflict with an already existing hardware model record.
*/
type HardwareModelCreateHardwareModelConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model create hardware model conflict response has a 2xx status code
func (o *HardwareModelCreateHardwareModelConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model create hardware model conflict response has a 3xx status code
func (o *HardwareModelCreateHardwareModelConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model create hardware model conflict response has a 4xx status code
func (o *HardwareModelCreateHardwareModelConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model create hardware model conflict response has a 5xx status code
func (o *HardwareModelCreateHardwareModelConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model create hardware model conflict response a status code equal to that given
func (o *HardwareModelCreateHardwareModelConflict) IsCode(code int) bool {
	return code == 409
}

func (o *HardwareModelCreateHardwareModelConflict) Error() string {
	return fmt.Sprintf("[POST /v1/sysmodels][%d] hardwareModelCreateHardwareModelConflict  %+v", 409, o.Payload)
}

func (o *HardwareModelCreateHardwareModelConflict) String() string {
	return fmt.Sprintf("[POST /v1/sysmodels][%d] hardwareModelCreateHardwareModelConflict  %+v", 409, o.Payload)
}

func (o *HardwareModelCreateHardwareModelConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelCreateHardwareModelConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelCreateHardwareModelInternalServerError creates a HardwareModelCreateHardwareModelInternalServerError with default headers values
func NewHardwareModelCreateHardwareModelInternalServerError() *HardwareModelCreateHardwareModelInternalServerError {
	return &HardwareModelCreateHardwareModelInternalServerError{}
}

/*
HardwareModelCreateHardwareModelInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type HardwareModelCreateHardwareModelInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model create hardware model internal server error response has a 2xx status code
func (o *HardwareModelCreateHardwareModelInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model create hardware model internal server error response has a 3xx status code
func (o *HardwareModelCreateHardwareModelInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model create hardware model internal server error response has a 4xx status code
func (o *HardwareModelCreateHardwareModelInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model create hardware model internal server error response has a 5xx status code
func (o *HardwareModelCreateHardwareModelInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model create hardware model internal server error response a status code equal to that given
func (o *HardwareModelCreateHardwareModelInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *HardwareModelCreateHardwareModelInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/sysmodels][%d] hardwareModelCreateHardwareModelInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelCreateHardwareModelInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/sysmodels][%d] hardwareModelCreateHardwareModelInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelCreateHardwareModelInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelCreateHardwareModelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelCreateHardwareModelGatewayTimeout creates a HardwareModelCreateHardwareModelGatewayTimeout with default headers values
func NewHardwareModelCreateHardwareModelGatewayTimeout() *HardwareModelCreateHardwareModelGatewayTimeout {
	return &HardwareModelCreateHardwareModelGatewayTimeout{}
}

/*
HardwareModelCreateHardwareModelGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type HardwareModelCreateHardwareModelGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model create hardware model gateway timeout response has a 2xx status code
func (o *HardwareModelCreateHardwareModelGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model create hardware model gateway timeout response has a 3xx status code
func (o *HardwareModelCreateHardwareModelGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model create hardware model gateway timeout response has a 4xx status code
func (o *HardwareModelCreateHardwareModelGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model create hardware model gateway timeout response has a 5xx status code
func (o *HardwareModelCreateHardwareModelGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model create hardware model gateway timeout response a status code equal to that given
func (o *HardwareModelCreateHardwareModelGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *HardwareModelCreateHardwareModelGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/sysmodels][%d] hardwareModelCreateHardwareModelGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelCreateHardwareModelGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/sysmodels][%d] hardwareModelCreateHardwareModelGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelCreateHardwareModelGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelCreateHardwareModelGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelCreateHardwareModelDefault creates a HardwareModelCreateHardwareModelDefault with default headers values
func NewHardwareModelCreateHardwareModelDefault(code int) *HardwareModelCreateHardwareModelDefault {
	return &HardwareModelCreateHardwareModelDefault{
		_statusCode: code,
	}
}

/*
HardwareModelCreateHardwareModelDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HardwareModelCreateHardwareModelDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the hardware model create hardware model default response
func (o *HardwareModelCreateHardwareModelDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this hardware model create hardware model default response has a 2xx status code
func (o *HardwareModelCreateHardwareModelDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this hardware model create hardware model default response has a 3xx status code
func (o *HardwareModelCreateHardwareModelDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this hardware model create hardware model default response has a 4xx status code
func (o *HardwareModelCreateHardwareModelDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this hardware model create hardware model default response has a 5xx status code
func (o *HardwareModelCreateHardwareModelDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this hardware model create hardware model default response a status code equal to that given
func (o *HardwareModelCreateHardwareModelDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *HardwareModelCreateHardwareModelDefault) Error() string {
	return fmt.Sprintf("[POST /v1/sysmodels][%d] HardwareModel_CreateHardwareModel default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelCreateHardwareModelDefault) String() string {
	return fmt.Sprintf("[POST /v1/sysmodels][%d] HardwareModel_CreateHardwareModel default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelCreateHardwareModelDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *HardwareModelCreateHardwareModelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
