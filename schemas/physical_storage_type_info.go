package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate PhysicalStorageTypeInfo resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PhysicalStorageTypeInfoModel(d *schema.ResourceData) *models.PhysicalStorageTypeInfo {
	physicalStorageTypeInfo, _ := d.Get("physical_storage_type_info").(models.PhysicalStorageTypeInfo)
	return &physicalStorageTypeInfo
}

func PhysicalStorageTypeInfoModelFromMap(m map[string]interface{}) *models.PhysicalStorageTypeInfo {
	physicalStorageTypeInfo := m["physical_storage_type_info"].(models.PhysicalStorageTypeInfo)
	return &physicalStorageTypeInfo
}

// Update the underlying PhysicalStorageTypeInfo resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPhysicalStorageTypeInfoResourceData(d *schema.ResourceData, m *models.PhysicalStorageTypeInfo) {
}

// Iterate through and update the PhysicalStorageTypeInfo resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPhysicalStorageTypeInfoSubResourceData(m []*models.PhysicalStorageTypeInfo) (d []*map[string]interface{}) {
	for _, PhysicalStorageTypeInfoModel := range m {
		if PhysicalStorageTypeInfoModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the PhysicalStorageTypeInfo resource defined in the Terraform configuration
func PhysicalStorageTypeInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the PhysicalStorageTypeInfo resource
func GetPhysicalStorageTypeInfoPropertyFields() (t []string) {
	return []string{}
}
