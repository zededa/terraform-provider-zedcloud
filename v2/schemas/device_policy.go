package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func DevicePolicyModel(d *schema.ResourceData) *models.DevicePolicy {
	var attestationPolicy *models.DeviceAttestationPolicy // DeviceAttestationPolicy
	attestationPolicyInterface, attestationPolicyIsSet := d.GetOk("attestation_policy")
	if attestationPolicyIsSet && attestationPolicyInterface != nil {
		attestationPolicyMap := attestationPolicyInterface.([]interface{})
		if len(attestationPolicyMap) > 0 {
			attestationPolicy = DeviceAttestationPolicyModelFromMap(attestationPolicyMap[0].(map[string]interface{}))
		}
	}
	var metaData *models.PolicyCommon // PolicyCommon
	metaDataInterface, metaDataIsSet := d.GetOk("meta_data")
	if metaDataIsSet && metaDataInterface != nil {
		metaDataMap := metaDataInterface.([]interface{})
		if len(metaDataMap) > 0 {
			metaData = PolicyCommonModelFromMap(metaDataMap[0].(map[string]interface{}))
		}
	}
	var policySubType *models.DevicePolicyType // DevicePolicyType
	policySubTypeInterface, policySubTypeIsSet := d.GetOk("policy_sub_type")
	if policySubTypeIsSet {
		policySubTypeModel := policySubTypeInterface.(string)
		policySubType = models.NewDevicePolicyType(models.DevicePolicyType(policySubTypeModel))
	}
	var portBasedNetworkAccessControlPolicy *models.DevicePortBasedNetworkAccessControlPolicy // DevicePortBasedNetworkAccessControlPolicy
	portBasedNetworkAccessControlPolicyInterface, portBasedNetworkAccessControlPolicyIsSet := d.GetOk("port_based_network_access_control_policy")
	if portBasedNetworkAccessControlPolicyIsSet && portBasedNetworkAccessControlPolicyInterface != nil {
		portBasedNetworkAccessControlPolicyMap := portBasedNetworkAccessControlPolicyInterface.([]interface{})
		if len(portBasedNetworkAccessControlPolicyMap) > 0 {
			portBasedNetworkAccessControlPolicy = DevicePortBasedNetworkAccessControlPolicyModelFromMap(portBasedNetworkAccessControlPolicyMap[0].(map[string]interface{}))
		}
	}
	var pcrPolicy *models.DevicePCRPolicy // DevicePCRPolicy
	if l, ok := d.Get("pcr_policy").([]interface{}); ok && len(l) > 0 && l[0] != nil {
		pcrPolicy = DevicePCRPolicyModelFromMap(l[0].(map[string]interface{}))
	}
	return &models.DevicePolicy{
		PcrPolicy:                           pcrPolicy,
		AttestationPolicy:                   attestationPolicy,
		MetaData:                            metaData,
		PolicySubType:                       policySubType,
		PortBasedNetworkAccessControlPolicy: portBasedNetworkAccessControlPolicy,
	}
}

func DevicePolicyModelFromMap(m map[string]interface{}) *models.DevicePolicy {
	var attestationPolicy *models.DeviceAttestationPolicy // DeviceAttestationPolicy
	attestationPolicyInterface, attestationPolicyIsSet := m["attestation_policy"]
	if attestationPolicyIsSet && attestationPolicyInterface != nil {
		attestationPolicyMap := attestationPolicyInterface.([]interface{})
		if len(attestationPolicyMap) > 0 {
			attestationPolicy = DeviceAttestationPolicyModelFromMap(attestationPolicyMap[0].(map[string]interface{}))
		}
	}
	//
	var metaData *models.PolicyCommon // PolicyCommon
	metaDataInterface, metaDataIsSet := m["meta_data"]
	if metaDataIsSet && metaDataInterface != nil {
		metaDataMap := metaDataInterface.([]interface{})
		if len(metaDataMap) > 0 {
			metaData = PolicyCommonModelFromMap(metaDataMap[0].(map[string]interface{}))
		}
	}
	//
	var policySubType *models.DevicePolicyType // DevicePolicyType
	policySubTypeInterface, policySubTypeIsSet := m["policy_sub_type"]
	if policySubTypeIsSet {
		policySubTypeModel := policySubTypeInterface.(string)
		policySubType = models.NewDevicePolicyType(models.DevicePolicyType(policySubTypeModel))
	}
	var portBasedNetworkAccessControlPolicy *models.DevicePortBasedNetworkAccessControlPolicy // DevicePortBasedNetworkAccessControlPolicy
	portBasedNetworkAccessControlPolicyInterface, portBasedNetworkAccessControlPolicyIsSet := m["port_based_network_access_control_policy"]
	if portBasedNetworkAccessControlPolicyIsSet && portBasedNetworkAccessControlPolicyInterface != nil {
		portBasedNetworkAccessControlPolicyMap := portBasedNetworkAccessControlPolicyInterface.([]interface{})
		if len(portBasedNetworkAccessControlPolicyMap) > 0 {
			portBasedNetworkAccessControlPolicy = DevicePortBasedNetworkAccessControlPolicyModelFromMap(portBasedNetworkAccessControlPolicyMap[0].(map[string]interface{}))
		}
	}
	var pcrPolicy *models.DevicePCRPolicy // DevicePCRPolicy
	if l, ok := m["pcr_policy"].([]interface{}); ok && len(l) > 0 && l[0] != nil {
		pcrPolicy = DevicePCRPolicyModelFromMap(l[0].(map[string]interface{}))
	}
	return &models.DevicePolicy{
		PcrPolicy:                           pcrPolicy,
		AttestationPolicy:                   attestationPolicy,
		MetaData:                            metaData,
		PolicySubType:                       policySubType,
		PortBasedNetworkAccessControlPolicy: portBasedNetworkAccessControlPolicy,
	}
}

func SetDevicePolicyResourceData(d *schema.ResourceData, m *models.DevicePolicy) {
	d.Set("attestation_policy", SetDeviceAttestationPolicySubResourceData([]*models.DeviceAttestationPolicy{m.AttestationPolicy}))
	d.Set("meta_data", SetPolicyCommonSubResourceData([]*models.PolicyCommon{m.MetaData}))
	d.Set("policy_sub_type", m.PolicySubType)
	d.Set("pcr_policy", SetDevicePCRPolicySubResourceData([]*models.DevicePCRPolicy{m.PcrPolicy}))
	d.Set("port_based_network_access_control_policy", SetDevicePortBasedNetworkAccessControlPolicySubResourceData([]*models.DevicePortBasedNetworkAccessControlPolicy{m.PortBasedNetworkAccessControlPolicy}))
}

func SetDevicePolicySubResourceData(m []*models.DevicePolicy) (d []*map[string]interface{}) {
	for _, DevicePolicyModel := range m {
		if DevicePolicyModel != nil {
			properties := make(map[string]interface{})
			properties["attestation_policy"] = SetDeviceAttestationPolicySubResourceData([]*models.DeviceAttestationPolicy{DevicePolicyModel.AttestationPolicy})
			properties["meta_data"] = SetPolicyCommonSubResourceData([]*models.PolicyCommon{DevicePolicyModel.MetaData})
			properties["policy_sub_type"] = DevicePolicyModel.PolicySubType
			properties["pcr_policy"] = SetDevicePCRPolicySubResourceData([]*models.DevicePCRPolicy{DevicePolicyModel.PcrPolicy})
			properties["port_based_network_access_control_policy"] = SetDevicePortBasedNetworkAccessControlPolicySubResourceData([]*models.DevicePortBasedNetworkAccessControlPolicy{DevicePolicyModel.PortBasedNetworkAccessControlPolicy})
			d = append(d, &properties)
		}
	}
	return
}

func DevicePolicy() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attestation_policy": {
			Description: `device attestation policy`,
			Type:        schema.TypeList, //GoType: DeviceAttestationPolicy
			Elem: &schema.Resource{
				Schema: DeviceAttestationPolicy(),
			},
			Optional: true,
		},

		"meta_data": {
			Description: `all the required metadata for a policy like id, name, different types of tags`,
			Type:        schema.TypeList, //GoType: PolicyCommon
			Elem: &schema.Resource{
				Schema: PolicyCommon(),
			},
			Optional: true,
		},

		"pcr_policy": {
			Description: `PCR policy for TPM volume key sealing on matching devices (set with policy_sub_type = "DEVICE_POLICY_TYPE_PCR"). ` +
				`For nodes in the deployment's project it takes precedence over a node-level policy; the project-level policy is not used for them.`,
			Type: schema.TypeList, //GoType: DevicePCRPolicy
			Elem: &schema.Resource{
				Schema: DevicePCRPolicySchema(),
			},
			MaxItems: 1,
			Optional: true,
		},

		"policy_sub_type": {
			Description: `device policy type`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"port_based_network_access_control_policy": {
			Description: `port based network access control (802.1X) policy to enforce on devices of this project`,
			Type:        schema.TypeList, //GoType: DevicePortBasedNetworkAccessControlPolicy
			Elem: &schema.Resource{
				Schema: DevicePortBasedNetworkAccessControlPolicySchema(),
			},
			MaxItems: 1,
			Optional: true,
		},
	}
}

func GetDevicePolicyPropertyFields() (t []string) {
	return []string{
		"attestation_policy",
		"meta_data",
		"pcr_policy",
		"policy_sub_type",
		"port_based_network_access_control_policy",
	}
}
