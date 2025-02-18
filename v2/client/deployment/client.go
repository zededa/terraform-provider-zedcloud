package deployment

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new resource group API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new resource group API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new resource group API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for resource group API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Create(params *CreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAccepted, error)
	DeleteAll(params *DeleteAllParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAllOK, error)
	Delete(params *DeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOK, error)
	GetByID(params *GetByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetByIDOK, error)
	GetListByProjectID(params *GetListbyIdParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetListbyIdOK, error)
	CreateNewVersion(params *CreateNewVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateNewVersionAccepted, error)
	SetTransport(transport runtime.ClientTransport)
}

/*
Create creates resource group

v2 API to create a resource group record.
*/
func (a *Client) Create(params *CreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateDeployment",
		Method:             "POST",
		PathPattern:        "/v2/projects/id/{projectId}/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateReader{formats: a.formats},
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
	success, ok := result.(*CreateAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteAll deletes all deployments in a project

v2 API to delete all deployments in a project by project ID.
*/
func (a *Client) DeleteAll(params *DeleteAllParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAllOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAllParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteDeploymentsAllV2",
		Method:             "DELETE",
		PathPattern:        "/v2/projects/id/{projectId}/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAllReader{formats: a.formats},
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
	success, ok := result.(*DeleteAllOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteAllDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
Delete deletes deployment in a project

v2 API to delete a specific deployment in a project by project ID and deployment ID.
*/
func (a *Client) Delete(params *DeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteDeploymentV2",
		Method:             "DELETE",
		PathPattern:        "/v2/projects/id/{projectId}/deployments/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteReader{formats: a.formats},
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
	success, ok := result.(*DeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetByID gets specific deployment status in a project by project ID and deployment ID

v2 API to get project status by project ID
*/
func (a *Client) GetByID(params *GetByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RGetDeploymentByIdV2",
		Method:             "GET",
		PathPattern:        "/v2/projects/id/{projectId}/deployments/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetByIDReader{formats: a.formats},
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
	success, ok := result.(*GetByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetListByProjectID gets deployment list in a project by ID

v2 API to get deployment list by project ID
*/
func (a *Client) GetListByProjectID(params *GetListbyIdParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetListbyIdOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetListbyIdParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDeploymentListbyProjectIdv2",
		Method:             "GET",
		PathPattern:        "/v2/projects/id/{projectId}/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetListbyIdReader{formats: a.formats},
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
	success, ok := result.(*GetListbyIdOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetListbyIdDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// CreateNewVersion creates a new version of deployment in a project
func (a *Client) CreateNewVersion(params *CreateNewVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateNewVersionAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNewVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateDeploymentGroupV2",
		Method:             "PUT",
		PathPattern:        "/v2/projects/id/{projectId}/deployments/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateNewVersionReader{formats: a.formats},
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
	success, ok := result.(*CreateNewVersionAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateNewVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
