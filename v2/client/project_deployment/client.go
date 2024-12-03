package project_deployment

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
	CreateProjectDeployment(params *CreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAccepted, error)
	DeleteProjectDeploymentAll(params *DeleteAllParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAllOK, error)
	DeleteProjectDeployment(params *DeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOK, error)
	GetProjectDeploymentByID(params *GetByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetByIDOK, error)
	GetProjectDeploymentListByID(params *GetListbyIdParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetListbyIdOK, error)
	UpdateProjectDeployment(params *UpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOK, error)
	SetTransport(transport runtime.ClientTransport)
}

/*
CreateProjectDeployment creates resource group

v2 API to create a resource group record.
*/
func (a *Client) CreateProjectDeployment(params *CreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResourceGroup_CreateResourceGroupV2",
		Method:             "POST",
		PathPattern:        "/v2/projects",
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
DeleteProjectDeploymentAll deletes all deployments in a project

v2 API to delete all deployments in a project by project ID.
*/
func (a *Client) DeleteProjectDeploymentAll(params *DeleteAllParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAllOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAllParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResourceGroup_DeleteResourceGroupAllV2",
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
DeleteProjectDeployment deletes deployment in a project

v2 API to delete a specific deployment in a project by project ID and deployment ID.
*/
func (a *Client) DeleteProjectDeployment(params *DeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResourceGroup_DeleteResourceGroupV2",
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
GetProjectDeploymentByID gets specific deployment status in a project by project ID and deployment ID

v2 API to get project status by project ID
*/
func (a *Client) GetProjectDeploymentByID(params *GetByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResourceGroup_GetDeploymentByIdV2",
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
GetProjectDeploymentListByID gets deployment list in a project by ID

v2 API to get deployment list by project ID
*/
func (a *Client) GetProjectDeploymentListByID(params *GetListbyIdParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetListbyIdOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetListbyIdParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResourceGroup_GetDeploymentListbyIdv2",
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

/*
UpdateProjectDeployment updates deployment in a project

v2 API to update deployment in a project. The usual pattern to update deployment is to retrieve the record and update with the modified values in a new body to update the deployment record.
*/
func (a *Client) UpdateProjectDeployment(params *UpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResourceGroup_UpdateResourceGroupV2",
		Method:             "PUT",
		PathPattern:        "/v2/projects/id/{projectId}/deployments/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateReader{formats: a.formats},
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
	success, ok := result.(*UpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
