package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SetEnrolledCertInfoSubResourceData(m []*models.EnrolledCertInfo) (d []*map[string]interface{}) {
	for _, v := range m {
		if v != nil {
			properties := make(map[string]interface{})
			properties["cert_enrollment_profile_name"] = v.CertEnrollmentProfileName
			properties["subject"] = SetCertDistinguishedNameSubResourceData([]*models.CertDistinguishedName{v.Subject})
			properties["issuer"] = SetCertDistinguishedNameSubResourceData([]*models.CertDistinguishedName{v.Issuer})
			properties["san_dns"] = v.SanDNS
			properties["san_ip"] = v.SanIP
			properties["san_uri"] = v.SanURI
			properties["san_email"] = v.SanEmail
			properties["issue_timestamp"] = v.IssueTimestamp.String()
			properties["expiration_timestamp"] = v.ExpirationTimestamp.String()
			properties["renew_period_percent"] = v.RenewPeriodPercent
			properties["sha256_fingerprint"] = v.Sha256Fingerprint
			if v.Status != nil {
				properties["status"] = string(*v.Status)
			}
			properties["err"] = SetDeviceErrorSubResourceData([]*models.DeviceError{v.Err})
			d = append(d, &properties)
		}
	}
	return
}

func EnrolledCertInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cert_enrollment_profile_name": {
			Description: `Certificate enrollment profile name used to obtain this certificate`,
			Type:        schema.TypeString,
			Computed:    true,
		},
		"subject": {
			Description: `Subject distinguished name (DN) of the certificate`,
			Type:        schema.TypeList,
			Elem:        &schema.Resource{Schema: CertDistinguishedNameSchema()},
			Computed:    true,
		},
		"issuer": {
			Description: `Issuer distinguished name (DN) of the certificate`,
			Type:        schema.TypeList,
			Elem:        &schema.Resource{Schema: CertDistinguishedNameSchema()},
			Computed:    true,
		},
		"san_dns": {
			Description: `DNS names from Subject Alternative Name`,
			Type:        schema.TypeList,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Computed:    true,
		},
		"san_ip": {
			Description: `IP addresses from Subject Alternative Name`,
			Type:        schema.TypeList,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Computed:    true,
		},
		"san_uri": {
			Description: `URIs from Subject Alternative Name`,
			Type:        schema.TypeList,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Computed:    true,
		},
		"san_email": {
			Description: `Email addresses from Subject Alternative Name`,
			Type:        schema.TypeList,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Computed:    true,
		},
		"issue_timestamp": {
			Description: `Certificate issue timestamp (notBefore)`,
			Type:        schema.TypeString,
			Computed:    true,
		},
		"expiration_timestamp": {
			Description: `Certificate expiration timestamp (notAfter)`,
			Type:        schema.TypeString,
			Computed:    true,
		},
		"renew_period_percent": {
			Description: `Percentage of validity period after which renewal begins`,
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"sha256_fingerprint": {
			Description: `SHA-256 fingerprint of the certificate (hex-encoded)`,
			Type:        schema.TypeString,
			Computed:    true,
		},
		"status": {
			Description: `Current status of the certificate as observed by the device`,
			Type:        schema.TypeString,
			Computed:    true,
		},
		"err": {
			Description: `Error details for enrollment or renewal failures`,
			Type:        schema.TypeList,
			Elem:        &schema.Resource{Schema: DeviceErrorSchema()},
			Computed:    true,
		},
	}
}
