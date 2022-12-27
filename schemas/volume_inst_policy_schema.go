package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate VolumeInstPolicy resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func VolumeInstPolicyModel(d *schema.ResourceData) *models.VolumeInstPolicy {
	var metaData *models.PolicyCommon // PolicyCommon
	metaDataInterface, metaDataIsSet := d.GetOk("meta_data")
	if metaDataIsSet {
		metaDataMap := metaDataInterface.([]interface{})[0].(map[string]interface{})
		metaData = PolicyCommonModelFromMap(metaDataMap)
	}
	var volInstConfig *models.VolumeInstConfig // VolumeInstConfig
	volInstConfigInterface, volInstConfigIsSet := d.GetOk("vol_inst_config")
	if volInstConfigIsSet {
		volInstConfigMap := volInstConfigInterface.([]interface{})[0].(map[string]interface{})
		volInstConfig = VolumeInstConfigModelFromMap(volInstConfigMap)
	}
	return &models.VolumeInstPolicy{
		MetaData:      metaData,
		VolInstConfig: volInstConfig,
	}
}

func VolumeInstPolicyModelFromMap(m map[string]interface{}) *models.VolumeInstPolicy {
	var metaData *models.PolicyCommon // PolicyCommon
	metaDataInterface, metaDataIsSet := m["meta_data"]
	if metaDataIsSet {
		metaDataMap := metaDataInterface.([]interface{})[0].(map[string]interface{})
		metaData = PolicyCommonModelFromMap(metaDataMap)
	}
	//
	var volInstConfig *models.VolumeInstConfig // VolumeInstConfig
	volInstConfigInterface, volInstConfigIsSet := m["vol_inst_config"]
	if volInstConfigIsSet {
		volInstConfigMap := volInstConfigInterface.([]interface{})[0].(map[string]interface{})
		volInstConfig = VolumeInstConfigModelFromMap(volInstConfigMap)
	}
	//
	return &models.VolumeInstPolicy{
		MetaData:      metaData,
		VolInstConfig: volInstConfig,
	}
}

// Update the underlying VolumeInstPolicy resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVolumeInstPolicyResourceData(d *schema.ResourceData, m *models.VolumeInstPolicy) {
	d.Set("meta_data", SetPolicyCommonSubResourceData([]*models.PolicyCommon{m.MetaData}))
	d.Set("vol_inst_config", SetVolumeInstConfigSubResourceData([]*models.VolumeInstConfig{m.VolInstConfig}))
}

// Iterate throught and update the VolumeInstPolicy resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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

// Schema mapping representing the VolumeInstPolicy resource defined in the Terraform configuration
func VolumeInstPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"meta_data": {
			Description: `all the required metadata for a policy like id, name, different types of tags`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"vol_inst_config": {
			Description: `volume instance config details`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the VolumeInstPolicy resource
func GetVolumeInstPolicyPropertyFields() (t []string) {
	return []string{
		"meta_data",
		"vol_inst_config",
	}
}
