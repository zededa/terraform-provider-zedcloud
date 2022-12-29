package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate PhysicalStorageDiskState resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PhysicalStorageDiskStateModel(d *schema.ResourceData) *models.PhysicalStorageDiskState {
	var disk *models.DiskDescription // DiskDescription
	diskInterface, diskIsSet := d.GetOk("disk")
	if diskIsSet {
		diskMap := diskInterface.([]interface{})[0].(map[string]interface{})
		disk = DiskDescriptionModelFromMap(diskMap)
	}
	statusModel, ok := d.Get("status").(models.PhysicalStorageStatus) // PhysicalStorageStatus
	status := &statusModel
	if !ok {
		status = nil
	}
	return &models.PhysicalStorageDiskState{
		Disk:   disk,
		Status: status,
	}
}

func PhysicalStorageDiskStateModelFromMap(m map[string]interface{}) *models.PhysicalStorageDiskState {
	var disk *models.DiskDescription // DiskDescription
	diskInterface, diskIsSet := m["disk"]
	if diskIsSet {
		diskMap := diskInterface.([]interface{})[0].(map[string]interface{})
		disk = DiskDescriptionModelFromMap(diskMap)
	}
	//
	status := m["status"].(*models.PhysicalStorageStatus) // PhysicalStorageStatus
	return &models.PhysicalStorageDiskState{
		Disk:   disk,
		Status: status,
	}
}

// Update the underlying PhysicalStorageDiskState resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPhysicalStorageDiskStateResourceData(d *schema.ResourceData, m *models.PhysicalStorageDiskState) {
	d.Set("disk", SetDiskDescriptionSubResourceData([]*models.DiskDescription{m.Disk}))
	d.Set("status", m.Status)
}

// Iterate through and update the PhysicalStorageDiskState resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPhysicalStorageDiskStateSubResourceData(m []*models.PhysicalStorageDiskState) (d []*map[string]interface{}) {
	for _, PhysicalStorageDiskStateModel := range m {
		if PhysicalStorageDiskStateModel != nil {
			properties := make(map[string]interface{})
			properties["disk"] = SetDiskDescriptionSubResourceData([]*models.DiskDescription{PhysicalStorageDiskStateModel.Disk})
			properties["status"] = PhysicalStorageDiskStateModel.Status
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the PhysicalStorageDiskState resource defined in the Terraform configuration
func PhysicalStorageDiskStateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"disk": {
			Description: ``,
			Type:        schema.TypeList, //GoType: DiskDescription
			Elem: &schema.Resource{
				Schema: DiskDescriptionSchema(),
			},
			Optional: true,
		},

		"status": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the PhysicalStorageDiskState resource
func GetPhysicalStorageDiskStatePropertyFields() (t []string) {
	return []string{
		"disk",
		"status",
	}
}
