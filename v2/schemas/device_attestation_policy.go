package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func DeviceAttestationPolicyModel(d *schema.ResourceData) *models.DeviceAttestationPolicy {
	var typeVar *models.DeviceAttestPolicyType // DeviceAttestPolicyType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewDeviceAttestPolicyType(models.DeviceAttestPolicyType(typeModel))
	}
	return &models.DeviceAttestationPolicy{
		Type: typeVar,
	}
}

func DeviceAttestationPolicyModelFromMap(m map[string]interface{}) *models.DeviceAttestationPolicy {
	var typeVar *models.DeviceAttestPolicyType // DeviceAttestPolicyType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewDeviceAttestPolicyType(models.DeviceAttestPolicyType(typeModel))
	}
	return &models.DeviceAttestationPolicy{
		Type: typeVar,
	}
}

func SetDeviceAttestationPolicyResourceData(d *schema.ResourceData, m *models.DeviceAttestationPolicy) {
	d.Set("type", m.Type)
}

func SetDeviceAttestationPolicySubResourceData(m []*models.DeviceAttestationPolicy) (d []*map[string]interface{}) {
	for _, DeviceAttestationPolicyModel := range m {
		if DeviceAttestationPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["type"] = DeviceAttestationPolicyModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func DeviceAttestationPolicy() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type": {
			Description: `Attestation policy type. Type values: 
	- DEVICE_ATTEST_POLICY_TYPE_ACCEPT: Do not enforce attestation. All devices are marked as successfully attested.
	- DEVICE_ATTEST_POLICY_TYPE_ENFORCE: Enforce attestation. Devices failing attestation are marked accordingly.`,
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func GetDeviceAttestationPolicyPropertyFields() (t []string) {
	return []string{
		"type",
	}
}
