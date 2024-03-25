package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func EnrollmentDetailModel(d *schema.ResourceData) *models.EnrollmentDetail {
	var allocationPolicy *models.AllocationPolicy // AllocationPolicy
	allocationPolicyInterface, allocationPolicyIsSet := d.GetOk("allocation_policy")
	if allocationPolicyIsSet {
		allocationPolicyModel := allocationPolicyInterface.(string)
		allocationPolicy = models.NewAllocationPolicy(models.AllocationPolicy(allocationPolicyModel))
	}
	var attachedIotHubsName []string
	attachedIotHubsNameInterface, attachedIotHubsNameIsSet := d.GetOk("attachedIotHubsName")
	if attachedIotHubsNameIsSet {
		var items []interface{}
		if listItems, isList := attachedIotHubsNameInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = attachedIotHubsNameInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			attachedIotHubsName = append(attachedIotHubsName, v.(string))
		}
	}
	var certificateEnrollment *models.CertificateEnrollmentDetail // CertificateEnrollmentDetail
	certificateEnrollmentInterface, certificateEnrollmentIsSet := d.GetOk("certificate_enrollment")
	if certificateEnrollmentIsSet && certificateEnrollmentInterface != nil {
		certificateEnrollmentMap := certificateEnrollmentInterface.([]interface{})
		if len(certificateEnrollmentMap) > 0 {
			certificateEnrollment = CertificateEnrollmentDetailModelFromMap(certificateEnrollmentMap[0].(map[string]interface{}))
		}
	}
	enableIotEdgeDevice, _ := d.Get("enable_iot_edge_device").(bool)
	var mechanism *models.EnrollmentMechanism // EnrollmentMechanism
	mechanismInterface, mechanismIsSet := d.GetOk("mechanism")
	if mechanismIsSet {
		mechanismModel := mechanismInterface.(string)
		mechanism = models.NewEnrollmentMechanism(models.EnrollmentMechanism(mechanismModel))
	}
	var symmetricKeyEnrollment *models.SymmetricKeyEnrollmentDetail // SymmetricKeyEnrollmentDetail
	symmetricKeyEnrollmentInterface, symmetricKeyEnrollmentIsSet := d.GetOk("symmetric_key_enrollment")
	if symmetricKeyEnrollmentIsSet && symmetricKeyEnrollmentInterface != nil {
		symmetricKeyEnrollmentMap := symmetricKeyEnrollmentInterface.([]interface{})
		if len(symmetricKeyEnrollmentMap) > 0 {
			symmetricKeyEnrollment = SymmetricKeyEnrollmentDetailModelFromMap(symmetricKeyEnrollmentMap[0].(map[string]interface{}))
		}
	}
	tags := map[string]string{}
	tagsInterface, tagsIsSet := d.GetOk("tags")
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	var tpmEnrollment *models.TPMEnrollmentDetail // TPMEnrollmentDetail
	tpmEnrollmentInterface, tpmEnrollmentIsSet := d.GetOk("tpm_enrollment")
	if tpmEnrollmentIsSet && tpmEnrollmentInterface != nil {
		tpmEnrollmentMap := tpmEnrollmentInterface.([]interface{})
		if len(tpmEnrollmentMap) > 0 {
			tpmEnrollment = TPMEnrollmentDetailModelFromMap(tpmEnrollmentMap[0].(map[string]interface{}))
		}
	}
	return &models.EnrollmentDetail{
		AllocationPolicy:       allocationPolicy,
		AttachedIotHubsName:    attachedIotHubsName,
		CertificateEnrollment:  *certificateEnrollment,
		EnableIotEdgeDevice:    enableIotEdgeDevice,
		Mechanism:              mechanism,
		SymmetricKeyEnrollment: symmetricKeyEnrollment,
		Tags:                   tags,
		TpmEnrollment:          tpmEnrollment,
	}
}

func EnrollmentDetailModelFromMap(m map[string]interface{}) *models.EnrollmentDetail {
	var allocationPolicy *models.AllocationPolicy // AllocationPolicy
	allocationPolicyInterface, allocationPolicyIsSet := m["allocation_policy"]
	if allocationPolicyIsSet {
		allocationPolicyModel := allocationPolicyInterface.(string)
		allocationPolicy = models.NewAllocationPolicy(models.AllocationPolicy(allocationPolicyModel))
	}
	var attachedIotHubsName []string
	attachedIotHubsNameInterface, attachedIotHubsNameIsSet := m["attachedIotHubsName"]
	if attachedIotHubsNameIsSet {
		var items []interface{}
		if listItems, isList := attachedIotHubsNameInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = attachedIotHubsNameInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			attachedIotHubsName = append(attachedIotHubsName, v.(string))
		}
	}
	var certificateEnrollment *models.CertificateEnrollmentDetail // CertificateEnrollmentDetail
	certificateEnrollmentInterface, certificateEnrollmentIsSet := m["certificate_enrollment"]
	if certificateEnrollmentIsSet && certificateEnrollmentInterface != nil {
		certificateEnrollmentMap := certificateEnrollmentInterface.([]interface{})
		if len(certificateEnrollmentMap) > 0 {
			certificateEnrollment = CertificateEnrollmentDetailModelFromMap(certificateEnrollmentMap[0].(map[string]interface{}))
		}
	}
	//
	enableIotEdgeDevice := m["enable_iot_edge_device"].(bool)
	var mechanism *models.EnrollmentMechanism // EnrollmentMechanism
	mechanismInterface, mechanismIsSet := m["mechanism"]
	if mechanismIsSet {
		mechanismModel := mechanismInterface.(string)
		mechanism = models.NewEnrollmentMechanism(models.EnrollmentMechanism(mechanismModel))
	}
	var symmetricKeyEnrollment *models.SymmetricKeyEnrollmentDetail // SymmetricKeyEnrollmentDetail
	symmetricKeyEnrollmentInterface, symmetricKeyEnrollmentIsSet := m["symmetric_key_enrollment"]
	if symmetricKeyEnrollmentIsSet && symmetricKeyEnrollmentInterface != nil {
		symmetricKeyEnrollmentMap := symmetricKeyEnrollmentInterface.([]interface{})
		if len(symmetricKeyEnrollmentMap) > 0 {
			symmetricKeyEnrollment = SymmetricKeyEnrollmentDetailModelFromMap(symmetricKeyEnrollmentMap[0].(map[string]interface{}))
		}
	}
	//
	tags := map[string]string{}
	tagsInterface, tagsIsSet := m["tags"]
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	var tpmEnrollment *models.TPMEnrollmentDetail // TPMEnrollmentDetail
	tpmEnrollmentInterface, tpmEnrollmentIsSet := m["tpm_enrollment"]
	if tpmEnrollmentIsSet && tpmEnrollmentInterface != nil {
		tpmEnrollmentMap := tpmEnrollmentInterface.([]interface{})
		if len(tpmEnrollmentMap) > 0 {
			tpmEnrollment = TPMEnrollmentDetailModelFromMap(tpmEnrollmentMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.EnrollmentDetail{
		AllocationPolicy:       allocationPolicy,
		AttachedIotHubsName:    attachedIotHubsName,
		CertificateEnrollment:  *certificateEnrollment,
		EnableIotEdgeDevice:    enableIotEdgeDevice,
		Mechanism:              mechanism,
		SymmetricKeyEnrollment: symmetricKeyEnrollment,
		Tags:                   tags,
		TpmEnrollment:          tpmEnrollment,
	}
}

func SetEnrollmentDetailResourceData(d *schema.ResourceData, m *models.EnrollmentDetail) {
	d.Set("allocation_policy", m.AllocationPolicy)
	d.Set("attached_iot_hubs_name", m.AttachedIotHubsName)
	d.Set("certificate_enrollment", SetCertificateEnrollmentDetailSubResourceData([]*models.CertificateEnrollmentDetail{&m.CertificateEnrollment}))
	d.Set("enable_iot_edge_device", m.EnableIotEdgeDevice)
	d.Set("mechanism", m.Mechanism)
	d.Set("symmetric_key_enrollment", SetSymmetricKeyEnrollmentDetailSubResourceData([]*models.SymmetricKeyEnrollmentDetail{m.SymmetricKeyEnrollment}))
	d.Set("tags", m.Tags)
	d.Set("tpm_enrollment", SetTPMEnrollmentDetailSubResourceData([]*models.TPMEnrollmentDetail{m.TpmEnrollment}))
}

func SetEnrollmentDetailSubResourceData(m []*models.EnrollmentDetail) (d []*map[string]interface{}) {
	for _, EnrollmentDetailModel := range m {
		if EnrollmentDetailModel != nil {
			properties := make(map[string]interface{})
			properties["allocation_policy"] = EnrollmentDetailModel.AllocationPolicy
			properties["attached_iot_hubs_name"] = EnrollmentDetailModel.AttachedIotHubsName
			properties["certificate_enrollment"] = SetCertificateEnrollmentDetailSubResourceData([]*models.CertificateEnrollmentDetail{&EnrollmentDetailModel.CertificateEnrollment})
			properties["enable_iot_edge_device"] = EnrollmentDetailModel.EnableIotEdgeDevice
			properties["mechanism"] = EnrollmentDetailModel.Mechanism
			properties["symmetric_key_enrollment"] = SetSymmetricKeyEnrollmentDetailSubResourceData([]*models.SymmetricKeyEnrollmentDetail{EnrollmentDetailModel.SymmetricKeyEnrollment})
			properties["tags"] = EnrollmentDetailModel.Tags
			properties["tpm_enrollment"] = SetTPMEnrollmentDetailSubResourceData([]*models.TPMEnrollmentDetail{EnrollmentDetailModel.TpmEnrollment})
			d = append(d, &properties)
		}
	}
	return
}

func EnrollmentDetail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allocation_policy": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"attached_iot_hubs_name": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"certificate_enrollment": {
			Description: ``,
			Type:        schema.TypeList, //GoType: CertificateEnrollmentDetail
			Elem: &schema.Resource{
				Schema: CertificateEnrollmentDetail(),
			},
			Optional: true,
		},

		"enable_iot_edge_device": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"mechanism": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"symmetric_key_enrollment": {
			Description: ``,
			Type:        schema.TypeList, //GoType: SymmetricKeyEnrollmentDetail
			Elem: &schema.Resource{
				Schema: SymmetricKeyEnrollmentDetail(),
			},
			Optional: true,
		},

		"tags": {
			Description: ``,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"tpm_enrollment": {
			Description: ``,
			Type:        schema.TypeList, //GoType: TPMEnrollmentDetail
			Elem: &schema.Resource{
				Schema: TPMEnrollmentDetail(),
			},
			Optional: true,
		},
	}
}

func GetEnrollmentDetailPropertyFields() (t []string) {
	return []string{
		"allocation_policy",
		"attached_iot_hubs_name",
		"certificate_enrollment",
		"enable_iot_edge_device",
		"mechanism",
		"symmetric_key_enrollment",
		"tags",
		"tpm_enrollment",
	}
}
