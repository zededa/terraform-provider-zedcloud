package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate StorageStatus resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func StorageStatusModel(d *schema.ResourceData) *models.StorageStatus {
	mountPath, _ := d.Get("mount_path").(string)
	name, _ := d.Get("name").(string)
	sizeMB, _ := d.Get("size_m_b").(string)
	return &models.StorageStatus{
		MountPath: mountPath,
		Name:      name,
		SizeMB:    sizeMB,
	}
}

func StorageStatusModelFromMap(m map[string]interface{}) *models.StorageStatus {
	mountPath := m["mount_path"].(string)
	name := m["name"].(string)
	sizeMB := m["size_m_b"].(string)
	return &models.StorageStatus{
		MountPath: mountPath,
		Name:      name,
		SizeMB:    sizeMB,
	}
}

// Update the underlying StorageStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetStorageStatusResourceData(d *schema.ResourceData, m *models.StorageStatus) {
	d.Set("mount_path", m.MountPath)
	d.Set("name", m.Name)
	d.Set("size_m_b", m.SizeMB)
}

// Iterate through and update the StorageStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetStorageStatusSubResourceData(m []*models.StorageStatus) (d []*map[string]interface{}) {
	for _, StorageStatusModel := range m {
		if StorageStatusModel != nil {
			properties := make(map[string]interface{})
			properties["mount_path"] = StorageStatusModel.MountPath
			properties["name"] = StorageStatusModel.Name
			properties["size_m_b"] = StorageStatusModel.SizeMB
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the StorageStatus resource defined in the Terraform configuration
func StorageStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"mount_path": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"size_m_b": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the StorageStatus resource
func GetStorageStatusPropertyFields() (t []string) {
	return []string{
		"mount_path",
		"name",
		"size_m_b",
	}
}
