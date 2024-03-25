// Code generated by go-swagger; DO NOT EDIT.

package patch_envelope

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/zededa/terraform-provider-zedcloud/models"
)

// PatchEnvelopeConfigurationUpdatePatchEnvelopeReader is a Reader for the PatchEnvelopeConfigurationUpdatePatchEnvelope structure.
type PatchEnvelopeConfigurationUpdatePatchEnvelopeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchEnvelopeConfigurationUpdatePatchEnvelopeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchEnvelopeConfigurationUpdatePatchEnvelopeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchEnvelopeConfigurationUpdatePatchEnvelopeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchEnvelopeConfigurationUpdatePatchEnvelopeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchEnvelopeConfigurationUpdatePatchEnvelopeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchEnvelopeConfigurationUpdatePatchEnvelopeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchEnvelopeConfigurationUpdatePatchEnvelopeOK creates a PatchEnvelopeConfigurationUpdatePatchEnvelopeOK with default headers values
func NewPatchEnvelopeConfigurationUpdatePatchEnvelopeOK() *PatchEnvelopeConfigurationUpdatePatchEnvelopeOK {
	return &PatchEnvelopeConfigurationUpdatePatchEnvelopeOK{}
}

/*
PatchEnvelopeConfigurationUpdatePatchEnvelopeOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchEnvelopeConfigurationUpdatePatchEnvelopeOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration update patch envelope o k response has a 2xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch envelope configuration update patch envelope o k response has a 3xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration update patch envelope o k response has a 4xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch envelope configuration update patch envelope o k response has a 5xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch envelope configuration update patch envelope o k response a status code equal to that given
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch envelope configuration update patch envelope o k response
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeOK) Code() int {
	return 200
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeOK) Error() string {
	return fmt.Sprintf("[PUT /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationUpdatePatchEnvelopeOK  %+v", 200, o.Payload)
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeOK) String() string {
	return fmt.Sprintf("[PUT /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationUpdatePatchEnvelopeOK  %+v", 200, o.Payload)
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationUpdatePatchEnvelopeUnauthorized creates a PatchEnvelopeConfigurationUpdatePatchEnvelopeUnauthorized with default headers values
func NewPatchEnvelopeConfigurationUpdatePatchEnvelopeUnauthorized() *PatchEnvelopeConfigurationUpdatePatchEnvelopeUnauthorized {
	return &PatchEnvelopeConfigurationUpdatePatchEnvelopeUnauthorized{}
}

/*
PatchEnvelopeConfigurationUpdatePatchEnvelopeUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type PatchEnvelopeConfigurationUpdatePatchEnvelopeUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration update patch envelope unauthorized response has a 2xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration update patch envelope unauthorized response has a 3xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration update patch envelope unauthorized response has a 4xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch envelope configuration update patch envelope unauthorized response has a 5xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch envelope configuration update patch envelope unauthorized response a status code equal to that given
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the patch envelope configuration update patch envelope unauthorized response
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeUnauthorized) Code() int {
	return 401
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationUpdatePatchEnvelopeUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationUpdatePatchEnvelopeUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationUpdatePatchEnvelopeForbidden creates a PatchEnvelopeConfigurationUpdatePatchEnvelopeForbidden with default headers values
func NewPatchEnvelopeConfigurationUpdatePatchEnvelopeForbidden() *PatchEnvelopeConfigurationUpdatePatchEnvelopeForbidden {
	return &PatchEnvelopeConfigurationUpdatePatchEnvelopeForbidden{}
}

/*
PatchEnvelopeConfigurationUpdatePatchEnvelopeForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type PatchEnvelopeConfigurationUpdatePatchEnvelopeForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration update patch envelope forbidden response has a 2xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration update patch envelope forbidden response has a 3xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration update patch envelope forbidden response has a 4xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch envelope configuration update patch envelope forbidden response has a 5xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch envelope configuration update patch envelope forbidden response a status code equal to that given
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the patch envelope configuration update patch envelope forbidden response
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeForbidden) Code() int {
	return 403
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationUpdatePatchEnvelopeForbidden  %+v", 403, o.Payload)
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationUpdatePatchEnvelopeForbidden  %+v", 403, o.Payload)
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationUpdatePatchEnvelopeInternalServerError creates a PatchEnvelopeConfigurationUpdatePatchEnvelopeInternalServerError with default headers values
func NewPatchEnvelopeConfigurationUpdatePatchEnvelopeInternalServerError() *PatchEnvelopeConfigurationUpdatePatchEnvelopeInternalServerError {
	return &PatchEnvelopeConfigurationUpdatePatchEnvelopeInternalServerError{}
}

/*
PatchEnvelopeConfigurationUpdatePatchEnvelopeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type PatchEnvelopeConfigurationUpdatePatchEnvelopeInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration update patch envelope internal server error response has a 2xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration update patch envelope internal server error response has a 3xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration update patch envelope internal server error response has a 4xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch envelope configuration update patch envelope internal server error response has a 5xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch envelope configuration update patch envelope internal server error response a status code equal to that given
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the patch envelope configuration update patch envelope internal server error response
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeInternalServerError) Code() int {
	return 500
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationUpdatePatchEnvelopeInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationUpdatePatchEnvelopeInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationUpdatePatchEnvelopeGatewayTimeout creates a PatchEnvelopeConfigurationUpdatePatchEnvelopeGatewayTimeout with default headers values
func NewPatchEnvelopeConfigurationUpdatePatchEnvelopeGatewayTimeout() *PatchEnvelopeConfigurationUpdatePatchEnvelopeGatewayTimeout {
	return &PatchEnvelopeConfigurationUpdatePatchEnvelopeGatewayTimeout{}
}

/*
PatchEnvelopeConfigurationUpdatePatchEnvelopeGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type PatchEnvelopeConfigurationUpdatePatchEnvelopeGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration update patch envelope gateway timeout response has a 2xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration update patch envelope gateway timeout response has a 3xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration update patch envelope gateway timeout response has a 4xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch envelope configuration update patch envelope gateway timeout response has a 5xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch envelope configuration update patch envelope gateway timeout response a status code equal to that given
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the patch envelope configuration update patch envelope gateway timeout response
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeGatewayTimeout) Code() int {
	return 504
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationUpdatePatchEnvelopeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationUpdatePatchEnvelopeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationUpdatePatchEnvelopeDefault creates a PatchEnvelopeConfigurationUpdatePatchEnvelopeDefault with default headers values
func NewPatchEnvelopeConfigurationUpdatePatchEnvelopeDefault(code int) *PatchEnvelopeConfigurationUpdatePatchEnvelopeDefault {
	return &PatchEnvelopeConfigurationUpdatePatchEnvelopeDefault{
		_statusCode: code,
	}
}

/*
PatchEnvelopeConfigurationUpdatePatchEnvelopeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PatchEnvelopeConfigurationUpdatePatchEnvelopeDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this patch envelope configuration update patch envelope default response has a 2xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch envelope configuration update patch envelope default response has a 3xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch envelope configuration update patch envelope default response has a 4xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch envelope configuration update patch envelope default response has a 5xx status code
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch envelope configuration update patch envelope default response a status code equal to that given
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch envelope configuration update patch envelope default response
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeDefault) Code() int {
	return o._statusCode
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/patch-envelope/id/{id}][%d] PatchEnvelopeConfiguration_UpdatePatchEnvelope default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeDefault) String() string {
	return fmt.Sprintf("[PUT /v1/patch-envelope/id/{id}][%d] PatchEnvelopeConfiguration_UpdatePatchEnvelope default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PatchEnvelopeConfigurationUpdatePatchEnvelopeBody Opaque data for application instances
//
// PatchEnvelope is a proto that should be used by user-agents to create/update opaque data for application instances.
swagger:model PatchEnvelopeConfigurationUpdatePatchEnvelopeBody
*/
type PatchEnvelopeConfigurationUpdatePatchEnvelopeBody struct {

	// if activate is false, device should only download
	// images and not present them to app instance
	//
	// Flag to represent whether device needs to present it to app instance
	// Required: true
	Action *models.PatchEnvelopeAction `json:"action"`

	// Patch envelope artifacts
	// Required: true
	Artifacts []*models.BinaryArtifact `json:"artifacts"`

	// Detailed description of the patch envelope.
	// Max Length: 256
	Description string `json:"description,omitempty"`

	// number of devices referencing this patch envelope
	DeviceCount int64 `json:"deviceCount,omitempty"`

	// User defined name of the patch envelope, unique across the enterprise. Once patch envelope is created, name can’t be changed.
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Name *string `json:"name"`

	// project id
	ProjectID string `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// system defined info
	// Read Only: true
	Revision *models.ObjectRevision `json:"revision,omitempty"`

	// User defined title of the patch envelope. Title can be changed at any time.
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+
	Title *string `json:"title"`

	// optional - arbitrary version string if any by user
	//
	// User defined version for the given patch envelope
	UserDefinedVersion string `json:"userDefinedVersion,omitempty"`
}

// Validate validates this patch envelope configuration update patch envelope body
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateArtifacts(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
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

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeBody) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"action", "body", o.Action); err != nil {
		return err
	}

	if err := validate.Required("body"+"."+"action", "body", o.Action); err != nil {
		return err
	}

	if o.Action != nil {
		if err := o.Action.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "action")
			}
			return err
		}
	}

	return nil
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeBody) validateArtifacts(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"artifacts", "body", o.Artifacts); err != nil {
		return err
	}

	for i := 0; i < len(o.Artifacts); i++ {
		if swag.IsZero(o.Artifacts[i]) { // not required
			continue
		}

		if o.Artifacts[i] != nil {
			if err := o.Artifacts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "artifacts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "artifacts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeBody) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(o.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("body"+"."+"description", "body", o.Description, 256); err != nil {
		return err
	}

	return nil
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	if err := validate.MinLength("body"+"."+"name", "body", *o.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("body"+"."+"name", "body", *o.Name, 256); err != nil {
		return err
	}

	if err := validate.Pattern("body"+"."+"name", "body", *o.Name, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeBody) validateRevision(formats strfmt.Registry) error {
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

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeBody) validateTitle(formats strfmt.Registry) error {

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

// ContextValidate validate this patch envelope configuration update patch envelope body based on the context it is used
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateArtifacts(ctx, formats); err != nil {
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

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeBody) contextValidateAction(ctx context.Context, formats strfmt.Registry) error {

	if o.Action != nil {

		if err := o.Action.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "action")
			}
			return err
		}
	}

	return nil
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeBody) contextValidateArtifacts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Artifacts); i++ {

		if o.Artifacts[i] != nil {

			if swag.IsZero(o.Artifacts[i]) { // not required
				return nil
			}

			if err := o.Artifacts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "artifacts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "artifacts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeBody) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

	if o.Revision != nil {

		if swag.IsZero(o.Revision) { // not required
			return nil
		}

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
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeBody) UnmarshalBinary(b []byte) error {
	var res PatchEnvelopeConfigurationUpdatePatchEnvelopeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
