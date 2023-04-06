package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func DatastoreTypeModel(d *schema.ResourceData) *models.DatastoreType {
	datastoreType, _ := d.Get("datastore_type").(models.DatastoreType)
	return &datastoreType
}

func DatastoreTypeModelFromMap(m map[string]interface{}) *models.DatastoreType {
	datastoreType := m["datastore_type"].(models.DatastoreType)
	return &datastoreType
}

// Update the underlying DatastoreType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDatastoreTypeResourceData(d *schema.ResourceData, m *models.DatastoreType) {
}

// Iterate through and update the DatastoreType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDatastoreTypeSubResourceData(m []*models.DatastoreType) (d []*map[string]interface{}) {
	for _, DatastoreTypeModel := range m {
		if DatastoreTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DatastoreType resource defined in the Terraform configuration
func DatastoreTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the DatastoreType resource
func GetDatastoreTypePropertyFields() (t []string) {
	return []string{}
}
