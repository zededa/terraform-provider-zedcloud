# terraform-provider-zedcloud

[![Terraform Registry](https://img.shields.io/badge/terraform-registry-blue.svg)](https://registry.terraform.io/providers/zededa/zedcloud/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/zededa/terraform-provider-zedcloud)](https://goreportcard.com/report/github.com/zededa/terraform-provider-zedcloud)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

## Overview

The Terraform Provider for Zededa Cloud enables infrastructure automation for Zededa's edge computing platform. It provides seamless integration with Zededa's public API, allowing you to manage your edge computing resources through Infrastructure as Code.

## Supported Resources

- [x] Edge-Node (Device-Config)
- [x] Edge-App
- [x] Edge-App-Instance
- [x] Network
- [x] Network-Instance
- [x] Datastore
- [x] Volume-Instance
- [x] User
- [x] Role
- [x] Credential
- [x] Patch Envelope
- [ ] Image

## Quickstart

1. **Installation**
```
go install github.com/zededa/terraform-provider-zedcloud/v2@latest
```

2. **Configuration**

Set the following environment variables to use the provider:
```
export TF_VAR_zedcloud_url="zedcontrol.zededa.net"
export TF_VAR_zedcloud_token=<YOUR-API-TOKEN>
export TF_LOG=ERROR
export TF_LOG_PATH=./terraform.log
```

## Documentation

### Schemas

Documentation of all API endpoints and data schemas can be found under https://zedcontrol.zededa.net/api/v1/docs/.

Schema documentation generated from the schema files contained in this repo can be found under:
1. [Provider](https://github.com/zededa/terraform-provider-zedcloud/blob/main/docs/index.md)
1. [Resources](https://github.com/zededa/terraform-provider-zedcloud/tree/main/docs/resources)
1. [Data Sources](https://github.com/zededa/terraform-provider-zedcloud/tree/main/docs/data-sources)

> Note, the resource schemas match the API schemas but support for some fields might be incomplete in the provider. The testdata directory contains examples with the full set of supported fields for the supported resources of the latest version deployed in the Terraform provider registry.

### Product

Product documentation including explanations of workflows and data schemas can be found under https://help.zededa.com/hc/en-us

## Breaking changes in v2

- Authentication uses the API-token only. Basic-Auth via username and password has been removed. You can find the API-token for your user under https://zedcontrol.zededa.net/profile/user
- The resource and data schemas have changed. The schemas now map to the ones in API documentation. For reference on supported fields see the testdata directory which contains Terraform configuration files for all supported resources. If you have trouble porting your configuration, please reach out to Zededa support.
- `adminstate_config` has been removed. `admin_state` can now be configured directly. Diffs will be suppressed if the `EdgeNode` is in `ADMINSTATE_REGISTERED` state and your configuration is set to `ADMINSTATE_ACTIVE`.
- CSRF checks have been removed.

## Project structure

The project structure has changed starting from the version 2.2.7. All the folders directly related to the provider are moved to the `v2` subfolder.

## Contribution

We welcome contributions! Please see our [contributing guidelines](CONTRIBUTING.md) for more details.

## Support

If you need help setting up the __terraform-provider-zedcloud__ please reach out to Zededa support.

The latest version of the provider can be found in the official Terraform provider registry under https://registry.terraform.io/providers/zededa/zedcloud/latest.

