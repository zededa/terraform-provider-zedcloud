package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// Function to perform the following actions:
// (1) Translate PhysicalStorageRaidType resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PhysicalStorageRaidTypeModel(d *schema.ResourceData) *models.PhysicalStorageRaidType {
	physicalStorageRaidType, _ := d.Get("physical_storage_raid_type").(models.PhysicalStorageRaidType)
	return &physicalStorageRaidType
}

func PhysicalStorageRaidTypeModelFromMap(m map[string]interface{}) *models.PhysicalStorageRaidType {
	physicalStorageRaidType := m["physical_storage_raid_type"].(models.PhysicalStorageRaidType)
	return &physicalStorageRaidType
}

// Update the underlying PhysicalStorageRaidType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPhysicalStorageRaidTypeResourceData(d *schema.ResourceData, m *models.PhysicalStorageRaidType) {
}

// Iterate through and update the PhysicalStorageRaidType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPhysicalStorageRaidTypeSubResourceData(m []*models.PhysicalStorageRaidType) (d []*map[string]interface{}) {
	for _, PhysicalStorageRaidTypeModel := range m {
		if PhysicalStorageRaidTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the PhysicalStorageRaidType resource defined in the Terraform configuration
func PhysicalStorageRaidTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the PhysicalStorageRaidType resource
func GetPhysicalStorageRaidTypePropertyFields() (t []string) {
	return []string{}
}
