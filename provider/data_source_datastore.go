// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
)

var dataStoreUrlExtension = "datastores"

var DataStoreDataSourceSchema = &schema.Resource{
	ReadContext: readDataSourceDataStore,
	Schema:      zschemas.DataStoreSchema,
	Description: "Schema for data source zedcloud_DataStore. Must specify id or name",
}

// The schema for a resource group data source
func getDataStoreDataSourceSchema() *schema.Resource {
	return DataStoreDataSourceSchema
}

func getDataStore(client *zedcloudapi.Client,
	name, id string) (*swagger_models.DatastoreInfo, error) {
	rspData := &swagger_models.DatastoreInfo{}
	err := client.GetObj(dataStoreUrlExtension, name, id, false, rspData)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] FAILED to get DataStore %s ( id: %s). Err: %s",
			name, id, err.Error())
	}
	return rspData, nil
}

func flattenDatastoreInfo(cfg *swagger_models.DatastoreInfo, computedOnly bool) map[string]interface{} {
	data := map[string]interface{}{
		"id": cfg.ID,
	}
	if !computedOnly {
		data["name"] = ptrValStr(cfg.Name)
	}
	flattenedDataCheckKeys(zschemas.DataStoreSchema, data, computedOnly)
	return data
}

func getDataStoreAndPublishData(client *zedcloudapi.Client, d *schema.ResourceData,
	name, id string, resource bool) error {
	cfg, err := getDataStore(client, name, id)
	if err != nil {
		return fmt.Errorf("[ERROR] App Instance %s (id: %s) not found. Err: %s",
			name, id, err.Error())
	}
	marshalData(d, flattenDatastoreInfo(cfg, resource))
	return nil
}

// Read the Resource Group
func readDataStore(ctx context.Context, d *schema.ResourceData, meta interface{}, resource bool) diag.Diagnostics {
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	id := rdEntryStr(d, "id")
	name := rdEntryStr(d, "name")
	if client == nil {
		return diag.Errorf("nil Client.")
	}
	if (id == "") && (name == "") {
		return diag.Errorf("The arguments \"id\" or \"name\" are required, " +
			"but no definition was found.")
	}
	err := getDataStoreAndPublishData(client, d, name, id, resource)
	if err != nil {
		return diag.Errorf("%s", err.Error())
	}
	return diags
}

func readDataSourceDataStore(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readDataStore(ctx, d, meta, false)
}
