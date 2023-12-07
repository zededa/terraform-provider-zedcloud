# Copyright (c) 2018-2021 Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

SWAGGER_FILE_LIST := zedge_node_service.swagger.json \
					 zedge_user_service.swagger.json \
					 zedge_storage_service.swagger.json \
					 zedge_job_service.swagger.json \
					 zedge_diag_service.swagger.json \
					 zedge_app_service.swagger.json \
					 zedge_network_service.swagger.json

NAME = terraform-provider-zedcloud
VERSION = $(shell git describe --always --abbrev=0 --tags)
SWAGGER_URL_BASE = https://zedcontrol.zededa.net/api/v1/docs/zapiservices/

.PHONY: download-specs
download-specs:
	# This target needs an env variable SWAGGER_URL_BASE to be set.
	# Example Value for SWAGGER_URL_BASE: https://zedcontrol.zededa.net/api/v1/docs/zapiservices/
	test -n "$(SWAGGER_URL_BASE)"  # Make sure SWAGGER_URL_BASE is set
	for f in $(SWAGGER_FILE_LIST); do \
	 	echo "downloading file: $$f"; \
		wget -O swagger/$$f $(SWAGGER_URL_BASE)/$$f; \
	done

.PHONY: gen
gen:
	swagger generate client -f swagger/$(src).swagger.json \
		-A zedcloudapi \
		-C swagger/config.yml

.PHONY: build
build:
	rm -f terraform-provider-zedcloud_v*
	go build -o $(NAME)_$(VERSION)
	chmod a+x  $(NAME)_$(VERSION)

# acceptance tests
.PHONY: test
test:
	TF_ACC=1 go test -v ./...

.PHONY: test-run
test-run:
	go vet ./...
	TF_ACC=1 go test -v ./... -run $(case)

.PHONY: udpatedeps
updatedeps:
	go get -f -t -u ./...
	go get -f -u ./...
