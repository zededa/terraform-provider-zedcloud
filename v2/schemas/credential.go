package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func CredentialModel(d *schema.ResourceData) *models.Credential {
	currentCred, _ := d.Get("current_cred").(string)
	forgot, _ := d.Get("forgot").(bool)
	id, _ := d.Get("id").(string)
	newCred, _ := d.Get("new_cred").(string)
	owner, _ := d.Get("owner").(string)
	saltInt, _ := d.Get("salt").(int)
	salt := int64(saltInt)
	var typeVar *models.CredentialType // CredentialType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewCredentialType(models.CredentialType(typeModel))
	}
	return &models.Credential{
		CurrentCred: currentCred,
		Forgot:      forgot,
		ID:          id,
		NewCred:     newCred,
		Owner:       owner,
		Salt:        salt,
		Type:        typeVar,
	}
}

func CredentialModelFromMap(m map[string]interface{}) *models.Credential {
	currentCred := m["current_cred"].(string)
	forgot := m["forgot"].(bool)
	id := m["id"].(string)
	newCred := m["new_cred"].(string)
	owner := m["owner"].(string)
	salt := int64(m["salt"].(int))     // int64
	var typeVar *models.CredentialType // CredentialType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewCredentialType(models.CredentialType(typeModel))
	}
	return &models.Credential{
		CurrentCred: currentCred,
		Forgot:      forgot,
		ID:          id,
		NewCred:     newCred,
		Owner:       owner,
		Salt:        salt,
		Type:        typeVar,
	}
}

func SetCredentialResourceData(d *schema.ResourceData, m *models.Credential) {
	d.Set("current_cred", m.CurrentCred)
	d.Set("forgot", m.Forgot)
	d.Set("id", m.ID)
	d.Set("new_cred", m.NewCred)
	d.Set("owner", m.Owner)
	d.Set("salt", m.Salt)
	d.Set("type", m.Type)
}

func SetCredentialSubResourceData(m []*models.Credential) (d []*map[string]interface{}) {
	for _, CredentialModel := range m {
		if CredentialModel != nil {
			properties := make(map[string]interface{})
			properties["current_cred"] = CredentialModel.CurrentCred
			properties["forgot"] = CredentialModel.Forgot
			properties["id"] = CredentialModel.ID
			properties["new_cred"] = CredentialModel.NewCred
			properties["owner"] = CredentialModel.Owner
			properties["salt"] = CredentialModel.Salt
			properties["type"] = CredentialModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func CredentialSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"current_cred": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
		},

		"forgot": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"id": {
			Description: ``,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"new_cred": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
		},

		"owner": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"salt": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetCredentialPropertyFields() (t []string) {
	return []string{
		"current_cred",
		"forgot",
		"id",
		"new_cred",
		"owner",
		"salt",
		"type",
	}
}
