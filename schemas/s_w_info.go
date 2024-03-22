package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func SWInfoModel(d *schema.ResourceData) *models.SWInfo {
	disk, _ := d.Get("disk").(string)
	downloadProgressInt, _ := d.Get("download_progress").(int)
	downloadProgress := int64(downloadProgressInt)
	hash, _ := d.Get("hash").(string)
	imageName, _ := d.Get("image_name").(string)
	var state *models.SWState // SWState
	stateInterface, stateIsSet := d.GetOk("state")
	if stateIsSet {
		stateModel := stateInterface.(string)
		state = models.NewSWState(models.SWState(stateModel))
	}
	version, _ := d.Get("version").(string)
	return &models.SWInfo{
		Disk:             disk,
		DownloadProgress: downloadProgress,
		Hash:             hash,
		ImageName:        imageName,
		State:            state,
		Version:          version,
	}
}

func SWInfoModelFromMap(m map[string]interface{}) *models.SWInfo {
	disk := m["disk"].(string)
	downloadProgress := int64(m["download_progress"].(int)) // int64
	hash := m["hash"].(string)
	imageName := m["image_name"].(string)
	var state *models.SWState // SWState
	stateInterface, stateIsSet := m["state"]
	if stateIsSet {
		stateModel := stateInterface.(string)
		state = models.NewSWState(models.SWState(stateModel))
	}
	version := m["version"].(string)
	return &models.SWInfo{
		Disk:             disk,
		DownloadProgress: downloadProgress,
		Hash:             hash,
		ImageName:        imageName,
		State:            state,
		Version:          version,
	}
}

// Update the underlying SWInfo resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetSWInfoResourceData(d *schema.ResourceData, m *models.SWInfo) {
	d.Set("disk", m.Disk)
	d.Set("download_progress", m.DownloadProgress)
	d.Set("hash", m.Hash)
	d.Set("image_name", m.ImageName)
	d.Set("state", m.State)
	d.Set("version", m.Version)
}

// Iterate through and update the SWInfo resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetSWInfoSubResourceData(m []*models.SWInfo) (d []*map[string]interface{}) {
	for _, SWInfoModel := range m {
		if SWInfoModel != nil {
			properties := make(map[string]interface{})
			properties["disk"] = SWInfoModel.Disk
			properties["download_progress"] = SWInfoModel.DownloadProgress
			properties["hash"] = SWInfoModel.Hash
			properties["image_name"] = SWInfoModel.ImageName
			properties["state"] = SWInfoModel.State
			properties["version"] = SWInfoModel.Version
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the SWInfo resource defined in the Terraform configuration
func SWInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"disk": {
			Description: `software disk`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"download_progress": {
			Description: `download progress`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"hash": {
			Description: `software hash`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"image_name": {
			Description: `image name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"state": {
			Description: `Software state`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"version": {
			Description: `software version`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the SWInfo resource
func GetSWInfoPropertyFields() (t []string) {
	return []string{
		"disk",
		"download_progress",
		"hash",
		"image_name",
		"state",
		"version",
	}
}
