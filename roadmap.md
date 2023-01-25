### New version of the zedcloud-terraform-provider

Version 2 is a rewrite of the entire code base. The purpose is:
- Have clean code that is easy to maintain and extend and follows Hashicorp best practices for Terraform providers
- Utilize internal state management only instead of custom one to provide expected and documented behaviour
- Introduce full test coverage for all resources via Terraform acceptance tests
- Extend the provider to support all resources in the Zedcloud API
- Fix several bugs and provide automatic retries in case of failure

### Roadmap

The roadmap should provide a rough overview of releases and features as well as bug-fixes.
It is intended to be used by PM (Product Management) as the base for communicating changes, requested features and progress in general.
It will be updated in case a milestones changes.
In this case, a notification will be sent to the slack channel and PM can update customers as well.

- __M1: 02/01/2023 - Release of v2__
    - This replaces the current version v1 and provides the same feature set with several improvements and fixes as well as full test coverage against a Zedcloud instance.
- __M2: 02/01-02/07/2023 - Test and release__
    - Beta testing with PM internally before releasing to customers.
	- Support for automatic retries with exponential back-off
- __M3: 02/07/2023 - Release__
	- Fixes for issues that came up in beta.
	- Indicate failure in app-image deletion + creation for high volume workloads (this issue is on API side, a fix is unlikely currently but it is on the agenda of the Zedcloud team).
    - Fix for network instance deletion errors.
- __M4: 02/07/2023 - Support__
    - Help customers/engineers porting their existing configurations to the new simplified schemas.
	- Assist in real-time (Keybase, email, slack/support channels) and keep time slots free for addressing issues and bug-fixes immediately.
- __M5 - 21/02/2023 - Release__
    - Full support for all API resources.
- __M6 - 28/02/2023 - Test__
	- CI/CD run tests  in GitHub actions against staging environment.
	- Add links to the documentation to the Zededa Help Center articles

### Issues

State of issues pointed out by customers:
- [x] admin_state
	- Status: fixed in v1 and v2
	- Deployed in version: v1.0.7
	- Reported by VW
- [ ] Support for automatic retries with exponential back-off
	- Status: in development for v2
	- Deployed in version: planned for milestone M2
	- Reported by VW
- [ ] Indicate failure in app image deletion + creation
	- Status: to do
	- Deployed in version: planned for milestone M3
	- Reported by VW
- [ ] Zedcloud API - network instance deletion errors
	- Status: to do
	- Deployed in version: planned for milestone M3
	- Reported by VW


### Documentation

The documentation is hosted by Hashicorp in the provider registry.
It gets updated automatically with each build and deployed along with the provider.
https://registry.terraform.io/providers/zededa/zedcloud/latest/docs


