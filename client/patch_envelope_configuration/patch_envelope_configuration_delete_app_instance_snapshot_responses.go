// Code generated by go-swagger; DO NOT EDIT.

package patch_envelope_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// PatchEnvelopeConfigurationDeleteAppInstanceSnapshotReader is a Reader for the PatchEnvelopeConfigurationDeleteAppInstanceSnapshot structure.
type PatchEnvelopeConfigurationDeleteAppInstanceSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchEnvelopeConfigurationDeleteAppInstanceSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchEnvelopeConfigurationDeleteAppInstanceSnapshotUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchEnvelopeConfigurationDeleteAppInstanceSnapshotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchEnvelopeConfigurationDeleteAppInstanceSnapshotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchEnvelopeConfigurationDeleteAppInstanceSnapshotGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchEnvelopeConfigurationDeleteAppInstanceSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchEnvelopeConfigurationDeleteAppInstanceSnapshotOK creates a PatchEnvelopeConfigurationDeleteAppInstanceSnapshotOK with default headers values
func NewPatchEnvelopeConfigurationDeleteAppInstanceSnapshotOK() *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotOK {
	return &PatchEnvelopeConfigurationDeleteAppInstanceSnapshotOK{}
}

/*
PatchEnvelopeConfigurationDeleteAppInstanceSnapshotOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchEnvelopeConfigurationDeleteAppInstanceSnapshotOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration delete app instance snapshot o k response has a 2xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch envelope configuration delete app instance snapshot o k response has a 3xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration delete app instance snapshot o k response has a 4xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch envelope configuration delete app instance snapshot o k response has a 5xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch envelope configuration delete app instance snapshot o k response a status code equal to that given
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch envelope configuration delete app instance snapshot o k response
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotOK) Code() int {
	return 200
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationDeleteAppInstanceSnapshotOK  %+v", 200, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotOK) String() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationDeleteAppInstanceSnapshotOK  %+v", 200, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationDeleteAppInstanceSnapshotUnauthorized creates a PatchEnvelopeConfigurationDeleteAppInstanceSnapshotUnauthorized with default headers values
func NewPatchEnvelopeConfigurationDeleteAppInstanceSnapshotUnauthorized() *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotUnauthorized {
	return &PatchEnvelopeConfigurationDeleteAppInstanceSnapshotUnauthorized{}
}

/*
PatchEnvelopeConfigurationDeleteAppInstanceSnapshotUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type PatchEnvelopeConfigurationDeleteAppInstanceSnapshotUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration delete app instance snapshot unauthorized response has a 2xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration delete app instance snapshot unauthorized response has a 3xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration delete app instance snapshot unauthorized response has a 4xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch envelope configuration delete app instance snapshot unauthorized response has a 5xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch envelope configuration delete app instance snapshot unauthorized response a status code equal to that given
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the patch envelope configuration delete app instance snapshot unauthorized response
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotUnauthorized) Code() int {
	return 401
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationDeleteAppInstanceSnapshotUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationDeleteAppInstanceSnapshotUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationDeleteAppInstanceSnapshotForbidden creates a PatchEnvelopeConfigurationDeleteAppInstanceSnapshotForbidden with default headers values
func NewPatchEnvelopeConfigurationDeleteAppInstanceSnapshotForbidden() *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotForbidden {
	return &PatchEnvelopeConfigurationDeleteAppInstanceSnapshotForbidden{}
}

/*
PatchEnvelopeConfigurationDeleteAppInstanceSnapshotForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type PatchEnvelopeConfigurationDeleteAppInstanceSnapshotForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration delete app instance snapshot forbidden response has a 2xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration delete app instance snapshot forbidden response has a 3xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration delete app instance snapshot forbidden response has a 4xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch envelope configuration delete app instance snapshot forbidden response has a 5xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch envelope configuration delete app instance snapshot forbidden response a status code equal to that given
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the patch envelope configuration delete app instance snapshot forbidden response
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotForbidden) Code() int {
	return 403
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationDeleteAppInstanceSnapshotForbidden  %+v", 403, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationDeleteAppInstanceSnapshotForbidden  %+v", 403, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationDeleteAppInstanceSnapshotInternalServerError creates a PatchEnvelopeConfigurationDeleteAppInstanceSnapshotInternalServerError with default headers values
func NewPatchEnvelopeConfigurationDeleteAppInstanceSnapshotInternalServerError() *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotInternalServerError {
	return &PatchEnvelopeConfigurationDeleteAppInstanceSnapshotInternalServerError{}
}

/*
PatchEnvelopeConfigurationDeleteAppInstanceSnapshotInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type PatchEnvelopeConfigurationDeleteAppInstanceSnapshotInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration delete app instance snapshot internal server error response has a 2xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration delete app instance snapshot internal server error response has a 3xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration delete app instance snapshot internal server error response has a 4xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch envelope configuration delete app instance snapshot internal server error response has a 5xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch envelope configuration delete app instance snapshot internal server error response a status code equal to that given
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the patch envelope configuration delete app instance snapshot internal server error response
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotInternalServerError) Code() int {
	return 500
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationDeleteAppInstanceSnapshotInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationDeleteAppInstanceSnapshotInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationDeleteAppInstanceSnapshotGatewayTimeout creates a PatchEnvelopeConfigurationDeleteAppInstanceSnapshotGatewayTimeout with default headers values
func NewPatchEnvelopeConfigurationDeleteAppInstanceSnapshotGatewayTimeout() *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotGatewayTimeout {
	return &PatchEnvelopeConfigurationDeleteAppInstanceSnapshotGatewayTimeout{}
}

/*
PatchEnvelopeConfigurationDeleteAppInstanceSnapshotGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type PatchEnvelopeConfigurationDeleteAppInstanceSnapshotGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration delete app instance snapshot gateway timeout response has a 2xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration delete app instance snapshot gateway timeout response has a 3xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration delete app instance snapshot gateway timeout response has a 4xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch envelope configuration delete app instance snapshot gateway timeout response has a 5xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch envelope configuration delete app instance snapshot gateway timeout response a status code equal to that given
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the patch envelope configuration delete app instance snapshot gateway timeout response
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotGatewayTimeout) Code() int {
	return 504
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationDeleteAppInstanceSnapshotGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationDeleteAppInstanceSnapshotGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationDeleteAppInstanceSnapshotDefault creates a PatchEnvelopeConfigurationDeleteAppInstanceSnapshotDefault with default headers values
func NewPatchEnvelopeConfigurationDeleteAppInstanceSnapshotDefault(code int) *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotDefault {
	return &PatchEnvelopeConfigurationDeleteAppInstanceSnapshotDefault{
		_statusCode: code,
	}
}

/*
PatchEnvelopeConfigurationDeleteAppInstanceSnapshotDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PatchEnvelopeConfigurationDeleteAppInstanceSnapshotDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this patch envelope configuration delete app instance snapshot default response has a 2xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch envelope configuration delete app instance snapshot default response has a 3xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch envelope configuration delete app instance snapshot default response has a 4xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch envelope configuration delete app instance snapshot default response has a 5xx status code
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch envelope configuration delete app instance snapshot default response a status code equal to that given
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch envelope configuration delete app instance snapshot default response
func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] PatchEnvelopeConfiguration_DeleteAppInstanceSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] PatchEnvelopeConfiguration_DeleteAppInstanceSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationDeleteAppInstanceSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
