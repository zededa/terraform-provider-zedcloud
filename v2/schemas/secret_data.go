package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SecretDataModel(d *schema.ResourceData) *models.SecretData {
	encodedPassword, _ := d.Get("encoded_password").(string)
	encodedUsername, _ := d.Get("encoded_username").(string)
	sSHPrivateKey, _ := d.Get("ssh_private_key").(string)
	sSHPublicKey, _ := d.Get("ssh_public_key").(string)
	return &models.SecretData{
		EncodedPassword: encodedPassword,
		EncodedUsername: encodedUsername,
		SSHPrivateKey:   sSHPrivateKey,
		SSHPublicKey:    sSHPublicKey,
	}
}

func SecretDataModelFromMap(m map[string]interface{}) *models.SecretData {
	encodedPassword := m["encoded_password"].(string)
	encodedUsername := m["encoded_username"].(string)
	sSHPrivateKey := m["ssh_private_key"].(string)
	sSHPublicKey := m["ssh_public_key"].(string)
	return &models.SecretData{
		EncodedPassword: encodedPassword,
		EncodedUsername: encodedUsername,
		SSHPrivateKey:   sSHPrivateKey,
		SSHPublicKey:    sSHPublicKey,
	}
}

func SetSecretDataResourceData(d *schema.ResourceData, m *models.SecretData) {
	d.Set("encoded_password", m.EncodedPassword)
	d.Set("encoded_username", m.EncodedUsername)
	d.Set("ssh_private_key", m.SSHPrivateKey)
	d.Set("ssh_public_key", m.SSHPublicKey)
}

func SetSecretDataSubResourceData(m []*models.SecretData) (d []*map[string]interface{}) {
	for _, SecretDataModel := range m {
		if SecretDataModel != nil {
			properties := make(map[string]interface{})
			properties["encoded_password"] = SecretDataModel.EncodedPassword
			properties["encoded_username"] = SecretDataModel.EncodedUsername
			properties["ssh_private_key"] = SecretDataModel.SSHPrivateKey
			properties["ssh_public_key"] = SecretDataModel.SSHPublicKey
			d = append(d, &properties)
		}
	}
	return
}

func SecretDataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"encoded_password": {
			Description: `Base64 encoded password for basic authentication, Required if Secret type is SECRET_TYPE_BASIC_AUTH`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"encoded_username": {
			Description: `Base64 encoded username for basic authentication, Required if Secret type is SECRET_TYPE_BASIC_AUTH`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"ssh_private_key": {
			Description: `Base64 encoded SSH private key for SSH authentication, Required if Secret type is SECRET_TYPE_SSH`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"ssh_public_key": {
			Description: `Base64 encoded SSH public key for SSH authentication, Required if Secret type is SECRET_TYPE_SSH`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetSecretDataPropertyFields() (t []string) {
	return []string{
		"encoded_password",
		"encoded_username",
		"ssh_private_key",
		"ssh_public_key",
	}
}
