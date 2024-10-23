package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func TagLevelSettingsModel(d *schema.ResourceData) *models.TagLevelSettings {
	var flowLogTransmission *models.NetworkInstanceFlowLogTransmission // NetworkInstanceFlowLogTransmission
	flowLogTransmissionInterface, flowLogTransmissionIsSet := d.GetOk("flow_log_transmission")
	if flowLogTransmissionIsSet {
		flowLogTransmissionModel := flowLogTransmissionInterface.(string)
		flowLogTransmission = models.NewNetworkInstanceFlowLogTransmission(models.NetworkInstanceFlowLogTransmission(flowLogTransmissionModel))
	}
	return &models.TagLevelSettings{
		FlowLogTransmission: flowLogTransmission,
	}
}

func TagLevelSettingsModelFromMap(m map[string]interface{}) *models.TagLevelSettings {
	var flowLogTransmission *models.NetworkInstanceFlowLogTransmission // NetworkInstanceFlowLogTransmission
	flowLogTransmissionInterface, flowLogTransmissionIsSet := m["flow_log_transmission"]
	if flowLogTransmissionIsSet {
		flowLogTransmissionModel := flowLogTransmissionInterface.(string)
		flowLogTransmission = models.NewNetworkInstanceFlowLogTransmission(models.NetworkInstanceFlowLogTransmission(flowLogTransmissionModel))
	}
	return &models.TagLevelSettings{
		FlowLogTransmission: flowLogTransmission,
	}
}

func SetTagLevelSettingsResourceData(d *schema.ResourceData, m *models.TagLevelSettings) {
	d.Set("flow_log_transmission", m.FlowLogTransmission)
}

func SetTagLevelSettingsSubResourceData(m []*models.TagLevelSettings) (d []*map[string]interface{}) {
	for _, TagLevelSettingsModel := range m {
		if TagLevelSettingsModel != nil {
			properties := make(map[string]interface{})
			properties["flow_log_transmission"] = TagLevelSettingsModel.FlowLogTransmission
			d = append(d, &properties)
		}
	}
	return
}

func TagLevelSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"flow_log_transmission": {
			Description: `Flow log transmission setting for the network instances`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetTagLevelSettingsPropertyFields() (t []string) {
	return []string{
		"flow_log_transmission",
	}
}
