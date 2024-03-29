package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new edge network configuration API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for edge network configuration API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Create(params *EdgeNetworkConfigurationCreateEdgeNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNetworkConfigurationCreateEdgeNetworkOK, error)

	DeleteNetwork(params *EdgeNetworkConfigurationDeleteEdgeNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNetworkConfigurationDeleteEdgeNetworkOK, error)

	GetByID(params *EdgeNetworkConfigurationGetEdgeNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNetworkConfigurationGetEdgeNetworkOK, error)

	GetByName(params *EdgeNetworkConfigurationGetEdgeNetworkByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNetworkConfigurationGetEdgeNetworkByNameOK, error)

	Update(params *EdgeNetworkConfigurationUpdateEdgeNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateResult, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Create creates edge network

Create an edge network record.
*/
func (a *Client) Create(params *EdgeNetworkConfigurationCreateEdgeNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNetworkConfigurationCreateEdgeNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = CreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNetworkConfiguration_CreateEdgeNetwork",
		Method:             "POST",
		PathPattern:        "/v1/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNetworkConfigurationCreateEdgeNetworkReader{formats: a.formats},
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
	success, ok := result.(*EdgeNetworkConfigurationCreateEdgeNetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNetworkConfigurationCreateEdgeNetworkDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteNetwork deletes edge network

Delete an edge network record.
*/
func (a *Client) DeleteNetwork(params *EdgeNetworkConfigurationDeleteEdgeNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNetworkConfigurationDeleteEdgeNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = DeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNetworkConfiguration_DeleteEdgeNetwork",
		Method:             "DELETE",
		PathPattern:        "/v1/networks/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNetworkConfigurationDeleteEdgeNetworkReader{formats: a.formats},
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
	success, ok := result.(*EdgeNetworkConfigurationDeleteEdgeNetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNetworkConfigurationDeleteEdgeNetworkDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetByID gets edge network

Get the configuration (without security details) of an edge network record.
*/
func (a *Client) GetByID(params *EdgeNetworkConfigurationGetEdgeNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNetworkConfigurationGetEdgeNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = GetByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNetworkConfiguration_GetEdgeNetwork",
		Method:             "GET",
		PathPattern:        "/v1/networks/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNetworkConfigurationGetEdgeNetworkReader{formats: a.formats},
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
	success, ok := result.(*EdgeNetworkConfigurationGetEdgeNetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNetworkConfigurationGetEdgeNetworkDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetByName gets edge network

Get the configuration (without security details) of an edge network record.
*/
func (a *Client) GetByName(params *EdgeNetworkConfigurationGetEdgeNetworkByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNetworkConfigurationGetEdgeNetworkByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = GetByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNetworkConfiguration_GetEdgeNetworkByName",
		Method:             "GET",
		PathPattern:        "/v1/networks/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNetworkConfigurationGetEdgeNetworkByNameReader{formats: a.formats},
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
	success, ok := result.(*EdgeNetworkConfigurationGetEdgeNetworkByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNetworkConfigurationGetEdgeNetworkByNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EdgeNetworkConfigurationQueryEdgeNetworks queries edge networks

Query the edge network records.
*/
func (a *Client) EdgeNetworkConfigurationQueryEdgeNetworks(params *EdgeNetworkConfigurationQueryEdgeNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeNetworkConfigurationQueryEdgeNetworksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeNetworkConfigurationQueryEdgeNetworksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNetworkConfiguration_QueryEdgeNetworks",
		Method:             "GET",
		PathPattern:        "/v1/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNetworkConfigurationQueryEdgeNetworksReader{formats: a.formats},
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
	success, ok := result.(*EdgeNetworkConfigurationQueryEdgeNetworksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNetworkConfigurationQueryEdgeNetworksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
Update updates edge network

Update an edge network. The usual pattern to update an edge network record is to retrieve the record and update with the modified values in a new body to update the edge network record.
*/
func (a *Client) Update(params *EdgeNetworkConfigurationUpdateEdgeNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateResult, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = UpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeNetworkConfiguration_UpdateEdgeNetwork",
		Method:             "PUT",
		PathPattern:        "/v1/networks/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EdgeNetworkConfigurationUpdateEdgeNetworkReader{formats: a.formats},
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
	success, ok := result.(*UpdateResult)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeNetworkConfigurationUpdateEdgeNetworkDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
