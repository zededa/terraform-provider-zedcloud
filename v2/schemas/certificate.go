package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func CertificateModel(d *schema.ResourceData) *models.Certificate {
	basicContraintsValid, _ := d.Get("basic_contraints_valid").(bool)
	cert, _ := d.Get("cert").(string)
	cryptoKey, _ := d.Get("crypto_key").(string)
	var ecdsaEncryption *models.ECDSA // ECDSA
	ecdsaEncryptionInterface, ecdsaEncryptionIsSet := d.GetOk("ecdsa_encryption")
	if ecdsaEncryptionIsSet && ecdsaEncryptionInterface != nil {
		ecdsaEncryptionMap := ecdsaEncryptionInterface.([]interface{})
		if len(ecdsaEncryptionMap) > 0 {
			ecdsaEncryption = ECDSAModelFromMap(ecdsaEncryptionMap[0].(map[string]interface{}))
		}
	}
	encryptedSecrets := map[string]string{}
	encryptedSecretsInterface, encryptedSecretsIsSet := d.GetOk("encryptedSecrets")
	if encryptedSecretsIsSet {
		encryptedSecretsMap := encryptedSecretsInterface.(map[string]interface{})
		for k, v := range encryptedSecretsMap {
			if v == nil {
				continue
			}
			encryptedSecrets[k] = v.(string)
		}
	}

	exportable, _ := d.Get("exportable").(bool)
	var extendedKeyUsage []string
	extendedKeyUsageInterface, extendedKeyUsageIsSet := d.GetOk("extendedKeyUsage")
	if extendedKeyUsageIsSet {
		var items []interface{}
		if listItems, isList := extendedKeyUsageInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = extendedKeyUsageInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			extendedKeyUsage = append(extendedKeyUsage, v.(string))
		}
	}
	var issuer *models.Subject // Subject
	issuerInterface, issuerIsSet := d.GetOk("issuer")
	if issuerIsSet && issuerInterface != nil {
		issuerMap := issuerInterface.([]interface{})
		if len(issuerMap) > 0 {
			issuer = SubjectModelFromMap(issuerMap[0].(map[string]interface{}))
		}
	}
	keyUsageInt, _ := d.Get("key_usage").(int)
	keyUsage := int32(keyUsageInt)
	passPhrase, _ := d.Get("pass_phrase").(string)
	publicKey, _ := d.Get("public_key").(string)
	publicKeyAlgorithm, _ := d.Get("public_key_algorithm").(string)
	pvtKey, _ := d.Get("pvt_key").(string)
	reuseKey, _ := d.Get("reuse_key").(bool)
	var rsaEcryption *models.RSA // RSA
	rsaEcryptionInterface, rsaEcryptionIsSet := d.GetOk("rsa_ecryption")
	if rsaEcryptionIsSet && rsaEcryptionInterface != nil {
		rsaEcryptionMap := rsaEcryptionInterface.([]interface{})
		if len(rsaEcryptionMap) > 0 {
			rsaEcryption = RSAModelFromMap(rsaEcryptionMap[0].(map[string]interface{}))
		}
	}
	var sanValues *models.SANValues // SANValues
	sanValuesInterface, sanValuesIsSet := d.GetOk("san_values")
	if sanValuesIsSet && sanValuesInterface != nil {
		sanValuesMap := sanValuesInterface.([]interface{})
		if len(sanValuesMap) > 0 {
			sanValues = SANValuesModelFromMap(sanValuesMap[0].(map[string]interface{}))
		}
	}
	serialNumber, _ := d.Get("serial_number").(string)
	signatureAlgorithm, _ := d.Get("signature_algorithm").(string)
	var subject *models.Subject // Subject
	subjectInterface, subjectIsSet := d.GetOk("subject")
	if subjectIsSet && subjectInterface != nil {
		subjectMap := subjectInterface.([]interface{})
		if len(subjectMap) > 0 {
			subject = SubjectModelFromMap(subjectMap[0].(map[string]interface{}))
		}
	}
	validFrom, _ := d.Get("valid_from").(strfmt.DateTime)
	validTill, _ := d.Get("valid_till").(strfmt.DateTime)
	return &models.Certificate{
		BasicContraintsValid: basicContraintsValid,
		Cert:                 cert,
		CryptoKey:            cryptoKey,
		EcdsaEncryption:      ecdsaEncryption,
		EncryptedSecrets:     encryptedSecrets,
		Exportable:           exportable,
		ExtendedKeyUsage:     extendedKeyUsage,
		Issuer:               issuer,
		KeyUsage:             keyUsage,
		PassPhrase:           passPhrase,
		PublicKey:            publicKey,
		PublicKeyAlgorithm:   publicKeyAlgorithm,
		PvtKey:               pvtKey,
		ReuseKey:             reuseKey,
		RsaEcryption:         rsaEcryption,
		SanValues:            sanValues,
		SerialNumber:         serialNumber,
		SignatureAlgorithm:   signatureAlgorithm,
		Subject:              subject,
		ValidFrom:            validFrom,
		ValidTill:            validTill,
	}
}

func CertificateModelFromMap(m map[string]interface{}) *models.Certificate {
	basicContraintsValid := m["basic_contraints_valid"].(bool)
	cert := m["cert"].(string)
	cryptoKey := m["crypto_key"].(string)
	var ecdsaEncryption *models.ECDSA // ECDSA
	ecdsaEncryptionInterface, ecdsaEncryptionIsSet := m["ecdsa_encryption"]
	if ecdsaEncryptionIsSet && ecdsaEncryptionInterface != nil {
		ecdsaEncryptionMap := ecdsaEncryptionInterface.([]interface{})
		if len(ecdsaEncryptionMap) > 0 {
			ecdsaEncryption = ECDSAModelFromMap(ecdsaEncryptionMap[0].(map[string]interface{}))
		}
	}
	//
	encryptedSecrets := map[string]string{}
	encryptedSecretsInterface, encryptedSecretsIsSet := m["encrypted_secrets"]
	if encryptedSecretsIsSet {
		encryptedSecretsMap := encryptedSecretsInterface.(map[string]interface{})
		for k, v := range encryptedSecretsMap {
			if v == nil {
				continue
			}
			encryptedSecrets[k] = v.(string)
		}
	}

	exportable := m["exportable"].(bool)
	var extendedKeyUsage []string
	extendedKeyUsageInterface, extendedKeyUsageIsSet := m["extendedKeyUsage"]
	if extendedKeyUsageIsSet {
		var items []interface{}
		if listItems, isList := extendedKeyUsageInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = extendedKeyUsageInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			extendedKeyUsage = append(extendedKeyUsage, v.(string))
		}
	}
	var issuer *models.Subject // Subject
	issuerInterface, issuerIsSet := m["issuer"]
	if issuerIsSet && issuerInterface != nil {
		issuerMap := issuerInterface.([]interface{})
		if len(issuerMap) > 0 {
			issuer = SubjectModelFromMap(issuerMap[0].(map[string]interface{}))
		}
	}
	//
	keyUsage := int32(m["key_usage"].(int)) // int32
	passPhrase := m["pass_phrase"].(string)
	publicKey := m["public_key"].(string)
	publicKeyAlgorithm := m["public_key_algorithm"].(string)
	pvtKey := m["pvt_key"].(string)
	reuseKey := m["reuse_key"].(bool)
	var rsaEcryption *models.RSA // RSA
	rsaEcryptionInterface, rsaEcryptionIsSet := m["rsa_ecryption"]
	if rsaEcryptionIsSet && rsaEcryptionInterface != nil {
		rsaEcryptionMap := rsaEcryptionInterface.([]interface{})
		if len(rsaEcryptionMap) > 0 {
			rsaEcryption = RSAModelFromMap(rsaEcryptionMap[0].(map[string]interface{}))
		}
	}
	//
	var sanValues *models.SANValues // SANValues
	sanValuesInterface, sanValuesIsSet := m["san_values"]
	if sanValuesIsSet && sanValuesInterface != nil {
		sanValuesMap := sanValuesInterface.([]interface{})
		if len(sanValuesMap) > 0 {
			sanValues = SANValuesModelFromMap(sanValuesMap[0].(map[string]interface{}))
		}
	}
	//
	serialNumber := m["serial_number"].(string)
	signatureAlgorithm := m["signature_algorithm"].(string)
	var subject *models.Subject // Subject
	subjectInterface, subjectIsSet := m["subject"]
	if subjectIsSet && subjectInterface != nil {
		subjectMap := subjectInterface.([]interface{})
		if len(subjectMap) > 0 {
			subject = SubjectModelFromMap(subjectMap[0].(map[string]interface{}))
		}
	}
	//
	validFrom := m["valid_from"].(strfmt.DateTime)
	validTill := m["valid_till"].(strfmt.DateTime)
	return &models.Certificate{
		BasicContraintsValid: basicContraintsValid,
		Cert:                 cert,
		CryptoKey:            cryptoKey,
		EcdsaEncryption:      ecdsaEncryption,
		EncryptedSecrets:     encryptedSecrets,
		Exportable:           exportable,
		ExtendedKeyUsage:     extendedKeyUsage,
		Issuer:               issuer,
		KeyUsage:             keyUsage,
		PassPhrase:           passPhrase,
		PublicKey:            publicKey,
		PublicKeyAlgorithm:   publicKeyAlgorithm,
		PvtKey:               pvtKey,
		ReuseKey:             reuseKey,
		RsaEcryption:         rsaEcryption,
		SanValues:            sanValues,
		SerialNumber:         serialNumber,
		SignatureAlgorithm:   signatureAlgorithm,
		Subject:              subject,
		ValidFrom:            validFrom,
		ValidTill:            validTill,
	}
}

func SetCertificateResourceData(d *schema.ResourceData, m *models.Certificate) {
	d.Set("basic_contraints_valid", m.BasicContraintsValid)
	d.Set("cert", m.Cert)
	d.Set("crypto_key", m.CryptoKey)
	d.Set("ecdsa_encryption", SetECDSASubResourceData([]*models.ECDSA{m.EcdsaEncryption}))
	d.Set("encrypted_secrets", m.EncryptedSecrets)
	d.Set("exportable", m.Exportable)
	d.Set("extended_key_usage", m.ExtendedKeyUsage)
	d.Set("issuer", SetSubjectSubResourceData([]*models.Subject{m.Issuer}))
	d.Set("key_usage", m.KeyUsage)
	d.Set("pass_phrase", m.PassPhrase)
	d.Set("public_key", m.PublicKey)
	d.Set("public_key_algorithm", m.PublicKeyAlgorithm)
	d.Set("pvt_key", m.PvtKey)
	d.Set("reuse_key", m.ReuseKey)
	d.Set("rsa_ecryption", SetRSASubResourceData([]*models.RSA{m.RsaEcryption}))
	d.Set("san_values", SetSANValuesSubResourceData([]*models.SANValues{m.SanValues}))
	d.Set("serial_number", m.SerialNumber)
	d.Set("signature_algorithm", m.SignatureAlgorithm)
	d.Set("subject", SetSubjectSubResourceData([]*models.Subject{m.Subject}))
	d.Set("valid_from", m.ValidFrom)
	d.Set("valid_till", m.ValidTill)
}

func SetCertificateSubResourceData(m []*models.Certificate) (d []*map[string]interface{}) {
	for _, CertificateModel := range m {
		if CertificateModel != nil {
			properties := make(map[string]interface{})
			properties["basic_contraints_valid"] = CertificateModel.BasicContraintsValid
			properties["cert"] = CertificateModel.Cert
			properties["crypto_key"] = CertificateModel.CryptoKey
			properties["ecdsa_encryption"] = SetECDSASubResourceData([]*models.ECDSA{CertificateModel.EcdsaEncryption})
			properties["encrypted_secrets"] = CertificateModel.EncryptedSecrets
			properties["exportable"] = CertificateModel.Exportable
			properties["extended_key_usage"] = CertificateModel.ExtendedKeyUsage
			properties["issuer"] = SetSubjectSubResourceData([]*models.Subject{CertificateModel.Issuer})
			properties["key_usage"] = CertificateModel.KeyUsage
			properties["pass_phrase"] = CertificateModel.PassPhrase
			properties["public_key"] = CertificateModel.PublicKey
			properties["public_key_algorithm"] = CertificateModel.PublicKeyAlgorithm
			properties["pvt_key"] = CertificateModel.PvtKey
			properties["reuse_key"] = CertificateModel.ReuseKey
			properties["rsa_ecryption"] = SetRSASubResourceData([]*models.RSA{CertificateModel.RsaEcryption})
			properties["san_values"] = SetSANValuesSubResourceData([]*models.SANValues{CertificateModel.SanValues})
			properties["serial_number"] = CertificateModel.SerialNumber
			properties["signature_algorithm"] = CertificateModel.SignatureAlgorithm
			properties["subject"] = SetSubjectSubResourceData([]*models.Subject{CertificateModel.Subject})
			properties["valid_from"] = CertificateModel.ValidFrom.String()
			properties["valid_till"] = CertificateModel.ValidTill.String()
			d = append(d, &properties)
		}
	}
	return
}

func Certificate() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"basic_contraints_valid": {
			Description: `This fields tells the basic constraints like isCA are correct.`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"cert": {
			Description: `base64 string of the parent certificate`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"crypto_key": {
			Description: `Crypto Key for decrypting user secret information`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"ecdsa_encryption": {
			Description: `ECDSA encryption algorithm of the certificate`,
			Type:        schema.TypeList, //GoType: ECDSA
			Elem: &schema.Resource{
				Schema: ECDSA(),
			},
			Optional: true,
		},

		"encrypted_secrets": {
			Description: `user encrypted secrets map`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"exportable": {
			Description: `Indicates if the private key can be exported.`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"extended_key_usage": {
			Description: `Sequence of extended key usages.`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"issuer": {
			Description: `Parameters for the issuer of the X509 component of a certificate.`,
			Type:        schema.TypeList, //GoType: Subject
			Elem: &schema.Resource{
				Schema: Subject(),
			},
			Optional: true,
		},

		"key_usage": {
			Description: `Key usage extensions define the purpose of the public key contained in a certificate.`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"pass_phrase": {
			Description: `pass phase for the pvt key, this has to be filled if pvt key is encrypted with a pass phrase`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"public_key": {
			Description: `base63 string of the public key`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"public_key_algorithm": {
			Description: `Public key algorithm.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"pvt_key": {
			Description: `base64 string of the parent pvt key`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"reuse_key": {
			Description: `Indicates if the same key pair will be used on certificate renewal.`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"rsa_ecryption": {
			Description: `RSA encryption algorithm of the certificate`,
			Type:        schema.TypeList, //GoType: RSA
			Elem: &schema.Resource{
				Schema: RSA(),
			},
			Optional: true,
		},

		"san_values": {
			Description: `This holds the alternative name values like URIs, domain names IPs etc.`,
			Type:        schema.TypeList, //GoType: SANValues
			Elem: &schema.Resource{
				Schema: SANValues(),
			},
			Optional: true,
		},

		"serial_number": {
			Description: `Unique identifier for each Certificate generated by an Certificate Issuer. `,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"signature_algorithm": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"subject": {
			Description: `Parameters for the subject of the X509 component of a certificate.`,
			Type:        schema.TypeList, //GoType: Subject
			Elem: &schema.Resource{
				Schema: Subject(),
			},
			Optional: true,
		},

		"valid_from": {
			Description:  `Certificate validatity start time`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"valid_till": {
			Description:  `Certificate validatity start time`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},
	}
}

func GetCertificatePropertyFields() (t []string) {
	return []string{
		"basic_contraints_valid",
		"cert",
		"crypto_key",
		"ecdsa_encryption",
		"encrypted_secrets",
		"exportable",
		"extended_key_usage",
		"issuer",
		"key_usage",
		"pass_phrase",
		"public_key",
		"public_key_algorithm",
		"pvt_key",
		"reuse_key",
		"rsa_ecryption",
		"san_values",
		"serial_number",
		"signature_algorithm",
		"subject",
		"valid_from",
		"valid_till",
	}
}
