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
	var interfaceOrdering *models.InterfaceOrdering // InterfaceOrdering
	interfaceOrderingInterface, interfaceOrderingIsSet := d.GetOk("interface_ordering")
	if interfaceOrderingIsSet {
		interfaceOrderingModel := interfaceOrderingInterface.(string)
		interfaceOrdering = models.NewInterfaceOrdering(models.InterfaceOrdering(interfaceOrderingModel))
	}
	return &models.TagLevelSettings{
		FlowLogTransmission: flowLogTransmission,
		InterfaceOrdering:   interfaceOrdering,
	}
}

func TagLevelSettingsModelFromMap(m map[string]interface{}) *models.TagLevelSettings {
	var flowLogTransmission *models.NetworkInstanceFlowLogTransmission // NetworkInstanceFlowLogTransmission
	flowLogTransmissionInterface, flowLogTransmissionIsSet := m["flow_log_transmission"]
	if flowLogTransmissionIsSet {
		flowLogTransmissionModel := flowLogTransmissionInterface.(string)
		flowLogTransmission = models.NewNetworkInstanceFlowLogTransmission(models.NetworkInstanceFlowLogTransmission(flowLogTransmissionModel))
	}
	var interfaceOrdering *models.InterfaceOrdering // InterfaceOrdering
	interfaceOrderingInterface, interfaceOrderingIsSet := m["interface_ordering"]
	if interfaceOrderingIsSet {
		interfaceOrderingModel := interfaceOrderingInterface.(string)
		interfaceOrdering = models.NewInterfaceOrdering(models.InterfaceOrdering(interfaceOrderingModel))
	}
	return &models.TagLevelSettings{
		FlowLogTransmission: flowLogTransmission,
		InterfaceOrdering:   interfaceOrdering,
	}
}

func SetTagLevelSettingsResourceData(d *schema.ResourceData, m *models.TagLevelSettings) {
	d.Set("flow_log_transmission", m.FlowLogTransmission)
	d.Set("interface_ordering", m.InterfaceOrdering)
}

func SetTagLevelSettingsSubResourceData(m []*models.TagLevelSettings) (d []*map[string]interface{}) {
	for _, TagLevelSettingsModel := range m {
		if TagLevelSettingsModel != nil {
			properties := make(map[string]interface{})
			properties["flow_log_transmission"] = TagLevelSettingsModel.FlowLogTransmission
			properties["interface_ordering"] = TagLevelSettingsModel.InterfaceOrdering
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

		"interface_ordering": {
			Description: `interface ordering for app instances`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetTagLevelSettingsPropertyFields() (t []string) {
	return []string{
		"flow_log_transmission",
		"interface_ordering",
	}
}
