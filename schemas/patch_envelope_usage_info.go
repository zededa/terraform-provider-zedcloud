package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func PatchEnvelopeUsageInfoModel(d *schema.ResourceData) *models.PatchEnvelopeUsageInfo {
	downloadCount, _ := d.Get("download_count").(string)
	patchAPICallCount, _ := d.Get("patch_api_call_count").(uint64)
	uuid, _ := d.Get("uuid").(string)
	version, _ := d.Get("version").(string)
	return &models.PatchEnvelopeUsageInfo{
		DownloadCount:     downloadCount,
		PatchAPICallCount: patchAPICallCount,
		UUID:              uuid,
		Version:           version,
	}
}

func PatchEnvelopeUsageInfoModelFromMap(m map[string]interface{}) *models.PatchEnvelopeUsageInfo {
	downloadCount := m["download_count"].(string)
	patchAPICallCount := m["patch_api_call_count"].(uint64)
	uuid := m["uuid"].(string)
	version := m["version"].(string)
	return &models.PatchEnvelopeUsageInfo{
		DownloadCount:     downloadCount,
		PatchAPICallCount: patchAPICallCount,
		UUID:              uuid,
		Version:           version,
	}
}

func SetPatchEnvelopeUsageInfoResourceData(d *schema.ResourceData, m *models.PatchEnvelopeUsageInfo) {
	d.Set("download_count", m.DownloadCount)
	d.Set("patch_api_call_count", m.PatchAPICallCount)
	d.Set("uuid", m.UUID)
	d.Set("version", m.Version)
}

func SetPatchEnvelopeUsageInfoSubResourceData(m []*models.PatchEnvelopeUsageInfo) (d []*map[string]interface{}) {
	for _, PatchEnvelopeUsageInfoModel := range m {
		if PatchEnvelopeUsageInfoModel != nil {
			properties := make(map[string]interface{})
			properties["download_count"] = PatchEnvelopeUsageInfoModel.DownloadCount
			properties["patch_api_call_count"] = PatchEnvelopeUsageInfoModel.PatchAPICallCount
			properties["uuid"] = PatchEnvelopeUsageInfoModel.UUID
			properties["version"] = PatchEnvelopeUsageInfoModel.Version
			d = append(d, &properties)
		}
	}
	return
}

func PatchEnvelopeUsageInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"download_count": {
			Description: ` number of times app instance download whole patch envelope or part of it`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"patch_api_call_count": {
			Description: `number of times app instance called patch APIs`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"uuid": {
			Description: `Id of patch envelope`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"version": {
			Description: `version of patch envelope`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetPatchEnvelopeUsageInfoPropertyFields() (t []string) {
	return []string{
		"download_count",
		"patch_api_call_count",
		"uuid",
		"version",
	}
}
