// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	zedcloudAPI "github.com/zededa/zedcloud-api"
	models "github.com/zededa/zedcloud-api/swagger_models"
)

var dataStoreUrlExtension = "datastores"
var resourceType = "datastore"

func dataSourceDataStore() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDataStoreRead,
		Schema:      zschemas.DataStore,
		Description: "Schema for data source zedcloud_DataStore. Must specify id or name",
	}
}

func dataSourceDataStoreRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := getStr(d, "id")
	name := getStr(d, "name")
	if id == "" && name == "" {
		return diag.Errorf("missing required fields \"id\" or \"name\" in resource data")
	}

	apiState, err := fetch(name, id, meta)
	if err != nil {
		// no object with this id exist in the api's state
		var notFoundErr *ObjectNotFound
		if errors.As(err, &notFoundErr) {
			removeFromLocalState(d, resourceType, id, name)
			return diag.FromErr(fmt.Errorf("[ERROR] could not find datastore %s (id=%s): %w", name, id, err))
		}
		return diag.FromErr(fmt.Errorf("[ERROR] could not fetch datastore %s (id=%s): %w", name, id, err))
	}

	flattenedAPIState, err := flatten(apiState)
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] could not flatten api state for datastore %s (id=%s): %w", name, id, err))
	}

	// we always need to set the state
	setLocalState(d, flattenedAPIState)

	return diags
}

func fetch(name, id string, meta interface{}) (*models.DatastoreInfo, error) {
	client, ok := meta.(Client)
	if !ok {
		return nil, fmt.Errorf("expect meta to be of type %T but is %T", Client{}, meta)
	}

	responseData := &models.DatastoreInfo{}
	resp, err := client.GetObj(dataStoreUrlExtension, name, id, false, responseData)
	if err != nil {
		return nil, err
	}

	if zedcloudAPI.IsObjectNotFound(resp) {
		return nil, &ObjectNotFound{id}
	}

	return responseData, nil
}

func flatten(data *models.DatastoreInfo) (map[string]interface{}, error) {
	flattened := map[string]interface{}{
		"id":   data.ID,
		"name": strPtrVal(data.Name),
	}
	if err := checkIfAllKeysExist(zschemas.DataStore, flattened); err != nil {
		return nil, err
	}
	return flattened, nil
}
