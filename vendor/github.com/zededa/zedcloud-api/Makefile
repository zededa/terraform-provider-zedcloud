# Copyright (c) 2018-2021 Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

SWAGGER_FILE_LIST := zedge_node_service.swagger.json zedge_user_service.swagger.json
SWAGGER_FILE_LIST += zedge_storage_service.swagger.json zedge_job_service.swagger.json
SWAGGER_FILE_LIST += zedge_diag_service.swagger.json zedge_app_service.swagger.json
SWAGGER_FILE_LIST += zedge_network_service.swagger.json

all: swagger

clean: swagger-clean

download-swagger:
	# This target needs an env variable SWAGGER_URL_BASE to be set.
	# Example Value for SWAGGER_URL_BASE: https://zedcontrol.zededa.net/api/v1/docs/zapiservices/
	test -n "$(SWAGGER_URL_BASE)"  # Make sure SWAGGER_URL_BASE is set
	for f in $(SWAGGER_FILE_LIST); do \
	 	echo "downloading file: $$f"; \
		wget -O zapiservices/$$f $(SWAGGER_URL_BASE)/$$f; \
	done

swagger: swagger-install swagger-generate

.PHONY: swagger
swagger-install:
	docker pull quay.io/goswagger/swagger

.PHONY: swagger-generate
swagger-generate:
	for sw_file in $(SWAGGER_FILE_LIST); do \
		docker run --rm -it --user $(shell id -u):$(shell id -g) -e GOPATH=$(HOME)/go:/go \
			-v $(HOME):$(HOME) -w $(shell pwd) quay.io/goswagger/swagger \
			generate client -f zapiservices/$$sw_file -A zedcloudapi \
			-c swagger_client -m swagger_models -a swagger_operations && \
		mv swagger_client/zedcloudapi_client.go swagger_client/$(sw_file).go.bkp; \
	done
	#Add CopyRight notice to the generated go files.
	tools/addCopyright
	# Verify build.
	go build

client:
	cd zedcloud_client && go build

swagger-clean:
	rm -rf swagger_client swagger_models swagger_operations
