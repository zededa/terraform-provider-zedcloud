package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate StorageDeviceMetrics resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func StorageDeviceMetricsModel(d *schema.ResourceData) *models.StorageDeviceMetrics {
	checksumErrors := int64(d.Get("checksum_errors").(int))
	readErrors := int64(d.Get("read_errors").(int))
	writeErrors := int64(d.Get("write_errors").(int))
	return &models.StorageDeviceMetrics{
		ChecksumErrors: checksumErrors,
		ReadErrors:     readErrors,
		WriteErrors:    writeErrors,
	}
}

func StorageDeviceMetricsModelFromMap(m map[string]interface{}) *models.StorageDeviceMetrics {
	checksumErrors := int64(m["checksum_errors"].(int)) // int64 false false false
	readErrors := int64(m["read_errors"].(int))         // int64 false false false
	writeErrors := int64(m["write_errors"].(int))       // int64 false false false
	return &models.StorageDeviceMetrics{
		ChecksumErrors: checksumErrors,
		ReadErrors:     readErrors,
		WriteErrors:    writeErrors,
	}
}

// Update the underlying StorageDeviceMetrics resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetStorageDeviceMetricsResourceData(d *schema.ResourceData, m *models.StorageDeviceMetrics) {
	d.Set("checksum_errors", m.ChecksumErrors)
	d.Set("read_errors", m.ReadErrors)
	d.Set("write_errors", m.WriteErrors)
}

// Iterate throught and update the StorageDeviceMetrics resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetStorageDeviceMetricsSubResourceData(m []*models.StorageDeviceMetrics) (d []*map[string]interface{}) {
	for _, StorageDeviceMetricsModel := range m {
		if StorageDeviceMetricsModel != nil {
			properties := make(map[string]interface{})
			properties["checksum_errors"] = StorageDeviceMetricsModel.ChecksumErrors
			properties["read_errors"] = StorageDeviceMetricsModel.ReadErrors
			properties["write_errors"] = StorageDeviceMetricsModel.WriteErrors
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the StorageDeviceMetrics resource defined in the Terraform configuration
func StorageDeviceMetricsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"checksum_errors": {
			Description: `checksum errors counter`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"read_errors": {
			Description: `read errors counter`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"write_errors": {
			Description: `write errors counter`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the StorageDeviceMetrics resource
func GetStorageDeviceMetricsPropertyFields() (t []string) {
	return []string{
		"checksum_errors",
		"read_errors",
		"write_errors",
	}
}
