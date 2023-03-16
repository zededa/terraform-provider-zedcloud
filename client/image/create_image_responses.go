package image

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// ImageConfigurationCreateImageReader is a Reader for the ImageConfigurationCreateImage structure.
type ImageConfigurationCreateImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageConfigurationCreateImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageConfigurationCreateImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImageConfigurationCreateImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImageConfigurationCreateImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImageConfigurationCreateImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewImageConfigurationCreateImageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageConfigurationCreateImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewImageConfigurationCreateImageGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewImageConfigurationCreateImageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImageConfigurationCreateImageOK creates a ImageConfigurationCreateImageOK with default headers values
func NewImageConfigurationCreateImageOK() *ImageConfigurationCreateImageOK {
	return &ImageConfigurationCreateImageOK{}
}

/*
ImageConfigurationCreateImageOK describes a response with status code 200, with default header values.

A successful response.
*/
type ImageConfigurationCreateImageOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration create image o k response has a 2xx status code
func (o *ImageConfigurationCreateImageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image configuration create image o k response has a 3xx status code
func (o *ImageConfigurationCreateImageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration create image o k response has a 4xx status code
func (o *ImageConfigurationCreateImageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration create image o k response has a 5xx status code
func (o *ImageConfigurationCreateImageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration create image o k response a status code equal to that given
func (o *ImageConfigurationCreateImageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the image configuration create image o k response
func (o *ImageConfigurationCreateImageOK) Code() int {
	return 200
}

func (o *ImageConfigurationCreateImageOK) Error() string {
	return fmt.Sprintf("[POST /v1/apps/images][%d] imageConfigurationCreateImageOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationCreateImageOK) String() string {
	return fmt.Sprintf("[POST /v1/apps/images][%d] imageConfigurationCreateImageOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationCreateImageOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationCreateImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationCreateImageBadRequest creates a ImageConfigurationCreateImageBadRequest with default headers values
func NewImageConfigurationCreateImageBadRequest() *ImageConfigurationCreateImageBadRequest {
	return &ImageConfigurationCreateImageBadRequest{}
}

/*
ImageConfigurationCreateImageBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type ImageConfigurationCreateImageBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration create image bad request response has a 2xx status code
func (o *ImageConfigurationCreateImageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration create image bad request response has a 3xx status code
func (o *ImageConfigurationCreateImageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration create image bad request response has a 4xx status code
func (o *ImageConfigurationCreateImageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration create image bad request response has a 5xx status code
func (o *ImageConfigurationCreateImageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration create image bad request response a status code equal to that given
func (o *ImageConfigurationCreateImageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the image configuration create image bad request response
func (o *ImageConfigurationCreateImageBadRequest) Code() int {
	return 400
}

func (o *ImageConfigurationCreateImageBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/apps/images][%d] imageConfigurationCreateImageBadRequest  %+v", 400, o.Payload)
}

func (o *ImageConfigurationCreateImageBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/apps/images][%d] imageConfigurationCreateImageBadRequest  %+v", 400, o.Payload)
}

func (o *ImageConfigurationCreateImageBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationCreateImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationCreateImageUnauthorized creates a ImageConfigurationCreateImageUnauthorized with default headers values
func NewImageConfigurationCreateImageUnauthorized() *ImageConfigurationCreateImageUnauthorized {
	return &ImageConfigurationCreateImageUnauthorized{}
}

/*
ImageConfigurationCreateImageUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ImageConfigurationCreateImageUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration create image unauthorized response has a 2xx status code
func (o *ImageConfigurationCreateImageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration create image unauthorized response has a 3xx status code
func (o *ImageConfigurationCreateImageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration create image unauthorized response has a 4xx status code
func (o *ImageConfigurationCreateImageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration create image unauthorized response has a 5xx status code
func (o *ImageConfigurationCreateImageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration create image unauthorized response a status code equal to that given
func (o *ImageConfigurationCreateImageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the image configuration create image unauthorized response
func (o *ImageConfigurationCreateImageUnauthorized) Code() int {
	return 401
}

func (o *ImageConfigurationCreateImageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/apps/images][%d] imageConfigurationCreateImageUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationCreateImageUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/apps/images][%d] imageConfigurationCreateImageUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationCreateImageUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationCreateImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationCreateImageForbidden creates a ImageConfigurationCreateImageForbidden with default headers values
func NewImageConfigurationCreateImageForbidden() *ImageConfigurationCreateImageForbidden {
	return &ImageConfigurationCreateImageForbidden{}
}

/*
ImageConfigurationCreateImageForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ImageConfigurationCreateImageForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration create image forbidden response has a 2xx status code
func (o *ImageConfigurationCreateImageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration create image forbidden response has a 3xx status code
func (o *ImageConfigurationCreateImageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration create image forbidden response has a 4xx status code
func (o *ImageConfigurationCreateImageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration create image forbidden response has a 5xx status code
func (o *ImageConfigurationCreateImageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration create image forbidden response a status code equal to that given
func (o *ImageConfigurationCreateImageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the image configuration create image forbidden response
func (o *ImageConfigurationCreateImageForbidden) Code() int {
	return 403
}

func (o *ImageConfigurationCreateImageForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/apps/images][%d] imageConfigurationCreateImageForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationCreateImageForbidden) String() string {
	return fmt.Sprintf("[POST /v1/apps/images][%d] imageConfigurationCreateImageForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationCreateImageForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationCreateImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationCreateImageConflict creates a ImageConfigurationCreateImageConflict with default headers values
func NewImageConfigurationCreateImageConflict() *ImageConfigurationCreateImageConflict {
	return &ImageConfigurationCreateImageConflict{}
}

/*
ImageConfigurationCreateImageConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this edge application image record will conflict with an already existing edge application image record.
*/
type ImageConfigurationCreateImageConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration create image conflict response has a 2xx status code
func (o *ImageConfigurationCreateImageConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration create image conflict response has a 3xx status code
func (o *ImageConfigurationCreateImageConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration create image conflict response has a 4xx status code
func (o *ImageConfigurationCreateImageConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration create image conflict response has a 5xx status code
func (o *ImageConfigurationCreateImageConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration create image conflict response a status code equal to that given
func (o *ImageConfigurationCreateImageConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the image configuration create image conflict response
func (o *ImageConfigurationCreateImageConflict) Code() int {
	return 409
}

func (o *ImageConfigurationCreateImageConflict) Error() string {
	return fmt.Sprintf("[POST /v1/apps/images][%d] imageConfigurationCreateImageConflict  %+v", 409, o.Payload)
}

func (o *ImageConfigurationCreateImageConflict) String() string {
	return fmt.Sprintf("[POST /v1/apps/images][%d] imageConfigurationCreateImageConflict  %+v", 409, o.Payload)
}

func (o *ImageConfigurationCreateImageConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationCreateImageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationCreateImageInternalServerError creates a ImageConfigurationCreateImageInternalServerError with default headers values
func NewImageConfigurationCreateImageInternalServerError() *ImageConfigurationCreateImageInternalServerError {
	return &ImageConfigurationCreateImageInternalServerError{}
}

/*
ImageConfigurationCreateImageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ImageConfigurationCreateImageInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration create image internal server error response has a 2xx status code
func (o *ImageConfigurationCreateImageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration create image internal server error response has a 3xx status code
func (o *ImageConfigurationCreateImageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration create image internal server error response has a 4xx status code
func (o *ImageConfigurationCreateImageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration create image internal server error response has a 5xx status code
func (o *ImageConfigurationCreateImageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration create image internal server error response a status code equal to that given
func (o *ImageConfigurationCreateImageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the image configuration create image internal server error response
func (o *ImageConfigurationCreateImageInternalServerError) Code() int {
	return 500
}

func (o *ImageConfigurationCreateImageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/apps/images][%d] imageConfigurationCreateImageInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationCreateImageInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/apps/images][%d] imageConfigurationCreateImageInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationCreateImageInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationCreateImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationCreateImageGatewayTimeout creates a ImageConfigurationCreateImageGatewayTimeout with default headers values
func NewImageConfigurationCreateImageGatewayTimeout() *ImageConfigurationCreateImageGatewayTimeout {
	return &ImageConfigurationCreateImageGatewayTimeout{}
}

/*
ImageConfigurationCreateImageGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ImageConfigurationCreateImageGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration create image gateway timeout response has a 2xx status code
func (o *ImageConfigurationCreateImageGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration create image gateway timeout response has a 3xx status code
func (o *ImageConfigurationCreateImageGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration create image gateway timeout response has a 4xx status code
func (o *ImageConfigurationCreateImageGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration create image gateway timeout response has a 5xx status code
func (o *ImageConfigurationCreateImageGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration create image gateway timeout response a status code equal to that given
func (o *ImageConfigurationCreateImageGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the image configuration create image gateway timeout response
func (o *ImageConfigurationCreateImageGatewayTimeout) Code() int {
	return 504
}

func (o *ImageConfigurationCreateImageGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/apps/images][%d] imageConfigurationCreateImageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationCreateImageGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/apps/images][%d] imageConfigurationCreateImageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationCreateImageGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationCreateImageGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationCreateImageDefault creates a ImageConfigurationCreateImageDefault with default headers values
func NewImageConfigurationCreateImageDefault(code int) *ImageConfigurationCreateImageDefault {
	return &ImageConfigurationCreateImageDefault{
		_statusCode: code,
	}
}

/*
ImageConfigurationCreateImageDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ImageConfigurationCreateImageDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this image configuration create image default response has a 2xx status code
func (o *ImageConfigurationCreateImageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this image configuration create image default response has a 3xx status code
func (o *ImageConfigurationCreateImageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this image configuration create image default response has a 4xx status code
func (o *ImageConfigurationCreateImageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this image configuration create image default response has a 5xx status code
func (o *ImageConfigurationCreateImageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this image configuration create image default response a status code equal to that given
func (o *ImageConfigurationCreateImageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the image configuration create image default response
func (o *ImageConfigurationCreateImageDefault) Code() int {
	return o._statusCode
}

func (o *ImageConfigurationCreateImageDefault) Error() string {
	return fmt.Sprintf("[POST /v1/apps/images][%d] ImageConfiguration_CreateImage default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationCreateImageDefault) String() string {
	return fmt.Sprintf("[POST /v1/apps/images][%d] ImageConfiguration_CreateImage default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationCreateImageDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ImageConfigurationCreateImageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
