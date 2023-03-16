package image

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// ImageConfigurationQueryImagesReader is a Reader for the ImageConfigurationQueryImages structure.
type ImageConfigurationQueryImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageConfigurationQueryImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageConfigurationQueryImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImageConfigurationQueryImagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImageConfigurationQueryImagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImageConfigurationQueryImagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageConfigurationQueryImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewImageConfigurationQueryImagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewImageConfigurationQueryImagesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImageConfigurationQueryImagesOK creates a ImageConfigurationQueryImagesOK with default headers values
func NewImageConfigurationQueryImagesOK() *ImageConfigurationQueryImagesOK {
	return &ImageConfigurationQueryImagesOK{}
}

/*
ImageConfigurationQueryImagesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ImageConfigurationQueryImagesOK struct {
	Payload *models.Images
}

// IsSuccess returns true when this image configuration query images o k response has a 2xx status code
func (o *ImageConfigurationQueryImagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image configuration query images o k response has a 3xx status code
func (o *ImageConfigurationQueryImagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration query images o k response has a 4xx status code
func (o *ImageConfigurationQueryImagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration query images o k response has a 5xx status code
func (o *ImageConfigurationQueryImagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration query images o k response a status code equal to that given
func (o *ImageConfigurationQueryImagesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the image configuration query images o k response
func (o *ImageConfigurationQueryImagesOK) Code() int {
	return 200
}

func (o *ImageConfigurationQueryImagesOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] imageConfigurationQueryImagesOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationQueryImagesOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] imageConfigurationQueryImagesOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationQueryImagesOK) GetPayload() *models.Images {
	return o.Payload
}

func (o *ImageConfigurationQueryImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Images)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryImagesBadRequest creates a ImageConfigurationQueryImagesBadRequest with default headers values
func NewImageConfigurationQueryImagesBadRequest() *ImageConfigurationQueryImagesBadRequest {
	return &ImageConfigurationQueryImagesBadRequest{}
}

/*
ImageConfigurationQueryImagesBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type ImageConfigurationQueryImagesBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration query images bad request response has a 2xx status code
func (o *ImageConfigurationQueryImagesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration query images bad request response has a 3xx status code
func (o *ImageConfigurationQueryImagesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration query images bad request response has a 4xx status code
func (o *ImageConfigurationQueryImagesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration query images bad request response has a 5xx status code
func (o *ImageConfigurationQueryImagesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration query images bad request response a status code equal to that given
func (o *ImageConfigurationQueryImagesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the image configuration query images bad request response
func (o *ImageConfigurationQueryImagesBadRequest) Code() int {
	return 400
}

func (o *ImageConfigurationQueryImagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] imageConfigurationQueryImagesBadRequest  %+v", 400, o.Payload)
}

func (o *ImageConfigurationQueryImagesBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] imageConfigurationQueryImagesBadRequest  %+v", 400, o.Payload)
}

func (o *ImageConfigurationQueryImagesBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationQueryImagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryImagesUnauthorized creates a ImageConfigurationQueryImagesUnauthorized with default headers values
func NewImageConfigurationQueryImagesUnauthorized() *ImageConfigurationQueryImagesUnauthorized {
	return &ImageConfigurationQueryImagesUnauthorized{}
}

/*
ImageConfigurationQueryImagesUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ImageConfigurationQueryImagesUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration query images unauthorized response has a 2xx status code
func (o *ImageConfigurationQueryImagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration query images unauthorized response has a 3xx status code
func (o *ImageConfigurationQueryImagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration query images unauthorized response has a 4xx status code
func (o *ImageConfigurationQueryImagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration query images unauthorized response has a 5xx status code
func (o *ImageConfigurationQueryImagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration query images unauthorized response a status code equal to that given
func (o *ImageConfigurationQueryImagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the image configuration query images unauthorized response
func (o *ImageConfigurationQueryImagesUnauthorized) Code() int {
	return 401
}

func (o *ImageConfigurationQueryImagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] imageConfigurationQueryImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationQueryImagesUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] imageConfigurationQueryImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationQueryImagesUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationQueryImagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryImagesForbidden creates a ImageConfigurationQueryImagesForbidden with default headers values
func NewImageConfigurationQueryImagesForbidden() *ImageConfigurationQueryImagesForbidden {
	return &ImageConfigurationQueryImagesForbidden{}
}

/*
ImageConfigurationQueryImagesForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ImageConfigurationQueryImagesForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration query images forbidden response has a 2xx status code
func (o *ImageConfigurationQueryImagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration query images forbidden response has a 3xx status code
func (o *ImageConfigurationQueryImagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration query images forbidden response has a 4xx status code
func (o *ImageConfigurationQueryImagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration query images forbidden response has a 5xx status code
func (o *ImageConfigurationQueryImagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration query images forbidden response a status code equal to that given
func (o *ImageConfigurationQueryImagesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the image configuration query images forbidden response
func (o *ImageConfigurationQueryImagesForbidden) Code() int {
	return 403
}

func (o *ImageConfigurationQueryImagesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] imageConfigurationQueryImagesForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationQueryImagesForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] imageConfigurationQueryImagesForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationQueryImagesForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationQueryImagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryImagesInternalServerError creates a ImageConfigurationQueryImagesInternalServerError with default headers values
func NewImageConfigurationQueryImagesInternalServerError() *ImageConfigurationQueryImagesInternalServerError {
	return &ImageConfigurationQueryImagesInternalServerError{}
}

/*
ImageConfigurationQueryImagesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ImageConfigurationQueryImagesInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration query images internal server error response has a 2xx status code
func (o *ImageConfigurationQueryImagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration query images internal server error response has a 3xx status code
func (o *ImageConfigurationQueryImagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration query images internal server error response has a 4xx status code
func (o *ImageConfigurationQueryImagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration query images internal server error response has a 5xx status code
func (o *ImageConfigurationQueryImagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration query images internal server error response a status code equal to that given
func (o *ImageConfigurationQueryImagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the image configuration query images internal server error response
func (o *ImageConfigurationQueryImagesInternalServerError) Code() int {
	return 500
}

func (o *ImageConfigurationQueryImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] imageConfigurationQueryImagesInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationQueryImagesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] imageConfigurationQueryImagesInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationQueryImagesInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationQueryImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryImagesGatewayTimeout creates a ImageConfigurationQueryImagesGatewayTimeout with default headers values
func NewImageConfigurationQueryImagesGatewayTimeout() *ImageConfigurationQueryImagesGatewayTimeout {
	return &ImageConfigurationQueryImagesGatewayTimeout{}
}

/*
ImageConfigurationQueryImagesGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ImageConfigurationQueryImagesGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration query images gateway timeout response has a 2xx status code
func (o *ImageConfigurationQueryImagesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration query images gateway timeout response has a 3xx status code
func (o *ImageConfigurationQueryImagesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration query images gateway timeout response has a 4xx status code
func (o *ImageConfigurationQueryImagesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration query images gateway timeout response has a 5xx status code
func (o *ImageConfigurationQueryImagesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration query images gateway timeout response a status code equal to that given
func (o *ImageConfigurationQueryImagesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the image configuration query images gateway timeout response
func (o *ImageConfigurationQueryImagesGatewayTimeout) Code() int {
	return 504
}

func (o *ImageConfigurationQueryImagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] imageConfigurationQueryImagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationQueryImagesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] imageConfigurationQueryImagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationQueryImagesGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationQueryImagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryImagesDefault creates a ImageConfigurationQueryImagesDefault with default headers values
func NewImageConfigurationQueryImagesDefault(code int) *ImageConfigurationQueryImagesDefault {
	return &ImageConfigurationQueryImagesDefault{
		_statusCode: code,
	}
}

/*
ImageConfigurationQueryImagesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ImageConfigurationQueryImagesDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this image configuration query images default response has a 2xx status code
func (o *ImageConfigurationQueryImagesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this image configuration query images default response has a 3xx status code
func (o *ImageConfigurationQueryImagesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this image configuration query images default response has a 4xx status code
func (o *ImageConfigurationQueryImagesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this image configuration query images default response has a 5xx status code
func (o *ImageConfigurationQueryImagesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this image configuration query images default response a status code equal to that given
func (o *ImageConfigurationQueryImagesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the image configuration query images default response
func (o *ImageConfigurationQueryImagesDefault) Code() int {
	return o._statusCode
}

func (o *ImageConfigurationQueryImagesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] ImageConfiguration_QueryImages default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationQueryImagesDefault) String() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] ImageConfiguration_QueryImages default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationQueryImagesDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ImageConfigurationQueryImagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
