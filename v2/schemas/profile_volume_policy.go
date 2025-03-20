package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ProfileVolumePolicyModel(d *schema.ResourceData) *models.ProfileVolumePolicy {
	var metaData *models.ProfilePolicyCommon // ProfilePolicyCommon
	metaDataInterface, metaDataIsSet := d.GetOk("meta_data")
	if metaDataIsSet && metaDataInterface != nil {
		metaDataMap := metaDataInterface.([]interface{})
		if len(metaDataMap) > 0 {
			metaData = ProfilePolicyCommonModelFromMap(metaDataMap[0].(map[string]interface{}))
		}
	}
	var volumeConfig *models.ProfileVolumeConfig // ProfileVolumeConfig
	volumeConfigInterface, volumeConfigIsSet := d.GetOk("volume_config")
	if volumeConfigIsSet && volumeConfigInterface != nil {
		volumeConfigMap := volumeConfigInterface.([]interface{})
		if len(volumeConfigMap) > 0 {
			volumeConfig = ProfileVolumeConfigModelFromMap(volumeConfigMap[0].(map[string]interface{}))
		}
	}
	return &models.ProfileVolumePolicy{
		MetaData:     metaData,
		VolumeConfig: volumeConfig,
	}
}

func ProfileVolumePolicyModelFromMap(m map[string]interface{}) *models.ProfileVolumePolicy {
	var metaData *models.ProfilePolicyCommon // ProfilePolicyCommon
	metaDataInterface, metaDataIsSet := m["meta_data"]
	if metaDataIsSet && metaDataInterface != nil {
		metaDataMap := metaDataInterface.([]interface{})
		if len(metaDataMap) > 0 {
			metaData = ProfilePolicyCommonModelFromMap(metaDataMap[0].(map[string]interface{}))
		}
	}
	//
	var volumeConfig *models.ProfileVolumeConfig // ProfileVolumeConfig
	volumeConfigInterface, volumeConfigIsSet := m["volume_config"]
	if volumeConfigIsSet && volumeConfigInterface != nil {
		volumeConfigMap := volumeConfigInterface.([]interface{})
		if len(volumeConfigMap) > 0 {
			volumeConfig = ProfileVolumeConfigModelFromMap(volumeConfigMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.ProfileVolumePolicy{
		MetaData:     metaData,
		VolumeConfig: volumeConfig,
	}
}

func SetProfileVolumePolicyResourceData(d *schema.ResourceData, m *models.ProfileVolumePolicy) {
	d.Set("meta_data", SetProfilePolicyCommonSubResourceData([]*models.ProfilePolicyCommon{m.MetaData}))
	d.Set("volume_config", SetProfileVolumeConfigSubResourceData([]*models.ProfileVolumeConfig{m.VolumeConfig}))
}

func SetProfileVolumePolicySubResourceData(m []*models.ProfileVolumePolicy) (d []*map[string]interface{}) {
	for _, ProfileVolumePolicyModel := range m {
		if ProfileVolumePolicyModel != nil {
			properties := make(map[string]interface{})
			properties["meta_data"] = SetProfilePolicyCommonSubResourceData([]*models.ProfilePolicyCommon{ProfileVolumePolicyModel.MetaData})
			properties["volume_config"] = SetProfileVolumeConfigSubResourceData([]*models.ProfileVolumeConfig{ProfileVolumePolicyModel.VolumeConfig})
			d = append(d, &properties)
		}
	}
	return
}

func ProfileVolumePolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"meta_data": {
			Description: `all the required metadata for a policy like id, name, different types of tags`,
			Type:        schema.TypeList, //GoType: ProfilePolicyCommon
			Elem: &schema.Resource{
				Schema: ProfilePolicyCommonSchema(),
			},
			Required: true,
		},

		"volume_config": {
			Description: `volume instance config details`,
			Type:        schema.TypeList, //GoType: ProfileVolumeConfig
			Elem: &schema.Resource{
				Schema: ProfileVolumeConfigSchema(),
			},
			Required: true,
		},
	}
}

func GetProfileVolumePolicyPropertyFields() (t []string) {
	return []string{
		"meta_data",
		"volume_config",
	}
}
