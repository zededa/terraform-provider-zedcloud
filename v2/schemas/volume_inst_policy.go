package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func VolumeInstPolicyModel(d *schema.ResourceData) *models.VolumeInstPolicy {
	var metaData *models.PolicyCommon // PolicyCommon
	metaDataInterface, metaDataIsSet := d.GetOk("meta_data")
	if metaDataIsSet && metaDataInterface != nil {
		metaDataMap := metaDataInterface.([]interface{})
		if len(metaDataMap) > 0 {
			metaData = PolicyCommonModelFromMap(metaDataMap[0].(map[string]interface{}))
		}
	}
	var volInstConfig *models.VolumeInstConfig // VolumeInstConfig
	volInstConfigInterface, volInstConfigIsSet := d.GetOk("vol_inst_config")
	if volInstConfigIsSet && volInstConfigInterface != nil {
		volInstConfigMap := volInstConfigInterface.([]interface{})
		if len(volInstConfigMap) > 0 {
			volInstConfig = VolumeInstConfigModelFromMap(volInstConfigMap[0].(map[string]interface{}))
		}
	}
	return &models.VolumeInstPolicy{
		MetaData:      metaData,
		VolInstConfig: volInstConfig,
	}
}

func VolumeInstPolicyModelFromMap(m map[string]interface{}) *models.VolumeInstPolicy {
	var metaData *models.PolicyCommon // PolicyCommon
	metaDataInterface, metaDataIsSet := m["meta_data"]
	if metaDataIsSet && metaDataInterface != nil {
		metaDataMap := metaDataInterface.([]interface{})
		if len(metaDataMap) > 0 {
			metaData = PolicyCommonModelFromMap(metaDataMap[0].(map[string]interface{}))
		}
	}
	//
	var volInstConfig *models.VolumeInstConfig // VolumeInstConfig
	volInstConfigInterface, volInstConfigIsSet := m["vol_inst_config"]
	if volInstConfigIsSet && volInstConfigInterface != nil {
		volInstConfigMap := volInstConfigInterface.([]interface{})
		if len(volInstConfigMap) > 0 {
			volInstConfig = VolumeInstConfigModelFromMap(volInstConfigMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.VolumeInstPolicy{
		MetaData:      metaData,
		VolInstConfig: volInstConfig,
	}
}

func SetVolumeInstPolicyResourceData(d *schema.ResourceData, m *models.VolumeInstPolicy) {
	d.Set("meta_data", SetPolicyCommonSubResourceData([]*models.PolicyCommon{m.MetaData}))
	d.Set("vol_inst_config", SetVolumeInstConfigSubResourceData([]*models.VolumeInstConfig{m.VolInstConfig}))
}

func SetVolumeInstPolicySubResourceData(m []*models.VolumeInstPolicy) (d []*map[string]interface{}) {
	for _, VolumeInstPolicyModel := range m {
		if VolumeInstPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["meta_data"] = SetPolicyCommonSubResourceData([]*models.PolicyCommon{VolumeInstPolicyModel.MetaData})
			properties["vol_inst_config"] = SetVolumeInstConfigSubResourceData([]*models.VolumeInstConfig{VolumeInstPolicyModel.VolInstConfig})
			d = append(d, &properties)
		}
	}
	return
}

func VolumeInstPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"meta_data": {
			Description: `all the required metadata for a policy like id, name, different types of tags`,
			Type:        schema.TypeList, //GoType: PolicyCommon
			Elem: &schema.Resource{
				Schema: PolicyCommon(),
			},
			Optional: true,
		},

		"vol_inst_config": {
			Description: `volume instance config details`,
			Type:        schema.TypeList, //GoType: VolumeInstConfig
			Elem: &schema.Resource{
				Schema: VolumeInstConfig(),
			},
			Optional: true,
		},
	}
}

func GetVolumeInstPolicyPropertyFields() (t []string) {
	return []string{
		"meta_data",
		"vol_inst_config",
	}
}
