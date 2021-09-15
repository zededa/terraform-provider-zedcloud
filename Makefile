# Copyright (c) 2018-2021 Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

all: plugin

plugin:
	go mod download
	go build -o terraform-provider-zedcloud_v0.0.0
	chmod a+x terraform-provider-zedcloud_v0.0.0
