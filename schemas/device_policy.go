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
	return &models.DevicePolicy{
		AttestationPolicy: attestationPolicy,
		MetaData:          metaData,
		PolicySubType:     policySubType,
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
	return &models.DevicePolicy{
		AttestationPolicy: attestationPolicy,
		MetaData:          metaData,
		PolicySubType:     policySubType,
	}
}

func SetDevicePolicyResourceData(d *schema.ResourceData, m *models.DevicePolicy) {
	d.Set("attestation_policy", SetDeviceAttestationPolicySubResourceData([]*models.DeviceAttestationPolicy{m.AttestationPolicy}))
	d.Set("meta_data", SetPolicyCommonSubResourceData([]*models.PolicyCommon{m.MetaData}))
	d.Set("policy_sub_type", m.PolicySubType)
}

func SetDevicePolicySubResourceData(m []*models.DevicePolicy) (d []*map[string]interface{}) {
	for _, DevicePolicyModel := range m {
		if DevicePolicyModel != nil {
			properties := make(map[string]interface{})
			properties["attestation_policy"] = SetDeviceAttestationPolicySubResourceData([]*models.DeviceAttestationPolicy{DevicePolicyModel.AttestationPolicy})
			properties["meta_data"] = SetPolicyCommonSubResourceData([]*models.PolicyCommon{DevicePolicyModel.MetaData})
			properties["policy_sub_type"] = DevicePolicyModel.PolicySubType
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

		"policy_sub_type": {
			Description: `device policy type`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetDevicePolicyPropertyFields() (t []string) {
	return []string{
		"attestation_policy",
		"meta_data",
		"policy_sub_type",
	}
}
