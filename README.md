# terraform-provider-zedcloud

The __terraform-provider-zedcloud__ provides access to Zededa's public API for the zedcloud cloud services.

Supported resources are:
- [x] Edge-Node (also called Device-Config)
- [x] Application
- [x] Application-Instance
- [x] Network
- [x] Network-Instance
- [x] Datastore
- [x] Volume-Instance
- [ ] Image

## Documentation

### Schemas

Documentation of all API endpoints and data schemas can be found under https://zedcontrol.zededa.net/api/v1/docs/.

Schema documentation generated from the schema files contained in this repo can be found under:
1. [Provider](https://github.com/zededa/terraform-provider-zedcloud/tree/main/docs/provider)
1. [Resources](https://github.com/zededa/terraform-provider-zedcloud/tree/main/docs/resources)
1. [Data Sources](https://github.com/zededa/terraform-provider-zedcloud/tree/main/docs/data-sources)

> Note, the resource schemas match the API schemas but support for some fields might be incomplete in the provider. The testdata directory contains examples with the full set of supported fields for the supported resources of the latest version deployed in the Terraform provider registry.

---

### Product

Product documentation including explanation of workflows and data schemas can be found under https://help.zededa.com/hc/en-us

## Installation and configuration

Information on how to install and configure a Terraform provider can be found under https://developer.hashicorp.com/terraform/language/providers. If you need help setting up the __terraform-provider-zedcloud__ please reach out to the Zededa support.
The latest version of the provider can be found in the official Terraform provider registry under https://registry.terraform.io/providers/zededa/zedcloud/latest.

## Breaking changes in v2

- Authentication uses the API-token only. Basic-Auth via username and password has been removed. You can find the API-token for your user under https://zedcontrol.zededa.net/profile/user
- The resource and data schemas have changed. The schemas now map to the ones in API documentation. For reference on supported fields see the testdata directory which contains Terraform configuration files for all supported resources. If you have trouble porting your configuration, please reach out to the Zededa support.
- `aminstate_config` has been removed. `admin_state` can now be configured directly. Diffs will be suppressed if the `EdgeNode` is in `ADMINSTATE_REGISTERED` state and your configuration is set to `ADMINSTATE_ACTIVE`.
- CSRF checks have been removed.

## Environment configuration

The following environment variables need to be set to use the provider.
```
export TF_VAR_zedcloud_url="zedcontrol.zededa.net"
export TF_VAR_zedcloud_token=<YOUR-API-TOKEN>
export TF_LOG=ERROR
export TF_LOG_PATH=./terraform.log
```

## Roadmap

- __M1: 02/14/2023 - Release of v2__
    - This replaces the current version v1 and provides the same feature set with several improvements and fixes as well as full test coverage against a Zedcloud instance.
- __M2: 02/14-02/20/2023 - Test and release__
    - Beta testing with PM internally before releasing to customers.
	- Support for automatic retries with exponential back-off
- __M3: 02/20/2023 - Release__
	- Fixes for issues that came up in beta.
	- Indicate failure in app-image deletion + creation for high volume workloads (this issue is on API side, a fix is unlikely currently but it is on the agenda of the Zedcloud team).
    - Fix for network instance deletion errors.
- __M4: 02/20/2023 - Support__
    - Help customers/engineers porting their existing configurations to the new simplified schemas.
	- Assist in real-time (Keybase, email, slack/support channels) and keep time slots free for addressing issues and bug-fixes immediately.
- __M5 - 02/27/2023 - Release__
    - Full support for all API resources.
- __M6 - 03/05/2023 - Test__
	- CI/CD run tests  in GitHub actions against staging environment.
	- Add links to the documentation to the Zededa Help Center articles

### Issues

State of issues pointed out by customers:
- [x] admin_state
	- Status: fixed in v1 and v2
	- Deployed in version: v1.0.7
- [x] Zedcloud API - network instance deletion errors
	- Status: fixed in v2
	- Deployed in version: v2.0.0
- [ ] Support for automatic retries with exponential back-off
	- Status: in development for v2
	- Deployed in version: planned for milestone M2
- [ ] Indicate failure in app image deletion + creation
	- Status: to do
	- Deployed in version: planned for milestone M3

