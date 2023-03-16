package hardware_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
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
	HardwareModelCreateHardwareBrand(params *HardwareModelCreateHardwareBrandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelCreateHardwareBrandOK, error)

	HardwareModelCreateHardwareModel(params *HardwareModelCreateHardwareModelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelCreateHardwareModelOK, error)

	HardwareModelDeleteHardwareBrand(params *HardwareModelDeleteHardwareBrandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelDeleteHardwareBrandOK, error)

	HardwareModelDeleteHardwareModel(params *HardwareModelDeleteHardwareModelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelDeleteHardwareModelOK, error)

	HardwareModelGetDeviceStatusConfig(params *HardwareModelGetDeviceStatusConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetDeviceStatusConfigOK, error)

	HardwareModelGetDeviceStatusLocation(params *HardwareModelGetDeviceStatusLocationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetDeviceStatusLocationOK, error)

	HardwareModelGetDeviceTags(params *HardwareModelGetDeviceTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetDeviceTagsOK, error)

	HardwareModelGetGlobalHardwareBrand(params *HardwareModelGetGlobalHardwareBrandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetGlobalHardwareBrandOK, error)

	HardwareModelGetGlobalHardwareBrandByName(params *HardwareModelGetGlobalHardwareBrandByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetGlobalHardwareBrandByNameOK, error)

	HardwareModelGetGlobalHardwareModel(params *HardwareModelGetGlobalHardwareModelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetGlobalHardwareModelOK, error)

	HardwareModelGetGlobalHardwareModelByName(params *HardwareModelGetGlobalHardwareModelByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetGlobalHardwareModelByNameOK, error)

	HardwareModelGetHardwareBrand(params *HardwareModelGetHardwareBrandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetHardwareBrandOK, error)

	HardwareModelGetHardwareBrandByName(params *HardwareModelGetHardwareBrandByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetHardwareBrandByNameOK, error)

	HardwareModelGetHardwareModel(params *HardwareModelGetHardwareModelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetHardwareModelOK, error)

	HardwareModelGetHardwareModelByName(params *HardwareModelGetHardwareModelByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetHardwareModelByNameOK, error)

	HardwareModelQueryGlobalHardwareBrands(params *HardwareModelQueryGlobalHardwareBrandsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelQueryGlobalHardwareBrandsOK, error)

	HardwareModelQueryGlobalHardwareModels(params *HardwareModelQueryGlobalHardwareModelsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelQueryGlobalHardwareModelsOK, error)

	HardwareModelQueryHardwareBrands(params *HardwareModelQueryHardwareBrandsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelQueryHardwareBrandsOK, error)

	HardwareModelQueryHardwareModels(params *HardwareModelQueryHardwareModelsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelQueryHardwareModelsOK, error)

	HardwareModelUpdateHardwareBrand(params *HardwareModelUpdateHardwareBrandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelUpdateHardwareBrandOK, error)

	HardwareModelUpdateHardwareModel(params *HardwareModelUpdateHardwareModelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelUpdateHardwareModelOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
HardwareModelCreateHardwareBrand creates hardware brand

Create a hardware brand record.
*/
func (a *Client) HardwareModelCreateHardwareBrand(params *HardwareModelCreateHardwareBrandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelCreateHardwareBrandOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelCreateHardwareBrandParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_CreateHardwareBrand",
		Method:             "POST",
		PathPattern:        "/v1/brands",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelCreateHardwareBrandReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelCreateHardwareBrandOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelCreateHardwareBrandDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HardwareModelCreateHardwareModel creates hardware model

Create a hardware model record.
*/
func (a *Client) HardwareModelCreateHardwareModel(params *HardwareModelCreateHardwareModelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelCreateHardwareModelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelCreateHardwareModelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_CreateHardwareModel",
		Method:             "POST",
		PathPattern:        "/v1/sysmodels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelCreateHardwareModelReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelCreateHardwareModelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelCreateHardwareModelDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HardwareModelDeleteHardwareBrand deletes hardware brand

Delete a hardware brand record.
*/
func (a *Client) HardwareModelDeleteHardwareBrand(params *HardwareModelDeleteHardwareBrandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelDeleteHardwareBrandOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelDeleteHardwareBrandParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_DeleteHardwareBrand",
		Method:             "DELETE",
		PathPattern:        "/v1/brands/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelDeleteHardwareBrandReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelDeleteHardwareBrandOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelDeleteHardwareBrandDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HardwareModelDeleteHardwareModel deletes hardware model

Delete a hardware model record.
*/
func (a *Client) HardwareModelDeleteHardwareModel(params *HardwareModelDeleteHardwareModelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelDeleteHardwareModelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelDeleteHardwareModelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_DeleteHardwareModel",
		Method:             "DELETE",
		PathPattern:        "/v1/sysmodels/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelDeleteHardwareModelReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelDeleteHardwareModelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelDeleteHardwareModelDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HardwareModelGetDeviceStatusConfig queries device status and config

Device status config API is a composite API for device config and device status
*/
func (a *Client) HardwareModelGetDeviceStatusConfig(params *HardwareModelGetDeviceStatusConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetDeviceStatusConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelGetDeviceStatusConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_GetDeviceStatusConfig",
		Method:             "GET",
		PathPattern:        "/v1/devices/status-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelGetDeviceStatusConfigReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelGetDeviceStatusConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelGetDeviceStatusConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HardwareModelGetDeviceStatusLocation queries device status location

Query device status locations
*/
func (a *Client) HardwareModelGetDeviceStatusLocation(params *HardwareModelGetDeviceStatusLocationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetDeviceStatusLocationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelGetDeviceStatusLocationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_GetDeviceStatusLocation",
		Method:             "GET",
		PathPattern:        "/v1/devices/status/locations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelGetDeviceStatusLocationReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelGetDeviceStatusLocationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelGetDeviceStatusLocationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HardwareModelGetDeviceTags queries device object tag key values

Query device object tag key-values
*/
func (a *Client) HardwareModelGetDeviceTags(params *HardwareModelGetDeviceTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetDeviceTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelGetDeviceTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_GetDeviceTags",
		Method:             "GET",
		PathPattern:        "/v1/devices/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelGetDeviceTagsReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelGetDeviceTagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelGetDeviceTagsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HardwareModelGetGlobalHardwareBrand gets global hardware brand

Get the configuration (without security details) of a global hardware brand record.
*/
func (a *Client) HardwareModelGetGlobalHardwareBrand(params *HardwareModelGetGlobalHardwareBrandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetGlobalHardwareBrandOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelGetGlobalHardwareBrandParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_GetGlobalHardwareBrand",
		Method:             "GET",
		PathPattern:        "/v1/brands/global/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelGetGlobalHardwareBrandReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelGetGlobalHardwareBrandOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelGetGlobalHardwareBrandDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HardwareModelGetGlobalHardwareBrandByName gets global hardware brand

Get the configuration (without security details) of a global hardware brand record.
*/
func (a *Client) HardwareModelGetGlobalHardwareBrandByName(params *HardwareModelGetGlobalHardwareBrandByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetGlobalHardwareBrandByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelGetGlobalHardwareBrandByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_GetGlobalHardwareBrandByName",
		Method:             "GET",
		PathPattern:        "/v1/brands/global/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelGetGlobalHardwareBrandByNameReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelGetGlobalHardwareBrandByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelGetGlobalHardwareBrandByNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HardwareModelGetGlobalHardwareModel gets global hardware model

Get the configuration (without security details) of a global hardware model record.
*/
func (a *Client) HardwareModelGetGlobalHardwareModel(params *HardwareModelGetGlobalHardwareModelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetGlobalHardwareModelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelGetGlobalHardwareModelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_GetGlobalHardwareModel",
		Method:             "GET",
		PathPattern:        "/v1/sysmodels/global/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelGetGlobalHardwareModelReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelGetGlobalHardwareModelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelGetGlobalHardwareModelDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HardwareModelGetGlobalHardwareModelByName gets global hardware model

Get the configuration (without security details) of a global hardware model record.
*/
func (a *Client) HardwareModelGetGlobalHardwareModelByName(params *HardwareModelGetGlobalHardwareModelByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetGlobalHardwareModelByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelGetGlobalHardwareModelByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_GetGlobalHardwareModelByName",
		Method:             "GET",
		PathPattern:        "/v1/sysmodels/global/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelGetGlobalHardwareModelByNameReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelGetGlobalHardwareModelByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelGetGlobalHardwareModelByNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HardwareModelGetHardwareBrand gets hardware brand

Get the configuration (without security details) of a hardware brand record.
*/
func (a *Client) HardwareModelGetHardwareBrand(params *HardwareModelGetHardwareBrandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetHardwareBrandOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelGetHardwareBrandParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_GetHardwareBrand",
		Method:             "GET",
		PathPattern:        "/v1/brands/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelGetHardwareBrandReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelGetHardwareBrandOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelGetHardwareBrandDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HardwareModelGetHardwareBrandByName gets hardware brand

Get the configuration (without security details) of a hardware brand record.
*/
func (a *Client) HardwareModelGetHardwareBrandByName(params *HardwareModelGetHardwareBrandByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetHardwareBrandByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelGetHardwareBrandByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_GetHardwareBrandByName",
		Method:             "GET",
		PathPattern:        "/v1/brands/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelGetHardwareBrandByNameReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelGetHardwareBrandByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelGetHardwareBrandByNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HardwareModelGetHardwareModel gets hardware model

Get the configuration (without security details) of a hardware model record.
*/
func (a *Client) HardwareModelGetHardwareModel(params *HardwareModelGetHardwareModelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetHardwareModelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelGetHardwareModelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_GetHardwareModel",
		Method:             "GET",
		PathPattern:        "/v1/sysmodels/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelGetHardwareModelReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelGetHardwareModelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelGetHardwareModelDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HardwareModelGetHardwareModelByName gets hardware model

Get the configuration (without security details) of a hardware model record.
*/
func (a *Client) HardwareModelGetHardwareModelByName(params *HardwareModelGetHardwareModelByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelGetHardwareModelByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelGetHardwareModelByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_GetHardwareModelByName",
		Method:             "GET",
		PathPattern:        "/v1/sysmodels/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelGetHardwareModelByNameReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelGetHardwareModelByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelGetHardwareModelByNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HardwareModelQueryGlobalHardwareBrands queries global hardware brands

Query the global hardware brand records.
*/
func (a *Client) HardwareModelQueryGlobalHardwareBrands(params *HardwareModelQueryGlobalHardwareBrandsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelQueryGlobalHardwareBrandsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelQueryGlobalHardwareBrandsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_QueryGlobalHardwareBrands",
		Method:             "GET",
		PathPattern:        "/v1/brands/global",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelQueryGlobalHardwareBrandsReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelQueryGlobalHardwareBrandsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelQueryGlobalHardwareBrandsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HardwareModelQueryGlobalHardwareModels queries global hardware models

Query the global hardware model records.
*/
func (a *Client) HardwareModelQueryGlobalHardwareModels(params *HardwareModelQueryGlobalHardwareModelsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelQueryGlobalHardwareModelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelQueryGlobalHardwareModelsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_QueryGlobalHardwareModels",
		Method:             "GET",
		PathPattern:        "/v1/sysmodels/global",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelQueryGlobalHardwareModelsReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelQueryGlobalHardwareModelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelQueryGlobalHardwareModelsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HardwareModelQueryHardwareBrands queries hardware brands

Query the hardware brand records.
*/
func (a *Client) HardwareModelQueryHardwareBrands(params *HardwareModelQueryHardwareBrandsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelQueryHardwareBrandsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelQueryHardwareBrandsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_QueryHardwareBrands",
		Method:             "GET",
		PathPattern:        "/v1/brands",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelQueryHardwareBrandsReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelQueryHardwareBrandsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelQueryHardwareBrandsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HardwareModelQueryHardwareModels queries hardware models

Query the hardware model records.
*/
func (a *Client) HardwareModelQueryHardwareModels(params *HardwareModelQueryHardwareModelsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelQueryHardwareModelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelQueryHardwareModelsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_QueryHardwareModels",
		Method:             "GET",
		PathPattern:        "/v1/sysmodels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelQueryHardwareModelsReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelQueryHardwareModelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelQueryHardwareModelsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HardwareModelUpdateHardwareBrand updates hardware brand

Update a hardware brand. The usual pattern to update a hardware brand record is to retrieve the record and update with the modified values in a new body to update the hardware brand record.
*/
func (a *Client) HardwareModelUpdateHardwareBrand(params *HardwareModelUpdateHardwareBrandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelUpdateHardwareBrandOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelUpdateHardwareBrandParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_UpdateHardwareBrand",
		Method:             "PUT",
		PathPattern:        "/v1/brands/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelUpdateHardwareBrandReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelUpdateHardwareBrandOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelUpdateHardwareBrandDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HardwareModelUpdateHardwareModel updates hardware model

Update a hardware model. The usual pattern to update a hardware model record is to retrieve the record and update with the modified values in a new body to update the hardware model record.
*/
func (a *Client) HardwareModelUpdateHardwareModel(params *HardwareModelUpdateHardwareModelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HardwareModelUpdateHardwareModelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHardwareModelUpdateHardwareModelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HardwareModel_UpdateHardwareModel",
		Method:             "PUT",
		PathPattern:        "/v1/sysmodels/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HardwareModelUpdateHardwareModelReader{formats: a.formats},
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
	success, ok := result.(*HardwareModelUpdateHardwareModelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HardwareModelUpdateHardwareModelDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
