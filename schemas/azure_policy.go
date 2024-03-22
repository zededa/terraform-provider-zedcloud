package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func AzurePolicyModel(d *schema.ResourceData) *models.AzurePolicy {
	appID, _ := d.Get("app_id").(string)
	appPassword, _ := d.Get("app_password").(string)
	var azureResourceAndServices *models.AzureResourceAndServices // AzureResourceAndServices
	azureResourceAndServicesInterface, azureResourceAndServicesIsSet := d.GetOk("azure_resource_and_services")
	if azureResourceAndServicesIsSet && azureResourceAndServicesInterface != nil {
		azureResourceAndServicesMap := azureResourceAndServicesInterface.([]interface{})
		if len(azureResourceAndServicesMap) > 0 {
			azureResourceAndServices = AzureResourceAndServicesModelFromMap(azureResourceAndServicesMap[0].(map[string]interface{}))
		}
	}
	var certificate *models.Certificate // Certificate
	certificateInterface, certificateIsSet := d.GetOk("certificate")
	if certificateIsSet && certificateInterface != nil {
		certificateMap := certificateInterface.([]interface{})
		if len(certificateMap) > 0 {
			certificate = CertificateModelFromMap(certificateMap[0].(map[string]interface{}))
		}
	}
	cryptoKey, _ := d.Get("crypto_key").(string)
	customDeploymentManaged, _ := d.Get("custom_deployment_managed").(bool)
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

	tenantID, _ := d.Get("tenant_id").(string)
	return &models.AzurePolicy{
		AppID:                    &appID,       // string true false false
		AppPassword:              &appPassword, // string true false false
		AzureResourceAndServices: azureResourceAndServices,
		Certificate:              certificate,
		CryptoKey:                cryptoKey,
		CustomDeploymentManaged:  customDeploymentManaged,
		EncryptedSecrets:         encryptedSecrets,
		TenantID:                 &tenantID, // string true false false
	}
}

func AzurePolicyModelFromMap(m map[string]interface{}) *models.AzurePolicy {
	appID := m["app_id"].(string)
	appPassword := m["app_password"].(string)
	var azureResourceAndServices *models.AzureResourceAndServices // AzureResourceAndServices
	azureResourceAndServicesInterface, azureResourceAndServicesIsSet := m["azure_resource_and_services"]
	if azureResourceAndServicesIsSet && azureResourceAndServicesInterface != nil {
		azureResourceAndServicesMap := azureResourceAndServicesInterface.([]interface{})
		if len(azureResourceAndServicesMap) > 0 {
			azureResourceAndServices = AzureResourceAndServicesModelFromMap(azureResourceAndServicesMap[0].(map[string]interface{}))
		}
	}
	//
	var certificate *models.Certificate // Certificate
	certificateInterface, certificateIsSet := m["certificate"]
	if certificateIsSet && certificateInterface != nil {
		certificateMap := certificateInterface.([]interface{})
		if len(certificateMap) > 0 {
			certificate = CertificateModelFromMap(certificateMap[0].(map[string]interface{}))
		}
	}
	//
	cryptoKey := m["crypto_key"].(string)
	customDeploymentManaged := m["custom_deployment_managed"].(bool)
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

	tenantID := m["tenant_id"].(string)
	return &models.AzurePolicy{
		AppID:                    &appID,
		AppPassword:              &appPassword,
		AzureResourceAndServices: azureResourceAndServices,
		Certificate:              certificate,
		CryptoKey:                cryptoKey,
		CustomDeploymentManaged:  customDeploymentManaged,
		EncryptedSecrets:         encryptedSecrets,
		TenantID:                 &tenantID,
	}
}

func SetAzurePolicyResourceData(d *schema.ResourceData, m *models.AzurePolicy) {
	d.Set("app_id", m.AppID)
	d.Set("app_password", m.AppPassword)
	d.Set("azure_resource_and_services", SetAzureResourceAndServicesSubResourceData([]*models.AzureResourceAndServices{m.AzureResourceAndServices}))
	d.Set("certificate", SetCertificateSubResourceData([]*models.Certificate{m.Certificate}))
	d.Set("crypto_key", m.CryptoKey)
	d.Set("custom_deployment_managed", m.CustomDeploymentManaged)
	d.Set("encrypted_secrets", m.EncryptedSecrets)
	d.Set("tenant_id", m.TenantID)
}

func SetAzurePolicySubResourceData(m []*models.AzurePolicy) (d []*map[string]interface{}) {
	for _, AzurePolicyModel := range m {
		if AzurePolicyModel != nil {
			properties := make(map[string]interface{})
			properties["app_id"] = AzurePolicyModel.AppID
			properties["app_password"] = AzurePolicyModel.AppPassword
			properties["azure_resource_and_services"] = SetAzureResourceAndServicesSubResourceData([]*models.AzureResourceAndServices{AzurePolicyModel.AzureResourceAndServices})
			properties["certificate"] = SetCertificateSubResourceData([]*models.Certificate{AzurePolicyModel.Certificate})
			properties["crypto_key"] = AzurePolicyModel.CryptoKey
			properties["custom_deployment_managed"] = AzurePolicyModel.CustomDeploymentManaged
			properties["encrypted_secrets"] = AzurePolicyModel.EncryptedSecrets
			properties["tenant_id"] = AzurePolicyModel.TenantID
			d = append(d, &properties)
		}
	}
	return
}

func AzurePolicy() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_id": {
			Description: `app id for rbac`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"app_password": {
			Description: `app password for rbac`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"azure_resource_and_services": {
			Description: `azure resource and service the policy will be interested in`,
			Type:        schema.TypeList, //GoType: AzureResourceAndServices
			Elem: &schema.Resource{
				Schema: AzureResourceAndServices(),
			},
			Optional: true,
		},

		"certificate": {
			Description: `Certificate object holds the details of certificate like encryption type, validity, subject etc`,
			Type:        schema.TypeList, //GoType: Certificate
			Elem: &schema.Resource{
				Schema: Certificate(),
			},
			Optional: true,
		},

		"crypto_key": {
			Description: `key to decrypt AppPassword`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"custom_deployment_managed": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"encrypted_secrets": {
			Description: `encrypted AppPassword`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"tenant_id": {
			Description: `tenant id for rbac`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func GetAzurePolicyPropertyFields() (t []string) {
	return []string{
		"app_id",
		"app_password",
		"azure_resource_and_services",
		"certificate",
		"crypto_key",
		"custom_deployment_managed",
		"encrypted_secrets",
		"tenant_id",
	}
}
