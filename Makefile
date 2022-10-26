# Copyright (c) 2018-2021 Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0
default: testacc

all: plugin

.PHONY: plugin
plugin:
	go build -o terraform-provider-zedcloud_v0.0.0
	chmod a+x terraform-provider-zedcloud_v0.0.0

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

.PHONY: gen
gen:
	go generate
