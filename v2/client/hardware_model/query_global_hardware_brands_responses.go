// Code generated by go-swagger; DO NOT EDIT.

package hardware_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// HardwareModelQueryGlobalHardwareBrandsReader is a Reader for the HardwareModelQueryGlobalHardwareBrands structure.
type HardwareModelQueryGlobalHardwareBrandsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareModelQueryGlobalHardwareBrandsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHardwareModelQueryGlobalHardwareBrandsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHardwareModelQueryGlobalHardwareBrandsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewHardwareModelQueryGlobalHardwareBrandsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHardwareModelQueryGlobalHardwareBrandsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHardwareModelQueryGlobalHardwareBrandsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewHardwareModelQueryGlobalHardwareBrandsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHardwareModelQueryGlobalHardwareBrandsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHardwareModelQueryGlobalHardwareBrandsOK creates a HardwareModelQueryGlobalHardwareBrandsOK with default headers values
func NewHardwareModelQueryGlobalHardwareBrandsOK() *HardwareModelQueryGlobalHardwareBrandsOK {
	return &HardwareModelQueryGlobalHardwareBrandsOK{}
}

/*
HardwareModelQueryGlobalHardwareBrandsOK describes a response with status code 200, with default header values.

A successful response.
*/
type HardwareModelQueryGlobalHardwareBrandsOK struct {
	Payload *models.SysBrands
}

// IsSuccess returns true when this hardware model query global hardware brands o k response has a 2xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hardware model query global hardware brands o k response has a 3xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model query global hardware brands o k response has a 4xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model query global hardware brands o k response has a 5xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model query global hardware brands o k response a status code equal to that given
func (o *HardwareModelQueryGlobalHardwareBrandsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the hardware model query global hardware brands o k response
func (o *HardwareModelQueryGlobalHardwareBrandsOK) Code() int {
	return 200
}

func (o *HardwareModelQueryGlobalHardwareBrandsOK) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global][%d] hardwareModelQueryGlobalHardwareBrandsOK  %+v", 200, o.Payload)
}

func (o *HardwareModelQueryGlobalHardwareBrandsOK) String() string {
	return fmt.Sprintf("[GET /v1/brands/global][%d] hardwareModelQueryGlobalHardwareBrandsOK  %+v", 200, o.Payload)
}

func (o *HardwareModelQueryGlobalHardwareBrandsOK) GetPayload() *models.SysBrands {
	return o.Payload
}

func (o *HardwareModelQueryGlobalHardwareBrandsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SysBrands)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryGlobalHardwareBrandsBadRequest creates a HardwareModelQueryGlobalHardwareBrandsBadRequest with default headers values
func NewHardwareModelQueryGlobalHardwareBrandsBadRequest() *HardwareModelQueryGlobalHardwareBrandsBadRequest {
	return &HardwareModelQueryGlobalHardwareBrandsBadRequest{}
}

/*
HardwareModelQueryGlobalHardwareBrandsBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type HardwareModelQueryGlobalHardwareBrandsBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model query global hardware brands bad request response has a 2xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model query global hardware brands bad request response has a 3xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model query global hardware brands bad request response has a 4xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model query global hardware brands bad request response has a 5xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model query global hardware brands bad request response a status code equal to that given
func (o *HardwareModelQueryGlobalHardwareBrandsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the hardware model query global hardware brands bad request response
func (o *HardwareModelQueryGlobalHardwareBrandsBadRequest) Code() int {
	return 400
}

func (o *HardwareModelQueryGlobalHardwareBrandsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global][%d] hardwareModelQueryGlobalHardwareBrandsBadRequest  %+v", 400, o.Payload)
}

func (o *HardwareModelQueryGlobalHardwareBrandsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/brands/global][%d] hardwareModelQueryGlobalHardwareBrandsBadRequest  %+v", 400, o.Payload)
}

func (o *HardwareModelQueryGlobalHardwareBrandsBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelQueryGlobalHardwareBrandsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryGlobalHardwareBrandsUnauthorized creates a HardwareModelQueryGlobalHardwareBrandsUnauthorized with default headers values
func NewHardwareModelQueryGlobalHardwareBrandsUnauthorized() *HardwareModelQueryGlobalHardwareBrandsUnauthorized {
	return &HardwareModelQueryGlobalHardwareBrandsUnauthorized{}
}

/*
HardwareModelQueryGlobalHardwareBrandsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type HardwareModelQueryGlobalHardwareBrandsUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model query global hardware brands unauthorized response has a 2xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model query global hardware brands unauthorized response has a 3xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model query global hardware brands unauthorized response has a 4xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model query global hardware brands unauthorized response has a 5xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model query global hardware brands unauthorized response a status code equal to that given
func (o *HardwareModelQueryGlobalHardwareBrandsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the hardware model query global hardware brands unauthorized response
func (o *HardwareModelQueryGlobalHardwareBrandsUnauthorized) Code() int {
	return 401
}

func (o *HardwareModelQueryGlobalHardwareBrandsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global][%d] hardwareModelQueryGlobalHardwareBrandsUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelQueryGlobalHardwareBrandsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/brands/global][%d] hardwareModelQueryGlobalHardwareBrandsUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelQueryGlobalHardwareBrandsUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelQueryGlobalHardwareBrandsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryGlobalHardwareBrandsForbidden creates a HardwareModelQueryGlobalHardwareBrandsForbidden with default headers values
func NewHardwareModelQueryGlobalHardwareBrandsForbidden() *HardwareModelQueryGlobalHardwareBrandsForbidden {
	return &HardwareModelQueryGlobalHardwareBrandsForbidden{}
}

/*
HardwareModelQueryGlobalHardwareBrandsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type HardwareModelQueryGlobalHardwareBrandsForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model query global hardware brands forbidden response has a 2xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model query global hardware brands forbidden response has a 3xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model query global hardware brands forbidden response has a 4xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model query global hardware brands forbidden response has a 5xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model query global hardware brands forbidden response a status code equal to that given
func (o *HardwareModelQueryGlobalHardwareBrandsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the hardware model query global hardware brands forbidden response
func (o *HardwareModelQueryGlobalHardwareBrandsForbidden) Code() int {
	return 403
}

func (o *HardwareModelQueryGlobalHardwareBrandsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global][%d] hardwareModelQueryGlobalHardwareBrandsForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelQueryGlobalHardwareBrandsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/brands/global][%d] hardwareModelQueryGlobalHardwareBrandsForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelQueryGlobalHardwareBrandsForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelQueryGlobalHardwareBrandsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryGlobalHardwareBrandsInternalServerError creates a HardwareModelQueryGlobalHardwareBrandsInternalServerError with default headers values
func NewHardwareModelQueryGlobalHardwareBrandsInternalServerError() *HardwareModelQueryGlobalHardwareBrandsInternalServerError {
	return &HardwareModelQueryGlobalHardwareBrandsInternalServerError{}
}

/*
HardwareModelQueryGlobalHardwareBrandsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type HardwareModelQueryGlobalHardwareBrandsInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model query global hardware brands internal server error response has a 2xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model query global hardware brands internal server error response has a 3xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model query global hardware brands internal server error response has a 4xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model query global hardware brands internal server error response has a 5xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model query global hardware brands internal server error response a status code equal to that given
func (o *HardwareModelQueryGlobalHardwareBrandsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the hardware model query global hardware brands internal server error response
func (o *HardwareModelQueryGlobalHardwareBrandsInternalServerError) Code() int {
	return 500
}

func (o *HardwareModelQueryGlobalHardwareBrandsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global][%d] hardwareModelQueryGlobalHardwareBrandsInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelQueryGlobalHardwareBrandsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/brands/global][%d] hardwareModelQueryGlobalHardwareBrandsInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelQueryGlobalHardwareBrandsInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelQueryGlobalHardwareBrandsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryGlobalHardwareBrandsGatewayTimeout creates a HardwareModelQueryGlobalHardwareBrandsGatewayTimeout with default headers values
func NewHardwareModelQueryGlobalHardwareBrandsGatewayTimeout() *HardwareModelQueryGlobalHardwareBrandsGatewayTimeout {
	return &HardwareModelQueryGlobalHardwareBrandsGatewayTimeout{}
}

/*
HardwareModelQueryGlobalHardwareBrandsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type HardwareModelQueryGlobalHardwareBrandsGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model query global hardware brands gateway timeout response has a 2xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model query global hardware brands gateway timeout response has a 3xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model query global hardware brands gateway timeout response has a 4xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model query global hardware brands gateway timeout response has a 5xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model query global hardware brands gateway timeout response a status code equal to that given
func (o *HardwareModelQueryGlobalHardwareBrandsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the hardware model query global hardware brands gateway timeout response
func (o *HardwareModelQueryGlobalHardwareBrandsGatewayTimeout) Code() int {
	return 504
}

func (o *HardwareModelQueryGlobalHardwareBrandsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global][%d] hardwareModelQueryGlobalHardwareBrandsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelQueryGlobalHardwareBrandsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/brands/global][%d] hardwareModelQueryGlobalHardwareBrandsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelQueryGlobalHardwareBrandsGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelQueryGlobalHardwareBrandsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryGlobalHardwareBrandsDefault creates a HardwareModelQueryGlobalHardwareBrandsDefault with default headers values
func NewHardwareModelQueryGlobalHardwareBrandsDefault(code int) *HardwareModelQueryGlobalHardwareBrandsDefault {
	return &HardwareModelQueryGlobalHardwareBrandsDefault{
		_statusCode: code,
	}
}

/*
HardwareModelQueryGlobalHardwareBrandsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HardwareModelQueryGlobalHardwareBrandsDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this hardware model query global hardware brands default response has a 2xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this hardware model query global hardware brands default response has a 3xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this hardware model query global hardware brands default response has a 4xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this hardware model query global hardware brands default response has a 5xx status code
func (o *HardwareModelQueryGlobalHardwareBrandsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this hardware model query global hardware brands default response a status code equal to that given
func (o *HardwareModelQueryGlobalHardwareBrandsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the hardware model query global hardware brands default response
func (o *HardwareModelQueryGlobalHardwareBrandsDefault) Code() int {
	return o._statusCode
}

func (o *HardwareModelQueryGlobalHardwareBrandsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global][%d] HardwareModel_QueryGlobalHardwareBrands default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelQueryGlobalHardwareBrandsDefault) String() string {
	return fmt.Sprintf("[GET /v1/brands/global][%d] HardwareModel_QueryGlobalHardwareBrands default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelQueryGlobalHardwareBrandsDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *HardwareModelQueryGlobalHardwareBrandsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
