package image

import (
	"context"
	"fmt"
	"io"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/zededa/terraform-provider-zedcloud/models"
)

// ImageConfigurationUplinkImageReader is a Reader for the ImageConfigurationUplinkImage structure.
type ImageConfigurationUplinkImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageConfigurationUplinkImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageConfigurationUplinkImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewImageConfigurationUplinkImageAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImageConfigurationUplinkImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImageConfigurationUplinkImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImageConfigurationUplinkImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImageConfigurationUplinkImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewImageConfigurationUplinkImageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageConfigurationUplinkImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewImageConfigurationUplinkImageGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewImageConfigurationUplinkImageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImageConfigurationUplinkImageOK creates a ImageConfigurationUplinkImageOK with default headers values
func NewImageConfigurationUplinkImageOK() *ImageConfigurationUplinkImageOK {
	return &ImageConfigurationUplinkImageOK{}
}

/*
ImageConfigurationUplinkImageOK describes a response with status code 200, with default header values.

A successful response.
*/
type ImageConfigurationUplinkImageOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration uplink image o k response has a 2xx status code
func (o *ImageConfigurationUplinkImageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image configuration uplink image o k response has a 3xx status code
func (o *ImageConfigurationUplinkImageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration uplink image o k response has a 4xx status code
func (o *ImageConfigurationUplinkImageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration uplink image o k response has a 5xx status code
func (o *ImageConfigurationUplinkImageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration uplink image o k response a status code equal to that given
func (o *ImageConfigurationUplinkImageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the image configuration uplink image o k response
func (o *ImageConfigurationUplinkImageOK) Code() int {
	return 200
}

func (o *ImageConfigurationUplinkImageOK) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] imageConfigurationUplinkImageOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationUplinkImageOK) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] imageConfigurationUplinkImageOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationUplinkImageOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUplinkImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUplinkImageAccepted creates a ImageConfigurationUplinkImageAccepted with default headers values
func NewImageConfigurationUplinkImageAccepted() *ImageConfigurationUplinkImageAccepted {
	return &ImageConfigurationUplinkImageAccepted{}
}

/*
ImageConfigurationUplinkImageAccepted describes a response with status code 202, with default header values.

Accepted. The API gateway accepted the request for uplinking but the uplinking process has not been completed. Please check ImageStatus and ImageError fields to track the status of uplinking process and any error messages.
*/
type ImageConfigurationUplinkImageAccepted struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration uplink image accepted response has a 2xx status code
func (o *ImageConfigurationUplinkImageAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image configuration uplink image accepted response has a 3xx status code
func (o *ImageConfigurationUplinkImageAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration uplink image accepted response has a 4xx status code
func (o *ImageConfigurationUplinkImageAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration uplink image accepted response has a 5xx status code
func (o *ImageConfigurationUplinkImageAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration uplink image accepted response a status code equal to that given
func (o *ImageConfigurationUplinkImageAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the image configuration uplink image accepted response
func (o *ImageConfigurationUplinkImageAccepted) Code() int {
	return 202
}

func (o *ImageConfigurationUplinkImageAccepted) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] imageConfigurationUplinkImageAccepted  %+v", 202, o.Payload)
}

func (o *ImageConfigurationUplinkImageAccepted) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] imageConfigurationUplinkImageAccepted  %+v", 202, o.Payload)
}

func (o *ImageConfigurationUplinkImageAccepted) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUplinkImageAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUplinkImageBadRequest creates a ImageConfigurationUplinkImageBadRequest with default headers values
func NewImageConfigurationUplinkImageBadRequest() *ImageConfigurationUplinkImageBadRequest {
	return &ImageConfigurationUplinkImageBadRequest{}
}

/*
ImageConfigurationUplinkImageBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type ImageConfigurationUplinkImageBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration uplink image bad request response has a 2xx status code
func (o *ImageConfigurationUplinkImageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration uplink image bad request response has a 3xx status code
func (o *ImageConfigurationUplinkImageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration uplink image bad request response has a 4xx status code
func (o *ImageConfigurationUplinkImageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration uplink image bad request response has a 5xx status code
func (o *ImageConfigurationUplinkImageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration uplink image bad request response a status code equal to that given
func (o *ImageConfigurationUplinkImageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the image configuration uplink image bad request response
func (o *ImageConfigurationUplinkImageBadRequest) Code() int {
	return 400
}

func (o *ImageConfigurationUplinkImageBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] imageConfigurationUplinkImageBadRequest  %+v", 400, o.Payload)
}

func (o *ImageConfigurationUplinkImageBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] imageConfigurationUplinkImageBadRequest  %+v", 400, o.Payload)
}

func (o *ImageConfigurationUplinkImageBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUplinkImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUplinkImageUnauthorized creates a ImageConfigurationUplinkImageUnauthorized with default headers values
func NewImageConfigurationUplinkImageUnauthorized() *ImageConfigurationUplinkImageUnauthorized {
	return &ImageConfigurationUplinkImageUnauthorized{}
}

/*
ImageConfigurationUplinkImageUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ImageConfigurationUplinkImageUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration uplink image unauthorized response has a 2xx status code
func (o *ImageConfigurationUplinkImageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration uplink image unauthorized response has a 3xx status code
func (o *ImageConfigurationUplinkImageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration uplink image unauthorized response has a 4xx status code
func (o *ImageConfigurationUplinkImageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration uplink image unauthorized response has a 5xx status code
func (o *ImageConfigurationUplinkImageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration uplink image unauthorized response a status code equal to that given
func (o *ImageConfigurationUplinkImageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the image configuration uplink image unauthorized response
func (o *ImageConfigurationUplinkImageUnauthorized) Code() int {
	return 401
}

func (o *ImageConfigurationUplinkImageUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] imageConfigurationUplinkImageUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationUplinkImageUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] imageConfigurationUplinkImageUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationUplinkImageUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUplinkImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUplinkImageForbidden creates a ImageConfigurationUplinkImageForbidden with default headers values
func NewImageConfigurationUplinkImageForbidden() *ImageConfigurationUplinkImageForbidden {
	return &ImageConfigurationUplinkImageForbidden{}
}

/*
ImageConfigurationUplinkImageForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ImageConfigurationUplinkImageForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration uplink image forbidden response has a 2xx status code
func (o *ImageConfigurationUplinkImageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration uplink image forbidden response has a 3xx status code
func (o *ImageConfigurationUplinkImageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration uplink image forbidden response has a 4xx status code
func (o *ImageConfigurationUplinkImageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration uplink image forbidden response has a 5xx status code
func (o *ImageConfigurationUplinkImageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration uplink image forbidden response a status code equal to that given
func (o *ImageConfigurationUplinkImageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the image configuration uplink image forbidden response
func (o *ImageConfigurationUplinkImageForbidden) Code() int {
	return 403
}

func (o *ImageConfigurationUplinkImageForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] imageConfigurationUplinkImageForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationUplinkImageForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] imageConfigurationUplinkImageForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationUplinkImageForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUplinkImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUplinkImageNotFound creates a ImageConfigurationUplinkImageNotFound with default headers values
func NewImageConfigurationUplinkImageNotFound() *ImageConfigurationUplinkImageNotFound {
	return &ImageConfigurationUplinkImageNotFound{}
}

/*
ImageConfigurationUplinkImageNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type ImageConfigurationUplinkImageNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration uplink image not found response has a 2xx status code
func (o *ImageConfigurationUplinkImageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration uplink image not found response has a 3xx status code
func (o *ImageConfigurationUplinkImageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration uplink image not found response has a 4xx status code
func (o *ImageConfigurationUplinkImageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration uplink image not found response has a 5xx status code
func (o *ImageConfigurationUplinkImageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration uplink image not found response a status code equal to that given
func (o *ImageConfigurationUplinkImageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the image configuration uplink image not found response
func (o *ImageConfigurationUplinkImageNotFound) Code() int {
	return 404
}

func (o *ImageConfigurationUplinkImageNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] Image Uplink: Not Found  %+v", 404, spew.Sdump(o.Payload))
}

func (o *ImageConfigurationUplinkImageNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] Image Uplink: Not Found  %+v", 404, spew.Sdump(o.Payload))
}

func (o *ImageConfigurationUplinkImageNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUplinkImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUplinkImageConflict creates a ImageConfigurationUplinkImageConflict with default headers values
func NewImageConfigurationUplinkImageConflict() *ImageConfigurationUplinkImageConflict {
	return &ImageConfigurationUplinkImageConflict{}
}

/*
ImageConfigurationUplinkImageConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because another request for uplink / upload is already in progress
*/
type ImageConfigurationUplinkImageConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration uplink image conflict response has a 2xx status code
func (o *ImageConfigurationUplinkImageConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration uplink image conflict response has a 3xx status code
func (o *ImageConfigurationUplinkImageConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration uplink image conflict response has a 4xx status code
func (o *ImageConfigurationUplinkImageConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration uplink image conflict response has a 5xx status code
func (o *ImageConfigurationUplinkImageConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration uplink image conflict response a status code equal to that given
func (o *ImageConfigurationUplinkImageConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the image configuration uplink image conflict response
func (o *ImageConfigurationUplinkImageConflict) Code() int {
	return 409
}

func (o *ImageConfigurationUplinkImageConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] imageConfigurationUplinkImageConflict  %+v", 409, o.Payload)
}

func (o *ImageConfigurationUplinkImageConflict) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] imageConfigurationUplinkImageConflict  %+v", 409, o.Payload)
}

func (o *ImageConfigurationUplinkImageConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUplinkImageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUplinkImageInternalServerError creates a ImageConfigurationUplinkImageInternalServerError with default headers values
func NewImageConfigurationUplinkImageInternalServerError() *ImageConfigurationUplinkImageInternalServerError {
	return &ImageConfigurationUplinkImageInternalServerError{}
}

/*
ImageConfigurationUplinkImageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ImageConfigurationUplinkImageInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration uplink image internal server error response has a 2xx status code
func (o *ImageConfigurationUplinkImageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration uplink image internal server error response has a 3xx status code
func (o *ImageConfigurationUplinkImageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration uplink image internal server error response has a 4xx status code
func (o *ImageConfigurationUplinkImageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration uplink image internal server error response has a 5xx status code
func (o *ImageConfigurationUplinkImageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration uplink image internal server error response a status code equal to that given
func (o *ImageConfigurationUplinkImageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the image configuration uplink image internal server error response
func (o *ImageConfigurationUplinkImageInternalServerError) Code() int {
	return 500
}

func (o *ImageConfigurationUplinkImageInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] imageConfigurationUplinkImageInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationUplinkImageInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] imageConfigurationUplinkImageInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationUplinkImageInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUplinkImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUplinkImageGatewayTimeout creates a ImageConfigurationUplinkImageGatewayTimeout with default headers values
func NewImageConfigurationUplinkImageGatewayTimeout() *ImageConfigurationUplinkImageGatewayTimeout {
	return &ImageConfigurationUplinkImageGatewayTimeout{}
}

/*
ImageConfigurationUplinkImageGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ImageConfigurationUplinkImageGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration uplink image gateway timeout response has a 2xx status code
func (o *ImageConfigurationUplinkImageGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration uplink image gateway timeout response has a 3xx status code
func (o *ImageConfigurationUplinkImageGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration uplink image gateway timeout response has a 4xx status code
func (o *ImageConfigurationUplinkImageGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration uplink image gateway timeout response has a 5xx status code
func (o *ImageConfigurationUplinkImageGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration uplink image gateway timeout response a status code equal to that given
func (o *ImageConfigurationUplinkImageGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the image configuration uplink image gateway timeout response
func (o *ImageConfigurationUplinkImageGatewayTimeout) Code() int {
	return 504
}

func (o *ImageConfigurationUplinkImageGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] imageConfigurationUplinkImageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationUplinkImageGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] imageConfigurationUplinkImageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationUplinkImageGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUplinkImageGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUplinkImageDefault creates a ImageConfigurationUplinkImageDefault with default headers values
func NewImageConfigurationUplinkImageDefault(code int) *ImageConfigurationUplinkImageDefault {
	return &ImageConfigurationUplinkImageDefault{
		_statusCode: code,
	}
}

/*
ImageConfigurationUplinkImageDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ImageConfigurationUplinkImageDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this image configuration uplink image default response has a 2xx status code
func (o *ImageConfigurationUplinkImageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this image configuration uplink image default response has a 3xx status code
func (o *ImageConfigurationUplinkImageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this image configuration uplink image default response has a 4xx status code
func (o *ImageConfigurationUplinkImageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this image configuration uplink image default response has a 5xx status code
func (o *ImageConfigurationUplinkImageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this image configuration uplink image default response a status code equal to that given
func (o *ImageConfigurationUplinkImageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the image configuration uplink image default response
func (o *ImageConfigurationUplinkImageDefault) Code() int {
	return o._statusCode
}

func (o *ImageConfigurationUplinkImageDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] ImageConfiguration_UplinkImage default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationUplinkImageDefault) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/uplink][%d] ImageConfiguration_UplinkImage default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationUplinkImageDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ImageConfigurationUplinkImageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ImageConfigurationUplinkImageBody Image metadata detail
//
// Image metadata for edge gateway Base OS or for eedge applications.
// Example: {"description":"My test image in QCOW2 format for Edge computing","dsId":"7927f6e3-484d-4105-a98e-868b21c1cb61","id":"d1125b0f-633d-459c-99c6-f47e7467cebc","imageArch":"AMD64","imageError":"Image uplinked successfully...","imageFormat":3,"imageLocal":"","imageRelUrl":"edge/computing/AMD64","imageSha256":"ADC5BB9DD39F83DD74C276B0BA119FB27925A5CDEA343FE1F2C8433F28AB345B","imageSizeBytes":142016512,"imageStatus":4,"imageType":2,"imageVersion":"","name":"my-test-image","originType":2,"revision":{"createdAt":{"seconds":1592068270},"createdBy":"admin@my-company.com","curr":"1","updatedAt":{"seconds":1592068271},"updatedBy":"admin@my-company.com"},"title":"My Test Image for Edge Computing"}
swagger:model ImageConfigurationUplinkImageBody
*/
type ImageConfigurationUplinkImageBody struct {

	// Datastore Id where image binary is located.
	// Required: true
	// Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}
	DatastoreID *string `json:"datastoreId"`

	// Detailed description of the image.
	// Max Length: 256
	Description string `json:"description,omitempty"`

	// System defined universally unique Id of the image.
	// Read Only: true
	// Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}
	ID string `json:"id,omitempty"`

	// Image Architecture.
	// Required: true
	ImageArch *models.ModelArchType `json:"imageArch"`

	// Image upload/uplink detailed error/status message
	// Read Only: true
	ImageError string `json:"imageError,omitempty"`

	// Image binary format.
	// Required: true
	ImageFormat *models.ConfigFormat `json:"imageFormat"`

	// Internal image location.
	// Read Only: true
	ImageLocal string `json:"imageLocal,omitempty"`

	// Image relative path w.r.t. Datastore
	ImageRelURL string `json:"imageRelUrl,omitempty"`

	// Image checksum in SHA256 format
	ImageSha256 string `json:"imageSha256,omitempty"`

	// Image size in KBytes.
	ImageSizeBytes string `json:"imageSizeBytes,omitempty"`

	// Image status
	// Read Only: true
	ImageStatus *models.ImageStatus `json:"imageStatus,omitempty"`

	// Image type
	// Required: true
	ImageType *models.ImageType `json:"imageType"`

	// system defined info
	ImageVersion string `json:"imageVersion,omitempty"`

	// Origin type of image.
	// Read Only: true
	OriginType *models.Origin `json:"originType,omitempty"`

	// project access list of the image
	ProjectAccessList []string `json:"projectAccessList"`

	// system defined info
	// Read Only: true
	Revision *models.ObjectRevision `json:"revision,omitempty"`

	// User defined title of the image. Title can be changed at any time.
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+
	Title *string `json:"title"`
}

// Validate validates this image configuration uplink image body
func (o *ImageConfigurationUplinkImageBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDatastoreID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateImageArch(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateImageFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateImageStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateImageType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOriginType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ImageConfigurationUplinkImageBody) validateDatastoreID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"datastoreId", "body", o.DatastoreID); err != nil {
		return err
	}

	if err := validate.Pattern("body"+"."+"datastoreId", "body", *o.DatastoreID, `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`); err != nil {
		return err
	}

	return nil
}

func (o *ImageConfigurationUplinkImageBody) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(o.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("body"+"."+"description", "body", o.Description, 256); err != nil {
		return err
	}

	return nil
}

func (o *ImageConfigurationUplinkImageBody) validateID(formats strfmt.Registry) error {
	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"id", "body", o.ID, `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`); err != nil {
		return err
	}

	return nil
}

func (o *ImageConfigurationUplinkImageBody) validateImageArch(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"imageArch", "body", o.ImageArch); err != nil {
		return err
	}

	if err := validate.Required("body"+"."+"imageArch", "body", o.ImageArch); err != nil {
		return err
	}

	if o.ImageArch != nil {
		if err := o.ImageArch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "imageArch")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "imageArch")
			}
			return err
		}
	}

	return nil
}

func (o *ImageConfigurationUplinkImageBody) validateImageFormat(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"imageFormat", "body", o.ImageFormat); err != nil {
		return err
	}

	if err := validate.Required("body"+"."+"imageFormat", "body", o.ImageFormat); err != nil {
		return err
	}

	if o.ImageFormat != nil {
		if err := o.ImageFormat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "imageFormat")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "imageFormat")
			}
			return err
		}
	}

	return nil
}

func (o *ImageConfigurationUplinkImageBody) validateImageStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.ImageStatus) { // not required
		return nil
	}

	if o.ImageStatus != nil {
		if err := o.ImageStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "imageStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "imageStatus")
			}
			return err
		}
	}

	return nil
}

func (o *ImageConfigurationUplinkImageBody) validateImageType(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"imageType", "body", o.ImageType); err != nil {
		return err
	}

	if err := validate.Required("body"+"."+"imageType", "body", o.ImageType); err != nil {
		return err
	}

	if o.ImageType != nil {
		if err := o.ImageType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "imageType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "imageType")
			}
			return err
		}
	}

	return nil
}

func (o *ImageConfigurationUplinkImageBody) validateOriginType(formats strfmt.Registry) error {
	if swag.IsZero(o.OriginType) { // not required
		return nil
	}

	if o.OriginType != nil {
		if err := o.OriginType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "originType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "originType")
			}
			return err
		}
	}

	return nil
}

func (o *ImageConfigurationUplinkImageBody) validateRevision(formats strfmt.Registry) error {
	if swag.IsZero(o.Revision) { // not required
		return nil
	}

	if o.Revision != nil {
		if err := o.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "revision")
			}
			return err
		}
	}

	return nil
}

func (o *ImageConfigurationUplinkImageBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"title", "body", o.Title); err != nil {
		return err
	}

	if err := validate.MinLength("body"+"."+"title", "body", *o.Title, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("body"+"."+"title", "body", *o.Title, 256); err != nil {
		return err
	}

	if err := validate.Pattern("body"+"."+"title", "body", *o.Title, `[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this image configuration uplink image body based on the context it is used
func (o *ImageConfigurationUplinkImageBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateImageArch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateImageError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateImageFormat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateImageLocal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateImageStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateImageType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOriginType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ImageConfigurationUplinkImageBody) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "body"+"."+"id", "body", string(o.ID)); err != nil {
		return err
	}

	return nil
}

func (o *ImageConfigurationUplinkImageBody) contextValidateImageArch(ctx context.Context, formats strfmt.Registry) error {

	if o.ImageArch != nil {
		if err := o.ImageArch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "imageArch")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "imageArch")
			}
			return err
		}
	}

	return nil
}

func (o *ImageConfigurationUplinkImageBody) contextValidateImageError(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "body"+"."+"imageError", "body", string(o.ImageError)); err != nil {
		return err
	}

	return nil
}

func (o *ImageConfigurationUplinkImageBody) contextValidateImageFormat(ctx context.Context, formats strfmt.Registry) error {

	if o.ImageFormat != nil {
		if err := o.ImageFormat.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "imageFormat")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "imageFormat")
			}
			return err
		}
	}

	return nil
}

func (o *ImageConfigurationUplinkImageBody) contextValidateImageLocal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "body"+"."+"imageLocal", "body", string(o.ImageLocal)); err != nil {
		return err
	}

	return nil
}

func (o *ImageConfigurationUplinkImageBody) contextValidateImageStatus(ctx context.Context, formats strfmt.Registry) error {

	if o.ImageStatus != nil {
		if err := o.ImageStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "imageStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "imageStatus")
			}
			return err
		}
	}

	return nil
}

func (o *ImageConfigurationUplinkImageBody) contextValidateImageType(ctx context.Context, formats strfmt.Registry) error {

	if o.ImageType != nil {
		if err := o.ImageType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "imageType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "imageType")
			}
			return err
		}
	}

	return nil
}

func (o *ImageConfigurationUplinkImageBody) contextValidateOriginType(ctx context.Context, formats strfmt.Registry) error {

	if o.OriginType != nil {
		if err := o.OriginType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "originType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "originType")
			}
			return err
		}
	}

	return nil
}

func (o *ImageConfigurationUplinkImageBody) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

	if o.Revision != nil {
		if err := o.Revision.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "revision")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ImageConfigurationUplinkImageBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageConfigurationUplinkImageBody) UnmarshalBinary(b []byte) error {
	var res ImageConfigurationUplinkImageBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
