package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func PortBasedNetworkAccessControlPolicyModel(d *schema.ResourceData) *models.PortBasedNetworkAccessControlPolicy {
	certificateEnrollmentID, _ := d.Get("certificate_enrollment_id").(string)
	eapIdentity, _ := d.Get("eap_identity").(string)
	var eapMethod *models.EAPMethod // EAPMethod
	eapMethodInterface, eapMethodIsSet := d.GetOk("eap_method")
	if eapMethodIsSet {
		eapMethodModel := eapMethodInterface.(string)
		eapMethod = models.NewEAPMethod(models.EAPMethod(eapMethodModel))
	}
	enablePortBasedNetworkAccessControl, _ := d.Get("enable_port_based_network_access_control").(bool)
	id, _ := d.Get("id").(string)
	return &models.PortBasedNetworkAccessControlPolicy{
		CertificateEnrollmentID:             certificateEnrollmentID,
		EapIdentity:                         eapIdentity,
		EapMethod:                           eapMethod,
		EnablePortBasedNetworkAccessControl: enablePortBasedNetworkAccessControl,
		ID:                                  id,
	}
}

func PortBasedNetworkAccessControlPolicyModelFromMap(m map[string]interface{}) *models.PortBasedNetworkAccessControlPolicy {
	certificateEnrollmentID := m["certificate_enrollment_id"].(string)
	eapIdentity := m["eap_identity"].(string)
	var eapMethod *models.EAPMethod // EAPMethod
	eapMethodInterface, eapMethodIsSet := m["eap_method"]
	if eapMethodIsSet {
		eapMethodModel := eapMethodInterface.(string)
		eapMethod = models.NewEAPMethod(models.EAPMethod(eapMethodModel))
	}
	enablePortBasedNetworkAccessControl := m["enable_port_based_network_access_control"].(bool)
	id := m["id"].(string)
	return &models.PortBasedNetworkAccessControlPolicy{
		CertificateEnrollmentID:             certificateEnrollmentID,
		EapIdentity:                         eapIdentity,
		EapMethod:                           eapMethod,
		EnablePortBasedNetworkAccessControl: enablePortBasedNetworkAccessControl,
		ID:                                  id,
	}
}

func SetPortBasedNetworkAccessControlPolicyResourceData(d *schema.ResourceData, m *models.PortBasedNetworkAccessControlPolicy) {
	d.Set("certificate_enrollment_id", m.CertificateEnrollmentID)
	d.Set("eap_identity", m.EapIdentity)
	d.Set("eap_method", m.EapMethod)
	d.Set("enable_port_based_network_access_control", m.EnablePortBasedNetworkAccessControl)
	d.Set("id", m.ID)
}

func SetPortBasedNetworkAccessControlPolicySubResourceData(m []*models.PortBasedNetworkAccessControlPolicy) (d []*map[string]interface{}) {
	for _, PortBasedNetworkAccessControlPolicyModel := range m {
		if PortBasedNetworkAccessControlPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["certificate_enrollment_id"] = PortBasedNetworkAccessControlPolicyModel.CertificateEnrollmentID
			properties["eap_identity"] = PortBasedNetworkAccessControlPolicyModel.EapIdentity
			properties["eap_method"] = PortBasedNetworkAccessControlPolicyModel.EapMethod
			properties["enable_port_based_network_access_control"] = PortBasedNetworkAccessControlPolicyModel.EnablePortBasedNetworkAccessControl
			properties["id"] = PortBasedNetworkAccessControlPolicyModel.ID
			d = append(d, &properties)
		}
	}
	return
}

func PortBasedNetworkAccessControlPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"certificate_enrollment_id": {
			Description: `ID of the certificate enrollment object used for 802.1X authentication`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"eap_identity": {
			Description: `EAP identity string used for authentication`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"eap_method": {
			Description: `EAP method to use for port based network access control. Only EAP-TLS is supported currently.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"enable_port_based_network_access_control": {
			Description: `Enable or disable port based network access control (802.1X). When disabled, other fields are ignored.`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"id": {
			Description: `unique policy id`,
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}

func GetPortBasedNetworkAccessControlPolicyPropertyFields() (t []string) {
	return []string{
		"certificate_enrollment_id",
		"eap_identity",
		"eap_method",
		"enable_port_based_network_access_control",
		"id",
	}
}

func DevicePortBasedNetworkAccessControlPolicyModel(d *schema.ResourceData) *models.DevicePortBasedNetworkAccessControlPolicy {
	certificateEnrollmentID, _ := d.Get("certificate_enrollment_id").(string)
	eapIdentity, _ := d.Get("eap_identity").(string)
	var eapMethod *models.EAPMethod // EAPMethod
	eapMethodInterface, eapMethodIsSet := d.GetOk("eap_method")
	if eapMethodIsSet {
		eapMethodModel := eapMethodInterface.(string)
		eapMethod = models.NewEAPMethod(models.EAPMethod(eapMethodModel))
	}
	enablePortBasedNetworkAccessControl, _ := d.Get("enable_port_based_network_access_control").(bool)
	return &models.DevicePortBasedNetworkAccessControlPolicy{
		CertificateEnrollmentID:             certificateEnrollmentID,
		EapIdentity:                         eapIdentity,
		EapMethod:                           eapMethod,
		EnablePortBasedNetworkAccessControl: enablePortBasedNetworkAccessControl,
	}
}

func DevicePortBasedNetworkAccessControlPolicyModelFromMap(m map[string]interface{}) *models.DevicePortBasedNetworkAccessControlPolicy {
	certificateEnrollmentID := m["certificate_enrollment_id"].(string)
	eapIdentity := m["eap_identity"].(string)
	var eapMethod *models.EAPMethod // EAPMethod
	eapMethodInterface, eapMethodIsSet := m["eap_method"]
	if eapMethodIsSet {
		eapMethodModel := eapMethodInterface.(string)
		eapMethod = models.NewEAPMethod(models.EAPMethod(eapMethodModel))
	}
	enablePortBasedNetworkAccessControl := m["enable_port_based_network_access_control"].(bool)
	return &models.DevicePortBasedNetworkAccessControlPolicy{
		CertificateEnrollmentID:             certificateEnrollmentID,
		EapIdentity:                         eapIdentity,
		EapMethod:                           eapMethod,
		EnablePortBasedNetworkAccessControl: enablePortBasedNetworkAccessControl,
	}
}

func SetDevicePortBasedNetworkAccessControlPolicyResourceData(d *schema.ResourceData, m *models.DevicePortBasedNetworkAccessControlPolicy) {
	d.Set("certificate_enrollment_id", m.CertificateEnrollmentID)
	d.Set("eap_identity", m.EapIdentity)
	d.Set("eap_method", m.EapMethod)
	d.Set("enable_port_based_network_access_control", m.EnablePortBasedNetworkAccessControl)
}

func SetDevicePortBasedNetworkAccessControlPolicySubResourceData(m []*models.DevicePortBasedNetworkAccessControlPolicy) (d []*map[string]interface{}) {
	for _, DevicePortBasedNetworkAccessControlPolicyModel := range m {
		if DevicePortBasedNetworkAccessControlPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["certificate_enrollment_id"] = DevicePortBasedNetworkAccessControlPolicyModel.CertificateEnrollmentID
			properties["eap_identity"] = DevicePortBasedNetworkAccessControlPolicyModel.EapIdentity
			properties["eap_method"] = DevicePortBasedNetworkAccessControlPolicyModel.EapMethod
			properties["enable_port_based_network_access_control"] = DevicePortBasedNetworkAccessControlPolicyModel.EnablePortBasedNetworkAccessControl
			d = append(d, &properties)
		}
	}
	return
}

func DevicePortBasedNetworkAccessControlPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"certificate_enrollment_id": {
			Description: `ID of the certificate enrollment object used for 802.1X authentication`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"eap_identity": {
			Description: `EAP identity string used for authentication`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"eap_method": {
			Description: `EAP method to use for port based network access control. Only EAP-TLS is supported currently.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"enable_port_based_network_access_control": {
			Description: `Enable or disable port based network access control (802.1X). When disabled, other fields are ignored.`,
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}

func GetDevicePortBasedNetworkAccessControlPolicyPropertyFields() (t []string) {
	return []string{
		"certificate_enrollment_id",
		"eap_identity",
		"eap_method",
		"enable_port_based_network_access_control",
	}
}
