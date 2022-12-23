package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate PhysicalStorage resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PhysicalStorageModel(d *schema.ResourceData) *models.PhysicalStorage {
	compressionRatio := d.Get("compression_ratio").(float64)
	countZvols := int64(d.Get("count_zvols").(int))
	currentRaid := d.Get("current_raid").(*models.PhysicalStorageRaidType) // PhysicalStorageRaidType
	disks := d.Get("disks").([]*models.PhysicalStorageDiskState)           // []*PhysicalStorageDiskState
	poolName := d.Get("pool_name").(string)
	poolStatusMsg := d.Get("pool_status_msg").(string)
	storageState := d.Get("storage_state").(*models.PhysicalStorageStatus) // PhysicalStorageStatus
	storageType := d.Get("storage_type").(*models.PhysicalStorageTypeInfo) // PhysicalStorageTypeInfo
	zfsVersion := d.Get("zfs_version").(string)
	zpoolSize := d.Get("zpool_size").(string)
	return &models.PhysicalStorage{
		CompressionRatio: compressionRatio,
		CountZvols:       countZvols,
		CurrentRaid:      currentRaid,
		Disks:            disks,
		PoolName:         poolName,
		PoolStatusMsg:    poolStatusMsg,
		StorageState:     storageState,
		StorageType:      storageType,
		ZfsVersion:       zfsVersion,
		ZpoolSize:        zpoolSize,
	}
}

func PhysicalStorageModelFromMap(m map[string]interface{}) *models.PhysicalStorage {
	compressionRatio := m["compression_ratio"].(float64)
	countZvols := int64(m["count_zvols"].(int))                        // int64 false false false
	currentRaid := m["current_raid"].(*models.PhysicalStorageRaidType) // PhysicalStorageRaidType
	disks := m["disks"].([]*models.PhysicalStorageDiskState)           // []*PhysicalStorageDiskState
	poolName := m["pool_name"].(string)
	poolStatusMsg := m["pool_status_msg"].(string)
	storageState := m["storage_state"].(*models.PhysicalStorageStatus) // PhysicalStorageStatus
	storageType := m["storage_type"].(*models.PhysicalStorageTypeInfo) // PhysicalStorageTypeInfo
	zfsVersion := m["zfs_version"].(string)
	zpoolSize := m["zpool_size"].(string)
	return &models.PhysicalStorage{
		CompressionRatio: compressionRatio,
		CountZvols:       countZvols,
		CurrentRaid:      currentRaid,
		Disks:            disks,
		PoolName:         poolName,
		PoolStatusMsg:    poolStatusMsg,
		StorageState:     storageState,
		StorageType:      storageType,
		ZfsVersion:       zfsVersion,
		ZpoolSize:        zpoolSize,
	}
}

// Update the underlying PhysicalStorage resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPhysicalStorageResourceData(d *schema.ResourceData, m *models.PhysicalStorage) {
	d.Set("compression_ratio", m.CompressionRatio)
	d.Set("count_zvols", m.CountZvols)
	d.Set("current_raid", m.CurrentRaid)
	d.Set("disks", SetPhysicalStorageDiskStateSubResourceData(m.Disks))
	d.Set("pool_name", m.PoolName)
	d.Set("pool_status_msg", m.PoolStatusMsg)
	d.Set("storage_state", m.StorageState)
	d.Set("storage_type", m.StorageType)
	d.Set("zfs_version", m.ZfsVersion)
	d.Set("zpool_size", m.ZpoolSize)
}

// Iterate throught and update the PhysicalStorage resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPhysicalStorageSubResourceData(m []*models.PhysicalStorage) (d []*map[string]interface{}) {
	for _, PhysicalStorageModel := range m {
		if PhysicalStorageModel != nil {
			properties := make(map[string]interface{})
			properties["compression_ratio"] = PhysicalStorageModel.CompressionRatio
			properties["count_zvols"] = PhysicalStorageModel.CountZvols
			properties["current_raid"] = PhysicalStorageModel.CurrentRaid
			properties["disks"] = SetPhysicalStorageDiskStateSubResourceData(PhysicalStorageModel.Disks)
			properties["pool_name"] = PhysicalStorageModel.PoolName
			properties["pool_status_msg"] = PhysicalStorageModel.PoolStatusMsg
			properties["storage_state"] = PhysicalStorageModel.StorageState
			properties["storage_type"] = PhysicalStorageModel.StorageType
			properties["zfs_version"] = PhysicalStorageModel.ZfsVersion
			properties["zpool_size"] = PhysicalStorageModel.ZpoolSize
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the PhysicalStorage resource defined in the Terraform configuration
func PhysicalStorageSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"compression_ratio": {
			Description: ``,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"count_zvols": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"current_raid": {
			Description: ``,
			Type:        schema.TypeList, //GoType: PhysicalStorageRaidType
			Elem: &schema.Resource{
				Schema: PhysicalStorageRaidTypeSchema(),
			},
			Optional: true,
		},

		"disks": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*PhysicalStorageDiskState
			Elem: &schema.Resource{
				Schema: PhysicalStorageDiskStateSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"pool_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"pool_status_msg": {
			Description: `Zpool Status message sent by EVE`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"storage_state": {
			Description: ``,
			Type:        schema.TypeList, //GoType: PhysicalStorageStatus
			Elem: &schema.Resource{
				Schema: PhysicalStorageStatusSchema(),
			},
			Optional: true,
		},

		"storage_type": {
			Description: ``,
			Type:        schema.TypeList, //GoType: PhysicalStorageTypeInfo
			Elem: &schema.Resource{
				Schema: PhysicalStorageTypeInfoSchema(),
			},
			Optional: true,
		},

		"zfs_version": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"zpool_size": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the PhysicalStorage resource
func GetPhysicalStoragePropertyFields() (t []string) {
	return []string{
		"compression_ratio",
		"count_zvols",
		"current_raid",
		"disks",
		"pool_name",
		"pool_status_msg",
		"storage_state",
		"storage_type",
		"zfs_version",
		"zpool_size",
	}
}
