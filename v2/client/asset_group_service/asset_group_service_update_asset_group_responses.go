// Code generated by go-swagger; DO NOT EDIT.

package asset_group_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// AssetGroupServiceUpdateAssetGroupReader is a Reader for the AssetGroupServiceUpdateAssetGroup structure.
type AssetGroupServiceUpdateAssetGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssetGroupServiceUpdateAssetGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssetGroupServiceUpdateAssetGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAssetGroupServiceUpdateAssetGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAssetGroupServiceUpdateAssetGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAssetGroupServiceUpdateAssetGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAssetGroupServiceUpdateAssetGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAssetGroupServiceUpdateAssetGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewAssetGroupServiceUpdateAssetGroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAssetGroupServiceUpdateAssetGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssetGroupServiceUpdateAssetGroupOK creates a AssetGroupServiceUpdateAssetGroupOK with default headers values
func NewAssetGroupServiceUpdateAssetGroupOK() *AssetGroupServiceUpdateAssetGroupOK {
	return &AssetGroupServiceUpdateAssetGroupOK{}
}

/*
AssetGroupServiceUpdateAssetGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type AssetGroupServiceUpdateAssetGroupOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this asset group service update asset group o k response has a 2xx status code
func (o *AssetGroupServiceUpdateAssetGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this asset group service update asset group o k response has a 3xx status code
func (o *AssetGroupServiceUpdateAssetGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this asset group service update asset group o k response has a 4xx status code
func (o *AssetGroupServiceUpdateAssetGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this asset group service update asset group o k response has a 5xx status code
func (o *AssetGroupServiceUpdateAssetGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this asset group service update asset group o k response a status code equal to that given
func (o *AssetGroupServiceUpdateAssetGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the asset group service update asset group o k response
func (o *AssetGroupServiceUpdateAssetGroupOK) Code() int {
	return 200
}

func (o *AssetGroupServiceUpdateAssetGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/assetgroups/id/{id}][%d] assetGroupServiceUpdateAssetGroupOK %s", 200, payload)
}

func (o *AssetGroupServiceUpdateAssetGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/assetgroups/id/{id}][%d] assetGroupServiceUpdateAssetGroupOK %s", 200, payload)
}

func (o *AssetGroupServiceUpdateAssetGroupOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AssetGroupServiceUpdateAssetGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetGroupServiceUpdateAssetGroupUnauthorized creates a AssetGroupServiceUpdateAssetGroupUnauthorized with default headers values
func NewAssetGroupServiceUpdateAssetGroupUnauthorized() *AssetGroupServiceUpdateAssetGroupUnauthorized {
	return &AssetGroupServiceUpdateAssetGroupUnauthorized{}
}

/*
AssetGroupServiceUpdateAssetGroupUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type AssetGroupServiceUpdateAssetGroupUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this asset group service update asset group unauthorized response has a 2xx status code
func (o *AssetGroupServiceUpdateAssetGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this asset group service update asset group unauthorized response has a 3xx status code
func (o *AssetGroupServiceUpdateAssetGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this asset group service update asset group unauthorized response has a 4xx status code
func (o *AssetGroupServiceUpdateAssetGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this asset group service update asset group unauthorized response has a 5xx status code
func (o *AssetGroupServiceUpdateAssetGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this asset group service update asset group unauthorized response a status code equal to that given
func (o *AssetGroupServiceUpdateAssetGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the asset group service update asset group unauthorized response
func (o *AssetGroupServiceUpdateAssetGroupUnauthorized) Code() int {
	return 401
}

func (o *AssetGroupServiceUpdateAssetGroupUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/assetgroups/id/{id}][%d] assetGroupServiceUpdateAssetGroupUnauthorized %s", 401, payload)
}

func (o *AssetGroupServiceUpdateAssetGroupUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/assetgroups/id/{id}][%d] assetGroupServiceUpdateAssetGroupUnauthorized %s", 401, payload)
}

func (o *AssetGroupServiceUpdateAssetGroupUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AssetGroupServiceUpdateAssetGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetGroupServiceUpdateAssetGroupForbidden creates a AssetGroupServiceUpdateAssetGroupForbidden with default headers values
func NewAssetGroupServiceUpdateAssetGroupForbidden() *AssetGroupServiceUpdateAssetGroupForbidden {
	return &AssetGroupServiceUpdateAssetGroupForbidden{}
}

/*
AssetGroupServiceUpdateAssetGroupForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type AssetGroupServiceUpdateAssetGroupForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this asset group service update asset group forbidden response has a 2xx status code
func (o *AssetGroupServiceUpdateAssetGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this asset group service update asset group forbidden response has a 3xx status code
func (o *AssetGroupServiceUpdateAssetGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this asset group service update asset group forbidden response has a 4xx status code
func (o *AssetGroupServiceUpdateAssetGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this asset group service update asset group forbidden response has a 5xx status code
func (o *AssetGroupServiceUpdateAssetGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this asset group service update asset group forbidden response a status code equal to that given
func (o *AssetGroupServiceUpdateAssetGroupForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the asset group service update asset group forbidden response
func (o *AssetGroupServiceUpdateAssetGroupForbidden) Code() int {
	return 403
}

func (o *AssetGroupServiceUpdateAssetGroupForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/assetgroups/id/{id}][%d] assetGroupServiceUpdateAssetGroupForbidden %s", 403, payload)
}

func (o *AssetGroupServiceUpdateAssetGroupForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/assetgroups/id/{id}][%d] assetGroupServiceUpdateAssetGroupForbidden %s", 403, payload)
}

func (o *AssetGroupServiceUpdateAssetGroupForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AssetGroupServiceUpdateAssetGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetGroupServiceUpdateAssetGroupNotFound creates a AssetGroupServiceUpdateAssetGroupNotFound with default headers values
func NewAssetGroupServiceUpdateAssetGroupNotFound() *AssetGroupServiceUpdateAssetGroupNotFound {
	return &AssetGroupServiceUpdateAssetGroupNotFound{}
}

/*
AssetGroupServiceUpdateAssetGroupNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type AssetGroupServiceUpdateAssetGroupNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this asset group service update asset group not found response has a 2xx status code
func (o *AssetGroupServiceUpdateAssetGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this asset group service update asset group not found response has a 3xx status code
func (o *AssetGroupServiceUpdateAssetGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this asset group service update asset group not found response has a 4xx status code
func (o *AssetGroupServiceUpdateAssetGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this asset group service update asset group not found response has a 5xx status code
func (o *AssetGroupServiceUpdateAssetGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this asset group service update asset group not found response a status code equal to that given
func (o *AssetGroupServiceUpdateAssetGroupNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the asset group service update asset group not found response
func (o *AssetGroupServiceUpdateAssetGroupNotFound) Code() int {
	return 404
}

func (o *AssetGroupServiceUpdateAssetGroupNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/assetgroups/id/{id}][%d] assetGroupServiceUpdateAssetGroupNotFound %s", 404, payload)
}

func (o *AssetGroupServiceUpdateAssetGroupNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/assetgroups/id/{id}][%d] assetGroupServiceUpdateAssetGroupNotFound %s", 404, payload)
}

func (o *AssetGroupServiceUpdateAssetGroupNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AssetGroupServiceUpdateAssetGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetGroupServiceUpdateAssetGroupConflict creates a AssetGroupServiceUpdateAssetGroupConflict with default headers values
func NewAssetGroupServiceUpdateAssetGroupConflict() *AssetGroupServiceUpdateAssetGroupConflict {
	return &AssetGroupServiceUpdateAssetGroupConflict{}
}

/*
AssetGroupServiceUpdateAssetGroupConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing asset group record.
*/
type AssetGroupServiceUpdateAssetGroupConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this asset group service update asset group conflict response has a 2xx status code
func (o *AssetGroupServiceUpdateAssetGroupConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this asset group service update asset group conflict response has a 3xx status code
func (o *AssetGroupServiceUpdateAssetGroupConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this asset group service update asset group conflict response has a 4xx status code
func (o *AssetGroupServiceUpdateAssetGroupConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this asset group service update asset group conflict response has a 5xx status code
func (o *AssetGroupServiceUpdateAssetGroupConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this asset group service update asset group conflict response a status code equal to that given
func (o *AssetGroupServiceUpdateAssetGroupConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the asset group service update asset group conflict response
func (o *AssetGroupServiceUpdateAssetGroupConflict) Code() int {
	return 409
}

func (o *AssetGroupServiceUpdateAssetGroupConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/assetgroups/id/{id}][%d] assetGroupServiceUpdateAssetGroupConflict %s", 409, payload)
}

func (o *AssetGroupServiceUpdateAssetGroupConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/assetgroups/id/{id}][%d] assetGroupServiceUpdateAssetGroupConflict %s", 409, payload)
}

func (o *AssetGroupServiceUpdateAssetGroupConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AssetGroupServiceUpdateAssetGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetGroupServiceUpdateAssetGroupInternalServerError creates a AssetGroupServiceUpdateAssetGroupInternalServerError with default headers values
func NewAssetGroupServiceUpdateAssetGroupInternalServerError() *AssetGroupServiceUpdateAssetGroupInternalServerError {
	return &AssetGroupServiceUpdateAssetGroupInternalServerError{}
}

/*
AssetGroupServiceUpdateAssetGroupInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type AssetGroupServiceUpdateAssetGroupInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this asset group service update asset group internal server error response has a 2xx status code
func (o *AssetGroupServiceUpdateAssetGroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this asset group service update asset group internal server error response has a 3xx status code
func (o *AssetGroupServiceUpdateAssetGroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this asset group service update asset group internal server error response has a 4xx status code
func (o *AssetGroupServiceUpdateAssetGroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this asset group service update asset group internal server error response has a 5xx status code
func (o *AssetGroupServiceUpdateAssetGroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this asset group service update asset group internal server error response a status code equal to that given
func (o *AssetGroupServiceUpdateAssetGroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the asset group service update asset group internal server error response
func (o *AssetGroupServiceUpdateAssetGroupInternalServerError) Code() int {
	return 500
}

func (o *AssetGroupServiceUpdateAssetGroupInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/assetgroups/id/{id}][%d] assetGroupServiceUpdateAssetGroupInternalServerError %s", 500, payload)
}

func (o *AssetGroupServiceUpdateAssetGroupInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/assetgroups/id/{id}][%d] assetGroupServiceUpdateAssetGroupInternalServerError %s", 500, payload)
}

func (o *AssetGroupServiceUpdateAssetGroupInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AssetGroupServiceUpdateAssetGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetGroupServiceUpdateAssetGroupGatewayTimeout creates a AssetGroupServiceUpdateAssetGroupGatewayTimeout with default headers values
func NewAssetGroupServiceUpdateAssetGroupGatewayTimeout() *AssetGroupServiceUpdateAssetGroupGatewayTimeout {
	return &AssetGroupServiceUpdateAssetGroupGatewayTimeout{}
}

/*
AssetGroupServiceUpdateAssetGroupGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type AssetGroupServiceUpdateAssetGroupGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this asset group service update asset group gateway timeout response has a 2xx status code
func (o *AssetGroupServiceUpdateAssetGroupGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this asset group service update asset group gateway timeout response has a 3xx status code
func (o *AssetGroupServiceUpdateAssetGroupGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this asset group service update asset group gateway timeout response has a 4xx status code
func (o *AssetGroupServiceUpdateAssetGroupGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this asset group service update asset group gateway timeout response has a 5xx status code
func (o *AssetGroupServiceUpdateAssetGroupGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this asset group service update asset group gateway timeout response a status code equal to that given
func (o *AssetGroupServiceUpdateAssetGroupGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the asset group service update asset group gateway timeout response
func (o *AssetGroupServiceUpdateAssetGroupGatewayTimeout) Code() int {
	return 504
}

func (o *AssetGroupServiceUpdateAssetGroupGatewayTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/assetgroups/id/{id}][%d] assetGroupServiceUpdateAssetGroupGatewayTimeout %s", 504, payload)
}

func (o *AssetGroupServiceUpdateAssetGroupGatewayTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/assetgroups/id/{id}][%d] assetGroupServiceUpdateAssetGroupGatewayTimeout %s", 504, payload)
}

func (o *AssetGroupServiceUpdateAssetGroupGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *AssetGroupServiceUpdateAssetGroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetGroupServiceUpdateAssetGroupDefault creates a AssetGroupServiceUpdateAssetGroupDefault with default headers values
func NewAssetGroupServiceUpdateAssetGroupDefault(code int) *AssetGroupServiceUpdateAssetGroupDefault {
	return &AssetGroupServiceUpdateAssetGroupDefault{
		_statusCode: code,
	}
}

/*
AssetGroupServiceUpdateAssetGroupDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AssetGroupServiceUpdateAssetGroupDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this asset group service update asset group default response has a 2xx status code
func (o *AssetGroupServiceUpdateAssetGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this asset group service update asset group default response has a 3xx status code
func (o *AssetGroupServiceUpdateAssetGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this asset group service update asset group default response has a 4xx status code
func (o *AssetGroupServiceUpdateAssetGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this asset group service update asset group default response has a 5xx status code
func (o *AssetGroupServiceUpdateAssetGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this asset group service update asset group default response a status code equal to that given
func (o *AssetGroupServiceUpdateAssetGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the asset group service update asset group default response
func (o *AssetGroupServiceUpdateAssetGroupDefault) Code() int {
	return o._statusCode
}

func (o *AssetGroupServiceUpdateAssetGroupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/assetgroups/id/{id}][%d] AssetGroupService_UpdateAssetGroup default %s", o._statusCode, payload)
}

func (o *AssetGroupServiceUpdateAssetGroupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/assetgroups/id/{id}][%d] AssetGroupService_UpdateAssetGroup default %s", o._statusCode, payload)
}

func (o *AssetGroupServiceUpdateAssetGroupDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *AssetGroupServiceUpdateAssetGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AssetGroupServiceUpdateAssetGroupBody Asset groups
//
// Assetgroups is a proto that should be used by user-agents to create/update asset groups to group various assets like devices
swagger:model AssetGroupServiceUpdateAssetGroupBody
*/
type AssetGroupServiceUpdateAssetGroupBody struct {

	// asset ids
	AssetIds *models.AssetIDs `json:"assetIds,omitempty"`

	// asset tags
	AssetTags *models.AssetTags `json:"assetTags,omitempty"`

	// Detailed description of the asset group.
	// Max Length: 256
	Description string `json:"description,omitempty"`

	// User defined name of the asset group, unique across the enterprise. Once asset group is created, name can’t be changed.
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Name *string `json:"name"`

	// project id
	ProjectID string `json:"projectId,omitempty"`

	// User defined title of the asset group. Title can be changed at any time.
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+
	Title string `json:"title,omitempty"`
}

// Validate validates this asset group service update asset group body
func (o *AssetGroupServiceUpdateAssetGroupBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAssetIds(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAssetTags(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
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

func (o *AssetGroupServiceUpdateAssetGroupBody) validateAssetIds(formats strfmt.Registry) error {
	if swag.IsZero(o.AssetIds) { // not required
		return nil
	}

	if o.AssetIds != nil {
		if err := o.AssetIds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "assetIds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "assetIds")
			}
			return err
		}
	}

	return nil
}

func (o *AssetGroupServiceUpdateAssetGroupBody) validateAssetTags(formats strfmt.Registry) error {
	if swag.IsZero(o.AssetTags) { // not required
		return nil
	}

	if o.AssetTags != nil {
		if err := o.AssetTags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "assetTags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "assetTags")
			}
			return err
		}
	}

	return nil
}

func (o *AssetGroupServiceUpdateAssetGroupBody) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(o.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("body"+"."+"description", "body", o.Description, 256); err != nil {
		return err
	}

	return nil
}

func (o *AssetGroupServiceUpdateAssetGroupBody) validateName(formats strfmt.Registry) error {

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

func (o *AssetGroupServiceUpdateAssetGroupBody) validateTitle(formats strfmt.Registry) error {
	if swag.IsZero(o.Title) { // not required
		return nil
	}

	if err := validate.MinLength("body"+"."+"title", "body", o.Title, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("body"+"."+"title", "body", o.Title, 256); err != nil {
		return err
	}

	if err := validate.Pattern("body"+"."+"title", "body", o.Title, `[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this asset group service update asset group body based on the context it is used
func (o *AssetGroupServiceUpdateAssetGroupBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAssetIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateAssetTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AssetGroupServiceUpdateAssetGroupBody) contextValidateAssetIds(ctx context.Context, formats strfmt.Registry) error {

	if o.AssetIds != nil {

		if swag.IsZero(o.AssetIds) { // not required
			return nil
		}

		if err := o.AssetIds.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "assetIds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "assetIds")
			}
			return err
		}
	}

	return nil
}

func (o *AssetGroupServiceUpdateAssetGroupBody) contextValidateAssetTags(ctx context.Context, formats strfmt.Registry) error {

	if o.AssetTags != nil {

		if swag.IsZero(o.AssetTags) { // not required
			return nil
		}

		if err := o.AssetTags.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "assetTags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "assetTags")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AssetGroupServiceUpdateAssetGroupBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AssetGroupServiceUpdateAssetGroupBody) UnmarshalBinary(b []byte) error {
	var res AssetGroupServiceUpdateAssetGroupBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
