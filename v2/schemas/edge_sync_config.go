package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func EdgeSyncConfigModel(d *schema.ResourceData) *models.EdgeSyncConfig {
	var datastoreID []string
	datastoreIDInterface, datastoreIDIsSet := d.GetOk("datastoreID")
	if datastoreIDIsSet && datastoreIDInterface != nil {
		var items []interface{}
		if listItems, isList := datastoreIDInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = datastoreIDInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			datastoreID = append(datastoreID, v.(string))
		}
	}
	edgeSyncURL, _ := d.Get("edge_sync_url").(string)
	return &models.EdgeSyncConfig{
		DatastoreID: datastoreID,
		EdgeSyncURL: edgeSyncURL,
	}
}

func EdgeSyncConfigModelFromMap(m map[string]interface{}) *models.EdgeSyncConfig {
	var datastoreID []string
	datastoreIDInterface, datastoreIDIsSet := m["datastoreID"]
	if datastoreIDIsSet && datastoreIDInterface != nil {
		var items []interface{}
		if listItems, isList := datastoreIDInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = datastoreIDInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			datastoreID = append(datastoreID, v.(string))
		}
	}
	edgeSyncURL := m["edge_sync_url"].(string)
	return &models.EdgeSyncConfig{
		DatastoreID: datastoreID,
		EdgeSyncURL: edgeSyncURL,
	}
}

func SetEdgeSyncConfigResourceData(d *schema.ResourceData, m *models.EdgeSyncConfig) {
	d.Set("datastore_id", m.DatastoreID)
	d.Set("edge_sync_url", m.EdgeSyncURL)
}

func SetEdgeSyncConfigSubResourceData(m []*models.EdgeSyncConfig) (d []*map[string]interface{}) {
	for _, EdgeSyncConfigModel := range m {
		if EdgeSyncConfigModel != nil {
			properties := make(map[string]interface{})
			properties["datastore_id"] = EdgeSyncConfigModel.DatastoreID
			properties["edge_sync_url"] = EdgeSyncConfigModel.EdgeSyncURL
			d = append(d, &properties)
		}
	}
	return
}

func EdgeSyncConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"datastore_id": {
			Description: `upload datastore list`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"edge_sync_url": {
			Description: `edge sync url`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetEdgeSyncConfigPropertyFields() (t []string) {
	return []string{
		"datastore_id",
		"edge_sync_url",
	}
}
