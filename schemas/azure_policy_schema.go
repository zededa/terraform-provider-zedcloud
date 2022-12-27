package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate AzurePolicy resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AzurePolicyModel(d *schema.ResourceData) *models.AzurePolicy {
	appID := d.Get("app_id").(string)
	appPassword := d.Get("app_password").(string)
	var azureResourceAndServices *models.AzureResourceAndServices // AzureResourceAndServices
	azureResourceAndServicesInterface, azureResourceAndServicesIsSet := d.GetOk("azure_resource_and_services")
	if azureResourceAndServicesIsSet {
		azureResourceAndServicesMap := azureResourceAndServicesInterface.([]interface{})[0].(map[string]interface{})
		azureResourceAndServices = AzureResourceAndServicesModelFromMap(azureResourceAndServicesMap)
	}
	var certificate *models.Certificate // Certificate
	certificateInterface, certificateIsSet := d.GetOk("certificate")
	if certificateIsSet {
		certificateMap := certificateInterface.([]interface{})[0].(map[string]interface{})
		certificate = CertificateModelFromMap(certificateMap)
	}
	cryptoKey := d.Get("crypto_key").(string)
	customDeploymentManaged := d.Get("custom_deployment_managed").(bool)
	encryptedSecrets := d.Get("encrypted_secrets").(map[string]string) // map[string]string
	tenantID := d.Get("tenant_id").(string)
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
	if azureResourceAndServicesIsSet {
		azureResourceAndServicesMap := azureResourceAndServicesInterface.([]interface{})[0].(map[string]interface{})
		azureResourceAndServices = AzureResourceAndServicesModelFromMap(azureResourceAndServicesMap)
	}
	//
	var certificate *models.Certificate // Certificate
	certificateInterface, certificateIsSet := m["certificate"]
	if certificateIsSet {
		certificateMap := certificateInterface.([]interface{})[0].(map[string]interface{})
		certificate = CertificateModelFromMap(certificateMap)
	}
	//
	cryptoKey := m["crypto_key"].(string)
	customDeploymentManaged := m["custom_deployment_managed"].(bool)
	encryptedSecrets := m["encrypted_secrets"].(map[string]string)
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

// Update the underlying AzurePolicy resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
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

// Iterate throught and update the AzurePolicy resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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

// Schema mapping representing the AzurePolicy resource defined in the Terraform configuration
func AzurePolicySchema() map[string]*schema.Schema {
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
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"certificate": {
			Description: `Certificate object holds the details of certificate like encryption type, validity, subject etc`,
			// We assume it's an enum type
			Type:     schema.TypeString,
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

// Retrieve property field names for updating the AzurePolicy resource
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
