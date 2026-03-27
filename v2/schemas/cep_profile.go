package schemas

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// validatePEMCertificate checks that the value is a valid PEM-encoded X.509 certificate.
func validatePEMCertificate(v interface{}, k string) (warnings []string, errors []error) {
	s := v.(string)
	block, _ := pem.Decode([]byte(s))
	if block == nil {
		errors = append(errors, fmt.Errorf("%q: not a valid PEM-encoded certificate (missing or malformed PEM block)", k))
		return
	}
	if _, err := x509.ParseCertificate(block.Bytes); err != nil {
		errors = append(errors, fmt.Errorf("%q: invalid X.509 certificate: %s", k, err))
	}
	return
}

func CEPProfileModel(d *schema.ResourceData) *models.CEPCommonSCEPProfile {
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	title, _ := d.Get("title").(string)
	description, _ := d.Get("description").(string)
	scepURL, _ := d.Get("scep_url").(string)
	useControllerProxy, _ := d.Get("use_controller_proxy").(bool)

	var caCertPem []strfmt.Base64
	caCertPemInterface, caCertPemIsSet := d.GetOk("ca_cert_pem")
	if caCertPemIsSet {
		for _, v := range caCertPemInterface.([]interface{}) {
			caCertPem = append(caCertPem, strfmt.Base64(v.(string)))
		}
	}

	var csrProfile *models.CEPCommonCSRProfile
	csrProfileInterface, csrProfileIsSet := d.GetOk("csr_profile")
	if csrProfileIsSet {
		csrProfileMap := csrProfileInterface.([]interface{})[0].(map[string]interface{})
		csrProfile = CEPCSRProfileModelFromMap(csrProfileMap)
	}

	var secret *models.CEPCommonSCEPProfileSecrets
	secretInterface, secretIsSet := d.GetOk("secret")
	if secretIsSet {
		secretList := secretInterface.([]interface{})
		if len(secretList) > 0 && secretList[0] != nil {
			secretMap := secretList[0].(map[string]interface{})
			secret = CEPSCEPProfileSecretsModelFromMap(secretMap)
		}
	}

	var revision *models.ObjectRevision
	revisionInterface, revisionIsSet := d.GetOk("revision")
	if revisionIsSet {
		revisionMap := revisionInterface.([]interface{})[0].(map[string]interface{})
		revision = ObjectRevisionModelFromMap(revisionMap)
	}

	return &models.CEPCommonSCEPProfile{
		ID:                 id,
		Name:               &name,
		Title:              title,
		Description:        description,
		ScepURL:            &scepURL,
		UseControllerProxy: useControllerProxy,
		CaCertPem:          caCertPem,
		CsrProfile:         csrProfile,
		Secret:             secret,
		Revision:           revision,
	}
}

func CEPProfileModelFromMap(m map[string]interface{}) *models.CEPCommonSCEPProfile {
	id := m["id"].(string)
	name := m["name"].(string)
	title := m["title"].(string)
	description := m["description"].(string)
	scepURL := m["scep_url"].(string)
	useControllerProxy := m["use_controller_proxy"].(bool)

	var caCertPem []strfmt.Base64
	caCertPemInterface, caCertPemIsSet := m["ca_cert_pem"]
	if caCertPemIsSet {
		for _, v := range caCertPemInterface.([]interface{}) {
			caCertPem = append(caCertPem, strfmt.Base64(v.(string)))
		}
	}

	var csrProfile *models.CEPCommonCSRProfile
	csrProfileInterface, csrProfileIsSet := m["csr_profile"]
	if csrProfileIsSet {
		csrProfileList := csrProfileInterface.([]interface{})
		if len(csrProfileList) > 0 && csrProfileList[0] != nil {
			csrProfile = CEPCSRProfileModelFromMap(csrProfileList[0].(map[string]interface{}))
		}
	}

	var secret *models.CEPCommonSCEPProfileSecrets
	secretInterface, secretIsSet := m["secret"]
	if secretIsSet {
		secretList := secretInterface.([]interface{})
		if len(secretList) > 0 && secretList[0] != nil {
			secret = CEPSCEPProfileSecretsModelFromMap(secretList[0].(map[string]interface{}))
		}
	}

	var revision *models.ObjectRevision
	revisionInterface, revisionIsSet := m["revision"]
	if revisionIsSet {
		revisionList := revisionInterface.([]interface{})
		if len(revisionList) > 0 && revisionList[0] != nil {
			revision = ObjectRevisionModelFromMap(revisionList[0].(map[string]interface{}))
		}
	}

	return &models.CEPCommonSCEPProfile{
		ID:                 id,
		Name:               &name,
		Title:              title,
		Description:        description,
		ScepURL:            &scepURL,
		UseControllerProxy: useControllerProxy,
		CaCertPem:          caCertPem,
		CsrProfile:         csrProfile,
		Secret:             secret,
		Revision:           revision,
	}
}

func CEPCSRProfileModelFromMap(m map[string]interface{}) *models.CEPCommonCSRProfile {
	commonName := m["common_name"].(string)
	organization := m["organization"].(string)
	organizationalUnit := m["organizational_unit"].(string)
	country := m["country"].(string)
	state := m["state"].(string)
	locality := m["locality"].(string)
	renewPeriodPercent := int64(m["renew_period_percent"].(int))

	var sanURI []string
	if v, ok := m["san_uri"]; ok {
		for _, s := range v.([]interface{}) {
			sanURI = append(sanURI, s.(string))
		}
	}

	var sanEmail []string
	if v, ok := m["san_email"]; ok {
		for _, s := range v.([]interface{}) {
			sanEmail = append(sanEmail, s.(string))
		}
	}

	var keyType *models.CEPCommonKeyType
	if v, ok := m["key_type"]; ok && v.(string) != "" {
		kt := models.CEPCommonKeyType(v.(string))
		keyType = &kt
	}

	var hashAlgorithm *models.CEPCommonHashAlgorithm
	if v, ok := m["hash_algorithm"]; ok && v.(string) != "" {
		ha := models.CEPCommonHashAlgorithm(v.(string))
		hashAlgorithm = &ha
	}

	return &models.CEPCommonCSRProfile{
		CommonName:         &commonName,
		Organization:       organization,
		OrganizationalUnit: organizationalUnit,
		Country:            country,
		State:              state,
		Locality:           locality,
		RenewPeriodPercent: renewPeriodPercent,
		SanURI:             sanURI,
		SanEmail:           sanEmail,
		KeyType:            keyType,
		HashAlgorithm:      hashAlgorithm,
	}
}

func CEPSCEPProfileSecretsModelFromMap(m map[string]interface{}) *models.CEPCommonSCEPProfileSecrets {
	challengePassword := m["challenge_password"].(string)
	return &models.CEPCommonSCEPProfileSecrets{
		ChallengePassword: challengePassword,
	}
}

// SetCEPProfileResourceData updates the underlying CEPProfile resource data in the Terraform configuration
func SetCEPProfileResourceData(d *schema.ResourceData, m *models.CEPCommonSCEPProfile) {
	d.Set("id", m.ID)
	if m.Name != nil {
		d.Set("name", *m.Name)
	}
	d.Set("title", m.Title)
	d.Set("description", m.Description)
	if m.ScepURL != nil {
		d.Set("scep_url", *m.ScepURL)
	}
	d.Set("use_controller_proxy", m.UseControllerProxy)

	var caCertPemStrings []string
	for _, b := range m.CaCertPem {
		caCertPemStrings = append(caCertPemStrings, string(b))
	}
	d.Set("ca_cert_pem", caCertPemStrings)

	d.Set("csr_profile", SetCEPCSRProfileSubResourceData([]*models.CEPCommonCSRProfile{m.CsrProfile}))
	// secret (challenge_password) is intentionally not set here: the API never returns it in responses.
	// Omitting d.Set("secret", ...) causes Terraform SDK v2 to preserve the user-configured value in
	// state across reads, avoiding a perpetual diff while keeping the value write-only/sensitive.
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("enterprise_id", m.EnterpriseID)
	d.Set("projects_list", SetCEPProjectReferenceSubResourceData(m.ProjectsList))
}

// SetCEPProfileSubResourceData iterates through and updates the CEPProfile resource data within a pagination response
func SetCEPProfileSubResourceData(m []*models.CEPCommonSCEPProfile) (d []*map[string]interface{}) {
	for _, CEPProfileModel := range m {
		if CEPProfileModel != nil {
			properties := make(map[string]interface{})
			properties["id"] = CEPProfileModel.ID
			if CEPProfileModel.Name != nil {
				properties["name"] = *CEPProfileModel.Name
			}
			properties["title"] = CEPProfileModel.Title
			properties["description"] = CEPProfileModel.Description
			if CEPProfileModel.ScepURL != nil {
				properties["scep_url"] = *CEPProfileModel.ScepURL
			}
			properties["use_controller_proxy"] = CEPProfileModel.UseControllerProxy
			var caCertPemStrings []string
			for _, b := range CEPProfileModel.CaCertPem {
				caCertPemStrings = append(caCertPemStrings, string(b))
			}
			properties["ca_cert_pem"] = caCertPemStrings
			properties["csr_profile"] = SetCEPCSRProfileSubResourceData([]*models.CEPCommonCSRProfile{CEPProfileModel.CsrProfile})
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{CEPProfileModel.Revision})
			properties["enterprise_id"] = CEPProfileModel.EnterpriseID
			properties["projects_list"] = SetCEPProjectReferenceSubResourceData(CEPProfileModel.ProjectsList)
			d = append(d, &properties)
		}
	}
	return
}

func SetCEPCSRProfileSubResourceData(m []*models.CEPCommonCSRProfile) (d []*map[string]interface{}) {
	for _, csrProfile := range m {
		if csrProfile != nil {
			properties := make(map[string]interface{})
			if csrProfile.CommonName != nil {
				properties["common_name"] = *csrProfile.CommonName
			}
			properties["organization"] = csrProfile.Organization
			properties["organizational_unit"] = csrProfile.OrganizationalUnit
			properties["country"] = csrProfile.Country
			properties["state"] = csrProfile.State
			properties["locality"] = csrProfile.Locality
			properties["renew_period_percent"] = int(csrProfile.RenewPeriodPercent)
			properties["san_uri"] = csrProfile.SanURI
			properties["san_email"] = csrProfile.SanEmail
			if csrProfile.KeyType != nil {
				properties["key_type"] = string(*csrProfile.KeyType)
			} else {
				properties["key_type"] = ""
			}
			if csrProfile.HashAlgorithm != nil {
				properties["hash_algorithm"] = string(*csrProfile.HashAlgorithm)
			} else {
				properties["hash_algorithm"] = ""
			}
			d = append(d, &properties)
		}
	}
	return
}

func SetCEPProjectReferenceSubResourceData(m []*models.CEPCommonProjectReference) (d []*map[string]interface{}) {
	for _, ref := range m {
		if ref != nil {
			properties := make(map[string]interface{})
			properties["project_id"] = ref.ProjectID
			properties["project_name"] = ref.ProjectName
			d = append(d, &properties)
		}
	}
	return
}

// CEPProfileSchema returns the schema mapping for the CEP profile resource
func CEPProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: `System-generated unique identifier (UUID)`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `User-defined unique name for the profile`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"title": {
			Description: `Human-readable display title`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"description": {
			Description: `Detailed description of the profile purpose`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"scep_url": {
			Description: `Full SCEP server URL including scheme, host, and path`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"use_controller_proxy": {
			Description: `If true, SCEP requests are routed through the controller proxy`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"ca_cert_pem": {
			Description: `PEM-encoded trusted CA certificates for SCEP server validation`,
			Type:        schema.TypeList,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validatePEMCertificate,
			},
			Required: true,
			MinItems: 1,
		},

		"csr_profile": {
			Description: `Certificate Signing Request configuration`,
			Type:        schema.TypeList,
			Required:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"common_name": {
						Description: `Common Name (CN) - typically the device identifier or hostname`,
						Type:        schema.TypeString,
						Required:    true,
					},
					"organization": {
						Description: `Organization (O) - company or entity name`,
						Type:        schema.TypeString,
						Optional:    true,
					},
					"organizational_unit": {
						Description: `Organizational Unit (OU) - department or division`,
						Type:        schema.TypeString,
						Optional:    true,
					},
					"country": {
						Description: `Country (C) - two-letter ISO country code`,
						Type:        schema.TypeString,
						Optional:    true,
					},
					"state": {
						Description: `State or Province (ST)`,
						Type:        schema.TypeString,
						Optional:    true,
					},
					"locality": {
						Description: `Locality (L) - city name`,
						Type:        schema.TypeString,
						Optional:    true,
					},
					"san_uri": {
						Description: `URIs for Subject Alternative Name extension`,
						Type:        schema.TypeList,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Optional: true,
					},
					"san_email": {
						Description: `Email addresses for Subject Alternative Name extension`,
						Type:        schema.TypeList,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Optional: true,
					},
					"renew_period_percent": {
						Description: `Percentage of certificate validity period after which renewal should be attempted`,
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"key_type": {
						Description: `Key type and size for certificate key pair generation`,
						Type:        schema.TypeString,
						Optional:    true,
					},
					"hash_algorithm": {
						Description: `Hash algorithm used for signing the CSR and SCEP messages`,
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},

		"secret": {
			Description: `Write-only secret fields (challenge password)`,
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"challenge_password": {
						Description: `Challenge password for SCEP enrollment`,
						Type:        schema.TypeString,
						Required:    true,
						Sensitive:   true,
					},
				},
			},
		},

		"revision": {
			Description: `Object revision information`,
			Type:        schema.TypeList,
			Computed:    true,
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
		},

		"enterprise_id": {
			Description: `Enterprise ID owning this profile`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"projects_list": {
			Description: `Projects referencing this CEP profile`,
			Type:        schema.TypeList,
			Computed:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"project_id": {
						Description: `Project ID`,
						Type:        schema.TypeString,
						Computed:    true,
					},
					"project_name": {
						Description: `Project name`,
						Type:        schema.TypeString,
						Computed:    true,
					},
				},
			},
		},
	}
}
