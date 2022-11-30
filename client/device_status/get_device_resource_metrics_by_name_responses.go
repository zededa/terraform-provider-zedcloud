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

// GetDeviceResourceMetricsByNameReader is a Reader for the GetDeviceResourceMetricsByName structure.
type GetDeviceResourceMetricsByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceResourceMetricsByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceResourceMetricsByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDeviceResourceMetricsByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDeviceResourceMetricsByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeviceResourceMetricsByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDeviceResourceMetricsByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetDeviceResourceMetricsByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeviceResourceMetricsByNameOK creates a GetDeviceResourceMetricsByNameOK with default headers values
func NewGetDeviceResourceMetricsByNameOK() *GetDeviceResourceMetricsByNameOK {
	return &GetDeviceResourceMetricsByNameOK{}
}

/*
GetDeviceResourceMetricsByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetDeviceResourceMetricsByNameOK struct {
	Payload *models.MetricQueryResponse
}

// IsSuccess returns true when this get device resource metrics by name o k response has a 2xx status code
func (o *GetDeviceResourceMetricsByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device resource metrics by name o k response has a 3xx status code
func (o *GetDeviceResourceMetricsByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device resource metrics by name o k response has a 4xx status code
func (o *GetDeviceResourceMetricsByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device resource metrics by name o k response has a 5xx status code
func (o *GetDeviceResourceMetricsByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device resource metrics by name o k response a status code equal to that given
func (o *GetDeviceResourceMetricsByNameOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDeviceResourceMetricsByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{objname}/timeSeries/{mType}][%d] getDeviceResourceMetricsByNameOK  %+v", 200, o.Payload)
}

func (o *GetDeviceResourceMetricsByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{objname}/timeSeries/{mType}][%d] getDeviceResourceMetricsByNameOK  %+v", 200, o.Payload)
}

func (o *GetDeviceResourceMetricsByNameOK) GetPayload() *models.MetricQueryResponse {
	return o.Payload
}

func (o *GetDeviceResourceMetricsByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetricQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceResourceMetricsByNameUnauthorized creates a GetDeviceResourceMetricsByNameUnauthorized with default headers values
func NewGetDeviceResourceMetricsByNameUnauthorized() *GetDeviceResourceMetricsByNameUnauthorized {
	return &GetDeviceResourceMetricsByNameUnauthorized{}
}

/*
GetDeviceResourceMetricsByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetDeviceResourceMetricsByNameUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get device resource metrics by name unauthorized response has a 2xx status code
func (o *GetDeviceResourceMetricsByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get device resource metrics by name unauthorized response has a 3xx status code
func (o *GetDeviceResourceMetricsByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device resource metrics by name unauthorized response has a 4xx status code
func (o *GetDeviceResourceMetricsByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get device resource metrics by name unauthorized response has a 5xx status code
func (o *GetDeviceResourceMetricsByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get device resource metrics by name unauthorized response a status code equal to that given
func (o *GetDeviceResourceMetricsByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetDeviceResourceMetricsByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{objname}/timeSeries/{mType}][%d] getDeviceResourceMetricsByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDeviceResourceMetricsByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{objname}/timeSeries/{mType}][%d] getDeviceResourceMetricsByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDeviceResourceMetricsByNameUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetDeviceResourceMetricsByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceResourceMetricsByNameForbidden creates a GetDeviceResourceMetricsByNameForbidden with default headers values
func NewGetDeviceResourceMetricsByNameForbidden() *GetDeviceResourceMetricsByNameForbidden {
	return &GetDeviceResourceMetricsByNameForbidden{}
}

/*
GetDeviceResourceMetricsByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetDeviceResourceMetricsByNameForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get device resource metrics by name forbidden response has a 2xx status code
func (o *GetDeviceResourceMetricsByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get device resource metrics by name forbidden response has a 3xx status code
func (o *GetDeviceResourceMetricsByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device resource metrics by name forbidden response has a 4xx status code
func (o *GetDeviceResourceMetricsByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get device resource metrics by name forbidden response has a 5xx status code
func (o *GetDeviceResourceMetricsByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get device resource metrics by name forbidden response a status code equal to that given
func (o *GetDeviceResourceMetricsByNameForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetDeviceResourceMetricsByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{objname}/timeSeries/{mType}][%d] getDeviceResourceMetricsByNameForbidden  %+v", 403, o.Payload)
}

func (o *GetDeviceResourceMetricsByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{objname}/timeSeries/{mType}][%d] getDeviceResourceMetricsByNameForbidden  %+v", 403, o.Payload)
}

func (o *GetDeviceResourceMetricsByNameForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetDeviceResourceMetricsByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceResourceMetricsByNameNotFound creates a GetDeviceResourceMetricsByNameNotFound with default headers values
func NewGetDeviceResourceMetricsByNameNotFound() *GetDeviceResourceMetricsByNameNotFound {
	return &GetDeviceResourceMetricsByNameNotFound{}
}

/*
GetDeviceResourceMetricsByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetDeviceResourceMetricsByNameNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get device resource metrics by name not found response has a 2xx status code
func (o *GetDeviceResourceMetricsByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get device resource metrics by name not found response has a 3xx status code
func (o *GetDeviceResourceMetricsByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device resource metrics by name not found response has a 4xx status code
func (o *GetDeviceResourceMetricsByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get device resource metrics by name not found response has a 5xx status code
func (o *GetDeviceResourceMetricsByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get device resource metrics by name not found response a status code equal to that given
func (o *GetDeviceResourceMetricsByNameNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetDeviceResourceMetricsByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{objname}/timeSeries/{mType}][%d] getDeviceResourceMetricsByNameNotFound  %+v", 404, o.Payload)
}

func (o *GetDeviceResourceMetricsByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{objname}/timeSeries/{mType}][%d] getDeviceResourceMetricsByNameNotFound  %+v", 404, o.Payload)
}

func (o *GetDeviceResourceMetricsByNameNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetDeviceResourceMetricsByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceResourceMetricsByNameInternalServerError creates a GetDeviceResourceMetricsByNameInternalServerError with default headers values
func NewGetDeviceResourceMetricsByNameInternalServerError() *GetDeviceResourceMetricsByNameInternalServerError {
	return &GetDeviceResourceMetricsByNameInternalServerError{}
}

/*
GetDeviceResourceMetricsByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetDeviceResourceMetricsByNameInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get device resource metrics by name internal server error response has a 2xx status code
func (o *GetDeviceResourceMetricsByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get device resource metrics by name internal server error response has a 3xx status code
func (o *GetDeviceResourceMetricsByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device resource metrics by name internal server error response has a 4xx status code
func (o *GetDeviceResourceMetricsByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device resource metrics by name internal server error response has a 5xx status code
func (o *GetDeviceResourceMetricsByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get device resource metrics by name internal server error response a status code equal to that given
func (o *GetDeviceResourceMetricsByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetDeviceResourceMetricsByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{objname}/timeSeries/{mType}][%d] getDeviceResourceMetricsByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDeviceResourceMetricsByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{objname}/timeSeries/{mType}][%d] getDeviceResourceMetricsByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDeviceResourceMetricsByNameInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetDeviceResourceMetricsByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceResourceMetricsByNameGatewayTimeout creates a GetDeviceResourceMetricsByNameGatewayTimeout with default headers values
func NewGetDeviceResourceMetricsByNameGatewayTimeout() *GetDeviceResourceMetricsByNameGatewayTimeout {
	return &GetDeviceResourceMetricsByNameGatewayTimeout{}
}

/*
GetDeviceResourceMetricsByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetDeviceResourceMetricsByNameGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this get device resource metrics by name gateway timeout response has a 2xx status code
func (o *GetDeviceResourceMetricsByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get device resource metrics by name gateway timeout response has a 3xx status code
func (o *GetDeviceResourceMetricsByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device resource metrics by name gateway timeout response has a 4xx status code
func (o *GetDeviceResourceMetricsByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device resource metrics by name gateway timeout response has a 5xx status code
func (o *GetDeviceResourceMetricsByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get device resource metrics by name gateway timeout response a status code equal to that given
func (o *GetDeviceResourceMetricsByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetDeviceResourceMetricsByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/name/{objname}/timeSeries/{mType}][%d] getDeviceResourceMetricsByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetDeviceResourceMetricsByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/devices/name/{objname}/timeSeries/{mType}][%d] getDeviceResourceMetricsByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetDeviceResourceMetricsByNameGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetDeviceResourceMetricsByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
