// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

module github.com/zededa/terraform-provider-zedcloud

go 1.16

require (
	github.com/agl/ed25519 v0.0.0-20170116200512-5312a6153412 // indirect
	github.com/alcortesm/tgz v0.0.0-20161220082320-9c5fe88206d7 // indirect
	github.com/go-openapi/strfmt v0.20.2
	github.com/go-test/deep v1.0.7
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	github.com/keybase/go-crypto v0.0.0-20161004153544-93f5b35093ba // indirect
	github.com/zededa/zedcloud-api v0.0.6-alpha
)

// replace github.com/zededa/zedcloud-api => ../zedcloud-api
