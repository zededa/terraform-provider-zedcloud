package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SSHCredentialsModel(d *schema.ResourceData) *models.SSHCredentials {
	sSHPrivateKey, _ := d.Get("ssh_private_key").(string)
	sSHPublicKey, _ := d.Get("ssh_public_key").(string)
	return &models.SSHCredentials{
		SSHPrivateKey: sSHPrivateKey,
		SSHPublicKey:  sSHPublicKey,
	}
}

func SSHCredentialsModelFromMap(m map[string]interface{}) *models.SSHCredentials {
	sSHPrivateKey := m["ssh_private_key"].(string)
	sSHPublicKey := m["ssh_public_key"].(string)
	return &models.SSHCredentials{
		SSHPrivateKey: sSHPrivateKey,
		SSHPublicKey:  sSHPublicKey,
	}
}

func SetSSHCredentialsResourceData(d *schema.ResourceData, m *models.SSHCredentials) {
	d.Set("ssh_private_key", m.SSHPrivateKey)
	d.Set("ssh_public_key", m.SSHPublicKey)
}

func SetSSHCredentialsSubResourceData(m []*models.SSHCredentials) (d []*map[string]interface{}) {
	for _, SSHCredentialsModel := range m {
		if SSHCredentialsModel != nil {
			properties := make(map[string]interface{})
			properties["ssh_private_key"] = SSHCredentialsModel.SSHPrivateKey
			properties["ssh_public_key"] = SSHCredentialsModel.SSHPublicKey
			d = append(d, &properties)
		}
	}
	return
}

func SSHCredentialsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ssh_private_key": {
			Description: `SSH private key for authentication`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"ssh_public_key": {
			Description: `SSH public key for authentication`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetSSHCredentialsPropertyFields() (t []string) {
	return []string{
		"ssh_private_key",
		"ssh_public_key",
	}
}
