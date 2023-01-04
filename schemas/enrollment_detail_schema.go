package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate EnrollmentDetail resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EnrollmentDetailModel(d *schema.ResourceData) *models.EnrollmentDetail {
	var allocationPolicy *models.AllocationPolicy // AllocationPolicy
	allocationPolicyInterface, allocationPolicyIsSet := d.GetOk("allocation_policy")
	if allocationPolicyIsSet {
		allocationPolicyModel := allocationPolicyInterface.(string)
		allocationPolicy = models.NewAllocationPolicy(models.AllocationPolicy(allocationPolicyModel))
	}
	attachedIotHubsName, _ := d.Get("attached_iot_hubs_name").([]string)
	certificateEnrollment, _ := d.Get("certificate_enrollment").(models.CertificateEnrollmentDetail) // CertificateEnrollmentDetail
	enableIotEdgeDevice, _ := d.Get("enable_iot_edge_device").(bool)
	var mechanism *models.EnrollmentMechanism // EnrollmentMechanism
	mechanismInterface, mechanismIsSet := d.GetOk("mechanism")
	if mechanismIsSet {
		mechanismModel := mechanismInterface.(string)
		mechanism = models.NewEnrollmentMechanism(models.EnrollmentMechanism(mechanismModel))
	}
	var symmetricKeyEnrollment *models.SymmetricKeyEnrollmentDetail // SymmetricKeyEnrollmentDetail
	symmetricKeyEnrollmentInterface, symmetricKeyEnrollmentIsSet := d.GetOk("symmetric_key_enrollment")
	if symmetricKeyEnrollmentIsSet {
		symmetricKeyEnrollmentMap := symmetricKeyEnrollmentInterface.([]interface{})[0].(map[string]interface{})
		symmetricKeyEnrollment = SymmetricKeyEnrollmentDetailModelFromMap(symmetricKeyEnrollmentMap)
	}
	tags, _ := d.Get("tags").(map[string]string)  // map[string]string
	var tpmEnrollment *models.TPMEnrollmentDetail // TPMEnrollmentDetail
	tpmEnrollmentInterface, tpmEnrollmentIsSet := d.GetOk("tpm_enrollment")
	if tpmEnrollmentIsSet {
		tpmEnrollmentMap := tpmEnrollmentInterface.([]interface{})[0].(map[string]interface{})
		tpmEnrollment = TPMEnrollmentDetailModelFromMap(tpmEnrollmentMap)
	}
	return &models.EnrollmentDetail{
		AllocationPolicy:       allocationPolicy,
		AttachedIotHubsName:    attachedIotHubsName,
		CertificateEnrollment:  certificateEnrollment,
		EnableIotEdgeDevice:    enableIotEdgeDevice,
		Mechanism:              mechanism,
		SymmetricKeyEnrollment: symmetricKeyEnrollment,
		Tags:                   tags,
		TpmEnrollment:          tpmEnrollment,
	}
}

func EnrollmentDetailModelFromMap(m map[string]interface{}) *models.EnrollmentDetail {
	allocationPolicy := m["allocation_policy"].(*models.AllocationPolicy) // AllocationPolicy
	attachedIotHubsName := m["attached_iot_hubs_name"].([]string)
	certificateEnrollment := m["certificate_enrollment"].(models.CertificateEnrollmentDetail) // CertificateEnrollmentDetail
	enableIotEdgeDevice := m["enable_iot_edge_device"].(bool)
	mechanism := m["mechanism"].(*models.EnrollmentMechanism)       // EnrollmentMechanism
	var symmetricKeyEnrollment *models.SymmetricKeyEnrollmentDetail // SymmetricKeyEnrollmentDetail
	symmetricKeyEnrollmentInterface, symmetricKeyEnrollmentIsSet := m["symmetric_key_enrollment"]
	if symmetricKeyEnrollmentIsSet {
		symmetricKeyEnrollmentMap := symmetricKeyEnrollmentInterface.([]interface{})[0].(map[string]interface{})
		symmetricKeyEnrollment = SymmetricKeyEnrollmentDetailModelFromMap(symmetricKeyEnrollmentMap)
	}
	//
	tags := m["tags"].(map[string]string)
	var tpmEnrollment *models.TPMEnrollmentDetail // TPMEnrollmentDetail
	tpmEnrollmentInterface, tpmEnrollmentIsSet := m["tpm_enrollment"]
	if tpmEnrollmentIsSet {
		tpmEnrollmentMap := tpmEnrollmentInterface.([]interface{})[0].(map[string]interface{})
		tpmEnrollment = TPMEnrollmentDetailModelFromMap(tpmEnrollmentMap)
	}
	//
	return &models.EnrollmentDetail{
		AllocationPolicy:       allocationPolicy,
		AttachedIotHubsName:    attachedIotHubsName,
		CertificateEnrollment:  certificateEnrollment,
		EnableIotEdgeDevice:    enableIotEdgeDevice,
		Mechanism:              mechanism,
		SymmetricKeyEnrollment: symmetricKeyEnrollment,
		Tags:                   tags,
		TpmEnrollment:          tpmEnrollment,
	}
}

// Update the underlying EnrollmentDetail resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEnrollmentDetailResourceData(d *schema.ResourceData, m *models.EnrollmentDetail) {
	d.Set("allocation_policy", m.AllocationPolicy)
	d.Set("attached_iot_hubs_name", m.AttachedIotHubsName)
	d.Set("certificate_enrollment", m.CertificateEnrollment)
	d.Set("enable_iot_edge_device", m.EnableIotEdgeDevice)
	d.Set("mechanism", m.Mechanism)
	d.Set("symmetric_key_enrollment", SetSymmetricKeyEnrollmentDetailSubResourceData([]*models.SymmetricKeyEnrollmentDetail{m.SymmetricKeyEnrollment}))
	d.Set("tags", m.Tags)
	d.Set("tpm_enrollment", SetTPMEnrollmentDetailSubResourceData([]*models.TPMEnrollmentDetail{m.TpmEnrollment}))
}

// Iterate through and update the EnrollmentDetail resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEnrollmentDetailSubResourceData(m []*models.EnrollmentDetail) (d []*map[string]interface{}) {
	for _, EnrollmentDetailModel := range m {
		if EnrollmentDetailModel != nil {
			properties := make(map[string]interface{})
			properties["allocation_policy"] = EnrollmentDetailModel.AllocationPolicy
			properties["attached_iot_hubs_name"] = EnrollmentDetailModel.AttachedIotHubsName
			properties["certificate_enrollment"] = EnrollmentDetailModel.CertificateEnrollment
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

// Schema mapping representing the EnrollmentDetail resource defined in the Terraform configuration
func EnrollmentDetailSchema() map[string]*schema.Schema {
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
			Type:        schema.TypeString,
			Optional:    true,
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
				Schema: SymmetricKeyEnrollmentDetailSchema(),
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
				Schema: TPMEnrollmentDetailSchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the EnrollmentDetail resource
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
