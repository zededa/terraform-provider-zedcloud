// Code generated by go-swagger; DO NOT EDIT.

package hardware_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new hardware model API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for hardware model API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateHardwareBrand(params *CreateHardwareBrandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateHardwareBrandOK, error)

	CreateHardwareModel(params *CreateHardwareModelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateHardwareModelOK, error)

	DeleteHardwareBrand(params *DeleteHardwareBrandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteHardwareBrandOK, error)

	DeleteHardwareModel(params *DeleteHardwareModelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteHardwareModelOK, error)

	GetDeviceTags(params *GetDeviceTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceTagsOK, error)

	GetGlobalHardwareBrand(params *GetGlobalHardwareBrandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGlobalHardwareBrandOK, error)

	GetGlobalHardwareBrandByName(params *GetGlobalHardwareBrandByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGlobalHardwareBrandByNameOK, error)

	GetGlobalHardwareModel(params *GetGlobalHardwareModelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGlobalHardwareModelOK, error)

	GetGlobalHardwareModelByName(params *GetGlobalHardwareModelByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGlobalHardwareModelByNameOK, error)

	GetHardwareBrand(params *GetHardwareBrandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHardwareBrandOK, error)

	GetHardwareBrandByName(params *GetHardwareBrandByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHardwareBrandByNameOK, error)

	GetHardwareModel(params *GetHardwareModelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHardwareModelOK, error)

	GetHardwareModelByName(params *GetHardwareModelByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHardwareModelByNameOK, error)

	QueryGlobalHardwareBrands(params *QueryGlobalHardwareBrandsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryGlobalHardwareBrandsOK, error)

	QueryGlobalHardwareModels(params *QueryGlobalHardwareModelsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryGlobalHardwareModelsOK, error)

	QueryHardwareBrands(params *QueryHardwareBrandsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryHardwareBrandsOK, error)

	QueryHardwareModels(params *QueryHardwareModelsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryHardwareModelsOK, error)

	UpdateHardwareBrand(params *UpdateHardwareBrandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateHardwareBrandOK, error)

	UpdateHardwareModel(params *UpdateHardwareModelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateHardwareModelOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateHardwareBrand creates hardware brand

Create a hardware brand record.
*/
func (a *Client) CreateHardwareBrand(params *CreateHardwareBrandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateHardwareBrandOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateHardwareBrandParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateHardwareBrand",
		Method:             "POST",
		PathPattern:        "/v1/brands",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateHardwareBrandReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateHardwareBrandOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateHardwareBrand: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateHardwareModel creates hardware model

Create a hardware model record.
*/
func (a *Client) CreateHardwareModel(params *CreateHardwareModelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateHardwareModelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateHardwareModelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateHardwareModel",
		Method:             "POST",
		PathPattern:        "/v1/sysmodels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateHardwareModelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateHardwareModelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateHardwareModel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteHardwareBrand deletes hardware brand

Delete a hardware brand record.
*/
func (a *Client) DeleteHardwareBrand(params *DeleteHardwareBrandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteHardwareBrandOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteHardwareBrandParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteHardwareBrand",
		Method:             "DELETE",
		PathPattern:        "/v1/brands/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteHardwareBrandReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteHardwareBrandOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteHardwareBrand: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteHardwareModel deletes hardware model

Delete a hardware model record.
*/
func (a *Client) DeleteHardwareModel(params *DeleteHardwareModelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteHardwareModelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteHardwareModelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteHardwareModel",
		Method:             "DELETE",
		PathPattern:        "/v1/sysmodels/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteHardwareModelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteHardwareModelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteHardwareModel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDeviceTags queries device object tag key values

Query device object tag key-values
*/
func (a *Client) GetDeviceTags(params *GetDeviceTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDeviceTags",
		Method:             "GET",
		PathPattern:        "/v1/devices/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceTagsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceTagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDeviceTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetGlobalHardwareBrand gets global hardware brand

Get the configuration (without security details) of a global hardware brand record.
*/
func (a *Client) GetGlobalHardwareBrand(params *GetGlobalHardwareBrandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGlobalHardwareBrandOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGlobalHardwareBrandParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetGlobalHardwareBrand",
		Method:             "GET",
		PathPattern:        "/v1/brands/global/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGlobalHardwareBrandReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGlobalHardwareBrandOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetGlobalHardwareBrand: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetGlobalHardwareBrandByName gets global hardware brand

Get the configuration (without security details) of a global hardware brand record.
*/
func (a *Client) GetGlobalHardwareBrandByName(params *GetGlobalHardwareBrandByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGlobalHardwareBrandByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGlobalHardwareBrandByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetGlobalHardwareBrandByName",
		Method:             "GET",
		PathPattern:        "/v1/brands/global/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGlobalHardwareBrandByNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGlobalHardwareBrandByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetGlobalHardwareBrandByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetGlobalHardwareModel gets global hardware model

Get the configuration (without security details) of a global hardware model record.
*/
func (a *Client) GetGlobalHardwareModel(params *GetGlobalHardwareModelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGlobalHardwareModelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGlobalHardwareModelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetGlobalHardwareModel",
		Method:             "GET",
		PathPattern:        "/v1/sysmodels/global/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGlobalHardwareModelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGlobalHardwareModelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetGlobalHardwareModel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetGlobalHardwareModelByName gets global hardware model

Get the configuration (without security details) of a global hardware model record.
*/
func (a *Client) GetGlobalHardwareModelByName(params *GetGlobalHardwareModelByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGlobalHardwareModelByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGlobalHardwareModelByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetGlobalHardwareModelByName",
		Method:             "GET",
		PathPattern:        "/v1/sysmodels/global/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGlobalHardwareModelByNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGlobalHardwareModelByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetGlobalHardwareModelByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetHardwareBrand gets hardware brand

Get the configuration (without security details) of a hardware brand record.
*/
func (a *Client) GetHardwareBrand(params *GetHardwareBrandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHardwareBrandOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHardwareBrandParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetHardwareBrand",
		Method:             "GET",
		PathPattern:        "/v1/brands/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHardwareBrandReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHardwareBrandOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHardwareBrand: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetHardwareBrandByName gets hardware brand

Get the configuration (without security details) of a hardware brand record.
*/
func (a *Client) GetHardwareBrandByName(params *GetHardwareBrandByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHardwareBrandByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHardwareBrandByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetHardwareBrandByName",
		Method:             "GET",
		PathPattern:        "/v1/brands/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHardwareBrandByNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHardwareBrandByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHardwareBrandByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetHardwareModel gets hardware model

Get the configuration (without security details) of a hardware model record.
*/
func (a *Client) GetHardwareModel(params *GetHardwareModelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHardwareModelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHardwareModelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetHardwareModel",
		Method:             "GET",
		PathPattern:        "/v1/sysmodels/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHardwareModelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHardwareModelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHardwareModel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetHardwareModelByName gets hardware model

Get the configuration (without security details) of a hardware model record.
*/
func (a *Client) GetHardwareModelByName(params *GetHardwareModelByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHardwareModelByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHardwareModelByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetHardwareModelByName",
		Method:             "GET",
		PathPattern:        "/v1/sysmodels/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHardwareModelByNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHardwareModelByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHardwareModelByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryGlobalHardwareBrands queries global hardware brands

Query the global hardware brand records.
*/
func (a *Client) QueryGlobalHardwareBrands(params *QueryGlobalHardwareBrandsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryGlobalHardwareBrandsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryGlobalHardwareBrandsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryGlobalHardwareBrands",
		Method:             "GET",
		PathPattern:        "/v1/brands/global",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryGlobalHardwareBrandsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryGlobalHardwareBrandsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryGlobalHardwareBrands: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryGlobalHardwareModels queries global hardware models

Query the global hardware model records.
*/
func (a *Client) QueryGlobalHardwareModels(params *QueryGlobalHardwareModelsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryGlobalHardwareModelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryGlobalHardwareModelsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryGlobalHardwareModels",
		Method:             "GET",
		PathPattern:        "/v1/sysmodels/global",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryGlobalHardwareModelsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryGlobalHardwareModelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryGlobalHardwareModels: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryHardwareBrands queries hardware brands

Query the hardware brand records.
*/
func (a *Client) QueryHardwareBrands(params *QueryHardwareBrandsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryHardwareBrandsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryHardwareBrandsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryHardwareBrands",
		Method:             "GET",
		PathPattern:        "/v1/brands",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryHardwareBrandsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryHardwareBrandsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryHardwareBrands: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryHardwareModels queries hardware models

Query the hardware model records.
*/
func (a *Client) QueryHardwareModels(params *QueryHardwareModelsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryHardwareModelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryHardwareModelsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryHardwareModels",
		Method:             "GET",
		PathPattern:        "/v1/sysmodels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryHardwareModelsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryHardwareModelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryHardwareModels: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateHardwareBrand updates hardware brand

Update a hardware brand. The usual pattern to update a hardware brand record is to retrieve the record and update with the modified values in a new body to update the hardware brand record.
*/
func (a *Client) UpdateHardwareBrand(params *UpdateHardwareBrandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateHardwareBrandOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateHardwareBrandParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateHardwareBrand",
		Method:             "PUT",
		PathPattern:        "/v1/brands/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateHardwareBrandReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateHardwareBrandOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateHardwareBrand: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateHardwareModel updates hardware model

Update a hardware model. The usual pattern to update a hardware model record is to retrieve the record and update with the modified values in a new body to update the hardware model record.
*/
func (a *Client) UpdateHardwareModel(params *UpdateHardwareModelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateHardwareModelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateHardwareModelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateHardwareModel",
		Method:             "PUT",
		PathPattern:        "/v1/sysmodels/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateHardwareModelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateHardwareModelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateHardwareModel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
