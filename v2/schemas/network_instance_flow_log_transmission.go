package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func NetworkInstanceFlowLogTransmissionModel(d *schema.ResourceData) *models.NetworkInstanceFlowLogTransmission {
	networkInstanceFlowLogTransmission, _ := d.Get("network_instance_flow_log_transmission").(models.NetworkInstanceFlowLogTransmission)
	return &networkInstanceFlowLogTransmission
}

func NetworkInstanceFlowLogTransmissionModelFromMap(m map[string]interface{}) *models.NetworkInstanceFlowLogTransmission {
	networkInstanceFlowLogTransmission := m["network_instance_flow_log_transmission"].(models.NetworkInstanceFlowLogTransmission)
	return &networkInstanceFlowLogTransmission
}

func SetNetworkInstanceFlowLogTransmissionResourceData(d *schema.ResourceData, m *models.NetworkInstanceFlowLogTransmission) {
}

func SetNetworkInstanceFlowLogTransmissionSubResourceData(m []*models.NetworkInstanceFlowLogTransmission) (d []*map[string]interface{}) {
	for _, NetworkInstanceFlowLogTransmissionModel := range m {
		if NetworkInstanceFlowLogTransmissionModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func NetworkInstanceFlowLogTransmissionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetNetworkInstanceFlowLogTransmissionPropertyFields() (t []string) {
	return []string{}
}
