package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate PhysicalStorageStatus resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PhysicalStorageStatusModel(d *schema.ResourceData) *models.PhysicalStorageStatus {
	physicalStorageStatus, _ := d.Get("physical_storage_status").(models.PhysicalStorageStatus)
	return &physicalStorageStatus
}

func PhysicalStorageStatusModelFromMap(m map[string]interface{}) *models.PhysicalStorageStatus {
	physicalStorageStatus := m["physical_storage_status"].(models.PhysicalStorageStatus)
	return &physicalStorageStatus
}

// Update the underlying PhysicalStorageStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPhysicalStorageStatusResourceData(d *schema.ResourceData, m *models.PhysicalStorageStatus) {
}

// Iterate through and update the PhysicalStorageStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPhysicalStorageStatusSubResourceData(m []*models.PhysicalStorageStatus) (d []*map[string]interface{}) {
	for _, PhysicalStorageStatusModel := range m {
		if PhysicalStorageStatusModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the PhysicalStorageStatus resource defined in the Terraform configuration
func PhysicalStorageStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the PhysicalStorageStatus resource
func GetPhysicalStorageStatusPropertyFields() (t []string) {
	return []string{}
}
