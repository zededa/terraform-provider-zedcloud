### Inconsistencies in the ZedCloudAPI
- This is not a REST API. It seems that partial updates of objects are not implemented via PATCH or at least idempotent PUT requests like in REST specification.
- Seems to be half RESTish/declarative half RPC/imperative. Weird design.
- Not cachable
- Not stateless
- Not a uniform interface
- PUT requests for partial updates without body but header params (RPC like flow)
- Partial updates don't return the updated object
- Logic for conditional updating based on state, is not done in the backend but required to be done by clients. This is not only error prone and requires redundant code in client, missing validation is also a security issue.
- Parameters are ignored (eve_os_image). To provide functionality for updating it seems that a RESTish endpoint is provided updating the config and a RPC-like  endpoint (or imperative PUT endpoint, not sure how to call such a construct) to apply the update is provided to trigger the update.
- Two different concepts are represented in a single field (adminstate). The dependency of the field's values doesn't justify violating separation of conerns. It's a wrong API design that imposes state handling on consumers.
- Endpoint naming is strange
- Several endpoints have same description
- Naming in general is inconsistent, doesn't follow any standards
- Spec files have many typos and copy paste errors (don't people use a spell-checker? How can this pass review?)

#### Implications of inconsistency and wrong API design
- We cannot generate tests for the API
- We cannot use the OpenAPI generated client
- We cannot generate the Terraform provider
- Code generation in general will be very difficult
- A chain of workarounds forms up that has severe impact on maintainability, stability and extensibility
- Software rot and spaghetti code (no wonder the servers are named like this)

#### Issues
- Imo, there is a lack of adhering to standards, consistency and conformity at zedcloud related code/developers. Not talking about EVE.
- The general approach seems to be to address issues with workarounds and conditions instead of identifying the cause and fixing it. These workaround then impose new issues, that get addressed in the same manner. Thus, a tree of conditional workarounds builds up that is responsible for most of the complexity in the zedcloud codebase. The business requirements themselves are not that complex. This makes the code and system very flaky, slow and bloated. And the developer (and probably user) experience very bad.

#### Changes in swagger spec files
- Rename tag/operation-group EdgeNodeConfiguration -> DeviceConfig
- Change /v1/devices endpoint operationId: EdgeNodeConfiguration_QueryEdgeNodes -> GetDeviceList

- Rename DeviceStatusMsg -> DeviceStatus

- Change /v1/devices/id/{id}/status endpoint operationId: EdgeNodeConfiguration_GetEdgeNode -> GetDeviceStatusByID

- Change /v1/devices/id/{objid}/events endpoint operationId: EdgeNodeStatus_GetEdgeNodeEvents -> GetDeviceEventListByID
- Change /v1/devices/name/{objname}/events endpoint operationId: -> EdgeNodeStatus_GetEdgeNodeEventsByName

### TODO
- rename Device to EdgeNode to stay consistent even though endpoints use device?

