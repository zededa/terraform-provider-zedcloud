package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func DatastoreStatusModel(d *schema.ResourceData) *models.DatastoreStatus {
	datastoreStatus, _ := d.Get("datastore_status").(models.DatastoreStatus)
	return &datastoreStatus
}

func DatastoreStatusModelFromMap(m map[string]interface{}) *models.DatastoreStatus {
	datastoreStatus := m["datastore_status"].(models.DatastoreStatus)
	return &datastoreStatus
}

// Update the underlying DatastoreStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDatastoreStatusResourceData(d *schema.ResourceData, m *models.DatastoreStatus) {
}

// Iterate through and update the DatastoreStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDatastoreStatusSubResourceData(m []*models.DatastoreStatus) (d []*map[string]interface{}) {
	for _, DatastoreStatusModel := range m {
		if DatastoreStatusModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DatastoreStatus resource defined in the Terraform configuration
func DatastoreStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the DatastoreStatus resource
func GetDatastoreStatusPropertyFields() (t []string) {
	return []string{}
}
