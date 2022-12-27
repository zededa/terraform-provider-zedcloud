package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DevicePolicy resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DevicePolicyModel(d *schema.ResourceData) *models.DevicePolicy {
	var attestationPolicy *models.DeviceAttestationPolicy // DeviceAttestationPolicy
	attestationPolicyInterface, attestationPolicyIsSet := d.GetOk("attestation_policy")
	if attestationPolicyIsSet {
		attestationPolicyMap := attestationPolicyInterface.([]interface{})[0].(map[string]interface{})
		attestationPolicy = DeviceAttestationPolicyModelFromMap(attestationPolicyMap)
	}
	var metaData *models.PolicyCommon // PolicyCommon
	metaDataInterface, metaDataIsSet := d.GetOk("meta_data")
	if metaDataIsSet {
		metaDataMap := metaDataInterface.([]interface{})[0].(map[string]interface{})
		metaData = PolicyCommonModelFromMap(metaDataMap)
	}
	policySubType := d.Get("policy_sub_type").(*models.DevicePolicyType) // DevicePolicyType
	return &models.DevicePolicy{
		AttestationPolicy: attestationPolicy,
		MetaData:          metaData,
		PolicySubType:     policySubType,
	}
}

func DevicePolicyModelFromMap(m map[string]interface{}) *models.DevicePolicy {
	var attestationPolicy *models.DeviceAttestationPolicy // DeviceAttestationPolicy
	attestationPolicyInterface, attestationPolicyIsSet := m["attestation_policy"]
	if attestationPolicyIsSet {
		attestationPolicyMap := attestationPolicyInterface.([]interface{})[0].(map[string]interface{})
		attestationPolicy = DeviceAttestationPolicyModelFromMap(attestationPolicyMap)
	}
	//
	var metaData *models.PolicyCommon // PolicyCommon
	metaDataInterface, metaDataIsSet := m["meta_data"]
	if metaDataIsSet {
		metaDataMap := metaDataInterface.([]interface{})[0].(map[string]interface{})
		metaData = PolicyCommonModelFromMap(metaDataMap)
	}
	//
	policySubType := m["policy_sub_type"].(*models.DevicePolicyType) // DevicePolicyType
	return &models.DevicePolicy{
		AttestationPolicy: attestationPolicy,
		MetaData:          metaData,
		PolicySubType:     policySubType,
	}
}

// Update the underlying DevicePolicy resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDevicePolicyResourceData(d *schema.ResourceData, m *models.DevicePolicy) {
	d.Set("attestation_policy", SetDeviceAttestationPolicySubResourceData([]*models.DeviceAttestationPolicy{m.AttestationPolicy}))
	d.Set("meta_data", SetPolicyCommonSubResourceData([]*models.PolicyCommon{m.MetaData}))
	d.Set("policy_sub_type", m.PolicySubType)
}

// Iterate throught and update the DevicePolicy resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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

// Schema mapping representing the DevicePolicy resource defined in the Terraform configuration
func DevicePolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attestation_policy": {
			Description: `device attestation policy`,
			Type:        schema.TypeList, //GoType: DeviceAttestationPolicy
			Elem: &schema.Resource{
				Schema: DeviceAttestationPolicySchema(),
			},
			Optional: true,
		},

		"meta_data": {
			Description: `all the required metadata for a policy like id, name, different types of tags`,
			Type:        schema.TypeList, //GoType: PolicyCommon
			Elem: &schema.Resource{
				Schema: PolicyCommonSchema(),
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

// Retrieve property field names for updating the DevicePolicy resource
func GetDevicePolicyPropertyFields() (t []string) {
	return []string{
		"attestation_policy",
		"meta_data",
		"policy_sub_type",
	}
}
