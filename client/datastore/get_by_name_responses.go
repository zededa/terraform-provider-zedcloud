package datastore

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// DatastoreConfigurationGetDatastoreByNameReader is a Reader for the DatastoreConfigurationGetDatastoreByName structure.
type DatastoreConfigurationGetDatastoreByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DatastoreConfigurationGetDatastoreByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDatastoreConfigurationGetDatastoreByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDatastoreConfigurationGetDatastoreByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDatastoreConfigurationGetDatastoreByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDatastoreConfigurationGetDatastoreByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDatastoreConfigurationGetDatastoreByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDatastoreConfigurationGetDatastoreByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDatastoreConfigurationGetDatastoreByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDatastoreConfigurationGetDatastoreByNameOK creates a DatastoreConfigurationGetDatastoreByNameOK with default headers values
func NewDatastoreConfigurationGetDatastoreByNameOK() *DatastoreConfigurationGetDatastoreByNameOK {
	return &DatastoreConfigurationGetDatastoreByNameOK{}
}

/*
DatastoreConfigurationGetDatastoreByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type DatastoreConfigurationGetDatastoreByNameOK struct {
	Payload *models.Datastore
}

// IsSuccess returns true when this datastore configuration get datastore by name o k response has a 2xx status code
func (o *DatastoreConfigurationGetDatastoreByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this datastore configuration get datastore by name o k response has a 3xx status code
func (o *DatastoreConfigurationGetDatastoreByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration get datastore by name o k response has a 4xx status code
func (o *DatastoreConfigurationGetDatastoreByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this datastore configuration get datastore by name o k response has a 5xx status code
func (o *DatastoreConfigurationGetDatastoreByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this datastore configuration get datastore by name o k response a status code equal to that given
func (o *DatastoreConfigurationGetDatastoreByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the datastore configuration get datastore by name o k response
func (o *DatastoreConfigurationGetDatastoreByNameOK) Code() int {
	return 200
}

func (o *DatastoreConfigurationGetDatastoreByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/name/{name}][%d] datastoreConfigurationGetDatastoreByNameOK  %+v", 200, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/datastores/name/{name}][%d] datastoreConfigurationGetDatastoreByNameOK  %+v", 200, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreByNameOK) GetPayload() *models.Datastore {
	return o.Payload
}

func (o *DatastoreConfigurationGetDatastoreByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Datastore)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationGetDatastoreByNameUnauthorized creates a DatastoreConfigurationGetDatastoreByNameUnauthorized with default headers values
func NewDatastoreConfigurationGetDatastoreByNameUnauthorized() *DatastoreConfigurationGetDatastoreByNameUnauthorized {
	return &DatastoreConfigurationGetDatastoreByNameUnauthorized{}
}

/*
DatastoreConfigurationGetDatastoreByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type DatastoreConfigurationGetDatastoreByNameUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this datastore configuration get datastore by name unauthorized response has a 2xx status code
func (o *DatastoreConfigurationGetDatastoreByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datastore configuration get datastore by name unauthorized response has a 3xx status code
func (o *DatastoreConfigurationGetDatastoreByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration get datastore by name unauthorized response has a 4xx status code
func (o *DatastoreConfigurationGetDatastoreByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this datastore configuration get datastore by name unauthorized response has a 5xx status code
func (o *DatastoreConfigurationGetDatastoreByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this datastore configuration get datastore by name unauthorized response a status code equal to that given
func (o *DatastoreConfigurationGetDatastoreByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the datastore configuration get datastore by name unauthorized response
func (o *DatastoreConfigurationGetDatastoreByNameUnauthorized) Code() int {
	return 401
}

func (o *DatastoreConfigurationGetDatastoreByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/name/{name}][%d] datastoreConfigurationGetDatastoreByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/datastores/name/{name}][%d] datastoreConfigurationGetDatastoreByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreByNameUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DatastoreConfigurationGetDatastoreByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationGetDatastoreByNameForbidden creates a DatastoreConfigurationGetDatastoreByNameForbidden with default headers values
func NewDatastoreConfigurationGetDatastoreByNameForbidden() *DatastoreConfigurationGetDatastoreByNameForbidden {
	return &DatastoreConfigurationGetDatastoreByNameForbidden{}
}

/*
DatastoreConfigurationGetDatastoreByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type DatastoreConfigurationGetDatastoreByNameForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this datastore configuration get datastore by name forbidden response has a 2xx status code
func (o *DatastoreConfigurationGetDatastoreByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datastore configuration get datastore by name forbidden response has a 3xx status code
func (o *DatastoreConfigurationGetDatastoreByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration get datastore by name forbidden response has a 4xx status code
func (o *DatastoreConfigurationGetDatastoreByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this datastore configuration get datastore by name forbidden response has a 5xx status code
func (o *DatastoreConfigurationGetDatastoreByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this datastore configuration get datastore by name forbidden response a status code equal to that given
func (o *DatastoreConfigurationGetDatastoreByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the datastore configuration get datastore by name forbidden response
func (o *DatastoreConfigurationGetDatastoreByNameForbidden) Code() int {
	return 403
}

func (o *DatastoreConfigurationGetDatastoreByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/name/{name}][%d] datastoreConfigurationGetDatastoreByNameForbidden  %+v", 403, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/datastores/name/{name}][%d] datastoreConfigurationGetDatastoreByNameForbidden  %+v", 403, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreByNameForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DatastoreConfigurationGetDatastoreByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationGetDatastoreByNameNotFound creates a DatastoreConfigurationGetDatastoreByNameNotFound with default headers values
func NewDatastoreConfigurationGetDatastoreByNameNotFound() *DatastoreConfigurationGetDatastoreByNameNotFound {
	return &DatastoreConfigurationGetDatastoreByNameNotFound{}
}

/*
DatastoreConfigurationGetDatastoreByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type DatastoreConfigurationGetDatastoreByNameNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this datastore configuration get datastore by name not found response has a 2xx status code
func (o *DatastoreConfigurationGetDatastoreByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datastore configuration get datastore by name not found response has a 3xx status code
func (o *DatastoreConfigurationGetDatastoreByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration get datastore by name not found response has a 4xx status code
func (o *DatastoreConfigurationGetDatastoreByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this datastore configuration get datastore by name not found response has a 5xx status code
func (o *DatastoreConfigurationGetDatastoreByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this datastore configuration get datastore by name not found response a status code equal to that given
func (o *DatastoreConfigurationGetDatastoreByNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the datastore configuration get datastore by name not found response
func (o *DatastoreConfigurationGetDatastoreByNameNotFound) Code() int {
	return 404
}

func (o *DatastoreConfigurationGetDatastoreByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/name/{name}][%d] datastoreConfigurationGetDatastoreByNameNotFound  %+v", 404, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/datastores/name/{name}][%d] datastoreConfigurationGetDatastoreByNameNotFound  %+v", 404, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreByNameNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DatastoreConfigurationGetDatastoreByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationGetDatastoreByNameInternalServerError creates a DatastoreConfigurationGetDatastoreByNameInternalServerError with default headers values
func NewDatastoreConfigurationGetDatastoreByNameInternalServerError() *DatastoreConfigurationGetDatastoreByNameInternalServerError {
	return &DatastoreConfigurationGetDatastoreByNameInternalServerError{}
}

/*
DatastoreConfigurationGetDatastoreByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type DatastoreConfigurationGetDatastoreByNameInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this datastore configuration get datastore by name internal server error response has a 2xx status code
func (o *DatastoreConfigurationGetDatastoreByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datastore configuration get datastore by name internal server error response has a 3xx status code
func (o *DatastoreConfigurationGetDatastoreByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration get datastore by name internal server error response has a 4xx status code
func (o *DatastoreConfigurationGetDatastoreByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this datastore configuration get datastore by name internal server error response has a 5xx status code
func (o *DatastoreConfigurationGetDatastoreByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this datastore configuration get datastore by name internal server error response a status code equal to that given
func (o *DatastoreConfigurationGetDatastoreByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the datastore configuration get datastore by name internal server error response
func (o *DatastoreConfigurationGetDatastoreByNameInternalServerError) Code() int {
	return 500
}

func (o *DatastoreConfigurationGetDatastoreByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/name/{name}][%d] datastoreConfigurationGetDatastoreByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/datastores/name/{name}][%d] datastoreConfigurationGetDatastoreByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreByNameInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DatastoreConfigurationGetDatastoreByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationGetDatastoreByNameGatewayTimeout creates a DatastoreConfigurationGetDatastoreByNameGatewayTimeout with default headers values
func NewDatastoreConfigurationGetDatastoreByNameGatewayTimeout() *DatastoreConfigurationGetDatastoreByNameGatewayTimeout {
	return &DatastoreConfigurationGetDatastoreByNameGatewayTimeout{}
}

/*
DatastoreConfigurationGetDatastoreByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type DatastoreConfigurationGetDatastoreByNameGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this datastore configuration get datastore by name gateway timeout response has a 2xx status code
func (o *DatastoreConfigurationGetDatastoreByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datastore configuration get datastore by name gateway timeout response has a 3xx status code
func (o *DatastoreConfigurationGetDatastoreByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration get datastore by name gateway timeout response has a 4xx status code
func (o *DatastoreConfigurationGetDatastoreByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this datastore configuration get datastore by name gateway timeout response has a 5xx status code
func (o *DatastoreConfigurationGetDatastoreByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this datastore configuration get datastore by name gateway timeout response a status code equal to that given
func (o *DatastoreConfigurationGetDatastoreByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the datastore configuration get datastore by name gateway timeout response
func (o *DatastoreConfigurationGetDatastoreByNameGatewayTimeout) Code() int {
	return 504
}

func (o *DatastoreConfigurationGetDatastoreByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/name/{name}][%d] datastoreConfigurationGetDatastoreByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/datastores/name/{name}][%d] datastoreConfigurationGetDatastoreByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreByNameGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *DatastoreConfigurationGetDatastoreByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationGetDatastoreByNameDefault creates a DatastoreConfigurationGetDatastoreByNameDefault with default headers values
func NewDatastoreConfigurationGetDatastoreByNameDefault(code int) *DatastoreConfigurationGetDatastoreByNameDefault {
	return &DatastoreConfigurationGetDatastoreByNameDefault{
		_statusCode: code,
	}
}

/*
DatastoreConfigurationGetDatastoreByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DatastoreConfigurationGetDatastoreByNameDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this datastore configuration get datastore by name default response has a 2xx status code
func (o *DatastoreConfigurationGetDatastoreByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this datastore configuration get datastore by name default response has a 3xx status code
func (o *DatastoreConfigurationGetDatastoreByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this datastore configuration get datastore by name default response has a 4xx status code
func (o *DatastoreConfigurationGetDatastoreByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this datastore configuration get datastore by name default response has a 5xx status code
func (o *DatastoreConfigurationGetDatastoreByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this datastore configuration get datastore by name default response a status code equal to that given
func (o *DatastoreConfigurationGetDatastoreByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the datastore configuration get datastore by name default response
func (o *DatastoreConfigurationGetDatastoreByNameDefault) Code() int {
	return o._statusCode
}

func (o *DatastoreConfigurationGetDatastoreByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/datastores/name/{name}][%d] DatastoreConfiguration_GetDatastoreByName default  %+v", o._statusCode, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/datastores/name/{name}][%d] DatastoreConfiguration_GetDatastoreByName default  %+v", o._statusCode, o.Payload)
}

func (o *DatastoreConfigurationGetDatastoreByNameDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *DatastoreConfigurationGetDatastoreByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
